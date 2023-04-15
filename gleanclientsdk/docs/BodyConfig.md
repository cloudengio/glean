# BodyConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatumConfigs** | Pointer to [**map[string]DatumConfig**](DatumConfig.md) |  | [optional] 
**Keys** | Pointer to **[]string** |  | [optional] 
**Separator** | Pointer to **string** |  | [optional] 
**Hide** | Pointer to **bool** |  | [optional] 

## Methods

### NewBodyConfig

`func NewBodyConfig() *BodyConfig`

NewBodyConfig instantiates a new BodyConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBodyConfigWithDefaults

`func NewBodyConfigWithDefaults() *BodyConfig`

NewBodyConfigWithDefaults instantiates a new BodyConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatumConfigs

`func (o *BodyConfig) GetDatumConfigs() map[string]DatumConfig`

GetDatumConfigs returns the DatumConfigs field if non-nil, zero value otherwise.

### GetDatumConfigsOk

`func (o *BodyConfig) GetDatumConfigsOk() (*map[string]DatumConfig, bool)`

GetDatumConfigsOk returns a tuple with the DatumConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatumConfigs

`func (o *BodyConfig) SetDatumConfigs(v map[string]DatumConfig)`

SetDatumConfigs sets DatumConfigs field to given value.

### HasDatumConfigs

`func (o *BodyConfig) HasDatumConfigs() bool`

HasDatumConfigs returns a boolean if a field has been set.

### GetKeys

`func (o *BodyConfig) GetKeys() []string`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *BodyConfig) GetKeysOk() (*[]string, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *BodyConfig) SetKeys(v []string)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *BodyConfig) HasKeys() bool`

HasKeys returns a boolean if a field has been set.

### GetSeparator

`func (o *BodyConfig) GetSeparator() string`

GetSeparator returns the Separator field if non-nil, zero value otherwise.

### GetSeparatorOk

`func (o *BodyConfig) GetSeparatorOk() (*string, bool)`

GetSeparatorOk returns a tuple with the Separator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeparator

`func (o *BodyConfig) SetSeparator(v string)`

SetSeparator sets Separator field to given value.

### HasSeparator

`func (o *BodyConfig) HasSeparator() bool`

HasSeparator returns a boolean if a field has been set.

### GetHide

`func (o *BodyConfig) GetHide() bool`

GetHide returns the Hide field if non-nil, zero value otherwise.

### GetHideOk

`func (o *BodyConfig) GetHideOk() (*bool, bool)`

GetHideOk returns a tuple with the Hide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHide

`func (o *BodyConfig) SetHide(v bool)`

SetHide sets Hide field to given value.

### HasHide

`func (o *BodyConfig) HasHide() bool`

HasHide returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


