# RenderConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to [**ActionsConfig**](ActionsConfig.md) |  | [optional] 
**AdditionalConfigs** | Pointer to **[]string** |  | [optional] 
**Body** | Pointer to [**BodyConfig**](BodyConfig.md) |  | [optional] 
**Contact** | Pointer to [**ContactConfig**](ContactConfig.md) |  | [optional] 
**Details** | Pointer to [**DetailsConfig**](DetailsConfig.md) |  | [optional] 
**FilePath** | Pointer to [**FilePathConfig**](FilePathConfig.md) |  | [optional] 
**Footer** | Pointer to [**RenderConfigFooter**](RenderConfigFooter.md) |  | [optional] 
**Related** | Pointer to [**RelatedConfig**](RelatedConfig.md) |  | [optional] 
**Icon** | Pointer to [**[]CompositeIconConfig**](CompositeIconConfig.md) |  | [optional] 
**Interactions** | Pointer to [**InteractionsConfig**](InteractionsConfig.md) |  | [optional] 
**Meta** | Pointer to [**MetaConfig**](MetaConfig.md) |  | [optional] 
**Preview** | Pointer to [**PreviewConfig**](PreviewConfig.md) |  | [optional] 
**Tags** | Pointer to [**TagsConfig**](TagsConfig.md) |  | [optional] 
**Title** | Pointer to [**TitleConfig**](TitleConfig.md) |  | [optional] 
**Sections** | Pointer to [**SectionsConfig**](SectionsConfig.md) |  | [optional] 

## Methods

### NewRenderConfig

`func NewRenderConfig() *RenderConfig`

NewRenderConfig instantiates a new RenderConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRenderConfigWithDefaults

`func NewRenderConfigWithDefaults() *RenderConfig`

NewRenderConfigWithDefaults instantiates a new RenderConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *RenderConfig) GetActions() ActionsConfig`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *RenderConfig) GetActionsOk() (*ActionsConfig, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *RenderConfig) SetActions(v ActionsConfig)`

SetActions sets Actions field to given value.

### HasActions

`func (o *RenderConfig) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetAdditionalConfigs

`func (o *RenderConfig) GetAdditionalConfigs() []string`

GetAdditionalConfigs returns the AdditionalConfigs field if non-nil, zero value otherwise.

### GetAdditionalConfigsOk

`func (o *RenderConfig) GetAdditionalConfigsOk() (*[]string, bool)`

GetAdditionalConfigsOk returns a tuple with the AdditionalConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalConfigs

`func (o *RenderConfig) SetAdditionalConfigs(v []string)`

SetAdditionalConfigs sets AdditionalConfigs field to given value.

### HasAdditionalConfigs

`func (o *RenderConfig) HasAdditionalConfigs() bool`

HasAdditionalConfigs returns a boolean if a field has been set.

### GetBody

`func (o *RenderConfig) GetBody() BodyConfig`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *RenderConfig) GetBodyOk() (*BodyConfig, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *RenderConfig) SetBody(v BodyConfig)`

SetBody sets Body field to given value.

### HasBody

`func (o *RenderConfig) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetContact

`func (o *RenderConfig) GetContact() ContactConfig`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *RenderConfig) GetContactOk() (*ContactConfig, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *RenderConfig) SetContact(v ContactConfig)`

SetContact sets Contact field to given value.

### HasContact

`func (o *RenderConfig) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetDetails

`func (o *RenderConfig) GetDetails() DetailsConfig`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *RenderConfig) GetDetailsOk() (*DetailsConfig, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *RenderConfig) SetDetails(v DetailsConfig)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *RenderConfig) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetFilePath

`func (o *RenderConfig) GetFilePath() FilePathConfig`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *RenderConfig) GetFilePathOk() (*FilePathConfig, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *RenderConfig) SetFilePath(v FilePathConfig)`

SetFilePath sets FilePath field to given value.

### HasFilePath

`func (o *RenderConfig) HasFilePath() bool`

HasFilePath returns a boolean if a field has been set.

### GetFooter

`func (o *RenderConfig) GetFooter() RenderConfigFooter`

GetFooter returns the Footer field if non-nil, zero value otherwise.

### GetFooterOk

`func (o *RenderConfig) GetFooterOk() (*RenderConfigFooter, bool)`

GetFooterOk returns a tuple with the Footer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFooter

`func (o *RenderConfig) SetFooter(v RenderConfigFooter)`

SetFooter sets Footer field to given value.

### HasFooter

`func (o *RenderConfig) HasFooter() bool`

HasFooter returns a boolean if a field has been set.

### GetRelated

`func (o *RenderConfig) GetRelated() RelatedConfig`

GetRelated returns the Related field if non-nil, zero value otherwise.

### GetRelatedOk

`func (o *RenderConfig) GetRelatedOk() (*RelatedConfig, bool)`

GetRelatedOk returns a tuple with the Related field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelated

`func (o *RenderConfig) SetRelated(v RelatedConfig)`

SetRelated sets Related field to given value.

### HasRelated

`func (o *RenderConfig) HasRelated() bool`

HasRelated returns a boolean if a field has been set.

### GetIcon

`func (o *RenderConfig) GetIcon() []CompositeIconConfig`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *RenderConfig) GetIconOk() (*[]CompositeIconConfig, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *RenderConfig) SetIcon(v []CompositeIconConfig)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *RenderConfig) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetInteractions

`func (o *RenderConfig) GetInteractions() InteractionsConfig`

GetInteractions returns the Interactions field if non-nil, zero value otherwise.

### GetInteractionsOk

`func (o *RenderConfig) GetInteractionsOk() (*InteractionsConfig, bool)`

GetInteractionsOk returns a tuple with the Interactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInteractions

`func (o *RenderConfig) SetInteractions(v InteractionsConfig)`

SetInteractions sets Interactions field to given value.

### HasInteractions

`func (o *RenderConfig) HasInteractions() bool`

HasInteractions returns a boolean if a field has been set.

### GetMeta

`func (o *RenderConfig) GetMeta() MetaConfig`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *RenderConfig) GetMetaOk() (*MetaConfig, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *RenderConfig) SetMeta(v MetaConfig)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *RenderConfig) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetPreview

`func (o *RenderConfig) GetPreview() PreviewConfig`

GetPreview returns the Preview field if non-nil, zero value otherwise.

### GetPreviewOk

`func (o *RenderConfig) GetPreviewOk() (*PreviewConfig, bool)`

GetPreviewOk returns a tuple with the Preview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreview

`func (o *RenderConfig) SetPreview(v PreviewConfig)`

SetPreview sets Preview field to given value.

### HasPreview

`func (o *RenderConfig) HasPreview() bool`

HasPreview returns a boolean if a field has been set.

### GetTags

`func (o *RenderConfig) GetTags() TagsConfig`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RenderConfig) GetTagsOk() (*TagsConfig, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *RenderConfig) SetTags(v TagsConfig)`

SetTags sets Tags field to given value.

### HasTags

`func (o *RenderConfig) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTitle

`func (o *RenderConfig) GetTitle() TitleConfig`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *RenderConfig) GetTitleOk() (*TitleConfig, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *RenderConfig) SetTitle(v TitleConfig)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *RenderConfig) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetSections

`func (o *RenderConfig) GetSections() SectionsConfig`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *RenderConfig) GetSectionsOk() (*SectionsConfig, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *RenderConfig) SetSections(v SectionsConfig)`

SetSections sets Sections field to given value.

### HasSections

`func (o *RenderConfig) HasSections() bool`

HasSections returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


