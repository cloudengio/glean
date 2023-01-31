// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

// Package builtin provides the set of builtin resources for a given
// instance of the Glean CLI.
package builtin

import (
	"cloudeng.io/aws/s3fs"
	"cloudeng.io/file"
	"cloudeng.io/file/content"
	"cloudeng.io/file/crawl/outlinks"
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

// Converters returns the converters for the given configuration.
func Converters(cfg config.Converters) (*content.Registry[converters.T], error) {
	return converters.CreateRegistry(
		converters.NewHTML(),
	)
}
