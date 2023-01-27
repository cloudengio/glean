// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package cmds

import (
	"context"

	"cloudeng.io/glean/crawlindex/crawl"
)

type Crawl struct{}

func (cmd *Crawl) Run(ctx context.Context, values interface{}, args []string) error {
	return crawl.Run(ctx, GlobalConfig, values.(*crawl.Flags), args[0])
}
