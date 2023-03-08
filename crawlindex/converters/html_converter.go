// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package converters

import (
	"bytes"
	"context"
	"fmt"

	"cloudeng.io/file/content"
	"cloudeng.io/file/content/processors"
	"cloudeng.io/file/download"
	"cloudeng.io/glean/crawlindex/config"
	"cloudeng.io/glean/gleansdk"
	"cloudeng.io/text/textutil"
)

// HTML represents an html to glean document converter.
type HTML struct{}

// NewHTML returns a new install of HTML.
func NewHTML() T {
	return &HTML{}
}

var htmlProc processors.HTML

func (cnv *HTML) Type() content.Type {
	return "text/html;charset=utf-8"
}

func (cnv *HTML) Convert(ctx context.Context, datasource string, cfg config.Conversion, ctype content.Type, data []byte) (gleansdk.DocumentDefinition, error) {
	var gd gleansdk.DocumentDefinition
	if content.Clean(ctype) != cfg.Type {
		return gd, fmt.Errorf("htmlConverter: expected %v, not %v", cfg.Type, ctype)
	}
	var obj content.Object[[]byte, download.Result]
	if err := obj.Decode(data); err != nil {
		return gd, fmt.Errorf("htmlConverter: converter: failed to decode object data: %v", err)
	}

	rwr, err := textutil.NewRewriteRules(cfg.Converter.ViewURLRewrites...)
	if err != nil {
		return gd, err
	}

	contents := obj.Value
	gd.Datasource = datasource

	dl := obj.Response

	gd.SetId(dl.Name)
	gd.SetViewURL(rwr.ReplaceAllStringFirst(dl.Name))

	var title string
	if doc, err := htmlProc.Parse(bytes.NewReader(contents)); err == nil {
		title = doc.Title()
	}
	gd.SetTitle(title)

	gd.Body = &gleansdk.ContentDefinition{}
	gd.Body.SetMimeType(string(obj.Type))
	gd.Body.SetTextContent(string(contents))

	gd.Author = &gleansdk.UserReferenceDefinition{}
	gd.Author.SetEmail(cfg.Converter.DefaultAuthor.Email)

	gd.Permissions = &gleansdk.DocumentPermissionsDefinition{}
	gd.Permissions.SetAllowAnonymousAccess(true)

	gd.UpdatedAt = new(int64)
	*gd.UpdatedAt = dl.FileInfo.ModTime().Unix()

	return gd, nil
}
