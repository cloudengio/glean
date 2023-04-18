# ManualFeedbackInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | The email address of the user who submitted the Feedback.event.MANUAL_FEEDBACK event. | [optional] 
**Source** | Pointer to **string** | The source associated with the Feedback.event.MANUAL_FEEDBACK event. | [optional] 
**Issue** | Pointer to **string** | The issue the user indicated in the feedback. | [optional] 
**Query** | Pointer to **string** | The query associated with the Feedback.event.MANUAL_FEEDBACK event. | [optional] 
**ObscuredQuery** | Pointer to **string** | The query associated with the Feedback.event.MANUAL_FEEDBACK event, but obscured such that the vowels are replaced with special characters. For search feedback events only. | [optional] 
**ActiveTab** | Pointer to **string** | Which tabs the user had chosen at the time of the Feedback.event.MANUAL_FEEDBACK event. For search feedback events only. | [optional] 
**Comments** | Pointer to **string** | The comments users can optionally add to the Feedback.event.MANUAL_FEEDBACK events. | [optional] 
**SearchResults** | Pointer to **[]string** | The array of search result docIds, ordered by top to bottom result. | [optional] 
**NumQueriesFromFirstRun** | Pointer to **int32** | How many times this query has been run in the past. | [optional] 

## Methods

### NewManualFeedbackInfo

`func NewManualFeedbackInfo() *ManualFeedbackInfo`

NewManualFeedbackInfo instantiates a new ManualFeedbackInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManualFeedbackInfoWithDefaults

`func NewManualFeedbackInfoWithDefaults() *ManualFeedbackInfo`

NewManualFeedbackInfoWithDefaults instantiates a new ManualFeedbackInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *ManualFeedbackInfo) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ManualFeedbackInfo) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ManualFeedbackInfo) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ManualFeedbackInfo) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetSource

`func (o *ManualFeedbackInfo) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ManualFeedbackInfo) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ManualFeedbackInfo) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ManualFeedbackInfo) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetIssue

`func (o *ManualFeedbackInfo) GetIssue() string`

GetIssue returns the Issue field if non-nil, zero value otherwise.

### GetIssueOk

`func (o *ManualFeedbackInfo) GetIssueOk() (*string, bool)`

GetIssueOk returns a tuple with the Issue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssue

`func (o *ManualFeedbackInfo) SetIssue(v string)`

SetIssue sets Issue field to given value.

### HasIssue

`func (o *ManualFeedbackInfo) HasIssue() bool`

HasIssue returns a boolean if a field has been set.

### GetQuery

`func (o *ManualFeedbackInfo) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *ManualFeedbackInfo) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *ManualFeedbackInfo) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *ManualFeedbackInfo) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetObscuredQuery

`func (o *ManualFeedbackInfo) GetObscuredQuery() string`

GetObscuredQuery returns the ObscuredQuery field if non-nil, zero value otherwise.

### GetObscuredQueryOk

`func (o *ManualFeedbackInfo) GetObscuredQueryOk() (*string, bool)`

GetObscuredQueryOk returns a tuple with the ObscuredQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObscuredQuery

`func (o *ManualFeedbackInfo) SetObscuredQuery(v string)`

SetObscuredQuery sets ObscuredQuery field to given value.

### HasObscuredQuery

`func (o *ManualFeedbackInfo) HasObscuredQuery() bool`

HasObscuredQuery returns a boolean if a field has been set.

### GetActiveTab

`func (o *ManualFeedbackInfo) GetActiveTab() string`

GetActiveTab returns the ActiveTab field if non-nil, zero value otherwise.

### GetActiveTabOk

`func (o *ManualFeedbackInfo) GetActiveTabOk() (*string, bool)`

GetActiveTabOk returns a tuple with the ActiveTab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveTab

`func (o *ManualFeedbackInfo) SetActiveTab(v string)`

SetActiveTab sets ActiveTab field to given value.

### HasActiveTab

`func (o *ManualFeedbackInfo) HasActiveTab() bool`

HasActiveTab returns a boolean if a field has been set.

### GetComments

`func (o *ManualFeedbackInfo) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *ManualFeedbackInfo) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *ManualFeedbackInfo) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *ManualFeedbackInfo) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetSearchResults

`func (o *ManualFeedbackInfo) GetSearchResults() []string`

GetSearchResults returns the SearchResults field if non-nil, zero value otherwise.

### GetSearchResultsOk

`func (o *ManualFeedbackInfo) GetSearchResultsOk() (*[]string, bool)`

GetSearchResultsOk returns a tuple with the SearchResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchResults

`func (o *ManualFeedbackInfo) SetSearchResults(v []string)`

SetSearchResults sets SearchResults field to given value.

### HasSearchResults

`func (o *ManualFeedbackInfo) HasSearchResults() bool`

HasSearchResults returns a boolean if a field has been set.

### GetNumQueriesFromFirstRun

`func (o *ManualFeedbackInfo) GetNumQueriesFromFirstRun() int32`

GetNumQueriesFromFirstRun returns the NumQueriesFromFirstRun field if non-nil, zero value otherwise.

### GetNumQueriesFromFirstRunOk

`func (o *ManualFeedbackInfo) GetNumQueriesFromFirstRunOk() (*int32, bool)`

GetNumQueriesFromFirstRunOk returns a tuple with the NumQueriesFromFirstRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumQueriesFromFirstRun

`func (o *ManualFeedbackInfo) SetNumQueriesFromFirstRun(v int32)`

SetNumQueriesFromFirstRun sets NumQueriesFromFirstRun field to given value.

### HasNumQueriesFromFirstRun

`func (o *ManualFeedbackInfo) HasNumQueriesFromFirstRun() bool`

HasNumQueriesFromFirstRun returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


