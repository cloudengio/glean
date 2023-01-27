# BulkIndexTeamsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UploadId** | **string** | Unique id that must be used for this instance of datasource employees upload | 
**IsFirstPage** | Pointer to **bool** | true if this is the first page of the upload. Defaults to false | [optional] 
**IsLastPage** | Pointer to **bool** | true if this is the last page of the upload. Defaults to false | [optional] 
**ForceRestartUpload** | Pointer to **bool** | Flag to discard previous upload attempts and start from scratch. Must be specified with isFirstPage&#x3D;true | [optional] 
**Teams** | [**[]TeamInfoDefinition**](TeamInfoDefinition.md) | Batch of team information | 

## Methods

### NewBulkIndexTeamsRequest

`func NewBulkIndexTeamsRequest(uploadId string, teams []TeamInfoDefinition, ) *BulkIndexTeamsRequest`

NewBulkIndexTeamsRequest instantiates a new BulkIndexTeamsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkIndexTeamsRequestWithDefaults

`func NewBulkIndexTeamsRequestWithDefaults() *BulkIndexTeamsRequest`

NewBulkIndexTeamsRequestWithDefaults instantiates a new BulkIndexTeamsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUploadId

`func (o *BulkIndexTeamsRequest) GetUploadId() string`

GetUploadId returns the UploadId field if non-nil, zero value otherwise.

### GetUploadIdOk

`func (o *BulkIndexTeamsRequest) GetUploadIdOk() (*string, bool)`

GetUploadIdOk returns a tuple with the UploadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadId

`func (o *BulkIndexTeamsRequest) SetUploadId(v string)`

SetUploadId sets UploadId field to given value.


### GetIsFirstPage

`func (o *BulkIndexTeamsRequest) GetIsFirstPage() bool`

GetIsFirstPage returns the IsFirstPage field if non-nil, zero value otherwise.

### GetIsFirstPageOk

`func (o *BulkIndexTeamsRequest) GetIsFirstPageOk() (*bool, bool)`

GetIsFirstPageOk returns a tuple with the IsFirstPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFirstPage

`func (o *BulkIndexTeamsRequest) SetIsFirstPage(v bool)`

SetIsFirstPage sets IsFirstPage field to given value.

### HasIsFirstPage

`func (o *BulkIndexTeamsRequest) HasIsFirstPage() bool`

HasIsFirstPage returns a boolean if a field has been set.

### GetIsLastPage

`func (o *BulkIndexTeamsRequest) GetIsLastPage() bool`

GetIsLastPage returns the IsLastPage field if non-nil, zero value otherwise.

### GetIsLastPageOk

`func (o *BulkIndexTeamsRequest) GetIsLastPageOk() (*bool, bool)`

GetIsLastPageOk returns a tuple with the IsLastPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLastPage

`func (o *BulkIndexTeamsRequest) SetIsLastPage(v bool)`

SetIsLastPage sets IsLastPage field to given value.

### HasIsLastPage

`func (o *BulkIndexTeamsRequest) HasIsLastPage() bool`

HasIsLastPage returns a boolean if a field has been set.

### GetForceRestartUpload

`func (o *BulkIndexTeamsRequest) GetForceRestartUpload() bool`

GetForceRestartUpload returns the ForceRestartUpload field if non-nil, zero value otherwise.

### GetForceRestartUploadOk

`func (o *BulkIndexTeamsRequest) GetForceRestartUploadOk() (*bool, bool)`

GetForceRestartUploadOk returns a tuple with the ForceRestartUpload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceRestartUpload

`func (o *BulkIndexTeamsRequest) SetForceRestartUpload(v bool)`

SetForceRestartUpload sets ForceRestartUpload field to given value.

### HasForceRestartUpload

`func (o *BulkIndexTeamsRequest) HasForceRestartUpload() bool`

HasForceRestartUpload returns a boolean if a field has been set.

### GetTeams

`func (o *BulkIndexTeamsRequest) GetTeams() []TeamInfoDefinition`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *BulkIndexTeamsRequest) GetTeamsOk() (*[]TeamInfoDefinition, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *BulkIndexTeamsRequest) SetTeams(v []TeamInfoDefinition)`

SetTeams sets Teams field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


