/*
Glean Client API

# Introduction These are the public APIs to enable implementing a custom client interface to the Glean system.  # Usage guidelines This API is evolving fast. Glean will provide advance notice of any planned backwards incompatible changes along with a 6-month sunset period for anything that requires developers to adopt the new versions. 

API version: 0.9.0
Contact: support@glean.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gleanclientsdk

import (
	"encoding/json"
)

// checks if the SearchDebugInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchDebugInfo{}

// SearchDebugInfo struct for SearchDebugInfo
type SearchDebugInfo struct {
	// A formatted string that represents the parsed query.
	FormattedDebugQuery *string `json:"formattedDebugQuery,omitempty"`
	// JSON of the search config
	SearchConfigurationJson *string `json:"searchConfigurationJson,omitempty"`
	// JSON of the extra args
	ExtraArgsJson *string `json:"extraArgsJson,omitempty"`
	// JSON for the parsed query, to be used as an override.
	ParsedQueryJson *string `json:"parsedQueryJson,omitempty"`
	// JSON for the parsed query with debugging signals (e.g. syns and spellchecks)
	DebugParsedQueryJson *string `json:"debugParsedQueryJson,omitempty"`
	// JSON containing Scholastic data (query embeddings, doc similarities).
	DebugScholasticJson *string `json:"debugScholasticJson,omitempty"`
	// JSON containing QP metadata
	DebugQPMetadataJson *string `json:"debugQPMetadataJson,omitempty"`
	// JSON containing Scholastic metadata
	DebugScholasticMetadataJson *string `json:"debugScholasticMetadataJson,omitempty"`
	// JSON containing mined Intelligence samples
	DebugMinedSamplesJson *string `json:"debugMinedSamplesJson,omitempty"`
	// JSON containing Elastic retrieval query
	DebugRetrievalElasticQuery *string `json:"debugRetrievalElasticQuery,omitempty"`
	// JSON containing Elastic snippets query
	DebugSnippetsElasticQuery *string `json:"debugSnippetsElasticQuery,omitempty"`
	// A string showing performance information returned by elastic.
	ElasticPerformanceString *string `json:"elasticPerformanceString,omitempty"`
	// A legend of what the functions are when computing the backend score
	ScoringLegendString *string `json:"scoringLegendString,omitempty"`
	// Additional debugging details associated with the request.
	ResultsDebugString *string `json:"resultsDebugString,omitempty"`
	// JSON containing Keyword Generation data for debugging purposes.
	DebugKeywordGenerationJson *string `json:"debugKeywordGenerationJson,omitempty"`
}

// NewSearchDebugInfo instantiates a new SearchDebugInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchDebugInfo() *SearchDebugInfo {
	this := SearchDebugInfo{}
	return &this
}

// NewSearchDebugInfoWithDefaults instantiates a new SearchDebugInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchDebugInfoWithDefaults() *SearchDebugInfo {
	this := SearchDebugInfo{}
	return &this
}

// GetFormattedDebugQuery returns the FormattedDebugQuery field value if set, zero value otherwise.
func (o *SearchDebugInfo) GetFormattedDebugQuery() string {
	if o == nil || IsNil(o.FormattedDebugQuery) {
		var ret string
		return ret
	}
	return *o.FormattedDebugQuery
}

// GetFormattedDebugQueryOk returns a tuple with the FormattedDebugQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchDebugInfo) GetFormattedDebugQueryOk() (*string, bool) {
	if o == nil || IsNil(o.FormattedDebugQuery) {
		return nil, false
	}
	return o.FormattedDebugQuery, true
}

// HasFormattedDebugQuery returns a boolean if a field has been set.
func (o *SearchDebugInfo) HasFormattedDebugQuery() bool {
	if o != nil && !IsNil(o.FormattedDebugQuery) {
		return true
	}

	return false
}

// SetFormattedDebugQuery gets a reference to the given string and assigns it to the FormattedDebugQuery field.
func (o *SearchDebugInfo) SetFormattedDebugQuery(v string) {
	o.FormattedDebugQuery = &v
}

// GetSearchConfigurationJson returns the SearchConfigurationJson field value if set, zero value otherwise.
func (o *SearchDebugInfo) GetSearchConfigurationJson() string {
	if o == nil || IsNil(o.SearchConfigurationJson) {
		var ret string
		return ret
	}
	return *o.SearchConfigurationJson
}

// GetSearchConfigurationJsonOk returns a tuple with the SearchConfigurationJson field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchDebugInfo) GetSearchConfigurationJsonOk() (*string, bool) {
	if o == nil || IsNil(o.SearchConfigurationJson) {
		return nil, false
	}
	return o.SearchConfigurationJson, true
}

// HasSearchConfigurationJson returns a boolean if a field has been set.
func (o *SearchDebugInfo) HasSearchConfigurationJson() bool {
	if o != nil && !IsNil(o.SearchConfigurationJson) {
		return true
	}

	return false
}

// SetSearchConfigurationJson gets a reference to the given string and assigns it to the SearchConfigurationJson field.
func (o *SearchDebugInfo) SetSearchConfigurationJson(v string) {
	o.SearchConfigurationJson = &v
}

// GetExtraArgsJson returns the ExtraArgsJson field value if set, zero value otherwise.
func (o *SearchDebugInfo) GetExtraArgsJson() string {
	if o == nil || IsNil(o.ExtraArgsJson) {
		var ret string
		return ret
	}
	return *o.ExtraArgsJson
}

// GetExtraArgsJsonOk returns a tuple with the ExtraArgsJson field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchDebugInfo) GetExtraArgsJsonOk() (*string, bool) {
	if o == nil || IsNil(o.ExtraArgsJson) {
		return nil, false
	}
	return o.ExtraArgsJson, true
}

// HasExtraArgsJson returns a boolean if a field has been set.
func (o *SearchDebugInfo) HasExtraArgsJson() bool {
	if o != nil && !IsNil(o.ExtraArgsJson) {
		return true
	}

	return false
}

// SetExtraArgsJson gets a reference to the given string and assigns it to the ExtraArgsJson field.
func (o *SearchDebugInfo) SetExtraArgsJson(v string) {
	o.ExtraArgsJson = &v
}

// GetParsedQueryJson returns the ParsedQueryJson field value if set, zero value otherwise.
func (o *SearchDebugInfo) GetParsedQueryJson() string {
	if o == nil || IsNil(o.ParsedQueryJson) {
		var ret string
		return ret
	}
	return *o.ParsedQueryJson
}

// GetParsedQueryJsonOk returns a tuple with the ParsedQueryJson field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchDebugInfo) GetParsedQueryJsonOk() (*string, bool) {
	if o == nil || IsNil(o.ParsedQueryJson) {
		return nil, false
	}
	return o.ParsedQueryJson, true
}

// HasParsedQueryJson returns a boolean if a field has been set.
func (o *SearchDebugInfo) HasParsedQueryJson() bool {
	if o != nil && !IsNil(o.ParsedQueryJson) {
		return true
	}

	return false
}

// SetParsedQueryJson gets a reference to the given string and assigns it to the ParsedQueryJson field.
func (o *SearchDebugInfo) SetParsedQueryJson(v string) {
	o.ParsedQueryJson = &v
}

// GetDebugParsedQueryJson returns the DebugParsedQueryJson field value if set, zero value otherwise.
func (o *SearchDebugInfo) GetDebugParsedQueryJson() string {
	if o == nil || IsNil(o.DebugParsedQueryJson) {
		var ret string
		return ret
	}
	return *o.DebugParsedQueryJson
}

// GetDebugParsedQueryJsonOk returns a tuple with the DebugParsedQueryJson field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchDebugInfo) GetDebugParsedQueryJsonOk() (*string, bool) {
	if o == nil || IsNil(o.DebugParsedQueryJson) {
		return nil, false
	}
	return o.DebugParsedQueryJson, true
}

// HasDebugParsedQueryJson returns a boolean if a field has been set.
func (o *SearchDebugInfo) HasDebugParsedQueryJson() bool {
	if o != nil && !IsNil(o.DebugParsedQueryJson) {
		return true
	}

	return false
}

// SetDebugParsedQueryJson gets a reference to the given string and assigns it to the DebugParsedQueryJson field.
func (o *SearchDebugInfo) SetDebugParsedQueryJson(v string) {
	o.DebugParsedQueryJson = &v
}

// GetDebugScholasticJson returns the DebugScholasticJson field value if set, zero value otherwise.
func (o *SearchDebugInfo) GetDebugScholasticJson() string {
	if o == nil || IsNil(o.DebugScholasticJson) {
		var ret string
		return ret
	}
	return *o.DebugScholasticJson
}

// GetDebugScholasticJsonOk returns a tuple with the DebugScholasticJson field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchDebugInfo) GetDebugScholasticJsonOk() (*string, bool) {
	if o == nil || IsNil(o.DebugScholasticJson) {
		return nil, false
	}
	return o.DebugScholasticJson, true
}

// HasDebugScholasticJson returns a boolean if a field has been set.
func (o *SearchDebugInfo) HasDebugScholasticJson() bool {
	if o != nil && !IsNil(o.DebugScholasticJson) {
		return true
	}

	return false
}

// SetDebugScholasticJson gets a reference to the given string and assigns it to the DebugScholasticJson field.
func (o *SearchDebugInfo) SetDebugScholasticJson(v string) {
	o.DebugScholasticJson = &v
}

// GetDebugQPMetadataJson returns the DebugQPMetadataJson field value if set, zero value otherwise.
func (o *SearchDebugInfo) GetDebugQPMetadataJson() string {
	if o == nil || IsNil(o.DebugQPMetadataJson) {
		var ret string
		return ret
	}
	return *o.DebugQPMetadataJson
}

// GetDebugQPMetadataJsonOk returns a tuple with the DebugQPMetadataJson field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchDebugInfo) GetDebugQPMetadataJsonOk() (*string, bool) {
	if o == nil || IsNil(o.DebugQPMetadataJson) {
		return nil, false
	}
	return o.DebugQPMetadataJson, true
}

// HasDebugQPMetadataJson returns a boolean if a field has been set.
func (o *SearchDebugInfo) HasDebugQPMetadataJson() bool {
	if o != nil && !IsNil(o.DebugQPMetadataJson) {
		return true
	}

	return false
}

// SetDebugQPMetadataJson gets a reference to the given string and assigns it to the DebugQPMetadataJson field.
func (o *SearchDebugInfo) SetDebugQPMetadataJson(v string) {
	o.DebugQPMetadataJson = &v
}

// GetDebugScholasticMetadataJson returns the DebugScholasticMetadataJson field value if set, zero value otherwise.
func (o *SearchDebugInfo) GetDebugScholasticMetadataJson() string {
	if o == nil || IsNil(o.DebugScholasticMetadataJson) {
		var ret string
		return ret
	}
	return *o.DebugScholasticMetadataJson
}

// GetDebugScholasticMetadataJsonOk returns a tuple with the DebugScholasticMetadataJson field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchDebugInfo) GetDebugScholasticMetadataJsonOk() (*string, bool) {
	if o == nil || IsNil(o.DebugScholasticMetadataJson) {
		return nil, false
	}
	return o.DebugScholasticMetadataJson, true
}

// HasDebugScholasticMetadataJson returns a boolean if a field has been set.
func (o *SearchDebugInfo) HasDebugScholasticMetadataJson() bool {
	if o != nil && !IsNil(o.DebugScholasticMetadataJson) {
		return true
	}

	return false
}

// SetDebugScholasticMetadataJson gets a reference to the given string and assigns it to the DebugScholasticMetadataJson field.
func (o *SearchDebugInfo) SetDebugScholasticMetadataJson(v string) {
	o.DebugScholasticMetadataJson = &v
}

// GetDebugMinedSamplesJson returns the DebugMinedSamplesJson field value if set, zero value otherwise.
func (o *SearchDebugInfo) GetDebugMinedSamplesJson() string {
	if o == nil || IsNil(o.DebugMinedSamplesJson) {
		var ret string
		return ret
	}
	return *o.DebugMinedSamplesJson
}

// GetDebugMinedSamplesJsonOk returns a tuple with the DebugMinedSamplesJson field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchDebugInfo) GetDebugMinedSamplesJsonOk() (*string, bool) {
	if o == nil || IsNil(o.DebugMinedSamplesJson) {
		return nil, false
	}
	return o.DebugMinedSamplesJson, true
}

// HasDebugMinedSamplesJson returns a boolean if a field has been set.
func (o *SearchDebugInfo) HasDebugMinedSamplesJson() bool {
	if o != nil && !IsNil(o.DebugMinedSamplesJson) {
		return true
	}

	return false
}

// SetDebugMinedSamplesJson gets a reference to the given string and assigns it to the DebugMinedSamplesJson field.
func (o *SearchDebugInfo) SetDebugMinedSamplesJson(v string) {
	o.DebugMinedSamplesJson = &v
}

// GetDebugRetrievalElasticQuery returns the DebugRetrievalElasticQuery field value if set, zero value otherwise.
func (o *SearchDebugInfo) GetDebugRetrievalElasticQuery() string {
	if o == nil || IsNil(o.DebugRetrievalElasticQuery) {
		var ret string
		return ret
	}
	return *o.DebugRetrievalElasticQuery
}

// GetDebugRetrievalElasticQueryOk returns a tuple with the DebugRetrievalElasticQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchDebugInfo) GetDebugRetrievalElasticQueryOk() (*string, bool) {
	if o == nil || IsNil(o.DebugRetrievalElasticQuery) {
		return nil, false
	}
	return o.DebugRetrievalElasticQuery, true
}

// HasDebugRetrievalElasticQuery returns a boolean if a field has been set.
func (o *SearchDebugInfo) HasDebugRetrievalElasticQuery() bool {
	if o != nil && !IsNil(o.DebugRetrievalElasticQuery) {
		return true
	}

	return false
}

// SetDebugRetrievalElasticQuery gets a reference to the given string and assigns it to the DebugRetrievalElasticQuery field.
func (o *SearchDebugInfo) SetDebugRetrievalElasticQuery(v string) {
	o.DebugRetrievalElasticQuery = &v
}

// GetDebugSnippetsElasticQuery returns the DebugSnippetsElasticQuery field value if set, zero value otherwise.
func (o *SearchDebugInfo) GetDebugSnippetsElasticQuery() string {
	if o == nil || IsNil(o.DebugSnippetsElasticQuery) {
		var ret string
		return ret
	}
	return *o.DebugSnippetsElasticQuery
}

// GetDebugSnippetsElasticQueryOk returns a tuple with the DebugSnippetsElasticQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchDebugInfo) GetDebugSnippetsElasticQueryOk() (*string, bool) {
	if o == nil || IsNil(o.DebugSnippetsElasticQuery) {
		return nil, false
	}
	return o.DebugSnippetsElasticQuery, true
}

// HasDebugSnippetsElasticQuery returns a boolean if a field has been set.
func (o *SearchDebugInfo) HasDebugSnippetsElasticQuery() bool {
	if o != nil && !IsNil(o.DebugSnippetsElasticQuery) {
		return true
	}

	return false
}

// SetDebugSnippetsElasticQuery gets a reference to the given string and assigns it to the DebugSnippetsElasticQuery field.
func (o *SearchDebugInfo) SetDebugSnippetsElasticQuery(v string) {
	o.DebugSnippetsElasticQuery = &v
}

// GetElasticPerformanceString returns the ElasticPerformanceString field value if set, zero value otherwise.
func (o *SearchDebugInfo) GetElasticPerformanceString() string {
	if o == nil || IsNil(o.ElasticPerformanceString) {
		var ret string
		return ret
	}
	return *o.ElasticPerformanceString
}

// GetElasticPerformanceStringOk returns a tuple with the ElasticPerformanceString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchDebugInfo) GetElasticPerformanceStringOk() (*string, bool) {
	if o == nil || IsNil(o.ElasticPerformanceString) {
		return nil, false
	}
	return o.ElasticPerformanceString, true
}

// HasElasticPerformanceString returns a boolean if a field has been set.
func (o *SearchDebugInfo) HasElasticPerformanceString() bool {
	if o != nil && !IsNil(o.ElasticPerformanceString) {
		return true
	}

	return false
}

// SetElasticPerformanceString gets a reference to the given string and assigns it to the ElasticPerformanceString field.
func (o *SearchDebugInfo) SetElasticPerformanceString(v string) {
	o.ElasticPerformanceString = &v
}

// GetScoringLegendString returns the ScoringLegendString field value if set, zero value otherwise.
func (o *SearchDebugInfo) GetScoringLegendString() string {
	if o == nil || IsNil(o.ScoringLegendString) {
		var ret string
		return ret
	}
	return *o.ScoringLegendString
}

// GetScoringLegendStringOk returns a tuple with the ScoringLegendString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchDebugInfo) GetScoringLegendStringOk() (*string, bool) {
	if o == nil || IsNil(o.ScoringLegendString) {
		return nil, false
	}
	return o.ScoringLegendString, true
}

// HasScoringLegendString returns a boolean if a field has been set.
func (o *SearchDebugInfo) HasScoringLegendString() bool {
	if o != nil && !IsNil(o.ScoringLegendString) {
		return true
	}

	return false
}

// SetScoringLegendString gets a reference to the given string and assigns it to the ScoringLegendString field.
func (o *SearchDebugInfo) SetScoringLegendString(v string) {
	o.ScoringLegendString = &v
}

// GetResultsDebugString returns the ResultsDebugString field value if set, zero value otherwise.
func (o *SearchDebugInfo) GetResultsDebugString() string {
	if o == nil || IsNil(o.ResultsDebugString) {
		var ret string
		return ret
	}
	return *o.ResultsDebugString
}

// GetResultsDebugStringOk returns a tuple with the ResultsDebugString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchDebugInfo) GetResultsDebugStringOk() (*string, bool) {
	if o == nil || IsNil(o.ResultsDebugString) {
		return nil, false
	}
	return o.ResultsDebugString, true
}

// HasResultsDebugString returns a boolean if a field has been set.
func (o *SearchDebugInfo) HasResultsDebugString() bool {
	if o != nil && !IsNil(o.ResultsDebugString) {
		return true
	}

	return false
}

// SetResultsDebugString gets a reference to the given string and assigns it to the ResultsDebugString field.
func (o *SearchDebugInfo) SetResultsDebugString(v string) {
	o.ResultsDebugString = &v
}

// GetDebugKeywordGenerationJson returns the DebugKeywordGenerationJson field value if set, zero value otherwise.
func (o *SearchDebugInfo) GetDebugKeywordGenerationJson() string {
	if o == nil || IsNil(o.DebugKeywordGenerationJson) {
		var ret string
		return ret
	}
	return *o.DebugKeywordGenerationJson
}

// GetDebugKeywordGenerationJsonOk returns a tuple with the DebugKeywordGenerationJson field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchDebugInfo) GetDebugKeywordGenerationJsonOk() (*string, bool) {
	if o == nil || IsNil(o.DebugKeywordGenerationJson) {
		return nil, false
	}
	return o.DebugKeywordGenerationJson, true
}

// HasDebugKeywordGenerationJson returns a boolean if a field has been set.
func (o *SearchDebugInfo) HasDebugKeywordGenerationJson() bool {
	if o != nil && !IsNil(o.DebugKeywordGenerationJson) {
		return true
	}

	return false
}

// SetDebugKeywordGenerationJson gets a reference to the given string and assigns it to the DebugKeywordGenerationJson field.
func (o *SearchDebugInfo) SetDebugKeywordGenerationJson(v string) {
	o.DebugKeywordGenerationJson = &v
}

func (o SearchDebugInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchDebugInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FormattedDebugQuery) {
		toSerialize["formattedDebugQuery"] = o.FormattedDebugQuery
	}
	if !IsNil(o.SearchConfigurationJson) {
		toSerialize["searchConfigurationJson"] = o.SearchConfigurationJson
	}
	if !IsNil(o.ExtraArgsJson) {
		toSerialize["extraArgsJson"] = o.ExtraArgsJson
	}
	if !IsNil(o.ParsedQueryJson) {
		toSerialize["parsedQueryJson"] = o.ParsedQueryJson
	}
	if !IsNil(o.DebugParsedQueryJson) {
		toSerialize["debugParsedQueryJson"] = o.DebugParsedQueryJson
	}
	if !IsNil(o.DebugScholasticJson) {
		toSerialize["debugScholasticJson"] = o.DebugScholasticJson
	}
	if !IsNil(o.DebugQPMetadataJson) {
		toSerialize["debugQPMetadataJson"] = o.DebugQPMetadataJson
	}
	if !IsNil(o.DebugScholasticMetadataJson) {
		toSerialize["debugScholasticMetadataJson"] = o.DebugScholasticMetadataJson
	}
	if !IsNil(o.DebugMinedSamplesJson) {
		toSerialize["debugMinedSamplesJson"] = o.DebugMinedSamplesJson
	}
	if !IsNil(o.DebugRetrievalElasticQuery) {
		toSerialize["debugRetrievalElasticQuery"] = o.DebugRetrievalElasticQuery
	}
	if !IsNil(o.DebugSnippetsElasticQuery) {
		toSerialize["debugSnippetsElasticQuery"] = o.DebugSnippetsElasticQuery
	}
	if !IsNil(o.ElasticPerformanceString) {
		toSerialize["elasticPerformanceString"] = o.ElasticPerformanceString
	}
	if !IsNil(o.ScoringLegendString) {
		toSerialize["scoringLegendString"] = o.ScoringLegendString
	}
	if !IsNil(o.ResultsDebugString) {
		toSerialize["resultsDebugString"] = o.ResultsDebugString
	}
	if !IsNil(o.DebugKeywordGenerationJson) {
		toSerialize["debugKeywordGenerationJson"] = o.DebugKeywordGenerationJson
	}
	return toSerialize, nil
}

type NullableSearchDebugInfo struct {
	value *SearchDebugInfo
	isSet bool
}

func (v NullableSearchDebugInfo) Get() *SearchDebugInfo {
	return v.value
}

func (v *NullableSearchDebugInfo) Set(val *SearchDebugInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchDebugInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchDebugInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchDebugInfo(val *SearchDebugInfo) *NullableSearchDebugInfo {
	return &NullableSearchDebugInfo{value: val, isSet: true}
}

func (v NullableSearchDebugInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchDebugInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

