# DatumConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | Pointer to [**[]DatumCondition**](DatumCondition.md) |  | [optional] 
**DatumConfigs** | Pointer to [**map[string]DatumConfig**](DatumConfig.md) |  | [optional] 
**Icon** | Pointer to [**IconConfig**](IconConfig.md) |  | [optional] 
**Keys** | Pointer to **[]string** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**Omit** | Pointer to **bool** |  | [optional] 
**Popover** | Pointer to [**PopoverConfig**](PopoverConfig.md) |  | [optional] 
**Prefix** | Pointer to **string** |  | [optional] 
**RoleConfig** | Pointer to [**DatumRoleConfig**](DatumRoleConfig.md) |  | [optional] 
**Separator** | Pointer to **string** |  | [optional] 
**Suffix** | Pointer to **string** |  | [optional] 
**Transforms** | Pointer to [**[]DatumTransform**](DatumTransform.md) |  | [optional] 
**Type** | Pointer to [**DatumType**](DatumType.md) |  | [optional] 

## Methods

### NewDatumConfig

`func NewDatumConfig() *DatumConfig`

NewDatumConfig instantiates a new DatumConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatumConfigWithDefaults

`func NewDatumConfigWithDefaults() *DatumConfig`

NewDatumConfigWithDefaults instantiates a new DatumConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditions

`func (o *DatumConfig) GetConditions() []DatumCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *DatumConfig) GetConditionsOk() (*[]DatumCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *DatumConfig) SetConditions(v []DatumCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *DatumConfig) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetDatumConfigs

`func (o *DatumConfig) GetDatumConfigs() map[string]DatumConfig`

GetDatumConfigs returns the DatumConfigs field if non-nil, zero value otherwise.

### GetDatumConfigsOk

`func (o *DatumConfig) GetDatumConfigsOk() (*map[string]DatumConfig, bool)`

GetDatumConfigsOk returns a tuple with the DatumConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatumConfigs

`func (o *DatumConfig) SetDatumConfigs(v map[string]DatumConfig)`

SetDatumConfigs sets DatumConfigs field to given value.

### HasDatumConfigs

`func (o *DatumConfig) HasDatumConfigs() bool`

HasDatumConfigs returns a boolean if a field has been set.

### GetIcon

`func (o *DatumConfig) GetIcon() IconConfig`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *DatumConfig) GetIconOk() (*IconConfig, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *DatumConfig) SetIcon(v IconConfig)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *DatumConfig) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetKeys

`func (o *DatumConfig) GetKeys() []string`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *DatumConfig) GetKeysOk() (*[]string, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *DatumConfig) SetKeys(v []string)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *DatumConfig) HasKeys() bool`

HasKeys returns a boolean if a field has been set.

### GetLabel

`func (o *DatumConfig) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *DatumConfig) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *DatumConfig) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *DatumConfig) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetOmit

`func (o *DatumConfig) GetOmit() bool`

GetOmit returns the Omit field if non-nil, zero value otherwise.

### GetOmitOk

`func (o *DatumConfig) GetOmitOk() (*bool, bool)`

GetOmitOk returns a tuple with the Omit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmit

`func (o *DatumConfig) SetOmit(v bool)`

SetOmit sets Omit field to given value.

### HasOmit

`func (o *DatumConfig) HasOmit() bool`

HasOmit returns a boolean if a field has been set.

### GetPopover

`func (o *DatumConfig) GetPopover() PopoverConfig`

GetPopover returns the Popover field if non-nil, zero value otherwise.

### GetPopoverOk

`func (o *DatumConfig) GetPopoverOk() (*PopoverConfig, bool)`

GetPopoverOk returns a tuple with the Popover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopover

`func (o *DatumConfig) SetPopover(v PopoverConfig)`

SetPopover sets Popover field to given value.

### HasPopover

`func (o *DatumConfig) HasPopover() bool`

HasPopover returns a boolean if a field has been set.

### GetPrefix

`func (o *DatumConfig) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *DatumConfig) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *DatumConfig) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *DatumConfig) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetRoleConfig

`func (o *DatumConfig) GetRoleConfig() DatumRoleConfig`

GetRoleConfig returns the RoleConfig field if non-nil, zero value otherwise.

### GetRoleConfigOk

`func (o *DatumConfig) GetRoleConfigOk() (*DatumRoleConfig, bool)`

GetRoleConfigOk returns a tuple with the RoleConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleConfig

`func (o *DatumConfig) SetRoleConfig(v DatumRoleConfig)`

SetRoleConfig sets RoleConfig field to given value.

### HasRoleConfig

`func (o *DatumConfig) HasRoleConfig() bool`

HasRoleConfig returns a boolean if a field has been set.

### GetSeparator

`func (o *DatumConfig) GetSeparator() string`

GetSeparator returns the Separator field if non-nil, zero value otherwise.

### GetSeparatorOk

`func (o *DatumConfig) GetSeparatorOk() (*string, bool)`

GetSeparatorOk returns a tuple with the Separator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeparator

`func (o *DatumConfig) SetSeparator(v string)`

SetSeparator sets Separator field to given value.

### HasSeparator

`func (o *DatumConfig) HasSeparator() bool`

HasSeparator returns a boolean if a field has been set.

### GetSuffix

`func (o *DatumConfig) GetSuffix() string`

GetSuffix returns the Suffix field if non-nil, zero value otherwise.

### GetSuffixOk

`func (o *DatumConfig) GetSuffixOk() (*string, bool)`

GetSuffixOk returns a tuple with the Suffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuffix

`func (o *DatumConfig) SetSuffix(v string)`

SetSuffix sets Suffix field to given value.

### HasSuffix

`func (o *DatumConfig) HasSuffix() bool`

HasSuffix returns a boolean if a field has been set.

### GetTransforms

`func (o *DatumConfig) GetTransforms() []DatumTransform`

GetTransforms returns the Transforms field if non-nil, zero value otherwise.

### GetTransformsOk

`func (o *DatumConfig) GetTransformsOk() (*[]DatumTransform, bool)`

GetTransformsOk returns a tuple with the Transforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransforms

`func (o *DatumConfig) SetTransforms(v []DatumTransform)`

SetTransforms sets Transforms field to given value.

### HasTransforms

`func (o *DatumConfig) HasTransforms() bool`

HasTransforms returns a boolean if a field has been set.

### GetType

`func (o *DatumConfig) GetType() DatumType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DatumConfig) GetTypeOk() (*DatumType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DatumConfig) SetType(v DatumType)`

SetType sets Type field to given value.

### HasType

`func (o *DatumConfig) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


