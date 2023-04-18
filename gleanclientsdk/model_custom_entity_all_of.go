/*
Glean Client API

# Introduction These are the public APIs to enable implementing a custom client interface to the Glean system.  # Usage guidelines This API is evolving fast. Glean will provide advance notice of any planned backwards incompatible changes along with a 6-month sunset period for anything that requires developers to adopt the new versions. 

API version: 0.9.0
Contact: support@glean.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gleanclientsdk

import (
	"encoding/json"
)

// checks if the CustomEntityAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomEntityAllOf{}

// CustomEntityAllOf struct for CustomEntityAllOf
type CustomEntityAllOf struct {
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
}

// NewCustomEntityAllOf instantiates a new CustomEntityAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomEntityAllOf() *CustomEntityAllOf {
	this := CustomEntityAllOf{}
	return &this
}

// NewCustomEntityAllOfWithDefaults instantiates a new CustomEntityAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomEntityAllOfWithDefaults() *CustomEntityAllOf {
	this := CustomEntityAllOf{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CustomEntityAllOf) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomEntityAllOf) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CustomEntityAllOf) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CustomEntityAllOf) SetId(v string) {
	o.Id = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *CustomEntityAllOf) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomEntityAllOf) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *CustomEntityAllOf) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *CustomEntityAllOf) SetTitle(v string) {
	o.Title = &v
}

// GetDatasource returns the Datasource field value if set, zero value otherwise.
func (o *CustomEntityAllOf) GetDatasource() string {
	if o == nil || IsNil(o.Datasource) {
		var ret string
		return ret
	}
	return *o.Datasource
}

// GetDatasourceOk returns a tuple with the Datasource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomEntityAllOf) GetDatasourceOk() (*string, bool) {
	if o == nil || IsNil(o.Datasource) {
		return nil, false
	}
	return o.Datasource, true
}

// HasDatasource returns a boolean if a field has been set.
func (o *CustomEntityAllOf) HasDatasource() bool {
	if o != nil && !IsNil(o.Datasource) {
		return true
	}

	return false
}

// SetDatasource gets a reference to the given string and assigns it to the Datasource field.
func (o *CustomEntityAllOf) SetDatasource(v string) {
	o.Datasource = &v
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *CustomEntityAllOf) GetObjectType() string {
	if o == nil || IsNil(o.ObjectType) {
		var ret string
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomEntityAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectType) {
		return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *CustomEntityAllOf) HasObjectType() bool {
	if o != nil && !IsNil(o.ObjectType) {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given string and assigns it to the ObjectType field.
func (o *CustomEntityAllOf) SetObjectType(v string) {
	o.ObjectType = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *CustomEntityAllOf) GetMetadata() CustomEntityMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret CustomEntityMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomEntityAllOf) GetMetadataOk() (*CustomEntityMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CustomEntityAllOf) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given CustomEntityMetadata and assigns it to the Metadata field.
func (o *CustomEntityAllOf) SetMetadata(v CustomEntityMetadata) {
	o.Metadata = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *CustomEntityAllOf) GetRoles() []UserRoleSpecification {
	if o == nil || IsNil(o.Roles) {
		var ret []UserRoleSpecification
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomEntityAllOf) GetRolesOk() ([]UserRoleSpecification, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *CustomEntityAllOf) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []UserRoleSpecification and assigns it to the Roles field.
func (o *CustomEntityAllOf) SetRoles(v []UserRoleSpecification) {
	o.Roles = v
}

func (o CustomEntityAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomEntityAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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
	return toSerialize, nil
}

type NullableCustomEntityAllOf struct {
	value *CustomEntityAllOf
	isSet bool
}

func (v NullableCustomEntityAllOf) Get() *CustomEntityAllOf {
	return v.value
}

func (v *NullableCustomEntityAllOf) Set(val *CustomEntityAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomEntityAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomEntityAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomEntityAllOf(val *CustomEntityAllOf) *NullableCustomEntityAllOf {
	return &NullableCustomEntityAllOf{value: val, isSet: true}
}

func (v NullableCustomEntityAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomEntityAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


