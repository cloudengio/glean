// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package cmds

import (
	"cloudeng.io/file"
	"cloudeng.io/file/content"
	"cloudeng.io/file/crawl/outlinks"
	gleancfg "cloudeng.io/glean/config"
	"cloudeng.io/glean/crawlindex/config"
	"cloudeng.io/glean/crawlindex/converters"
)

type Option func(o *options)

type options struct {
	apiExtensions  []gleancfg.Extension
	extractors     func() map[content.Type]outlinks.Extractor
	fsForCrawl     func(cfg config.Crawl) map[string]file.FSFactory
	converters     *content.Registry[converters.T]
	gleanConfig    gleancfg.Glean
	useGleanConfig bool
}

func WithAPIExtensions(extensions []gleancfg.Extension) Option {
	return func(o *options) {
		o.apiExtensions = extensions
	}
}

func WithFSForCrawl(fsForCrawl func(cfg config.Crawl) map[string]file.FSFactory) Option {
	return func(o *options) {
		o.fsForCrawl = fsForCrawl
	}
}

func WithOutlinkExtractors(extractors func() map[content.Type]outlinks.Extractor) Option {
	return func(o *options) {
		o.extractors = extractors
	}
}

func WithConverters(converters *content.Registry[converters.T]) Option {
	return func(o *options) {
		o.converters = converters
	}
}

func WithGleanConfig(cfg gleancfg.Glean) Option {
	return func(o *options) {
		o.gleanConfig = cfg
		o.useGleanConfig = true
	}
}
