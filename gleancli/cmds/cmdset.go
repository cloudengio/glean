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
	"cloudeng.io/cmdutil/subcmd"
	"cloudeng.io/glean/crawlindex/config"
	"cloudeng.io/glean/crawlindex/crawl"
	"cloudeng.io/glean/crawlindex/datasources"
	"cloudeng.io/glean/crawlindex/index"
)

type GlobalFlags struct {
	config.GleanConfigFlags
}

var (
	globalFlags  GlobalFlags
	GlobalConfig config.GleanConfig
	cmdSet       *subcmd.CommandSetYAML
)

func MustNew() *subcmd.CommandSetYAML {
	cmdSet := subcmd.MustFromYAML(`name: gleancli
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
        summary: add the named datasource to a glean instance using the configuration defined in the datasource configuration file specified using the --datasource-config flag.
        arguments:
          - datasource-name - the datasource to be downloaded
      - name: show-config
        summary: load and display the specified config file
        arguments:
          - datasource-config-file

  - name: crawl
    summary: crawl a web site or filesystem
    commands:
      - name: run
        summary: run the crawls for a datasource
        arguments:
          - datasource-name - the datasource to be crawled

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
`)

	var ds Datasources
	cmdSet.Set("datasources", "download").MustRunnerAndFlags(
		ds.Download, subcmd.MustRegisteredFlagSet(&struct{}{}))

	cmdSet.Set("datasources", "register").MustRunnerAndFlags(
		ds.Register, subcmd.MustRegisteredFlagSet(&datasources.Flags{}))

	cmdSet.Set("datasources", "show-config").MustRunnerAndFlags(
		ds.ShowConfig, subcmd.MustRegisteredFlagSet(&struct{}{}))

	var cc Crawl
	cmdSet.Set("crawl", "run").MustRunnerAndFlags(
		cc.Run, subcmd.MustRegisteredFlagSet(&crawl.Flags{}))

	var idx Index
	cmdSet.Set("index", "bulk").MustRunnerAndFlags(
		idx.bulk, subcmd.MustRegisteredFlagSet(&index.BulkFlags{}))

	cmdSet.Set("index", "stats").MustRunnerAndFlags(
		idx.stats, subcmd.MustRegisteredFlagSet(&index.StatsFlags{}))

	globals := subcmd.GlobalFlagSet()
	globals.MustRegisterFlagStruct(&globalFlags, nil, nil)
	cmdSet.WithGlobalFlags(globals)
	cmdSet.WithMain(mainWrapper)
	return cmdSet
}

func mainWrapper(ctx context.Context, cmdRunner func(ctx context.Context) error) error {
	ctx, cancel := context.WithCancel(ctx)
	cmdutil.HandleSignals(cancel, os.Interrupt, os.Kill)
	err := cmdutil.ParseYAMLConfigFile(globalFlags.Config, &GlobalConfig)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return err
		}
		fmt.Printf("warning: %q not found\n", globalFlags.Config)
	}
	return cmdRunner(ctx)
}
