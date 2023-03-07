// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package config

import (
	"gopkg.in/yaml.v3"
)

type APICrawls map[string]yaml.Node

// ParseAPICrawlConfig parses an API specific crawl
// config with the specified name.
func ParseAPICrawlConfig[T any](crawls APICrawls, name string, crawlConfig *T) (bool, error) {
	cfg, ok := crawls[name]
	if !ok {
		return false, nil
	}
	if err := cfg.Decode(crawlConfig); err != nil {
		return false, err
	}
	return true, nil
}
