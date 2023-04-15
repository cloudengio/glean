# SearchDebugInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FormattedDebugQuery** | Pointer to **string** | A formatted string that represents the parsed query. | [optional] 
**SearchConfigurationJson** | Pointer to **string** | JSON of the search config | [optional] 
**ExtraArgsJson** | Pointer to **string** | JSON of the extra args | [optional] 
**ParsedQueryJson** | Pointer to **string** | JSON for the parsed query, to be used as an override. | [optional] 
**DebugParsedQueryJson** | Pointer to **string** | JSON for the parsed query with debugging signals (e.g. syns and spellchecks) | [optional] 
**DebugScholasticJson** | Pointer to **string** | JSON containing Scholastic data (query embeddings, doc similarities). | [optional] 
**DebugQPMetadataJson** | Pointer to **string** | JSON containing QP metadata | [optional] 
**DebugScholasticMetadataJson** | Pointer to **string** | JSON containing Scholastic metadata | [optional] 
**DebugMinedSamplesJson** | Pointer to **string** | JSON containing mined Intelligence samples | [optional] 
**DebugRetrievalElasticQuery** | Pointer to **string** | JSON containing Elastic retrieval query | [optional] 
**DebugSnippetsElasticQuery** | Pointer to **string** | JSON containing Elastic snippets query | [optional] 
**ElasticPerformanceString** | Pointer to **string** | A string showing performance information returned by elastic. | [optional] 
**ScoringLegendString** | Pointer to **string** | A legend of what the functions are when computing the backend score | [optional] 
**ResultsDebugString** | Pointer to **string** | Additional debugging details associated with the request. | [optional] 
**DebugKeywordGenerationJson** | Pointer to **string** | JSON containing Keyword Generation data for debugging purposes. | [optional] 

## Methods

### NewSearchDebugInfo

`func NewSearchDebugInfo() *SearchDebugInfo`

NewSearchDebugInfo instantiates a new SearchDebugInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchDebugInfoWithDefaults

`func NewSearchDebugInfoWithDefaults() *SearchDebugInfo`

NewSearchDebugInfoWithDefaults instantiates a new SearchDebugInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormattedDebugQuery

`func (o *SearchDebugInfo) GetFormattedDebugQuery() string`

GetFormattedDebugQuery returns the FormattedDebugQuery field if non-nil, zero value otherwise.

### GetFormattedDebugQueryOk

`func (o *SearchDebugInfo) GetFormattedDebugQueryOk() (*string, bool)`

GetFormattedDebugQueryOk returns a tuple with the FormattedDebugQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedDebugQuery

`func (o *SearchDebugInfo) SetFormattedDebugQuery(v string)`

SetFormattedDebugQuery sets FormattedDebugQuery field to given value.

### HasFormattedDebugQuery

`func (o *SearchDebugInfo) HasFormattedDebugQuery() bool`

HasFormattedDebugQuery returns a boolean if a field has been set.

### GetSearchConfigurationJson

`func (o *SearchDebugInfo) GetSearchConfigurationJson() string`

GetSearchConfigurationJson returns the SearchConfigurationJson field if non-nil, zero value otherwise.

### GetSearchConfigurationJsonOk

`func (o *SearchDebugInfo) GetSearchConfigurationJsonOk() (*string, bool)`

GetSearchConfigurationJsonOk returns a tuple with the SearchConfigurationJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchConfigurationJson

`func (o *SearchDebugInfo) SetSearchConfigurationJson(v string)`

SetSearchConfigurationJson sets SearchConfigurationJson field to given value.

### HasSearchConfigurationJson

`func (o *SearchDebugInfo) HasSearchConfigurationJson() bool`

HasSearchConfigurationJson returns a boolean if a field has been set.

### GetExtraArgsJson

`func (o *SearchDebugInfo) GetExtraArgsJson() string`

GetExtraArgsJson returns the ExtraArgsJson field if non-nil, zero value otherwise.

### GetExtraArgsJsonOk

`func (o *SearchDebugInfo) GetExtraArgsJsonOk() (*string, bool)`

GetExtraArgsJsonOk returns a tuple with the ExtraArgsJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraArgsJson

`func (o *SearchDebugInfo) SetExtraArgsJson(v string)`

SetExtraArgsJson sets ExtraArgsJson field to given value.

### HasExtraArgsJson

`func (o *SearchDebugInfo) HasExtraArgsJson() bool`

HasExtraArgsJson returns a boolean if a field has been set.

### GetParsedQueryJson

`func (o *SearchDebugInfo) GetParsedQueryJson() string`

GetParsedQueryJson returns the ParsedQueryJson field if non-nil, zero value otherwise.

### GetParsedQueryJsonOk

`func (o *SearchDebugInfo) GetParsedQueryJsonOk() (*string, bool)`

GetParsedQueryJsonOk returns a tuple with the ParsedQueryJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParsedQueryJson

`func (o *SearchDebugInfo) SetParsedQueryJson(v string)`

SetParsedQueryJson sets ParsedQueryJson field to given value.

### HasParsedQueryJson

`func (o *SearchDebugInfo) HasParsedQueryJson() bool`

HasParsedQueryJson returns a boolean if a field has been set.

### GetDebugParsedQueryJson

`func (o *SearchDebugInfo) GetDebugParsedQueryJson() string`

GetDebugParsedQueryJson returns the DebugParsedQueryJson field if non-nil, zero value otherwise.

### GetDebugParsedQueryJsonOk

`func (o *SearchDebugInfo) GetDebugParsedQueryJsonOk() (*string, bool)`

GetDebugParsedQueryJsonOk returns a tuple with the DebugParsedQueryJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugParsedQueryJson

`func (o *SearchDebugInfo) SetDebugParsedQueryJson(v string)`

SetDebugParsedQueryJson sets DebugParsedQueryJson field to given value.

### HasDebugParsedQueryJson

`func (o *SearchDebugInfo) HasDebugParsedQueryJson() bool`

HasDebugParsedQueryJson returns a boolean if a field has been set.

### GetDebugScholasticJson

`func (o *SearchDebugInfo) GetDebugScholasticJson() string`

GetDebugScholasticJson returns the DebugScholasticJson field if non-nil, zero value otherwise.

### GetDebugScholasticJsonOk

`func (o *SearchDebugInfo) GetDebugScholasticJsonOk() (*string, bool)`

GetDebugScholasticJsonOk returns a tuple with the DebugScholasticJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugScholasticJson

`func (o *SearchDebugInfo) SetDebugScholasticJson(v string)`

SetDebugScholasticJson sets DebugScholasticJson field to given value.

### HasDebugScholasticJson

`func (o *SearchDebugInfo) HasDebugScholasticJson() bool`

HasDebugScholasticJson returns a boolean if a field has been set.

### GetDebugQPMetadataJson

`func (o *SearchDebugInfo) GetDebugQPMetadataJson() string`

GetDebugQPMetadataJson returns the DebugQPMetadataJson field if non-nil, zero value otherwise.

### GetDebugQPMetadataJsonOk

`func (o *SearchDebugInfo) GetDebugQPMetadataJsonOk() (*string, bool)`

GetDebugQPMetadataJsonOk returns a tuple with the DebugQPMetadataJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugQPMetadataJson

`func (o *SearchDebugInfo) SetDebugQPMetadataJson(v string)`

SetDebugQPMetadataJson sets DebugQPMetadataJson field to given value.

### HasDebugQPMetadataJson

`func (o *SearchDebugInfo) HasDebugQPMetadataJson() bool`

HasDebugQPMetadataJson returns a boolean if a field has been set.

### GetDebugScholasticMetadataJson

`func (o *SearchDebugInfo) GetDebugScholasticMetadataJson() string`

GetDebugScholasticMetadataJson returns the DebugScholasticMetadataJson field if non-nil, zero value otherwise.

### GetDebugScholasticMetadataJsonOk

`func (o *SearchDebugInfo) GetDebugScholasticMetadataJsonOk() (*string, bool)`

GetDebugScholasticMetadataJsonOk returns a tuple with the DebugScholasticMetadataJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugScholasticMetadataJson

`func (o *SearchDebugInfo) SetDebugScholasticMetadataJson(v string)`

SetDebugScholasticMetadataJson sets DebugScholasticMetadataJson field to given value.

### HasDebugScholasticMetadataJson

`func (o *SearchDebugInfo) HasDebugScholasticMetadataJson() bool`

HasDebugScholasticMetadataJson returns a boolean if a field has been set.

### GetDebugMinedSamplesJson

`func (o *SearchDebugInfo) GetDebugMinedSamplesJson() string`

GetDebugMinedSamplesJson returns the DebugMinedSamplesJson field if non-nil, zero value otherwise.

### GetDebugMinedSamplesJsonOk

`func (o *SearchDebugInfo) GetDebugMinedSamplesJsonOk() (*string, bool)`

GetDebugMinedSamplesJsonOk returns a tuple with the DebugMinedSamplesJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugMinedSamplesJson

`func (o *SearchDebugInfo) SetDebugMinedSamplesJson(v string)`

SetDebugMinedSamplesJson sets DebugMinedSamplesJson field to given value.

### HasDebugMinedSamplesJson

`func (o *SearchDebugInfo) HasDebugMinedSamplesJson() bool`

HasDebugMinedSamplesJson returns a boolean if a field has been set.

### GetDebugRetrievalElasticQuery

`func (o *SearchDebugInfo) GetDebugRetrievalElasticQuery() string`

GetDebugRetrievalElasticQuery returns the DebugRetrievalElasticQuery field if non-nil, zero value otherwise.

### GetDebugRetrievalElasticQueryOk

`func (o *SearchDebugInfo) GetDebugRetrievalElasticQueryOk() (*string, bool)`

GetDebugRetrievalElasticQueryOk returns a tuple with the DebugRetrievalElasticQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugRetrievalElasticQuery

`func (o *SearchDebugInfo) SetDebugRetrievalElasticQuery(v string)`

SetDebugRetrievalElasticQuery sets DebugRetrievalElasticQuery field to given value.

### HasDebugRetrievalElasticQuery

`func (o *SearchDebugInfo) HasDebugRetrievalElasticQuery() bool`

HasDebugRetrievalElasticQuery returns a boolean if a field has been set.

### GetDebugSnippetsElasticQuery

`func (o *SearchDebugInfo) GetDebugSnippetsElasticQuery() string`

GetDebugSnippetsElasticQuery returns the DebugSnippetsElasticQuery field if non-nil, zero value otherwise.

### GetDebugSnippetsElasticQueryOk

`func (o *SearchDebugInfo) GetDebugSnippetsElasticQueryOk() (*string, bool)`

GetDebugSnippetsElasticQueryOk returns a tuple with the DebugSnippetsElasticQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugSnippetsElasticQuery

`func (o *SearchDebugInfo) SetDebugSnippetsElasticQuery(v string)`

SetDebugSnippetsElasticQuery sets DebugSnippetsElasticQuery field to given value.

### HasDebugSnippetsElasticQuery

`func (o *SearchDebugInfo) HasDebugSnippetsElasticQuery() bool`

HasDebugSnippetsElasticQuery returns a boolean if a field has been set.

### GetElasticPerformanceString

`func (o *SearchDebugInfo) GetElasticPerformanceString() string`

GetElasticPerformanceString returns the ElasticPerformanceString field if non-nil, zero value otherwise.

### GetElasticPerformanceStringOk

`func (o *SearchDebugInfo) GetElasticPerformanceStringOk() (*string, bool)`

GetElasticPerformanceStringOk returns a tuple with the ElasticPerformanceString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElasticPerformanceString

`func (o *SearchDebugInfo) SetElasticPerformanceString(v string)`

SetElasticPerformanceString sets ElasticPerformanceString field to given value.

### HasElasticPerformanceString

`func (o *SearchDebugInfo) HasElasticPerformanceString() bool`

HasElasticPerformanceString returns a boolean if a field has been set.

### GetScoringLegendString

`func (o *SearchDebugInfo) GetScoringLegendString() string`

GetScoringLegendString returns the ScoringLegendString field if non-nil, zero value otherwise.

### GetScoringLegendStringOk

`func (o *SearchDebugInfo) GetScoringLegendStringOk() (*string, bool)`

GetScoringLegendStringOk returns a tuple with the ScoringLegendString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoringLegendString

`func (o *SearchDebugInfo) SetScoringLegendString(v string)`

SetScoringLegendString sets ScoringLegendString field to given value.

### HasScoringLegendString

`func (o *SearchDebugInfo) HasScoringLegendString() bool`

HasScoringLegendString returns a boolean if a field has been set.

### GetResultsDebugString

`func (o *SearchDebugInfo) GetResultsDebugString() string`

GetResultsDebugString returns the ResultsDebugString field if non-nil, zero value otherwise.

### GetResultsDebugStringOk

`func (o *SearchDebugInfo) GetResultsDebugStringOk() (*string, bool)`

GetResultsDebugStringOk returns a tuple with the ResultsDebugString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsDebugString

`func (o *SearchDebugInfo) SetResultsDebugString(v string)`

SetResultsDebugString sets ResultsDebugString field to given value.

### HasResultsDebugString

`func (o *SearchDebugInfo) HasResultsDebugString() bool`

HasResultsDebugString returns a boolean if a field has been set.

### GetDebugKeywordGenerationJson

`func (o *SearchDebugInfo) GetDebugKeywordGenerationJson() string`

GetDebugKeywordGenerationJson returns the DebugKeywordGenerationJson field if non-nil, zero value otherwise.

### GetDebugKeywordGenerationJsonOk

`func (o *SearchDebugInfo) GetDebugKeywordGenerationJsonOk() (*string, bool)`

GetDebugKeywordGenerationJsonOk returns a tuple with the DebugKeywordGenerationJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugKeywordGenerationJson

`func (o *SearchDebugInfo) SetDebugKeywordGenerationJson(v string)`

SetDebugKeywordGenerationJson sets DebugKeywordGenerationJson field to given value.

### HasDebugKeywordGenerationJson

`func (o *SearchDebugInfo) HasDebugKeywordGenerationJson() bool`

HasDebugKeywordGenerationJson returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


