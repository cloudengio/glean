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

// checks if the MetadataListConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetadataListConfig{}

// MetadataListConfig struct for MetadataListConfig
type MetadataListConfig struct {
	DatumConfigs *map[string]DatumConfig `json:"datumConfigs,omitempty"`
	Keys []string `json:"keys,omitempty"`
	Separator *string `json:"separator,omitempty"`
}

// NewMetadataListConfig instantiates a new MetadataListConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetadataListConfig() *MetadataListConfig {
	this := MetadataListConfig{}
	return &this
}

// NewMetadataListConfigWithDefaults instantiates a new MetadataListConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetadataListConfigWithDefaults() *MetadataListConfig {
	this := MetadataListConfig{}
	return &this
}

// GetDatumConfigs returns the DatumConfigs field value if set, zero value otherwise.
func (o *MetadataListConfig) GetDatumConfigs() map[string]DatumConfig {
	if o == nil || IsNil(o.DatumConfigs) {
		var ret map[string]DatumConfig
		return ret
	}
	return *o.DatumConfigs
}

// GetDatumConfigsOk returns a tuple with the DatumConfigs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataListConfig) GetDatumConfigsOk() (*map[string]DatumConfig, bool) {
	if o == nil || IsNil(o.DatumConfigs) {
		return nil, false
	}
	return o.DatumConfigs, true
}

// HasDatumConfigs returns a boolean if a field has been set.
func (o *MetadataListConfig) HasDatumConfigs() bool {
	if o != nil && !IsNil(o.DatumConfigs) {
		return true
	}

	return false
}

// SetDatumConfigs gets a reference to the given map[string]DatumConfig and assigns it to the DatumConfigs field.
func (o *MetadataListConfig) SetDatumConfigs(v map[string]DatumConfig) {
	o.DatumConfigs = &v
}

// GetKeys returns the Keys field value if set, zero value otherwise.
func (o *MetadataListConfig) GetKeys() []string {
	if o == nil || IsNil(o.Keys) {
		var ret []string
		return ret
	}
	return o.Keys
}

// GetKeysOk returns a tuple with the Keys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataListConfig) GetKeysOk() ([]string, bool) {
	if o == nil || IsNil(o.Keys) {
		return nil, false
	}
	return o.Keys, true
}

// HasKeys returns a boolean if a field has been set.
func (o *MetadataListConfig) HasKeys() bool {
	if o != nil && !IsNil(o.Keys) {
		return true
	}

	return false
}

// SetKeys gets a reference to the given []string and assigns it to the Keys field.
func (o *MetadataListConfig) SetKeys(v []string) {
	o.Keys = v
}

// GetSeparator returns the Separator field value if set, zero value otherwise.
func (o *MetadataListConfig) GetSeparator() string {
	if o == nil || IsNil(o.Separator) {
		var ret string
		return ret
	}
	return *o.Separator
}

// GetSeparatorOk returns a tuple with the Separator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataListConfig) GetSeparatorOk() (*string, bool) {
	if o == nil || IsNil(o.Separator) {
		return nil, false
	}
	return o.Separator, true
}

// HasSeparator returns a boolean if a field has been set.
func (o *MetadataListConfig) HasSeparator() bool {
	if o != nil && !IsNil(o.Separator) {
		return true
	}

	return false
}

// SetSeparator gets a reference to the given string and assigns it to the Separator field.
func (o *MetadataListConfig) SetSeparator(v string) {
	o.Separator = &v
}

func (o MetadataListConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetadataListConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DatumConfigs) {
		toSerialize["datumConfigs"] = o.DatumConfigs
	}
	if !IsNil(o.Keys) {
		toSerialize["keys"] = o.Keys
	}
	if !IsNil(o.Separator) {
		toSerialize["separator"] = o.Separator
	}
	return toSerialize, nil
}

type NullableMetadataListConfig struct {
	value *MetadataListConfig
	isSet bool
}

func (v NullableMetadataListConfig) Get() *MetadataListConfig {
	return v.value
}

func (v *NullableMetadataListConfig) Set(val *MetadataListConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadataListConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadataListConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadataListConfig(val *MetadataListConfig) *NullableMetadataListConfig {
	return &NullableMetadataListConfig{value: val, isSet: true}
}

func (v NullableMetadataListConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadataListConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


