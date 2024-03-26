// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package index

import (
	"context"
	"log"
	"path/filepath"
	"sync"

	"cloudeng.io/file/content"
	"cloudeng.io/file/content/stores"
	"cloudeng.io/file/crawl/crawlcmd"
	"cloudeng.io/file/filewalk"
	"cloudeng.io/glean/crawlindex/config"
	"cloudeng.io/glean/crawlindex/converters"
	"cloudeng.io/glean/gleansdk"
	"cloudeng.io/sync/errgroup"
	"cloudeng.io/webapi/operations"
)

type walker struct {
	resources   Resources
	datasource  config.Datasource
	cnvConfig   map[content.Type]config.Conversion
	docCnv      *content.Registry[converters.Document]
	empCnv      *content.Registry[converters.User]
	idxCh       chan<- Request
	concurrency int
	scanSize    int
}

func (w *walker) run(ctx context.Context) error {
	var caches []crawlcmd.CrawlCacheConfig
	for _, c := range w.datasource.Crawls {
		caches = append(caches, c.Cache)
	}
	for _, c := range w.datasource.APICrawls {
		caches = append(caches, c.Cache)
	}
	var cacheWalkers []cacheWalker
	for _, c := range caches {
		w, err := w.newCacheWalker(ctx, c, w.concurrency)
		if err != nil {
			return err
		}
		cacheWalkers = append(cacheWalkers, *w)
	}
	opts := []filewalk.Option{filewalk.WithScanSize(w.scanSize)}
	var g errgroup.T
	for _, w := range cacheWalkers {
		w := w
		g.Go(func() error {
			return filewalk.ContentsOnly(ctx, w.fs, w.root, w.ContentsHandler, opts...)
		})
	}
	return g.Wait()
}

func (w *walker) newCacheWalker(ctx context.Context, crawl crawlcmd.CrawlCacheConfig, concurrency int) (*cacheWalker, error) {
	fs, err := w.resources.NewOperationsFS(ctx, crawl)
	store := stores.New(fs, concurrency)
	return &cacheWalker{
		idxCh:      w.idxCh,
		root:       crawl.DownloadPath(),
		fs:         fs,
		store:      store,
		datasource: w.datasource.GleanDatasource.GetName(),
		cnvConfig:  w.cnvConfig,
		docCnv:     w.docCnv,
		empCnv:     w.empCnv,
	}, err
}

type cacheWalker struct {
	idxCh      chan<- Request
	root       string
	fs         operations.FS
	store      stores.T
	datasource string
	cnvConfig  map[content.Type]config.Conversion
	docCnv     *content.Registry[converters.Document]
	empCnv     *content.Registry[converters.User]
}

func (w *cacheWalker) ContentsHandler(ctx context.Context, prefix string, contents []filewalk.Entry, err error) error {
	if err != nil {
		if w.fs.IsPermissionError(err) {
			log.Printf("permission error: %v: %v\n", prefix, err)
			return nil
		}
		return err
	}
	var req Request
	var mu sync.Mutex

	names := make([]string, len(contents))
	for i, entry := range contents {
		names[i] = entry.Name
	}

	err = w.store.ReadV(ctx, prefix, names, func(ctx context.Context, prefix, name string, ctype content.Type, data []byte, err error) error {
		if err != nil {
			log.Printf("failed to read: %v: %v\n", filepath.Join(prefix, name), err)
			return err
		}
		gd, ok, err := w.convertDocument(ctx, ctype, data)
		if err != nil {
			log.Printf("failed to convert: %v: %v\n", filepath.Join(prefix, name), err)
			return err
		}
		if ok {
			mu.Lock()
			req.Documents = append(req.Documents, gd)
			mu.Unlock()
		}
		ge, ok, err := w.convertUser(ctx, ctype, data)
		if err != nil {
			log.Printf("failed to convert: %v: %v\n", filepath.Join(prefix, name), err)
			return err
		}
		if ok {
			mu.Lock()
			req.Users = append(req.Users, ge)
			mu.Unlock()
		}
		return nil
	})
	if err != nil {
		return err
	}

	if len(req.Documents) > 0 || len(req.Users) > 0 {
		select {
		case w.idxCh <- req:
		case <-ctx.Done():
			return ctx.Err()
		}
	}
	return nil
}

func (w *cacheWalker) convertDocument(ctx context.Context, ctype content.Type, obj []byte) (gleansdk.DocumentDefinition, bool, error) {
	converter, err := w.docCnv.LookupConverters(ctype, converters.GleanDocumentType)
	if err != nil {
		return gleansdk.DocumentDefinition{}, false, nil
	}
	doc, err := converter.Convert(ctx, w.datasource, w.cnvConfig[content.Clean(ctype)], ctype, obj)
	return doc, true, err
}

func (w *cacheWalker) convertUser(ctx context.Context, ctype content.Type, obj []byte) (gleansdk.DatasourceUserDefinition, bool, error) {
	converter, err := w.empCnv.LookupConverters(ctype, converters.GleanUserType)
	if err != nil {
		return gleansdk.DatasourceUserDefinition{}, false, nil
	}
	user, err := converter.Convert(ctx, w.datasource, w.cnvConfig[content.Clean(ctype)], ctype, obj)
	return user, true, err
}
