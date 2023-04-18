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

// checks if the FacetResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FacetResult{}

// FacetResult struct for FacetResult
type FacetResult struct {
	// The source of this facet (e.g. container_name, doc_type, last_updated_at).
	SourceName *string `json:"sourceName,omitempty"`
	// How to display this facet. Currently supportes 'SelectSingle' and 'SelectMultiple'.
	OperatorName *string `json:"operatorName,omitempty"`
	// A list of unique buckets that exist within this result set.
	Buckets []FacetBucket `json:"buckets,omitempty"`
	// Returns true if more buckets exist than those returned. Additional buckets can be retrieve by requesting again with a higher facetBucketSize.
	HasMoreBuckets *bool `json:"hasMoreBuckets,omitempty"`
	// For most facets this will be the empty string, meaning the facet is high-level and applies to all documents for the datasource. When non-empty, this is used to group facets together (i.e. group facets for each doctype for a certain datasource)
	GroupName *string `json:"groupName,omitempty"`
}

// NewFacetResult instantiates a new FacetResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFacetResult() *FacetResult {
	this := FacetResult{}
	return &this
}

// NewFacetResultWithDefaults instantiates a new FacetResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFacetResultWithDefaults() *FacetResult {
	this := FacetResult{}
	return &this
}

// GetSourceName returns the SourceName field value if set, zero value otherwise.
func (o *FacetResult) GetSourceName() string {
	if o == nil || IsNil(o.SourceName) {
		var ret string
		return ret
	}
	return *o.SourceName
}

// GetSourceNameOk returns a tuple with the SourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FacetResult) GetSourceNameOk() (*string, bool) {
	if o == nil || IsNil(o.SourceName) {
		return nil, false
	}
	return o.SourceName, true
}

// HasSourceName returns a boolean if a field has been set.
func (o *FacetResult) HasSourceName() bool {
	if o != nil && !IsNil(o.SourceName) {
		return true
	}

	return false
}

// SetSourceName gets a reference to the given string and assigns it to the SourceName field.
func (o *FacetResult) SetSourceName(v string) {
	o.SourceName = &v
}

// GetOperatorName returns the OperatorName field value if set, zero value otherwise.
func (o *FacetResult) GetOperatorName() string {
	if o == nil || IsNil(o.OperatorName) {
		var ret string
		return ret
	}
	return *o.OperatorName
}

// GetOperatorNameOk returns a tuple with the OperatorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FacetResult) GetOperatorNameOk() (*string, bool) {
	if o == nil || IsNil(o.OperatorName) {
		return nil, false
	}
	return o.OperatorName, true
}

// HasOperatorName returns a boolean if a field has been set.
func (o *FacetResult) HasOperatorName() bool {
	if o != nil && !IsNil(o.OperatorName) {
		return true
	}

	return false
}

// SetOperatorName gets a reference to the given string and assigns it to the OperatorName field.
func (o *FacetResult) SetOperatorName(v string) {
	o.OperatorName = &v
}

// GetBuckets returns the Buckets field value if set, zero value otherwise.
func (o *FacetResult) GetBuckets() []FacetBucket {
	if o == nil || IsNil(o.Buckets) {
		var ret []FacetBucket
		return ret
	}
	return o.Buckets
}

// GetBucketsOk returns a tuple with the Buckets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FacetResult) GetBucketsOk() ([]FacetBucket, bool) {
	if o == nil || IsNil(o.Buckets) {
		return nil, false
	}
	return o.Buckets, true
}

// HasBuckets returns a boolean if a field has been set.
func (o *FacetResult) HasBuckets() bool {
	if o != nil && !IsNil(o.Buckets) {
		return true
	}

	return false
}

// SetBuckets gets a reference to the given []FacetBucket and assigns it to the Buckets field.
func (o *FacetResult) SetBuckets(v []FacetBucket) {
	o.Buckets = v
}

// GetHasMoreBuckets returns the HasMoreBuckets field value if set, zero value otherwise.
func (o *FacetResult) GetHasMoreBuckets() bool {
	if o == nil || IsNil(o.HasMoreBuckets) {
		var ret bool
		return ret
	}
	return *o.HasMoreBuckets
}

// GetHasMoreBucketsOk returns a tuple with the HasMoreBuckets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FacetResult) GetHasMoreBucketsOk() (*bool, bool) {
	if o == nil || IsNil(o.HasMoreBuckets) {
		return nil, false
	}
	return o.HasMoreBuckets, true
}

// HasHasMoreBuckets returns a boolean if a field has been set.
func (o *FacetResult) HasHasMoreBuckets() bool {
	if o != nil && !IsNil(o.HasMoreBuckets) {
		return true
	}

	return false
}

// SetHasMoreBuckets gets a reference to the given bool and assigns it to the HasMoreBuckets field.
func (o *FacetResult) SetHasMoreBuckets(v bool) {
	o.HasMoreBuckets = &v
}

// GetGroupName returns the GroupName field value if set, zero value otherwise.
func (o *FacetResult) GetGroupName() string {
	if o == nil || IsNil(o.GroupName) {
		var ret string
		return ret
	}
	return *o.GroupName
}

// GetGroupNameOk returns a tuple with the GroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FacetResult) GetGroupNameOk() (*string, bool) {
	if o == nil || IsNil(o.GroupName) {
		return nil, false
	}
	return o.GroupName, true
}

// HasGroupName returns a boolean if a field has been set.
func (o *FacetResult) HasGroupName() bool {
	if o != nil && !IsNil(o.GroupName) {
		return true
	}

	return false
}

// SetGroupName gets a reference to the given string and assigns it to the GroupName field.
func (o *FacetResult) SetGroupName(v string) {
	o.GroupName = &v
}

func (o FacetResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FacetResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SourceName) {
		toSerialize["sourceName"] = o.SourceName
	}
	if !IsNil(o.OperatorName) {
		toSerialize["operatorName"] = o.OperatorName
	}
	if !IsNil(o.Buckets) {
		toSerialize["buckets"] = o.Buckets
	}
	if !IsNil(o.HasMoreBuckets) {
		toSerialize["hasMoreBuckets"] = o.HasMoreBuckets
	}
	if !IsNil(o.GroupName) {
		toSerialize["groupName"] = o.GroupName
	}
	return toSerialize, nil
}

type NullableFacetResult struct {
	value *FacetResult
	isSet bool
}

func (v NullableFacetResult) Get() *FacetResult {
	return v.value
}

func (v *NullableFacetResult) Set(val *FacetResult) {
	v.value = val
	v.isSet = true
}

func (v NullableFacetResult) IsSet() bool {
	return v.isSet
}

func (v *NullableFacetResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFacetResult(val *FacetResult) *NullableFacetResult {
	return &NullableFacetResult{value: val, isSet: true}
}

func (v NullableFacetResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFacetResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

