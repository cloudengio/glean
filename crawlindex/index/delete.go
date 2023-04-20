// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package index

import (
	"context"
	"fmt"
	"strings"
	"time"

	"cloudeng.io/glean/crawlindex/config"
	"cloudeng.io/glean/gleanclientsdk"
	"cloudeng.io/glean/gleansdk"
)

// DeleteFlags represents the flags to the indexing delete command.
type DeleteFlags struct {
	config.FileFlags
	NumDocuments int    `subcmd:"num-documents,100,number of documents to return"`
	Type         string `subcmd:"doc-type,,type of object to delete"`
}

// DeleteAllFlags represents the flags to the indexing delete-all command.
type DeleteAllFlags struct {
	config.FileFlags
}

func (idx *Indexer) Delete(ctx context.Context, fv *DeleteFlags, datasource string, query string) error {
	if len(fv.Type) == 0 {
		return fmt.Errorf("must specify object type to delete")
	}
	cfg, err := config.DatasourceForName(ctx, fv.ConfigFile, datasource)
	if err != nil {
		return err
	}
	results, err := idx.query(ctx, fv.NumDocuments, cfg, query)
	if err != nil {
		return err
	}
	return idx.delete(ctx, cfg, results.Results, fv.Type)
}

func (idx *Indexer) delete(ctx context.Context, cfg config.Datasource, results []gleanclientsdk.SearchResult, objectType string) error {
	ctx, client, err := idx.GleanConfig.NewIndexingAPIClient(ctx, cfg.GleanInstance)
	if err != nil {
		return err
	}
	for _, r := range results {
		if r.Document.GetDocType() != objectType {
			continue
		}
		objType := strings.ToLower(r.Document.GetDocType())
		docid := r.Document.GetId()
		prefix := fmt.Sprintf("CUSTOM_%v_%v_", strings.ToUpper(cfg.DatasourceConfig.Name), objType)
		docid = strings.TrimPrefix(docid, prefix)
		fmt.Printf("deleting: %v %v\n", objType, docid)
		sr := gleansdk.NewDeleteDocumentRequest(cfg.DatasourceConfig.Name,
			objType, docid)
		_, err := client.DocumentsApi.DeletedocumentPost(ctx).DeleteDocumentRequest(*sr).Execute()
		if err != nil {
			fmt.Printf("ERR: %v: docid %v, objType: %v\n", err, docid, objType)
			return err
		}
		time.Sleep(time.Second)
	}
	return nil
}

func (idx *Indexer) DeleteAll(ctx context.Context, fv *DeleteAllFlags, datasource string) error {
	cfg, err := config.DatasourceForName(ctx, fv.ConfigFile, datasource)
	if err != nil {
		return err
	}
	ctx, client, err := idx.GleanConfig.NewIndexingAPIClient(ctx, cfg.GleanInstance)
	if err != nil {
		return err
	}
	bulkReq := gleansdk.BulkIndexDocumentsRequest{}
	bulkReq.SetDatasource(cfg.CustomDatasourceConfig.Name)
	bulkReq.SetForceRestartUpload(true)
	bulkReq.SetIsFirstPage(true)
	bulkReq.SetIsLastPage(true)
	bulkReq.SetUploadId(time.Now().Round(0).String())
	bulkReq.SetDisableStaleDocumentDeletionCheck(true)
	bulkReq.Documents = []gleansdk.DocumentDefinition{}
	resp, err := client.DocumentsApi.BulkindexdocumentsPost(ctx).BulkIndexDocumentsRequest(bulkReq).Execute()
	if err := handleHTTPError(resp, err); err != nil {
		return fmt.Errorf("delete all: %v", err)
	}
	return nil
}
