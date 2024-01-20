// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package cmds

import (
	"context"
	"fmt"
	"strings"

	gleancfg "cloudeng.io/glean/config"
	"cloudeng.io/glean/crawlindex/datasources"
)

type Datasources struct {
	gleanCfgFile string
	extensions   []gleancfg.Extension
}

func (ds Datasources) Download(ctx context.Context, _ interface{}, args []string) error {
	d := datasources.T{GleanConfig: globalConfig}
	return d.Download(ctx, args[0], args[1])
}

func (ds Datasources) Register(ctx context.Context, values interface{}, args []string) error {
	d := datasources.T{GleanConfig: globalConfig}
	return d.Register(ctx, values.(*datasources.Flags), args[0])
}

func (ds Datasources) ShowConfig(ctx context.Context, _ interface{}, args []string) error {
	d := datasources.T{GleanConfig: globalConfig}
	return d.ShowConfig(ctx, args[0])
}

func (ds Datasources) ExplainConfig(ctx context.Context, _ interface{}, _ []string) error {
	d := datasources.T{GleanConfig: globalConfig}
	var out strings.Builder

	if err := d.ExplainConfig(ctx, &out, ds.gleanCfgFile, "connectors.yaml", ds.extensions); err != nil {
		return err
	}
	fmt.Println(out.String())
	return nil
}
