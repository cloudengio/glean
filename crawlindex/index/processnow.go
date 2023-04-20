// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package index

import (
	"context"

	"cloudeng.io/glean/crawlindex/config"
	"cloudeng.io/glean/gleansdk"
)

// ProcessNowFlags represents the flags to the indexing process now command.
type ProcessNowFlags struct {
	config.FileFlags
}

func (idx *Indexer) ProcessNow(ctx context.Context, fv *ProcessNowFlags, datasource string) error {
	cfg, err := config.DatasourceForName(ctx, fv.ConfigFile, datasource)
	if err != nil {
		return err
	}
	ctx, client, err := idx.GleanConfig.NewIndexingAPIClient(ctx, cfg.GleanInstance)
	if err != nil {
		return err
	}
	req := gleansdk.NewProcessAllDocumentsRequest()
	req.SetDatasource(cfg.DatasourceConfig.Name)
	_, err = client.DocumentsApi.ProcessalldocumentsPost(ctx).ProcessAllDocumentsRequest(*req).Execute()
	return err
}
