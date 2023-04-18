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

// checks if the QueryInsight type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QueryInsight{}

// QueryInsight struct for QueryInsight
type QueryInsight struct {
	// The query string the information is about.
	Query string `json:"query"`
	SearchCount *CountInfo `json:"searchCount,omitempty"`
	SearchorCount *CountInfo `json:"searchorCount,omitempty"`
	SearchWithClickCount *CountInfo `json:"searchWithClickCount,omitempty"`
	ClickCount *CountInfo `json:"clickCount,omitempty"`
	// list of similar queries to current one.
	SimilarQueries []QueryInsight `json:"similarQueries,omitempty"`
}

// NewQueryInsight instantiates a new QueryInsight object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryInsight(query string) *QueryInsight {
	this := QueryInsight{}
	this.Query = query
	return &this
}

// NewQueryInsightWithDefaults instantiates a new QueryInsight object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryInsightWithDefaults() *QueryInsight {
	this := QueryInsight{}
	return &this
}

// GetQuery returns the Query field value
func (o *QueryInsight) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *QueryInsight) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *QueryInsight) SetQuery(v string) {
	o.Query = v
}

// GetSearchCount returns the SearchCount field value if set, zero value otherwise.
func (o *QueryInsight) GetSearchCount() CountInfo {
	if o == nil || IsNil(o.SearchCount) {
		var ret CountInfo
		return ret
	}
	return *o.SearchCount
}

// GetSearchCountOk returns a tuple with the SearchCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryInsight) GetSearchCountOk() (*CountInfo, bool) {
	if o == nil || IsNil(o.SearchCount) {
		return nil, false
	}
	return o.SearchCount, true
}

// HasSearchCount returns a boolean if a field has been set.
func (o *QueryInsight) HasSearchCount() bool {
	if o != nil && !IsNil(o.SearchCount) {
		return true
	}

	return false
}

// SetSearchCount gets a reference to the given CountInfo and assigns it to the SearchCount field.
func (o *QueryInsight) SetSearchCount(v CountInfo) {
	o.SearchCount = &v
}

// GetSearchorCount returns the SearchorCount field value if set, zero value otherwise.
func (o *QueryInsight) GetSearchorCount() CountInfo {
	if o == nil || IsNil(o.SearchorCount) {
		var ret CountInfo
		return ret
	}
	return *o.SearchorCount
}

// GetSearchorCountOk returns a tuple with the SearchorCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryInsight) GetSearchorCountOk() (*CountInfo, bool) {
	if o == nil || IsNil(o.SearchorCount) {
		return nil, false
	}
	return o.SearchorCount, true
}

// HasSearchorCount returns a boolean if a field has been set.
func (o *QueryInsight) HasSearchorCount() bool {
	if o != nil && !IsNil(o.SearchorCount) {
		return true
	}

	return false
}

// SetSearchorCount gets a reference to the given CountInfo and assigns it to the SearchorCount field.
func (o *QueryInsight) SetSearchorCount(v CountInfo) {
	o.SearchorCount = &v
}

// GetSearchWithClickCount returns the SearchWithClickCount field value if set, zero value otherwise.
func (o *QueryInsight) GetSearchWithClickCount() CountInfo {
	if o == nil || IsNil(o.SearchWithClickCount) {
		var ret CountInfo
		return ret
	}
	return *o.SearchWithClickCount
}

// GetSearchWithClickCountOk returns a tuple with the SearchWithClickCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryInsight) GetSearchWithClickCountOk() (*CountInfo, bool) {
	if o == nil || IsNil(o.SearchWithClickCount) {
		return nil, false
	}
	return o.SearchWithClickCount, true
}

// HasSearchWithClickCount returns a boolean if a field has been set.
func (o *QueryInsight) HasSearchWithClickCount() bool {
	if o != nil && !IsNil(o.SearchWithClickCount) {
		return true
	}

	return false
}

// SetSearchWithClickCount gets a reference to the given CountInfo and assigns it to the SearchWithClickCount field.
func (o *QueryInsight) SetSearchWithClickCount(v CountInfo) {
	o.SearchWithClickCount = &v
}

// GetClickCount returns the ClickCount field value if set, zero value otherwise.
func (o *QueryInsight) GetClickCount() CountInfo {
	if o == nil || IsNil(o.ClickCount) {
		var ret CountInfo
		return ret
	}
	return *o.ClickCount
}

// GetClickCountOk returns a tuple with the ClickCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryInsight) GetClickCountOk() (*CountInfo, bool) {
	if o == nil || IsNil(o.ClickCount) {
		return nil, false
	}
	return o.ClickCount, true
}

// HasClickCount returns a boolean if a field has been set.
func (o *QueryInsight) HasClickCount() bool {
	if o != nil && !IsNil(o.ClickCount) {
		return true
	}

	return false
}

// SetClickCount gets a reference to the given CountInfo and assigns it to the ClickCount field.
func (o *QueryInsight) SetClickCount(v CountInfo) {
	o.ClickCount = &v
}

// GetSimilarQueries returns the SimilarQueries field value if set, zero value otherwise.
func (o *QueryInsight) GetSimilarQueries() []QueryInsight {
	if o == nil || IsNil(o.SimilarQueries) {
		var ret []QueryInsight
		return ret
	}
	return o.SimilarQueries
}

// GetSimilarQueriesOk returns a tuple with the SimilarQueries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryInsight) GetSimilarQueriesOk() ([]QueryInsight, bool) {
	if o == nil || IsNil(o.SimilarQueries) {
		return nil, false
	}
	return o.SimilarQueries, true
}

// HasSimilarQueries returns a boolean if a field has been set.
func (o *QueryInsight) HasSimilarQueries() bool {
	if o != nil && !IsNil(o.SimilarQueries) {
		return true
	}

	return false
}

// SetSimilarQueries gets a reference to the given []QueryInsight and assigns it to the SimilarQueries field.
func (o *QueryInsight) SetSimilarQueries(v []QueryInsight) {
	o.SimilarQueries = v
}

func (o QueryInsight) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryInsight) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["query"] = o.Query
	if !IsNil(o.SearchCount) {
		toSerialize["searchCount"] = o.SearchCount
	}
	if !IsNil(o.SearchorCount) {
		toSerialize["searchorCount"] = o.SearchorCount
	}
	if !IsNil(o.SearchWithClickCount) {
		toSerialize["searchWithClickCount"] = o.SearchWithClickCount
	}
	if !IsNil(o.ClickCount) {
		toSerialize["clickCount"] = o.ClickCount
	}
	if !IsNil(o.SimilarQueries) {
		toSerialize["similarQueries"] = o.SimilarQueries
	}
	return toSerialize, nil
}

type NullableQueryInsight struct {
	value *QueryInsight
	isSet bool
}

func (v NullableQueryInsight) Get() *QueryInsight {
	return v.value
}

func (v *NullableQueryInsight) Set(val *QueryInsight) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryInsight) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryInsight) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryInsight(val *QueryInsight) *NullableQueryInsight {
	return &NullableQueryInsight{value: val, isSet: true}
}

func (v NullableQueryInsight) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryInsight) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

