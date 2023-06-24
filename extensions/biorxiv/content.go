// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package biorxiv

import (
	"context"
	"fmt"
	"strings"
	"time"

	"cloudeng.io/file/content"
	"cloudeng.io/glean/crawlindex/config"
	"cloudeng.io/glean/crawlindex/converters"
	"cloudeng.io/glean/gleansdk"
	"cloudeng.io/webapi/biorxiv"
)

func NewDocumentConverter() converters.Document {
	return &docConverter{}
}

type docConverter struct{}

func (c *docConverter) Type() content.Type {
	return biorxiv.DocumentType
}

func (c *docConverter) Convert(_ context.Context, datasourceName string, cfg config.Conversion, ctype content.Type, data []byte) (gleansdk.DocumentDefinition, error) {
	var gd gleansdk.DocumentDefinition
	if ctype != biorxiv.DocumentType {
		return gd, fmt.Errorf("api.biorxiv.org article converter: expected %v, not %v", cfg.Type, ctype)
	}
	var obj content.Object[biorxiv.PreprintDetail, struct{}]
	if err := obj.Decode(data); err != nil {
		return gd, fmt.Errorf("api.biorxiv.org article converter: failed to decode object data: %v", err)
	}

	doc := obj.Value
	gd.SetDatasource(datasourceName)
	gd.SetId(doc.PreprintDOI)

	host := "https://www.biorxiv.org"
	if strings.ToLower(doc.PreprintPlatform) == "medrxiv" {
		host = "https://www.medrxiv.org"
	}
	gd.SetViewURL(fmt.Sprintf("%s/content/%s", host, doc.PreprintDOI))
	gd.SetContainer(doc.PreprintPlatform)

	gd.SetTitle(doc.PreprintTitle)
	gd.SetObjectType("preprint")

	preprintDate, err := time.Parse(time.DateOnly, doc.PreprintDate)
	if err == nil {
		gd.SetCreatedAt(preprintDate.Unix())
	}

	summary := &strings.Builder{}
	summary.WriteString(doc.PreprintAuthors)
	summary.WriteString(": ")
	if len(doc.PublishedJournal) > 0 {
		summary.WriteString(doc.PublishedJournal)
	}
	str := summary.String()
	gd.SetSummary(gleansdk.ContentDefinition{
		MimeType:    "text/plain",
		TextContent: &str,
	})

	gd.SetBody(gleansdk.ContentDefinition{
		MimeType:    "text/plain",
		TextContent: &doc.PreprintAbstract,
	})

	gd.Author.SetName(doc.PreprintAuthors)
	gd.Author.SetEmail(cfg.Converter.DefaultAuthor.Email)

	gd.UpdatedBy = gd.Author

	gd.Permissions = &gleansdk.DocumentPermissionsDefinition{}
	gd.Permissions.SetAllowAnonymousAccess(true)

	return gd, nil
}
