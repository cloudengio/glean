# ClientAnalyticsEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrackingParams** | Pointer to [**ClientAnalyticsEventTrackingParams**](ClientAnalyticsEventTrackingParams.md) |  | [optional] 
**StringParams** | Pointer to **map[string]string** | Additional string parameters associated with the analytics event, contents will depend on event category and type. | [optional] 
**NumberParams** | Pointer to **map[string]float32** | Additional numerical parameters associated with the analytics event, contents will depend on event category and type. | [optional] 
**BoolParams** | Pointer to **map[string]bool** | Additional boolean parameters associated with the analytics event, contents will depend on event category and type. | [optional] 

## Methods

### NewClientAnalyticsEvent

`func NewClientAnalyticsEvent() *ClientAnalyticsEvent`

NewClientAnalyticsEvent instantiates a new ClientAnalyticsEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientAnalyticsEventWithDefaults

`func NewClientAnalyticsEventWithDefaults() *ClientAnalyticsEvent`

NewClientAnalyticsEventWithDefaults instantiates a new ClientAnalyticsEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrackingParams

`func (o *ClientAnalyticsEvent) GetTrackingParams() ClientAnalyticsEventTrackingParams`

GetTrackingParams returns the TrackingParams field if non-nil, zero value otherwise.

### GetTrackingParamsOk

`func (o *ClientAnalyticsEvent) GetTrackingParamsOk() (*ClientAnalyticsEventTrackingParams, bool)`

GetTrackingParamsOk returns a tuple with the TrackingParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingParams

`func (o *ClientAnalyticsEvent) SetTrackingParams(v ClientAnalyticsEventTrackingParams)`

SetTrackingParams sets TrackingParams field to given value.

### HasTrackingParams

`func (o *ClientAnalyticsEvent) HasTrackingParams() bool`

HasTrackingParams returns a boolean if a field has been set.

### GetStringParams

`func (o *ClientAnalyticsEvent) GetStringParams() map[string]string`

GetStringParams returns the StringParams field if non-nil, zero value otherwise.

### GetStringParamsOk

`func (o *ClientAnalyticsEvent) GetStringParamsOk() (*map[string]string, bool)`

GetStringParamsOk returns a tuple with the StringParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStringParams

`func (o *ClientAnalyticsEvent) SetStringParams(v map[string]string)`

SetStringParams sets StringParams field to given value.

### HasStringParams

`func (o *ClientAnalyticsEvent) HasStringParams() bool`

HasStringParams returns a boolean if a field has been set.

### GetNumberParams

`func (o *ClientAnalyticsEvent) GetNumberParams() map[string]float32`

GetNumberParams returns the NumberParams field if non-nil, zero value otherwise.

### GetNumberParamsOk

`func (o *ClientAnalyticsEvent) GetNumberParamsOk() (*map[string]float32, bool)`

GetNumberParamsOk returns a tuple with the NumberParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberParams

`func (o *ClientAnalyticsEvent) SetNumberParams(v map[string]float32)`

SetNumberParams sets NumberParams field to given value.

### HasNumberParams

`func (o *ClientAnalyticsEvent) HasNumberParams() bool`

HasNumberParams returns a boolean if a field has been set.

### GetBoolParams

`func (o *ClientAnalyticsEvent) GetBoolParams() map[string]bool`

GetBoolParams returns the BoolParams field if non-nil, zero value otherwise.

### GetBoolParamsOk

`func (o *ClientAnalyticsEvent) GetBoolParamsOk() (*map[string]bool, bool)`

GetBoolParamsOk returns a tuple with the BoolParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoolParams

`func (o *ClientAnalyticsEvent) SetBoolParams(v map[string]bool)`

SetBoolParams sets BoolParams field to given value.

### HasBoolParams

`func (o *ClientAnalyticsEvent) HasBoolParams() bool`

HasBoolParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


