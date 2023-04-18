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

// checks if the RelatedSectionConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RelatedSectionConfig{}

// RelatedSectionConfig struct for RelatedSectionConfig
type RelatedSectionConfig struct {
	DefaultExpanded *bool `json:"defaultExpanded,omitempty"`
	Title *string `json:"title,omitempty"`
	Type *string `json:"type,omitempty"`
	TypeConfig *RelatedSectionConfigTypeConfig `json:"typeConfig,omitempty"`
	Icon *IconConfig `json:"icon,omitempty"`
}

// NewRelatedSectionConfig instantiates a new RelatedSectionConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelatedSectionConfig() *RelatedSectionConfig {
	this := RelatedSectionConfig{}
	return &this
}

// NewRelatedSectionConfigWithDefaults instantiates a new RelatedSectionConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelatedSectionConfigWithDefaults() *RelatedSectionConfig {
	this := RelatedSectionConfig{}
	return &this
}

// GetDefaultExpanded returns the DefaultExpanded field value if set, zero value otherwise.
func (o *RelatedSectionConfig) GetDefaultExpanded() bool {
	if o == nil || IsNil(o.DefaultExpanded) {
		var ret bool
		return ret
	}
	return *o.DefaultExpanded
}

// GetDefaultExpandedOk returns a tuple with the DefaultExpanded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelatedSectionConfig) GetDefaultExpandedOk() (*bool, bool) {
	if o == nil || IsNil(o.DefaultExpanded) {
		return nil, false
	}
	return o.DefaultExpanded, true
}

// HasDefaultExpanded returns a boolean if a field has been set.
func (o *RelatedSectionConfig) HasDefaultExpanded() bool {
	if o != nil && !IsNil(o.DefaultExpanded) {
		return true
	}

	return false
}

// SetDefaultExpanded gets a reference to the given bool and assigns it to the DefaultExpanded field.
func (o *RelatedSectionConfig) SetDefaultExpanded(v bool) {
	o.DefaultExpanded = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *RelatedSectionConfig) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelatedSectionConfig) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *RelatedSectionConfig) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *RelatedSectionConfig) SetTitle(v string) {
	o.Title = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RelatedSectionConfig) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelatedSectionConfig) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RelatedSectionConfig) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *RelatedSectionConfig) SetType(v string) {
	o.Type = &v
}

// GetTypeConfig returns the TypeConfig field value if set, zero value otherwise.
func (o *RelatedSectionConfig) GetTypeConfig() RelatedSectionConfigTypeConfig {
	if o == nil || IsNil(o.TypeConfig) {
		var ret RelatedSectionConfigTypeConfig
		return ret
	}
	return *o.TypeConfig
}

// GetTypeConfigOk returns a tuple with the TypeConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelatedSectionConfig) GetTypeConfigOk() (*RelatedSectionConfigTypeConfig, bool) {
	if o == nil || IsNil(o.TypeConfig) {
		return nil, false
	}
	return o.TypeConfig, true
}

// HasTypeConfig returns a boolean if a field has been set.
func (o *RelatedSectionConfig) HasTypeConfig() bool {
	if o != nil && !IsNil(o.TypeConfig) {
		return true
	}

	return false
}

// SetTypeConfig gets a reference to the given RelatedSectionConfigTypeConfig and assigns it to the TypeConfig field.
func (o *RelatedSectionConfig) SetTypeConfig(v RelatedSectionConfigTypeConfig) {
	o.TypeConfig = &v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *RelatedSectionConfig) GetIcon() IconConfig {
	if o == nil || IsNil(o.Icon) {
		var ret IconConfig
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelatedSectionConfig) GetIconOk() (*IconConfig, bool) {
	if o == nil || IsNil(o.Icon) {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *RelatedSectionConfig) HasIcon() bool {
	if o != nil && !IsNil(o.Icon) {
		return true
	}

	return false
}

// SetIcon gets a reference to the given IconConfig and assigns it to the Icon field.
func (o *RelatedSectionConfig) SetIcon(v IconConfig) {
	o.Icon = &v
}

func (o RelatedSectionConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RelatedSectionConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DefaultExpanded) {
		toSerialize["defaultExpanded"] = o.DefaultExpanded
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.TypeConfig) {
		toSerialize["typeConfig"] = o.TypeConfig
	}
	if !IsNil(o.Icon) {
		toSerialize["icon"] = o.Icon
	}
	return toSerialize, nil
}

type NullableRelatedSectionConfig struct {
	value *RelatedSectionConfig
	isSet bool
}

func (v NullableRelatedSectionConfig) Get() *RelatedSectionConfig {
	return v.value
}

func (v *NullableRelatedSectionConfig) Set(val *RelatedSectionConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableRelatedSectionConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableRelatedSectionConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRelatedSectionConfig(val *RelatedSectionConfig) *NullableRelatedSectionConfig {
	return &NullableRelatedSectionConfig{value: val, isSet: true}
}

func (v NullableRelatedSectionConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRelatedSectionConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

