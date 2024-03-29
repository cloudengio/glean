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

// checks if the FeedRequestOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeedRequestOptions{}

// FeedRequestOptions struct for FeedRequestOptions
type FeedRequestOptions struct {
	// Number of results asked in response. If a result is a collection, counts as one.
	ResultSize int32 `json:"resultSize"`
	// The offset of the client's timezone in minutes from UTC. e.g. PDT is -420 because it's 7 hours behind UTC.
	TimezoneOffset *int32 `json:"timezoneOffset,omitempty"`
	// Mapping from category to number of results asked for the category.
	CategoryToResultSize *map[string]FeedRequestOptionsCategoryToResultSizeValue `json:"categoryToResultSize,omitempty"`
	// Datasources for which content should be included. Empty is for all.
	DatasourceFilter []string `json:"datasourceFilter,omitempty"`
	// Auth tokens which may be used for federated retrieval of Feed entries.
	AuthTokens []AuthToken `json:"authTokens,omitempty"`
}

// NewFeedRequestOptions instantiates a new FeedRequestOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeedRequestOptions(resultSize int32) *FeedRequestOptions {
	this := FeedRequestOptions{}
	this.ResultSize = resultSize
	return &this
}

// NewFeedRequestOptionsWithDefaults instantiates a new FeedRequestOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeedRequestOptionsWithDefaults() *FeedRequestOptions {
	this := FeedRequestOptions{}
	return &this
}

// GetResultSize returns the ResultSize field value
func (o *FeedRequestOptions) GetResultSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ResultSize
}

// GetResultSizeOk returns a tuple with the ResultSize field value
// and a boolean to check if the value has been set.
func (o *FeedRequestOptions) GetResultSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResultSize, true
}

// SetResultSize sets field value
func (o *FeedRequestOptions) SetResultSize(v int32) {
	o.ResultSize = v
}

// GetTimezoneOffset returns the TimezoneOffset field value if set, zero value otherwise.
func (o *FeedRequestOptions) GetTimezoneOffset() int32 {
	if o == nil || IsNil(o.TimezoneOffset) {
		var ret int32
		return ret
	}
	return *o.TimezoneOffset
}

// GetTimezoneOffsetOk returns a tuple with the TimezoneOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedRequestOptions) GetTimezoneOffsetOk() (*int32, bool) {
	if o == nil || IsNil(o.TimezoneOffset) {
		return nil, false
	}
	return o.TimezoneOffset, true
}

// HasTimezoneOffset returns a boolean if a field has been set.
func (o *FeedRequestOptions) HasTimezoneOffset() bool {
	if o != nil && !IsNil(o.TimezoneOffset) {
		return true
	}

	return false
}

// SetTimezoneOffset gets a reference to the given int32 and assigns it to the TimezoneOffset field.
func (o *FeedRequestOptions) SetTimezoneOffset(v int32) {
	o.TimezoneOffset = &v
}

// GetCategoryToResultSize returns the CategoryToResultSize field value if set, zero value otherwise.
func (o *FeedRequestOptions) GetCategoryToResultSize() map[string]FeedRequestOptionsCategoryToResultSizeValue {
	if o == nil || IsNil(o.CategoryToResultSize) {
		var ret map[string]FeedRequestOptionsCategoryToResultSizeValue
		return ret
	}
	return *o.CategoryToResultSize
}

// GetCategoryToResultSizeOk returns a tuple with the CategoryToResultSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedRequestOptions) GetCategoryToResultSizeOk() (*map[string]FeedRequestOptionsCategoryToResultSizeValue, bool) {
	if o == nil || IsNil(o.CategoryToResultSize) {
		return nil, false
	}
	return o.CategoryToResultSize, true
}

// HasCategoryToResultSize returns a boolean if a field has been set.
func (o *FeedRequestOptions) HasCategoryToResultSize() bool {
	if o != nil && !IsNil(o.CategoryToResultSize) {
		return true
	}

	return false
}

// SetCategoryToResultSize gets a reference to the given map[string]FeedRequestOptionsCategoryToResultSizeValue and assigns it to the CategoryToResultSize field.
func (o *FeedRequestOptions) SetCategoryToResultSize(v map[string]FeedRequestOptionsCategoryToResultSizeValue) {
	o.CategoryToResultSize = &v
}

// GetDatasourceFilter returns the DatasourceFilter field value if set, zero value otherwise.
func (o *FeedRequestOptions) GetDatasourceFilter() []string {
	if o == nil || IsNil(o.DatasourceFilter) {
		var ret []string
		return ret
	}
	return o.DatasourceFilter
}

// GetDatasourceFilterOk returns a tuple with the DatasourceFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedRequestOptions) GetDatasourceFilterOk() ([]string, bool) {
	if o == nil || IsNil(o.DatasourceFilter) {
		return nil, false
	}
	return o.DatasourceFilter, true
}

// HasDatasourceFilter returns a boolean if a field has been set.
func (o *FeedRequestOptions) HasDatasourceFilter() bool {
	if o != nil && !IsNil(o.DatasourceFilter) {
		return true
	}

	return false
}

// SetDatasourceFilter gets a reference to the given []string and assigns it to the DatasourceFilter field.
func (o *FeedRequestOptions) SetDatasourceFilter(v []string) {
	o.DatasourceFilter = v
}

// GetAuthTokens returns the AuthTokens field value if set, zero value otherwise.
func (o *FeedRequestOptions) GetAuthTokens() []AuthToken {
	if o == nil || IsNil(o.AuthTokens) {
		var ret []AuthToken
		return ret
	}
	return o.AuthTokens
}

// GetAuthTokensOk returns a tuple with the AuthTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedRequestOptions) GetAuthTokensOk() ([]AuthToken, bool) {
	if o == nil || IsNil(o.AuthTokens) {
		return nil, false
	}
	return o.AuthTokens, true
}

// HasAuthTokens returns a boolean if a field has been set.
func (o *FeedRequestOptions) HasAuthTokens() bool {
	if o != nil && !IsNil(o.AuthTokens) {
		return true
	}

	return false
}

// SetAuthTokens gets a reference to the given []AuthToken and assigns it to the AuthTokens field.
func (o *FeedRequestOptions) SetAuthTokens(v []AuthToken) {
	o.AuthTokens = v
}

func (o FeedRequestOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FeedRequestOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["resultSize"] = o.ResultSize
	if !IsNil(o.TimezoneOffset) {
		toSerialize["timezoneOffset"] = o.TimezoneOffset
	}
	if !IsNil(o.CategoryToResultSize) {
		toSerialize["categoryToResultSize"] = o.CategoryToResultSize
	}
	if !IsNil(o.DatasourceFilter) {
		toSerialize["datasourceFilter"] = o.DatasourceFilter
	}
	if !IsNil(o.AuthTokens) {
		toSerialize["authTokens"] = o.AuthTokens
	}
	return toSerialize, nil
}

type NullableFeedRequestOptions struct {
	value *FeedRequestOptions
	isSet bool
}

func (v NullableFeedRequestOptions) Get() *FeedRequestOptions {
	return v.value
}

func (v *NullableFeedRequestOptions) Set(val *FeedRequestOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableFeedRequestOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableFeedRequestOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeedRequestOptions(val *FeedRequestOptions) *NullableFeedRequestOptions {
	return &NullableFeedRequestOptions{value: val, isSet: true}
}

func (v NullableFeedRequestOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeedRequestOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


