# GetDocumentStatusRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datasource** | **string** | Datasource to get fetch document status for | 
**ObjectType** | **string** | Object type of the document to get the status for | 
**DocId** | **string** | Document ID within the datasource to get the status for | 

## Methods

### NewGetDocumentStatusRequest

`func NewGetDocumentStatusRequest(datasource string, objectType string, docId string, ) *GetDocumentStatusRequest`

NewGetDocumentStatusRequest instantiates a new GetDocumentStatusRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDocumentStatusRequestWithDefaults

`func NewGetDocumentStatusRequestWithDefaults() *GetDocumentStatusRequest`

NewGetDocumentStatusRequestWithDefaults instantiates a new GetDocumentStatusRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatasource

`func (o *GetDocumentStatusRequest) GetDatasource() string`

GetDatasource returns the Datasource field if non-nil, zero value otherwise.

### GetDatasourceOk

`func (o *GetDocumentStatusRequest) GetDatasourceOk() (*string, bool)`

GetDatasourceOk returns a tuple with the Datasource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasource

`func (o *GetDocumentStatusRequest) SetDatasource(v string)`

SetDatasource sets Datasource field to given value.


### GetObjectType

`func (o *GetDocumentStatusRequest) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *GetDocumentStatusRequest) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *GetDocumentStatusRequest) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDocId

`func (o *GetDocumentStatusRequest) GetDocId() string`

GetDocId returns the DocId field if non-nil, zero value otherwise.

### GetDocIdOk

`func (o *GetDocumentStatusRequest) GetDocIdOk() (*string, bool)`

GetDocIdOk returns a tuple with the DocId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocId

`func (o *GetDocumentStatusRequest) SetDocId(v string)`

SetDocId sets DocId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


