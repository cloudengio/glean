// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package config_test

import (
	"testing"

	"cloudeng.io/cmdutil"
	"cloudeng.io/glean/crawlindex/config"
)

const datasourcesSpec = `- datasource: static-corpus
  crawls:
    - name: crawl1
      depth: 3
      nofollow:
        - "^https://.*"
      seeds:
        - s3://foo/bar
        - https://yahoo.com
      download:
        default_concurrency: 4 # 0 will default to all available CPUs
        default_request_chan_size: 100
        default_downloads_chan_size: 100
        concurrency: [1, 2, 4]
  cache:
    path: $HOME/.cache/test-crawl
  index:
    force_restart: true
    force_deletion: true
  datasource_config:
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
	var ds []config.Datasource
	if err := cmdutil.ParseYAMLConfigString(datasourcesSpec, &ds); err != nil {
		t.Fatal(err)
	}

	if got, want := len(ds), 1; got != want {
		t.Fatalf("got %v, want %v", got, want)
	}

	ds0 := ds[0]
	if got, want := ds0.Datasource, "static-corpus"; got != want {
		t.Errorf("got %v, want %v", got, want)
	}

	if got, want := ds0.Cache.Path, "$HOME/.cache/test-crawl"; got != want {
		t.Errorf("got %v, want %v", got, want)
	}

	if got, want := len(ds0.Crawls), 1; got != want {
		t.Fatalf("got %v, want %v", got, want)
	}

	crawls0 := ds0.Crawls[0]
	if got, want := crawls0.Download.DefaultConcurrency, 4; got != want {
		t.Errorf("got %v, want %v", got, want)
	}

	if got, want := crawls0.Depth, 3; got != want {
		t.Errorf("got %v, want %v", got, want)
	}

	if got, want := ds0.CustomDatasourceConfig.GetHomeUrl(), "static.example.com"; got != want {
		t.Errorf("got %v, want %v", got, want)
	}

	if got, want := ds0.CustomDatasourceConfig.GetDatasourceCategory(), "PUBLISHED_CONTENT"; got != want {
		t.Errorf("got %v, want %v", got, want)
	}

}
