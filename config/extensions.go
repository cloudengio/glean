// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

// Package extensions contains a set of subcmd extensions that can
// be used by gleancli implementations. Each such extension is a
// submodule so that its dependencies are imported only by
// gleancli implementations that use it.
package config

import (
	"cloudeng.io/cmdutil/subcmd"
)

type Extension interface {
	subcmd.Extension
	SetGleanConfig(*Glean)
	GleanConfig() *Glean
}

type extension struct {
	subcmd.Extension
	cfg *Glean
}

func (e *extension) SetGleanConfig(cfg *Glean) {
	e.cfg = cfg
}

func (e *extension) GleanConfig() *Glean {
	return e.cfg
}

func NewExtension(name, spec string, appendFn func(cmdSet *subcmd.CommandSetYAML) error) Extension {
	return &extension{
		Extension: subcmd.NewExtension(name, spec, appendFn),
	}
}
