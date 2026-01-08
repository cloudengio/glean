// Copyright 2024 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package crawlindex

/*
type AuthFileFlag struct {
	AuthFile string `subcmd:"glean-key-info,,'file or keychain item containing authentication tokens for glean instances'"`
}

func (f AuthFileFlag) KeyInfo() string {
	return f.AuthFile
}

type Auth []struct {
	Name string `yaml:"name" cmd:"name of the glean token instance"`
	Auth struct {
		IndexingTokenName string `yaml:"indexing_token_name" cmd:"name of the indexing token for the glean instance"`
		ClientTokenName   string `yaml:"client_token_name" cmd:"name of the client token for the glean instance"`
	}
}

func (a Auth) TokensForName(name, domain string) (indexingTokenOrName, clientTokenOrName string) {
	if len(name) == 0 {
		name = domain
	}
	for _, cfg := range a {
		if cfg.Name == name {
			return cfg.Auth.IndexingTokenName, cfg.Auth.ClientTokenName
		}
	}
	return "", ""
}
*/
