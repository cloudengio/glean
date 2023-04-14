// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package index

import (
	"context"

	"cloudeng.io/glean/crawlindex/config"
)

// QueryFlags represents the flags to the indexing query command.
type QueryFlags struct {
	config.FileFlags
}

func (idx *Indexer) Query(ctx context.Context, fv *StatsFlags, datasource string, query []string) error {

	return nil
}
