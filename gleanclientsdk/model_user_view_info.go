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

// checks if the UserViewInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserViewInfo{}

// UserViewInfo struct for UserViewInfo
type UserViewInfo struct {
	// Unique identifier of associated document
	DocId *string `json:"docId,omitempty"`
	// Title of associated document
	DocTitle *string `json:"docTitle,omitempty"`
	// URL of associated document
	DocUrl *string `json:"docUrl,omitempty"`
}

// NewUserViewInfo instantiates a new UserViewInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserViewInfo() *UserViewInfo {
	this := UserViewInfo{}
	return &this
}

// NewUserViewInfoWithDefaults instantiates a new UserViewInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserViewInfoWithDefaults() *UserViewInfo {
	this := UserViewInfo{}
	return &this
}

// GetDocId returns the DocId field value if set, zero value otherwise.
func (o *UserViewInfo) GetDocId() string {
	if o == nil || IsNil(o.DocId) {
		var ret string
		return ret
	}
	return *o.DocId
}

// GetDocIdOk returns a tuple with the DocId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserViewInfo) GetDocIdOk() (*string, bool) {
	if o == nil || IsNil(o.DocId) {
		return nil, false
	}
	return o.DocId, true
}

// HasDocId returns a boolean if a field has been set.
func (o *UserViewInfo) HasDocId() bool {
	if o != nil && !IsNil(o.DocId) {
		return true
	}

	return false
}

// SetDocId gets a reference to the given string and assigns it to the DocId field.
func (o *UserViewInfo) SetDocId(v string) {
	o.DocId = &v
}

// GetDocTitle returns the DocTitle field value if set, zero value otherwise.
func (o *UserViewInfo) GetDocTitle() string {
	if o == nil || IsNil(o.DocTitle) {
		var ret string
		return ret
	}
	return *o.DocTitle
}

// GetDocTitleOk returns a tuple with the DocTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserViewInfo) GetDocTitleOk() (*string, bool) {
	if o == nil || IsNil(o.DocTitle) {
		return nil, false
	}
	return o.DocTitle, true
}

// HasDocTitle returns a boolean if a field has been set.
func (o *UserViewInfo) HasDocTitle() bool {
	if o != nil && !IsNil(o.DocTitle) {
		return true
	}

	return false
}

// SetDocTitle gets a reference to the given string and assigns it to the DocTitle field.
func (o *UserViewInfo) SetDocTitle(v string) {
	o.DocTitle = &v
}

// GetDocUrl returns the DocUrl field value if set, zero value otherwise.
func (o *UserViewInfo) GetDocUrl() string {
	if o == nil || IsNil(o.DocUrl) {
		var ret string
		return ret
	}
	return *o.DocUrl
}

// GetDocUrlOk returns a tuple with the DocUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserViewInfo) GetDocUrlOk() (*string, bool) {
	if o == nil || IsNil(o.DocUrl) {
		return nil, false
	}
	return o.DocUrl, true
}

// HasDocUrl returns a boolean if a field has been set.
func (o *UserViewInfo) HasDocUrl() bool {
	if o != nil && !IsNil(o.DocUrl) {
		return true
	}

	return false
}

// SetDocUrl gets a reference to the given string and assigns it to the DocUrl field.
func (o *UserViewInfo) SetDocUrl(v string) {
	o.DocUrl = &v
}

func (o UserViewInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserViewInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DocId) {
		toSerialize["docId"] = o.DocId
	}
	if !IsNil(o.DocTitle) {
		toSerialize["docTitle"] = o.DocTitle
	}
	if !IsNil(o.DocUrl) {
		toSerialize["docUrl"] = o.DocUrl
	}
	return toSerialize, nil
}

type NullableUserViewInfo struct {
	value *UserViewInfo
	isSet bool
}

func (v NullableUserViewInfo) Get() *UserViewInfo {
	return v.value
}

func (v *NullableUserViewInfo) Set(val *UserViewInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableUserViewInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableUserViewInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserViewInfo(val *UserViewInfo) *NullableUserViewInfo {
	return &NullableUserViewInfo{value: val, isSet: true}
}

func (v NullableUserViewInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserViewInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


