// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

// Package bioxrv provides gleancli extensions for bioRxiv and medRxiv.
package biorxiv

import (
	"context"
	"fmt"

	"cloudeng.io/cmdutil/subcmd"
	"cloudeng.io/glean/crawlindex/config"
	"cloudeng.io/glean/gleancli/extensions"
	"cloudeng.io/webapi/clients/biorxiv/biorxivcmd"
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

var ExtensionSpec = extensions.ExtensionSpec{
	Name:       cmdName,
	CmdSpec:    cmdSpec,
	AuthCfg:    struct{}{},
	ServiceCfg: biorxivcmd.Service{},
	AddFunc:    AddExtension,
}

func AddExtension(extension extensions.Extension, cmdSet *subcmd.CommandSetYAML, parents []string) error {
	c := &command{parent: extension}
	cmdSet.Set(append(parents, cmdName, "crawl")...).MustRunner(c.crawlCmd, &CrawlFlags{})
	cmdSet.Set(append(parents, cmdName, "scan-downloaded")...).MustRunner(c.scanDownloadsCmd, &ScanFlags{})
	cmdSet.Set(append(parents, cmdName, "lookup-downloaded")...).MustRunner(c.lookupDownloadsCmd, &LookupFlags{})
	return nil
}

type command struct {
	parent extensions.Extension
}

func (cmd *command) new(ctx context.Context, fv CommonFlags, datasource string) (*biorxivcmd.Command, error) {
	cfg, resources, err := cmd.parent.Options().ResourcesForDatasource(ctx, fv.ConfigFile, "", datasource)
	if err != nil {
		return nil, err
	}
	first, ok := extensions.FirstAPICrawl(cfg.APICrawls)
	if !ok {
		return nil, fmt.Errorf("no api crawl defined for %v", datasource)
	}
	return biorxivcmd.NewCommand(ctx, first, resources)
}

func (cmd *command) crawlCmd(ctx context.Context, values interface{}, args []string) error {
	fv := values.(*CrawlFlags)
	c, err := cmd.new(ctx, fv.CommonFlags, args[0])
	if err != nil {
		return err
	}
	return c.Crawl(ctx, fv.CrawlFlags)
}

func (cmd *command) scanDownloadsCmd(ctx context.Context, values interface{}, args []string) error {
	fv := values.(*ScanFlags)
	c, err := cmd.new(ctx, fv.CommonFlags, args[0])
	if err != nil {
		return err
	}
	return c.ScanDownloaded(ctx, &fv.ScanFlags)
}

func (cmd *command) lookupDownloadsCmd(ctx context.Context, values interface{}, args []string) error {
	fv := values.(*LookupFlags)
	c, err := cmd.new(ctx, fv.CommonFlags, args[0])
	if err != nil {
		return err
	}
	return c.LookupDownloaded(ctx, &fv.LookupFlags, args[1:]...)
}
