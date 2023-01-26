# BulkIndexDocumentsRequestAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datasource** | **string** | Datasource of the documents | 
**Documents** | [**[]DocumentDefinition**](DocumentDefinition.md) | Batch of documents for the datasource | 
**DisableStaleDocumentDeletionCheck** | Pointer to **bool** | True if older documents need to be force deleted after the upload completes. Defaults to older documents being deleted asynchronously. This must only be set when &#x60;isLastPage &#x3D; true&#x60; | [optional] 

## Methods

### NewBulkIndexDocumentsRequestAllOf

`func NewBulkIndexDocumentsRequestAllOf(datasource string, documents []DocumentDefinition, ) *BulkIndexDocumentsRequestAllOf`

NewBulkIndexDocumentsRequestAllOf instantiates a new BulkIndexDocumentsRequestAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkIndexDocumentsRequestAllOfWithDefaults

`func NewBulkIndexDocumentsRequestAllOfWithDefaults() *BulkIndexDocumentsRequestAllOf`

NewBulkIndexDocumentsRequestAllOfWithDefaults instantiates a new BulkIndexDocumentsRequestAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatasource

`func (o *BulkIndexDocumentsRequestAllOf) GetDatasource() string`

GetDatasource returns the Datasource field if non-nil, zero value otherwise.

### GetDatasourceOk

`func (o *BulkIndexDocumentsRequestAllOf) GetDatasourceOk() (*string, bool)`

GetDatasourceOk returns a tuple with the Datasource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasource

`func (o *BulkIndexDocumentsRequestAllOf) SetDatasource(v string)`

SetDatasource sets Datasource field to given value.


### GetDocuments

`func (o *BulkIndexDocumentsRequestAllOf) GetDocuments() []DocumentDefinition`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *BulkIndexDocumentsRequestAllOf) GetDocumentsOk() (*[]DocumentDefinition, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *BulkIndexDocumentsRequestAllOf) SetDocuments(v []DocumentDefinition)`

SetDocuments sets Documents field to given value.


### GetDisableStaleDocumentDeletionCheck

`func (o *BulkIndexDocumentsRequestAllOf) GetDisableStaleDocumentDeletionCheck() bool`

GetDisableStaleDocumentDeletionCheck returns the DisableStaleDocumentDeletionCheck field if non-nil, zero value otherwise.

### GetDisableStaleDocumentDeletionCheckOk

`func (o *BulkIndexDocumentsRequestAllOf) GetDisableStaleDocumentDeletionCheckOk() (*bool, bool)`

GetDisableStaleDocumentDeletionCheckOk returns a tuple with the DisableStaleDocumentDeletionCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableStaleDocumentDeletionCheck

`func (o *BulkIndexDocumentsRequestAllOf) SetDisableStaleDocumentDeletionCheck(v bool)`

SetDisableStaleDocumentDeletionCheck sets DisableStaleDocumentDeletionCheck field to given value.

### HasDisableStaleDocumentDeletionCheck

`func (o *BulkIndexDocumentsRequestAllOf) HasDisableStaleDocumentDeletionCheck() bool`

HasDisableStaleDocumentDeletionCheck returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


