# CreateCustomEntityRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Permissions** | Pointer to [**ObjectPermissions**](ObjectPermissions.md) |  | [optional] 
**Id** | Pointer to **string** | Unique identifier. | [optional] 
**Title** | Pointer to **string** | Title or name of the custom entity. | [optional] 
**Datasource** | Pointer to **string** | The datasource the custom entity is from. | [optional] 
**ObjectType** | Pointer to **string** | The type of the entity. Interpretation is specific to each datasource | [optional] 
**Metadata** | Pointer to [**CustomEntityMetadata**](CustomEntityMetadata.md) |  | [optional] 
**Roles** | Pointer to [**[]UserRoleSpecification**](UserRoleSpecification.md) | A list of user roles for the custom entity explicitly granted by the owner. | [optional] 
**AddedRoles** | Pointer to [**[]UserRoleSpecification**](UserRoleSpecification.md) | A list of user roles for the custom entity added by the owner. | [optional] 
**RemovedRoles** | Pointer to [**[]UserRoleSpecification**](UserRoleSpecification.md) | A list of user roles for the custom entity removed by the owner. | [optional] 

## Methods

### NewCreateCustomEntityRequest

`func NewCreateCustomEntityRequest() *CreateCustomEntityRequest`

NewCreateCustomEntityRequest instantiates a new CreateCustomEntityRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCustomEntityRequestWithDefaults

`func NewCreateCustomEntityRequestWithDefaults() *CreateCustomEntityRequest`

NewCreateCustomEntityRequestWithDefaults instantiates a new CreateCustomEntityRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermissions

`func (o *CreateCustomEntityRequest) GetPermissions() ObjectPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *CreateCustomEntityRequest) GetPermissionsOk() (*ObjectPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *CreateCustomEntityRequest) SetPermissions(v ObjectPermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *CreateCustomEntityRequest) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetId

`func (o *CreateCustomEntityRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateCustomEntityRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateCustomEntityRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateCustomEntityRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *CreateCustomEntityRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateCustomEntityRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateCustomEntityRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CreateCustomEntityRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDatasource

`func (o *CreateCustomEntityRequest) GetDatasource() string`

GetDatasource returns the Datasource field if non-nil, zero value otherwise.

### GetDatasourceOk

`func (o *CreateCustomEntityRequest) GetDatasourceOk() (*string, bool)`

GetDatasourceOk returns a tuple with the Datasource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasource

`func (o *CreateCustomEntityRequest) SetDatasource(v string)`

SetDatasource sets Datasource field to given value.

### HasDatasource

`func (o *CreateCustomEntityRequest) HasDatasource() bool`

HasDatasource returns a boolean if a field has been set.

### GetObjectType

`func (o *CreateCustomEntityRequest) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CreateCustomEntityRequest) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CreateCustomEntityRequest) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *CreateCustomEntityRequest) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateCustomEntityRequest) GetMetadata() CustomEntityMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateCustomEntityRequest) GetMetadataOk() (*CustomEntityMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateCustomEntityRequest) SetMetadata(v CustomEntityMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateCustomEntityRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetRoles

`func (o *CreateCustomEntityRequest) GetRoles() []UserRoleSpecification`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *CreateCustomEntityRequest) GetRolesOk() (*[]UserRoleSpecification, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *CreateCustomEntityRequest) SetRoles(v []UserRoleSpecification)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *CreateCustomEntityRequest) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetAddedRoles

`func (o *CreateCustomEntityRequest) GetAddedRoles() []UserRoleSpecification`

GetAddedRoles returns the AddedRoles field if non-nil, zero value otherwise.

### GetAddedRolesOk

`func (o *CreateCustomEntityRequest) GetAddedRolesOk() (*[]UserRoleSpecification, bool)`

GetAddedRolesOk returns a tuple with the AddedRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedRoles

`func (o *CreateCustomEntityRequest) SetAddedRoles(v []UserRoleSpecification)`

SetAddedRoles sets AddedRoles field to given value.

### HasAddedRoles

`func (o *CreateCustomEntityRequest) HasAddedRoles() bool`

HasAddedRoles returns a boolean if a field has been set.

### GetRemovedRoles

`func (o *CreateCustomEntityRequest) GetRemovedRoles() []UserRoleSpecification`

GetRemovedRoles returns the RemovedRoles field if non-nil, zero value otherwise.

### GetRemovedRolesOk

`func (o *CreateCustomEntityRequest) GetRemovedRolesOk() (*[]UserRoleSpecification, bool)`

GetRemovedRolesOk returns a tuple with the RemovedRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovedRoles

`func (o *CreateCustomEntityRequest) SetRemovedRoles(v []UserRoleSpecification)`

SetRemovedRoles sets RemovedRoles field to given value.

### HasRemovedRoles

`func (o *CreateCustomEntityRequest) HasRemovedRoles() bool`

HasRemovedRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


