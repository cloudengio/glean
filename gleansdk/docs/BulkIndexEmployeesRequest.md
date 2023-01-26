# BulkIndexEmployeesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UploadId** | **string** | Unique id that must be used for this instance of datasource employees upload | 
**IsFirstPage** | Pointer to **bool** | true if this is the first page of the upload. Defaults to false | [optional] 
**IsLastPage** | Pointer to **bool** | true if this is the last page of the upload. Defaults to false | [optional] 
**ForceRestartUpload** | Pointer to **bool** | Flag to discard previous upload attempts and start from scratch. Must be specified with isFirstPage&#x3D;true | [optional] 
**Employees** | [**[]EmployeeInfoDefinition**](EmployeeInfoDefinition.md) | Batch of employee information | 

## Methods

### NewBulkIndexEmployeesRequest

`func NewBulkIndexEmployeesRequest(uploadId string, employees []EmployeeInfoDefinition, ) *BulkIndexEmployeesRequest`

NewBulkIndexEmployeesRequest instantiates a new BulkIndexEmployeesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkIndexEmployeesRequestWithDefaults

`func NewBulkIndexEmployeesRequestWithDefaults() *BulkIndexEmployeesRequest`

NewBulkIndexEmployeesRequestWithDefaults instantiates a new BulkIndexEmployeesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUploadId

`func (o *BulkIndexEmployeesRequest) GetUploadId() string`

GetUploadId returns the UploadId field if non-nil, zero value otherwise.

### GetUploadIdOk

`func (o *BulkIndexEmployeesRequest) GetUploadIdOk() (*string, bool)`

GetUploadIdOk returns a tuple with the UploadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadId

`func (o *BulkIndexEmployeesRequest) SetUploadId(v string)`

SetUploadId sets UploadId field to given value.


### GetIsFirstPage

`func (o *BulkIndexEmployeesRequest) GetIsFirstPage() bool`

GetIsFirstPage returns the IsFirstPage field if non-nil, zero value otherwise.

### GetIsFirstPageOk

`func (o *BulkIndexEmployeesRequest) GetIsFirstPageOk() (*bool, bool)`

GetIsFirstPageOk returns a tuple with the IsFirstPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFirstPage

`func (o *BulkIndexEmployeesRequest) SetIsFirstPage(v bool)`

SetIsFirstPage sets IsFirstPage field to given value.

### HasIsFirstPage

`func (o *BulkIndexEmployeesRequest) HasIsFirstPage() bool`

HasIsFirstPage returns a boolean if a field has been set.

### GetIsLastPage

`func (o *BulkIndexEmployeesRequest) GetIsLastPage() bool`

GetIsLastPage returns the IsLastPage field if non-nil, zero value otherwise.

### GetIsLastPageOk

`func (o *BulkIndexEmployeesRequest) GetIsLastPageOk() (*bool, bool)`

GetIsLastPageOk returns a tuple with the IsLastPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLastPage

`func (o *BulkIndexEmployeesRequest) SetIsLastPage(v bool)`

SetIsLastPage sets IsLastPage field to given value.

### HasIsLastPage

`func (o *BulkIndexEmployeesRequest) HasIsLastPage() bool`

HasIsLastPage returns a boolean if a field has been set.

### GetForceRestartUpload

`func (o *BulkIndexEmployeesRequest) GetForceRestartUpload() bool`

GetForceRestartUpload returns the ForceRestartUpload field if non-nil, zero value otherwise.

### GetForceRestartUploadOk

`func (o *BulkIndexEmployeesRequest) GetForceRestartUploadOk() (*bool, bool)`

GetForceRestartUploadOk returns a tuple with the ForceRestartUpload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceRestartUpload

`func (o *BulkIndexEmployeesRequest) SetForceRestartUpload(v bool)`

SetForceRestartUpload sets ForceRestartUpload field to given value.

### HasForceRestartUpload

`func (o *BulkIndexEmployeesRequest) HasForceRestartUpload() bool`

HasForceRestartUpload returns a boolean if a field has been set.

### GetEmployees

`func (o *BulkIndexEmployeesRequest) GetEmployees() []EmployeeInfoDefinition`

GetEmployees returns the Employees field if non-nil, zero value otherwise.

### GetEmployeesOk

`func (o *BulkIndexEmployeesRequest) GetEmployeesOk() (*[]EmployeeInfoDefinition, bool)`

GetEmployeesOk returns a tuple with the Employees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployees

`func (o *BulkIndexEmployeesRequest) SetEmployees(v []EmployeeInfoDefinition)`

SetEmployees sets Employees field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


