// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package cmds

import (
	"context"

	"cloudeng.io/glean/crawlindex/crawl"
)

type Crawl struct {
	options
}

func (cmd *Crawl) Run(ctx context.Context, values interface{}, args []string) error {
	crawler := crawl.Crawler{
		GleanConfig: globalConfig,
		Extractors:  cmd.extractors,
		FSForCrawl:  cmd.fsForCrawl,
	}
	return crawler.Run(ctx, values.(*crawl.Flags), args[0])
}
