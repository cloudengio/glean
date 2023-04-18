# PreviewConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Icon** | Pointer to [**CompositeIconConfig**](CompositeIconConfig.md) |  | [optional] 
**Meta** | Pointer to [**MetaConfig**](MetaConfig.md) |  | [optional] 
**ModalWidth** | Pointer to **float32** |  | [optional] 
**Title** | Pointer to [**MetadataListConfig**](MetadataListConfig.md) |  | [optional] 
**Type** | Pointer to [**DocumentPreviewType**](DocumentPreviewType.md) |  | [optional] 

## Methods

### NewPreviewConfig

`func NewPreviewConfig() *PreviewConfig`

NewPreviewConfig instantiates a new PreviewConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPreviewConfigWithDefaults

`func NewPreviewConfigWithDefaults() *PreviewConfig`

NewPreviewConfigWithDefaults instantiates a new PreviewConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIcon

`func (o *PreviewConfig) GetIcon() CompositeIconConfig`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *PreviewConfig) GetIconOk() (*CompositeIconConfig, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *PreviewConfig) SetIcon(v CompositeIconConfig)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *PreviewConfig) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetMeta

`func (o *PreviewConfig) GetMeta() MetaConfig`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *PreviewConfig) GetMetaOk() (*MetaConfig, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *PreviewConfig) SetMeta(v MetaConfig)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *PreviewConfig) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetModalWidth

`func (o *PreviewConfig) GetModalWidth() float32`

GetModalWidth returns the ModalWidth field if non-nil, zero value otherwise.

### GetModalWidthOk

`func (o *PreviewConfig) GetModalWidthOk() (*float32, bool)`

GetModalWidthOk returns a tuple with the ModalWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModalWidth

`func (o *PreviewConfig) SetModalWidth(v float32)`

SetModalWidth sets ModalWidth field to given value.

### HasModalWidth

`func (o *PreviewConfig) HasModalWidth() bool`

HasModalWidth returns a boolean if a field has been set.

### GetTitle

`func (o *PreviewConfig) GetTitle() MetadataListConfig`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PreviewConfig) GetTitleOk() (*MetadataListConfig, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PreviewConfig) SetTitle(v MetadataListConfig)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *PreviewConfig) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *PreviewConfig) GetType() DocumentPreviewType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PreviewConfig) GetTypeOk() (*DocumentPreviewType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PreviewConfig) SetType(v DocumentPreviewType)`

SetType sets Type field to given value.

### HasType

`func (o *PreviewConfig) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


