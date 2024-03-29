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

// checks if the ContactConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContactConfig{}

// ContactConfig struct for ContactConfig
type ContactConfig struct {
	Show *bool `json:"show,omitempty"`
	EmailKey *string `json:"emailKey,omitempty"`
	PhoneKey *string `json:"phoneKey,omitempty"`
}

// NewContactConfig instantiates a new ContactConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContactConfig() *ContactConfig {
	this := ContactConfig{}
	return &this
}

// NewContactConfigWithDefaults instantiates a new ContactConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContactConfigWithDefaults() *ContactConfig {
	this := ContactConfig{}
	return &this
}

// GetShow returns the Show field value if set, zero value otherwise.
func (o *ContactConfig) GetShow() bool {
	if o == nil || IsNil(o.Show) {
		var ret bool
		return ret
	}
	return *o.Show
}

// GetShowOk returns a tuple with the Show field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactConfig) GetShowOk() (*bool, bool) {
	if o == nil || IsNil(o.Show) {
		return nil, false
	}
	return o.Show, true
}

// HasShow returns a boolean if a field has been set.
func (o *ContactConfig) HasShow() bool {
	if o != nil && !IsNil(o.Show) {
		return true
	}

	return false
}

// SetShow gets a reference to the given bool and assigns it to the Show field.
func (o *ContactConfig) SetShow(v bool) {
	o.Show = &v
}

// GetEmailKey returns the EmailKey field value if set, zero value otherwise.
func (o *ContactConfig) GetEmailKey() string {
	if o == nil || IsNil(o.EmailKey) {
		var ret string
		return ret
	}
	return *o.EmailKey
}

// GetEmailKeyOk returns a tuple with the EmailKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactConfig) GetEmailKeyOk() (*string, bool) {
	if o == nil || IsNil(o.EmailKey) {
		return nil, false
	}
	return o.EmailKey, true
}

// HasEmailKey returns a boolean if a field has been set.
func (o *ContactConfig) HasEmailKey() bool {
	if o != nil && !IsNil(o.EmailKey) {
		return true
	}

	return false
}

// SetEmailKey gets a reference to the given string and assigns it to the EmailKey field.
func (o *ContactConfig) SetEmailKey(v string) {
	o.EmailKey = &v
}

// GetPhoneKey returns the PhoneKey field value if set, zero value otherwise.
func (o *ContactConfig) GetPhoneKey() string {
	if o == nil || IsNil(o.PhoneKey) {
		var ret string
		return ret
	}
	return *o.PhoneKey
}

// GetPhoneKeyOk returns a tuple with the PhoneKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactConfig) GetPhoneKeyOk() (*string, bool) {
	if o == nil || IsNil(o.PhoneKey) {
		return nil, false
	}
	return o.PhoneKey, true
}

// HasPhoneKey returns a boolean if a field has been set.
func (o *ContactConfig) HasPhoneKey() bool {
	if o != nil && !IsNil(o.PhoneKey) {
		return true
	}

	return false
}

// SetPhoneKey gets a reference to the given string and assigns it to the PhoneKey field.
func (o *ContactConfig) SetPhoneKey(v string) {
	o.PhoneKey = &v
}

func (o ContactConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContactConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Show) {
		toSerialize["show"] = o.Show
	}
	if !IsNil(o.EmailKey) {
		toSerialize["emailKey"] = o.EmailKey
	}
	if !IsNil(o.PhoneKey) {
		toSerialize["phoneKey"] = o.PhoneKey
	}
	return toSerialize, nil
}

type NullableContactConfig struct {
	value *ContactConfig
	isSet bool
}

func (v NullableContactConfig) Get() *ContactConfig {
	return v.value
}

func (v *NullableContactConfig) Set(val *ContactConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableContactConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableContactConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContactConfig(val *ContactConfig) *NullableContactConfig {
	return &NullableContactConfig{value: val, isSet: true}
}

func (v NullableContactConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContactConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


