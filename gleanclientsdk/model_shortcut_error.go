/*
Glean Client API

# Introduction These are the public APIs to enable implementing a custom client interface to the Glean system.  # Usage guidelines This API is evolving fast. Glean will provide advance notice of any planned backwards incompatible changes along with a 6-month sunset period for anything that requires developers to adopt the new versions. 

API version: 0.9.0
Contact: support@glean.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gleanclientsdk

import (
	"encoding/json"
)

// checks if the ShortcutError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShortcutError{}

// ShortcutError struct for ShortcutError
type ShortcutError struct {
	ErrorType *string `json:"errorType,omitempty"`
}

// NewShortcutError instantiates a new ShortcutError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShortcutError() *ShortcutError {
	this := ShortcutError{}
	return &this
}

// NewShortcutErrorWithDefaults instantiates a new ShortcutError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShortcutErrorWithDefaults() *ShortcutError {
	this := ShortcutError{}
	return &this
}

// GetErrorType returns the ErrorType field value if set, zero value otherwise.
func (o *ShortcutError) GetErrorType() string {
	if o == nil || IsNil(o.ErrorType) {
		var ret string
		return ret
	}
	return *o.ErrorType
}

// GetErrorTypeOk returns a tuple with the ErrorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShortcutError) GetErrorTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorType) {
		return nil, false
	}
	return o.ErrorType, true
}

// HasErrorType returns a boolean if a field has been set.
func (o *ShortcutError) HasErrorType() bool {
	if o != nil && !IsNil(o.ErrorType) {
		return true
	}

	return false
}

// SetErrorType gets a reference to the given string and assigns it to the ErrorType field.
func (o *ShortcutError) SetErrorType(v string) {
	o.ErrorType = &v
}

func (o ShortcutError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShortcutError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ErrorType) {
		toSerialize["errorType"] = o.ErrorType
	}
	return toSerialize, nil
}

type NullableShortcutError struct {
	value *ShortcutError
	isSet bool
}

func (v NullableShortcutError) Get() *ShortcutError {
	return v.value
}

func (v *NullableShortcutError) Set(val *ShortcutError) {
	v.value = val
	v.isSet = true
}

func (v NullableShortcutError) IsSet() bool {
	return v.isSet
}

func (v *NullableShortcutError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShortcutError(val *ShortcutError) *NullableShortcutError {
	return &NullableShortcutError{value: val, isSet: true}
}

func (v NullableShortcutError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShortcutError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

