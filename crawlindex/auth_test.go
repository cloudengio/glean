// Copyright 2024 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package crawlindex_test

/*
func TestAuthCfg(t *testing.T) {
	var cfg crawlindex.Auth
	if err := cmdyaml.ParseConfig([]byte(`
- name: a
  auth:
    indexing_token_name: ia
    client_token_name: ca
- name: b
  auth:
    indexing_token_name: ib
    client_token_name: cb
`), &cfg); err != nil {
		t.Fatal(err)
	}

	i1, c1 := cfg.TokensForName("a", "a")
	i2, c2 := cfg.TokensForName("", "a")
	if got, want := i1, "ia"; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	if got, want := c1, "ca"; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	if got, want := i1, i2; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	if got, want := c1, c2; got != want {
		t.Errorf("got %v, want %v", got, want)
	}

	i3, c3 := cfg.TokensForName("", "b")
	if got, want := i3, "ib"; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	if got, want := c3, "cb"; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
*/
