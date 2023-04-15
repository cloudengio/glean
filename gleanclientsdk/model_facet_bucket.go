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

// checks if the FacetBucket type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FacetBucket{}

// FacetBucket struct for FacetBucket
type FacetBucket struct {
	// Estimated number of results in this facet.
	Count *int32 `json:"count,omitempty"`
	// The datasource the value belongs to. This will be used by the all tab to show types across all datasources.
	Datasource *string `json:"datasource,omitempty"`
	// Estimated percentage of results in this facet.
	Percentage *int32 `json:"percentage,omitempty"`
	Value *FacetValue `json:"value,omitempty"`
}

// NewFacetBucket instantiates a new FacetBucket object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFacetBucket() *FacetBucket {
	this := FacetBucket{}
	return &this
}

// NewFacetBucketWithDefaults instantiates a new FacetBucket object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFacetBucketWithDefaults() *FacetBucket {
	this := FacetBucket{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *FacetBucket) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FacetBucket) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *FacetBucket) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *FacetBucket) SetCount(v int32) {
	o.Count = &v
}

// GetDatasource returns the Datasource field value if set, zero value otherwise.
func (o *FacetBucket) GetDatasource() string {
	if o == nil || IsNil(o.Datasource) {
		var ret string
		return ret
	}
	return *o.Datasource
}

// GetDatasourceOk returns a tuple with the Datasource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FacetBucket) GetDatasourceOk() (*string, bool) {
	if o == nil || IsNil(o.Datasource) {
		return nil, false
	}
	return o.Datasource, true
}

// HasDatasource returns a boolean if a field has been set.
func (o *FacetBucket) HasDatasource() bool {
	if o != nil && !IsNil(o.Datasource) {
		return true
	}

	return false
}

// SetDatasource gets a reference to the given string and assigns it to the Datasource field.
func (o *FacetBucket) SetDatasource(v string) {
	o.Datasource = &v
}

// GetPercentage returns the Percentage field value if set, zero value otherwise.
func (o *FacetBucket) GetPercentage() int32 {
	if o == nil || IsNil(o.Percentage) {
		var ret int32
		return ret
	}
	return *o.Percentage
}

// GetPercentageOk returns a tuple with the Percentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FacetBucket) GetPercentageOk() (*int32, bool) {
	if o == nil || IsNil(o.Percentage) {
		return nil, false
	}
	return o.Percentage, true
}

// HasPercentage returns a boolean if a field has been set.
func (o *FacetBucket) HasPercentage() bool {
	if o != nil && !IsNil(o.Percentage) {
		return true
	}

	return false
}

// SetPercentage gets a reference to the given int32 and assigns it to the Percentage field.
func (o *FacetBucket) SetPercentage(v int32) {
	o.Percentage = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *FacetBucket) GetValue() FacetValue {
	if o == nil || IsNil(o.Value) {
		var ret FacetValue
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FacetBucket) GetValueOk() (*FacetValue, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *FacetBucket) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given FacetValue and assigns it to the Value field.
func (o *FacetBucket) SetValue(v FacetValue) {
	o.Value = &v
}

func (o FacetBucket) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FacetBucket) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Datasource) {
		toSerialize["datasource"] = o.Datasource
	}
	if !IsNil(o.Percentage) {
		toSerialize["percentage"] = o.Percentage
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableFacetBucket struct {
	value *FacetBucket
	isSet bool
}

func (v NullableFacetBucket) Get() *FacetBucket {
	return v.value
}

func (v *NullableFacetBucket) Set(val *FacetBucket) {
	v.value = val
	v.isSet = true
}

func (v NullableFacetBucket) IsSet() bool {
	return v.isSet
}

func (v *NullableFacetBucket) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFacetBucket(val *FacetBucket) *NullableFacetBucket {
	return &NullableFacetBucket{value: val, isSet: true}
}

func (v NullableFacetBucket) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFacetBucket) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


