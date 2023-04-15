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

// CollectionPinnableTargets What targets can a collection be pinned to
type CollectionPinnableTargets string

// List of CollectionPinnableTargets
const (
	RESOURCE_CARD CollectionPinnableTargets = "RESOURCE_CARD"
	TEAM_PROFILE_PAGE CollectionPinnableTargets = "TEAM_PROFILE_PAGE"
)

// All allowed values of CollectionPinnableTargets enum
var AllowedCollectionPinnableTargetsEnumValues = []CollectionPinnableTargets{
	"RESOURCE_CARD",
	"TEAM_PROFILE_PAGE",
}

func (v *CollectionPinnableTargets) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CollectionPinnableTargets(value)
	for _, existing := range AllowedCollectionPinnableTargetsEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CollectionPinnableTargets", value)
}

// NewCollectionPinnableTargetsFromValue returns a pointer to a valid CollectionPinnableTargets
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCollectionPinnableTargetsFromValue(v string) (*CollectionPinnableTargets, error) {
	ev := CollectionPinnableTargets(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CollectionPinnableTargets: valid values are %v", v, AllowedCollectionPinnableTargetsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CollectionPinnableTargets) IsValid() bool {
	for _, existing := range AllowedCollectionPinnableTargetsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CollectionPinnableTargets value
func (v CollectionPinnableTargets) Ptr() *CollectionPinnableTargets {
	return &v
}

type NullableCollectionPinnableTargets struct {
	value *CollectionPinnableTargets
	isSet bool
}

func (v NullableCollectionPinnableTargets) Get() *CollectionPinnableTargets {
	return v.value
}

func (v *NullableCollectionPinnableTargets) Set(val *CollectionPinnableTargets) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionPinnableTargets) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionPinnableTargets) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionPinnableTargets(val *CollectionPinnableTargets) *NullableCollectionPinnableTargets {
	return &NullableCollectionPinnableTargets{value: val, isSet: true}
}

func (v NullableCollectionPinnableTargets) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionPinnableTargets) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

