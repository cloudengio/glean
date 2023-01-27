# GetUserCountResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserCount** | Pointer to **int32** | Number of users corresponding to the specified custom datasource. | [optional] 

## Methods

### NewGetUserCountResponse

`func NewGetUserCountResponse() *GetUserCountResponse`

NewGetUserCountResponse instantiates a new GetUserCountResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserCountResponseWithDefaults

`func NewGetUserCountResponseWithDefaults() *GetUserCountResponse`

NewGetUserCountResponseWithDefaults instantiates a new GetUserCountResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserCount

`func (o *GetUserCountResponse) GetUserCount() int32`

GetUserCount returns the UserCount field if non-nil, zero value otherwise.

### GetUserCountOk

`func (o *GetUserCountResponse) GetUserCountOk() (*int32, bool)`

GetUserCountOk returns a tuple with the UserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCount

`func (o *GetUserCountResponse) SetUserCount(v int32)`

SetUserCount sets UserCount field to given value.

### HasUserCount

`func (o *GetUserCountResponse) HasUserCount() bool`

HasUserCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


