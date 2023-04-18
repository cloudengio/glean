# ClientAnalyticsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommonParams** | Pointer to [**ClientAnalyticsCommonParams**](ClientAnalyticsCommonParams.md) |  | [optional] 
**Events** | Pointer to [**[]ClientAnalyticsEvent**](ClientAnalyticsEvent.md) |  | [optional] 

## Methods

### NewClientAnalyticsRequest

`func NewClientAnalyticsRequest() *ClientAnalyticsRequest`

NewClientAnalyticsRequest instantiates a new ClientAnalyticsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientAnalyticsRequestWithDefaults

`func NewClientAnalyticsRequestWithDefaults() *ClientAnalyticsRequest`

NewClientAnalyticsRequestWithDefaults instantiates a new ClientAnalyticsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommonParams

`func (o *ClientAnalyticsRequest) GetCommonParams() ClientAnalyticsCommonParams`

GetCommonParams returns the CommonParams field if non-nil, zero value otherwise.

### GetCommonParamsOk

`func (o *ClientAnalyticsRequest) GetCommonParamsOk() (*ClientAnalyticsCommonParams, bool)`

GetCommonParamsOk returns a tuple with the CommonParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonParams

`func (o *ClientAnalyticsRequest) SetCommonParams(v ClientAnalyticsCommonParams)`

SetCommonParams sets CommonParams field to given value.

### HasCommonParams

`func (o *ClientAnalyticsRequest) HasCommonParams() bool`

HasCommonParams returns a boolean if a field has been set.

### GetEvents

`func (o *ClientAnalyticsRequest) GetEvents() []ClientAnalyticsEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *ClientAnalyticsRequest) GetEventsOk() (*[]ClientAnalyticsEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *ClientAnalyticsRequest) SetEvents(v []ClientAnalyticsEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *ClientAnalyticsRequest) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


