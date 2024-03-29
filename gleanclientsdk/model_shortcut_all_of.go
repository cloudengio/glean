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

// checks if the ShortcutAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShortcutAllOf{}

// ShortcutAllOf struct for ShortcutAllOf
type ShortcutAllOf struct {
	// canonical link text following go/ prefix where hyphen/underscore is removed.
	Alias *string `json:"alias,omitempty"`
	// Title for the Go Link
	Title *string `json:"title,omitempty"`
	// A list of user roles for the Go Link.
	Roles []UserRoleSpecification `json:"roles,omitempty"`
}

// NewShortcutAllOf instantiates a new ShortcutAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShortcutAllOf() *ShortcutAllOf {
	this := ShortcutAllOf{}
	return &this
}

// NewShortcutAllOfWithDefaults instantiates a new ShortcutAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShortcutAllOfWithDefaults() *ShortcutAllOf {
	this := ShortcutAllOf{}
	return &this
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *ShortcutAllOf) GetAlias() string {
	if o == nil || IsNil(o.Alias) {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShortcutAllOf) GetAliasOk() (*string, bool) {
	if o == nil || IsNil(o.Alias) {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *ShortcutAllOf) HasAlias() bool {
	if o != nil && !IsNil(o.Alias) {
		return true
	}

	return false
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *ShortcutAllOf) SetAlias(v string) {
	o.Alias = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ShortcutAllOf) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShortcutAllOf) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ShortcutAllOf) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ShortcutAllOf) SetTitle(v string) {
	o.Title = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *ShortcutAllOf) GetRoles() []UserRoleSpecification {
	if o == nil || IsNil(o.Roles) {
		var ret []UserRoleSpecification
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShortcutAllOf) GetRolesOk() ([]UserRoleSpecification, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *ShortcutAllOf) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []UserRoleSpecification and assigns it to the Roles field.
func (o *ShortcutAllOf) SetRoles(v []UserRoleSpecification) {
	o.Roles = v
}

func (o ShortcutAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShortcutAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Alias) {
		toSerialize["alias"] = o.Alias
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	return toSerialize, nil
}

type NullableShortcutAllOf struct {
	value *ShortcutAllOf
	isSet bool
}

func (v NullableShortcutAllOf) Get() *ShortcutAllOf {
	return v.value
}

func (v *NullableShortcutAllOf) Set(val *ShortcutAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableShortcutAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableShortcutAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShortcutAllOf(val *ShortcutAllOf) *NullableShortcutAllOf {
	return &NullableShortcutAllOf{value: val, isSet: true}
}

func (v NullableShortcutAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShortcutAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


