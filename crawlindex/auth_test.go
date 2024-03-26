// Copyright 2024 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package crawlindex_test

import (
	"testing"

	"cloudeng.io/cmdutil/cmdyaml"
	"cloudeng.io/glean/crawlindex"
)

func TestAuthCfg(t *testing.T) {
	var cfg crawlindex.Auth
	if err := cmdyaml.ParseConfig([]byte(`
- name: a
  auth:
    indexing_token: ia
    client_token: ca
- name: b
  auth:
    indexing_token: ib
    client_token: cb
`), &cfg); err != nil {
		t.Fatal(err)
	}

	i1, c1 := cfg.TokensForName("a", "a")
	i2, c2 := cfg.TokensForName("", "a")
	if got, want := i1.Path, "ia"; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	if got, want := c1.Path, "ca"; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	if got, want := i1.Path, i2.Path; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	if got, want := c1.Path, c2.Path; got != want {
		t.Errorf("got %v, want %v", got, want)
	}

	i3, c3 := cfg.TokensForName("", "b")
	if got, want := i3.Path, "ib"; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	if got, want := c3.Path, "cb"; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
