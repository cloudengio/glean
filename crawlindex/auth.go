// Copyright 2024 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package crawlindex

import (
	"cloudeng.io/webapi/operations/apitokens"
)

type AuthFileFlag struct {
	AuthFile string `subcmd:"glean-auth,$HOME/.glean-auth.yaml,'file containing authentication tokens for glean instances'"`
}

type Auth []struct {
	Name string `yaml:"name" cmd:"name of the glean token instance"`
	Auth struct {
		BearerToken       string `yaml:"indexing_token" cmd:"indexing token for the glean instance"`
		ClientBearerToken string `yaml:"client_token" cmd:"client bearer token for the glean instance"`
	}
}

func (a Auth) TokensForName(name, domain string) (indexingToken, clientToken *apitokens.T) {
	if len(name) == 0 {
		name = domain
	}
	for _, cfg := range a {
		if cfg.Name == name {
			return apitokens.New(cfg.Auth.BearerToken), apitokens.New(cfg.Auth.ClientBearerToken)
		}
	}
	return nil, nil
}
