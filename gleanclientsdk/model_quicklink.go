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

// checks if the Quicklink type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Quicklink{}

// Quicklink An action for a specific datasource that will show up in autocomplete and app card, e.g. \"Create new issue\" for jira
type Quicklink struct {
	// Full action name. Used in autocomplete
	Name *string `json:"name,omitempty"`
	// Shortened name. Used in app card
	ShortName *string `json:"shortName,omitempty"`
	// The URL for the action
	Url *string `json:"url,omitempty"`
	IconConfig *IconConfig `json:"iconConfig,omitempty"`
	// Unique identifier of this quicklink
	Id *string `json:"id,omitempty"`
	// The scopes for which this quicklink is applicable
	Scopes []string `json:"scopes,omitempty"`
}

// NewQuicklink instantiates a new Quicklink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuicklink() *Quicklink {
	this := Quicklink{}
	return &this
}

// NewQuicklinkWithDefaults instantiates a new Quicklink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuicklinkWithDefaults() *Quicklink {
	this := Quicklink{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Quicklink) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Quicklink) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Quicklink) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Quicklink) SetName(v string) {
	o.Name = &v
}

// GetShortName returns the ShortName field value if set, zero value otherwise.
func (o *Quicklink) GetShortName() string {
	if o == nil || IsNil(o.ShortName) {
		var ret string
		return ret
	}
	return *o.ShortName
}

// GetShortNameOk returns a tuple with the ShortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Quicklink) GetShortNameOk() (*string, bool) {
	if o == nil || IsNil(o.ShortName) {
		return nil, false
	}
	return o.ShortName, true
}

// HasShortName returns a boolean if a field has been set.
func (o *Quicklink) HasShortName() bool {
	if o != nil && !IsNil(o.ShortName) {
		return true
	}

	return false
}

// SetShortName gets a reference to the given string and assigns it to the ShortName field.
func (o *Quicklink) SetShortName(v string) {
	o.ShortName = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *Quicklink) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Quicklink) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *Quicklink) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *Quicklink) SetUrl(v string) {
	o.Url = &v
}

// GetIconConfig returns the IconConfig field value if set, zero value otherwise.
func (o *Quicklink) GetIconConfig() IconConfig {
	if o == nil || IsNil(o.IconConfig) {
		var ret IconConfig
		return ret
	}
	return *o.IconConfig
}

// GetIconConfigOk returns a tuple with the IconConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Quicklink) GetIconConfigOk() (*IconConfig, bool) {
	if o == nil || IsNil(o.IconConfig) {
		return nil, false
	}
	return o.IconConfig, true
}

// HasIconConfig returns a boolean if a field has been set.
func (o *Quicklink) HasIconConfig() bool {
	if o != nil && !IsNil(o.IconConfig) {
		return true
	}

	return false
}

// SetIconConfig gets a reference to the given IconConfig and assigns it to the IconConfig field.
func (o *Quicklink) SetIconConfig(v IconConfig) {
	o.IconConfig = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Quicklink) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Quicklink) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Quicklink) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Quicklink) SetId(v string) {
	o.Id = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *Quicklink) GetScopes() []string {
	if o == nil || IsNil(o.Scopes) {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Quicklink) GetScopesOk() ([]string, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *Quicklink) HasScopes() bool {
	if o != nil && !IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *Quicklink) SetScopes(v []string) {
	o.Scopes = v
}

func (o Quicklink) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Quicklink) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ShortName) {
		toSerialize["shortName"] = o.ShortName
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.IconConfig) {
		toSerialize["iconConfig"] = o.IconConfig
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Scopes) {
		toSerialize["scopes"] = o.Scopes
	}
	return toSerialize, nil
}

type NullableQuicklink struct {
	value *Quicklink
	isSet bool
}

func (v NullableQuicklink) Get() *Quicklink {
	return v.value
}

func (v *NullableQuicklink) Set(val *Quicklink) {
	v.value = val
	v.isSet = true
}

func (v NullableQuicklink) IsSet() bool {
	return v.isSet
}

func (v *NullableQuicklink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuicklink(val *Quicklink) *NullableQuicklink {
	return &NullableQuicklink{value: val, isSet: true}
}

func (v NullableQuicklink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuicklink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

