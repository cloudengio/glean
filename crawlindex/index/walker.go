// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package index

import (
	"context"
	"fmt"
	"path/filepath"

	"cloudeng.io/file/content"
	"cloudeng.io/file/filewalk"
	"cloudeng.io/glean/crawlindex"
	"cloudeng.io/glean/crawlindex/config"
	"cloudeng.io/glean/crawlindex/converters"
	"cloudeng.io/glean/gleansdk"
)

type walker struct {
	datasource string
	cfg        map[content.Type]config.Conversion
	cnvs       *content.Registry[converters.T]
	wk         *filewalk.Walker
	ch         chan<- Request
}

func newWalker(datasource string, cfg map[content.Type]config.Conversion, cnvs *content.Registry[converters.T], scanSize int, ch chan<- Request) *walker {
	sc := filewalk.LocalFilesystem(scanSize)
	wk := filewalk.New(sc)
	return &walker{
		datasource: datasource,
		wk:         wk,
		cfg:        cfg,
		cnvs:       cnvs,
		ch:         ch,
	}
}

func (w *walker) dirs(ctx context.Context, prefix string, info *filewalk.Info, err error) (bool, []filewalk.Info, error) {
	return false, nil, nil
}

func (w *walker) files(ctx context.Context, prefix string, info *filewalk.Info, ch <-chan filewalk.Contents) ([]filewalk.Info, error) {
	children := make([]filewalk.Info, 0, 10)
	var req Request
	for {
		var contents filewalk.Contents
		var ok bool
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case contents, ok = <-ch:
			if !ok {
				if len(req.Documents) > 0 {
					select {
					case <-ctx.Done():
						return nil, nil
					case w.ch <- req:
					}
				}
				return children, nil
			}
		}
		for _, file := range contents.Files {
			dl, err := crawlindex.ReadDocument(prefix, file.Name)
			if err != nil {
				fmt.Printf("failed to read file: %v: %v\n", filepath.Join(prefix, file.Name), err)
				continue
			}
			gd, err := w.convert(dl)
			if err != nil {
				fmt.Printf("failed to convert: %v: %v\n", filepath.Join(prefix, file.Name), err)
				continue
			}
			req.Documents = append(req.Documents, gd)
		}
		children = append(children, contents.Children...)
	}
}

func (w *walker) convert(doc crawlindex.Document) (*gleansdk.DocumentDefinition, error) {
	converter, err := w.cnvs.LookupConverters(doc.Type, converters.GleanContentType)
	if err != nil {
		return nil, fmt.Errorf("no content converters found for %v to %v", doc.Type, converters.GleanContentType)
	}
	return converter.GleanDocument(w.datasource, w.cfg[content.Clean(doc.Type)], doc)
}

func (w *walker) Run(ctx context.Context, dir string) error {
	return w.wk.Walk(ctx, w.dirs, w.files, dir)
}
