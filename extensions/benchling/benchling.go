// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package benchling

import (
	"context"
	"os"
	"strings"

	"cloudeng.io/cmdutil/subcmd"
	gleancfg "cloudeng.io/glean/config"
	"cloudeng.io/glean/crawlindex/config"
	"cloudeng.io/webapi/benchling/benchlingcmd"
)

type CommonFlags struct {
	config.FileFlags
}

type CrawlFlags struct {
	CommonFlags
	benchlingcmd.CrawlFlags
	Entities string `subcmd:"entities,'users,entries,folders,projects','comma separated list of entities to crawl: users, entries, folders, projects'"`
}

type IndexFlags struct {
	CommonFlags
	benchlingcmd.IndexFlags
}

const (
	cmdName = "benchling"
	cmdSpec = `
- name: benchling
  summary: benchling commands
  commands:
    - name: crawl
      summary: crawl benchling, i.e. download all users, folders etc.
      arguments:
        - datasource-name - the datasource to be crawled
    - name: create-indexable-documents
      summary: create-indexable-documennts from all downloaded benchling data
      arguments:
        - datasource-name - the datasource to be crawled
`
)

func Extension(parents ...string) gleancfg.Extension {
	c := &command{}
	return gleancfg.NewExtension(cmdName, cmdSpec,
		benchlingcmd.Auth{}, benchlingcmd.Service{},
		func(cmdSet *subcmd.CommandSetYAML) error {
			cmdSet.Set(append(parents, cmdName, "crawl")...).MustRunner(c.crawlCmd, &CrawlFlags{})
			cmdSet.Set(append(parents, cmdName, "create-indexable-documents")...).MustRunner(c.indexCmd, &IndexFlags{})
			return nil
		})
}

type command struct{ cacheRoot string }

func (cmd *command) new(ctx context.Context, fv CommonFlags, pfv benchlingcmd.CommonFlags, datasource string) (*benchlingcmd.Command, error) {
	cfg, err := config.DatasourceForName(ctx, fv.ConfigFile, datasource)
	if err != nil {
		return nil, err
	}
	cmd.cacheRoot = os.ExpandEnv(cfg.Cache.Path)
	return benchlingcmd.NewCommand(ctx, cfg.APICrawls, cmdName, pfv.BenchlingConfig)
}

func (cmd *command) crawlCmd(ctx context.Context, values interface{}, args []string) error {
	fv := values.(*CrawlFlags)
	c, err := cmd.new(ctx, fv.CommonFlags, fv.CrawlFlags.CommonFlags, args[0])
	if err != nil {
		return err
	}
	entities := strings.Split(fv.Entities, ",")
	return c.Crawl(ctx, cmd.cacheRoot, fv.CrawlFlags, entities...)
}

func (cmd *command) indexCmd(ctx context.Context, values interface{}, args []string) error {
	fv := values.(*IndexFlags)
	c, err := cmd.new(ctx, fv.CommonFlags, fv.IndexFlags.CommonFlags, args[0])
	if err != nil {
		return err
	}
	return c.CreateIndexableDocuments(ctx, cmd.cacheRoot, fv.IndexFlags)
}
