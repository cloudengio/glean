# IndexGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **int64** | Version number for document for optimistic concurrency control. If absent or 0 then no version checks are done. | [optional] 
**Datasource** | **string** | The datasource for which the group is added | 
**Group** | [**DatasourceGroupDefinition**](DatasourceGroupDefinition.md) |  | 

## Methods

### NewIndexGroupRequest

`func NewIndexGroupRequest(datasource string, group DatasourceGroupDefinition, ) *IndexGroupRequest`

NewIndexGroupRequest instantiates a new IndexGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndexGroupRequestWithDefaults

`func NewIndexGroupRequestWithDefaults() *IndexGroupRequest`

NewIndexGroupRequestWithDefaults instantiates a new IndexGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *IndexGroupRequest) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *IndexGroupRequest) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *IndexGroupRequest) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *IndexGroupRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetDatasource

`func (o *IndexGroupRequest) GetDatasource() string`

GetDatasource returns the Datasource field if non-nil, zero value otherwise.

### GetDatasourceOk

`func (o *IndexGroupRequest) GetDatasourceOk() (*string, bool)`

GetDatasourceOk returns a tuple with the Datasource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasource

`func (o *IndexGroupRequest) SetDatasource(v string)`

SetDatasource sets Datasource field to given value.


### GetGroup

`func (o *IndexGroupRequest) GetGroup() DatasourceGroupDefinition`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *IndexGroupRequest) GetGroupOk() (*DatasourceGroupDefinition, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *IndexGroupRequest) SetGroup(v DatasourceGroupDefinition)`

SetGroup sets Group field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


