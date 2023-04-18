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

// checks if the ListEntitiesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListEntitiesResponse{}

// ListEntitiesResponse struct for ListEntitiesResponse
type ListEntitiesResponse struct {
	Results []Person `json:"results,omitempty"`
	TeamResults []Team `json:"teamResults,omitempty"`
	CustomEntityResults []CustomEntity `json:"customEntityResults,omitempty"`
	FacetResults []FacetResult `json:"facetResults,omitempty"`
	// Pagination cursor. A previously received opaque token representing the position in the overall results at which to start.
	Cursor *string `json:"cursor,omitempty"`
	// The total number of entities available
	TotalCount *int32 `json:"totalCount,omitempty"`
	// Whether or not more entities can be fetched.
	HasMoreResults *bool `json:"hasMoreResults,omitempty"`
	// Sort options from EntitiesSortOrder supported for this response. Default is empty list.
	SortOptions []EntitiesSortOrder `json:"sortOptions,omitempty"`
	// list of Person attributes that are custom setup by deployment
	CustomFacetNames []string `json:"customFacetNames,omitempty"`
}

// NewListEntitiesResponse instantiates a new ListEntitiesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListEntitiesResponse() *ListEntitiesResponse {
	this := ListEntitiesResponse{}
	return &this
}

// NewListEntitiesResponseWithDefaults instantiates a new ListEntitiesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListEntitiesResponseWithDefaults() *ListEntitiesResponse {
	this := ListEntitiesResponse{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *ListEntitiesResponse) GetResults() []Person {
	if o == nil || IsNil(o.Results) {
		var ret []Person
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListEntitiesResponse) GetResultsOk() ([]Person, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *ListEntitiesResponse) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []Person and assigns it to the Results field.
func (o *ListEntitiesResponse) SetResults(v []Person) {
	o.Results = v
}

// GetTeamResults returns the TeamResults field value if set, zero value otherwise.
func (o *ListEntitiesResponse) GetTeamResults() []Team {
	if o == nil || IsNil(o.TeamResults) {
		var ret []Team
		return ret
	}
	return o.TeamResults
}

// GetTeamResultsOk returns a tuple with the TeamResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListEntitiesResponse) GetTeamResultsOk() ([]Team, bool) {
	if o == nil || IsNil(o.TeamResults) {
		return nil, false
	}
	return o.TeamResults, true
}

// HasTeamResults returns a boolean if a field has been set.
func (o *ListEntitiesResponse) HasTeamResults() bool {
	if o != nil && !IsNil(o.TeamResults) {
		return true
	}

	return false
}

// SetTeamResults gets a reference to the given []Team and assigns it to the TeamResults field.
func (o *ListEntitiesResponse) SetTeamResults(v []Team) {
	o.TeamResults = v
}

// GetCustomEntityResults returns the CustomEntityResults field value if set, zero value otherwise.
func (o *ListEntitiesResponse) GetCustomEntityResults() []CustomEntity {
	if o == nil || IsNil(o.CustomEntityResults) {
		var ret []CustomEntity
		return ret
	}
	return o.CustomEntityResults
}

// GetCustomEntityResultsOk returns a tuple with the CustomEntityResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListEntitiesResponse) GetCustomEntityResultsOk() ([]CustomEntity, bool) {
	if o == nil || IsNil(o.CustomEntityResults) {
		return nil, false
	}
	return o.CustomEntityResults, true
}

// HasCustomEntityResults returns a boolean if a field has been set.
func (o *ListEntitiesResponse) HasCustomEntityResults() bool {
	if o != nil && !IsNil(o.CustomEntityResults) {
		return true
	}

	return false
}

// SetCustomEntityResults gets a reference to the given []CustomEntity and assigns it to the CustomEntityResults field.
func (o *ListEntitiesResponse) SetCustomEntityResults(v []CustomEntity) {
	o.CustomEntityResults = v
}

// GetFacetResults returns the FacetResults field value if set, zero value otherwise.
func (o *ListEntitiesResponse) GetFacetResults() []FacetResult {
	if o == nil || IsNil(o.FacetResults) {
		var ret []FacetResult
		return ret
	}
	return o.FacetResults
}

// GetFacetResultsOk returns a tuple with the FacetResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListEntitiesResponse) GetFacetResultsOk() ([]FacetResult, bool) {
	if o == nil || IsNil(o.FacetResults) {
		return nil, false
	}
	return o.FacetResults, true
}

// HasFacetResults returns a boolean if a field has been set.
func (o *ListEntitiesResponse) HasFacetResults() bool {
	if o != nil && !IsNil(o.FacetResults) {
		return true
	}

	return false
}

// SetFacetResults gets a reference to the given []FacetResult and assigns it to the FacetResults field.
func (o *ListEntitiesResponse) SetFacetResults(v []FacetResult) {
	o.FacetResults = v
}

// GetCursor returns the Cursor field value if set, zero value otherwise.
func (o *ListEntitiesResponse) GetCursor() string {
	if o == nil || IsNil(o.Cursor) {
		var ret string
		return ret
	}
	return *o.Cursor
}

// GetCursorOk returns a tuple with the Cursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListEntitiesResponse) GetCursorOk() (*string, bool) {
	if o == nil || IsNil(o.Cursor) {
		return nil, false
	}
	return o.Cursor, true
}

// HasCursor returns a boolean if a field has been set.
func (o *ListEntitiesResponse) HasCursor() bool {
	if o != nil && !IsNil(o.Cursor) {
		return true
	}

	return false
}

// SetCursor gets a reference to the given string and assigns it to the Cursor field.
func (o *ListEntitiesResponse) SetCursor(v string) {
	o.Cursor = &v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *ListEntitiesResponse) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListEntitiesResponse) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *ListEntitiesResponse) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *ListEntitiesResponse) SetTotalCount(v int32) {
	o.TotalCount = &v
}

// GetHasMoreResults returns the HasMoreResults field value if set, zero value otherwise.
func (o *ListEntitiesResponse) GetHasMoreResults() bool {
	if o == nil || IsNil(o.HasMoreResults) {
		var ret bool
		return ret
	}
	return *o.HasMoreResults
}

// GetHasMoreResultsOk returns a tuple with the HasMoreResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListEntitiesResponse) GetHasMoreResultsOk() (*bool, bool) {
	if o == nil || IsNil(o.HasMoreResults) {
		return nil, false
	}
	return o.HasMoreResults, true
}

// HasHasMoreResults returns a boolean if a field has been set.
func (o *ListEntitiesResponse) HasHasMoreResults() bool {
	if o != nil && !IsNil(o.HasMoreResults) {
		return true
	}

	return false
}

// SetHasMoreResults gets a reference to the given bool and assigns it to the HasMoreResults field.
func (o *ListEntitiesResponse) SetHasMoreResults(v bool) {
	o.HasMoreResults = &v
}

// GetSortOptions returns the SortOptions field value if set, zero value otherwise.
func (o *ListEntitiesResponse) GetSortOptions() []EntitiesSortOrder {
	if o == nil || IsNil(o.SortOptions) {
		var ret []EntitiesSortOrder
		return ret
	}
	return o.SortOptions
}

// GetSortOptionsOk returns a tuple with the SortOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListEntitiesResponse) GetSortOptionsOk() ([]EntitiesSortOrder, bool) {
	if o == nil || IsNil(o.SortOptions) {
		return nil, false
	}
	return o.SortOptions, true
}

// HasSortOptions returns a boolean if a field has been set.
func (o *ListEntitiesResponse) HasSortOptions() bool {
	if o != nil && !IsNil(o.SortOptions) {
		return true
	}

	return false
}

// SetSortOptions gets a reference to the given []EntitiesSortOrder and assigns it to the SortOptions field.
func (o *ListEntitiesResponse) SetSortOptions(v []EntitiesSortOrder) {
	o.SortOptions = v
}

// GetCustomFacetNames returns the CustomFacetNames field value if set, zero value otherwise.
func (o *ListEntitiesResponse) GetCustomFacetNames() []string {
	if o == nil || IsNil(o.CustomFacetNames) {
		var ret []string
		return ret
	}
	return o.CustomFacetNames
}

// GetCustomFacetNamesOk returns a tuple with the CustomFacetNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListEntitiesResponse) GetCustomFacetNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.CustomFacetNames) {
		return nil, false
	}
	return o.CustomFacetNames, true
}

// HasCustomFacetNames returns a boolean if a field has been set.
func (o *ListEntitiesResponse) HasCustomFacetNames() bool {
	if o != nil && !IsNil(o.CustomFacetNames) {
		return true
	}

	return false
}

// SetCustomFacetNames gets a reference to the given []string and assigns it to the CustomFacetNames field.
func (o *ListEntitiesResponse) SetCustomFacetNames(v []string) {
	o.CustomFacetNames = v
}

func (o ListEntitiesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListEntitiesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	if !IsNil(o.TeamResults) {
		toSerialize["teamResults"] = o.TeamResults
	}
	if !IsNil(o.CustomEntityResults) {
		toSerialize["customEntityResults"] = o.CustomEntityResults
	}
	if !IsNil(o.FacetResults) {
		toSerialize["facetResults"] = o.FacetResults
	}
	if !IsNil(o.Cursor) {
		toSerialize["cursor"] = o.Cursor
	}
	if !IsNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	if !IsNil(o.HasMoreResults) {
		toSerialize["hasMoreResults"] = o.HasMoreResults
	}
	if !IsNil(o.SortOptions) {
		toSerialize["sortOptions"] = o.SortOptions
	}
	if !IsNil(o.CustomFacetNames) {
		toSerialize["customFacetNames"] = o.CustomFacetNames
	}
	return toSerialize, nil
}

type NullableListEntitiesResponse struct {
	value *ListEntitiesResponse
	isSet bool
}

func (v NullableListEntitiesResponse) Get() *ListEntitiesResponse {
	return v.value
}

func (v *NullableListEntitiesResponse) Set(val *ListEntitiesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListEntitiesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListEntitiesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListEntitiesResponse(val *ListEntitiesResponse) *NullableListEntitiesResponse {
	return &NullableListEntitiesResponse{value: val, isSet: true}
}

func (v NullableListEntitiesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListEntitiesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

