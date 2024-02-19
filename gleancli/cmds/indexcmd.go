// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package cmds

import (
	"context"

	"cloudeng.io/glean/crawlindex/config"
	"cloudeng.io/glean/crawlindex/index"
)

type Index struct {
	Options
}

func (cmd *Index) bulk(ctx context.Context, values interface{}, args []string) error {
	fv := values.(*index.BulkFlags)
	datasource := args[0]
	cfg, err := config.DatasourceForName(ctx, fv.ConfigFile, datasource)
	if err != nil {
		return err
	}
	indexer := &index.Indexer{
		GleanConfig:        globalConfig,
		Config:             cfg,
		CreateStoreFS:      cmd.CreateStoreFS,
		DocumentConverters: cmd.IndexProcessors.DocumentConverters,
		UserConverters:     cmd.IndexProcessors.UserConverters,
	}
	return indexer.Bulk(ctx, fv)
}

func (cmd *Index) stats(ctx context.Context, values interface{}, args []string) error {
	indexer := index.Indexer{
		GleanConfig: globalConfig,
	}
	return indexer.Stats(ctx, values.(*index.StatsFlags), args[0])
}

func (cmd *Index) query(ctx context.Context, values interface{}, args []string) error {
	indexer := index.Indexer{
		GleanConfig: globalConfig,
	}
	return indexer.Query(ctx, values.(*index.QueryFlags), args[0], args[1])
}

func (cmd *Index) delete(ctx context.Context, values interface{}, args []string) error {
	indexer := index.Indexer{
		GleanConfig: globalConfig,
	}
	return indexer.Delete(ctx, values.(*index.DeleteFlags), args[0], args[1])
}

func (cmd *Index) deleteAll(ctx context.Context, values interface{}, args []string) error {
	indexer := index.Indexer{
		GleanConfig: globalConfig,
	}
	return indexer.DeleteAll(ctx, values.(*index.DeleteAllFlags), args[0])
}

func (cmd *Index) processNow(ctx context.Context, values interface{}, args []string) error {
	indexer := index.Indexer{
		GleanConfig: globalConfig,
	}
	return indexer.ProcessNow(ctx, values.(*index.ProcessNowFlags), args[0])
}
