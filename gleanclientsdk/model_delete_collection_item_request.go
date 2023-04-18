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

// checks if the DeleteCollectionItemRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteCollectionItemRequest{}

// DeleteCollectionItemRequest struct for DeleteCollectionItemRequest
type DeleteCollectionItemRequest struct {
	// The ID of the Collection to remove an item in.
	CollectionId float32 `json:"collectionId"`
	// The item ID of the CollectionItem to remove from this Collection.
	ItemId string `json:"itemId"`
	// The (optional) document ID of the CollectionItem to remove from this Collection if this is a Glean-indexed document.
	DocumentId *string `json:"documentId,omitempty"`
}

// NewDeleteCollectionItemRequest instantiates a new DeleteCollectionItemRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteCollectionItemRequest(collectionId float32, itemId string) *DeleteCollectionItemRequest {
	this := DeleteCollectionItemRequest{}
	this.CollectionId = collectionId
	this.ItemId = itemId
	return &this
}

// NewDeleteCollectionItemRequestWithDefaults instantiates a new DeleteCollectionItemRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteCollectionItemRequestWithDefaults() *DeleteCollectionItemRequest {
	this := DeleteCollectionItemRequest{}
	return &this
}

// GetCollectionId returns the CollectionId field value
func (o *DeleteCollectionItemRequest) GetCollectionId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CollectionId
}

// GetCollectionIdOk returns a tuple with the CollectionId field value
// and a boolean to check if the value has been set.
func (o *DeleteCollectionItemRequest) GetCollectionIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CollectionId, true
}

// SetCollectionId sets field value
func (o *DeleteCollectionItemRequest) SetCollectionId(v float32) {
	o.CollectionId = v
}

// GetItemId returns the ItemId field value
func (o *DeleteCollectionItemRequest) GetItemId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value
// and a boolean to check if the value has been set.
func (o *DeleteCollectionItemRequest) GetItemIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ItemId, true
}

// SetItemId sets field value
func (o *DeleteCollectionItemRequest) SetItemId(v string) {
	o.ItemId = v
}

// GetDocumentId returns the DocumentId field value if set, zero value otherwise.
func (o *DeleteCollectionItemRequest) GetDocumentId() string {
	if o == nil || IsNil(o.DocumentId) {
		var ret string
		return ret
	}
	return *o.DocumentId
}

// GetDocumentIdOk returns a tuple with the DocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteCollectionItemRequest) GetDocumentIdOk() (*string, bool) {
	if o == nil || IsNil(o.DocumentId) {
		return nil, false
	}
	return o.DocumentId, true
}

// HasDocumentId returns a boolean if a field has been set.
func (o *DeleteCollectionItemRequest) HasDocumentId() bool {
	if o != nil && !IsNil(o.DocumentId) {
		return true
	}

	return false
}

// SetDocumentId gets a reference to the given string and assigns it to the DocumentId field.
func (o *DeleteCollectionItemRequest) SetDocumentId(v string) {
	o.DocumentId = &v
}

func (o DeleteCollectionItemRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteCollectionItemRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["collectionId"] = o.CollectionId
	toSerialize["itemId"] = o.ItemId
	if !IsNil(o.DocumentId) {
		toSerialize["documentId"] = o.DocumentId
	}
	return toSerialize, nil
}

type NullableDeleteCollectionItemRequest struct {
	value *DeleteCollectionItemRequest
	isSet bool
}

func (v NullableDeleteCollectionItemRequest) Get() *DeleteCollectionItemRequest {
	return v.value
}

func (v *NullableDeleteCollectionItemRequest) Set(val *DeleteCollectionItemRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteCollectionItemRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteCollectionItemRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteCollectionItemRequest(val *DeleteCollectionItemRequest) *NullableDeleteCollectionItemRequest {
	return &NullableDeleteCollectionItemRequest{value: val, isSet: true}
}

func (v NullableDeleteCollectionItemRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteCollectionItemRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

