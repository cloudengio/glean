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

// checks if the CollectionBaseMutableProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionBaseMutableProperties{}

// CollectionBaseMutableProperties struct for CollectionBaseMutableProperties
type CollectionBaseMutableProperties struct {
	// The unique name of the Collection.
	Name string `json:"name"`
	// A brief summary of the Collection's contents.
	Description *string `json:"description,omitempty"`
	// A list of added user roles for the collection.
	AddedRoles []UserRoleSpecification `json:"addedRoles,omitempty"`
	// A list of removed user roles for the collection.
	RemovedRoles []UserRoleSpecification `json:"removedRoles,omitempty"`
	// Filters which restrict who should see this collection. Values are taken from the corresponding filters in people search.
	AudienceFilters []FacetFilter `json:"audienceFilters,omitempty"`
}

// NewCollectionBaseMutableProperties instantiates a new CollectionBaseMutableProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionBaseMutableProperties(name string) *CollectionBaseMutableProperties {
	this := CollectionBaseMutableProperties{}
	this.Name = name
	return &this
}

// NewCollectionBaseMutablePropertiesWithDefaults instantiates a new CollectionBaseMutableProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionBaseMutablePropertiesWithDefaults() *CollectionBaseMutableProperties {
	this := CollectionBaseMutableProperties{}
	return &this
}

// GetName returns the Name field value
func (o *CollectionBaseMutableProperties) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CollectionBaseMutableProperties) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CollectionBaseMutableProperties) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CollectionBaseMutableProperties) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionBaseMutableProperties) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CollectionBaseMutableProperties) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CollectionBaseMutableProperties) SetDescription(v string) {
	o.Description = &v
}

// GetAddedRoles returns the AddedRoles field value if set, zero value otherwise.
func (o *CollectionBaseMutableProperties) GetAddedRoles() []UserRoleSpecification {
	if o == nil || IsNil(o.AddedRoles) {
		var ret []UserRoleSpecification
		return ret
	}
	return o.AddedRoles
}

// GetAddedRolesOk returns a tuple with the AddedRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionBaseMutableProperties) GetAddedRolesOk() ([]UserRoleSpecification, bool) {
	if o == nil || IsNil(o.AddedRoles) {
		return nil, false
	}
	return o.AddedRoles, true
}

// HasAddedRoles returns a boolean if a field has been set.
func (o *CollectionBaseMutableProperties) HasAddedRoles() bool {
	if o != nil && !IsNil(o.AddedRoles) {
		return true
	}

	return false
}

// SetAddedRoles gets a reference to the given []UserRoleSpecification and assigns it to the AddedRoles field.
func (o *CollectionBaseMutableProperties) SetAddedRoles(v []UserRoleSpecification) {
	o.AddedRoles = v
}

// GetRemovedRoles returns the RemovedRoles field value if set, zero value otherwise.
func (o *CollectionBaseMutableProperties) GetRemovedRoles() []UserRoleSpecification {
	if o == nil || IsNil(o.RemovedRoles) {
		var ret []UserRoleSpecification
		return ret
	}
	return o.RemovedRoles
}

// GetRemovedRolesOk returns a tuple with the RemovedRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionBaseMutableProperties) GetRemovedRolesOk() ([]UserRoleSpecification, bool) {
	if o == nil || IsNil(o.RemovedRoles) {
		return nil, false
	}
	return o.RemovedRoles, true
}

// HasRemovedRoles returns a boolean if a field has been set.
func (o *CollectionBaseMutableProperties) HasRemovedRoles() bool {
	if o != nil && !IsNil(o.RemovedRoles) {
		return true
	}

	return false
}

// SetRemovedRoles gets a reference to the given []UserRoleSpecification and assigns it to the RemovedRoles field.
func (o *CollectionBaseMutableProperties) SetRemovedRoles(v []UserRoleSpecification) {
	o.RemovedRoles = v
}

// GetAudienceFilters returns the AudienceFilters field value if set, zero value otherwise.
func (o *CollectionBaseMutableProperties) GetAudienceFilters() []FacetFilter {
	if o == nil || IsNil(o.AudienceFilters) {
		var ret []FacetFilter
		return ret
	}
	return o.AudienceFilters
}

// GetAudienceFiltersOk returns a tuple with the AudienceFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionBaseMutableProperties) GetAudienceFiltersOk() ([]FacetFilter, bool) {
	if o == nil || IsNil(o.AudienceFilters) {
		return nil, false
	}
	return o.AudienceFilters, true
}

// HasAudienceFilters returns a boolean if a field has been set.
func (o *CollectionBaseMutableProperties) HasAudienceFilters() bool {
	if o != nil && !IsNil(o.AudienceFilters) {
		return true
	}

	return false
}

// SetAudienceFilters gets a reference to the given []FacetFilter and assigns it to the AudienceFilters field.
func (o *CollectionBaseMutableProperties) SetAudienceFilters(v []FacetFilter) {
	o.AudienceFilters = v
}

func (o CollectionBaseMutableProperties) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionBaseMutableProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.AddedRoles) {
		toSerialize["addedRoles"] = o.AddedRoles
	}
	if !IsNil(o.RemovedRoles) {
		toSerialize["removedRoles"] = o.RemovedRoles
	}
	if !IsNil(o.AudienceFilters) {
		toSerialize["audienceFilters"] = o.AudienceFilters
	}
	return toSerialize, nil
}

type NullableCollectionBaseMutableProperties struct {
	value *CollectionBaseMutableProperties
	isSet bool
}

func (v NullableCollectionBaseMutableProperties) Get() *CollectionBaseMutableProperties {
	return v.value
}

func (v *NullableCollectionBaseMutableProperties) Set(val *CollectionBaseMutableProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionBaseMutableProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionBaseMutableProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionBaseMutableProperties(val *CollectionBaseMutableProperties) *NullableCollectionBaseMutableProperties {
	return &NullableCollectionBaseMutableProperties{value: val, isSet: true}
}

func (v NullableCollectionBaseMutableProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionBaseMutableProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


