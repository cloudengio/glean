/*
Glean Client API

# Introduction These are the public APIs to enable implementing a custom client interface to the Glean system.  # Usage guidelines This API is evolving fast. Glean will provide advance notice of any planned backwards incompatible changes along with a 6-month sunset period for anything that requires developers to adopt the new versions. 

API version: 0.9.0
Contact: support@glean.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gleanclientsdk

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// AnnouncementsApiService AnnouncementsApi service
type AnnouncementsApiService service

type ApiCreateannouncementRequest struct {
	ctx context.Context
	ApiService *AnnouncementsApiService
	payload *CreateAnnouncementRequest
	actas *string
	clientVersion *string
	domain *string
	eids *[]int64
}

// Announcement content
func (r ApiCreateannouncementRequest) Payload(payload CreateAnnouncementRequest) ApiCreateannouncementRequest {
	r.payload = &payload
	return r
}

// Email of another user to act as for debugging purposes. Requires sufficient permissions.
func (r ApiCreateannouncementRequest) Actas(actas string) ApiCreateannouncementRequest {
	r.actas = &actas
	return r
}

// The version of the client making the request.
func (r ApiCreateannouncementRequest) ClientVersion(clientVersion string) ApiCreateannouncementRequest {
	r.clientVersion = &clientVersion
	return r
}

// The domain of the top-level page in which the client is being run. For embedded search and NSR, it will be the domain of the embedding page. For sidebar, it will be the domain of the embedding page if it&#39;s a recognized workplace app. For NTP and app.glean.com requests, it will be empty.
func (r ApiCreateannouncementRequest) Domain(domain string) ApiCreateannouncementRequest {
	r.domain = &domain
	return r
}

// List of experiment ids to force for incoming request.
func (r ApiCreateannouncementRequest) Eids(eids []int64) ApiCreateannouncementRequest {
	r.eids = &eids
	return r
}

func (r ApiCreateannouncementRequest) Execute() (*Announcement, *http.Response, error) {
	return r.ApiService.CreateannouncementExecute(r)
}

/*
Createannouncement Create announcement

Creates a textual announcement visible to some set of users based on department and location.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateannouncementRequest
*/
func (a *AnnouncementsApiService) Createannouncement(ctx context.Context) ApiCreateannouncementRequest {
	return ApiCreateannouncementRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Announcement
func (a *AnnouncementsApiService) CreateannouncementExecute(r ApiCreateannouncementRequest) (*Announcement, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Announcement
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AnnouncementsApiService.Createannouncement")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/createannouncement"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.payload == nil {
		return localVarReturnValue, nil, reportError("payload is required and must be specified")
	}

	if r.actas != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "actas", r.actas, "")
	}
	if r.clientVersion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "clientVersion", r.clientVersion, "")
	}
	if r.domain != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "domain", r.domain, "")
	}
	if r.eids != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "eids", r.eids, "csv")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.payload
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDeleteannouncementRequest struct {
	ctx context.Context
	ApiService *AnnouncementsApiService
	payload *DeleteAnnouncementRequest
	actas *string
	clientVersion *string
	domain *string
	eids *[]int64
}

// Delete announcement request
func (r ApiDeleteannouncementRequest) Payload(payload DeleteAnnouncementRequest) ApiDeleteannouncementRequest {
	r.payload = &payload
	return r
}

// Email of another user to act as for debugging purposes. Requires sufficient permissions.
func (r ApiDeleteannouncementRequest) Actas(actas string) ApiDeleteannouncementRequest {
	r.actas = &actas
	return r
}

// The version of the client making the request.
func (r ApiDeleteannouncementRequest) ClientVersion(clientVersion string) ApiDeleteannouncementRequest {
	r.clientVersion = &clientVersion
	return r
}

// The domain of the top-level page in which the client is being run. For embedded search and NSR, it will be the domain of the embedding page. For sidebar, it will be the domain of the embedding page if it&#39;s a recognized workplace app. For NTP and app.glean.com requests, it will be empty.
func (r ApiDeleteannouncementRequest) Domain(domain string) ApiDeleteannouncementRequest {
	r.domain = &domain
	return r
}

// List of experiment ids to force for incoming request.
func (r ApiDeleteannouncementRequest) Eids(eids []int64) ApiDeleteannouncementRequest {
	r.eids = &eids
	return r
}

func (r ApiDeleteannouncementRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteannouncementExecute(r)
}

/*
Deleteannouncement Delete announcement

Deletes an existing user-generated announcement.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDeleteannouncementRequest
*/
func (a *AnnouncementsApiService) Deleteannouncement(ctx context.Context) ApiDeleteannouncementRequest {
	return ApiDeleteannouncementRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *AnnouncementsApiService) DeleteannouncementExecute(r ApiDeleteannouncementRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AnnouncementsApiService.Deleteannouncement")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/deleteannouncement"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.payload == nil {
		return nil, reportError("payload is required and must be specified")
	}

	if r.actas != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "actas", r.actas, "")
	}
	if r.clientVersion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "clientVersion", r.clientVersion, "")
	}
	if r.domain != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "domain", r.domain, "")
	}
	if r.eids != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "eids", r.eids, "csv")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	localVarPostBody = r.payload
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiUpdateannouncementRequest struct {
	ctx context.Context
	ApiService *AnnouncementsApiService
	payload *UpdateAnnouncementRequest
	actas *string
	clientVersion *string
	domain *string
	eids *[]int64
}

// Announcement content. Id need to be specified for the announcement.
func (r ApiUpdateannouncementRequest) Payload(payload UpdateAnnouncementRequest) ApiUpdateannouncementRequest {
	r.payload = &payload
	return r
}

// Email of another user to act as for debugging purposes. Requires sufficient permissions.
func (r ApiUpdateannouncementRequest) Actas(actas string) ApiUpdateannouncementRequest {
	r.actas = &actas
	return r
}

// The version of the client making the request.
func (r ApiUpdateannouncementRequest) ClientVersion(clientVersion string) ApiUpdateannouncementRequest {
	r.clientVersion = &clientVersion
	return r
}

// The domain of the top-level page in which the client is being run. For embedded search and NSR, it will be the domain of the embedding page. For sidebar, it will be the domain of the embedding page if it&#39;s a recognized workplace app. For NTP and app.glean.com requests, it will be empty.
func (r ApiUpdateannouncementRequest) Domain(domain string) ApiUpdateannouncementRequest {
	r.domain = &domain
	return r
}

// List of experiment ids to force for incoming request.
func (r ApiUpdateannouncementRequest) Eids(eids []int64) ApiUpdateannouncementRequest {
	r.eids = &eids
	return r
}

func (r ApiUpdateannouncementRequest) Execute() (*Announcement, *http.Response, error) {
	return r.ApiService.UpdateannouncementExecute(r)
}

/*
Updateannouncement Update announcement

Updates a textual announcement visible to some set of users based on department and location.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUpdateannouncementRequest
*/
func (a *AnnouncementsApiService) Updateannouncement(ctx context.Context) ApiUpdateannouncementRequest {
	return ApiUpdateannouncementRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Announcement
func (a *AnnouncementsApiService) UpdateannouncementExecute(r ApiUpdateannouncementRequest) (*Announcement, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Announcement
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AnnouncementsApiService.Updateannouncement")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/updateannouncement"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.payload == nil {
		return localVarReturnValue, nil, reportError("payload is required and must be specified")
	}

	if r.actas != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "actas", r.actas, "")
	}
	if r.clientVersion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "clientVersion", r.clientVersion, "")
	}
	if r.domain != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "domain", r.domain, "")
	}
	if r.eids != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "eids", r.eids, "csv")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.payload
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
