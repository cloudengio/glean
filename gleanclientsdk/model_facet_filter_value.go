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

// checks if the FacetFilterValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FacetFilterValue{}

// FacetFilterValue struct for FacetFilterValue
type FacetFilterValue struct {
	Value *string `json:"value,omitempty"`
	RelationType *string `json:"relationType,omitempty"`
	IsNegated *bool `json:"isNegated,omitempty"`
}

// NewFacetFilterValue instantiates a new FacetFilterValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFacetFilterValue() *FacetFilterValue {
	this := FacetFilterValue{}
	return &this
}

// NewFacetFilterValueWithDefaults instantiates a new FacetFilterValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFacetFilterValueWithDefaults() *FacetFilterValue {
	this := FacetFilterValue{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *FacetFilterValue) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FacetFilterValue) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *FacetFilterValue) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *FacetFilterValue) SetValue(v string) {
	o.Value = &v
}

// GetRelationType returns the RelationType field value if set, zero value otherwise.
func (o *FacetFilterValue) GetRelationType() string {
	if o == nil || IsNil(o.RelationType) {
		var ret string
		return ret
	}
	return *o.RelationType
}

// GetRelationTypeOk returns a tuple with the RelationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FacetFilterValue) GetRelationTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RelationType) {
		return nil, false
	}
	return o.RelationType, true
}

// HasRelationType returns a boolean if a field has been set.
func (o *FacetFilterValue) HasRelationType() bool {
	if o != nil && !IsNil(o.RelationType) {
		return true
	}

	return false
}

// SetRelationType gets a reference to the given string and assigns it to the RelationType field.
func (o *FacetFilterValue) SetRelationType(v string) {
	o.RelationType = &v
}

// GetIsNegated returns the IsNegated field value if set, zero value otherwise.
func (o *FacetFilterValue) GetIsNegated() bool {
	if o == nil || IsNil(o.IsNegated) {
		var ret bool
		return ret
	}
	return *o.IsNegated
}

// GetIsNegatedOk returns a tuple with the IsNegated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FacetFilterValue) GetIsNegatedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsNegated) {
		return nil, false
	}
	return o.IsNegated, true
}

// HasIsNegated returns a boolean if a field has been set.
func (o *FacetFilterValue) HasIsNegated() bool {
	if o != nil && !IsNil(o.IsNegated) {
		return true
	}

	return false
}

// SetIsNegated gets a reference to the given bool and assigns it to the IsNegated field.
func (o *FacetFilterValue) SetIsNegated(v bool) {
	o.IsNegated = &v
}

func (o FacetFilterValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FacetFilterValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.RelationType) {
		toSerialize["relationType"] = o.RelationType
	}
	if !IsNil(o.IsNegated) {
		toSerialize["isNegated"] = o.IsNegated
	}
	return toSerialize, nil
}

type NullableFacetFilterValue struct {
	value *FacetFilterValue
	isSet bool
}

func (v NullableFacetFilterValue) Get() *FacetFilterValue {
	return v.value
}

func (v *NullableFacetFilterValue) Set(val *FacetFilterValue) {
	v.value = val
	v.isSet = true
}

func (v NullableFacetFilterValue) IsSet() bool {
	return v.isSet
}

func (v *NullableFacetFilterValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFacetFilterValue(val *FacetFilterValue) *NullableFacetFilterValue {
	return &NullableFacetFilterValue{value: val, isSet: true}
}

func (v NullableFacetFilterValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFacetFilterValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


