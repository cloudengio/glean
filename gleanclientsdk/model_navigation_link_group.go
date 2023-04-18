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

// checks if the NavigationLinkGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NavigationLinkGroup{}

// NavigationLinkGroup A link group is a collection of navigation links, with an optional default panel link target if a group link is provided, the links are displayed as second level links otherwise, the group is separated from other section via line separators
type NavigationLinkGroup struct {
	Title *string `json:"title,omitempty"`
	DefaultLink *NavigationLink `json:"defaultLink,omitempty"`
	// Links within the group
	Links []NavigationLink `json:"links"`
	// Behavior hints such as whether the group should be separated.  May or may not be used by the client during rendering.
	Behaviors []NavigationLinkGroupBehavior `json:"behaviors,omitempty"`
}

// NewNavigationLinkGroup instantiates a new NavigationLinkGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNavigationLinkGroup(links []NavigationLink) *NavigationLinkGroup {
	this := NavigationLinkGroup{}
	this.Links = links
	return &this
}

// NewNavigationLinkGroupWithDefaults instantiates a new NavigationLinkGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNavigationLinkGroupWithDefaults() *NavigationLinkGroup {
	this := NavigationLinkGroup{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *NavigationLinkGroup) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NavigationLinkGroup) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *NavigationLinkGroup) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *NavigationLinkGroup) SetTitle(v string) {
	o.Title = &v
}

// GetDefaultLink returns the DefaultLink field value if set, zero value otherwise.
func (o *NavigationLinkGroup) GetDefaultLink() NavigationLink {
	if o == nil || IsNil(o.DefaultLink) {
		var ret NavigationLink
		return ret
	}
	return *o.DefaultLink
}

// GetDefaultLinkOk returns a tuple with the DefaultLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NavigationLinkGroup) GetDefaultLinkOk() (*NavigationLink, bool) {
	if o == nil || IsNil(o.DefaultLink) {
		return nil, false
	}
	return o.DefaultLink, true
}

// HasDefaultLink returns a boolean if a field has been set.
func (o *NavigationLinkGroup) HasDefaultLink() bool {
	if o != nil && !IsNil(o.DefaultLink) {
		return true
	}

	return false
}

// SetDefaultLink gets a reference to the given NavigationLink and assigns it to the DefaultLink field.
func (o *NavigationLinkGroup) SetDefaultLink(v NavigationLink) {
	o.DefaultLink = &v
}

// GetLinks returns the Links field value
func (o *NavigationLinkGroup) GetLinks() []NavigationLink {
	if o == nil {
		var ret []NavigationLink
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *NavigationLinkGroup) GetLinksOk() ([]NavigationLink, bool) {
	if o == nil {
		return nil, false
	}
	return o.Links, true
}

// SetLinks sets field value
func (o *NavigationLinkGroup) SetLinks(v []NavigationLink) {
	o.Links = v
}

// GetBehaviors returns the Behaviors field value if set, zero value otherwise.
func (o *NavigationLinkGroup) GetBehaviors() []NavigationLinkGroupBehavior {
	if o == nil || IsNil(o.Behaviors) {
		var ret []NavigationLinkGroupBehavior
		return ret
	}
	return o.Behaviors
}

// GetBehaviorsOk returns a tuple with the Behaviors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NavigationLinkGroup) GetBehaviorsOk() ([]NavigationLinkGroupBehavior, bool) {
	if o == nil || IsNil(o.Behaviors) {
		return nil, false
	}
	return o.Behaviors, true
}

// HasBehaviors returns a boolean if a field has been set.
func (o *NavigationLinkGroup) HasBehaviors() bool {
	if o != nil && !IsNil(o.Behaviors) {
		return true
	}

	return false
}

// SetBehaviors gets a reference to the given []NavigationLinkGroupBehavior and assigns it to the Behaviors field.
func (o *NavigationLinkGroup) SetBehaviors(v []NavigationLinkGroupBehavior) {
	o.Behaviors = v
}

func (o NavigationLinkGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NavigationLinkGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.DefaultLink) {
		toSerialize["defaultLink"] = o.DefaultLink
	}
	toSerialize["links"] = o.Links
	if !IsNil(o.Behaviors) {
		toSerialize["behaviors"] = o.Behaviors
	}
	return toSerialize, nil
}

type NullableNavigationLinkGroup struct {
	value *NavigationLinkGroup
	isSet bool
}

func (v NullableNavigationLinkGroup) Get() *NavigationLinkGroup {
	return v.value
}

func (v *NullableNavigationLinkGroup) Set(val *NavigationLinkGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableNavigationLinkGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableNavigationLinkGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNavigationLinkGroup(val *NavigationLinkGroup) *NullableNavigationLinkGroup {
	return &NullableNavigationLinkGroup{value: val, isSet: true}
}

func (v NullableNavigationLinkGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNavigationLinkGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


