# DatasourceMembershipDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupName** | **string** | The group for which the membership is specified | 
**MemberUserId** | Pointer to **string** | If the member is a user, then the email or datasource id for the user | [optional] 
**MemberGroupName** | Pointer to **string** | If the member is a group, then the name of the member group | [optional] 

## Methods

### NewDatasourceMembershipDefinition

`func NewDatasourceMembershipDefinition(groupName string, ) *DatasourceMembershipDefinition`

NewDatasourceMembershipDefinition instantiates a new DatasourceMembershipDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatasourceMembershipDefinitionWithDefaults

`func NewDatasourceMembershipDefinitionWithDefaults() *DatasourceMembershipDefinition`

NewDatasourceMembershipDefinitionWithDefaults instantiates a new DatasourceMembershipDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupName

`func (o *DatasourceMembershipDefinition) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *DatasourceMembershipDefinition) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *DatasourceMembershipDefinition) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.


### GetMemberUserId

`func (o *DatasourceMembershipDefinition) GetMemberUserId() string`

GetMemberUserId returns the MemberUserId field if non-nil, zero value otherwise.

### GetMemberUserIdOk

`func (o *DatasourceMembershipDefinition) GetMemberUserIdOk() (*string, bool)`

GetMemberUserIdOk returns a tuple with the MemberUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberUserId

`func (o *DatasourceMembershipDefinition) SetMemberUserId(v string)`

SetMemberUserId sets MemberUserId field to given value.

### HasMemberUserId

`func (o *DatasourceMembershipDefinition) HasMemberUserId() bool`

HasMemberUserId returns a boolean if a field has been set.

### GetMemberGroupName

`func (o *DatasourceMembershipDefinition) GetMemberGroupName() string`

GetMemberGroupName returns the MemberGroupName field if non-nil, zero value otherwise.

### GetMemberGroupNameOk

`func (o *DatasourceMembershipDefinition) GetMemberGroupNameOk() (*string, bool)`

GetMemberGroupNameOk returns a tuple with the MemberGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberGroupName

`func (o *DatasourceMembershipDefinition) SetMemberGroupName(v string)`

SetMemberGroupName sets MemberGroupName field to given value.

### HasMemberGroupName

`func (o *DatasourceMembershipDefinition) HasMemberGroupName() bool`

HasMemberGroupName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


