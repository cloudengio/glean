// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

// Package dynamic provides dynamically instantiated resources for a given
// instance of the Glean CLI. These resources are generally returned as 'factories'.
package dynamic

import (
	"context"
	"fmt"

	"cloudeng.io/file/checkpoint"
	"cloudeng.io/file/crawl/crawlcmd"
	"cloudeng.io/path/cloudpath"
	"cloudeng.io/webapi/operations"
	"gopkg.in/yaml.v3"
)

// PopulateFS fills a map of URI schemes to crawlcmd.FS factories
func PopulateFS(name string, cfg yaml.Node, factories map[string]crawlcmd.FSFactory) error {
	switch name {
	case "s3", "aws":
		factories["s3"] = s3FS(cfg)
		return nil
	case "unix":
		factories["unix"] = localFS
		return nil
	}
	return fmt.Errorf("unsupported file system: %q", name)
}

// StoreFS returns a factory to creates an operations.FS for the
// given crawl/indexing cache configuration.
func StoreFS(ctx context.Context, path string, cfg yaml.Node) (operations.FS, error) {
	match := cloudpath.DefaultMatchers.Match(path)
	if len(match.Matched) == 0 {
		return nil, fmt.Errorf("unsupported cache path: %v", path)
	}
	switch match.Scheme {
	case "s3":
		return newS3FS(ctx, cfg)
	case "unix":
		return newLocalFS(ctx)
	}
	return nil, fmt.Errorf("unsupported cache path: %v", path)
}

func CheckpointOp(ctx context.Context, path string, cfg yaml.Node) (checkpoint.Operation, error) {
	match := cloudpath.DefaultMatchers.Match(path)
	if len(match.Matched) == 0 {
		return nil, fmt.Errorf("unsupported cache path: %v", path)
	}
	switch match.Scheme {
	case "s3":
		return newS3Checkpoint(ctx, cfg)
	case "unix":
		return checkpoint.NewDirectoryOperation(), nil
	}
	return nil, fmt.Errorf("unsupported cache path: %v", path)
}
