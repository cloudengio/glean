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

// checks if the StructuredText type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StructuredText{}

// StructuredText struct for StructuredText
type StructuredText struct {
	Text string `json:"text"`
	// An array of objects each of which contains either a string or a link which optionally corresponds to a document.
	StructuredList []StructuredTextItem `json:"structuredList,omitempty"`
}

// NewStructuredText instantiates a new StructuredText object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStructuredText(text string) *StructuredText {
	this := StructuredText{}
	this.Text = text
	return &this
}

// NewStructuredTextWithDefaults instantiates a new StructuredText object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStructuredTextWithDefaults() *StructuredText {
	this := StructuredText{}
	return &this
}

// GetText returns the Text field value
func (o *StructuredText) GetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *StructuredText) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value
func (o *StructuredText) SetText(v string) {
	o.Text = v
}

// GetStructuredList returns the StructuredList field value if set, zero value otherwise.
func (o *StructuredText) GetStructuredList() []StructuredTextItem {
	if o == nil || IsNil(o.StructuredList) {
		var ret []StructuredTextItem
		return ret
	}
	return o.StructuredList
}

// GetStructuredListOk returns a tuple with the StructuredList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuredText) GetStructuredListOk() ([]StructuredTextItem, bool) {
	if o == nil || IsNil(o.StructuredList) {
		return nil, false
	}
	return o.StructuredList, true
}

// HasStructuredList returns a boolean if a field has been set.
func (o *StructuredText) HasStructuredList() bool {
	if o != nil && !IsNil(o.StructuredList) {
		return true
	}

	return false
}

// SetStructuredList gets a reference to the given []StructuredTextItem and assigns it to the StructuredList field.
func (o *StructuredText) SetStructuredList(v []StructuredTextItem) {
	o.StructuredList = v
}

func (o StructuredText) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StructuredText) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["text"] = o.Text
	if !IsNil(o.StructuredList) {
		toSerialize["structuredList"] = o.StructuredList
	}
	return toSerialize, nil
}

type NullableStructuredText struct {
	value *StructuredText
	isSet bool
}

func (v NullableStructuredText) Get() *StructuredText {
	return v.value
}

func (v *NullableStructuredText) Set(val *StructuredText) {
	v.value = val
	v.isSet = true
}

func (v NullableStructuredText) IsSet() bool {
	return v.isSet
}

func (v *NullableStructuredText) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStructuredText(val *StructuredText) *NullableStructuredText {
	return &NullableStructuredText{value: val, isSet: true}
}

func (v NullableStructuredText) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStructuredText) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


