# SearchResultDebugInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FinalScore** | Pointer to **float64** |  | [optional] 
**ScholasticScore** | Pointer to **float64** |  | [optional] 
**ScholasticContent** | Pointer to **string** |  | [optional] 
**ScholasticRetrievalOnly** | Pointer to **bool** |  | [optional] 
**Explanation** | Pointer to **string** |  | [optional] 
**SnippetScore** | Pointer to **float64** |  | [optional] 
**Position** | Pointer to **int64** |  | [optional] 
**EvalDocumentDescriptorEncoded** | Pointer to **string** | Information about boundaries / endpoints of documents, e.g., messages in a conversation - only used for evals and always scrubbed + without PII. | [optional] 
**IdHash** | Pointer to **string** | The document id hash. | [optional] 

## Methods

### NewSearchResultDebugInfo

`func NewSearchResultDebugInfo() *SearchResultDebugInfo`

NewSearchResultDebugInfo instantiates a new SearchResultDebugInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchResultDebugInfoWithDefaults

`func NewSearchResultDebugInfoWithDefaults() *SearchResultDebugInfo`

NewSearchResultDebugInfoWithDefaults instantiates a new SearchResultDebugInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFinalScore

`func (o *SearchResultDebugInfo) GetFinalScore() float64`

GetFinalScore returns the FinalScore field if non-nil, zero value otherwise.

### GetFinalScoreOk

`func (o *SearchResultDebugInfo) GetFinalScoreOk() (*float64, bool)`

GetFinalScoreOk returns a tuple with the FinalScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalScore

`func (o *SearchResultDebugInfo) SetFinalScore(v float64)`

SetFinalScore sets FinalScore field to given value.

### HasFinalScore

`func (o *SearchResultDebugInfo) HasFinalScore() bool`

HasFinalScore returns a boolean if a field has been set.

### GetScholasticScore

`func (o *SearchResultDebugInfo) GetScholasticScore() float64`

GetScholasticScore returns the ScholasticScore field if non-nil, zero value otherwise.

### GetScholasticScoreOk

`func (o *SearchResultDebugInfo) GetScholasticScoreOk() (*float64, bool)`

GetScholasticScoreOk returns a tuple with the ScholasticScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScholasticScore

`func (o *SearchResultDebugInfo) SetScholasticScore(v float64)`

SetScholasticScore sets ScholasticScore field to given value.

### HasScholasticScore

`func (o *SearchResultDebugInfo) HasScholasticScore() bool`

HasScholasticScore returns a boolean if a field has been set.

### GetScholasticContent

`func (o *SearchResultDebugInfo) GetScholasticContent() string`

GetScholasticContent returns the ScholasticContent field if non-nil, zero value otherwise.

### GetScholasticContentOk

`func (o *SearchResultDebugInfo) GetScholasticContentOk() (*string, bool)`

GetScholasticContentOk returns a tuple with the ScholasticContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScholasticContent

`func (o *SearchResultDebugInfo) SetScholasticContent(v string)`

SetScholasticContent sets ScholasticContent field to given value.

### HasScholasticContent

`func (o *SearchResultDebugInfo) HasScholasticContent() bool`

HasScholasticContent returns a boolean if a field has been set.

### GetScholasticRetrievalOnly

`func (o *SearchResultDebugInfo) GetScholasticRetrievalOnly() bool`

GetScholasticRetrievalOnly returns the ScholasticRetrievalOnly field if non-nil, zero value otherwise.

### GetScholasticRetrievalOnlyOk

`func (o *SearchResultDebugInfo) GetScholasticRetrievalOnlyOk() (*bool, bool)`

GetScholasticRetrievalOnlyOk returns a tuple with the ScholasticRetrievalOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScholasticRetrievalOnly

`func (o *SearchResultDebugInfo) SetScholasticRetrievalOnly(v bool)`

SetScholasticRetrievalOnly sets ScholasticRetrievalOnly field to given value.

### HasScholasticRetrievalOnly

`func (o *SearchResultDebugInfo) HasScholasticRetrievalOnly() bool`

HasScholasticRetrievalOnly returns a boolean if a field has been set.

### GetExplanation

`func (o *SearchResultDebugInfo) GetExplanation() string`

GetExplanation returns the Explanation field if non-nil, zero value otherwise.

### GetExplanationOk

`func (o *SearchResultDebugInfo) GetExplanationOk() (*string, bool)`

GetExplanationOk returns a tuple with the Explanation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplanation

`func (o *SearchResultDebugInfo) SetExplanation(v string)`

SetExplanation sets Explanation field to given value.

### HasExplanation

`func (o *SearchResultDebugInfo) HasExplanation() bool`

HasExplanation returns a boolean if a field has been set.

### GetSnippetScore

`func (o *SearchResultDebugInfo) GetSnippetScore() float64`

GetSnippetScore returns the SnippetScore field if non-nil, zero value otherwise.

### GetSnippetScoreOk

`func (o *SearchResultDebugInfo) GetSnippetScoreOk() (*float64, bool)`

GetSnippetScoreOk returns a tuple with the SnippetScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnippetScore

`func (o *SearchResultDebugInfo) SetSnippetScore(v float64)`

SetSnippetScore sets SnippetScore field to given value.

### HasSnippetScore

`func (o *SearchResultDebugInfo) HasSnippetScore() bool`

HasSnippetScore returns a boolean if a field has been set.

### GetPosition

`func (o *SearchResultDebugInfo) GetPosition() int64`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *SearchResultDebugInfo) GetPositionOk() (*int64, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *SearchResultDebugInfo) SetPosition(v int64)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *SearchResultDebugInfo) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetEvalDocumentDescriptorEncoded

`func (o *SearchResultDebugInfo) GetEvalDocumentDescriptorEncoded() string`

GetEvalDocumentDescriptorEncoded returns the EvalDocumentDescriptorEncoded field if non-nil, zero value otherwise.

### GetEvalDocumentDescriptorEncodedOk

`func (o *SearchResultDebugInfo) GetEvalDocumentDescriptorEncodedOk() (*string, bool)`

GetEvalDocumentDescriptorEncodedOk returns a tuple with the EvalDocumentDescriptorEncoded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvalDocumentDescriptorEncoded

`func (o *SearchResultDebugInfo) SetEvalDocumentDescriptorEncoded(v string)`

SetEvalDocumentDescriptorEncoded sets EvalDocumentDescriptorEncoded field to given value.

### HasEvalDocumentDescriptorEncoded

`func (o *SearchResultDebugInfo) HasEvalDocumentDescriptorEncoded() bool`

HasEvalDocumentDescriptorEncoded returns a boolean if a field has been set.

### GetIdHash

`func (o *SearchResultDebugInfo) GetIdHash() string`

GetIdHash returns the IdHash field if non-nil, zero value otherwise.

### GetIdHashOk

`func (o *SearchResultDebugInfo) GetIdHashOk() (*string, bool)`

GetIdHashOk returns a tuple with the IdHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdHash

`func (o *SearchResultDebugInfo) SetIdHash(v string)`

SetIdHash sets IdHash field to given value.

### HasIdHash

`func (o *SearchResultDebugInfo) HasIdHash() bool`

HasIdHash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


