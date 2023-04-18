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

// checks if the Person type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Person{}

// Person struct for Person
type Person struct {
	// The display name.
	Name string `json:"name"`
	// An opaque identifier that can be used to request metadata for a Person.
	ObfuscatedId string `json:"obfuscatedId"`
	// A list of documents related to this person.
	RelatedDocuments []RelatedDocuments `json:"relatedDocuments,omitempty"`
	Metadata *PersonMetadata `json:"metadata,omitempty"`
}

// NewPerson instantiates a new Person object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPerson(name string, obfuscatedId string) *Person {
	this := Person{}
	this.Name = name
	this.ObfuscatedId = obfuscatedId
	return &this
}

// NewPersonWithDefaults instantiates a new Person object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPersonWithDefaults() *Person {
	this := Person{}
	return &this
}

// GetName returns the Name field value
func (o *Person) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Person) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Person) SetName(v string) {
	o.Name = v
}

// GetObfuscatedId returns the ObfuscatedId field value
func (o *Person) GetObfuscatedId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObfuscatedId
}

// GetObfuscatedIdOk returns a tuple with the ObfuscatedId field value
// and a boolean to check if the value has been set.
func (o *Person) GetObfuscatedIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObfuscatedId, true
}

// SetObfuscatedId sets field value
func (o *Person) SetObfuscatedId(v string) {
	o.ObfuscatedId = v
}

// GetRelatedDocuments returns the RelatedDocuments field value if set, zero value otherwise.
func (o *Person) GetRelatedDocuments() []RelatedDocuments {
	if o == nil || IsNil(o.RelatedDocuments) {
		var ret []RelatedDocuments
		return ret
	}
	return o.RelatedDocuments
}

// GetRelatedDocumentsOk returns a tuple with the RelatedDocuments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Person) GetRelatedDocumentsOk() ([]RelatedDocuments, bool) {
	if o == nil || IsNil(o.RelatedDocuments) {
		return nil, false
	}
	return o.RelatedDocuments, true
}

// HasRelatedDocuments returns a boolean if a field has been set.
func (o *Person) HasRelatedDocuments() bool {
	if o != nil && !IsNil(o.RelatedDocuments) {
		return true
	}

	return false
}

// SetRelatedDocuments gets a reference to the given []RelatedDocuments and assigns it to the RelatedDocuments field.
func (o *Person) SetRelatedDocuments(v []RelatedDocuments) {
	o.RelatedDocuments = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *Person) GetMetadata() PersonMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret PersonMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Person) GetMetadataOk() (*PersonMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Person) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given PersonMetadata and assigns it to the Metadata field.
func (o *Person) SetMetadata(v PersonMetadata) {
	o.Metadata = &v
}

func (o Person) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Person) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["obfuscatedId"] = o.ObfuscatedId
	if !IsNil(o.RelatedDocuments) {
		toSerialize["relatedDocuments"] = o.RelatedDocuments
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

type NullablePerson struct {
	value *Person
	isSet bool
}

func (v NullablePerson) Get() *Person {
	return v.value
}

func (v *NullablePerson) Set(val *Person) {
	v.value = val
	v.isSet = true
}

func (v NullablePerson) IsSet() bool {
	return v.isSet
}

func (v *NullablePerson) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePerson(val *Person) *NullablePerson {
	return &NullablePerson{value: val, isSet: true}
}

func (v NullablePerson) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePerson) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

