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
	datasource config.Datasource
}

type BulkFlags struct {
	index.BulkFlags
}

type StatsFlags struct {
	index.StatsFlags
}

type QueryFlags struct {
	index.QueryFlags
}

type DeleteFlags struct {
	index.DeleteFlags
}

type DeleteAllFlags struct {
	index.DeleteAllFlags
}

type ProcessNowFlags struct {
	index.ProcessNowFlags
}

func initConfig(ctx context.Context, configFile, datasource string) (cfg config.Datasource, err error) {
	cfg, err = config.DatasourceForName(ctx, configFile, datasource)
	if err != nil {
		return
	}

	return
}

func (cmd *Index) init(ctx context.Context, configFile, datasource string) (err error) {
	cmd.datasource, err = initConfig(ctx, configFile, datasource)
	return
}

func (cmd *Index) bulk(ctx context.Context, values interface{}, args []string) error {
	fv := values.(*BulkFlags)
	if err := cmd.init(ctx, fv.ConfigFile, args[0]); err != nil {
		return err
	}
	indexer, err := index.New(ctx, fv.FileFlags, args[0], index.Resources{
		DocumentConverters: cmd.DocumentConverters,
		UserConverters:     cmd.UserConverters,
		NewOperationsFS:    cmd.NewOperationsFS,
	})
	if err != nil {
		return err
	}
	return indexer.Bulk(ctx, &fv.BulkFlags)
}

func (cmd *Index) stats(ctx context.Context, values interface{}, args []string) error {
	fv := values.(*StatsFlags)
	if err := cmd.init(ctx, fv.ConfigFile, args[0]); err != nil {
		return err
	}
	indexer, err := index.New(ctx, fv.FileFlags, args[0], index.Resources{})
	if err != nil {
		return err
	}
	return indexer.Stats(ctx, &fv.StatsFlags)
}

func (cmd *Index) query(ctx context.Context, values interface{}, args []string) error {
	fv := values.(*QueryFlags)
	if err := cmd.init(ctx, fv.ConfigFile, args[0]); err != nil {
		return err
	}
	indexer, err := index.New(ctx, fv.FileFlags, args[0], index.Resources{})
	if err != nil {
		return err
	}
	return indexer.Query(ctx, &fv.QueryFlags, args[0], args[1])
}

func (cmd *Index) delete(ctx context.Context, values interface{}, args []string) error {
	fv := values.(*DeleteFlags)
	if err := cmd.init(ctx, fv.ConfigFile, args[0]); err != nil {
		return err
	}
	indexer, err := index.New(ctx, fv.FileFlags, args[0], index.Resources{})
	if err != nil {
		return err
	}
	return indexer.Delete(ctx, &fv.DeleteFlags, args[1])
}

func (cmd *Index) deleteAll(ctx context.Context, values interface{}, args []string) error {
	fv := values.(*DeleteAllFlags)
	if err := cmd.init(ctx, fv.ConfigFile, args[0]); err != nil {
		return err
	}
	indexer, err := index.New(ctx, fv.FileFlags, args[0], index.Resources{})
	if err != nil {
		return err
	}
	return indexer.DeleteAll(ctx, &fv.DeleteAllFlags)
}

func (cmd *Index) processNow(ctx context.Context, values interface{}, args []string) error {
	fv := values.(*ProcessNowFlags)
	if err := cmd.init(ctx, fv.ConfigFile, args[0]); err != nil {
		return err
	}
	indexer, err := index.New(ctx, fv.FileFlags, args[0], index.Resources{})
	if err != nil {
		return err
	}
	return indexer.ProcessNow(ctx, &fv.ProcessNowFlags)
}
