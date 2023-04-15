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

// checks if the Unpin type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Unpin{}

// Unpin struct for Unpin
type Unpin struct {
	// DEPRECATED - Prefer use of `id`
	PinId *int32 `json:"pinId,omitempty"`
	// The opaque id of the pin to be unpinned.
	Id *string `json:"id,omitempty"`
}

// NewUnpin instantiates a new Unpin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnpin() *Unpin {
	this := Unpin{}
	return &this
}

// NewUnpinWithDefaults instantiates a new Unpin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnpinWithDefaults() *Unpin {
	this := Unpin{}
	return &this
}

// GetPinId returns the PinId field value if set, zero value otherwise.
func (o *Unpin) GetPinId() int32 {
	if o == nil || IsNil(o.PinId) {
		var ret int32
		return ret
	}
	return *o.PinId
}

// GetPinIdOk returns a tuple with the PinId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Unpin) GetPinIdOk() (*int32, bool) {
	if o == nil || IsNil(o.PinId) {
		return nil, false
	}
	return o.PinId, true
}

// HasPinId returns a boolean if a field has been set.
func (o *Unpin) HasPinId() bool {
	if o != nil && !IsNil(o.PinId) {
		return true
	}

	return false
}

// SetPinId gets a reference to the given int32 and assigns it to the PinId field.
func (o *Unpin) SetPinId(v int32) {
	o.PinId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Unpin) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Unpin) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Unpin) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Unpin) SetId(v string) {
	o.Id = &v
}

func (o Unpin) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Unpin) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PinId) {
		toSerialize["pinId"] = o.PinId
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableUnpin struct {
	value *Unpin
	isSet bool
}

func (v NullableUnpin) Get() *Unpin {
	return v.value
}

func (v *NullableUnpin) Set(val *Unpin) {
	v.value = val
	v.isSet = true
}

func (v NullableUnpin) IsSet() bool {
	return v.isSet
}

func (v *NullableUnpin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnpin(val *Unpin) *NullableUnpin {
	return &NullableUnpin{value: val, isSet: true}
}

func (v NullableUnpin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnpin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


