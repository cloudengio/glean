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

// checks if the ConferenceData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConferenceData{}

// ConferenceData struct for ConferenceData
type ConferenceData struct {
	Provider string `json:"provider"`
	// A permalink for the conference.
	Uri string `json:"uri"`
	Source *string `json:"source,omitempty"`
}

// NewConferenceData instantiates a new ConferenceData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConferenceData(provider string, uri string) *ConferenceData {
	this := ConferenceData{}
	this.Provider = provider
	this.Uri = uri
	return &this
}

// NewConferenceDataWithDefaults instantiates a new ConferenceData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConferenceDataWithDefaults() *ConferenceData {
	this := ConferenceData{}
	return &this
}

// GetProvider returns the Provider field value
func (o *ConferenceData) GetProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *ConferenceData) GetProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *ConferenceData) SetProvider(v string) {
	o.Provider = v
}

// GetUri returns the Uri field value
func (o *ConferenceData) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *ConferenceData) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *ConferenceData) SetUri(v string) {
	o.Uri = v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *ConferenceData) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConferenceData) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *ConferenceData) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *ConferenceData) SetSource(v string) {
	o.Source = &v
}

func (o ConferenceData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConferenceData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["provider"] = o.Provider
	toSerialize["uri"] = o.Uri
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	return toSerialize, nil
}

type NullableConferenceData struct {
	value *ConferenceData
	isSet bool
}

func (v NullableConferenceData) Get() *ConferenceData {
	return v.value
}

func (v *NullableConferenceData) Set(val *ConferenceData) {
	v.value = val
	v.isSet = true
}

func (v NullableConferenceData) IsSet() bool {
	return v.isSet
}

func (v *NullableConferenceData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConferenceData(val *ConferenceData) *NullableConferenceData {
	return &NullableConferenceData{value: val, isSet: true}
}

func (v NullableConferenceData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConferenceData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


