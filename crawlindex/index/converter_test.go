// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package index_test

import (
	"testing"
	"time"

	"cloudeng.io/file/download"
	"cloudeng.io/file/filetestutil"
	"cloudeng.io/glean/crawlindex/index"
	"cloudeng.io/glean/gleansdk"
)

const htmlExample = `
<!DOCTYPE html>
<html>
<body>

<h1>My First Heading</h1>
<title>My title</title>
<p><link href="https://sample.css">CSS</a></p>
<a hread=xx>
</body>
</html>

`

func TestConverter(t *testing.T) {
	dsURL := "https://example.com"
	gds := &gleansdk.CustomDatasourceConfig{
		Name:    "test-datasource",
		HomeUrl: &dsURL,
	}
	ds := index.NewConverter(gds)
	modTime := time.Now()
	dl := download.Result{
		Name:     "foo",
		Contents: []byte(htmlExample),
		FileInfo: filetestutil.NewInfo("foo", 100, 0600, modTime, false, nil),
	}
	gd := ds.GleanDocument("text/html", dl)

	if got, want := gd.GetTitle(), "My title"; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	if got, want := gd.GetUpdatedAt(), modTime.Unix(); got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	if got, want := gd.GetViewURL(), "https://example.com/foo"; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
