# FooterConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attachment** | Pointer to [**RelatedDocumentsConfig**](RelatedDocumentsConfig.md) |  | [optional] 
**FilePath** | Pointer to **bool** | Use RenderConfig.filePath instead | [optional] 

## Methods

### NewFooterConfig

`func NewFooterConfig() *FooterConfig`

NewFooterConfig instantiates a new FooterConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFooterConfigWithDefaults

`func NewFooterConfigWithDefaults() *FooterConfig`

NewFooterConfigWithDefaults instantiates a new FooterConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachment

`func (o *FooterConfig) GetAttachment() RelatedDocumentsConfig`

GetAttachment returns the Attachment field if non-nil, zero value otherwise.

### GetAttachmentOk

`func (o *FooterConfig) GetAttachmentOk() (*RelatedDocumentsConfig, bool)`

GetAttachmentOk returns a tuple with the Attachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachment

`func (o *FooterConfig) SetAttachment(v RelatedDocumentsConfig)`

SetAttachment sets Attachment field to given value.

### HasAttachment

`func (o *FooterConfig) HasAttachment() bool`

HasAttachment returns a boolean if a field has been set.

### GetFilePath

`func (o *FooterConfig) GetFilePath() bool`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *FooterConfig) GetFilePathOk() (*bool, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *FooterConfig) SetFilePath(v bool)`

SetFilePath sets FilePath field to given value.

### HasFilePath

`func (o *FooterConfig) HasFilePath() bool`

HasFilePath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


