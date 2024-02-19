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

	"cloudeng.io/cmdutil/subcmd"
	"cloudeng.io/file/checkpoint"
	"cloudeng.io/webapi/operations"
	"gopkg.in/yaml.v3"
)

type ExtensionOptions struct {
	Glean
	CreateStoreFS      func(ctx context.Context, path string, cfg yaml.Node) (operations.FS, error)
	CreateCheckpointOp func(ctx context.Context, path string, cfg yaml.Node) (checkpoint.Operation, error)
}

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

type ExtensionSpec struct {
	Name       string // Name of the extension
	CmdSpec    string // YAML subcmd spec
	AuthCfg    any    // used for describing the auth configuration
	ServiceCfg any    // used for describing the service configuration
	// called to add the command to its parent cmdSet.
	AddFunc func(extension Extension, cmdSet *subcmd.CommandSetYAML, parents []string) error
}

func NewExtensionC(spec ExtensionSpec, parents []string) Extension {
	ext := &extension{
		authCfgType:    spec.AuthCfg,
		serviceCfgType: spec.ServiceCfg,
	}
	ext.Extension = subcmd.NewExtension(spec.Name, spec.CmdSpec, func(cmdSet *subcmd.CommandSetYAML) error {
		return spec.AddFunc(ext, cmdSet, parents)
	})
	return ext
}

func NewExtension(name, spec string, authCfgType, serviceCfgType any, appendFn func(cmdSet *subcmd.CommandSetYAML) error) Extension {
	return &extension{
		Extension:      subcmd.NewExtension(name, spec, appendFn),
		serviceCfgType: serviceCfgType,
		authCfgType:    authCfgType,
	}
}
