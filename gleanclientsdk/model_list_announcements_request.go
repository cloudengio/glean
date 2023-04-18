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

// checks if the ListAnnouncementsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListAnnouncementsRequest{}

// ListAnnouncementsRequest struct for ListAnnouncementsRequest
type ListAnnouncementsRequest struct {
	Channel *AnnouncementChannel `json:"channel,omitempty"`
}

// NewListAnnouncementsRequest instantiates a new ListAnnouncementsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListAnnouncementsRequest() *ListAnnouncementsRequest {
	this := ListAnnouncementsRequest{}
	return &this
}

// NewListAnnouncementsRequestWithDefaults instantiates a new ListAnnouncementsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListAnnouncementsRequestWithDefaults() *ListAnnouncementsRequest {
	this := ListAnnouncementsRequest{}
	return &this
}

// GetChannel returns the Channel field value if set, zero value otherwise.
func (o *ListAnnouncementsRequest) GetChannel() AnnouncementChannel {
	if o == nil || IsNil(o.Channel) {
		var ret AnnouncementChannel
		return ret
	}
	return *o.Channel
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAnnouncementsRequest) GetChannelOk() (*AnnouncementChannel, bool) {
	if o == nil || IsNil(o.Channel) {
		return nil, false
	}
	return o.Channel, true
}

// HasChannel returns a boolean if a field has been set.
func (o *ListAnnouncementsRequest) HasChannel() bool {
	if o != nil && !IsNil(o.Channel) {
		return true
	}

	return false
}

// SetChannel gets a reference to the given AnnouncementChannel and assigns it to the Channel field.
func (o *ListAnnouncementsRequest) SetChannel(v AnnouncementChannel) {
	o.Channel = &v
}

func (o ListAnnouncementsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListAnnouncementsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Channel) {
		toSerialize["channel"] = o.Channel
	}
	return toSerialize, nil
}

type NullableListAnnouncementsRequest struct {
	value *ListAnnouncementsRequest
	isSet bool
}

func (v NullableListAnnouncementsRequest) Get() *ListAnnouncementsRequest {
	return v.value
}

func (v *NullableListAnnouncementsRequest) Set(val *ListAnnouncementsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableListAnnouncementsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableListAnnouncementsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListAnnouncementsRequest(val *ListAnnouncementsRequest) *NullableListAnnouncementsRequest {
	return &NullableListAnnouncementsRequest{value: val, isSet: true}
}

func (v NullableListAnnouncementsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListAnnouncementsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


