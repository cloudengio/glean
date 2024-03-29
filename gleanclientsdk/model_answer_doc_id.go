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

// checks if the AnswerDocId type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnswerDocId{}

// AnswerDocId struct for AnswerDocId
type AnswerDocId struct {
	// Internal document id of the answer. We support using the document id for cases where the client doesn't have the integer id available. If both are available, using the integer id is preferred.
	DocId *string `json:"docId,omitempty"`
}

// NewAnswerDocId instantiates a new AnswerDocId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnswerDocId() *AnswerDocId {
	this := AnswerDocId{}
	return &this
}

// NewAnswerDocIdWithDefaults instantiates a new AnswerDocId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnswerDocIdWithDefaults() *AnswerDocId {
	this := AnswerDocId{}
	return &this
}

// GetDocId returns the DocId field value if set, zero value otherwise.
func (o *AnswerDocId) GetDocId() string {
	if o == nil || IsNil(o.DocId) {
		var ret string
		return ret
	}
	return *o.DocId
}

// GetDocIdOk returns a tuple with the DocId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnswerDocId) GetDocIdOk() (*string, bool) {
	if o == nil || IsNil(o.DocId) {
		return nil, false
	}
	return o.DocId, true
}

// HasDocId returns a boolean if a field has been set.
func (o *AnswerDocId) HasDocId() bool {
	if o != nil && !IsNil(o.DocId) {
		return true
	}

	return false
}

// SetDocId gets a reference to the given string and assigns it to the DocId field.
func (o *AnswerDocId) SetDocId(v string) {
	o.DocId = &v
}

func (o AnswerDocId) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnswerDocId) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DocId) {
		toSerialize["docId"] = o.DocId
	}
	return toSerialize, nil
}

type NullableAnswerDocId struct {
	value *AnswerDocId
	isSet bool
}

func (v NullableAnswerDocId) Get() *AnswerDocId {
	return v.value
}

func (v *NullableAnswerDocId) Set(val *AnswerDocId) {
	v.value = val
	v.isSet = true
}

func (v NullableAnswerDocId) IsSet() bool {
	return v.isSet
}

func (v *NullableAnswerDocId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnswerDocId(val *AnswerDocId) *NullableAnswerDocId {
	return &NullableAnswerDocId{value: val, isSet: true}
}

func (v NullableAnswerDocId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnswerDocId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


