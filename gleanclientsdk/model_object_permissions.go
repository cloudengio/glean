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

// checks if the ObjectPermissions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ObjectPermissions{}

// ObjectPermissions struct for ObjectPermissions
type ObjectPermissions struct {
	Write *WritePermission `json:"write,omitempty"`
}

// NewObjectPermissions instantiates a new ObjectPermissions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectPermissions() *ObjectPermissions {
	this := ObjectPermissions{}
	return &this
}

// NewObjectPermissionsWithDefaults instantiates a new ObjectPermissions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectPermissionsWithDefaults() *ObjectPermissions {
	this := ObjectPermissions{}
	return &this
}

// GetWrite returns the Write field value if set, zero value otherwise.
func (o *ObjectPermissions) GetWrite() WritePermission {
	if o == nil || IsNil(o.Write) {
		var ret WritePermission
		return ret
	}
	return *o.Write
}

// GetWriteOk returns a tuple with the Write field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectPermissions) GetWriteOk() (*WritePermission, bool) {
	if o == nil || IsNil(o.Write) {
		return nil, false
	}
	return o.Write, true
}

// HasWrite returns a boolean if a field has been set.
func (o *ObjectPermissions) HasWrite() bool {
	if o != nil && !IsNil(o.Write) {
		return true
	}

	return false
}

// SetWrite gets a reference to the given WritePermission and assigns it to the Write field.
func (o *ObjectPermissions) SetWrite(v WritePermission) {
	o.Write = &v
}

func (o ObjectPermissions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ObjectPermissions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Write) {
		toSerialize["write"] = o.Write
	}
	return toSerialize, nil
}

type NullableObjectPermissions struct {
	value *ObjectPermissions
	isSet bool
}

func (v NullableObjectPermissions) Get() *ObjectPermissions {
	return v.value
}

func (v *NullableObjectPermissions) Set(val *ObjectPermissions) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectPermissions) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectPermissions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectPermissions(val *ObjectPermissions) *NullableObjectPermissions {
	return &NullableObjectPermissions{value: val, isSet: true}
}

func (v NullableObjectPermissions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectPermissions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


