// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package papersapp

import (
	"context"
	"fmt"

	"cloudeng.io/cmdutil/subcmd"
	"cloudeng.io/glean/crawlindex/config"
	"cloudeng.io/glean/gleancli/extensions"
	"cloudeng.io/webapi/apis/papersapp/papersappcmd"
)

type CommonFlags struct {
	config.FileFlags
	AuthFile string `subcmd:"papersapp-auth,$HOME/.benchling.yaml,'papersapp.io auth config file'"`
}

type CrawlFlags struct {
	CommonFlags
	papersappcmd.CrawlFlags
}

type ScanFlags struct {
	CommonFlags
	papersappcmd.ScanFlags
}

const (
	cmdName = "papersapp"
	cmdSpec = `
- name: papersapp
  summary: papersapp commands
  commands:
    - name: crawl
      summary: crawl papersapp, i.e. download all items etc.
      arguments:
        - datasource-name - the datasource to be crawled
    - name: scan-downloaded
      summary: scan downloaded papersapp collections and items,
      arguments:
        - datasource-name - the datasource whose downloaded collections and items are to be scanned
`
)

var ExtensionSpec = extensions.ExtensionSpec{
	Name:       cmdName,
	CmdSpec:    cmdSpec,
	AuthCfg:    extensions.APIKey{},
	ServiceCfg: papersappcmd.Service{},
	AddFunc:    AddExtension,
}

func AddExtension(extension extensions.Extension, cmdSet *subcmd.CommandSetYAML, parents []string) error {
	c := &command{parent: extension}
	cmdSet.Set(append(parents, cmdName, "crawl")...).MustRunner(c.crawlCmd, &CrawlFlags{})
	cmdSet.Set(append(parents, cmdName, "scan-downloaded")...).MustRunner(c.scanCmd, &ScanFlags{})
	return nil
}

type command struct {
	parent extensions.Extension
}

func (cmd *command) newCommand(ctx context.Context, fv CommonFlags, datasource string) (*papersappcmd.Command, error) {
	cfg, resources, err := cmd.parent.Options().ResourcesForDatasource(ctx, fv.ConfigFile, fv.AuthFile, datasource)
	if err != nil {
		return nil, err
	}
	first, ok := extensions.FirstAPICrawl(cfg.APICrawls)
	if !ok {
		return nil, fmt.Errorf("no api crawl defined for %v", datasource)
	}
	return papersappcmd.NewCommand(ctx, first, resources)
}

func (cmd *command) crawlCmd(ctx context.Context, values interface{}, args []string) error {
	fv := values.(*CrawlFlags)
	c, err := cmd.newCommand(ctx, fv.CommonFlags, args[0])
	if err != nil {
		return err
	}
	return c.Crawl(ctx, &fv.CrawlFlags)
}

func (cmd *command) scanCmd(ctx context.Context, values interface{}, args []string) error {
	fv := values.(*ScanFlags)
	c, err := cmd.newCommand(ctx, fv.CommonFlags, args[0])
	if err != nil {
		return err
	}
	return c.ScanDownloaded(ctx, &fv.ScanFlags)
}
