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

// checks if the NavigationLink type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NavigationLink{}

// NavigationLink Basic navigation link item.  These represent the leaves in the tree
type NavigationLink struct {
	// Title to show for the navigation link
	Title string `json:"title"`
	// Destination for the link
	Target string `json:"target"`
	// Hint indicating the link points to an external site
	External *bool `json:"external,omitempty"`
	// Client side checks to be performed before render
	ClientSideChecks []ClientSideCheck `json:"clientSideChecks,omitempty"`
}

// NewNavigationLink instantiates a new NavigationLink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNavigationLink(title string, target string) *NavigationLink {
	this := NavigationLink{}
	this.Title = title
	this.Target = target
	return &this
}

// NewNavigationLinkWithDefaults instantiates a new NavigationLink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNavigationLinkWithDefaults() *NavigationLink {
	this := NavigationLink{}
	return &this
}

// GetTitle returns the Title field value
func (o *NavigationLink) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *NavigationLink) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *NavigationLink) SetTitle(v string) {
	o.Title = v
}

// GetTarget returns the Target field value
func (o *NavigationLink) GetTarget() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *NavigationLink) GetTargetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value
func (o *NavigationLink) SetTarget(v string) {
	o.Target = v
}

// GetExternal returns the External field value if set, zero value otherwise.
func (o *NavigationLink) GetExternal() bool {
	if o == nil || IsNil(o.External) {
		var ret bool
		return ret
	}
	return *o.External
}

// GetExternalOk returns a tuple with the External field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NavigationLink) GetExternalOk() (*bool, bool) {
	if o == nil || IsNil(o.External) {
		return nil, false
	}
	return o.External, true
}

// HasExternal returns a boolean if a field has been set.
func (o *NavigationLink) HasExternal() bool {
	if o != nil && !IsNil(o.External) {
		return true
	}

	return false
}

// SetExternal gets a reference to the given bool and assigns it to the External field.
func (o *NavigationLink) SetExternal(v bool) {
	o.External = &v
}

// GetClientSideChecks returns the ClientSideChecks field value if set, zero value otherwise.
func (o *NavigationLink) GetClientSideChecks() []ClientSideCheck {
	if o == nil || IsNil(o.ClientSideChecks) {
		var ret []ClientSideCheck
		return ret
	}
	return o.ClientSideChecks
}

// GetClientSideChecksOk returns a tuple with the ClientSideChecks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NavigationLink) GetClientSideChecksOk() ([]ClientSideCheck, bool) {
	if o == nil || IsNil(o.ClientSideChecks) {
		return nil, false
	}
	return o.ClientSideChecks, true
}

// HasClientSideChecks returns a boolean if a field has been set.
func (o *NavigationLink) HasClientSideChecks() bool {
	if o != nil && !IsNil(o.ClientSideChecks) {
		return true
	}

	return false
}

// SetClientSideChecks gets a reference to the given []ClientSideCheck and assigns it to the ClientSideChecks field.
func (o *NavigationLink) SetClientSideChecks(v []ClientSideCheck) {
	o.ClientSideChecks = v
}

func (o NavigationLink) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NavigationLink) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["title"] = o.Title
	toSerialize["target"] = o.Target
	if !IsNil(o.External) {
		toSerialize["external"] = o.External
	}
	if !IsNil(o.ClientSideChecks) {
		toSerialize["clientSideChecks"] = o.ClientSideChecks
	}
	return toSerialize, nil
}

type NullableNavigationLink struct {
	value *NavigationLink
	isSet bool
}

func (v NullableNavigationLink) Get() *NavigationLink {
	return v.value
}

func (v *NullableNavigationLink) Set(val *NavigationLink) {
	v.value = val
	v.isSet = true
}

func (v NullableNavigationLink) IsSet() bool {
	return v.isSet
}

func (v *NullableNavigationLink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNavigationLink(val *NavigationLink) *NullableNavigationLink {
	return &NullableNavigationLink{value: val, isSet: true}
}

func (v NullableNavigationLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNavigationLink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

