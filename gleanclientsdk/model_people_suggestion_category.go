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
	"fmt"
)

// PeopleSuggestionCategory the model 'PeopleSuggestionCategory'
type PeopleSuggestionCategory string

// List of PeopleSuggestionCategory
const (
	INACTIVE_PROMO PeopleSuggestionCategory = "INVITE_INACTIVE_PROMO"
	NONUSERS PeopleSuggestionCategory = "INVITE_NONUSERS"
)

// All allowed values of PeopleSuggestionCategory enum
var AllowedPeopleSuggestionCategoryEnumValues = []PeopleSuggestionCategory{
	"INVITE_INACTIVE_PROMO",
	"INVITE_NONUSERS",
}

func (v *PeopleSuggestionCategory) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PeopleSuggestionCategory(value)
	for _, existing := range AllowedPeopleSuggestionCategoryEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PeopleSuggestionCategory", value)
}

// NewPeopleSuggestionCategoryFromValue returns a pointer to a valid PeopleSuggestionCategory
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPeopleSuggestionCategoryFromValue(v string) (*PeopleSuggestionCategory, error) {
	ev := PeopleSuggestionCategory(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PeopleSuggestionCategory: valid values are %v", v, AllowedPeopleSuggestionCategoryEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PeopleSuggestionCategory) IsValid() bool {
	for _, existing := range AllowedPeopleSuggestionCategoryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PeopleSuggestionCategory value
func (v PeopleSuggestionCategory) Ptr() *PeopleSuggestionCategory {
	return &v
}

type NullablePeopleSuggestionCategory struct {
	value *PeopleSuggestionCategory
	isSet bool
}

func (v NullablePeopleSuggestionCategory) Get() *PeopleSuggestionCategory {
	return v.value
}

func (v *NullablePeopleSuggestionCategory) Set(val *PeopleSuggestionCategory) {
	v.value = val
	v.isSet = true
}

func (v NullablePeopleSuggestionCategory) IsSet() bool {
	return v.isSet
}

func (v *NullablePeopleSuggestionCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePeopleSuggestionCategory(val *PeopleSuggestionCategory) *NullablePeopleSuggestionCategory {
	return &NullablePeopleSuggestionCategory{value: val, isSet: true}
}

func (v NullablePeopleSuggestionCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePeopleSuggestionCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

