# ContentDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MimeType** | **string** |  | 
**TextContent** | Pointer to **string** | text content. Only one of textContent or binary content can be specified | [optional] 
**BinaryContent** | Pointer to **string** | base64 encoded binary content. Only one of textContent or binary content can be specified | [optional] 

## Methods

### NewContentDefinition

`func NewContentDefinition(mimeType string, ) *ContentDefinition`

NewContentDefinition instantiates a new ContentDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentDefinitionWithDefaults

`func NewContentDefinitionWithDefaults() *ContentDefinition`

NewContentDefinitionWithDefaults instantiates a new ContentDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMimeType

`func (o *ContentDefinition) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *ContentDefinition) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *ContentDefinition) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.


### GetTextContent

`func (o *ContentDefinition) GetTextContent() string`

GetTextContent returns the TextContent field if non-nil, zero value otherwise.

### GetTextContentOk

`func (o *ContentDefinition) GetTextContentOk() (*string, bool)`

GetTextContentOk returns a tuple with the TextContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextContent

`func (o *ContentDefinition) SetTextContent(v string)`

SetTextContent sets TextContent field to given value.

### HasTextContent

`func (o *ContentDefinition) HasTextContent() bool`

HasTextContent returns a boolean if a field has been set.

### GetBinaryContent

`func (o *ContentDefinition) GetBinaryContent() string`

GetBinaryContent returns the BinaryContent field if non-nil, zero value otherwise.

### GetBinaryContentOk

`func (o *ContentDefinition) GetBinaryContentOk() (*string, bool)`

GetBinaryContentOk returns a tuple with the BinaryContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryContent

`func (o *ContentDefinition) SetBinaryContent(v string)`

SetBinaryContent sets BinaryContent field to given value.

### HasBinaryContent

`func (o *ContentDefinition) HasBinaryContent() bool`

HasBinaryContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


