// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package index

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func httpBody(resp *http.Response) string {
	if resp == nil {
		return ""
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	return string(body)
}

func handleHTTPError(resp *http.Response, err error) error {
	if err != nil {
		body := httpBody(resp)
		log.Printf("err: %v, HTTP response: %v ", err, body)
		return fmt.Errorf("err: %v, HTTP response: %v", err, body)
	}
	if resp.StatusCode != http.StatusOK {
		body := httpBody(resp)
		log.Printf("unexpected HTTP status code: %v, HTTP response: %v", resp.Status, body)
		return fmt.Errorf("unexpected HTTP status code: %v, HTTP response: %v", resp.Status, body)
	}
	return nil
}
