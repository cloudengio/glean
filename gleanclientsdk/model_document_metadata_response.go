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

// checks if the DocumentMetadataResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocumentMetadataResponse{}

// DocumentMetadataResponse struct for DocumentMetadataResponse
type DocumentMetadataResponse struct {
	// List of document metadata requested.
	Documents []DocumentMetadata `json:"documents,omitempty"`
}

// NewDocumentMetadataResponse instantiates a new DocumentMetadataResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentMetadataResponse() *DocumentMetadataResponse {
	this := DocumentMetadataResponse{}
	return &this
}

// NewDocumentMetadataResponseWithDefaults instantiates a new DocumentMetadataResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentMetadataResponseWithDefaults() *DocumentMetadataResponse {
	this := DocumentMetadataResponse{}
	return &this
}

// GetDocuments returns the Documents field value if set, zero value otherwise.
func (o *DocumentMetadataResponse) GetDocuments() []DocumentMetadata {
	if o == nil || IsNil(o.Documents) {
		var ret []DocumentMetadata
		return ret
	}
	return o.Documents
}

// GetDocumentsOk returns a tuple with the Documents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentMetadataResponse) GetDocumentsOk() ([]DocumentMetadata, bool) {
	if o == nil || IsNil(o.Documents) {
		return nil, false
	}
	return o.Documents, true
}

// HasDocuments returns a boolean if a field has been set.
func (o *DocumentMetadataResponse) HasDocuments() bool {
	if o != nil && !IsNil(o.Documents) {
		return true
	}

	return false
}

// SetDocuments gets a reference to the given []DocumentMetadata and assigns it to the Documents field.
func (o *DocumentMetadataResponse) SetDocuments(v []DocumentMetadata) {
	o.Documents = v
}

func (o DocumentMetadataResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocumentMetadataResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Documents) {
		toSerialize["documents"] = o.Documents
	}
	return toSerialize, nil
}

type NullableDocumentMetadataResponse struct {
	value *DocumentMetadataResponse
	isSet bool
}

func (v NullableDocumentMetadataResponse) Get() *DocumentMetadataResponse {
	return v.value
}

func (v *NullableDocumentMetadataResponse) Set(val *DocumentMetadataResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentMetadataResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentMetadataResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentMetadataResponse(val *DocumentMetadataResponse) *NullableDocumentMetadataResponse {
	return &NullableDocumentMetadataResponse{value: val, isSet: true}
}

func (v NullableDocumentMetadataResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentMetadataResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


