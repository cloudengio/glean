# CheckActasAuthResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsValid** | **bool** | Whether the request supplied a valid actas auth token. | 
**ErrMsg** | Pointer to **string** | Error message | [optional] 

## Methods

### NewCheckActasAuthResponse

`func NewCheckActasAuthResponse(isValid bool, ) *CheckActasAuthResponse`

NewCheckActasAuthResponse instantiates a new CheckActasAuthResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckActasAuthResponseWithDefaults

`func NewCheckActasAuthResponseWithDefaults() *CheckActasAuthResponse`

NewCheckActasAuthResponseWithDefaults instantiates a new CheckActasAuthResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsValid

`func (o *CheckActasAuthResponse) GetIsValid() bool`

GetIsValid returns the IsValid field if non-nil, zero value otherwise.

### GetIsValidOk

`func (o *CheckActasAuthResponse) GetIsValidOk() (*bool, bool)`

GetIsValidOk returns a tuple with the IsValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsValid

`func (o *CheckActasAuthResponse) SetIsValid(v bool)`

SetIsValid sets IsValid field to given value.


### GetErrMsg

`func (o *CheckActasAuthResponse) GetErrMsg() string`

GetErrMsg returns the ErrMsg field if non-nil, zero value otherwise.

### GetErrMsgOk

`func (o *CheckActasAuthResponse) GetErrMsgOk() (*string, bool)`

GetErrMsgOk returns a tuple with the ErrMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrMsg

`func (o *CheckActasAuthResponse) SetErrMsg(v string)`

SetErrMsg sets ErrMsg field to given value.

### HasErrMsg

`func (o *CheckActasAuthResponse) HasErrMsg() bool`

HasErrMsg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


