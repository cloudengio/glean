// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package config

import (
	"cloudeng.io/file/content"
	"gopkg.in/yaml.v3"
)

// User represents a user in the system/datasource beinfg indexed.
type User struct {
	Email  string `yaml:"email"`
	UserID string `yaml:"user_id"`
	Name   string `yaml:"name"`
}

type Converter struct {
	FromContentType []content.Type `yaml:"from_content_types"` // Content types that this converter can handle.

	ViewURLRewrites []string `yaml:"view_url_rewrites"` // Rewrite rules for viewurls specified as textutil.RewriteRules

	AllowAnonymousAccess bool `yaml:"allow_anonymous_access"` // allow anonymous access to the converted documents.

	// Default author to use if none can be obtained from the document itself.
	DefaultAuthor User `yaml:"default_author"`

	CustomConfig yaml.Node `yaml:"custom"`
}

type Converters []Converter
