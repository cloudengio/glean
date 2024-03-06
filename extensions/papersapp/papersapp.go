// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package papersapp

import (
	"context"

	"cloudeng.io/cmdutil/subcmd"
	gleancfg "cloudeng.io/glean/config"
	"cloudeng.io/glean/crawlindex/config"
	"cloudeng.io/webapi/apis/papersapp/papersappcmd"
	"cloudeng.io/webapi/operations/apicrawlcmd"
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

var ExtensionSpec = gleancfg.ExtensionSpec{
	Name:       cmdName,
	CmdSpec:    cmdSpec,
	AuthCfg:    gleancfg.APIKey{},
	ServiceCfg: papersappcmd.Service{},
	AddFunc:    AddExtension,
}

func AddExtension(extension gleancfg.Extension, cmdSet *subcmd.CommandSetYAML, parents []string) error {
	c := &command{parent: extension}
	cmdSet.Set(append(parents, cmdName, "crawl")...).MustRunner(c.crawlCmd, &CrawlFlags{})
	cmdSet.Set(append(parents, cmdName, "scan-downloaded")...).MustRunner(c.scanCmd, &ScanFlags{})
	return nil
}

type command struct {
	parent gleancfg.Extension
}

func (cmd *command) new(ctx context.Context, fv CommonFlags, datasource string) (*papersappcmd.Command, error) {
	cfg, err := config.DatasourceForName(ctx, fv.ConfigFile, datasource)
	if err != nil {
		return nil, err
	}
	var key gleancfg.APIKey
	token, err := key.ParseAndRead(ctx, fv.AuthFile, cmd.parent.Options().TokenReaders)
	if err != nil {
		return nil, err
	}
	resources := apicrawlcmd.Resources{
		Token:           token,
		NewOperationsFS: cmd.parent.Options().NewOperationsFS,
	}
	return papersappcmd.NewCommand(ctx, cfg.APICrawls[datasource], resources)
}

func (cmd *command) crawlCmd(ctx context.Context, values interface{}, args []string) error {
	fv := values.(*CrawlFlags)
	c, err := cmd.new(ctx, fv.CommonFlags, args[0])
	if err != nil {
		return err
	}
	return c.Crawl(ctx, &fv.CrawlFlags)
}

func (cmd *command) scanCmd(ctx context.Context, values interface{}, args []string) error {
	fv := values.(*ScanFlags)
	c, err := cmd.new(ctx, fv.CommonFlags, args[0])
	if err != nil {
		return err
	}
	return c.ScanDownloaded(ctx, &fv.ScanFlags)
}
