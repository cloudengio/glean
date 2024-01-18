// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package config_test

import (
	"testing"

	"cloudeng.io/cmdutil/cmdyaml"
	"cloudeng.io/glean/config"
)

const datasourcesSpec = `
glean_instance: glean-dev
name: internal-documentation"
displayname: display-internal-documentation"
homeurl: static.example.com
datasourcecategory: PUBLISHED_CONTENT
urlregex: "^https://static.example.com/.*"
iconurl: "icon.jpg"
icondarkurl: "dark.jpg"
`

func TestDataSource(t *testing.T) {
	var ds config.DatasourceConfig
	if err := cmdyaml.ParseConfigString(datasourcesSpec, &ds); err != nil {
		t.Fatal(err)
	}

	if got, want := ds.GetHomeUrl(), "static.example.com"; got != want {
		t.Errorf("got %v, want %v", got, want)
	}

	if got, want := ds.GetDatasourceCategory(), "PUBLISHED_CONTENT"; got != want {
		t.Errorf("got %v, want %v", got, want)
	}

}
