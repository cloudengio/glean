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

// checks if the GetDocumentAnalyticsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDocumentAnalyticsRequest{}

// GetDocumentAnalyticsRequest struct for GetDocumentAnalyticsRequest
type GetDocumentAnalyticsRequest struct {
	// The specification for the documents for which analytics will be retrieved.
	DocumentSpecs []DocumentSpec `json:"documentSpecs"`
	DayRange Period `json:"dayRange"`
	// Whether response should include click information or not. Default is to not include click information.
	WithClickerCounts *bool `json:"withClickerCounts,omitempty"`
	// Whether the results will include aggregate counts/info for facets like location, department, etc.
	WithFacetAggregations *bool `json:"withFacetAggregations,omitempty"`
	// Whether response should include visit counts or not. Default is to return only visitor counts.
	WithVisitCounts *bool `json:"withVisitCounts,omitempty"`
}

// NewGetDocumentAnalyticsRequest instantiates a new GetDocumentAnalyticsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDocumentAnalyticsRequest(documentSpecs []DocumentSpec, dayRange Period) *GetDocumentAnalyticsRequest {
	this := GetDocumentAnalyticsRequest{}
	this.DocumentSpecs = documentSpecs
	this.DayRange = dayRange
	return &this
}

// NewGetDocumentAnalyticsRequestWithDefaults instantiates a new GetDocumentAnalyticsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDocumentAnalyticsRequestWithDefaults() *GetDocumentAnalyticsRequest {
	this := GetDocumentAnalyticsRequest{}
	return &this
}

// GetDocumentSpecs returns the DocumentSpecs field value
func (o *GetDocumentAnalyticsRequest) GetDocumentSpecs() []DocumentSpec {
	if o == nil {
		var ret []DocumentSpec
		return ret
	}

	return o.DocumentSpecs
}

// GetDocumentSpecsOk returns a tuple with the DocumentSpecs field value
// and a boolean to check if the value has been set.
func (o *GetDocumentAnalyticsRequest) GetDocumentSpecsOk() ([]DocumentSpec, bool) {
	if o == nil {
		return nil, false
	}
	return o.DocumentSpecs, true
}

// SetDocumentSpecs sets field value
func (o *GetDocumentAnalyticsRequest) SetDocumentSpecs(v []DocumentSpec) {
	o.DocumentSpecs = v
}

// GetDayRange returns the DayRange field value
func (o *GetDocumentAnalyticsRequest) GetDayRange() Period {
	if o == nil {
		var ret Period
		return ret
	}

	return o.DayRange
}

// GetDayRangeOk returns a tuple with the DayRange field value
// and a boolean to check if the value has been set.
func (o *GetDocumentAnalyticsRequest) GetDayRangeOk() (*Period, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DayRange, true
}

// SetDayRange sets field value
func (o *GetDocumentAnalyticsRequest) SetDayRange(v Period) {
	o.DayRange = v
}

// GetWithClickerCounts returns the WithClickerCounts field value if set, zero value otherwise.
func (o *GetDocumentAnalyticsRequest) GetWithClickerCounts() bool {
	if o == nil || IsNil(o.WithClickerCounts) {
		var ret bool
		return ret
	}
	return *o.WithClickerCounts
}

// GetWithClickerCountsOk returns a tuple with the WithClickerCounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDocumentAnalyticsRequest) GetWithClickerCountsOk() (*bool, bool) {
	if o == nil || IsNil(o.WithClickerCounts) {
		return nil, false
	}
	return o.WithClickerCounts, true
}

// HasWithClickerCounts returns a boolean if a field has been set.
func (o *GetDocumentAnalyticsRequest) HasWithClickerCounts() bool {
	if o != nil && !IsNil(o.WithClickerCounts) {
		return true
	}

	return false
}

// SetWithClickerCounts gets a reference to the given bool and assigns it to the WithClickerCounts field.
func (o *GetDocumentAnalyticsRequest) SetWithClickerCounts(v bool) {
	o.WithClickerCounts = &v
}

// GetWithFacetAggregations returns the WithFacetAggregations field value if set, zero value otherwise.
func (o *GetDocumentAnalyticsRequest) GetWithFacetAggregations() bool {
	if o == nil || IsNil(o.WithFacetAggregations) {
		var ret bool
		return ret
	}
	return *o.WithFacetAggregations
}

// GetWithFacetAggregationsOk returns a tuple with the WithFacetAggregations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDocumentAnalyticsRequest) GetWithFacetAggregationsOk() (*bool, bool) {
	if o == nil || IsNil(o.WithFacetAggregations) {
		return nil, false
	}
	return o.WithFacetAggregations, true
}

// HasWithFacetAggregations returns a boolean if a field has been set.
func (o *GetDocumentAnalyticsRequest) HasWithFacetAggregations() bool {
	if o != nil && !IsNil(o.WithFacetAggregations) {
		return true
	}

	return false
}

// SetWithFacetAggregations gets a reference to the given bool and assigns it to the WithFacetAggregations field.
func (o *GetDocumentAnalyticsRequest) SetWithFacetAggregations(v bool) {
	o.WithFacetAggregations = &v
}

// GetWithVisitCounts returns the WithVisitCounts field value if set, zero value otherwise.
func (o *GetDocumentAnalyticsRequest) GetWithVisitCounts() bool {
	if o == nil || IsNil(o.WithVisitCounts) {
		var ret bool
		return ret
	}
	return *o.WithVisitCounts
}

// GetWithVisitCountsOk returns a tuple with the WithVisitCounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDocumentAnalyticsRequest) GetWithVisitCountsOk() (*bool, bool) {
	if o == nil || IsNil(o.WithVisitCounts) {
		return nil, false
	}
	return o.WithVisitCounts, true
}

// HasWithVisitCounts returns a boolean if a field has been set.
func (o *GetDocumentAnalyticsRequest) HasWithVisitCounts() bool {
	if o != nil && !IsNil(o.WithVisitCounts) {
		return true
	}

	return false
}

// SetWithVisitCounts gets a reference to the given bool and assigns it to the WithVisitCounts field.
func (o *GetDocumentAnalyticsRequest) SetWithVisitCounts(v bool) {
	o.WithVisitCounts = &v
}

func (o GetDocumentAnalyticsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDocumentAnalyticsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["documentSpecs"] = o.DocumentSpecs
	toSerialize["dayRange"] = o.DayRange
	if !IsNil(o.WithClickerCounts) {
		toSerialize["withClickerCounts"] = o.WithClickerCounts
	}
	if !IsNil(o.WithFacetAggregations) {
		toSerialize["withFacetAggregations"] = o.WithFacetAggregations
	}
	if !IsNil(o.WithVisitCounts) {
		toSerialize["withVisitCounts"] = o.WithVisitCounts
	}
	return toSerialize, nil
}

type NullableGetDocumentAnalyticsRequest struct {
	value *GetDocumentAnalyticsRequest
	isSet bool
}

func (v NullableGetDocumentAnalyticsRequest) Get() *GetDocumentAnalyticsRequest {
	return v.value
}

func (v *NullableGetDocumentAnalyticsRequest) Set(val *GetDocumentAnalyticsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDocumentAnalyticsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDocumentAnalyticsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDocumentAnalyticsRequest(val *GetDocumentAnalyticsRequest) *NullableGetDocumentAnalyticsRequest {
	return &NullableGetDocumentAnalyticsRequest{value: val, isSet: true}
}

func (v NullableGetDocumentAnalyticsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDocumentAnalyticsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


