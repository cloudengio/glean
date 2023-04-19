// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package config_test

import (
	"testing"

	"cloudeng.io/cmdutil"
	"cloudeng.io/glean/config"
)

const configSpec = `
- name: glean-dev
  auth:
    token: "token"
  api:
    domain: glean-dev
- name: another-instance
  auth:
    indexing_token: "another token"
    client_token: "client token"
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
	if got, want := cfg[1].Auth.BearerToken, "another token"; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	if got, want := cfg[1].Auth.ClientBearerToken, "client token"; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
