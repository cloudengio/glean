# DocumentPermissionsDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedUsers** | Pointer to [**[]UserReferenceDefinition**](UserReferenceDefinition.md) | List of users who can view the document | [optional] 
**AllowedGroups** | Pointer to **[]string** | List of groups that can view the document | [optional] 
**AllowedGroupIntersections** | Pointer to [**[]PermissionsGroupIntersectionDefinition**](PermissionsGroupIntersectionDefinition.md) | List of allowed group intersections. This describes a permissions constraint of the form ((GroupA AND GroupB AND GroupC) OR (GroupX AND GroupY) OR ... | [optional] 
**AllowAnonymousAccess** | Pointer to **bool** | If true, then any Glean user can view the document | [optional] 
**AllowAllDatasourceUsersAccess** | Pointer to **bool** | If true, then any user who has an account in the datasource can view the document. | [optional] 

## Methods

### NewDocumentPermissionsDefinition

`func NewDocumentPermissionsDefinition() *DocumentPermissionsDefinition`

NewDocumentPermissionsDefinition instantiates a new DocumentPermissionsDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentPermissionsDefinitionWithDefaults

`func NewDocumentPermissionsDefinitionWithDefaults() *DocumentPermissionsDefinition`

NewDocumentPermissionsDefinitionWithDefaults instantiates a new DocumentPermissionsDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedUsers

`func (o *DocumentPermissionsDefinition) GetAllowedUsers() []UserReferenceDefinition`

GetAllowedUsers returns the AllowedUsers field if non-nil, zero value otherwise.

### GetAllowedUsersOk

`func (o *DocumentPermissionsDefinition) GetAllowedUsersOk() (*[]UserReferenceDefinition, bool)`

GetAllowedUsersOk returns a tuple with the AllowedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedUsers

`func (o *DocumentPermissionsDefinition) SetAllowedUsers(v []UserReferenceDefinition)`

SetAllowedUsers sets AllowedUsers field to given value.

### HasAllowedUsers

`func (o *DocumentPermissionsDefinition) HasAllowedUsers() bool`

HasAllowedUsers returns a boolean if a field has been set.

### GetAllowedGroups

`func (o *DocumentPermissionsDefinition) GetAllowedGroups() []string`

GetAllowedGroups returns the AllowedGroups field if non-nil, zero value otherwise.

### GetAllowedGroupsOk

`func (o *DocumentPermissionsDefinition) GetAllowedGroupsOk() (*[]string, bool)`

GetAllowedGroupsOk returns a tuple with the AllowedGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedGroups

`func (o *DocumentPermissionsDefinition) SetAllowedGroups(v []string)`

SetAllowedGroups sets AllowedGroups field to given value.

### HasAllowedGroups

`func (o *DocumentPermissionsDefinition) HasAllowedGroups() bool`

HasAllowedGroups returns a boolean if a field has been set.

### GetAllowedGroupIntersections

`func (o *DocumentPermissionsDefinition) GetAllowedGroupIntersections() []PermissionsGroupIntersectionDefinition`

GetAllowedGroupIntersections returns the AllowedGroupIntersections field if non-nil, zero value otherwise.

### GetAllowedGroupIntersectionsOk

`func (o *DocumentPermissionsDefinition) GetAllowedGroupIntersectionsOk() (*[]PermissionsGroupIntersectionDefinition, bool)`

GetAllowedGroupIntersectionsOk returns a tuple with the AllowedGroupIntersections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedGroupIntersections

`func (o *DocumentPermissionsDefinition) SetAllowedGroupIntersections(v []PermissionsGroupIntersectionDefinition)`

SetAllowedGroupIntersections sets AllowedGroupIntersections field to given value.

### HasAllowedGroupIntersections

`func (o *DocumentPermissionsDefinition) HasAllowedGroupIntersections() bool`

HasAllowedGroupIntersections returns a boolean if a field has been set.

### GetAllowAnonymousAccess

`func (o *DocumentPermissionsDefinition) GetAllowAnonymousAccess() bool`

GetAllowAnonymousAccess returns the AllowAnonymousAccess field if non-nil, zero value otherwise.

### GetAllowAnonymousAccessOk

`func (o *DocumentPermissionsDefinition) GetAllowAnonymousAccessOk() (*bool, bool)`

GetAllowAnonymousAccessOk returns a tuple with the AllowAnonymousAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAnonymousAccess

`func (o *DocumentPermissionsDefinition) SetAllowAnonymousAccess(v bool)`

SetAllowAnonymousAccess sets AllowAnonymousAccess field to given value.

### HasAllowAnonymousAccess

`func (o *DocumentPermissionsDefinition) HasAllowAnonymousAccess() bool`

HasAllowAnonymousAccess returns a boolean if a field has been set.

### GetAllowAllDatasourceUsersAccess

`func (o *DocumentPermissionsDefinition) GetAllowAllDatasourceUsersAccess() bool`

GetAllowAllDatasourceUsersAccess returns the AllowAllDatasourceUsersAccess field if non-nil, zero value otherwise.

### GetAllowAllDatasourceUsersAccessOk

`func (o *DocumentPermissionsDefinition) GetAllowAllDatasourceUsersAccessOk() (*bool, bool)`

GetAllowAllDatasourceUsersAccessOk returns a tuple with the AllowAllDatasourceUsersAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAllDatasourceUsersAccess

`func (o *DocumentPermissionsDefinition) SetAllowAllDatasourceUsersAccess(v bool)`

SetAllowAllDatasourceUsersAccess sets AllowAllDatasourceUsersAccess field to given value.

### HasAllowAllDatasourceUsersAccess

`func (o *DocumentPermissionsDefinition) HasAllowAllDatasourceUsersAccess() bool`

HasAllowAllDatasourceUsersAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


