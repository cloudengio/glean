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

// checks if the GetSimilarShortcutsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSimilarShortcutsRequest{}

// GetSimilarShortcutsRequest struct for GetSimilarShortcutsRequest
type GetSimilarShortcutsRequest struct {
	// Link text following go/ prefix.
	Alias string `json:"alias"`
}

// NewGetSimilarShortcutsRequest instantiates a new GetSimilarShortcutsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSimilarShortcutsRequest(alias string) *GetSimilarShortcutsRequest {
	this := GetSimilarShortcutsRequest{}
	this.Alias = alias
	return &this
}

// NewGetSimilarShortcutsRequestWithDefaults instantiates a new GetSimilarShortcutsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSimilarShortcutsRequestWithDefaults() *GetSimilarShortcutsRequest {
	this := GetSimilarShortcutsRequest{}
	return &this
}

// GetAlias returns the Alias field value
func (o *GetSimilarShortcutsRequest) GetAlias() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Alias
}

// GetAliasOk returns a tuple with the Alias field value
// and a boolean to check if the value has been set.
func (o *GetSimilarShortcutsRequest) GetAliasOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Alias, true
}

// SetAlias sets field value
func (o *GetSimilarShortcutsRequest) SetAlias(v string) {
	o.Alias = v
}

func (o GetSimilarShortcutsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSimilarShortcutsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["alias"] = o.Alias
	return toSerialize, nil
}

type NullableGetSimilarShortcutsRequest struct {
	value *GetSimilarShortcutsRequest
	isSet bool
}

func (v NullableGetSimilarShortcutsRequest) Get() *GetSimilarShortcutsRequest {
	return v.value
}

func (v *NullableGetSimilarShortcutsRequest) Set(val *GetSimilarShortcutsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSimilarShortcutsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSimilarShortcutsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSimilarShortcutsRequest(val *GetSimilarShortcutsRequest) *NullableGetSimilarShortcutsRequest {
	return &NullableGetSimilarShortcutsRequest{value: val, isSet: true}
}

func (v NullableGetSimilarShortcutsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSimilarShortcutsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

