// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package index

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"cloudeng.io/errors"
	"cloudeng.io/glean/crawlindex/config"
	"cloudeng.io/glean/gleansdk"
)

// Bulk indexes a datasource in bulk mode.
func (idx *Indexer) Bulk(ctx context.Context, fv *BulkFlags) error {
	ctx, client := idx.newGleanIndexingClient(ctx)
	if idx.datasource.BulkIndex == nil {
		return fmt.Errorf("bulk_index block is missing from config file")
	}

	size := idx.datasource.BulkIndex.ReaddirEntries
	if size == 0 {
		size = 100
	}
	reqCh := make(chan Request, size)

	forceRestart, forceDeletion := idx.datasource.ForceRestart, idx.datasource.ForceDeletion
	if fv.ForceDeletion {
		forceDeletion = true
	}

	if fv.ForceRestart {
		forceRestart = true
	}

	cnvmap := idx.datasource.ConfigForContentType()
	if len(cnvmap) == 0 {
		return fmt.Errorf("no converters specified/found in config file")
	}

	walker := &walker{
		resources:   idx.resources,
		datasource:  idx.datasource,
		cnvConfig:   cnvmap,
		docCnv:      idx.resources.DocumentConverters,
		empCnv:      idx.resources.UserConverters,
		idxCh:       reqCh,
		concurrency: idx.datasource.BulkIndex.CacheConcurrency,
		scanSize:    idx.datasource.BulkIndex.ReaddirEntries,
	}

	indexer := newBulkIndexer(client, idx.datasource,
		WithForceDelete(forceDeletion),
		WithForceRestart(forceRestart),
		WithBulkID(fv.UploadID),
		WithReqSizes(idx.datasource.BulkIndex.DocumentRequestSize, idx.datasource.BulkIndex.UserRequestSize),
		WithUsers(idx.datasource.BulkIndex.UserRequestSize > 0),
		WithDryRun(fv.DryRun),
	)

	errs := &errors.M{}
	wg := &sync.WaitGroup{}
	wg.Add(2)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	go func() {
		errs.Append(walker.run(ctx))
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
	forceDeletion          bool
	forceRestart           bool
	id                     string
	dryRun                 bool
	docReqSize, empReqSize int
	users                  bool
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

func WithDryRun(v bool) BulkIndexOption {
	return func(o *bulkOptions) {
		o.dryRun = v
	}
}

func WithUsers(v bool) BulkIndexOption {
	return func(o *bulkOptions) {
		o.users = v
	}
}

func WithReqSizes(doc, emp int) BulkIndexOption {
	return func(o *bulkOptions) {
		o.docReqSize = doc
		o.empReqSize = emp
	}
}

// bulkIndexer represents a bulk indexer.
type bulkIndexer struct {
	documentIndexer *bulkDocumentIndexer
	userIndexer     *bulkUserIndexer
}

// bewBulkIndexer creates a new bulk indexer.
func newBulkIndexer(client *gleansdk.APIClient, datasource config.Datasource, opts ...BulkIndexOption) *bulkIndexer {
	var options bulkOptions
	for _, fn := range opts {
		fn(&options)
	}
	if options.docReqSize == 0 {
		options.docReqSize = 50
	}
	if options.empReqSize == 0 {
		options.empReqSize = 50
	}

	datasourceName := datasource.GleanDatasource.GetName()
	b := &bulkIndexer{
		documentIndexer: newBulkDocumentIndexer(
			options, client, datasourceName),
	}
	if options.users {
		b.userIndexer = newBulkUserIndexer(options, client, datasourceName)
	}
	return b
}

// Run runs the a bulk index operation, receiving requests to be indexed over
// the specified channel.
func (b *bulkIndexer) Run(ctx context.Context, ch <-chan Request) error {
	then := time.Now()
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case req, ok := <-ch:
			if !ok {
				if err := b.finish(ctx); err != nil {
					return err
				}
				if b.userIndexer != nil {
					timings(then, b.documentIndexer.indexed, b.userIndexer.indexed)
				} else {
					timings(then, b.documentIndexer.indexed, 0)
				}
				return nil
			}
			if err := b.documentIndexer.handleDocuments(ctx, req.Documents); err != nil {
				return err
			}
			if b.userIndexer == nil {
				continue
			}
			if err := b.userIndexer.handleUsers(ctx, req.Users); err != nil {
				return err
			}
		}
	}
}

func timings(then time.Time, docs, employees int) {
	taken := time.Since(then)
	if taken == 0 || docs+employees == 0 {
		return
	}
	avg := time.Duration(int64(taken) / int64(docs+employees))
	log.Printf("indexed: % 5v docs, % 5v employees in % 8v, (avg: %8v)\n", docs, employees, taken, avg)
}

// finish will send any buffered documents and then the last
// page.
func (b *bulkIndexer) finish(ctx context.Context) error {
	if err := b.documentIndexer.finish(ctx); err != nil {
		return err
	}
	if b.userIndexer == nil {
		return nil
	}
	return b.userIndexer.finish(ctx)
}
