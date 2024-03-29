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

// checks if the InviteInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InviteInfo{}

// InviteInfo Information regarding the invite status of a person.
type InviteInfo struct {
	Inviter *Person `json:"inviter,omitempty"`
	// The time this person was invited in ISO format (ISO 8601).
	InviteTime *time.Time `json:"inviteTime,omitempty"`
	// The time this person signed up in ISO format (ISO 8601).
	SignUpTime *time.Time `json:"signUpTime,omitempty"`
	// The time this person was reminded in ISO format (ISO 8601) if a reminder was sent.
	ReminderTime *time.Time `json:"reminderTime,omitempty"`
}

// NewInviteInfo instantiates a new InviteInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInviteInfo() *InviteInfo {
	this := InviteInfo{}
	return &this
}

// NewInviteInfoWithDefaults instantiates a new InviteInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInviteInfoWithDefaults() *InviteInfo {
	this := InviteInfo{}
	return &this
}

// GetInviter returns the Inviter field value if set, zero value otherwise.
func (o *InviteInfo) GetInviter() Person {
	if o == nil || IsNil(o.Inviter) {
		var ret Person
		return ret
	}
	return *o.Inviter
}

// GetInviterOk returns a tuple with the Inviter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InviteInfo) GetInviterOk() (*Person, bool) {
	if o == nil || IsNil(o.Inviter) {
		return nil, false
	}
	return o.Inviter, true
}

// HasInviter returns a boolean if a field has been set.
func (o *InviteInfo) HasInviter() bool {
	if o != nil && !IsNil(o.Inviter) {
		return true
	}

	return false
}

// SetInviter gets a reference to the given Person and assigns it to the Inviter field.
func (o *InviteInfo) SetInviter(v Person) {
	o.Inviter = &v
}

// GetInviteTime returns the InviteTime field value if set, zero value otherwise.
func (o *InviteInfo) GetInviteTime() time.Time {
	if o == nil || IsNil(o.InviteTime) {
		var ret time.Time
		return ret
	}
	return *o.InviteTime
}

// GetInviteTimeOk returns a tuple with the InviteTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InviteInfo) GetInviteTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.InviteTime) {
		return nil, false
	}
	return o.InviteTime, true
}

// HasInviteTime returns a boolean if a field has been set.
func (o *InviteInfo) HasInviteTime() bool {
	if o != nil && !IsNil(o.InviteTime) {
		return true
	}

	return false
}

// SetInviteTime gets a reference to the given time.Time and assigns it to the InviteTime field.
func (o *InviteInfo) SetInviteTime(v time.Time) {
	o.InviteTime = &v
}

// GetSignUpTime returns the SignUpTime field value if set, zero value otherwise.
func (o *InviteInfo) GetSignUpTime() time.Time {
	if o == nil || IsNil(o.SignUpTime) {
		var ret time.Time
		return ret
	}
	return *o.SignUpTime
}

// GetSignUpTimeOk returns a tuple with the SignUpTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InviteInfo) GetSignUpTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.SignUpTime) {
		return nil, false
	}
	return o.SignUpTime, true
}

// HasSignUpTime returns a boolean if a field has been set.
func (o *InviteInfo) HasSignUpTime() bool {
	if o != nil && !IsNil(o.SignUpTime) {
		return true
	}

	return false
}

// SetSignUpTime gets a reference to the given time.Time and assigns it to the SignUpTime field.
func (o *InviteInfo) SetSignUpTime(v time.Time) {
	o.SignUpTime = &v
}

// GetReminderTime returns the ReminderTime field value if set, zero value otherwise.
func (o *InviteInfo) GetReminderTime() time.Time {
	if o == nil || IsNil(o.ReminderTime) {
		var ret time.Time
		return ret
	}
	return *o.ReminderTime
}

// GetReminderTimeOk returns a tuple with the ReminderTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InviteInfo) GetReminderTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ReminderTime) {
		return nil, false
	}
	return o.ReminderTime, true
}

// HasReminderTime returns a boolean if a field has been set.
func (o *InviteInfo) HasReminderTime() bool {
	if o != nil && !IsNil(o.ReminderTime) {
		return true
	}

	return false
}

// SetReminderTime gets a reference to the given time.Time and assigns it to the ReminderTime field.
func (o *InviteInfo) SetReminderTime(v time.Time) {
	o.ReminderTime = &v
}

func (o InviteInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InviteInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Inviter) {
		toSerialize["inviter"] = o.Inviter
	}
	if !IsNil(o.InviteTime) {
		toSerialize["inviteTime"] = o.InviteTime
	}
	if !IsNil(o.SignUpTime) {
		toSerialize["signUpTime"] = o.SignUpTime
	}
	if !IsNil(o.ReminderTime) {
		toSerialize["reminderTime"] = o.ReminderTime
	}
	return toSerialize, nil
}

type NullableInviteInfo struct {
	value *InviteInfo
	isSet bool
}

func (v NullableInviteInfo) Get() *InviteInfo {
	return v.value
}

func (v *NullableInviteInfo) Set(val *InviteInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableInviteInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableInviteInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInviteInfo(val *InviteInfo) *NullableInviteInfo {
	return &NullableInviteInfo{value: val, isSet: true}
}

func (v NullableInviteInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInviteInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


