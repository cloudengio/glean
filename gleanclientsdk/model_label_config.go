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

// checks if the LabelConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LabelConfig{}

// LabelConfig struct for LabelConfig
type LabelConfig struct {
	Background *string `json:"background,omitempty"`
	Color *string `json:"color,omitempty"`
	Icon *string `json:"icon,omitempty"`
	Label *string `json:"label,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewLabelConfig instantiates a new LabelConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLabelConfig() *LabelConfig {
	this := LabelConfig{}
	return &this
}

// NewLabelConfigWithDefaults instantiates a new LabelConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLabelConfigWithDefaults() *LabelConfig {
	this := LabelConfig{}
	return &this
}

// GetBackground returns the Background field value if set, zero value otherwise.
func (o *LabelConfig) GetBackground() string {
	if o == nil || IsNil(o.Background) {
		var ret string
		return ret
	}
	return *o.Background
}

// GetBackgroundOk returns a tuple with the Background field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabelConfig) GetBackgroundOk() (*string, bool) {
	if o == nil || IsNil(o.Background) {
		return nil, false
	}
	return o.Background, true
}

// HasBackground returns a boolean if a field has been set.
func (o *LabelConfig) HasBackground() bool {
	if o != nil && !IsNil(o.Background) {
		return true
	}

	return false
}

// SetBackground gets a reference to the given string and assigns it to the Background field.
func (o *LabelConfig) SetBackground(v string) {
	o.Background = &v
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *LabelConfig) GetColor() string {
	if o == nil || IsNil(o.Color) {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabelConfig) GetColorOk() (*string, bool) {
	if o == nil || IsNil(o.Color) {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *LabelConfig) HasColor() bool {
	if o != nil && !IsNil(o.Color) {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *LabelConfig) SetColor(v string) {
	o.Color = &v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *LabelConfig) GetIcon() string {
	if o == nil || IsNil(o.Icon) {
		var ret string
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabelConfig) GetIconOk() (*string, bool) {
	if o == nil || IsNil(o.Icon) {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *LabelConfig) HasIcon() bool {
	if o != nil && !IsNil(o.Icon) {
		return true
	}

	return false
}

// SetIcon gets a reference to the given string and assigns it to the Icon field.
func (o *LabelConfig) SetIcon(v string) {
	o.Icon = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *LabelConfig) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabelConfig) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *LabelConfig) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *LabelConfig) SetLabel(v string) {
	o.Label = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *LabelConfig) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabelConfig) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *LabelConfig) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *LabelConfig) SetValue(v string) {
	o.Value = &v
}

func (o LabelConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LabelConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Background) {
		toSerialize["background"] = o.Background
	}
	if !IsNil(o.Color) {
		toSerialize["color"] = o.Color
	}
	if !IsNil(o.Icon) {
		toSerialize["icon"] = o.Icon
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableLabelConfig struct {
	value *LabelConfig
	isSet bool
}

func (v NullableLabelConfig) Get() *LabelConfig {
	return v.value
}

func (v *NullableLabelConfig) Set(val *LabelConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableLabelConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableLabelConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLabelConfig(val *LabelConfig) *NullableLabelConfig {
	return &NullableLabelConfig{value: val, isSet: true}
}

func (v NullableLabelConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLabelConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

