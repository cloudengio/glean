// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package index

import (
	"context"
	"log"
	"path/filepath"
	"strings"
	"sync"

	"cloudeng.io/file"
	"cloudeng.io/file/content"
	"cloudeng.io/file/content/stores"
	"cloudeng.io/file/filewalk"
	"cloudeng.io/glean/crawlindex/config"
	"cloudeng.io/glean/crawlindex/converters"
	"cloudeng.io/glean/gleansdk"
	"cloudeng.io/webapi/operations"
)

type skipdirs []string

func (sd skipdirs) skip(dir string) bool {
	dir = filepath.Clean(dir)
	for _, s := range sd {
		if strings.HasSuffix(dir, s) {
			return true
		}
	}
	return false
}

type walker struct {
	datasource string
	fs         operations.FS
	store      stores.T
	cfg        map[content.Type]config.Conversion
	docCnv     *content.Registry[converters.Document]
	empCnv     *content.Registry[converters.User]
	skip       skipdirs
	ch         chan<- Request
	wk         *filewalk.Walker[struct{}]
}

func newWalker(datasource string, fs operations.FS, cfg map[content.Type]config.Conversion, docCnv *content.Registry[converters.Document], empCnv *content.Registry[converters.User], skip skipdirs, concurrency, scanSize int, ch chan<- Request) *walker {
	idxWalker := &walker{
		datasource: datasource,
		fs:         fs,
		store:      stores.New(fs, concurrency),
		cfg:        cfg,
		docCnv:     docCnv,
		empCnv:     empCnv,
		ch:         ch,
		skip:       skip,
	}
	idxWalker.wk = filewalk.New(fs, idxWalker, filewalk.WithScanSize(scanSize))
	return idxWalker
}

func (w *walker) Prefix(_ context.Context, _ *struct{}, prefix string, _ file.Info, err error) (stop bool, children file.InfoList, returnErr error) {
	if err != nil {
		return false, nil, err
	}
	return w.skip.skip(prefix), nil, nil
}

func (w *walker) handleChildren(ctx context.Context, prefix string, contents []filewalk.Entry) (file.InfoList, error) {
	children := make(file.InfoList, 0, 10)
	for _, entry := range contents {
		if entry.IsDir() {
			stat, err := w.fs.Stat(ctx, w.fs.Join(prefix, entry.Name))
			if err != nil {
				return nil, err
			}
			children = append(children, stat)
		}
	}
	return children, nil
}

func (w *walker) Contents(ctx context.Context, _ *struct{}, prefix string, contents []filewalk.Entry) (file.InfoList, error) {
	doneCh := make(chan struct{})
	var children file.InfoList
	var childrenErr error
	go func() {
		children, childrenErr = w.handleChildren(ctx, prefix, contents)
		close(doneCh)
	}()

	names := make([]string, 0, len(contents))
	for _, entry := range contents {
		if entry.Type.IsRegular() {
			names = append(names, entry.Name)
		}
	}
	var req Request
	var mu sync.Mutex

	err := w.store.ReadV(ctx, prefix, names, func(ctx context.Context, prefix, name string, ctype content.Type, data []byte, err error) error {
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
		return nil, err
	}

	if len(req.Documents) > 0 || len(req.Users) > 0 {
		select {
		case w.ch <- req:
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}

	<-doneCh
	return children, childrenErr
}

func (w *walker) Done(_ context.Context, _ *struct{}, _ string, err error) error {
	return err
}

func (w *walker) convertDocument(ctx context.Context, ctype content.Type, obj []byte) (gleansdk.DocumentDefinition, bool, error) {
	converter, err := w.docCnv.LookupConverters(ctype, converters.GleanDocumentType)
	if err != nil {
		return gleansdk.DocumentDefinition{}, false, nil
	}
	doc, err := converter.Convert(ctx, w.datasource, w.cfg[content.Clean(ctype)], ctype, obj)
	return doc, true, err
}

func (w *walker) convertUser(ctx context.Context, ctype content.Type, obj []byte) (gleansdk.DatasourceUserDefinition, bool, error) {
	converter, err := w.empCnv.LookupConverters(ctype, converters.GleanUserType)
	if err != nil {
		return gleansdk.DatasourceUserDefinition{}, false, nil
	}
	user, err := converter.Convert(ctx, w.datasource, w.cfg[content.Clean(ctype)], ctype, obj)
	return user, true, err
}

func (w *walker) Run(ctx context.Context, dir string) error {
	return w.wk.Walk(ctx, dir)
}
