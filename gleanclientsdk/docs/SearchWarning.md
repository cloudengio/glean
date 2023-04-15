# SearchWarning

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WarningType** | **string** | The type of the warning. | 
**LastUsedTerm** | Pointer to **string** | The last term we considered in the user&#39;s long query. | [optional] 
**QuotesIgnoredQuery** | Pointer to **string** | The query after ignoring/removing quotes. | [optional] 

## Methods

### NewSearchWarning

`func NewSearchWarning(warningType string, ) *SearchWarning`

NewSearchWarning instantiates a new SearchWarning object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchWarningWithDefaults

`func NewSearchWarningWithDefaults() *SearchWarning`

NewSearchWarningWithDefaults instantiates a new SearchWarning object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWarningType

`func (o *SearchWarning) GetWarningType() string`

GetWarningType returns the WarningType field if non-nil, zero value otherwise.

### GetWarningTypeOk

`func (o *SearchWarning) GetWarningTypeOk() (*string, bool)`

GetWarningTypeOk returns a tuple with the WarningType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningType

`func (o *SearchWarning) SetWarningType(v string)`

SetWarningType sets WarningType field to given value.


### GetLastUsedTerm

`func (o *SearchWarning) GetLastUsedTerm() string`

GetLastUsedTerm returns the LastUsedTerm field if non-nil, zero value otherwise.

### GetLastUsedTermOk

`func (o *SearchWarning) GetLastUsedTermOk() (*string, bool)`

GetLastUsedTermOk returns a tuple with the LastUsedTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedTerm

`func (o *SearchWarning) SetLastUsedTerm(v string)`

SetLastUsedTerm sets LastUsedTerm field to given value.

### HasLastUsedTerm

`func (o *SearchWarning) HasLastUsedTerm() bool`

HasLastUsedTerm returns a boolean if a field has been set.

### GetQuotesIgnoredQuery

`func (o *SearchWarning) GetQuotesIgnoredQuery() string`

GetQuotesIgnoredQuery returns the QuotesIgnoredQuery field if non-nil, zero value otherwise.

### GetQuotesIgnoredQueryOk

`func (o *SearchWarning) GetQuotesIgnoredQueryOk() (*string, bool)`

GetQuotesIgnoredQueryOk returns a tuple with the QuotesIgnoredQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotesIgnoredQuery

`func (o *SearchWarning) SetQuotesIgnoredQuery(v string)`

SetQuotesIgnoredQuery sets QuotesIgnoredQuery field to given value.

### HasQuotesIgnoredQuery

`func (o *SearchWarning) HasQuotesIgnoredQuery() bool`

HasQuotesIgnoredQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


