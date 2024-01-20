// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

// Package builtin provides the set of builtin resources for a given
// instance of the Glean CLI.
package builtin

import (
	"fmt"

	"cloudeng.io/aws/awsconfig"
	"cloudeng.io/aws/s3fs"
	"cloudeng.io/file"
	"cloudeng.io/file/content"
	"cloudeng.io/file/crawl/outlinks"
	gleancfg "cloudeng.io/glean/config"
	"cloudeng.io/glean/crawlindex/config"
	"cloudeng.io/glean/crawlindex/converters"
	"cloudeng.io/glean/extensions/benchling"
	"cloudeng.io/glean/extensions/biorxiv"
	"cloudeng.io/glean/extensions/papersapp"
	"cloudeng.io/glean/extensions/protocolsio"
)

// Extractors represents the set of available outlink extractors.
func Extractors() map[content.Type]outlinks.Extractor {
	return map[content.Type]outlinks.Extractor{
		"text/html;charset=utf-8": outlinks.NewHTML(),
	}
}

// FSForCrawl returns a map of filesystem schemes to filesystem factories
// for use when creating crawling requests.
func FSForCrawl(cfg config.Crawl) (map[string]file.FSFactory, error) {
	if cfg.ServiceName == "s3" || cfg.ServiceName == "aws" {
		var awscfg awsconfig.AWSFlags
		if err := cfg.ServiceConfig.Decode(&awscfg); err != nil {
			return nil, err
		}
		return map[string]file.FSFactory{
			"s3": &s3fs.Factory{Config: awscfg},
		}, nil
	}
	return nil, nil
}

// MustDocumentConverters returns the available converters,
// it will panic on encountering an error.
func MustDocumentConverters() *content.Registry[converters.Document] {
	cnv, err := converters.CreateDocumentRegistry(
		converters.NewHTML(),
		protocolsio.NewDocumentConverter(),
		benchling.NewDocumentConverter(),
		papersapp.NewDocumentConverter(),
		biorxiv.NewDocumentConverter(),
	)
	if err != nil {
		panic(fmt.Errorf("failed to load document converters: %v", err))
	}
	return cnv
}

// MustUserConverters returns the available converters,
// it will panic on encountering an error.
func MustUserConverters() *content.Registry[converters.User] {
	cnv, err := converters.CreateUserRegistry(
		benchling.NewUserConverter(),
	)
	if err != nil {
		panic(fmt.Errorf("failed to load employee converters: %v", err))
	}
	return cnv
}

// APIExtensions returns the builtin API related commands.
func APIExtensions(parents ...string) []gleancfg.Extension {
	var exts []gleancfg.Extension
	exts = append(exts,
		protocolsio.Extension(parents...),
		benchling.Extension(parents...),
		papersapp.Extension(parents...),
		biorxiv.Extension(parents...),
	)
	return exts
}
