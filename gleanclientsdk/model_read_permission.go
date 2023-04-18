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

// checks if the ReadPermission type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReadPermission{}

// ReadPermission Describes the read permission level that a user has for a specific feature
type ReadPermission struct {
	ScopeType *ScopeType `json:"scopeType,omitempty"`
}

// NewReadPermission instantiates a new ReadPermission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadPermission() *ReadPermission {
	this := ReadPermission{}
	return &this
}

// NewReadPermissionWithDefaults instantiates a new ReadPermission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadPermissionWithDefaults() *ReadPermission {
	this := ReadPermission{}
	return &this
}

// GetScopeType returns the ScopeType field value if set, zero value otherwise.
func (o *ReadPermission) GetScopeType() ScopeType {
	if o == nil || IsNil(o.ScopeType) {
		var ret ScopeType
		return ret
	}
	return *o.ScopeType
}

// GetScopeTypeOk returns a tuple with the ScopeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadPermission) GetScopeTypeOk() (*ScopeType, bool) {
	if o == nil || IsNil(o.ScopeType) {
		return nil, false
	}
	return o.ScopeType, true
}

// HasScopeType returns a boolean if a field has been set.
func (o *ReadPermission) HasScopeType() bool {
	if o != nil && !IsNil(o.ScopeType) {
		return true
	}

	return false
}

// SetScopeType gets a reference to the given ScopeType and assigns it to the ScopeType field.
func (o *ReadPermission) SetScopeType(v ScopeType) {
	o.ScopeType = &v
}

func (o ReadPermission) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReadPermission) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ScopeType) {
		toSerialize["scopeType"] = o.ScopeType
	}
	return toSerialize, nil
}

type NullableReadPermission struct {
	value *ReadPermission
	isSet bool
}

func (v NullableReadPermission) Get() *ReadPermission {
	return v.value
}

func (v *NullableReadPermission) Set(val *ReadPermission) {
	v.value = val
	v.isSet = true
}

func (v NullableReadPermission) IsSet() bool {
	return v.isSet
}

func (v *NullableReadPermission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadPermission(val *ReadPermission) *NullableReadPermission {
	return &NullableReadPermission{value: val, isSet: true}
}

func (v NullableReadPermission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadPermission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

