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

// checks if the CreateCollectionResponseAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateCollectionResponseAllOf{}

// CreateCollectionResponseAllOf struct for CreateCollectionResponseAllOf
type CreateCollectionResponseAllOf struct {
	Collection *Collection `json:"collection,omitempty"`
	Error *CollectionError `json:"error,omitempty"`
}

// NewCreateCollectionResponseAllOf instantiates a new CreateCollectionResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCollectionResponseAllOf() *CreateCollectionResponseAllOf {
	this := CreateCollectionResponseAllOf{}
	return &this
}

// NewCreateCollectionResponseAllOfWithDefaults instantiates a new CreateCollectionResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCollectionResponseAllOfWithDefaults() *CreateCollectionResponseAllOf {
	this := CreateCollectionResponseAllOf{}
	return &this
}

// GetCollection returns the Collection field value if set, zero value otherwise.
func (o *CreateCollectionResponseAllOf) GetCollection() Collection {
	if o == nil || IsNil(o.Collection) {
		var ret Collection
		return ret
	}
	return *o.Collection
}

// GetCollectionOk returns a tuple with the Collection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCollectionResponseAllOf) GetCollectionOk() (*Collection, bool) {
	if o == nil || IsNil(o.Collection) {
		return nil, false
	}
	return o.Collection, true
}

// HasCollection returns a boolean if a field has been set.
func (o *CreateCollectionResponseAllOf) HasCollection() bool {
	if o != nil && !IsNil(o.Collection) {
		return true
	}

	return false
}

// SetCollection gets a reference to the given Collection and assigns it to the Collection field.
func (o *CreateCollectionResponseAllOf) SetCollection(v Collection) {
	o.Collection = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *CreateCollectionResponseAllOf) GetError() CollectionError {
	if o == nil || IsNil(o.Error) {
		var ret CollectionError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCollectionResponseAllOf) GetErrorOk() (*CollectionError, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *CreateCollectionResponseAllOf) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given CollectionError and assigns it to the Error field.
func (o *CreateCollectionResponseAllOf) SetError(v CollectionError) {
	o.Error = &v
}

func (o CreateCollectionResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateCollectionResponseAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Collection) {
		toSerialize["collection"] = o.Collection
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableCreateCollectionResponseAllOf struct {
	value *CreateCollectionResponseAllOf
	isSet bool
}

func (v NullableCreateCollectionResponseAllOf) Get() *CreateCollectionResponseAllOf {
	return v.value
}

func (v *NullableCreateCollectionResponseAllOf) Set(val *CreateCollectionResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCollectionResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCollectionResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCollectionResponseAllOf(val *CreateCollectionResponseAllOf) *NullableCreateCollectionResponseAllOf {
	return &NullableCreateCollectionResponseAllOf{value: val, isSet: true}
}

func (v NullableCreateCollectionResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCollectionResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


