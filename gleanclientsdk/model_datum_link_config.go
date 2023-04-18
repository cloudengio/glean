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

// checks if the DatumLinkConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatumLinkConfig{}

// DatumLinkConfig struct for DatumLinkConfig
type DatumLinkConfig struct {
	Key *string `json:"key,omitempty"`
	Role *string `json:"role,omitempty"`
	LinkSource *DatumLinkSource `json:"linkSource,omitempty"`
}

// NewDatumLinkConfig instantiates a new DatumLinkConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatumLinkConfig() *DatumLinkConfig {
	this := DatumLinkConfig{}
	return &this
}

// NewDatumLinkConfigWithDefaults instantiates a new DatumLinkConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatumLinkConfigWithDefaults() *DatumLinkConfig {
	this := DatumLinkConfig{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *DatumLinkConfig) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatumLinkConfig) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *DatumLinkConfig) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *DatumLinkConfig) SetKey(v string) {
	o.Key = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *DatumLinkConfig) GetRole() string {
	if o == nil || IsNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatumLinkConfig) GetRoleOk() (*string, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *DatumLinkConfig) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *DatumLinkConfig) SetRole(v string) {
	o.Role = &v
}

// GetLinkSource returns the LinkSource field value if set, zero value otherwise.
func (o *DatumLinkConfig) GetLinkSource() DatumLinkSource {
	if o == nil || IsNil(o.LinkSource) {
		var ret DatumLinkSource
		return ret
	}
	return *o.LinkSource
}

// GetLinkSourceOk returns a tuple with the LinkSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatumLinkConfig) GetLinkSourceOk() (*DatumLinkSource, bool) {
	if o == nil || IsNil(o.LinkSource) {
		return nil, false
	}
	return o.LinkSource, true
}

// HasLinkSource returns a boolean if a field has been set.
func (o *DatumLinkConfig) HasLinkSource() bool {
	if o != nil && !IsNil(o.LinkSource) {
		return true
	}

	return false
}

// SetLinkSource gets a reference to the given DatumLinkSource and assigns it to the LinkSource field.
func (o *DatumLinkConfig) SetLinkSource(v DatumLinkSource) {
	o.LinkSource = &v
}

func (o DatumLinkConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatumLinkConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	if !IsNil(o.LinkSource) {
		toSerialize["linkSource"] = o.LinkSource
	}
	return toSerialize, nil
}

type NullableDatumLinkConfig struct {
	value *DatumLinkConfig
	isSet bool
}

func (v NullableDatumLinkConfig) Get() *DatumLinkConfig {
	return v.value
}

func (v *NullableDatumLinkConfig) Set(val *DatumLinkConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableDatumLinkConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableDatumLinkConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatumLinkConfig(val *DatumLinkConfig) *NullableDatumLinkConfig {
	return &NullableDatumLinkConfig{value: val, isSet: true}
}

func (v NullableDatumLinkConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatumLinkConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


