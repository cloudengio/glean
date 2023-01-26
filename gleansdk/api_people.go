/*
Glean Indexing API

# Introduction In addition to the data sources that Glean has built-in support for, Glean also provides a REST API that enables customers to put arbitrary content in the search index. This is useful, for example, for doing permissions-aware search over content in internal tools that reside on-prem as well as for searching over applications that Glean does not currently support first class. In addition these APIs allow the customer to push organization data (people info, organization structure etc) into Glean.  # Early Access Please note that we are continually evolving the system so these APIs are considered “early access” and are subject to change. 

API version: 0.9.0
Contact: support@glean.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gleansdk

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
)


// PeopleApiService PeopleApi service
type PeopleApiService service

type ApiBulkindexemployeesPostRequest struct {
	ctx context.Context
	ApiService *PeopleApiService
	bulkIndexEmployeesRequest *BulkIndexEmployeesRequest
}

func (r ApiBulkindexemployeesPostRequest) BulkIndexEmployeesRequest(bulkIndexEmployeesRequest BulkIndexEmployeesRequest) ApiBulkindexemployeesPostRequest {
	r.bulkIndexEmployeesRequest = &bulkIndexEmployeesRequest
	return r
}

func (r ApiBulkindexemployeesPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.BulkindexemployeesPostExecute(r)
}

/*
BulkindexemployeesPost Bulk index employees

Replaces all the currently indexed employees using paginated batch API calls.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiBulkindexemployeesPostRequest
*/
func (a *PeopleApiService) BulkindexemployeesPost(ctx context.Context) ApiBulkindexemployeesPostRequest {
	return ApiBulkindexemployeesPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *PeopleApiService) BulkindexemployeesPostExecute(r ApiBulkindexemployeesPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PeopleApiService.BulkindexemployeesPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/bulkindexemployees"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.bulkIndexEmployeesRequest == nil {
		return nil, reportError("bulkIndexEmployeesRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json; charset=UTF-8"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.bulkIndexEmployeesRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiBulkindexteamsPostRequest struct {
	ctx context.Context
	ApiService *PeopleApiService
	bulkIndexTeamsRequest *BulkIndexTeamsRequest
}

func (r ApiBulkindexteamsPostRequest) BulkIndexTeamsRequest(bulkIndexTeamsRequest BulkIndexTeamsRequest) ApiBulkindexteamsPostRequest {
	r.bulkIndexTeamsRequest = &bulkIndexTeamsRequest
	return r
}

func (r ApiBulkindexteamsPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.BulkindexteamsPostExecute(r)
}

/*
BulkindexteamsPost Bulk index teams

Replaces all the currently indexed teams using paginated batch API calls.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiBulkindexteamsPostRequest
*/
func (a *PeopleApiService) BulkindexteamsPost(ctx context.Context) ApiBulkindexteamsPostRequest {
	return ApiBulkindexteamsPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *PeopleApiService) BulkindexteamsPostExecute(r ApiBulkindexteamsPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PeopleApiService.BulkindexteamsPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/bulkindexteams"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.bulkIndexTeamsRequest == nil {
		return nil, reportError("bulkIndexTeamsRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json; charset=UTF-8"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.bulkIndexTeamsRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiDeleteemployeePostRequest struct {
	ctx context.Context
	ApiService *PeopleApiService
	deleteEmployeeRequest *DeleteEmployeeRequest
}

func (r ApiDeleteemployeePostRequest) DeleteEmployeeRequest(deleteEmployeeRequest DeleteEmployeeRequest) ApiDeleteemployeePostRequest {
	r.deleteEmployeeRequest = &deleteEmployeeRequest
	return r
}

func (r ApiDeleteemployeePostRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteemployeePostExecute(r)
}

/*
DeleteemployeePost Delete employee

Delete an employee. Silently succeeds if employee is not present.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDeleteemployeePostRequest
*/
func (a *PeopleApiService) DeleteemployeePost(ctx context.Context) ApiDeleteemployeePostRequest {
	return ApiDeleteemployeePostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *PeopleApiService) DeleteemployeePostExecute(r ApiDeleteemployeePostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PeopleApiService.DeleteemployeePost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/deleteemployee"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.deleteEmployeeRequest == nil {
		return nil, reportError("deleteEmployeeRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json; charset=UTF-8"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.deleteEmployeeRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiIndexemployeePostRequest struct {
	ctx context.Context
	ApiService *PeopleApiService
	indexEmployeeRequest *IndexEmployeeRequest
}

func (r ApiIndexemployeePostRequest) IndexEmployeeRequest(indexEmployeeRequest IndexEmployeeRequest) ApiIndexemployeePostRequest {
	r.indexEmployeeRequest = &indexEmployeeRequest
	return r
}

func (r ApiIndexemployeePostRequest) Execute() (*http.Response, error) {
	return r.ApiService.IndexemployeePostExecute(r)
}

/*
IndexemployeePost Index employee

Adds an employee or updates information about an employee

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiIndexemployeePostRequest
*/
func (a *PeopleApiService) IndexemployeePost(ctx context.Context) ApiIndexemployeePostRequest {
	return ApiIndexemployeePostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *PeopleApiService) IndexemployeePostExecute(r ApiIndexemployeePostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PeopleApiService.IndexemployeePost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/indexemployee"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.indexEmployeeRequest == nil {
		return nil, reportError("indexEmployeeRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json; charset=UTF-8"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.indexEmployeeRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiIndexemployeelistPostRequest struct {
	ctx context.Context
	ApiService *PeopleApiService
	indexEmployeeListRequest *IndexEmployeeListRequest
}

func (r ApiIndexemployeelistPostRequest) IndexEmployeeListRequest(indexEmployeeListRequest IndexEmployeeListRequest) ApiIndexemployeelistPostRequest {
	r.indexEmployeeListRequest = &indexEmployeeListRequest
	return r
}

func (r ApiIndexemployeelistPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.IndexemployeelistPostExecute(r)
}

/*
IndexemployeelistPost Bulk index employees

Bulk upload details of all the employees. This deletes all employees uploaded in the prior batch. SOON TO BE DEPRECATED in favor of /bulkindexemployees.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiIndexemployeelistPostRequest

Deprecated
*/
func (a *PeopleApiService) IndexemployeelistPost(ctx context.Context) ApiIndexemployeelistPostRequest {
	return ApiIndexemployeelistPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
// Deprecated
func (a *PeopleApiService) IndexemployeelistPostExecute(r ApiIndexemployeelistPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PeopleApiService.IndexemployeelistPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/indexemployeelist"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.indexEmployeeListRequest == nil {
		return nil, reportError("indexEmployeeListRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json; charset=UTF-8"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.indexEmployeeListRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiIndexteamPostRequest struct {
	ctx context.Context
	ApiService *PeopleApiService
	indexTeamRequest *IndexTeamRequest
}

func (r ApiIndexteamPostRequest) IndexTeamRequest(indexTeamRequest IndexTeamRequest) ApiIndexteamPostRequest {
	r.indexTeamRequest = &indexTeamRequest
	return r
}

func (r ApiIndexteamPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.IndexteamPostExecute(r)
}

/*
IndexteamPost Index team

Adds a team or updates information about a team

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiIndexteamPostRequest
*/
func (a *PeopleApiService) IndexteamPost(ctx context.Context) ApiIndexteamPostRequest {
	return ApiIndexteamPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *PeopleApiService) IndexteamPostExecute(r ApiIndexteamPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PeopleApiService.IndexteamPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/indexteam"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.indexTeamRequest == nil {
		return nil, reportError("indexTeamRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json; charset=UTF-8"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.indexTeamRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
