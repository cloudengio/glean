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

// checks if the CreateShortcutResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateShortcutResponse{}

// CreateShortcutResponse struct for CreateShortcutResponse
type CreateShortcutResponse struct {
	Shortcut *Shortcut `json:"shortcut,omitempty"`
	Error *ShortcutError `json:"error,omitempty"`
}

// NewCreateShortcutResponse instantiates a new CreateShortcutResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateShortcutResponse() *CreateShortcutResponse {
	this := CreateShortcutResponse{}
	return &this
}

// NewCreateShortcutResponseWithDefaults instantiates a new CreateShortcutResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateShortcutResponseWithDefaults() *CreateShortcutResponse {
	this := CreateShortcutResponse{}
	return &this
}

// GetShortcut returns the Shortcut field value if set, zero value otherwise.
func (o *CreateShortcutResponse) GetShortcut() Shortcut {
	if o == nil || IsNil(o.Shortcut) {
		var ret Shortcut
		return ret
	}
	return *o.Shortcut
}

// GetShortcutOk returns a tuple with the Shortcut field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShortcutResponse) GetShortcutOk() (*Shortcut, bool) {
	if o == nil || IsNil(o.Shortcut) {
		return nil, false
	}
	return o.Shortcut, true
}

// HasShortcut returns a boolean if a field has been set.
func (o *CreateShortcutResponse) HasShortcut() bool {
	if o != nil && !IsNil(o.Shortcut) {
		return true
	}

	return false
}

// SetShortcut gets a reference to the given Shortcut and assigns it to the Shortcut field.
func (o *CreateShortcutResponse) SetShortcut(v Shortcut) {
	o.Shortcut = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *CreateShortcutResponse) GetError() ShortcutError {
	if o == nil || IsNil(o.Error) {
		var ret ShortcutError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShortcutResponse) GetErrorOk() (*ShortcutError, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *CreateShortcutResponse) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given ShortcutError and assigns it to the Error field.
func (o *CreateShortcutResponse) SetError(v ShortcutError) {
	o.Error = &v
}

func (o CreateShortcutResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateShortcutResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Shortcut) {
		toSerialize["shortcut"] = o.Shortcut
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableCreateShortcutResponse struct {
	value *CreateShortcutResponse
	isSet bool
}

func (v NullableCreateShortcutResponse) Get() *CreateShortcutResponse {
	return v.value
}

func (v *NullableCreateShortcutResponse) Set(val *CreateShortcutResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateShortcutResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateShortcutResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateShortcutResponse(val *CreateShortcutResponse) *NullableCreateShortcutResponse {
	return &NullableCreateShortcutResponse{value: val, isSet: true}
}

func (v NullableCreateShortcutResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateShortcutResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


