# SearchDebugOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ForceParseQuery** | Pointer to **string** | Debug only. When present, uses this parsed query JSON string instead of the query in the request. Requires elevated permission. | [optional] 

## Methods

### NewSearchDebugOptions

`func NewSearchDebugOptions() *SearchDebugOptions`

NewSearchDebugOptions instantiates a new SearchDebugOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchDebugOptionsWithDefaults

`func NewSearchDebugOptionsWithDefaults() *SearchDebugOptions`

NewSearchDebugOptionsWithDefaults instantiates a new SearchDebugOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForceParseQuery

`func (o *SearchDebugOptions) GetForceParseQuery() string`

GetForceParseQuery returns the ForceParseQuery field if non-nil, zero value otherwise.

### GetForceParseQueryOk

`func (o *SearchDebugOptions) GetForceParseQueryOk() (*string, bool)`

GetForceParseQueryOk returns a tuple with the ForceParseQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceParseQuery

`func (o *SearchDebugOptions) SetForceParseQuery(v string)`

SetForceParseQuery sets ForceParseQuery field to given value.

### HasForceParseQuery

`func (o *SearchDebugOptions) HasForceParseQuery() bool`

HasForceParseQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


