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

// checks if the GetAnswerBoardResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAnswerBoardResponse{}

// GetAnswerBoardResponse struct for GetAnswerBoardResponse
type GetAnswerBoardResponse struct {
	BoardResult *AnswerBoardResult `json:"boardResult,omitempty"`
	Error *AnswerBoardError `json:"error,omitempty"`
}

// NewGetAnswerBoardResponse instantiates a new GetAnswerBoardResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAnswerBoardResponse() *GetAnswerBoardResponse {
	this := GetAnswerBoardResponse{}
	return &this
}

// NewGetAnswerBoardResponseWithDefaults instantiates a new GetAnswerBoardResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAnswerBoardResponseWithDefaults() *GetAnswerBoardResponse {
	this := GetAnswerBoardResponse{}
	return &this
}

// GetBoardResult returns the BoardResult field value if set, zero value otherwise.
func (o *GetAnswerBoardResponse) GetBoardResult() AnswerBoardResult {
	if o == nil || IsNil(o.BoardResult) {
		var ret AnswerBoardResult
		return ret
	}
	return *o.BoardResult
}

// GetBoardResultOk returns a tuple with the BoardResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAnswerBoardResponse) GetBoardResultOk() (*AnswerBoardResult, bool) {
	if o == nil || IsNil(o.BoardResult) {
		return nil, false
	}
	return o.BoardResult, true
}

// HasBoardResult returns a boolean if a field has been set.
func (o *GetAnswerBoardResponse) HasBoardResult() bool {
	if o != nil && !IsNil(o.BoardResult) {
		return true
	}

	return false
}

// SetBoardResult gets a reference to the given AnswerBoardResult and assigns it to the BoardResult field.
func (o *GetAnswerBoardResponse) SetBoardResult(v AnswerBoardResult) {
	o.BoardResult = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *GetAnswerBoardResponse) GetError() AnswerBoardError {
	if o == nil || IsNil(o.Error) {
		var ret AnswerBoardError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAnswerBoardResponse) GetErrorOk() (*AnswerBoardError, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *GetAnswerBoardResponse) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given AnswerBoardError and assigns it to the Error field.
func (o *GetAnswerBoardResponse) SetError(v AnswerBoardError) {
	o.Error = &v
}

func (o GetAnswerBoardResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAnswerBoardResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BoardResult) {
		toSerialize["boardResult"] = o.BoardResult
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableGetAnswerBoardResponse struct {
	value *GetAnswerBoardResponse
	isSet bool
}

func (v NullableGetAnswerBoardResponse) Get() *GetAnswerBoardResponse {
	return v.value
}

func (v *NullableGetAnswerBoardResponse) Set(val *GetAnswerBoardResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAnswerBoardResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAnswerBoardResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAnswerBoardResponse(val *GetAnswerBoardResponse) *NullableGetAnswerBoardResponse {
	return &NullableGetAnswerBoardResponse{value: val, isSet: true}
}

func (v NullableGetAnswerBoardResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAnswerBoardResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


