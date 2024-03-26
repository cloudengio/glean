// Copyright 2024 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package main

import (
	"context"

	"cloudeng.io/glean/gleancli/runner/runnercmd"
	"cloudeng.io/glean/gleancli/runner/runnerdata"
	"cloudeng.io/glean/gleancli/runner/runnerglobals"
)

var GlobalFlagValues runnerglobals.Flags

func main() {
	ctx := context.Background()
	cmdspecs := &runnercmd.T{
		Datasources:        runnerdata.Datasources,
		CrawlCommands:      runnerdata.CrawlCommands,
		ProcessCommands:    runnerdata.ProcessCommands,
		IndexCommands:      runnerdata.IndexCommands,
		IndexStatsCommands: runnerdata.IndexStatsCommands,
		TestCacheCommands:  runnerdata.TestCacheCommands,
		AuthFiles:          runnerdata.AuthFiles,
	}
	yamlSpec, cmds := cmdspecs.Spec()
	cmdset := runnercmd.NewCmdSet(yamlSpec, cmds, &GlobalFlagValues)
	cmdset.WithMain(func(ctx context.Context, cmdRunner func(ctx context.Context) error) error {
		cmdspecs.GlobalExecOpts = GlobalFlagValues.CmdExecOptions()
		cmdspecs.DatasourceConfigFile = GlobalFlagValues.DatasourceConfigFile
		return cmdRunner(ctx)
	})
	cmdset.MustDispatch(ctx)
}
