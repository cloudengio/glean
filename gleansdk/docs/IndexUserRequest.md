# IndexUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **int64** | Version number for document for optimistic concurrency control. If absent or 0 then no version checks are done. | [optional] 
**Datasource** | **string** | The datasource for which the user is added | 
**User** | [**DatasourceUserDefinition**](DatasourceUserDefinition.md) |  | 

## Methods

### NewIndexUserRequest

`func NewIndexUserRequest(datasource string, user DatasourceUserDefinition, ) *IndexUserRequest`

NewIndexUserRequest instantiates a new IndexUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndexUserRequestWithDefaults

`func NewIndexUserRequestWithDefaults() *IndexUserRequest`

NewIndexUserRequestWithDefaults instantiates a new IndexUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *IndexUserRequest) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *IndexUserRequest) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *IndexUserRequest) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *IndexUserRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetDatasource

`func (o *IndexUserRequest) GetDatasource() string`

GetDatasource returns the Datasource field if non-nil, zero value otherwise.

### GetDatasourceOk

`func (o *IndexUserRequest) GetDatasourceOk() (*string, bool)`

GetDatasourceOk returns a tuple with the Datasource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasource

`func (o *IndexUserRequest) SetDatasource(v string)`

SetDatasource sets Datasource field to given value.


### GetUser

`func (o *IndexUserRequest) GetUser() DatasourceUserDefinition`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *IndexUserRequest) GetUserOk() (*DatasourceUserDefinition, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *IndexUserRequest) SetUser(v DatasourceUserDefinition)`

SetUser sets User field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


