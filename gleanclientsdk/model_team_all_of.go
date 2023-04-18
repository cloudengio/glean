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
	"time"
)

// checks if the TeamAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TeamAllOf{}

// TeamAllOf struct for TeamAllOf
type TeamAllOf struct {
	// Unique identifier
	Id string `json:"id"`
	// Team name
	Name string `json:"name"`
	// A description of the team
	Description *string `json:"description,omitempty"`
	// Typically the highest level organizational unit; generally applies to bigger companies with multiple distinct businesses.
	BusinessUnit *string `json:"businessUnit,omitempty"`
	// An organizational unit where everyone has a similar task, e.g. `Engineering`.
	Department *string `json:"department,omitempty"`
	// A link to the team's photo
	PhotoUrl *string `json:"photoUrl,omitempty"`
	// A link to the team's banner photo
	BannerUrl *string `json:"bannerUrl,omitempty"`
	// Link to a team page on the internet or your company's intranet
	ExternalLink *string `json:"externalLink,omitempty"`
	// The members on this team
	Members []PersonToTeamRelationship `json:"members,omitempty"`
	// Number of members on this team (recursive; includes all individuals that belong to this team, and all individuals that belong to a subteam within this team)
	MemberCount *int32 `json:"memberCount,omitempty"`
	// The emails for this team
	Emails []TeamEmail `json:"emails,omitempty"`
	// The datasource profiles of the team
	DatasourceProfiles []DatasourceProfile `json:"datasourceProfiles,omitempty"`
	// the data source of the team, e.g. GDRIVE
	Datasource *string `json:"datasource,omitempty"`
	// For teams created from docs, the doc title. Otherwise empty.
	CreatedFrom *string `json:"createdFrom,omitempty"`
	// when this team was last updated.
	LastUpdatedAt *time.Time `json:"lastUpdatedAt,omitempty"`
	// whether this team is fully processed or there are still unprocessed operations that'll affect it
	Status *string `json:"status,omitempty"`
	// can this team be deleted. Some manually ingested teams like GCS_CSV or PUSH_API cannot
	CanBeDeleted *bool `json:"canBeDeleted,omitempty"`
	// The logging id of the team used in scrubbed logs, client analytics, and metrics.
	LoggingId *string `json:"loggingId,omitempty"`
}

// NewTeamAllOf instantiates a new TeamAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeamAllOf(id string, name string) *TeamAllOf {
	this := TeamAllOf{}
	this.Id = id
	this.Name = name
	var status string = "PROCESSED"
	this.Status = &status
	var canBeDeleted bool = true
	this.CanBeDeleted = &canBeDeleted
	return &this
}

// NewTeamAllOfWithDefaults instantiates a new TeamAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeamAllOfWithDefaults() *TeamAllOf {
	this := TeamAllOf{}
	var status string = "PROCESSED"
	this.Status = &status
	var canBeDeleted bool = true
	this.CanBeDeleted = &canBeDeleted
	return &this
}

// GetId returns the Id field value
func (o *TeamAllOf) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TeamAllOf) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TeamAllOf) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *TeamAllOf) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TeamAllOf) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TeamAllOf) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TeamAllOf) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TeamAllOf) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TeamAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetBusinessUnit returns the BusinessUnit field value if set, zero value otherwise.
func (o *TeamAllOf) GetBusinessUnit() string {
	if o == nil || IsNil(o.BusinessUnit) {
		var ret string
		return ret
	}
	return *o.BusinessUnit
}

// GetBusinessUnitOk returns a tuple with the BusinessUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamAllOf) GetBusinessUnitOk() (*string, bool) {
	if o == nil || IsNil(o.BusinessUnit) {
		return nil, false
	}
	return o.BusinessUnit, true
}

// HasBusinessUnit returns a boolean if a field has been set.
func (o *TeamAllOf) HasBusinessUnit() bool {
	if o != nil && !IsNil(o.BusinessUnit) {
		return true
	}

	return false
}

// SetBusinessUnit gets a reference to the given string and assigns it to the BusinessUnit field.
func (o *TeamAllOf) SetBusinessUnit(v string) {
	o.BusinessUnit = &v
}

// GetDepartment returns the Department field value if set, zero value otherwise.
func (o *TeamAllOf) GetDepartment() string {
	if o == nil || IsNil(o.Department) {
		var ret string
		return ret
	}
	return *o.Department
}

// GetDepartmentOk returns a tuple with the Department field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamAllOf) GetDepartmentOk() (*string, bool) {
	if o == nil || IsNil(o.Department) {
		return nil, false
	}
	return o.Department, true
}

// HasDepartment returns a boolean if a field has been set.
func (o *TeamAllOf) HasDepartment() bool {
	if o != nil && !IsNil(o.Department) {
		return true
	}

	return false
}

// SetDepartment gets a reference to the given string and assigns it to the Department field.
func (o *TeamAllOf) SetDepartment(v string) {
	o.Department = &v
}

// GetPhotoUrl returns the PhotoUrl field value if set, zero value otherwise.
func (o *TeamAllOf) GetPhotoUrl() string {
	if o == nil || IsNil(o.PhotoUrl) {
		var ret string
		return ret
	}
	return *o.PhotoUrl
}

// GetPhotoUrlOk returns a tuple with the PhotoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamAllOf) GetPhotoUrlOk() (*string, bool) {
	if o == nil || IsNil(o.PhotoUrl) {
		return nil, false
	}
	return o.PhotoUrl, true
}

// HasPhotoUrl returns a boolean if a field has been set.
func (o *TeamAllOf) HasPhotoUrl() bool {
	if o != nil && !IsNil(o.PhotoUrl) {
		return true
	}

	return false
}

// SetPhotoUrl gets a reference to the given string and assigns it to the PhotoUrl field.
func (o *TeamAllOf) SetPhotoUrl(v string) {
	o.PhotoUrl = &v
}

// GetBannerUrl returns the BannerUrl field value if set, zero value otherwise.
func (o *TeamAllOf) GetBannerUrl() string {
	if o == nil || IsNil(o.BannerUrl) {
		var ret string
		return ret
	}
	return *o.BannerUrl
}

// GetBannerUrlOk returns a tuple with the BannerUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamAllOf) GetBannerUrlOk() (*string, bool) {
	if o == nil || IsNil(o.BannerUrl) {
		return nil, false
	}
	return o.BannerUrl, true
}

// HasBannerUrl returns a boolean if a field has been set.
func (o *TeamAllOf) HasBannerUrl() bool {
	if o != nil && !IsNil(o.BannerUrl) {
		return true
	}

	return false
}

// SetBannerUrl gets a reference to the given string and assigns it to the BannerUrl field.
func (o *TeamAllOf) SetBannerUrl(v string) {
	o.BannerUrl = &v
}

// GetExternalLink returns the ExternalLink field value if set, zero value otherwise.
func (o *TeamAllOf) GetExternalLink() string {
	if o == nil || IsNil(o.ExternalLink) {
		var ret string
		return ret
	}
	return *o.ExternalLink
}

// GetExternalLinkOk returns a tuple with the ExternalLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamAllOf) GetExternalLinkOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalLink) {
		return nil, false
	}
	return o.ExternalLink, true
}

// HasExternalLink returns a boolean if a field has been set.
func (o *TeamAllOf) HasExternalLink() bool {
	if o != nil && !IsNil(o.ExternalLink) {
		return true
	}

	return false
}

// SetExternalLink gets a reference to the given string and assigns it to the ExternalLink field.
func (o *TeamAllOf) SetExternalLink(v string) {
	o.ExternalLink = &v
}

// GetMembers returns the Members field value if set, zero value otherwise.
func (o *TeamAllOf) GetMembers() []PersonToTeamRelationship {
	if o == nil || IsNil(o.Members) {
		var ret []PersonToTeamRelationship
		return ret
	}
	return o.Members
}

// GetMembersOk returns a tuple with the Members field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamAllOf) GetMembersOk() ([]PersonToTeamRelationship, bool) {
	if o == nil || IsNil(o.Members) {
		return nil, false
	}
	return o.Members, true
}

// HasMembers returns a boolean if a field has been set.
func (o *TeamAllOf) HasMembers() bool {
	if o != nil && !IsNil(o.Members) {
		return true
	}

	return false
}

// SetMembers gets a reference to the given []PersonToTeamRelationship and assigns it to the Members field.
func (o *TeamAllOf) SetMembers(v []PersonToTeamRelationship) {
	o.Members = v
}

// GetMemberCount returns the MemberCount field value if set, zero value otherwise.
func (o *TeamAllOf) GetMemberCount() int32 {
	if o == nil || IsNil(o.MemberCount) {
		var ret int32
		return ret
	}
	return *o.MemberCount
}

// GetMemberCountOk returns a tuple with the MemberCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamAllOf) GetMemberCountOk() (*int32, bool) {
	if o == nil || IsNil(o.MemberCount) {
		return nil, false
	}
	return o.MemberCount, true
}

// HasMemberCount returns a boolean if a field has been set.
func (o *TeamAllOf) HasMemberCount() bool {
	if o != nil && !IsNil(o.MemberCount) {
		return true
	}

	return false
}

// SetMemberCount gets a reference to the given int32 and assigns it to the MemberCount field.
func (o *TeamAllOf) SetMemberCount(v int32) {
	o.MemberCount = &v
}

// GetEmails returns the Emails field value if set, zero value otherwise.
func (o *TeamAllOf) GetEmails() []TeamEmail {
	if o == nil || IsNil(o.Emails) {
		var ret []TeamEmail
		return ret
	}
	return o.Emails
}

// GetEmailsOk returns a tuple with the Emails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamAllOf) GetEmailsOk() ([]TeamEmail, bool) {
	if o == nil || IsNil(o.Emails) {
		return nil, false
	}
	return o.Emails, true
}

// HasEmails returns a boolean if a field has been set.
func (o *TeamAllOf) HasEmails() bool {
	if o != nil && !IsNil(o.Emails) {
		return true
	}

	return false
}

// SetEmails gets a reference to the given []TeamEmail and assigns it to the Emails field.
func (o *TeamAllOf) SetEmails(v []TeamEmail) {
	o.Emails = v
}

// GetDatasourceProfiles returns the DatasourceProfiles field value if set, zero value otherwise.
func (o *TeamAllOf) GetDatasourceProfiles() []DatasourceProfile {
	if o == nil || IsNil(o.DatasourceProfiles) {
		var ret []DatasourceProfile
		return ret
	}
	return o.DatasourceProfiles
}

// GetDatasourceProfilesOk returns a tuple with the DatasourceProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamAllOf) GetDatasourceProfilesOk() ([]DatasourceProfile, bool) {
	if o == nil || IsNil(o.DatasourceProfiles) {
		return nil, false
	}
	return o.DatasourceProfiles, true
}

// HasDatasourceProfiles returns a boolean if a field has been set.
func (o *TeamAllOf) HasDatasourceProfiles() bool {
	if o != nil && !IsNil(o.DatasourceProfiles) {
		return true
	}

	return false
}

// SetDatasourceProfiles gets a reference to the given []DatasourceProfile and assigns it to the DatasourceProfiles field.
func (o *TeamAllOf) SetDatasourceProfiles(v []DatasourceProfile) {
	o.DatasourceProfiles = v
}

// GetDatasource returns the Datasource field value if set, zero value otherwise.
func (o *TeamAllOf) GetDatasource() string {
	if o == nil || IsNil(o.Datasource) {
		var ret string
		return ret
	}
	return *o.Datasource
}

// GetDatasourceOk returns a tuple with the Datasource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamAllOf) GetDatasourceOk() (*string, bool) {
	if o == nil || IsNil(o.Datasource) {
		return nil, false
	}
	return o.Datasource, true
}

// HasDatasource returns a boolean if a field has been set.
func (o *TeamAllOf) HasDatasource() bool {
	if o != nil && !IsNil(o.Datasource) {
		return true
	}

	return false
}

// SetDatasource gets a reference to the given string and assigns it to the Datasource field.
func (o *TeamAllOf) SetDatasource(v string) {
	o.Datasource = &v
}

// GetCreatedFrom returns the CreatedFrom field value if set, zero value otherwise.
func (o *TeamAllOf) GetCreatedFrom() string {
	if o == nil || IsNil(o.CreatedFrom) {
		var ret string
		return ret
	}
	return *o.CreatedFrom
}

// GetCreatedFromOk returns a tuple with the CreatedFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamAllOf) GetCreatedFromOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedFrom) {
		return nil, false
	}
	return o.CreatedFrom, true
}

// HasCreatedFrom returns a boolean if a field has been set.
func (o *TeamAllOf) HasCreatedFrom() bool {
	if o != nil && !IsNil(o.CreatedFrom) {
		return true
	}

	return false
}

// SetCreatedFrom gets a reference to the given string and assigns it to the CreatedFrom field.
func (o *TeamAllOf) SetCreatedFrom(v string) {
	o.CreatedFrom = &v
}

// GetLastUpdatedAt returns the LastUpdatedAt field value if set, zero value otherwise.
func (o *TeamAllOf) GetLastUpdatedAt() time.Time {
	if o == nil || IsNil(o.LastUpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdatedAt
}

// GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamAllOf) GetLastUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdatedAt) {
		return nil, false
	}
	return o.LastUpdatedAt, true
}

// HasLastUpdatedAt returns a boolean if a field has been set.
func (o *TeamAllOf) HasLastUpdatedAt() bool {
	if o != nil && !IsNil(o.LastUpdatedAt) {
		return true
	}

	return false
}

// SetLastUpdatedAt gets a reference to the given time.Time and assigns it to the LastUpdatedAt field.
func (o *TeamAllOf) SetLastUpdatedAt(v time.Time) {
	o.LastUpdatedAt = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *TeamAllOf) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamAllOf) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *TeamAllOf) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *TeamAllOf) SetStatus(v string) {
	o.Status = &v
}

// GetCanBeDeleted returns the CanBeDeleted field value if set, zero value otherwise.
func (o *TeamAllOf) GetCanBeDeleted() bool {
	if o == nil || IsNil(o.CanBeDeleted) {
		var ret bool
		return ret
	}
	return *o.CanBeDeleted
}

// GetCanBeDeletedOk returns a tuple with the CanBeDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamAllOf) GetCanBeDeletedOk() (*bool, bool) {
	if o == nil || IsNil(o.CanBeDeleted) {
		return nil, false
	}
	return o.CanBeDeleted, true
}

// HasCanBeDeleted returns a boolean if a field has been set.
func (o *TeamAllOf) HasCanBeDeleted() bool {
	if o != nil && !IsNil(o.CanBeDeleted) {
		return true
	}

	return false
}

// SetCanBeDeleted gets a reference to the given bool and assigns it to the CanBeDeleted field.
func (o *TeamAllOf) SetCanBeDeleted(v bool) {
	o.CanBeDeleted = &v
}

// GetLoggingId returns the LoggingId field value if set, zero value otherwise.
func (o *TeamAllOf) GetLoggingId() string {
	if o == nil || IsNil(o.LoggingId) {
		var ret string
		return ret
	}
	return *o.LoggingId
}

// GetLoggingIdOk returns a tuple with the LoggingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamAllOf) GetLoggingIdOk() (*string, bool) {
	if o == nil || IsNil(o.LoggingId) {
		return nil, false
	}
	return o.LoggingId, true
}

// HasLoggingId returns a boolean if a field has been set.
func (o *TeamAllOf) HasLoggingId() bool {
	if o != nil && !IsNil(o.LoggingId) {
		return true
	}

	return false
}

// SetLoggingId gets a reference to the given string and assigns it to the LoggingId field.
func (o *TeamAllOf) SetLoggingId(v string) {
	o.LoggingId = &v
}

func (o TeamAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TeamAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.BusinessUnit) {
		toSerialize["businessUnit"] = o.BusinessUnit
	}
	if !IsNil(o.Department) {
		toSerialize["department"] = o.Department
	}
	if !IsNil(o.PhotoUrl) {
		toSerialize["photoUrl"] = o.PhotoUrl
	}
	if !IsNil(o.BannerUrl) {
		toSerialize["bannerUrl"] = o.BannerUrl
	}
	if !IsNil(o.ExternalLink) {
		toSerialize["externalLink"] = o.ExternalLink
	}
	if !IsNil(o.Members) {
		toSerialize["members"] = o.Members
	}
	if !IsNil(o.MemberCount) {
		toSerialize["memberCount"] = o.MemberCount
	}
	if !IsNil(o.Emails) {
		toSerialize["emails"] = o.Emails
	}
	if !IsNil(o.DatasourceProfiles) {
		toSerialize["datasourceProfiles"] = o.DatasourceProfiles
	}
	if !IsNil(o.Datasource) {
		toSerialize["datasource"] = o.Datasource
	}
	if !IsNil(o.CreatedFrom) {
		toSerialize["createdFrom"] = o.CreatedFrom
	}
	if !IsNil(o.LastUpdatedAt) {
		toSerialize["lastUpdatedAt"] = o.LastUpdatedAt
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.CanBeDeleted) {
		toSerialize["canBeDeleted"] = o.CanBeDeleted
	}
	if !IsNil(o.LoggingId) {
		toSerialize["loggingId"] = o.LoggingId
	}
	return toSerialize, nil
}

type NullableTeamAllOf struct {
	value *TeamAllOf
	isSet bool
}

func (v NullableTeamAllOf) Get() *TeamAllOf {
	return v.value
}

func (v *NullableTeamAllOf) Set(val *TeamAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTeamAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTeamAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeamAllOf(val *TeamAllOf) *NullableTeamAllOf {
	return &NullableTeamAllOf{value: val, isSet: true}
}

func (v NullableTeamAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeamAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

