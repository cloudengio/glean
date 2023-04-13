// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package papersapp

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"cloudeng.io/file/content"
	"cloudeng.io/glean/crawlindex/config"
	"cloudeng.io/glean/crawlindex/converters"
	"cloudeng.io/glean/gleansdk"
	"cloudeng.io/webapi/operations"
	"cloudeng.io/webapi/papersapp"
	"cloudeng.io/webapi/papersapp/papersappsdk"
)

type docConverter struct {
}

func (c *docConverter) Type() content.Type {
	return papersapp.ItemType
}

func (c *docConverter) Convert(_ context.Context, datasourceName string, cfg config.Conversion, ctype content.Type, data []byte) (gleansdk.DocumentDefinition, error) {
	var gd gleansdk.DocumentDefinition
	switch ctype {
	case papersapp.ItemType:
	case papersapp.CollectionType:
		return gd, converters.IgnoreContentType(ctype)
	default:
		return gd, fmt.Errorf("app.readcube.com item converter: expected %v, not %v", cfg.Type, ctype)
	}
	var obj content.Object[papersapp.Item, operations.Response]
	if err := obj.Decode(data); err != nil {
		return gd, fmt.Errorf("app.readcube.com item converter: failed to decode object data: %v", err)
	}
	doc := obj.Value
	gd.SetDatasource(datasourceName)
	gd.SetId(doc.Item.ID)

	viewURL := fmt.Sprintf("%s/library/%v/all/(sidepanel:details)?item_id=%v&collection_id=%v",
		*cfg.Datasource.HomeUrl, doc.Collection.ID, doc.Item.ID, doc.Collection.ID)
	gd.SetViewURL(viewURL)
	gd.SetContainer(doc.Collection.Name)

	gd.Permissions = &gleansdk.DocumentPermissionsDefinition{}
	gd.Permissions.SetAllowAnonymousAccess(true)

	switch doc.Item.ItemType {
	case "article":
		c.article(cfg, &gd, doc.Item)
	default:
		// full list is in papersapp.ItemType which can be obtained from:
		// curl -s https://api.papers.ai/schema | jq .definitions.ItemType
		log.Printf("unsupported papersapp item type: %v", doc.Item.ItemType)
	}
	return gd, nil
}

func (c *docConverter) article(cfg config.Conversion, gd *gleansdk.DocumentDefinition, item *papersappsdk.Item) {

	article := item.Article

	gd.SetTitle(article.Title)

	summary := &strings.Builder{}
	for _, author := range article.Authors {
		summary.WriteString(author)
		summary.WriteRune('\n')
	}
	summary.WriteString(article.Journal)
	summary.WriteRune('\n')
	str := summary.String()
	gd.SetSummary(gleansdk.ContentDefinition{
		MimeType:    "text/plain",
		TextContent: &str,
	})

	gd.SetBody(gleansdk.ContentDefinition{
		MimeType:    "text/plain",
		TextContent: &article.Abstract,
	})

	year := time.Date(article.Year, time.January, 1, 0, 0, 0, 0, time.UTC)
	gd.SetCreatedAt(year.Unix())
	gd.Author = &gleansdk.UserReferenceDefinition{}
	if len(article.Authors) > 0 {
		gd.Author.SetName(article.Authors[0])
		gd.Author.SetEmail(cfg.Converter.DefaultAuthor.Email)
	} else {
		gd.Author.SetName(cfg.Converter.DefaultAuthor.Name)
		gd.Author.SetEmail(cfg.Converter.DefaultAuthor.Email)
	}
	gd.UpdatedBy = gd.Author
}

func NewDocumentConverter() converters.Document {
	return &docConverter{}
}
