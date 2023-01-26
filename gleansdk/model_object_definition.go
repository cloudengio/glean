/*
Glean Indexing API

# Introduction In addition to the data sources that Glean has built-in support for, Glean also provides a REST API that enables customers to put arbitrary content in the search index. This is useful, for example, for doing permissions-aware search over content in internal tools that reside on-prem as well as for searching over applications that Glean does not currently support first class. In addition these APIs allow the customer to push organization data (people info, organization structure etc) into Glean.  # Early Access Please note that we are continually evolving the system so these APIs are considered “early access” and are subject to change. 

API version: 0.9.0
Contact: support@glean.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gleansdk

import (
	"encoding/json"
)

// ObjectDefinition The definition for an `DocumentMetadata.objectType` within a datasource.
type ObjectDefinition struct {
	// Unique identifier for this `DocumentMetadata.objectType`. If omitted, this definition is used as a default for all unmatched `DocumentMetadata.objectType`s in this datasource.
	Name *string `json:"name,omitempty"`
	// The user-friendly name of the object for display.
	DisplayLabel *string `json:"displayLabel,omitempty"`
	// The document category of this object type.
	DocCategory *string `json:"docCategory,omitempty"`
	PropertyDefinitions []PropertyDefinition `json:"propertyDefinitions,omitempty"`
	// A list of `PropertyGroup`s belonging to this object type, which will group properties to be displayed together in the UI.
	PropertyGroups []PropertyGroup `json:"propertyGroups,omitempty"`
}

// NewObjectDefinition instantiates a new ObjectDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectDefinition() *ObjectDefinition {
	this := ObjectDefinition{}
	return &this
}

// NewObjectDefinitionWithDefaults instantiates a new ObjectDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectDefinitionWithDefaults() *ObjectDefinition {
	this := ObjectDefinition{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ObjectDefinition) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectDefinition) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ObjectDefinition) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ObjectDefinition) SetName(v string) {
	o.Name = &v
}

// GetDisplayLabel returns the DisplayLabel field value if set, zero value otherwise.
func (o *ObjectDefinition) GetDisplayLabel() string {
	if o == nil || o.DisplayLabel == nil {
		var ret string
		return ret
	}
	return *o.DisplayLabel
}

// GetDisplayLabelOk returns a tuple with the DisplayLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectDefinition) GetDisplayLabelOk() (*string, bool) {
	if o == nil || o.DisplayLabel == nil {
		return nil, false
	}
	return o.DisplayLabel, true
}

// HasDisplayLabel returns a boolean if a field has been set.
func (o *ObjectDefinition) HasDisplayLabel() bool {
	if o != nil && o.DisplayLabel != nil {
		return true
	}

	return false
}

// SetDisplayLabel gets a reference to the given string and assigns it to the DisplayLabel field.
func (o *ObjectDefinition) SetDisplayLabel(v string) {
	o.DisplayLabel = &v
}

// GetDocCategory returns the DocCategory field value if set, zero value otherwise.
func (o *ObjectDefinition) GetDocCategory() string {
	if o == nil || o.DocCategory == nil {
		var ret string
		return ret
	}
	return *o.DocCategory
}

// GetDocCategoryOk returns a tuple with the DocCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectDefinition) GetDocCategoryOk() (*string, bool) {
	if o == nil || o.DocCategory == nil {
		return nil, false
	}
	return o.DocCategory, true
}

// HasDocCategory returns a boolean if a field has been set.
func (o *ObjectDefinition) HasDocCategory() bool {
	if o != nil && o.DocCategory != nil {
		return true
	}

	return false
}

// SetDocCategory gets a reference to the given string and assigns it to the DocCategory field.
func (o *ObjectDefinition) SetDocCategory(v string) {
	o.DocCategory = &v
}

// GetPropertyDefinitions returns the PropertyDefinitions field value if set, zero value otherwise.
func (o *ObjectDefinition) GetPropertyDefinitions() []PropertyDefinition {
	if o == nil || o.PropertyDefinitions == nil {
		var ret []PropertyDefinition
		return ret
	}
	return o.PropertyDefinitions
}

// GetPropertyDefinitionsOk returns a tuple with the PropertyDefinitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectDefinition) GetPropertyDefinitionsOk() ([]PropertyDefinition, bool) {
	if o == nil || o.PropertyDefinitions == nil {
		return nil, false
	}
	return o.PropertyDefinitions, true
}

// HasPropertyDefinitions returns a boolean if a field has been set.
func (o *ObjectDefinition) HasPropertyDefinitions() bool {
	if o != nil && o.PropertyDefinitions != nil {
		return true
	}

	return false
}

// SetPropertyDefinitions gets a reference to the given []PropertyDefinition and assigns it to the PropertyDefinitions field.
func (o *ObjectDefinition) SetPropertyDefinitions(v []PropertyDefinition) {
	o.PropertyDefinitions = v
}

// GetPropertyGroups returns the PropertyGroups field value if set, zero value otherwise.
func (o *ObjectDefinition) GetPropertyGroups() []PropertyGroup {
	if o == nil || o.PropertyGroups == nil {
		var ret []PropertyGroup
		return ret
	}
	return o.PropertyGroups
}

// GetPropertyGroupsOk returns a tuple with the PropertyGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectDefinition) GetPropertyGroupsOk() ([]PropertyGroup, bool) {
	if o == nil || o.PropertyGroups == nil {
		return nil, false
	}
	return o.PropertyGroups, true
}

// HasPropertyGroups returns a boolean if a field has been set.
func (o *ObjectDefinition) HasPropertyGroups() bool {
	if o != nil && o.PropertyGroups != nil {
		return true
	}

	return false
}

// SetPropertyGroups gets a reference to the given []PropertyGroup and assigns it to the PropertyGroups field.
func (o *ObjectDefinition) SetPropertyGroups(v []PropertyGroup) {
	o.PropertyGroups = v
}

func (o ObjectDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.DisplayLabel != nil {
		toSerialize["displayLabel"] = o.DisplayLabel
	}
	if o.DocCategory != nil {
		toSerialize["docCategory"] = o.DocCategory
	}
	if o.PropertyDefinitions != nil {
		toSerialize["propertyDefinitions"] = o.PropertyDefinitions
	}
	if o.PropertyGroups != nil {
		toSerialize["propertyGroups"] = o.PropertyGroups
	}
	return json.Marshal(toSerialize)
}

type NullableObjectDefinition struct {
	value *ObjectDefinition
	isSet bool
}

func (v NullableObjectDefinition) Get() *ObjectDefinition {
	return v.value
}

func (v *NullableObjectDefinition) Set(val *ObjectDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectDefinition(val *ObjectDefinition) *NullableObjectDefinition {
	return &NullableObjectDefinition{value: val, isSet: true}
}

func (v NullableObjectDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


