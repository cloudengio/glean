// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package cmds

import (
	"context"
	"errors"
	"fmt"
	"os"

	"cloudeng.io/cmdutil"
	"cloudeng.io/cmdutil/cmdyaml"
	"cloudeng.io/cmdutil/subcmd"
	"cloudeng.io/file/checkpoint"
	"cloudeng.io/file/crawl/crawlcmd"
	gleancfg "cloudeng.io/glean/config"
	"cloudeng.io/glean/crawlindex/crawl"
	"cloudeng.io/glean/crawlindex/datasources"
	"cloudeng.io/glean/crawlindex/index"
	"cloudeng.io/glean/gleancli/builtin/static"
	"cloudeng.io/webapi/operations"
	"gopkg.in/yaml.v3"
)

type GlobalFlags struct {
	gleancfg.GleanFlags
}

var (
	globalFlags  GlobalFlags
	globalConfig gleancfg.Glean
)

const baseCommands = `name: gleancli
summary: command line for working with the Glean API
commands:
  - name: datasources
    summary: register or list Glean SDK data sources
    commands:
      - name: download
        summary: download and display the configuration for the specified data source from its glean instance
        arguments:
          - glean-instance-name - the name of the glean instance to use
          - datasource-name     - the datasource to be downloaded
      - name: register
        summary: add the named datasource to a glean instance using the configuration defined in the datasource configuration file specified using the --datasource-configs flag.
        arguments:
          - datasource-name - the datasource to be registered
      - name: show-config
        summary: load and display the specified config file
        arguments:
          - datasource-config-file
      - name: explain-config
        summary: explain the configuration options for configured connectors

  - name: crawl
    summary: crawl a web site or filesystem
    commands:
      - name: run
        summary: run the crawls for a datasource
        arguments:
          - datasource-name - the datasource to be crawled

  - name: api
    summary: API releated commands
    commands:
      {{range subcmdExtension "APIExtensions"}}
      {{.}}{{end}}

  - name: index
    summary: crawl a web site or filesystem
    commands:
      - name: bulk
        summary: index the results of a prior crawl as a single bulk update
        arguments:
          - datasource-name - the datasource to be indexed

      - name: stats
        summary: display statistics for the specified datasource
        arguments:
          - datasource-name - the datasource to be indexed

      - name: query
        summary: run the specified query against the specified datasource
        arguments:
          - datasource-name - the name of the datasource to use
          - query           - the query to run

      - name: process-now
        summary: request processing of all documents in the specified datasource as soon as possible.
        arguments:
          - datasource-name - the name of the datasource to process

      - name: delete-all
        summary: delete all of the documents from a datasource
        arguments:
          - datasource-name - the name of the datasource to use

      - name: delete
        summary: delete indexed documents returned by a query based on their type
        arguments:
          - datasource-name - the name of the datasource to use
          - query           - the query to be used to find documents to be deleted

  {{range subcmdExtension "Extensions"}}
  {{.}}{{end}}
`

var cmdExtensions []gleancfg.Extension

func asSubcmdExtensions(exts []gleancfg.Extension) []subcmd.Extension {
	subext := []subcmd.Extension{}
	for _, e := range exts {
		subext = append(subext, e)
	}
	return subext
}

type Options struct {
	CrawlProcessors    static.CrawlProcessors
	IndexProcessors    static.IndexProcessors
	CreateCrawlFS      func(name string, cfg yaml.Node, factories map[string]crawlcmd.FSFactory) error
	CreateStoreFS      func(ctx context.Context, path string, cfg yaml.Node) (operations.FS, error)
	CreateCheckpointOp func(ctx context.Context, path string, cfg yaml.Node) (checkpoint.Operation, error)
	Extensions         []gleancfg.Extension
	APIExtensions      []gleancfg.Extension
	MainWrapper        func(ctx context.Context, extOpts gleancfg.ExtensionOptions, cmdRunner func(ctx context.Context) error) error
}

func MustNew(options Options) *subcmd.CommandSetYAML {

	cmdExtensions = append([]gleancfg.Extension{}, options.Extensions...)
	cmdExtensions = append(cmdExtensions, options.APIExtensions...)

	cmdSet := subcmd.MustFromYAMLTemplate(baseCommands,
		subcmd.MergeExtensions("Extensions", asSubcmdExtensions(options.Extensions)...),
		subcmd.MergeExtensions("APIExtensions", asSubcmdExtensions(options.APIExtensions)...),
	)

	ds := Datasources{
		extensions: cmdExtensions,
	}
	cmdSet.Set("datasources", "download").MustRunner(ds.Download, &struct{}{})

	cmdSet.Set("datasources", "register").MustRunner(ds.Register, &datasources.Flags{})

	cmdSet.Set("datasources", "show-config").MustRunner(ds.ShowConfig, &struct{}{})

	cmdSet.Set("datasources", "explain-config").MustRunner(ds.ExplainConfig, &struct{}{})

	cc := Crawl{Options: options}
	cmdSet.Set("crawl", "run").MustRunner(cc.Run, &crawl.Flags{})

	idx := Index{Options: options}
	cmdSet.Set("index", "bulk").MustRunner(idx.bulk, &index.BulkFlags{})

	cmdSet.Set("index", "stats").MustRunner(idx.stats, &index.StatsFlags{})

	cmdSet.Set("index", "query").MustRunner(idx.query, &index.QueryFlags{})

	cmdSet.Set("index", "delete").MustRunner(idx.delete, &index.DeleteFlags{})

	cmdSet.Set("index", "delete-all").MustRunner(idx.deleteAll, &index.DeleteAllFlags{})

	cmdSet.Set("index", "process-now").MustRunner(idx.processNow, &index.ProcessNowFlags{})

	cmdSet.MustAddExtensions()

	globals := subcmd.GlobalFlagSet()
	globals.MustRegisterFlagStruct(&globalFlags, nil, nil)
	cmdSet.WithGlobalFlags(globals)

	extOpts := gleancfg.ExtensionOptions{
		CreateStoreFS:      options.CreateStoreFS,
		CreateCheckpointOp: options.CreateCheckpointOp,
	}
	wrapper := options.MainWrapper
	if wrapper == nil {
		wrapper = mainWrapper
	}
	cmdSet.WithMain(func(ctx context.Context, cmdRunner func(ctx context.Context) error) error {
		return wrapper(ctx, extOpts, cmdRunner)
	})
	return cmdSet
}

func mainWrapper(ctx context.Context, extOpts gleancfg.ExtensionOptions, cmdRunner func(ctx context.Context) error) error {
	ctx, cancel := context.WithCancel(ctx)
	cmdutil.HandleSignals(cancel, os.Interrupt, os.Kill)

	err := cmdyaml.ParseConfigFile(ctx, globalFlags.Config, &globalConfig)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return err
		}
		fmt.Printf("warning: %q not found\n", globalFlags.Config)
	}
	extOpts.Glean = globalConfig
	for _, ext := range cmdExtensions {
		ext.SetOptions(extOpts)
	}
	return cmdRunner(ctx)
}
