// Copyright 2022 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package crawl

import (
	"context"
	"os"

	"cloudeng.io/file/content"
	"cloudeng.io/file/crawl/crawlcmd"
	"cloudeng.io/file/crawl/outlinks"
	gleancfg "cloudeng.io/glean/config"
	"cloudeng.io/glean/crawlindex/config"
	"cloudeng.io/sync/errgroup"
	"cloudeng.io/webapi/operations"
	"gopkg.in/yaml.v3"
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
	GleanConfig   gleancfg.Glean
	Extractors    map[content.Type]outlinks.Extractor
	CreateCrawlFS func(service string, cfg yaml.Node, fsmap map[string]crawlcmd.FSFactory) error
	CreateStoreFS func(ctx context.Context, path string, cfg yaml.Node) (operations.FS, error)
}

func (c *Crawler) Run(ctx context.Context, fv *Flags, datasource string) error {
	cfg, err := config.DatasourceForName(ctx, fv.ConfigFile, datasource)
	if err != nil {
		return err
	}

	ofs, err := c.CreateStoreFS(ctx, cfg.Cache.Path, cfg.Cache.ServiceConfig)
	if err != nil {
		return err
	}

	fsmap := map[string]crawlcmd.FSFactory{}
	for _, crawl := range cfg.Crawls {
		if err := c.CreateCrawlFS(crawl.ServiceName, crawl.ServiceConfig, fsmap); err != nil {
			return err
		}
	}

	cacheRoot := os.ExpandEnv(cfg.Cache.Path)
	// Run all of the crawlers concurrently.
	g := errgroup.T{}
	for _, crawl := range cfg.Crawls {
		crawler := &crawlcmd.Crawler{
			Config:     crawl.Config,
			Extractors: c.Extractors,
		}
		g.Go(func() error {
			return crawler.Run(ctx, fsmap, ofs, cacheRoot, fv.Outlinks, fv.Progress)
		})
	}
	return g.Wait()
}
