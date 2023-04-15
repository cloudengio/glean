# RecommendationsRequestOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatasourceFilter** | Pointer to **string** | Filter results to a single datasource name (e.g. gmail, slack). All results are returned if missing. | [optional] 
**ContextType** | Pointer to **string** | Identifier for what JSON schema to expect for &#x60;context&#x60;. | [optional] 
**Context** | Pointer to **string** | JSON containing client context like case subject input.  Expects schema based on &#x60;contextType&#x60;. | [optional] 
**ResultProminence** | Pointer to [**[]SearchResultProminenceEnum**](SearchResultProminenceEnum.md) | The types of prominence wanted in results returned. Default is any type. | [optional] 

## Methods

### NewRecommendationsRequestOptions

`func NewRecommendationsRequestOptions() *RecommendationsRequestOptions`

NewRecommendationsRequestOptions instantiates a new RecommendationsRequestOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecommendationsRequestOptionsWithDefaults

`func NewRecommendationsRequestOptionsWithDefaults() *RecommendationsRequestOptions`

NewRecommendationsRequestOptionsWithDefaults instantiates a new RecommendationsRequestOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatasourceFilter

`func (o *RecommendationsRequestOptions) GetDatasourceFilter() string`

GetDatasourceFilter returns the DatasourceFilter field if non-nil, zero value otherwise.

### GetDatasourceFilterOk

`func (o *RecommendationsRequestOptions) GetDatasourceFilterOk() (*string, bool)`

GetDatasourceFilterOk returns a tuple with the DatasourceFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasourceFilter

`func (o *RecommendationsRequestOptions) SetDatasourceFilter(v string)`

SetDatasourceFilter sets DatasourceFilter field to given value.

### HasDatasourceFilter

`func (o *RecommendationsRequestOptions) HasDatasourceFilter() bool`

HasDatasourceFilter returns a boolean if a field has been set.

### GetContextType

`func (o *RecommendationsRequestOptions) GetContextType() string`

GetContextType returns the ContextType field if non-nil, zero value otherwise.

### GetContextTypeOk

`func (o *RecommendationsRequestOptions) GetContextTypeOk() (*string, bool)`

GetContextTypeOk returns a tuple with the ContextType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextType

`func (o *RecommendationsRequestOptions) SetContextType(v string)`

SetContextType sets ContextType field to given value.

### HasContextType

`func (o *RecommendationsRequestOptions) HasContextType() bool`

HasContextType returns a boolean if a field has been set.

### GetContext

`func (o *RecommendationsRequestOptions) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *RecommendationsRequestOptions) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *RecommendationsRequestOptions) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *RecommendationsRequestOptions) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetResultProminence

`func (o *RecommendationsRequestOptions) GetResultProminence() []SearchResultProminenceEnum`

GetResultProminence returns the ResultProminence field if non-nil, zero value otherwise.

### GetResultProminenceOk

`func (o *RecommendationsRequestOptions) GetResultProminenceOk() (*[]SearchResultProminenceEnum, bool)`

GetResultProminenceOk returns a tuple with the ResultProminence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultProminence

`func (o *RecommendationsRequestOptions) SetResultProminence(v []SearchResultProminenceEnum)`

SetResultProminence sets ResultProminence field to given value.

### HasResultProminence

`func (o *RecommendationsRequestOptions) HasResultProminence() bool`

HasResultProminence returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


