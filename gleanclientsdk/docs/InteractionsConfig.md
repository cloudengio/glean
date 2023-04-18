# InteractionsConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatumConfigs** | Pointer to [**map[string]DatumConfig**](DatumConfig.md) |  | [optional] 
**Keys** | Pointer to **[]string** |  | [optional] 
**Separator** | Pointer to **string** |  | [optional] 
**Show** | Pointer to **bool** |  | [optional] 

## Methods

### NewInteractionsConfig

`func NewInteractionsConfig() *InteractionsConfig`

NewInteractionsConfig instantiates a new InteractionsConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInteractionsConfigWithDefaults

`func NewInteractionsConfigWithDefaults() *InteractionsConfig`

NewInteractionsConfigWithDefaults instantiates a new InteractionsConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatumConfigs

`func (o *InteractionsConfig) GetDatumConfigs() map[string]DatumConfig`

GetDatumConfigs returns the DatumConfigs field if non-nil, zero value otherwise.

### GetDatumConfigsOk

`func (o *InteractionsConfig) GetDatumConfigsOk() (*map[string]DatumConfig, bool)`

GetDatumConfigsOk returns a tuple with the DatumConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatumConfigs

`func (o *InteractionsConfig) SetDatumConfigs(v map[string]DatumConfig)`

SetDatumConfigs sets DatumConfigs field to given value.

### HasDatumConfigs

`func (o *InteractionsConfig) HasDatumConfigs() bool`

HasDatumConfigs returns a boolean if a field has been set.

### GetKeys

`func (o *InteractionsConfig) GetKeys() []string`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *InteractionsConfig) GetKeysOk() (*[]string, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *InteractionsConfig) SetKeys(v []string)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *InteractionsConfig) HasKeys() bool`

HasKeys returns a boolean if a field has been set.

### GetSeparator

`func (o *InteractionsConfig) GetSeparator() string`

GetSeparator returns the Separator field if non-nil, zero value otherwise.

### GetSeparatorOk

`func (o *InteractionsConfig) GetSeparatorOk() (*string, bool)`

GetSeparatorOk returns a tuple with the Separator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeparator

`func (o *InteractionsConfig) SetSeparator(v string)`

SetSeparator sets Separator field to given value.

### HasSeparator

`func (o *InteractionsConfig) HasSeparator() bool`

HasSeparator returns a boolean if a field has been set.

### GetShow

`func (o *InteractionsConfig) GetShow() bool`

GetShow returns the Show field if non-nil, zero value otherwise.

### GetShowOk

`func (o *InteractionsConfig) GetShowOk() (*bool, bool)`

GetShowOk returns a tuple with the Show field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShow

`func (o *InteractionsConfig) SetShow(v bool)`

SetShow sets Show field to given value.

### HasShow

`func (o *InteractionsConfig) HasShow() bool`

HasShow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


