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

// checks if the FilePathConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FilePathConfig{}

// FilePathConfig struct for FilePathConfig
type FilePathConfig struct {
	Show *bool `json:"show,omitempty"`
}

// NewFilePathConfig instantiates a new FilePathConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilePathConfig() *FilePathConfig {
	this := FilePathConfig{}
	return &this
}

// NewFilePathConfigWithDefaults instantiates a new FilePathConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilePathConfigWithDefaults() *FilePathConfig {
	this := FilePathConfig{}
	return &this
}

// GetShow returns the Show field value if set, zero value otherwise.
func (o *FilePathConfig) GetShow() bool {
	if o == nil || IsNil(o.Show) {
		var ret bool
		return ret
	}
	return *o.Show
}

// GetShowOk returns a tuple with the Show field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilePathConfig) GetShowOk() (*bool, bool) {
	if o == nil || IsNil(o.Show) {
		return nil, false
	}
	return o.Show, true
}

// HasShow returns a boolean if a field has been set.
func (o *FilePathConfig) HasShow() bool {
	if o != nil && !IsNil(o.Show) {
		return true
	}

	return false
}

// SetShow gets a reference to the given bool and assigns it to the Show field.
func (o *FilePathConfig) SetShow(v bool) {
	o.Show = &v
}

func (o FilePathConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FilePathConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Show) {
		toSerialize["show"] = o.Show
	}
	return toSerialize, nil
}

type NullableFilePathConfig struct {
	value *FilePathConfig
	isSet bool
}

func (v NullableFilePathConfig) Get() *FilePathConfig {
	return v.value
}

func (v *NullableFilePathConfig) Set(val *FilePathConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableFilePathConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableFilePathConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilePathConfig(val *FilePathConfig) *NullableFilePathConfig {
	return &NullableFilePathConfig{value: val, isSet: true}
}

func (v NullableFilePathConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilePathConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


