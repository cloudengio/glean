# BulkIndexRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UploadId** | **string** | Unique id that must be used for this instance of datasource employees upload | 
**IsFirstPage** | Pointer to **bool** | true if this is the first page of the upload. Defaults to false | [optional] 
**IsLastPage** | Pointer to **bool** | true if this is the last page of the upload. Defaults to false | [optional] 
**ForceRestartUpload** | Pointer to **bool** | Flag to discard previous upload attempts and start from scratch. Must be specified with isFirstPage&#x3D;true | [optional] 

## Methods

### NewBulkIndexRequest

`func NewBulkIndexRequest(uploadId string, ) *BulkIndexRequest`

NewBulkIndexRequest instantiates a new BulkIndexRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkIndexRequestWithDefaults

`func NewBulkIndexRequestWithDefaults() *BulkIndexRequest`

NewBulkIndexRequestWithDefaults instantiates a new BulkIndexRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUploadId

`func (o *BulkIndexRequest) GetUploadId() string`

GetUploadId returns the UploadId field if non-nil, zero value otherwise.

### GetUploadIdOk

`func (o *BulkIndexRequest) GetUploadIdOk() (*string, bool)`

GetUploadIdOk returns a tuple with the UploadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadId

`func (o *BulkIndexRequest) SetUploadId(v string)`

SetUploadId sets UploadId field to given value.


### GetIsFirstPage

`func (o *BulkIndexRequest) GetIsFirstPage() bool`

GetIsFirstPage returns the IsFirstPage field if non-nil, zero value otherwise.

### GetIsFirstPageOk

`func (o *BulkIndexRequest) GetIsFirstPageOk() (*bool, bool)`

GetIsFirstPageOk returns a tuple with the IsFirstPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFirstPage

`func (o *BulkIndexRequest) SetIsFirstPage(v bool)`

SetIsFirstPage sets IsFirstPage field to given value.

### HasIsFirstPage

`func (o *BulkIndexRequest) HasIsFirstPage() bool`

HasIsFirstPage returns a boolean if a field has been set.

### GetIsLastPage

`func (o *BulkIndexRequest) GetIsLastPage() bool`

GetIsLastPage returns the IsLastPage field if non-nil, zero value otherwise.

### GetIsLastPageOk

`func (o *BulkIndexRequest) GetIsLastPageOk() (*bool, bool)`

GetIsLastPageOk returns a tuple with the IsLastPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLastPage

`func (o *BulkIndexRequest) SetIsLastPage(v bool)`

SetIsLastPage sets IsLastPage field to given value.

### HasIsLastPage

`func (o *BulkIndexRequest) HasIsLastPage() bool`

HasIsLastPage returns a boolean if a field has been set.

### GetForceRestartUpload

`func (o *BulkIndexRequest) GetForceRestartUpload() bool`

GetForceRestartUpload returns the ForceRestartUpload field if non-nil, zero value otherwise.

### GetForceRestartUploadOk

`func (o *BulkIndexRequest) GetForceRestartUploadOk() (*bool, bool)`

GetForceRestartUploadOk returns a tuple with the ForceRestartUpload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceRestartUpload

`func (o *BulkIndexRequest) SetForceRestartUpload(v bool)`

SetForceRestartUpload sets ForceRestartUpload field to given value.

### HasForceRestartUpload

`func (o *BulkIndexRequest) HasForceRestartUpload() bool`

HasForceRestartUpload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


