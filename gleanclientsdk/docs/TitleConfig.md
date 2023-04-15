# TitleConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatumConfigs** | Pointer to [**map[string]DatumConfig**](DatumConfig.md) |  | [optional] 
**Keys** | Pointer to **[]string** |  | [optional] 
**Separator** | Pointer to **string** |  | [optional] 
**Popover** | Pointer to [**PopoverConfig**](PopoverConfig.md) |  | [optional] 

## Methods

### NewTitleConfig

`func NewTitleConfig() *TitleConfig`

NewTitleConfig instantiates a new TitleConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTitleConfigWithDefaults

`func NewTitleConfigWithDefaults() *TitleConfig`

NewTitleConfigWithDefaults instantiates a new TitleConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatumConfigs

`func (o *TitleConfig) GetDatumConfigs() map[string]DatumConfig`

GetDatumConfigs returns the DatumConfigs field if non-nil, zero value otherwise.

### GetDatumConfigsOk

`func (o *TitleConfig) GetDatumConfigsOk() (*map[string]DatumConfig, bool)`

GetDatumConfigsOk returns a tuple with the DatumConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatumConfigs

`func (o *TitleConfig) SetDatumConfigs(v map[string]DatumConfig)`

SetDatumConfigs sets DatumConfigs field to given value.

### HasDatumConfigs

`func (o *TitleConfig) HasDatumConfigs() bool`

HasDatumConfigs returns a boolean if a field has been set.

### GetKeys

`func (o *TitleConfig) GetKeys() []string`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *TitleConfig) GetKeysOk() (*[]string, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *TitleConfig) SetKeys(v []string)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *TitleConfig) HasKeys() bool`

HasKeys returns a boolean if a field has been set.

### GetSeparator

`func (o *TitleConfig) GetSeparator() string`

GetSeparator returns the Separator field if non-nil, zero value otherwise.

### GetSeparatorOk

`func (o *TitleConfig) GetSeparatorOk() (*string, bool)`

GetSeparatorOk returns a tuple with the Separator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeparator

`func (o *TitleConfig) SetSeparator(v string)`

SetSeparator sets Separator field to given value.

### HasSeparator

`func (o *TitleConfig) HasSeparator() bool`

HasSeparator returns a boolean if a field has been set.

### GetPopover

`func (o *TitleConfig) GetPopover() PopoverConfig`

GetPopover returns the Popover field if non-nil, zero value otherwise.

### GetPopoverOk

`func (o *TitleConfig) GetPopoverOk() (*PopoverConfig, bool)`

GetPopoverOk returns a tuple with the Popover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopover

`func (o *TitleConfig) SetPopover(v PopoverConfig)`

SetPopover sets Popover field to given value.

### HasPopover

`func (o *TitleConfig) HasPopover() bool`

HasPopover returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


