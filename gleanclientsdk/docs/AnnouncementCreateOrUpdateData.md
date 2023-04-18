# AnnouncementCreateOrUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | If editing an existing announcement, the announcement&#39;s ID. | [optional] 
**Title** | **string** | The headline of the announcement. | 
**StructuredText** | Pointer to [**StructuredText**](StructuredText.md) |  | [optional] 
**Emoji** | Pointer to **string** | An emoji used to indicate the nature of the announcement. | [optional] 
**StartTime** | Pointer to **time.Time** | The date and time at which the announcement becomes active. If omitted, the announement will become active immediately. | [optional] 
**EndTime** | **time.Time** | The date and time at which the announcement expires. | 
**AudienceFilters** | Pointer to [**[]FacetFilter**](FacetFilter.md) | Filters which restrict who should see the announcement. Values are taken from the corresponding filters in people search. | [optional] 

## Methods

### NewAnnouncementCreateOrUpdateData

`func NewAnnouncementCreateOrUpdateData(title string, endTime time.Time, ) *AnnouncementCreateOrUpdateData`

NewAnnouncementCreateOrUpdateData instantiates a new AnnouncementCreateOrUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnnouncementCreateOrUpdateDataWithDefaults

`func NewAnnouncementCreateOrUpdateDataWithDefaults() *AnnouncementCreateOrUpdateData`

NewAnnouncementCreateOrUpdateDataWithDefaults instantiates a new AnnouncementCreateOrUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AnnouncementCreateOrUpdateData) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AnnouncementCreateOrUpdateData) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AnnouncementCreateOrUpdateData) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AnnouncementCreateOrUpdateData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *AnnouncementCreateOrUpdateData) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AnnouncementCreateOrUpdateData) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AnnouncementCreateOrUpdateData) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetStructuredText

`func (o *AnnouncementCreateOrUpdateData) GetStructuredText() StructuredText`

GetStructuredText returns the StructuredText field if non-nil, zero value otherwise.

### GetStructuredTextOk

`func (o *AnnouncementCreateOrUpdateData) GetStructuredTextOk() (*StructuredText, bool)`

GetStructuredTextOk returns a tuple with the StructuredText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructuredText

`func (o *AnnouncementCreateOrUpdateData) SetStructuredText(v StructuredText)`

SetStructuredText sets StructuredText field to given value.

### HasStructuredText

`func (o *AnnouncementCreateOrUpdateData) HasStructuredText() bool`

HasStructuredText returns a boolean if a field has been set.

### GetEmoji

`func (o *AnnouncementCreateOrUpdateData) GetEmoji() string`

GetEmoji returns the Emoji field if non-nil, zero value otherwise.

### GetEmojiOk

`func (o *AnnouncementCreateOrUpdateData) GetEmojiOk() (*string, bool)`

GetEmojiOk returns a tuple with the Emoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmoji

`func (o *AnnouncementCreateOrUpdateData) SetEmoji(v string)`

SetEmoji sets Emoji field to given value.

### HasEmoji

`func (o *AnnouncementCreateOrUpdateData) HasEmoji() bool`

HasEmoji returns a boolean if a field has been set.

### GetStartTime

`func (o *AnnouncementCreateOrUpdateData) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *AnnouncementCreateOrUpdateData) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *AnnouncementCreateOrUpdateData) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *AnnouncementCreateOrUpdateData) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *AnnouncementCreateOrUpdateData) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *AnnouncementCreateOrUpdateData) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *AnnouncementCreateOrUpdateData) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.


### GetAudienceFilters

`func (o *AnnouncementCreateOrUpdateData) GetAudienceFilters() []FacetFilter`

GetAudienceFilters returns the AudienceFilters field if non-nil, zero value otherwise.

### GetAudienceFiltersOk

`func (o *AnnouncementCreateOrUpdateData) GetAudienceFiltersOk() (*[]FacetFilter, bool)`

GetAudienceFiltersOk returns a tuple with the AudienceFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudienceFilters

`func (o *AnnouncementCreateOrUpdateData) SetAudienceFilters(v []FacetFilter)`

SetAudienceFilters sets AudienceFilters field to given value.

### HasAudienceFilters

`func (o *AnnouncementCreateOrUpdateData) HasAudienceFilters() bool`

HasAudienceFilters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


