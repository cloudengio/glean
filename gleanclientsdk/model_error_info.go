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

// checks if the ErrorInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorInfo{}

// ErrorInfo struct for ErrorInfo
type ErrorInfo struct {
	// Indicates the gmail results could not be fetched due to bad token.
	BadGmailToken *bool `json:"badGmailToken,omitempty"`
	// Indicates the outlook results could not be fetched due to bad token.
	BadOutlookToken *bool `json:"badOutlookToken,omitempty"`
	// Indicates results could not be fetched due to invalid operators in the query.
	InvalidOperators []InvalidOperatorValueError `json:"invalidOperators,omitempty"`
	ErrorMessages []ErrorMessage `json:"errorMessages,omitempty"`
}

// NewErrorInfo instantiates a new ErrorInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorInfo() *ErrorInfo {
	this := ErrorInfo{}
	return &this
}

// NewErrorInfoWithDefaults instantiates a new ErrorInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorInfoWithDefaults() *ErrorInfo {
	this := ErrorInfo{}
	return &this
}

// GetBadGmailToken returns the BadGmailToken field value if set, zero value otherwise.
func (o *ErrorInfo) GetBadGmailToken() bool {
	if o == nil || IsNil(o.BadGmailToken) {
		var ret bool
		return ret
	}
	return *o.BadGmailToken
}

// GetBadGmailTokenOk returns a tuple with the BadGmailToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorInfo) GetBadGmailTokenOk() (*bool, bool) {
	if o == nil || IsNil(o.BadGmailToken) {
		return nil, false
	}
	return o.BadGmailToken, true
}

// HasBadGmailToken returns a boolean if a field has been set.
func (o *ErrorInfo) HasBadGmailToken() bool {
	if o != nil && !IsNil(o.BadGmailToken) {
		return true
	}

	return false
}

// SetBadGmailToken gets a reference to the given bool and assigns it to the BadGmailToken field.
func (o *ErrorInfo) SetBadGmailToken(v bool) {
	o.BadGmailToken = &v
}

// GetBadOutlookToken returns the BadOutlookToken field value if set, zero value otherwise.
func (o *ErrorInfo) GetBadOutlookToken() bool {
	if o == nil || IsNil(o.BadOutlookToken) {
		var ret bool
		return ret
	}
	return *o.BadOutlookToken
}

// GetBadOutlookTokenOk returns a tuple with the BadOutlookToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorInfo) GetBadOutlookTokenOk() (*bool, bool) {
	if o == nil || IsNil(o.BadOutlookToken) {
		return nil, false
	}
	return o.BadOutlookToken, true
}

// HasBadOutlookToken returns a boolean if a field has been set.
func (o *ErrorInfo) HasBadOutlookToken() bool {
	if o != nil && !IsNil(o.BadOutlookToken) {
		return true
	}

	return false
}

// SetBadOutlookToken gets a reference to the given bool and assigns it to the BadOutlookToken field.
func (o *ErrorInfo) SetBadOutlookToken(v bool) {
	o.BadOutlookToken = &v
}

// GetInvalidOperators returns the InvalidOperators field value if set, zero value otherwise.
func (o *ErrorInfo) GetInvalidOperators() []InvalidOperatorValueError {
	if o == nil || IsNil(o.InvalidOperators) {
		var ret []InvalidOperatorValueError
		return ret
	}
	return o.InvalidOperators
}

// GetInvalidOperatorsOk returns a tuple with the InvalidOperators field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorInfo) GetInvalidOperatorsOk() ([]InvalidOperatorValueError, bool) {
	if o == nil || IsNil(o.InvalidOperators) {
		return nil, false
	}
	return o.InvalidOperators, true
}

// HasInvalidOperators returns a boolean if a field has been set.
func (o *ErrorInfo) HasInvalidOperators() bool {
	if o != nil && !IsNil(o.InvalidOperators) {
		return true
	}

	return false
}

// SetInvalidOperators gets a reference to the given []InvalidOperatorValueError and assigns it to the InvalidOperators field.
func (o *ErrorInfo) SetInvalidOperators(v []InvalidOperatorValueError) {
	o.InvalidOperators = v
}

// GetErrorMessages returns the ErrorMessages field value if set, zero value otherwise.
func (o *ErrorInfo) GetErrorMessages() []ErrorMessage {
	if o == nil || IsNil(o.ErrorMessages) {
		var ret []ErrorMessage
		return ret
	}
	return o.ErrorMessages
}

// GetErrorMessagesOk returns a tuple with the ErrorMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorInfo) GetErrorMessagesOk() ([]ErrorMessage, bool) {
	if o == nil || IsNil(o.ErrorMessages) {
		return nil, false
	}
	return o.ErrorMessages, true
}

// HasErrorMessages returns a boolean if a field has been set.
func (o *ErrorInfo) HasErrorMessages() bool {
	if o != nil && !IsNil(o.ErrorMessages) {
		return true
	}

	return false
}

// SetErrorMessages gets a reference to the given []ErrorMessage and assigns it to the ErrorMessages field.
func (o *ErrorInfo) SetErrorMessages(v []ErrorMessage) {
	o.ErrorMessages = v
}

func (o ErrorInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BadGmailToken) {
		toSerialize["badGmailToken"] = o.BadGmailToken
	}
	if !IsNil(o.BadOutlookToken) {
		toSerialize["badOutlookToken"] = o.BadOutlookToken
	}
	if !IsNil(o.InvalidOperators) {
		toSerialize["invalidOperators"] = o.InvalidOperators
	}
	if !IsNil(o.ErrorMessages) {
		toSerialize["errorMessages"] = o.ErrorMessages
	}
	return toSerialize, nil
}

type NullableErrorInfo struct {
	value *ErrorInfo
	isSet bool
}

func (v NullableErrorInfo) Get() *ErrorInfo {
	return v.value
}

func (v *NullableErrorInfo) Set(val *ErrorInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorInfo(val *ErrorInfo) *NullableErrorInfo {
	return &NullableErrorInfo{value: val, isSet: true}
}

func (v NullableErrorInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


