// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

// Package extensions contains a set of subcmd extensions that can
// be used by gleancli implementations. Each such extension is a
// submodule so that its dependencies are imported only by
// gleancli implementations that use it.
package config

import (
	"context"

	"cloudeng.io/cmdutil/cmdyaml"
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

// CrawlProcessors represents the set of available Extractors for crawling.
type CrawlProcessors struct {
	Extractors map[content.Type]outlinks.Extractor
}

// IndexProcessor represents the set of available converters for indexing.
type IndexProcessors struct {
	DocumentConverters *content.Registry[converters.Document]
	UserConverters     *content.Registry[converters.User]
}

// StaticResources provides a set of resources that are typically required
// by commands and command extensions that are statically configured
// per application instance.
type StaticResources struct {
	CrawlProcessors CrawlProcessors
	IndexProcessors IndexProcessors
	TokenReaders    *apitokens.Readers
}

// ExtensionsOptions are the options that are passed to each extension.
type ExtensionOptions struct {
	Glean

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

type APIKey struct {
	APIKey string `yaml:"api_key" cmd:"API key in apitokens format"`
}

func (k *APIKey) ParseAndRead(ctx context.Context, filename string, readers *apitokens.Readers) (*apitokens.T, error) {
	if err := cmdyaml.ParseConfigFile(ctx, filename, &k); err != nil {
		return nil, err
	}
	tok := apitokens.New(k.APIKey)
	if err := tok.Read(ctx, readers); err != nil {
		return nil, err
	}
	return tok, nil
}
