// Copyright 2022 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package crawl

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"cloudeng.io/errors"
	"cloudeng.io/file"
	"cloudeng.io/file/content"
	"cloudeng.io/file/crawl"
	"cloudeng.io/file/crawl/outlinks"
	"cloudeng.io/file/download"
	"cloudeng.io/glean/crawlindex"
	"cloudeng.io/glean/crawlindex/config"
	"cloudeng.io/path"
	"cloudeng.io/path/cloudpath"
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
	GleanConfig config.Glean
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
	for _, crawl := range cfg.Crawls {
		if err := crawl.CreateAndCleanCache(cachePath); err != nil {
			return err
		}
	}

	// Run all of the crawlers concurrently.
	g := errgroup.T{}
	for _, crawl := range cfg.Crawls {
		crawlCache := filepath.Join(cachePath, crawl.Cache.Prefix)
		crawler := crawler{
			name:        crawl.Name,
			gleanConfig: c.GleanConfig,
			extractors:  c.Extractors,
			fsForCrawl:  c.FSForCrawl,
			cachePath:   crawlCache,
			Flags:       fv,
			Crawl:       crawl,
			Datasource:  cfg,
		}
		g.Go(func() error {
			return crawler.run(ctx)
		})
	}
	return g.Wait()
}

type crawler struct {
	*Flags
	gleanConfig config.Glean
	extractors  func() map[content.Type]outlinks.Extractor
	fsForCrawl  func(config.Crawl) map[string]file.FSFactory
	name        string
	cachePath   string
	config.Crawl
	config.Datasource
}

func displayProgress(ctx context.Context, name string, progress <-chan download.Progress) {
	for {
		select {
		case <-ctx.Done():
			return
		case p := <-progress:
			fmt.Printf(" 16%v: % 8v: % 8v\n", name, p.Downloaded, p.Outstanding)
		}
	}
}

func (c *crawler) run(ctx context.Context) error {
	seedsByScheme, rejected := c.Crawl.SeedsByScheme(cloudpath.DefaultMatchers)
	if len(rejected) > 0 {
		return fmt.Errorf("unable to determine a file system for seeds: %v", rejected)
	}

	requests, err := c.Crawl.CreateSeedCrawlRequests(ctx, c.fsForCrawl(c.Crawl), seedsByScheme)
	if err != nil {
		return err
	}

	var progressCh chan download.Progress
	if c.Progress {
		progressCh = make(chan download.Progress, 100)
		go displayProgress(ctx, c.Name, progressCh)
	}

	dlFactory := c.Crawl.Download.NewFactory(progressCh)

	reqCh, crawledCh := c.Crawl.Download.Depth0Chans()

	extractorErrCh := make(chan outlinks.Errors, 100)

	crawler := crawl.New(crawl.WithNumExtractors(c.NumExtractors),
		crawl.WithCrawlDepth(c.Depth))

	linkProcessor, err := c.Crawl.NewLinkProcessor()
	if err != nil {
		return fmt.Errorf("failed to compile link processing rules: %v", err)
	}

	extractorRegistry, err := c.Crawl.ExtractorRegistry(c.extractors())
	if err != nil {
		return fmt.Errorf("failed to create extractor registry: %v", err)
	}

	extractor := outlinks.NewExtractors(extractorErrCh, linkProcessor, extractorRegistry)

	var errs errors.M
	var wg sync.WaitGroup
	wg.Add(3)

	go func(ch chan crawl.Crawled) {
		errs.Append(c.saveCrawled(ctx, c.name, ch))
		wg.Done()
	}(crawledCh)

	go func() {
		errs.Append(crawler.Run(ctx, dlFactory, extractor, reqCh, crawledCh))
		wg.Done()
	}()

	go func() {
		defer wg.Done()
		defer close(reqCh)
		for _, req := range requests {
			select {
			case <-ctx.Done():
				errs.Append(ctx.Err())
				return
			case reqCh <- req:
			}
		}
	}()

	go func() {
		for err := range extractorErrCh {
			if len(err.Errors) > 0 {
				fmt.Printf("extractor error: %v\n", err)
			}
		}
	}()

	wg.Wait()
	close(extractorErrCh)
	return errs.Err()
}

func (c crawler) saveCrawled(ctx context.Context, prefix string, crawledCh chan crawl.Crawled) error {
	sharder := path.NewSharder(path.WithSHA1PrefixLength(3))

	for crawled := range crawledCh {
		if c.Outlinks {
			for _, req := range crawled.Outlinks {
				fmt.Printf("%v\n", strings.Join(crawled.Request.Names(), " "))
				for _, name := range req.Names() {
					fmt.Printf("\t-> %v\n", name)
				}
			}
		}
		for _, dld := range crawled.Downloads {
			if dld.Err != nil {
				fmt.Printf("download error: %v: %v\n", dld.Name, dld.Err)
				continue
			}
			prefix, suffix := sharder.Assign(prefix + dld.Name)
			path := filepath.Join(c.cachePath, prefix, suffix)
			doc := crawlindex.NewDownloadDocument(&dld, content.TypeForPath(dld.Name))
			err := doc.WriteToFile(filepath.Join(c.cachePath, prefix), suffix)
			if err != nil {
				fmt.Printf("failed to write: %v as %v: %v\n", dld.Name, path, err)
				continue
			}
			fmt.Printf("%v -> %v\n", dld.Name, path)
		}
	}
	return nil
}
