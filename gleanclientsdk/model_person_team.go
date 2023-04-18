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

// checks if the PersonTeam type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PersonTeam{}

// PersonTeam Use `id` if you index teams via Glean, and use `name` and `externalLink` if you want to use your own team pages
type PersonTeam struct {
	// Unique identifier
	Id *string `json:"id,omitempty"`
	// Team name
	Name *string `json:"name,omitempty"`
	// Link to a team page on the internet or your company's intranet
	ExternalLink *string `json:"externalLink,omitempty"`
	// The team member's relationship to the team. This defaults to MEMBER if not set.
	Relationship *string `json:"relationship,omitempty"`
	// The team member's start date
	JoinDate *time.Time `json:"joinDate,omitempty"`
}

// NewPersonTeam instantiates a new PersonTeam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPersonTeam() *PersonTeam {
	this := PersonTeam{}
	var relationship string = "MEMBER"
	this.Relationship = &relationship
	return &this
}

// NewPersonTeamWithDefaults instantiates a new PersonTeam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPersonTeamWithDefaults() *PersonTeam {
	this := PersonTeam{}
	var relationship string = "MEMBER"
	this.Relationship = &relationship
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PersonTeam) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonTeam) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PersonTeam) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PersonTeam) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PersonTeam) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonTeam) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PersonTeam) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PersonTeam) SetName(v string) {
	o.Name = &v
}

// GetExternalLink returns the ExternalLink field value if set, zero value otherwise.
func (o *PersonTeam) GetExternalLink() string {
	if o == nil || IsNil(o.ExternalLink) {
		var ret string
		return ret
	}
	return *o.ExternalLink
}

// GetExternalLinkOk returns a tuple with the ExternalLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonTeam) GetExternalLinkOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalLink) {
		return nil, false
	}
	return o.ExternalLink, true
}

// HasExternalLink returns a boolean if a field has been set.
func (o *PersonTeam) HasExternalLink() bool {
	if o != nil && !IsNil(o.ExternalLink) {
		return true
	}

	return false
}

// SetExternalLink gets a reference to the given string and assigns it to the ExternalLink field.
func (o *PersonTeam) SetExternalLink(v string) {
	o.ExternalLink = &v
}

// GetRelationship returns the Relationship field value if set, zero value otherwise.
func (o *PersonTeam) GetRelationship() string {
	if o == nil || IsNil(o.Relationship) {
		var ret string
		return ret
	}
	return *o.Relationship
}

// GetRelationshipOk returns a tuple with the Relationship field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonTeam) GetRelationshipOk() (*string, bool) {
	if o == nil || IsNil(o.Relationship) {
		return nil, false
	}
	return o.Relationship, true
}

// HasRelationship returns a boolean if a field has been set.
func (o *PersonTeam) HasRelationship() bool {
	if o != nil && !IsNil(o.Relationship) {
		return true
	}

	return false
}

// SetRelationship gets a reference to the given string and assigns it to the Relationship field.
func (o *PersonTeam) SetRelationship(v string) {
	o.Relationship = &v
}

// GetJoinDate returns the JoinDate field value if set, zero value otherwise.
func (o *PersonTeam) GetJoinDate() time.Time {
	if o == nil || IsNil(o.JoinDate) {
		var ret time.Time
		return ret
	}
	return *o.JoinDate
}

// GetJoinDateOk returns a tuple with the JoinDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonTeam) GetJoinDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.JoinDate) {
		return nil, false
	}
	return o.JoinDate, true
}

// HasJoinDate returns a boolean if a field has been set.
func (o *PersonTeam) HasJoinDate() bool {
	if o != nil && !IsNil(o.JoinDate) {
		return true
	}

	return false
}

// SetJoinDate gets a reference to the given time.Time and assigns it to the JoinDate field.
func (o *PersonTeam) SetJoinDate(v time.Time) {
	o.JoinDate = &v
}

func (o PersonTeam) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PersonTeam) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ExternalLink) {
		toSerialize["externalLink"] = o.ExternalLink
	}
	if !IsNil(o.Relationship) {
		toSerialize["relationship"] = o.Relationship
	}
	if !IsNil(o.JoinDate) {
		toSerialize["joinDate"] = o.JoinDate
	}
	return toSerialize, nil
}

type NullablePersonTeam struct {
	value *PersonTeam
	isSet bool
}

func (v NullablePersonTeam) Get() *PersonTeam {
	return v.value
}

func (v *NullablePersonTeam) Set(val *PersonTeam) {
	v.value = val
	v.isSet = true
}

func (v NullablePersonTeam) IsSet() bool {
	return v.isSet
}

func (v *NullablePersonTeam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePersonTeam(val *PersonTeam) *NullablePersonTeam {
	return &NullablePersonTeam{value: val, isSet: true}
}

func (v NullablePersonTeam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePersonTeam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

