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

// checks if the RecommendationsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecommendationsResponse{}

// RecommendationsResponse struct for RecommendationsResponse
type RecommendationsResponse struct {
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
}

// NewRecommendationsResponse instantiates a new RecommendationsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecommendationsResponse() *RecommendationsResponse {
	this := RecommendationsResponse{}
	return &this
}

// NewRecommendationsResponseWithDefaults instantiates a new RecommendationsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecommendationsResponseWithDefaults() *RecommendationsResponse {
	this := RecommendationsResponse{}
	return &this
}

// GetTrackingToken returns the TrackingToken field value if set, zero value otherwise.
func (o *RecommendationsResponse) GetTrackingToken() string {
	if o == nil || IsNil(o.TrackingToken) {
		var ret string
		return ret
	}
	return *o.TrackingToken
}

// GetTrackingTokenOk returns a tuple with the TrackingToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationsResponse) GetTrackingTokenOk() (*string, bool) {
	if o == nil || IsNil(o.TrackingToken) {
		return nil, false
	}
	return o.TrackingToken, true
}

// HasTrackingToken returns a boolean if a field has been set.
func (o *RecommendationsResponse) HasTrackingToken() bool {
	if o != nil && !IsNil(o.TrackingToken) {
		return true
	}

	return false
}

// SetTrackingToken gets a reference to the given string and assigns it to the TrackingToken field.
func (o *RecommendationsResponse) SetTrackingToken(v string) {
	o.TrackingToken = &v
}

// GetSessionInfo returns the SessionInfo field value if set, zero value otherwise.
func (o *RecommendationsResponse) GetSessionInfo() SessionInfo {
	if o == nil || IsNil(o.SessionInfo) {
		var ret SessionInfo
		return ret
	}
	return *o.SessionInfo
}

// GetSessionInfoOk returns a tuple with the SessionInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationsResponse) GetSessionInfoOk() (*SessionInfo, bool) {
	if o == nil || IsNil(o.SessionInfo) {
		return nil, false
	}
	return o.SessionInfo, true
}

// HasSessionInfo returns a boolean if a field has been set.
func (o *RecommendationsResponse) HasSessionInfo() bool {
	if o != nil && !IsNil(o.SessionInfo) {
		return true
	}

	return false
}

// SetSessionInfo gets a reference to the given SessionInfo and assigns it to the SessionInfo field.
func (o *RecommendationsResponse) SetSessionInfo(v SessionInfo) {
	o.SessionInfo = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *RecommendationsResponse) GetResults() []SearchResult {
	if o == nil || IsNil(o.Results) {
		var ret []SearchResult
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationsResponse) GetResultsOk() ([]SearchResult, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *RecommendationsResponse) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []SearchResult and assigns it to the Results field.
func (o *RecommendationsResponse) SetResults(v []SearchResult) {
	o.Results = v
}

// GetStructuredResults returns the StructuredResults field value if set, zero value otherwise.
func (o *RecommendationsResponse) GetStructuredResults() []StructuredResult {
	if o == nil || IsNil(o.StructuredResults) {
		var ret []StructuredResult
		return ret
	}
	return o.StructuredResults
}

// GetStructuredResultsOk returns a tuple with the StructuredResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationsResponse) GetStructuredResultsOk() ([]StructuredResult, bool) {
	if o == nil || IsNil(o.StructuredResults) {
		return nil, false
	}
	return o.StructuredResults, true
}

// HasStructuredResults returns a boolean if a field has been set.
func (o *RecommendationsResponse) HasStructuredResults() bool {
	if o != nil && !IsNil(o.StructuredResults) {
		return true
	}

	return false
}

// SetStructuredResults gets a reference to the given []StructuredResult and assigns it to the StructuredResults field.
func (o *RecommendationsResponse) SetStructuredResults(v []StructuredResult) {
	o.StructuredResults = v
}

// GetGeneratedQnaResult returns the GeneratedQnaResult field value if set, zero value otherwise.
func (o *RecommendationsResponse) GetGeneratedQnaResult() GeneratedQna {
	if o == nil || IsNil(o.GeneratedQnaResult) {
		var ret GeneratedQna
		return ret
	}
	return *o.GeneratedQnaResult
}

// GetGeneratedQnaResultOk returns a tuple with the GeneratedQnaResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationsResponse) GetGeneratedQnaResultOk() (*GeneratedQna, bool) {
	if o == nil || IsNil(o.GeneratedQnaResult) {
		return nil, false
	}
	return o.GeneratedQnaResult, true
}

// HasGeneratedQnaResult returns a boolean if a field has been set.
func (o *RecommendationsResponse) HasGeneratedQnaResult() bool {
	if o != nil && !IsNil(o.GeneratedQnaResult) {
		return true
	}

	return false
}

// SetGeneratedQnaResult gets a reference to the given GeneratedQna and assigns it to the GeneratedQnaResult field.
func (o *RecommendationsResponse) SetGeneratedQnaResult(v GeneratedQna) {
	o.GeneratedQnaResult = &v
}

// GetDebugInfo returns the DebugInfo field value if set, zero value otherwise.
func (o *RecommendationsResponse) GetDebugInfo() SearchDebugInfo {
	if o == nil || IsNil(o.DebugInfo) {
		var ret SearchDebugInfo
		return ret
	}
	return *o.DebugInfo
}

// GetDebugInfoOk returns a tuple with the DebugInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationsResponse) GetDebugInfoOk() (*SearchDebugInfo, bool) {
	if o == nil || IsNil(o.DebugInfo) {
		return nil, false
	}
	return o.DebugInfo, true
}

// HasDebugInfo returns a boolean if a field has been set.
func (o *RecommendationsResponse) HasDebugInfo() bool {
	if o != nil && !IsNil(o.DebugInfo) {
		return true
	}

	return false
}

// SetDebugInfo gets a reference to the given SearchDebugInfo and assigns it to the DebugInfo field.
func (o *RecommendationsResponse) SetDebugInfo(v SearchDebugInfo) {
	o.DebugInfo = &v
}

// GetErrorInfo returns the ErrorInfo field value if set, zero value otherwise.
func (o *RecommendationsResponse) GetErrorInfo() ErrorInfo {
	if o == nil || IsNil(o.ErrorInfo) {
		var ret ErrorInfo
		return ret
	}
	return *o.ErrorInfo
}

// GetErrorInfoOk returns a tuple with the ErrorInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationsResponse) GetErrorInfoOk() (*ErrorInfo, bool) {
	if o == nil || IsNil(o.ErrorInfo) {
		return nil, false
	}
	return o.ErrorInfo, true
}

// HasErrorInfo returns a boolean if a field has been set.
func (o *RecommendationsResponse) HasErrorInfo() bool {
	if o != nil && !IsNil(o.ErrorInfo) {
		return true
	}

	return false
}

// SetErrorInfo gets a reference to the given ErrorInfo and assigns it to the ErrorInfo field.
func (o *RecommendationsResponse) SetErrorInfo(v ErrorInfo) {
	o.ErrorInfo = &v
}

// GetRequestID returns the RequestID field value if set, zero value otherwise.
func (o *RecommendationsResponse) GetRequestID() string {
	if o == nil || IsNil(o.RequestID) {
		var ret string
		return ret
	}
	return *o.RequestID
}

// GetRequestIDOk returns a tuple with the RequestID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationsResponse) GetRequestIDOk() (*string, bool) {
	if o == nil || IsNil(o.RequestID) {
		return nil, false
	}
	return o.RequestID, true
}

// HasRequestID returns a boolean if a field has been set.
func (o *RecommendationsResponse) HasRequestID() bool {
	if o != nil && !IsNil(o.RequestID) {
		return true
	}

	return false
}

// SetRequestID gets a reference to the given string and assigns it to the RequestID field.
func (o *RecommendationsResponse) SetRequestID(v string) {
	o.RequestID = &v
}

// GetBackendTimeMillis returns the BackendTimeMillis field value if set, zero value otherwise.
func (o *RecommendationsResponse) GetBackendTimeMillis() int64 {
	if o == nil || IsNil(o.BackendTimeMillis) {
		var ret int64
		return ret
	}
	return *o.BackendTimeMillis
}

// GetBackendTimeMillisOk returns a tuple with the BackendTimeMillis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationsResponse) GetBackendTimeMillisOk() (*int64, bool) {
	if o == nil || IsNil(o.BackendTimeMillis) {
		return nil, false
	}
	return o.BackendTimeMillis, true
}

// HasBackendTimeMillis returns a boolean if a field has been set.
func (o *RecommendationsResponse) HasBackendTimeMillis() bool {
	if o != nil && !IsNil(o.BackendTimeMillis) {
		return true
	}

	return false
}

// SetBackendTimeMillis gets a reference to the given int64 and assigns it to the BackendTimeMillis field.
func (o *RecommendationsResponse) SetBackendTimeMillis(v int64) {
	o.BackendTimeMillis = &v
}

func (o RecommendationsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecommendationsResponse) ToMap() (map[string]interface{}, error) {
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
	return toSerialize, nil
}

type NullableRecommendationsResponse struct {
	value *RecommendationsResponse
	isSet bool
}

func (v NullableRecommendationsResponse) Get() *RecommendationsResponse {
	return v.value
}

func (v *NullableRecommendationsResponse) Set(val *RecommendationsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRecommendationsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRecommendationsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecommendationsResponse(val *RecommendationsResponse) *NullableRecommendationsResponse {
	return &NullableRecommendationsResponse{value: val, isSet: true}
}

func (v NullableRecommendationsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecommendationsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


