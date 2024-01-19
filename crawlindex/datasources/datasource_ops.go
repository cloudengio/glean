// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package datasources

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"cloudeng.io/cmdutil/cmdyaml"
	"cloudeng.io/cmdutil/structdoc"
	"cloudeng.io/errors"
	gleancfg "cloudeng.io/glean/config"
	"cloudeng.io/glean/crawlindex/config"
	"cloudeng.io/glean/crawlindex/internal"
	"cloudeng.io/glean/gleansdk"
	"gopkg.in/yaml.v3"
)

type Flags struct {
	config.FileFlags
}

type T struct {
	GleanConfig gleancfg.Glean
}

func (d *T) Download(ctx context.Context, instance, datasource string) error {
	ctx, client, err := d.GleanConfig.NewIndexingAPIClient(ctx, instance)
	if err != nil {
		return err
	}
	return downloadDataSource(ctx, client, datasource)
}

func (d *T) Register(ctx context.Context, fv *Flags, datasource string) error {
	cfg, err := config.DatasourceForName(ctx, fv.ConfigFile, datasource)
	if err != nil {
		return err
	}
	buf, err := yaml.Marshal(cfg.CustomDatasourceConfig)
	if err != nil {
		return err
	}
	fmt.Printf("Registering custom datasource:\n%s\n", buf)

	ctx, client, err := d.GleanConfig.NewIndexingAPIClient(ctx, cfg.GleanInstance)
	if err != nil {
		return err
	}
	getDatasourceConfigRequest := gleansdk.NewGetDatasourceConfigRequest()
	getDatasourceConfigRequest.Name = &datasource
	r, err := client.DatasourcesApi.AdddatasourcePost(ctx).CustomDatasourceConfig(cfg.CustomDatasourceConfig).Execute()
	return internal.ParseGleanError(r, err)
}

func (d *T) ShowConfig(ctx context.Context, filename string) error {
	var cfg config.Datasources
	if err := cmdyaml.ParseConfigFile(ctx, filename, &cfg); err != nil {
		return err
	}
	buf, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", buf)
	return nil
}

func explain(out *strings.Builder, indent int, cfg any) error {
	desc, err := structdoc.Describe(cfg, "cmd", "YAML configuration file options\n")
	if err != nil {
		return err
	}
	out.WriteString(structdoc.FormatFields(indent, indent+2, desc.Fields))
	fmt.Println(out.String())
	return nil
}

func (d *T) ExplainConfig(ctx context.Context, out *strings.Builder,
	gleanCfgGile, connectorsCfgFile string,
	extensions []gleancfg.Extension) error {
	var errs errors.M

	fmt.Fprintf(out, "Configuration for Glean instances, typically stored in %v\n", gleanCfgGile)
	type gleanConfigs struct {
		GleanConfig gleancfg.Glean `yaml:"glean_config" cmd:"Glean configuration for authentication and API access"`
	}
	errs.Append(explain(out, 0, gleanConfigs{}))

	fmt.Fprintf(out, "Configuration for datasources/connectors, typically stored in %v specified as a command line flag\n", connectorsCfgFile)
	type datasources struct {
		Datasources config.Datasources `yaml:"datasources"`
	}
	errs.Append(explain(out, 0, datasources{}))

	fmt.Fprintf(out, "\nAPI Crawls. Each API Crawl has an authentication and a service section with the authentication info generally being stored separately to the service configuration.\n")
	for _, ext := range extensions {
		fmt.Fprintf(out, "\nAuthentication configuration for %v\n", ext.Name())
		errs.Append(explain(out, 2, ext.AuthConfigType()))
		fmt.Fprintf(out, "\nService configuration for %v\n", ext.Name())
		errs.Append(explain(out, 2, ext.ServiceConfigType()))
	}

	return errs.Err()
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
