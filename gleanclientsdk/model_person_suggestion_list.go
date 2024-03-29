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

// checks if the PersonSuggestionList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PersonSuggestionList{}

// PersonSuggestionList struct for PersonSuggestionList
type PersonSuggestionList struct {
	Category PeopleSuggestionCategory `json:"category"`
	// Information about suggested users.
	People []Person `json:"people,omitempty"`
}

// NewPersonSuggestionList instantiates a new PersonSuggestionList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPersonSuggestionList(category PeopleSuggestionCategory) *PersonSuggestionList {
	this := PersonSuggestionList{}
	this.Category = category
	return &this
}

// NewPersonSuggestionListWithDefaults instantiates a new PersonSuggestionList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPersonSuggestionListWithDefaults() *PersonSuggestionList {
	this := PersonSuggestionList{}
	return &this
}

// GetCategory returns the Category field value
func (o *PersonSuggestionList) GetCategory() PeopleSuggestionCategory {
	if o == nil {
		var ret PeopleSuggestionCategory
		return ret
	}

	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *PersonSuggestionList) GetCategoryOk() (*PeopleSuggestionCategory, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value
func (o *PersonSuggestionList) SetCategory(v PeopleSuggestionCategory) {
	o.Category = v
}

// GetPeople returns the People field value if set, zero value otherwise.
func (o *PersonSuggestionList) GetPeople() []Person {
	if o == nil || IsNil(o.People) {
		var ret []Person
		return ret
	}
	return o.People
}

// GetPeopleOk returns a tuple with the People field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonSuggestionList) GetPeopleOk() ([]Person, bool) {
	if o == nil || IsNil(o.People) {
		return nil, false
	}
	return o.People, true
}

// HasPeople returns a boolean if a field has been set.
func (o *PersonSuggestionList) HasPeople() bool {
	if o != nil && !IsNil(o.People) {
		return true
	}

	return false
}

// SetPeople gets a reference to the given []Person and assigns it to the People field.
func (o *PersonSuggestionList) SetPeople(v []Person) {
	o.People = v
}

func (o PersonSuggestionList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PersonSuggestionList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["category"] = o.Category
	if !IsNil(o.People) {
		toSerialize["people"] = o.People
	}
	return toSerialize, nil
}

type NullablePersonSuggestionList struct {
	value *PersonSuggestionList
	isSet bool
}

func (v NullablePersonSuggestionList) Get() *PersonSuggestionList {
	return v.value
}

func (v *NullablePersonSuggestionList) Set(val *PersonSuggestionList) {
	v.value = val
	v.isSet = true
}

func (v NullablePersonSuggestionList) IsSet() bool {
	return v.isSet
}

func (v *NullablePersonSuggestionList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePersonSuggestionList(val *PersonSuggestionList) *NullablePersonSuggestionList {
	return &NullablePersonSuggestionList{value: val, isSet: true}
}

func (v NullablePersonSuggestionList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePersonSuggestionList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


