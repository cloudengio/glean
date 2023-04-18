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

// checks if the CreateTeamsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTeamsResponse{}

// CreateTeamsResponse struct for CreateTeamsResponse
type CreateTeamsResponse struct {
	// Number of teams that failed to be created
	NumErrors *int32 `json:"numErrors,omitempty"`
}

// NewCreateTeamsResponse instantiates a new CreateTeamsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTeamsResponse() *CreateTeamsResponse {
	this := CreateTeamsResponse{}
	return &this
}

// NewCreateTeamsResponseWithDefaults instantiates a new CreateTeamsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTeamsResponseWithDefaults() *CreateTeamsResponse {
	this := CreateTeamsResponse{}
	return &this
}

// GetNumErrors returns the NumErrors field value if set, zero value otherwise.
func (o *CreateTeamsResponse) GetNumErrors() int32 {
	if o == nil || IsNil(o.NumErrors) {
		var ret int32
		return ret
	}
	return *o.NumErrors
}

// GetNumErrorsOk returns a tuple with the NumErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTeamsResponse) GetNumErrorsOk() (*int32, bool) {
	if o == nil || IsNil(o.NumErrors) {
		return nil, false
	}
	return o.NumErrors, true
}

// HasNumErrors returns a boolean if a field has been set.
func (o *CreateTeamsResponse) HasNumErrors() bool {
	if o != nil && !IsNil(o.NumErrors) {
		return true
	}

	return false
}

// SetNumErrors gets a reference to the given int32 and assigns it to the NumErrors field.
func (o *CreateTeamsResponse) SetNumErrors(v int32) {
	o.NumErrors = &v
}

func (o CreateTeamsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTeamsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NumErrors) {
		toSerialize["numErrors"] = o.NumErrors
	}
	return toSerialize, nil
}

type NullableCreateTeamsResponse struct {
	value *CreateTeamsResponse
	isSet bool
}

func (v NullableCreateTeamsResponse) Get() *CreateTeamsResponse {
	return v.value
}

func (v *NullableCreateTeamsResponse) Set(val *CreateTeamsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTeamsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTeamsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTeamsResponse(val *CreateTeamsResponse) *NullableCreateTeamsResponse {
	return &NullableCreateTeamsResponse{value: val, isSet: true}
}

func (v NullableCreateTeamsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTeamsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


