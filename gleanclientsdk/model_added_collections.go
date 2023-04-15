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

// checks if the AddedCollections type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddedCollections{}

// AddedCollections struct for AddedCollections
type AddedCollections struct {
	// IDs of collections to which a document is added.
	AddedCollections []int32 `json:"addedCollections,omitempty"`
}

// NewAddedCollections instantiates a new AddedCollections object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddedCollections() *AddedCollections {
	this := AddedCollections{}
	return &this
}

// NewAddedCollectionsWithDefaults instantiates a new AddedCollections object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddedCollectionsWithDefaults() *AddedCollections {
	this := AddedCollections{}
	return &this
}

// GetAddedCollections returns the AddedCollections field value if set, zero value otherwise.
func (o *AddedCollections) GetAddedCollections() []int32 {
	if o == nil || IsNil(o.AddedCollections) {
		var ret []int32
		return ret
	}
	return o.AddedCollections
}

// GetAddedCollectionsOk returns a tuple with the AddedCollections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddedCollections) GetAddedCollectionsOk() ([]int32, bool) {
	if o == nil || IsNil(o.AddedCollections) {
		return nil, false
	}
	return o.AddedCollections, true
}

// HasAddedCollections returns a boolean if a field has been set.
func (o *AddedCollections) HasAddedCollections() bool {
	if o != nil && !IsNil(o.AddedCollections) {
		return true
	}

	return false
}

// SetAddedCollections gets a reference to the given []int32 and assigns it to the AddedCollections field.
func (o *AddedCollections) SetAddedCollections(v []int32) {
	o.AddedCollections = v
}

func (o AddedCollections) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddedCollections) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AddedCollections) {
		toSerialize["addedCollections"] = o.AddedCollections
	}
	return toSerialize, nil
}

type NullableAddedCollections struct {
	value *AddedCollections
	isSet bool
}

func (v NullableAddedCollections) Get() *AddedCollections {
	return v.value
}

func (v *NullableAddedCollections) Set(val *AddedCollections) {
	v.value = val
	v.isSet = true
}

func (v NullableAddedCollections) IsSet() bool {
	return v.isSet
}

func (v *NullableAddedCollections) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddedCollections(val *AddedCollections) *NullableAddedCollections {
	return &NullableAddedCollections{value: val, isSet: true}
}

func (v NullableAddedCollections) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddedCollections) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


