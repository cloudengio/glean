// Copyright 2022 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package main

import (
	"context"

	"cloudeng.io/glean/gleancli/cmds"
)

func main() {
	cmds.MustNew().MustDispatch(context.Background())
}
