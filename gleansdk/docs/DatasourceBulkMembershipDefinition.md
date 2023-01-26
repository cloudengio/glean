# DatasourceBulkMembershipDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MemberUserId** | Pointer to **string** | If the member is a user, then the email or datasource id for the user | [optional] 
**MemberGroupName** | Pointer to **string** | If the member is a group, then the name of the member group | [optional] 

## Methods

### NewDatasourceBulkMembershipDefinition

`func NewDatasourceBulkMembershipDefinition() *DatasourceBulkMembershipDefinition`

NewDatasourceBulkMembershipDefinition instantiates a new DatasourceBulkMembershipDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatasourceBulkMembershipDefinitionWithDefaults

`func NewDatasourceBulkMembershipDefinitionWithDefaults() *DatasourceBulkMembershipDefinition`

NewDatasourceBulkMembershipDefinitionWithDefaults instantiates a new DatasourceBulkMembershipDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMemberUserId

`func (o *DatasourceBulkMembershipDefinition) GetMemberUserId() string`

GetMemberUserId returns the MemberUserId field if non-nil, zero value otherwise.

### GetMemberUserIdOk

`func (o *DatasourceBulkMembershipDefinition) GetMemberUserIdOk() (*string, bool)`

GetMemberUserIdOk returns a tuple with the MemberUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberUserId

`func (o *DatasourceBulkMembershipDefinition) SetMemberUserId(v string)`

SetMemberUserId sets MemberUserId field to given value.

### HasMemberUserId

`func (o *DatasourceBulkMembershipDefinition) HasMemberUserId() bool`

HasMemberUserId returns a boolean if a field has been set.

### GetMemberGroupName

`func (o *DatasourceBulkMembershipDefinition) GetMemberGroupName() string`

GetMemberGroupName returns the MemberGroupName field if non-nil, zero value otherwise.

### GetMemberGroupNameOk

`func (o *DatasourceBulkMembershipDefinition) GetMemberGroupNameOk() (*string, bool)`

GetMemberGroupNameOk returns a tuple with the MemberGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberGroupName

`func (o *DatasourceBulkMembershipDefinition) SetMemberGroupName(v string)`

SetMemberGroupName sets MemberGroupName field to given value.

### HasMemberGroupName

`func (o *DatasourceBulkMembershipDefinition) HasMemberGroupName() bool`

HasMemberGroupName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


