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
	"fmt"
)

// DocumentVisibility The level of visibility of the document as understood by our system. PRIVATE - Only one person is able to see the document SPECIFIC_PEOPLE_AND_GROUPS - Only specific people and/or groups can see the document DOMAIN_LINK - Anyone in the domain with the link can see the document DOMAIN_VISIBLE - Anyone in the domain can search for the document PUBLIC_LINK - Anyone with the link can see the document PUBLIC_VISIBLE - Anyone on the internet can search for the document
type DocumentVisibility string

// List of DocumentVisibility
const (
	PRIVATE DocumentVisibility = "PRIVATE"
	SPECIFIC_PEOPLE_AND_GROUPS DocumentVisibility = "SPECIFIC_PEOPLE_AND_GROUPS"
	DOMAIN_LINK DocumentVisibility = "DOMAIN_LINK"
	DOMAIN_VISIBLE DocumentVisibility = "DOMAIN_VISIBLE"
	PUBLIC_LINK DocumentVisibility = "PUBLIC_LINK"
	PUBLIC_VISIBLE DocumentVisibility = "PUBLIC_VISIBLE"
)

// All allowed values of DocumentVisibility enum
var AllowedDocumentVisibilityEnumValues = []DocumentVisibility{
	"PRIVATE",
	"SPECIFIC_PEOPLE_AND_GROUPS",
	"DOMAIN_LINK",
	"DOMAIN_VISIBLE",
	"PUBLIC_LINK",
	"PUBLIC_VISIBLE",
}

func (v *DocumentVisibility) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DocumentVisibility(value)
	for _, existing := range AllowedDocumentVisibilityEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DocumentVisibility", value)
}

// NewDocumentVisibilityFromValue returns a pointer to a valid DocumentVisibility
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDocumentVisibilityFromValue(v string) (*DocumentVisibility, error) {
	ev := DocumentVisibility(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DocumentVisibility: valid values are %v", v, AllowedDocumentVisibilityEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DocumentVisibility) IsValid() bool {
	for _, existing := range AllowedDocumentVisibilityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DocumentVisibility value
func (v DocumentVisibility) Ptr() *DocumentVisibility {
	return &v
}

type NullableDocumentVisibility struct {
	value *DocumentVisibility
	isSet bool
}

func (v NullableDocumentVisibility) Get() *DocumentVisibility {
	return v.value
}

func (v *NullableDocumentVisibility) Set(val *DocumentVisibility) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentVisibility) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentVisibility) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentVisibility(val *DocumentVisibility) *NullableDocumentVisibility {
	return &NullableDocumentVisibility{value: val, isSet: true}
}

func (v NullableDocumentVisibility) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentVisibility) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

