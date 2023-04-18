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

// checks if the SearchRequestAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchRequestAllOf{}

// SearchRequestAllOf struct for SearchRequestAllOf
type SearchRequestAllOf struct {
	// The search terms.
	Query *string `json:"query,omitempty"`
	// Pagination cursor. A previously received opaque token representing the position in the overall results at which to start.
	Cursor *string `json:"cursor,omitempty"`
	// The unique IDs of the result tabs for which to fetch results. This will have precedence over datasource filters if both are specified and in conflict.
	ResultTabIds []string `json:"resultTabIds,omitempty"`
	InputDetails *SearchRequestInputDetails `json:"inputDetails,omitempty"`
	RequestOptions *SearchRequestOptions `json:"requestOptions,omitempty"`
	// Timeout in milliseconds for the request. Backend should throw a 408 if request takes longer than this.
	TimeoutMillis *int32 `json:"timeoutMillis,omitempty"`
	// People associated with the search request. Hints to the server to fetch additional information for these people. Note that in this request, an email may be used as a person's obfuscatedId value.
	People []Person `json:"people,omitempty"`
	// Whether or not to disable spellcheck.
	DisableSpellcheck *bool `json:"disableSpellcheck,omitempty"`
}

// NewSearchRequestAllOf instantiates a new SearchRequestAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchRequestAllOf() *SearchRequestAllOf {
	this := SearchRequestAllOf{}
	return &this
}

// NewSearchRequestAllOfWithDefaults instantiates a new SearchRequestAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchRequestAllOfWithDefaults() *SearchRequestAllOf {
	this := SearchRequestAllOf{}
	return &this
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *SearchRequestAllOf) GetQuery() string {
	if o == nil || IsNil(o.Query) {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchRequestAllOf) GetQueryOk() (*string, bool) {
	if o == nil || IsNil(o.Query) {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *SearchRequestAllOf) HasQuery() bool {
	if o != nil && !IsNil(o.Query) {
		return true
	}

	return false
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *SearchRequestAllOf) SetQuery(v string) {
	o.Query = &v
}

// GetCursor returns the Cursor field value if set, zero value otherwise.
func (o *SearchRequestAllOf) GetCursor() string {
	if o == nil || IsNil(o.Cursor) {
		var ret string
		return ret
	}
	return *o.Cursor
}

// GetCursorOk returns a tuple with the Cursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchRequestAllOf) GetCursorOk() (*string, bool) {
	if o == nil || IsNil(o.Cursor) {
		return nil, false
	}
	return o.Cursor, true
}

// HasCursor returns a boolean if a field has been set.
func (o *SearchRequestAllOf) HasCursor() bool {
	if o != nil && !IsNil(o.Cursor) {
		return true
	}

	return false
}

// SetCursor gets a reference to the given string and assigns it to the Cursor field.
func (o *SearchRequestAllOf) SetCursor(v string) {
	o.Cursor = &v
}

// GetResultTabIds returns the ResultTabIds field value if set, zero value otherwise.
func (o *SearchRequestAllOf) GetResultTabIds() []string {
	if o == nil || IsNil(o.ResultTabIds) {
		var ret []string
		return ret
	}
	return o.ResultTabIds
}

// GetResultTabIdsOk returns a tuple with the ResultTabIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchRequestAllOf) GetResultTabIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ResultTabIds) {
		return nil, false
	}
	return o.ResultTabIds, true
}

// HasResultTabIds returns a boolean if a field has been set.
func (o *SearchRequestAllOf) HasResultTabIds() bool {
	if o != nil && !IsNil(o.ResultTabIds) {
		return true
	}

	return false
}

// SetResultTabIds gets a reference to the given []string and assigns it to the ResultTabIds field.
func (o *SearchRequestAllOf) SetResultTabIds(v []string) {
	o.ResultTabIds = v
}

// GetInputDetails returns the InputDetails field value if set, zero value otherwise.
func (o *SearchRequestAllOf) GetInputDetails() SearchRequestInputDetails {
	if o == nil || IsNil(o.InputDetails) {
		var ret SearchRequestInputDetails
		return ret
	}
	return *o.InputDetails
}

// GetInputDetailsOk returns a tuple with the InputDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchRequestAllOf) GetInputDetailsOk() (*SearchRequestInputDetails, bool) {
	if o == nil || IsNil(o.InputDetails) {
		return nil, false
	}
	return o.InputDetails, true
}

// HasInputDetails returns a boolean if a field has been set.
func (o *SearchRequestAllOf) HasInputDetails() bool {
	if o != nil && !IsNil(o.InputDetails) {
		return true
	}

	return false
}

// SetInputDetails gets a reference to the given SearchRequestInputDetails and assigns it to the InputDetails field.
func (o *SearchRequestAllOf) SetInputDetails(v SearchRequestInputDetails) {
	o.InputDetails = &v
}

// GetRequestOptions returns the RequestOptions field value if set, zero value otherwise.
func (o *SearchRequestAllOf) GetRequestOptions() SearchRequestOptions {
	if o == nil || IsNil(o.RequestOptions) {
		var ret SearchRequestOptions
		return ret
	}
	return *o.RequestOptions
}

// GetRequestOptionsOk returns a tuple with the RequestOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchRequestAllOf) GetRequestOptionsOk() (*SearchRequestOptions, bool) {
	if o == nil || IsNil(o.RequestOptions) {
		return nil, false
	}
	return o.RequestOptions, true
}

// HasRequestOptions returns a boolean if a field has been set.
func (o *SearchRequestAllOf) HasRequestOptions() bool {
	if o != nil && !IsNil(o.RequestOptions) {
		return true
	}

	return false
}

// SetRequestOptions gets a reference to the given SearchRequestOptions and assigns it to the RequestOptions field.
func (o *SearchRequestAllOf) SetRequestOptions(v SearchRequestOptions) {
	o.RequestOptions = &v
}

// GetTimeoutMillis returns the TimeoutMillis field value if set, zero value otherwise.
func (o *SearchRequestAllOf) GetTimeoutMillis() int32 {
	if o == nil || IsNil(o.TimeoutMillis) {
		var ret int32
		return ret
	}
	return *o.TimeoutMillis
}

// GetTimeoutMillisOk returns a tuple with the TimeoutMillis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchRequestAllOf) GetTimeoutMillisOk() (*int32, bool) {
	if o == nil || IsNil(o.TimeoutMillis) {
		return nil, false
	}
	return o.TimeoutMillis, true
}

// HasTimeoutMillis returns a boolean if a field has been set.
func (o *SearchRequestAllOf) HasTimeoutMillis() bool {
	if o != nil && !IsNil(o.TimeoutMillis) {
		return true
	}

	return false
}

// SetTimeoutMillis gets a reference to the given int32 and assigns it to the TimeoutMillis field.
func (o *SearchRequestAllOf) SetTimeoutMillis(v int32) {
	o.TimeoutMillis = &v
}

// GetPeople returns the People field value if set, zero value otherwise.
func (o *SearchRequestAllOf) GetPeople() []Person {
	if o == nil || IsNil(o.People) {
		var ret []Person
		return ret
	}
	return o.People
}

// GetPeopleOk returns a tuple with the People field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchRequestAllOf) GetPeopleOk() ([]Person, bool) {
	if o == nil || IsNil(o.People) {
		return nil, false
	}
	return o.People, true
}

// HasPeople returns a boolean if a field has been set.
func (o *SearchRequestAllOf) HasPeople() bool {
	if o != nil && !IsNil(o.People) {
		return true
	}

	return false
}

// SetPeople gets a reference to the given []Person and assigns it to the People field.
func (o *SearchRequestAllOf) SetPeople(v []Person) {
	o.People = v
}

// GetDisableSpellcheck returns the DisableSpellcheck field value if set, zero value otherwise.
func (o *SearchRequestAllOf) GetDisableSpellcheck() bool {
	if o == nil || IsNil(o.DisableSpellcheck) {
		var ret bool
		return ret
	}
	return *o.DisableSpellcheck
}

// GetDisableSpellcheckOk returns a tuple with the DisableSpellcheck field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchRequestAllOf) GetDisableSpellcheckOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableSpellcheck) {
		return nil, false
	}
	return o.DisableSpellcheck, true
}

// HasDisableSpellcheck returns a boolean if a field has been set.
func (o *SearchRequestAllOf) HasDisableSpellcheck() bool {
	if o != nil && !IsNil(o.DisableSpellcheck) {
		return true
	}

	return false
}

// SetDisableSpellcheck gets a reference to the given bool and assigns it to the DisableSpellcheck field.
func (o *SearchRequestAllOf) SetDisableSpellcheck(v bool) {
	o.DisableSpellcheck = &v
}

func (o SearchRequestAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchRequestAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Query) {
		toSerialize["query"] = o.Query
	}
	if !IsNil(o.Cursor) {
		toSerialize["cursor"] = o.Cursor
	}
	if !IsNil(o.ResultTabIds) {
		toSerialize["resultTabIds"] = o.ResultTabIds
	}
	if !IsNil(o.InputDetails) {
		toSerialize["inputDetails"] = o.InputDetails
	}
	if !IsNil(o.RequestOptions) {
		toSerialize["requestOptions"] = o.RequestOptions
	}
	if !IsNil(o.TimeoutMillis) {
		toSerialize["timeoutMillis"] = o.TimeoutMillis
	}
	if !IsNil(o.People) {
		toSerialize["people"] = o.People
	}
	if !IsNil(o.DisableSpellcheck) {
		toSerialize["disableSpellcheck"] = o.DisableSpellcheck
	}
	return toSerialize, nil
}

type NullableSearchRequestAllOf struct {
	value *SearchRequestAllOf
	isSet bool
}

func (v NullableSearchRequestAllOf) Get() *SearchRequestAllOf {
	return v.value
}

func (v *NullableSearchRequestAllOf) Set(val *SearchRequestAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchRequestAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchRequestAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchRequestAllOf(val *SearchRequestAllOf) *NullableSearchRequestAllOf {
	return &NullableSearchRequestAllOf{value: val, isSet: true}
}

func (v NullableSearchRequestAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchRequestAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

