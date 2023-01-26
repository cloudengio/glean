# DeleteDocumentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **int64** | Version number for document for optimistic concurrency control. If absent or 0 then no version checks are done. | [optional] 
**Datasource** | **string** | datasource of the document | 
**ObjectType** | **string** | object type of the document | 
**Id** | **string** | The id of the document | 

## Methods

### NewDeleteDocumentRequest

`func NewDeleteDocumentRequest(datasource string, objectType string, id string, ) *DeleteDocumentRequest`

NewDeleteDocumentRequest instantiates a new DeleteDocumentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteDocumentRequestWithDefaults

`func NewDeleteDocumentRequestWithDefaults() *DeleteDocumentRequest`

NewDeleteDocumentRequestWithDefaults instantiates a new DeleteDocumentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *DeleteDocumentRequest) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DeleteDocumentRequest) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DeleteDocumentRequest) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DeleteDocumentRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetDatasource

`func (o *DeleteDocumentRequest) GetDatasource() string`

GetDatasource returns the Datasource field if non-nil, zero value otherwise.

### GetDatasourceOk

`func (o *DeleteDocumentRequest) GetDatasourceOk() (*string, bool)`

GetDatasourceOk returns a tuple with the Datasource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasource

`func (o *DeleteDocumentRequest) SetDatasource(v string)`

SetDatasource sets Datasource field to given value.


### GetObjectType

`func (o *DeleteDocumentRequest) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *DeleteDocumentRequest) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *DeleteDocumentRequest) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetId

`func (o *DeleteDocumentRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeleteDocumentRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeleteDocumentRequest) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


