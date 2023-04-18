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

// ClusterTypeEnum The reason for inclusion of clusteredResults.
type ClusterTypeEnum string

// List of ClusterTypeEnum
const (
	SIMILAR ClusterTypeEnum = "SIMILAR"
	FRESHNESS ClusterTypeEnum = "FRESHNESS"
	TITLE ClusterTypeEnum = "TITLE"
	CONTENT ClusterTypeEnum = "CONTENT"
	NONE ClusterTypeEnum = "NONE"
	THREAD_REPLY ClusterTypeEnum = "THREAD_REPLY"
	THREAD_ROOT ClusterTypeEnum = "THREAD_ROOT"
)

// All allowed values of ClusterTypeEnum enum
var AllowedClusterTypeEnumEnumValues = []ClusterTypeEnum{
	"SIMILAR",
	"FRESHNESS",
	"TITLE",
	"CONTENT",
	"NONE",
	"THREAD_REPLY",
	"THREAD_ROOT",
}

func (v *ClusterTypeEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ClusterTypeEnum(value)
	for _, existing := range AllowedClusterTypeEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ClusterTypeEnum", value)
}

// NewClusterTypeEnumFromValue returns a pointer to a valid ClusterTypeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClusterTypeEnumFromValue(v string) (*ClusterTypeEnum, error) {
	ev := ClusterTypeEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClusterTypeEnum: valid values are %v", v, AllowedClusterTypeEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClusterTypeEnum) IsValid() bool {
	for _, existing := range AllowedClusterTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ClusterTypeEnum value
func (v ClusterTypeEnum) Ptr() *ClusterTypeEnum {
	return &v
}

type NullableClusterTypeEnum struct {
	value *ClusterTypeEnum
	isSet bool
}

func (v NullableClusterTypeEnum) Get() *ClusterTypeEnum {
	return v.value
}

func (v *NullableClusterTypeEnum) Set(val *ClusterTypeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterTypeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterTypeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterTypeEnum(val *ClusterTypeEnum) *NullableClusterTypeEnum {
	return &NullableClusterTypeEnum{value: val, isSet: true}
}

func (v NullableClusterTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterTypeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

