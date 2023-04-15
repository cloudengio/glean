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

// checks if the CollectionPinnedMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionPinnedMetadata{}

// CollectionPinnedMetadata struct for CollectionPinnedMetadata
type CollectionPinnedMetadata struct {
	// List of targets this collection is pinned to
	ExistingPins []CollectionPinTarget `json:"existingPins,omitempty"`
	// List of targets this collection can be pinned to, excluding the targets this collection is already pinned to. We also include collection id already is pinned to each eligible target, which will be 0 if the target has no pinned collection.
	EligiblePins []CollectionPinMetadata `json:"eligiblePins,omitempty"`
	// DEPRECATED - List of categories this collection is pinned to. Use existingPins instead.
	// Deprecated
	PinnedCategories []CollectionPinnableCategories `json:"pinnedCategories"`
	// DEPRECATED - A map of {category, collectionId to bump out} pairs of eligible categories to pin. We exclude categories the collection is already pinned to. CollectionId will be 0 if the the cateogry has no pinned collection. Use eligiblePins instead.
	// Deprecated
	EligibleCategoriesForPinning map[string]int32 `json:"eligibleCategoriesForPinning"`
}

// NewCollectionPinnedMetadata instantiates a new CollectionPinnedMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionPinnedMetadata(pinnedCategories []CollectionPinnableCategories, eligibleCategoriesForPinning map[string]int32) *CollectionPinnedMetadata {
	this := CollectionPinnedMetadata{}
	this.PinnedCategories = pinnedCategories
	this.EligibleCategoriesForPinning = eligibleCategoriesForPinning
	return &this
}

// NewCollectionPinnedMetadataWithDefaults instantiates a new CollectionPinnedMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionPinnedMetadataWithDefaults() *CollectionPinnedMetadata {
	this := CollectionPinnedMetadata{}
	return &this
}

// GetExistingPins returns the ExistingPins field value if set, zero value otherwise.
func (o *CollectionPinnedMetadata) GetExistingPins() []CollectionPinTarget {
	if o == nil || IsNil(o.ExistingPins) {
		var ret []CollectionPinTarget
		return ret
	}
	return o.ExistingPins
}

// GetExistingPinsOk returns a tuple with the ExistingPins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionPinnedMetadata) GetExistingPinsOk() ([]CollectionPinTarget, bool) {
	if o == nil || IsNil(o.ExistingPins) {
		return nil, false
	}
	return o.ExistingPins, true
}

// HasExistingPins returns a boolean if a field has been set.
func (o *CollectionPinnedMetadata) HasExistingPins() bool {
	if o != nil && !IsNil(o.ExistingPins) {
		return true
	}

	return false
}

// SetExistingPins gets a reference to the given []CollectionPinTarget and assigns it to the ExistingPins field.
func (o *CollectionPinnedMetadata) SetExistingPins(v []CollectionPinTarget) {
	o.ExistingPins = v
}

// GetEligiblePins returns the EligiblePins field value if set, zero value otherwise.
func (o *CollectionPinnedMetadata) GetEligiblePins() []CollectionPinMetadata {
	if o == nil || IsNil(o.EligiblePins) {
		var ret []CollectionPinMetadata
		return ret
	}
	return o.EligiblePins
}

// GetEligiblePinsOk returns a tuple with the EligiblePins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionPinnedMetadata) GetEligiblePinsOk() ([]CollectionPinMetadata, bool) {
	if o == nil || IsNil(o.EligiblePins) {
		return nil, false
	}
	return o.EligiblePins, true
}

// HasEligiblePins returns a boolean if a field has been set.
func (o *CollectionPinnedMetadata) HasEligiblePins() bool {
	if o != nil && !IsNil(o.EligiblePins) {
		return true
	}

	return false
}

// SetEligiblePins gets a reference to the given []CollectionPinMetadata and assigns it to the EligiblePins field.
func (o *CollectionPinnedMetadata) SetEligiblePins(v []CollectionPinMetadata) {
	o.EligiblePins = v
}

// GetPinnedCategories returns the PinnedCategories field value
// Deprecated
func (o *CollectionPinnedMetadata) GetPinnedCategories() []CollectionPinnableCategories {
	if o == nil {
		var ret []CollectionPinnableCategories
		return ret
	}

	return o.PinnedCategories
}

// GetPinnedCategoriesOk returns a tuple with the PinnedCategories field value
// and a boolean to check if the value has been set.
// Deprecated
func (o *CollectionPinnedMetadata) GetPinnedCategoriesOk() ([]CollectionPinnableCategories, bool) {
	if o == nil {
		return nil, false
	}
	return o.PinnedCategories, true
}

// SetPinnedCategories sets field value
// Deprecated
func (o *CollectionPinnedMetadata) SetPinnedCategories(v []CollectionPinnableCategories) {
	o.PinnedCategories = v
}

// GetEligibleCategoriesForPinning returns the EligibleCategoriesForPinning field value
// Deprecated
func (o *CollectionPinnedMetadata) GetEligibleCategoriesForPinning() map[string]int32 {
	if o == nil {
		var ret map[string]int32
		return ret
	}

	return o.EligibleCategoriesForPinning
}

// GetEligibleCategoriesForPinningOk returns a tuple with the EligibleCategoriesForPinning field value
// and a boolean to check if the value has been set.
// Deprecated
func (o *CollectionPinnedMetadata) GetEligibleCategoriesForPinningOk() (*map[string]int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EligibleCategoriesForPinning, true
}

// SetEligibleCategoriesForPinning sets field value
// Deprecated
func (o *CollectionPinnedMetadata) SetEligibleCategoriesForPinning(v map[string]int32) {
	o.EligibleCategoriesForPinning = v
}

func (o CollectionPinnedMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionPinnedMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExistingPins) {
		toSerialize["existingPins"] = o.ExistingPins
	}
	if !IsNil(o.EligiblePins) {
		toSerialize["eligiblePins"] = o.EligiblePins
	}
	toSerialize["pinnedCategories"] = o.PinnedCategories
	toSerialize["eligibleCategoriesForPinning"] = o.EligibleCategoriesForPinning
	return toSerialize, nil
}

type NullableCollectionPinnedMetadata struct {
	value *CollectionPinnedMetadata
	isSet bool
}

func (v NullableCollectionPinnedMetadata) Get() *CollectionPinnedMetadata {
	return v.value
}

func (v *NullableCollectionPinnedMetadata) Set(val *CollectionPinnedMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionPinnedMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionPinnedMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionPinnedMetadata(val *CollectionPinnedMetadata) *NullableCollectionPinnedMetadata {
	return &NullableCollectionPinnedMetadata{value: val, isSet: true}
}

func (v NullableCollectionPinnedMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionPinnedMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


