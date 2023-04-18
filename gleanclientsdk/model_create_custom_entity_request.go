/*
Glean Client API - Platform Preview

# Introduction These are all the APIs used by Glean to implement the Glean client. These are available as platform preview for implementing a custom client to the Glean system.  # Usage guidelines A subset of these endpoints are also in the developer ready section, which is available for public use. The rest of the endpoints are subject to prior agreement with Glean before usage. Please contact support@glean.com if you would like to use an API that is not currently available in the developer ready section. 

API version: 0.9.0
Contact: support@glean.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gleanclientsdk

import (
	"encoding/json"
)

// checks if the CreateCustomEntityRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateCustomEntityRequest{}

// CreateCustomEntityRequest struct for CreateCustomEntityRequest
type CreateCustomEntityRequest struct {
	Permissions *ObjectPermissions `json:"permissions,omitempty"`
	// Unique identifier.
	Id *string `json:"id,omitempty"`
	// Title or name of the custom entity.
	Title *string `json:"title,omitempty"`
	// The datasource the custom entity is from.
	Datasource *string `json:"datasource,omitempty"`
	// The type of the entity. Interpretation is specific to each datasource
	ObjectType *string `json:"objectType,omitempty"`
	Metadata *CustomEntityMetadata `json:"metadata,omitempty"`
	// A list of user roles for the custom entity explicitly granted by the owner.
	Roles []UserRoleSpecification `json:"roles,omitempty"`
	// A list of user roles for the custom entity added by the owner.
	AddedRoles []UserRoleSpecification `json:"addedRoles,omitempty"`
	// A list of user roles for the custom entity removed by the owner.
	RemovedRoles []UserRoleSpecification `json:"removedRoles,omitempty"`
}

// NewCreateCustomEntityRequest instantiates a new CreateCustomEntityRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCustomEntityRequest() *CreateCustomEntityRequest {
	this := CreateCustomEntityRequest{}
	return &this
}

// NewCreateCustomEntityRequestWithDefaults instantiates a new CreateCustomEntityRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCustomEntityRequestWithDefaults() *CreateCustomEntityRequest {
	this := CreateCustomEntityRequest{}
	return &this
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *CreateCustomEntityRequest) GetPermissions() ObjectPermissions {
	if o == nil || IsNil(o.Permissions) {
		var ret ObjectPermissions
		return ret
	}
	return *o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCustomEntityRequest) GetPermissionsOk() (*ObjectPermissions, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *CreateCustomEntityRequest) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given ObjectPermissions and assigns it to the Permissions field.
func (o *CreateCustomEntityRequest) SetPermissions(v ObjectPermissions) {
	o.Permissions = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CreateCustomEntityRequest) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCustomEntityRequest) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CreateCustomEntityRequest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CreateCustomEntityRequest) SetId(v string) {
	o.Id = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *CreateCustomEntityRequest) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCustomEntityRequest) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *CreateCustomEntityRequest) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *CreateCustomEntityRequest) SetTitle(v string) {
	o.Title = &v
}

// GetDatasource returns the Datasource field value if set, zero value otherwise.
func (o *CreateCustomEntityRequest) GetDatasource() string {
	if o == nil || IsNil(o.Datasource) {
		var ret string
		return ret
	}
	return *o.Datasource
}

// GetDatasourceOk returns a tuple with the Datasource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCustomEntityRequest) GetDatasourceOk() (*string, bool) {
	if o == nil || IsNil(o.Datasource) {
		return nil, false
	}
	return o.Datasource, true
}

// HasDatasource returns a boolean if a field has been set.
func (o *CreateCustomEntityRequest) HasDatasource() bool {
	if o != nil && !IsNil(o.Datasource) {
		return true
	}

	return false
}

// SetDatasource gets a reference to the given string and assigns it to the Datasource field.
func (o *CreateCustomEntityRequest) SetDatasource(v string) {
	o.Datasource = &v
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *CreateCustomEntityRequest) GetObjectType() string {
	if o == nil || IsNil(o.ObjectType) {
		var ret string
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCustomEntityRequest) GetObjectTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectType) {
		return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *CreateCustomEntityRequest) HasObjectType() bool {
	if o != nil && !IsNil(o.ObjectType) {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given string and assigns it to the ObjectType field.
func (o *CreateCustomEntityRequest) SetObjectType(v string) {
	o.ObjectType = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *CreateCustomEntityRequest) GetMetadata() CustomEntityMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret CustomEntityMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCustomEntityRequest) GetMetadataOk() (*CustomEntityMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CreateCustomEntityRequest) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given CustomEntityMetadata and assigns it to the Metadata field.
func (o *CreateCustomEntityRequest) SetMetadata(v CustomEntityMetadata) {
	o.Metadata = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *CreateCustomEntityRequest) GetRoles() []UserRoleSpecification {
	if o == nil || IsNil(o.Roles) {
		var ret []UserRoleSpecification
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCustomEntityRequest) GetRolesOk() ([]UserRoleSpecification, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *CreateCustomEntityRequest) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []UserRoleSpecification and assigns it to the Roles field.
func (o *CreateCustomEntityRequest) SetRoles(v []UserRoleSpecification) {
	o.Roles = v
}

// GetAddedRoles returns the AddedRoles field value if set, zero value otherwise.
func (o *CreateCustomEntityRequest) GetAddedRoles() []UserRoleSpecification {
	if o == nil || IsNil(o.AddedRoles) {
		var ret []UserRoleSpecification
		return ret
	}
	return o.AddedRoles
}

// GetAddedRolesOk returns a tuple with the AddedRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCustomEntityRequest) GetAddedRolesOk() ([]UserRoleSpecification, bool) {
	if o == nil || IsNil(o.AddedRoles) {
		return nil, false
	}
	return o.AddedRoles, true
}

// HasAddedRoles returns a boolean if a field has been set.
func (o *CreateCustomEntityRequest) HasAddedRoles() bool {
	if o != nil && !IsNil(o.AddedRoles) {
		return true
	}

	return false
}

// SetAddedRoles gets a reference to the given []UserRoleSpecification and assigns it to the AddedRoles field.
func (o *CreateCustomEntityRequest) SetAddedRoles(v []UserRoleSpecification) {
	o.AddedRoles = v
}

// GetRemovedRoles returns the RemovedRoles field value if set, zero value otherwise.
func (o *CreateCustomEntityRequest) GetRemovedRoles() []UserRoleSpecification {
	if o == nil || IsNil(o.RemovedRoles) {
		var ret []UserRoleSpecification
		return ret
	}
	return o.RemovedRoles
}

// GetRemovedRolesOk returns a tuple with the RemovedRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCustomEntityRequest) GetRemovedRolesOk() ([]UserRoleSpecification, bool) {
	if o == nil || IsNil(o.RemovedRoles) {
		return nil, false
	}
	return o.RemovedRoles, true
}

// HasRemovedRoles returns a boolean if a field has been set.
func (o *CreateCustomEntityRequest) HasRemovedRoles() bool {
	if o != nil && !IsNil(o.RemovedRoles) {
		return true
	}

	return false
}

// SetRemovedRoles gets a reference to the given []UserRoleSpecification and assigns it to the RemovedRoles field.
func (o *CreateCustomEntityRequest) SetRemovedRoles(v []UserRoleSpecification) {
	o.RemovedRoles = v
}

func (o CreateCustomEntityRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateCustomEntityRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Datasource) {
		toSerialize["datasource"] = o.Datasource
	}
	if !IsNil(o.ObjectType) {
		toSerialize["objectType"] = o.ObjectType
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	if !IsNil(o.AddedRoles) {
		toSerialize["addedRoles"] = o.AddedRoles
	}
	if !IsNil(o.RemovedRoles) {
		toSerialize["removedRoles"] = o.RemovedRoles
	}
	return toSerialize, nil
}

type NullableCreateCustomEntityRequest struct {
	value *CreateCustomEntityRequest
	isSet bool
}

func (v NullableCreateCustomEntityRequest) Get() *CreateCustomEntityRequest {
	return v.value
}

func (v *NullableCreateCustomEntityRequest) Set(val *CreateCustomEntityRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCustomEntityRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCustomEntityRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCustomEntityRequest(val *CreateCustomEntityRequest) *NullableCreateCustomEntityRequest {
	return &NullableCreateCustomEntityRequest{value: val, isSet: true}
}

func (v NullableCreateCustomEntityRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCustomEntityRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


