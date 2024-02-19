// Copyright 2024 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package dynamic

import (
	"context"

	"cloudeng.io/aws/awsconfig"
	"cloudeng.io/aws/s3fs"
	"cloudeng.io/file"
	"cloudeng.io/file/checkpoint"
	"cloudeng.io/file/localfs"
	"gopkg.in/yaml.v3"
)

func newS3FS(ctx context.Context, cfg yaml.Node) (*s3fs.T, error) {
	var awscfg awsconfig.AWSFlags
	if err := cfg.Decode(&awscfg); err != nil {
		return nil, err
	}
	s3f := &s3fs.Factory{Config: awscfg}
	return s3f.New(ctx)
}

func s3FS(cfg yaml.Node) func(ctx context.Context) (file.FS, error) {
	return func(ctx context.Context) (file.FS, error) {
		return newS3FS(ctx, cfg)
	}
}

func newLocalFS(ctx context.Context) (*localfs.T, error) {
	return localfs.New(), nil
}

func localFS(ctx context.Context) (file.FS, error) {
	return localfs.New(), nil
}

func newS3Checkpoint(ctx context.Context, cfg yaml.Node) (checkpoint.Operation, error) {
	fs, err := newS3FS(ctx, cfg)
	if err != nil {
		return nil, err
	}
	return s3fs.NewCheckpointOperation(fs), nil
}
