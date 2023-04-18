# CustomEntityAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier. | [optional] 
**Title** | Pointer to **string** | Title or name of the custom entity. | [optional] 
**Datasource** | Pointer to **string** | The datasource the custom entity is from. | [optional] 
**ObjectType** | Pointer to **string** | The type of the entity. Interpretation is specific to each datasource | [optional] 
**Metadata** | Pointer to [**CustomEntityMetadata**](CustomEntityMetadata.md) |  | [optional] 
**Roles** | Pointer to [**[]UserRoleSpecification**](UserRoleSpecification.md) | A list of user roles for the custom entity explicitly granted by the owner. | [optional] 

## Methods

### NewCustomEntityAllOf

`func NewCustomEntityAllOf() *CustomEntityAllOf`

NewCustomEntityAllOf instantiates a new CustomEntityAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomEntityAllOfWithDefaults

`func NewCustomEntityAllOfWithDefaults() *CustomEntityAllOf`

NewCustomEntityAllOfWithDefaults instantiates a new CustomEntityAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomEntityAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomEntityAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomEntityAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustomEntityAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *CustomEntityAllOf) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CustomEntityAllOf) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CustomEntityAllOf) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CustomEntityAllOf) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDatasource

`func (o *CustomEntityAllOf) GetDatasource() string`

GetDatasource returns the Datasource field if non-nil, zero value otherwise.

### GetDatasourceOk

`func (o *CustomEntityAllOf) GetDatasourceOk() (*string, bool)`

GetDatasourceOk returns a tuple with the Datasource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasource

`func (o *CustomEntityAllOf) SetDatasource(v string)`

SetDatasource sets Datasource field to given value.

### HasDatasource

`func (o *CustomEntityAllOf) HasDatasource() bool`

HasDatasource returns a boolean if a field has been set.

### GetObjectType

`func (o *CustomEntityAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CustomEntityAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CustomEntityAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *CustomEntityAllOf) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetMetadata

`func (o *CustomEntityAllOf) GetMetadata() CustomEntityMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CustomEntityAllOf) GetMetadataOk() (*CustomEntityMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CustomEntityAllOf) SetMetadata(v CustomEntityMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CustomEntityAllOf) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetRoles

`func (o *CustomEntityAllOf) GetRoles() []UserRoleSpecification`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *CustomEntityAllOf) GetRolesOk() (*[]UserRoleSpecification, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *CustomEntityAllOf) SetRoles(v []UserRoleSpecification)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *CustomEntityAllOf) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


