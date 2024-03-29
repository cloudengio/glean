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

// checks if the FacetValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FacetValue{}

// FacetValue struct for FacetValue
type FacetValue struct {
	// The value that should be set in the FacetFilter when applying this filter to a search request.
	StringValue *string `json:"stringValue,omitempty"`
	IntegerValue *int32 `json:"integerValue,omitempty"`
	// An optional user-friendly label to display in place of the facet value.
	DisplayLabel *string `json:"displayLabel,omitempty"`
	IconConfig *IconConfig `json:"iconConfig,omitempty"`
}

// NewFacetValue instantiates a new FacetValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFacetValue() *FacetValue {
	this := FacetValue{}
	return &this
}

// NewFacetValueWithDefaults instantiates a new FacetValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFacetValueWithDefaults() *FacetValue {
	this := FacetValue{}
	return &this
}

// GetStringValue returns the StringValue field value if set, zero value otherwise.
func (o *FacetValue) GetStringValue() string {
	if o == nil || IsNil(o.StringValue) {
		var ret string
		return ret
	}
	return *o.StringValue
}

// GetStringValueOk returns a tuple with the StringValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FacetValue) GetStringValueOk() (*string, bool) {
	if o == nil || IsNil(o.StringValue) {
		return nil, false
	}
	return o.StringValue, true
}

// HasStringValue returns a boolean if a field has been set.
func (o *FacetValue) HasStringValue() bool {
	if o != nil && !IsNil(o.StringValue) {
		return true
	}

	return false
}

// SetStringValue gets a reference to the given string and assigns it to the StringValue field.
func (o *FacetValue) SetStringValue(v string) {
	o.StringValue = &v
}

// GetIntegerValue returns the IntegerValue field value if set, zero value otherwise.
func (o *FacetValue) GetIntegerValue() int32 {
	if o == nil || IsNil(o.IntegerValue) {
		var ret int32
		return ret
	}
	return *o.IntegerValue
}

// GetIntegerValueOk returns a tuple with the IntegerValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FacetValue) GetIntegerValueOk() (*int32, bool) {
	if o == nil || IsNil(o.IntegerValue) {
		return nil, false
	}
	return o.IntegerValue, true
}

// HasIntegerValue returns a boolean if a field has been set.
func (o *FacetValue) HasIntegerValue() bool {
	if o != nil && !IsNil(o.IntegerValue) {
		return true
	}

	return false
}

// SetIntegerValue gets a reference to the given int32 and assigns it to the IntegerValue field.
func (o *FacetValue) SetIntegerValue(v int32) {
	o.IntegerValue = &v
}

// GetDisplayLabel returns the DisplayLabel field value if set, zero value otherwise.
func (o *FacetValue) GetDisplayLabel() string {
	if o == nil || IsNil(o.DisplayLabel) {
		var ret string
		return ret
	}
	return *o.DisplayLabel
}

// GetDisplayLabelOk returns a tuple with the DisplayLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FacetValue) GetDisplayLabelOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayLabel) {
		return nil, false
	}
	return o.DisplayLabel, true
}

// HasDisplayLabel returns a boolean if a field has been set.
func (o *FacetValue) HasDisplayLabel() bool {
	if o != nil && !IsNil(o.DisplayLabel) {
		return true
	}

	return false
}

// SetDisplayLabel gets a reference to the given string and assigns it to the DisplayLabel field.
func (o *FacetValue) SetDisplayLabel(v string) {
	o.DisplayLabel = &v
}

// GetIconConfig returns the IconConfig field value if set, zero value otherwise.
func (o *FacetValue) GetIconConfig() IconConfig {
	if o == nil || IsNil(o.IconConfig) {
		var ret IconConfig
		return ret
	}
	return *o.IconConfig
}

// GetIconConfigOk returns a tuple with the IconConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FacetValue) GetIconConfigOk() (*IconConfig, bool) {
	if o == nil || IsNil(o.IconConfig) {
		return nil, false
	}
	return o.IconConfig, true
}

// HasIconConfig returns a boolean if a field has been set.
func (o *FacetValue) HasIconConfig() bool {
	if o != nil && !IsNil(o.IconConfig) {
		return true
	}

	return false
}

// SetIconConfig gets a reference to the given IconConfig and assigns it to the IconConfig field.
func (o *FacetValue) SetIconConfig(v IconConfig) {
	o.IconConfig = &v
}

func (o FacetValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FacetValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StringValue) {
		toSerialize["stringValue"] = o.StringValue
	}
	if !IsNil(o.IntegerValue) {
		toSerialize["integerValue"] = o.IntegerValue
	}
	if !IsNil(o.DisplayLabel) {
		toSerialize["displayLabel"] = o.DisplayLabel
	}
	if !IsNil(o.IconConfig) {
		toSerialize["iconConfig"] = o.IconConfig
	}
	return toSerialize, nil
}

type NullableFacetValue struct {
	value *FacetValue
	isSet bool
}

func (v NullableFacetValue) Get() *FacetValue {
	return v.value
}

func (v *NullableFacetValue) Set(val *FacetValue) {
	v.value = val
	v.isSet = true
}

func (v NullableFacetValue) IsSet() bool {
	return v.isSet
}

func (v *NullableFacetValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFacetValue(val *FacetValue) *NullableFacetValue {
	return &NullableFacetValue{value: val, isSet: true}
}

func (v NullableFacetValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFacetValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


