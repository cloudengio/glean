# RenderConfigFooter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attachment** | Pointer to [**RelatedDocumentsConfig**](RelatedDocumentsConfig.md) |  | [optional] 
**FilePath** | Pointer to **bool** | Use RenderConfig.filePath instead | [optional] 

## Methods

### NewRenderConfigFooter

`func NewRenderConfigFooter() *RenderConfigFooter`

NewRenderConfigFooter instantiates a new RenderConfigFooter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRenderConfigFooterWithDefaults

`func NewRenderConfigFooterWithDefaults() *RenderConfigFooter`

NewRenderConfigFooterWithDefaults instantiates a new RenderConfigFooter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachment

`func (o *RenderConfigFooter) GetAttachment() RelatedDocumentsConfig`

GetAttachment returns the Attachment field if non-nil, zero value otherwise.

### GetAttachmentOk

`func (o *RenderConfigFooter) GetAttachmentOk() (*RelatedDocumentsConfig, bool)`

GetAttachmentOk returns a tuple with the Attachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachment

`func (o *RenderConfigFooter) SetAttachment(v RelatedDocumentsConfig)`

SetAttachment sets Attachment field to given value.

### HasAttachment

`func (o *RenderConfigFooter) HasAttachment() bool`

HasAttachment returns a boolean if a field has been set.

### GetFilePath

`func (o *RenderConfigFooter) GetFilePath() bool`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *RenderConfigFooter) GetFilePathOk() (*bool, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *RenderConfigFooter) SetFilePath(v bool)`

SetFilePath sets FilePath field to given value.

### HasFilePath

`func (o *RenderConfigFooter) HasFilePath() bool`

HasFilePath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


