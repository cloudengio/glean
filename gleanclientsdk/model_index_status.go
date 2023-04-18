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
	"time"
)

// checks if the IndexStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IndexStatus{}

// IndexStatus struct for IndexStatus
type IndexStatus struct {
	// When the document was last crawled
	LastCrawledTime *time.Time `json:"lastCrawledTime,omitempty"`
	// When the document was last indexed
	LastIndexedTime *time.Time `json:"lastIndexedTime,omitempty"`
}

// NewIndexStatus instantiates a new IndexStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndexStatus() *IndexStatus {
	this := IndexStatus{}
	return &this
}

// NewIndexStatusWithDefaults instantiates a new IndexStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndexStatusWithDefaults() *IndexStatus {
	this := IndexStatus{}
	return &this
}

// GetLastCrawledTime returns the LastCrawledTime field value if set, zero value otherwise.
func (o *IndexStatus) GetLastCrawledTime() time.Time {
	if o == nil || IsNil(o.LastCrawledTime) {
		var ret time.Time
		return ret
	}
	return *o.LastCrawledTime
}

// GetLastCrawledTimeOk returns a tuple with the LastCrawledTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexStatus) GetLastCrawledTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastCrawledTime) {
		return nil, false
	}
	return o.LastCrawledTime, true
}

// HasLastCrawledTime returns a boolean if a field has been set.
func (o *IndexStatus) HasLastCrawledTime() bool {
	if o != nil && !IsNil(o.LastCrawledTime) {
		return true
	}

	return false
}

// SetLastCrawledTime gets a reference to the given time.Time and assigns it to the LastCrawledTime field.
func (o *IndexStatus) SetLastCrawledTime(v time.Time) {
	o.LastCrawledTime = &v
}

// GetLastIndexedTime returns the LastIndexedTime field value if set, zero value otherwise.
func (o *IndexStatus) GetLastIndexedTime() time.Time {
	if o == nil || IsNil(o.LastIndexedTime) {
		var ret time.Time
		return ret
	}
	return *o.LastIndexedTime
}

// GetLastIndexedTimeOk returns a tuple with the LastIndexedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexStatus) GetLastIndexedTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastIndexedTime) {
		return nil, false
	}
	return o.LastIndexedTime, true
}

// HasLastIndexedTime returns a boolean if a field has been set.
func (o *IndexStatus) HasLastIndexedTime() bool {
	if o != nil && !IsNil(o.LastIndexedTime) {
		return true
	}

	return false
}

// SetLastIndexedTime gets a reference to the given time.Time and assigns it to the LastIndexedTime field.
func (o *IndexStatus) SetLastIndexedTime(v time.Time) {
	o.LastIndexedTime = &v
}

func (o IndexStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IndexStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LastCrawledTime) {
		toSerialize["lastCrawledTime"] = o.LastCrawledTime
	}
	if !IsNil(o.LastIndexedTime) {
		toSerialize["lastIndexedTime"] = o.LastIndexedTime
	}
	return toSerialize, nil
}

type NullableIndexStatus struct {
	value *IndexStatus
	isSet bool
}

func (v NullableIndexStatus) Get() *IndexStatus {
	return v.value
}

func (v *NullableIndexStatus) Set(val *IndexStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableIndexStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableIndexStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndexStatus(val *IndexStatus) *NullableIndexStatus {
	return &NullableIndexStatus{value: val, isSet: true}
}

func (v NullableIndexStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndexStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


