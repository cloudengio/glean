// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package converters_test

import (
	"context"
	"testing"

	"cloudeng.io/file/content"
	"cloudeng.io/glean/crawlindex/config"
	"cloudeng.io/glean/crawlindex/converters"
	"cloudeng.io/glean/gleansdk"
)

type cnv struct{}

func (c *cnv) Convert(_ context.Context, _ string, _ config.Conversion, _ content.Type, _ []byte) (gleansdk.DocumentDefinition, error) {
	return gleansdk.DocumentDefinition{}, converters.IgnoreContentType("foo")
}

func TestIgnore(t *testing.T) {
	c := &cnv{}
	_, err := c.Convert(context.Background(), "", config.Conversion{}, content.Type("foo"), nil)
	if !converters.IsIgnoreContentType(err) {
		t.Errorf("unexpected error: %v", err)
	}
}
