# MetaConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatumConfigs** | Pointer to [**map[string]DatumConfig**](DatumConfig.md) |  | [optional] 
**Keys** | Pointer to **[]string** |  | [optional] 
**Separator** | Pointer to **string** |  | [optional] 
**HideInAttachment** | Pointer to **bool** |  | [optional] 
**Label** | Pointer to [**StatusLabelConfig**](StatusLabelConfig.md) |  | [optional] 
**CommonKeys** | Pointer to **[]string** |  | [optional] 

## Methods

### NewMetaConfig

`func NewMetaConfig() *MetaConfig`

NewMetaConfig instantiates a new MetaConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaConfigWithDefaults

`func NewMetaConfigWithDefaults() *MetaConfig`

NewMetaConfigWithDefaults instantiates a new MetaConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatumConfigs

`func (o *MetaConfig) GetDatumConfigs() map[string]DatumConfig`

GetDatumConfigs returns the DatumConfigs field if non-nil, zero value otherwise.

### GetDatumConfigsOk

`func (o *MetaConfig) GetDatumConfigsOk() (*map[string]DatumConfig, bool)`

GetDatumConfigsOk returns a tuple with the DatumConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatumConfigs

`func (o *MetaConfig) SetDatumConfigs(v map[string]DatumConfig)`

SetDatumConfigs sets DatumConfigs field to given value.

### HasDatumConfigs

`func (o *MetaConfig) HasDatumConfigs() bool`

HasDatumConfigs returns a boolean if a field has been set.

### GetKeys

`func (o *MetaConfig) GetKeys() []string`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *MetaConfig) GetKeysOk() (*[]string, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *MetaConfig) SetKeys(v []string)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *MetaConfig) HasKeys() bool`

HasKeys returns a boolean if a field has been set.

### GetSeparator

`func (o *MetaConfig) GetSeparator() string`

GetSeparator returns the Separator field if non-nil, zero value otherwise.

### GetSeparatorOk

`func (o *MetaConfig) GetSeparatorOk() (*string, bool)`

GetSeparatorOk returns a tuple with the Separator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeparator

`func (o *MetaConfig) SetSeparator(v string)`

SetSeparator sets Separator field to given value.

### HasSeparator

`func (o *MetaConfig) HasSeparator() bool`

HasSeparator returns a boolean if a field has been set.

### GetHideInAttachment

`func (o *MetaConfig) GetHideInAttachment() bool`

GetHideInAttachment returns the HideInAttachment field if non-nil, zero value otherwise.

### GetHideInAttachmentOk

`func (o *MetaConfig) GetHideInAttachmentOk() (*bool, bool)`

GetHideInAttachmentOk returns a tuple with the HideInAttachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideInAttachment

`func (o *MetaConfig) SetHideInAttachment(v bool)`

SetHideInAttachment sets HideInAttachment field to given value.

### HasHideInAttachment

`func (o *MetaConfig) HasHideInAttachment() bool`

HasHideInAttachment returns a boolean if a field has been set.

### GetLabel

`func (o *MetaConfig) GetLabel() StatusLabelConfig`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *MetaConfig) GetLabelOk() (*StatusLabelConfig, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *MetaConfig) SetLabel(v StatusLabelConfig)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *MetaConfig) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetCommonKeys

`func (o *MetaConfig) GetCommonKeys() []string`

GetCommonKeys returns the CommonKeys field if non-nil, zero value otherwise.

### GetCommonKeysOk

`func (o *MetaConfig) GetCommonKeysOk() (*[]string, bool)`

GetCommonKeysOk returns a tuple with the CommonKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonKeys

`func (o *MetaConfig) SetCommonKeys(v []string)`

SetCommonKeys sets CommonKeys field to given value.

### HasCommonKeys

`func (o *MetaConfig) HasCommonKeys() bool`

HasCommonKeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


