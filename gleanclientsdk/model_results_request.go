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
	"time"
)

// checks if the ResultsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResultsRequest{}

// ResultsRequest struct for ResultsRequest
type ResultsRequest struct {
	// The ISO 8601 timestamp associated with the client request.
	Timestamp *time.Time `json:"timestamp,omitempty"`
	// A previously received trackingToken for a search associated with the same query. Useful for more requests and requests for other tabs.
	TrackingToken *string `json:"trackingToken,omitempty"`
	SessionInfo *SessionInfo `json:"sessionInfo,omitempty"`
	SourceInfo SearchRequestSourceInfo `json:"sourceInfo"`
	SourceDocument *Document `json:"sourceDocument,omitempty"`
	// Debug only search params to to change the flow of control in request handling. It will be passed around service boundaries/endpoints. For more details, https://docs.google.com/document/d/1e6taTfWUL8KNUC9de8kmmG2MG2L6cTx4ulOJfAshDTM/edit. Requires sufficient permissions.
	Sc *string `json:"sc,omitempty"`
	// Hint to the server about how many results to send back. Server may return less or more. Structured results and clustered results don't count towards pageSize.
	PageSize *int32 `json:"pageSize,omitempty"`
	// Hint to the server about how many characters long a snippet may be. Server may return less or more.
	MaxSnippetSize *int32 `json:"maxSnippetSize,omitempty"`
}

// NewResultsRequest instantiates a new ResultsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResultsRequest(sourceInfo SearchRequestSourceInfo) *ResultsRequest {
	this := ResultsRequest{}
	this.SourceInfo = sourceInfo
	return &this
}

// NewResultsRequestWithDefaults instantiates a new ResultsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResultsRequestWithDefaults() *ResultsRequest {
	this := ResultsRequest{}
	return &this
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *ResultsRequest) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultsRequest) GetTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *ResultsRequest) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *ResultsRequest) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetTrackingToken returns the TrackingToken field value if set, zero value otherwise.
func (o *ResultsRequest) GetTrackingToken() string {
	if o == nil || IsNil(o.TrackingToken) {
		var ret string
		return ret
	}
	return *o.TrackingToken
}

// GetTrackingTokenOk returns a tuple with the TrackingToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultsRequest) GetTrackingTokenOk() (*string, bool) {
	if o == nil || IsNil(o.TrackingToken) {
		return nil, false
	}
	return o.TrackingToken, true
}

// HasTrackingToken returns a boolean if a field has been set.
func (o *ResultsRequest) HasTrackingToken() bool {
	if o != nil && !IsNil(o.TrackingToken) {
		return true
	}

	return false
}

// SetTrackingToken gets a reference to the given string and assigns it to the TrackingToken field.
func (o *ResultsRequest) SetTrackingToken(v string) {
	o.TrackingToken = &v
}

// GetSessionInfo returns the SessionInfo field value if set, zero value otherwise.
func (o *ResultsRequest) GetSessionInfo() SessionInfo {
	if o == nil || IsNil(o.SessionInfo) {
		var ret SessionInfo
		return ret
	}
	return *o.SessionInfo
}

// GetSessionInfoOk returns a tuple with the SessionInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultsRequest) GetSessionInfoOk() (*SessionInfo, bool) {
	if o == nil || IsNil(o.SessionInfo) {
		return nil, false
	}
	return o.SessionInfo, true
}

// HasSessionInfo returns a boolean if a field has been set.
func (o *ResultsRequest) HasSessionInfo() bool {
	if o != nil && !IsNil(o.SessionInfo) {
		return true
	}

	return false
}

// SetSessionInfo gets a reference to the given SessionInfo and assigns it to the SessionInfo field.
func (o *ResultsRequest) SetSessionInfo(v SessionInfo) {
	o.SessionInfo = &v
}

// GetSourceInfo returns the SourceInfo field value
func (o *ResultsRequest) GetSourceInfo() SearchRequestSourceInfo {
	if o == nil {
		var ret SearchRequestSourceInfo
		return ret
	}

	return o.SourceInfo
}

// GetSourceInfoOk returns a tuple with the SourceInfo field value
// and a boolean to check if the value has been set.
func (o *ResultsRequest) GetSourceInfoOk() (*SearchRequestSourceInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceInfo, true
}

// SetSourceInfo sets field value
func (o *ResultsRequest) SetSourceInfo(v SearchRequestSourceInfo) {
	o.SourceInfo = v
}

// GetSourceDocument returns the SourceDocument field value if set, zero value otherwise.
func (o *ResultsRequest) GetSourceDocument() Document {
	if o == nil || IsNil(o.SourceDocument) {
		var ret Document
		return ret
	}
	return *o.SourceDocument
}

// GetSourceDocumentOk returns a tuple with the SourceDocument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultsRequest) GetSourceDocumentOk() (*Document, bool) {
	if o == nil || IsNil(o.SourceDocument) {
		return nil, false
	}
	return o.SourceDocument, true
}

// HasSourceDocument returns a boolean if a field has been set.
func (o *ResultsRequest) HasSourceDocument() bool {
	if o != nil && !IsNil(o.SourceDocument) {
		return true
	}

	return false
}

// SetSourceDocument gets a reference to the given Document and assigns it to the SourceDocument field.
func (o *ResultsRequest) SetSourceDocument(v Document) {
	o.SourceDocument = &v
}

// GetSc returns the Sc field value if set, zero value otherwise.
func (o *ResultsRequest) GetSc() string {
	if o == nil || IsNil(o.Sc) {
		var ret string
		return ret
	}
	return *o.Sc
}

// GetScOk returns a tuple with the Sc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultsRequest) GetScOk() (*string, bool) {
	if o == nil || IsNil(o.Sc) {
		return nil, false
	}
	return o.Sc, true
}

// HasSc returns a boolean if a field has been set.
func (o *ResultsRequest) HasSc() bool {
	if o != nil && !IsNil(o.Sc) {
		return true
	}

	return false
}

// SetSc gets a reference to the given string and assigns it to the Sc field.
func (o *ResultsRequest) SetSc(v string) {
	o.Sc = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *ResultsRequest) GetPageSize() int32 {
	if o == nil || IsNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultsRequest) GetPageSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *ResultsRequest) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *ResultsRequest) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetMaxSnippetSize returns the MaxSnippetSize field value if set, zero value otherwise.
func (o *ResultsRequest) GetMaxSnippetSize() int32 {
	if o == nil || IsNil(o.MaxSnippetSize) {
		var ret int32
		return ret
	}
	return *o.MaxSnippetSize
}

// GetMaxSnippetSizeOk returns a tuple with the MaxSnippetSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultsRequest) GetMaxSnippetSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxSnippetSize) {
		return nil, false
	}
	return o.MaxSnippetSize, true
}

// HasMaxSnippetSize returns a boolean if a field has been set.
func (o *ResultsRequest) HasMaxSnippetSize() bool {
	if o != nil && !IsNil(o.MaxSnippetSize) {
		return true
	}

	return false
}

// SetMaxSnippetSize gets a reference to the given int32 and assigns it to the MaxSnippetSize field.
func (o *ResultsRequest) SetMaxSnippetSize(v int32) {
	o.MaxSnippetSize = &v
}

func (o ResultsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResultsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !IsNil(o.TrackingToken) {
		toSerialize["trackingToken"] = o.TrackingToken
	}
	if !IsNil(o.SessionInfo) {
		toSerialize["sessionInfo"] = o.SessionInfo
	}
	toSerialize["sourceInfo"] = o.SourceInfo
	if !IsNil(o.SourceDocument) {
		toSerialize["sourceDocument"] = o.SourceDocument
	}
	if !IsNil(o.Sc) {
		toSerialize["sc"] = o.Sc
	}
	if !IsNil(o.PageSize) {
		toSerialize["pageSize"] = o.PageSize
	}
	if !IsNil(o.MaxSnippetSize) {
		toSerialize["maxSnippetSize"] = o.MaxSnippetSize
	}
	return toSerialize, nil
}

type NullableResultsRequest struct {
	value *ResultsRequest
	isSet bool
}

func (v NullableResultsRequest) Get() *ResultsRequest {
	return v.value
}

func (v *NullableResultsRequest) Set(val *ResultsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableResultsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableResultsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResultsRequest(val *ResultsRequest) *NullableResultsRequest {
	return &NullableResultsRequest{value: val, isSet: true}
}

func (v NullableResultsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResultsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


