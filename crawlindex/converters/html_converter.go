// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package converters

import (
	"bytes"

	"cloudeng.io/file/content"
	"cloudeng.io/file/content/processors"
	"cloudeng.io/glean/crawlindex"
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

func (cnv *HTML) GleanDocument(datasource string, cfg config.Conversion, doc crawlindex.Document) (*gleansdk.DocumentDefinition, error) {
	rwr, err := textutil.NewRewriteRules(cfg.Converter.ViewURLRewrites...)
	if err != nil {
		return nil, err
	}

	gd := &gleansdk.DocumentDefinition{}
	gd.Datasource = datasource

	dl := doc.Download
	gd.SetId(dl.Name)
	gd.SetViewURL(rwr.ReplaceAllStringFirst(dl.Name))

	var title string
	if doc, err := htmlProc.Parse(bytes.NewReader(dl.Contents)); err == nil {
		title = doc.Title()
	}
	gd.SetTitle(title)

	gd.Body = &gleansdk.ContentDefinition{}
	gd.Body.SetMimeType(string(doc.Type))
	gd.Body.SetTextContent(string(dl.Contents))

	gd.Author = &gleansdk.UserReferenceDefinition{}
	gd.Author.SetEmail(cfg.Converter.DefaultAuthor.Email)

	gd.Permissions = &gleansdk.DocumentPermissionsDefinition{}
	gd.Permissions.SetAllowAnonymousAccess(true)

	gd.UpdatedAt = new(int64)
	*gd.UpdatedAt = dl.FileInfo.ModTime().Unix()

	return gd, nil
}
