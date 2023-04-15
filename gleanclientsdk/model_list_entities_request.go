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

// checks if the ListEntitiesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListEntitiesRequest{}

// ListEntitiesRequest struct for ListEntitiesRequest
type ListEntitiesRequest struct {
	Filter []FacetFilter `json:"filter,omitempty"`
	// Use EntitiesSortOrder enum for SortOptions.sortBy
	Sort []SortOptions `json:"sort,omitempty"`
	EntityType *string `json:"entityType,omitempty"`
	// The datasource associated with the entity type, most commonly used with CUSTOM_ENTITIES
	Datasource *string `json:"datasource,omitempty"`
	// A query string to search for entities that each entity in the response must conform to. An empty query does not filter any entities.
	Query *string `json:"query,omitempty"`
	// Debug only search params to to change the flow of control in request handling. It will be passed around service boundaries/endpoints. For more details, https://docs.google.com/document/d/1e6taTfWUL8KNUC9de8kmmG2MG2L6cTx4ulOJfAshDTM/edit. Requires sufficient permissions.
	Sc *string `json:"sc,omitempty"`
	// List of entity fields to return (that aren't returned by default)
	IncludeFields []string `json:"includeFields,omitempty"`
	// Hint to the server about how many results to send back. Server may return less.
	PageSize *int32 `json:"pageSize,omitempty"`
	// Pagination cursor. A previously received opaque token representing the position in the overall results at which to start.
	Cursor *string `json:"cursor,omitempty"`
	// A string denoting the search surface from which the endpoint is called.
	Source *string `json:"source,omitempty"`
}

// NewListEntitiesRequest instantiates a new ListEntitiesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListEntitiesRequest() *ListEntitiesRequest {
	this := ListEntitiesRequest{}
	var entityType string = "PEOPLE"
	this.EntityType = &entityType
	return &this
}

// NewListEntitiesRequestWithDefaults instantiates a new ListEntitiesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListEntitiesRequestWithDefaults() *ListEntitiesRequest {
	this := ListEntitiesRequest{}
	var entityType string = "PEOPLE"
	this.EntityType = &entityType
	return &this
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *ListEntitiesRequest) GetFilter() []FacetFilter {
	if o == nil || IsNil(o.Filter) {
		var ret []FacetFilter
		return ret
	}
	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListEntitiesRequest) GetFilterOk() ([]FacetFilter, bool) {
	if o == nil || IsNil(o.Filter) {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *ListEntitiesRequest) HasFilter() bool {
	if o != nil && !IsNil(o.Filter) {
		return true
	}

	return false
}

// SetFilter gets a reference to the given []FacetFilter and assigns it to the Filter field.
func (o *ListEntitiesRequest) SetFilter(v []FacetFilter) {
	o.Filter = v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *ListEntitiesRequest) GetSort() []SortOptions {
	if o == nil || IsNil(o.Sort) {
		var ret []SortOptions
		return ret
	}
	return o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListEntitiesRequest) GetSortOk() ([]SortOptions, bool) {
	if o == nil || IsNil(o.Sort) {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *ListEntitiesRequest) HasSort() bool {
	if o != nil && !IsNil(o.Sort) {
		return true
	}

	return false
}

// SetSort gets a reference to the given []SortOptions and assigns it to the Sort field.
func (o *ListEntitiesRequest) SetSort(v []SortOptions) {
	o.Sort = v
}

// GetEntityType returns the EntityType field value if set, zero value otherwise.
func (o *ListEntitiesRequest) GetEntityType() string {
	if o == nil || IsNil(o.EntityType) {
		var ret string
		return ret
	}
	return *o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListEntitiesRequest) GetEntityTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EntityType) {
		return nil, false
	}
	return o.EntityType, true
}

// HasEntityType returns a boolean if a field has been set.
func (o *ListEntitiesRequest) HasEntityType() bool {
	if o != nil && !IsNil(o.EntityType) {
		return true
	}

	return false
}

// SetEntityType gets a reference to the given string and assigns it to the EntityType field.
func (o *ListEntitiesRequest) SetEntityType(v string) {
	o.EntityType = &v
}

// GetDatasource returns the Datasource field value if set, zero value otherwise.
func (o *ListEntitiesRequest) GetDatasource() string {
	if o == nil || IsNil(o.Datasource) {
		var ret string
		return ret
	}
	return *o.Datasource
}

// GetDatasourceOk returns a tuple with the Datasource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListEntitiesRequest) GetDatasourceOk() (*string, bool) {
	if o == nil || IsNil(o.Datasource) {
		return nil, false
	}
	return o.Datasource, true
}

// HasDatasource returns a boolean if a field has been set.
func (o *ListEntitiesRequest) HasDatasource() bool {
	if o != nil && !IsNil(o.Datasource) {
		return true
	}

	return false
}

// SetDatasource gets a reference to the given string and assigns it to the Datasource field.
func (o *ListEntitiesRequest) SetDatasource(v string) {
	o.Datasource = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *ListEntitiesRequest) GetQuery() string {
	if o == nil || IsNil(o.Query) {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListEntitiesRequest) GetQueryOk() (*string, bool) {
	if o == nil || IsNil(o.Query) {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *ListEntitiesRequest) HasQuery() bool {
	if o != nil && !IsNil(o.Query) {
		return true
	}

	return false
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *ListEntitiesRequest) SetQuery(v string) {
	o.Query = &v
}

// GetSc returns the Sc field value if set, zero value otherwise.
func (o *ListEntitiesRequest) GetSc() string {
	if o == nil || IsNil(o.Sc) {
		var ret string
		return ret
	}
	return *o.Sc
}

// GetScOk returns a tuple with the Sc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListEntitiesRequest) GetScOk() (*string, bool) {
	if o == nil || IsNil(o.Sc) {
		return nil, false
	}
	return o.Sc, true
}

// HasSc returns a boolean if a field has been set.
func (o *ListEntitiesRequest) HasSc() bool {
	if o != nil && !IsNil(o.Sc) {
		return true
	}

	return false
}

// SetSc gets a reference to the given string and assigns it to the Sc field.
func (o *ListEntitiesRequest) SetSc(v string) {
	o.Sc = &v
}

// GetIncludeFields returns the IncludeFields field value if set, zero value otherwise.
func (o *ListEntitiesRequest) GetIncludeFields() []string {
	if o == nil || IsNil(o.IncludeFields) {
		var ret []string
		return ret
	}
	return o.IncludeFields
}

// GetIncludeFieldsOk returns a tuple with the IncludeFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListEntitiesRequest) GetIncludeFieldsOk() ([]string, bool) {
	if o == nil || IsNil(o.IncludeFields) {
		return nil, false
	}
	return o.IncludeFields, true
}

// HasIncludeFields returns a boolean if a field has been set.
func (o *ListEntitiesRequest) HasIncludeFields() bool {
	if o != nil && !IsNil(o.IncludeFields) {
		return true
	}

	return false
}

// SetIncludeFields gets a reference to the given []string and assigns it to the IncludeFields field.
func (o *ListEntitiesRequest) SetIncludeFields(v []string) {
	o.IncludeFields = v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *ListEntitiesRequest) GetPageSize() int32 {
	if o == nil || IsNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListEntitiesRequest) GetPageSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *ListEntitiesRequest) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *ListEntitiesRequest) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetCursor returns the Cursor field value if set, zero value otherwise.
func (o *ListEntitiesRequest) GetCursor() string {
	if o == nil || IsNil(o.Cursor) {
		var ret string
		return ret
	}
	return *o.Cursor
}

// GetCursorOk returns a tuple with the Cursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListEntitiesRequest) GetCursorOk() (*string, bool) {
	if o == nil || IsNil(o.Cursor) {
		return nil, false
	}
	return o.Cursor, true
}

// HasCursor returns a boolean if a field has been set.
func (o *ListEntitiesRequest) HasCursor() bool {
	if o != nil && !IsNil(o.Cursor) {
		return true
	}

	return false
}

// SetCursor gets a reference to the given string and assigns it to the Cursor field.
func (o *ListEntitiesRequest) SetCursor(v string) {
	o.Cursor = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *ListEntitiesRequest) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListEntitiesRequest) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *ListEntitiesRequest) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *ListEntitiesRequest) SetSource(v string) {
	o.Source = &v
}

func (o ListEntitiesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListEntitiesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Filter) {
		toSerialize["filter"] = o.Filter
	}
	if !IsNil(o.Sort) {
		toSerialize["sort"] = o.Sort
	}
	if !IsNil(o.EntityType) {
		toSerialize["entityType"] = o.EntityType
	}
	if !IsNil(o.Datasource) {
		toSerialize["datasource"] = o.Datasource
	}
	if !IsNil(o.Query) {
		toSerialize["query"] = o.Query
	}
	if !IsNil(o.Sc) {
		toSerialize["sc"] = o.Sc
	}
	if !IsNil(o.IncludeFields) {
		toSerialize["includeFields"] = o.IncludeFields
	}
	if !IsNil(o.PageSize) {
		toSerialize["pageSize"] = o.PageSize
	}
	if !IsNil(o.Cursor) {
		toSerialize["cursor"] = o.Cursor
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	return toSerialize, nil
}

type NullableListEntitiesRequest struct {
	value *ListEntitiesRequest
	isSet bool
}

func (v NullableListEntitiesRequest) Get() *ListEntitiesRequest {
	return v.value
}

func (v *NullableListEntitiesRequest) Set(val *ListEntitiesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableListEntitiesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableListEntitiesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListEntitiesRequest(val *ListEntitiesRequest) *NullableListEntitiesRequest {
	return &NullableListEntitiesRequest{value: val, isSet: true}
}

func (v NullableListEntitiesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListEntitiesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


