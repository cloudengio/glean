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

// checks if the ListPinsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListPinsResponse{}

// ListPinsResponse struct for ListPinsResponse
type ListPinsResponse struct {
	// List of pinned documents.
	Pins []PinDocument `json:"pins"`
}

// NewListPinsResponse instantiates a new ListPinsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListPinsResponse(pins []PinDocument) *ListPinsResponse {
	this := ListPinsResponse{}
	this.Pins = pins
	return &this
}

// NewListPinsResponseWithDefaults instantiates a new ListPinsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListPinsResponseWithDefaults() *ListPinsResponse {
	this := ListPinsResponse{}
	return &this
}

// GetPins returns the Pins field value
func (o *ListPinsResponse) GetPins() []PinDocument {
	if o == nil {
		var ret []PinDocument
		return ret
	}

	return o.Pins
}

// GetPinsOk returns a tuple with the Pins field value
// and a boolean to check if the value has been set.
func (o *ListPinsResponse) GetPinsOk() ([]PinDocument, bool) {
	if o == nil {
		return nil, false
	}
	return o.Pins, true
}

// SetPins sets field value
func (o *ListPinsResponse) SetPins(v []PinDocument) {
	o.Pins = v
}

func (o ListPinsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListPinsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pins"] = o.Pins
	return toSerialize, nil
}

type NullableListPinsResponse struct {
	value *ListPinsResponse
	isSet bool
}

func (v NullableListPinsResponse) Get() *ListPinsResponse {
	return v.value
}

func (v *NullableListPinsResponse) Set(val *ListPinsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListPinsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListPinsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListPinsResponse(val *ListPinsResponse) *NullableListPinsResponse {
	return &NullableListPinsResponse{value: val, isSet: true}
}

func (v NullableListPinsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListPinsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


