# UserReferenceDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** |  | [optional] 
**DatasourceUserId** | Pointer to **string** | some datasources refer to the user by the datasource user id in the document | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewUserReferenceDefinition

`func NewUserReferenceDefinition() *UserReferenceDefinition`

NewUserReferenceDefinition instantiates a new UserReferenceDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserReferenceDefinitionWithDefaults

`func NewUserReferenceDefinitionWithDefaults() *UserReferenceDefinition`

NewUserReferenceDefinitionWithDefaults instantiates a new UserReferenceDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *UserReferenceDefinition) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserReferenceDefinition) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserReferenceDefinition) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserReferenceDefinition) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetDatasourceUserId

`func (o *UserReferenceDefinition) GetDatasourceUserId() string`

GetDatasourceUserId returns the DatasourceUserId field if non-nil, zero value otherwise.

### GetDatasourceUserIdOk

`func (o *UserReferenceDefinition) GetDatasourceUserIdOk() (*string, bool)`

GetDatasourceUserIdOk returns a tuple with the DatasourceUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasourceUserId

`func (o *UserReferenceDefinition) SetDatasourceUserId(v string)`

SetDatasourceUserId sets DatasourceUserId field to given value.

### HasDatasourceUserId

`func (o *UserReferenceDefinition) HasDatasourceUserId() bool`

HasDatasourceUserId returns a boolean if a field has been set.

### GetName

`func (o *UserReferenceDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserReferenceDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserReferenceDefinition) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserReferenceDefinition) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


