// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

// Package builtin provides the set of builtin resources for a given
// instance of the Glean CLI. These resources fall into two categories:
//  1. static, these can be instantiated once at init time.
//  2. dynamic, these are instantiate dynamically, usually by other
//     packages when the application is running. They are always returned
//     as 'factories' that can be used to create the actual resources.
package builtin
