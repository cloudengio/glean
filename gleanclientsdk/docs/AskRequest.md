# AskRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DetectOnly** | Pointer to **bool** | Whether to apply apply only question detection and not answering. | [optional] 
**SearchRequest** | [**SearchRequest**](SearchRequest.md) |  | 
**Operators** | Pointer to **string** | Search operators to append to the query | [optional] 

## Methods

### NewAskRequest

`func NewAskRequest(searchRequest SearchRequest, ) *AskRequest`

NewAskRequest instantiates a new AskRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAskRequestWithDefaults

`func NewAskRequestWithDefaults() *AskRequest`

NewAskRequestWithDefaults instantiates a new AskRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetectOnly

`func (o *AskRequest) GetDetectOnly() bool`

GetDetectOnly returns the DetectOnly field if non-nil, zero value otherwise.

### GetDetectOnlyOk

`func (o *AskRequest) GetDetectOnlyOk() (*bool, bool)`

GetDetectOnlyOk returns a tuple with the DetectOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectOnly

`func (o *AskRequest) SetDetectOnly(v bool)`

SetDetectOnly sets DetectOnly field to given value.

### HasDetectOnly

`func (o *AskRequest) HasDetectOnly() bool`

HasDetectOnly returns a boolean if a field has been set.

### GetSearchRequest

`func (o *AskRequest) GetSearchRequest() SearchRequest`

GetSearchRequest returns the SearchRequest field if non-nil, zero value otherwise.

### GetSearchRequestOk

`func (o *AskRequest) GetSearchRequestOk() (*SearchRequest, bool)`

GetSearchRequestOk returns a tuple with the SearchRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchRequest

`func (o *AskRequest) SetSearchRequest(v SearchRequest)`

SetSearchRequest sets SearchRequest field to given value.


### GetOperators

`func (o *AskRequest) GetOperators() string`

GetOperators returns the Operators field if non-nil, zero value otherwise.

### GetOperatorsOk

`func (o *AskRequest) GetOperatorsOk() (*string, bool)`

GetOperatorsOk returns a tuple with the Operators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperators

`func (o *AskRequest) SetOperators(v string)`

SetOperators sets Operators field to given value.

### HasOperators

`func (o *AskRequest) HasOperators() bool`

HasOperators returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


