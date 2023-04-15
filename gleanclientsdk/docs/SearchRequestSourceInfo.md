# SearchRequestSourceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Initiator** | Pointer to **string** | The means by which the search request was initiated. EVAL - Internal usage for automated evaluation FACETS - A background request to get bucket counts for facets following a user search request MORE - The infinite scroll requested more results for an existing search ONBOARDING - A demo query was issued as part of the onboarding flow ONBOARDING_CHECKLIST - The user performed a search from the search step of the new user onboarding checklist PAGE_LOAD - A SERP was visited without going through regular product UI, e.g. from a bookmark, page refresh, or hitting the browser back button DISCARDED_PAGE_LOAD - as PAGE_LOAD but the page was previously discarded by the browser PREFETCH - Results for a non-active tab were requested, e.g. gmail USER - The user initiated a search by manually typing a query, clicking a suggestion, etc. RECOMMENDATION - A query intent is detected from the user&#39;s activity and a search request is issued proactively AUTOMATION - Initiated from an API used for automation by an external developer or integration. | [optional] 
**Modality** | **string** | The UI paradigm from which the search request was sent. FULLPAGE - The standard web app (including mobile) OVERLAY - An iframe that&#39;s not Embedded Search / NSR (No such frame type as of now) OMNIBOX - The browser omnibox CONTEXT_MENU - The browser right-click context menu (powered by the browser extension) EMBEDDED_SEARCH - The embedded search added as an iframe NSR - Native search replacement provided by extension injected iframe SIDEBAR - The extension sidebar GLEANBOT - Gleanbot in Slack, MS Teams, Discord, etc. | 
**Domain** | Pointer to **string** | The domain from/on behalf of which the request is being issued. Currently only being used for tracking / logging purposes. | [optional] 
**Platform** | Pointer to **string** | Platform from which the search request was sent. Optional field. | [optional] 
**UiElement** | Pointer to **string** | The (optional) UI element within the paradigm from which the search request was sent. Each modality will have a dedicated uiElement enum (e.g., SearchRequestGleanbotUIElementEnum) | [optional] 
**IsDebug** | Pointer to **bool** | Whether the query is for debugging purposes and, as such, should not be included in usage metrics and quality pipelines. | [optional] 
**ClientVersion** | Pointer to **string** | An opaque version identifier for the client. This is meant to be used for logging and debugging purposes only. | [optional] 

## Methods

### NewSearchRequestSourceInfo

`func NewSearchRequestSourceInfo(modality string, ) *SearchRequestSourceInfo`

NewSearchRequestSourceInfo instantiates a new SearchRequestSourceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchRequestSourceInfoWithDefaults

`func NewSearchRequestSourceInfoWithDefaults() *SearchRequestSourceInfo`

NewSearchRequestSourceInfoWithDefaults instantiates a new SearchRequestSourceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInitiator

`func (o *SearchRequestSourceInfo) GetInitiator() string`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *SearchRequestSourceInfo) GetInitiatorOk() (*string, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *SearchRequestSourceInfo) SetInitiator(v string)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *SearchRequestSourceInfo) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.

### GetModality

`func (o *SearchRequestSourceInfo) GetModality() string`

GetModality returns the Modality field if non-nil, zero value otherwise.

### GetModalityOk

`func (o *SearchRequestSourceInfo) GetModalityOk() (*string, bool)`

GetModalityOk returns a tuple with the Modality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModality

`func (o *SearchRequestSourceInfo) SetModality(v string)`

SetModality sets Modality field to given value.


### GetDomain

`func (o *SearchRequestSourceInfo) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *SearchRequestSourceInfo) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *SearchRequestSourceInfo) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *SearchRequestSourceInfo) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetPlatform

`func (o *SearchRequestSourceInfo) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *SearchRequestSourceInfo) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *SearchRequestSourceInfo) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *SearchRequestSourceInfo) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetUiElement

`func (o *SearchRequestSourceInfo) GetUiElement() string`

GetUiElement returns the UiElement field if non-nil, zero value otherwise.

### GetUiElementOk

`func (o *SearchRequestSourceInfo) GetUiElementOk() (*string, bool)`

GetUiElementOk returns a tuple with the UiElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiElement

`func (o *SearchRequestSourceInfo) SetUiElement(v string)`

SetUiElement sets UiElement field to given value.

### HasUiElement

`func (o *SearchRequestSourceInfo) HasUiElement() bool`

HasUiElement returns a boolean if a field has been set.

### GetIsDebug

`func (o *SearchRequestSourceInfo) GetIsDebug() bool`

GetIsDebug returns the IsDebug field if non-nil, zero value otherwise.

### GetIsDebugOk

`func (o *SearchRequestSourceInfo) GetIsDebugOk() (*bool, bool)`

GetIsDebugOk returns a tuple with the IsDebug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDebug

`func (o *SearchRequestSourceInfo) SetIsDebug(v bool)`

SetIsDebug sets IsDebug field to given value.

### HasIsDebug

`func (o *SearchRequestSourceInfo) HasIsDebug() bool`

HasIsDebug returns a boolean if a field has been set.

### GetClientVersion

`func (o *SearchRequestSourceInfo) GetClientVersion() string`

GetClientVersion returns the ClientVersion field if non-nil, zero value otherwise.

### GetClientVersionOk

`func (o *SearchRequestSourceInfo) GetClientVersionOk() (*string, bool)`

GetClientVersionOk returns a tuple with the ClientVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientVersion

`func (o *SearchRequestSourceInfo) SetClientVersion(v string)`

SetClientVersion sets ClientVersion field to given value.

### HasClientVersion

`func (o *SearchRequestSourceInfo) HasClientVersion() bool`

HasClientVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


