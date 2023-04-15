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

// checks if the CalendarEventAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CalendarEventAllOf{}

// CalendarEventAllOf struct for CalendarEventAllOf
type CalendarEventAllOf struct {
	// The calendar event id
	Id *string `json:"id,omitempty"`
	// A permalink for this calendar event
	Url *string `json:"url,omitempty"`
	Attendees *CalendarAttendees `json:"attendees,omitempty"`
	// The location that this event is taking place at.
	Location *string `json:"location,omitempty"`
	ConferenceData *ConferenceData `json:"conferenceData,omitempty"`
	// The HTML description of the event.
	Description *string `json:"description,omitempty"`
	// The app or other repository type from which the event was extracted
	Datasource *string `json:"datasource,omitempty"`
}

// NewCalendarEventAllOf instantiates a new CalendarEventAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCalendarEventAllOf() *CalendarEventAllOf {
	this := CalendarEventAllOf{}
	return &this
}

// NewCalendarEventAllOfWithDefaults instantiates a new CalendarEventAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCalendarEventAllOfWithDefaults() *CalendarEventAllOf {
	this := CalendarEventAllOf{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CalendarEventAllOf) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalendarEventAllOf) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CalendarEventAllOf) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CalendarEventAllOf) SetId(v string) {
	o.Id = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *CalendarEventAllOf) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalendarEventAllOf) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *CalendarEventAllOf) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *CalendarEventAllOf) SetUrl(v string) {
	o.Url = &v
}

// GetAttendees returns the Attendees field value if set, zero value otherwise.
func (o *CalendarEventAllOf) GetAttendees() CalendarAttendees {
	if o == nil || IsNil(o.Attendees) {
		var ret CalendarAttendees
		return ret
	}
	return *o.Attendees
}

// GetAttendeesOk returns a tuple with the Attendees field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalendarEventAllOf) GetAttendeesOk() (*CalendarAttendees, bool) {
	if o == nil || IsNil(o.Attendees) {
		return nil, false
	}
	return o.Attendees, true
}

// HasAttendees returns a boolean if a field has been set.
func (o *CalendarEventAllOf) HasAttendees() bool {
	if o != nil && !IsNil(o.Attendees) {
		return true
	}

	return false
}

// SetAttendees gets a reference to the given CalendarAttendees and assigns it to the Attendees field.
func (o *CalendarEventAllOf) SetAttendees(v CalendarAttendees) {
	o.Attendees = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *CalendarEventAllOf) GetLocation() string {
	if o == nil || IsNil(o.Location) {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalendarEventAllOf) GetLocationOk() (*string, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *CalendarEventAllOf) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *CalendarEventAllOf) SetLocation(v string) {
	o.Location = &v
}

// GetConferenceData returns the ConferenceData field value if set, zero value otherwise.
func (o *CalendarEventAllOf) GetConferenceData() ConferenceData {
	if o == nil || IsNil(o.ConferenceData) {
		var ret ConferenceData
		return ret
	}
	return *o.ConferenceData
}

// GetConferenceDataOk returns a tuple with the ConferenceData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalendarEventAllOf) GetConferenceDataOk() (*ConferenceData, bool) {
	if o == nil || IsNil(o.ConferenceData) {
		return nil, false
	}
	return o.ConferenceData, true
}

// HasConferenceData returns a boolean if a field has been set.
func (o *CalendarEventAllOf) HasConferenceData() bool {
	if o != nil && !IsNil(o.ConferenceData) {
		return true
	}

	return false
}

// SetConferenceData gets a reference to the given ConferenceData and assigns it to the ConferenceData field.
func (o *CalendarEventAllOf) SetConferenceData(v ConferenceData) {
	o.ConferenceData = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CalendarEventAllOf) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalendarEventAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CalendarEventAllOf) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CalendarEventAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetDatasource returns the Datasource field value if set, zero value otherwise.
func (o *CalendarEventAllOf) GetDatasource() string {
	if o == nil || IsNil(o.Datasource) {
		var ret string
		return ret
	}
	return *o.Datasource
}

// GetDatasourceOk returns a tuple with the Datasource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalendarEventAllOf) GetDatasourceOk() (*string, bool) {
	if o == nil || IsNil(o.Datasource) {
		return nil, false
	}
	return o.Datasource, true
}

// HasDatasource returns a boolean if a field has been set.
func (o *CalendarEventAllOf) HasDatasource() bool {
	if o != nil && !IsNil(o.Datasource) {
		return true
	}

	return false
}

// SetDatasource gets a reference to the given string and assigns it to the Datasource field.
func (o *CalendarEventAllOf) SetDatasource(v string) {
	o.Datasource = &v
}

func (o CalendarEventAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CalendarEventAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Attendees) {
		toSerialize["attendees"] = o.Attendees
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !IsNil(o.ConferenceData) {
		toSerialize["conferenceData"] = o.ConferenceData
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Datasource) {
		toSerialize["datasource"] = o.Datasource
	}
	return toSerialize, nil
}

type NullableCalendarEventAllOf struct {
	value *CalendarEventAllOf
	isSet bool
}

func (v NullableCalendarEventAllOf) Get() *CalendarEventAllOf {
	return v.value
}

func (v *NullableCalendarEventAllOf) Set(val *CalendarEventAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCalendarEventAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCalendarEventAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCalendarEventAllOf(val *CalendarEventAllOf) *NullableCalendarEventAllOf {
	return &NullableCalendarEventAllOf{value: val, isSet: true}
}

func (v NullableCalendarEventAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCalendarEventAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


