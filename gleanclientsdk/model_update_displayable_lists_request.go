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

// checks if the UpdateDisplayableListsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateDisplayableListsRequest{}

// UpdateDisplayableListsRequest struct for UpdateDisplayableListsRequest
type UpdateDisplayableListsRequest struct {
	Items []DisplayableList `json:"items"`
}

// NewUpdateDisplayableListsRequest instantiates a new UpdateDisplayableListsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDisplayableListsRequest(items []DisplayableList) *UpdateDisplayableListsRequest {
	this := UpdateDisplayableListsRequest{}
	this.Items = items
	return &this
}

// NewUpdateDisplayableListsRequestWithDefaults instantiates a new UpdateDisplayableListsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDisplayableListsRequestWithDefaults() *UpdateDisplayableListsRequest {
	this := UpdateDisplayableListsRequest{}
	return &this
}

// GetItems returns the Items field value
func (o *UpdateDisplayableListsRequest) GetItems() []DisplayableList {
	if o == nil {
		var ret []DisplayableList
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *UpdateDisplayableListsRequest) GetItemsOk() ([]DisplayableList, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *UpdateDisplayableListsRequest) SetItems(v []DisplayableList) {
	o.Items = v
}

func (o UpdateDisplayableListsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateDisplayableListsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

type NullableUpdateDisplayableListsRequest struct {
	value *UpdateDisplayableListsRequest
	isSet bool
}

func (v NullableUpdateDisplayableListsRequest) Get() *UpdateDisplayableListsRequest {
	return v.value
}

func (v *NullableUpdateDisplayableListsRequest) Set(val *UpdateDisplayableListsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDisplayableListsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDisplayableListsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDisplayableListsRequest(val *UpdateDisplayableListsRequest) *NullableUpdateDisplayableListsRequest {
	return &NullableUpdateDisplayableListsRequest{value: val, isSet: true}
}

func (v NullableUpdateDisplayableListsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDisplayableListsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


