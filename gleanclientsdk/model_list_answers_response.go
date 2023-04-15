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

// checks if the ListAnswersResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListAnswersResponse{}

// ListAnswersResponse struct for ListAnswersResponse
type ListAnswersResponse struct {
	// List of answers with tracking tokens.
	AnswerResults []AnswerResult `json:"answerResults"`
	// DEPRECATED - use permissions instead. User's role for Answers in their workplace.
	// Deprecated
	UserRole *string `json:"userRole,omitempty"`
}

// NewListAnswersResponse instantiates a new ListAnswersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListAnswersResponse(answerResults []AnswerResult) *ListAnswersResponse {
	this := ListAnswersResponse{}
	this.AnswerResults = answerResults
	return &this
}

// NewListAnswersResponseWithDefaults instantiates a new ListAnswersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListAnswersResponseWithDefaults() *ListAnswersResponse {
	this := ListAnswersResponse{}
	return &this
}

// GetAnswerResults returns the AnswerResults field value
func (o *ListAnswersResponse) GetAnswerResults() []AnswerResult {
	if o == nil {
		var ret []AnswerResult
		return ret
	}

	return o.AnswerResults
}

// GetAnswerResultsOk returns a tuple with the AnswerResults field value
// and a boolean to check if the value has been set.
func (o *ListAnswersResponse) GetAnswerResultsOk() ([]AnswerResult, bool) {
	if o == nil {
		return nil, false
	}
	return o.AnswerResults, true
}

// SetAnswerResults sets field value
func (o *ListAnswersResponse) SetAnswerResults(v []AnswerResult) {
	o.AnswerResults = v
}

// GetUserRole returns the UserRole field value if set, zero value otherwise.
// Deprecated
func (o *ListAnswersResponse) GetUserRole() string {
	if o == nil || IsNil(o.UserRole) {
		var ret string
		return ret
	}
	return *o.UserRole
}

// GetUserRoleOk returns a tuple with the UserRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *ListAnswersResponse) GetUserRoleOk() (*string, bool) {
	if o == nil || IsNil(o.UserRole) {
		return nil, false
	}
	return o.UserRole, true
}

// HasUserRole returns a boolean if a field has been set.
func (o *ListAnswersResponse) HasUserRole() bool {
	if o != nil && !IsNil(o.UserRole) {
		return true
	}

	return false
}

// SetUserRole gets a reference to the given string and assigns it to the UserRole field.
// Deprecated
func (o *ListAnswersResponse) SetUserRole(v string) {
	o.UserRole = &v
}

func (o ListAnswersResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListAnswersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["answerResults"] = o.AnswerResults
	if !IsNil(o.UserRole) {
		toSerialize["userRole"] = o.UserRole
	}
	return toSerialize, nil
}

type NullableListAnswersResponse struct {
	value *ListAnswersResponse
	isSet bool
}

func (v NullableListAnswersResponse) Get() *ListAnswersResponse {
	return v.value
}

func (v *NullableListAnswersResponse) Set(val *ListAnswersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListAnswersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListAnswersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListAnswersResponse(val *ListAnswersResponse) *NullableListAnswersResponse {
	return &NullableListAnswersResponse{value: val, isSet: true}
}

func (v NullableListAnswersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListAnswersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


