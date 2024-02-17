// Copyright 2022 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package main

import (
	"context"

	"cloudeng.io/glean/gleancli/builtin/dynamic"
	"cloudeng.io/glean/gleancli/builtin/static"
	"cloudeng.io/glean/gleancli/cmds"
)

func main() {
	opts := cmds.Options{
		CrawlProcessors:    static.MustCrawlProcessors(),
		IndexProcessors:    static.MustIndexProcessors(),
		CreateCrawlFS:      dynamic.PopulateFS,
		CreateStoreFS:      dynamic.StoreFS,
		CreateCheckpointOp: dynamic.CheckpointOp,
		APIExtensions:      static.APIExtensions("api"),
	}
	cmds.MustNew(opts).MustDispatch(context.Background())
}
