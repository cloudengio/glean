/*
Glean Client API - Platform Preview

# Introduction These are all the APIs used by Glean to implement the Glean client. These are available as platform preview for implementing a custom client to the Glean system.  # Usage guidelines A subset of these endpoints are also in the developer ready section, which is available for public use. The rest of the endpoints are subject to prior agreement with Glean before usage. Please contact support@glean.com if you would like to use an API that is not currently available in the developer ready section. 

API version: 0.9.0
Contact: support@glean.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gleanclientsdk

import (
	"encoding/json"
	"time"
)

// checks if the UpdateAnnouncementRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateAnnouncementRequest{}

// UpdateAnnouncementRequest struct for UpdateAnnouncementRequest
type UpdateAnnouncementRequest struct {
	// The date and time at which the announcement becomes active.
	StartTime time.Time `json:"startTime"`
	// The date and time at which the announcement expires.
	EndTime time.Time `json:"endTime"`
	// The headline of the announcement.
	Title string `json:"title"`
	Body *StructuredText `json:"body,omitempty"`
	// An emoji used to indicate the nature of the announcement.
	Emoji *string `json:"emoji,omitempty"`
	Thumbnail *Thumbnail `json:"thumbnail,omitempty"`
	Banner *Thumbnail `json:"banner,omitempty"`
	// Filters which restrict who should see the announcement. Values are taken from the corresponding filters in people search.
	AudienceFilters []FacetFilter `json:"audienceFilters,omitempty"`
	// The Document ID of the Source Document this Announcement was created from (e.g. Slack thread).
	SourceDocumentId *string `json:"sourceDocumentId,omitempty"`
	// Whether or not to hide an author attribution.
	HideAttribution *bool `json:"hideAttribution,omitempty"`
	// This determines whether this is a Social Feed post or a regular announcement.
	Channel *string `json:"channel,omitempty"`
	// Used by the Social Feed to pin posts to the front of the feed.
	IsPrioritized *bool `json:"isPrioritized,omitempty"`
	// Url for viewing the announcement. It will be set to document url for announcements from other datasources e.g. simpplr. Can only be written when channel=\"SOCIAL_FEED\".
	ViewUrl *string `json:"viewUrl,omitempty"`
	// The opaque id of the announcement.
	Id int32 `json:"id"`
	Data *AnnouncementCreateOrUpdateData `json:"data,omitempty"`
}

// NewUpdateAnnouncementRequest instantiates a new UpdateAnnouncementRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateAnnouncementRequest(startTime time.Time, endTime time.Time, title string, id int32) *UpdateAnnouncementRequest {
	this := UpdateAnnouncementRequest{}
	this.StartTime = startTime
	this.EndTime = endTime
	this.Title = title
	this.Id = id
	return &this
}

// NewUpdateAnnouncementRequestWithDefaults instantiates a new UpdateAnnouncementRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateAnnouncementRequestWithDefaults() *UpdateAnnouncementRequest {
	this := UpdateAnnouncementRequest{}
	return &this
}

// GetStartTime returns the StartTime field value
func (o *UpdateAnnouncementRequest) GetStartTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *UpdateAnnouncementRequest) GetStartTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value
func (o *UpdateAnnouncementRequest) SetStartTime(v time.Time) {
	o.StartTime = v
}

// GetEndTime returns the EndTime field value
func (o *UpdateAnnouncementRequest) GetEndTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value
// and a boolean to check if the value has been set.
func (o *UpdateAnnouncementRequest) GetEndTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndTime, true
}

// SetEndTime sets field value
func (o *UpdateAnnouncementRequest) SetEndTime(v time.Time) {
	o.EndTime = v
}

// GetTitle returns the Title field value
func (o *UpdateAnnouncementRequest) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *UpdateAnnouncementRequest) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *UpdateAnnouncementRequest) SetTitle(v string) {
	o.Title = v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *UpdateAnnouncementRequest) GetBody() StructuredText {
	if o == nil || IsNil(o.Body) {
		var ret StructuredText
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAnnouncementRequest) GetBodyOk() (*StructuredText, bool) {
	if o == nil || IsNil(o.Body) {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *UpdateAnnouncementRequest) HasBody() bool {
	if o != nil && !IsNil(o.Body) {
		return true
	}

	return false
}

// SetBody gets a reference to the given StructuredText and assigns it to the Body field.
func (o *UpdateAnnouncementRequest) SetBody(v StructuredText) {
	o.Body = &v
}

// GetEmoji returns the Emoji field value if set, zero value otherwise.
func (o *UpdateAnnouncementRequest) GetEmoji() string {
	if o == nil || IsNil(o.Emoji) {
		var ret string
		return ret
	}
	return *o.Emoji
}

// GetEmojiOk returns a tuple with the Emoji field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAnnouncementRequest) GetEmojiOk() (*string, bool) {
	if o == nil || IsNil(o.Emoji) {
		return nil, false
	}
	return o.Emoji, true
}

// HasEmoji returns a boolean if a field has been set.
func (o *UpdateAnnouncementRequest) HasEmoji() bool {
	if o != nil && !IsNil(o.Emoji) {
		return true
	}

	return false
}

// SetEmoji gets a reference to the given string and assigns it to the Emoji field.
func (o *UpdateAnnouncementRequest) SetEmoji(v string) {
	o.Emoji = &v
}

// GetThumbnail returns the Thumbnail field value if set, zero value otherwise.
func (o *UpdateAnnouncementRequest) GetThumbnail() Thumbnail {
	if o == nil || IsNil(o.Thumbnail) {
		var ret Thumbnail
		return ret
	}
	return *o.Thumbnail
}

// GetThumbnailOk returns a tuple with the Thumbnail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAnnouncementRequest) GetThumbnailOk() (*Thumbnail, bool) {
	if o == nil || IsNil(o.Thumbnail) {
		return nil, false
	}
	return o.Thumbnail, true
}

// HasThumbnail returns a boolean if a field has been set.
func (o *UpdateAnnouncementRequest) HasThumbnail() bool {
	if o != nil && !IsNil(o.Thumbnail) {
		return true
	}

	return false
}

// SetThumbnail gets a reference to the given Thumbnail and assigns it to the Thumbnail field.
func (o *UpdateAnnouncementRequest) SetThumbnail(v Thumbnail) {
	o.Thumbnail = &v
}

// GetBanner returns the Banner field value if set, zero value otherwise.
func (o *UpdateAnnouncementRequest) GetBanner() Thumbnail {
	if o == nil || IsNil(o.Banner) {
		var ret Thumbnail
		return ret
	}
	return *o.Banner
}

// GetBannerOk returns a tuple with the Banner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAnnouncementRequest) GetBannerOk() (*Thumbnail, bool) {
	if o == nil || IsNil(o.Banner) {
		return nil, false
	}
	return o.Banner, true
}

// HasBanner returns a boolean if a field has been set.
func (o *UpdateAnnouncementRequest) HasBanner() bool {
	if o != nil && !IsNil(o.Banner) {
		return true
	}

	return false
}

// SetBanner gets a reference to the given Thumbnail and assigns it to the Banner field.
func (o *UpdateAnnouncementRequest) SetBanner(v Thumbnail) {
	o.Banner = &v
}

// GetAudienceFilters returns the AudienceFilters field value if set, zero value otherwise.
func (o *UpdateAnnouncementRequest) GetAudienceFilters() []FacetFilter {
	if o == nil || IsNil(o.AudienceFilters) {
		var ret []FacetFilter
		return ret
	}
	return o.AudienceFilters
}

// GetAudienceFiltersOk returns a tuple with the AudienceFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAnnouncementRequest) GetAudienceFiltersOk() ([]FacetFilter, bool) {
	if o == nil || IsNil(o.AudienceFilters) {
		return nil, false
	}
	return o.AudienceFilters, true
}

// HasAudienceFilters returns a boolean if a field has been set.
func (o *UpdateAnnouncementRequest) HasAudienceFilters() bool {
	if o != nil && !IsNil(o.AudienceFilters) {
		return true
	}

	return false
}

// SetAudienceFilters gets a reference to the given []FacetFilter and assigns it to the AudienceFilters field.
func (o *UpdateAnnouncementRequest) SetAudienceFilters(v []FacetFilter) {
	o.AudienceFilters = v
}

// GetSourceDocumentId returns the SourceDocumentId field value if set, zero value otherwise.
func (o *UpdateAnnouncementRequest) GetSourceDocumentId() string {
	if o == nil || IsNil(o.SourceDocumentId) {
		var ret string
		return ret
	}
	return *o.SourceDocumentId
}

// GetSourceDocumentIdOk returns a tuple with the SourceDocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAnnouncementRequest) GetSourceDocumentIdOk() (*string, bool) {
	if o == nil || IsNil(o.SourceDocumentId) {
		return nil, false
	}
	return o.SourceDocumentId, true
}

// HasSourceDocumentId returns a boolean if a field has been set.
func (o *UpdateAnnouncementRequest) HasSourceDocumentId() bool {
	if o != nil && !IsNil(o.SourceDocumentId) {
		return true
	}

	return false
}

// SetSourceDocumentId gets a reference to the given string and assigns it to the SourceDocumentId field.
func (o *UpdateAnnouncementRequest) SetSourceDocumentId(v string) {
	o.SourceDocumentId = &v
}

// GetHideAttribution returns the HideAttribution field value if set, zero value otherwise.
func (o *UpdateAnnouncementRequest) GetHideAttribution() bool {
	if o == nil || IsNil(o.HideAttribution) {
		var ret bool
		return ret
	}
	return *o.HideAttribution
}

// GetHideAttributionOk returns a tuple with the HideAttribution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAnnouncementRequest) GetHideAttributionOk() (*bool, bool) {
	if o == nil || IsNil(o.HideAttribution) {
		return nil, false
	}
	return o.HideAttribution, true
}

// HasHideAttribution returns a boolean if a field has been set.
func (o *UpdateAnnouncementRequest) HasHideAttribution() bool {
	if o != nil && !IsNil(o.HideAttribution) {
		return true
	}

	return false
}

// SetHideAttribution gets a reference to the given bool and assigns it to the HideAttribution field.
func (o *UpdateAnnouncementRequest) SetHideAttribution(v bool) {
	o.HideAttribution = &v
}

// GetChannel returns the Channel field value if set, zero value otherwise.
func (o *UpdateAnnouncementRequest) GetChannel() string {
	if o == nil || IsNil(o.Channel) {
		var ret string
		return ret
	}
	return *o.Channel
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAnnouncementRequest) GetChannelOk() (*string, bool) {
	if o == nil || IsNil(o.Channel) {
		return nil, false
	}
	return o.Channel, true
}

// HasChannel returns a boolean if a field has been set.
func (o *UpdateAnnouncementRequest) HasChannel() bool {
	if o != nil && !IsNil(o.Channel) {
		return true
	}

	return false
}

// SetChannel gets a reference to the given string and assigns it to the Channel field.
func (o *UpdateAnnouncementRequest) SetChannel(v string) {
	o.Channel = &v
}

// GetIsPrioritized returns the IsPrioritized field value if set, zero value otherwise.
func (o *UpdateAnnouncementRequest) GetIsPrioritized() bool {
	if o == nil || IsNil(o.IsPrioritized) {
		var ret bool
		return ret
	}
	return *o.IsPrioritized
}

// GetIsPrioritizedOk returns a tuple with the IsPrioritized field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAnnouncementRequest) GetIsPrioritizedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPrioritized) {
		return nil, false
	}
	return o.IsPrioritized, true
}

// HasIsPrioritized returns a boolean if a field has been set.
func (o *UpdateAnnouncementRequest) HasIsPrioritized() bool {
	if o != nil && !IsNil(o.IsPrioritized) {
		return true
	}

	return false
}

// SetIsPrioritized gets a reference to the given bool and assigns it to the IsPrioritized field.
func (o *UpdateAnnouncementRequest) SetIsPrioritized(v bool) {
	o.IsPrioritized = &v
}

// GetViewUrl returns the ViewUrl field value if set, zero value otherwise.
func (o *UpdateAnnouncementRequest) GetViewUrl() string {
	if o == nil || IsNil(o.ViewUrl) {
		var ret string
		return ret
	}
	return *o.ViewUrl
}

// GetViewUrlOk returns a tuple with the ViewUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAnnouncementRequest) GetViewUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ViewUrl) {
		return nil, false
	}
	return o.ViewUrl, true
}

// HasViewUrl returns a boolean if a field has been set.
func (o *UpdateAnnouncementRequest) HasViewUrl() bool {
	if o != nil && !IsNil(o.ViewUrl) {
		return true
	}

	return false
}

// SetViewUrl gets a reference to the given string and assigns it to the ViewUrl field.
func (o *UpdateAnnouncementRequest) SetViewUrl(v string) {
	o.ViewUrl = &v
}

// GetId returns the Id field value
func (o *UpdateAnnouncementRequest) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UpdateAnnouncementRequest) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UpdateAnnouncementRequest) SetId(v int32) {
	o.Id = v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *UpdateAnnouncementRequest) GetData() AnnouncementCreateOrUpdateData {
	if o == nil || IsNil(o.Data) {
		var ret AnnouncementCreateOrUpdateData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAnnouncementRequest) GetDataOk() (*AnnouncementCreateOrUpdateData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *UpdateAnnouncementRequest) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given AnnouncementCreateOrUpdateData and assigns it to the Data field.
func (o *UpdateAnnouncementRequest) SetData(v AnnouncementCreateOrUpdateData) {
	o.Data = &v
}

func (o UpdateAnnouncementRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateAnnouncementRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["startTime"] = o.StartTime
	toSerialize["endTime"] = o.EndTime
	toSerialize["title"] = o.Title
	if !IsNil(o.Body) {
		toSerialize["body"] = o.Body
	}
	if !IsNil(o.Emoji) {
		toSerialize["emoji"] = o.Emoji
	}
	if !IsNil(o.Thumbnail) {
		toSerialize["thumbnail"] = o.Thumbnail
	}
	if !IsNil(o.Banner) {
		toSerialize["banner"] = o.Banner
	}
	if !IsNil(o.AudienceFilters) {
		toSerialize["audienceFilters"] = o.AudienceFilters
	}
	if !IsNil(o.SourceDocumentId) {
		toSerialize["sourceDocumentId"] = o.SourceDocumentId
	}
	if !IsNil(o.HideAttribution) {
		toSerialize["hideAttribution"] = o.HideAttribution
	}
	if !IsNil(o.Channel) {
		toSerialize["channel"] = o.Channel
	}
	if !IsNil(o.IsPrioritized) {
		toSerialize["isPrioritized"] = o.IsPrioritized
	}
	if !IsNil(o.ViewUrl) {
		toSerialize["viewUrl"] = o.ViewUrl
	}
	toSerialize["id"] = o.Id
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableUpdateAnnouncementRequest struct {
	value *UpdateAnnouncementRequest
	isSet bool
}

func (v NullableUpdateAnnouncementRequest) Get() *UpdateAnnouncementRequest {
	return v.value
}

func (v *NullableUpdateAnnouncementRequest) Set(val *UpdateAnnouncementRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateAnnouncementRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateAnnouncementRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateAnnouncementRequest(val *UpdateAnnouncementRequest) *NullableUpdateAnnouncementRequest {
	return &NullableUpdateAnnouncementRequest{value: val, isSet: true}
}

func (v NullableUpdateAnnouncementRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateAnnouncementRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


