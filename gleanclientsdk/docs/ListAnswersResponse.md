# ListAnswersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnswerResults** | [**[]AnswerResult**](AnswerResult.md) | List of answers with tracking tokens. | 
**UserRole** | Pointer to **string** | DEPRECATED - use permissions instead. User&#39;s role for Answers in their workplace. | [optional] 

## Methods

### NewListAnswersResponse

`func NewListAnswersResponse(answerResults []AnswerResult, ) *ListAnswersResponse`

NewListAnswersResponse instantiates a new ListAnswersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAnswersResponseWithDefaults

`func NewListAnswersResponseWithDefaults() *ListAnswersResponse`

NewListAnswersResponseWithDefaults instantiates a new ListAnswersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnswerResults

`func (o *ListAnswersResponse) GetAnswerResults() []AnswerResult`

GetAnswerResults returns the AnswerResults field if non-nil, zero value otherwise.

### GetAnswerResultsOk

`func (o *ListAnswersResponse) GetAnswerResultsOk() (*[]AnswerResult, bool)`

GetAnswerResultsOk returns a tuple with the AnswerResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswerResults

`func (o *ListAnswersResponse) SetAnswerResults(v []AnswerResult)`

SetAnswerResults sets AnswerResults field to given value.


### GetUserRole

`func (o *ListAnswersResponse) GetUserRole() string`

GetUserRole returns the UserRole field if non-nil, zero value otherwise.

### GetUserRoleOk

`func (o *ListAnswersResponse) GetUserRoleOk() (*string, bool)`

GetUserRoleOk returns a tuple with the UserRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserRole

`func (o *ListAnswersResponse) SetUserRole(v string)`

SetUserRole sets UserRole field to given value.

### HasUserRole

`func (o *ListAnswersResponse) HasUserRole() bool`

HasUserRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


