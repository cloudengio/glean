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

// checks if the Reminder type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Reminder{}

// Reminder struct for Reminder
type Reminder struct {
	Assignee Person `json:"assignee"`
	Requestor *Person `json:"requestor,omitempty"`
	// Unix timestamp for when the reminder should trigger (in seconds since epoch UTC).
	RemindAt int32 `json:"remindAt"`
	// Unix timestamp for when the reminder was first created (in seconds since epoch UTC).
	CreatedAt *int32 `json:"createdAt,omitempty"`
	// An optional free-text reason for the reminder. This is particularly useful when a reminder is used to ask for verification from another user (for example, \"Duplicate\", \"Incomplete\", \"Incorrect\").
	Reason *string `json:"reason,omitempty"`
}

// NewReminder instantiates a new Reminder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReminder(assignee Person, remindAt int32) *Reminder {
	this := Reminder{}
	this.Assignee = assignee
	this.RemindAt = remindAt
	return &this
}

// NewReminderWithDefaults instantiates a new Reminder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReminderWithDefaults() *Reminder {
	this := Reminder{}
	return &this
}

// GetAssignee returns the Assignee field value
func (o *Reminder) GetAssignee() Person {
	if o == nil {
		var ret Person
		return ret
	}

	return o.Assignee
}

// GetAssigneeOk returns a tuple with the Assignee field value
// and a boolean to check if the value has been set.
func (o *Reminder) GetAssigneeOk() (*Person, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Assignee, true
}

// SetAssignee sets field value
func (o *Reminder) SetAssignee(v Person) {
	o.Assignee = v
}

// GetRequestor returns the Requestor field value if set, zero value otherwise.
func (o *Reminder) GetRequestor() Person {
	if o == nil || IsNil(o.Requestor) {
		var ret Person
		return ret
	}
	return *o.Requestor
}

// GetRequestorOk returns a tuple with the Requestor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Reminder) GetRequestorOk() (*Person, bool) {
	if o == nil || IsNil(o.Requestor) {
		return nil, false
	}
	return o.Requestor, true
}

// HasRequestor returns a boolean if a field has been set.
func (o *Reminder) HasRequestor() bool {
	if o != nil && !IsNil(o.Requestor) {
		return true
	}

	return false
}

// SetRequestor gets a reference to the given Person and assigns it to the Requestor field.
func (o *Reminder) SetRequestor(v Person) {
	o.Requestor = &v
}

// GetRemindAt returns the RemindAt field value
func (o *Reminder) GetRemindAt() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RemindAt
}

// GetRemindAtOk returns a tuple with the RemindAt field value
// and a boolean to check if the value has been set.
func (o *Reminder) GetRemindAtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RemindAt, true
}

// SetRemindAt sets field value
func (o *Reminder) SetRemindAt(v int32) {
	o.RemindAt = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Reminder) GetCreatedAt() int32 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret int32
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Reminder) GetCreatedAtOk() (*int32, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Reminder) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int32 and assigns it to the CreatedAt field.
func (o *Reminder) SetCreatedAt(v int32) {
	o.CreatedAt = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *Reminder) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Reminder) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *Reminder) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *Reminder) SetReason(v string) {
	o.Reason = &v
}

func (o Reminder) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Reminder) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["assignee"] = o.Assignee
	if !IsNil(o.Requestor) {
		toSerialize["requestor"] = o.Requestor
	}
	toSerialize["remindAt"] = o.RemindAt
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	return toSerialize, nil
}

type NullableReminder struct {
	value *Reminder
	isSet bool
}

func (v NullableReminder) Get() *Reminder {
	return v.value
}

func (v *NullableReminder) Set(val *Reminder) {
	v.value = val
	v.isSet = true
}

func (v NullableReminder) IsSet() bool {
	return v.isSet
}

func (v *NullableReminder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReminder(val *Reminder) *NullableReminder {
	return &NullableReminder{value: val, isSet: true}
}

func (v NullableReminder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReminder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


