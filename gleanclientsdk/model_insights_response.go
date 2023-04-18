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

// checks if the InsightsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InsightsResponse{}

// InsightsResponse struct for InsightsResponse
type InsightsResponse struct {
	// List of timeseries to make charts (if applicable).
	Timeseries []LabeledCountInfo `json:"timeseries,omitempty"`
	Users *UserInsightsResponse `json:"users,omitempty"`
	Content *ContentInsightsResponse `json:"content,omitempty"`
	Queries *QueryInsightsResponse `json:"queries,omitempty"`
	Collections *ContentInsightsResponse `json:"collections,omitempty"`
	CollectionsV2 *ContentInsightsResponse `json:"collectionsV2,omitempty"`
	Shortcuts *ShortcutInsightsResponse `json:"shortcuts,omitempty"`
	Announcements *ContentInsightsResponse `json:"announcements,omitempty"`
	Answers *ContentInsightsResponse `json:"answers,omitempty"`
	// list of all departments.
	Departments []string `json:"departments,omitempty"`
}

// NewInsightsResponse instantiates a new InsightsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInsightsResponse() *InsightsResponse {
	this := InsightsResponse{}
	return &this
}

// NewInsightsResponseWithDefaults instantiates a new InsightsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInsightsResponseWithDefaults() *InsightsResponse {
	this := InsightsResponse{}
	return &this
}

// GetTimeseries returns the Timeseries field value if set, zero value otherwise.
func (o *InsightsResponse) GetTimeseries() []LabeledCountInfo {
	if o == nil || IsNil(o.Timeseries) {
		var ret []LabeledCountInfo
		return ret
	}
	return o.Timeseries
}

// GetTimeseriesOk returns a tuple with the Timeseries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InsightsResponse) GetTimeseriesOk() ([]LabeledCountInfo, bool) {
	if o == nil || IsNil(o.Timeseries) {
		return nil, false
	}
	return o.Timeseries, true
}

// HasTimeseries returns a boolean if a field has been set.
func (o *InsightsResponse) HasTimeseries() bool {
	if o != nil && !IsNil(o.Timeseries) {
		return true
	}

	return false
}

// SetTimeseries gets a reference to the given []LabeledCountInfo and assigns it to the Timeseries field.
func (o *InsightsResponse) SetTimeseries(v []LabeledCountInfo) {
	o.Timeseries = v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *InsightsResponse) GetUsers() UserInsightsResponse {
	if o == nil || IsNil(o.Users) {
		var ret UserInsightsResponse
		return ret
	}
	return *o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InsightsResponse) GetUsersOk() (*UserInsightsResponse, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *InsightsResponse) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given UserInsightsResponse and assigns it to the Users field.
func (o *InsightsResponse) SetUsers(v UserInsightsResponse) {
	o.Users = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *InsightsResponse) GetContent() ContentInsightsResponse {
	if o == nil || IsNil(o.Content) {
		var ret ContentInsightsResponse
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InsightsResponse) GetContentOk() (*ContentInsightsResponse, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *InsightsResponse) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given ContentInsightsResponse and assigns it to the Content field.
func (o *InsightsResponse) SetContent(v ContentInsightsResponse) {
	o.Content = &v
}

// GetQueries returns the Queries field value if set, zero value otherwise.
func (o *InsightsResponse) GetQueries() QueryInsightsResponse {
	if o == nil || IsNil(o.Queries) {
		var ret QueryInsightsResponse
		return ret
	}
	return *o.Queries
}

// GetQueriesOk returns a tuple with the Queries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InsightsResponse) GetQueriesOk() (*QueryInsightsResponse, bool) {
	if o == nil || IsNil(o.Queries) {
		return nil, false
	}
	return o.Queries, true
}

// HasQueries returns a boolean if a field has been set.
func (o *InsightsResponse) HasQueries() bool {
	if o != nil && !IsNil(o.Queries) {
		return true
	}

	return false
}

// SetQueries gets a reference to the given QueryInsightsResponse and assigns it to the Queries field.
func (o *InsightsResponse) SetQueries(v QueryInsightsResponse) {
	o.Queries = &v
}

// GetCollections returns the Collections field value if set, zero value otherwise.
func (o *InsightsResponse) GetCollections() ContentInsightsResponse {
	if o == nil || IsNil(o.Collections) {
		var ret ContentInsightsResponse
		return ret
	}
	return *o.Collections
}

// GetCollectionsOk returns a tuple with the Collections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InsightsResponse) GetCollectionsOk() (*ContentInsightsResponse, bool) {
	if o == nil || IsNil(o.Collections) {
		return nil, false
	}
	return o.Collections, true
}

// HasCollections returns a boolean if a field has been set.
func (o *InsightsResponse) HasCollections() bool {
	if o != nil && !IsNil(o.Collections) {
		return true
	}

	return false
}

// SetCollections gets a reference to the given ContentInsightsResponse and assigns it to the Collections field.
func (o *InsightsResponse) SetCollections(v ContentInsightsResponse) {
	o.Collections = &v
}

// GetCollectionsV2 returns the CollectionsV2 field value if set, zero value otherwise.
func (o *InsightsResponse) GetCollectionsV2() ContentInsightsResponse {
	if o == nil || IsNil(o.CollectionsV2) {
		var ret ContentInsightsResponse
		return ret
	}
	return *o.CollectionsV2
}

// GetCollectionsV2Ok returns a tuple with the CollectionsV2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InsightsResponse) GetCollectionsV2Ok() (*ContentInsightsResponse, bool) {
	if o == nil || IsNil(o.CollectionsV2) {
		return nil, false
	}
	return o.CollectionsV2, true
}

// HasCollectionsV2 returns a boolean if a field has been set.
func (o *InsightsResponse) HasCollectionsV2() bool {
	if o != nil && !IsNil(o.CollectionsV2) {
		return true
	}

	return false
}

// SetCollectionsV2 gets a reference to the given ContentInsightsResponse and assigns it to the CollectionsV2 field.
func (o *InsightsResponse) SetCollectionsV2(v ContentInsightsResponse) {
	o.CollectionsV2 = &v
}

// GetShortcuts returns the Shortcuts field value if set, zero value otherwise.
func (o *InsightsResponse) GetShortcuts() ShortcutInsightsResponse {
	if o == nil || IsNil(o.Shortcuts) {
		var ret ShortcutInsightsResponse
		return ret
	}
	return *o.Shortcuts
}

// GetShortcutsOk returns a tuple with the Shortcuts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InsightsResponse) GetShortcutsOk() (*ShortcutInsightsResponse, bool) {
	if o == nil || IsNil(o.Shortcuts) {
		return nil, false
	}
	return o.Shortcuts, true
}

// HasShortcuts returns a boolean if a field has been set.
func (o *InsightsResponse) HasShortcuts() bool {
	if o != nil && !IsNil(o.Shortcuts) {
		return true
	}

	return false
}

// SetShortcuts gets a reference to the given ShortcutInsightsResponse and assigns it to the Shortcuts field.
func (o *InsightsResponse) SetShortcuts(v ShortcutInsightsResponse) {
	o.Shortcuts = &v
}

// GetAnnouncements returns the Announcements field value if set, zero value otherwise.
func (o *InsightsResponse) GetAnnouncements() ContentInsightsResponse {
	if o == nil || IsNil(o.Announcements) {
		var ret ContentInsightsResponse
		return ret
	}
	return *o.Announcements
}

// GetAnnouncementsOk returns a tuple with the Announcements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InsightsResponse) GetAnnouncementsOk() (*ContentInsightsResponse, bool) {
	if o == nil || IsNil(o.Announcements) {
		return nil, false
	}
	return o.Announcements, true
}

// HasAnnouncements returns a boolean if a field has been set.
func (o *InsightsResponse) HasAnnouncements() bool {
	if o != nil && !IsNil(o.Announcements) {
		return true
	}

	return false
}

// SetAnnouncements gets a reference to the given ContentInsightsResponse and assigns it to the Announcements field.
func (o *InsightsResponse) SetAnnouncements(v ContentInsightsResponse) {
	o.Announcements = &v
}

// GetAnswers returns the Answers field value if set, zero value otherwise.
func (o *InsightsResponse) GetAnswers() ContentInsightsResponse {
	if o == nil || IsNil(o.Answers) {
		var ret ContentInsightsResponse
		return ret
	}
	return *o.Answers
}

// GetAnswersOk returns a tuple with the Answers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InsightsResponse) GetAnswersOk() (*ContentInsightsResponse, bool) {
	if o == nil || IsNil(o.Answers) {
		return nil, false
	}
	return o.Answers, true
}

// HasAnswers returns a boolean if a field has been set.
func (o *InsightsResponse) HasAnswers() bool {
	if o != nil && !IsNil(o.Answers) {
		return true
	}

	return false
}

// SetAnswers gets a reference to the given ContentInsightsResponse and assigns it to the Answers field.
func (o *InsightsResponse) SetAnswers(v ContentInsightsResponse) {
	o.Answers = &v
}

// GetDepartments returns the Departments field value if set, zero value otherwise.
func (o *InsightsResponse) GetDepartments() []string {
	if o == nil || IsNil(o.Departments) {
		var ret []string
		return ret
	}
	return o.Departments
}

// GetDepartmentsOk returns a tuple with the Departments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InsightsResponse) GetDepartmentsOk() ([]string, bool) {
	if o == nil || IsNil(o.Departments) {
		return nil, false
	}
	return o.Departments, true
}

// HasDepartments returns a boolean if a field has been set.
func (o *InsightsResponse) HasDepartments() bool {
	if o != nil && !IsNil(o.Departments) {
		return true
	}

	return false
}

// SetDepartments gets a reference to the given []string and assigns it to the Departments field.
func (o *InsightsResponse) SetDepartments(v []string) {
	o.Departments = v
}

func (o InsightsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InsightsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Timeseries) {
		toSerialize["timeseries"] = o.Timeseries
	}
	if !IsNil(o.Users) {
		toSerialize["users"] = o.Users
	}
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	if !IsNil(o.Queries) {
		toSerialize["queries"] = o.Queries
	}
	if !IsNil(o.Collections) {
		toSerialize["collections"] = o.Collections
	}
	if !IsNil(o.CollectionsV2) {
		toSerialize["collectionsV2"] = o.CollectionsV2
	}
	if !IsNil(o.Shortcuts) {
		toSerialize["shortcuts"] = o.Shortcuts
	}
	if !IsNil(o.Announcements) {
		toSerialize["announcements"] = o.Announcements
	}
	if !IsNil(o.Answers) {
		toSerialize["answers"] = o.Answers
	}
	if !IsNil(o.Departments) {
		toSerialize["departments"] = o.Departments
	}
	return toSerialize, nil
}

type NullableInsightsResponse struct {
	value *InsightsResponse
	isSet bool
}

func (v NullableInsightsResponse) Get() *InsightsResponse {
	return v.value
}

func (v *NullableInsightsResponse) Set(val *InsightsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableInsightsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableInsightsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInsightsResponse(val *InsightsResponse) *NullableInsightsResponse {
	return &NullableInsightsResponse{value: val, isSet: true}
}

func (v NullableInsightsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInsightsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

