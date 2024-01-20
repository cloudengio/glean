// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package config

import (
	"cloudeng.io/file/crawl/crawlcmd"
	"gopkg.in/yaml.v3"
)

// Crawl represents a single crawl that contributes data to a datasource.
type Crawl struct {
	crawlcmd.Config `yaml:",inline" cmd:"crawl configuration"`
	ServiceName     string    `yaml:"service_name" cmd:"name of service to crawl, eg. s3"`
	ServiceConfig   yaml.Node `yaml:"service_config" cmd:"service specific configuration"`
}
