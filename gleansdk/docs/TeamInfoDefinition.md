# TeamInfoDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique ID of the team | 
**Name** | **string** | Human-readable team name | 
**Description** | Pointer to **string** | The description of this team | [optional] 
**BusinessUnit** | Pointer to **string** | Typically the highest level organizational unit; generally applies to bigger companies with multiple distinct businesses. | [optional] 
**Department** | Pointer to **string** | An organizational unit where everyone has a similar task, e.g. &#x60;Engineering&#x60;. | [optional] 
**PhotoUrl** | Pointer to **string** | A link to the team&#39;s photo | [optional] 
**ExternalLink** | Pointer to **string** | A link to an external team page. If set, team results will link to it.  | [optional] 
**Emails** | Pointer to [**[]TeamEmail**](TeamEmail.md) | The emails of the team | [optional] 
**DatasourceProfiles** | Pointer to [**[]DatasourceProfile**](DatasourceProfile.md) | The datasource profiles of the team | [optional] 
**Members** | [**[]TeamMember**](TeamMember.md) | The members of the team | 

## Methods

### NewTeamInfoDefinition

`func NewTeamInfoDefinition(id string, name string, members []TeamMember, ) *TeamInfoDefinition`

NewTeamInfoDefinition instantiates a new TeamInfoDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamInfoDefinitionWithDefaults

`func NewTeamInfoDefinitionWithDefaults() *TeamInfoDefinition`

NewTeamInfoDefinitionWithDefaults instantiates a new TeamInfoDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TeamInfoDefinition) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TeamInfoDefinition) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TeamInfoDefinition) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *TeamInfoDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TeamInfoDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TeamInfoDefinition) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *TeamInfoDefinition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TeamInfoDefinition) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TeamInfoDefinition) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TeamInfoDefinition) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetBusinessUnit

`func (o *TeamInfoDefinition) GetBusinessUnit() string`

GetBusinessUnit returns the BusinessUnit field if non-nil, zero value otherwise.

### GetBusinessUnitOk

`func (o *TeamInfoDefinition) GetBusinessUnitOk() (*string, bool)`

GetBusinessUnitOk returns a tuple with the BusinessUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessUnit

`func (o *TeamInfoDefinition) SetBusinessUnit(v string)`

SetBusinessUnit sets BusinessUnit field to given value.

### HasBusinessUnit

`func (o *TeamInfoDefinition) HasBusinessUnit() bool`

HasBusinessUnit returns a boolean if a field has been set.

### GetDepartment

`func (o *TeamInfoDefinition) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *TeamInfoDefinition) GetDepartmentOk() (*string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *TeamInfoDefinition) SetDepartment(v string)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *TeamInfoDefinition) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### GetPhotoUrl

`func (o *TeamInfoDefinition) GetPhotoUrl() string`

GetPhotoUrl returns the PhotoUrl field if non-nil, zero value otherwise.

### GetPhotoUrlOk

`func (o *TeamInfoDefinition) GetPhotoUrlOk() (*string, bool)`

GetPhotoUrlOk returns a tuple with the PhotoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotoUrl

`func (o *TeamInfoDefinition) SetPhotoUrl(v string)`

SetPhotoUrl sets PhotoUrl field to given value.

### HasPhotoUrl

`func (o *TeamInfoDefinition) HasPhotoUrl() bool`

HasPhotoUrl returns a boolean if a field has been set.

### GetExternalLink

`func (o *TeamInfoDefinition) GetExternalLink() string`

GetExternalLink returns the ExternalLink field if non-nil, zero value otherwise.

### GetExternalLinkOk

`func (o *TeamInfoDefinition) GetExternalLinkOk() (*string, bool)`

GetExternalLinkOk returns a tuple with the ExternalLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalLink

`func (o *TeamInfoDefinition) SetExternalLink(v string)`

SetExternalLink sets ExternalLink field to given value.

### HasExternalLink

`func (o *TeamInfoDefinition) HasExternalLink() bool`

HasExternalLink returns a boolean if a field has been set.

### GetEmails

`func (o *TeamInfoDefinition) GetEmails() []TeamEmail`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *TeamInfoDefinition) GetEmailsOk() (*[]TeamEmail, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *TeamInfoDefinition) SetEmails(v []TeamEmail)`

SetEmails sets Emails field to given value.

### HasEmails

`func (o *TeamInfoDefinition) HasEmails() bool`

HasEmails returns a boolean if a field has been set.

### GetDatasourceProfiles

`func (o *TeamInfoDefinition) GetDatasourceProfiles() []DatasourceProfile`

GetDatasourceProfiles returns the DatasourceProfiles field if non-nil, zero value otherwise.

### GetDatasourceProfilesOk

`func (o *TeamInfoDefinition) GetDatasourceProfilesOk() (*[]DatasourceProfile, bool)`

GetDatasourceProfilesOk returns a tuple with the DatasourceProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasourceProfiles

`func (o *TeamInfoDefinition) SetDatasourceProfiles(v []DatasourceProfile)`

SetDatasourceProfiles sets DatasourceProfiles field to given value.

### HasDatasourceProfiles

`func (o *TeamInfoDefinition) HasDatasourceProfiles() bool`

HasDatasourceProfiles returns a boolean if a field has been set.

### GetMembers

`func (o *TeamInfoDefinition) GetMembers() []TeamMember`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *TeamInfoDefinition) GetMembersOk() (*[]TeamMember, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *TeamInfoDefinition) SetMembers(v []TeamMember)`

SetMembers sets Members field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


