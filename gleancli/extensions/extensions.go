// Copyright 2024 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

// Package extensions defines an extension mechanism for gleancli
// implementations and utilities for creating, managing and
// implementing extensions.
package extensions

import (
	"context"

	"cloudeng.io/cmdutil/subcmd"
	"cloudeng.io/file/checkpoint"
	"cloudeng.io/file/content"
	"cloudeng.io/file/crawl/crawlcmd"
	"cloudeng.io/file/crawl/outlinks"
	"cloudeng.io/glean/crawlindex/config"
	"cloudeng.io/glean/crawlindex/converters"
	"cloudeng.io/webapi/operations"
	"cloudeng.io/webapi/operations/apitokens"
)

// DynamicResources provides a set of functions that can be used to create
// the various resources required by commands and command extensions.
type DynamicResources struct {
	PopulateCrawlFS func(ctx context.Context, cfg config.CrawlService, factories map[string]crawlcmd.FSFactory) error

	NewContentFS func(ctx context.Context, cfg crawlcmd.CrawlCacheConfig) (content.FS, error)

	NewCheckpointOp func(ctx context.Context, cfg crawlcmd.CrawlCacheConfig) (checkpoint.Operation, error)

	NewOperationsFS func(ctx context.Context, cfg crawlcmd.CrawlCacheConfig) (operations.FS, error)
}

// StaticResources provides a set of resources that are typically required
// by commands and command extensions that are statically configured
// per application instance.
type StaticResources struct {
	DocumentConverters *content.Registry[converters.Document]
	UserConverters     *content.Registry[converters.User]
	Extractors         map[content.Type]outlinks.Extractor

	// TokenReaders provides a set of readers for reading API tokens.
	TokenReaders *apitokens.Readers
}

// ExtensionsOptions are the options that are passed to each extension.
type ExtensionOptions struct {
	StaticResources

	DynamicResources
}

// Extension is implemented by all extensions.
type Extension interface {
	subcmd.Extension
	SetOptions(ExtensionOptions)
	Options() ExtensionOptions
	AuthConfigType() any
	ServiceConfigType() any
}

type extension struct {
	subcmd.Extension
	options        ExtensionOptions
	authCfgType    any
	serviceCfgType any
}

func (e *extension) SetOptions(opts ExtensionOptions) {
	e.options = opts
}

func (e *extension) Options() ExtensionOptions {
	return e.options
}

func (e *extension) AuthConfigType() any {
	return e.authCfgType
}

func (e *extension) ServiceConfigType() any {
	return e.serviceCfgType
}

// ExtensionSpec defines an extension.
type ExtensionSpec struct {
	Name       string // Name of the extension
	CmdSpec    string // YAML subcmd spec
	AuthCfg    any    // used for describing the auth configuration
	ServiceCfg any    // used for describing the service configuration
	// called to add the command to its parent cmdSet.
	AddFunc func(extension Extension, cmdSet *subcmd.CommandSetYAML, parents []string) error
}

// NewExtensions creates all of the command extensions specified by specs.
func NewExtensions(parents []string, specs ...ExtensionSpec) []Extension {
	var exts []Extension
	for _, spec := range specs {
		ext := NewExtension(spec, parents)
		exts = append(exts, ext)
	}
	return exts
}

// NewExtension creates a new extension from the supplied spec.
func NewExtension(spec ExtensionSpec, parents []string) Extension {
	ext := &extension{
		authCfgType:    spec.AuthCfg,
		serviceCfgType: spec.ServiceCfg,
	}
	ext.Extension = subcmd.NewExtension(spec.Name, spec.CmdSpec, func(cmdSet *subcmd.CommandSetYAML) error {
		return spec.AddFunc(ext, cmdSet, parents)
	})
	return ext
}
