// Copyright 2022 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package crawl

import (
	"context"

	"cloudeng.io/file/crawl/crawlcmd"
	gleancfg "cloudeng.io/glean/config"
	"cloudeng.io/glean/crawlindex/config"
	"cloudeng.io/sync/errgroup"
)

// Flags represents the flags that are used to control the crawl.
type Flags struct {
	config.FileFlags
	Outlinks bool `subcmd:"outlinks,false,display extracted outlinks"`
	Progress bool `subcmd:"progress,true,'display progress of downloads'"`
}

// Crawler represents a crawler instance that contains global configuration
// information.
type Crawler struct {
	gleancfg.CrawlProcessors
	gleancfg.DynamicResources
}

func (c *Crawler) Run(ctx context.Context, fv *Flags, datasource string) error {
	cfg, err := config.DatasourceForName(ctx, fv.ConfigFile, datasource)
	if err != nil {
		return err
	}

	fsmap := map[string]crawlcmd.FSFactory{}
	for _, crawl := range cfg.Crawls {
		if err := c.PopulateCrawlFS(ctx, crawl.Service, fsmap); err != nil {
			return err
		}
	}
	resources := crawlcmd.Resources{
		Extractors:          c.CrawlProcessors.Extractors,
		CrawlStoreFactories: fsmap,
		NewContentFS:        c.NewContentFS,
	}

	// Run all of the crawlers concurrently.
	g := errgroup.T{}
	for _, crawl := range cfg.Crawls {
		crawler := crawlcmd.NewCrawler(crawl.Config, resources)
		g.Go(func() error {
			return crawler.Run(ctx, fv.Outlinks, fv.Progress)
		})
	}
	return g.Wait()
}
