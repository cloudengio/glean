# SearchRequestAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | Pointer to **string** | The search terms. | [optional] 
**Cursor** | Pointer to **string** | Pagination cursor. A previously received opaque token representing the position in the overall results at which to start. | [optional] 
**ResultTabIds** | Pointer to **[]string** | The unique IDs of the result tabs for which to fetch results. This will have precedence over datasource filters if both are specified and in conflict. | [optional] 
**InputDetails** | Pointer to [**SearchRequestInputDetails**](SearchRequestInputDetails.md) |  | [optional] 
**RequestOptions** | Pointer to [**SearchRequestOptions**](SearchRequestOptions.md) |  | [optional] 
**TimeoutMillis** | Pointer to **int32** | Timeout in milliseconds for the request. Backend should throw a 408 if request takes longer than this. | [optional] 
**People** | Pointer to [**[]Person**](Person.md) | People associated with the search request. Hints to the server to fetch additional information for these people. Note that in this request, an email may be used as a person&#39;s obfuscatedId value. | [optional] 
**DisableSpellcheck** | Pointer to **bool** | Whether or not to disable spellcheck. | [optional] 

## Methods

### NewSearchRequestAllOf

`func NewSearchRequestAllOf() *SearchRequestAllOf`

NewSearchRequestAllOf instantiates a new SearchRequestAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchRequestAllOfWithDefaults

`func NewSearchRequestAllOfWithDefaults() *SearchRequestAllOf`

NewSearchRequestAllOfWithDefaults instantiates a new SearchRequestAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *SearchRequestAllOf) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *SearchRequestAllOf) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *SearchRequestAllOf) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *SearchRequestAllOf) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetCursor

`func (o *SearchRequestAllOf) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *SearchRequestAllOf) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *SearchRequestAllOf) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *SearchRequestAllOf) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### GetResultTabIds

`func (o *SearchRequestAllOf) GetResultTabIds() []string`

GetResultTabIds returns the ResultTabIds field if non-nil, zero value otherwise.

### GetResultTabIdsOk

`func (o *SearchRequestAllOf) GetResultTabIdsOk() (*[]string, bool)`

GetResultTabIdsOk returns a tuple with the ResultTabIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultTabIds

`func (o *SearchRequestAllOf) SetResultTabIds(v []string)`

SetResultTabIds sets ResultTabIds field to given value.

### HasResultTabIds

`func (o *SearchRequestAllOf) HasResultTabIds() bool`

HasResultTabIds returns a boolean if a field has been set.

### GetInputDetails

`func (o *SearchRequestAllOf) GetInputDetails() SearchRequestInputDetails`

GetInputDetails returns the InputDetails field if non-nil, zero value otherwise.

### GetInputDetailsOk

`func (o *SearchRequestAllOf) GetInputDetailsOk() (*SearchRequestInputDetails, bool)`

GetInputDetailsOk returns a tuple with the InputDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputDetails

`func (o *SearchRequestAllOf) SetInputDetails(v SearchRequestInputDetails)`

SetInputDetails sets InputDetails field to given value.

### HasInputDetails

`func (o *SearchRequestAllOf) HasInputDetails() bool`

HasInputDetails returns a boolean if a field has been set.

### GetRequestOptions

`func (o *SearchRequestAllOf) GetRequestOptions() SearchRequestOptions`

GetRequestOptions returns the RequestOptions field if non-nil, zero value otherwise.

### GetRequestOptionsOk

`func (o *SearchRequestAllOf) GetRequestOptionsOk() (*SearchRequestOptions, bool)`

GetRequestOptionsOk returns a tuple with the RequestOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestOptions

`func (o *SearchRequestAllOf) SetRequestOptions(v SearchRequestOptions)`

SetRequestOptions sets RequestOptions field to given value.

### HasRequestOptions

`func (o *SearchRequestAllOf) HasRequestOptions() bool`

HasRequestOptions returns a boolean if a field has been set.

### GetTimeoutMillis

`func (o *SearchRequestAllOf) GetTimeoutMillis() int32`

GetTimeoutMillis returns the TimeoutMillis field if non-nil, zero value otherwise.

### GetTimeoutMillisOk

`func (o *SearchRequestAllOf) GetTimeoutMillisOk() (*int32, bool)`

GetTimeoutMillisOk returns a tuple with the TimeoutMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutMillis

`func (o *SearchRequestAllOf) SetTimeoutMillis(v int32)`

SetTimeoutMillis sets TimeoutMillis field to given value.

### HasTimeoutMillis

`func (o *SearchRequestAllOf) HasTimeoutMillis() bool`

HasTimeoutMillis returns a boolean if a field has been set.

### GetPeople

`func (o *SearchRequestAllOf) GetPeople() []Person`

GetPeople returns the People field if non-nil, zero value otherwise.

### GetPeopleOk

`func (o *SearchRequestAllOf) GetPeopleOk() (*[]Person, bool)`

GetPeopleOk returns a tuple with the People field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeople

`func (o *SearchRequestAllOf) SetPeople(v []Person)`

SetPeople sets People field to given value.

### HasPeople

`func (o *SearchRequestAllOf) HasPeople() bool`

HasPeople returns a boolean if a field has been set.

### GetDisableSpellcheck

`func (o *SearchRequestAllOf) GetDisableSpellcheck() bool`

GetDisableSpellcheck returns the DisableSpellcheck field if non-nil, zero value otherwise.

### GetDisableSpellcheckOk

`func (o *SearchRequestAllOf) GetDisableSpellcheckOk() (*bool, bool)`

GetDisableSpellcheckOk returns a tuple with the DisableSpellcheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableSpellcheck

`func (o *SearchRequestAllOf) SetDisableSpellcheck(v bool)`

SetDisableSpellcheck sets DisableSpellcheck field to given value.

### HasDisableSpellcheck

`func (o *SearchRequestAllOf) HasDisableSpellcheck() bool`

HasDisableSpellcheck returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


