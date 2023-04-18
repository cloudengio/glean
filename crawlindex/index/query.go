// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package index

import (
	"context"
	"fmt"
	"math"
	"strings"
	"time"

	"cloudeng.io/glean/crawlindex/config"
	"cloudeng.io/glean/gleanclientsdk"
	"cloudeng.io/glean/gleansdk"
)

// QueryFlags represents the flags to the indexing query command.
type QueryFlags struct {
	config.FileFlags
	NumDocuments       int  `subcmd:"num-documents,10,number of documents to return"`
	ShowIndexingStatus bool `subcmd:"show-indexing-status,true,show indexing status for all returned documents"`
}

func (idx *Indexer) Query(ctx context.Context, fv *QueryFlags, datasource string, query string) error {
	cfg, err := config.DatasourceForName(ctx, fv.ConfigFile, datasource)
	if err != nil {
		return err

	}
	ctx, client, err := idx.GleanConfig.NewClientAPIClient(ctx, cfg.GleanInstance)
	if err != nil {
		return err
	}
	si := gleanclientsdk.NewSearchRequestSourceInfo("FULLPAGE")
	for _, c := range idx.GleanConfig {
		if c.Name == cfg.GleanInstance {
			si.SetDomain(c.API.Domain)
			break
		}
	}
	si.SetIsDebug(true)
	sreq := gleanclientsdk.NewSearchRequest(*si)
	sreq.SetQuery(query)
	if fv.NumDocuments >= math.MaxInt32 {
		fv.NumDocuments = 1000
	}
	sreq.SetPageSize(int32(fv.NumDocuments))
	sreq.RequestOptions = gleanclientsdk.NewSearchRequestOptions(100)
	sreq.RequestOptions.SetDatasourceFilter(cfg.DatasourceConfig.Name)
	sreq.RequestOptions.SetDisableSpellcheck(true)
	sreq.RequestOptions.SetDisableQueryAutocorrect(true)
	results, _, err := client.SearchApi.Search(ctx).Payload(*sreq).Execute()
	if err != nil {
		return err
	}
	for _, fr := range results.FacetResults {
		if fr.SourceName == nil {
			continue
		}
		switch *fr.SourceName {
		case "datasource":
			for _, b := range fr.Buckets {
				fmt.Printf("datasource: %v: %v\n", b.Value.GetStringValue(), b.GetCount())
			}
		case "last_updated_at":
			for _, b := range fr.Buckets {
				fmt.Printf("updated: %v: %v\n", b.Value.GetStringValue(), b.GetCount())
			}
		default:
			continue
		}
	}
	return idx.indexingStatus(ctx, cfg, results.Results)
}

func (idx *Indexer) indexingStatus(ctx context.Context, cfg config.Datasource, results []gleanclientsdk.SearchResult) error {
	ctx, client, err := idx.GleanConfig.NewIndexingAPIClient(ctx, cfg.GleanInstance)
	if err != nil {
		return err
	}
	for _, r := range results {
		docid := r.Document.GetId()
		if idx := strings.Index(docid, "_etr_"); idx != -1 {
			docid = docid[idx+1:]
		}
		sr := gleansdk.NewGetDocumentStatusRequest(cfg.DatasourceConfig.Name,
			strings.ToLower(r.Document.GetDocType()),
			docid)

		resp, _, err := client.DocumentsApi.GetdocumentstatusPost(ctx).GetDocumentStatusRequest(*sr).Execute()
		if err != nil {
			return err
		}
		uploaded := time.Unix(resp.GetLastUploadedAt(), 0)
		indexed := time.Unix(resp.GetLastIndexedAt(), 0)
		if resp.GetIndexingStatus() == "INDEXED" {
			fmt.Printf("%v: %v\t%v\t%v\t(%v) - %v : %v\n",
				docid, resp.GetIndexingStatus(), uploaded, indexed, indexed.Sub(uploaded), r.Document.GetDocType(), r.Document.GetId())
		} else {
			fmt.Printf("WARNING: %v: %v - %v : %v\n",
				docid, resp.GetIndexingStatus(), r.Document.GetDocType(), r.Document.GetId())
		}
		time.Sleep(time.Second)
	}
	return nil
}
