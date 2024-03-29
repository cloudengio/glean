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

// checks if the AnswerLikes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnswerLikes{}

// AnswerLikes struct for AnswerLikes
type AnswerLikes struct {
	LikedBy []AnswerLike `json:"likedBy"`
	// Whether the user in context liked the answer.
	LikedByUser bool `json:"likedByUser"`
	// The total number of likes for the answer.
	NumLikes int32 `json:"numLikes"`
}

// NewAnswerLikes instantiates a new AnswerLikes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnswerLikes(likedBy []AnswerLike, likedByUser bool, numLikes int32) *AnswerLikes {
	this := AnswerLikes{}
	this.LikedBy = likedBy
	this.LikedByUser = likedByUser
	this.NumLikes = numLikes
	return &this
}

// NewAnswerLikesWithDefaults instantiates a new AnswerLikes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnswerLikesWithDefaults() *AnswerLikes {
	this := AnswerLikes{}
	return &this
}

// GetLikedBy returns the LikedBy field value
func (o *AnswerLikes) GetLikedBy() []AnswerLike {
	if o == nil {
		var ret []AnswerLike
		return ret
	}

	return o.LikedBy
}

// GetLikedByOk returns a tuple with the LikedBy field value
// and a boolean to check if the value has been set.
func (o *AnswerLikes) GetLikedByOk() ([]AnswerLike, bool) {
	if o == nil {
		return nil, false
	}
	return o.LikedBy, true
}

// SetLikedBy sets field value
func (o *AnswerLikes) SetLikedBy(v []AnswerLike) {
	o.LikedBy = v
}

// GetLikedByUser returns the LikedByUser field value
func (o *AnswerLikes) GetLikedByUser() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.LikedByUser
}

// GetLikedByUserOk returns a tuple with the LikedByUser field value
// and a boolean to check if the value has been set.
func (o *AnswerLikes) GetLikedByUserOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LikedByUser, true
}

// SetLikedByUser sets field value
func (o *AnswerLikes) SetLikedByUser(v bool) {
	o.LikedByUser = v
}

// GetNumLikes returns the NumLikes field value
func (o *AnswerLikes) GetNumLikes() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumLikes
}

// GetNumLikesOk returns a tuple with the NumLikes field value
// and a boolean to check if the value has been set.
func (o *AnswerLikes) GetNumLikesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumLikes, true
}

// SetNumLikes sets field value
func (o *AnswerLikes) SetNumLikes(v int32) {
	o.NumLikes = v
}

func (o AnswerLikes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnswerLikes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["likedBy"] = o.LikedBy
	toSerialize["likedByUser"] = o.LikedByUser
	toSerialize["numLikes"] = o.NumLikes
	return toSerialize, nil
}

type NullableAnswerLikes struct {
	value *AnswerLikes
	isSet bool
}

func (v NullableAnswerLikes) Get() *AnswerLikes {
	return v.value
}

func (v *NullableAnswerLikes) Set(val *AnswerLikes) {
	v.value = val
	v.isSet = true
}

func (v NullableAnswerLikes) IsSet() bool {
	return v.isSet
}

func (v *NullableAnswerLikes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnswerLikes(val *AnswerLikes) *NullableAnswerLikes {
	return &NullableAnswerLikes{value: val, isSet: true}
}

func (v NullableAnswerLikes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnswerLikes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


