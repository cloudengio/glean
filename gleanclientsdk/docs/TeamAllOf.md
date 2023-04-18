# TeamAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier | 
**Name** | **string** | Team name | 
**Description** | Pointer to **string** | A description of the team | [optional] 
**BusinessUnit** | Pointer to **string** | Typically the highest level organizational unit; generally applies to bigger companies with multiple distinct businesses. | [optional] 
**Department** | Pointer to **string** | An organizational unit where everyone has a similar task, e.g. &#x60;Engineering&#x60;. | [optional] 
**PhotoUrl** | Pointer to **string** | A link to the team&#39;s photo | [optional] 
**BannerUrl** | Pointer to **string** | A link to the team&#39;s banner photo | [optional] 
**ExternalLink** | Pointer to **string** | Link to a team page on the internet or your company&#39;s intranet | [optional] 
**Members** | Pointer to [**[]PersonToTeamRelationship**](PersonToTeamRelationship.md) | The members on this team | [optional] 
**MemberCount** | Pointer to **int32** | Number of members on this team (recursive; includes all individuals that belong to this team, and all individuals that belong to a subteam within this team) | [optional] 
**Emails** | Pointer to [**[]TeamEmail**](TeamEmail.md) | The emails for this team | [optional] 
**DatasourceProfiles** | Pointer to [**[]DatasourceProfile**](DatasourceProfile.md) | The datasource profiles of the team | [optional] 
**Datasource** | Pointer to **string** | the data source of the team, e.g. GDRIVE | [optional] 
**CreatedFrom** | Pointer to **string** | For teams created from docs, the doc title. Otherwise empty. | [optional] 
**LastUpdatedAt** | Pointer to **time.Time** | when this team was last updated. | [optional] 
**Status** | Pointer to **string** | whether this team is fully processed or there are still unprocessed operations that&#39;ll affect it | [optional] [default to "PROCESSED"]
**CanBeDeleted** | Pointer to **bool** | can this team be deleted. Some manually ingested teams like GCS_CSV or PUSH_API cannot | [optional] [default to true]
**LoggingId** | Pointer to **string** | The logging id of the team used in scrubbed logs, client analytics, and metrics. | [optional] 

## Methods

### NewTeamAllOf

`func NewTeamAllOf(id string, name string, ) *TeamAllOf`

NewTeamAllOf instantiates a new TeamAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamAllOfWithDefaults

`func NewTeamAllOfWithDefaults() *TeamAllOf`

NewTeamAllOfWithDefaults instantiates a new TeamAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TeamAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TeamAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TeamAllOf) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *TeamAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TeamAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TeamAllOf) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *TeamAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TeamAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TeamAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TeamAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetBusinessUnit

`func (o *TeamAllOf) GetBusinessUnit() string`

GetBusinessUnit returns the BusinessUnit field if non-nil, zero value otherwise.

### GetBusinessUnitOk

`func (o *TeamAllOf) GetBusinessUnitOk() (*string, bool)`

GetBusinessUnitOk returns a tuple with the BusinessUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessUnit

`func (o *TeamAllOf) SetBusinessUnit(v string)`

SetBusinessUnit sets BusinessUnit field to given value.

### HasBusinessUnit

`func (o *TeamAllOf) HasBusinessUnit() bool`

HasBusinessUnit returns a boolean if a field has been set.

### GetDepartment

`func (o *TeamAllOf) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *TeamAllOf) GetDepartmentOk() (*string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *TeamAllOf) SetDepartment(v string)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *TeamAllOf) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### GetPhotoUrl

`func (o *TeamAllOf) GetPhotoUrl() string`

GetPhotoUrl returns the PhotoUrl field if non-nil, zero value otherwise.

### GetPhotoUrlOk

`func (o *TeamAllOf) GetPhotoUrlOk() (*string, bool)`

GetPhotoUrlOk returns a tuple with the PhotoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotoUrl

`func (o *TeamAllOf) SetPhotoUrl(v string)`

SetPhotoUrl sets PhotoUrl field to given value.

### HasPhotoUrl

`func (o *TeamAllOf) HasPhotoUrl() bool`

HasPhotoUrl returns a boolean if a field has been set.

### GetBannerUrl

`func (o *TeamAllOf) GetBannerUrl() string`

GetBannerUrl returns the BannerUrl field if non-nil, zero value otherwise.

### GetBannerUrlOk

`func (o *TeamAllOf) GetBannerUrlOk() (*string, bool)`

GetBannerUrlOk returns a tuple with the BannerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannerUrl

`func (o *TeamAllOf) SetBannerUrl(v string)`

SetBannerUrl sets BannerUrl field to given value.

### HasBannerUrl

`func (o *TeamAllOf) HasBannerUrl() bool`

HasBannerUrl returns a boolean if a field has been set.

### GetExternalLink

`func (o *TeamAllOf) GetExternalLink() string`

GetExternalLink returns the ExternalLink field if non-nil, zero value otherwise.

### GetExternalLinkOk

`func (o *TeamAllOf) GetExternalLinkOk() (*string, bool)`

GetExternalLinkOk returns a tuple with the ExternalLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalLink

`func (o *TeamAllOf) SetExternalLink(v string)`

SetExternalLink sets ExternalLink field to given value.

### HasExternalLink

`func (o *TeamAllOf) HasExternalLink() bool`

HasExternalLink returns a boolean if a field has been set.

### GetMembers

`func (o *TeamAllOf) GetMembers() []PersonToTeamRelationship`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *TeamAllOf) GetMembersOk() (*[]PersonToTeamRelationship, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *TeamAllOf) SetMembers(v []PersonToTeamRelationship)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *TeamAllOf) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetMemberCount

`func (o *TeamAllOf) GetMemberCount() int32`

GetMemberCount returns the MemberCount field if non-nil, zero value otherwise.

### GetMemberCountOk

`func (o *TeamAllOf) GetMemberCountOk() (*int32, bool)`

GetMemberCountOk returns a tuple with the MemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberCount

`func (o *TeamAllOf) SetMemberCount(v int32)`

SetMemberCount sets MemberCount field to given value.

### HasMemberCount

`func (o *TeamAllOf) HasMemberCount() bool`

HasMemberCount returns a boolean if a field has been set.

### GetEmails

`func (o *TeamAllOf) GetEmails() []TeamEmail`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *TeamAllOf) GetEmailsOk() (*[]TeamEmail, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *TeamAllOf) SetEmails(v []TeamEmail)`

SetEmails sets Emails field to given value.

### HasEmails

`func (o *TeamAllOf) HasEmails() bool`

HasEmails returns a boolean if a field has been set.

### GetDatasourceProfiles

`func (o *TeamAllOf) GetDatasourceProfiles() []DatasourceProfile`

GetDatasourceProfiles returns the DatasourceProfiles field if non-nil, zero value otherwise.

### GetDatasourceProfilesOk

`func (o *TeamAllOf) GetDatasourceProfilesOk() (*[]DatasourceProfile, bool)`

GetDatasourceProfilesOk returns a tuple with the DatasourceProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasourceProfiles

`func (o *TeamAllOf) SetDatasourceProfiles(v []DatasourceProfile)`

SetDatasourceProfiles sets DatasourceProfiles field to given value.

### HasDatasourceProfiles

`func (o *TeamAllOf) HasDatasourceProfiles() bool`

HasDatasourceProfiles returns a boolean if a field has been set.

### GetDatasource

`func (o *TeamAllOf) GetDatasource() string`

GetDatasource returns the Datasource field if non-nil, zero value otherwise.

### GetDatasourceOk

`func (o *TeamAllOf) GetDatasourceOk() (*string, bool)`

GetDatasourceOk returns a tuple with the Datasource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasource

`func (o *TeamAllOf) SetDatasource(v string)`

SetDatasource sets Datasource field to given value.

### HasDatasource

`func (o *TeamAllOf) HasDatasource() bool`

HasDatasource returns a boolean if a field has been set.

### GetCreatedFrom

`func (o *TeamAllOf) GetCreatedFrom() string`

GetCreatedFrom returns the CreatedFrom field if non-nil, zero value otherwise.

### GetCreatedFromOk

`func (o *TeamAllOf) GetCreatedFromOk() (*string, bool)`

GetCreatedFromOk returns a tuple with the CreatedFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedFrom

`func (o *TeamAllOf) SetCreatedFrom(v string)`

SetCreatedFrom sets CreatedFrom field to given value.

### HasCreatedFrom

`func (o *TeamAllOf) HasCreatedFrom() bool`

HasCreatedFrom returns a boolean if a field has been set.

### GetLastUpdatedAt

`func (o *TeamAllOf) GetLastUpdatedAt() time.Time`

GetLastUpdatedAt returns the LastUpdatedAt field if non-nil, zero value otherwise.

### GetLastUpdatedAtOk

`func (o *TeamAllOf) GetLastUpdatedAtOk() (*time.Time, bool)`

GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedAt

`func (o *TeamAllOf) SetLastUpdatedAt(v time.Time)`

SetLastUpdatedAt sets LastUpdatedAt field to given value.

### HasLastUpdatedAt

`func (o *TeamAllOf) HasLastUpdatedAt() bool`

HasLastUpdatedAt returns a boolean if a field has been set.

### GetStatus

`func (o *TeamAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TeamAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TeamAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TeamAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCanBeDeleted

`func (o *TeamAllOf) GetCanBeDeleted() bool`

GetCanBeDeleted returns the CanBeDeleted field if non-nil, zero value otherwise.

### GetCanBeDeletedOk

`func (o *TeamAllOf) GetCanBeDeletedOk() (*bool, bool)`

GetCanBeDeletedOk returns a tuple with the CanBeDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanBeDeleted

`func (o *TeamAllOf) SetCanBeDeleted(v bool)`

SetCanBeDeleted sets CanBeDeleted field to given value.

### HasCanBeDeleted

`func (o *TeamAllOf) HasCanBeDeleted() bool`

HasCanBeDeleted returns a boolean if a field has been set.

### GetLoggingId

`func (o *TeamAllOf) GetLoggingId() string`

GetLoggingId returns the LoggingId field if non-nil, zero value otherwise.

### GetLoggingIdOk

`func (o *TeamAllOf) GetLoggingIdOk() (*string, bool)`

GetLoggingIdOk returns a tuple with the LoggingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggingId

`func (o *TeamAllOf) SetLoggingId(v string)`

SetLoggingId sets LoggingId field to given value.

### HasLoggingId

`func (o *TeamAllOf) HasLoggingId() bool`

HasLoggingId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


