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

// checks if the UpdateDraftAnnouncementRequestAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateDraftAnnouncementRequestAllOf{}

// UpdateDraftAnnouncementRequestAllOf struct for UpdateDraftAnnouncementRequestAllOf
type UpdateDraftAnnouncementRequestAllOf struct {
	// The opaque id of the announcement.
	Id *int32 `json:"id,omitempty"`
	// The opaque id of the draft.
	DraftId int32 `json:"draftId"`
}

// NewUpdateDraftAnnouncementRequestAllOf instantiates a new UpdateDraftAnnouncementRequestAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDraftAnnouncementRequestAllOf(draftId int32) *UpdateDraftAnnouncementRequestAllOf {
	this := UpdateDraftAnnouncementRequestAllOf{}
	this.DraftId = draftId
	return &this
}

// NewUpdateDraftAnnouncementRequestAllOfWithDefaults instantiates a new UpdateDraftAnnouncementRequestAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDraftAnnouncementRequestAllOfWithDefaults() *UpdateDraftAnnouncementRequestAllOf {
	this := UpdateDraftAnnouncementRequestAllOf{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UpdateDraftAnnouncementRequestAllOf) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDraftAnnouncementRequestAllOf) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UpdateDraftAnnouncementRequestAllOf) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *UpdateDraftAnnouncementRequestAllOf) SetId(v int32) {
	o.Id = &v
}

// GetDraftId returns the DraftId field value
func (o *UpdateDraftAnnouncementRequestAllOf) GetDraftId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DraftId
}

// GetDraftIdOk returns a tuple with the DraftId field value
// and a boolean to check if the value has been set.
func (o *UpdateDraftAnnouncementRequestAllOf) GetDraftIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DraftId, true
}

// SetDraftId sets field value
func (o *UpdateDraftAnnouncementRequestAllOf) SetDraftId(v int32) {
	o.DraftId = v
}

func (o UpdateDraftAnnouncementRequestAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateDraftAnnouncementRequestAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["draftId"] = o.DraftId
	return toSerialize, nil
}

type NullableUpdateDraftAnnouncementRequestAllOf struct {
	value *UpdateDraftAnnouncementRequestAllOf
	isSet bool
}

func (v NullableUpdateDraftAnnouncementRequestAllOf) Get() *UpdateDraftAnnouncementRequestAllOf {
	return v.value
}

func (v *NullableUpdateDraftAnnouncementRequestAllOf) Set(val *UpdateDraftAnnouncementRequestAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDraftAnnouncementRequestAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDraftAnnouncementRequestAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDraftAnnouncementRequestAllOf(val *UpdateDraftAnnouncementRequestAllOf) *NullableUpdateDraftAnnouncementRequestAllOf {
	return &NullableUpdateDraftAnnouncementRequestAllOf{value: val, isSet: true}
}

func (v NullableUpdateDraftAnnouncementRequestAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDraftAnnouncementRequestAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

