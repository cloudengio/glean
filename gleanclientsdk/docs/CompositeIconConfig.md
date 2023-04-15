# CompositeIconConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | Pointer to [**[]DatumCondition**](DatumCondition.md) |  | [optional] 
**OverlayColor** | Pointer to **string** |  | [optional] 
**OverlaySizeRatio** | Pointer to **float32** | The size of the secondary icon overlay, expressed as a fraction of the primary icon size. Value should be less than or equal to 0.75, the default is 0.75. | [optional] 
**PreferSecondaryForNoOverlay** | Pointer to **bool** | If true, the secondary config will be used in cases where an icon overlay can&#39;t be displayed, e.g. for small icons. | [optional] 
**Primary** | Pointer to [**IconConfig**](IconConfig.md) |  | [optional] 
**Secondary** | Pointer to [**IconConfig**](IconConfig.md) |  | [optional] 
**SecondarySizeRatio** | Pointer to **float32** | The size of the secondary icon, expressed as a fraction of the primary icon size. For example, if the primary icon is displayed at 40px and secondarySize is 0.5, the secondary icon will be 20px. Should be less than or equal to overlaySize if both are specified. Default is 0.5. | [optional] 

## Methods

### NewCompositeIconConfig

`func NewCompositeIconConfig() *CompositeIconConfig`

NewCompositeIconConfig instantiates a new CompositeIconConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompositeIconConfigWithDefaults

`func NewCompositeIconConfigWithDefaults() *CompositeIconConfig`

NewCompositeIconConfigWithDefaults instantiates a new CompositeIconConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditions

`func (o *CompositeIconConfig) GetConditions() []DatumCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *CompositeIconConfig) GetConditionsOk() (*[]DatumCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *CompositeIconConfig) SetConditions(v []DatumCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *CompositeIconConfig) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetOverlayColor

`func (o *CompositeIconConfig) GetOverlayColor() string`

GetOverlayColor returns the OverlayColor field if non-nil, zero value otherwise.

### GetOverlayColorOk

`func (o *CompositeIconConfig) GetOverlayColorOk() (*string, bool)`

GetOverlayColorOk returns a tuple with the OverlayColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverlayColor

`func (o *CompositeIconConfig) SetOverlayColor(v string)`

SetOverlayColor sets OverlayColor field to given value.

### HasOverlayColor

`func (o *CompositeIconConfig) HasOverlayColor() bool`

HasOverlayColor returns a boolean if a field has been set.

### GetOverlaySizeRatio

`func (o *CompositeIconConfig) GetOverlaySizeRatio() float32`

GetOverlaySizeRatio returns the OverlaySizeRatio field if non-nil, zero value otherwise.

### GetOverlaySizeRatioOk

`func (o *CompositeIconConfig) GetOverlaySizeRatioOk() (*float32, bool)`

GetOverlaySizeRatioOk returns a tuple with the OverlaySizeRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverlaySizeRatio

`func (o *CompositeIconConfig) SetOverlaySizeRatio(v float32)`

SetOverlaySizeRatio sets OverlaySizeRatio field to given value.

### HasOverlaySizeRatio

`func (o *CompositeIconConfig) HasOverlaySizeRatio() bool`

HasOverlaySizeRatio returns a boolean if a field has been set.

### GetPreferSecondaryForNoOverlay

`func (o *CompositeIconConfig) GetPreferSecondaryForNoOverlay() bool`

GetPreferSecondaryForNoOverlay returns the PreferSecondaryForNoOverlay field if non-nil, zero value otherwise.

### GetPreferSecondaryForNoOverlayOk

`func (o *CompositeIconConfig) GetPreferSecondaryForNoOverlayOk() (*bool, bool)`

GetPreferSecondaryForNoOverlayOk returns a tuple with the PreferSecondaryForNoOverlay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferSecondaryForNoOverlay

`func (o *CompositeIconConfig) SetPreferSecondaryForNoOverlay(v bool)`

SetPreferSecondaryForNoOverlay sets PreferSecondaryForNoOverlay field to given value.

### HasPreferSecondaryForNoOverlay

`func (o *CompositeIconConfig) HasPreferSecondaryForNoOverlay() bool`

HasPreferSecondaryForNoOverlay returns a boolean if a field has been set.

### GetPrimary

`func (o *CompositeIconConfig) GetPrimary() IconConfig`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *CompositeIconConfig) GetPrimaryOk() (*IconConfig, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *CompositeIconConfig) SetPrimary(v IconConfig)`

SetPrimary sets Primary field to given value.

### HasPrimary

`func (o *CompositeIconConfig) HasPrimary() bool`

HasPrimary returns a boolean if a field has been set.

### GetSecondary

`func (o *CompositeIconConfig) GetSecondary() IconConfig`

GetSecondary returns the Secondary field if non-nil, zero value otherwise.

### GetSecondaryOk

`func (o *CompositeIconConfig) GetSecondaryOk() (*IconConfig, bool)`

GetSecondaryOk returns a tuple with the Secondary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondary

`func (o *CompositeIconConfig) SetSecondary(v IconConfig)`

SetSecondary sets Secondary field to given value.

### HasSecondary

`func (o *CompositeIconConfig) HasSecondary() bool`

HasSecondary returns a boolean if a field has been set.

### GetSecondarySizeRatio

`func (o *CompositeIconConfig) GetSecondarySizeRatio() float32`

GetSecondarySizeRatio returns the SecondarySizeRatio field if non-nil, zero value otherwise.

### GetSecondarySizeRatioOk

`func (o *CompositeIconConfig) GetSecondarySizeRatioOk() (*float32, bool)`

GetSecondarySizeRatioOk returns a tuple with the SecondarySizeRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondarySizeRatio

`func (o *CompositeIconConfig) SetSecondarySizeRatio(v float32)`

SetSecondarySizeRatio sets SecondarySizeRatio field to given value.

### HasSecondarySizeRatio

`func (o *CompositeIconConfig) HasSecondarySizeRatio() bool`

HasSecondarySizeRatio returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


