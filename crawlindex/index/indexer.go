// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package index

import (
	"context"

	"cloudeng.io/file/content"
	"cloudeng.io/file/crawl/crawlcmd"
	"cloudeng.io/glean/crawlindex/config"
	"cloudeng.io/glean/crawlindex/converters"
	"cloudeng.io/glean/crawlindex/internal"
	"cloudeng.io/glean/gleanclientsdk"
	"cloudeng.io/glean/gleansdk"
	"cloudeng.io/webapi/operations"
	"cloudeng.io/webapi/operations/apitokens"
)

// BulkFlags represents the flags to the bulk indexing command.
type BulkFlags struct {
	config.FileFlags
	UploadID      string `subcmd:"upload-id,upload,id to use for this bulk upload"`
	ForceRestart  bool   `subcmd:"force-restart,false,restart the bulk upload"`
	ForceDeletion bool   `subcmd:"force-sync-deletion,false,synchronously delete stale documents on upload of last bulk indexing batch"`
	DryRun        bool   `subcmd:"dry-run,false,'process but not do index documents'"`
}

// Resources represents the resources needed by an indexer.
type Resources struct {
	IndexingToken      *apitokens.T
	ClientToken        *apitokens.T
	DocumentConverters *content.Registry[converters.Document]
	UserConverters     *content.Registry[converters.User]
	NewOperationsFS    func(ctx context.Context, cfg crawlcmd.CrawlCacheConfig) (operations.FS, error)
}

// Indexer represents a Glean indexer.
type Indexer struct {
	datasource     config.Datasource
	resources      Resources
	datasourceName string
}

func New(ctx context.Context, fv config.FileFlags, datasource string, resources Resources) (*Indexer, error) {
	cfg, err := config.DatasourceForName(ctx, fv.ConfigFile, datasource)
	if err != nil {
		return nil, err
	}
	return &Indexer{
		datasource:     cfg,
		resources:      resources,
		datasourceName: cfg.GleanDatasource.GetName(),
	}, nil
}

func (idx *Indexer) newGleanIndexingClient(ctx context.Context) (context.Context, *gleansdk.APIClient) {
	return internal.NewIndexingClient(ctx, idx.datasource.GleanDomain, idx.resources.IndexingToken)
}

func (idx *Indexer) newGleanClient(ctx context.Context) (context.Context, *gleanclientsdk.APIClient) {
	return internal.NewClient(ctx, idx.datasource.GleanDomain, idx.resources.ClientToken)
}
