// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package index

import "cloudeng.io/glean/gleansdk"

// Request contains the entities to be indexed and a flag indicating if
// these are the last entities in a bulk index operation. The indexer
// will stop when it receives a request with LastPage set. If LastPage
// flag is never set, the indexer will assume that the indexing operation
// is complete when its input channel is closed.
type Request struct {
	Documents []gleansdk.DocumentDefinition
	Users     []gleansdk.DatasourceUserDefinition
	LastPage  bool
}
