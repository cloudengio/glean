// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"cloudeng.io/glean/gleanclientsdk"
	"cloudeng.io/glean/gleansdk"
	"cloudeng.io/webapi/operations/apitokens"
)

// ParseGleanError parses the error returned by the glean API and returns a
// more useful error message.
func ParseGleanError(r *http.Response, err error) error {
	if err == nil {
		return nil
	}
	oapiErr, ok := err.(*gleansdk.GenericOpenAPIError)
	if !ok {
		if r != nil {
			return fmt.Errorf("%v: %v: %v", r.Request.URL, r.StatusCode, err)
		}
		return err
	}
	if r == nil {
		return err
	}
	var tmp any
	body := oapiErr.Body()
	if json.Unmarshal(body, &tmp) == nil {
		return fmt.Errorf("%s: %v", body, err)
	}
	if body, nerr := io.ReadAll(r.Body); nerr == nil {
		return fmt.Errorf("%s: %v", body, err)
	}
	return err
}

func NewIndexingClient(ctx context.Context, domain, tokenName string) (context.Context, *gleansdk.APIClient, error) {
	templateVars := map[string]string{
		"domain": domain,
	}
	ctx = context.WithValue(ctx, gleansdk.ContextServerVariables, templateVars)
	token, ok := apitokens.TokenFromContext(ctx, tokenName)
	if !ok {
		return ctx, nil, fmt.Errorf("no token found for name %q", tokenName)
	}
	ctx = context.WithValue(ctx, gleansdk.ContextAccessToken, string(token.Value()))
	return ctx, gleansdk.NewAPIClient(gleansdk.NewConfiguration()), nil
}

func NewClient(ctx context.Context, domain, tokenName string) (context.Context, *gleanclientsdk.APIClient, error) {
	templateVars := map[string]string{
		"domain": domain,
	}
	ctx = context.WithValue(ctx, gleansdk.ContextServerVariables, templateVars)
	token, ok := apitokens.TokenFromContext(ctx, tokenName)
	if !ok {
		return ctx, nil, fmt.Errorf("no token found for name %q", tokenName)
	}
	ctx = context.WithValue(ctx, gleansdk.ContextAccessToken, string(token.Value()))
	return ctx, gleanclientsdk.NewAPIClient(gleanclientsdk.NewConfiguration()), nil
}
