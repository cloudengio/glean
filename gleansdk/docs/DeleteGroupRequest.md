# DeleteGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **int64** | Version number for document for optimistic concurrency control. If absent or 0 then no version checks are done. | [optional] 
**Datasource** | **string** | The datasource for which the group is removed | 
**GroupName** | **string** | the name of the group to be deleted | 

## Methods

### NewDeleteGroupRequest

`func NewDeleteGroupRequest(datasource string, groupName string, ) *DeleteGroupRequest`

NewDeleteGroupRequest instantiates a new DeleteGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteGroupRequestWithDefaults

`func NewDeleteGroupRequestWithDefaults() *DeleteGroupRequest`

NewDeleteGroupRequestWithDefaults instantiates a new DeleteGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *DeleteGroupRequest) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DeleteGroupRequest) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DeleteGroupRequest) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DeleteGroupRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetDatasource

`func (o *DeleteGroupRequest) GetDatasource() string`

GetDatasource returns the Datasource field if non-nil, zero value otherwise.

### GetDatasourceOk

`func (o *DeleteGroupRequest) GetDatasourceOk() (*string, bool)`

GetDatasourceOk returns a tuple with the Datasource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasource

`func (o *DeleteGroupRequest) SetDatasource(v string)`

SetDatasource sets Datasource field to given value.


### GetGroupName

`func (o *DeleteGroupRequest) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *DeleteGroupRequest) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *DeleteGroupRequest) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


