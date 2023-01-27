# DatasourceUserDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** | To be supplied if the user id in the datasource is not the email | [optional] 
**Name** | **string** |  | 
**IsActive** | Pointer to **bool** | set to false if the user is a former employee or a bot | [optional] 

## Methods

### NewDatasourceUserDefinition

`func NewDatasourceUserDefinition(name string, ) *DatasourceUserDefinition`

NewDatasourceUserDefinition instantiates a new DatasourceUserDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatasourceUserDefinitionWithDefaults

`func NewDatasourceUserDefinitionWithDefaults() *DatasourceUserDefinition`

NewDatasourceUserDefinitionWithDefaults instantiates a new DatasourceUserDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *DatasourceUserDefinition) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *DatasourceUserDefinition) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *DatasourceUserDefinition) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *DatasourceUserDefinition) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetUserId

`func (o *DatasourceUserDefinition) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *DatasourceUserDefinition) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *DatasourceUserDefinition) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *DatasourceUserDefinition) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetName

`func (o *DatasourceUserDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DatasourceUserDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DatasourceUserDefinition) SetName(v string)`

SetName sets Name field to given value.


### GetIsActive

`func (o *DatasourceUserDefinition) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *DatasourceUserDefinition) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *DatasourceUserDefinition) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *DatasourceUserDefinition) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


