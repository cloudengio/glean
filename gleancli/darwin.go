// Copyright 2025 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

//go:build darwin

package main

import (
	"context"
	"fmt"

	"cloudeng.io/cmdutil/keys"
	"cloudeng.io/cmdutil/subcmd"
	"cloudeng.io/file/localfs"
	"cloudeng.io/glean/gleancli/cmds"
	"cloudeng.io/logging/ctxlog"
	"cloudeng.io/macos/keychain/plugin"
	"cloudeng.io/security/keys/keychain/plugins"
	"github.com/oasdiff/yaml"
)

func platformSpecificConfig() cmds.PlatformSpecificConfig {
	return &macosConfig{}
}

type macosFlags struct {
	plugin.ReadFlags
	cmds.LocalKeysAndLoggingFlags
	Name string `subcmd:"keychain-item,,'name of the macOS keychain item to use for storing/retrieving tokens'"`
}

type macosConfig struct {
	macosFlags
}

func (m *macosConfig) RegisterGlobalFlags(cmdset *subcmd.CommandSetYAML) {
	fs := subcmd.NewFlagSet().MustRegisterFlagStruct(&m.macosFlags, nil, nil)
	cmdset.WithGlobalFlags(fs)
}

func (m *macosConfig) InitContext(ctx context.Context) (context.Context, error) {
	logger := m.LoggingConfig().NewLoggerMust()
	ctx = ctxlog.WithLogger(ctx, logger.Logger)
	ims := keys.NewInMemoryKeyStore()
	cfg := m.Config()
	fs := plugins.NewFS(cfg.Binary, cfg)
	contents, err := fs.ReadFileCtx(ctx, m.Name)
	if err == nil {
		if err := yaml.Unmarshal(contents, ims); err != nil {
			return nil, err
		}
		fmt.Printf("using keys from key chain item %v\n", m.Name)
		return keys.ContextWithKeyStore(ctx, ims), nil
	}
	if m.LocalKeyFile == "" {
		return ctx, nil
	}
	lfs := localfs.New()
	contents, err = lfs.ReadFileCtx(ctx, m.LocalKeyFile)
	if err != nil {
		return nil, err
	}
	if err := yaml.Unmarshal(contents, ims); err != nil {
		return nil, err
	}
	fmt.Printf("WARNING: using keys from local key file %v\n", m.LocalKeyFile)
	return keys.ContextWithKeyStore(ctx, ims), nil
}
