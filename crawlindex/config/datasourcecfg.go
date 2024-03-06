// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package config

import (
	"context"
	"fmt"

	"cloudeng.io/cmdutil/cmdyaml"
	"cloudeng.io/file/content"
	"cloudeng.io/glean/gleansdk"
	"cloudeng.io/webapi/operations/apicrawlcmd"
)

type DatasourceName struct {
	Datasource string `subcmd:"datasource,,name of the datasource"`
}

// FileFlags represents a command line flag for the datasource config file.
type FileFlags struct {
	ConfigFile string `subcmd:"datasource-configs,,datasource config file"`
}

// GleanConfig represents the configuration of the datasource with
// Glean's API.
type GleanConfig struct {
	GleanInstance                   string `yaml:"glean_instance" cmd:"glean instance to use"`
	gleansdk.CustomDatasourceConfig `yaml:",inline" cmd:"glean custom datasource configuration"`
}

// Datasource represents a single datasource or corpus to be crawled and
// indexed.
type Datasource struct {
	// Datasource name.
	Datasource string `yaml:"datasource" cmd:"name of the datasource"`

	Crawls []Crawl `yaml:"crawls,omitempty" cmd:"file based crawls to run for this datasource"`
	// File/download orientend Crawls that obtain data for this datasource.

	// API based 'crawls' that obtain data for this datasource.
	APICrawls apicrawlcmd.Crawls `yaml:"api_crawls,omitempty" cmd:"api crawls to run for this datasource"`

	// Bulk index configuration for this datasource.
	*BulkIndex `yaml:"bulk_index,omitempty" cmd:"bulk index configuration for this datasource"`

	// Incremental index configuration for this datasource.
	*IncrementalIndex `yaml:"incremental_index,omitempty" cmd:"incremental index configuration for this datasource"`

	// Converters (from download.Result to Glean document) configuration.
	Converters []Converter `yaml:"converters,omitempty" cmd:"converters for this datasource"`

	// The Glean datasource configuration in YAML as opposed to JSON
	// format.
	GleanConfig `yaml:"datasource_config" cmd:"glean datasource configuration, ie. the glean datasource to be indexed"`
}

// Datasources represents a list of named datasources.
type Datasources []Datasource

// ConfigForName for returns the configuration for the named datasource.
func (d Datasources) ConfigForName(name string) (Datasource, bool) {
	for _, ds := range d {
		if ds.Datasource == name {
			return ds, true
		}
	}
	return Datasource{}, false
}

// DatasourceForName returns the datasource configuration for the named datasource
// read from the specified config file.
func DatasourceForName(ctx context.Context, filename string, name string) (Datasource, error) {
	if len(filename) == 0 {
		return Datasource{}, fmt.Errorf("no datasource config file specified")
	}
	var cfg Datasources
	if err := cmdyaml.ParseConfigFile(ctx, filename, &cfg); err != nil {
		return Datasource{}, err
	}
	ds, ok := cfg.ConfigForName(name)
	if !ok {
		return Datasource{}, fmt.Errorf("no datasource config found for %q in %q", name, filename)
	}
	return ds, nil
}

// Conversion represents the configuration for a single content conversion operation.
type Conversion struct {
	Type       content.Type
	Converter  *Converter
	Datasource *gleansdk.CustomDatasourceConfig
}

// ConfigForContentType returns a map from content type to all of the configuration
// information that pertains to	that content type.
func (d Datasource) ConfigForContentType() map[content.Type]Conversion {
	cnvmap := make(map[content.Type]Conversion)
	for i, c := range d.Converters {
		for _, ft := range c.FromContentType {
			cnvmap[content.Clean(ft)] = Conversion{
				ft,
				&d.Converters[i],
				&d.GleanConfig.CustomDatasourceConfig,
			}
		}
	}
	return cnvmap
}
