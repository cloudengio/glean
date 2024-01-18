// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package config_test

import (
	"testing"

	"cloudeng.io/cmdutil/cmdyaml"
	"cloudeng.io/glean/crawlindex/config"
)

const crawlsSpec = `
  name: test
  depth: 3
  seeds:
    - s3://foo/bar

  download:
    default_concurrency: 4 # 0 will default to all available CPUs
    default_request_chan_size: 100
    default_downloads_chan_size: 100
    per_depth_concurrency: [1, 2, 4]

  aws:
    aws: true
    aws_region: us-east-1
`

func TestCrawlConfig(t *testing.T) {
	var crawl config.Crawl
	if err := cmdyaml.ParseConfigString(crawlsSpec, &crawl); err != nil {
		t.Fatal(err)
	}

	if got, want := crawl.Name, "test"; got != want {
		t.Errorf("got %v, want %v", got, want)
	}

	if got, want := len(crawl.Seeds), 1; got != want {
		t.Errorf("got %v, want %v", got, want)
	}

	if got, want := crawl.AWS.AWS, true; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
	if got, want := crawl.AWS.AWSRegion, "us-east-1"; got != want {
		t.Errorf("got %v, want %v", got, want)
	}

}
