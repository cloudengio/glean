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

// checks if the TeamEmail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TeamEmail{}

// TeamEmail struct for TeamEmail
type TeamEmail struct {
	// An email address
	Email *string `json:"email,omitempty"`
	Type *string `json:"type,omitempty"`
	// true iff the email was manually added by a user from within Glean (aka not from the original data source)
	IsUserGenerated *bool `json:"isUserGenerated,omitempty"`
}

// NewTeamEmail instantiates a new TeamEmail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeamEmail() *TeamEmail {
	this := TeamEmail{}
	var type_ string = "OTHER"
	this.Type = &type_
	return &this
}

// NewTeamEmailWithDefaults instantiates a new TeamEmail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeamEmailWithDefaults() *TeamEmail {
	this := TeamEmail{}
	var type_ string = "OTHER"
	this.Type = &type_
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *TeamEmail) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamEmail) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *TeamEmail) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *TeamEmail) SetEmail(v string) {
	o.Email = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TeamEmail) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamEmail) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TeamEmail) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TeamEmail) SetType(v string) {
	o.Type = &v
}

// GetIsUserGenerated returns the IsUserGenerated field value if set, zero value otherwise.
func (o *TeamEmail) GetIsUserGenerated() bool {
	if o == nil || IsNil(o.IsUserGenerated) {
		var ret bool
		return ret
	}
	return *o.IsUserGenerated
}

// GetIsUserGeneratedOk returns a tuple with the IsUserGenerated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamEmail) GetIsUserGeneratedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsUserGenerated) {
		return nil, false
	}
	return o.IsUserGenerated, true
}

// HasIsUserGenerated returns a boolean if a field has been set.
func (o *TeamEmail) HasIsUserGenerated() bool {
	if o != nil && !IsNil(o.IsUserGenerated) {
		return true
	}

	return false
}

// SetIsUserGenerated gets a reference to the given bool and assigns it to the IsUserGenerated field.
func (o *TeamEmail) SetIsUserGenerated(v bool) {
	o.IsUserGenerated = &v
}

func (o TeamEmail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TeamEmail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.IsUserGenerated) {
		toSerialize["isUserGenerated"] = o.IsUserGenerated
	}
	return toSerialize, nil
}

type NullableTeamEmail struct {
	value *TeamEmail
	isSet bool
}

func (v NullableTeamEmail) Get() *TeamEmail {
	return v.value
}

func (v *NullableTeamEmail) Set(val *TeamEmail) {
	v.value = val
	v.isSet = true
}

func (v NullableTeamEmail) IsSet() bool {
	return v.isSet
}

func (v *NullableTeamEmail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeamEmail(val *TeamEmail) *NullableTeamEmail {
	return &NullableTeamEmail{value: val, isSet: true}
}

func (v NullableTeamEmail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeamEmail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


