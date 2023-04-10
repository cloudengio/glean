// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package index

import (
	"context"
	"fmt"
	"time"

	"cloudeng.io/algo/container/circular"
	"cloudeng.io/glean/gleansdk"
)

type bulkDocumentIndexer struct {
	bulkOptions
	client      *gleansdk.APIClient
	datasource  string
	id          string
	docs        *circular.Buffer[gleansdk.DocumentDefinition]
	indexed     int
	invocations int
	firstPage   bool
}

func newBulkDocumentIndexer(opts bulkOptions, client *gleansdk.APIClient, datasource string) *bulkDocumentIndexer {
	return &bulkDocumentIndexer{
		bulkOptions: opts,
		client:      client,
		id:          time.Now().Round(0).String(),
		firstPage:   true,
		datasource:  datasource,
		docs:        circular.NewBuffer[gleansdk.DocumentDefinition](opts.docReqSize),
	}
}

func (idx *bulkDocumentIndexer) handleDocuments(ctx context.Context, docs []gleansdk.DocumentDefinition) error {
	if len(docs) == 0 {
		return nil
	}
	idx.docs.Append(docs)
	return idx.handleNextRequests(ctx, idx.docReqSize, idx.docReqSize)
}

func (idx *bulkDocumentIndexer) handleNextRequests(ctx context.Context, atleast, atmost int) error {
	for {
		done, err := idx.handleNextRequest(ctx, atleast, atmost)
		if err != nil {
			return err
		}
		if done {
			return nil
		}
	}
}

func (idx *bulkDocumentIndexer) handleNextRequest(ctx context.Context, atleast, atmost int) (bool, error) {
	var req gleansdk.BulkIndexDocumentsRequest
	docs := itemsToIndex(idx.docs, atleast, atmost)
	if len(docs) == 0 {
		return true, nil
	}
	req.Datasource = idx.datasource
	req.SetIsFirstPage(idx.firstPage)
	if idx.firstPage {
		req.SetForceRestartUpload(idx.forceRestart)
	}
	req.SetIsLastPage(false)
	req.SetUploadId(idx.id)
	req.SetDisableStaleDocumentDeletionCheck(idx.forceDeletion)
	req.Documents = docs

	fmt.Printf("documents: sending request with %v (req size: %v)\n", len(req.Documents), idx.docReqSize)
	reqStart := time.Now()
	if idx.dryRun {
		fmt.Printf("documents: would index %v documents, %v invocations\n", len(req.Documents), idx.invocations+1)
	} else {
		resp, err := idx.client.DocumentsApi.BulkindexdocumentsPost(ctx).BulkIndexDocumentsRequest(req).Execute()
		if err != nil {
			return true, handleHTTPError(resp, err)
		}
	}
	idx.firstPage = false
	took := time.Since(reqStart)
	idx.indexed += len(req.Documents)
	idx.invocations++

	fmt.Printf("documents: indexed: total # docs: % 5v, per req # docs: % 3v in % 8v\n", len(req.Documents), len(req.Documents), took)

	return false, nil
}

func (idx *bulkDocumentIndexer) finish(ctx context.Context) error {
	if err := idx.handleNextRequests(ctx, 0, idx.docReqSize); err != nil {
		return err
	}
	if idx.dryRun {
		fmt.Printf("document: last page: sent %v documents over %v invocations\n", idx.indexed, idx.invocations)
		return nil
	}
	bulkReq := gleansdk.BulkIndexDocumentsRequest{}
	bulkReq.SetDatasource(idx.datasource)
	bulkReq.SetIsFirstPage(false)
	bulkReq.SetIsLastPage(true)
	bulkReq.SetUploadId(idx.id)
	bulkReq.Documents = []gleansdk.DocumentDefinition{}
	resp, err := idx.client.DocumentsApi.BulkindexdocumentsPost(ctx).BulkIndexDocumentsRequest(bulkReq).Execute()
	if err := handleHTTPError(resp, err); err != nil {
		return fmt.Errorf("document: last page request: %v", err)
	}
	return nil
}
