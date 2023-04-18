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
	"time"
)

// checks if the ShortcutMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShortcutMetadata{}

// ShortcutMetadata struct for ShortcutMetadata
type ShortcutMetadata struct {
	CreatedBy *Person `json:"createdBy,omitempty"`
	// The time the shortcut was created in ISO format (ISO 8601).
	CreateTime *time.Time `json:"createTime,omitempty"`
	UpdatedBy *Person `json:"updatedBy,omitempty"`
	// The time the shortcut was updated in ISO format (ISO 8601).
	UpdateTime *time.Time `json:"updateTime,omitempty"`
	DestinationDocument *Document `json:"destinationDocument,omitempty"`
	// The URL from which the user is then redirected to the destination URL. Full replacement for https://go/<inputAlias>.
	IntermediateUrl *string `json:"intermediateUrl,omitempty"`
	// The part of the shortcut preceding the input alias when used for showing shortcuts to users. Should end with \"/\". e.g. \"go/\" for native shortcuts.
	ViewPrefix *string `json:"viewPrefix,omitempty"`
	// Indicates whether a shortcut is native or external.
	IsExternal *bool `json:"isExternal,omitempty"`
	// The URL using which the user can access the edit page of the shortcut.
	EditUrl *string `json:"editUrl,omitempty"`
}

// NewShortcutMetadata instantiates a new ShortcutMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShortcutMetadata() *ShortcutMetadata {
	this := ShortcutMetadata{}
	return &this
}

// NewShortcutMetadataWithDefaults instantiates a new ShortcutMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShortcutMetadataWithDefaults() *ShortcutMetadata {
	this := ShortcutMetadata{}
	return &this
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *ShortcutMetadata) GetCreatedBy() Person {
	if o == nil || IsNil(o.CreatedBy) {
		var ret Person
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShortcutMetadata) GetCreatedByOk() (*Person, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *ShortcutMetadata) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given Person and assigns it to the CreatedBy field.
func (o *ShortcutMetadata) SetCreatedBy(v Person) {
	o.CreatedBy = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *ShortcutMetadata) GetCreateTime() time.Time {
	if o == nil || IsNil(o.CreateTime) {
		var ret time.Time
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShortcutMetadata) GetCreateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *ShortcutMetadata) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given time.Time and assigns it to the CreateTime field.
func (o *ShortcutMetadata) SetCreateTime(v time.Time) {
	o.CreateTime = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *ShortcutMetadata) GetUpdatedBy() Person {
	if o == nil || IsNil(o.UpdatedBy) {
		var ret Person
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShortcutMetadata) GetUpdatedByOk() (*Person, bool) {
	if o == nil || IsNil(o.UpdatedBy) {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *ShortcutMetadata) HasUpdatedBy() bool {
	if o != nil && !IsNil(o.UpdatedBy) {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given Person and assigns it to the UpdatedBy field.
func (o *ShortcutMetadata) SetUpdatedBy(v Person) {
	o.UpdatedBy = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *ShortcutMetadata) GetUpdateTime() time.Time {
	if o == nil || IsNil(o.UpdateTime) {
		var ret time.Time
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShortcutMetadata) GetUpdateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *ShortcutMetadata) HasUpdateTime() bool {
	if o != nil && !IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given time.Time and assigns it to the UpdateTime field.
func (o *ShortcutMetadata) SetUpdateTime(v time.Time) {
	o.UpdateTime = &v
}

// GetDestinationDocument returns the DestinationDocument field value if set, zero value otherwise.
func (o *ShortcutMetadata) GetDestinationDocument() Document {
	if o == nil || IsNil(o.DestinationDocument) {
		var ret Document
		return ret
	}
	return *o.DestinationDocument
}

// GetDestinationDocumentOk returns a tuple with the DestinationDocument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShortcutMetadata) GetDestinationDocumentOk() (*Document, bool) {
	if o == nil || IsNil(o.DestinationDocument) {
		return nil, false
	}
	return o.DestinationDocument, true
}

// HasDestinationDocument returns a boolean if a field has been set.
func (o *ShortcutMetadata) HasDestinationDocument() bool {
	if o != nil && !IsNil(o.DestinationDocument) {
		return true
	}

	return false
}

// SetDestinationDocument gets a reference to the given Document and assigns it to the DestinationDocument field.
func (o *ShortcutMetadata) SetDestinationDocument(v Document) {
	o.DestinationDocument = &v
}

// GetIntermediateUrl returns the IntermediateUrl field value if set, zero value otherwise.
func (o *ShortcutMetadata) GetIntermediateUrl() string {
	if o == nil || IsNil(o.IntermediateUrl) {
		var ret string
		return ret
	}
	return *o.IntermediateUrl
}

// GetIntermediateUrlOk returns a tuple with the IntermediateUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShortcutMetadata) GetIntermediateUrlOk() (*string, bool) {
	if o == nil || IsNil(o.IntermediateUrl) {
		return nil, false
	}
	return o.IntermediateUrl, true
}

// HasIntermediateUrl returns a boolean if a field has been set.
func (o *ShortcutMetadata) HasIntermediateUrl() bool {
	if o != nil && !IsNil(o.IntermediateUrl) {
		return true
	}

	return false
}

// SetIntermediateUrl gets a reference to the given string and assigns it to the IntermediateUrl field.
func (o *ShortcutMetadata) SetIntermediateUrl(v string) {
	o.IntermediateUrl = &v
}

// GetViewPrefix returns the ViewPrefix field value if set, zero value otherwise.
func (o *ShortcutMetadata) GetViewPrefix() string {
	if o == nil || IsNil(o.ViewPrefix) {
		var ret string
		return ret
	}
	return *o.ViewPrefix
}

// GetViewPrefixOk returns a tuple with the ViewPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShortcutMetadata) GetViewPrefixOk() (*string, bool) {
	if o == nil || IsNil(o.ViewPrefix) {
		return nil, false
	}
	return o.ViewPrefix, true
}

// HasViewPrefix returns a boolean if a field has been set.
func (o *ShortcutMetadata) HasViewPrefix() bool {
	if o != nil && !IsNil(o.ViewPrefix) {
		return true
	}

	return false
}

// SetViewPrefix gets a reference to the given string and assigns it to the ViewPrefix field.
func (o *ShortcutMetadata) SetViewPrefix(v string) {
	o.ViewPrefix = &v
}

// GetIsExternal returns the IsExternal field value if set, zero value otherwise.
func (o *ShortcutMetadata) GetIsExternal() bool {
	if o == nil || IsNil(o.IsExternal) {
		var ret bool
		return ret
	}
	return *o.IsExternal
}

// GetIsExternalOk returns a tuple with the IsExternal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShortcutMetadata) GetIsExternalOk() (*bool, bool) {
	if o == nil || IsNil(o.IsExternal) {
		return nil, false
	}
	return o.IsExternal, true
}

// HasIsExternal returns a boolean if a field has been set.
func (o *ShortcutMetadata) HasIsExternal() bool {
	if o != nil && !IsNil(o.IsExternal) {
		return true
	}

	return false
}

// SetIsExternal gets a reference to the given bool and assigns it to the IsExternal field.
func (o *ShortcutMetadata) SetIsExternal(v bool) {
	o.IsExternal = &v
}

// GetEditUrl returns the EditUrl field value if set, zero value otherwise.
func (o *ShortcutMetadata) GetEditUrl() string {
	if o == nil || IsNil(o.EditUrl) {
		var ret string
		return ret
	}
	return *o.EditUrl
}

// GetEditUrlOk returns a tuple with the EditUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShortcutMetadata) GetEditUrlOk() (*string, bool) {
	if o == nil || IsNil(o.EditUrl) {
		return nil, false
	}
	return o.EditUrl, true
}

// HasEditUrl returns a boolean if a field has been set.
func (o *ShortcutMetadata) HasEditUrl() bool {
	if o != nil && !IsNil(o.EditUrl) {
		return true
	}

	return false
}

// SetEditUrl gets a reference to the given string and assigns it to the EditUrl field.
func (o *ShortcutMetadata) SetEditUrl(v string) {
	o.EditUrl = &v
}

func (o ShortcutMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShortcutMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedBy) {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if !IsNil(o.CreateTime) {
		toSerialize["createTime"] = o.CreateTime
	}
	if !IsNil(o.UpdatedBy) {
		toSerialize["updatedBy"] = o.UpdatedBy
	}
	if !IsNil(o.UpdateTime) {
		toSerialize["updateTime"] = o.UpdateTime
	}
	if !IsNil(o.DestinationDocument) {
		toSerialize["destinationDocument"] = o.DestinationDocument
	}
	if !IsNil(o.IntermediateUrl) {
		toSerialize["intermediateUrl"] = o.IntermediateUrl
	}
	if !IsNil(o.ViewPrefix) {
		toSerialize["viewPrefix"] = o.ViewPrefix
	}
	if !IsNil(o.IsExternal) {
		toSerialize["isExternal"] = o.IsExternal
	}
	if !IsNil(o.EditUrl) {
		toSerialize["editUrl"] = o.EditUrl
	}
	return toSerialize, nil
}

type NullableShortcutMetadata struct {
	value *ShortcutMetadata
	isSet bool
}

func (v NullableShortcutMetadata) Get() *ShortcutMetadata {
	return v.value
}

func (v *NullableShortcutMetadata) Set(val *ShortcutMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableShortcutMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableShortcutMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShortcutMetadata(val *ShortcutMetadata) *NullableShortcutMetadata {
	return &NullableShortcutMetadata{value: val, isSet: true}
}

func (v NullableShortcutMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShortcutMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

