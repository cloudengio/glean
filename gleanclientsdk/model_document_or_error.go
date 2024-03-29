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

// DocumentOrError - struct for DocumentOrError
type DocumentOrError struct {
	Document *Document
	DocumentOrErrorOneOf *DocumentOrErrorOneOf
}

// DocumentAsDocumentOrError is a convenience function that returns Document wrapped in DocumentOrError
func DocumentAsDocumentOrError(v *Document) DocumentOrError {
	return DocumentOrError{
		Document: v,
	}
}

// DocumentOrErrorOneOfAsDocumentOrError is a convenience function that returns DocumentOrErrorOneOf wrapped in DocumentOrError
func DocumentOrErrorOneOfAsDocumentOrError(v *DocumentOrErrorOneOf) DocumentOrError {
	return DocumentOrError{
		DocumentOrErrorOneOf: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *DocumentOrError) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Document
	err = newStrictDecoder(data).Decode(&dst.Document)
	if err == nil {
		jsonDocument, _ := json.Marshal(dst.Document)
		if string(jsonDocument) == "{}" { // empty struct
			dst.Document = nil
		} else {
			match++
		}
	} else {
		dst.Document = nil
	}

	// try to unmarshal data into DocumentOrErrorOneOf
	err = newStrictDecoder(data).Decode(&dst.DocumentOrErrorOneOf)
	if err == nil {
		jsonDocumentOrErrorOneOf, _ := json.Marshal(dst.DocumentOrErrorOneOf)
		if string(jsonDocumentOrErrorOneOf) == "{}" { // empty struct
			dst.DocumentOrErrorOneOf = nil
		} else {
			match++
		}
	} else {
		dst.DocumentOrErrorOneOf = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Document = nil
		dst.DocumentOrErrorOneOf = nil

		return fmt.Errorf("data matches more than one schema in oneOf(DocumentOrError)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(DocumentOrError)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src DocumentOrError) MarshalJSON() ([]byte, error) {
	if src.Document != nil {
		return json.Marshal(&src.Document)
	}

	if src.DocumentOrErrorOneOf != nil {
		return json.Marshal(&src.DocumentOrErrorOneOf)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *DocumentOrError) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Document != nil {
		return obj.Document
	}

	if obj.DocumentOrErrorOneOf != nil {
		return obj.DocumentOrErrorOneOf
	}

	// all schemas are nil
	return nil
}

type NullableDocumentOrError struct {
	value *DocumentOrError
	isSet bool
}

func (v NullableDocumentOrError) Get() *DocumentOrError {
	return v.value
}

func (v *NullableDocumentOrError) Set(val *DocumentOrError) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentOrError) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentOrError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentOrError(val *DocumentOrError) *NullableDocumentOrError {
	return &NullableDocumentOrError{value: val, isSet: true}
}

func (v NullableDocumentOrError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentOrError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


