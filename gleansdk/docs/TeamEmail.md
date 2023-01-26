# TeamEmail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | An email address | 
**Type** | **string** | An enum of &#x60;PRIMARY&#x60;, &#x60;SECONDARY&#x60;, &#x60;ONCALL&#x60;, &#x60;OTHER&#x60; | [default to "OTHER"]

## Methods

### NewTeamEmail

`func NewTeamEmail(email string, type_ string, ) *TeamEmail`

NewTeamEmail instantiates a new TeamEmail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamEmailWithDefaults

`func NewTeamEmailWithDefaults() *TeamEmail`

NewTeamEmailWithDefaults instantiates a new TeamEmail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *TeamEmail) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *TeamEmail) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *TeamEmail) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetType

`func (o *TeamEmail) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TeamEmail) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TeamEmail) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


