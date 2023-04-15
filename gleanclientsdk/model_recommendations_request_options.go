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

// checks if the RecommendationsRequestOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecommendationsRequestOptions{}

// RecommendationsRequestOptions struct for RecommendationsRequestOptions
type RecommendationsRequestOptions struct {
	// Filter results to a single datasource name (e.g. gmail, slack). All results are returned if missing.
	DatasourceFilter *string `json:"datasourceFilter,omitempty"`
	// Identifier for what JSON schema to expect for `context`.
	ContextType *string `json:"contextType,omitempty"`
	// JSON containing client context like case subject input.  Expects schema based on `contextType`.
	Context *string `json:"context,omitempty"`
	// The types of prominence wanted in results returned. Default is any type.
	ResultProminence []SearchResultProminenceEnum `json:"resultProminence,omitempty"`
}

// NewRecommendationsRequestOptions instantiates a new RecommendationsRequestOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecommendationsRequestOptions() *RecommendationsRequestOptions {
	this := RecommendationsRequestOptions{}
	return &this
}

// NewRecommendationsRequestOptionsWithDefaults instantiates a new RecommendationsRequestOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecommendationsRequestOptionsWithDefaults() *RecommendationsRequestOptions {
	this := RecommendationsRequestOptions{}
	return &this
}

// GetDatasourceFilter returns the DatasourceFilter field value if set, zero value otherwise.
func (o *RecommendationsRequestOptions) GetDatasourceFilter() string {
	if o == nil || IsNil(o.DatasourceFilter) {
		var ret string
		return ret
	}
	return *o.DatasourceFilter
}

// GetDatasourceFilterOk returns a tuple with the DatasourceFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationsRequestOptions) GetDatasourceFilterOk() (*string, bool) {
	if o == nil || IsNil(o.DatasourceFilter) {
		return nil, false
	}
	return o.DatasourceFilter, true
}

// HasDatasourceFilter returns a boolean if a field has been set.
func (o *RecommendationsRequestOptions) HasDatasourceFilter() bool {
	if o != nil && !IsNil(o.DatasourceFilter) {
		return true
	}

	return false
}

// SetDatasourceFilter gets a reference to the given string and assigns it to the DatasourceFilter field.
func (o *RecommendationsRequestOptions) SetDatasourceFilter(v string) {
	o.DatasourceFilter = &v
}

// GetContextType returns the ContextType field value if set, zero value otherwise.
func (o *RecommendationsRequestOptions) GetContextType() string {
	if o == nil || IsNil(o.ContextType) {
		var ret string
		return ret
	}
	return *o.ContextType
}

// GetContextTypeOk returns a tuple with the ContextType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationsRequestOptions) GetContextTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ContextType) {
		return nil, false
	}
	return o.ContextType, true
}

// HasContextType returns a boolean if a field has been set.
func (o *RecommendationsRequestOptions) HasContextType() bool {
	if o != nil && !IsNil(o.ContextType) {
		return true
	}

	return false
}

// SetContextType gets a reference to the given string and assigns it to the ContextType field.
func (o *RecommendationsRequestOptions) SetContextType(v string) {
	o.ContextType = &v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *RecommendationsRequestOptions) GetContext() string {
	if o == nil || IsNil(o.Context) {
		var ret string
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationsRequestOptions) GetContextOk() (*string, bool) {
	if o == nil || IsNil(o.Context) {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *RecommendationsRequestOptions) HasContext() bool {
	if o != nil && !IsNil(o.Context) {
		return true
	}

	return false
}

// SetContext gets a reference to the given string and assigns it to the Context field.
func (o *RecommendationsRequestOptions) SetContext(v string) {
	o.Context = &v
}

// GetResultProminence returns the ResultProminence field value if set, zero value otherwise.
func (o *RecommendationsRequestOptions) GetResultProminence() []SearchResultProminenceEnum {
	if o == nil || IsNil(o.ResultProminence) {
		var ret []SearchResultProminenceEnum
		return ret
	}
	return o.ResultProminence
}

// GetResultProminenceOk returns a tuple with the ResultProminence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationsRequestOptions) GetResultProminenceOk() ([]SearchResultProminenceEnum, bool) {
	if o == nil || IsNil(o.ResultProminence) {
		return nil, false
	}
	return o.ResultProminence, true
}

// HasResultProminence returns a boolean if a field has been set.
func (o *RecommendationsRequestOptions) HasResultProminence() bool {
	if o != nil && !IsNil(o.ResultProminence) {
		return true
	}

	return false
}

// SetResultProminence gets a reference to the given []SearchResultProminenceEnum and assigns it to the ResultProminence field.
func (o *RecommendationsRequestOptions) SetResultProminence(v []SearchResultProminenceEnum) {
	o.ResultProminence = v
}

func (o RecommendationsRequestOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecommendationsRequestOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DatasourceFilter) {
		toSerialize["datasourceFilter"] = o.DatasourceFilter
	}
	if !IsNil(o.ContextType) {
		toSerialize["contextType"] = o.ContextType
	}
	if !IsNil(o.Context) {
		toSerialize["context"] = o.Context
	}
	if !IsNil(o.ResultProminence) {
		toSerialize["resultProminence"] = o.ResultProminence
	}
	return toSerialize, nil
}

type NullableRecommendationsRequestOptions struct {
	value *RecommendationsRequestOptions
	isSet bool
}

func (v NullableRecommendationsRequestOptions) Get() *RecommendationsRequestOptions {
	return v.value
}

func (v *NullableRecommendationsRequestOptions) Set(val *RecommendationsRequestOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableRecommendationsRequestOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableRecommendationsRequestOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecommendationsRequestOptions(val *RecommendationsRequestOptions) *NullableRecommendationsRequestOptions {
	return &NullableRecommendationsRequestOptions{value: val, isSet: true}
}

func (v NullableRecommendationsRequestOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecommendationsRequestOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


