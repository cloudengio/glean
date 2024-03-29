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

// PermissionedFeatureOrObject A feature or object type with restricted access
type PermissionedFeatureOrObject string

// List of PermissionedFeatureOrObject
const (
	ADMIN_SEARCH PermissionedFeatureOrObject = "ADMIN_SEARCH"
	ANNOUNCEMENTS PermissionedFeatureOrObject = "ANNOUNCEMENTS"
	ANSWERS PermissionedFeatureOrObject = "ANSWERS"
	COLLECTIONS PermissionedFeatureOrObject = "COLLECTIONS"
	COLLECTION_ITEMS PermissionedFeatureOrObject = "COLLECTION_ITEMS"
	INSIGHTS_COLLECTIONS PermissionedFeatureOrObject = "INSIGHTS_COLLECTIONS"
	INSIGHTS_CONTENT PermissionedFeatureOrObject = "INSIGHTS_CONTENT"
	INSIGHTS_INVITE PermissionedFeatureOrObject = "INSIGHTS_INVITE"
	INSIGHTS_SEARCHES PermissionedFeatureOrObject = "INSIGHTS_SEARCHES"
	INSIGHTS_TEAMMATES PermissionedFeatureOrObject = "INSIGHTS_TEAMMATES"
	PINS PermissionedFeatureOrObject = "PINS"
	SENSITIVE_CONTENT_REPORTS PermissionedFeatureOrObject = "SENSITIVE_CONTENT_REPORTS"
	SHORTCUTS PermissionedFeatureOrObject = "SHORTCUTS"
	TEAMS PermissionedFeatureOrObject = "TEAMS"
	VERIFICATIONS PermissionedFeatureOrObject = "VERIFICATIONS"
	WORKSPACE_APPS PermissionedFeatureOrObject = "WORKSPACE_APPS"
	WORKSPACE_PERMISSIONS PermissionedFeatureOrObject = "WORKSPACE_PERMISSIONS"
)

// All allowed values of PermissionedFeatureOrObject enum
var AllowedPermissionedFeatureOrObjectEnumValues = []PermissionedFeatureOrObject{
	"ADMIN_SEARCH",
	"ANNOUNCEMENTS",
	"ANSWERS",
	"COLLECTIONS",
	"COLLECTION_ITEMS",
	"INSIGHTS_COLLECTIONS",
	"INSIGHTS_CONTENT",
	"INSIGHTS_INVITE",
	"INSIGHTS_SEARCHES",
	"INSIGHTS_TEAMMATES",
	"PINS",
	"SENSITIVE_CONTENT_REPORTS",
	"SHORTCUTS",
	"TEAMS",
	"VERIFICATIONS",
	"WORKSPACE_APPS",
	"WORKSPACE_PERMISSIONS",
}

func (v *PermissionedFeatureOrObject) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PermissionedFeatureOrObject(value)
	for _, existing := range AllowedPermissionedFeatureOrObjectEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PermissionedFeatureOrObject", value)
}

// NewPermissionedFeatureOrObjectFromValue returns a pointer to a valid PermissionedFeatureOrObject
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPermissionedFeatureOrObjectFromValue(v string) (*PermissionedFeatureOrObject, error) {
	ev := PermissionedFeatureOrObject(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PermissionedFeatureOrObject: valid values are %v", v, AllowedPermissionedFeatureOrObjectEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PermissionedFeatureOrObject) IsValid() bool {
	for _, existing := range AllowedPermissionedFeatureOrObjectEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PermissionedFeatureOrObject value
func (v PermissionedFeatureOrObject) Ptr() *PermissionedFeatureOrObject {
	return &v
}

type NullablePermissionedFeatureOrObject struct {
	value *PermissionedFeatureOrObject
	isSet bool
}

func (v NullablePermissionedFeatureOrObject) Get() *PermissionedFeatureOrObject {
	return v.value
}

func (v *NullablePermissionedFeatureOrObject) Set(val *PermissionedFeatureOrObject) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissionedFeatureOrObject) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissionedFeatureOrObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissionedFeatureOrObject(val *PermissionedFeatureOrObject) *NullablePermissionedFeatureOrObject {
	return &NullablePermissionedFeatureOrObject{value: val, isSet: true}
}

func (v NullablePermissionedFeatureOrObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissionedFeatureOrObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

