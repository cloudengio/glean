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

// checks if the CalendarAttendee type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CalendarAttendee{}

// CalendarAttendee struct for CalendarAttendee
type CalendarAttendee struct {
	// Whether or not this attendee is an organizer.
	IsOrganizer *bool `json:"isOrganizer,omitempty"`
	// Whether or not this attendee is in a group. Needed temporarily at least to support both flat attendees and tree for compatibility.
	IsInGroup *bool `json:"isInGroup,omitempty"`
	Person Person `json:"person"`
	// If this attendee is a group, represents the list of individual attendees in the group.
	GroupAttendees []CalendarAttendee `json:"groupAttendees,omitempty"`
	ResponseStatus *string `json:"responseStatus,omitempty"`
}

// NewCalendarAttendee instantiates a new CalendarAttendee object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCalendarAttendee(person Person) *CalendarAttendee {
	this := CalendarAttendee{}
	this.Person = person
	return &this
}

// NewCalendarAttendeeWithDefaults instantiates a new CalendarAttendee object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCalendarAttendeeWithDefaults() *CalendarAttendee {
	this := CalendarAttendee{}
	return &this
}

// GetIsOrganizer returns the IsOrganizer field value if set, zero value otherwise.
func (o *CalendarAttendee) GetIsOrganizer() bool {
	if o == nil || IsNil(o.IsOrganizer) {
		var ret bool
		return ret
	}
	return *o.IsOrganizer
}

// GetIsOrganizerOk returns a tuple with the IsOrganizer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalendarAttendee) GetIsOrganizerOk() (*bool, bool) {
	if o == nil || IsNil(o.IsOrganizer) {
		return nil, false
	}
	return o.IsOrganizer, true
}

// HasIsOrganizer returns a boolean if a field has been set.
func (o *CalendarAttendee) HasIsOrganizer() bool {
	if o != nil && !IsNil(o.IsOrganizer) {
		return true
	}

	return false
}

// SetIsOrganizer gets a reference to the given bool and assigns it to the IsOrganizer field.
func (o *CalendarAttendee) SetIsOrganizer(v bool) {
	o.IsOrganizer = &v
}

// GetIsInGroup returns the IsInGroup field value if set, zero value otherwise.
func (o *CalendarAttendee) GetIsInGroup() bool {
	if o == nil || IsNil(o.IsInGroup) {
		var ret bool
		return ret
	}
	return *o.IsInGroup
}

// GetIsInGroupOk returns a tuple with the IsInGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalendarAttendee) GetIsInGroupOk() (*bool, bool) {
	if o == nil || IsNil(o.IsInGroup) {
		return nil, false
	}
	return o.IsInGroup, true
}

// HasIsInGroup returns a boolean if a field has been set.
func (o *CalendarAttendee) HasIsInGroup() bool {
	if o != nil && !IsNil(o.IsInGroup) {
		return true
	}

	return false
}

// SetIsInGroup gets a reference to the given bool and assigns it to the IsInGroup field.
func (o *CalendarAttendee) SetIsInGroup(v bool) {
	o.IsInGroup = &v
}

// GetPerson returns the Person field value
func (o *CalendarAttendee) GetPerson() Person {
	if o == nil {
		var ret Person
		return ret
	}

	return o.Person
}

// GetPersonOk returns a tuple with the Person field value
// and a boolean to check if the value has been set.
func (o *CalendarAttendee) GetPersonOk() (*Person, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Person, true
}

// SetPerson sets field value
func (o *CalendarAttendee) SetPerson(v Person) {
	o.Person = v
}

// GetGroupAttendees returns the GroupAttendees field value if set, zero value otherwise.
func (o *CalendarAttendee) GetGroupAttendees() []CalendarAttendee {
	if o == nil || IsNil(o.GroupAttendees) {
		var ret []CalendarAttendee
		return ret
	}
	return o.GroupAttendees
}

// GetGroupAttendeesOk returns a tuple with the GroupAttendees field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalendarAttendee) GetGroupAttendeesOk() ([]CalendarAttendee, bool) {
	if o == nil || IsNil(o.GroupAttendees) {
		return nil, false
	}
	return o.GroupAttendees, true
}

// HasGroupAttendees returns a boolean if a field has been set.
func (o *CalendarAttendee) HasGroupAttendees() bool {
	if o != nil && !IsNil(o.GroupAttendees) {
		return true
	}

	return false
}

// SetGroupAttendees gets a reference to the given []CalendarAttendee and assigns it to the GroupAttendees field.
func (o *CalendarAttendee) SetGroupAttendees(v []CalendarAttendee) {
	o.GroupAttendees = v
}

// GetResponseStatus returns the ResponseStatus field value if set, zero value otherwise.
func (o *CalendarAttendee) GetResponseStatus() string {
	if o == nil || IsNil(o.ResponseStatus) {
		var ret string
		return ret
	}
	return *o.ResponseStatus
}

// GetResponseStatusOk returns a tuple with the ResponseStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalendarAttendee) GetResponseStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ResponseStatus) {
		return nil, false
	}
	return o.ResponseStatus, true
}

// HasResponseStatus returns a boolean if a field has been set.
func (o *CalendarAttendee) HasResponseStatus() bool {
	if o != nil && !IsNil(o.ResponseStatus) {
		return true
	}

	return false
}

// SetResponseStatus gets a reference to the given string and assigns it to the ResponseStatus field.
func (o *CalendarAttendee) SetResponseStatus(v string) {
	o.ResponseStatus = &v
}

func (o CalendarAttendee) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CalendarAttendee) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsOrganizer) {
		toSerialize["isOrganizer"] = o.IsOrganizer
	}
	if !IsNil(o.IsInGroup) {
		toSerialize["isInGroup"] = o.IsInGroup
	}
	toSerialize["person"] = o.Person
	if !IsNil(o.GroupAttendees) {
		toSerialize["groupAttendees"] = o.GroupAttendees
	}
	if !IsNil(o.ResponseStatus) {
		toSerialize["responseStatus"] = o.ResponseStatus
	}
	return toSerialize, nil
}

type NullableCalendarAttendee struct {
	value *CalendarAttendee
	isSet bool
}

func (v NullableCalendarAttendee) Get() *CalendarAttendee {
	return v.value
}

func (v *NullableCalendarAttendee) Set(val *CalendarAttendee) {
	v.value = val
	v.isSet = true
}

func (v NullableCalendarAttendee) IsSet() bool {
	return v.isSet
}

func (v *NullableCalendarAttendee) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCalendarAttendee(val *CalendarAttendee) *NullableCalendarAttendee {
	return &NullableCalendarAttendee{value: val, isSet: true}
}

func (v NullableCalendarAttendee) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCalendarAttendee) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


