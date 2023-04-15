# DatumCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaseInsensitive** | Pointer to **bool** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**DatumConditionType**](DatumConditionType.md) |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 

## Methods

### NewDatumCondition

`func NewDatumCondition() *DatumCondition`

NewDatumCondition instantiates a new DatumCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatumConditionWithDefaults

`func NewDatumConditionWithDefaults() *DatumCondition`

NewDatumConditionWithDefaults instantiates a new DatumCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaseInsensitive

`func (o *DatumCondition) GetCaseInsensitive() bool`

GetCaseInsensitive returns the CaseInsensitive field if non-nil, zero value otherwise.

### GetCaseInsensitiveOk

`func (o *DatumCondition) GetCaseInsensitiveOk() (*bool, bool)`

GetCaseInsensitiveOk returns a tuple with the CaseInsensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseInsensitive

`func (o *DatumCondition) SetCaseInsensitive(v bool)`

SetCaseInsensitive sets CaseInsensitive field to given value.

### HasCaseInsensitive

`func (o *DatumCondition) HasCaseInsensitive() bool`

HasCaseInsensitive returns a boolean if a field has been set.

### GetKey

`func (o *DatumCondition) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *DatumCondition) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *DatumCondition) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *DatumCondition) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetType

`func (o *DatumCondition) GetType() DatumConditionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DatumCondition) GetTypeOk() (*DatumConditionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DatumCondition) SetType(v DatumConditionType)`

SetType sets Type field to given value.

### HasType

`func (o *DatumCondition) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *DatumCondition) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DatumCondition) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DatumCondition) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *DatumCondition) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


