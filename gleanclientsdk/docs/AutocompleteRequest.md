# AutocompleteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrackingToken** | Pointer to **string** |  | [optional] 
**SessionInfo** | Pointer to [**SessionInfo**](SessionInfo.md) |  | [optional] 
**SourceInfo** | Pointer to [**SearchRequestSourceInfo**](SearchRequestSourceInfo.md) |  | [optional] 
**Query** | Pointer to **string** | Partially typed query. | [optional] 
**DatasourcesFilter** | Pointer to **[]string** | Filter results to only those relevant to one or more datasources (e.g. jira, gdrive). Results are unfiltered if missing. | [optional] 
**Datasource** | Pointer to **string** | Filter to only return results relevant to the given datasource. | [optional] 
**ResultTypes** | Pointer to **[]string** | Filter to only return results of the given type(s). All types may be returned if omitted. | [optional] 
**ResultSize** | Pointer to **int32** | Maximum number of results to be returned. If no value is provided, the backend will cap at 200.  | [optional] 
**AuthTokens** | Pointer to [**[]AuthToken**](AuthToken.md) | Auth tokens which may be used for federated results. | [optional] 
**Sc** | Pointer to **string** | Debug only search params to to change the flow of control in request handling. It will be passed around service boundaries/endpoints. For more details, | [optional] 

## Methods

### NewAutocompleteRequest

`func NewAutocompleteRequest() *AutocompleteRequest`

NewAutocompleteRequest instantiates a new AutocompleteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutocompleteRequestWithDefaults

`func NewAutocompleteRequestWithDefaults() *AutocompleteRequest`

NewAutocompleteRequestWithDefaults instantiates a new AutocompleteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrackingToken

`func (o *AutocompleteRequest) GetTrackingToken() string`

GetTrackingToken returns the TrackingToken field if non-nil, zero value otherwise.

### GetTrackingTokenOk

`func (o *AutocompleteRequest) GetTrackingTokenOk() (*string, bool)`

GetTrackingTokenOk returns a tuple with the TrackingToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingToken

`func (o *AutocompleteRequest) SetTrackingToken(v string)`

SetTrackingToken sets TrackingToken field to given value.

### HasTrackingToken

`func (o *AutocompleteRequest) HasTrackingToken() bool`

HasTrackingToken returns a boolean if a field has been set.

### GetSessionInfo

`func (o *AutocompleteRequest) GetSessionInfo() SessionInfo`

GetSessionInfo returns the SessionInfo field if non-nil, zero value otherwise.

### GetSessionInfoOk

`func (o *AutocompleteRequest) GetSessionInfoOk() (*SessionInfo, bool)`

GetSessionInfoOk returns a tuple with the SessionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionInfo

`func (o *AutocompleteRequest) SetSessionInfo(v SessionInfo)`

SetSessionInfo sets SessionInfo field to given value.

### HasSessionInfo

`func (o *AutocompleteRequest) HasSessionInfo() bool`

HasSessionInfo returns a boolean if a field has been set.

### GetSourceInfo

`func (o *AutocompleteRequest) GetSourceInfo() SearchRequestSourceInfo`

GetSourceInfo returns the SourceInfo field if non-nil, zero value otherwise.

### GetSourceInfoOk

`func (o *AutocompleteRequest) GetSourceInfoOk() (*SearchRequestSourceInfo, bool)`

GetSourceInfoOk returns a tuple with the SourceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceInfo

`func (o *AutocompleteRequest) SetSourceInfo(v SearchRequestSourceInfo)`

SetSourceInfo sets SourceInfo field to given value.

### HasSourceInfo

`func (o *AutocompleteRequest) HasSourceInfo() bool`

HasSourceInfo returns a boolean if a field has been set.

### GetQuery

`func (o *AutocompleteRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *AutocompleteRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *AutocompleteRequest) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *AutocompleteRequest) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetDatasourcesFilter

`func (o *AutocompleteRequest) GetDatasourcesFilter() []string`

GetDatasourcesFilter returns the DatasourcesFilter field if non-nil, zero value otherwise.

### GetDatasourcesFilterOk

`func (o *AutocompleteRequest) GetDatasourcesFilterOk() (*[]string, bool)`

GetDatasourcesFilterOk returns a tuple with the DatasourcesFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasourcesFilter

`func (o *AutocompleteRequest) SetDatasourcesFilter(v []string)`

SetDatasourcesFilter sets DatasourcesFilter field to given value.

### HasDatasourcesFilter

`func (o *AutocompleteRequest) HasDatasourcesFilter() bool`

HasDatasourcesFilter returns a boolean if a field has been set.

### GetDatasource

`func (o *AutocompleteRequest) GetDatasource() string`

GetDatasource returns the Datasource field if non-nil, zero value otherwise.

### GetDatasourceOk

`func (o *AutocompleteRequest) GetDatasourceOk() (*string, bool)`

GetDatasourceOk returns a tuple with the Datasource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasource

`func (o *AutocompleteRequest) SetDatasource(v string)`

SetDatasource sets Datasource field to given value.

### HasDatasource

`func (o *AutocompleteRequest) HasDatasource() bool`

HasDatasource returns a boolean if a field has been set.

### GetResultTypes

`func (o *AutocompleteRequest) GetResultTypes() []string`

GetResultTypes returns the ResultTypes field if non-nil, zero value otherwise.

### GetResultTypesOk

`func (o *AutocompleteRequest) GetResultTypesOk() (*[]string, bool)`

GetResultTypesOk returns a tuple with the ResultTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultTypes

`func (o *AutocompleteRequest) SetResultTypes(v []string)`

SetResultTypes sets ResultTypes field to given value.

### HasResultTypes

`func (o *AutocompleteRequest) HasResultTypes() bool`

HasResultTypes returns a boolean if a field has been set.

### GetResultSize

`func (o *AutocompleteRequest) GetResultSize() int32`

GetResultSize returns the ResultSize field if non-nil, zero value otherwise.

### GetResultSizeOk

`func (o *AutocompleteRequest) GetResultSizeOk() (*int32, bool)`

GetResultSizeOk returns a tuple with the ResultSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultSize

`func (o *AutocompleteRequest) SetResultSize(v int32)`

SetResultSize sets ResultSize field to given value.

### HasResultSize

`func (o *AutocompleteRequest) HasResultSize() bool`

HasResultSize returns a boolean if a field has been set.

### GetAuthTokens

`func (o *AutocompleteRequest) GetAuthTokens() []AuthToken`

GetAuthTokens returns the AuthTokens field if non-nil, zero value otherwise.

### GetAuthTokensOk

`func (o *AutocompleteRequest) GetAuthTokensOk() (*[]AuthToken, bool)`

GetAuthTokensOk returns a tuple with the AuthTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthTokens

`func (o *AutocompleteRequest) SetAuthTokens(v []AuthToken)`

SetAuthTokens sets AuthTokens field to given value.

### HasAuthTokens

`func (o *AutocompleteRequest) HasAuthTokens() bool`

HasAuthTokens returns a boolean if a field has been set.

### GetSc

`func (o *AutocompleteRequest) GetSc() string`

GetSc returns the Sc field if non-nil, zero value otherwise.

### GetScOk

`func (o *AutocompleteRequest) GetScOk() (*string, bool)`

GetScOk returns a tuple with the Sc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSc

`func (o *AutocompleteRequest) SetSc(v string)`

SetSc sets Sc field to given value.

### HasSc

`func (o *AutocompleteRequest) HasSc() bool`

HasSc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


