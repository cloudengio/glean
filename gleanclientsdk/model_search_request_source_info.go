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

// checks if the SearchRequestSourceInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchRequestSourceInfo{}

// SearchRequestSourceInfo struct for SearchRequestSourceInfo
type SearchRequestSourceInfo struct {
	// The means by which the search request was initiated. EVAL - Internal usage for automated evaluation FACETS - A background request to get bucket counts for facets following a user search request MORE - The infinite scroll requested more results for an existing search ONBOARDING - A demo query was issued as part of the onboarding flow ONBOARDING_CHECKLIST - The user performed a search from the search step of the new user onboarding checklist PAGE_LOAD - A SERP was visited without going through regular product UI, e.g. from a bookmark, page refresh, or hitting the browser back button DISCARDED_PAGE_LOAD - as PAGE_LOAD but the page was previously discarded by the browser PREFETCH - Results for a non-active tab were requested, e.g. gmail USER - The user initiated a search by manually typing a query, clicking a suggestion, etc. RECOMMENDATION - A query intent is detected from the user's activity and a search request is issued proactively AUTOMATION - Initiated from an API used for automation by an external developer or integration.
	Initiator *string `json:"initiator,omitempty"`
	// The UI paradigm from which the search request was sent. FULLPAGE - The standard web app (including mobile) OVERLAY - An iframe that's not Embedded Search / NSR (No such frame type as of now) OMNIBOX - The browser omnibox CONTEXT_MENU - The browser right-click context menu (powered by the browser extension) EMBEDDED_SEARCH - The embedded search added as an iframe NSR - Native search replacement provided by extension injected iframe SIDEBAR - The extension sidebar GLEANBOT - Gleanbot in Slack, MS Teams, Discord, etc.
	Modality string `json:"modality"`
	// The domain from/on behalf of which the request is being issued. Currently only being used for tracking / logging purposes.
	Domain *string `json:"domain,omitempty"`
	// Platform from which the search request was sent. Optional field.
	Platform *string `json:"platform,omitempty"`
	// The (optional) UI element within the paradigm from which the search request was sent. Each modality will have a dedicated uiElement enum (e.g., SearchRequestGleanbotUIElementEnum)
	UiElement *string `json:"uiElement,omitempty"`
	// Whether the query is for debugging purposes and, as such, should not be included in usage metrics and quality pipelines.
	IsDebug *bool `json:"isDebug,omitempty"`
	// An opaque version identifier for the client. This is meant to be used for logging and debugging purposes only.
	ClientVersion *string `json:"clientVersion,omitempty"`
}

// NewSearchRequestSourceInfo instantiates a new SearchRequestSourceInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchRequestSourceInfo(modality string) *SearchRequestSourceInfo {
	this := SearchRequestSourceInfo{}
	this.Modality = modality
	return &this
}

// NewSearchRequestSourceInfoWithDefaults instantiates a new SearchRequestSourceInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchRequestSourceInfoWithDefaults() *SearchRequestSourceInfo {
	this := SearchRequestSourceInfo{}
	return &this
}

// GetInitiator returns the Initiator field value if set, zero value otherwise.
func (o *SearchRequestSourceInfo) GetInitiator() string {
	if o == nil || IsNil(o.Initiator) {
		var ret string
		return ret
	}
	return *o.Initiator
}

// GetInitiatorOk returns a tuple with the Initiator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchRequestSourceInfo) GetInitiatorOk() (*string, bool) {
	if o == nil || IsNil(o.Initiator) {
		return nil, false
	}
	return o.Initiator, true
}

// HasInitiator returns a boolean if a field has been set.
func (o *SearchRequestSourceInfo) HasInitiator() bool {
	if o != nil && !IsNil(o.Initiator) {
		return true
	}

	return false
}

// SetInitiator gets a reference to the given string and assigns it to the Initiator field.
func (o *SearchRequestSourceInfo) SetInitiator(v string) {
	o.Initiator = &v
}

// GetModality returns the Modality field value
func (o *SearchRequestSourceInfo) GetModality() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Modality
}

// GetModalityOk returns a tuple with the Modality field value
// and a boolean to check if the value has been set.
func (o *SearchRequestSourceInfo) GetModalityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Modality, true
}

// SetModality sets field value
func (o *SearchRequestSourceInfo) SetModality(v string) {
	o.Modality = v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *SearchRequestSourceInfo) GetDomain() string {
	if o == nil || IsNil(o.Domain) {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchRequestSourceInfo) GetDomainOk() (*string, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *SearchRequestSourceInfo) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *SearchRequestSourceInfo) SetDomain(v string) {
	o.Domain = &v
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *SearchRequestSourceInfo) GetPlatform() string {
	if o == nil || IsNil(o.Platform) {
		var ret string
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchRequestSourceInfo) GetPlatformOk() (*string, bool) {
	if o == nil || IsNil(o.Platform) {
		return nil, false
	}
	return o.Platform, true
}

// HasPlatform returns a boolean if a field has been set.
func (o *SearchRequestSourceInfo) HasPlatform() bool {
	if o != nil && !IsNil(o.Platform) {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given string and assigns it to the Platform field.
func (o *SearchRequestSourceInfo) SetPlatform(v string) {
	o.Platform = &v
}

// GetUiElement returns the UiElement field value if set, zero value otherwise.
func (o *SearchRequestSourceInfo) GetUiElement() string {
	if o == nil || IsNil(o.UiElement) {
		var ret string
		return ret
	}
	return *o.UiElement
}

// GetUiElementOk returns a tuple with the UiElement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchRequestSourceInfo) GetUiElementOk() (*string, bool) {
	if o == nil || IsNil(o.UiElement) {
		return nil, false
	}
	return o.UiElement, true
}

// HasUiElement returns a boolean if a field has been set.
func (o *SearchRequestSourceInfo) HasUiElement() bool {
	if o != nil && !IsNil(o.UiElement) {
		return true
	}

	return false
}

// SetUiElement gets a reference to the given string and assigns it to the UiElement field.
func (o *SearchRequestSourceInfo) SetUiElement(v string) {
	o.UiElement = &v
}

// GetIsDebug returns the IsDebug field value if set, zero value otherwise.
func (o *SearchRequestSourceInfo) GetIsDebug() bool {
	if o == nil || IsNil(o.IsDebug) {
		var ret bool
		return ret
	}
	return *o.IsDebug
}

// GetIsDebugOk returns a tuple with the IsDebug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchRequestSourceInfo) GetIsDebugOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDebug) {
		return nil, false
	}
	return o.IsDebug, true
}

// HasIsDebug returns a boolean if a field has been set.
func (o *SearchRequestSourceInfo) HasIsDebug() bool {
	if o != nil && !IsNil(o.IsDebug) {
		return true
	}

	return false
}

// SetIsDebug gets a reference to the given bool and assigns it to the IsDebug field.
func (o *SearchRequestSourceInfo) SetIsDebug(v bool) {
	o.IsDebug = &v
}

// GetClientVersion returns the ClientVersion field value if set, zero value otherwise.
func (o *SearchRequestSourceInfo) GetClientVersion() string {
	if o == nil || IsNil(o.ClientVersion) {
		var ret string
		return ret
	}
	return *o.ClientVersion
}

// GetClientVersionOk returns a tuple with the ClientVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchRequestSourceInfo) GetClientVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ClientVersion) {
		return nil, false
	}
	return o.ClientVersion, true
}

// HasClientVersion returns a boolean if a field has been set.
func (o *SearchRequestSourceInfo) HasClientVersion() bool {
	if o != nil && !IsNil(o.ClientVersion) {
		return true
	}

	return false
}

// SetClientVersion gets a reference to the given string and assigns it to the ClientVersion field.
func (o *SearchRequestSourceInfo) SetClientVersion(v string) {
	o.ClientVersion = &v
}

func (o SearchRequestSourceInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchRequestSourceInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Initiator) {
		toSerialize["initiator"] = o.Initiator
	}
	toSerialize["modality"] = o.Modality
	if !IsNil(o.Domain) {
		toSerialize["domain"] = o.Domain
	}
	if !IsNil(o.Platform) {
		toSerialize["platform"] = o.Platform
	}
	if !IsNil(o.UiElement) {
		toSerialize["uiElement"] = o.UiElement
	}
	if !IsNil(o.IsDebug) {
		toSerialize["isDebug"] = o.IsDebug
	}
	if !IsNil(o.ClientVersion) {
		toSerialize["clientVersion"] = o.ClientVersion
	}
	return toSerialize, nil
}

type NullableSearchRequestSourceInfo struct {
	value *SearchRequestSourceInfo
	isSet bool
}

func (v NullableSearchRequestSourceInfo) Get() *SearchRequestSourceInfo {
	return v.value
}

func (v *NullableSearchRequestSourceInfo) Set(val *SearchRequestSourceInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchRequestSourceInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchRequestSourceInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchRequestSourceInfo(val *SearchRequestSourceInfo) *NullableSearchRequestSourceInfo {
	return &NullableSearchRequestSourceInfo{value: val, isSet: true}
}

func (v NullableSearchRequestSourceInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchRequestSourceInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


