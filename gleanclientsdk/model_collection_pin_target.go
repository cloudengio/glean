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

// checks if the CollectionPinTarget type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionPinTarget{}

// CollectionPinTarget struct for CollectionPinTarget
type CollectionPinTarget struct {
	Category CollectionPinnableCategories `json:"category"`
	// Optional. If category supports values, then the additional value for the category e.g. department name for DEPARTMENT_RESOURCE, team name/id for TEAM_RESOURCE and so on.
	Value *string `json:"value,omitempty"`
	Target *CollectionPinnableTargets `json:"target,omitempty"`
}

// NewCollectionPinTarget instantiates a new CollectionPinTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionPinTarget(category CollectionPinnableCategories) *CollectionPinTarget {
	this := CollectionPinTarget{}
	this.Category = category
	return &this
}

// NewCollectionPinTargetWithDefaults instantiates a new CollectionPinTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionPinTargetWithDefaults() *CollectionPinTarget {
	this := CollectionPinTarget{}
	return &this
}

// GetCategory returns the Category field value
func (o *CollectionPinTarget) GetCategory() CollectionPinnableCategories {
	if o == nil {
		var ret CollectionPinnableCategories
		return ret
	}

	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *CollectionPinTarget) GetCategoryOk() (*CollectionPinnableCategories, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value
func (o *CollectionPinTarget) SetCategory(v CollectionPinnableCategories) {
	o.Category = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionPinTarget) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionPinTarget) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionPinTarget) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *CollectionPinTarget) SetValue(v string) {
	o.Value = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *CollectionPinTarget) GetTarget() CollectionPinnableTargets {
	if o == nil || IsNil(o.Target) {
		var ret CollectionPinnableTargets
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionPinTarget) GetTargetOk() (*CollectionPinnableTargets, bool) {
	if o == nil || IsNil(o.Target) {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *CollectionPinTarget) HasTarget() bool {
	if o != nil && !IsNil(o.Target) {
		return true
	}

	return false
}

// SetTarget gets a reference to the given CollectionPinnableTargets and assigns it to the Target field.
func (o *CollectionPinTarget) SetTarget(v CollectionPinnableTargets) {
	o.Target = &v
}

func (o CollectionPinTarget) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionPinTarget) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["category"] = o.Category
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Target) {
		toSerialize["target"] = o.Target
	}
	return toSerialize, nil
}

type NullableCollectionPinTarget struct {
	value *CollectionPinTarget
	isSet bool
}

func (v NullableCollectionPinTarget) Get() *CollectionPinTarget {
	return v.value
}

func (v *NullableCollectionPinTarget) Set(val *CollectionPinTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionPinTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionPinTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionPinTarget(val *CollectionPinTarget) *NullableCollectionPinTarget {
	return &NullableCollectionPinTarget{value: val, isSet: true}
}

func (v NullableCollectionPinTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionPinTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

