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
		StaticResources:  static.New(),
		Extensions:       static.Extensions("test"),
		APIExtensions:    static.APIExtensions("api"),
		DynamicResources: dynamic.New(),
	}
	cmds.MustNew(opts).MustDispatch(context.Background())
}
