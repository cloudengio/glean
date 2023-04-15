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

// DocumentSpec - struct for DocumentSpec
type DocumentSpec struct {
	DocumentSpecOneOf *DocumentSpecOneOf
	DocumentSpecOneOf1 *DocumentSpecOneOf1
	DocumentSpecOneOf2 *DocumentSpecOneOf2
}

// DocumentSpecOneOfAsDocumentSpec is a convenience function that returns DocumentSpecOneOf wrapped in DocumentSpec
func DocumentSpecOneOfAsDocumentSpec(v *DocumentSpecOneOf) DocumentSpec {
	return DocumentSpec{
		DocumentSpecOneOf: v,
	}
}

// DocumentSpecOneOf1AsDocumentSpec is a convenience function that returns DocumentSpecOneOf1 wrapped in DocumentSpec
func DocumentSpecOneOf1AsDocumentSpec(v *DocumentSpecOneOf1) DocumentSpec {
	return DocumentSpec{
		DocumentSpecOneOf1: v,
	}
}

// DocumentSpecOneOf2AsDocumentSpec is a convenience function that returns DocumentSpecOneOf2 wrapped in DocumentSpec
func DocumentSpecOneOf2AsDocumentSpec(v *DocumentSpecOneOf2) DocumentSpec {
	return DocumentSpec{
		DocumentSpecOneOf2: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *DocumentSpec) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into DocumentSpecOneOf
	err = newStrictDecoder(data).Decode(&dst.DocumentSpecOneOf)
	if err == nil {
		jsonDocumentSpecOneOf, _ := json.Marshal(dst.DocumentSpecOneOf)
		if string(jsonDocumentSpecOneOf) == "{}" { // empty struct
			dst.DocumentSpecOneOf = nil
		} else {
			match++
		}
	} else {
		dst.DocumentSpecOneOf = nil
	}

	// try to unmarshal data into DocumentSpecOneOf1
	err = newStrictDecoder(data).Decode(&dst.DocumentSpecOneOf1)
	if err == nil {
		jsonDocumentSpecOneOf1, _ := json.Marshal(dst.DocumentSpecOneOf1)
		if string(jsonDocumentSpecOneOf1) == "{}" { // empty struct
			dst.DocumentSpecOneOf1 = nil
		} else {
			match++
		}
	} else {
		dst.DocumentSpecOneOf1 = nil
	}

	// try to unmarshal data into DocumentSpecOneOf2
	err = newStrictDecoder(data).Decode(&dst.DocumentSpecOneOf2)
	if err == nil {
		jsonDocumentSpecOneOf2, _ := json.Marshal(dst.DocumentSpecOneOf2)
		if string(jsonDocumentSpecOneOf2) == "{}" { // empty struct
			dst.DocumentSpecOneOf2 = nil
		} else {
			match++
		}
	} else {
		dst.DocumentSpecOneOf2 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.DocumentSpecOneOf = nil
		dst.DocumentSpecOneOf1 = nil
		dst.DocumentSpecOneOf2 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(DocumentSpec)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(DocumentSpec)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src DocumentSpec) MarshalJSON() ([]byte, error) {
	if src.DocumentSpecOneOf != nil {
		return json.Marshal(&src.DocumentSpecOneOf)
	}

	if src.DocumentSpecOneOf1 != nil {
		return json.Marshal(&src.DocumentSpecOneOf1)
	}

	if src.DocumentSpecOneOf2 != nil {
		return json.Marshal(&src.DocumentSpecOneOf2)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *DocumentSpec) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.DocumentSpecOneOf != nil {
		return obj.DocumentSpecOneOf
	}

	if obj.DocumentSpecOneOf1 != nil {
		return obj.DocumentSpecOneOf1
	}

	if obj.DocumentSpecOneOf2 != nil {
		return obj.DocumentSpecOneOf2
	}

	// all schemas are nil
	return nil
}

type NullableDocumentSpec struct {
	value *DocumentSpec
	isSet bool
}

func (v NullableDocumentSpec) Get() *DocumentSpec {
	return v.value
}

func (v *NullableDocumentSpec) Set(val *DocumentSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentSpec(val *DocumentSpec) *NullableDocumentSpec {
	return &NullableDocumentSpec{value: val, isSet: true}
}

func (v NullableDocumentSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


