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
	AuthConfigType() any
	ServiceConfigType() any
}

type extension struct {
	subcmd.Extension
	cfg            *Glean
	authCfgType    any
	serviceCfgType any
}

func (e *extension) SetGleanConfig(cfg *Glean) {
	e.cfg = cfg
}

func (e *extension) GleanConfig() *Glean {
	return e.cfg
}

func (e *extension) AuthConfigType() any {
	return e.authCfgType
}

func (e *extension) ServiceConfigType() any {
	return e.serviceCfgType
}

func NewExtension(name, spec string, authCfgType, serviceCfgType any, appendFn func(cmdSet *subcmd.CommandSetYAML) error) Extension {
	return &extension{
		Extension:      subcmd.NewExtension(name, spec, appendFn),
		serviceCfgType: serviceCfgType,
		authCfgType:    authCfgType,
	}
}
