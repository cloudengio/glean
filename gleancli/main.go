// Copyright 2022 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package main

import (
	"context"

	"cloudeng.io/glean/gleancli/builtin"
	"cloudeng.io/glean/gleancli/cmds"
)

func main() {
	cmds.MustNew(
		cmds.WithDocumentConverters(builtin.MustDocumentConverters()),
		cmds.WithUserConverters(builtin.MustUserConverters()),
		cmds.WithFSForCrawl(builtin.FSForCrawl),
		cmds.WithOutlinkExtractors(builtin.Extractors),
		cmds.WithAPIExtensions(builtin.APIExtensions("api")),
	).MustDispatch(context.Background())
}
