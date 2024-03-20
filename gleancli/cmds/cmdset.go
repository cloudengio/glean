// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package cmds

import (
	"context"
	"os"

	"cloudeng.io/cmdutil"
	"cloudeng.io/cmdutil/subcmd"
	"cloudeng.io/glean/crawlindex/crawl"
	"cloudeng.io/glean/gleancli/extensions"
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
          - glean-domain-name   - the glean instance to use
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

var cmdExtensions []extensions.Extension

func asSubcmdExtensions(exts []extensions.Extension) []subcmd.Extension {
	subext := []subcmd.Extension{}
	for _, e := range exts {
		subext = append(subext, e)
	}
	return subext
}

type Options struct {
	extensions.StaticResources

	extensions.DynamicResources

	Extensions    []extensions.Extension
	APIExtensions []extensions.Extension

	InitContext func(ctx context.Context) (context.Context, error)
}

func MustNew(options Options) *subcmd.CommandSetYAML {

	cmdExtensions = append([]extensions.Extension{}, options.Extensions...)
	cmdExtensions = append(cmdExtensions, options.APIExtensions...)

	cmdSet := subcmd.MustFromYAMLTemplate(baseCommands,
		subcmd.MergeExtensions("Extensions", asSubcmdExtensions(options.Extensions)...),
		subcmd.MergeExtensions("APIExtensions", asSubcmdExtensions(options.APIExtensions)...),
	)

	ds := Datasources{
		Options:    options,
		extensions: cmdExtensions,
	}
	cmdSet.Set("datasources", "download").MustRunner(ds.Download, &DownloadFlags{})

	cmdSet.Set("datasources", "register").MustRunner(ds.Register, &RegisterFlags{})

	cmdSet.Set("datasources", "show-config").MustRunner(ds.ShowConfig, &struct{}{})

	cmdSet.Set("datasources", "explain-config").MustRunner(ds.ExplainConfig, &struct{}{})

	cc := Crawl{Options: options}
	cmdSet.Set("crawl", "run").MustRunner(cc.Run, &crawl.Flags{})

	idx := Index{Options: options}
	cmdSet.Set("index", "bulk").MustRunner(idx.bulk, &BulkFlags{})

	cmdSet.Set("index", "stats").MustRunner(idx.stats, &StatsFlags{})

	cmdSet.Set("index", "query").MustRunner(idx.query, &QueryFlags{})

	cmdSet.Set("index", "delete").MustRunner(idx.delete, &DeleteFlags{})

	cmdSet.Set("index", "delete-all").MustRunner(idx.deleteAll, &DeleteAllFlags{})

	cmdSet.Set("index", "process-now").MustRunner(idx.processNow, &ProcessNowFlags{})

	cmdSet.MustAddExtensions()

	initState := options.InitContext
	if initState == nil {
		initState = func(ctx context.Context) (context.Context, error) {
			return ctx, nil
		}
	}
	cmdSet.WithMain(func(ctx context.Context, cmdRunner func(ctx context.Context) error) error {
		ctx, cancel := context.WithCancel(ctx)
		cmdutil.HandleSignals(cancel, os.Interrupt, os.Kill)
		ctx, err := initState(ctx)
		if err != nil {
			return err
		}
		extOpts := extensions.ExtensionOptions{
			StaticResources:  options.StaticResources,
			DynamicResources: options.DynamicResources,
		}
		for _, ext := range cmdExtensions {
			ext.SetOptions(extOpts)
		}
		return cmdRunner(ctx)
	})
	return cmdSet
}
