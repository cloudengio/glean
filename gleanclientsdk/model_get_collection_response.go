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

// checks if the GetCollectionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCollectionResponse{}

// GetCollectionResponse struct for GetCollectionResponse
type GetCollectionResponse struct {
	Collection *Collection `json:"collection,omitempty"`
	RootCollection *Collection `json:"rootCollection,omitempty"`
	// An opaque token that represents this particular collection. To be used for /feedback reporting.
	TrackingToken *string `json:"trackingToken,omitempty"`
	Error *CollectionError `json:"error,omitempty"`
}

// NewGetCollectionResponse instantiates a new GetCollectionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCollectionResponse() *GetCollectionResponse {
	this := GetCollectionResponse{}
	return &this
}

// NewGetCollectionResponseWithDefaults instantiates a new GetCollectionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCollectionResponseWithDefaults() *GetCollectionResponse {
	this := GetCollectionResponse{}
	return &this
}

// GetCollection returns the Collection field value if set, zero value otherwise.
func (o *GetCollectionResponse) GetCollection() Collection {
	if o == nil || IsNil(o.Collection) {
		var ret Collection
		return ret
	}
	return *o.Collection
}

// GetCollectionOk returns a tuple with the Collection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCollectionResponse) GetCollectionOk() (*Collection, bool) {
	if o == nil || IsNil(o.Collection) {
		return nil, false
	}
	return o.Collection, true
}

// HasCollection returns a boolean if a field has been set.
func (o *GetCollectionResponse) HasCollection() bool {
	if o != nil && !IsNil(o.Collection) {
		return true
	}

	return false
}

// SetCollection gets a reference to the given Collection and assigns it to the Collection field.
func (o *GetCollectionResponse) SetCollection(v Collection) {
	o.Collection = &v
}

// GetRootCollection returns the RootCollection field value if set, zero value otherwise.
func (o *GetCollectionResponse) GetRootCollection() Collection {
	if o == nil || IsNil(o.RootCollection) {
		var ret Collection
		return ret
	}
	return *o.RootCollection
}

// GetRootCollectionOk returns a tuple with the RootCollection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCollectionResponse) GetRootCollectionOk() (*Collection, bool) {
	if o == nil || IsNil(o.RootCollection) {
		return nil, false
	}
	return o.RootCollection, true
}

// HasRootCollection returns a boolean if a field has been set.
func (o *GetCollectionResponse) HasRootCollection() bool {
	if o != nil && !IsNil(o.RootCollection) {
		return true
	}

	return false
}

// SetRootCollection gets a reference to the given Collection and assigns it to the RootCollection field.
func (o *GetCollectionResponse) SetRootCollection(v Collection) {
	o.RootCollection = &v
}

// GetTrackingToken returns the TrackingToken field value if set, zero value otherwise.
func (o *GetCollectionResponse) GetTrackingToken() string {
	if o == nil || IsNil(o.TrackingToken) {
		var ret string
		return ret
	}
	return *o.TrackingToken
}

// GetTrackingTokenOk returns a tuple with the TrackingToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCollectionResponse) GetTrackingTokenOk() (*string, bool) {
	if o == nil || IsNil(o.TrackingToken) {
		return nil, false
	}
	return o.TrackingToken, true
}

// HasTrackingToken returns a boolean if a field has been set.
func (o *GetCollectionResponse) HasTrackingToken() bool {
	if o != nil && !IsNil(o.TrackingToken) {
		return true
	}

	return false
}

// SetTrackingToken gets a reference to the given string and assigns it to the TrackingToken field.
func (o *GetCollectionResponse) SetTrackingToken(v string) {
	o.TrackingToken = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *GetCollectionResponse) GetError() CollectionError {
	if o == nil || IsNil(o.Error) {
		var ret CollectionError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCollectionResponse) GetErrorOk() (*CollectionError, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *GetCollectionResponse) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given CollectionError and assigns it to the Error field.
func (o *GetCollectionResponse) SetError(v CollectionError) {
	o.Error = &v
}

func (o GetCollectionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCollectionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Collection) {
		toSerialize["collection"] = o.Collection
	}
	if !IsNil(o.RootCollection) {
		toSerialize["rootCollection"] = o.RootCollection
	}
	if !IsNil(o.TrackingToken) {
		toSerialize["trackingToken"] = o.TrackingToken
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableGetCollectionResponse struct {
	value *GetCollectionResponse
	isSet bool
}

func (v NullableGetCollectionResponse) Get() *GetCollectionResponse {
	return v.value
}

func (v *NullableGetCollectionResponse) Set(val *GetCollectionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCollectionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCollectionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCollectionResponse(val *GetCollectionResponse) *NullableGetCollectionResponse {
	return &NullableGetCollectionResponse{value: val, isSet: true}
}

func (v NullableGetCollectionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCollectionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


