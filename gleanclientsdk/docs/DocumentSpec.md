# DocumentSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | The url for document. | 
**Id** | **string** | The id for document. | 
**UgcType** | **string** | The type of the user generated content (UGC datasource). | 
**ContentId** | **int32** | The id for user generated content. | 
**DocType** | Pointer to **string** | The specific type of the user generated content type. | [optional] 

## Methods

### NewDocumentSpec

`func NewDocumentSpec(url string, id string, ugcType string, contentId int32, ) *DocumentSpec`

NewDocumentSpec instantiates a new DocumentSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentSpecWithDefaults

`func NewDocumentSpecWithDefaults() *DocumentSpec`

NewDocumentSpecWithDefaults instantiates a new DocumentSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *DocumentSpec) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DocumentSpec) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DocumentSpec) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetId

`func (o *DocumentSpec) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DocumentSpec) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DocumentSpec) SetId(v string)`

SetId sets Id field to given value.


### GetUgcType

`func (o *DocumentSpec) GetUgcType() string`

GetUgcType returns the UgcType field if non-nil, zero value otherwise.

### GetUgcTypeOk

`func (o *DocumentSpec) GetUgcTypeOk() (*string, bool)`

GetUgcTypeOk returns a tuple with the UgcType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUgcType

`func (o *DocumentSpec) SetUgcType(v string)`

SetUgcType sets UgcType field to given value.


### GetContentId

`func (o *DocumentSpec) GetContentId() int32`

GetContentId returns the ContentId field if non-nil, zero value otherwise.

### GetContentIdOk

`func (o *DocumentSpec) GetContentIdOk() (*int32, bool)`

GetContentIdOk returns a tuple with the ContentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentId

`func (o *DocumentSpec) SetContentId(v int32)`

SetContentId sets ContentId field to given value.


### GetDocType

`func (o *DocumentSpec) GetDocType() string`

GetDocType returns the DocType field if non-nil, zero value otherwise.

### GetDocTypeOk

`func (o *DocumentSpec) GetDocTypeOk() (*string, bool)`

GetDocTypeOk returns a tuple with the DocType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocType

`func (o *DocumentSpec) SetDocType(v string)`

SetDocType sets DocType field to given value.

### HasDocType

`func (o *DocumentSpec) HasDocType() bool`

HasDocType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


