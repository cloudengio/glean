# DeleteUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **int64** | Version number for document for optimistic concurrency control. If absent or 0 then no version checks are done. | [optional] 
**Datasource** | **string** | The datasource for which the user is removed | 
**Email** | **string** | The email of the user to be deleted | 

## Methods

### NewDeleteUserRequest

`func NewDeleteUserRequest(datasource string, email string, ) *DeleteUserRequest`

NewDeleteUserRequest instantiates a new DeleteUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteUserRequestWithDefaults

`func NewDeleteUserRequestWithDefaults() *DeleteUserRequest`

NewDeleteUserRequestWithDefaults instantiates a new DeleteUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *DeleteUserRequest) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DeleteUserRequest) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DeleteUserRequest) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DeleteUserRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetDatasource

`func (o *DeleteUserRequest) GetDatasource() string`

GetDatasource returns the Datasource field if non-nil, zero value otherwise.

### GetDatasourceOk

`func (o *DeleteUserRequest) GetDatasourceOk() (*string, bool)`

GetDatasourceOk returns a tuple with the Datasource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasource

`func (o *DeleteUserRequest) SetDatasource(v string)`

SetDatasource sets Datasource field to given value.


### GetEmail

`func (o *DeleteUserRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *DeleteUserRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *DeleteUserRequest) SetEmail(v string)`

SetEmail sets Email field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


