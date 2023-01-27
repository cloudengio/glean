# IndexTeamRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Team** | [**TeamInfoDefinition**](TeamInfoDefinition.md) |  | 
**Version** | Pointer to **int64** | Version number for the team object. If absent or 0 then no version checks are done | [optional] 

## Methods

### NewIndexTeamRequest

`func NewIndexTeamRequest(team TeamInfoDefinition, ) *IndexTeamRequest`

NewIndexTeamRequest instantiates a new IndexTeamRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndexTeamRequestWithDefaults

`func NewIndexTeamRequestWithDefaults() *IndexTeamRequest`

NewIndexTeamRequestWithDefaults instantiates a new IndexTeamRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTeam

`func (o *IndexTeamRequest) GetTeam() TeamInfoDefinition`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *IndexTeamRequest) GetTeamOk() (*TeamInfoDefinition, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *IndexTeamRequest) SetTeam(v TeamInfoDefinition)`

SetTeam sets Team field to given value.


### GetVersion

`func (o *IndexTeamRequest) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *IndexTeamRequest) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *IndexTeamRequest) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *IndexTeamRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


