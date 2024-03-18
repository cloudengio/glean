// Copyright 2024 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package testcmd

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"

	"cloudeng.io/cmdutil/subcmd"
	"cloudeng.io/file/checkpoint"
	"cloudeng.io/file/crawl/crawlcmd"
	"cloudeng.io/file/filewalk"
	"cloudeng.io/glean/crawlindex/config"
	"cloudeng.io/glean/gleancli/extensions"
	"cloudeng.io/webapi/operations"
)

type CommonFlags struct {
	config.FileFlags
}

type CacheFlags struct {
	CommonFlags
	Items int `subcmd:"show-n,100,'show this many items from the cache'"`
}

const (
	cmdName = "test"
	cmdSpec = `
- name: test
  summary: test configuration for the glean cli
  commands:
    - name: cache
      summary: test access to cache locations for a datasource
      arguments:
        - datasource-name - the datasource to be tested
`
)

var ExtensionSpec = extensions.ExtensionSpec{
	Name:       cmdName,
	CmdSpec:    cmdSpec,
	AuthCfg:    struct{}{},
	ServiceCfg: struct{}{},
	AddFunc:    AddExtension,
}

func AddExtension(extension extensions.Extension, cmdSet *subcmd.CommandSetYAML, parents []string) error {
	c := &command{parent: extension}
	cmdSet.Set(append(parents, "cache")...).MustRunner(c.testCaches, &CacheFlags{})
	return nil
}

type command struct {
	parent extensions.Extension
	fs     operations.FS
	chkpt  checkpoint.Operation
}

func (cmd *command) testCaches(ctx context.Context, values interface{}, args []string) error {
	fv := values.(*CacheFlags)
	datasource := args[0]
	cfg, err := config.DatasourceForName(ctx, fv.ConfigFile, datasource)
	if err != nil {
		return err
	}

	var caches []crawlcmd.CrawlCacheConfig
	for _, c := range cfg.Crawls {
		caches = append(caches, c.Cache)
	}
	for _, c := range cfg.APICrawls {
		caches = append(caches, c.Cache)
	}

	for _, cache := range caches {
		fs, err := cmd.parent.Options().NewOperationsFS(ctx, cache)
		if err != nil {
			return err
		}
		n, err := scanCache(ctx, datasource, fs, cache, fv.Items)
		if err != nil {
			return err
		}
		downloads, _ := cache.Paths()
		log.Printf("found %v total items in cache for %v in %v\n", n, datasource, downloads)
	}
	return nil
}

func scanCache(ctx context.Context, datasource string, fs operations.FS, cache crawlcmd.CrawlCacheConfig, showN int) (int, error) {
	items := 0
	downloads, _ := cache.Paths()

	if _, err := fs.Stat(ctx, downloads); err != nil {
		return 0, fmt.Errorf("failed to stat %v: %w", downloads, err)
	}

	handler := func(ctx context.Context, prefix string, contents []filewalk.Entry, err error) error {
		if err != nil {
			return err
		}
		for _, c := range contents {
			if items >= showN {
				return filewalk.SkipAll
			}
			items++
			p := fs.Join(prefix, c.Name)
			fmt.Printf("%v\n", p)
			if items%100 == 0 {
				log.Printf("found %v items in cache for %v\n", items, datasource)
			}

		}
		return nil
	}
	err := filewalk.ContentsOnly(ctx, fs, downloads, handler,
		filewalk.WithScanSize(showN))
	if !errors.Is(err, io.EOF) {
		return items, err
	}
	return items, nil
}
