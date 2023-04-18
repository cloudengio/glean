# GetEventsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | **[]string** | The ids of the calendar events to be retrieved. | 
**AuthTokens** | Pointer to [**[]AuthToken**](AuthToken.md) | Auth tokens if client-side authentication. | [optional] 

## Methods

### NewGetEventsRequest

`func NewGetEventsRequest(ids []string, ) *GetEventsRequest`

NewGetEventsRequest instantiates a new GetEventsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEventsRequestWithDefaults

`func NewGetEventsRequestWithDefaults() *GetEventsRequest`

NewGetEventsRequestWithDefaults instantiates a new GetEventsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *GetEventsRequest) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *GetEventsRequest) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *GetEventsRequest) SetIds(v []string)`

SetIds sets Ids field to given value.


### GetAuthTokens

`func (o *GetEventsRequest) GetAuthTokens() []AuthToken`

GetAuthTokens returns the AuthTokens field if non-nil, zero value otherwise.

### GetAuthTokensOk

`func (o *GetEventsRequest) GetAuthTokensOk() (*[]AuthToken, bool)`

GetAuthTokensOk returns a tuple with the AuthTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthTokens

`func (o *GetEventsRequest) SetAuthTokens(v []AuthToken)`

SetAuthTokens sets AuthTokens field to given value.

### HasAuthTokens

`func (o *GetEventsRequest) HasAuthTokens() bool`

HasAuthTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


