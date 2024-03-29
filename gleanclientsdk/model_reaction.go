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

// checks if the Reaction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Reaction{}

// Reaction struct for Reaction
type Reaction struct {
	Type *string `json:"type,omitempty"`
	// The count of the reaction type on the document.
	Count *int32 `json:"count,omitempty"`
	Reactors []Person `json:"reactors,omitempty"`
	// Whether the user in context reacted with this type to the document.
	ReactedByViewer *bool `json:"reactedByViewer,omitempty"`
}

// NewReaction instantiates a new Reaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReaction() *Reaction {
	this := Reaction{}
	return &this
}

// NewReactionWithDefaults instantiates a new Reaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReactionWithDefaults() *Reaction {
	this := Reaction{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Reaction) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Reaction) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Reaction) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Reaction) SetType(v string) {
	o.Type = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *Reaction) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Reaction) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *Reaction) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *Reaction) SetCount(v int32) {
	o.Count = &v
}

// GetReactors returns the Reactors field value if set, zero value otherwise.
func (o *Reaction) GetReactors() []Person {
	if o == nil || IsNil(o.Reactors) {
		var ret []Person
		return ret
	}
	return o.Reactors
}

// GetReactorsOk returns a tuple with the Reactors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Reaction) GetReactorsOk() ([]Person, bool) {
	if o == nil || IsNil(o.Reactors) {
		return nil, false
	}
	return o.Reactors, true
}

// HasReactors returns a boolean if a field has been set.
func (o *Reaction) HasReactors() bool {
	if o != nil && !IsNil(o.Reactors) {
		return true
	}

	return false
}

// SetReactors gets a reference to the given []Person and assigns it to the Reactors field.
func (o *Reaction) SetReactors(v []Person) {
	o.Reactors = v
}

// GetReactedByViewer returns the ReactedByViewer field value if set, zero value otherwise.
func (o *Reaction) GetReactedByViewer() bool {
	if o == nil || IsNil(o.ReactedByViewer) {
		var ret bool
		return ret
	}
	return *o.ReactedByViewer
}

// GetReactedByViewerOk returns a tuple with the ReactedByViewer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Reaction) GetReactedByViewerOk() (*bool, bool) {
	if o == nil || IsNil(o.ReactedByViewer) {
		return nil, false
	}
	return o.ReactedByViewer, true
}

// HasReactedByViewer returns a boolean if a field has been set.
func (o *Reaction) HasReactedByViewer() bool {
	if o != nil && !IsNil(o.ReactedByViewer) {
		return true
	}

	return false
}

// SetReactedByViewer gets a reference to the given bool and assigns it to the ReactedByViewer field.
func (o *Reaction) SetReactedByViewer(v bool) {
	o.ReactedByViewer = &v
}

func (o Reaction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Reaction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Reactors) {
		toSerialize["reactors"] = o.Reactors
	}
	if !IsNil(o.ReactedByViewer) {
		toSerialize["reactedByViewer"] = o.ReactedByViewer
	}
	return toSerialize, nil
}

type NullableReaction struct {
	value *Reaction
	isSet bool
}

func (v NullableReaction) Get() *Reaction {
	return v.value
}

func (v *NullableReaction) Set(val *Reaction) {
	v.value = val
	v.isSet = true
}

func (v NullableReaction) IsSet() bool {
	return v.isSet
}

func (v *NullableReaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReaction(val *Reaction) *NullableReaction {
	return &NullableReaction{value: val, isSet: true}
}

func (v NullableReaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


