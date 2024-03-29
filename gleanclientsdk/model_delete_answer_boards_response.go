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

// checks if the DeleteAnswerBoardsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteAnswerBoardsResponse{}

// DeleteAnswerBoardsResponse struct for DeleteAnswerBoardsResponse
type DeleteAnswerBoardsResponse struct {
	Error *AnswerBoardError `json:"error,omitempty"`
}

// NewDeleteAnswerBoardsResponse instantiates a new DeleteAnswerBoardsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteAnswerBoardsResponse() *DeleteAnswerBoardsResponse {
	this := DeleteAnswerBoardsResponse{}
	return &this
}

// NewDeleteAnswerBoardsResponseWithDefaults instantiates a new DeleteAnswerBoardsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteAnswerBoardsResponseWithDefaults() *DeleteAnswerBoardsResponse {
	this := DeleteAnswerBoardsResponse{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *DeleteAnswerBoardsResponse) GetError() AnswerBoardError {
	if o == nil || IsNil(o.Error) {
		var ret AnswerBoardError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteAnswerBoardsResponse) GetErrorOk() (*AnswerBoardError, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *DeleteAnswerBoardsResponse) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given AnswerBoardError and assigns it to the Error field.
func (o *DeleteAnswerBoardsResponse) SetError(v AnswerBoardError) {
	o.Error = &v
}

func (o DeleteAnswerBoardsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteAnswerBoardsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableDeleteAnswerBoardsResponse struct {
	value *DeleteAnswerBoardsResponse
	isSet bool
}

func (v NullableDeleteAnswerBoardsResponse) Get() *DeleteAnswerBoardsResponse {
	return v.value
}

func (v *NullableDeleteAnswerBoardsResponse) Set(val *DeleteAnswerBoardsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteAnswerBoardsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteAnswerBoardsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteAnswerBoardsResponse(val *DeleteAnswerBoardsResponse) *NullableDeleteAnswerBoardsResponse {
	return &NullableDeleteAnswerBoardsResponse{value: val, isSet: true}
}

func (v NullableDeleteAnswerBoardsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteAnswerBoardsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


