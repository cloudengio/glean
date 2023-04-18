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

// checks if the PinDocumentMutableProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PinDocumentMutableProperties{}

// PinDocumentMutableProperties struct for PinDocumentMutableProperties
type PinDocumentMutableProperties struct {
	// DEPRECATED - The query string to be set for the pin
	Query *string `json:"query,omitempty"`
	// The query strings for which the pinned result will show.
	Queries []string `json:"queries,omitempty"`
	// Filters which restrict who should see the pinned document. Values are taken from the corresponding filters in people search.
	AudienceFilters []FacetFilter `json:"audienceFilters,omitempty"`
}

// NewPinDocumentMutableProperties instantiates a new PinDocumentMutableProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPinDocumentMutableProperties() *PinDocumentMutableProperties {
	this := PinDocumentMutableProperties{}
	return &this
}

// NewPinDocumentMutablePropertiesWithDefaults instantiates a new PinDocumentMutableProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPinDocumentMutablePropertiesWithDefaults() *PinDocumentMutableProperties {
	this := PinDocumentMutableProperties{}
	return &this
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *PinDocumentMutableProperties) GetQuery() string {
	if o == nil || IsNil(o.Query) {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PinDocumentMutableProperties) GetQueryOk() (*string, bool) {
	if o == nil || IsNil(o.Query) {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *PinDocumentMutableProperties) HasQuery() bool {
	if o != nil && !IsNil(o.Query) {
		return true
	}

	return false
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *PinDocumentMutableProperties) SetQuery(v string) {
	o.Query = &v
}

// GetQueries returns the Queries field value if set, zero value otherwise.
func (o *PinDocumentMutableProperties) GetQueries() []string {
	if o == nil || IsNil(o.Queries) {
		var ret []string
		return ret
	}
	return o.Queries
}

// GetQueriesOk returns a tuple with the Queries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PinDocumentMutableProperties) GetQueriesOk() ([]string, bool) {
	if o == nil || IsNil(o.Queries) {
		return nil, false
	}
	return o.Queries, true
}

// HasQueries returns a boolean if a field has been set.
func (o *PinDocumentMutableProperties) HasQueries() bool {
	if o != nil && !IsNil(o.Queries) {
		return true
	}

	return false
}

// SetQueries gets a reference to the given []string and assigns it to the Queries field.
func (o *PinDocumentMutableProperties) SetQueries(v []string) {
	o.Queries = v
}

// GetAudienceFilters returns the AudienceFilters field value if set, zero value otherwise.
func (o *PinDocumentMutableProperties) GetAudienceFilters() []FacetFilter {
	if o == nil || IsNil(o.AudienceFilters) {
		var ret []FacetFilter
		return ret
	}
	return o.AudienceFilters
}

// GetAudienceFiltersOk returns a tuple with the AudienceFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PinDocumentMutableProperties) GetAudienceFiltersOk() ([]FacetFilter, bool) {
	if o == nil || IsNil(o.AudienceFilters) {
		return nil, false
	}
	return o.AudienceFilters, true
}

// HasAudienceFilters returns a boolean if a field has been set.
func (o *PinDocumentMutableProperties) HasAudienceFilters() bool {
	if o != nil && !IsNil(o.AudienceFilters) {
		return true
	}

	return false
}

// SetAudienceFilters gets a reference to the given []FacetFilter and assigns it to the AudienceFilters field.
func (o *PinDocumentMutableProperties) SetAudienceFilters(v []FacetFilter) {
	o.AudienceFilters = v
}

func (o PinDocumentMutableProperties) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PinDocumentMutableProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Query) {
		toSerialize["query"] = o.Query
	}
	if !IsNil(o.Queries) {
		toSerialize["queries"] = o.Queries
	}
	if !IsNil(o.AudienceFilters) {
		toSerialize["audienceFilters"] = o.AudienceFilters
	}
	return toSerialize, nil
}

type NullablePinDocumentMutableProperties struct {
	value *PinDocumentMutableProperties
	isSet bool
}

func (v NullablePinDocumentMutableProperties) Get() *PinDocumentMutableProperties {
	return v.value
}

func (v *NullablePinDocumentMutableProperties) Set(val *PinDocumentMutableProperties) {
	v.value = val
	v.isSet = true
}

func (v NullablePinDocumentMutableProperties) IsSet() bool {
	return v.isSet
}

func (v *NullablePinDocumentMutableProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePinDocumentMutableProperties(val *PinDocumentMutableProperties) *NullablePinDocumentMutableProperties {
	return &NullablePinDocumentMutableProperties{value: val, isSet: true}
}

func (v NullablePinDocumentMutableProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePinDocumentMutableProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


