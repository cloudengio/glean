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


// AnswersApiService AnswersApi service
type AnswersApiService service

type ApiCreateanswerRequest struct {
	ctx context.Context
	ApiService *AnswersApiService
	payload *CreateAnswerRequest
	actas *string
	clientVersion *string
	domain *string
	eids *[]int64
}

// CreateAnswer request
func (r ApiCreateanswerRequest) Payload(payload CreateAnswerRequest) ApiCreateanswerRequest {
	r.payload = &payload
	return r
}

// Email of another user to act as for debugging purposes. Requires sufficient permissions.
func (r ApiCreateanswerRequest) Actas(actas string) ApiCreateanswerRequest {
	r.actas = &actas
	return r
}

// The version of the client making the request.
func (r ApiCreateanswerRequest) ClientVersion(clientVersion string) ApiCreateanswerRequest {
	r.clientVersion = &clientVersion
	return r
}

// The domain of the top-level page in which the client is being run. For embedded search and NSR, it will be the domain of the embedding page. For sidebar, it will be the domain of the embedding page if it&#39;s a recognized workplace app. For NTP and app.glean.com requests, it will be empty.
func (r ApiCreateanswerRequest) Domain(domain string) ApiCreateanswerRequest {
	r.domain = &domain
	return r
}

// List of experiment ids to force for incoming request.
func (r ApiCreateanswerRequest) Eids(eids []int64) ApiCreateanswerRequest {
	r.eids = &eids
	return r
}

func (r ApiCreateanswerRequest) Execute() (*Answer, *http.Response, error) {
	return r.ApiService.CreateanswerExecute(r)
}

/*
Createanswer Create answer

Creates a user-generated answer that contains a question and answer.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateanswerRequest
*/
func (a *AnswersApiService) Createanswer(ctx context.Context) ApiCreateanswerRequest {
	return ApiCreateanswerRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Answer
func (a *AnswersApiService) CreateanswerExecute(r ApiCreateanswerRequest) (*Answer, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Answer
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AnswersApiService.Createanswer")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/createanswer"

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

type ApiDeleteanswerRequest struct {
	ctx context.Context
	ApiService *AnswersApiService
	payload *DeleteAnswerRequest
	actas *string
	clientVersion *string
	domain *string
	eids *[]int64
}

// DeleteAnswer request
func (r ApiDeleteanswerRequest) Payload(payload DeleteAnswerRequest) ApiDeleteanswerRequest {
	r.payload = &payload
	return r
}

// Email of another user to act as for debugging purposes. Requires sufficient permissions.
func (r ApiDeleteanswerRequest) Actas(actas string) ApiDeleteanswerRequest {
	r.actas = &actas
	return r
}

// The version of the client making the request.
func (r ApiDeleteanswerRequest) ClientVersion(clientVersion string) ApiDeleteanswerRequest {
	r.clientVersion = &clientVersion
	return r
}

// The domain of the top-level page in which the client is being run. For embedded search and NSR, it will be the domain of the embedding page. For sidebar, it will be the domain of the embedding page if it&#39;s a recognized workplace app. For NTP and app.glean.com requests, it will be empty.
func (r ApiDeleteanswerRequest) Domain(domain string) ApiDeleteanswerRequest {
	r.domain = &domain
	return r
}

// List of experiment ids to force for incoming request.
func (r ApiDeleteanswerRequest) Eids(eids []int64) ApiDeleteanswerRequest {
	r.eids = &eids
	return r
}

func (r ApiDeleteanswerRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteanswerExecute(r)
}

/*
Deleteanswer Delete answer

Deletes an existing user-generated answer.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDeleteanswerRequest
*/
func (a *AnswersApiService) Deleteanswer(ctx context.Context) ApiDeleteanswerRequest {
	return ApiDeleteanswerRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *AnswersApiService) DeleteanswerExecute(r ApiDeleteanswerRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AnswersApiService.Deleteanswer")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/deleteanswer"

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

type ApiEditanswerRequest struct {
	ctx context.Context
	ApiService *AnswersApiService
	payload *EditAnswerRequest
	actas *string
	clientVersion *string
	domain *string
	eids *[]int64
}

// EditAnswer request
func (r ApiEditanswerRequest) Payload(payload EditAnswerRequest) ApiEditanswerRequest {
	r.payload = &payload
	return r
}

// Email of another user to act as for debugging purposes. Requires sufficient permissions.
func (r ApiEditanswerRequest) Actas(actas string) ApiEditanswerRequest {
	r.actas = &actas
	return r
}

// The version of the client making the request.
func (r ApiEditanswerRequest) ClientVersion(clientVersion string) ApiEditanswerRequest {
	r.clientVersion = &clientVersion
	return r
}

// The domain of the top-level page in which the client is being run. For embedded search and NSR, it will be the domain of the embedding page. For sidebar, it will be the domain of the embedding page if it&#39;s a recognized workplace app. For NTP and app.glean.com requests, it will be empty.
func (r ApiEditanswerRequest) Domain(domain string) ApiEditanswerRequest {
	r.domain = &domain
	return r
}

// List of experiment ids to force for incoming request.
func (r ApiEditanswerRequest) Eids(eids []int64) ApiEditanswerRequest {
	r.eids = &eids
	return r
}

func (r ApiEditanswerRequest) Execute() (*Answer, *http.Response, error) {
	return r.ApiService.EditanswerExecute(r)
}

/*
Editanswer Edit answer

Edits an existing user-generated answer.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiEditanswerRequest
*/
func (a *AnswersApiService) Editanswer(ctx context.Context) ApiEditanswerRequest {
	return ApiEditanswerRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Answer
func (a *AnswersApiService) EditanswerExecute(r ApiEditanswerRequest) (*Answer, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Answer
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AnswersApiService.Editanswer")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/editanswer"

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

type ApiListanswersRequest struct {
	ctx context.Context
	ApiService *AnswersApiService
	payload *ListAnswersRequest
	actas *string
	clientVersion *string
	domain *string
	eids *[]int64
}

// ListAnswers request
func (r ApiListanswersRequest) Payload(payload ListAnswersRequest) ApiListanswersRequest {
	r.payload = &payload
	return r
}

// Email of another user to act as for debugging purposes. Requires sufficient permissions.
func (r ApiListanswersRequest) Actas(actas string) ApiListanswersRequest {
	r.actas = &actas
	return r
}

// The version of the client making the request.
func (r ApiListanswersRequest) ClientVersion(clientVersion string) ApiListanswersRequest {
	r.clientVersion = &clientVersion
	return r
}

// The domain of the top-level page in which the client is being run. For embedded search and NSR, it will be the domain of the embedding page. For sidebar, it will be the domain of the embedding page if it&#39;s a recognized workplace app. For NTP and app.glean.com requests, it will be empty.
func (r ApiListanswersRequest) Domain(domain string) ApiListanswersRequest {
	r.domain = &domain
	return r
}

// List of experiment ids to force for incoming request.
func (r ApiListanswersRequest) Eids(eids []int64) ApiListanswersRequest {
	r.eids = &eids
	return r
}

func (r ApiListanswersRequest) Execute() (*ListAnswersResponse, *http.Response, error) {
	return r.ApiService.ListanswersExecute(r)
}

/*
Listanswers List answers created by the authed user

Lists answers created by the authed user.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListanswersRequest
*/
func (a *AnswersApiService) Listanswers(ctx context.Context) ApiListanswersRequest {
	return ApiListanswersRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ListAnswersResponse
func (a *AnswersApiService) ListanswersExecute(r ApiListanswersRequest) (*ListAnswersResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListAnswersResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AnswersApiService.Listanswers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/listanswers"

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