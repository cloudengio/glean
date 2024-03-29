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

// checks if the ContentDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContentDefinition{}

// ContentDefinition Describes text content or base64 encoded binary content
type ContentDefinition struct {
	MimeType string `json:"mimeType"`
	// text content. Only one of textContent or binary content can be specified
	TextContent *string `json:"textContent,omitempty"`
	// base64 encoded binary content. Only one of textContent or binary content can be specified
	BinaryContent *string `json:"binaryContent,omitempty"`
}

// NewContentDefinition instantiates a new ContentDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentDefinition(mimeType string) *ContentDefinition {
	this := ContentDefinition{}
	this.MimeType = mimeType
	return &this
}

// NewContentDefinitionWithDefaults instantiates a new ContentDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentDefinitionWithDefaults() *ContentDefinition {
	this := ContentDefinition{}
	return &this
}

// GetMimeType returns the MimeType field value
func (o *ContentDefinition) GetMimeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MimeType
}

// GetMimeTypeOk returns a tuple with the MimeType field value
// and a boolean to check if the value has been set.
func (o *ContentDefinition) GetMimeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MimeType, true
}

// SetMimeType sets field value
func (o *ContentDefinition) SetMimeType(v string) {
	o.MimeType = v
}

// GetTextContent returns the TextContent field value if set, zero value otherwise.
func (o *ContentDefinition) GetTextContent() string {
	if o == nil || IsNil(o.TextContent) {
		var ret string
		return ret
	}
	return *o.TextContent
}

// GetTextContentOk returns a tuple with the TextContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentDefinition) GetTextContentOk() (*string, bool) {
	if o == nil || IsNil(o.TextContent) {
		return nil, false
	}
	return o.TextContent, true
}

// HasTextContent returns a boolean if a field has been set.
func (o *ContentDefinition) HasTextContent() bool {
	if o != nil && !IsNil(o.TextContent) {
		return true
	}

	return false
}

// SetTextContent gets a reference to the given string and assigns it to the TextContent field.
func (o *ContentDefinition) SetTextContent(v string) {
	o.TextContent = &v
}

// GetBinaryContent returns the BinaryContent field value if set, zero value otherwise.
func (o *ContentDefinition) GetBinaryContent() string {
	if o == nil || IsNil(o.BinaryContent) {
		var ret string
		return ret
	}
	return *o.BinaryContent
}

// GetBinaryContentOk returns a tuple with the BinaryContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentDefinition) GetBinaryContentOk() (*string, bool) {
	if o == nil || IsNil(o.BinaryContent) {
		return nil, false
	}
	return o.BinaryContent, true
}

// HasBinaryContent returns a boolean if a field has been set.
func (o *ContentDefinition) HasBinaryContent() bool {
	if o != nil && !IsNil(o.BinaryContent) {
		return true
	}

	return false
}

// SetBinaryContent gets a reference to the given string and assigns it to the BinaryContent field.
func (o *ContentDefinition) SetBinaryContent(v string) {
	o.BinaryContent = &v
}

func (o ContentDefinition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContentDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mimeType"] = o.MimeType
	if !IsNil(o.TextContent) {
		toSerialize["textContent"] = o.TextContent
	}
	if !IsNil(o.BinaryContent) {
		toSerialize["binaryContent"] = o.BinaryContent
	}
	return toSerialize, nil
}

type NullableContentDefinition struct {
	value *ContentDefinition
	isSet bool
}

func (v NullableContentDefinition) Get() *ContentDefinition {
	return v.value
}

func (v *NullableContentDefinition) Set(val *ContentDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableContentDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableContentDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentDefinition(val *ContentDefinition) *NullableContentDefinition {
	return &NullableContentDefinition{value: val, isSet: true}
}

func (v NullableContentDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


