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

// checks if the TextRange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TextRange{}

// TextRange A subsection of a given string to which some special formatting should be applied.
type TextRange struct {
	// The inclusive start index of the range.
	StartIndex int32 `json:"startIndex"`
	// The exclusive end index of the range.
	EndIndex *int32 `json:"endIndex,omitempty"`
	Type *string `json:"type,omitempty"`
	Document *Document `json:"document,omitempty"`
}

// NewTextRange instantiates a new TextRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTextRange(startIndex int32) *TextRange {
	this := TextRange{}
	this.StartIndex = startIndex
	return &this
}

// NewTextRangeWithDefaults instantiates a new TextRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTextRangeWithDefaults() *TextRange {
	this := TextRange{}
	return &this
}

// GetStartIndex returns the StartIndex field value
func (o *TextRange) GetStartIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.StartIndex
}

// GetStartIndexOk returns a tuple with the StartIndex field value
// and a boolean to check if the value has been set.
func (o *TextRange) GetStartIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartIndex, true
}

// SetStartIndex sets field value
func (o *TextRange) SetStartIndex(v int32) {
	o.StartIndex = v
}

// GetEndIndex returns the EndIndex field value if set, zero value otherwise.
func (o *TextRange) GetEndIndex() int32 {
	if o == nil || IsNil(o.EndIndex) {
		var ret int32
		return ret
	}
	return *o.EndIndex
}

// GetEndIndexOk returns a tuple with the EndIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TextRange) GetEndIndexOk() (*int32, bool) {
	if o == nil || IsNil(o.EndIndex) {
		return nil, false
	}
	return o.EndIndex, true
}

// HasEndIndex returns a boolean if a field has been set.
func (o *TextRange) HasEndIndex() bool {
	if o != nil && !IsNil(o.EndIndex) {
		return true
	}

	return false
}

// SetEndIndex gets a reference to the given int32 and assigns it to the EndIndex field.
func (o *TextRange) SetEndIndex(v int32) {
	o.EndIndex = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TextRange) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TextRange) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TextRange) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TextRange) SetType(v string) {
	o.Type = &v
}

// GetDocument returns the Document field value if set, zero value otherwise.
func (o *TextRange) GetDocument() Document {
	if o == nil || IsNil(o.Document) {
		var ret Document
		return ret
	}
	return *o.Document
}

// GetDocumentOk returns a tuple with the Document field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TextRange) GetDocumentOk() (*Document, bool) {
	if o == nil || IsNil(o.Document) {
		return nil, false
	}
	return o.Document, true
}

// HasDocument returns a boolean if a field has been set.
func (o *TextRange) HasDocument() bool {
	if o != nil && !IsNil(o.Document) {
		return true
	}

	return false
}

// SetDocument gets a reference to the given Document and assigns it to the Document field.
func (o *TextRange) SetDocument(v Document) {
	o.Document = &v
}

func (o TextRange) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TextRange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["startIndex"] = o.StartIndex
	if !IsNil(o.EndIndex) {
		toSerialize["endIndex"] = o.EndIndex
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Document) {
		toSerialize["document"] = o.Document
	}
	return toSerialize, nil
}

type NullableTextRange struct {
	value *TextRange
	isSet bool
}

func (v NullableTextRange) Get() *TextRange {
	return v.value
}

func (v *NullableTextRange) Set(val *TextRange) {
	v.value = val
	v.isSet = true
}

func (v NullableTextRange) IsSet() bool {
	return v.isSet
}

func (v *NullableTextRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTextRange(val *TextRange) *NullableTextRange {
	return &NullableTextRange{value: val, isSet: true}
}

func (v NullableTextRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTextRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

