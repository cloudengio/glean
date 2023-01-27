/*
Glean Indexing API

# Introduction In addition to the data sources that Glean has built-in support for, Glean also provides a REST API that enables customers to put arbitrary content in the search index. This is useful, for example, for doing permissions-aware search over content in internal tools that reside on-prem as well as for searching over applications that Glean does not currently support first class. In addition these APIs allow the customer to push organization data (people info, organization structure etc) into Glean.  # Early Access Please note that we are continually evolving the system so these APIs are considered “early access” and are subject to change. 

API version: 0.9.0
Contact: support@glean.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gleansdk

import (
	"encoding/json"
)

// GetDocumentCountResponse Describes the response body of the /getdocumentcount API call
type GetDocumentCountResponse struct {
	// Number of documents corresponding to the specified custom datasource.
	DocumentCount *int32 `json:"documentCount,omitempty"`
}

// NewGetDocumentCountResponse instantiates a new GetDocumentCountResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDocumentCountResponse() *GetDocumentCountResponse {
	this := GetDocumentCountResponse{}
	return &this
}

// NewGetDocumentCountResponseWithDefaults instantiates a new GetDocumentCountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDocumentCountResponseWithDefaults() *GetDocumentCountResponse {
	this := GetDocumentCountResponse{}
	return &this
}

// GetDocumentCount returns the DocumentCount field value if set, zero value otherwise.
func (o *GetDocumentCountResponse) GetDocumentCount() int32 {
	if o == nil || o.DocumentCount == nil {
		var ret int32
		return ret
	}
	return *o.DocumentCount
}

// GetDocumentCountOk returns a tuple with the DocumentCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDocumentCountResponse) GetDocumentCountOk() (*int32, bool) {
	if o == nil || o.DocumentCount == nil {
		return nil, false
	}
	return o.DocumentCount, true
}

// HasDocumentCount returns a boolean if a field has been set.
func (o *GetDocumentCountResponse) HasDocumentCount() bool {
	if o != nil && o.DocumentCount != nil {
		return true
	}

	return false
}

// SetDocumentCount gets a reference to the given int32 and assigns it to the DocumentCount field.
func (o *GetDocumentCountResponse) SetDocumentCount(v int32) {
	o.DocumentCount = &v
}

func (o GetDocumentCountResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DocumentCount != nil {
		toSerialize["documentCount"] = o.DocumentCount
	}
	return json.Marshal(toSerialize)
}

type NullableGetDocumentCountResponse struct {
	value *GetDocumentCountResponse
	isSet bool
}

func (v NullableGetDocumentCountResponse) Get() *GetDocumentCountResponse {
	return v.value
}

func (v *NullableGetDocumentCountResponse) Set(val *GetDocumentCountResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDocumentCountResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDocumentCountResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDocumentCountResponse(val *GetDocumentCountResponse) *NullableGetDocumentCountResponse {
	return &NullableGetDocumentCountResponse{value: val, isSet: true}
}

func (v NullableGetDocumentCountResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDocumentCountResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


