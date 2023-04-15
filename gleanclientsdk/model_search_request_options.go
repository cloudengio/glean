/*
Glean Client API - Platform Preview

# Introduction These are all the APIs used by Glean to implement the Glean client. These are available as platform preview for implementing a custom client to the Glean system.  # Usage guidelines A subset of these endpoints are also in the developer ready section, which is available for public use. The rest of the endpoints are subject to prior agreement with Glean before usage. Please contact support@glean.com if you would like to use an API that is not currently available in the developer ready section. 

API version: 0.9.0
Contact: support@glean.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gleanclientsdk

import (
	"encoding/json"
)

// checks if the SearchRequestOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchRequestOptions{}

// SearchRequestOptions struct for SearchRequestOptions
type SearchRequestOptions struct {
	DebugOptions *SearchDebugOptions `json:"debugOptions,omitempty"`
	// Filter results to a single datasource name (e.g. gmail, slack). All results are returned if missing.
	DatasourceFilter *string `json:"datasourceFilter,omitempty"`
	// Filter results to one or more datasources (e.g. gmail, slack). All results are returned if missing.
	DatasourcesFilter []string `json:"datasourcesFilter,omitempty"`
	// If true, the operators in the query are taken to override any operators in facetFilters in the case of conflict. This is used to correctly set rewrittenFacetFilters and rewrittenQuery.
	QueryOverridesFacetFilters *bool `json:"queryOverridesFacetFilters,omitempty"`
	// A list of filters for the query. An AND is assumed between different facetFilters. For example, owner Sumeet and doc_type Spreadsheet shows documents that are by Sumeet AND are Spreadsheets.
	FacetFilters []FacetFilter `json:"facetFilters,omitempty"`
	FacetBucketFilter *FacetBucketFilter `json:"facetBucketFilter,omitempty"`
	// The maximum number of FacetBuckets to return in each FacetResult.
	FacetBucketSize int32 `json:"facetBucketSize"`
	// Auth tokens which may be used for non-indexed, federated results (e.g. Gmail).
	AuthTokens []AuthToken `json:"authTokens,omitempty"`
	// Hints that the QE should return result counts (via the datasource facet result) for all supported datasources, rather than just those specified in the datasource[s]Filter
	FetchAllDatasourceCounts *bool `json:"fetchAllDatasourceCounts,omitempty"`
	// Array of hints containing what QE should return back to FE.
	ResponseHints []string `json:"responseHints,omitempty"`
	// The offset of the client's timezone in minutes from UTC. e.g. PDT is -420 because it's 7 hours behind UTC.
	TimezoneOffset *int32 `json:"timezoneOffset,omitempty"`
	// Whether or not to force not ignoring of negation, i.e. force negated terms to be negated.
	ForceNegation *bool `json:"forceNegation,omitempty"`
	// Whether or not to disable spellcheck.
	DisableSpellcheck *bool `json:"disableSpellcheck,omitempty"`
	// Disables automatic adjustment of the input query for spelling corrections or other reasons.
	DisableQueryAutocorrect *bool `json:"disableQueryAutocorrect,omitempty"`
	// The number of characters to include in the expanded snippet.
	ExpandedSnippetSize *int32 `json:"expandedSnippetSize,omitempty"`
}

// NewSearchRequestOptions instantiates a new SearchRequestOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchRequestOptions(facetBucketSize int32) *SearchRequestOptions {
	this := SearchRequestOptions{}
	this.FacetBucketSize = facetBucketSize
	return &this
}

// NewSearchRequestOptionsWithDefaults instantiates a new SearchRequestOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchRequestOptionsWithDefaults() *SearchRequestOptions {
	this := SearchRequestOptions{}
	return &this
}

// GetDebugOptions returns the DebugOptions field value if set, zero value otherwise.
func (o *SearchRequestOptions) GetDebugOptions() SearchDebugOptions {
	if o == nil || IsNil(o.DebugOptions) {
		var ret SearchDebugOptions
		return ret
	}
	return *o.DebugOptions
}

// GetDebugOptionsOk returns a tuple with the DebugOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchRequestOptions) GetDebugOptionsOk() (*SearchDebugOptions, bool) {
	if o == nil || IsNil(o.DebugOptions) {
		return nil, false
	}
	return o.DebugOptions, true
}

// HasDebugOptions returns a boolean if a field has been set.
func (o *SearchRequestOptions) HasDebugOptions() bool {
	if o != nil && !IsNil(o.DebugOptions) {
		return true
	}

	return false
}

// SetDebugOptions gets a reference to the given SearchDebugOptions and assigns it to the DebugOptions field.
func (o *SearchRequestOptions) SetDebugOptions(v SearchDebugOptions) {
	o.DebugOptions = &v
}

// GetDatasourceFilter returns the DatasourceFilter field value if set, zero value otherwise.
func (o *SearchRequestOptions) GetDatasourceFilter() string {
	if o == nil || IsNil(o.DatasourceFilter) {
		var ret string
		return ret
	}
	return *o.DatasourceFilter
}

// GetDatasourceFilterOk returns a tuple with the DatasourceFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchRequestOptions) GetDatasourceFilterOk() (*string, bool) {
	if o == nil || IsNil(o.DatasourceFilter) {
		return nil, false
	}
	return o.DatasourceFilter, true
}

// HasDatasourceFilter returns a boolean if a field has been set.
func (o *SearchRequestOptions) HasDatasourceFilter() bool {
	if o != nil && !IsNil(o.DatasourceFilter) {
		return true
	}

	return false
}

// SetDatasourceFilter gets a reference to the given string and assigns it to the DatasourceFilter field.
func (o *SearchRequestOptions) SetDatasourceFilter(v string) {
	o.DatasourceFilter = &v
}

// GetDatasourcesFilter returns the DatasourcesFilter field value if set, zero value otherwise.
func (o *SearchRequestOptions) GetDatasourcesFilter() []string {
	if o == nil || IsNil(o.DatasourcesFilter) {
		var ret []string
		return ret
	}
	return o.DatasourcesFilter
}

// GetDatasourcesFilterOk returns a tuple with the DatasourcesFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchRequestOptions) GetDatasourcesFilterOk() ([]string, bool) {
	if o == nil || IsNil(o.DatasourcesFilter) {
		return nil, false
	}
	return o.DatasourcesFilter, true
}

// HasDatasourcesFilter returns a boolean if a field has been set.
func (o *SearchRequestOptions) HasDatasourcesFilter() bool {
	if o != nil && !IsNil(o.DatasourcesFilter) {
		return true
	}

	return false
}

// SetDatasourcesFilter gets a reference to the given []string and assigns it to the DatasourcesFilter field.
func (o *SearchRequestOptions) SetDatasourcesFilter(v []string) {
	o.DatasourcesFilter = v
}

// GetQueryOverridesFacetFilters returns the QueryOverridesFacetFilters field value if set, zero value otherwise.
func (o *SearchRequestOptions) GetQueryOverridesFacetFilters() bool {
	if o == nil || IsNil(o.QueryOverridesFacetFilters) {
		var ret bool
		return ret
	}
	return *o.QueryOverridesFacetFilters
}

// GetQueryOverridesFacetFiltersOk returns a tuple with the QueryOverridesFacetFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchRequestOptions) GetQueryOverridesFacetFiltersOk() (*bool, bool) {
	if o == nil || IsNil(o.QueryOverridesFacetFilters) {
		return nil, false
	}
	return o.QueryOverridesFacetFilters, true
}

// HasQueryOverridesFacetFilters returns a boolean if a field has been set.
func (o *SearchRequestOptions) HasQueryOverridesFacetFilters() bool {
	if o != nil && !IsNil(o.QueryOverridesFacetFilters) {
		return true
	}

	return false
}

// SetQueryOverridesFacetFilters gets a reference to the given bool and assigns it to the QueryOverridesFacetFilters field.
func (o *SearchRequestOptions) SetQueryOverridesFacetFilters(v bool) {
	o.QueryOverridesFacetFilters = &v
}

// GetFacetFilters returns the FacetFilters field value if set, zero value otherwise.
func (o *SearchRequestOptions) GetFacetFilters() []FacetFilter {
	if o == nil || IsNil(o.FacetFilters) {
		var ret []FacetFilter
		return ret
	}
	return o.FacetFilters
}

// GetFacetFiltersOk returns a tuple with the FacetFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchRequestOptions) GetFacetFiltersOk() ([]FacetFilter, bool) {
	if o == nil || IsNil(o.FacetFilters) {
		return nil, false
	}
	return o.FacetFilters, true
}

// HasFacetFilters returns a boolean if a field has been set.
func (o *SearchRequestOptions) HasFacetFilters() bool {
	if o != nil && !IsNil(o.FacetFilters) {
		return true
	}

	return false
}

// SetFacetFilters gets a reference to the given []FacetFilter and assigns it to the FacetFilters field.
func (o *SearchRequestOptions) SetFacetFilters(v []FacetFilter) {
	o.FacetFilters = v
}

// GetFacetBucketFilter returns the FacetBucketFilter field value if set, zero value otherwise.
func (o *SearchRequestOptions) GetFacetBucketFilter() FacetBucketFilter {
	if o == nil || IsNil(o.FacetBucketFilter) {
		var ret FacetBucketFilter
		return ret
	}
	return *o.FacetBucketFilter
}

// GetFacetBucketFilterOk returns a tuple with the FacetBucketFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchRequestOptions) GetFacetBucketFilterOk() (*FacetBucketFilter, bool) {
	if o == nil || IsNil(o.FacetBucketFilter) {
		return nil, false
	}
	return o.FacetBucketFilter, true
}

// HasFacetBucketFilter returns a boolean if a field has been set.
func (o *SearchRequestOptions) HasFacetBucketFilter() bool {
	if o != nil && !IsNil(o.FacetBucketFilter) {
		return true
	}

	return false
}

// SetFacetBucketFilter gets a reference to the given FacetBucketFilter and assigns it to the FacetBucketFilter field.
func (o *SearchRequestOptions) SetFacetBucketFilter(v FacetBucketFilter) {
	o.FacetBucketFilter = &v
}

// GetFacetBucketSize returns the FacetBucketSize field value
func (o *SearchRequestOptions) GetFacetBucketSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FacetBucketSize
}

// GetFacetBucketSizeOk returns a tuple with the FacetBucketSize field value
// and a boolean to check if the value has been set.
func (o *SearchRequestOptions) GetFacetBucketSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FacetBucketSize, true
}

// SetFacetBucketSize sets field value
func (o *SearchRequestOptions) SetFacetBucketSize(v int32) {
	o.FacetBucketSize = v
}

// GetAuthTokens returns the AuthTokens field value if set, zero value otherwise.
func (o *SearchRequestOptions) GetAuthTokens() []AuthToken {
	if o == nil || IsNil(o.AuthTokens) {
		var ret []AuthToken
		return ret
	}
	return o.AuthTokens
}

// GetAuthTokensOk returns a tuple with the AuthTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchRequestOptions) GetAuthTokensOk() ([]AuthToken, bool) {
	if o == nil || IsNil(o.AuthTokens) {
		return nil, false
	}
	return o.AuthTokens, true
}

// HasAuthTokens returns a boolean if a field has been set.
func (o *SearchRequestOptions) HasAuthTokens() bool {
	if o != nil && !IsNil(o.AuthTokens) {
		return true
	}

	return false
}

// SetAuthTokens gets a reference to the given []AuthToken and assigns it to the AuthTokens field.
func (o *SearchRequestOptions) SetAuthTokens(v []AuthToken) {
	o.AuthTokens = v
}

// GetFetchAllDatasourceCounts returns the FetchAllDatasourceCounts field value if set, zero value otherwise.
func (o *SearchRequestOptions) GetFetchAllDatasourceCounts() bool {
	if o == nil || IsNil(o.FetchAllDatasourceCounts) {
		var ret bool
		return ret
	}
	return *o.FetchAllDatasourceCounts
}

// GetFetchAllDatasourceCountsOk returns a tuple with the FetchAllDatasourceCounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchRequestOptions) GetFetchAllDatasourceCountsOk() (*bool, bool) {
	if o == nil || IsNil(o.FetchAllDatasourceCounts) {
		return nil, false
	}
	return o.FetchAllDatasourceCounts, true
}

// HasFetchAllDatasourceCounts returns a boolean if a field has been set.
func (o *SearchRequestOptions) HasFetchAllDatasourceCounts() bool {
	if o != nil && !IsNil(o.FetchAllDatasourceCounts) {
		return true
	}

	return false
}

// SetFetchAllDatasourceCounts gets a reference to the given bool and assigns it to the FetchAllDatasourceCounts field.
func (o *SearchRequestOptions) SetFetchAllDatasourceCounts(v bool) {
	o.FetchAllDatasourceCounts = &v
}

// GetResponseHints returns the ResponseHints field value if set, zero value otherwise.
func (o *SearchRequestOptions) GetResponseHints() []string {
	if o == nil || IsNil(o.ResponseHints) {
		var ret []string
		return ret
	}
	return o.ResponseHints
}

// GetResponseHintsOk returns a tuple with the ResponseHints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchRequestOptions) GetResponseHintsOk() ([]string, bool) {
	if o == nil || IsNil(o.ResponseHints) {
		return nil, false
	}
	return o.ResponseHints, true
}

// HasResponseHints returns a boolean if a field has been set.
func (o *SearchRequestOptions) HasResponseHints() bool {
	if o != nil && !IsNil(o.ResponseHints) {
		return true
	}

	return false
}

// SetResponseHints gets a reference to the given []string and assigns it to the ResponseHints field.
func (o *SearchRequestOptions) SetResponseHints(v []string) {
	o.ResponseHints = v
}

// GetTimezoneOffset returns the TimezoneOffset field value if set, zero value otherwise.
func (o *SearchRequestOptions) GetTimezoneOffset() int32 {
	if o == nil || IsNil(o.TimezoneOffset) {
		var ret int32
		return ret
	}
	return *o.TimezoneOffset
}

// GetTimezoneOffsetOk returns a tuple with the TimezoneOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchRequestOptions) GetTimezoneOffsetOk() (*int32, bool) {
	if o == nil || IsNil(o.TimezoneOffset) {
		return nil, false
	}
	return o.TimezoneOffset, true
}

// HasTimezoneOffset returns a boolean if a field has been set.
func (o *SearchRequestOptions) HasTimezoneOffset() bool {
	if o != nil && !IsNil(o.TimezoneOffset) {
		return true
	}

	return false
}

// SetTimezoneOffset gets a reference to the given int32 and assigns it to the TimezoneOffset field.
func (o *SearchRequestOptions) SetTimezoneOffset(v int32) {
	o.TimezoneOffset = &v
}

// GetForceNegation returns the ForceNegation field value if set, zero value otherwise.
func (o *SearchRequestOptions) GetForceNegation() bool {
	if o == nil || IsNil(o.ForceNegation) {
		var ret bool
		return ret
	}
	return *o.ForceNegation
}

// GetForceNegationOk returns a tuple with the ForceNegation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchRequestOptions) GetForceNegationOk() (*bool, bool) {
	if o == nil || IsNil(o.ForceNegation) {
		return nil, false
	}
	return o.ForceNegation, true
}

// HasForceNegation returns a boolean if a field has been set.
func (o *SearchRequestOptions) HasForceNegation() bool {
	if o != nil && !IsNil(o.ForceNegation) {
		return true
	}

	return false
}

// SetForceNegation gets a reference to the given bool and assigns it to the ForceNegation field.
func (o *SearchRequestOptions) SetForceNegation(v bool) {
	o.ForceNegation = &v
}

// GetDisableSpellcheck returns the DisableSpellcheck field value if set, zero value otherwise.
func (o *SearchRequestOptions) GetDisableSpellcheck() bool {
	if o == nil || IsNil(o.DisableSpellcheck) {
		var ret bool
		return ret
	}
	return *o.DisableSpellcheck
}

// GetDisableSpellcheckOk returns a tuple with the DisableSpellcheck field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchRequestOptions) GetDisableSpellcheckOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableSpellcheck) {
		return nil, false
	}
	return o.DisableSpellcheck, true
}

// HasDisableSpellcheck returns a boolean if a field has been set.
func (o *SearchRequestOptions) HasDisableSpellcheck() bool {
	if o != nil && !IsNil(o.DisableSpellcheck) {
		return true
	}

	return false
}

// SetDisableSpellcheck gets a reference to the given bool and assigns it to the DisableSpellcheck field.
func (o *SearchRequestOptions) SetDisableSpellcheck(v bool) {
	o.DisableSpellcheck = &v
}

// GetDisableQueryAutocorrect returns the DisableQueryAutocorrect field value if set, zero value otherwise.
func (o *SearchRequestOptions) GetDisableQueryAutocorrect() bool {
	if o == nil || IsNil(o.DisableQueryAutocorrect) {
		var ret bool
		return ret
	}
	return *o.DisableQueryAutocorrect
}

// GetDisableQueryAutocorrectOk returns a tuple with the DisableQueryAutocorrect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchRequestOptions) GetDisableQueryAutocorrectOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableQueryAutocorrect) {
		return nil, false
	}
	return o.DisableQueryAutocorrect, true
}

// HasDisableQueryAutocorrect returns a boolean if a field has been set.
func (o *SearchRequestOptions) HasDisableQueryAutocorrect() bool {
	if o != nil && !IsNil(o.DisableQueryAutocorrect) {
		return true
	}

	return false
}

// SetDisableQueryAutocorrect gets a reference to the given bool and assigns it to the DisableQueryAutocorrect field.
func (o *SearchRequestOptions) SetDisableQueryAutocorrect(v bool) {
	o.DisableQueryAutocorrect = &v
}

// GetExpandedSnippetSize returns the ExpandedSnippetSize field value if set, zero value otherwise.
func (o *SearchRequestOptions) GetExpandedSnippetSize() int32 {
	if o == nil || IsNil(o.ExpandedSnippetSize) {
		var ret int32
		return ret
	}
	return *o.ExpandedSnippetSize
}

// GetExpandedSnippetSizeOk returns a tuple with the ExpandedSnippetSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchRequestOptions) GetExpandedSnippetSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.ExpandedSnippetSize) {
		return nil, false
	}
	return o.ExpandedSnippetSize, true
}

// HasExpandedSnippetSize returns a boolean if a field has been set.
func (o *SearchRequestOptions) HasExpandedSnippetSize() bool {
	if o != nil && !IsNil(o.ExpandedSnippetSize) {
		return true
	}

	return false
}

// SetExpandedSnippetSize gets a reference to the given int32 and assigns it to the ExpandedSnippetSize field.
func (o *SearchRequestOptions) SetExpandedSnippetSize(v int32) {
	o.ExpandedSnippetSize = &v
}

func (o SearchRequestOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchRequestOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DebugOptions) {
		toSerialize["debugOptions"] = o.DebugOptions
	}
	if !IsNil(o.DatasourceFilter) {
		toSerialize["datasourceFilter"] = o.DatasourceFilter
	}
	if !IsNil(o.DatasourcesFilter) {
		toSerialize["datasourcesFilter"] = o.DatasourcesFilter
	}
	if !IsNil(o.QueryOverridesFacetFilters) {
		toSerialize["queryOverridesFacetFilters"] = o.QueryOverridesFacetFilters
	}
	if !IsNil(o.FacetFilters) {
		toSerialize["facetFilters"] = o.FacetFilters
	}
	if !IsNil(o.FacetBucketFilter) {
		toSerialize["facetBucketFilter"] = o.FacetBucketFilter
	}
	toSerialize["facetBucketSize"] = o.FacetBucketSize
	if !IsNil(o.AuthTokens) {
		toSerialize["authTokens"] = o.AuthTokens
	}
	if !IsNil(o.FetchAllDatasourceCounts) {
		toSerialize["fetchAllDatasourceCounts"] = o.FetchAllDatasourceCounts
	}
	if !IsNil(o.ResponseHints) {
		toSerialize["responseHints"] = o.ResponseHints
	}
	if !IsNil(o.TimezoneOffset) {
		toSerialize["timezoneOffset"] = o.TimezoneOffset
	}
	if !IsNil(o.ForceNegation) {
		toSerialize["forceNegation"] = o.ForceNegation
	}
	if !IsNil(o.DisableSpellcheck) {
		toSerialize["disableSpellcheck"] = o.DisableSpellcheck
	}
	if !IsNil(o.DisableQueryAutocorrect) {
		toSerialize["disableQueryAutocorrect"] = o.DisableQueryAutocorrect
	}
	if !IsNil(o.ExpandedSnippetSize) {
		toSerialize["expandedSnippetSize"] = o.ExpandedSnippetSize
	}
	return toSerialize, nil
}

type NullableSearchRequestOptions struct {
	value *SearchRequestOptions
	isSet bool
}

func (v NullableSearchRequestOptions) Get() *SearchRequestOptions {
	return v.value
}

func (v *NullableSearchRequestOptions) Set(val *SearchRequestOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchRequestOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchRequestOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchRequestOptions(val *SearchRequestOptions) *NullableSearchRequestOptions {
	return &NullableSearchRequestOptions{value: val, isSet: true}
}

func (v NullableSearchRequestOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchRequestOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


