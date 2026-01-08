// Copyright 2024 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package extensions

import (
	"context"

	"cloudeng.io/glean/crawlindex/config"
	"cloudeng.io/webapi/operations/apicrawlcmd"
	"gopkg.in/yaml.v3"
)


func (eo ExtensionOptions) ResourcesForDatasource(ctx context.Context, configFile, datasource string) (config.Datasource, apicrawlcmd.Resources, error) {
	cfg, err := config.DatasourceForName(ctx, configFile, datasource)
	if err != nil {
		return config.Datasource{}, apicrawlcmd.Resources{}, err
	}
	return cfg, apicrawlcmd.Resources{
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
