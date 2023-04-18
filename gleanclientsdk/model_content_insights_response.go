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

// checks if the ContentInsightsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContentInsightsResponse{}

// ContentInsightsResponse struct for ContentInsightsResponse
type ContentInsightsResponse struct {
	// Unix timestamp of the last activity processed to make the response (in seconds since epoch UTC).
	LastLogTimestamp *int32 `json:"lastLogTimestamp,omitempty"`
	// Insights for documents.
	DocumentInsights []DocumentInsight `json:"documentInsights,omitempty"`
	// list of departments applicable for contents tab.
	Departments []string `json:"departments,omitempty"`
	// Min threshold in size of departments while populating results, otherwise 0.
	MinDepartmentSizeThreshold *int32 `json:"minDepartmentSizeThreshold,omitempty"`
	// Minimum number of visitors to a document required to be included in insights.
	MinVisitorThreshold *int32 `json:"minVisitorThreshold,omitempty"`
}

// NewContentInsightsResponse instantiates a new ContentInsightsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentInsightsResponse() *ContentInsightsResponse {
	this := ContentInsightsResponse{}
	return &this
}

// NewContentInsightsResponseWithDefaults instantiates a new ContentInsightsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentInsightsResponseWithDefaults() *ContentInsightsResponse {
	this := ContentInsightsResponse{}
	return &this
}

// GetLastLogTimestamp returns the LastLogTimestamp field value if set, zero value otherwise.
func (o *ContentInsightsResponse) GetLastLogTimestamp() int32 {
	if o == nil || IsNil(o.LastLogTimestamp) {
		var ret int32
		return ret
	}
	return *o.LastLogTimestamp
}

// GetLastLogTimestampOk returns a tuple with the LastLogTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentInsightsResponse) GetLastLogTimestampOk() (*int32, bool) {
	if o == nil || IsNil(o.LastLogTimestamp) {
		return nil, false
	}
	return o.LastLogTimestamp, true
}

// HasLastLogTimestamp returns a boolean if a field has been set.
func (o *ContentInsightsResponse) HasLastLogTimestamp() bool {
	if o != nil && !IsNil(o.LastLogTimestamp) {
		return true
	}

	return false
}

// SetLastLogTimestamp gets a reference to the given int32 and assigns it to the LastLogTimestamp field.
func (o *ContentInsightsResponse) SetLastLogTimestamp(v int32) {
	o.LastLogTimestamp = &v
}

// GetDocumentInsights returns the DocumentInsights field value if set, zero value otherwise.
func (o *ContentInsightsResponse) GetDocumentInsights() []DocumentInsight {
	if o == nil || IsNil(o.DocumentInsights) {
		var ret []DocumentInsight
		return ret
	}
	return o.DocumentInsights
}

// GetDocumentInsightsOk returns a tuple with the DocumentInsights field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentInsightsResponse) GetDocumentInsightsOk() ([]DocumentInsight, bool) {
	if o == nil || IsNil(o.DocumentInsights) {
		return nil, false
	}
	return o.DocumentInsights, true
}

// HasDocumentInsights returns a boolean if a field has been set.
func (o *ContentInsightsResponse) HasDocumentInsights() bool {
	if o != nil && !IsNil(o.DocumentInsights) {
		return true
	}

	return false
}

// SetDocumentInsights gets a reference to the given []DocumentInsight and assigns it to the DocumentInsights field.
func (o *ContentInsightsResponse) SetDocumentInsights(v []DocumentInsight) {
	o.DocumentInsights = v
}

// GetDepartments returns the Departments field value if set, zero value otherwise.
func (o *ContentInsightsResponse) GetDepartments() []string {
	if o == nil || IsNil(o.Departments) {
		var ret []string
		return ret
	}
	return o.Departments
}

// GetDepartmentsOk returns a tuple with the Departments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentInsightsResponse) GetDepartmentsOk() ([]string, bool) {
	if o == nil || IsNil(o.Departments) {
		return nil, false
	}
	return o.Departments, true
}

// HasDepartments returns a boolean if a field has been set.
func (o *ContentInsightsResponse) HasDepartments() bool {
	if o != nil && !IsNil(o.Departments) {
		return true
	}

	return false
}

// SetDepartments gets a reference to the given []string and assigns it to the Departments field.
func (o *ContentInsightsResponse) SetDepartments(v []string) {
	o.Departments = v
}

// GetMinDepartmentSizeThreshold returns the MinDepartmentSizeThreshold field value if set, zero value otherwise.
func (o *ContentInsightsResponse) GetMinDepartmentSizeThreshold() int32 {
	if o == nil || IsNil(o.MinDepartmentSizeThreshold) {
		var ret int32
		return ret
	}
	return *o.MinDepartmentSizeThreshold
}

// GetMinDepartmentSizeThresholdOk returns a tuple with the MinDepartmentSizeThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentInsightsResponse) GetMinDepartmentSizeThresholdOk() (*int32, bool) {
	if o == nil || IsNil(o.MinDepartmentSizeThreshold) {
		return nil, false
	}
	return o.MinDepartmentSizeThreshold, true
}

// HasMinDepartmentSizeThreshold returns a boolean if a field has been set.
func (o *ContentInsightsResponse) HasMinDepartmentSizeThreshold() bool {
	if o != nil && !IsNil(o.MinDepartmentSizeThreshold) {
		return true
	}

	return false
}

// SetMinDepartmentSizeThreshold gets a reference to the given int32 and assigns it to the MinDepartmentSizeThreshold field.
func (o *ContentInsightsResponse) SetMinDepartmentSizeThreshold(v int32) {
	o.MinDepartmentSizeThreshold = &v
}

// GetMinVisitorThreshold returns the MinVisitorThreshold field value if set, zero value otherwise.
func (o *ContentInsightsResponse) GetMinVisitorThreshold() int32 {
	if o == nil || IsNil(o.MinVisitorThreshold) {
		var ret int32
		return ret
	}
	return *o.MinVisitorThreshold
}

// GetMinVisitorThresholdOk returns a tuple with the MinVisitorThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentInsightsResponse) GetMinVisitorThresholdOk() (*int32, bool) {
	if o == nil || IsNil(o.MinVisitorThreshold) {
		return nil, false
	}
	return o.MinVisitorThreshold, true
}

// HasMinVisitorThreshold returns a boolean if a field has been set.
func (o *ContentInsightsResponse) HasMinVisitorThreshold() bool {
	if o != nil && !IsNil(o.MinVisitorThreshold) {
		return true
	}

	return false
}

// SetMinVisitorThreshold gets a reference to the given int32 and assigns it to the MinVisitorThreshold field.
func (o *ContentInsightsResponse) SetMinVisitorThreshold(v int32) {
	o.MinVisitorThreshold = &v
}

func (o ContentInsightsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContentInsightsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LastLogTimestamp) {
		toSerialize["lastLogTimestamp"] = o.LastLogTimestamp
	}
	if !IsNil(o.DocumentInsights) {
		toSerialize["documentInsights"] = o.DocumentInsights
	}
	if !IsNil(o.Departments) {
		toSerialize["departments"] = o.Departments
	}
	if !IsNil(o.MinDepartmentSizeThreshold) {
		toSerialize["minDepartmentSizeThreshold"] = o.MinDepartmentSizeThreshold
	}
	if !IsNil(o.MinVisitorThreshold) {
		toSerialize["minVisitorThreshold"] = o.MinVisitorThreshold
	}
	return toSerialize, nil
}

type NullableContentInsightsResponse struct {
	value *ContentInsightsResponse
	isSet bool
}

func (v NullableContentInsightsResponse) Get() *ContentInsightsResponse {
	return v.value
}

func (v *NullableContentInsightsResponse) Set(val *ContentInsightsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableContentInsightsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableContentInsightsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentInsightsResponse(val *ContentInsightsResponse) *NullableContentInsightsResponse {
	return &NullableContentInsightsResponse{value: val, isSet: true}
}

func (v NullableContentInsightsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentInsightsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


