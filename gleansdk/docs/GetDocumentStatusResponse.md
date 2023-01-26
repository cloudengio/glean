# GetDocumentStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UploadStatus** | Pointer to **string** | Upload status, enum of NOT_UPLOADED, UPLOADED, STATUS_UNKNOWN | [optional] 
**LastUploadedAt** | Pointer to **int64** | Time of last successful upload, in epoch seconds | [optional] 
**IndexingStatus** | Pointer to **string** | Indexing status, enum of NOT_INDEXED, INDEXED, STATUS_UNKNOWN | [optional] 
**LastIndexedAt** | Pointer to **int64** | Time of last successful indexing, in epoch seconds | [optional] 

## Methods

### NewGetDocumentStatusResponse

`func NewGetDocumentStatusResponse() *GetDocumentStatusResponse`

NewGetDocumentStatusResponse instantiates a new GetDocumentStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDocumentStatusResponseWithDefaults

`func NewGetDocumentStatusResponseWithDefaults() *GetDocumentStatusResponse`

NewGetDocumentStatusResponseWithDefaults instantiates a new GetDocumentStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUploadStatus

`func (o *GetDocumentStatusResponse) GetUploadStatus() string`

GetUploadStatus returns the UploadStatus field if non-nil, zero value otherwise.

### GetUploadStatusOk

`func (o *GetDocumentStatusResponse) GetUploadStatusOk() (*string, bool)`

GetUploadStatusOk returns a tuple with the UploadStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadStatus

`func (o *GetDocumentStatusResponse) SetUploadStatus(v string)`

SetUploadStatus sets UploadStatus field to given value.

### HasUploadStatus

`func (o *GetDocumentStatusResponse) HasUploadStatus() bool`

HasUploadStatus returns a boolean if a field has been set.

### GetLastUploadedAt

`func (o *GetDocumentStatusResponse) GetLastUploadedAt() int64`

GetLastUploadedAt returns the LastUploadedAt field if non-nil, zero value otherwise.

### GetLastUploadedAtOk

`func (o *GetDocumentStatusResponse) GetLastUploadedAtOk() (*int64, bool)`

GetLastUploadedAtOk returns a tuple with the LastUploadedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUploadedAt

`func (o *GetDocumentStatusResponse) SetLastUploadedAt(v int64)`

SetLastUploadedAt sets LastUploadedAt field to given value.

### HasLastUploadedAt

`func (o *GetDocumentStatusResponse) HasLastUploadedAt() bool`

HasLastUploadedAt returns a boolean if a field has been set.

### GetIndexingStatus

`func (o *GetDocumentStatusResponse) GetIndexingStatus() string`

GetIndexingStatus returns the IndexingStatus field if non-nil, zero value otherwise.

### GetIndexingStatusOk

`func (o *GetDocumentStatusResponse) GetIndexingStatusOk() (*string, bool)`

GetIndexingStatusOk returns a tuple with the IndexingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexingStatus

`func (o *GetDocumentStatusResponse) SetIndexingStatus(v string)`

SetIndexingStatus sets IndexingStatus field to given value.

### HasIndexingStatus

`func (o *GetDocumentStatusResponse) HasIndexingStatus() bool`

HasIndexingStatus returns a boolean if a field has been set.

### GetLastIndexedAt

`func (o *GetDocumentStatusResponse) GetLastIndexedAt() int64`

GetLastIndexedAt returns the LastIndexedAt field if non-nil, zero value otherwise.

### GetLastIndexedAtOk

`func (o *GetDocumentStatusResponse) GetLastIndexedAtOk() (*int64, bool)`

GetLastIndexedAtOk returns a tuple with the LastIndexedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastIndexedAt

`func (o *GetDocumentStatusResponse) SetLastIndexedAt(v int64)`

SetLastIndexedAt sets LastIndexedAt field to given value.

### HasLastIndexedAt

`func (o *GetDocumentStatusResponse) HasLastIndexedAt() bool`

HasLastIndexedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


