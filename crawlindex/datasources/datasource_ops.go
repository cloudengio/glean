// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package datasources

import (
	"context"
	"encoding/json"
	"fmt"

	"cloudeng.io/cmdutil"
	"cloudeng.io/glean/crawlindex/config"
	"cloudeng.io/glean/crawlindex/internal"
	"cloudeng.io/glean/gleansdk"
	"gopkg.in/yaml.v3"
)

type Flags struct {
	config.FileFlags
}

type Datasource struct {
	GleanConfig config.Glean
}

func (d *Datasource) Download(ctx context.Context, instance, datasource string) error {
	ctx, client, err := d.GleanConfig.NewAPIClient(ctx, instance)
	if err != nil {
		return err
	}
	return downloadDataSource(ctx, client, datasource)
}

func (d *Datasource) Register(ctx context.Context, fv *Flags, datasource string) error {
	cfg, err := config.DatasourceForName(fv.ConfigFile, datasource)
	if err != nil {
		return err
	}
	buf, err := yaml.Marshal(cfg.CustomDatasourceConfig)
	if err != nil {
		return err
	}
	fmt.Printf("Registering custom datasource:\n%s\n", buf)

	ctx, client, err := d.GleanConfig.NewAPIClient(ctx, cfg.GleanInstance)
	if err != nil {
		return err
	}
	getDatasourceConfigRequest := gleansdk.NewGetDatasourceConfigRequest()
	getDatasourceConfigRequest.Name = &datasource
	r, err := client.DatasourcesApi.AdddatasourcePost(ctx).CustomDatasourceConfig(cfg.CustomDatasourceConfig).Execute()
	return internal.ParseGleanError(r, err)
}

func (d *Datasource) ShowConfig(ctx context.Context, filename string) error {
	var cfg config.Datasources
	if err := cmdutil.ParseYAMLConfigFile(filename, &cfg); err != nil {
		return err
	}
	buf, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", buf)
	return nil
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
