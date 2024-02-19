// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

// Package bioxrv provides gleancli extensions for bioRxiv and medRxiv.
package biorxiv

import (
	"context"
	"os"

	"cloudeng.io/cmdutil/subcmd"
	"cloudeng.io/file/checkpoint"
	gleancfg "cloudeng.io/glean/config"
	"cloudeng.io/glean/crawlindex/config"
	"cloudeng.io/webapi/biorxiv/biorxivcmd"
	"cloudeng.io/webapi/operations"
)

type CommonFlags struct {
	config.FileFlags
}

type CrawlFlags struct {
	CommonFlags
	biorxivcmd.CrawlFlags
}

type ScanFlags struct {
	CommonFlags
	biorxivcmd.ScanFlags
}

type LookupFlags struct {
	CommonFlags
	biorxivcmd.LookupFlags
}

type IndexFlags struct {
	CommonFlags
	biorxivcmd.IndexFlags
}

const (
	cmdName = "biorxiv"
	cmdSpec = `
- name: biorxiv
  summary: biorxiv commands
  commands:
    - name: crawl
      summary: crawl biorxiv/medrxiv, i.e. download article metadata and abstracts
      arguments:
        - datasource-name - the datasource to be crawled
    - name: scan-downloaded
      summary: scan downloaded articles
      arguments:
        - datasource-name - the datasource whose downloaded protocols are to be scanned
    - name: lookup-downloaded
      summary: lookup downloaded articles
      arguments:
        - datasource-name - the datasource whose downloaded articles are to be accessed
        - DOI - the DOI of the article to be looked up
        - ...
`
)

var ExtensionSpec = gleancfg.ExtensionSpec{
	Name:       cmdName,
	CmdSpec:    cmdSpec,
	AuthCfg:    biorxivcmd.Auth{},
	ServiceCfg: biorxivcmd.Service{},
	AddFunc:    AddExtension,
}

func AddExtension(extension gleancfg.Extension, cmdSet *subcmd.CommandSetYAML, parents []string) error {
	c := &command{parent: extension}
	cmdSet.Set(append(parents, cmdName, "crawl")...).MustRunner(c.crawlCmd, &CrawlFlags{})
	cmdSet.Set(append(parents, cmdName, "scan-downloaded")...).MustRunner(c.scanDownloadsCmd, &ScanFlags{})
	cmdSet.Set(append(parents, cmdName, "lookup-downloaded")...).MustRunner(c.lookupDownloadsCmd, &LookupFlags{})
	return nil
}

type command struct {
	parent gleancfg.Extension
	fs     operations.FS
	chkpt  checkpoint.Operation
}

func (cmd *command) new(ctx context.Context, fv CommonFlags, _ biorxivcmd.CommonFlags, datasource string) (*biorxivcmd.Command, error) {
	cfg, err := config.DatasourceForName(ctx, fv.ConfigFile, datasource)
	if err != nil {
		return nil, err
	}
	cacheRoot := os.ExpandEnv(cfg.Cache.Path)
	cmd.fs, err = cmd.parent.Options().CreateStoreFS(ctx, cacheRoot, cfg.Cache.ServiceConfig)
	if err != nil {
		return nil, err
	}
	cmd.chkpt, err = cmd.parent.Options().CreateCheckpointOp(ctx, cacheRoot, cfg.Cache.ServiceConfig)
	if err != nil {
		return nil, err
	}
	return biorxivcmd.NewCommand(ctx, cfg.APICrawls, cmd.fs, cacheRoot, cmd.chkpt, cmdName)
}

func (cmd *command) crawlCmd(ctx context.Context, values interface{}, args []string) error {
	fv := values.(*CrawlFlags)
	c, err := cmd.new(ctx, fv.CommonFlags, fv.CrawlFlags.CommonFlags, args[0])
	if err != nil {
		return err
	}
	return c.Crawl(ctx, fv.CrawlFlags)
}

func (cmd *command) scanDownloadsCmd(ctx context.Context, values interface{}, args []string) error {
	fv := values.(*ScanFlags)
	c, err := cmd.new(ctx, fv.CommonFlags, fv.ScanFlags.CommonFlags, args[0])
	if err != nil {
		return err
	}
	return c.ScanDownloaded(ctx, &fv.ScanFlags)
}

func (cmd *command) lookupDownloadsCmd(ctx context.Context, values interface{}, args []string) error {
	fv := values.(*LookupFlags)
	c, err := cmd.new(ctx, fv.CommonFlags, fv.LookupFlags.CommonFlags, args[0])
	if err != nil {
		return err
	}
	return c.LookupDownloaded(ctx, &fv.LookupFlags, args[1:]...)
}
