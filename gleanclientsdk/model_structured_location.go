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

// checks if the StructuredLocation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StructuredLocation{}

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
	if o == nil || IsNil(o.DeskLocation) {
		var ret string
		return ret
	}
	return *o.DeskLocation
}

// GetDeskLocationOk returns a tuple with the DeskLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuredLocation) GetDeskLocationOk() (*string, bool) {
	if o == nil || IsNil(o.DeskLocation) {
		return nil, false
	}
	return o.DeskLocation, true
}

// HasDeskLocation returns a boolean if a field has been set.
func (o *StructuredLocation) HasDeskLocation() bool {
	if o != nil && !IsNil(o.DeskLocation) {
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
	if o == nil || IsNil(o.Timezone) {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuredLocation) GetTimezoneOk() (*string, bool) {
	if o == nil || IsNil(o.Timezone) {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *StructuredLocation) HasTimezone() bool {
	if o != nil && !IsNil(o.Timezone) {
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
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuredLocation) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *StructuredLocation) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
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
	if o == nil || IsNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuredLocation) GetCityOk() (*string, bool) {
	if o == nil || IsNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *StructuredLocation) HasCity() bool {
	if o != nil && !IsNil(o.City) {
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
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuredLocation) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *StructuredLocation) HasState() bool {
	if o != nil && !IsNil(o.State) {
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
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuredLocation) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *StructuredLocation) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
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
	if o == nil || IsNil(o.ZipCode) {
		var ret string
		return ret
	}
	return *o.ZipCode
}

// GetZipCodeOk returns a tuple with the ZipCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuredLocation) GetZipCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ZipCode) {
		return nil, false
	}
	return o.ZipCode, true
}

// HasZipCode returns a boolean if a field has been set.
func (o *StructuredLocation) HasZipCode() bool {
	if o != nil && !IsNil(o.ZipCode) {
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
	if o == nil || IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuredLocation) GetCountryOk() (*string, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *StructuredLocation) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
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
	if o == nil || IsNil(o.CountryCode) {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuredLocation) GetCountryCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CountryCode) {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *StructuredLocation) HasCountryCode() bool {
	if o != nil && !IsNil(o.CountryCode) {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *StructuredLocation) SetCountryCode(v string) {
	o.CountryCode = &v
}

func (o StructuredLocation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StructuredLocation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeskLocation) {
		toSerialize["deskLocation"] = o.DeskLocation
	}
	if !IsNil(o.Timezone) {
		toSerialize["timezone"] = o.Timezone
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.City) {
		toSerialize["city"] = o.City
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !IsNil(o.ZipCode) {
		toSerialize["zipCode"] = o.ZipCode
	}
	if !IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !IsNil(o.CountryCode) {
		toSerialize["countryCode"] = o.CountryCode
	}
	return toSerialize, nil
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


