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

// checks if the UpdateAnswerLikesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateAnswerLikesRequest{}

// UpdateAnswerLikesRequest struct for UpdateAnswerLikesRequest
type UpdateAnswerLikesRequest struct {
	// The opaque id of the answer to like.
	AnswerId int32 `json:"answerId"`
	// Internal document id of the answer. We support using the document id for cases where the client doesn't have the integer id available. If both are available, using the integer id is preferred.
	AnswerDocId *string `json:"answerDocId,omitempty"`
	Action string `json:"action"`
}

// NewUpdateAnswerLikesRequest instantiates a new UpdateAnswerLikesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateAnswerLikesRequest(answerId int32, action string) *UpdateAnswerLikesRequest {
	this := UpdateAnswerLikesRequest{}
	this.AnswerId = answerId
	this.Action = action
	return &this
}

// NewUpdateAnswerLikesRequestWithDefaults instantiates a new UpdateAnswerLikesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateAnswerLikesRequestWithDefaults() *UpdateAnswerLikesRequest {
	this := UpdateAnswerLikesRequest{}
	return &this
}

// GetAnswerId returns the AnswerId field value
func (o *UpdateAnswerLikesRequest) GetAnswerId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AnswerId
}

// GetAnswerIdOk returns a tuple with the AnswerId field value
// and a boolean to check if the value has been set.
func (o *UpdateAnswerLikesRequest) GetAnswerIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AnswerId, true
}

// SetAnswerId sets field value
func (o *UpdateAnswerLikesRequest) SetAnswerId(v int32) {
	o.AnswerId = v
}

// GetAnswerDocId returns the AnswerDocId field value if set, zero value otherwise.
func (o *UpdateAnswerLikesRequest) GetAnswerDocId() string {
	if o == nil || IsNil(o.AnswerDocId) {
		var ret string
		return ret
	}
	return *o.AnswerDocId
}

// GetAnswerDocIdOk returns a tuple with the AnswerDocId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAnswerLikesRequest) GetAnswerDocIdOk() (*string, bool) {
	if o == nil || IsNil(o.AnswerDocId) {
		return nil, false
	}
	return o.AnswerDocId, true
}

// HasAnswerDocId returns a boolean if a field has been set.
func (o *UpdateAnswerLikesRequest) HasAnswerDocId() bool {
	if o != nil && !IsNil(o.AnswerDocId) {
		return true
	}

	return false
}

// SetAnswerDocId gets a reference to the given string and assigns it to the AnswerDocId field.
func (o *UpdateAnswerLikesRequest) SetAnswerDocId(v string) {
	o.AnswerDocId = &v
}

// GetAction returns the Action field value
func (o *UpdateAnswerLikesRequest) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *UpdateAnswerLikesRequest) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *UpdateAnswerLikesRequest) SetAction(v string) {
	o.Action = v
}

func (o UpdateAnswerLikesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateAnswerLikesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["answerId"] = o.AnswerId
	if !IsNil(o.AnswerDocId) {
		toSerialize["answerDocId"] = o.AnswerDocId
	}
	toSerialize["action"] = o.Action
	return toSerialize, nil
}

type NullableUpdateAnswerLikesRequest struct {
	value *UpdateAnswerLikesRequest
	isSet bool
}

func (v NullableUpdateAnswerLikesRequest) Get() *UpdateAnswerLikesRequest {
	return v.value
}

func (v *NullableUpdateAnswerLikesRequest) Set(val *UpdateAnswerLikesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateAnswerLikesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateAnswerLikesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateAnswerLikesRequest(val *UpdateAnswerLikesRequest) *NullableUpdateAnswerLikesRequest {
	return &NullableUpdateAnswerLikesRequest{value: val, isSet: true}
}

func (v NullableUpdateAnswerLikesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateAnswerLikesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


