// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package crawlindex

/*
type objectWrapper struct {
	Type  content.Type
	Value any
}

func WriteObject(path string, ctype content.Type, object any) error {
	ow := objectWrapper{
		Type:  ctype,
		Value: object,
	}
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0700); err != nil {
		return err
	}
	wr, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer wr.Close()
	enc := gob.NewEncoder(wr)
	return enc.Encode(ow)
}

func ReadObject(path string) (ctype content.Type, object any, err error) {
	rd, err := os.Open(path)
	if err != nil {
		return "", nil, err
	}
	defer rd.Close()
	enc := gob.NewDecoder(rd)
	ow := objectWrapper{}
	err = enc.Decode(&ow)
	return ow.Type, ow.Value, err
}
*/
