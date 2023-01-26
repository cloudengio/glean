# BulkIndexUsersRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UploadId** | **string** | Unique id that must be used for this instance of datasource users upload | 
**IsFirstPage** | Pointer to **bool** | true if this is the first page of the upload. Defaults to false | [optional] 
**IsLastPage** | Pointer to **bool** | true if this is the last page of the upload. Defaults to false | [optional] 
**ForceRestartUpload** | Pointer to **bool** | Flag to discard previous upload attempts and start from scratch. Must be specified with isFirstPage&#x3D;true | [optional] 
**Datasource** | **string** | datasource of the users | 
**Users** | [**[]DatasourceUserDefinition**](DatasourceUserDefinition.md) | batch of users for the datasource | 

## Methods

### NewBulkIndexUsersRequest

`func NewBulkIndexUsersRequest(uploadId string, datasource string, users []DatasourceUserDefinition, ) *BulkIndexUsersRequest`

NewBulkIndexUsersRequest instantiates a new BulkIndexUsersRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkIndexUsersRequestWithDefaults

`func NewBulkIndexUsersRequestWithDefaults() *BulkIndexUsersRequest`

NewBulkIndexUsersRequestWithDefaults instantiates a new BulkIndexUsersRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUploadId

`func (o *BulkIndexUsersRequest) GetUploadId() string`

GetUploadId returns the UploadId field if non-nil, zero value otherwise.

### GetUploadIdOk

`func (o *BulkIndexUsersRequest) GetUploadIdOk() (*string, bool)`

GetUploadIdOk returns a tuple with the UploadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadId

`func (o *BulkIndexUsersRequest) SetUploadId(v string)`

SetUploadId sets UploadId field to given value.


### GetIsFirstPage

`func (o *BulkIndexUsersRequest) GetIsFirstPage() bool`

GetIsFirstPage returns the IsFirstPage field if non-nil, zero value otherwise.

### GetIsFirstPageOk

`func (o *BulkIndexUsersRequest) GetIsFirstPageOk() (*bool, bool)`

GetIsFirstPageOk returns a tuple with the IsFirstPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFirstPage

`func (o *BulkIndexUsersRequest) SetIsFirstPage(v bool)`

SetIsFirstPage sets IsFirstPage field to given value.

### HasIsFirstPage

`func (o *BulkIndexUsersRequest) HasIsFirstPage() bool`

HasIsFirstPage returns a boolean if a field has been set.

### GetIsLastPage

`func (o *BulkIndexUsersRequest) GetIsLastPage() bool`

GetIsLastPage returns the IsLastPage field if non-nil, zero value otherwise.

### GetIsLastPageOk

`func (o *BulkIndexUsersRequest) GetIsLastPageOk() (*bool, bool)`

GetIsLastPageOk returns a tuple with the IsLastPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLastPage

`func (o *BulkIndexUsersRequest) SetIsLastPage(v bool)`

SetIsLastPage sets IsLastPage field to given value.

### HasIsLastPage

`func (o *BulkIndexUsersRequest) HasIsLastPage() bool`

HasIsLastPage returns a boolean if a field has been set.

### GetForceRestartUpload

`func (o *BulkIndexUsersRequest) GetForceRestartUpload() bool`

GetForceRestartUpload returns the ForceRestartUpload field if non-nil, zero value otherwise.

### GetForceRestartUploadOk

`func (o *BulkIndexUsersRequest) GetForceRestartUploadOk() (*bool, bool)`

GetForceRestartUploadOk returns a tuple with the ForceRestartUpload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceRestartUpload

`func (o *BulkIndexUsersRequest) SetForceRestartUpload(v bool)`

SetForceRestartUpload sets ForceRestartUpload field to given value.

### HasForceRestartUpload

`func (o *BulkIndexUsersRequest) HasForceRestartUpload() bool`

HasForceRestartUpload returns a boolean if a field has been set.

### GetDatasource

`func (o *BulkIndexUsersRequest) GetDatasource() string`

GetDatasource returns the Datasource field if non-nil, zero value otherwise.

### GetDatasourceOk

`func (o *BulkIndexUsersRequest) GetDatasourceOk() (*string, bool)`

GetDatasourceOk returns a tuple with the Datasource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasource

`func (o *BulkIndexUsersRequest) SetDatasource(v string)`

SetDatasource sets Datasource field to given value.


### GetUsers

`func (o *BulkIndexUsersRequest) GetUsers() []DatasourceUserDefinition`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *BulkIndexUsersRequest) GetUsersOk() (*[]DatasourceUserDefinition, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *BulkIndexUsersRequest) SetUsers(v []DatasourceUserDefinition)`

SetUsers sets Users field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


