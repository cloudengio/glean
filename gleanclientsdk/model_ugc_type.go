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

// UgcType the model 'UgcType'
type UgcType string

// List of UgcType
const (
	ANNOUNCEMENTS_TYPE UgcType = "ANNOUNCEMENTS_TYPE"
	ANSWERS_TYPE UgcType = "ANSWERS_TYPE"
	COLLECTIONS_TYPE UgcType = "COLLECTIONS_TYPE"
	SHORTCUTS_TYPE UgcType = "SHORTCUTS_TYPE"
)

// All allowed values of UgcType enum
var AllowedUgcTypeEnumValues = []UgcType{
	"ANNOUNCEMENTS_TYPE",
	"ANSWERS_TYPE",
	"COLLECTIONS_TYPE",
	"SHORTCUTS_TYPE",
}

func (v *UgcType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UgcType(value)
	for _, existing := range AllowedUgcTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UgcType", value)
}

// NewUgcTypeFromValue returns a pointer to a valid UgcType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUgcTypeFromValue(v string) (*UgcType, error) {
	ev := UgcType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UgcType: valid values are %v", v, AllowedUgcTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UgcType) IsValid() bool {
	for _, existing := range AllowedUgcTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UgcType value
func (v UgcType) Ptr() *UgcType {
	return &v
}

type NullableUgcType struct {
	value *UgcType
	isSet bool
}

func (v NullableUgcType) Get() *UgcType {
	return v.value
}

func (v *NullableUgcType) Set(val *UgcType) {
	v.value = val
	v.isSet = true
}

func (v NullableUgcType) IsSet() bool {
	return v.isSet
}

func (v *NullableUgcType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUgcType(val *UgcType) *NullableUgcType {
	return &NullableUgcType{value: val, isSet: true}
}

func (v NullableUgcType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUgcType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

