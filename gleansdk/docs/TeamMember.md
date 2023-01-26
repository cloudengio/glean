# TeamMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | The member&#39;s email | 
**Relationship** | Pointer to **string** | The member&#39;s relationship to the team, an enum of &#x60;MEMBER&#x60;, &#x60;MANAGER&#x60;, &#x60;LEAD&#x60;, &#x60;POINT_OF_CONTACT&#x60;, &#x60;OTHER&#x60; | [optional] [default to "MEMBER"]
**JoinDate** | Pointer to **string** | The member&#39;s start date | [optional] 

## Methods

### NewTeamMember

`func NewTeamMember(email string, ) *TeamMember`

NewTeamMember instantiates a new TeamMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamMemberWithDefaults

`func NewTeamMemberWithDefaults() *TeamMember`

NewTeamMemberWithDefaults instantiates a new TeamMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *TeamMember) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *TeamMember) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *TeamMember) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetRelationship

`func (o *TeamMember) GetRelationship() string`

GetRelationship returns the Relationship field if non-nil, zero value otherwise.

### GetRelationshipOk

`func (o *TeamMember) GetRelationshipOk() (*string, bool)`

GetRelationshipOk returns a tuple with the Relationship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationship

`func (o *TeamMember) SetRelationship(v string)`

SetRelationship sets Relationship field to given value.

### HasRelationship

`func (o *TeamMember) HasRelationship() bool`

HasRelationship returns a boolean if a field has been set.

### GetJoinDate

`func (o *TeamMember) GetJoinDate() string`

GetJoinDate returns the JoinDate field if non-nil, zero value otherwise.

### GetJoinDateOk

`func (o *TeamMember) GetJoinDateOk() (*string, bool)`

GetJoinDateOk returns a tuple with the JoinDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinDate

`func (o *TeamMember) SetJoinDate(v string)`

SetJoinDate sets JoinDate field to given value.

### HasJoinDate

`func (o *TeamMember) HasJoinDate() bool`

HasJoinDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


