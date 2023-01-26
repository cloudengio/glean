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

// CheckDocumentAccessRequest Describes the request body of the /checkdocumentaccess API call
type CheckDocumentAccessRequest struct {
	// Datasource of document to get check access for
	Datasource string `json:"datasource"`
	// Object type of document to get check access for
	ObjectType string `json:"objectType"`
	// ID of document to get check access for
	DocId string `json:"docId"`
	// Email of user to check access for
	UserEmail string `json:"userEmail"`
}

// NewCheckDocumentAccessRequest instantiates a new CheckDocumentAccessRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckDocumentAccessRequest(datasource string, objectType string, docId string, userEmail string) *CheckDocumentAccessRequest {
	this := CheckDocumentAccessRequest{}
	this.Datasource = datasource
	this.ObjectType = objectType
	this.DocId = docId
	this.UserEmail = userEmail
	return &this
}

// NewCheckDocumentAccessRequestWithDefaults instantiates a new CheckDocumentAccessRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckDocumentAccessRequestWithDefaults() *CheckDocumentAccessRequest {
	this := CheckDocumentAccessRequest{}
	return &this
}

// GetDatasource returns the Datasource field value
func (o *CheckDocumentAccessRequest) GetDatasource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datasource
}

// GetDatasourceOk returns a tuple with the Datasource field value
// and a boolean to check if the value has been set.
func (o *CheckDocumentAccessRequest) GetDatasourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datasource, true
}

// SetDatasource sets field value
func (o *CheckDocumentAccessRequest) SetDatasource(v string) {
	o.Datasource = v
}

// GetObjectType returns the ObjectType field value
func (o *CheckDocumentAccessRequest) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CheckDocumentAccessRequest) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CheckDocumentAccessRequest) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDocId returns the DocId field value
func (o *CheckDocumentAccessRequest) GetDocId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DocId
}

// GetDocIdOk returns a tuple with the DocId field value
// and a boolean to check if the value has been set.
func (o *CheckDocumentAccessRequest) GetDocIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DocId, true
}

// SetDocId sets field value
func (o *CheckDocumentAccessRequest) SetDocId(v string) {
	o.DocId = v
}

// GetUserEmail returns the UserEmail field value
func (o *CheckDocumentAccessRequest) GetUserEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserEmail
}

// GetUserEmailOk returns a tuple with the UserEmail field value
// and a boolean to check if the value has been set.
func (o *CheckDocumentAccessRequest) GetUserEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserEmail, true
}

// SetUserEmail sets field value
func (o *CheckDocumentAccessRequest) SetUserEmail(v string) {
	o.UserEmail = v
}

func (o CheckDocumentAccessRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["datasource"] = o.Datasource
	}
	if true {
		toSerialize["objectType"] = o.ObjectType
	}
	if true {
		toSerialize["docId"] = o.DocId
	}
	if true {
		toSerialize["userEmail"] = o.UserEmail
	}
	return json.Marshal(toSerialize)
}

type NullableCheckDocumentAccessRequest struct {
	value *CheckDocumentAccessRequest
	isSet bool
}

func (v NullableCheckDocumentAccessRequest) Get() *CheckDocumentAccessRequest {
	return v.value
}

func (v *NullableCheckDocumentAccessRequest) Set(val *CheckDocumentAccessRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckDocumentAccessRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckDocumentAccessRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckDocumentAccessRequest(val *CheckDocumentAccessRequest) *NullableCheckDocumentAccessRequest {
	return &NullableCheckDocumentAccessRequest{value: val, isSet: true}
}

func (v NullableCheckDocumentAccessRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckDocumentAccessRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


