// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package config

import "time"

// BulkIndex represents the configuration info for a Glean builk index operation.
type BulkIndex struct {
	ForceDeletion       bool `yaml:"force_deletion"`        // Glean's force deletion flag
	ForceRestart        bool `yaml:"force_restart"`         // Glean's force restart flag
	ReaddirEntries      int  `yaml:"readdir_entries"`       // number of entries per Readdir call.
	DocumentRequestSize int  `yaml:"document_request_size"` // number of documents to include in a single request in a single bulk index request.
	EmployeeRequestSize int  `yaml:"employee_request_size"` // number of employees to include in a single request in a single bulk index request.
}

// IncrementalIndex represents the configuration info for incremental, document
// at a time, indexing.
type IncrementalIndex struct {
	DeletionDelay time.Duration `yaml:"deletion_delay"` // Documents that have not been updated within deletion delay will be removed the Glean index.
}
