// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package config_test

import (
	"testing"

	"cloudeng.io/cmdutil/cmdyaml"
	"cloudeng.io/file/content"
	"cloudeng.io/glean/crawlindex/config"
)

const convertersSpec = `
- from_content_types:
  - "text/html; charset=utf-8"
  default_author:
    email: "user@example.com"
  view_url_rewrites:
    - "s%a%b%"
  custom:
    foo: "bar"
    bar: "baz"
 `

func TestConverterCfg(t *testing.T) {
	var converters config.Converters
	if err := cmdyaml.ParseConfigString(convertersSpec, &converters); err != nil {
		t.Fatal(err)
	}
	if got, want := len(converters), 1; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	cnv0 := converters[0]
	if got, want := len(cnv0.FromContentType), 1; got != want {
		t.Errorf("got %v, want %v", got, want)
	}

	if got, want := cnv0.FromContentType[0], content.Type("text/html; charset=utf-8"); got != want {
		t.Errorf("got %v, want %v", got, want)
	}

	if got, want := cnv0.DefaultAuthor.Email, "user@example.com"; got != want {
		t.Errorf("got %v, want %v", got, want)
	}

	custom := struct {
		Foo string `yaml:"foo"`
		Bar string `yaml:"bar"`
	}{}

	if err := cnv0.CustomConfig.Decode(&custom); err != nil {
		t.Fatal(err)
	}

	if got, want := custom.Foo, "bar"; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	if got, want := custom.Bar, "baz"; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
