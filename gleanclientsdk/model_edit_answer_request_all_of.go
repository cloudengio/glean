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

// checks if the EditAnswerRequestAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EditAnswerRequestAllOf{}

// EditAnswerRequestAllOf struct for EditAnswerRequestAllOf
type EditAnswerRequestAllOf struct {
	CombinedAnswerText *StructuredTextMutableProperties `json:"combinedAnswerText,omitempty"`
}

// NewEditAnswerRequestAllOf instantiates a new EditAnswerRequestAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditAnswerRequestAllOf() *EditAnswerRequestAllOf {
	this := EditAnswerRequestAllOf{}
	return &this
}

// NewEditAnswerRequestAllOfWithDefaults instantiates a new EditAnswerRequestAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditAnswerRequestAllOfWithDefaults() *EditAnswerRequestAllOf {
	this := EditAnswerRequestAllOf{}
	return &this
}

// GetCombinedAnswerText returns the CombinedAnswerText field value if set, zero value otherwise.
func (o *EditAnswerRequestAllOf) GetCombinedAnswerText() StructuredTextMutableProperties {
	if o == nil || IsNil(o.CombinedAnswerText) {
		var ret StructuredTextMutableProperties
		return ret
	}
	return *o.CombinedAnswerText
}

// GetCombinedAnswerTextOk returns a tuple with the CombinedAnswerText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditAnswerRequestAllOf) GetCombinedAnswerTextOk() (*StructuredTextMutableProperties, bool) {
	if o == nil || IsNil(o.CombinedAnswerText) {
		return nil, false
	}
	return o.CombinedAnswerText, true
}

// HasCombinedAnswerText returns a boolean if a field has been set.
func (o *EditAnswerRequestAllOf) HasCombinedAnswerText() bool {
	if o != nil && !IsNil(o.CombinedAnswerText) {
		return true
	}

	return false
}

// SetCombinedAnswerText gets a reference to the given StructuredTextMutableProperties and assigns it to the CombinedAnswerText field.
func (o *EditAnswerRequestAllOf) SetCombinedAnswerText(v StructuredTextMutableProperties) {
	o.CombinedAnswerText = &v
}

func (o EditAnswerRequestAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EditAnswerRequestAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CombinedAnswerText) {
		toSerialize["combinedAnswerText"] = o.CombinedAnswerText
	}
	return toSerialize, nil
}

type NullableEditAnswerRequestAllOf struct {
	value *EditAnswerRequestAllOf
	isSet bool
}

func (v NullableEditAnswerRequestAllOf) Get() *EditAnswerRequestAllOf {
	return v.value
}

func (v *NullableEditAnswerRequestAllOf) Set(val *EditAnswerRequestAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEditAnswerRequestAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEditAnswerRequestAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditAnswerRequestAllOf(val *EditAnswerRequestAllOf) *NullableEditAnswerRequestAllOf {
	return &NullableEditAnswerRequestAllOf{value: val, isSet: true}
}

func (v NullableEditAnswerRequestAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditAnswerRequestAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


