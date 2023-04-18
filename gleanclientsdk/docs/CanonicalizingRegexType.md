# CanonicalizingRegexType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MatchRegex** | Pointer to **string** | Regular expression to match to an arbitrary string. | [optional] 
**RewriteRegex** | Pointer to **string** | Regular expression to transform into a canonical string. | [optional] 

## Methods

### NewCanonicalizingRegexType

`func NewCanonicalizingRegexType() *CanonicalizingRegexType`

NewCanonicalizingRegexType instantiates a new CanonicalizingRegexType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCanonicalizingRegexTypeWithDefaults

`func NewCanonicalizingRegexTypeWithDefaults() *CanonicalizingRegexType`

NewCanonicalizingRegexTypeWithDefaults instantiates a new CanonicalizingRegexType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMatchRegex

`func (o *CanonicalizingRegexType) GetMatchRegex() string`

GetMatchRegex returns the MatchRegex field if non-nil, zero value otherwise.

### GetMatchRegexOk

`func (o *CanonicalizingRegexType) GetMatchRegexOk() (*string, bool)`

GetMatchRegexOk returns a tuple with the MatchRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchRegex

`func (o *CanonicalizingRegexType) SetMatchRegex(v string)`

SetMatchRegex sets MatchRegex field to given value.

### HasMatchRegex

`func (o *CanonicalizingRegexType) HasMatchRegex() bool`

HasMatchRegex returns a boolean if a field has been set.

### GetRewriteRegex

`func (o *CanonicalizingRegexType) GetRewriteRegex() string`

GetRewriteRegex returns the RewriteRegex field if non-nil, zero value otherwise.

### GetRewriteRegexOk

`func (o *CanonicalizingRegexType) GetRewriteRegexOk() (*string, bool)`

GetRewriteRegexOk returns a tuple with the RewriteRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewriteRegex

`func (o *CanonicalizingRegexType) SetRewriteRegex(v string)`

SetRewriteRegex sets RewriteRegex field to given value.

### HasRewriteRegex

`func (o *CanonicalizingRegexType) HasRewriteRegex() bool`

HasRewriteRegex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


