/*
Glean Client API

# Introduction These are the public APIs to enable implementing a custom client interface to the Glean system.  # Usage guidelines This API is evolving fast. Glean will provide advance notice of any planned backwards incompatible changes along with a 6-month sunset period for anything that requires developers to adopt the new versions. 

API version: 0.9.0
Contact: support@glean.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gleanclientsdk

import (
	"encoding/json"
	"time"
)

// checks if the Feedback type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Feedback{}

// Feedback struct for Feedback
type Feedback struct {
	// The feature category to which the feedback applies. These should be broad product areas such as Announcements, Answers, Search, etc. rather than specific components or UI treatments within those areas.
	Category *string `json:"category,omitempty"`
	// A list of server-generated trackingTokens to which this event applies.
	TrackingTokens []string `json:"trackingTokens"`
	// The action the user took within a Glean client with respect to the object referred to by the given `trackingToken``. CLICK - The object's primary link was clicked with the intent to view its full representation. Depending on the object type, this may imply an external navigation or navigating to a new page or view within the Glean app. CONTAINER_CLICK - A link to the object's parent container (e.g. the folder in which it's located) was clicked. COPY_LINK - The user copied a link to the primary link. CREATE - The user creates a document. DISMISS - The user dismissed the object such that it was hidden from view. DOWNVOTE - The user gave feedback that the object was not useful. EMAIL - The user attempted to send an email. FOCUS_IN - The user clicked into an interactive element, e.g. the search box. MANUAL_FEEDBACK - The user submitted textual manual feedback regarding the object. MESSAGE - The user attempted to send a message using their default messaing app. MIDDLE_CLICK - The user middle clicked the object's primary link with the intent to open its full representation in a new tab. PREVIEW - The user clicked the object's inline preview affordance. RIGHT_CLICK - The user right clicked the object's primary link. This may indicate an intent to open it in a new tab or copy it. SECTION_CLICK - The user clicked a link to a subsection of the primary object. SEEN - The user has likely seen the object (e.g. took action to make the object visible within the user's viewport). SHARE - The user shared the object with another user. SHOW_MORE - The user clicked the object's show more affordance. UPVOTE - The user gave feedback that the object was useful. VIEW - The object was visible within the user's viewport. VISIBLE - The object was visible within the user's viewport.
	Event string `json:"event"`
	// Position of the element in the case that the client controls order (such as feed and autocomplete).
	Position *int32 `json:"position,omitempty"`
	// For type MANUAL_FEEDBACK, contains string of user feedback. For autocomplete, partial query string. For feed, string of user feedback in addition to manual feedback signals extracted from all suggested content.
	Payload *string `json:"payload,omitempty"`
	SessionInfo *SessionInfo `json:"sessionInfo,omitempty"`
	// The ISO 8601 timestamp when the event occured.
	Timestamp *time.Time `json:"timestamp,omitempty"`
	User *User `json:"user,omitempty"`
	// The path the client was at when the feedback event triggered.
	Pathname *string `json:"pathname,omitempty"`
	// Where the feedback will be sent, e.g. to Glean, the user's company, or both. If no channels are specified, feedback will go only to Glean.
	Channels []string `json:"channels,omitempty"`
	// The url the client was at when the feedback event triggered.
	Url *string `json:"url,omitempty"`
	// The UI element associated with the event, if any.
	UiElement *string `json:"uiElement,omitempty"`
	ManualFeedbackInfo *ManualFeedbackInfo `json:"manualFeedbackInfo,omitempty"`
	SeenFeedbackInfo *SeenFeedbackInfo `json:"seenFeedbackInfo,omitempty"`
	UserViewInfo *UserViewInfo `json:"userViewInfo,omitempty"`
}

// NewFeedback instantiates a new Feedback object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeedback(trackingTokens []string, event string) *Feedback {
	this := Feedback{}
	this.TrackingTokens = trackingTokens
	this.Event = event
	return &this
}

// NewFeedbackWithDefaults instantiates a new Feedback object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeedbackWithDefaults() *Feedback {
	this := Feedback{}
	return &this
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *Feedback) GetCategory() string {
	if o == nil || IsNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Feedback) GetCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *Feedback) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *Feedback) SetCategory(v string) {
	o.Category = &v
}

// GetTrackingTokens returns the TrackingTokens field value
func (o *Feedback) GetTrackingTokens() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.TrackingTokens
}

// GetTrackingTokensOk returns a tuple with the TrackingTokens field value
// and a boolean to check if the value has been set.
func (o *Feedback) GetTrackingTokensOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TrackingTokens, true
}

// SetTrackingTokens sets field value
func (o *Feedback) SetTrackingTokens(v []string) {
	o.TrackingTokens = v
}

// GetEvent returns the Event field value
func (o *Feedback) GetEvent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *Feedback) GetEventOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value
func (o *Feedback) SetEvent(v string) {
	o.Event = v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *Feedback) GetPosition() int32 {
	if o == nil || IsNil(o.Position) {
		var ret int32
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Feedback) GetPositionOk() (*int32, bool) {
	if o == nil || IsNil(o.Position) {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *Feedback) HasPosition() bool {
	if o != nil && !IsNil(o.Position) {
		return true
	}

	return false
}

// SetPosition gets a reference to the given int32 and assigns it to the Position field.
func (o *Feedback) SetPosition(v int32) {
	o.Position = &v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *Feedback) GetPayload() string {
	if o == nil || IsNil(o.Payload) {
		var ret string
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Feedback) GetPayloadOk() (*string, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *Feedback) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given string and assigns it to the Payload field.
func (o *Feedback) SetPayload(v string) {
	o.Payload = &v
}

// GetSessionInfo returns the SessionInfo field value if set, zero value otherwise.
func (o *Feedback) GetSessionInfo() SessionInfo {
	if o == nil || IsNil(o.SessionInfo) {
		var ret SessionInfo
		return ret
	}
	return *o.SessionInfo
}

// GetSessionInfoOk returns a tuple with the SessionInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Feedback) GetSessionInfoOk() (*SessionInfo, bool) {
	if o == nil || IsNil(o.SessionInfo) {
		return nil, false
	}
	return o.SessionInfo, true
}

// HasSessionInfo returns a boolean if a field has been set.
func (o *Feedback) HasSessionInfo() bool {
	if o != nil && !IsNil(o.SessionInfo) {
		return true
	}

	return false
}

// SetSessionInfo gets a reference to the given SessionInfo and assigns it to the SessionInfo field.
func (o *Feedback) SetSessionInfo(v SessionInfo) {
	o.SessionInfo = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *Feedback) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Feedback) GetTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *Feedback) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *Feedback) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *Feedback) GetUser() User {
	if o == nil || IsNil(o.User) {
		var ret User
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Feedback) GetUserOk() (*User, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *Feedback) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given User and assigns it to the User field.
func (o *Feedback) SetUser(v User) {
	o.User = &v
}

// GetPathname returns the Pathname field value if set, zero value otherwise.
func (o *Feedback) GetPathname() string {
	if o == nil || IsNil(o.Pathname) {
		var ret string
		return ret
	}
	return *o.Pathname
}

// GetPathnameOk returns a tuple with the Pathname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Feedback) GetPathnameOk() (*string, bool) {
	if o == nil || IsNil(o.Pathname) {
		return nil, false
	}
	return o.Pathname, true
}

// HasPathname returns a boolean if a field has been set.
func (o *Feedback) HasPathname() bool {
	if o != nil && !IsNil(o.Pathname) {
		return true
	}

	return false
}

// SetPathname gets a reference to the given string and assigns it to the Pathname field.
func (o *Feedback) SetPathname(v string) {
	o.Pathname = &v
}

// GetChannels returns the Channels field value if set, zero value otherwise.
func (o *Feedback) GetChannels() []string {
	if o == nil || IsNil(o.Channels) {
		var ret []string
		return ret
	}
	return o.Channels
}

// GetChannelsOk returns a tuple with the Channels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Feedback) GetChannelsOk() ([]string, bool) {
	if o == nil || IsNil(o.Channels) {
		return nil, false
	}
	return o.Channels, true
}

// HasChannels returns a boolean if a field has been set.
func (o *Feedback) HasChannels() bool {
	if o != nil && !IsNil(o.Channels) {
		return true
	}

	return false
}

// SetChannels gets a reference to the given []string and assigns it to the Channels field.
func (o *Feedback) SetChannels(v []string) {
	o.Channels = v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *Feedback) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Feedback) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *Feedback) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *Feedback) SetUrl(v string) {
	o.Url = &v
}

// GetUiElement returns the UiElement field value if set, zero value otherwise.
func (o *Feedback) GetUiElement() string {
	if o == nil || IsNil(o.UiElement) {
		var ret string
		return ret
	}
	return *o.UiElement
}

// GetUiElementOk returns a tuple with the UiElement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Feedback) GetUiElementOk() (*string, bool) {
	if o == nil || IsNil(o.UiElement) {
		return nil, false
	}
	return o.UiElement, true
}

// HasUiElement returns a boolean if a field has been set.
func (o *Feedback) HasUiElement() bool {
	if o != nil && !IsNil(o.UiElement) {
		return true
	}

	return false
}

// SetUiElement gets a reference to the given string and assigns it to the UiElement field.
func (o *Feedback) SetUiElement(v string) {
	o.UiElement = &v
}

// GetManualFeedbackInfo returns the ManualFeedbackInfo field value if set, zero value otherwise.
func (o *Feedback) GetManualFeedbackInfo() ManualFeedbackInfo {
	if o == nil || IsNil(o.ManualFeedbackInfo) {
		var ret ManualFeedbackInfo
		return ret
	}
	return *o.ManualFeedbackInfo
}

// GetManualFeedbackInfoOk returns a tuple with the ManualFeedbackInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Feedback) GetManualFeedbackInfoOk() (*ManualFeedbackInfo, bool) {
	if o == nil || IsNil(o.ManualFeedbackInfo) {
		return nil, false
	}
	return o.ManualFeedbackInfo, true
}

// HasManualFeedbackInfo returns a boolean if a field has been set.
func (o *Feedback) HasManualFeedbackInfo() bool {
	if o != nil && !IsNil(o.ManualFeedbackInfo) {
		return true
	}

	return false
}

// SetManualFeedbackInfo gets a reference to the given ManualFeedbackInfo and assigns it to the ManualFeedbackInfo field.
func (o *Feedback) SetManualFeedbackInfo(v ManualFeedbackInfo) {
	o.ManualFeedbackInfo = &v
}

// GetSeenFeedbackInfo returns the SeenFeedbackInfo field value if set, zero value otherwise.
func (o *Feedback) GetSeenFeedbackInfo() SeenFeedbackInfo {
	if o == nil || IsNil(o.SeenFeedbackInfo) {
		var ret SeenFeedbackInfo
		return ret
	}
	return *o.SeenFeedbackInfo
}

// GetSeenFeedbackInfoOk returns a tuple with the SeenFeedbackInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Feedback) GetSeenFeedbackInfoOk() (*SeenFeedbackInfo, bool) {
	if o == nil || IsNil(o.SeenFeedbackInfo) {
		return nil, false
	}
	return o.SeenFeedbackInfo, true
}

// HasSeenFeedbackInfo returns a boolean if a field has been set.
func (o *Feedback) HasSeenFeedbackInfo() bool {
	if o != nil && !IsNil(o.SeenFeedbackInfo) {
		return true
	}

	return false
}

// SetSeenFeedbackInfo gets a reference to the given SeenFeedbackInfo and assigns it to the SeenFeedbackInfo field.
func (o *Feedback) SetSeenFeedbackInfo(v SeenFeedbackInfo) {
	o.SeenFeedbackInfo = &v
}

// GetUserViewInfo returns the UserViewInfo field value if set, zero value otherwise.
func (o *Feedback) GetUserViewInfo() UserViewInfo {
	if o == nil || IsNil(o.UserViewInfo) {
		var ret UserViewInfo
		return ret
	}
	return *o.UserViewInfo
}

// GetUserViewInfoOk returns a tuple with the UserViewInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Feedback) GetUserViewInfoOk() (*UserViewInfo, bool) {
	if o == nil || IsNil(o.UserViewInfo) {
		return nil, false
	}
	return o.UserViewInfo, true
}

// HasUserViewInfo returns a boolean if a field has been set.
func (o *Feedback) HasUserViewInfo() bool {
	if o != nil && !IsNil(o.UserViewInfo) {
		return true
	}

	return false
}

// SetUserViewInfo gets a reference to the given UserViewInfo and assigns it to the UserViewInfo field.
func (o *Feedback) SetUserViewInfo(v UserViewInfo) {
	o.UserViewInfo = &v
}

func (o Feedback) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Feedback) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	toSerialize["trackingTokens"] = o.TrackingTokens
	toSerialize["event"] = o.Event
	if !IsNil(o.Position) {
		toSerialize["position"] = o.Position
	}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !IsNil(o.SessionInfo) {
		toSerialize["sessionInfo"] = o.SessionInfo
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	if !IsNil(o.Pathname) {
		toSerialize["pathname"] = o.Pathname
	}
	if !IsNil(o.Channels) {
		toSerialize["channels"] = o.Channels
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.UiElement) {
		toSerialize["uiElement"] = o.UiElement
	}
	if !IsNil(o.ManualFeedbackInfo) {
		toSerialize["manualFeedbackInfo"] = o.ManualFeedbackInfo
	}
	if !IsNil(o.SeenFeedbackInfo) {
		toSerialize["seenFeedbackInfo"] = o.SeenFeedbackInfo
	}
	if !IsNil(o.UserViewInfo) {
		toSerialize["userViewInfo"] = o.UserViewInfo
	}
	return toSerialize, nil
}

type NullableFeedback struct {
	value *Feedback
	isSet bool
}

func (v NullableFeedback) Get() *Feedback {
	return v.value
}

func (v *NullableFeedback) Set(val *Feedback) {
	v.value = val
	v.isSet = true
}

func (v NullableFeedback) IsSet() bool {
	return v.isSet
}

func (v *NullableFeedback) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeedback(val *Feedback) *NullableFeedback {
	return &NullableFeedback{value: val, isSet: true}
}

func (v NullableFeedback) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeedback) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


