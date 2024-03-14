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
	"cloudeng.io/glean/crawlindex/converters"
	"cloudeng.io/glean/extensions/benchling"
	"cloudeng.io/glean/extensions/biorxiv"
	"cloudeng.io/glean/extensions/papersapp"
	"cloudeng.io/glean/extensions/protocolsio"
	"cloudeng.io/glean/extensions/testcmd"
	"cloudeng.io/glean/gleancli/extensions"
	"cloudeng.io/webapi/operations/apitokens"
)

// LinkExtractors represents the set of available outlink extractors.
func LinkExtractors() map[content.Type]outlinks.Extractor {
	return map[content.Type]outlinks.Extractor{
		"text/html;charset=utf-8": outlinks.NewHTML(),
	}
}

// MustDocumentConverters returns the available converters,
// it will panic on encountering an error.
func MustDocumentConverters() *content.Registry[converters.Document] {
	cnv, err := converters.CreateDocumentRegistry(
		NewHTML(), // Use the Glean HTML converter.
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
func APIExtensions(parents ...string) []extensions.Extension {
	return extensions.NewExtensions(parents,
		protocolsio.ExtensionSpec,
		benchling.ExtensionSpec,
		papersapp.ExtensionSpec,
		biorxiv.ExtensionSpec,
	)
}

func TokenReaders() *apitokens.Readers {
	def := apitokens.CloneReaders(apitokens.DefaultReaders)
	// Add any additional token readers go here, eg. for reading
	// from AWS secrets manager.
	return def
}

func New() extensions.StaticResources {
	return extensions.StaticResources{
		Extractors:         LinkExtractors(),
		DocumentConverters: MustDocumentConverters(),
		UserConverters:     MustUserConverters(),
		TokenReaders:       TokenReaders(),
	}
}

// Extensions returns any top-level commands.
func Extensions(parents ...string) []extensions.Extension {
	return extensions.NewExtensions(parents, testcmd.ExtensionSpec)
}
