// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package cmds

import (
	"context"

	"cloudeng.io/file/content"
	"cloudeng.io/glean/crawlindex/config"
	"cloudeng.io/glean/crawlindex/converters"
	"cloudeng.io/glean/crawlindex/index"
)

type Index struct {
	options
}

func newIndexer(ctx context.Context, configFile string, converters *content.Registry[converters.T], datasource string) (*index.Indexer, error) {
	cfg, err := config.DatasourceForName(ctx, configFile, datasource)
	if err != nil {
		return nil, err
	}
	indexer := &index.Indexer{
		GleanConfig: globalConfig,
		Config:      cfg,
		Converters:  converters,
	}
	return indexer, nil
}

func (cmd *Index) bulk(ctx context.Context, values interface{}, args []string) error {
	fv := values.(*index.BulkFlags)
	datasource := args[0]
	indexer, err := newIndexer(ctx, fv.ConfigFile, cmd.converters, datasource)
	if err != nil {
		return err
	}
	return indexer.Bulk(ctx, fv, args[0])
}

func (cmd *Index) stats(ctx context.Context, values interface{}, args []string) error {
	indexer := index.Indexer{
		GleanConfig: globalConfig,
	}
	return indexer.Stats(ctx, values.(*index.StatsFlags), args[0])
}
