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

// checks if the RelatedObjects type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RelatedObjects{}

// RelatedObjects struct for RelatedObjects
type RelatedObjects struct {
	// A list of objects related to a source object.
	RelatedObjects *map[string]RelatedObjectEdge `json:"relatedObjects,omitempty"`
}

// NewRelatedObjects instantiates a new RelatedObjects object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelatedObjects() *RelatedObjects {
	this := RelatedObjects{}
	return &this
}

// NewRelatedObjectsWithDefaults instantiates a new RelatedObjects object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelatedObjectsWithDefaults() *RelatedObjects {
	this := RelatedObjects{}
	return &this
}

// GetRelatedObjects returns the RelatedObjects field value if set, zero value otherwise.
func (o *RelatedObjects) GetRelatedObjects() map[string]RelatedObjectEdge {
	if o == nil || IsNil(o.RelatedObjects) {
		var ret map[string]RelatedObjectEdge
		return ret
	}
	return *o.RelatedObjects
}

// GetRelatedObjectsOk returns a tuple with the RelatedObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelatedObjects) GetRelatedObjectsOk() (*map[string]RelatedObjectEdge, bool) {
	if o == nil || IsNil(o.RelatedObjects) {
		return nil, false
	}
	return o.RelatedObjects, true
}

// HasRelatedObjects returns a boolean if a field has been set.
func (o *RelatedObjects) HasRelatedObjects() bool {
	if o != nil && !IsNil(o.RelatedObjects) {
		return true
	}

	return false
}

// SetRelatedObjects gets a reference to the given map[string]RelatedObjectEdge and assigns it to the RelatedObjects field.
func (o *RelatedObjects) SetRelatedObjects(v map[string]RelatedObjectEdge) {
	o.RelatedObjects = &v
}

func (o RelatedObjects) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RelatedObjects) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RelatedObjects) {
		toSerialize["relatedObjects"] = o.RelatedObjects
	}
	return toSerialize, nil
}

type NullableRelatedObjects struct {
	value *RelatedObjects
	isSet bool
}

func (v NullableRelatedObjects) Get() *RelatedObjects {
	return v.value
}

func (v *NullableRelatedObjects) Set(val *RelatedObjects) {
	v.value = val
	v.isSet = true
}

func (v NullableRelatedObjects) IsSet() bool {
	return v.isSet
}

func (v *NullableRelatedObjects) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRelatedObjects(val *RelatedObjects) *NullableRelatedObjects {
	return &NullableRelatedObjects{value: val, isSet: true}
}

func (v NullableRelatedObjects) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRelatedObjects) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


