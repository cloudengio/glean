// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

// Package static provides the set of builtin resources for a given
// instance of the Glean CLI.
package static

import (
	"fmt"

	"cloudeng.io/file/content"
	"cloudeng.io/file/crawl/outlinks"
	gleancfg "cloudeng.io/glean/config"
	"cloudeng.io/glean/crawlindex/converters"
	"cloudeng.io/glean/extensions/benchling"
	"cloudeng.io/glean/extensions/biorxiv"
	"cloudeng.io/glean/extensions/papersapp"
	"cloudeng.io/glean/extensions/protocolsio"
	"cloudeng.io/webapi/operations/apitokens"
)

// CrawlProcessors represents the set of available Extractors for crawling.
type CrawlProcessors struct {
	Extractors map[content.Type]outlinks.Extractor
}

// IndexProcessor represents the set of available converters for indexing.
type IndexProcessors struct {
	DocumentConverters *content.Registry[converters.Document]
	UserConverters     *content.Registry[converters.User]
}

// Extractors represents the set of available outlink extractors.
func Extractors() map[content.Type]outlinks.Extractor {
	return map[content.Type]outlinks.Extractor{
		"text/html;charset=utf-8": outlinks.NewHTML(),
	}
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

func MustCrawlProcessors() CrawlProcessors {
	return CrawlProcessors{
		Extractors: Extractors(),
	}
}

func MustIndexProcessors() IndexProcessors {
	return IndexProcessors{
		DocumentConverters: MustDocumentConverters(),
		UserConverters:     MustUserConverters(),
	}
}

// APIExtensions returns the builtin API related commands.
func APIExtensions(parents ...string) []gleancfg.Extension {
	return NewExtensions(parents,
		protocolsio.ExtensionSpec,
		benchling.ExtensionSpec,
		papersapp.ExtensionSpec,
		biorxiv.ExtensionSpec,
	)
}

// NewExtensions creates the command specified extensions.
func NewExtensions(parents []string, specs ...gleancfg.ExtensionSpec) []gleancfg.Extension {
	var exts []gleancfg.Extension
	for _, spec := range specs {
		ext := gleancfg.NewExtension(spec, parents)
		exts = append(exts, ext)
	}
	return exts
}

func TokenReaders() *apitokens.Readers {
	def := apitokens.CloneReaders(apitokens.DefaultReaders)
	// Add any additional token readers go here.
	return def
}
