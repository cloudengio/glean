// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package cmds

import (
	"context"
	"fmt"
	"strings"

	"cloudeng.io/cmdutil/cmdyaml"
	"cloudeng.io/cmdutil/structdoc"
	"cloudeng.io/errors"
	"cloudeng.io/glean/crawlindex"
	"cloudeng.io/glean/crawlindex/config"
	"cloudeng.io/glean/crawlindex/datasources"
	"cloudeng.io/glean/gleancli/extensions"
	"gopkg.in/yaml.v3"
)

type DownloadFlags struct {
	crawlindex.AuthFileFlag
}

type RegisterFlags struct {
	config.FileFlags
	crawlindex.AuthFileFlag
}

type Datasources struct {
	Options
	extensions []extensions.Extension
}

func (ds Datasources) Download(ctx context.Context, values any, args []string) error {
	fv := values.(*DownloadFlags)
	domain, name := args[0], args[1]
	indexingToken, _, err := initTokens(ctx, domain, ds.TokenReaders, fv.AuthFile)
	if err != nil {
		return err
	}
	return datasources.Download(ctx, domain, name, indexingToken)
}

func (ds Datasources) Register(ctx context.Context, values interface{}, args []string) error {
	fv := values.(*RegisterFlags)
	datasource := args[0]
	cfg, err := config.DatasourceForName(ctx, fv.ConfigFile, datasource)
	if err != nil {
		return err
	}
	indexingToken, _, err := initTokens(ctx, cfg.GleanDomain, ds.TokenReaders, fv.AuthFile)
	if err != nil {
		return err
	}
	return datasources.Register(ctx, cfg, datasource, indexingToken)
}

func (ds Datasources) ShowConfig(ctx context.Context, _ interface{}, args []string) error {
	var cfg config.Datasources
	if err := cmdyaml.ParseConfigFile(ctx, args[0], &cfg); err != nil {
		return err
	}
	buf, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", buf)
	return nil

}

func (ds Datasources) ExplainConfig(_ context.Context, _ interface{}, _ []string) error {
	var out strings.Builder
	if err := explainConfig(&out, ds.extensions); err != nil {
		return err
	}
	fmt.Println(out.String())
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

func explainConfig(out *strings.Builder, extensions []extensions.Extension) error {
	var errs errors.M

	fmt.Fprintf(out, "Authentication for Glean instances")

	type gleanConfigs struct {
		crawlindex.AuthFileFlag
	}
	errs.Append(explain(out, 0, gleanConfigs{}))

	fmt.Fprintf(out, "Configuration for datasources/connectors, typically specified as a command line flag\n")
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
