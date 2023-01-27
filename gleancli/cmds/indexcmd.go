// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package cmds

import (
	"context"

	"cloudeng.io/glean/crawlindex/index"
)

type Index struct{}

func (cmd *Index) bulk(ctx context.Context, values interface{}, args []string) error {
	return index.Bulk(ctx, GlobalConfig, values.(*index.BulkFlags), args[0])
}

func (cmd *Index) stats(ctx context.Context, values interface{}, args []string) error {
	return index.Stats(ctx, GlobalConfig, values.(*index.StatsFlags), args[0])
}
