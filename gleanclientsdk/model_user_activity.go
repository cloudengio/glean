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

// checks if the UserActivity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserActivity{}

// UserActivity struct for UserActivity
type UserActivity struct {
	Actor *Person `json:"actor,omitempty"`
	// Unix timestamp of the activity (in seconds since epoch UTC).
	Timestamp *int32 `json:"timestamp,omitempty"`
	// The action for the activity
	Action *string `json:"action,omitempty"`
	AggregateVisitCount *CountInfo `json:"aggregateVisitCount,omitempty"`
}

// NewUserActivity instantiates a new UserActivity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserActivity() *UserActivity {
	this := UserActivity{}
	return &this
}

// NewUserActivityWithDefaults instantiates a new UserActivity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserActivityWithDefaults() *UserActivity {
	this := UserActivity{}
	return &this
}

// GetActor returns the Actor field value if set, zero value otherwise.
func (o *UserActivity) GetActor() Person {
	if o == nil || IsNil(o.Actor) {
		var ret Person
		return ret
	}
	return *o.Actor
}

// GetActorOk returns a tuple with the Actor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserActivity) GetActorOk() (*Person, bool) {
	if o == nil || IsNil(o.Actor) {
		return nil, false
	}
	return o.Actor, true
}

// HasActor returns a boolean if a field has been set.
func (o *UserActivity) HasActor() bool {
	if o != nil && !IsNil(o.Actor) {
		return true
	}

	return false
}

// SetActor gets a reference to the given Person and assigns it to the Actor field.
func (o *UserActivity) SetActor(v Person) {
	o.Actor = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *UserActivity) GetTimestamp() int32 {
	if o == nil || IsNil(o.Timestamp) {
		var ret int32
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserActivity) GetTimestampOk() (*int32, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *UserActivity) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int32 and assigns it to the Timestamp field.
func (o *UserActivity) SetTimestamp(v int32) {
	o.Timestamp = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *UserActivity) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserActivity) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *UserActivity) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *UserActivity) SetAction(v string) {
	o.Action = &v
}

// GetAggregateVisitCount returns the AggregateVisitCount field value if set, zero value otherwise.
func (o *UserActivity) GetAggregateVisitCount() CountInfo {
	if o == nil || IsNil(o.AggregateVisitCount) {
		var ret CountInfo
		return ret
	}
	return *o.AggregateVisitCount
}

// GetAggregateVisitCountOk returns a tuple with the AggregateVisitCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserActivity) GetAggregateVisitCountOk() (*CountInfo, bool) {
	if o == nil || IsNil(o.AggregateVisitCount) {
		return nil, false
	}
	return o.AggregateVisitCount, true
}

// HasAggregateVisitCount returns a boolean if a field has been set.
func (o *UserActivity) HasAggregateVisitCount() bool {
	if o != nil && !IsNil(o.AggregateVisitCount) {
		return true
	}

	return false
}

// SetAggregateVisitCount gets a reference to the given CountInfo and assigns it to the AggregateVisitCount field.
func (o *UserActivity) SetAggregateVisitCount(v CountInfo) {
	o.AggregateVisitCount = &v
}

func (o UserActivity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserActivity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Actor) {
		toSerialize["actor"] = o.Actor
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.AggregateVisitCount) {
		toSerialize["aggregateVisitCount"] = o.AggregateVisitCount
	}
	return toSerialize, nil
}

type NullableUserActivity struct {
	value *UserActivity
	isSet bool
}

func (v NullableUserActivity) Get() *UserActivity {
	return v.value
}

func (v *NullableUserActivity) Set(val *UserActivity) {
	v.value = val
	v.isSet = true
}

func (v NullableUserActivity) IsSet() bool {
	return v.isSet
}

func (v *NullableUserActivity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserActivity(val *UserActivity) *NullableUserActivity {
	return &NullableUserActivity{value: val, isSet: true}
}

func (v NullableUserActivity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserActivity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


