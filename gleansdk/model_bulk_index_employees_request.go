/*
Glean Indexing API

# Introduction In addition to the data sources that Glean has built-in support for, Glean also provides a REST API that enables customers to put arbitrary content in the search index. This is useful, for example, for doing permissions-aware search over content in internal tools that reside on-prem as well as for searching over applications that Glean does not currently support first class. In addition these APIs allow the customer to push organization data (people info, organization structure etc) into Glean.  # Early Access Please note that we are continually evolving the system so these APIs are considered “early access” and are subject to change. 

API version: 0.9.0
Contact: support@glean.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gleansdk

import (
	"encoding/json"
)

// BulkIndexEmployeesRequest Describes the request body of the /bulkindexemployees API call
type BulkIndexEmployeesRequest struct {
	// Unique id that must be used for this instance of datasource employees upload
	UploadId string `json:"uploadId"`
	// true if this is the first page of the upload. Defaults to false
	IsFirstPage *bool `json:"isFirstPage,omitempty"`
	// true if this is the last page of the upload. Defaults to false
	IsLastPage *bool `json:"isLastPage,omitempty"`
	// Flag to discard previous upload attempts and start from scratch. Must be specified with isFirstPage=true
	ForceRestartUpload *bool `json:"forceRestartUpload,omitempty"`
	// Batch of employee information
	Employees []EmployeeInfoDefinition `json:"employees"`
}

// NewBulkIndexEmployeesRequest instantiates a new BulkIndexEmployeesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkIndexEmployeesRequest(uploadId string, employees []EmployeeInfoDefinition) *BulkIndexEmployeesRequest {
	this := BulkIndexEmployeesRequest{}
	this.UploadId = uploadId
	this.Employees = employees
	return &this
}

// NewBulkIndexEmployeesRequestWithDefaults instantiates a new BulkIndexEmployeesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkIndexEmployeesRequestWithDefaults() *BulkIndexEmployeesRequest {
	this := BulkIndexEmployeesRequest{}
	return &this
}

// GetUploadId returns the UploadId field value
func (o *BulkIndexEmployeesRequest) GetUploadId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UploadId
}

// GetUploadIdOk returns a tuple with the UploadId field value
// and a boolean to check if the value has been set.
func (o *BulkIndexEmployeesRequest) GetUploadIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UploadId, true
}

// SetUploadId sets field value
func (o *BulkIndexEmployeesRequest) SetUploadId(v string) {
	o.UploadId = v
}

// GetIsFirstPage returns the IsFirstPage field value if set, zero value otherwise.
func (o *BulkIndexEmployeesRequest) GetIsFirstPage() bool {
	if o == nil || o.IsFirstPage == nil {
		var ret bool
		return ret
	}
	return *o.IsFirstPage
}

// GetIsFirstPageOk returns a tuple with the IsFirstPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkIndexEmployeesRequest) GetIsFirstPageOk() (*bool, bool) {
	if o == nil || o.IsFirstPage == nil {
		return nil, false
	}
	return o.IsFirstPage, true
}

// HasIsFirstPage returns a boolean if a field has been set.
func (o *BulkIndexEmployeesRequest) HasIsFirstPage() bool {
	if o != nil && o.IsFirstPage != nil {
		return true
	}

	return false
}

// SetIsFirstPage gets a reference to the given bool and assigns it to the IsFirstPage field.
func (o *BulkIndexEmployeesRequest) SetIsFirstPage(v bool) {
	o.IsFirstPage = &v
}

// GetIsLastPage returns the IsLastPage field value if set, zero value otherwise.
func (o *BulkIndexEmployeesRequest) GetIsLastPage() bool {
	if o == nil || o.IsLastPage == nil {
		var ret bool
		return ret
	}
	return *o.IsLastPage
}

// GetIsLastPageOk returns a tuple with the IsLastPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkIndexEmployeesRequest) GetIsLastPageOk() (*bool, bool) {
	if o == nil || o.IsLastPage == nil {
		return nil, false
	}
	return o.IsLastPage, true
}

// HasIsLastPage returns a boolean if a field has been set.
func (o *BulkIndexEmployeesRequest) HasIsLastPage() bool {
	if o != nil && o.IsLastPage != nil {
		return true
	}

	return false
}

// SetIsLastPage gets a reference to the given bool and assigns it to the IsLastPage field.
func (o *BulkIndexEmployeesRequest) SetIsLastPage(v bool) {
	o.IsLastPage = &v
}

// GetForceRestartUpload returns the ForceRestartUpload field value if set, zero value otherwise.
func (o *BulkIndexEmployeesRequest) GetForceRestartUpload() bool {
	if o == nil || o.ForceRestartUpload == nil {
		var ret bool
		return ret
	}
	return *o.ForceRestartUpload
}

// GetForceRestartUploadOk returns a tuple with the ForceRestartUpload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkIndexEmployeesRequest) GetForceRestartUploadOk() (*bool, bool) {
	if o == nil || o.ForceRestartUpload == nil {
		return nil, false
	}
	return o.ForceRestartUpload, true
}

// HasForceRestartUpload returns a boolean if a field has been set.
func (o *BulkIndexEmployeesRequest) HasForceRestartUpload() bool {
	if o != nil && o.ForceRestartUpload != nil {
		return true
	}

	return false
}

// SetForceRestartUpload gets a reference to the given bool and assigns it to the ForceRestartUpload field.
func (o *BulkIndexEmployeesRequest) SetForceRestartUpload(v bool) {
	o.ForceRestartUpload = &v
}

// GetEmployees returns the Employees field value
func (o *BulkIndexEmployeesRequest) GetEmployees() []EmployeeInfoDefinition {
	if o == nil {
		var ret []EmployeeInfoDefinition
		return ret
	}

	return o.Employees
}

// GetEmployeesOk returns a tuple with the Employees field value
// and a boolean to check if the value has been set.
func (o *BulkIndexEmployeesRequest) GetEmployeesOk() ([]EmployeeInfoDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return o.Employees, true
}

// SetEmployees sets field value
func (o *BulkIndexEmployeesRequest) SetEmployees(v []EmployeeInfoDefinition) {
	o.Employees = v
}

func (o BulkIndexEmployeesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["uploadId"] = o.UploadId
	}
	if o.IsFirstPage != nil {
		toSerialize["isFirstPage"] = o.IsFirstPage
	}
	if o.IsLastPage != nil {
		toSerialize["isLastPage"] = o.IsLastPage
	}
	if o.ForceRestartUpload != nil {
		toSerialize["forceRestartUpload"] = o.ForceRestartUpload
	}
	if true {
		toSerialize["employees"] = o.Employees
	}
	return json.Marshal(toSerialize)
}

type NullableBulkIndexEmployeesRequest struct {
	value *BulkIndexEmployeesRequest
	isSet bool
}

func (v NullableBulkIndexEmployeesRequest) Get() *BulkIndexEmployeesRequest {
	return v.value
}

func (v *NullableBulkIndexEmployeesRequest) Set(val *BulkIndexEmployeesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkIndexEmployeesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkIndexEmployeesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkIndexEmployeesRequest(val *BulkIndexEmployeesRequest) *NullableBulkIndexEmployeesRequest {
	return &NullableBulkIndexEmployeesRequest{value: val, isSet: true}
}

func (v NullableBulkIndexEmployeesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkIndexEmployeesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


