// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package datasources

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"cloudeng.io/glean/crawlindex/config"
	"cloudeng.io/glean/crawlindex/internal"
	"cloudeng.io/glean/gleansdk"
	"cloudeng.io/webapi/operations/apitokens"
	"gopkg.in/yaml.v3"
)

func Download(ctx context.Context, domain, datasource string, token *apitokens.T) error {
	ctx, client := internal.NewIndexingClient(ctx, domain, token)
	return downloadDataSource(ctx, client, datasource)
}

func downloadDataSource(ctx context.Context, client *gleansdk.APIClient, datasource string) error {
	getDatasourceConfigRequest := gleansdk.NewGetDatasourceConfigRequest()
	getDatasourceConfigRequest.Name = &datasource
	resp, r, err := client.DatasourcesApi.GetdatasourceconfigPost(ctx).GetDatasourceConfigRequest(*getDatasourceConfigRequest).Execute()
	if err != nil {
		return internal.ParseGleanError(r, err)
	}
	out, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(out))
	return nil
}

func Register(ctx context.Context, cfg config.Datasource, name string, token *apitokens.T) error {
	buf, err := yaml.Marshal(cfg.GleanDatasource.CustomDatasourceConfig)
	if err != nil {
		return err
	}
	log.Printf("Registering custom datasource:\n%s\n", buf)

	ctx, client := internal.NewIndexingClient(ctx, cfg.GleanDomain, token)
	getDatasourceConfigRequest := gleansdk.NewGetDatasourceConfigRequest()
	getDatasourceConfigRequest.Name = &name
	r, err := client.DatasourcesApi.AdddatasourcePost(ctx).CustomDatasourceConfig(cfg.GleanDatasource.CustomDatasourceConfig).Execute()
	return internal.ParseGleanError(r, err)
}
