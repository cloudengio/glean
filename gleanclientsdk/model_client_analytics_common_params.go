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

// checks if the ClientAnalyticsCommonParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClientAnalyticsCommonParams{}

// ClientAnalyticsCommonParams struct for ClientAnalyticsCommonParams
type ClientAnalyticsCommonParams struct {
	DebugMode *bool `json:"debugMode,omitempty"`
	ExtensionVersion *string `json:"extensionVersion,omitempty"`
	PageReferrer *string `json:"pageReferrer,omitempty"`
	Theme *string `json:"theme,omitempty"`
	WebAppVersion *string `json:"webAppVersion,omitempty"`
	Modality *string `json:"modality,omitempty"`
}

// NewClientAnalyticsCommonParams instantiates a new ClientAnalyticsCommonParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientAnalyticsCommonParams() *ClientAnalyticsCommonParams {
	this := ClientAnalyticsCommonParams{}
	return &this
}

// NewClientAnalyticsCommonParamsWithDefaults instantiates a new ClientAnalyticsCommonParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientAnalyticsCommonParamsWithDefaults() *ClientAnalyticsCommonParams {
	this := ClientAnalyticsCommonParams{}
	return &this
}

// GetDebugMode returns the DebugMode field value if set, zero value otherwise.
func (o *ClientAnalyticsCommonParams) GetDebugMode() bool {
	if o == nil || IsNil(o.DebugMode) {
		var ret bool
		return ret
	}
	return *o.DebugMode
}

// GetDebugModeOk returns a tuple with the DebugMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientAnalyticsCommonParams) GetDebugModeOk() (*bool, bool) {
	if o == nil || IsNil(o.DebugMode) {
		return nil, false
	}
	return o.DebugMode, true
}

// HasDebugMode returns a boolean if a field has been set.
func (o *ClientAnalyticsCommonParams) HasDebugMode() bool {
	if o != nil && !IsNil(o.DebugMode) {
		return true
	}

	return false
}

// SetDebugMode gets a reference to the given bool and assigns it to the DebugMode field.
func (o *ClientAnalyticsCommonParams) SetDebugMode(v bool) {
	o.DebugMode = &v
}

// GetExtensionVersion returns the ExtensionVersion field value if set, zero value otherwise.
func (o *ClientAnalyticsCommonParams) GetExtensionVersion() string {
	if o == nil || IsNil(o.ExtensionVersion) {
		var ret string
		return ret
	}
	return *o.ExtensionVersion
}

// GetExtensionVersionOk returns a tuple with the ExtensionVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientAnalyticsCommonParams) GetExtensionVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ExtensionVersion) {
		return nil, false
	}
	return o.ExtensionVersion, true
}

// HasExtensionVersion returns a boolean if a field has been set.
func (o *ClientAnalyticsCommonParams) HasExtensionVersion() bool {
	if o != nil && !IsNil(o.ExtensionVersion) {
		return true
	}

	return false
}

// SetExtensionVersion gets a reference to the given string and assigns it to the ExtensionVersion field.
func (o *ClientAnalyticsCommonParams) SetExtensionVersion(v string) {
	o.ExtensionVersion = &v
}

// GetPageReferrer returns the PageReferrer field value if set, zero value otherwise.
func (o *ClientAnalyticsCommonParams) GetPageReferrer() string {
	if o == nil || IsNil(o.PageReferrer) {
		var ret string
		return ret
	}
	return *o.PageReferrer
}

// GetPageReferrerOk returns a tuple with the PageReferrer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientAnalyticsCommonParams) GetPageReferrerOk() (*string, bool) {
	if o == nil || IsNil(o.PageReferrer) {
		return nil, false
	}
	return o.PageReferrer, true
}

// HasPageReferrer returns a boolean if a field has been set.
func (o *ClientAnalyticsCommonParams) HasPageReferrer() bool {
	if o != nil && !IsNil(o.PageReferrer) {
		return true
	}

	return false
}

// SetPageReferrer gets a reference to the given string and assigns it to the PageReferrer field.
func (o *ClientAnalyticsCommonParams) SetPageReferrer(v string) {
	o.PageReferrer = &v
}

// GetTheme returns the Theme field value if set, zero value otherwise.
func (o *ClientAnalyticsCommonParams) GetTheme() string {
	if o == nil || IsNil(o.Theme) {
		var ret string
		return ret
	}
	return *o.Theme
}

// GetThemeOk returns a tuple with the Theme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientAnalyticsCommonParams) GetThemeOk() (*string, bool) {
	if o == nil || IsNil(o.Theme) {
		return nil, false
	}
	return o.Theme, true
}

// HasTheme returns a boolean if a field has been set.
func (o *ClientAnalyticsCommonParams) HasTheme() bool {
	if o != nil && !IsNil(o.Theme) {
		return true
	}

	return false
}

// SetTheme gets a reference to the given string and assigns it to the Theme field.
func (o *ClientAnalyticsCommonParams) SetTheme(v string) {
	o.Theme = &v
}

// GetWebAppVersion returns the WebAppVersion field value if set, zero value otherwise.
func (o *ClientAnalyticsCommonParams) GetWebAppVersion() string {
	if o == nil || IsNil(o.WebAppVersion) {
		var ret string
		return ret
	}
	return *o.WebAppVersion
}

// GetWebAppVersionOk returns a tuple with the WebAppVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientAnalyticsCommonParams) GetWebAppVersionOk() (*string, bool) {
	if o == nil || IsNil(o.WebAppVersion) {
		return nil, false
	}
	return o.WebAppVersion, true
}

// HasWebAppVersion returns a boolean if a field has been set.
func (o *ClientAnalyticsCommonParams) HasWebAppVersion() bool {
	if o != nil && !IsNil(o.WebAppVersion) {
		return true
	}

	return false
}

// SetWebAppVersion gets a reference to the given string and assigns it to the WebAppVersion field.
func (o *ClientAnalyticsCommonParams) SetWebAppVersion(v string) {
	o.WebAppVersion = &v
}

// GetModality returns the Modality field value if set, zero value otherwise.
func (o *ClientAnalyticsCommonParams) GetModality() string {
	if o == nil || IsNil(o.Modality) {
		var ret string
		return ret
	}
	return *o.Modality
}

// GetModalityOk returns a tuple with the Modality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientAnalyticsCommonParams) GetModalityOk() (*string, bool) {
	if o == nil || IsNil(o.Modality) {
		return nil, false
	}
	return o.Modality, true
}

// HasModality returns a boolean if a field has been set.
func (o *ClientAnalyticsCommonParams) HasModality() bool {
	if o != nil && !IsNil(o.Modality) {
		return true
	}

	return false
}

// SetModality gets a reference to the given string and assigns it to the Modality field.
func (o *ClientAnalyticsCommonParams) SetModality(v string) {
	o.Modality = &v
}

func (o ClientAnalyticsCommonParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientAnalyticsCommonParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DebugMode) {
		toSerialize["debugMode"] = o.DebugMode
	}
	if !IsNil(o.ExtensionVersion) {
		toSerialize["extensionVersion"] = o.ExtensionVersion
	}
	if !IsNil(o.PageReferrer) {
		toSerialize["pageReferrer"] = o.PageReferrer
	}
	if !IsNil(o.Theme) {
		toSerialize["theme"] = o.Theme
	}
	if !IsNil(o.WebAppVersion) {
		toSerialize["webAppVersion"] = o.WebAppVersion
	}
	if !IsNil(o.Modality) {
		toSerialize["modality"] = o.Modality
	}
	return toSerialize, nil
}

type NullableClientAnalyticsCommonParams struct {
	value *ClientAnalyticsCommonParams
	isSet bool
}

func (v NullableClientAnalyticsCommonParams) Get() *ClientAnalyticsCommonParams {
	return v.value
}

func (v *NullableClientAnalyticsCommonParams) Set(val *ClientAnalyticsCommonParams) {
	v.value = val
	v.isSet = true
}

func (v NullableClientAnalyticsCommonParams) IsSet() bool {
	return v.isSet
}

func (v *NullableClientAnalyticsCommonParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientAnalyticsCommonParams(val *ClientAnalyticsCommonParams) *NullableClientAnalyticsCommonParams {
	return &NullableClientAnalyticsCommonParams{value: val, isSet: true}
}

func (v NullableClientAnalyticsCommonParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientAnalyticsCommonParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

