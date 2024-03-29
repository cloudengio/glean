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

// checks if the GetSimilarShortcutsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSimilarShortcutsResponse{}

// GetSimilarShortcutsResponse struct for GetSimilarShortcutsResponse
type GetSimilarShortcutsResponse struct {
	// Shortcuts with similar aliases to the given. Includes shortcut with the exact alias if it exists.
	Shortcuts []Shortcut `json:"shortcuts"`
}

// NewGetSimilarShortcutsResponse instantiates a new GetSimilarShortcutsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSimilarShortcutsResponse(shortcuts []Shortcut) *GetSimilarShortcutsResponse {
	this := GetSimilarShortcutsResponse{}
	this.Shortcuts = shortcuts
	return &this
}

// NewGetSimilarShortcutsResponseWithDefaults instantiates a new GetSimilarShortcutsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSimilarShortcutsResponseWithDefaults() *GetSimilarShortcutsResponse {
	this := GetSimilarShortcutsResponse{}
	return &this
}

// GetShortcuts returns the Shortcuts field value
func (o *GetSimilarShortcutsResponse) GetShortcuts() []Shortcut {
	if o == nil {
		var ret []Shortcut
		return ret
	}

	return o.Shortcuts
}

// GetShortcutsOk returns a tuple with the Shortcuts field value
// and a boolean to check if the value has been set.
func (o *GetSimilarShortcutsResponse) GetShortcutsOk() ([]Shortcut, bool) {
	if o == nil {
		return nil, false
	}
	return o.Shortcuts, true
}

// SetShortcuts sets field value
func (o *GetSimilarShortcutsResponse) SetShortcuts(v []Shortcut) {
	o.Shortcuts = v
}

func (o GetSimilarShortcutsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSimilarShortcutsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["shortcuts"] = o.Shortcuts
	return toSerialize, nil
}

type NullableGetSimilarShortcutsResponse struct {
	value *GetSimilarShortcutsResponse
	isSet bool
}

func (v NullableGetSimilarShortcutsResponse) Get() *GetSimilarShortcutsResponse {
	return v.value
}

func (v *NullableGetSimilarShortcutsResponse) Set(val *GetSimilarShortcutsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSimilarShortcutsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSimilarShortcutsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSimilarShortcutsResponse(val *GetSimilarShortcutsResponse) *NullableGetSimilarShortcutsResponse {
	return &NullableGetSimilarShortcutsResponse{value: val, isSet: true}
}

func (v NullableGetSimilarShortcutsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSimilarShortcutsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


