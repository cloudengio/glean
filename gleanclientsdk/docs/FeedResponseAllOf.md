# FeedResponseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrackingToken** | Pointer to **string** | An opaque token that represents this particular feed response. | [optional] 
**ServerTimestamp** | Pointer to **int32** | Server unix timestamp (in seconds since epoch UTC). | [optional] 
**Results** | Pointer to [**[]FeedResult**](FeedResult.md) |  | [optional] 
**BackendTimeMillis** | Pointer to **int64** | Time in milliseconds the backend took to respond to the request. | [optional] 
**ServerBuildVersion** | Pointer to **string** | Build versions to be rendered in debug mode. | [optional] 
**DatasourceAffinity** | Pointer to **map[string]float32** | A mapping from datasources to affinity of the user to each with scores. | [optional] 
**DebugInfo** | Pointer to [**FeedDebugInfo**](FeedDebugInfo.md) |  | [optional] 
**ManualFeedbackSignals** | Pointer to [**FeedManualFeedback**](FeedManualFeedback.md) |  | [optional] 
**CompanyResourcesCollectionId** | Pointer to **int32** | The unique ID of the collection for company resources. | [optional] 
**FacetResults** | Pointer to [**map[string][]FacetResult**](array.md) | Map from category to the list of facets that can be used to filter the entry&#39;s content. | [optional] 
**MentionsTimeWindowInHours** | Pointer to **int32** | The time window (in hours) used for generating user mentions. | [optional] 

## Methods

### NewFeedResponseAllOf

`func NewFeedResponseAllOf() *FeedResponseAllOf`

NewFeedResponseAllOf instantiates a new FeedResponseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeedResponseAllOfWithDefaults

`func NewFeedResponseAllOfWithDefaults() *FeedResponseAllOf`

NewFeedResponseAllOfWithDefaults instantiates a new FeedResponseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrackingToken

`func (o *FeedResponseAllOf) GetTrackingToken() string`

GetTrackingToken returns the TrackingToken field if non-nil, zero value otherwise.

### GetTrackingTokenOk

`func (o *FeedResponseAllOf) GetTrackingTokenOk() (*string, bool)`

GetTrackingTokenOk returns a tuple with the TrackingToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingToken

`func (o *FeedResponseAllOf) SetTrackingToken(v string)`

SetTrackingToken sets TrackingToken field to given value.

### HasTrackingToken

`func (o *FeedResponseAllOf) HasTrackingToken() bool`

HasTrackingToken returns a boolean if a field has been set.

### GetServerTimestamp

`func (o *FeedResponseAllOf) GetServerTimestamp() int32`

GetServerTimestamp returns the ServerTimestamp field if non-nil, zero value otherwise.

### GetServerTimestampOk

`func (o *FeedResponseAllOf) GetServerTimestampOk() (*int32, bool)`

GetServerTimestampOk returns a tuple with the ServerTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTimestamp

`func (o *FeedResponseAllOf) SetServerTimestamp(v int32)`

SetServerTimestamp sets ServerTimestamp field to given value.

### HasServerTimestamp

`func (o *FeedResponseAllOf) HasServerTimestamp() bool`

HasServerTimestamp returns a boolean if a field has been set.

### GetResults

`func (o *FeedResponseAllOf) GetResults() []FeedResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *FeedResponseAllOf) GetResultsOk() (*[]FeedResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *FeedResponseAllOf) SetResults(v []FeedResult)`

SetResults sets Results field to given value.

### HasResults

`func (o *FeedResponseAllOf) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetBackendTimeMillis

`func (o *FeedResponseAllOf) GetBackendTimeMillis() int64`

GetBackendTimeMillis returns the BackendTimeMillis field if non-nil, zero value otherwise.

### GetBackendTimeMillisOk

`func (o *FeedResponseAllOf) GetBackendTimeMillisOk() (*int64, bool)`

GetBackendTimeMillisOk returns a tuple with the BackendTimeMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendTimeMillis

`func (o *FeedResponseAllOf) SetBackendTimeMillis(v int64)`

SetBackendTimeMillis sets BackendTimeMillis field to given value.

### HasBackendTimeMillis

`func (o *FeedResponseAllOf) HasBackendTimeMillis() bool`

HasBackendTimeMillis returns a boolean if a field has been set.

### GetServerBuildVersion

`func (o *FeedResponseAllOf) GetServerBuildVersion() string`

GetServerBuildVersion returns the ServerBuildVersion field if non-nil, zero value otherwise.

### GetServerBuildVersionOk

`func (o *FeedResponseAllOf) GetServerBuildVersionOk() (*string, bool)`

GetServerBuildVersionOk returns a tuple with the ServerBuildVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerBuildVersion

`func (o *FeedResponseAllOf) SetServerBuildVersion(v string)`

SetServerBuildVersion sets ServerBuildVersion field to given value.

### HasServerBuildVersion

`func (o *FeedResponseAllOf) HasServerBuildVersion() bool`

HasServerBuildVersion returns a boolean if a field has been set.

### GetDatasourceAffinity

`func (o *FeedResponseAllOf) GetDatasourceAffinity() map[string]float32`

GetDatasourceAffinity returns the DatasourceAffinity field if non-nil, zero value otherwise.

### GetDatasourceAffinityOk

`func (o *FeedResponseAllOf) GetDatasourceAffinityOk() (*map[string]float32, bool)`

GetDatasourceAffinityOk returns a tuple with the DatasourceAffinity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasourceAffinity

`func (o *FeedResponseAllOf) SetDatasourceAffinity(v map[string]float32)`

SetDatasourceAffinity sets DatasourceAffinity field to given value.

### HasDatasourceAffinity

`func (o *FeedResponseAllOf) HasDatasourceAffinity() bool`

HasDatasourceAffinity returns a boolean if a field has been set.

### GetDebugInfo

`func (o *FeedResponseAllOf) GetDebugInfo() FeedDebugInfo`

GetDebugInfo returns the DebugInfo field if non-nil, zero value otherwise.

### GetDebugInfoOk

`func (o *FeedResponseAllOf) GetDebugInfoOk() (*FeedDebugInfo, bool)`

GetDebugInfoOk returns a tuple with the DebugInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugInfo

`func (o *FeedResponseAllOf) SetDebugInfo(v FeedDebugInfo)`

SetDebugInfo sets DebugInfo field to given value.

### HasDebugInfo

`func (o *FeedResponseAllOf) HasDebugInfo() bool`

HasDebugInfo returns a boolean if a field has been set.

### GetManualFeedbackSignals

`func (o *FeedResponseAllOf) GetManualFeedbackSignals() FeedManualFeedback`

GetManualFeedbackSignals returns the ManualFeedbackSignals field if non-nil, zero value otherwise.

### GetManualFeedbackSignalsOk

`func (o *FeedResponseAllOf) GetManualFeedbackSignalsOk() (*FeedManualFeedback, bool)`

GetManualFeedbackSignalsOk returns a tuple with the ManualFeedbackSignals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualFeedbackSignals

`func (o *FeedResponseAllOf) SetManualFeedbackSignals(v FeedManualFeedback)`

SetManualFeedbackSignals sets ManualFeedbackSignals field to given value.

### HasManualFeedbackSignals

`func (o *FeedResponseAllOf) HasManualFeedbackSignals() bool`

HasManualFeedbackSignals returns a boolean if a field has been set.

### GetCompanyResourcesCollectionId

`func (o *FeedResponseAllOf) GetCompanyResourcesCollectionId() int32`

GetCompanyResourcesCollectionId returns the CompanyResourcesCollectionId field if non-nil, zero value otherwise.

### GetCompanyResourcesCollectionIdOk

`func (o *FeedResponseAllOf) GetCompanyResourcesCollectionIdOk() (*int32, bool)`

GetCompanyResourcesCollectionIdOk returns a tuple with the CompanyResourcesCollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyResourcesCollectionId

`func (o *FeedResponseAllOf) SetCompanyResourcesCollectionId(v int32)`

SetCompanyResourcesCollectionId sets CompanyResourcesCollectionId field to given value.

### HasCompanyResourcesCollectionId

`func (o *FeedResponseAllOf) HasCompanyResourcesCollectionId() bool`

HasCompanyResourcesCollectionId returns a boolean if a field has been set.

### GetFacetResults

`func (o *FeedResponseAllOf) GetFacetResults() map[string][]FacetResult`

GetFacetResults returns the FacetResults field if non-nil, zero value otherwise.

### GetFacetResultsOk

`func (o *FeedResponseAllOf) GetFacetResultsOk() (*map[string][]FacetResult, bool)`

GetFacetResultsOk returns a tuple with the FacetResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacetResults

`func (o *FeedResponseAllOf) SetFacetResults(v map[string][]FacetResult)`

SetFacetResults sets FacetResults field to given value.

### HasFacetResults

`func (o *FeedResponseAllOf) HasFacetResults() bool`

HasFacetResults returns a boolean if a field has been set.

### GetMentionsTimeWindowInHours

`func (o *FeedResponseAllOf) GetMentionsTimeWindowInHours() int32`

GetMentionsTimeWindowInHours returns the MentionsTimeWindowInHours field if non-nil, zero value otherwise.

### GetMentionsTimeWindowInHoursOk

`func (o *FeedResponseAllOf) GetMentionsTimeWindowInHoursOk() (*int32, bool)`

GetMentionsTimeWindowInHoursOk returns a tuple with the MentionsTimeWindowInHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentionsTimeWindowInHours

`func (o *FeedResponseAllOf) SetMentionsTimeWindowInHours(v int32)`

SetMentionsTimeWindowInHours sets MentionsTimeWindowInHours field to given value.

### HasMentionsTimeWindowInHours

`func (o *FeedResponseAllOf) HasMentionsTimeWindowInHours() bool`

HasMentionsTimeWindowInHours returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


