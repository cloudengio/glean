// Copyright 2023 cloudeng llc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package index

import (
	"context"
	"fmt"
	"log"
	"time"

	"cloudeng.io/algo/container/circular"
	"cloudeng.io/glean/gleansdk"
)

func itemsToIndex[T any](buf *circular.Buffer[T], atleast, atmost int) []T {
	if buf.Len() >= atleast {
		return buf.Head(atmost)
	}
	return nil
}

type bulkUserIndexer struct {
	bulkOptions
	client      *gleansdk.APIClient
	datasource  string
	id          string
	users       *circular.Buffer[gleansdk.DatasourceUserDefinition]
	indexed     int
	invocations int
	firstPage   bool
}

func newBulkUserIndexer(opts bulkOptions, client *gleansdk.APIClient, datasource string) *bulkUserIndexer {
	return &bulkUserIndexer{
		bulkOptions: opts,
		client:      client,
		id:          time.Now().Round(0).String(),
		firstPage:   true,
		datasource:  datasource,
		users:       circular.NewBuffer[gleansdk.DatasourceUserDefinition](opts.empReqSize),
	}
}

func (idx *bulkUserIndexer) handleUsers(ctx context.Context, users []gleansdk.DatasourceUserDefinition) error {
	if len(users) == 0 {
		return nil
	}
	idx.users.Append(users)
	return idx.handleNextRequests(ctx, idx.empReqSize, idx.empReqSize)
}

func (idx *bulkUserIndexer) handleNextRequests(ctx context.Context, atleast, atmost int) error {
	for {
		done, err := idx.handleNextRequest(ctx, atleast, atmost)
		if err != nil {
			return err
		}
		if done {
			return nil
		}
	}
}

func (idx *bulkUserIndexer) handleNextRequest(ctx context.Context, atleast, atmost int) (bool, error) {
	var req gleansdk.BulkIndexUsersRequest
	users := itemsToIndex(idx.users, atleast, atmost)
	if len(users) == 0 {
		return true, nil
	}
	req.SetDatasource(idx.datasource)
	req.SetIsFirstPage(idx.firstPage)
	if idx.firstPage {
		req.SetForceRestartUpload(idx.forceRestart)
	}
	req.SetIsLastPage(false)
	req.SetUploadId(idx.id)
	req.Users = users

	log.Printf("user: request with %v\n", len(req.Users))
	reqStart := time.Now()
	if idx.dryRun {
		log.Printf("user: would index %v users, %v invocations\n", len(req.Users), idx.invocations+1)
	} else {
		resp, err := idx.client.PermissionsApi.BulkindexusersPost(ctx).BulkIndexUsersRequest(req).Execute()
		if err := handleHTTPError(resp, err); err != nil {
			return true, fmt.Errorf("user: request: %v", err)
		}
	}

	idx.firstPage = false
	took := time.Since(reqStart)
	idx.indexed += len(req.Users)
	idx.invocations++

	log.Printf("user: indexed: total # docs: % 5v, per req # docs: % 3v in % 8v\n", idx.indexed, len(req.Users), took)

	return false, nil
}

func (idx *bulkUserIndexer) finish(ctx context.Context) error {
	if err := idx.handleNextRequests(ctx, 0, idx.empReqSize); err != nil {
		return err
	}
	if idx.dryRun {
		log.Printf("user: last page: sent %v users over %v invocations\n", idx.indexed, idx.invocations)
		return nil
	}
	bulkReq := gleansdk.BulkIndexUsersRequest{}
	bulkReq.SetDatasource(idx.datasource)
	bulkReq.SetIsFirstPage(false)
	bulkReq.SetIsLastPage(true)
	bulkReq.SetUploadId(idx.id)
	bulkReq.Users = []gleansdk.DatasourceUserDefinition{}
	resp, err := idx.client.PermissionsApi.BulkindexusersPost(ctx).BulkIndexUsersRequest(bulkReq).Execute()
	if err := handleHTTPError(resp, err); err != nil {
		return fmt.Errorf("user: last page request: %v", err)
	}
	return nil
}
