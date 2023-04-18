# RelatedConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SectionConfigs** | Pointer to [**map[string]RelatedSectionConfig**](RelatedSectionConfig.md) |  | [optional] 
**Keys** | Pointer to **[]string** |  | [optional] 

## Methods

### NewRelatedConfig

`func NewRelatedConfig() *RelatedConfig`

NewRelatedConfig instantiates a new RelatedConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelatedConfigWithDefaults

`func NewRelatedConfigWithDefaults() *RelatedConfig`

NewRelatedConfigWithDefaults instantiates a new RelatedConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSectionConfigs

`func (o *RelatedConfig) GetSectionConfigs() map[string]RelatedSectionConfig`

GetSectionConfigs returns the SectionConfigs field if non-nil, zero value otherwise.

### GetSectionConfigsOk

`func (o *RelatedConfig) GetSectionConfigsOk() (*map[string]RelatedSectionConfig, bool)`

GetSectionConfigsOk returns a tuple with the SectionConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionConfigs

`func (o *RelatedConfig) SetSectionConfigs(v map[string]RelatedSectionConfig)`

SetSectionConfigs sets SectionConfigs field to given value.

### HasSectionConfigs

`func (o *RelatedConfig) HasSectionConfigs() bool`

HasSectionConfigs returns a boolean if a field has been set.

### GetKeys

`func (o *RelatedConfig) GetKeys() []string`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *RelatedConfig) GetKeysOk() (*[]string, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *RelatedConfig) SetKeys(v []string)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *RelatedConfig) HasKeys() bool`

HasKeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


