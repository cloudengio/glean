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

// checks if the EditDocumentCollectionsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EditDocumentCollectionsRequest{}

// EditDocumentCollectionsRequest struct for EditDocumentCollectionsRequest
type EditDocumentCollectionsRequest struct {
	// IDs of collections to which a document is added.
	AddedCollections []int32 `json:"addedCollections,omitempty"`
	// IDs of collections from which a document is removed.
	RemovedCollections []int32 `json:"removedCollections,omitempty"`
	// The document ID of the item being added to or removed from collections if it's a Glean-indexed document.
	DocumentId *string `json:"documentId,omitempty"`
	// The url of the item being added to or removed from collections.
	Url *string `json:"url,omitempty"`
	// Custom title of the document if adding a non-indexed URL.
	Name *string `json:"name,omitempty"`
	// The description of this CollectionItem.
	Description *string `json:"description,omitempty"`
}

// NewEditDocumentCollectionsRequest instantiates a new EditDocumentCollectionsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditDocumentCollectionsRequest() *EditDocumentCollectionsRequest {
	this := EditDocumentCollectionsRequest{}
	return &this
}

// NewEditDocumentCollectionsRequestWithDefaults instantiates a new EditDocumentCollectionsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditDocumentCollectionsRequestWithDefaults() *EditDocumentCollectionsRequest {
	this := EditDocumentCollectionsRequest{}
	return &this
}

// GetAddedCollections returns the AddedCollections field value if set, zero value otherwise.
func (o *EditDocumentCollectionsRequest) GetAddedCollections() []int32 {
	if o == nil || IsNil(o.AddedCollections) {
		var ret []int32
		return ret
	}
	return o.AddedCollections
}

// GetAddedCollectionsOk returns a tuple with the AddedCollections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditDocumentCollectionsRequest) GetAddedCollectionsOk() ([]int32, bool) {
	if o == nil || IsNil(o.AddedCollections) {
		return nil, false
	}
	return o.AddedCollections, true
}

// HasAddedCollections returns a boolean if a field has been set.
func (o *EditDocumentCollectionsRequest) HasAddedCollections() bool {
	if o != nil && !IsNil(o.AddedCollections) {
		return true
	}

	return false
}

// SetAddedCollections gets a reference to the given []int32 and assigns it to the AddedCollections field.
func (o *EditDocumentCollectionsRequest) SetAddedCollections(v []int32) {
	o.AddedCollections = v
}

// GetRemovedCollections returns the RemovedCollections field value if set, zero value otherwise.
func (o *EditDocumentCollectionsRequest) GetRemovedCollections() []int32 {
	if o == nil || IsNil(o.RemovedCollections) {
		var ret []int32
		return ret
	}
	return o.RemovedCollections
}

// GetRemovedCollectionsOk returns a tuple with the RemovedCollections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditDocumentCollectionsRequest) GetRemovedCollectionsOk() ([]int32, bool) {
	if o == nil || IsNil(o.RemovedCollections) {
		return nil, false
	}
	return o.RemovedCollections, true
}

// HasRemovedCollections returns a boolean if a field has been set.
func (o *EditDocumentCollectionsRequest) HasRemovedCollections() bool {
	if o != nil && !IsNil(o.RemovedCollections) {
		return true
	}

	return false
}

// SetRemovedCollections gets a reference to the given []int32 and assigns it to the RemovedCollections field.
func (o *EditDocumentCollectionsRequest) SetRemovedCollections(v []int32) {
	o.RemovedCollections = v
}

// GetDocumentId returns the DocumentId field value if set, zero value otherwise.
func (o *EditDocumentCollectionsRequest) GetDocumentId() string {
	if o == nil || IsNil(o.DocumentId) {
		var ret string
		return ret
	}
	return *o.DocumentId
}

// GetDocumentIdOk returns a tuple with the DocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditDocumentCollectionsRequest) GetDocumentIdOk() (*string, bool) {
	if o == nil || IsNil(o.DocumentId) {
		return nil, false
	}
	return o.DocumentId, true
}

// HasDocumentId returns a boolean if a field has been set.
func (o *EditDocumentCollectionsRequest) HasDocumentId() bool {
	if o != nil && !IsNil(o.DocumentId) {
		return true
	}

	return false
}

// SetDocumentId gets a reference to the given string and assigns it to the DocumentId field.
func (o *EditDocumentCollectionsRequest) SetDocumentId(v string) {
	o.DocumentId = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *EditDocumentCollectionsRequest) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditDocumentCollectionsRequest) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *EditDocumentCollectionsRequest) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *EditDocumentCollectionsRequest) SetUrl(v string) {
	o.Url = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EditDocumentCollectionsRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditDocumentCollectionsRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EditDocumentCollectionsRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EditDocumentCollectionsRequest) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EditDocumentCollectionsRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditDocumentCollectionsRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EditDocumentCollectionsRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EditDocumentCollectionsRequest) SetDescription(v string) {
	o.Description = &v
}

func (o EditDocumentCollectionsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EditDocumentCollectionsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AddedCollections) {
		toSerialize["addedCollections"] = o.AddedCollections
	}
	if !IsNil(o.RemovedCollections) {
		toSerialize["removedCollections"] = o.RemovedCollections
	}
	if !IsNil(o.DocumentId) {
		toSerialize["documentId"] = o.DocumentId
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableEditDocumentCollectionsRequest struct {
	value *EditDocumentCollectionsRequest
	isSet bool
}

func (v NullableEditDocumentCollectionsRequest) Get() *EditDocumentCollectionsRequest {
	return v.value
}

func (v *NullableEditDocumentCollectionsRequest) Set(val *EditDocumentCollectionsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEditDocumentCollectionsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEditDocumentCollectionsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditDocumentCollectionsRequest(val *EditDocumentCollectionsRequest) *NullableEditDocumentCollectionsRequest {
	return &NullableEditDocumentCollectionsRequest{value: val, isSet: true}
}

func (v NullableEditDocumentCollectionsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditDocumentCollectionsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


