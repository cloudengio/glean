// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package config

import (
	"cloudeng.io/aws/awsconfig"
	"cloudeng.io/file/crawl/crawlcmd"
)

// Crawl represents a single crawl that contributes data to a datasource.
type Crawl struct {
	crawlcmd.Config `yaml:",inline"`
	AWS             awsconfig.AWSFlags `yaml:"aws,omitempty"`
}
