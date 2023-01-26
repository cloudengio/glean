# IndexEmployeeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Employee** | [**EmployeeInfoDefinition**](EmployeeInfoDefinition.md) |  | 
**Version** | Pointer to **int64** | Version number for the employee object. If absent or 0 then no version checks are done | [optional] 

## Methods

### NewIndexEmployeeRequest

`func NewIndexEmployeeRequest(employee EmployeeInfoDefinition, ) *IndexEmployeeRequest`

NewIndexEmployeeRequest instantiates a new IndexEmployeeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndexEmployeeRequestWithDefaults

`func NewIndexEmployeeRequestWithDefaults() *IndexEmployeeRequest`

NewIndexEmployeeRequestWithDefaults instantiates a new IndexEmployeeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmployee

`func (o *IndexEmployeeRequest) GetEmployee() EmployeeInfoDefinition`

GetEmployee returns the Employee field if non-nil, zero value otherwise.

### GetEmployeeOk

`func (o *IndexEmployeeRequest) GetEmployeeOk() (*EmployeeInfoDefinition, bool)`

GetEmployeeOk returns a tuple with the Employee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployee

`func (o *IndexEmployeeRequest) SetEmployee(v EmployeeInfoDefinition)`

SetEmployee sets Employee field to given value.


### GetVersion

`func (o *IndexEmployeeRequest) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *IndexEmployeeRequest) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *IndexEmployeeRequest) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *IndexEmployeeRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


