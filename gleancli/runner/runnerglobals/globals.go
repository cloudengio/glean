// Copyright 2024 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package runnerglobals

import (
	"strings"

	"cloudeng.io/cmdutil/cmdexec"
)

type Flags struct {
	WorkingDir           string `subcmd:"working-dir,.,'working directory'"`
	DatasourceConfigFile string `subcmd:"datasource-configs,connectors.yaml,'datasource configuration file'"`
	Executable           string `subcmd:"executable,gleancli,'comma separated list of commands to run for gleancli, eg. go run .'"`
	DryRun               bool   `subcmd:"dry-run,false,'do not actually run the commands'"`
	Verbose              bool   `subcmd:"verbose,false,'print verbose output'"`
}

func (f *Flags) CmdExecOptions() []cmdexec.Option {
	return []cmdexec.Option{
		cmdexec.WithCommandsPrefix(strings.Split(f.Executable, ",")...),
		cmdexec.WithWorkingDir(f.WorkingDir),
		cmdexec.WithDryRun(f.DryRun),
		cmdexec.WithVerbose(f.Verbose),
	}
}
