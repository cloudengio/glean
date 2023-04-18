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

// checks if the PersonToTeamRelationship type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PersonToTeamRelationship{}

// PersonToTeamRelationship Metadata about the relationship of a person to a team.
type PersonToTeamRelationship struct {
	Person Person `json:"person"`
	// The team member's relationship to the team. This defaults to MEMBER if not set.
	Relationship *string `json:"relationship,omitempty"`
	// Displayed name for the relationship if relationship is set to `OTHER`.
	CustomRelationshipStr *string `json:"customRelationshipStr,omitempty"`
	// The team member's start date
	JoinDate *time.Time `json:"joinDate,omitempty"`
}

// NewPersonToTeamRelationship instantiates a new PersonToTeamRelationship object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPersonToTeamRelationship(person Person) *PersonToTeamRelationship {
	this := PersonToTeamRelationship{}
	this.Person = person
	var relationship string = "MEMBER"
	this.Relationship = &relationship
	return &this
}

// NewPersonToTeamRelationshipWithDefaults instantiates a new PersonToTeamRelationship object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPersonToTeamRelationshipWithDefaults() *PersonToTeamRelationship {
	this := PersonToTeamRelationship{}
	var relationship string = "MEMBER"
	this.Relationship = &relationship
	return &this
}

// GetPerson returns the Person field value
func (o *PersonToTeamRelationship) GetPerson() Person {
	if o == nil {
		var ret Person
		return ret
	}

	return o.Person
}

// GetPersonOk returns a tuple with the Person field value
// and a boolean to check if the value has been set.
func (o *PersonToTeamRelationship) GetPersonOk() (*Person, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Person, true
}

// SetPerson sets field value
func (o *PersonToTeamRelationship) SetPerson(v Person) {
	o.Person = v
}

// GetRelationship returns the Relationship field value if set, zero value otherwise.
func (o *PersonToTeamRelationship) GetRelationship() string {
	if o == nil || IsNil(o.Relationship) {
		var ret string
		return ret
	}
	return *o.Relationship
}

// GetRelationshipOk returns a tuple with the Relationship field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonToTeamRelationship) GetRelationshipOk() (*string, bool) {
	if o == nil || IsNil(o.Relationship) {
		return nil, false
	}
	return o.Relationship, true
}

// HasRelationship returns a boolean if a field has been set.
func (o *PersonToTeamRelationship) HasRelationship() bool {
	if o != nil && !IsNil(o.Relationship) {
		return true
	}

	return false
}

// SetRelationship gets a reference to the given string and assigns it to the Relationship field.
func (o *PersonToTeamRelationship) SetRelationship(v string) {
	o.Relationship = &v
}

// GetCustomRelationshipStr returns the CustomRelationshipStr field value if set, zero value otherwise.
func (o *PersonToTeamRelationship) GetCustomRelationshipStr() string {
	if o == nil || IsNil(o.CustomRelationshipStr) {
		var ret string
		return ret
	}
	return *o.CustomRelationshipStr
}

// GetCustomRelationshipStrOk returns a tuple with the CustomRelationshipStr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonToTeamRelationship) GetCustomRelationshipStrOk() (*string, bool) {
	if o == nil || IsNil(o.CustomRelationshipStr) {
		return nil, false
	}
	return o.CustomRelationshipStr, true
}

// HasCustomRelationshipStr returns a boolean if a field has been set.
func (o *PersonToTeamRelationship) HasCustomRelationshipStr() bool {
	if o != nil && !IsNil(o.CustomRelationshipStr) {
		return true
	}

	return false
}

// SetCustomRelationshipStr gets a reference to the given string and assigns it to the CustomRelationshipStr field.
func (o *PersonToTeamRelationship) SetCustomRelationshipStr(v string) {
	o.CustomRelationshipStr = &v
}

// GetJoinDate returns the JoinDate field value if set, zero value otherwise.
func (o *PersonToTeamRelationship) GetJoinDate() time.Time {
	if o == nil || IsNil(o.JoinDate) {
		var ret time.Time
		return ret
	}
	return *o.JoinDate
}

// GetJoinDateOk returns a tuple with the JoinDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonToTeamRelationship) GetJoinDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.JoinDate) {
		return nil, false
	}
	return o.JoinDate, true
}

// HasJoinDate returns a boolean if a field has been set.
func (o *PersonToTeamRelationship) HasJoinDate() bool {
	if o != nil && !IsNil(o.JoinDate) {
		return true
	}

	return false
}

// SetJoinDate gets a reference to the given time.Time and assigns it to the JoinDate field.
func (o *PersonToTeamRelationship) SetJoinDate(v time.Time) {
	o.JoinDate = &v
}

func (o PersonToTeamRelationship) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PersonToTeamRelationship) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["person"] = o.Person
	if !IsNil(o.Relationship) {
		toSerialize["relationship"] = o.Relationship
	}
	if !IsNil(o.CustomRelationshipStr) {
		toSerialize["customRelationshipStr"] = o.CustomRelationshipStr
	}
	if !IsNil(o.JoinDate) {
		toSerialize["joinDate"] = o.JoinDate
	}
	return toSerialize, nil
}

type NullablePersonToTeamRelationship struct {
	value *PersonToTeamRelationship
	isSet bool
}

func (v NullablePersonToTeamRelationship) Get() *PersonToTeamRelationship {
	return v.value
}

func (v *NullablePersonToTeamRelationship) Set(val *PersonToTeamRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullablePersonToTeamRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullablePersonToTeamRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePersonToTeamRelationship(val *PersonToTeamRelationship) *NullablePersonToTeamRelationship {
	return &NullablePersonToTeamRelationship{value: val, isSet: true}
}

func (v NullablePersonToTeamRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePersonToTeamRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


