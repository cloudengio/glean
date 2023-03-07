// Copyright 2022 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package crawl

import (
	"context"
	"fmt"
	"os"

	"cloudeng.io/file"
	"cloudeng.io/file/content"
	"cloudeng.io/file/crawl/crawlcmd"
	"cloudeng.io/file/crawl/outlinks"
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
	GleanConfig gleancfg.Glean
	Extractors  func() map[content.Type]outlinks.Extractor
	FSForCrawl  func(config.Crawl) map[string]file.FSFactory
}

func (c *Crawler) Run(ctx context.Context, fv *Flags, datasource string) error {
	cfg, err := config.DatasourceForName(fv.ConfigFile, datasource)
	if err != nil {
		return err
	}

	// Initialize the cache storage for crawled files.
	if len(cfg.Cache.Path) == 0 {
		return fmt.Errorf("no path specified for the cache to stored downloaded files")
	}

	cachePath := os.ExpandEnv(cfg.Cache.Path)
	if err := os.MkdirAll(cachePath, 0755); err != nil {
		return fmt.Errorf("failedto ensure that %v exists: %v", cfg.Cache.Path, err)
	}

	// Run all of the crawlers concurrently.
	g := errgroup.T{}
	for _, crawl := range cfg.Crawls {
		crawler := &crawlcmd.Crawler{
			Config:     crawl.Config,
			Extractors: c.Extractors,
		}
		g.Go(func() error {
			return crawler.Run(ctx, c.FSForCrawl(crawl), cachePath, fv.Outlinks, fv.Progress)
		})
	}
	return g.Wait()
}
