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

// checks if the DocumentSpecOneOf2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocumentSpecOneOf2{}

// DocumentSpecOneOf2 struct for DocumentSpecOneOf2
type DocumentSpecOneOf2 struct {
	// The type of the user generated content (UGC datasource).
	UgcType string `json:"ugcType"`
	// The id for user generated content.
	ContentId int32 `json:"contentId"`
	// The specific type of the user generated content type.
	DocType *string `json:"docType,omitempty"`
}

// NewDocumentSpecOneOf2 instantiates a new DocumentSpecOneOf2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentSpecOneOf2(ugcType string, contentId int32) *DocumentSpecOneOf2 {
	this := DocumentSpecOneOf2{}
	this.UgcType = ugcType
	this.ContentId = contentId
	return &this
}

// NewDocumentSpecOneOf2WithDefaults instantiates a new DocumentSpecOneOf2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentSpecOneOf2WithDefaults() *DocumentSpecOneOf2 {
	this := DocumentSpecOneOf2{}
	return &this
}

// GetUgcType returns the UgcType field value
func (o *DocumentSpecOneOf2) GetUgcType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UgcType
}

// GetUgcTypeOk returns a tuple with the UgcType field value
// and a boolean to check if the value has been set.
func (o *DocumentSpecOneOf2) GetUgcTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UgcType, true
}

// SetUgcType sets field value
func (o *DocumentSpecOneOf2) SetUgcType(v string) {
	o.UgcType = v
}

// GetContentId returns the ContentId field value
func (o *DocumentSpecOneOf2) GetContentId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ContentId
}

// GetContentIdOk returns a tuple with the ContentId field value
// and a boolean to check if the value has been set.
func (o *DocumentSpecOneOf2) GetContentIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentId, true
}

// SetContentId sets field value
func (o *DocumentSpecOneOf2) SetContentId(v int32) {
	o.ContentId = v
}

// GetDocType returns the DocType field value if set, zero value otherwise.
func (o *DocumentSpecOneOf2) GetDocType() string {
	if o == nil || IsNil(o.DocType) {
		var ret string
		return ret
	}
	return *o.DocType
}

// GetDocTypeOk returns a tuple with the DocType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentSpecOneOf2) GetDocTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DocType) {
		return nil, false
	}
	return o.DocType, true
}

// HasDocType returns a boolean if a field has been set.
func (o *DocumentSpecOneOf2) HasDocType() bool {
	if o != nil && !IsNil(o.DocType) {
		return true
	}

	return false
}

// SetDocType gets a reference to the given string and assigns it to the DocType field.
func (o *DocumentSpecOneOf2) SetDocType(v string) {
	o.DocType = &v
}

func (o DocumentSpecOneOf2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocumentSpecOneOf2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ugcType"] = o.UgcType
	toSerialize["contentId"] = o.ContentId
	if !IsNil(o.DocType) {
		toSerialize["docType"] = o.DocType
	}
	return toSerialize, nil
}

type NullableDocumentSpecOneOf2 struct {
	value *DocumentSpecOneOf2
	isSet bool
}

func (v NullableDocumentSpecOneOf2) Get() *DocumentSpecOneOf2 {
	return v.value
}

func (v *NullableDocumentSpecOneOf2) Set(val *DocumentSpecOneOf2) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentSpecOneOf2) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentSpecOneOf2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentSpecOneOf2(val *DocumentSpecOneOf2) *NullableDocumentSpecOneOf2 {
	return &NullableDocumentSpecOneOf2{value: val, isSet: true}
}

func (v NullableDocumentSpecOneOf2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentSpecOneOf2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


