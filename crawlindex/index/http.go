// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package index

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func httpBody(resp *http.Response) string {
	if resp == nil {
		return ""
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}

func handleHTTPError(resp *http.Response, err error) error {
	if err != nil {
		return fmt.Errorf("err: %v, HTTP response: %v", err, httpBody(resp))
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected HTTP status code: %v", resp.Status)
	}
	return nil
}
