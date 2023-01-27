# ProcessAllMembershipsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datasource** | Pointer to **string** | If provided, process group memberships only for this custom datasource. Otherwise all uploaded memberships are processed. | [optional] 

## Methods

### NewProcessAllMembershipsRequest

`func NewProcessAllMembershipsRequest() *ProcessAllMembershipsRequest`

NewProcessAllMembershipsRequest instantiates a new ProcessAllMembershipsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessAllMembershipsRequestWithDefaults

`func NewProcessAllMembershipsRequestWithDefaults() *ProcessAllMembershipsRequest`

NewProcessAllMembershipsRequestWithDefaults instantiates a new ProcessAllMembershipsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatasource

`func (o *ProcessAllMembershipsRequest) GetDatasource() string`

GetDatasource returns the Datasource field if non-nil, zero value otherwise.

### GetDatasourceOk

`func (o *ProcessAllMembershipsRequest) GetDatasourceOk() (*string, bool)`

GetDatasourceOk returns a tuple with the Datasource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasource

`func (o *ProcessAllMembershipsRequest) SetDatasource(v string)`

SetDatasource sets Datasource field to given value.

### HasDatasource

`func (o *ProcessAllMembershipsRequest) HasDatasource() bool`

HasDatasource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


