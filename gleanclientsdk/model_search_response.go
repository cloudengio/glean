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

// checks if the SearchResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchResponse{}

// SearchResponse struct for SearchResponse
type SearchResponse struct {
	// A token that should be passed for additional requests related to this request (such as more results requests).
	TrackingToken *string `json:"trackingToken,omitempty"`
	SessionInfo *SessionInfo `json:"sessionInfo,omitempty"`
	Results []SearchResult `json:"results,omitempty"`
	StructuredResults []StructuredResult `json:"structuredResults,omitempty"`
	GeneratedQnaResult *GeneratedQna `json:"generatedQnaResult,omitempty"`
	DebugInfo *SearchDebugInfo `json:"debugInfo,omitempty"`
	ErrorInfo *ErrorInfo `json:"errorInfo,omitempty"`
	// A platform-generated request ID to correlate backend logs.
	RequestID *string `json:"requestID,omitempty"`
	// Time in milliseconds the backend took to respond to the request.
	BackendTimeMillis *int64 `json:"backendTimeMillis,omitempty"`
	// List of experiment ids for the corresponding request.
	ExperimentIds []int64 `json:"experimentIds,omitempty"`
	Metadata *SearchResponseMetadata `json:"metadata,omitempty"`
	FacetResults []FacetResult `json:"facetResults,omitempty"`
	// All result tabs available for the current query. Populated if QUERY_METADATA is specified in the request.
	ResultTabs []ResultTab `json:"resultTabs,omitempty"`
	// The unique IDs of the result tabs to which this response belongs.
	ResultTabIds []string `json:"resultTabIds,omitempty"`
	ResultsDescription *ResultsDescription `json:"resultsDescription,omitempty"`
	// The actual applied facet filters based on the operators and facetFilters in the query. Useful for mapping typed operators to visual facets.
	RewrittenFacetFilters []FacetFilter `json:"rewrittenFacetFilters,omitempty"`
	// Cursor that indicates the start of the next page of results. To be passed in \"more\" requests for this query.
	Cursor *string `json:"cursor,omitempty"`
	// Whether more results are available. Use cursor to retrieve them.
	HasMoreResults *bool `json:"hasMoreResults,omitempty"`
}

// NewSearchResponse instantiates a new SearchResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchResponse() *SearchResponse {
	this := SearchResponse{}
	return &this
}

// NewSearchResponseWithDefaults instantiates a new SearchResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchResponseWithDefaults() *SearchResponse {
	this := SearchResponse{}
	return &this
}

// GetTrackingToken returns the TrackingToken field value if set, zero value otherwise.
func (o *SearchResponse) GetTrackingToken() string {
	if o == nil || IsNil(o.TrackingToken) {
		var ret string
		return ret
	}
	return *o.TrackingToken
}

// GetTrackingTokenOk returns a tuple with the TrackingToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResponse) GetTrackingTokenOk() (*string, bool) {
	if o == nil || IsNil(o.TrackingToken) {
		return nil, false
	}
	return o.TrackingToken, true
}

// HasTrackingToken returns a boolean if a field has been set.
func (o *SearchResponse) HasTrackingToken() bool {
	if o != nil && !IsNil(o.TrackingToken) {
		return true
	}

	return false
}

// SetTrackingToken gets a reference to the given string and assigns it to the TrackingToken field.
func (o *SearchResponse) SetTrackingToken(v string) {
	o.TrackingToken = &v
}

// GetSessionInfo returns the SessionInfo field value if set, zero value otherwise.
func (o *SearchResponse) GetSessionInfo() SessionInfo {
	if o == nil || IsNil(o.SessionInfo) {
		var ret SessionInfo
		return ret
	}
	return *o.SessionInfo
}

// GetSessionInfoOk returns a tuple with the SessionInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResponse) GetSessionInfoOk() (*SessionInfo, bool) {
	if o == nil || IsNil(o.SessionInfo) {
		return nil, false
	}
	return o.SessionInfo, true
}

// HasSessionInfo returns a boolean if a field has been set.
func (o *SearchResponse) HasSessionInfo() bool {
	if o != nil && !IsNil(o.SessionInfo) {
		return true
	}

	return false
}

// SetSessionInfo gets a reference to the given SessionInfo and assigns it to the SessionInfo field.
func (o *SearchResponse) SetSessionInfo(v SessionInfo) {
	o.SessionInfo = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *SearchResponse) GetResults() []SearchResult {
	if o == nil || IsNil(o.Results) {
		var ret []SearchResult
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResponse) GetResultsOk() ([]SearchResult, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *SearchResponse) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []SearchResult and assigns it to the Results field.
func (o *SearchResponse) SetResults(v []SearchResult) {
	o.Results = v
}

// GetStructuredResults returns the StructuredResults field value if set, zero value otherwise.
func (o *SearchResponse) GetStructuredResults() []StructuredResult {
	if o == nil || IsNil(o.StructuredResults) {
		var ret []StructuredResult
		return ret
	}
	return o.StructuredResults
}

// GetStructuredResultsOk returns a tuple with the StructuredResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResponse) GetStructuredResultsOk() ([]StructuredResult, bool) {
	if o == nil || IsNil(o.StructuredResults) {
		return nil, false
	}
	return o.StructuredResults, true
}

// HasStructuredResults returns a boolean if a field has been set.
func (o *SearchResponse) HasStructuredResults() bool {
	if o != nil && !IsNil(o.StructuredResults) {
		return true
	}

	return false
}

// SetStructuredResults gets a reference to the given []StructuredResult and assigns it to the StructuredResults field.
func (o *SearchResponse) SetStructuredResults(v []StructuredResult) {
	o.StructuredResults = v
}

// GetGeneratedQnaResult returns the GeneratedQnaResult field value if set, zero value otherwise.
func (o *SearchResponse) GetGeneratedQnaResult() GeneratedQna {
	if o == nil || IsNil(o.GeneratedQnaResult) {
		var ret GeneratedQna
		return ret
	}
	return *o.GeneratedQnaResult
}

// GetGeneratedQnaResultOk returns a tuple with the GeneratedQnaResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResponse) GetGeneratedQnaResultOk() (*GeneratedQna, bool) {
	if o == nil || IsNil(o.GeneratedQnaResult) {
		return nil, false
	}
	return o.GeneratedQnaResult, true
}

// HasGeneratedQnaResult returns a boolean if a field has been set.
func (o *SearchResponse) HasGeneratedQnaResult() bool {
	if o != nil && !IsNil(o.GeneratedQnaResult) {
		return true
	}

	return false
}

// SetGeneratedQnaResult gets a reference to the given GeneratedQna and assigns it to the GeneratedQnaResult field.
func (o *SearchResponse) SetGeneratedQnaResult(v GeneratedQna) {
	o.GeneratedQnaResult = &v
}

// GetDebugInfo returns the DebugInfo field value if set, zero value otherwise.
func (o *SearchResponse) GetDebugInfo() SearchDebugInfo {
	if o == nil || IsNil(o.DebugInfo) {
		var ret SearchDebugInfo
		return ret
	}
	return *o.DebugInfo
}

// GetDebugInfoOk returns a tuple with the DebugInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResponse) GetDebugInfoOk() (*SearchDebugInfo, bool) {
	if o == nil || IsNil(o.DebugInfo) {
		return nil, false
	}
	return o.DebugInfo, true
}

// HasDebugInfo returns a boolean if a field has been set.
func (o *SearchResponse) HasDebugInfo() bool {
	if o != nil && !IsNil(o.DebugInfo) {
		return true
	}

	return false
}

// SetDebugInfo gets a reference to the given SearchDebugInfo and assigns it to the DebugInfo field.
func (o *SearchResponse) SetDebugInfo(v SearchDebugInfo) {
	o.DebugInfo = &v
}

// GetErrorInfo returns the ErrorInfo field value if set, zero value otherwise.
func (o *SearchResponse) GetErrorInfo() ErrorInfo {
	if o == nil || IsNil(o.ErrorInfo) {
		var ret ErrorInfo
		return ret
	}
	return *o.ErrorInfo
}

// GetErrorInfoOk returns a tuple with the ErrorInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResponse) GetErrorInfoOk() (*ErrorInfo, bool) {
	if o == nil || IsNil(o.ErrorInfo) {
		return nil, false
	}
	return o.ErrorInfo, true
}

// HasErrorInfo returns a boolean if a field has been set.
func (o *SearchResponse) HasErrorInfo() bool {
	if o != nil && !IsNil(o.ErrorInfo) {
		return true
	}

	return false
}

// SetErrorInfo gets a reference to the given ErrorInfo and assigns it to the ErrorInfo field.
func (o *SearchResponse) SetErrorInfo(v ErrorInfo) {
	o.ErrorInfo = &v
}

// GetRequestID returns the RequestID field value if set, zero value otherwise.
func (o *SearchResponse) GetRequestID() string {
	if o == nil || IsNil(o.RequestID) {
		var ret string
		return ret
	}
	return *o.RequestID
}

// GetRequestIDOk returns a tuple with the RequestID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResponse) GetRequestIDOk() (*string, bool) {
	if o == nil || IsNil(o.RequestID) {
		return nil, false
	}
	return o.RequestID, true
}

// HasRequestID returns a boolean if a field has been set.
func (o *SearchResponse) HasRequestID() bool {
	if o != nil && !IsNil(o.RequestID) {
		return true
	}

	return false
}

// SetRequestID gets a reference to the given string and assigns it to the RequestID field.
func (o *SearchResponse) SetRequestID(v string) {
	o.RequestID = &v
}

// GetBackendTimeMillis returns the BackendTimeMillis field value if set, zero value otherwise.
func (o *SearchResponse) GetBackendTimeMillis() int64 {
	if o == nil || IsNil(o.BackendTimeMillis) {
		var ret int64
		return ret
	}
	return *o.BackendTimeMillis
}

// GetBackendTimeMillisOk returns a tuple with the BackendTimeMillis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResponse) GetBackendTimeMillisOk() (*int64, bool) {
	if o == nil || IsNil(o.BackendTimeMillis) {
		return nil, false
	}
	return o.BackendTimeMillis, true
}

// HasBackendTimeMillis returns a boolean if a field has been set.
func (o *SearchResponse) HasBackendTimeMillis() bool {
	if o != nil && !IsNil(o.BackendTimeMillis) {
		return true
	}

	return false
}

// SetBackendTimeMillis gets a reference to the given int64 and assigns it to the BackendTimeMillis field.
func (o *SearchResponse) SetBackendTimeMillis(v int64) {
	o.BackendTimeMillis = &v
}

// GetExperimentIds returns the ExperimentIds field value if set, zero value otherwise.
func (o *SearchResponse) GetExperimentIds() []int64 {
	if o == nil || IsNil(o.ExperimentIds) {
		var ret []int64
		return ret
	}
	return o.ExperimentIds
}

// GetExperimentIdsOk returns a tuple with the ExperimentIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResponse) GetExperimentIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.ExperimentIds) {
		return nil, false
	}
	return o.ExperimentIds, true
}

// HasExperimentIds returns a boolean if a field has been set.
func (o *SearchResponse) HasExperimentIds() bool {
	if o != nil && !IsNil(o.ExperimentIds) {
		return true
	}

	return false
}

// SetExperimentIds gets a reference to the given []int64 and assigns it to the ExperimentIds field.
func (o *SearchResponse) SetExperimentIds(v []int64) {
	o.ExperimentIds = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *SearchResponse) GetMetadata() SearchResponseMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret SearchResponseMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResponse) GetMetadataOk() (*SearchResponseMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *SearchResponse) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given SearchResponseMetadata and assigns it to the Metadata field.
func (o *SearchResponse) SetMetadata(v SearchResponseMetadata) {
	o.Metadata = &v
}

// GetFacetResults returns the FacetResults field value if set, zero value otherwise.
func (o *SearchResponse) GetFacetResults() []FacetResult {
	if o == nil || IsNil(o.FacetResults) {
		var ret []FacetResult
		return ret
	}
	return o.FacetResults
}

// GetFacetResultsOk returns a tuple with the FacetResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResponse) GetFacetResultsOk() ([]FacetResult, bool) {
	if o == nil || IsNil(o.FacetResults) {
		return nil, false
	}
	return o.FacetResults, true
}

// HasFacetResults returns a boolean if a field has been set.
func (o *SearchResponse) HasFacetResults() bool {
	if o != nil && !IsNil(o.FacetResults) {
		return true
	}

	return false
}

// SetFacetResults gets a reference to the given []FacetResult and assigns it to the FacetResults field.
func (o *SearchResponse) SetFacetResults(v []FacetResult) {
	o.FacetResults = v
}

// GetResultTabs returns the ResultTabs field value if set, zero value otherwise.
func (o *SearchResponse) GetResultTabs() []ResultTab {
	if o == nil || IsNil(o.ResultTabs) {
		var ret []ResultTab
		return ret
	}
	return o.ResultTabs
}

// GetResultTabsOk returns a tuple with the ResultTabs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResponse) GetResultTabsOk() ([]ResultTab, bool) {
	if o == nil || IsNil(o.ResultTabs) {
		return nil, false
	}
	return o.ResultTabs, true
}

// HasResultTabs returns a boolean if a field has been set.
func (o *SearchResponse) HasResultTabs() bool {
	if o != nil && !IsNil(o.ResultTabs) {
		return true
	}

	return false
}

// SetResultTabs gets a reference to the given []ResultTab and assigns it to the ResultTabs field.
func (o *SearchResponse) SetResultTabs(v []ResultTab) {
	o.ResultTabs = v
}

// GetResultTabIds returns the ResultTabIds field value if set, zero value otherwise.
func (o *SearchResponse) GetResultTabIds() []string {
	if o == nil || IsNil(o.ResultTabIds) {
		var ret []string
		return ret
	}
	return o.ResultTabIds
}

// GetResultTabIdsOk returns a tuple with the ResultTabIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResponse) GetResultTabIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ResultTabIds) {
		return nil, false
	}
	return o.ResultTabIds, true
}

// HasResultTabIds returns a boolean if a field has been set.
func (o *SearchResponse) HasResultTabIds() bool {
	if o != nil && !IsNil(o.ResultTabIds) {
		return true
	}

	return false
}

// SetResultTabIds gets a reference to the given []string and assigns it to the ResultTabIds field.
func (o *SearchResponse) SetResultTabIds(v []string) {
	o.ResultTabIds = v
}

// GetResultsDescription returns the ResultsDescription field value if set, zero value otherwise.
func (o *SearchResponse) GetResultsDescription() ResultsDescription {
	if o == nil || IsNil(o.ResultsDescription) {
		var ret ResultsDescription
		return ret
	}
	return *o.ResultsDescription
}

// GetResultsDescriptionOk returns a tuple with the ResultsDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResponse) GetResultsDescriptionOk() (*ResultsDescription, bool) {
	if o == nil || IsNil(o.ResultsDescription) {
		return nil, false
	}
	return o.ResultsDescription, true
}

// HasResultsDescription returns a boolean if a field has been set.
func (o *SearchResponse) HasResultsDescription() bool {
	if o != nil && !IsNil(o.ResultsDescription) {
		return true
	}

	return false
}

// SetResultsDescription gets a reference to the given ResultsDescription and assigns it to the ResultsDescription field.
func (o *SearchResponse) SetResultsDescription(v ResultsDescription) {
	o.ResultsDescription = &v
}

// GetRewrittenFacetFilters returns the RewrittenFacetFilters field value if set, zero value otherwise.
func (o *SearchResponse) GetRewrittenFacetFilters() []FacetFilter {
	if o == nil || IsNil(o.RewrittenFacetFilters) {
		var ret []FacetFilter
		return ret
	}
	return o.RewrittenFacetFilters
}

// GetRewrittenFacetFiltersOk returns a tuple with the RewrittenFacetFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResponse) GetRewrittenFacetFiltersOk() ([]FacetFilter, bool) {
	if o == nil || IsNil(o.RewrittenFacetFilters) {
		return nil, false
	}
	return o.RewrittenFacetFilters, true
}

// HasRewrittenFacetFilters returns a boolean if a field has been set.
func (o *SearchResponse) HasRewrittenFacetFilters() bool {
	if o != nil && !IsNil(o.RewrittenFacetFilters) {
		return true
	}

	return false
}

// SetRewrittenFacetFilters gets a reference to the given []FacetFilter and assigns it to the RewrittenFacetFilters field.
func (o *SearchResponse) SetRewrittenFacetFilters(v []FacetFilter) {
	o.RewrittenFacetFilters = v
}

// GetCursor returns the Cursor field value if set, zero value otherwise.
func (o *SearchResponse) GetCursor() string {
	if o == nil || IsNil(o.Cursor) {
		var ret string
		return ret
	}
	return *o.Cursor
}

// GetCursorOk returns a tuple with the Cursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResponse) GetCursorOk() (*string, bool) {
	if o == nil || IsNil(o.Cursor) {
		return nil, false
	}
	return o.Cursor, true
}

// HasCursor returns a boolean if a field has been set.
func (o *SearchResponse) HasCursor() bool {
	if o != nil && !IsNil(o.Cursor) {
		return true
	}

	return false
}

// SetCursor gets a reference to the given string and assigns it to the Cursor field.
func (o *SearchResponse) SetCursor(v string) {
	o.Cursor = &v
}

// GetHasMoreResults returns the HasMoreResults field value if set, zero value otherwise.
func (o *SearchResponse) GetHasMoreResults() bool {
	if o == nil || IsNil(o.HasMoreResults) {
		var ret bool
		return ret
	}
	return *o.HasMoreResults
}

// GetHasMoreResultsOk returns a tuple with the HasMoreResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResponse) GetHasMoreResultsOk() (*bool, bool) {
	if o == nil || IsNil(o.HasMoreResults) {
		return nil, false
	}
	return o.HasMoreResults, true
}

// HasHasMoreResults returns a boolean if a field has been set.
func (o *SearchResponse) HasHasMoreResults() bool {
	if o != nil && !IsNil(o.HasMoreResults) {
		return true
	}

	return false
}

// SetHasMoreResults gets a reference to the given bool and assigns it to the HasMoreResults field.
func (o *SearchResponse) SetHasMoreResults(v bool) {
	o.HasMoreResults = &v
}

func (o SearchResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TrackingToken) {
		toSerialize["trackingToken"] = o.TrackingToken
	}
	if !IsNil(o.SessionInfo) {
		toSerialize["sessionInfo"] = o.SessionInfo
	}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	if !IsNil(o.StructuredResults) {
		toSerialize["structuredResults"] = o.StructuredResults
	}
	if !IsNil(o.GeneratedQnaResult) {
		toSerialize["generatedQnaResult"] = o.GeneratedQnaResult
	}
	if !IsNil(o.DebugInfo) {
		toSerialize["debugInfo"] = o.DebugInfo
	}
	if !IsNil(o.ErrorInfo) {
		toSerialize["errorInfo"] = o.ErrorInfo
	}
	if !IsNil(o.RequestID) {
		toSerialize["requestID"] = o.RequestID
	}
	if !IsNil(o.BackendTimeMillis) {
		toSerialize["backendTimeMillis"] = o.BackendTimeMillis
	}
	if !IsNil(o.ExperimentIds) {
		toSerialize["experimentIds"] = o.ExperimentIds
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.FacetResults) {
		toSerialize["facetResults"] = o.FacetResults
	}
	if !IsNil(o.ResultTabs) {
		toSerialize["resultTabs"] = o.ResultTabs
	}
	if !IsNil(o.ResultTabIds) {
		toSerialize["resultTabIds"] = o.ResultTabIds
	}
	if !IsNil(o.ResultsDescription) {
		toSerialize["resultsDescription"] = o.ResultsDescription
	}
	if !IsNil(o.RewrittenFacetFilters) {
		toSerialize["rewrittenFacetFilters"] = o.RewrittenFacetFilters
	}
	if !IsNil(o.Cursor) {
		toSerialize["cursor"] = o.Cursor
	}
	if !IsNil(o.HasMoreResults) {
		toSerialize["hasMoreResults"] = o.HasMoreResults
	}
	return toSerialize, nil
}

type NullableSearchResponse struct {
	value *SearchResponse
	isSet bool
}

func (v NullableSearchResponse) Get() *SearchResponse {
	return v.value
}

func (v *NullableSearchResponse) Set(val *SearchResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchResponse(val *SearchResponse) *NullableSearchResponse {
	return &NullableSearchResponse{value: val, isSet: true}
}

func (v NullableSearchResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


