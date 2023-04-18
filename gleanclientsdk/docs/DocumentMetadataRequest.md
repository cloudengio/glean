# DocumentMetadataRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | **[]string** | The IDs of the documents for which metadata is requested | 
**DocumentFields** | Pointer to **[]string** | The requested metadata fields | [optional] 

## Methods

### NewDocumentMetadataRequest

`func NewDocumentMetadataRequest(ids []string, ) *DocumentMetadataRequest`

NewDocumentMetadataRequest instantiates a new DocumentMetadataRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentMetadataRequestWithDefaults

`func NewDocumentMetadataRequestWithDefaults() *DocumentMetadataRequest`

NewDocumentMetadataRequestWithDefaults instantiates a new DocumentMetadataRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *DocumentMetadataRequest) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *DocumentMetadataRequest) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *DocumentMetadataRequest) SetIds(v []string)`

SetIds sets Ids field to given value.


### GetDocumentFields

`func (o *DocumentMetadataRequest) GetDocumentFields() []string`

GetDocumentFields returns the DocumentFields field if non-nil, zero value otherwise.

### GetDocumentFieldsOk

`func (o *DocumentMetadataRequest) GetDocumentFieldsOk() (*[]string, bool)`

GetDocumentFieldsOk returns a tuple with the DocumentFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentFields

`func (o *DocumentMetadataRequest) SetDocumentFields(v []string)`

SetDocumentFields sets DocumentFields field to given value.

### HasDocumentFields

`func (o *DocumentMetadataRequest) HasDocumentFields() bool`

HasDocumentFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


