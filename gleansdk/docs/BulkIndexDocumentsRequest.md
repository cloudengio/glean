# BulkIndexDocumentsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UploadId** | **string** | Unique id that must be used for this instance of datasource employees upload | 
**IsFirstPage** | Pointer to **bool** | true if this is the first page of the upload. Defaults to false | [optional] 
**IsLastPage** | Pointer to **bool** | true if this is the last page of the upload. Defaults to false | [optional] 
**ForceRestartUpload** | Pointer to **bool** | Flag to discard previous upload attempts and start from scratch. Must be specified with isFirstPage&#x3D;true | [optional] 
**Datasource** | **string** | Datasource of the documents | 
**Documents** | [**[]DocumentDefinition**](DocumentDefinition.md) | Batch of documents for the datasource | 
**DisableStaleDocumentDeletionCheck** | Pointer to **bool** | True if older documents need to be force deleted after the upload completes. Defaults to older documents being deleted asynchronously. This must only be set when &#x60;isLastPage &#x3D; true&#x60; | [optional] 

## Methods

### NewBulkIndexDocumentsRequest

`func NewBulkIndexDocumentsRequest(uploadId string, datasource string, documents []DocumentDefinition, ) *BulkIndexDocumentsRequest`

NewBulkIndexDocumentsRequest instantiates a new BulkIndexDocumentsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkIndexDocumentsRequestWithDefaults

`func NewBulkIndexDocumentsRequestWithDefaults() *BulkIndexDocumentsRequest`

NewBulkIndexDocumentsRequestWithDefaults instantiates a new BulkIndexDocumentsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUploadId

`func (o *BulkIndexDocumentsRequest) GetUploadId() string`

GetUploadId returns the UploadId field if non-nil, zero value otherwise.

### GetUploadIdOk

`func (o *BulkIndexDocumentsRequest) GetUploadIdOk() (*string, bool)`

GetUploadIdOk returns a tuple with the UploadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadId

`func (o *BulkIndexDocumentsRequest) SetUploadId(v string)`

SetUploadId sets UploadId field to given value.


### GetIsFirstPage

`func (o *BulkIndexDocumentsRequest) GetIsFirstPage() bool`

GetIsFirstPage returns the IsFirstPage field if non-nil, zero value otherwise.

### GetIsFirstPageOk

`func (o *BulkIndexDocumentsRequest) GetIsFirstPageOk() (*bool, bool)`

GetIsFirstPageOk returns a tuple with the IsFirstPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFirstPage

`func (o *BulkIndexDocumentsRequest) SetIsFirstPage(v bool)`

SetIsFirstPage sets IsFirstPage field to given value.

### HasIsFirstPage

`func (o *BulkIndexDocumentsRequest) HasIsFirstPage() bool`

HasIsFirstPage returns a boolean if a field has been set.

### GetIsLastPage

`func (o *BulkIndexDocumentsRequest) GetIsLastPage() bool`

GetIsLastPage returns the IsLastPage field if non-nil, zero value otherwise.

### GetIsLastPageOk

`func (o *BulkIndexDocumentsRequest) GetIsLastPageOk() (*bool, bool)`

GetIsLastPageOk returns a tuple with the IsLastPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLastPage

`func (o *BulkIndexDocumentsRequest) SetIsLastPage(v bool)`

SetIsLastPage sets IsLastPage field to given value.

### HasIsLastPage

`func (o *BulkIndexDocumentsRequest) HasIsLastPage() bool`

HasIsLastPage returns a boolean if a field has been set.

### GetForceRestartUpload

`func (o *BulkIndexDocumentsRequest) GetForceRestartUpload() bool`

GetForceRestartUpload returns the ForceRestartUpload field if non-nil, zero value otherwise.

### GetForceRestartUploadOk

`func (o *BulkIndexDocumentsRequest) GetForceRestartUploadOk() (*bool, bool)`

GetForceRestartUploadOk returns a tuple with the ForceRestartUpload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceRestartUpload

`func (o *BulkIndexDocumentsRequest) SetForceRestartUpload(v bool)`

SetForceRestartUpload sets ForceRestartUpload field to given value.

### HasForceRestartUpload

`func (o *BulkIndexDocumentsRequest) HasForceRestartUpload() bool`

HasForceRestartUpload returns a boolean if a field has been set.

### GetDatasource

`func (o *BulkIndexDocumentsRequest) GetDatasource() string`

GetDatasource returns the Datasource field if non-nil, zero value otherwise.

### GetDatasourceOk

`func (o *BulkIndexDocumentsRequest) GetDatasourceOk() (*string, bool)`

GetDatasourceOk returns a tuple with the Datasource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasource

`func (o *BulkIndexDocumentsRequest) SetDatasource(v string)`

SetDatasource sets Datasource field to given value.


### GetDocuments

`func (o *BulkIndexDocumentsRequest) GetDocuments() []DocumentDefinition`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *BulkIndexDocumentsRequest) GetDocumentsOk() (*[]DocumentDefinition, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *BulkIndexDocumentsRequest) SetDocuments(v []DocumentDefinition)`

SetDocuments sets Documents field to given value.


### GetDisableStaleDocumentDeletionCheck

`func (o *BulkIndexDocumentsRequest) GetDisableStaleDocumentDeletionCheck() bool`

GetDisableStaleDocumentDeletionCheck returns the DisableStaleDocumentDeletionCheck field if non-nil, zero value otherwise.

### GetDisableStaleDocumentDeletionCheckOk

`func (o *BulkIndexDocumentsRequest) GetDisableStaleDocumentDeletionCheckOk() (*bool, bool)`

GetDisableStaleDocumentDeletionCheckOk returns a tuple with the DisableStaleDocumentDeletionCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableStaleDocumentDeletionCheck

`func (o *BulkIndexDocumentsRequest) SetDisableStaleDocumentDeletionCheck(v bool)`

SetDisableStaleDocumentDeletionCheck sets DisableStaleDocumentDeletionCheck field to given value.

### HasDisableStaleDocumentDeletionCheck

`func (o *BulkIndexDocumentsRequest) HasDisableStaleDocumentDeletionCheck() bool`

HasDisableStaleDocumentDeletionCheck returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


