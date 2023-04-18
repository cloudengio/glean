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

// checks if the DocumentOrErrorOneOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocumentOrErrorOneOf{}

// DocumentOrErrorOneOf struct for DocumentOrErrorOneOf
type DocumentOrErrorOneOf struct {
	// The text for error, reason.
	Error string `json:"error"`
}

// NewDocumentOrErrorOneOf instantiates a new DocumentOrErrorOneOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentOrErrorOneOf(error_ string) *DocumentOrErrorOneOf {
	this := DocumentOrErrorOneOf{}
	this.Error = error_
	return &this
}

// NewDocumentOrErrorOneOfWithDefaults instantiates a new DocumentOrErrorOneOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentOrErrorOneOfWithDefaults() *DocumentOrErrorOneOf {
	this := DocumentOrErrorOneOf{}
	return &this
}

// GetError returns the Error field value
func (o *DocumentOrErrorOneOf) GetError() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *DocumentOrErrorOneOf) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *DocumentOrErrorOneOf) SetError(v string) {
	o.Error = v
}

func (o DocumentOrErrorOneOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocumentOrErrorOneOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["error"] = o.Error
	return toSerialize, nil
}

type NullableDocumentOrErrorOneOf struct {
	value *DocumentOrErrorOneOf
	isSet bool
}

func (v NullableDocumentOrErrorOneOf) Get() *DocumentOrErrorOneOf {
	return v.value
}

func (v *NullableDocumentOrErrorOneOf) Set(val *DocumentOrErrorOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentOrErrorOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentOrErrorOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentOrErrorOneOf(val *DocumentOrErrorOneOf) *NullableDocumentOrErrorOneOf {
	return &NullableDocumentOrErrorOneOf{value: val, isSet: true}
}

func (v NullableDocumentOrErrorOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentOrErrorOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

