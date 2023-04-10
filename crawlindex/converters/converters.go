// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package converters

import (
	"context"

	"cloudeng.io/file/content"
	"cloudeng.io/glean/crawlindex/config"
	"cloudeng.io/glean/gleansdk"
)

const (
	GleanDocumentType = content.Type("glean/document")
	GleanUserType     = content.Type("glean/user")
)

type ignoreError struct {
	ctype content.Type
}

func (e ignoreError) Error() string {
	return "ignore content type: " + string(e.ctype)
}

func (e ignoreError) Is(target error) bool {
	_, ok := target.(ignoreError)
	return ok
}

func IgnoreContentType(ctype content.Type) error {
	return ignoreError{ctype: ctype}
}

func IsIgnoreContentType(err error) bool {
	_, ok := err.(ignoreError)
	return ok
}

type Document interface {
	Type() content.Type
	Convert(ctx context.Context, datasource string, cfg config.Conversion, ctype content.Type, data []byte) (gleansdk.DocumentDefinition, error)
}

func CreateDocumentRegistry(converters ...Document) (*content.Registry[Document], error) {
	reg := content.NewRegistry[Document]()
	for _, cnv := range converters {
		from := cnv.Type()
		if err := reg.RegisterConverters(from, GleanDocumentType, cnv); err != nil {
			return nil, err
		}
	}
	return reg, nil
}

type User interface {
	Type() content.Type
	Convert(ctx context.Context, datasource string, cfg config.Conversion, ctype content.Type, data []byte) (gleansdk.DatasourceUserDefinition, error)
}

func CreateUserRegistry(converters ...User) (*content.Registry[User], error) {
	reg := content.NewRegistry[User]()
	for _, cnv := range converters {
		from := cnv.Type()
		if err := reg.RegisterConverters(from, GleanUserType, cnv); err != nil {
			return nil, err
		}
	}
	return reg, nil
}
