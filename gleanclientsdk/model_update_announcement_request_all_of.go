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

// checks if the UpdateAnnouncementRequestAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateAnnouncementRequestAllOf{}

// UpdateAnnouncementRequestAllOf struct for UpdateAnnouncementRequestAllOf
type UpdateAnnouncementRequestAllOf struct {
	// The opaque id of the announcement.
	Id int32 `json:"id"`
	Data *AnnouncementCreateOrUpdateData `json:"data,omitempty"`
}

// NewUpdateAnnouncementRequestAllOf instantiates a new UpdateAnnouncementRequestAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateAnnouncementRequestAllOf(id int32) *UpdateAnnouncementRequestAllOf {
	this := UpdateAnnouncementRequestAllOf{}
	this.Id = id
	return &this
}

// NewUpdateAnnouncementRequestAllOfWithDefaults instantiates a new UpdateAnnouncementRequestAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateAnnouncementRequestAllOfWithDefaults() *UpdateAnnouncementRequestAllOf {
	this := UpdateAnnouncementRequestAllOf{}
	return &this
}

// GetId returns the Id field value
func (o *UpdateAnnouncementRequestAllOf) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UpdateAnnouncementRequestAllOf) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UpdateAnnouncementRequestAllOf) SetId(v int32) {
	o.Id = v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *UpdateAnnouncementRequestAllOf) GetData() AnnouncementCreateOrUpdateData {
	if o == nil || IsNil(o.Data) {
		var ret AnnouncementCreateOrUpdateData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAnnouncementRequestAllOf) GetDataOk() (*AnnouncementCreateOrUpdateData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *UpdateAnnouncementRequestAllOf) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given AnnouncementCreateOrUpdateData and assigns it to the Data field.
func (o *UpdateAnnouncementRequestAllOf) SetData(v AnnouncementCreateOrUpdateData) {
	o.Data = &v
}

func (o UpdateAnnouncementRequestAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateAnnouncementRequestAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableUpdateAnnouncementRequestAllOf struct {
	value *UpdateAnnouncementRequestAllOf
	isSet bool
}

func (v NullableUpdateAnnouncementRequestAllOf) Get() *UpdateAnnouncementRequestAllOf {
	return v.value
}

func (v *NullableUpdateAnnouncementRequestAllOf) Set(val *UpdateAnnouncementRequestAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateAnnouncementRequestAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateAnnouncementRequestAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateAnnouncementRequestAllOf(val *UpdateAnnouncementRequestAllOf) *NullableUpdateAnnouncementRequestAllOf {
	return &NullableUpdateAnnouncementRequestAllOf{value: val, isSet: true}
}

func (v NullableUpdateAnnouncementRequestAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateAnnouncementRequestAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


