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

// checks if the GetDocumentsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDocumentsResponse{}

// GetDocumentsResponse struct for GetDocumentsResponse
type GetDocumentsResponse struct {
	// The document details or the error if document is not found.
	Documents *map[string]DocumentOrError `json:"documents,omitempty"`
}

// NewGetDocumentsResponse instantiates a new GetDocumentsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDocumentsResponse() *GetDocumentsResponse {
	this := GetDocumentsResponse{}
	return &this
}

// NewGetDocumentsResponseWithDefaults instantiates a new GetDocumentsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDocumentsResponseWithDefaults() *GetDocumentsResponse {
	this := GetDocumentsResponse{}
	return &this
}

// GetDocuments returns the Documents field value if set, zero value otherwise.
func (o *GetDocumentsResponse) GetDocuments() map[string]DocumentOrError {
	if o == nil || IsNil(o.Documents) {
		var ret map[string]DocumentOrError
		return ret
	}
	return *o.Documents
}

// GetDocumentsOk returns a tuple with the Documents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDocumentsResponse) GetDocumentsOk() (*map[string]DocumentOrError, bool) {
	if o == nil || IsNil(o.Documents) {
		return nil, false
	}
	return o.Documents, true
}

// HasDocuments returns a boolean if a field has been set.
func (o *GetDocumentsResponse) HasDocuments() bool {
	if o != nil && !IsNil(o.Documents) {
		return true
	}

	return false
}

// SetDocuments gets a reference to the given map[string]DocumentOrError and assigns it to the Documents field.
func (o *GetDocumentsResponse) SetDocuments(v map[string]DocumentOrError) {
	o.Documents = &v
}

func (o GetDocumentsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDocumentsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Documents) {
		toSerialize["documents"] = o.Documents
	}
	return toSerialize, nil
}

type NullableGetDocumentsResponse struct {
	value *GetDocumentsResponse
	isSet bool
}

func (v NullableGetDocumentsResponse) Get() *GetDocumentsResponse {
	return v.value
}

func (v *NullableGetDocumentsResponse) Set(val *GetDocumentsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDocumentsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDocumentsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDocumentsResponse(val *GetDocumentsResponse) *NullableGetDocumentsResponse {
	return &NullableGetDocumentsResponse{value: val, isSet: true}
}

func (v NullableGetDocumentsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDocumentsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

