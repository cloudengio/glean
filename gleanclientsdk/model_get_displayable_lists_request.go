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

// checks if the GetDisplayableListsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDisplayableListsRequest{}

// GetDisplayableListsRequest struct for GetDisplayableListsRequest
type GetDisplayableListsRequest struct {
	Ids []int32 `json:"ids"`
}

// NewGetDisplayableListsRequest instantiates a new GetDisplayableListsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDisplayableListsRequest(ids []int32) *GetDisplayableListsRequest {
	this := GetDisplayableListsRequest{}
	this.Ids = ids
	return &this
}

// NewGetDisplayableListsRequestWithDefaults instantiates a new GetDisplayableListsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDisplayableListsRequestWithDefaults() *GetDisplayableListsRequest {
	this := GetDisplayableListsRequest{}
	return &this
}

// GetIds returns the Ids field value
func (o *GetDisplayableListsRequest) GetIds() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value
// and a boolean to check if the value has been set.
func (o *GetDisplayableListsRequest) GetIdsOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ids, true
}

// SetIds sets field value
func (o *GetDisplayableListsRequest) SetIds(v []int32) {
	o.Ids = v
}

func (o GetDisplayableListsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDisplayableListsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ids"] = o.Ids
	return toSerialize, nil
}

type NullableGetDisplayableListsRequest struct {
	value *GetDisplayableListsRequest
	isSet bool
}

func (v NullableGetDisplayableListsRequest) Get() *GetDisplayableListsRequest {
	return v.value
}

func (v *NullableGetDisplayableListsRequest) Set(val *GetDisplayableListsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDisplayableListsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDisplayableListsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDisplayableListsRequest(val *GetDisplayableListsRequest) *NullableGetDisplayableListsRequest {
	return &NullableGetDisplayableListsRequest{value: val, isSet: true}
}

func (v NullableGetDisplayableListsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDisplayableListsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


