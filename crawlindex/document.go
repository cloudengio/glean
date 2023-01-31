// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package crawlindex

import (
	"time"

	"cloudeng.io/file/content"
	"cloudeng.io/file/download"
)

// Document represents a downloaded document, whether it be a file, a http get
// or the result of a Rest API call.
type Document struct {
	Time     time.Time
	Type     content.Type
	Download *download.Result
	// Fill in for http downloads, API Calls etc.
}

func (d Document) WriteToFile(dir, file string) error {
	return WriteDocument(d, dir, file)
}

func NewDownloadDocument(d *download.Result, ctype content.Type) Document {
	return Document{
		Time:     time.Now(),
		Type:     ctype,
		Download: d,
	}
}
