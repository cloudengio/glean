// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package config

import (
	"cloudeng.io/glean/gleansdk"
)

// DatasourceConfig represents the configuration of the datasource with
// Glean's API.
type DatasourceConfig struct {
	GleanInstance                   string `yaml:"glean_instance"`
	gleansdk.CustomDatasourceConfig `yaml:",inline"`
}
