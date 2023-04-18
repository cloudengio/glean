# GetAnswerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnswerResult** | Pointer to [**AnswerResult**](AnswerResult.md) |  | [optional] 
**Error** | Pointer to [**GetAnswerError**](GetAnswerError.md) |  | [optional] 
**Answer** | Pointer to [**Answer**](Answer.md) |  | [optional] 
**TrackingToken** | Pointer to **string** | DEPRECATED - use answerResult tracking token instead. An opaque token that represents this particular answer. To be used for /feedback reporting. | [optional] 

## Methods

### NewGetAnswerResponse

`func NewGetAnswerResponse() *GetAnswerResponse`

NewGetAnswerResponse instantiates a new GetAnswerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAnswerResponseWithDefaults

`func NewGetAnswerResponseWithDefaults() *GetAnswerResponse`

NewGetAnswerResponseWithDefaults instantiates a new GetAnswerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnswerResult

`func (o *GetAnswerResponse) GetAnswerResult() AnswerResult`

GetAnswerResult returns the AnswerResult field if non-nil, zero value otherwise.

### GetAnswerResultOk

`func (o *GetAnswerResponse) GetAnswerResultOk() (*AnswerResult, bool)`

GetAnswerResultOk returns a tuple with the AnswerResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswerResult

`func (o *GetAnswerResponse) SetAnswerResult(v AnswerResult)`

SetAnswerResult sets AnswerResult field to given value.

### HasAnswerResult

`func (o *GetAnswerResponse) HasAnswerResult() bool`

HasAnswerResult returns a boolean if a field has been set.

### GetError

`func (o *GetAnswerResponse) GetError() GetAnswerError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetAnswerResponse) GetErrorOk() (*GetAnswerError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetAnswerResponse) SetError(v GetAnswerError)`

SetError sets Error field to given value.

### HasError

`func (o *GetAnswerResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### GetAnswer

`func (o *GetAnswerResponse) GetAnswer() Answer`

GetAnswer returns the Answer field if non-nil, zero value otherwise.

### GetAnswerOk

`func (o *GetAnswerResponse) GetAnswerOk() (*Answer, bool)`

GetAnswerOk returns a tuple with the Answer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswer

`func (o *GetAnswerResponse) SetAnswer(v Answer)`

SetAnswer sets Answer field to given value.

### HasAnswer

`func (o *GetAnswerResponse) HasAnswer() bool`

HasAnswer returns a boolean if a field has been set.

### GetTrackingToken

`func (o *GetAnswerResponse) GetTrackingToken() string`

GetTrackingToken returns the TrackingToken field if non-nil, zero value otherwise.

### GetTrackingTokenOk

`func (o *GetAnswerResponse) GetTrackingTokenOk() (*string, bool)`

GetTrackingTokenOk returns a tuple with the TrackingToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingToken

`func (o *GetAnswerResponse) SetTrackingToken(v string)`

SetTrackingToken sets TrackingToken field to given value.

### HasTrackingToken

`func (o *GetAnswerResponse) HasTrackingToken() bool`

HasTrackingToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


