// Copyright 2022 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package crawl

import (
	"context"

	"cloudeng.io/file/content"
	"cloudeng.io/file/crawl/crawlcmd"
	"cloudeng.io/file/crawl/outlinks"
	"cloudeng.io/glean/crawlindex/config"
	"cloudeng.io/sync/errgroup"
)

// Flags represents the flags that are used to control the crawl.
type Flags struct {
	config.FileFlags
	Outlinks bool `subcmd:"outlinks,false,display extracted outlinks"`
	Progress bool `subcmd:"progress,true,'display progress of downloads'"`
}

// Resources represents the resources that are used by the crawler.
type Resources struct {
	Extractors      map[content.Type]outlinks.Extractor
	PopulateCrawlFS func(ctx context.Context, cfg config.CrawlService, factories map[string]crawlcmd.FSFactory) error
	NewContentFS    func(ctx context.Context, cfg crawlcmd.CrawlCacheConfig) (content.FS, error)
}

// Crawler represents a crawler instance that contains global configuration
// information.
type Crawler struct {
	resourses Resources
}

// New creates a new Crawler instance.
func New(resources Resources) *Crawler {
	return &Crawler{resources}
}

func (c *Crawler) Run(ctx context.Context, fv *Flags, datasource string) error {
	cfg, err := config.DatasourceForName(ctx, fv.ConfigFile, datasource)
	if err != nil {
		return err
	}

	fsmap := map[string]crawlcmd.FSFactory{}
	for _, crawl := range cfg.Crawls {
		if err := c.resourses.PopulateCrawlFS(ctx, crawl.Service, fsmap); err != nil {
			return err
		}
	}
	resources := crawlcmd.Resources{
		Extractors:          c.resourses.Extractors,
		CrawlStoreFactories: fsmap,
		NewContentFS:        c.resourses.NewContentFS,
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
