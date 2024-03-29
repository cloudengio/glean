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
)

// checks if the CustomerMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerMetadata{}

// CustomerMetadata struct for CustomerMetadata
type CustomerMetadata struct {
	// The user visible id of the salesforce customer account.
	DatasourceId *string `json:"datasourceId,omitempty"`
}

// NewCustomerMetadata instantiates a new CustomerMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerMetadata() *CustomerMetadata {
	this := CustomerMetadata{}
	return &this
}

// NewCustomerMetadataWithDefaults instantiates a new CustomerMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerMetadataWithDefaults() *CustomerMetadata {
	this := CustomerMetadata{}
	return &this
}

// GetDatasourceId returns the DatasourceId field value if set, zero value otherwise.
func (o *CustomerMetadata) GetDatasourceId() string {
	if o == nil || IsNil(o.DatasourceId) {
		var ret string
		return ret
	}
	return *o.DatasourceId
}

// GetDatasourceIdOk returns a tuple with the DatasourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerMetadata) GetDatasourceIdOk() (*string, bool) {
	if o == nil || IsNil(o.DatasourceId) {
		return nil, false
	}
	return o.DatasourceId, true
}

// HasDatasourceId returns a boolean if a field has been set.
func (o *CustomerMetadata) HasDatasourceId() bool {
	if o != nil && !IsNil(o.DatasourceId) {
		return true
	}

	return false
}

// SetDatasourceId gets a reference to the given string and assigns it to the DatasourceId field.
func (o *CustomerMetadata) SetDatasourceId(v string) {
	o.DatasourceId = &v
}

func (o CustomerMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DatasourceId) {
		toSerialize["datasourceId"] = o.DatasourceId
	}
	return toSerialize, nil
}

type NullableCustomerMetadata struct {
	value *CustomerMetadata
	isSet bool
}

func (v NullableCustomerMetadata) Get() *CustomerMetadata {
	return v.value
}

func (v *NullableCustomerMetadata) Set(val *CustomerMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerMetadata(val *CustomerMetadata) *NullableCustomerMetadata {
	return &NullableCustomerMetadata{value: val, isSet: true}
}

func (v NullableCustomerMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


