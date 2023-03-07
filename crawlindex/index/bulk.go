// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package index

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"cloudeng.io/algo/container/circular"
	"cloudeng.io/errors"
	"cloudeng.io/file/content"
	"cloudeng.io/glean/crawlindex/config"
	"cloudeng.io/glean/crawlindex/converters"
	"cloudeng.io/glean/gleansdk"
)

// BulkFlags represents the flags to the bulk indexing command.
type BulkFlags struct {
	config.FileFlags
	UploadID      string `subcmd:"upload-id,upload,id to use for this bulk upload"`
	ForceRestart  bool   `subcmd:"force-restart,false,restart the bulk upload"`
	ForceDeletion bool   `subcmd:"force-sync-deletion,false,synchronously delete stale documents on upload of last bulk indexing batch"`
}

// Indexer represents a Glean indexer.
type Indexer struct {
	GleanConfig config.Glean
	Config      config.Datasource
	Converters  *content.Registry[converters.T]
}

// Bulk indexes a datasource in bulk mode.
func (idx *Indexer) Bulk(ctx context.Context, fv *BulkFlags, datasource string) error {
	cachePath := os.ExpandEnv(idx.Config.Cache.Path)
	if len(cachePath) == 0 {
		return fmt.Errorf("no path specified for the cache to be indexed")
	}

	ctx, client, err := idx.GleanConfig.NewAPIClient(ctx, idx.Config.GleanInstance)
	if err != nil {
		return err
	}

	if idx.Config.BulkIndex == nil {
		return fmt.Errorf("bulk_index block is missing from config file")
	}

	size := idx.Config.BulkIndex.ReaddirEntries
	if size == 0 {
		size = 100
	}
	reqCh := make(chan Request, size)

	forceRestart, forceDeletion := idx.Config.ForceRestart, idx.Config.ForceDeletion
	if fv.ForceDeletion {
		forceDeletion = true
	}

	if fv.ForceRestart {
		forceRestart = true
	}

	cnvmap := idx.Config.ConfigForContentType()
	if len(cnvmap) == 0 {
		return fmt.Errorf("no converters specified/found in config file")
	}

	walker := newWalker(idx.Config.CustomDatasourceConfig.GetName(), cnvmap, idx.Converters, idx.Config.BulkIndex.ReaddirEntries, reqCh)
	indexer := newBulkIndexer(client, idx.Config,
		WithForceDelete(forceDeletion),
		WithForceRestart(forceRestart),
		WithBulkID(fv.UploadID))

	errs := &errors.M{}
	wg := &sync.WaitGroup{}
	wg.Add(2)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	go func() {
		errs.Append(walker.Run(ctx, cachePath))
		close(reqCh)
		wg.Done()
	}()

	go func() {
		err := indexer.Run(ctx, reqCh)
		if err != nil {
			cancel()
		}
		errs.Append(err)
		wg.Done()

	}()
	wg.Wait()
	return errs.Err()
}

// BulkIndexOption represents an option to NewBulkIndexer.
type BulkIndexOption func(o *bulkOptions)

type bulkOptions struct {
	forceDeletion bool
	forceRestart  bool
	id            string
}

// WithForceDelete disables the deletion of too many documents test.
func WithForceDelete(forceDelete bool) BulkIndexOption {
	return func(o *bulkOptions) {
		o.forceDeletion = forceDelete
	}
}

// WithForceRestart sets the force restart options.
func WithForceRestart(forceRestart bool) BulkIndexOption {
	return func(o *bulkOptions) {
		o.forceRestart = forceRestart
	}
}

// WithBulkID specifies a custom id to use for the bulk upload. If one
// is not specified the current date and time are used.
func WithBulkID(id string) BulkIndexOption {
	return func(o *bulkOptions) {
		o.id = id
	}
}

// bulkIndexer represents a bulk indexer.
type bulkIndexer struct {
	bulkOptions
	id       string
	req_size int
	cfg      config.Datasource
	client   *gleansdk.APIClient
	docs     *circular.Buffer[gleansdk.DocumentDefinition]
}

// bewBulkIndexer creates a new bulk indexer.
func newBulkIndexer(client *gleansdk.APIClient, datasource config.Datasource, opts ...BulkIndexOption) *bulkIndexer {
	b := &bulkIndexer{
		id:       time.Now().Round(0).String(),
		cfg:      datasource,
		req_size: datasource.BulkIndex.RequestSize,
		client:   client,
	}
	for _, fn := range opts {
		fn(&b.bulkOptions)
	}
	if b.req_size == 0 {
		b.req_size = 50
	}
	b.docs = circular.NewBuffer[gleansdk.DocumentDefinition](b.req_size)
	return b
}

// Run runs the a bulk index operation, receiving requests to be indexed over
// the specified channel.
func (b *bulkIndexer) Run(ctx context.Context, ch <-chan Request) error {
	var (
		firstPage = true
		indexed   = 0
		duration  time.Duration
	)
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case req, ok := <-ch:
			if !ok {
				if err := b.finish(ctx); err != nil {
					return err
				}
				if duration > 0 && indexed > 0 {
					avg := time.Duration(int64(duration) / int64(indexed))
					fmt.Printf("indexed: all # docs: % 5v docs in % 8v, (avg: %8v)\n", indexed, duration, avg)
				}
				return nil
			}
			if len(req.Documents) == 0 {
				continue
			}
			b.docs.Append(req.Documents)
			for {
				bulkReq, ok := b.nextBulkRequest(firstPage)
				if !ok {
					break
				}
				if err := b.sendAndTimeRequst(ctx, bulkReq); err != nil {
					return err
				}
				firstPage = false
			}
		}
	}
}

func (b *bulkIndexer) sendAndTimeRequst(ctx context.Context, req gleansdk.BulkIndexDocumentsRequest) error {
	fmt.Printf("sending request with %v\n", len(req.Documents))
	reqStart := time.Now()
	resp, err := b.client.DocumentsApi.BulkindexdocumentsPost(ctx).BulkIndexDocumentsRequest(req).Execute()
	if err := handleHTTPError(resp, err); err != nil {
		return err
	}
	took := time.Since(reqStart)
	fmt.Printf("indexed: total # docs: % 5v, per req # docs: % 3v in % 8v\n", len(req.Documents), len(req.Documents), took)
	return nil
}

// nextBulkRequest creates and returns a bulk index request
// with exactly 0 or b.req_size documents. If the boolean return
// value is false then zero documents are available (and up to
// n_req may still be buffered).
func (b *bulkIndexer) nextBulkRequest(firstPage bool) (gleansdk.BulkIndexDocumentsRequest, bool) {
	var req gleansdk.BulkIndexDocumentsRequest
	if b.docs.Len() < b.req_size {
		return req, false
	}
	return b.nextRequest(firstPage)
}

func (b *bulkIndexer) nextRequest(firstPage bool) (gleansdk.BulkIndexDocumentsRequest, bool) {
	var req gleansdk.BulkIndexDocumentsRequest
	docs := b.docs.Head(b.req_size)
	if len(docs) == 0 {
		return req, false
	}
	req.Datasource = b.cfg.CustomDatasourceConfig.GetName()
	req.SetIsFirstPage(firstPage)
	if firstPage {
		req.SetForceRestartUpload(b.forceRestart)
		firstPage = false
	}
	req.SetIsLastPage(false)
	req.SetUploadId(b.id)
	req.SetDisableStaleDocumentDeletionCheck(b.forceDeletion)
	req.Documents = docs
	return req, true
}

// finish will send any buffered documents and then the last
// page.
func (b *bulkIndexer) finish(ctx context.Context) error {
	for {
		bulkReq, ok := b.nextRequest(false)
		if !ok {
			break
		}
		if err := b.sendAndTimeRequst(ctx, bulkReq); err != nil {
			return err
		}
	}
	bulkReq := gleansdk.BulkIndexDocumentsRequest{}
	bulkReq.SetDatasource(b.cfg.CustomDatasourceConfig.GetName())
	bulkReq.SetIsFirstPage(false)
	bulkReq.SetIsLastPage(true)
	bulkReq.SetUploadId(b.id)
	bulkReq.Documents = []gleansdk.DocumentDefinition{}
	resp, err := b.client.DocumentsApi.BulkindexdocumentsPost(ctx).BulkIndexDocumentsRequest(bulkReq).Execute()
	if err := handleHTTPError(resp, err); err != nil {
		return fmt.Errorf("last page request: %v", err)
	}
	return nil
}
