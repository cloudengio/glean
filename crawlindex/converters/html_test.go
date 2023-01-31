// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package converters_test

import (
	"testing"
	"time"

	"cloudeng.io/file"
	"cloudeng.io/file/content"
	"cloudeng.io/file/download"
	"cloudeng.io/glean/crawlindex"
	"cloudeng.io/glean/crawlindex/config"
	"cloudeng.io/glean/crawlindex/converters"
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

	ds := converters.NewHTML()
	modTime := time.Now()
	dl := download.Result{
		Name:     "foo",
		Contents: []byte(htmlExample),
		FileInfo: file.NewInfo("foo", 100, 0600, modTime, file.InfoOption{}),
	}
	crawled := crawlindex.Document{
		Time:     modTime,
		Type:     content.Type("text/html;charset=utf-8"),
		Download: &dl,
	}
	dsURL := "https://example.com"
	cfg := config.Conversion{
		Converter: &config.Converter{
			ViewURLRewrites: []string{"s%(.*)%https://example.com/$1%"},
		},
		Datasource: &gleansdk.CustomDatasourceConfig{
			Name:    "test-datasource",
			HomeUrl: &dsURL,
		},
	}
	gd, err := ds.GleanDocument("text/html", cfg, crawled)
	if err != nil {
		t.Fatal(err)
	}

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