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

// checks if the GetAnnouncementResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAnnouncementResponse{}

// GetAnnouncementResponse struct for GetAnnouncementResponse
type GetAnnouncementResponse struct {
	Announcement *Announcement `json:"announcement,omitempty"`
	// An opaque token that represents this particular announcement. To be used for /feedback reporting.
	TrackingToken *string `json:"trackingToken,omitempty"`
	Error *AnnouncementError `json:"error,omitempty"`
}

// NewGetAnnouncementResponse instantiates a new GetAnnouncementResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAnnouncementResponse() *GetAnnouncementResponse {
	this := GetAnnouncementResponse{}
	return &this
}

// NewGetAnnouncementResponseWithDefaults instantiates a new GetAnnouncementResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAnnouncementResponseWithDefaults() *GetAnnouncementResponse {
	this := GetAnnouncementResponse{}
	return &this
}

// GetAnnouncement returns the Announcement field value if set, zero value otherwise.
func (o *GetAnnouncementResponse) GetAnnouncement() Announcement {
	if o == nil || IsNil(o.Announcement) {
		var ret Announcement
		return ret
	}
	return *o.Announcement
}

// GetAnnouncementOk returns a tuple with the Announcement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAnnouncementResponse) GetAnnouncementOk() (*Announcement, bool) {
	if o == nil || IsNil(o.Announcement) {
		return nil, false
	}
	return o.Announcement, true
}

// HasAnnouncement returns a boolean if a field has been set.
func (o *GetAnnouncementResponse) HasAnnouncement() bool {
	if o != nil && !IsNil(o.Announcement) {
		return true
	}

	return false
}

// SetAnnouncement gets a reference to the given Announcement and assigns it to the Announcement field.
func (o *GetAnnouncementResponse) SetAnnouncement(v Announcement) {
	o.Announcement = &v
}

// GetTrackingToken returns the TrackingToken field value if set, zero value otherwise.
func (o *GetAnnouncementResponse) GetTrackingToken() string {
	if o == nil || IsNil(o.TrackingToken) {
		var ret string
		return ret
	}
	return *o.TrackingToken
}

// GetTrackingTokenOk returns a tuple with the TrackingToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAnnouncementResponse) GetTrackingTokenOk() (*string, bool) {
	if o == nil || IsNil(o.TrackingToken) {
		return nil, false
	}
	return o.TrackingToken, true
}

// HasTrackingToken returns a boolean if a field has been set.
func (o *GetAnnouncementResponse) HasTrackingToken() bool {
	if o != nil && !IsNil(o.TrackingToken) {
		return true
	}

	return false
}

// SetTrackingToken gets a reference to the given string and assigns it to the TrackingToken field.
func (o *GetAnnouncementResponse) SetTrackingToken(v string) {
	o.TrackingToken = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *GetAnnouncementResponse) GetError() AnnouncementError {
	if o == nil || IsNil(o.Error) {
		var ret AnnouncementError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAnnouncementResponse) GetErrorOk() (*AnnouncementError, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *GetAnnouncementResponse) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given AnnouncementError and assigns it to the Error field.
func (o *GetAnnouncementResponse) SetError(v AnnouncementError) {
	o.Error = &v
}

func (o GetAnnouncementResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAnnouncementResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Announcement) {
		toSerialize["announcement"] = o.Announcement
	}
	if !IsNil(o.TrackingToken) {
		toSerialize["trackingToken"] = o.TrackingToken
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableGetAnnouncementResponse struct {
	value *GetAnnouncementResponse
	isSet bool
}

func (v NullableGetAnnouncementResponse) Get() *GetAnnouncementResponse {
	return v.value
}

func (v *NullableGetAnnouncementResponse) Set(val *GetAnnouncementResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAnnouncementResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAnnouncementResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAnnouncementResponse(val *GetAnnouncementResponse) *NullableGetAnnouncementResponse {
	return &NullableGetAnnouncementResponse{value: val, isSet: true}
}

func (v NullableGetAnnouncementResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAnnouncementResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

