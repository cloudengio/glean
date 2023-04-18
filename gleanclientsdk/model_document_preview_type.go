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

// DocumentPreviewType the model 'DocumentPreviewType'
type DocumentPreviewType string

// List of DocumentPreviewType
const (
	ANNOUNCEMENT DocumentPreviewType = "ANNOUNCEMENT"
	ANSWER DocumentPreviewType = "ANSWER"
	ANSWER_BOARD DocumentPreviewType = "ANSWER_BOARD"
	COLLECTION DocumentPreviewType = "COLLECTION"
	CONVERSATION DocumentPreviewType = "CONVERSATION"
	DETAILS DocumentPreviewType = "DETAILS"
	NO_PREVIEW DocumentPreviewType = "NO_PREVIEW"
	THREAD DocumentPreviewType = "THREAD"
	URL DocumentPreviewType = "URL"
)

// All allowed values of DocumentPreviewType enum
var AllowedDocumentPreviewTypeEnumValues = []DocumentPreviewType{
	"ANNOUNCEMENT",
	"ANSWER",
	"ANSWER_BOARD",
	"COLLECTION",
	"CONVERSATION",
	"DETAILS",
	"NO_PREVIEW",
	"THREAD",
	"URL",
}

func (v *DocumentPreviewType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DocumentPreviewType(value)
	for _, existing := range AllowedDocumentPreviewTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DocumentPreviewType", value)
}

// NewDocumentPreviewTypeFromValue returns a pointer to a valid DocumentPreviewType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDocumentPreviewTypeFromValue(v string) (*DocumentPreviewType, error) {
	ev := DocumentPreviewType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DocumentPreviewType: valid values are %v", v, AllowedDocumentPreviewTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DocumentPreviewType) IsValid() bool {
	for _, existing := range AllowedDocumentPreviewTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DocumentPreviewType value
func (v DocumentPreviewType) Ptr() *DocumentPreviewType {
	return &v
}

type NullableDocumentPreviewType struct {
	value *DocumentPreviewType
	isSet bool
}

func (v NullableDocumentPreviewType) Get() *DocumentPreviewType {
	return v.value
}

func (v *NullableDocumentPreviewType) Set(val *DocumentPreviewType) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentPreviewType) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentPreviewType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentPreviewType(val *DocumentPreviewType) *NullableDocumentPreviewType {
	return &NullableDocumentPreviewType{value: val, isSet: true}
}

func (v NullableDocumentPreviewType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentPreviewType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

