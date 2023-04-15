// Copyright 2022 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package config

import (
	"context"
	"fmt"
	"strings"

	"cloudeng.io/glean/gleanclientsdk"
	"cloudeng.io/glean/gleansdk"
)

type GleanFlags struct {
	Config string `subcmd:"config,$HOME/.glean.yaml,'glean config file'"`
}

type Glean []struct {
	Name string `yaml:"name"`
	Auth struct {
		BearerToken       string `yaml:"indexing_token"`
		ClientBearerToken string `yaml:"client_token"`
	}
	API struct {
		Domain string `yaml:"domain"`
	}
}

func (c Glean) String() string {
	var out strings.Builder
	for _, cfg := range c {
		fmt.Fprintf(&out, "name: %s\n  auth:\n", cfg.Name)
		if len(cfg.Auth.BearerToken) > 0 {
			fmt.Fprintf(&out, "  token: **redacted**\n")
		}
		fmt.Fprintf(&out, "api\n  url: %s\n", cfg.API.Domain)
	}
	return out.String()
}

func (c Glean) NewIndexingAPIClient(ctx context.Context, name string) (context.Context, *gleansdk.APIClient, error) {
	for _, cfg := range c {
		if cfg.Name == name {
			templateVars := map[string]string{
				"domain": cfg.API.Domain,
			}
			ctx = context.WithValue(ctx, gleansdk.ContextAccessToken, cfg.Auth.BearerToken)
			ctx = context.WithValue(ctx, gleansdk.ContextServerVariables, templateVars)
			return ctx, gleansdk.NewAPIClient(gleansdk.NewConfiguration()), nil
		}
	}
	return ctx, nil, fmt.Errorf("Glean.NewIndexingAPIClient: failed to find config for %q", name)
}

func (c Glean) NewClientAPIClient(ctx context.Context, name string) (context.Context, *gleanclientsdk.APIClient, error) {
	for _, cfg := range c {
		if cfg.Name == name {
			templateVars := map[string]string{
				"domain": cfg.API.Domain,
			}
			ctx = context.WithValue(ctx, gleanclientsdk.ContextAccessToken, cfg.Auth.ClientBearerToken)
			ctx = context.WithValue(ctx, gleanclientsdk.ContextServerVariables, templateVars)
			return ctx, gleanclientsdk.NewAPIClient(gleanclientsdk.NewConfiguration()), nil
		}
	}
	return ctx, nil, fmt.Errorf("Glean.NewClientAPIClient: failed to find config for %q", name)
}
