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

// checks if the RelatedObjectEdge type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RelatedObjectEdge{}

// RelatedObjectEdge struct for RelatedObjectEdge
type RelatedObjectEdge struct {
	Objects []RelatedObject `json:"objects,omitempty"`
}

// NewRelatedObjectEdge instantiates a new RelatedObjectEdge object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelatedObjectEdge() *RelatedObjectEdge {
	this := RelatedObjectEdge{}
	return &this
}

// NewRelatedObjectEdgeWithDefaults instantiates a new RelatedObjectEdge object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelatedObjectEdgeWithDefaults() *RelatedObjectEdge {
	this := RelatedObjectEdge{}
	return &this
}

// GetObjects returns the Objects field value if set, zero value otherwise.
func (o *RelatedObjectEdge) GetObjects() []RelatedObject {
	if o == nil || IsNil(o.Objects) {
		var ret []RelatedObject
		return ret
	}
	return o.Objects
}

// GetObjectsOk returns a tuple with the Objects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelatedObjectEdge) GetObjectsOk() ([]RelatedObject, bool) {
	if o == nil || IsNil(o.Objects) {
		return nil, false
	}
	return o.Objects, true
}

// HasObjects returns a boolean if a field has been set.
func (o *RelatedObjectEdge) HasObjects() bool {
	if o != nil && !IsNil(o.Objects) {
		return true
	}

	return false
}

// SetObjects gets a reference to the given []RelatedObject and assigns it to the Objects field.
func (o *RelatedObjectEdge) SetObjects(v []RelatedObject) {
	o.Objects = v
}

func (o RelatedObjectEdge) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RelatedObjectEdge) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Objects) {
		toSerialize["objects"] = o.Objects
	}
	return toSerialize, nil
}

type NullableRelatedObjectEdge struct {
	value *RelatedObjectEdge
	isSet bool
}

func (v NullableRelatedObjectEdge) Get() *RelatedObjectEdge {
	return v.value
}

func (v *NullableRelatedObjectEdge) Set(val *RelatedObjectEdge) {
	v.value = val
	v.isSet = true
}

func (v NullableRelatedObjectEdge) IsSet() bool {
	return v.isSet
}

func (v *NullableRelatedObjectEdge) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRelatedObjectEdge(val *RelatedObjectEdge) *NullableRelatedObjectEdge {
	return &NullableRelatedObjectEdge{value: val, isSet: true}
}

func (v NullableRelatedObjectEdge) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRelatedObjectEdge) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


