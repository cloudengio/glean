# IndexEmployeeListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Employees** | Pointer to [**[]IndexEmployeeRequest**](IndexEmployeeRequest.md) | List of employee info and version. | [optional] 

## Methods

### NewIndexEmployeeListRequest

`func NewIndexEmployeeListRequest() *IndexEmployeeListRequest`

NewIndexEmployeeListRequest instantiates a new IndexEmployeeListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndexEmployeeListRequestWithDefaults

`func NewIndexEmployeeListRequestWithDefaults() *IndexEmployeeListRequest`

NewIndexEmployeeListRequestWithDefaults instantiates a new IndexEmployeeListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmployees

`func (o *IndexEmployeeListRequest) GetEmployees() []IndexEmployeeRequest`

GetEmployees returns the Employees field if non-nil, zero value otherwise.

### GetEmployeesOk

`func (o *IndexEmployeeListRequest) GetEmployeesOk() (*[]IndexEmployeeRequest, bool)`

GetEmployeesOk returns a tuple with the Employees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployees

`func (o *IndexEmployeeListRequest) SetEmployees(v []IndexEmployeeRequest)`

SetEmployees sets Employees field to given value.

### HasEmployees

`func (o *IndexEmployeeListRequest) HasEmployees() bool`

HasEmployees returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


