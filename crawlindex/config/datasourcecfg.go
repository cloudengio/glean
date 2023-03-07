// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package config

import (
	"fmt"

	"cloudeng.io/cmdutil"
	"cloudeng.io/file/content"
	"cloudeng.io/glean/config"
	"cloudeng.io/glean/gleansdk"
)

type DatasourceName struct {
	Datasource string `subcmd:"datasource,,name of the datasource"`
}

// FileFlags represents a command line flag for the datasource config file.
type FileFlags struct {
	ConfigFile string `subcmd:"datasource-configs,,datasource config file"`
}

// Datasource represents a single datasource or corpus to be crawled and
// indexed.
type Datasource struct {
	Datasource string // Datasource name.

	Crawls []Crawl `yaml:"crawls,omitempty"`
	// File/download orientend Crawls that obtain data for this datasource.

	// API based 'crawls' that obtain data for this datasource.
	APICrawls APICrawls `yaml:"api_crawls,omitempty"`

	// Bulk index configuration for this datasource.
	*BulkIndex `yaml:"bulk_index,omitempty"`
	// Incremental index configuration for this datasource.
	*IncrementalIndex `yaml:"incremental_index,omitempty"`

	Cache                  // Cache configuration for this datasource.
	Converters []Converter // Converters (from download.Result to Glean document) configuration.

	// The Glean datasource configuration in YAML as opposed to JSON
	// format.
	config.DatasourceConfig `yaml:"datasource_config"`
}

// Cache represents a cache configuration.
type Cache struct {
	Path string // Path is the location of the cache for this datasource.
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
func DatasourceForName(filename string, name string) (Datasource, error) {
	if len(filename) == 0 {
		return Datasource{}, fmt.Errorf("no datasource config file specified")
	}
	var cfg Datasources
	if err := cmdutil.ParseYAMLConfigFile(filename, &cfg); err != nil {
		fmt.Printf("oops: %v\n", filename)
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
				&d.DatasourceConfig.CustomDatasourceConfig,
			}
		}
	}
	return cnvmap
}
