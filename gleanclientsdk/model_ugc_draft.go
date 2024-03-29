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

// checks if the UgcDraft type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UgcDraft{}

// UgcDraft struct for UgcDraft
type UgcDraft struct {
	Announcement *AnnouncementMutableProperties `json:"announcement,omitempty"`
	Answer *AnswerMutableProperties `json:"answer,omitempty"`
}

// NewUgcDraft instantiates a new UgcDraft object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUgcDraft() *UgcDraft {
	this := UgcDraft{}
	return &this
}

// NewUgcDraftWithDefaults instantiates a new UgcDraft object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUgcDraftWithDefaults() *UgcDraft {
	this := UgcDraft{}
	return &this
}

// GetAnnouncement returns the Announcement field value if set, zero value otherwise.
func (o *UgcDraft) GetAnnouncement() AnnouncementMutableProperties {
	if o == nil || IsNil(o.Announcement) {
		var ret AnnouncementMutableProperties
		return ret
	}
	return *o.Announcement
}

// GetAnnouncementOk returns a tuple with the Announcement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UgcDraft) GetAnnouncementOk() (*AnnouncementMutableProperties, bool) {
	if o == nil || IsNil(o.Announcement) {
		return nil, false
	}
	return o.Announcement, true
}

// HasAnnouncement returns a boolean if a field has been set.
func (o *UgcDraft) HasAnnouncement() bool {
	if o != nil && !IsNil(o.Announcement) {
		return true
	}

	return false
}

// SetAnnouncement gets a reference to the given AnnouncementMutableProperties and assigns it to the Announcement field.
func (o *UgcDraft) SetAnnouncement(v AnnouncementMutableProperties) {
	o.Announcement = &v
}

// GetAnswer returns the Answer field value if set, zero value otherwise.
func (o *UgcDraft) GetAnswer() AnswerMutableProperties {
	if o == nil || IsNil(o.Answer) {
		var ret AnswerMutableProperties
		return ret
	}
	return *o.Answer
}

// GetAnswerOk returns a tuple with the Answer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UgcDraft) GetAnswerOk() (*AnswerMutableProperties, bool) {
	if o == nil || IsNil(o.Answer) {
		return nil, false
	}
	return o.Answer, true
}

// HasAnswer returns a boolean if a field has been set.
func (o *UgcDraft) HasAnswer() bool {
	if o != nil && !IsNil(o.Answer) {
		return true
	}

	return false
}

// SetAnswer gets a reference to the given AnswerMutableProperties and assigns it to the Answer field.
func (o *UgcDraft) SetAnswer(v AnswerMutableProperties) {
	o.Answer = &v
}

func (o UgcDraft) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UgcDraft) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Announcement) {
		toSerialize["announcement"] = o.Announcement
	}
	if !IsNil(o.Answer) {
		toSerialize["answer"] = o.Answer
	}
	return toSerialize, nil
}

type NullableUgcDraft struct {
	value *UgcDraft
	isSet bool
}

func (v NullableUgcDraft) Get() *UgcDraft {
	return v.value
}

func (v *NullableUgcDraft) Set(val *UgcDraft) {
	v.value = val
	v.isSet = true
}

func (v NullableUgcDraft) IsSet() bool {
	return v.isSet
}

func (v *NullableUgcDraft) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUgcDraft(val *UgcDraft) *NullableUgcDraft {
	return &NullableUgcDraft{value: val, isSet: true}
}

func (v NullableUgcDraft) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUgcDraft) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


