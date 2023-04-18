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

// checks if the ListShortcutsPaginatedRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListShortcutsPaginatedRequest{}

// ListShortcutsPaginatedRequest struct for ListShortcutsPaginatedRequest
type ListShortcutsPaginatedRequest struct {
	// Array of fields/data to be included in response that are not included by default
	IncludeFields []string `json:"includeFields,omitempty"`
	PageSize int32 `json:"pageSize"`
	// A token specifying the position in the overall results to start at. Received from the endpoint and iterated back. Currently being used as page no (as we implement offset pagination)
	Cursor *string `json:"cursor,omitempty"`
	// A list of filters for the query. An AND is assumed between different filters. We support filters on Go Link name, author, department and type.
	Filters []FacetFilter `json:"filters,omitempty"`
	Sort *SortOptions `json:"sort,omitempty"`
	// Search query that should be a substring in atleast one of the fields (alias , inputAlias, destinationUrl, description). Empty query does not filter shortcuts.
	Query *string `json:"query,omitempty"`
}

// NewListShortcutsPaginatedRequest instantiates a new ListShortcutsPaginatedRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListShortcutsPaginatedRequest(pageSize int32) *ListShortcutsPaginatedRequest {
	this := ListShortcutsPaginatedRequest{}
	this.PageSize = pageSize
	return &this
}

// NewListShortcutsPaginatedRequestWithDefaults instantiates a new ListShortcutsPaginatedRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListShortcutsPaginatedRequestWithDefaults() *ListShortcutsPaginatedRequest {
	this := ListShortcutsPaginatedRequest{}
	return &this
}

// GetIncludeFields returns the IncludeFields field value if set, zero value otherwise.
func (o *ListShortcutsPaginatedRequest) GetIncludeFields() []string {
	if o == nil || IsNil(o.IncludeFields) {
		var ret []string
		return ret
	}
	return o.IncludeFields
}

// GetIncludeFieldsOk returns a tuple with the IncludeFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListShortcutsPaginatedRequest) GetIncludeFieldsOk() ([]string, bool) {
	if o == nil || IsNil(o.IncludeFields) {
		return nil, false
	}
	return o.IncludeFields, true
}

// HasIncludeFields returns a boolean if a field has been set.
func (o *ListShortcutsPaginatedRequest) HasIncludeFields() bool {
	if o != nil && !IsNil(o.IncludeFields) {
		return true
	}

	return false
}

// SetIncludeFields gets a reference to the given []string and assigns it to the IncludeFields field.
func (o *ListShortcutsPaginatedRequest) SetIncludeFields(v []string) {
	o.IncludeFields = v
}

// GetPageSize returns the PageSize field value
func (o *ListShortcutsPaginatedRequest) GetPageSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value
// and a boolean to check if the value has been set.
func (o *ListShortcutsPaginatedRequest) GetPageSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageSize, true
}

// SetPageSize sets field value
func (o *ListShortcutsPaginatedRequest) SetPageSize(v int32) {
	o.PageSize = v
}

// GetCursor returns the Cursor field value if set, zero value otherwise.
func (o *ListShortcutsPaginatedRequest) GetCursor() string {
	if o == nil || IsNil(o.Cursor) {
		var ret string
		return ret
	}
	return *o.Cursor
}

// GetCursorOk returns a tuple with the Cursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListShortcutsPaginatedRequest) GetCursorOk() (*string, bool) {
	if o == nil || IsNil(o.Cursor) {
		return nil, false
	}
	return o.Cursor, true
}

// HasCursor returns a boolean if a field has been set.
func (o *ListShortcutsPaginatedRequest) HasCursor() bool {
	if o != nil && !IsNil(o.Cursor) {
		return true
	}

	return false
}

// SetCursor gets a reference to the given string and assigns it to the Cursor field.
func (o *ListShortcutsPaginatedRequest) SetCursor(v string) {
	o.Cursor = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *ListShortcutsPaginatedRequest) GetFilters() []FacetFilter {
	if o == nil || IsNil(o.Filters) {
		var ret []FacetFilter
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListShortcutsPaginatedRequest) GetFiltersOk() ([]FacetFilter, bool) {
	if o == nil || IsNil(o.Filters) {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *ListShortcutsPaginatedRequest) HasFilters() bool {
	if o != nil && !IsNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given []FacetFilter and assigns it to the Filters field.
func (o *ListShortcutsPaginatedRequest) SetFilters(v []FacetFilter) {
	o.Filters = v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *ListShortcutsPaginatedRequest) GetSort() SortOptions {
	if o == nil || IsNil(o.Sort) {
		var ret SortOptions
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListShortcutsPaginatedRequest) GetSortOk() (*SortOptions, bool) {
	if o == nil || IsNil(o.Sort) {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *ListShortcutsPaginatedRequest) HasSort() bool {
	if o != nil && !IsNil(o.Sort) {
		return true
	}

	return false
}

// SetSort gets a reference to the given SortOptions and assigns it to the Sort field.
func (o *ListShortcutsPaginatedRequest) SetSort(v SortOptions) {
	o.Sort = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *ListShortcutsPaginatedRequest) GetQuery() string {
	if o == nil || IsNil(o.Query) {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListShortcutsPaginatedRequest) GetQueryOk() (*string, bool) {
	if o == nil || IsNil(o.Query) {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *ListShortcutsPaginatedRequest) HasQuery() bool {
	if o != nil && !IsNil(o.Query) {
		return true
	}

	return false
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *ListShortcutsPaginatedRequest) SetQuery(v string) {
	o.Query = &v
}

func (o ListShortcutsPaginatedRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListShortcutsPaginatedRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IncludeFields) {
		toSerialize["includeFields"] = o.IncludeFields
	}
	toSerialize["pageSize"] = o.PageSize
	if !IsNil(o.Cursor) {
		toSerialize["cursor"] = o.Cursor
	}
	if !IsNil(o.Filters) {
		toSerialize["filters"] = o.Filters
	}
	if !IsNil(o.Sort) {
		toSerialize["sort"] = o.Sort
	}
	if !IsNil(o.Query) {
		toSerialize["query"] = o.Query
	}
	return toSerialize, nil
}

type NullableListShortcutsPaginatedRequest struct {
	value *ListShortcutsPaginatedRequest
	isSet bool
}

func (v NullableListShortcutsPaginatedRequest) Get() *ListShortcutsPaginatedRequest {
	return v.value
}

func (v *NullableListShortcutsPaginatedRequest) Set(val *ListShortcutsPaginatedRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableListShortcutsPaginatedRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableListShortcutsPaginatedRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListShortcutsPaginatedRequest(val *ListShortcutsPaginatedRequest) *NullableListShortcutsPaginatedRequest {
	return &NullableListShortcutsPaginatedRequest{value: val, isSet: true}
}

func (v NullableListShortcutsPaginatedRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListShortcutsPaginatedRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


