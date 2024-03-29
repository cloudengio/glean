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

// checks if the ActionsConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActionsConfig{}

// ActionsConfig struct for ActionsConfig
type ActionsConfig struct {
	Answer *AnswerActionsConfig `json:"answer,omitempty"`
	DisableActions *bool `json:"disableActions,omitempty"`
	DisableCopyLink *bool `json:"disableCopyLink,omitempty"`
}

// NewActionsConfig instantiates a new ActionsConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActionsConfig() *ActionsConfig {
	this := ActionsConfig{}
	return &this
}

// NewActionsConfigWithDefaults instantiates a new ActionsConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActionsConfigWithDefaults() *ActionsConfig {
	this := ActionsConfig{}
	return &this
}

// GetAnswer returns the Answer field value if set, zero value otherwise.
func (o *ActionsConfig) GetAnswer() AnswerActionsConfig {
	if o == nil || IsNil(o.Answer) {
		var ret AnswerActionsConfig
		return ret
	}
	return *o.Answer
}

// GetAnswerOk returns a tuple with the Answer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionsConfig) GetAnswerOk() (*AnswerActionsConfig, bool) {
	if o == nil || IsNil(o.Answer) {
		return nil, false
	}
	return o.Answer, true
}

// HasAnswer returns a boolean if a field has been set.
func (o *ActionsConfig) HasAnswer() bool {
	if o != nil && !IsNil(o.Answer) {
		return true
	}

	return false
}

// SetAnswer gets a reference to the given AnswerActionsConfig and assigns it to the Answer field.
func (o *ActionsConfig) SetAnswer(v AnswerActionsConfig) {
	o.Answer = &v
}

// GetDisableActions returns the DisableActions field value if set, zero value otherwise.
func (o *ActionsConfig) GetDisableActions() bool {
	if o == nil || IsNil(o.DisableActions) {
		var ret bool
		return ret
	}
	return *o.DisableActions
}

// GetDisableActionsOk returns a tuple with the DisableActions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionsConfig) GetDisableActionsOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableActions) {
		return nil, false
	}
	return o.DisableActions, true
}

// HasDisableActions returns a boolean if a field has been set.
func (o *ActionsConfig) HasDisableActions() bool {
	if o != nil && !IsNil(o.DisableActions) {
		return true
	}

	return false
}

// SetDisableActions gets a reference to the given bool and assigns it to the DisableActions field.
func (o *ActionsConfig) SetDisableActions(v bool) {
	o.DisableActions = &v
}

// GetDisableCopyLink returns the DisableCopyLink field value if set, zero value otherwise.
func (o *ActionsConfig) GetDisableCopyLink() bool {
	if o == nil || IsNil(o.DisableCopyLink) {
		var ret bool
		return ret
	}
	return *o.DisableCopyLink
}

// GetDisableCopyLinkOk returns a tuple with the DisableCopyLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionsConfig) GetDisableCopyLinkOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableCopyLink) {
		return nil, false
	}
	return o.DisableCopyLink, true
}

// HasDisableCopyLink returns a boolean if a field has been set.
func (o *ActionsConfig) HasDisableCopyLink() bool {
	if o != nil && !IsNil(o.DisableCopyLink) {
		return true
	}

	return false
}

// SetDisableCopyLink gets a reference to the given bool and assigns it to the DisableCopyLink field.
func (o *ActionsConfig) SetDisableCopyLink(v bool) {
	o.DisableCopyLink = &v
}

func (o ActionsConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActionsConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Answer) {
		toSerialize["answer"] = o.Answer
	}
	if !IsNil(o.DisableActions) {
		toSerialize["disableActions"] = o.DisableActions
	}
	if !IsNil(o.DisableCopyLink) {
		toSerialize["disableCopyLink"] = o.DisableCopyLink
	}
	return toSerialize, nil
}

type NullableActionsConfig struct {
	value *ActionsConfig
	isSet bool
}

func (v NullableActionsConfig) Get() *ActionsConfig {
	return v.value
}

func (v *NullableActionsConfig) Set(val *ActionsConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableActionsConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableActionsConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActionsConfig(val *ActionsConfig) *NullableActionsConfig {
	return &NullableActionsConfig{value: val, isSet: true}
}

func (v NullableActionsConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActionsConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


