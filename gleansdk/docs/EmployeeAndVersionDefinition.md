# EmployeeAndVersionDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Employee** | Pointer to [**EmployeeInfoDefinition**](EmployeeInfoDefinition.md) |  | [optional] 
**Version** | Pointer to **int64** | Version number for the employee object. If absent or 0 then no version checks are done | [optional] 

## Methods

### NewEmployeeAndVersionDefinition

`func NewEmployeeAndVersionDefinition() *EmployeeAndVersionDefinition`

NewEmployeeAndVersionDefinition instantiates a new EmployeeAndVersionDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmployeeAndVersionDefinitionWithDefaults

`func NewEmployeeAndVersionDefinitionWithDefaults() *EmployeeAndVersionDefinition`

NewEmployeeAndVersionDefinitionWithDefaults instantiates a new EmployeeAndVersionDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmployee

`func (o *EmployeeAndVersionDefinition) GetEmployee() EmployeeInfoDefinition`

GetEmployee returns the Employee field if non-nil, zero value otherwise.

### GetEmployeeOk

`func (o *EmployeeAndVersionDefinition) GetEmployeeOk() (*EmployeeInfoDefinition, bool)`

GetEmployeeOk returns a tuple with the Employee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployee

`func (o *EmployeeAndVersionDefinition) SetEmployee(v EmployeeInfoDefinition)`

SetEmployee sets Employee field to given value.

### HasEmployee

`func (o *EmployeeAndVersionDefinition) HasEmployee() bool`

HasEmployee returns a boolean if a field has been set.

### GetVersion

`func (o *EmployeeAndVersionDefinition) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *EmployeeAndVersionDefinition) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *EmployeeAndVersionDefinition) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *EmployeeAndVersionDefinition) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


