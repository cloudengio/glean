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

// checks if the DeleteAnswerRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteAnswerRequest{}

// DeleteAnswerRequest struct for DeleteAnswerRequest
type DeleteAnswerRequest struct {
	// The opaque id of the answer.
	Id int32 `json:"id"`
	// Internal document id of the answer. We support using the document id for cases where the client doesn't have the integer id available. If both are available, using the integer id is preferred.
	DocId *string `json:"docId,omitempty"`
}

// NewDeleteAnswerRequest instantiates a new DeleteAnswerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteAnswerRequest(id int32) *DeleteAnswerRequest {
	this := DeleteAnswerRequest{}
	this.Id = id
	return &this
}

// NewDeleteAnswerRequestWithDefaults instantiates a new DeleteAnswerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteAnswerRequestWithDefaults() *DeleteAnswerRequest {
	this := DeleteAnswerRequest{}
	return &this
}

// GetId returns the Id field value
func (o *DeleteAnswerRequest) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DeleteAnswerRequest) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DeleteAnswerRequest) SetId(v int32) {
	o.Id = v
}

// GetDocId returns the DocId field value if set, zero value otherwise.
func (o *DeleteAnswerRequest) GetDocId() string {
	if o == nil || IsNil(o.DocId) {
		var ret string
		return ret
	}
	return *o.DocId
}

// GetDocIdOk returns a tuple with the DocId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteAnswerRequest) GetDocIdOk() (*string, bool) {
	if o == nil || IsNil(o.DocId) {
		return nil, false
	}
	return o.DocId, true
}

// HasDocId returns a boolean if a field has been set.
func (o *DeleteAnswerRequest) HasDocId() bool {
	if o != nil && !IsNil(o.DocId) {
		return true
	}

	return false
}

// SetDocId gets a reference to the given string and assigns it to the DocId field.
func (o *DeleteAnswerRequest) SetDocId(v string) {
	o.DocId = &v
}

func (o DeleteAnswerRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteAnswerRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.DocId) {
		toSerialize["docId"] = o.DocId
	}
	return toSerialize, nil
}

type NullableDeleteAnswerRequest struct {
	value *DeleteAnswerRequest
	isSet bool
}

func (v NullableDeleteAnswerRequest) Get() *DeleteAnswerRequest {
	return v.value
}

func (v *NullableDeleteAnswerRequest) Set(val *DeleteAnswerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteAnswerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteAnswerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteAnswerRequest(val *DeleteAnswerRequest) *NullableDeleteAnswerRequest {
	return &NullableDeleteAnswerRequest{value: val, isSet: true}
}

func (v NullableDeleteAnswerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteAnswerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


