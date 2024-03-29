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

// checks if the PropertyGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PropertyGroup{}

// PropertyGroup A grouping for multiple PropertyDefinition. Grouped properties will be displayed together in the UI.
type PropertyGroup struct {
	// The unique identifier of the group.
	Name *string `json:"name,omitempty"`
	// The user-friendly group label to display.
	DisplayLabel *string `json:"displayLabel,omitempty"`
}

// NewPropertyGroup instantiates a new PropertyGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPropertyGroup() *PropertyGroup {
	this := PropertyGroup{}
	return &this
}

// NewPropertyGroupWithDefaults instantiates a new PropertyGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPropertyGroupWithDefaults() *PropertyGroup {
	this := PropertyGroup{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PropertyGroup) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyGroup) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PropertyGroup) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PropertyGroup) SetName(v string) {
	o.Name = &v
}

// GetDisplayLabel returns the DisplayLabel field value if set, zero value otherwise.
func (o *PropertyGroup) GetDisplayLabel() string {
	if o == nil || IsNil(o.DisplayLabel) {
		var ret string
		return ret
	}
	return *o.DisplayLabel
}

// GetDisplayLabelOk returns a tuple with the DisplayLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyGroup) GetDisplayLabelOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayLabel) {
		return nil, false
	}
	return o.DisplayLabel, true
}

// HasDisplayLabel returns a boolean if a field has been set.
func (o *PropertyGroup) HasDisplayLabel() bool {
	if o != nil && !IsNil(o.DisplayLabel) {
		return true
	}

	return false
}

// SetDisplayLabel gets a reference to the given string and assigns it to the DisplayLabel field.
func (o *PropertyGroup) SetDisplayLabel(v string) {
	o.DisplayLabel = &v
}

func (o PropertyGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PropertyGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.DisplayLabel) {
		toSerialize["displayLabel"] = o.DisplayLabel
	}
	return toSerialize, nil
}

type NullablePropertyGroup struct {
	value *PropertyGroup
	isSet bool
}

func (v NullablePropertyGroup) Get() *PropertyGroup {
	return v.value
}

func (v *NullablePropertyGroup) Set(val *PropertyGroup) {
	v.value = val
	v.isSet = true
}

func (v NullablePropertyGroup) IsSet() bool {
	return v.isSet
}

func (v *NullablePropertyGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePropertyGroup(val *PropertyGroup) *NullablePropertyGroup {
	return &NullablePropertyGroup{value: val, isSet: true}
}

func (v NullablePropertyGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePropertyGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


