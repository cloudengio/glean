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

// checks if the UserGeneratedContentId type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserGeneratedContentId{}

// UserGeneratedContentId struct for UserGeneratedContentId
type UserGeneratedContentId struct {
	// The opaque id of the user generated content.
	Id *int32 `json:"id,omitempty"`
}

// NewUserGeneratedContentId instantiates a new UserGeneratedContentId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserGeneratedContentId() *UserGeneratedContentId {
	this := UserGeneratedContentId{}
	return &this
}

// NewUserGeneratedContentIdWithDefaults instantiates a new UserGeneratedContentId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserGeneratedContentIdWithDefaults() *UserGeneratedContentId {
	this := UserGeneratedContentId{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UserGeneratedContentId) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserGeneratedContentId) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UserGeneratedContentId) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *UserGeneratedContentId) SetId(v int32) {
	o.Id = &v
}

func (o UserGeneratedContentId) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserGeneratedContentId) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableUserGeneratedContentId struct {
	value *UserGeneratedContentId
	isSet bool
}

func (v NullableUserGeneratedContentId) Get() *UserGeneratedContentId {
	return v.value
}

func (v *NullableUserGeneratedContentId) Set(val *UserGeneratedContentId) {
	v.value = val
	v.isSet = true
}

func (v NullableUserGeneratedContentId) IsSet() bool {
	return v.isSet
}

func (v *NullableUserGeneratedContentId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserGeneratedContentId(val *UserGeneratedContentId) *NullableUserGeneratedContentId {
	return &NullableUserGeneratedContentId{value: val, isSet: true}
}

func (v NullableUserGeneratedContentId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserGeneratedContentId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


