// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

// Package bioxrv provides gleancli extensions for bioRxiv and medRxiv.
package biorxiv

import (
	"context"
	"os"

	"cloudeng.io/cmdutil/subcmd"
	gleancfg "cloudeng.io/glean/config"
	"cloudeng.io/glean/crawlindex/config"
	"cloudeng.io/webapi/biorxiv/biorxivcmd"
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

func Extension(parents ...string) gleancfg.Extension {
	c := &command{}
	return gleancfg.NewExtension(cmdName, cmdSpec, func(cmdSet *subcmd.CommandSetYAML) error {
		cmdSet.Set(append(parents, cmdName, "crawl")...).MustRunnerAndFlags(c.crawlCmd, subcmd.MustRegisteredFlagSet(&CrawlFlags{}))
		cmdSet.Set(append(parents, cmdName, "scan-downloaded")...).MustRunnerAndFlags(c.scanDownloadsCmd, subcmd.MustRegisteredFlagSet(&ScanFlags{}))
		cmdSet.Set(append(parents, cmdName, "lookup-downloaded")...).MustRunnerAndFlags(c.lookupDownloadsCmd, subcmd.MustRegisteredFlagSet(&LookupFlags{}))
		return nil
	})
}

type command struct{ cacheRoot string }

func (cmd *command) new(ctx context.Context, fv CommonFlags, pfv biorxivcmd.CommonFlags, datasource string) (*biorxivcmd.Command, error) {
	cfg, err := config.DatasourceForName(ctx, fv.ConfigFile, datasource)
	if err != nil {
		return nil, err
	}
	cmd.cacheRoot = os.ExpandEnv(cfg.Cache.Path)
	return biorxivcmd.NewCommand(ctx, cfg.APICrawls, cmdName)
}

func (cmd *command) crawlCmd(ctx context.Context, values interface{}, args []string) error {
	fv := values.(*CrawlFlags)
	c, err := cmd.new(ctx, fv.CommonFlags, fv.CrawlFlags.CommonFlags, args[0])
	if err != nil {
		return err
	}
	return c.Crawl(ctx, cmd.cacheRoot, fv.CrawlFlags)
}

func (cmd *command) scanDownloadsCmd(ctx context.Context, values interface{}, args []string) error {
	fv := values.(*ScanFlags)
	c, err := cmd.new(ctx, fv.CommonFlags, fv.ScanFlags.CommonFlags, args[0])
	if err != nil {
		return err
	}
	return c.ScanDownloaded(ctx, cmd.cacheRoot, &fv.ScanFlags)
}

func (cmd *command) lookupDownloadsCmd(ctx context.Context, values interface{}, args []string) error {
	fv := values.(*LookupFlags)
	c, err := cmd.new(ctx, fv.CommonFlags, fv.LookupFlags.CommonFlags, args[0])
	if err != nil {
		return err
	}
	return c.LookupDownloaded(ctx, cmd.cacheRoot, &fv.LookupFlags, args[1:]...)
}
