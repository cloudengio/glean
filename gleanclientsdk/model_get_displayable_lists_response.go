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

// checks if the GetDisplayableListsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDisplayableListsResponse{}

// GetDisplayableListsResponse struct for GetDisplayableListsResponse
type GetDisplayableListsResponse struct {
	Items []DisplayableList `json:"items"`
}

// NewGetDisplayableListsResponse instantiates a new GetDisplayableListsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDisplayableListsResponse(items []DisplayableList) *GetDisplayableListsResponse {
	this := GetDisplayableListsResponse{}
	this.Items = items
	return &this
}

// NewGetDisplayableListsResponseWithDefaults instantiates a new GetDisplayableListsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDisplayableListsResponseWithDefaults() *GetDisplayableListsResponse {
	this := GetDisplayableListsResponse{}
	return &this
}

// GetItems returns the Items field value
func (o *GetDisplayableListsResponse) GetItems() []DisplayableList {
	if o == nil {
		var ret []DisplayableList
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *GetDisplayableListsResponse) GetItemsOk() ([]DisplayableList, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *GetDisplayableListsResponse) SetItems(v []DisplayableList) {
	o.Items = v
}

func (o GetDisplayableListsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDisplayableListsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

type NullableGetDisplayableListsResponse struct {
	value *GetDisplayableListsResponse
	isSet bool
}

func (v NullableGetDisplayableListsResponse) Get() *GetDisplayableListsResponse {
	return v.value
}

func (v *NullableGetDisplayableListsResponse) Set(val *GetDisplayableListsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDisplayableListsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDisplayableListsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDisplayableListsResponse(val *GetDisplayableListsResponse) *NullableGetDisplayableListsResponse {
	return &NullableGetDisplayableListsResponse{value: val, isSet: true}
}

func (v NullableGetDisplayableListsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDisplayableListsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


