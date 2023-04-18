# Feedback

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** | The feature category to which the feedback applies. These should be broad product areas such as Announcements, Answers, Search, etc. rather than specific components or UI treatments within those areas. | [optional] 
**TrackingTokens** | **[]string** | A list of server-generated trackingTokens to which this event applies. | 
**Event** | **string** | The action the user took within a Glean client with respect to the object referred to by the given &#x60;trackingToken&#x60;&#x60;. CLICK - The object&#39;s primary link was clicked with the intent to view its full representation. Depending on the object type, this may imply an external navigation or navigating to a new page or view within the Glean app. CONTAINER_CLICK - A link to the object&#39;s parent container (e.g. the folder in which it&#39;s located) was clicked. COPY_LINK - The user copied a link to the primary link. CREATE - The user creates a document. DISMISS - The user dismissed the object such that it was hidden from view. DOWNVOTE - The user gave feedback that the object was not useful. EMAIL - The user attempted to send an email. FOCUS_IN - The user clicked into an interactive element, e.g. the search box. MANUAL_FEEDBACK - The user submitted textual manual feedback regarding the object. MESSAGE - The user attempted to send a message using their default messaing app. MIDDLE_CLICK - The user middle clicked the object&#39;s primary link with the intent to open its full representation in a new tab. PREVIEW - The user clicked the object&#39;s inline preview affordance. RIGHT_CLICK - The user right clicked the object&#39;s primary link. This may indicate an intent to open it in a new tab or copy it. SECTION_CLICK - The user clicked a link to a subsection of the primary object. SEEN - The user has likely seen the object (e.g. took action to make the object visible within the user&#39;s viewport). SHARE - The user shared the object with another user. SHOW_MORE - The user clicked the object&#39;s show more affordance. UPVOTE - The user gave feedback that the object was useful. VIEW - The object was visible within the user&#39;s viewport. VISIBLE - The object was visible within the user&#39;s viewport. | 
**Position** | Pointer to **int32** | Position of the element in the case that the client controls order (such as feed and autocomplete). | [optional] 
**Payload** | Pointer to **string** | For type MANUAL_FEEDBACK, contains string of user feedback. For autocomplete, partial query string. For feed, string of user feedback in addition to manual feedback signals extracted from all suggested content. | [optional] 
**SessionInfo** | Pointer to [**SessionInfo**](SessionInfo.md) |  | [optional] 
**Timestamp** | Pointer to **time.Time** | The ISO 8601 timestamp when the event occured. | [optional] 
**User** | Pointer to [**User**](User.md) |  | [optional] 
**Pathname** | Pointer to **string** | The path the client was at when the feedback event triggered. | [optional] 
**Channels** | Pointer to **[]string** | Where the feedback will be sent, e.g. to Glean, the user&#39;s company, or both. If no channels are specified, feedback will go only to Glean. | [optional] 
**Url** | Pointer to **string** | The url the client was at when the feedback event triggered. | [optional] 
**UiElement** | Pointer to **string** | The UI element associated with the event, if any. | [optional] 
**ManualFeedbackInfo** | Pointer to [**ManualFeedbackInfo**](ManualFeedbackInfo.md) |  | [optional] 
**SeenFeedbackInfo** | Pointer to [**SeenFeedbackInfo**](SeenFeedbackInfo.md) |  | [optional] 
**UserViewInfo** | Pointer to [**UserViewInfo**](UserViewInfo.md) |  | [optional] 

## Methods

### NewFeedback

`func NewFeedback(trackingTokens []string, event string, ) *Feedback`

NewFeedback instantiates a new Feedback object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeedbackWithDefaults

`func NewFeedbackWithDefaults() *Feedback`

NewFeedbackWithDefaults instantiates a new Feedback object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *Feedback) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *Feedback) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *Feedback) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *Feedback) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetTrackingTokens

`func (o *Feedback) GetTrackingTokens() []string`

GetTrackingTokens returns the TrackingTokens field if non-nil, zero value otherwise.

### GetTrackingTokensOk

`func (o *Feedback) GetTrackingTokensOk() (*[]string, bool)`

GetTrackingTokensOk returns a tuple with the TrackingTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingTokens

`func (o *Feedback) SetTrackingTokens(v []string)`

SetTrackingTokens sets TrackingTokens field to given value.


### GetEvent

`func (o *Feedback) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *Feedback) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *Feedback) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetPosition

`func (o *Feedback) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *Feedback) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *Feedback) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *Feedback) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetPayload

`func (o *Feedback) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *Feedback) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *Feedback) SetPayload(v string)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *Feedback) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetSessionInfo

`func (o *Feedback) GetSessionInfo() SessionInfo`

GetSessionInfo returns the SessionInfo field if non-nil, zero value otherwise.

### GetSessionInfoOk

`func (o *Feedback) GetSessionInfoOk() (*SessionInfo, bool)`

GetSessionInfoOk returns a tuple with the SessionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionInfo

`func (o *Feedback) SetSessionInfo(v SessionInfo)`

SetSessionInfo sets SessionInfo field to given value.

### HasSessionInfo

`func (o *Feedback) HasSessionInfo() bool`

HasSessionInfo returns a boolean if a field has been set.

### GetTimestamp

`func (o *Feedback) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Feedback) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Feedback) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *Feedback) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetUser

`func (o *Feedback) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Feedback) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Feedback) SetUser(v User)`

SetUser sets User field to given value.

### HasUser

`func (o *Feedback) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetPathname

`func (o *Feedback) GetPathname() string`

GetPathname returns the Pathname field if non-nil, zero value otherwise.

### GetPathnameOk

`func (o *Feedback) GetPathnameOk() (*string, bool)`

GetPathnameOk returns a tuple with the Pathname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathname

`func (o *Feedback) SetPathname(v string)`

SetPathname sets Pathname field to given value.

### HasPathname

`func (o *Feedback) HasPathname() bool`

HasPathname returns a boolean if a field has been set.

### GetChannels

`func (o *Feedback) GetChannels() []string`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *Feedback) GetChannelsOk() (*[]string, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *Feedback) SetChannels(v []string)`

SetChannels sets Channels field to given value.

### HasChannels

`func (o *Feedback) HasChannels() bool`

HasChannels returns a boolean if a field has been set.

### GetUrl

`func (o *Feedback) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Feedback) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Feedback) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Feedback) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUiElement

`func (o *Feedback) GetUiElement() string`

GetUiElement returns the UiElement field if non-nil, zero value otherwise.

### GetUiElementOk

`func (o *Feedback) GetUiElementOk() (*string, bool)`

GetUiElementOk returns a tuple with the UiElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiElement

`func (o *Feedback) SetUiElement(v string)`

SetUiElement sets UiElement field to given value.

### HasUiElement

`func (o *Feedback) HasUiElement() bool`

HasUiElement returns a boolean if a field has been set.

### GetManualFeedbackInfo

`func (o *Feedback) GetManualFeedbackInfo() ManualFeedbackInfo`

GetManualFeedbackInfo returns the ManualFeedbackInfo field if non-nil, zero value otherwise.

### GetManualFeedbackInfoOk

`func (o *Feedback) GetManualFeedbackInfoOk() (*ManualFeedbackInfo, bool)`

GetManualFeedbackInfoOk returns a tuple with the ManualFeedbackInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualFeedbackInfo

`func (o *Feedback) SetManualFeedbackInfo(v ManualFeedbackInfo)`

SetManualFeedbackInfo sets ManualFeedbackInfo field to given value.

### HasManualFeedbackInfo

`func (o *Feedback) HasManualFeedbackInfo() bool`

HasManualFeedbackInfo returns a boolean if a field has been set.

### GetSeenFeedbackInfo

`func (o *Feedback) GetSeenFeedbackInfo() SeenFeedbackInfo`

GetSeenFeedbackInfo returns the SeenFeedbackInfo field if non-nil, zero value otherwise.

### GetSeenFeedbackInfoOk

`func (o *Feedback) GetSeenFeedbackInfoOk() (*SeenFeedbackInfo, bool)`

GetSeenFeedbackInfoOk returns a tuple with the SeenFeedbackInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeenFeedbackInfo

`func (o *Feedback) SetSeenFeedbackInfo(v SeenFeedbackInfo)`

SetSeenFeedbackInfo sets SeenFeedbackInfo field to given value.

### HasSeenFeedbackInfo

`func (o *Feedback) HasSeenFeedbackInfo() bool`

HasSeenFeedbackInfo returns a boolean if a field has been set.

### GetUserViewInfo

`func (o *Feedback) GetUserViewInfo() UserViewInfo`

GetUserViewInfo returns the UserViewInfo field if non-nil, zero value otherwise.

### GetUserViewInfoOk

`func (o *Feedback) GetUserViewInfoOk() (*UserViewInfo, bool)`

GetUserViewInfoOk returns a tuple with the UserViewInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserViewInfo

`func (o *Feedback) SetUserViewInfo(v UserViewInfo)`

SetUserViewInfo sets UserViewInfo field to given value.

### HasUserViewInfo

`func (o *Feedback) HasUserViewInfo() bool`

HasUserViewInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


