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
)

// checks if the AddCollectionItemsError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddCollectionItemsError{}

// AddCollectionItemsError struct for AddCollectionItemsError
type AddCollectionItemsError struct {
	ErrorType *string `json:"errorType,omitempty"`
}

// NewAddCollectionItemsError instantiates a new AddCollectionItemsError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddCollectionItemsError() *AddCollectionItemsError {
	this := AddCollectionItemsError{}
	return &this
}

// NewAddCollectionItemsErrorWithDefaults instantiates a new AddCollectionItemsError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddCollectionItemsErrorWithDefaults() *AddCollectionItemsError {
	this := AddCollectionItemsError{}
	return &this
}

// GetErrorType returns the ErrorType field value if set, zero value otherwise.
func (o *AddCollectionItemsError) GetErrorType() string {
	if o == nil || IsNil(o.ErrorType) {
		var ret string
		return ret
	}
	return *o.ErrorType
}

// GetErrorTypeOk returns a tuple with the ErrorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCollectionItemsError) GetErrorTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorType) {
		return nil, false
	}
	return o.ErrorType, true
}

// HasErrorType returns a boolean if a field has been set.
func (o *AddCollectionItemsError) HasErrorType() bool {
	if o != nil && !IsNil(o.ErrorType) {
		return true
	}

	return false
}

// SetErrorType gets a reference to the given string and assigns it to the ErrorType field.
func (o *AddCollectionItemsError) SetErrorType(v string) {
	o.ErrorType = &v
}

func (o AddCollectionItemsError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddCollectionItemsError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ErrorType) {
		toSerialize["errorType"] = o.ErrorType
	}
	return toSerialize, nil
}

type NullableAddCollectionItemsError struct {
	value *AddCollectionItemsError
	isSet bool
}

func (v NullableAddCollectionItemsError) Get() *AddCollectionItemsError {
	return v.value
}

func (v *NullableAddCollectionItemsError) Set(val *AddCollectionItemsError) {
	v.value = val
	v.isSet = true
}

func (v NullableAddCollectionItemsError) IsSet() bool {
	return v.isSet
}

func (v *NullableAddCollectionItemsError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddCollectionItemsError(val *AddCollectionItemsError) *NullableAddCollectionItemsError {
	return &NullableAddCollectionItemsError{value: val, isSet: true}
}

func (v NullableAddCollectionItemsError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddCollectionItemsError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


