# BulkIndexGroupsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UploadId** | **string** | Unique id that must be used for this instance of datasource groups upload | 
**IsFirstPage** | Pointer to **bool** | true if this is the first page of the upload. Defaults to false | [optional] 
**IsLastPage** | Pointer to **bool** | true if this is the last page of the upload. Defaults to false | [optional] 
**ForceRestartUpload** | Pointer to **bool** | Flag to discard previous upload attempts and start from scratch. Must be specified with isFirstPage&#x3D;true | [optional] 
**Datasource** | **string** | datasource of the groups | 
**Groups** | [**[]DatasourceGroupDefinition**](DatasourceGroupDefinition.md) | batch of groups for the datasource | 

## Methods

### NewBulkIndexGroupsRequest

`func NewBulkIndexGroupsRequest(uploadId string, datasource string, groups []DatasourceGroupDefinition, ) *BulkIndexGroupsRequest`

NewBulkIndexGroupsRequest instantiates a new BulkIndexGroupsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkIndexGroupsRequestWithDefaults

`func NewBulkIndexGroupsRequestWithDefaults() *BulkIndexGroupsRequest`

NewBulkIndexGroupsRequestWithDefaults instantiates a new BulkIndexGroupsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUploadId

`func (o *BulkIndexGroupsRequest) GetUploadId() string`

GetUploadId returns the UploadId field if non-nil, zero value otherwise.

### GetUploadIdOk

`func (o *BulkIndexGroupsRequest) GetUploadIdOk() (*string, bool)`

GetUploadIdOk returns a tuple with the UploadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadId

`func (o *BulkIndexGroupsRequest) SetUploadId(v string)`

SetUploadId sets UploadId field to given value.


### GetIsFirstPage

`func (o *BulkIndexGroupsRequest) GetIsFirstPage() bool`

GetIsFirstPage returns the IsFirstPage field if non-nil, zero value otherwise.

### GetIsFirstPageOk

`func (o *BulkIndexGroupsRequest) GetIsFirstPageOk() (*bool, bool)`

GetIsFirstPageOk returns a tuple with the IsFirstPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFirstPage

`func (o *BulkIndexGroupsRequest) SetIsFirstPage(v bool)`

SetIsFirstPage sets IsFirstPage field to given value.

### HasIsFirstPage

`func (o *BulkIndexGroupsRequest) HasIsFirstPage() bool`

HasIsFirstPage returns a boolean if a field has been set.

### GetIsLastPage

`func (o *BulkIndexGroupsRequest) GetIsLastPage() bool`

GetIsLastPage returns the IsLastPage field if non-nil, zero value otherwise.

### GetIsLastPageOk

`func (o *BulkIndexGroupsRequest) GetIsLastPageOk() (*bool, bool)`

GetIsLastPageOk returns a tuple with the IsLastPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLastPage

`func (o *BulkIndexGroupsRequest) SetIsLastPage(v bool)`

SetIsLastPage sets IsLastPage field to given value.

### HasIsLastPage

`func (o *BulkIndexGroupsRequest) HasIsLastPage() bool`

HasIsLastPage returns a boolean if a field has been set.

### GetForceRestartUpload

`func (o *BulkIndexGroupsRequest) GetForceRestartUpload() bool`

GetForceRestartUpload returns the ForceRestartUpload field if non-nil, zero value otherwise.

### GetForceRestartUploadOk

`func (o *BulkIndexGroupsRequest) GetForceRestartUploadOk() (*bool, bool)`

GetForceRestartUploadOk returns a tuple with the ForceRestartUpload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceRestartUpload

`func (o *BulkIndexGroupsRequest) SetForceRestartUpload(v bool)`

SetForceRestartUpload sets ForceRestartUpload field to given value.

### HasForceRestartUpload

`func (o *BulkIndexGroupsRequest) HasForceRestartUpload() bool`

HasForceRestartUpload returns a boolean if a field has been set.

### GetDatasource

`func (o *BulkIndexGroupsRequest) GetDatasource() string`

GetDatasource returns the Datasource field if non-nil, zero value otherwise.

### GetDatasourceOk

`func (o *BulkIndexGroupsRequest) GetDatasourceOk() (*string, bool)`

GetDatasourceOk returns a tuple with the Datasource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasource

`func (o *BulkIndexGroupsRequest) SetDatasource(v string)`

SetDatasource sets Datasource field to given value.


### GetGroups

`func (o *BulkIndexGroupsRequest) GetGroups() []DatasourceGroupDefinition`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *BulkIndexGroupsRequest) GetGroupsOk() (*[]DatasourceGroupDefinition, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *BulkIndexGroupsRequest) SetGroups(v []DatasourceGroupDefinition)`

SetGroups sets Groups field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


