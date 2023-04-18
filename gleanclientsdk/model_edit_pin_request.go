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

// checks if the EditPinRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EditPinRequest{}

// EditPinRequest struct for EditPinRequest
type EditPinRequest struct {
	// DEPRECATED - The query string to be set for the pin
	Query *string `json:"query,omitempty"`
	// The query strings for which the pinned result will show.
	Queries []string `json:"queries,omitempty"`
	// Filters which restrict who should see the pinned document. Values are taken from the corresponding filters in people search.
	AudienceFilters []FacetFilter `json:"audienceFilters,omitempty"`
	// DEPRECATED - Prefer use of `id`
	PinId *int32 `json:"pinId,omitempty"`
	// The opaque id of the pin to be edited
	Id *string `json:"id,omitempty"`
}

// NewEditPinRequest instantiates a new EditPinRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditPinRequest() *EditPinRequest {
	this := EditPinRequest{}
	return &this
}

// NewEditPinRequestWithDefaults instantiates a new EditPinRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditPinRequestWithDefaults() *EditPinRequest {
	this := EditPinRequest{}
	return &this
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *EditPinRequest) GetQuery() string {
	if o == nil || IsNil(o.Query) {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditPinRequest) GetQueryOk() (*string, bool) {
	if o == nil || IsNil(o.Query) {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *EditPinRequest) HasQuery() bool {
	if o != nil && !IsNil(o.Query) {
		return true
	}

	return false
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *EditPinRequest) SetQuery(v string) {
	o.Query = &v
}

// GetQueries returns the Queries field value if set, zero value otherwise.
func (o *EditPinRequest) GetQueries() []string {
	if o == nil || IsNil(o.Queries) {
		var ret []string
		return ret
	}
	return o.Queries
}

// GetQueriesOk returns a tuple with the Queries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditPinRequest) GetQueriesOk() ([]string, bool) {
	if o == nil || IsNil(o.Queries) {
		return nil, false
	}
	return o.Queries, true
}

// HasQueries returns a boolean if a field has been set.
func (o *EditPinRequest) HasQueries() bool {
	if o != nil && !IsNil(o.Queries) {
		return true
	}

	return false
}

// SetQueries gets a reference to the given []string and assigns it to the Queries field.
func (o *EditPinRequest) SetQueries(v []string) {
	o.Queries = v
}

// GetAudienceFilters returns the AudienceFilters field value if set, zero value otherwise.
func (o *EditPinRequest) GetAudienceFilters() []FacetFilter {
	if o == nil || IsNil(o.AudienceFilters) {
		var ret []FacetFilter
		return ret
	}
	return o.AudienceFilters
}

// GetAudienceFiltersOk returns a tuple with the AudienceFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditPinRequest) GetAudienceFiltersOk() ([]FacetFilter, bool) {
	if o == nil || IsNil(o.AudienceFilters) {
		return nil, false
	}
	return o.AudienceFilters, true
}

// HasAudienceFilters returns a boolean if a field has been set.
func (o *EditPinRequest) HasAudienceFilters() bool {
	if o != nil && !IsNil(o.AudienceFilters) {
		return true
	}

	return false
}

// SetAudienceFilters gets a reference to the given []FacetFilter and assigns it to the AudienceFilters field.
func (o *EditPinRequest) SetAudienceFilters(v []FacetFilter) {
	o.AudienceFilters = v
}

// GetPinId returns the PinId field value if set, zero value otherwise.
func (o *EditPinRequest) GetPinId() int32 {
	if o == nil || IsNil(o.PinId) {
		var ret int32
		return ret
	}
	return *o.PinId
}

// GetPinIdOk returns a tuple with the PinId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditPinRequest) GetPinIdOk() (*int32, bool) {
	if o == nil || IsNil(o.PinId) {
		return nil, false
	}
	return o.PinId, true
}

// HasPinId returns a boolean if a field has been set.
func (o *EditPinRequest) HasPinId() bool {
	if o != nil && !IsNil(o.PinId) {
		return true
	}

	return false
}

// SetPinId gets a reference to the given int32 and assigns it to the PinId field.
func (o *EditPinRequest) SetPinId(v int32) {
	o.PinId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EditPinRequest) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditPinRequest) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EditPinRequest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EditPinRequest) SetId(v string) {
	o.Id = &v
}

func (o EditPinRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EditPinRequest) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.PinId) {
		toSerialize["pinId"] = o.PinId
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableEditPinRequest struct {
	value *EditPinRequest
	isSet bool
}

func (v NullableEditPinRequest) Get() *EditPinRequest {
	return v.value
}

func (v *NullableEditPinRequest) Set(val *EditPinRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEditPinRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEditPinRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditPinRequest(val *EditPinRequest) *NullableEditPinRequest {
	return &NullableEditPinRequest{value: val, isSet: true}
}

func (v NullableEditPinRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditPinRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


