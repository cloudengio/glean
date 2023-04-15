# AutocompleteResponseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrackingToken** | Pointer to **string** | An opaque token that represents this particular set of autocomplete results. To be used for /feedback reporting. | [optional] 
**SessionInfo** | Pointer to [**SessionInfo**](SessionInfo.md) |  | [optional] 
**Results** | Pointer to [**[]AutocompleteResult**](AutocompleteResult.md) |  | [optional] 
**Groups** | Pointer to [**[]AutocompleteResultGroup**](AutocompleteResultGroup.md) | Subsections of the results list from which distinct sections should be created. | [optional] 
**ErrorInfo** | Pointer to [**ErrorInfo**](ErrorInfo.md) |  | [optional] 
**BackendTimeMillis** | Pointer to **int64** | Time in milliseconds the backend took to respond to the request. | [optional] 

## Methods

### NewAutocompleteResponseAllOf

`func NewAutocompleteResponseAllOf() *AutocompleteResponseAllOf`

NewAutocompleteResponseAllOf instantiates a new AutocompleteResponseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutocompleteResponseAllOfWithDefaults

`func NewAutocompleteResponseAllOfWithDefaults() *AutocompleteResponseAllOf`

NewAutocompleteResponseAllOfWithDefaults instantiates a new AutocompleteResponseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrackingToken

`func (o *AutocompleteResponseAllOf) GetTrackingToken() string`

GetTrackingToken returns the TrackingToken field if non-nil, zero value otherwise.

### GetTrackingTokenOk

`func (o *AutocompleteResponseAllOf) GetTrackingTokenOk() (*string, bool)`

GetTrackingTokenOk returns a tuple with the TrackingToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingToken

`func (o *AutocompleteResponseAllOf) SetTrackingToken(v string)`

SetTrackingToken sets TrackingToken field to given value.

### HasTrackingToken

`func (o *AutocompleteResponseAllOf) HasTrackingToken() bool`

HasTrackingToken returns a boolean if a field has been set.

### GetSessionInfo

`func (o *AutocompleteResponseAllOf) GetSessionInfo() SessionInfo`

GetSessionInfo returns the SessionInfo field if non-nil, zero value otherwise.

### GetSessionInfoOk

`func (o *AutocompleteResponseAllOf) GetSessionInfoOk() (*SessionInfo, bool)`

GetSessionInfoOk returns a tuple with the SessionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionInfo

`func (o *AutocompleteResponseAllOf) SetSessionInfo(v SessionInfo)`

SetSessionInfo sets SessionInfo field to given value.

### HasSessionInfo

`func (o *AutocompleteResponseAllOf) HasSessionInfo() bool`

HasSessionInfo returns a boolean if a field has been set.

### GetResults

`func (o *AutocompleteResponseAllOf) GetResults() []AutocompleteResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *AutocompleteResponseAllOf) GetResultsOk() (*[]AutocompleteResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *AutocompleteResponseAllOf) SetResults(v []AutocompleteResult)`

SetResults sets Results field to given value.

### HasResults

`func (o *AutocompleteResponseAllOf) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetGroups

`func (o *AutocompleteResponseAllOf) GetGroups() []AutocompleteResultGroup`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *AutocompleteResponseAllOf) GetGroupsOk() (*[]AutocompleteResultGroup, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *AutocompleteResponseAllOf) SetGroups(v []AutocompleteResultGroup)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *AutocompleteResponseAllOf) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetErrorInfo

`func (o *AutocompleteResponseAllOf) GetErrorInfo() ErrorInfo`

GetErrorInfo returns the ErrorInfo field if non-nil, zero value otherwise.

### GetErrorInfoOk

`func (o *AutocompleteResponseAllOf) GetErrorInfoOk() (*ErrorInfo, bool)`

GetErrorInfoOk returns a tuple with the ErrorInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorInfo

`func (o *AutocompleteResponseAllOf) SetErrorInfo(v ErrorInfo)`

SetErrorInfo sets ErrorInfo field to given value.

### HasErrorInfo

`func (o *AutocompleteResponseAllOf) HasErrorInfo() bool`

HasErrorInfo returns a boolean if a field has been set.

### GetBackendTimeMillis

`func (o *AutocompleteResponseAllOf) GetBackendTimeMillis() int64`

GetBackendTimeMillis returns the BackendTimeMillis field if non-nil, zero value otherwise.

### GetBackendTimeMillisOk

`func (o *AutocompleteResponseAllOf) GetBackendTimeMillisOk() (*int64, bool)`

GetBackendTimeMillisOk returns a tuple with the BackendTimeMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendTimeMillis

`func (o *AutocompleteResponseAllOf) SetBackendTimeMillis(v int64)`

SetBackendTimeMillis sets BackendTimeMillis field to given value.

### HasBackendTimeMillis

`func (o *AutocompleteResponseAllOf) HasBackendTimeMillis() bool`

HasBackendTimeMillis returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


