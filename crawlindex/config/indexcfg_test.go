// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package config_test

import (
	"reflect"
	"testing"
	"time"

	"cloudeng.io/cmdutil"
	"github.com/cloudengio/glean/crawlindex/config"
)

const bulkIndexSpec = `
  glean_instance: glean-dev
  force_restart: true
  force_deletion: true
  readdir_entries: 72
`

func TestBulkIndexConfig(t *testing.T) {
	var index config.BulkIndex
	if err := cmdutil.ParseYAMLConfigString(bulkIndexSpec, &index); err != nil {
		t.Fatal(err)
	}

	if got, want := index.ForceDeletion, true; got != want {
		t.Errorf("got %v, want %v", got, want)
	}

	if got, want := index.ForceRestart, true; got != want {
		t.Errorf("got %v, want %v", got, want)
	}

	if got, want := index.ReaddirEntries, 72; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

const incrementalIndexSpec = `
  deletion_delay: 48h
`

func TestIncrementalIndexConfig(t *testing.T) {
	var index config.IncrementalIndex
	if err := cmdutil.ParseYAMLConfigString(incrementalIndexSpec, &index); err != nil {
		t.Fatal(err)
	}
	if got, want := index.DeletionDelay, 48*time.Hour; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

const bulkOnly = `
bulk_index:
` + bulkIndexSpec

const incrementalOnly = `
incremental_index:
` + incrementalIndexSpec

type bothIndices struct {
	*config.BulkIndex        `yaml:"bulk_index,omitempty"`
	*config.IncrementalIndex `yaml:"incremental_index,omitempty"`
}

func TestIndexModeSelection(t *testing.T) {
	var index bothIndices
	if err := cmdutil.ParseYAMLConfigString(bulkOnly, &index); err != nil {
		t.Fatal(err)
	}

	actualBulk := &config.BulkIndex{
		ForceDeletion:  true,
		ForceRestart:   true,
		ReaddirEntries: 72,
	}
	actualIncremental := &config.IncrementalIndex{DeletionDelay: 48 * time.Hour}

	if got, want := index.BulkIndex, actualBulk; !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}

	if got, want := index.IncrementalIndex, (*config.IncrementalIndex)(nil); !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}

	index = bothIndices{}
	if err := cmdutil.ParseYAMLConfigString(incrementalOnly, &index); err != nil {
		t.Fatal(err)
	}
	if got, want := index.BulkIndex, (*config.BulkIndex)(nil); !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}

	if got, want := index.IncrementalIndex, actualIncremental; !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
