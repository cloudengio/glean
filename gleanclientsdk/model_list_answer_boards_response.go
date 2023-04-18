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

// checks if the ListAnswerBoardsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListAnswerBoardsResponse{}

// ListAnswerBoardsResponse struct for ListAnswerBoardsResponse
type ListAnswerBoardsResponse struct {
	// List of all Answer Boards, no Answers are included.
	BoardResults []AnswerBoardResult `json:"boardResults"`
}

// NewListAnswerBoardsResponse instantiates a new ListAnswerBoardsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListAnswerBoardsResponse(boardResults []AnswerBoardResult) *ListAnswerBoardsResponse {
	this := ListAnswerBoardsResponse{}
	this.BoardResults = boardResults
	return &this
}

// NewListAnswerBoardsResponseWithDefaults instantiates a new ListAnswerBoardsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListAnswerBoardsResponseWithDefaults() *ListAnswerBoardsResponse {
	this := ListAnswerBoardsResponse{}
	return &this
}

// GetBoardResults returns the BoardResults field value
func (o *ListAnswerBoardsResponse) GetBoardResults() []AnswerBoardResult {
	if o == nil {
		var ret []AnswerBoardResult
		return ret
	}

	return o.BoardResults
}

// GetBoardResultsOk returns a tuple with the BoardResults field value
// and a boolean to check if the value has been set.
func (o *ListAnswerBoardsResponse) GetBoardResultsOk() ([]AnswerBoardResult, bool) {
	if o == nil {
		return nil, false
	}
	return o.BoardResults, true
}

// SetBoardResults sets field value
func (o *ListAnswerBoardsResponse) SetBoardResults(v []AnswerBoardResult) {
	o.BoardResults = v
}

func (o ListAnswerBoardsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListAnswerBoardsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["boardResults"] = o.BoardResults
	return toSerialize, nil
}

type NullableListAnswerBoardsResponse struct {
	value *ListAnswerBoardsResponse
	isSet bool
}

func (v NullableListAnswerBoardsResponse) Get() *ListAnswerBoardsResponse {
	return v.value
}

func (v *NullableListAnswerBoardsResponse) Set(val *ListAnswerBoardsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListAnswerBoardsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListAnswerBoardsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListAnswerBoardsResponse(val *ListAnswerBoardsResponse) *NullableListAnswerBoardsResponse {
	return &NullableListAnswerBoardsResponse{value: val, isSet: true}
}

func (v NullableListAnswerBoardsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListAnswerBoardsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


