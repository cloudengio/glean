// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package benchling

import (
	"context"
	"fmt"
	"strings"
	"time"

	"cloudeng.io/file/content"
	"cloudeng.io/glean/crawlindex/config"
	"cloudeng.io/glean/crawlindex/converters"
	"cloudeng.io/glean/gleansdk"
	"cloudeng.io/webapi/benchling"
	"cloudeng.io/webapi/benchling/benchlingsdk"
	"cloudeng.io/webapi/operations"
)

type docConverter struct{}

func (c *docConverter) Type() content.Type {
	return benchling.DocumentType
}

func (c *docConverter) Convert(_ context.Context, datasourceName string, cfg config.Conversion, ctype content.Type, data []byte) (gleansdk.DocumentDefinition, error) {
	var gd gleansdk.DocumentDefinition
	switch ctype {
	case benchling.DocumentType:
	case benchling.UserType, benchling.EntryType, benchling.FolderType:
		return gd, converters.IgnoreContentType(ctype)
	default:
		return gd, fmt.Errorf("benchling.com entry converter: expected %v, not %v", cfg.Type, ctype)
	}
	var obj content.Object[benchling.Document, struct{}]
	if err := obj.Decode(data); err != nil {
		return gd, fmt.Errorf("benchling.com entry converter: failed to decode object data: %v", err)
	}
	doc := obj.Value

	gd.SetDatasource(datasourceName)
	gd.SetId(*doc.Entry.Id)
	gd.SetObjectType(string(benchling.DocumentType))
	gd.SetViewURL(*doc.Entry.WebURL)
	gd.SetTitle(*doc.Entry.Name)
	gd.SetContainer(strings.Join(doc.Parents, "/"))
	gd.SetCreatedAt(doc.Entry.CreatedAt.Unix())
	if when, err := time.Parse(time.RFC3339, *doc.Entry.ModifiedAt); err == nil {
		gd.SetUpdatedAt(when.Unix())
	}
	gd.Author = &gleansdk.UserReferenceDefinition{}
	if doc.Entry.Authors != nil && len(*doc.Entry.Authors) > 0 {
		usr := doc.Users[*(*doc.Entry.Authors)[0].Id]
		gd.Author.SetName(*usr.Name)
		gd.Author.SetEmail(*usr.Email) // need email from
		gd.Author.SetDatasourceUserId(*usr.Id)
	} else {
		gd.Author.SetName(cfg.Converter.DefaultAuthor.Name)
		gd.Author.SetEmail(cfg.Converter.DefaultAuthor.Email)
		gd.Author.SetDatasourceUserId(cfg.Converter.DefaultAuthor.UserID)
	}
	gd.UpdatedBy = gd.Author

	tmp := doc.DayNotes
	gd.SetBody(gleansdk.ContentDefinition{
		MimeType:    "text/plain",
		TextContent: &tmp,
	})

	// Allow benchling users to view benchling results.
	gd.Permissions = &gleansdk.DocumentPermissionsDefinition{}
	gd.Permissions.SetAllowAllDatasourceUsersAccess(true)
	gd.Permissions.AllowedUsers = make([]gleansdk.UserReferenceDefinition, 0, len(doc.Users))
	return gd, nil
}

type userConverter struct{}

func (c *userConverter) Type() content.Type {
	return benchling.UserType
}

func (c *userConverter) Convert(_ context.Context, datasourceName string, cfg config.Conversion, ctype content.Type, data []byte) (gleansdk.DatasourceUserDefinition, error) {
	var gd gleansdk.DatasourceUserDefinition
	switch ctype {
	case benchling.UserType:
	case benchling.DocumentType, benchling.EntryType, benchling.FolderType:
		return gd, converters.IgnoreContentType(ctype)
	default:
		return gd, fmt.Errorf("benchling.com user content converter: expected %v, not %v", cfg.Type, ctype)
	}

	var obj content.Object[benchlingsdk.User, *operations.Response]
	if err := obj.Decode(data); err != nil {
		return gd, fmt.Errorf("benchling.com user converter: failed to decode object data: %v", err)
	}
	usr := obj.Value
	if usr.Name == nil {
		return gd, fmt.Errorf("benchling.com user converter: user name is nil")
	}
	gd.SetName(*usr.Name)
	gd.SetEmail(*usr.Email)
	active := !(usr.IsSuspended != nil && *usr.IsSuspended)
	gd.SetIsActive(active)
	gd.SetUserId(*usr.Id)
	return gd, nil
}

func NewDocumentConverter() converters.Document {
	return &docConverter{}
}

func NewUserConverter() converters.User {
	return &userConverter{}
}
