# OAuthConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKey** | Pointer to **string** | An API key required to authenticate. | [optional] 
**ClientId** | Pointer to **string** | Client ID of the app registered in the provider&#39;s portal. | [optional] 
**Scope** | Pointer to **string** | Space separated list of required scopes for this datasource. | [optional] 

## Methods

### NewOAuthConfig

`func NewOAuthConfig() *OAuthConfig`

NewOAuthConfig instantiates a new OAuthConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuthConfigWithDefaults

`func NewOAuthConfigWithDefaults() *OAuthConfig`

NewOAuthConfigWithDefaults instantiates a new OAuthConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKey

`func (o *OAuthConfig) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *OAuthConfig) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *OAuthConfig) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *OAuthConfig) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetClientId

`func (o *OAuthConfig) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *OAuthConfig) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *OAuthConfig) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *OAuthConfig) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetScope

`func (o *OAuthConfig) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *OAuthConfig) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *OAuthConfig) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *OAuthConfig) HasScope() bool`

HasScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


