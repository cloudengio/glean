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

// checks if the DisplayableListSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DisplayableListSource{}

// DisplayableListSource struct for DisplayableListSource
type DisplayableListSource struct {
	// The type of data that backs this displayable list
	Source *string `json:"source,omitempty"`
}

// NewDisplayableListSource instantiates a new DisplayableListSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDisplayableListSource() *DisplayableListSource {
	this := DisplayableListSource{}
	return &this
}

// NewDisplayableListSourceWithDefaults instantiates a new DisplayableListSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDisplayableListSourceWithDefaults() *DisplayableListSource {
	this := DisplayableListSource{}
	return &this
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *DisplayableListSource) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DisplayableListSource) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *DisplayableListSource) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *DisplayableListSource) SetSource(v string) {
	o.Source = &v
}

func (o DisplayableListSource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DisplayableListSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	return toSerialize, nil
}

type NullableDisplayableListSource struct {
	value *DisplayableListSource
	isSet bool
}

func (v NullableDisplayableListSource) Get() *DisplayableListSource {
	return v.value
}

func (v *NullableDisplayableListSource) Set(val *DisplayableListSource) {
	v.value = val
	v.isSet = true
}

func (v NullableDisplayableListSource) IsSet() bool {
	return v.isSet
}

func (v *NullableDisplayableListSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDisplayableListSource(val *DisplayableListSource) *NullableDisplayableListSource {
	return &NullableDisplayableListSource{value: val, isSet: true}
}

func (v NullableDisplayableListSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDisplayableListSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


