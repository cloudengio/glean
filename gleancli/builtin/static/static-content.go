// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package static

import (
	"context"

	"cloudeng.io/file/content"
	"cloudeng.io/glean/crawlindex/config"
	"cloudeng.io/glean/crawlindex/converters"
	"cloudeng.io/glean/gleansdk"
)

// HTML represents an html to glean document converter. It differs
// from converters.HTML only in that it sets the object type to statichtml.
type HTML struct {
	cnv converters.Document
}

// NewHTML returns a new install of HTML.
func NewHTML() converters.Document {
	return &HTML{cnv: converters.NewHTML()}
}

func (cnv *HTML) Type() content.Type {
	return cnv.cnv.Type()
}

func (cnv *HTML) Convert(ctx context.Context, datasource string, cfg config.Conversion, ctype content.Type, data []byte) (gleansdk.DocumentDefinition, error) {
	gd, err := cnv.cnv.Convert(ctx, datasource, cfg, ctype, data)
	if err != nil {
		return gleansdk.DocumentDefinition{}, err
	}
	gd.SetObjectType("statichtml")
	return gd, nil
}
