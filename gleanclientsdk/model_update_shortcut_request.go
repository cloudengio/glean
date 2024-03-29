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

// checks if the UpdateShortcutRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateShortcutRequest{}

// UpdateShortcutRequest struct for UpdateShortcutRequest
type UpdateShortcutRequest struct {
	// The opaque id of the user generated content.
	Id int32 `json:"id"`
	// link text following go/ prefix as entered by the user.
	InputAlias *string `json:"inputAlias,omitempty"`
	// destination URL for the shortcut.
	DestinationUrl *string `json:"destinationUrl,omitempty"`
	// document id for the url, if known.
	DestinationDocumentId *string `json:"destinationDocumentId,omitempty"`
	// A short, plain text blurb to help people understand the intent of the shortcut.
	Description *string `json:"description,omitempty"`
	// Whether this shortcut is unlisted or not. Unlisted shortcuts are visible to author + admins only.
	Unlisted *bool `json:"unlisted,omitempty"`
	// For variable shortcuts, contains the url template; note, destinationUrl contains default url
	UrlTemplate *string `json:"urlTemplate,omitempty"`
	// A list of user roles added for the Shortcut.
	AddedRoles []UserRoleSpecification `json:"addedRoles,omitempty"`
	// A list of user roles removed for the Shortcut.
	RemovedRoles []UserRoleSpecification `json:"removedRoles,omitempty"`
}

// NewUpdateShortcutRequest instantiates a new UpdateShortcutRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateShortcutRequest(id int32) *UpdateShortcutRequest {
	this := UpdateShortcutRequest{}
	this.Id = id
	return &this
}

// NewUpdateShortcutRequestWithDefaults instantiates a new UpdateShortcutRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateShortcutRequestWithDefaults() *UpdateShortcutRequest {
	this := UpdateShortcutRequest{}
	return &this
}

// GetId returns the Id field value
func (o *UpdateShortcutRequest) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UpdateShortcutRequest) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UpdateShortcutRequest) SetId(v int32) {
	o.Id = v
}

// GetInputAlias returns the InputAlias field value if set, zero value otherwise.
func (o *UpdateShortcutRequest) GetInputAlias() string {
	if o == nil || IsNil(o.InputAlias) {
		var ret string
		return ret
	}
	return *o.InputAlias
}

// GetInputAliasOk returns a tuple with the InputAlias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateShortcutRequest) GetInputAliasOk() (*string, bool) {
	if o == nil || IsNil(o.InputAlias) {
		return nil, false
	}
	return o.InputAlias, true
}

// HasInputAlias returns a boolean if a field has been set.
func (o *UpdateShortcutRequest) HasInputAlias() bool {
	if o != nil && !IsNil(o.InputAlias) {
		return true
	}

	return false
}

// SetInputAlias gets a reference to the given string and assigns it to the InputAlias field.
func (o *UpdateShortcutRequest) SetInputAlias(v string) {
	o.InputAlias = &v
}

// GetDestinationUrl returns the DestinationUrl field value if set, zero value otherwise.
func (o *UpdateShortcutRequest) GetDestinationUrl() string {
	if o == nil || IsNil(o.DestinationUrl) {
		var ret string
		return ret
	}
	return *o.DestinationUrl
}

// GetDestinationUrlOk returns a tuple with the DestinationUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateShortcutRequest) GetDestinationUrlOk() (*string, bool) {
	if o == nil || IsNil(o.DestinationUrl) {
		return nil, false
	}
	return o.DestinationUrl, true
}

// HasDestinationUrl returns a boolean if a field has been set.
func (o *UpdateShortcutRequest) HasDestinationUrl() bool {
	if o != nil && !IsNil(o.DestinationUrl) {
		return true
	}

	return false
}

// SetDestinationUrl gets a reference to the given string and assigns it to the DestinationUrl field.
func (o *UpdateShortcutRequest) SetDestinationUrl(v string) {
	o.DestinationUrl = &v
}

// GetDestinationDocumentId returns the DestinationDocumentId field value if set, zero value otherwise.
func (o *UpdateShortcutRequest) GetDestinationDocumentId() string {
	if o == nil || IsNil(o.DestinationDocumentId) {
		var ret string
		return ret
	}
	return *o.DestinationDocumentId
}

// GetDestinationDocumentIdOk returns a tuple with the DestinationDocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateShortcutRequest) GetDestinationDocumentIdOk() (*string, bool) {
	if o == nil || IsNil(o.DestinationDocumentId) {
		return nil, false
	}
	return o.DestinationDocumentId, true
}

// HasDestinationDocumentId returns a boolean if a field has been set.
func (o *UpdateShortcutRequest) HasDestinationDocumentId() bool {
	if o != nil && !IsNil(o.DestinationDocumentId) {
		return true
	}

	return false
}

// SetDestinationDocumentId gets a reference to the given string and assigns it to the DestinationDocumentId field.
func (o *UpdateShortcutRequest) SetDestinationDocumentId(v string) {
	o.DestinationDocumentId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateShortcutRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateShortcutRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateShortcutRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateShortcutRequest) SetDescription(v string) {
	o.Description = &v
}

// GetUnlisted returns the Unlisted field value if set, zero value otherwise.
func (o *UpdateShortcutRequest) GetUnlisted() bool {
	if o == nil || IsNil(o.Unlisted) {
		var ret bool
		return ret
	}
	return *o.Unlisted
}

// GetUnlistedOk returns a tuple with the Unlisted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateShortcutRequest) GetUnlistedOk() (*bool, bool) {
	if o == nil || IsNil(o.Unlisted) {
		return nil, false
	}
	return o.Unlisted, true
}

// HasUnlisted returns a boolean if a field has been set.
func (o *UpdateShortcutRequest) HasUnlisted() bool {
	if o != nil && !IsNil(o.Unlisted) {
		return true
	}

	return false
}

// SetUnlisted gets a reference to the given bool and assigns it to the Unlisted field.
func (o *UpdateShortcutRequest) SetUnlisted(v bool) {
	o.Unlisted = &v
}

// GetUrlTemplate returns the UrlTemplate field value if set, zero value otherwise.
func (o *UpdateShortcutRequest) GetUrlTemplate() string {
	if o == nil || IsNil(o.UrlTemplate) {
		var ret string
		return ret
	}
	return *o.UrlTemplate
}

// GetUrlTemplateOk returns a tuple with the UrlTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateShortcutRequest) GetUrlTemplateOk() (*string, bool) {
	if o == nil || IsNil(o.UrlTemplate) {
		return nil, false
	}
	return o.UrlTemplate, true
}

// HasUrlTemplate returns a boolean if a field has been set.
func (o *UpdateShortcutRequest) HasUrlTemplate() bool {
	if o != nil && !IsNil(o.UrlTemplate) {
		return true
	}

	return false
}

// SetUrlTemplate gets a reference to the given string and assigns it to the UrlTemplate field.
func (o *UpdateShortcutRequest) SetUrlTemplate(v string) {
	o.UrlTemplate = &v
}

// GetAddedRoles returns the AddedRoles field value if set, zero value otherwise.
func (o *UpdateShortcutRequest) GetAddedRoles() []UserRoleSpecification {
	if o == nil || IsNil(o.AddedRoles) {
		var ret []UserRoleSpecification
		return ret
	}
	return o.AddedRoles
}

// GetAddedRolesOk returns a tuple with the AddedRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateShortcutRequest) GetAddedRolesOk() ([]UserRoleSpecification, bool) {
	if o == nil || IsNil(o.AddedRoles) {
		return nil, false
	}
	return o.AddedRoles, true
}

// HasAddedRoles returns a boolean if a field has been set.
func (o *UpdateShortcutRequest) HasAddedRoles() bool {
	if o != nil && !IsNil(o.AddedRoles) {
		return true
	}

	return false
}

// SetAddedRoles gets a reference to the given []UserRoleSpecification and assigns it to the AddedRoles field.
func (o *UpdateShortcutRequest) SetAddedRoles(v []UserRoleSpecification) {
	o.AddedRoles = v
}

// GetRemovedRoles returns the RemovedRoles field value if set, zero value otherwise.
func (o *UpdateShortcutRequest) GetRemovedRoles() []UserRoleSpecification {
	if o == nil || IsNil(o.RemovedRoles) {
		var ret []UserRoleSpecification
		return ret
	}
	return o.RemovedRoles
}

// GetRemovedRolesOk returns a tuple with the RemovedRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateShortcutRequest) GetRemovedRolesOk() ([]UserRoleSpecification, bool) {
	if o == nil || IsNil(o.RemovedRoles) {
		return nil, false
	}
	return o.RemovedRoles, true
}

// HasRemovedRoles returns a boolean if a field has been set.
func (o *UpdateShortcutRequest) HasRemovedRoles() bool {
	if o != nil && !IsNil(o.RemovedRoles) {
		return true
	}

	return false
}

// SetRemovedRoles gets a reference to the given []UserRoleSpecification and assigns it to the RemovedRoles field.
func (o *UpdateShortcutRequest) SetRemovedRoles(v []UserRoleSpecification) {
	o.RemovedRoles = v
}

func (o UpdateShortcutRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateShortcutRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.InputAlias) {
		toSerialize["inputAlias"] = o.InputAlias
	}
	if !IsNil(o.DestinationUrl) {
		toSerialize["destinationUrl"] = o.DestinationUrl
	}
	if !IsNil(o.DestinationDocumentId) {
		toSerialize["destinationDocumentId"] = o.DestinationDocumentId
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Unlisted) {
		toSerialize["unlisted"] = o.Unlisted
	}
	if !IsNil(o.UrlTemplate) {
		toSerialize["urlTemplate"] = o.UrlTemplate
	}
	if !IsNil(o.AddedRoles) {
		toSerialize["addedRoles"] = o.AddedRoles
	}
	if !IsNil(o.RemovedRoles) {
		toSerialize["removedRoles"] = o.RemovedRoles
	}
	return toSerialize, nil
}

type NullableUpdateShortcutRequest struct {
	value *UpdateShortcutRequest
	isSet bool
}

func (v NullableUpdateShortcutRequest) Get() *UpdateShortcutRequest {
	return v.value
}

func (v *NullableUpdateShortcutRequest) Set(val *UpdateShortcutRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateShortcutRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateShortcutRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateShortcutRequest(val *UpdateShortcutRequest) *NullableUpdateShortcutRequest {
	return &NullableUpdateShortcutRequest{value: val, isSet: true}
}

func (v NullableUpdateShortcutRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateShortcutRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


