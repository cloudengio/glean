/*
Glean Client API - Platform Preview

# Introduction These are all the APIs used by Glean to implement the Glean client. These are available as platform preview for implementing a custom client to the Glean system.  # Usage guidelines A subset of these endpoints are also in the developer ready section, which is available for public use. The rest of the endpoints are subject to prior agreement with Glean before usage. Please contact support@glean.com if you would like to use an API that is not currently available in the developer ready section. 

API version: 0.9.0
Contact: support@glean.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gleanclientsdk

import (
	"encoding/json"
)

// checks if the DatumTransform type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatumTransform{}

// DatumTransform struct for DatumTransform
type DatumTransform struct {
	Match *string `json:"match,omitempty"`
	Replace *string `json:"replace,omitempty"`
	Type *string `json:"type,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewDatumTransform instantiates a new DatumTransform object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatumTransform() *DatumTransform {
	this := DatumTransform{}
	return &this
}

// NewDatumTransformWithDefaults instantiates a new DatumTransform object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatumTransformWithDefaults() *DatumTransform {
	this := DatumTransform{}
	return &this
}

// GetMatch returns the Match field value if set, zero value otherwise.
func (o *DatumTransform) GetMatch() string {
	if o == nil || IsNil(o.Match) {
		var ret string
		return ret
	}
	return *o.Match
}

// GetMatchOk returns a tuple with the Match field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatumTransform) GetMatchOk() (*string, bool) {
	if o == nil || IsNil(o.Match) {
		return nil, false
	}
	return o.Match, true
}

// HasMatch returns a boolean if a field has been set.
func (o *DatumTransform) HasMatch() bool {
	if o != nil && !IsNil(o.Match) {
		return true
	}

	return false
}

// SetMatch gets a reference to the given string and assigns it to the Match field.
func (o *DatumTransform) SetMatch(v string) {
	o.Match = &v
}

// GetReplace returns the Replace field value if set, zero value otherwise.
func (o *DatumTransform) GetReplace() string {
	if o == nil || IsNil(o.Replace) {
		var ret string
		return ret
	}
	return *o.Replace
}

// GetReplaceOk returns a tuple with the Replace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatumTransform) GetReplaceOk() (*string, bool) {
	if o == nil || IsNil(o.Replace) {
		return nil, false
	}
	return o.Replace, true
}

// HasReplace returns a boolean if a field has been set.
func (o *DatumTransform) HasReplace() bool {
	if o != nil && !IsNil(o.Replace) {
		return true
	}

	return false
}

// SetReplace gets a reference to the given string and assigns it to the Replace field.
func (o *DatumTransform) SetReplace(v string) {
	o.Replace = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DatumTransform) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatumTransform) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DatumTransform) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DatumTransform) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *DatumTransform) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatumTransform) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *DatumTransform) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *DatumTransform) SetValue(v string) {
	o.Value = &v
}

func (o DatumTransform) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatumTransform) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Match) {
		toSerialize["match"] = o.Match
	}
	if !IsNil(o.Replace) {
		toSerialize["replace"] = o.Replace
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableDatumTransform struct {
	value *DatumTransform
	isSet bool
}

func (v NullableDatumTransform) Get() *DatumTransform {
	return v.value
}

func (v *NullableDatumTransform) Set(val *DatumTransform) {
	v.value = val
	v.isSet = true
}

func (v NullableDatumTransform) IsSet() bool {
	return v.isSet
}

func (v *NullableDatumTransform) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatumTransform(val *DatumTransform) *NullableDatumTransform {
	return &NullableDatumTransform{value: val, isSet: true}
}

func (v NullableDatumTransform) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatumTransform) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


