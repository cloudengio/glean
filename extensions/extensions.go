// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

// Package extensions contains a set of subcmd extensions that can
// be used by gleancli implementations. Each such extension is a
// submodule so that its dependencies are imported only by
// gleancli implementations that use it.
package extensions

import (
	"cloudeng.io/cmdutil/subcmd"
	"cloudeng.io/glean/crawlindex/config"
)

type T interface {
	subcmd.Extension
	SetGleanConfig(*config.Glean)
	GleanConfig() *config.Glean
}

type extension struct {
	subcmd.Extension
	cfg *config.Glean
}

func (e *extension) SetGleanConfig(cfg *config.Glean) {
	e.cfg = cfg
}

func (e *extension) GleanConfig() *config.Glean {
	return e.cfg
}

func New(name, spec string, appendFn func(cmdSet *subcmd.CommandSetYAML) error) T {
	return &extension{
		Extension: subcmd.NewExtension(name, spec, appendFn),
	}
}
