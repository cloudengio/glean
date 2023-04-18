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

// checks if the PreviewConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PreviewConfig{}

// PreviewConfig struct for PreviewConfig
type PreviewConfig struct {
	Icon *CompositeIconConfig `json:"icon,omitempty"`
	Meta *MetaConfig `json:"meta,omitempty"`
	ModalWidth *float32 `json:"modalWidth,omitempty"`
	Title *MetadataListConfig `json:"title,omitempty"`
	Type *DocumentPreviewType `json:"type,omitempty"`
}

// NewPreviewConfig instantiates a new PreviewConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPreviewConfig() *PreviewConfig {
	this := PreviewConfig{}
	return &this
}

// NewPreviewConfigWithDefaults instantiates a new PreviewConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPreviewConfigWithDefaults() *PreviewConfig {
	this := PreviewConfig{}
	return &this
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *PreviewConfig) GetIcon() CompositeIconConfig {
	if o == nil || IsNil(o.Icon) {
		var ret CompositeIconConfig
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreviewConfig) GetIconOk() (*CompositeIconConfig, bool) {
	if o == nil || IsNil(o.Icon) {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *PreviewConfig) HasIcon() bool {
	if o != nil && !IsNil(o.Icon) {
		return true
	}

	return false
}

// SetIcon gets a reference to the given CompositeIconConfig and assigns it to the Icon field.
func (o *PreviewConfig) SetIcon(v CompositeIconConfig) {
	o.Icon = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *PreviewConfig) GetMeta() MetaConfig {
	if o == nil || IsNil(o.Meta) {
		var ret MetaConfig
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreviewConfig) GetMetaOk() (*MetaConfig, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *PreviewConfig) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaConfig and assigns it to the Meta field.
func (o *PreviewConfig) SetMeta(v MetaConfig) {
	o.Meta = &v
}

// GetModalWidth returns the ModalWidth field value if set, zero value otherwise.
func (o *PreviewConfig) GetModalWidth() float32 {
	if o == nil || IsNil(o.ModalWidth) {
		var ret float32
		return ret
	}
	return *o.ModalWidth
}

// GetModalWidthOk returns a tuple with the ModalWidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreviewConfig) GetModalWidthOk() (*float32, bool) {
	if o == nil || IsNil(o.ModalWidth) {
		return nil, false
	}
	return o.ModalWidth, true
}

// HasModalWidth returns a boolean if a field has been set.
func (o *PreviewConfig) HasModalWidth() bool {
	if o != nil && !IsNil(o.ModalWidth) {
		return true
	}

	return false
}

// SetModalWidth gets a reference to the given float32 and assigns it to the ModalWidth field.
func (o *PreviewConfig) SetModalWidth(v float32) {
	o.ModalWidth = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *PreviewConfig) GetTitle() MetadataListConfig {
	if o == nil || IsNil(o.Title) {
		var ret MetadataListConfig
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreviewConfig) GetTitleOk() (*MetadataListConfig, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *PreviewConfig) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given MetadataListConfig and assigns it to the Title field.
func (o *PreviewConfig) SetTitle(v MetadataListConfig) {
	o.Title = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PreviewConfig) GetType() DocumentPreviewType {
	if o == nil || IsNil(o.Type) {
		var ret DocumentPreviewType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreviewConfig) GetTypeOk() (*DocumentPreviewType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PreviewConfig) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given DocumentPreviewType and assigns it to the Type field.
func (o *PreviewConfig) SetType(v DocumentPreviewType) {
	o.Type = &v
}

func (o PreviewConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PreviewConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Icon) {
		toSerialize["icon"] = o.Icon
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.ModalWidth) {
		toSerialize["modalWidth"] = o.ModalWidth
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullablePreviewConfig struct {
	value *PreviewConfig
	isSet bool
}

func (v NullablePreviewConfig) Get() *PreviewConfig {
	return v.value
}

func (v *NullablePreviewConfig) Set(val *PreviewConfig) {
	v.value = val
	v.isSet = true
}

func (v NullablePreviewConfig) IsSet() bool {
	return v.isSet
}

func (v *NullablePreviewConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePreviewConfig(val *PreviewConfig) *NullablePreviewConfig {
	return &NullablePreviewConfig{value: val, isSet: true}
}

func (v NullablePreviewConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePreviewConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

