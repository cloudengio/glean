// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package benchling

import (
	"context"
	"fmt"
	"strings"

	"cloudeng.io/cmdutil/subcmd"
	"cloudeng.io/glean/crawlindex/config"
	"cloudeng.io/glean/gleancli/extensions"
	"cloudeng.io/webapi/clients/benchling/benchlingcmd"
)

type CommonFlags struct {
	config.FileFlags
	AuthFile string `subcmd:"benchling-auth,$HOME/.benchling.yaml,'benchling.io auth config file'"`
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

var ExtensionSpec = extensions.ExtensionSpec{
	Name:       cmdName,
	CmdSpec:    cmdSpec,
	AuthCfg:    extensions.APIKey{},
	ServiceCfg: benchlingcmd.Service{},
	AddFunc:    AddExtension,
}

func AddExtension(extension extensions.Extension, cmdSet *subcmd.CommandSetYAML, parents []string) error {
	c := &command{parent: extension}
	cmdSet.Set(append(parents, cmdName, "crawl")...).MustRunner(c.crawlCmd, &CrawlFlags{})
	cmdSet.Set(append(parents, cmdName, "create-indexable-documents")...).MustRunner(c.indexCmd, &IndexFlags{})
	return nil
}

type command struct {
	parent extensions.Extension
}

func (cmd *command) newCommand(ctx context.Context, fv CommonFlags, datasource string) (*benchlingcmd.Command, error) {
	cfg, resources, err := cmd.parent.Options().ResourcesForDatasource(ctx, fv.ConfigFile, fv.AuthFile, datasource)
	if err != nil {
		return nil, err
	}
	first, ok := extensions.FirstAPICrawl(cfg.APICrawls)
	if !ok {
		return nil, fmt.Errorf("no api crawl defined for %v", datasource)
	}
	return benchlingcmd.NewCommand(ctx, first, resources)
}

func (cmd *command) crawlCmd(ctx context.Context, values interface{}, args []string) error {
	fv := values.(*CrawlFlags)
	c, err := cmd.newCommand(ctx, fv.CommonFlags, args[0])
	if err != nil {
		return err
	}
	entities := strings.Split(fv.Entities, ",")
	return c.Crawl(ctx, fv.CrawlFlags, entities...)
}

func (cmd *command) indexCmd(ctx context.Context, values interface{}, args []string) error {
	fv := values.(*IndexFlags)
	c, err := cmd.newCommand(ctx, fv.CommonFlags, args[0])
	if err != nil {
		return err
	}
	return c.CreateIndexableDocuments(ctx, fv.IndexFlags)
}
