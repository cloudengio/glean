// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

// Package builtin provides the set of builtin resources for a given
// instance of the Glean CLI.
package builtin

import (
	"fmt"

	"cloudeng.io/aws/s3fs"
	"cloudeng.io/file"
	"cloudeng.io/file/content"
	"cloudeng.io/file/crawl/outlinks"
	gleancfg "cloudeng.io/glean/config"
	"cloudeng.io/glean/crawlindex/config"
	"cloudeng.io/glean/crawlindex/converters"
)

// Extractors represents the set of available outlink extractors.
func Extractors() map[content.Type]outlinks.Extractor {
	return map[content.Type]outlinks.Extractor{
		"text/html;charset=utf-8": outlinks.NewHTML(),
	}
}

// FSForCrawl returns a map of filesystem schemes to filesystem factories
// for use when creating crawling requests.
func FSForCrawl(cfg config.Crawl) map[string]file.FSFactory {
	return map[string]file.FSFactory{
		"s3": &s3fs.Factory{Config: cfg.AWS},
	}
}

//	MustConverters returns the available converters, it will panic
//
// on encountering an error.
func MustConverters() *content.Registry[converters.T] {
	cnv, err := converters.CreateRegistry(
		converters.NewHTML(),
		// Uncomment to add protocols.io support.
		//protocolsio.NewConverter(),
	)
	if err != nil {
		panic(fmt.Errorf("failed to load converters: %v", err))
	}
	return cnv
}

// APIExtensions returns the builtin API related commands.
func APIExtensions(parents ...string) []gleancfg.Extension {
	var exts []gleancfg.Extension
	// Uncomment to add protocols.io support.
	//exts = append(exts, protocolsio.Extension(parents...))
	return exts
}
