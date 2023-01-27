// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package cmds

import (
	"context"

	"cloudeng.io/glean/crawlindex/datasources"
)

type Datasources struct{}

func (ds Datasources) Download(ctx context.Context, values interface{}, args []string) error {
	return datasources.Download(ctx, GlobalConfig, args[0], args[1])
}

func (ds Datasources) Register(ctx context.Context, values interface{}, args []string) error {
	return datasources.Register(ctx, GlobalConfig, values.(*datasources.Flags), args[0])
}

func (ds Datasources) ShowConfig(ctx context.Context, values interface{}, args []string) error {
	return datasources.ShowConfig(ctx, args[0])
}
