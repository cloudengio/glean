/*
Glean Indexing API

# Introduction In addition to the data sources that Glean has built-in support for, Glean also provides a REST API that enables customers to put arbitrary content in the search index. This is useful, for example, for doing permissions-aware search over content in internal tools that reside on-prem as well as for searching over applications that Glean does not currently support first class. In addition these APIs allow the customer to push organization data (people info, organization structure etc) into Glean.  # Early Access Please note that we are continually evolving the system so these APIs are considered “early access” and are subject to change. 

API version: 0.9.0
Contact: support@glean.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gleansdk

import (
	"encoding/json"
)

// StructuredLocation Detailed location with information about country, state, city etc.
type StructuredLocation struct {
	// Desk number.
	DeskLocation *string `json:"deskLocation,omitempty"`
	// Location's timezone, e.g. UTC, PST.
	Timezone *string `json:"timezone,omitempty"`
	// Office address or name.
	Address *string `json:"address,omitempty"`
	// Name of the city.
	City *string `json:"city,omitempty"`
	// State code.
	State *string `json:"state,omitempty"`
	// Region information, e.g. NORAM, APAC.
	Region *string `json:"region,omitempty"`
	// ZIP Code for the address.
	ZipCode *string `json:"zipCode,omitempty"`
	// Country name.
	Country *string `json:"country,omitempty"`
	// Alpha-2 or Alpha-3 ISO 3166 country code, e.g. US or USA.
	CountryCode *string `json:"countryCode,omitempty"`
}

// NewStructuredLocation instantiates a new StructuredLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStructuredLocation() *StructuredLocation {
	this := StructuredLocation{}
	return &this
}

// NewStructuredLocationWithDefaults instantiates a new StructuredLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStructuredLocationWithDefaults() *StructuredLocation {
	this := StructuredLocation{}
	return &this
}

// GetDeskLocation returns the DeskLocation field value if set, zero value otherwise.
func (o *StructuredLocation) GetDeskLocation() string {
	if o == nil || o.DeskLocation == nil {
		var ret string
		return ret
	}
	return *o.DeskLocation
}

// GetDeskLocationOk returns a tuple with the DeskLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuredLocation) GetDeskLocationOk() (*string, bool) {
	if o == nil || o.DeskLocation == nil {
		return nil, false
	}
	return o.DeskLocation, true
}

// HasDeskLocation returns a boolean if a field has been set.
func (o *StructuredLocation) HasDeskLocation() bool {
	if o != nil && o.DeskLocation != nil {
		return true
	}

	return false
}

// SetDeskLocation gets a reference to the given string and assigns it to the DeskLocation field.
func (o *StructuredLocation) SetDeskLocation(v string) {
	o.DeskLocation = &v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *StructuredLocation) GetTimezone() string {
	if o == nil || o.Timezone == nil {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuredLocation) GetTimezoneOk() (*string, bool) {
	if o == nil || o.Timezone == nil {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *StructuredLocation) HasTimezone() bool {
	if o != nil && o.Timezone != nil {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *StructuredLocation) SetTimezone(v string) {
	o.Timezone = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *StructuredLocation) GetAddress() string {
	if o == nil || o.Address == nil {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuredLocation) GetAddressOk() (*string, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *StructuredLocation) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *StructuredLocation) SetAddress(v string) {
	o.Address = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *StructuredLocation) GetCity() string {
	if o == nil || o.City == nil {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuredLocation) GetCityOk() (*string, bool) {
	if o == nil || o.City == nil {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *StructuredLocation) HasCity() bool {
	if o != nil && o.City != nil {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *StructuredLocation) SetCity(v string) {
	o.City = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *StructuredLocation) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuredLocation) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *StructuredLocation) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *StructuredLocation) SetState(v string) {
	o.State = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *StructuredLocation) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuredLocation) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *StructuredLocation) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *StructuredLocation) SetRegion(v string) {
	o.Region = &v
}

// GetZipCode returns the ZipCode field value if set, zero value otherwise.
func (o *StructuredLocation) GetZipCode() string {
	if o == nil || o.ZipCode == nil {
		var ret string
		return ret
	}
	return *o.ZipCode
}

// GetZipCodeOk returns a tuple with the ZipCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuredLocation) GetZipCodeOk() (*string, bool) {
	if o == nil || o.ZipCode == nil {
		return nil, false
	}
	return o.ZipCode, true
}

// HasZipCode returns a boolean if a field has been set.
func (o *StructuredLocation) HasZipCode() bool {
	if o != nil && o.ZipCode != nil {
		return true
	}

	return false
}

// SetZipCode gets a reference to the given string and assigns it to the ZipCode field.
func (o *StructuredLocation) SetZipCode(v string) {
	o.ZipCode = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *StructuredLocation) GetCountry() string {
	if o == nil || o.Country == nil {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuredLocation) GetCountryOk() (*string, bool) {
	if o == nil || o.Country == nil {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *StructuredLocation) HasCountry() bool {
	if o != nil && o.Country != nil {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *StructuredLocation) SetCountry(v string) {
	o.Country = &v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *StructuredLocation) GetCountryCode() string {
	if o == nil || o.CountryCode == nil {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuredLocation) GetCountryCodeOk() (*string, bool) {
	if o == nil || o.CountryCode == nil {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *StructuredLocation) HasCountryCode() bool {
	if o != nil && o.CountryCode != nil {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *StructuredLocation) SetCountryCode(v string) {
	o.CountryCode = &v
}

func (o StructuredLocation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeskLocation != nil {
		toSerialize["deskLocation"] = o.DeskLocation
	}
	if o.Timezone != nil {
		toSerialize["timezone"] = o.Timezone
	}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if o.City != nil {
		toSerialize["city"] = o.City
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	if o.ZipCode != nil {
		toSerialize["zipCode"] = o.ZipCode
	}
	if o.Country != nil {
		toSerialize["country"] = o.Country
	}
	if o.CountryCode != nil {
		toSerialize["countryCode"] = o.CountryCode
	}
	return json.Marshal(toSerialize)
}

type NullableStructuredLocation struct {
	value *StructuredLocation
	isSet bool
}

func (v NullableStructuredLocation) Get() *StructuredLocation {
	return v.value
}

func (v *NullableStructuredLocation) Set(val *StructuredLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableStructuredLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableStructuredLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStructuredLocation(val *StructuredLocation) *NullableStructuredLocation {
	return &NullableStructuredLocation{value: val, isSet: true}
}

func (v NullableStructuredLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStructuredLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


