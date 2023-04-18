# CalendarEventAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The calendar event id | [optional] 
**Url** | Pointer to **string** | A permalink for this calendar event | [optional] 
**Attendees** | Pointer to [**CalendarAttendees**](CalendarAttendees.md) |  | [optional] 
**Location** | Pointer to **string** | The location that this event is taking place at. | [optional] 
**ConferenceData** | Pointer to [**ConferenceData**](ConferenceData.md) |  | [optional] 
**Description** | Pointer to **string** | The HTML description of the event. | [optional] 
**Datasource** | Pointer to **string** | The app or other repository type from which the event was extracted | [optional] 

## Methods

### NewCalendarEventAllOf

`func NewCalendarEventAllOf() *CalendarEventAllOf`

NewCalendarEventAllOf instantiates a new CalendarEventAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCalendarEventAllOfWithDefaults

`func NewCalendarEventAllOfWithDefaults() *CalendarEventAllOf`

NewCalendarEventAllOfWithDefaults instantiates a new CalendarEventAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CalendarEventAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CalendarEventAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CalendarEventAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CalendarEventAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *CalendarEventAllOf) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CalendarEventAllOf) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CalendarEventAllOf) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CalendarEventAllOf) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetAttendees

`func (o *CalendarEventAllOf) GetAttendees() CalendarAttendees`

GetAttendees returns the Attendees field if non-nil, zero value otherwise.

### GetAttendeesOk

`func (o *CalendarEventAllOf) GetAttendeesOk() (*CalendarAttendees, bool)`

GetAttendeesOk returns a tuple with the Attendees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttendees

`func (o *CalendarEventAllOf) SetAttendees(v CalendarAttendees)`

SetAttendees sets Attendees field to given value.

### HasAttendees

`func (o *CalendarEventAllOf) HasAttendees() bool`

HasAttendees returns a boolean if a field has been set.

### GetLocation

`func (o *CalendarEventAllOf) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *CalendarEventAllOf) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *CalendarEventAllOf) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *CalendarEventAllOf) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetConferenceData

`func (o *CalendarEventAllOf) GetConferenceData() ConferenceData`

GetConferenceData returns the ConferenceData field if non-nil, zero value otherwise.

### GetConferenceDataOk

`func (o *CalendarEventAllOf) GetConferenceDataOk() (*ConferenceData, bool)`

GetConferenceDataOk returns a tuple with the ConferenceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConferenceData

`func (o *CalendarEventAllOf) SetConferenceData(v ConferenceData)`

SetConferenceData sets ConferenceData field to given value.

### HasConferenceData

`func (o *CalendarEventAllOf) HasConferenceData() bool`

HasConferenceData returns a boolean if a field has been set.

### GetDescription

`func (o *CalendarEventAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CalendarEventAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CalendarEventAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CalendarEventAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDatasource

`func (o *CalendarEventAllOf) GetDatasource() string`

GetDatasource returns the Datasource field if non-nil, zero value otherwise.

### GetDatasourceOk

`func (o *CalendarEventAllOf) GetDatasourceOk() (*string, bool)`

GetDatasourceOk returns a tuple with the Datasource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasource

`func (o *CalendarEventAllOf) SetDatasource(v string)`

SetDatasource sets Datasource field to given value.

### HasDatasource

`func (o *CalendarEventAllOf) HasDatasource() bool`

HasDatasource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


