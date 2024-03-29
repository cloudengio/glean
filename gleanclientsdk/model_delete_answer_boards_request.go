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

// checks if the DeleteAnswerBoardsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteAnswerBoardsRequest{}

// DeleteAnswerBoardsRequest struct for DeleteAnswerBoardsRequest
type DeleteAnswerBoardsRequest struct {
	// The IDs of the Answer Boards to delete.
	Ids []int32 `json:"ids"`
}

// NewDeleteAnswerBoardsRequest instantiates a new DeleteAnswerBoardsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteAnswerBoardsRequest(ids []int32) *DeleteAnswerBoardsRequest {
	this := DeleteAnswerBoardsRequest{}
	this.Ids = ids
	return &this
}

// NewDeleteAnswerBoardsRequestWithDefaults instantiates a new DeleteAnswerBoardsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteAnswerBoardsRequestWithDefaults() *DeleteAnswerBoardsRequest {
	this := DeleteAnswerBoardsRequest{}
	return &this
}

// GetIds returns the Ids field value
func (o *DeleteAnswerBoardsRequest) GetIds() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value
// and a boolean to check if the value has been set.
func (o *DeleteAnswerBoardsRequest) GetIdsOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ids, true
}

// SetIds sets field value
func (o *DeleteAnswerBoardsRequest) SetIds(v []int32) {
	o.Ids = v
}

func (o DeleteAnswerBoardsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteAnswerBoardsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ids"] = o.Ids
	return toSerialize, nil
}

type NullableDeleteAnswerBoardsRequest struct {
	value *DeleteAnswerBoardsRequest
	isSet bool
}

func (v NullableDeleteAnswerBoardsRequest) Get() *DeleteAnswerBoardsRequest {
	return v.value
}

func (v *NullableDeleteAnswerBoardsRequest) Set(val *DeleteAnswerBoardsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteAnswerBoardsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteAnswerBoardsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteAnswerBoardsRequest(val *DeleteAnswerBoardsRequest) *NullableDeleteAnswerBoardsRequest {
	return &NullableDeleteAnswerBoardsRequest{value: val, isSet: true}
}

func (v NullableDeleteAnswerBoardsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteAnswerBoardsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


