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
	"cloudeng.io/file/content"
	"cloudeng.io/file/crawl/crawlcmd"
	"cloudeng.io/glean/crawlindex/config"
	"cloudeng.io/glean/gleancli/extensions"
	"cloudeng.io/path/cloudpath"
	"cloudeng.io/webapi/operations"
)

// PopulateCrawlFS fills a map of URI schemes to crawlcmd.FS factories
func PopulateCrawlFS(_ context.Context, cfg config.CrawlService, factories map[string]crawlcmd.FSFactory) error {
	switch cfg.Name {
	case "s3", "aws":
		factories["s3"] = s3FS(cfg.Config)
		return nil
	case "unix", "file":
		factories["unix"] = localFS
		return nil
	}
	return fmt.Errorf("unsupported service  %q", cfg.Name)
}

// NewContentFS creates a content.FS appropriate for the given crawl downloads
// cache configuration.
func NewContentFS(ctx context.Context, cfg crawlcmd.CrawlCacheConfig) (content.FS, error) {
	dp, _ := cfg.Paths()
	match := cloudpath.DefaultMatchers.Match(dp)
	switch match.Scheme {
	case "s3":
		return newS3FS(ctx, cfg.ServiceConfig)
	case "unix":
		return newLocalFS(ctx)
	}
	return nil, fmt.Errorf("unsupported cache downloads path: %v", dp)
}

// NewCheckpointOp creates a checkpoint.Operation appropriate for the given
// crawl checkpoint configuration.
func NewCheckpointOp(ctx context.Context, cfg crawlcmd.CrawlCacheConfig) (checkpoint.Operation, error) {
	_, cp := cfg.Paths()
	match := cloudpath.DefaultMatchers.Match(cp)
	switch match.Scheme {
	case "s3":
		return newS3Checkpoint(ctx, cfg.ServiceConfig)
	case "unix":
		return checkpoint.NewDirectoryOperation(), nil
	}
	return nil, fmt.Errorf("unsupported cache checkpoint path: %v", cp)
}

// NewOperationsFS returns an operations.FS appropriate for the given path and
// configuration.
func NewOperationsFS(ctx context.Context, cfg crawlcmd.CrawlCacheConfig) (operations.FS, error) {
	dp, _ := cfg.Paths()
	match := cloudpath.DefaultMatchers.Match(dp)
	switch match.Scheme {
	case "s3":
		return newS3FS(ctx, cfg.ServiceConfig)
	case "unix":
		return newLocalFS(ctx)
	}
	return nil, fmt.Errorf("unsupported cache path: %v", dp)
}

// New returns a set of dynamic resources for use with the Glean CLI.
func New() extensions.DynamicResources {
	return extensions.DynamicResources{
		PopulateCrawlFS: PopulateCrawlFS,
		NewContentFS:    NewContentFS,
		NewCheckpointOp: NewCheckpointOp,
		NewOperationsFS: NewOperationsFS,
	}
}
