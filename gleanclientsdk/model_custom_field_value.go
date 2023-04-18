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

// CustomFieldValue - struct for CustomFieldValue
type CustomFieldValue struct {
	CustomFieldValueHyperlink *CustomFieldValueHyperlink
	CustomFieldValuePerson *CustomFieldValuePerson
	CustomFieldValueStr *CustomFieldValueStr
}

// CustomFieldValueHyperlinkAsCustomFieldValue is a convenience function that returns CustomFieldValueHyperlink wrapped in CustomFieldValue
func CustomFieldValueHyperlinkAsCustomFieldValue(v *CustomFieldValueHyperlink) CustomFieldValue {
	return CustomFieldValue{
		CustomFieldValueHyperlink: v,
	}
}

// CustomFieldValuePersonAsCustomFieldValue is a convenience function that returns CustomFieldValuePerson wrapped in CustomFieldValue
func CustomFieldValuePersonAsCustomFieldValue(v *CustomFieldValuePerson) CustomFieldValue {
	return CustomFieldValue{
		CustomFieldValuePerson: v,
	}
}

// CustomFieldValueStrAsCustomFieldValue is a convenience function that returns CustomFieldValueStr wrapped in CustomFieldValue
func CustomFieldValueStrAsCustomFieldValue(v *CustomFieldValueStr) CustomFieldValue {
	return CustomFieldValue{
		CustomFieldValueStr: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *CustomFieldValue) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CustomFieldValueHyperlink
	err = newStrictDecoder(data).Decode(&dst.CustomFieldValueHyperlink)
	if err == nil {
		jsonCustomFieldValueHyperlink, _ := json.Marshal(dst.CustomFieldValueHyperlink)
		if string(jsonCustomFieldValueHyperlink) == "{}" { // empty struct
			dst.CustomFieldValueHyperlink = nil
		} else {
			match++
		}
	} else {
		dst.CustomFieldValueHyperlink = nil
	}

	// try to unmarshal data into CustomFieldValuePerson
	err = newStrictDecoder(data).Decode(&dst.CustomFieldValuePerson)
	if err == nil {
		jsonCustomFieldValuePerson, _ := json.Marshal(dst.CustomFieldValuePerson)
		if string(jsonCustomFieldValuePerson) == "{}" { // empty struct
			dst.CustomFieldValuePerson = nil
		} else {
			match++
		}
	} else {
		dst.CustomFieldValuePerson = nil
	}

	// try to unmarshal data into CustomFieldValueStr
	err = newStrictDecoder(data).Decode(&dst.CustomFieldValueStr)
	if err == nil {
		jsonCustomFieldValueStr, _ := json.Marshal(dst.CustomFieldValueStr)
		if string(jsonCustomFieldValueStr) == "{}" { // empty struct
			dst.CustomFieldValueStr = nil
		} else {
			match++
		}
	} else {
		dst.CustomFieldValueStr = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.CustomFieldValueHyperlink = nil
		dst.CustomFieldValuePerson = nil
		dst.CustomFieldValueStr = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CustomFieldValue)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CustomFieldValue)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CustomFieldValue) MarshalJSON() ([]byte, error) {
	if src.CustomFieldValueHyperlink != nil {
		return json.Marshal(&src.CustomFieldValueHyperlink)
	}

	if src.CustomFieldValuePerson != nil {
		return json.Marshal(&src.CustomFieldValuePerson)
	}

	if src.CustomFieldValueStr != nil {
		return json.Marshal(&src.CustomFieldValueStr)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CustomFieldValue) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.CustomFieldValueHyperlink != nil {
		return obj.CustomFieldValueHyperlink
	}

	if obj.CustomFieldValuePerson != nil {
		return obj.CustomFieldValuePerson
	}

	if obj.CustomFieldValueStr != nil {
		return obj.CustomFieldValueStr
	}

	// all schemas are nil
	return nil
}

type NullableCustomFieldValue struct {
	value *CustomFieldValue
	isSet bool
}

func (v NullableCustomFieldValue) Get() *CustomFieldValue {
	return v.value
}

func (v *NullableCustomFieldValue) Set(val *CustomFieldValue) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomFieldValue) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomFieldValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomFieldValue(val *CustomFieldValue) *NullableCustomFieldValue {
	return &NullableCustomFieldValue{value: val, isSet: true}
}

func (v NullableCustomFieldValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomFieldValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


