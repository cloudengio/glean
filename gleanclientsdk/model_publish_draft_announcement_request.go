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

// checks if the PublishDraftAnnouncementRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublishDraftAnnouncementRequest{}

// PublishDraftAnnouncementRequest struct for PublishDraftAnnouncementRequest
type PublishDraftAnnouncementRequest struct {
	// The opaque id of the draft announcement to be published.
	Id int32 `json:"id"`
}

// NewPublishDraftAnnouncementRequest instantiates a new PublishDraftAnnouncementRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublishDraftAnnouncementRequest(id int32) *PublishDraftAnnouncementRequest {
	this := PublishDraftAnnouncementRequest{}
	this.Id = id
	return &this
}

// NewPublishDraftAnnouncementRequestWithDefaults instantiates a new PublishDraftAnnouncementRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublishDraftAnnouncementRequestWithDefaults() *PublishDraftAnnouncementRequest {
	this := PublishDraftAnnouncementRequest{}
	return &this
}

// GetId returns the Id field value
func (o *PublishDraftAnnouncementRequest) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PublishDraftAnnouncementRequest) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PublishDraftAnnouncementRequest) SetId(v int32) {
	o.Id = v
}

func (o PublishDraftAnnouncementRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublishDraftAnnouncementRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullablePublishDraftAnnouncementRequest struct {
	value *PublishDraftAnnouncementRequest
	isSet bool
}

func (v NullablePublishDraftAnnouncementRequest) Get() *PublishDraftAnnouncementRequest {
	return v.value
}

func (v *NullablePublishDraftAnnouncementRequest) Set(val *PublishDraftAnnouncementRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePublishDraftAnnouncementRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePublishDraftAnnouncementRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublishDraftAnnouncementRequest(val *PublishDraftAnnouncementRequest) *NullablePublishDraftAnnouncementRequest {
	return &NullablePublishDraftAnnouncementRequest{value: val, isSet: true}
}

func (v NullablePublishDraftAnnouncementRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublishDraftAnnouncementRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


