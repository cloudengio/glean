# SummarizeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Summary** | Pointer to [**Summary**](Summary.md) |  | [optional] 
**DebugInfo** | Pointer to **string** | Debug details for this summary if debug is enabled | [optional] 
**TrackingToken** | Pointer to **string** | An opaque token that represents this summary in this particular query. To be used for /feedback reporting. | [optional] 

## Methods

### NewSummarizeResponse

`func NewSummarizeResponse() *SummarizeResponse`

NewSummarizeResponse instantiates a new SummarizeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSummarizeResponseWithDefaults

`func NewSummarizeResponseWithDefaults() *SummarizeResponse`

NewSummarizeResponseWithDefaults instantiates a new SummarizeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSummary

`func (o *SummarizeResponse) GetSummary() Summary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *SummarizeResponse) GetSummaryOk() (*Summary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *SummarizeResponse) SetSummary(v Summary)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *SummarizeResponse) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetDebugInfo

`func (o *SummarizeResponse) GetDebugInfo() string`

GetDebugInfo returns the DebugInfo field if non-nil, zero value otherwise.

### GetDebugInfoOk

`func (o *SummarizeResponse) GetDebugInfoOk() (*string, bool)`

GetDebugInfoOk returns a tuple with the DebugInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugInfo

`func (o *SummarizeResponse) SetDebugInfo(v string)`

SetDebugInfo sets DebugInfo field to given value.

### HasDebugInfo

`func (o *SummarizeResponse) HasDebugInfo() bool`

HasDebugInfo returns a boolean if a field has been set.

### GetTrackingToken

`func (o *SummarizeResponse) GetTrackingToken() string`

GetTrackingToken returns the TrackingToken field if non-nil, zero value otherwise.

### GetTrackingTokenOk

`func (o *SummarizeResponse) GetTrackingTokenOk() (*string, bool)`

GetTrackingTokenOk returns a tuple with the TrackingToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingToken

`func (o *SummarizeResponse) SetTrackingToken(v string)`

SetTrackingToken sets TrackingToken field to given value.

### HasTrackingToken

`func (o *SummarizeResponse) HasTrackingToken() bool`

HasTrackingToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


