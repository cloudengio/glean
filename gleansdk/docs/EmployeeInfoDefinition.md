# EmployeeInfoDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | The employee&#39;s email | 
**FirstName** | Pointer to **string** | The first name of the employee | [optional] 
**LastName** | Pointer to **string** | The last name of the employee | [optional] 
**PreferredName** | Pointer to **string** | The preferred name or nickname of the employee | [optional] 
**Id** | Pointer to **string** | **[Advanced]** A unique universal internal identifier for the employee. This is solely used for understanding manager relationships along with &#x60;managerId&#x60;.  | [optional] 
**PhoneNumber** | Pointer to **string** | The employee&#39;s phone number. | [optional] 
**Location** | Pointer to **string** | The employee&#39;s location (city/office name etc). | [optional] 
**StructuredLocation** | Pointer to [**StructuredLocation**](StructuredLocation.md) |  | [optional] 
**Title** | Pointer to **string** | The employee&#39;s role title. | [optional] 
**PhotoUrl** | Pointer to **string** | The employee&#39;s profile pic | [optional] 
**BusinessUnit** | Pointer to **string** | Typically the highest level organizational unit; generally applies to bigger companies with multiple distinct businesses. | [optional] 
**Department** | **string** | An organizational unit where everyone has a similar task, e.g. &#x60;Engineering&#x60;. | 
**DatasourceProfiles** | Pointer to [**[]DatasourceProfile**](DatasourceProfile.md) | The datasource profiles of the employee, e.g. &#x60;Slack&#x60;,&#x60;Github&#x60;. | [optional] 
**Teams** | Pointer to [**[]EmployeeTeamInfo**](EmployeeTeamInfo.md) | Info about the employee&#39;s team(s) | [optional] 
**StartDate** | Pointer to **string** | The date when the employee started | [optional] 
**EndDate** | Pointer to **string** | If a former employee, the last date of employment. | [optional] 
**Bio** | Pointer to **string** | Short biography or mission statement of the employee. | [optional] 
**Pronoun** | Pointer to **string** | She/her, He/his or other pronoun. | [optional] 
**AlsoKnownAs** | Pointer to **[]string** | Other names associated with the employee. | [optional] 
**ProfileUrl** | Pointer to **string** | Link to internal company person profile. | [optional] 
**SocialNetworks** | Pointer to [**[]SocialNetworkDefinition**](SocialNetworkDefinition.md) | List of social network profiles. | [optional] 
**ManagerEmail** | Pointer to **string** | The email of the employee&#39;s manager | [optional] 
**ManagerId** | Pointer to **string** | **[Advanced]** A unique universal internal identifier for the employee&#39;s manager. This is solely used in conjunction with &#x60;id&#x60;.  | [optional] 
**Status** | Pointer to **string** | The status of the employee, an enum of &#x60;CURRENT&#x60;, &#x60;FUTURE&#x60;, &#x60;EX&#x60; | [optional] [default to "CURRENT"]
**AdditionalFields** | Pointer to [**[]AdditionalFieldDefinition**](AdditionalFieldDefinition.md) | List of additional fields with more information about the employee. | [optional] 

## Methods

### NewEmployeeInfoDefinition

`func NewEmployeeInfoDefinition(email string, department string, ) *EmployeeInfoDefinition`

NewEmployeeInfoDefinition instantiates a new EmployeeInfoDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmployeeInfoDefinitionWithDefaults

`func NewEmployeeInfoDefinitionWithDefaults() *EmployeeInfoDefinition`

NewEmployeeInfoDefinitionWithDefaults instantiates a new EmployeeInfoDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *EmployeeInfoDefinition) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *EmployeeInfoDefinition) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *EmployeeInfoDefinition) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetFirstName

`func (o *EmployeeInfoDefinition) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *EmployeeInfoDefinition) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *EmployeeInfoDefinition) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *EmployeeInfoDefinition) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *EmployeeInfoDefinition) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *EmployeeInfoDefinition) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *EmployeeInfoDefinition) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *EmployeeInfoDefinition) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetPreferredName

`func (o *EmployeeInfoDefinition) GetPreferredName() string`

GetPreferredName returns the PreferredName field if non-nil, zero value otherwise.

### GetPreferredNameOk

`func (o *EmployeeInfoDefinition) GetPreferredNameOk() (*string, bool)`

GetPreferredNameOk returns a tuple with the PreferredName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredName

`func (o *EmployeeInfoDefinition) SetPreferredName(v string)`

SetPreferredName sets PreferredName field to given value.

### HasPreferredName

`func (o *EmployeeInfoDefinition) HasPreferredName() bool`

HasPreferredName returns a boolean if a field has been set.

### GetId

`func (o *EmployeeInfoDefinition) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EmployeeInfoDefinition) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EmployeeInfoDefinition) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EmployeeInfoDefinition) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *EmployeeInfoDefinition) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *EmployeeInfoDefinition) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *EmployeeInfoDefinition) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *EmployeeInfoDefinition) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetLocation

`func (o *EmployeeInfoDefinition) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *EmployeeInfoDefinition) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *EmployeeInfoDefinition) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *EmployeeInfoDefinition) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetStructuredLocation

`func (o *EmployeeInfoDefinition) GetStructuredLocation() StructuredLocation`

GetStructuredLocation returns the StructuredLocation field if non-nil, zero value otherwise.

### GetStructuredLocationOk

`func (o *EmployeeInfoDefinition) GetStructuredLocationOk() (*StructuredLocation, bool)`

GetStructuredLocationOk returns a tuple with the StructuredLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructuredLocation

`func (o *EmployeeInfoDefinition) SetStructuredLocation(v StructuredLocation)`

SetStructuredLocation sets StructuredLocation field to given value.

### HasStructuredLocation

`func (o *EmployeeInfoDefinition) HasStructuredLocation() bool`

HasStructuredLocation returns a boolean if a field has been set.

### GetTitle

`func (o *EmployeeInfoDefinition) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *EmployeeInfoDefinition) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *EmployeeInfoDefinition) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *EmployeeInfoDefinition) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetPhotoUrl

`func (o *EmployeeInfoDefinition) GetPhotoUrl() string`

GetPhotoUrl returns the PhotoUrl field if non-nil, zero value otherwise.

### GetPhotoUrlOk

`func (o *EmployeeInfoDefinition) GetPhotoUrlOk() (*string, bool)`

GetPhotoUrlOk returns a tuple with the PhotoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotoUrl

`func (o *EmployeeInfoDefinition) SetPhotoUrl(v string)`

SetPhotoUrl sets PhotoUrl field to given value.

### HasPhotoUrl

`func (o *EmployeeInfoDefinition) HasPhotoUrl() bool`

HasPhotoUrl returns a boolean if a field has been set.

### GetBusinessUnit

`func (o *EmployeeInfoDefinition) GetBusinessUnit() string`

GetBusinessUnit returns the BusinessUnit field if non-nil, zero value otherwise.

### GetBusinessUnitOk

`func (o *EmployeeInfoDefinition) GetBusinessUnitOk() (*string, bool)`

GetBusinessUnitOk returns a tuple with the BusinessUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessUnit

`func (o *EmployeeInfoDefinition) SetBusinessUnit(v string)`

SetBusinessUnit sets BusinessUnit field to given value.

### HasBusinessUnit

`func (o *EmployeeInfoDefinition) HasBusinessUnit() bool`

HasBusinessUnit returns a boolean if a field has been set.

### GetDepartment

`func (o *EmployeeInfoDefinition) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *EmployeeInfoDefinition) GetDepartmentOk() (*string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *EmployeeInfoDefinition) SetDepartment(v string)`

SetDepartment sets Department field to given value.


### GetDatasourceProfiles

`func (o *EmployeeInfoDefinition) GetDatasourceProfiles() []DatasourceProfile`

GetDatasourceProfiles returns the DatasourceProfiles field if non-nil, zero value otherwise.

### GetDatasourceProfilesOk

`func (o *EmployeeInfoDefinition) GetDatasourceProfilesOk() (*[]DatasourceProfile, bool)`

GetDatasourceProfilesOk returns a tuple with the DatasourceProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasourceProfiles

`func (o *EmployeeInfoDefinition) SetDatasourceProfiles(v []DatasourceProfile)`

SetDatasourceProfiles sets DatasourceProfiles field to given value.

### HasDatasourceProfiles

`func (o *EmployeeInfoDefinition) HasDatasourceProfiles() bool`

HasDatasourceProfiles returns a boolean if a field has been set.

### GetTeams

`func (o *EmployeeInfoDefinition) GetTeams() []EmployeeTeamInfo`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *EmployeeInfoDefinition) GetTeamsOk() (*[]EmployeeTeamInfo, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *EmployeeInfoDefinition) SetTeams(v []EmployeeTeamInfo)`

SetTeams sets Teams field to given value.

### HasTeams

`func (o *EmployeeInfoDefinition) HasTeams() bool`

HasTeams returns a boolean if a field has been set.

### GetStartDate

`func (o *EmployeeInfoDefinition) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *EmployeeInfoDefinition) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *EmployeeInfoDefinition) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *EmployeeInfoDefinition) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *EmployeeInfoDefinition) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *EmployeeInfoDefinition) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *EmployeeInfoDefinition) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *EmployeeInfoDefinition) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetBio

`func (o *EmployeeInfoDefinition) GetBio() string`

GetBio returns the Bio field if non-nil, zero value otherwise.

### GetBioOk

`func (o *EmployeeInfoDefinition) GetBioOk() (*string, bool)`

GetBioOk returns a tuple with the Bio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBio

`func (o *EmployeeInfoDefinition) SetBio(v string)`

SetBio sets Bio field to given value.

### HasBio

`func (o *EmployeeInfoDefinition) HasBio() bool`

HasBio returns a boolean if a field has been set.

### GetPronoun

`func (o *EmployeeInfoDefinition) GetPronoun() string`

GetPronoun returns the Pronoun field if non-nil, zero value otherwise.

### GetPronounOk

`func (o *EmployeeInfoDefinition) GetPronounOk() (*string, bool)`

GetPronounOk returns a tuple with the Pronoun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPronoun

`func (o *EmployeeInfoDefinition) SetPronoun(v string)`

SetPronoun sets Pronoun field to given value.

### HasPronoun

`func (o *EmployeeInfoDefinition) HasPronoun() bool`

HasPronoun returns a boolean if a field has been set.

### GetAlsoKnownAs

`func (o *EmployeeInfoDefinition) GetAlsoKnownAs() []string`

GetAlsoKnownAs returns the AlsoKnownAs field if non-nil, zero value otherwise.

### GetAlsoKnownAsOk

`func (o *EmployeeInfoDefinition) GetAlsoKnownAsOk() (*[]string, bool)`

GetAlsoKnownAsOk returns a tuple with the AlsoKnownAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlsoKnownAs

`func (o *EmployeeInfoDefinition) SetAlsoKnownAs(v []string)`

SetAlsoKnownAs sets AlsoKnownAs field to given value.

### HasAlsoKnownAs

`func (o *EmployeeInfoDefinition) HasAlsoKnownAs() bool`

HasAlsoKnownAs returns a boolean if a field has been set.

### GetProfileUrl

`func (o *EmployeeInfoDefinition) GetProfileUrl() string`

GetProfileUrl returns the ProfileUrl field if non-nil, zero value otherwise.

### GetProfileUrlOk

`func (o *EmployeeInfoDefinition) GetProfileUrlOk() (*string, bool)`

GetProfileUrlOk returns a tuple with the ProfileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileUrl

`func (o *EmployeeInfoDefinition) SetProfileUrl(v string)`

SetProfileUrl sets ProfileUrl field to given value.

### HasProfileUrl

`func (o *EmployeeInfoDefinition) HasProfileUrl() bool`

HasProfileUrl returns a boolean if a field has been set.

### GetSocialNetworks

`func (o *EmployeeInfoDefinition) GetSocialNetworks() []SocialNetworkDefinition`

GetSocialNetworks returns the SocialNetworks field if non-nil, zero value otherwise.

### GetSocialNetworksOk

`func (o *EmployeeInfoDefinition) GetSocialNetworksOk() (*[]SocialNetworkDefinition, bool)`

GetSocialNetworksOk returns a tuple with the SocialNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialNetworks

`func (o *EmployeeInfoDefinition) SetSocialNetworks(v []SocialNetworkDefinition)`

SetSocialNetworks sets SocialNetworks field to given value.

### HasSocialNetworks

`func (o *EmployeeInfoDefinition) HasSocialNetworks() bool`

HasSocialNetworks returns a boolean if a field has been set.

### GetManagerEmail

`func (o *EmployeeInfoDefinition) GetManagerEmail() string`

GetManagerEmail returns the ManagerEmail field if non-nil, zero value otherwise.

### GetManagerEmailOk

`func (o *EmployeeInfoDefinition) GetManagerEmailOk() (*string, bool)`

GetManagerEmailOk returns a tuple with the ManagerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagerEmail

`func (o *EmployeeInfoDefinition) SetManagerEmail(v string)`

SetManagerEmail sets ManagerEmail field to given value.

### HasManagerEmail

`func (o *EmployeeInfoDefinition) HasManagerEmail() bool`

HasManagerEmail returns a boolean if a field has been set.

### GetManagerId

`func (o *EmployeeInfoDefinition) GetManagerId() string`

GetManagerId returns the ManagerId field if non-nil, zero value otherwise.

### GetManagerIdOk

`func (o *EmployeeInfoDefinition) GetManagerIdOk() (*string, bool)`

GetManagerIdOk returns a tuple with the ManagerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagerId

`func (o *EmployeeInfoDefinition) SetManagerId(v string)`

SetManagerId sets ManagerId field to given value.

### HasManagerId

`func (o *EmployeeInfoDefinition) HasManagerId() bool`

HasManagerId returns a boolean if a field has been set.

### GetStatus

`func (o *EmployeeInfoDefinition) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EmployeeInfoDefinition) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EmployeeInfoDefinition) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EmployeeInfoDefinition) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *EmployeeInfoDefinition) GetAdditionalFields() []AdditionalFieldDefinition`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *EmployeeInfoDefinition) GetAdditionalFieldsOk() (*[]AdditionalFieldDefinition, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *EmployeeInfoDefinition) SetAdditionalFields(v []AdditionalFieldDefinition)`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *EmployeeInfoDefinition) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


