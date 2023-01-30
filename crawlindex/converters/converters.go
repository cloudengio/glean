// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package converters

import (
	"cloudeng.io/file/content"
	"cloudeng.io/glean/crawlindex"
	"cloudeng.io/glean/crawlindex/config"
	"cloudeng.io/glean/gleansdk"
)

const GleanContentType = content.Type("glean/document")

type T interface {
	Type() content.Type
	GleanDocument(datasource string, cfg config.Conversion, dl crawlindex.Document) (*gleansdk.DocumentDefinition, error)
}

func CreateRegistry(converters ...T) (*content.Registry[T], error) {
	reg := content.NewRegistry[T]()
	for _, cnv := range converters {
		from := cnv.Type()
		if err := reg.RegisterConverters(from, GleanContentType, cnv); err != nil {
			return nil, err
		}
	}
	return reg, nil
}
