// Copyright 2024 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package extensions

import (
	"context"

	"cloudeng.io/cmdutil/cmdyaml"
	"cloudeng.io/glean/crawlindex/config"
	"cloudeng.io/webapi/operations/apicrawlcmd"
	"cloudeng.io/webapi/operations/apitokens"
	"gopkg.in/yaml.v3"
)

// APIKey represents an API key in apitokens format.
type APIKey struct {
	APIKey string `yaml:"api_key" cmd:"API key in apitokens format (ie. scheme://<value>)"`
}

// ParseAndRead reads the API key from the specified file and returns it in apitokens format.
func (k *APIKey) ParseAndRead(ctx context.Context, filename string, readers *apitokens.Readers) (*apitokens.T, error) {
	if err := cmdyaml.ParseConfigFile(ctx, filename, &k); err != nil {
		return nil, err
	}
	tok := apitokens.New(k.APIKey)
	if err := tok.Read(ctx, readers); err != nil {
		return nil, err
	}
	return tok, nil
}

func (eo ExtensionOptions) ResourcesForDatasource(ctx context.Context, configFile, authFile, datasource string) (config.Datasource, apicrawlcmd.Resources, error) {
	cfg, err := config.DatasourceForName(ctx, configFile, datasource)
	if err != nil {
		return config.Datasource{}, apicrawlcmd.Resources{}, err
	}
	var token *apitokens.T
	if len(authFile) > 0 {
		var key APIKey
		token, err = key.ParseAndRead(ctx, authFile, eo.TokenReaders)
		if err != nil {
			return config.Datasource{}, apicrawlcmd.Resources{}, err
		}
	}
	return cfg, apicrawlcmd.Resources{
		Token:           token,
		NewOperationsFS: eo.NewOperationsFS,
		NewCheckpointOp: eo.NewCheckpointOp,
	}, nil
}

func FirstAPICrawl(crawls apicrawlcmd.Crawls) (apicrawlcmd.Crawl[yaml.Node], bool) {
	for _, crawl := range crawls {
		return crawl, true
	}
	return apicrawlcmd.Crawl[yaml.Node]{}, false
}
