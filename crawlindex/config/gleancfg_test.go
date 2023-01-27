// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package config_test

import (
	"testing"

	"cloudeng.io/cmdutil"
	"github.com/cloudengio/glean/crawlindex/config"
)

const configSpec = `
- name: glean-dev
  auth:
    token: "token"
  api:
    domain: glean-dev
- name: another-instance
  auth:
    token: "another token"
  api:
    domain: another-instance.com
`

func TestConfig(t *testing.T) {
	var cfg config.Glean
	if err := cmdutil.ParseYAMLConfigString(configSpec, &cfg); err != nil {
		t.Fatal(err)
	}
	if got, want := len(cfg), 2; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	if got, want := cfg[1].Name, "another-instance"; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
