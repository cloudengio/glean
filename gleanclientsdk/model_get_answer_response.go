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

// checks if the GetAnswerResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAnswerResponse{}

// GetAnswerResponse struct for GetAnswerResponse
type GetAnswerResponse struct {
	AnswerResult *AnswerResult `json:"answerResult,omitempty"`
	Error *GetAnswerError `json:"error,omitempty"`
	Answer *Answer `json:"answer,omitempty"`
	// DEPRECATED - use answerResult tracking token instead. An opaque token that represents this particular answer. To be used for /feedback reporting.
	// Deprecated
	TrackingToken *string `json:"trackingToken,omitempty"`
}

// NewGetAnswerResponse instantiates a new GetAnswerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAnswerResponse() *GetAnswerResponse {
	this := GetAnswerResponse{}
	return &this
}

// NewGetAnswerResponseWithDefaults instantiates a new GetAnswerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAnswerResponseWithDefaults() *GetAnswerResponse {
	this := GetAnswerResponse{}
	return &this
}

// GetAnswerResult returns the AnswerResult field value if set, zero value otherwise.
func (o *GetAnswerResponse) GetAnswerResult() AnswerResult {
	if o == nil || IsNil(o.AnswerResult) {
		var ret AnswerResult
		return ret
	}
	return *o.AnswerResult
}

// GetAnswerResultOk returns a tuple with the AnswerResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAnswerResponse) GetAnswerResultOk() (*AnswerResult, bool) {
	if o == nil || IsNil(o.AnswerResult) {
		return nil, false
	}
	return o.AnswerResult, true
}

// HasAnswerResult returns a boolean if a field has been set.
func (o *GetAnswerResponse) HasAnswerResult() bool {
	if o != nil && !IsNil(o.AnswerResult) {
		return true
	}

	return false
}

// SetAnswerResult gets a reference to the given AnswerResult and assigns it to the AnswerResult field.
func (o *GetAnswerResponse) SetAnswerResult(v AnswerResult) {
	o.AnswerResult = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *GetAnswerResponse) GetError() GetAnswerError {
	if o == nil || IsNil(o.Error) {
		var ret GetAnswerError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAnswerResponse) GetErrorOk() (*GetAnswerError, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *GetAnswerResponse) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given GetAnswerError and assigns it to the Error field.
func (o *GetAnswerResponse) SetError(v GetAnswerError) {
	o.Error = &v
}

// GetAnswer returns the Answer field value if set, zero value otherwise.
func (o *GetAnswerResponse) GetAnswer() Answer {
	if o == nil || IsNil(o.Answer) {
		var ret Answer
		return ret
	}
	return *o.Answer
}

// GetAnswerOk returns a tuple with the Answer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAnswerResponse) GetAnswerOk() (*Answer, bool) {
	if o == nil || IsNil(o.Answer) {
		return nil, false
	}
	return o.Answer, true
}

// HasAnswer returns a boolean if a field has been set.
func (o *GetAnswerResponse) HasAnswer() bool {
	if o != nil && !IsNil(o.Answer) {
		return true
	}

	return false
}

// SetAnswer gets a reference to the given Answer and assigns it to the Answer field.
func (o *GetAnswerResponse) SetAnswer(v Answer) {
	o.Answer = &v
}

// GetTrackingToken returns the TrackingToken field value if set, zero value otherwise.
// Deprecated
func (o *GetAnswerResponse) GetTrackingToken() string {
	if o == nil || IsNil(o.TrackingToken) {
		var ret string
		return ret
	}
	return *o.TrackingToken
}

// GetTrackingTokenOk returns a tuple with the TrackingToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *GetAnswerResponse) GetTrackingTokenOk() (*string, bool) {
	if o == nil || IsNil(o.TrackingToken) {
		return nil, false
	}
	return o.TrackingToken, true
}

// HasTrackingToken returns a boolean if a field has been set.
func (o *GetAnswerResponse) HasTrackingToken() bool {
	if o != nil && !IsNil(o.TrackingToken) {
		return true
	}

	return false
}

// SetTrackingToken gets a reference to the given string and assigns it to the TrackingToken field.
// Deprecated
func (o *GetAnswerResponse) SetTrackingToken(v string) {
	o.TrackingToken = &v
}

func (o GetAnswerResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAnswerResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AnswerResult) {
		toSerialize["answerResult"] = o.AnswerResult
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !IsNil(o.Answer) {
		toSerialize["answer"] = o.Answer
	}
	if !IsNil(o.TrackingToken) {
		toSerialize["trackingToken"] = o.TrackingToken
	}
	return toSerialize, nil
}

type NullableGetAnswerResponse struct {
	value *GetAnswerResponse
	isSet bool
}

func (v NullableGetAnswerResponse) Get() *GetAnswerResponse {
	return v.value
}

func (v *NullableGetAnswerResponse) Set(val *GetAnswerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAnswerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAnswerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAnswerResponse(val *GetAnswerResponse) *NullableGetAnswerResponse {
	return &NullableGetAnswerResponse{value: val, isSet: true}
}

func (v NullableGetAnswerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAnswerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


