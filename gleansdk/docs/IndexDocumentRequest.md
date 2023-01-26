# IndexDocumentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **int64** | Version number for document for optimistic concurrency control. If absent or 0 then no version checks are done. | [optional] 
**Document** | [**DocumentDefinition**](DocumentDefinition.md) |  | 

## Methods

### NewIndexDocumentRequest

`func NewIndexDocumentRequest(document DocumentDefinition, ) *IndexDocumentRequest`

NewIndexDocumentRequest instantiates a new IndexDocumentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndexDocumentRequestWithDefaults

`func NewIndexDocumentRequestWithDefaults() *IndexDocumentRequest`

NewIndexDocumentRequestWithDefaults instantiates a new IndexDocumentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *IndexDocumentRequest) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *IndexDocumentRequest) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *IndexDocumentRequest) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *IndexDocumentRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetDocument

`func (o *IndexDocumentRequest) GetDocument() DocumentDefinition`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *IndexDocumentRequest) GetDocumentOk() (*DocumentDefinition, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *IndexDocumentRequest) SetDocument(v DocumentDefinition)`

SetDocument sets Document field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


