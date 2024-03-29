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

// checks if the PeopleResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PeopleResponse{}

// PeopleResponse struct for PeopleResponse
type PeopleResponse struct {
	// A Person for each ID in the request, each with PersonMetadata populated.
	Results []Person `json:"results,omitempty"`
	// A list of documents related to this people response. This is only included if DOCUMENT_ACTIVITY is requested and only 1 person is included in the request.
	RelatedDocuments []RelatedDocuments `json:"relatedDocuments,omitempty"`
	// A list of IDs that could not be found.
	Errors []string `json:"errors,omitempty"`
}

// NewPeopleResponse instantiates a new PeopleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPeopleResponse() *PeopleResponse {
	this := PeopleResponse{}
	return &this
}

// NewPeopleResponseWithDefaults instantiates a new PeopleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPeopleResponseWithDefaults() *PeopleResponse {
	this := PeopleResponse{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *PeopleResponse) GetResults() []Person {
	if o == nil || IsNil(o.Results) {
		var ret []Person
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PeopleResponse) GetResultsOk() ([]Person, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *PeopleResponse) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []Person and assigns it to the Results field.
func (o *PeopleResponse) SetResults(v []Person) {
	o.Results = v
}

// GetRelatedDocuments returns the RelatedDocuments field value if set, zero value otherwise.
func (o *PeopleResponse) GetRelatedDocuments() []RelatedDocuments {
	if o == nil || IsNil(o.RelatedDocuments) {
		var ret []RelatedDocuments
		return ret
	}
	return o.RelatedDocuments
}

// GetRelatedDocumentsOk returns a tuple with the RelatedDocuments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PeopleResponse) GetRelatedDocumentsOk() ([]RelatedDocuments, bool) {
	if o == nil || IsNil(o.RelatedDocuments) {
		return nil, false
	}
	return o.RelatedDocuments, true
}

// HasRelatedDocuments returns a boolean if a field has been set.
func (o *PeopleResponse) HasRelatedDocuments() bool {
	if o != nil && !IsNil(o.RelatedDocuments) {
		return true
	}

	return false
}

// SetRelatedDocuments gets a reference to the given []RelatedDocuments and assigns it to the RelatedDocuments field.
func (o *PeopleResponse) SetRelatedDocuments(v []RelatedDocuments) {
	o.RelatedDocuments = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *PeopleResponse) GetErrors() []string {
	if o == nil || IsNil(o.Errors) {
		var ret []string
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PeopleResponse) GetErrorsOk() ([]string, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *PeopleResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []string and assigns it to the Errors field.
func (o *PeopleResponse) SetErrors(v []string) {
	o.Errors = v
}

func (o PeopleResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PeopleResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	if !IsNil(o.RelatedDocuments) {
		toSerialize["relatedDocuments"] = o.RelatedDocuments
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullablePeopleResponse struct {
	value *PeopleResponse
	isSet bool
}

func (v NullablePeopleResponse) Get() *PeopleResponse {
	return v.value
}

func (v *NullablePeopleResponse) Set(val *PeopleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePeopleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePeopleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePeopleResponse(val *PeopleResponse) *NullablePeopleResponse {
	return &NullablePeopleResponse{value: val, isSet: true}
}

func (v NullablePeopleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePeopleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


