// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package cmds

import (
	"context"
	"fmt"

	"cloudeng.io/cmdutil/cmdyaml"
	"cloudeng.io/glean/crawlindex"
	"cloudeng.io/glean/crawlindex/config"
	"cloudeng.io/glean/crawlindex/index"
	"cloudeng.io/webapi/operations/apitokens"
)

type Index struct {
	Options
	datasource                 config.Datasource
	indexingToken, clientToken *apitokens.T
}

type BulkFlags struct {
	index.BulkFlags
	crawlindex.AuthFileFlag
}

type StatsFlags struct {
	index.StatsFlags
	crawlindex.AuthFileFlag
}

type QueryFlags struct {
	index.QueryFlags
	crawlindex.AuthFileFlag
}

type DeleteFlags struct {
	index.DeleteFlags
	crawlindex.AuthFileFlag
}

type DeleteAllFlags struct {
	index.DeleteAllFlags
	crawlindex.AuthFileFlag
}

type ProcessNowFlags struct {
	index.ProcessNowFlags
	crawlindex.AuthFileFlag
}

func initTokens(ctx context.Context, domain string, tokenReaders *apitokens.Readers, authFile string) (indexingToken, clientToken *apitokens.T, err error) {

	var auth crawlindex.Auth
	if err = cmdyaml.ParseConfigFile(ctx, authFile, &auth); err != nil {
		return
	}
	indexingToken, clientToken = auth.TokensForDomain(domain)
	if indexingToken == nil || clientToken == nil {
		err = fmt.Errorf("no tokens found for domain: %v", domain)
		return
	}

	if indexingToken.Scheme == "" || clientToken.Scheme == "" {
		err = fmt.Errorf("invalid indexing or client token found for domain: %v", domain)
	}

	if err = indexingToken.Read(ctx, tokenReaders); err != nil {
		return
	}

	err = clientToken.Read(ctx, tokenReaders)
	return
}

func initConfigAndTokens(ctx context.Context, tokenReaders *apitokens.Readers, configFile, authFile, datasource string) (cfg config.Datasource, indexingToken, clientToken *apitokens.T, err error) {
	cfg, err = config.DatasourceForName(ctx, configFile, datasource)
	if err != nil {
		return
	}
	indexingToken, clientToken, err = initTokens(ctx, cfg.GleanDomain, tokenReaders, authFile)
	return
}

func (cmd *Index) init(ctx context.Context, configFile, authFile, datasource string) (err error) {
	cmd.datasource, cmd.indexingToken, cmd.clientToken, err = initConfigAndTokens(ctx, cmd.TokenReaders, configFile, authFile, datasource)
	return
}

func (cmd *Index) bulk(ctx context.Context, values interface{}, args []string) error {
	fv := values.(*BulkFlags)
	if err := cmd.init(ctx, fv.ConfigFile, fv.AuthFile, args[0]); err != nil {
		return err
	}
	indexer, err := index.New(ctx, fv.FileFlags, args[0], index.Resources{
		IndexingToken:    cmd.indexingToken,
		ClientToken:      cmd.clientToken,
		IndexProcessors:  cmd.StaticResources.IndexProcessors,
		DynamicResources: cmd.DynamicResources,
	})
	if err != nil {
		return err
	}
	return indexer.Bulk(ctx, &fv.BulkFlags)
}

func (cmd *Index) stats(ctx context.Context, values interface{}, args []string) error {
	fv := values.(*StatsFlags)
	if err := cmd.init(ctx, fv.ConfigFile, fv.AuthFile, args[0]); err != nil {
		return err
	}
	indexer, err := index.New(ctx, fv.FileFlags, args[0], index.Resources{
		IndexingToken: cmd.indexingToken,
		ClientToken:   cmd.clientToken,
	})
	if err != nil {
		return err
	}
	return indexer.Stats(ctx, &fv.StatsFlags)
}

func (cmd *Index) query(ctx context.Context, values interface{}, args []string) error {
	fv := values.(*QueryFlags)
	if err := cmd.init(ctx, fv.ConfigFile, fv.AuthFile, args[0]); err != nil {
		return err
	}
	indexer, err := index.New(ctx, fv.FileFlags, args[0], index.Resources{
		IndexingToken: cmd.indexingToken,
		ClientToken:   cmd.clientToken,
	})
	if err != nil {
		return err
	}
	return indexer.Query(ctx, &fv.QueryFlags, args[0], args[1])
}

func (cmd *Index) delete(ctx context.Context, values interface{}, args []string) error {
	fv := values.(*DeleteFlags)
	if err := cmd.init(ctx, fv.ConfigFile, fv.AuthFile, args[0]); err != nil {
		return err
	}
	indexer, err := index.New(ctx, fv.FileFlags, args[0], index.Resources{
		IndexingToken: cmd.indexingToken,
		ClientToken:   cmd.clientToken,
	})
	if err != nil {
		return err
	}
	return indexer.Delete(ctx, &fv.DeleteFlags, args[1])
}

func (cmd *Index) deleteAll(ctx context.Context, values interface{}, args []string) error {
	fv := values.(*DeleteAllFlags)
	if err := cmd.init(ctx, fv.ConfigFile, fv.AuthFile, args[0]); err != nil {
		return err
	}
	indexer, err := index.New(ctx, fv.FileFlags, args[0], index.Resources{
		IndexingToken: cmd.indexingToken,
		ClientToken:   cmd.clientToken,
	})
	if err != nil {
		return err
	}
	return indexer.DeleteAll(ctx, &fv.DeleteAllFlags)
}

func (cmd *Index) processNow(ctx context.Context, values interface{}, args []string) error {
	fv := values.(*ProcessNowFlags)
	if err := cmd.init(ctx, fv.ConfigFile, fv.AuthFile, args[0]); err != nil {
		return err
	}
	indexer, err := index.New(ctx, fv.FileFlags, args[0], index.Resources{
		IndexingToken: cmd.indexingToken,
		ClientToken:   cmd.clientToken,
	})
	if err != nil {
		return err
	}
	return indexer.ProcessNow(ctx, &fv.ProcessNowFlags)
}
