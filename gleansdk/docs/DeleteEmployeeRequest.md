# DeleteEmployeeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **int64** | Version number for document for optimistic concurrency control. If absent or 0 then no version checks are done. | [optional] 
**EmployeeEmail** | **string** | The deleted employee&#39;s email | 

## Methods

### NewDeleteEmployeeRequest

`func NewDeleteEmployeeRequest(employeeEmail string, ) *DeleteEmployeeRequest`

NewDeleteEmployeeRequest instantiates a new DeleteEmployeeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteEmployeeRequestWithDefaults

`func NewDeleteEmployeeRequestWithDefaults() *DeleteEmployeeRequest`

NewDeleteEmployeeRequestWithDefaults instantiates a new DeleteEmployeeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *DeleteEmployeeRequest) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DeleteEmployeeRequest) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DeleteEmployeeRequest) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DeleteEmployeeRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetEmployeeEmail

`func (o *DeleteEmployeeRequest) GetEmployeeEmail() string`

GetEmployeeEmail returns the EmployeeEmail field if non-nil, zero value otherwise.

### GetEmployeeEmailOk

`func (o *DeleteEmployeeRequest) GetEmployeeEmailOk() (*string, bool)`

GetEmployeeEmailOk returns a tuple with the EmployeeEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeEmail

`func (o *DeleteEmployeeRequest) SetEmployeeEmail(v string)`

SetEmployeeEmail sets EmployeeEmail field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


