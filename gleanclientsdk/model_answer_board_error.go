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

// checks if the AnswerBoardError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnswerBoardError{}

// AnswerBoardError struct for AnswerBoardError
type AnswerBoardError struct {
	ErrorCode string `json:"errorCode"`
}

// NewAnswerBoardError instantiates a new AnswerBoardError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnswerBoardError(errorCode string) *AnswerBoardError {
	this := AnswerBoardError{}
	this.ErrorCode = errorCode
	return &this
}

// NewAnswerBoardErrorWithDefaults instantiates a new AnswerBoardError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnswerBoardErrorWithDefaults() *AnswerBoardError {
	this := AnswerBoardError{}
	return &this
}

// GetErrorCode returns the ErrorCode field value
func (o *AnswerBoardError) GetErrorCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value
// and a boolean to check if the value has been set.
func (o *AnswerBoardError) GetErrorCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorCode, true
}

// SetErrorCode sets field value
func (o *AnswerBoardError) SetErrorCode(v string) {
	o.ErrorCode = v
}

func (o AnswerBoardError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnswerBoardError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["errorCode"] = o.ErrorCode
	return toSerialize, nil
}

type NullableAnswerBoardError struct {
	value *AnswerBoardError
	isSet bool
}

func (v NullableAnswerBoardError) Get() *AnswerBoardError {
	return v.value
}

func (v *NullableAnswerBoardError) Set(val *AnswerBoardError) {
	v.value = val
	v.isSet = true
}

func (v NullableAnswerBoardError) IsSet() bool {
	return v.isSet
}

func (v *NullableAnswerBoardError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnswerBoardError(val *AnswerBoardError) *NullableAnswerBoardError {
	return &NullableAnswerBoardError{value: val, isSet: true}
}

func (v NullableAnswerBoardError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnswerBoardError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

