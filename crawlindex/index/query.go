// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package index

import (
	"context"
	"fmt"

	"cloudeng.io/glean/crawlindex/config"
	"cloudeng.io/glean/gleanclientsdk"
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
	fmt.Printf("CFG: %v\n", idx.Config)
	ctx, client, err := idx.GleanConfig.NewClientAPIClient(ctx, cfg.GleanInstance)
	if err != nil {
		return err
	}
	si := gleanclientsdk.SearchRequestSourceInfo{
		Modality: "EMBEDDED_SEARCH",
	}
	sreq := gleanclientsdk.NewSearchRequest(si)
	results, _, err := client.SearchApi.Search(ctx).Payload(*sreq).Execute()
	if err != nil {
		return err
	}
	fmt.Printf("RESULTS: %v\n", results)
	return nil
}
