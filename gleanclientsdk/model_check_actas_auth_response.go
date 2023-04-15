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

// checks if the CheckActasAuthResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CheckActasAuthResponse{}

// CheckActasAuthResponse struct for CheckActasAuthResponse
type CheckActasAuthResponse struct {
	// Whether the request supplied a valid actas auth token.
	IsValid bool `json:"isValid"`
	// Error message
	ErrMsg *string `json:"errMsg,omitempty"`
}

// NewCheckActasAuthResponse instantiates a new CheckActasAuthResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckActasAuthResponse(isValid bool) *CheckActasAuthResponse {
	this := CheckActasAuthResponse{}
	this.IsValid = isValid
	return &this
}

// NewCheckActasAuthResponseWithDefaults instantiates a new CheckActasAuthResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckActasAuthResponseWithDefaults() *CheckActasAuthResponse {
	this := CheckActasAuthResponse{}
	return &this
}

// GetIsValid returns the IsValid field value
func (o *CheckActasAuthResponse) GetIsValid() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsValid
}

// GetIsValidOk returns a tuple with the IsValid field value
// and a boolean to check if the value has been set.
func (o *CheckActasAuthResponse) GetIsValidOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsValid, true
}

// SetIsValid sets field value
func (o *CheckActasAuthResponse) SetIsValid(v bool) {
	o.IsValid = v
}

// GetErrMsg returns the ErrMsg field value if set, zero value otherwise.
func (o *CheckActasAuthResponse) GetErrMsg() string {
	if o == nil || IsNil(o.ErrMsg) {
		var ret string
		return ret
	}
	return *o.ErrMsg
}

// GetErrMsgOk returns a tuple with the ErrMsg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckActasAuthResponse) GetErrMsgOk() (*string, bool) {
	if o == nil || IsNil(o.ErrMsg) {
		return nil, false
	}
	return o.ErrMsg, true
}

// HasErrMsg returns a boolean if a field has been set.
func (o *CheckActasAuthResponse) HasErrMsg() bool {
	if o != nil && !IsNil(o.ErrMsg) {
		return true
	}

	return false
}

// SetErrMsg gets a reference to the given string and assigns it to the ErrMsg field.
func (o *CheckActasAuthResponse) SetErrMsg(v string) {
	o.ErrMsg = &v
}

func (o CheckActasAuthResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckActasAuthResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["isValid"] = o.IsValid
	if !IsNil(o.ErrMsg) {
		toSerialize["errMsg"] = o.ErrMsg
	}
	return toSerialize, nil
}

type NullableCheckActasAuthResponse struct {
	value *CheckActasAuthResponse
	isSet bool
}

func (v NullableCheckActasAuthResponse) Get() *CheckActasAuthResponse {
	return v.value
}

func (v *NullableCheckActasAuthResponse) Set(val *CheckActasAuthResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckActasAuthResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckActasAuthResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckActasAuthResponse(val *CheckActasAuthResponse) *NullableCheckActasAuthResponse {
	return &NullableCheckActasAuthResponse{value: val, isSet: true}
}

func (v NullableCheckActasAuthResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckActasAuthResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


