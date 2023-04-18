# MetadataListConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatumConfigs** | Pointer to [**map[string]DatumConfig**](DatumConfig.md) |  | [optional] 
**Keys** | Pointer to **[]string** |  | [optional] 
**Separator** | Pointer to **string** |  | [optional] 

## Methods

### NewMetadataListConfig

`func NewMetadataListConfig() *MetadataListConfig`

NewMetadataListConfig instantiates a new MetadataListConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataListConfigWithDefaults

`func NewMetadataListConfigWithDefaults() *MetadataListConfig`

NewMetadataListConfigWithDefaults instantiates a new MetadataListConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatumConfigs

`func (o *MetadataListConfig) GetDatumConfigs() map[string]DatumConfig`

GetDatumConfigs returns the DatumConfigs field if non-nil, zero value otherwise.

### GetDatumConfigsOk

`func (o *MetadataListConfig) GetDatumConfigsOk() (*map[string]DatumConfig, bool)`

GetDatumConfigsOk returns a tuple with the DatumConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatumConfigs

`func (o *MetadataListConfig) SetDatumConfigs(v map[string]DatumConfig)`

SetDatumConfigs sets DatumConfigs field to given value.

### HasDatumConfigs

`func (o *MetadataListConfig) HasDatumConfigs() bool`

HasDatumConfigs returns a boolean if a field has been set.

### GetKeys

`func (o *MetadataListConfig) GetKeys() []string`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *MetadataListConfig) GetKeysOk() (*[]string, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *MetadataListConfig) SetKeys(v []string)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *MetadataListConfig) HasKeys() bool`

HasKeys returns a boolean if a field has been set.

### GetSeparator

`func (o *MetadataListConfig) GetSeparator() string`

GetSeparator returns the Separator field if non-nil, zero value otherwise.

### GetSeparatorOk

`func (o *MetadataListConfig) GetSeparatorOk() (*string, bool)`

GetSeparatorOk returns a tuple with the Separator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeparator

`func (o *MetadataListConfig) SetSeparator(v string)`

SetSeparator sets Separator field to given value.

### HasSeparator

`func (o *MetadataListConfig) HasSeparator() bool`

HasSeparator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


