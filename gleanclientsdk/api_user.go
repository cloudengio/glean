/*
Glean Client API - Platform Preview

# Introduction These are all the APIs used by Glean to implement the Glean client. These are available as platform preview for implementing a custom client to the Glean system.  # Usage guidelines A subset of these endpoints are also in the developer ready section, which is available for public use. The rest of the endpoints are subject to prior agreement with Glean before usage. Please contact support@glean.com if you would like to use an API that is not currently available in the developer ready section. 

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


// UserApiService UserApi service
type UserApiService service

type ApiAddcredentialRequest struct {
	ctx context.Context
	ApiService *UserApiService
	payload *AddCredentialRequest
	actas *string
	clientVersion *string
	domain *string
	eids *[]int64
}

// Credential content
func (r ApiAddcredentialRequest) Payload(payload AddCredentialRequest) ApiAddcredentialRequest {
	r.payload = &payload
	return r
}

// Email of another user to act as for debugging purposes. Requires sufficient permissions.
func (r ApiAddcredentialRequest) Actas(actas string) ApiAddcredentialRequest {
	r.actas = &actas
	return r
}

// The version of the client making the request.
func (r ApiAddcredentialRequest) ClientVersion(clientVersion string) ApiAddcredentialRequest {
	r.clientVersion = &clientVersion
	return r
}

// The domain of the top-level page in which the client is being run. For embedded search and NSR, it will be the domain of the embedding page. For sidebar, it will be the domain of the embedding page if it&#39;s a recognized workplace app. For NTP and app.glean.com requests, it will be empty.
func (r ApiAddcredentialRequest) Domain(domain string) ApiAddcredentialRequest {
	r.domain = &domain
	return r
}

// List of experiment ids to force for incoming request.
func (r ApiAddcredentialRequest) Eids(eids []int64) ApiAddcredentialRequest {
	r.eids = &eids
	return r
}

func (r ApiAddcredentialRequest) Execute() (*http.Response, error) {
	return r.ApiService.AddcredentialExecute(r)
}

/*
Addcredential Create credentials

API to save a user's credentials. Used for Confluence restricted pages and Tableau per-user auth, for example

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAddcredentialRequest
*/
func (a *UserApiService) Addcredential(ctx context.Context) ApiAddcredentialRequest {
	return ApiAddcredentialRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *UserApiService) AddcredentialExecute(r ApiAddcredentialRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserApiService.Addcredential")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/addcredential"

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

type ApiConfigRequest struct {
	ctx context.Context
	ApiService *UserApiService
	actas *string
	clientVersion *string
	domain *string
	eids *[]int64
	payload *ConfigRequest
}

// Email of another user to act as for debugging purposes. Requires sufficient permissions.
func (r ApiConfigRequest) Actas(actas string) ApiConfigRequest {
	r.actas = &actas
	return r
}

// The version of the client making the request.
func (r ApiConfigRequest) ClientVersion(clientVersion string) ApiConfigRequest {
	r.clientVersion = &clientVersion
	return r
}

// The domain of the top-level page in which the client is being run. For embedded search and NSR, it will be the domain of the embedding page. For sidebar, it will be the domain of the embedding page if it&#39;s a recognized workplace app. For NTP and app.glean.com requests, it will be empty.
func (r ApiConfigRequest) Domain(domain string) ApiConfigRequest {
	r.domain = &domain
	return r
}

// List of experiment ids to force for incoming request.
func (r ApiConfigRequest) Eids(eids []int64) ApiConfigRequest {
	r.eids = &eids
	return r
}

// Config request
func (r ApiConfigRequest) Payload(payload ConfigRequest) ApiConfigRequest {
	r.payload = &payload
	return r
}

func (r ApiConfigRequest) Execute() (*ConfigResponse, *http.Response, error) {
	return r.ApiService.ConfigExecute(r)
}

/*
Config Read client configuration

Retrieves configuration information such as the intended display of results, settings, etc.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiConfigRequest
*/
func (a *UserApiService) Config(ctx context.Context) ApiConfigRequest {
	return ApiConfigRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ConfigResponse
func (a *UserApiService) ConfigExecute(r ApiConfigRequest) (*ConfigResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ConfigResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserApiService.Config")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/config"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiInviteRequest struct {
	ctx context.Context
	ApiService *UserApiService
	payload *EmailRequest
	actas *string
	clientVersion *string
	domain *string
	eids *[]int64
}

// Invite request
func (r ApiInviteRequest) Payload(payload EmailRequest) ApiInviteRequest {
	r.payload = &payload
	return r
}

// Email of another user to act as for debugging purposes. Requires sufficient permissions.
func (r ApiInviteRequest) Actas(actas string) ApiInviteRequest {
	r.actas = &actas
	return r
}

// The version of the client making the request.
func (r ApiInviteRequest) ClientVersion(clientVersion string) ApiInviteRequest {
	r.clientVersion = &clientVersion
	return r
}

// The domain of the top-level page in which the client is being run. For embedded search and NSR, it will be the domain of the embedding page. For sidebar, it will be the domain of the embedding page if it&#39;s a recognized workplace app. For NTP and app.glean.com requests, it will be empty.
func (r ApiInviteRequest) Domain(domain string) ApiInviteRequest {
	r.domain = &domain
	return r
}

// List of experiment ids to force for incoming request.
func (r ApiInviteRequest) Eids(eids []int64) ApiInviteRequest {
	r.eids = &eids
	return r
}

func (r ApiInviteRequest) Execute() (*http.Response, error) {
	return r.ApiService.InviteExecute(r)
}

/*
Invite Send email invitation

Invites people to Glean via email

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiInviteRequest
*/
func (a *UserApiService) Invite(ctx context.Context) ApiInviteRequest {
	return ApiInviteRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *UserApiService) InviteExecute(r ApiInviteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserApiService.Invite")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/invite"

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

type ApiListusersettingsRequest struct {
	ctx context.Context
	ApiService *UserApiService
	actas *string
	clientVersion *string
	domain *string
	eids *[]int64
}

// Email of another user to act as for debugging purposes. Requires sufficient permissions.
func (r ApiListusersettingsRequest) Actas(actas string) ApiListusersettingsRequest {
	r.actas = &actas
	return r
}

// The version of the client making the request.
func (r ApiListusersettingsRequest) ClientVersion(clientVersion string) ApiListusersettingsRequest {
	r.clientVersion = &clientVersion
	return r
}

// The domain of the top-level page in which the client is being run. For embedded search and NSR, it will be the domain of the embedding page. For sidebar, it will be the domain of the embedding page if it&#39;s a recognized workplace app. For NTP and app.glean.com requests, it will be empty.
func (r ApiListusersettingsRequest) Domain(domain string) ApiListusersettingsRequest {
	r.domain = &domain
	return r
}

// List of experiment ids to force for incoming request.
func (r ApiListusersettingsRequest) Eids(eids []int64) ApiListusersettingsRequest {
	r.eids = &eids
	return r
}

func (r ApiListusersettingsRequest) Execute() (*UserSettings, *http.Response, error) {
	return r.ApiService.ListusersettingsExecute(r)
}

/*
Listusersettings Read user settings

Returns the complete list of user settings for the requestor.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListusersettingsRequest
*/
func (a *UserApiService) Listusersettings(ctx context.Context) ApiListusersettingsRequest {
	return ApiListusersettingsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return UserSettings
func (a *UserApiService) ListusersettingsExecute(r ApiListusersettingsRequest) (*UserSettings, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UserSettings
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserApiService.Listusersettings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/listusersettings"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	localVarHTTPContentTypes := []string{}

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

type ApiPublicconfigRequest struct {
	ctx context.Context
	ApiService *UserApiService
	clientVersion *string
	domain *string
	eids *[]int64
	payload *PublicConfigRequest
}

// The version of the client making the request.
func (r ApiPublicconfigRequest) ClientVersion(clientVersion string) ApiPublicconfigRequest {
	r.clientVersion = &clientVersion
	return r
}

// The domain of the top-level page in which the client is being run. For embedded search and NSR, it will be the domain of the embedding page. For sidebar, it will be the domain of the embedding page if it&#39;s a recognized workplace app. For NTP and app.glean.com requests, it will be empty.
func (r ApiPublicconfigRequest) Domain(domain string) ApiPublicconfigRequest {
	r.domain = &domain
	return r
}

// List of experiment ids to force for incoming request.
func (r ApiPublicconfigRequest) Eids(eids []int64) ApiPublicconfigRequest {
	r.eids = &eids
	return r
}

// Public Config request
func (r ApiPublicconfigRequest) Payload(payload PublicConfigRequest) ApiPublicconfigRequest {
	r.payload = &payload
	return r
}

func (r ApiPublicconfigRequest) Execute() (*ClientConfig, *http.Response, error) {
	return r.ApiService.PublicconfigExecute(r)
}

/*
Publicconfig Reads the client configuration that is public, no PII.

Retrieves configuration information such as the company name, logo and etc that is public and is not considered as PII.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPublicconfigRequest
*/
func (a *UserApiService) Publicconfig(ctx context.Context) ApiPublicconfigRequest {
	return ApiPublicconfigRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ClientConfig
func (a *UserApiService) PublicconfigExecute(r ApiPublicconfigRequest) (*ClientConfig, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ClientConfig
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserApiService.Publicconfig")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/publicclientconfig"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiRemovecredentialRequest struct {
	ctx context.Context
	ApiService *UserApiService
	payload *RemoveCredentialRequest
	actas *string
	clientVersion *string
	domain *string
	eids *[]int64
}

// Credential content
func (r ApiRemovecredentialRequest) Payload(payload RemoveCredentialRequest) ApiRemovecredentialRequest {
	r.payload = &payload
	return r
}

// Email of another user to act as for debugging purposes. Requires sufficient permissions.
func (r ApiRemovecredentialRequest) Actas(actas string) ApiRemovecredentialRequest {
	r.actas = &actas
	return r
}

// The version of the client making the request.
func (r ApiRemovecredentialRequest) ClientVersion(clientVersion string) ApiRemovecredentialRequest {
	r.clientVersion = &clientVersion
	return r
}

// The domain of the top-level page in which the client is being run. For embedded search and NSR, it will be the domain of the embedding page. For sidebar, it will be the domain of the embedding page if it&#39;s a recognized workplace app. For NTP and app.glean.com requests, it will be empty.
func (r ApiRemovecredentialRequest) Domain(domain string) ApiRemovecredentialRequest {
	r.domain = &domain
	return r
}

// List of experiment ids to force for incoming request.
func (r ApiRemovecredentialRequest) Eids(eids []int64) ApiRemovecredentialRequest {
	r.eids = &eids
	return r
}

func (r ApiRemovecredentialRequest) Execute() (*http.Response, error) {
	return r.ApiService.RemovecredentialExecute(r)
}

/*
Removecredential Delete credentials

API to remove a user's credentials. Used for Confluence restricted pages and Tableau per-user auth, for example

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiRemovecredentialRequest
*/
func (a *UserApiService) Removecredential(ctx context.Context) ApiRemovecredentialRequest {
	return ApiRemovecredentialRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *UserApiService) RemovecredentialExecute(r ApiRemovecredentialRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserApiService.Removecredential")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/removecredential"

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

type ApiSaveusersettingsRequest struct {
	ctx context.Context
	ApiService *UserApiService
	payload *UserSettings
	clientVersion *string
	domain *string
	eids *[]int64
}

// Set of user settings.
func (r ApiSaveusersettingsRequest) Payload(payload UserSettings) ApiSaveusersettingsRequest {
	r.payload = &payload
	return r
}

// The version of the client making the request.
func (r ApiSaveusersettingsRequest) ClientVersion(clientVersion string) ApiSaveusersettingsRequest {
	r.clientVersion = &clientVersion
	return r
}

// The domain of the top-level page in which the client is being run. For embedded search and NSR, it will be the domain of the embedding page. For sidebar, it will be the domain of the embedding page if it&#39;s a recognized workplace app. For NTP and app.glean.com requests, it will be empty.
func (r ApiSaveusersettingsRequest) Domain(domain string) ApiSaveusersettingsRequest {
	r.domain = &domain
	return r
}

// List of experiment ids to force for incoming request.
func (r ApiSaveusersettingsRequest) Eids(eids []int64) ApiSaveusersettingsRequest {
	r.eids = &eids
	return r
}

func (r ApiSaveusersettingsRequest) Execute() (*http.Response, error) {
	return r.ApiService.SaveusersettingsExecute(r)
}

/*
Saveusersettings Update user settings

Saves the user settings included in the payload, can be partial, does not need to be all settings.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSaveusersettingsRequest
*/
func (a *UserApiService) Saveusersettings(ctx context.Context) ApiSaveusersettingsRequest {
	return ApiSaveusersettingsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *UserApiService) SaveusersettingsExecute(r ApiSaveusersettingsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserApiService.Saveusersettings")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/saveusersettings"

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

type ApiSupportEmailRequest struct {
	ctx context.Context
	ApiService *UserApiService
	payload *EmailRequest
	actas *string
	clientVersion *string
	domain *string
	eids *[]int64
}

// Support request
func (r ApiSupportEmailRequest) Payload(payload EmailRequest) ApiSupportEmailRequest {
	r.payload = &payload
	return r
}

// Email of another user to act as for debugging purposes. Requires sufficient permissions.
func (r ApiSupportEmailRequest) Actas(actas string) ApiSupportEmailRequest {
	r.actas = &actas
	return r
}

// The version of the client making the request.
func (r ApiSupportEmailRequest) ClientVersion(clientVersion string) ApiSupportEmailRequest {
	r.clientVersion = &clientVersion
	return r
}

// The domain of the top-level page in which the client is being run. For embedded search and NSR, it will be the domain of the embedding page. For sidebar, it will be the domain of the embedding page if it&#39;s a recognized workplace app. For NTP and app.glean.com requests, it will be empty.
func (r ApiSupportEmailRequest) Domain(domain string) ApiSupportEmailRequest {
	r.domain = &domain
	return r
}

// List of experiment ids to force for incoming request.
func (r ApiSupportEmailRequest) Eids(eids []int64) ApiSupportEmailRequest {
	r.eids = &eids
	return r
}

func (r ApiSupportEmailRequest) Execute() (*http.Response, error) {
	return r.ApiService.SupportEmailExecute(r)
}

/*
SupportEmail Send support email to Glean support team

Sends a support email based on a template to the Glean support team

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSupportEmailRequest
*/
func (a *UserApiService) SupportEmail(ctx context.Context) ApiSupportEmailRequest {
	return ApiSupportEmailRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *UserApiService) SupportEmailExecute(r ApiSupportEmailRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserApiService.SupportEmail")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/support"

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
