// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package config_test

import (
	"testing"

	"cloudeng.io/cmdutil"
	"cloudeng.io/glean/crawlindex/config"
)

const apiCrawlSpec = `
api1:
  something: 1
api2:
  else: 2
`

type api1 struct {
	Something int `yaml:"something"`
}

type api2 struct {
	Else int `yaml:"else"`
}

func TestAPICrawlConfig(t *testing.T) {
	var crawls config.APICrawls
	if err := cmdutil.ParseYAMLConfigString(apiCrawlSpec, &crawls); err != nil {
		t.Fatal(err)
	}

	if got, want := len(crawls), 2; got != want {
		t.Errorf("got %v, want %v", got, want)
	}

	var a1 api1
	if present, err := config.ParseCrawlConfig(crawls, "api1", &a1); !present || err != nil {
		t.Fatalf("not present: %v, or err: %v", present, err)
	}
	if got, want := a1.Something, 1; got != want {
		t.Errorf("got %v, want %v", got, want)
	}

	var a2 api2
	if present, err := config.ParseCrawlConfig(crawls, "api2", &a2); !present || err != nil {
		t.Fatalf("not present: %v, or err: %v", present, err)
	}
	if got, want := a2.Else, 2; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
