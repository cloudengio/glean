// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package index

import (
	"context"
	"fmt"

	"cloudeng.io/glean/crawlindex/config"
	"cloudeng.io/glean/gleansdk"
)

// StatsFlags represents the flags to the bulk indexing command.
type StatsFlags struct {
	config.FileFlags
}

func (idx *Indexer) Stats(ctx context.Context, fv *StatsFlags) error {
	ctx, client := idx.newGleanIndexingClient(ctx)

	var req gleansdk.GetDocumentCountRequest
	req.SetName(idx.datasourceName)

	count, resp, err := client.DocumentsApi.GetdocumentcountPost(ctx).GetDocumentCountRequest(req).Execute()
	if err := handleHTTPError(resp, err); err != nil {
		return err
	}

	fmt.Printf("Datasource: %v\n", req.GetName())
	fmt.Printf("\t# documents: %d\n", count.GetDocumentCount())

	return nil
}
