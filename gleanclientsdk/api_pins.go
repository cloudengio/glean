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


// PinsApiService PinsApi service
type PinsApiService service

type ApiEditpinRequest struct {
	ctx context.Context
	ApiService *PinsApiService
	payload *EditPinRequest
	clientVersion *string
	domain *string
	eids *[]int64
}

// Edit pins request
func (r ApiEditpinRequest) Payload(payload EditPinRequest) ApiEditpinRequest {
	r.payload = &payload
	return r
}

// The version of the client making the request.
func (r ApiEditpinRequest) ClientVersion(clientVersion string) ApiEditpinRequest {
	r.clientVersion = &clientVersion
	return r
}

// The domain of the top-level page in which the client is being run. For embedded search and NSR, it will be the domain of the embedding page. For sidebar, it will be the domain of the embedding page if it&#39;s a recognized workplace app. For NTP and app.glean.com requests, it will be empty.
func (r ApiEditpinRequest) Domain(domain string) ApiEditpinRequest {
	r.domain = &domain
	return r
}

// List of experiment ids to force for incoming request.
func (r ApiEditpinRequest) Eids(eids []int64) ApiEditpinRequest {
	r.eids = &eids
	return r
}

func (r ApiEditpinRequest) Execute() (*PinDocument, *http.Response, error) {
	return r.ApiService.EditpinExecute(r)
}

/*
Editpin Edit a pin

Edit an existing user-generated pin.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiEditpinRequest
*/
func (a *PinsApiService) Editpin(ctx context.Context) ApiEditpinRequest {
	return ApiEditpinRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PinDocument
func (a *PinsApiService) EditpinExecute(r ApiEditpinRequest) (*PinDocument, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PinDocument
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PinsApiService.Editpin")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/editpin"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.payload == nil {
		return localVarReturnValue, nil, reportError("payload is required and must be specified")
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

type ApiGetpinRequest struct {
	ctx context.Context
	ApiService *PinsApiService
	payload *GetPinRequest
	actas *string
	clientVersion *string
	domain *string
	eids *[]int64
}

// Get pin request
func (r ApiGetpinRequest) Payload(payload GetPinRequest) ApiGetpinRequest {
	r.payload = &payload
	return r
}

// Email of another user to act as for debugging purposes. Requires sufficient permissions.
func (r ApiGetpinRequest) Actas(actas string) ApiGetpinRequest {
	r.actas = &actas
	return r
}

// The version of the client making the request.
func (r ApiGetpinRequest) ClientVersion(clientVersion string) ApiGetpinRequest {
	r.clientVersion = &clientVersion
	return r
}

// The domain of the top-level page in which the client is being run. For embedded search and NSR, it will be the domain of the embedding page. For sidebar, it will be the domain of the embedding page if it&#39;s a recognized workplace app. For NTP and app.glean.com requests, it will be empty.
func (r ApiGetpinRequest) Domain(domain string) ApiGetpinRequest {
	r.domain = &domain
	return r
}

// List of experiment ids to force for incoming request.
func (r ApiGetpinRequest) Eids(eids []int64) ApiGetpinRequest {
	r.eids = &eids
	return r
}

func (r ApiGetpinRequest) Execute() (*GetPinResponse, *http.Response, error) {
	return r.ApiService.GetpinExecute(r)
}

/*
Getpin Read pin details.

Gets a pin given its ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetpinRequest
*/
func (a *PinsApiService) Getpin(ctx context.Context) ApiGetpinRequest {
	return ApiGetpinRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetPinResponse
func (a *PinsApiService) GetpinExecute(r ApiGetpinRequest) (*GetPinResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetPinResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PinsApiService.Getpin")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/getpin"

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

type ApiListpinsRequest struct {
	ctx context.Context
	ApiService *PinsApiService
	payload *map[string]interface{}
	actas *string
	clientVersion *string
	domain *string
	eids *[]int64
}

// List pins request
func (r ApiListpinsRequest) Payload(payload map[string]interface{}) ApiListpinsRequest {
	r.payload = &payload
	return r
}

// Email of another user to act as for debugging purposes. Requires sufficient permissions.
func (r ApiListpinsRequest) Actas(actas string) ApiListpinsRequest {
	r.actas = &actas
	return r
}

// The version of the client making the request.
func (r ApiListpinsRequest) ClientVersion(clientVersion string) ApiListpinsRequest {
	r.clientVersion = &clientVersion
	return r
}

// The domain of the top-level page in which the client is being run. For embedded search and NSR, it will be the domain of the embedding page. For sidebar, it will be the domain of the embedding page if it&#39;s a recognized workplace app. For NTP and app.glean.com requests, it will be empty.
func (r ApiListpinsRequest) Domain(domain string) ApiListpinsRequest {
	r.domain = &domain
	return r
}

// List of experiment ids to force for incoming request.
func (r ApiListpinsRequest) Eids(eids []int64) ApiListpinsRequest {
	r.eids = &eids
	return r
}

func (r ApiListpinsRequest) Execute() (*ListPinsResponse, *http.Response, error) {
	return r.ApiService.ListpinsExecute(r)
}

/*
Listpins List all pins.

Lists all pins.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListpinsRequest
*/
func (a *PinsApiService) Listpins(ctx context.Context) ApiListpinsRequest {
	return ApiListpinsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ListPinsResponse
func (a *PinsApiService) ListpinsExecute(r ApiListpinsRequest) (*ListPinsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListPinsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PinsApiService.Listpins")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/listpins"

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

type ApiPinRequest struct {
	ctx context.Context
	ApiService *PinsApiService
	payload *PinRequest
	clientVersion *string
	domain *string
	eids *[]int64
}

// Details about the document and query for the pin.
func (r ApiPinRequest) Payload(payload PinRequest) ApiPinRequest {
	r.payload = &payload
	return r
}

// The version of the client making the request.
func (r ApiPinRequest) ClientVersion(clientVersion string) ApiPinRequest {
	r.clientVersion = &clientVersion
	return r
}

// The domain of the top-level page in which the client is being run. For embedded search and NSR, it will be the domain of the embedding page. For sidebar, it will be the domain of the embedding page if it&#39;s a recognized workplace app. For NTP and app.glean.com requests, it will be empty.
func (r ApiPinRequest) Domain(domain string) ApiPinRequest {
	r.domain = &domain
	return r
}

// List of experiment ids to force for incoming request.
func (r ApiPinRequest) Eids(eids []int64) ApiPinRequest {
	r.eids = &eids
	return r
}

func (r ApiPinRequest) Execute() (*PinDocument, *http.Response, error) {
	return r.ApiService.PinExecute(r)
}

/*
Pin Create pin

Pin a document as a result for a given search query.Pin results that are known to be a good match.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPinRequest
*/
func (a *PinsApiService) Pin(ctx context.Context) ApiPinRequest {
	return ApiPinRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PinDocument
func (a *PinsApiService) PinExecute(r ApiPinRequest) (*PinDocument, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PinDocument
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PinsApiService.Pin")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pin"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.payload == nil {
		return localVarReturnValue, nil, reportError("payload is required and must be specified")
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

type ApiUnpinRequest struct {
	ctx context.Context
	ApiService *PinsApiService
	payload *Unpin
	clientVersion *string
	domain *string
	eids *[]int64
}

// Details about the pin being unpinned.
func (r ApiUnpinRequest) Payload(payload Unpin) ApiUnpinRequest {
	r.payload = &payload
	return r
}

// The version of the client making the request.
func (r ApiUnpinRequest) ClientVersion(clientVersion string) ApiUnpinRequest {
	r.clientVersion = &clientVersion
	return r
}

// The domain of the top-level page in which the client is being run. For embedded search and NSR, it will be the domain of the embedding page. For sidebar, it will be the domain of the embedding page if it&#39;s a recognized workplace app. For NTP and app.glean.com requests, it will be empty.
func (r ApiUnpinRequest) Domain(domain string) ApiUnpinRequest {
	r.domain = &domain
	return r
}

// List of experiment ids to force for incoming request.
func (r ApiUnpinRequest) Eids(eids []int64) ApiUnpinRequest {
	r.eids = &eids
	return r
}

func (r ApiUnpinRequest) Execute() (*http.Response, error) {
	return r.ApiService.UnpinExecute(r)
}

/*
Unpin Delete pin

Unpin a previously pinned result.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUnpinRequest
*/
func (a *PinsApiService) Unpin(ctx context.Context) ApiUnpinRequest {
	return ApiUnpinRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *PinsApiService) UnpinExecute(r ApiUnpinRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PinsApiService.Unpin")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/unpin"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.payload == nil {
		return nil, reportError("payload is required and must be specified")
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
