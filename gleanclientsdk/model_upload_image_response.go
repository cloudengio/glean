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

// checks if the UploadImageResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UploadImageResponse{}

// UploadImageResponse struct for UploadImageResponse
type UploadImageResponse struct {
	// Url of the uploaded image.
	Url string `json:"url"`
	Metadata *ImageMetadata `json:"metadata,omitempty"`
}

// NewUploadImageResponse instantiates a new UploadImageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUploadImageResponse(url string) *UploadImageResponse {
	this := UploadImageResponse{}
	this.Url = url
	return &this
}

// NewUploadImageResponseWithDefaults instantiates a new UploadImageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUploadImageResponseWithDefaults() *UploadImageResponse {
	this := UploadImageResponse{}
	return &this
}

// GetUrl returns the Url field value
func (o *UploadImageResponse) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *UploadImageResponse) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *UploadImageResponse) SetUrl(v string) {
	o.Url = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *UploadImageResponse) GetMetadata() ImageMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret ImageMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadImageResponse) GetMetadataOk() (*ImageMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *UploadImageResponse) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given ImageMetadata and assigns it to the Metadata field.
func (o *UploadImageResponse) SetMetadata(v ImageMetadata) {
	o.Metadata = &v
}

func (o UploadImageResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UploadImageResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["url"] = o.Url
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

type NullableUploadImageResponse struct {
	value *UploadImageResponse
	isSet bool
}

func (v NullableUploadImageResponse) Get() *UploadImageResponse {
	return v.value
}

func (v *NullableUploadImageResponse) Set(val *UploadImageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUploadImageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUploadImageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUploadImageResponse(val *UploadImageResponse) *NullableUploadImageResponse {
	return &NullableUploadImageResponse{value: val, isSet: true}
}

func (v NullableUploadImageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUploadImageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


