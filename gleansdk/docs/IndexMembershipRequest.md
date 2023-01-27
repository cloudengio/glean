# IndexMembershipRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **int64** | Version number for document for optimistic concurrency control. If absent or 0 then no version checks are done. | [optional] 
**Datasource** | **string** | The datasource for which the membership is added | 
**Membership** | [**DatasourceMembershipDefinition**](DatasourceMembershipDefinition.md) |  | 

## Methods

### NewIndexMembershipRequest

`func NewIndexMembershipRequest(datasource string, membership DatasourceMembershipDefinition, ) *IndexMembershipRequest`

NewIndexMembershipRequest instantiates a new IndexMembershipRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndexMembershipRequestWithDefaults

`func NewIndexMembershipRequestWithDefaults() *IndexMembershipRequest`

NewIndexMembershipRequestWithDefaults instantiates a new IndexMembershipRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *IndexMembershipRequest) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *IndexMembershipRequest) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *IndexMembershipRequest) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *IndexMembershipRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetDatasource

`func (o *IndexMembershipRequest) GetDatasource() string`

GetDatasource returns the Datasource field if non-nil, zero value otherwise.

### GetDatasourceOk

`func (o *IndexMembershipRequest) GetDatasourceOk() (*string, bool)`

GetDatasourceOk returns a tuple with the Datasource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasource

`func (o *IndexMembershipRequest) SetDatasource(v string)`

SetDatasource sets Datasource field to given value.


### GetMembership

`func (o *IndexMembershipRequest) GetMembership() DatasourceMembershipDefinition`

GetMembership returns the Membership field if non-nil, zero value otherwise.

### GetMembershipOk

`func (o *IndexMembershipRequest) GetMembershipOk() (*DatasourceMembershipDefinition, bool)`

GetMembershipOk returns a tuple with the Membership field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembership

`func (o *IndexMembershipRequest) SetMembership(v DatasourceMembershipDefinition)`

SetMembership sets Membership field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


