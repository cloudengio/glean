// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package crawlindex

import (
	"bytes"
	"encoding/gob"
	"os"
	"path/filepath"
)

func marshal(v any) ([]byte, error) {
	buf := &bytes.Buffer{}
	enc := gob.NewEncoder(buf)
	if err := enc.Encode(v); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func unmarshal(data []byte, v any) error {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	err := dec.Decode(v)
	return err
}

// WriteDocument writes the provided document to the specified file.
func WriteDocument(v Document, dir, file string) error {
	if err := os.MkdirAll(dir, 0700); err != nil {
		return err
	}
	data, err := marshal(v)
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(dir, file), data, 0600)
}

// ReadDocument reads a Document from the specified file.
func ReadDocument(dir, file string) (Document, error) {
	data, err := os.ReadFile(filepath.Join(dir, file))
	if err != nil {
		return Document{}, err
	}
	var v Document
	if err := unmarshal(data, &v); err != nil {
		return Document{}, err
	}
	return v, nil
}
