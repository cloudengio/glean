# ClientAnalyticsEventTrackingParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to **int64** | Unix timestamp in millis for the client event. | [optional] 
**EventName** | Pointer to **string** |  | [optional] 
**Attribution** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**Datasource** | Pointer to **string** |  | [optional] 
**DocType** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**PagePath** | Pointer to **string** |  | [optional] 
**UiElement** | Pointer to **string** |  | [optional] 
**UtmSource** | Pointer to **string** |  | [optional] 
**Rate** | Pointer to **float32** | Sample rate applicable for this event. | [optional] 
**Stt** | Pointer to **string** | Session identifier. | [optional] 

## Methods

### NewClientAnalyticsEventTrackingParams

`func NewClientAnalyticsEventTrackingParams() *ClientAnalyticsEventTrackingParams`

NewClientAnalyticsEventTrackingParams instantiates a new ClientAnalyticsEventTrackingParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientAnalyticsEventTrackingParamsWithDefaults

`func NewClientAnalyticsEventTrackingParamsWithDefaults() *ClientAnalyticsEventTrackingParams`

NewClientAnalyticsEventTrackingParamsWithDefaults instantiates a new ClientAnalyticsEventTrackingParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *ClientAnalyticsEventTrackingParams) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ClientAnalyticsEventTrackingParams) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ClientAnalyticsEventTrackingParams) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ClientAnalyticsEventTrackingParams) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetEventName

`func (o *ClientAnalyticsEventTrackingParams) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *ClientAnalyticsEventTrackingParams) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *ClientAnalyticsEventTrackingParams) SetEventName(v string)`

SetEventName sets EventName field to given value.

### HasEventName

`func (o *ClientAnalyticsEventTrackingParams) HasEventName() bool`

HasEventName returns a boolean if a field has been set.

### GetAttribution

`func (o *ClientAnalyticsEventTrackingParams) GetAttribution() string`

GetAttribution returns the Attribution field if non-nil, zero value otherwise.

### GetAttributionOk

`func (o *ClientAnalyticsEventTrackingParams) GetAttributionOk() (*string, bool)`

GetAttributionOk returns a tuple with the Attribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribution

`func (o *ClientAnalyticsEventTrackingParams) SetAttribution(v string)`

SetAttribution sets Attribution field to given value.

### HasAttribution

`func (o *ClientAnalyticsEventTrackingParams) HasAttribution() bool`

HasAttribution returns a boolean if a field has been set.

### GetCategory

`func (o *ClientAnalyticsEventTrackingParams) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ClientAnalyticsEventTrackingParams) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ClientAnalyticsEventTrackingParams) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ClientAnalyticsEventTrackingParams) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDatasource

`func (o *ClientAnalyticsEventTrackingParams) GetDatasource() string`

GetDatasource returns the Datasource field if non-nil, zero value otherwise.

### GetDatasourceOk

`func (o *ClientAnalyticsEventTrackingParams) GetDatasourceOk() (*string, bool)`

GetDatasourceOk returns a tuple with the Datasource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasource

`func (o *ClientAnalyticsEventTrackingParams) SetDatasource(v string)`

SetDatasource sets Datasource field to given value.

### HasDatasource

`func (o *ClientAnalyticsEventTrackingParams) HasDatasource() bool`

HasDatasource returns a boolean if a field has been set.

### GetDocType

`func (o *ClientAnalyticsEventTrackingParams) GetDocType() string`

GetDocType returns the DocType field if non-nil, zero value otherwise.

### GetDocTypeOk

`func (o *ClientAnalyticsEventTrackingParams) GetDocTypeOk() (*string, bool)`

GetDocTypeOk returns a tuple with the DocType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocType

`func (o *ClientAnalyticsEventTrackingParams) SetDocType(v string)`

SetDocType sets DocType field to given value.

### HasDocType

`func (o *ClientAnalyticsEventTrackingParams) HasDocType() bool`

HasDocType returns a boolean if a field has been set.

### GetLabel

`func (o *ClientAnalyticsEventTrackingParams) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ClientAnalyticsEventTrackingParams) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ClientAnalyticsEventTrackingParams) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ClientAnalyticsEventTrackingParams) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetPagePath

`func (o *ClientAnalyticsEventTrackingParams) GetPagePath() string`

GetPagePath returns the PagePath field if non-nil, zero value otherwise.

### GetPagePathOk

`func (o *ClientAnalyticsEventTrackingParams) GetPagePathOk() (*string, bool)`

GetPagePathOk returns a tuple with the PagePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagePath

`func (o *ClientAnalyticsEventTrackingParams) SetPagePath(v string)`

SetPagePath sets PagePath field to given value.

### HasPagePath

`func (o *ClientAnalyticsEventTrackingParams) HasPagePath() bool`

HasPagePath returns a boolean if a field has been set.

### GetUiElement

`func (o *ClientAnalyticsEventTrackingParams) GetUiElement() string`

GetUiElement returns the UiElement field if non-nil, zero value otherwise.

### GetUiElementOk

`func (o *ClientAnalyticsEventTrackingParams) GetUiElementOk() (*string, bool)`

GetUiElementOk returns a tuple with the UiElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiElement

`func (o *ClientAnalyticsEventTrackingParams) SetUiElement(v string)`

SetUiElement sets UiElement field to given value.

### HasUiElement

`func (o *ClientAnalyticsEventTrackingParams) HasUiElement() bool`

HasUiElement returns a boolean if a field has been set.

### GetUtmSource

`func (o *ClientAnalyticsEventTrackingParams) GetUtmSource() string`

GetUtmSource returns the UtmSource field if non-nil, zero value otherwise.

### GetUtmSourceOk

`func (o *ClientAnalyticsEventTrackingParams) GetUtmSourceOk() (*string, bool)`

GetUtmSourceOk returns a tuple with the UtmSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtmSource

`func (o *ClientAnalyticsEventTrackingParams) SetUtmSource(v string)`

SetUtmSource sets UtmSource field to given value.

### HasUtmSource

`func (o *ClientAnalyticsEventTrackingParams) HasUtmSource() bool`

HasUtmSource returns a boolean if a field has been set.

### GetRate

`func (o *ClientAnalyticsEventTrackingParams) GetRate() float32`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *ClientAnalyticsEventTrackingParams) GetRateOk() (*float32, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *ClientAnalyticsEventTrackingParams) SetRate(v float32)`

SetRate sets Rate field to given value.

### HasRate

`func (o *ClientAnalyticsEventTrackingParams) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetStt

`func (o *ClientAnalyticsEventTrackingParams) GetStt() string`

GetStt returns the Stt field if non-nil, zero value otherwise.

### GetSttOk

`func (o *ClientAnalyticsEventTrackingParams) GetSttOk() (*string, bool)`

GetSttOk returns a tuple with the Stt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStt

`func (o *ClientAnalyticsEventTrackingParams) SetStt(v string)`

SetStt sets Stt field to given value.

### HasStt

`func (o *ClientAnalyticsEventTrackingParams) HasStt() bool`

HasStt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


