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

// checks if the GetGeneratedQnaResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetGeneratedQnaResponse{}

// GetGeneratedQnaResponse struct for GetGeneratedQnaResponse
type GetGeneratedQnaResponse struct {
	GeneratedQnaResult *GeneratedQna `json:"generatedQnaResult,omitempty"`
}

// NewGetGeneratedQnaResponse instantiates a new GetGeneratedQnaResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetGeneratedQnaResponse() *GetGeneratedQnaResponse {
	this := GetGeneratedQnaResponse{}
	return &this
}

// NewGetGeneratedQnaResponseWithDefaults instantiates a new GetGeneratedQnaResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetGeneratedQnaResponseWithDefaults() *GetGeneratedQnaResponse {
	this := GetGeneratedQnaResponse{}
	return &this
}

// GetGeneratedQnaResult returns the GeneratedQnaResult field value if set, zero value otherwise.
func (o *GetGeneratedQnaResponse) GetGeneratedQnaResult() GeneratedQna {
	if o == nil || IsNil(o.GeneratedQnaResult) {
		var ret GeneratedQna
		return ret
	}
	return *o.GeneratedQnaResult
}

// GetGeneratedQnaResultOk returns a tuple with the GeneratedQnaResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetGeneratedQnaResponse) GetGeneratedQnaResultOk() (*GeneratedQna, bool) {
	if o == nil || IsNil(o.GeneratedQnaResult) {
		return nil, false
	}
	return o.GeneratedQnaResult, true
}

// HasGeneratedQnaResult returns a boolean if a field has been set.
func (o *GetGeneratedQnaResponse) HasGeneratedQnaResult() bool {
	if o != nil && !IsNil(o.GeneratedQnaResult) {
		return true
	}

	return false
}

// SetGeneratedQnaResult gets a reference to the given GeneratedQna and assigns it to the GeneratedQnaResult field.
func (o *GetGeneratedQnaResponse) SetGeneratedQnaResult(v GeneratedQna) {
	o.GeneratedQnaResult = &v
}

func (o GetGeneratedQnaResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetGeneratedQnaResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GeneratedQnaResult) {
		toSerialize["generatedQnaResult"] = o.GeneratedQnaResult
	}
	return toSerialize, nil
}

type NullableGetGeneratedQnaResponse struct {
	value *GetGeneratedQnaResponse
	isSet bool
}

func (v NullableGetGeneratedQnaResponse) Get() *GetGeneratedQnaResponse {
	return v.value
}

func (v *NullableGetGeneratedQnaResponse) Set(val *GetGeneratedQnaResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetGeneratedQnaResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetGeneratedQnaResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetGeneratedQnaResponse(val *GetGeneratedQnaResponse) *NullableGetGeneratedQnaResponse {
	return &NullableGetGeneratedQnaResponse{value: val, isSet: true}
}

func (v NullableGetGeneratedQnaResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetGeneratedQnaResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


