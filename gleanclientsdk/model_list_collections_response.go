/*
Glean Client API

# Introduction These are the public APIs to enable implementing a custom client interface to the Glean system.  # Usage guidelines This API is evolving fast. Glean will provide advance notice of any planned backwards incompatible changes along with a 6-month sunset period for anything that requires developers to adopt the new versions. 

API version: 0.9.0
Contact: support@glean.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gleanclientsdk

import (
	"encoding/json"
)

// checks if the ListCollectionsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListCollectionsResponse{}

// ListCollectionsResponse struct for ListCollectionsResponse
type ListCollectionsResponse struct {
	// List of all collections, no collection items are fetched.
	Collections []Collection `json:"collections"`
}

// NewListCollectionsResponse instantiates a new ListCollectionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListCollectionsResponse(collections []Collection) *ListCollectionsResponse {
	this := ListCollectionsResponse{}
	this.Collections = collections
	return &this
}

// NewListCollectionsResponseWithDefaults instantiates a new ListCollectionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListCollectionsResponseWithDefaults() *ListCollectionsResponse {
	this := ListCollectionsResponse{}
	return &this
}

// GetCollections returns the Collections field value
func (o *ListCollectionsResponse) GetCollections() []Collection {
	if o == nil {
		var ret []Collection
		return ret
	}

	return o.Collections
}

// GetCollectionsOk returns a tuple with the Collections field value
// and a boolean to check if the value has been set.
func (o *ListCollectionsResponse) GetCollectionsOk() ([]Collection, bool) {
	if o == nil {
		return nil, false
	}
	return o.Collections, true
}

// SetCollections sets field value
func (o *ListCollectionsResponse) SetCollections(v []Collection) {
	o.Collections = v
}

func (o ListCollectionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListCollectionsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["collections"] = o.Collections
	return toSerialize, nil
}

type NullableListCollectionsResponse struct {
	value *ListCollectionsResponse
	isSet bool
}

func (v NullableListCollectionsResponse) Get() *ListCollectionsResponse {
	return v.value
}

func (v *NullableListCollectionsResponse) Set(val *ListCollectionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListCollectionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListCollectionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListCollectionsResponse(val *ListCollectionsResponse) *NullableListCollectionsResponse {
	return &NullableListCollectionsResponse{value: val, isSet: true}
}

func (v NullableListCollectionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListCollectionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


