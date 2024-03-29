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

// checks if the LabeledCountInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LabeledCountInfo{}

// LabeledCountInfo struct for LabeledCountInfo
type LabeledCountInfo struct {
	// Label for the included count information.
	Label string `json:"label"`
	// List of data points for counts for a given date period.
	CountInfo []CountInfo `json:"countInfo,omitempty"`
}

// NewLabeledCountInfo instantiates a new LabeledCountInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLabeledCountInfo(label string) *LabeledCountInfo {
	this := LabeledCountInfo{}
	this.Label = label
	return &this
}

// NewLabeledCountInfoWithDefaults instantiates a new LabeledCountInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLabeledCountInfoWithDefaults() *LabeledCountInfo {
	this := LabeledCountInfo{}
	return &this
}

// GetLabel returns the Label field value
func (o *LabeledCountInfo) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *LabeledCountInfo) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *LabeledCountInfo) SetLabel(v string) {
	o.Label = v
}

// GetCountInfo returns the CountInfo field value if set, zero value otherwise.
func (o *LabeledCountInfo) GetCountInfo() []CountInfo {
	if o == nil || IsNil(o.CountInfo) {
		var ret []CountInfo
		return ret
	}
	return o.CountInfo
}

// GetCountInfoOk returns a tuple with the CountInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabeledCountInfo) GetCountInfoOk() ([]CountInfo, bool) {
	if o == nil || IsNil(o.CountInfo) {
		return nil, false
	}
	return o.CountInfo, true
}

// HasCountInfo returns a boolean if a field has been set.
func (o *LabeledCountInfo) HasCountInfo() bool {
	if o != nil && !IsNil(o.CountInfo) {
		return true
	}

	return false
}

// SetCountInfo gets a reference to the given []CountInfo and assigns it to the CountInfo field.
func (o *LabeledCountInfo) SetCountInfo(v []CountInfo) {
	o.CountInfo = v
}

func (o LabeledCountInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LabeledCountInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["label"] = o.Label
	if !IsNil(o.CountInfo) {
		toSerialize["countInfo"] = o.CountInfo
	}
	return toSerialize, nil
}

type NullableLabeledCountInfo struct {
	value *LabeledCountInfo
	isSet bool
}

func (v NullableLabeledCountInfo) Get() *LabeledCountInfo {
	return v.value
}

func (v *NullableLabeledCountInfo) Set(val *LabeledCountInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableLabeledCountInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableLabeledCountInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLabeledCountInfo(val *LabeledCountInfo) *NullableLabeledCountInfo {
	return &NullableLabeledCountInfo{value: val, isSet: true}
}

func (v NullableLabeledCountInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLabeledCountInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


