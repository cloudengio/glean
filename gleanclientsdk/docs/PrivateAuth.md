# PrivateAuth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | The authorization status of the user. | [optional] 
**IsDomainDelegated** | Pointer to **bool** | Whether the datasource is authorized through domain delegation. | [optional] 
**ResultsAuthPromptSubtitle** | Pointer to **string** | The subtitle of the results auth prompt, required when the datasource has results even if the user has not authed privately. | [optional] 
**Title** | Pointer to **string** | The title of the auth prompt shown to the user | [optional] 
**Content** | Pointer to **string** | The content of the auth prompt shown to the user | [optional] 
**Questions** | Pointer to **[][]string** | FAQs to show when asking the user to authorize the datasource | [optional] 
**AuthUrlRelativePath** | Pointer to **string** | The relative path of the url to take the user to to authorize the datasource | [optional] 

## Methods

### NewPrivateAuth

`func NewPrivateAuth() *PrivateAuth`

NewPrivateAuth instantiates a new PrivateAuth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateAuthWithDefaults

`func NewPrivateAuthWithDefaults() *PrivateAuth`

NewPrivateAuthWithDefaults instantiates a new PrivateAuth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *PrivateAuth) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PrivateAuth) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PrivateAuth) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PrivateAuth) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetIsDomainDelegated

`func (o *PrivateAuth) GetIsDomainDelegated() bool`

GetIsDomainDelegated returns the IsDomainDelegated field if non-nil, zero value otherwise.

### GetIsDomainDelegatedOk

`func (o *PrivateAuth) GetIsDomainDelegatedOk() (*bool, bool)`

GetIsDomainDelegatedOk returns a tuple with the IsDomainDelegated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDomainDelegated

`func (o *PrivateAuth) SetIsDomainDelegated(v bool)`

SetIsDomainDelegated sets IsDomainDelegated field to given value.

### HasIsDomainDelegated

`func (o *PrivateAuth) HasIsDomainDelegated() bool`

HasIsDomainDelegated returns a boolean if a field has been set.

### GetResultsAuthPromptSubtitle

`func (o *PrivateAuth) GetResultsAuthPromptSubtitle() string`

GetResultsAuthPromptSubtitle returns the ResultsAuthPromptSubtitle field if non-nil, zero value otherwise.

### GetResultsAuthPromptSubtitleOk

`func (o *PrivateAuth) GetResultsAuthPromptSubtitleOk() (*string, bool)`

GetResultsAuthPromptSubtitleOk returns a tuple with the ResultsAuthPromptSubtitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsAuthPromptSubtitle

`func (o *PrivateAuth) SetResultsAuthPromptSubtitle(v string)`

SetResultsAuthPromptSubtitle sets ResultsAuthPromptSubtitle field to given value.

### HasResultsAuthPromptSubtitle

`func (o *PrivateAuth) HasResultsAuthPromptSubtitle() bool`

HasResultsAuthPromptSubtitle returns a boolean if a field has been set.

### GetTitle

`func (o *PrivateAuth) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PrivateAuth) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PrivateAuth) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *PrivateAuth) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetContent

`func (o *PrivateAuth) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *PrivateAuth) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *PrivateAuth) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *PrivateAuth) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetQuestions

`func (o *PrivateAuth) GetQuestions() [][]string`

GetQuestions returns the Questions field if non-nil, zero value otherwise.

### GetQuestionsOk

`func (o *PrivateAuth) GetQuestionsOk() (*[][]string, bool)`

GetQuestionsOk returns a tuple with the Questions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestions

`func (o *PrivateAuth) SetQuestions(v [][]string)`

SetQuestions sets Questions field to given value.

### HasQuestions

`func (o *PrivateAuth) HasQuestions() bool`

HasQuestions returns a boolean if a field has been set.

### GetAuthUrlRelativePath

`func (o *PrivateAuth) GetAuthUrlRelativePath() string`

GetAuthUrlRelativePath returns the AuthUrlRelativePath field if non-nil, zero value otherwise.

### GetAuthUrlRelativePathOk

`func (o *PrivateAuth) GetAuthUrlRelativePathOk() (*string, bool)`

GetAuthUrlRelativePathOk returns a tuple with the AuthUrlRelativePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUrlRelativePath

`func (o *PrivateAuth) SetAuthUrlRelativePath(v string)`

SetAuthUrlRelativePath sets AuthUrlRelativePath field to given value.

### HasAuthUrlRelativePath

`func (o *PrivateAuth) HasAuthUrlRelativePath() bool`

HasAuthUrlRelativePath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


