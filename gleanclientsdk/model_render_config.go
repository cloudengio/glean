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

// checks if the RenderConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RenderConfig{}

// RenderConfig struct for RenderConfig
type RenderConfig struct {
	Actions *ActionsConfig `json:"actions,omitempty"`
	AdditionalConfigs []string `json:"additionalConfigs,omitempty"`
	Body *BodyConfig `json:"body,omitempty"`
	Contact *ContactConfig `json:"contact,omitempty"`
	Details *DetailsConfig `json:"details,omitempty"`
	FilePath *FilePathConfig `json:"filePath,omitempty"`
	// Deprecated
	Footer *RenderConfigFooter `json:"footer,omitempty"`
	Related *RelatedConfig `json:"related,omitempty"`
	Icon []CompositeIconConfig `json:"icon,omitempty"`
	Interactions *InteractionsConfig `json:"interactions,omitempty"`
	Meta *MetaConfig `json:"meta,omitempty"`
	Preview *PreviewConfig `json:"preview,omitempty"`
	Tags *TagsConfig `json:"tags,omitempty"`
	Title *TitleConfig `json:"title,omitempty"`
	Sections *SectionsConfig `json:"sections,omitempty"`
}

// NewRenderConfig instantiates a new RenderConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRenderConfig() *RenderConfig {
	this := RenderConfig{}
	return &this
}

// NewRenderConfigWithDefaults instantiates a new RenderConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRenderConfigWithDefaults() *RenderConfig {
	this := RenderConfig{}
	return &this
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *RenderConfig) GetActions() ActionsConfig {
	if o == nil || IsNil(o.Actions) {
		var ret ActionsConfig
		return ret
	}
	return *o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenderConfig) GetActionsOk() (*ActionsConfig, bool) {
	if o == nil || IsNil(o.Actions) {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *RenderConfig) HasActions() bool {
	if o != nil && !IsNil(o.Actions) {
		return true
	}

	return false
}

// SetActions gets a reference to the given ActionsConfig and assigns it to the Actions field.
func (o *RenderConfig) SetActions(v ActionsConfig) {
	o.Actions = &v
}

// GetAdditionalConfigs returns the AdditionalConfigs field value if set, zero value otherwise.
func (o *RenderConfig) GetAdditionalConfigs() []string {
	if o == nil || IsNil(o.AdditionalConfigs) {
		var ret []string
		return ret
	}
	return o.AdditionalConfigs
}

// GetAdditionalConfigsOk returns a tuple with the AdditionalConfigs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenderConfig) GetAdditionalConfigsOk() ([]string, bool) {
	if o == nil || IsNil(o.AdditionalConfigs) {
		return nil, false
	}
	return o.AdditionalConfigs, true
}

// HasAdditionalConfigs returns a boolean if a field has been set.
func (o *RenderConfig) HasAdditionalConfigs() bool {
	if o != nil && !IsNil(o.AdditionalConfigs) {
		return true
	}

	return false
}

// SetAdditionalConfigs gets a reference to the given []string and assigns it to the AdditionalConfigs field.
func (o *RenderConfig) SetAdditionalConfigs(v []string) {
	o.AdditionalConfigs = v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *RenderConfig) GetBody() BodyConfig {
	if o == nil || IsNil(o.Body) {
		var ret BodyConfig
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenderConfig) GetBodyOk() (*BodyConfig, bool) {
	if o == nil || IsNil(o.Body) {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *RenderConfig) HasBody() bool {
	if o != nil && !IsNil(o.Body) {
		return true
	}

	return false
}

// SetBody gets a reference to the given BodyConfig and assigns it to the Body field.
func (o *RenderConfig) SetBody(v BodyConfig) {
	o.Body = &v
}

// GetContact returns the Contact field value if set, zero value otherwise.
func (o *RenderConfig) GetContact() ContactConfig {
	if o == nil || IsNil(o.Contact) {
		var ret ContactConfig
		return ret
	}
	return *o.Contact
}

// GetContactOk returns a tuple with the Contact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenderConfig) GetContactOk() (*ContactConfig, bool) {
	if o == nil || IsNil(o.Contact) {
		return nil, false
	}
	return o.Contact, true
}

// HasContact returns a boolean if a field has been set.
func (o *RenderConfig) HasContact() bool {
	if o != nil && !IsNil(o.Contact) {
		return true
	}

	return false
}

// SetContact gets a reference to the given ContactConfig and assigns it to the Contact field.
func (o *RenderConfig) SetContact(v ContactConfig) {
	o.Contact = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *RenderConfig) GetDetails() DetailsConfig {
	if o == nil || IsNil(o.Details) {
		var ret DetailsConfig
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenderConfig) GetDetailsOk() (*DetailsConfig, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *RenderConfig) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given DetailsConfig and assigns it to the Details field.
func (o *RenderConfig) SetDetails(v DetailsConfig) {
	o.Details = &v
}

// GetFilePath returns the FilePath field value if set, zero value otherwise.
func (o *RenderConfig) GetFilePath() FilePathConfig {
	if o == nil || IsNil(o.FilePath) {
		var ret FilePathConfig
		return ret
	}
	return *o.FilePath
}

// GetFilePathOk returns a tuple with the FilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenderConfig) GetFilePathOk() (*FilePathConfig, bool) {
	if o == nil || IsNil(o.FilePath) {
		return nil, false
	}
	return o.FilePath, true
}

// HasFilePath returns a boolean if a field has been set.
func (o *RenderConfig) HasFilePath() bool {
	if o != nil && !IsNil(o.FilePath) {
		return true
	}

	return false
}

// SetFilePath gets a reference to the given FilePathConfig and assigns it to the FilePath field.
func (o *RenderConfig) SetFilePath(v FilePathConfig) {
	o.FilePath = &v
}

// GetFooter returns the Footer field value if set, zero value otherwise.
// Deprecated
func (o *RenderConfig) GetFooter() RenderConfigFooter {
	if o == nil || IsNil(o.Footer) {
		var ret RenderConfigFooter
		return ret
	}
	return *o.Footer
}

// GetFooterOk returns a tuple with the Footer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *RenderConfig) GetFooterOk() (*RenderConfigFooter, bool) {
	if o == nil || IsNil(o.Footer) {
		return nil, false
	}
	return o.Footer, true
}

// HasFooter returns a boolean if a field has been set.
func (o *RenderConfig) HasFooter() bool {
	if o != nil && !IsNil(o.Footer) {
		return true
	}

	return false
}

// SetFooter gets a reference to the given RenderConfigFooter and assigns it to the Footer field.
// Deprecated
func (o *RenderConfig) SetFooter(v RenderConfigFooter) {
	o.Footer = &v
}

// GetRelated returns the Related field value if set, zero value otherwise.
func (o *RenderConfig) GetRelated() RelatedConfig {
	if o == nil || IsNil(o.Related) {
		var ret RelatedConfig
		return ret
	}
	return *o.Related
}

// GetRelatedOk returns a tuple with the Related field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenderConfig) GetRelatedOk() (*RelatedConfig, bool) {
	if o == nil || IsNil(o.Related) {
		return nil, false
	}
	return o.Related, true
}

// HasRelated returns a boolean if a field has been set.
func (o *RenderConfig) HasRelated() bool {
	if o != nil && !IsNil(o.Related) {
		return true
	}

	return false
}

// SetRelated gets a reference to the given RelatedConfig and assigns it to the Related field.
func (o *RenderConfig) SetRelated(v RelatedConfig) {
	o.Related = &v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *RenderConfig) GetIcon() []CompositeIconConfig {
	if o == nil || IsNil(o.Icon) {
		var ret []CompositeIconConfig
		return ret
	}
	return o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenderConfig) GetIconOk() ([]CompositeIconConfig, bool) {
	if o == nil || IsNil(o.Icon) {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *RenderConfig) HasIcon() bool {
	if o != nil && !IsNil(o.Icon) {
		return true
	}

	return false
}

// SetIcon gets a reference to the given []CompositeIconConfig and assigns it to the Icon field.
func (o *RenderConfig) SetIcon(v []CompositeIconConfig) {
	o.Icon = v
}

// GetInteractions returns the Interactions field value if set, zero value otherwise.
func (o *RenderConfig) GetInteractions() InteractionsConfig {
	if o == nil || IsNil(o.Interactions) {
		var ret InteractionsConfig
		return ret
	}
	return *o.Interactions
}

// GetInteractionsOk returns a tuple with the Interactions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenderConfig) GetInteractionsOk() (*InteractionsConfig, bool) {
	if o == nil || IsNil(o.Interactions) {
		return nil, false
	}
	return o.Interactions, true
}

// HasInteractions returns a boolean if a field has been set.
func (o *RenderConfig) HasInteractions() bool {
	if o != nil && !IsNil(o.Interactions) {
		return true
	}

	return false
}

// SetInteractions gets a reference to the given InteractionsConfig and assigns it to the Interactions field.
func (o *RenderConfig) SetInteractions(v InteractionsConfig) {
	o.Interactions = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *RenderConfig) GetMeta() MetaConfig {
	if o == nil || IsNil(o.Meta) {
		var ret MetaConfig
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenderConfig) GetMetaOk() (*MetaConfig, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *RenderConfig) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaConfig and assigns it to the Meta field.
func (o *RenderConfig) SetMeta(v MetaConfig) {
	o.Meta = &v
}

// GetPreview returns the Preview field value if set, zero value otherwise.
func (o *RenderConfig) GetPreview() PreviewConfig {
	if o == nil || IsNil(o.Preview) {
		var ret PreviewConfig
		return ret
	}
	return *o.Preview
}

// GetPreviewOk returns a tuple with the Preview field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenderConfig) GetPreviewOk() (*PreviewConfig, bool) {
	if o == nil || IsNil(o.Preview) {
		return nil, false
	}
	return o.Preview, true
}

// HasPreview returns a boolean if a field has been set.
func (o *RenderConfig) HasPreview() bool {
	if o != nil && !IsNil(o.Preview) {
		return true
	}

	return false
}

// SetPreview gets a reference to the given PreviewConfig and assigns it to the Preview field.
func (o *RenderConfig) SetPreview(v PreviewConfig) {
	o.Preview = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *RenderConfig) GetTags() TagsConfig {
	if o == nil || IsNil(o.Tags) {
		var ret TagsConfig
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenderConfig) GetTagsOk() (*TagsConfig, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *RenderConfig) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given TagsConfig and assigns it to the Tags field.
func (o *RenderConfig) SetTags(v TagsConfig) {
	o.Tags = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *RenderConfig) GetTitle() TitleConfig {
	if o == nil || IsNil(o.Title) {
		var ret TitleConfig
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenderConfig) GetTitleOk() (*TitleConfig, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *RenderConfig) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given TitleConfig and assigns it to the Title field.
func (o *RenderConfig) SetTitle(v TitleConfig) {
	o.Title = &v
}

// GetSections returns the Sections field value if set, zero value otherwise.
func (o *RenderConfig) GetSections() SectionsConfig {
	if o == nil || IsNil(o.Sections) {
		var ret SectionsConfig
		return ret
	}
	return *o.Sections
}

// GetSectionsOk returns a tuple with the Sections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenderConfig) GetSectionsOk() (*SectionsConfig, bool) {
	if o == nil || IsNil(o.Sections) {
		return nil, false
	}
	return o.Sections, true
}

// HasSections returns a boolean if a field has been set.
func (o *RenderConfig) HasSections() bool {
	if o != nil && !IsNil(o.Sections) {
		return true
	}

	return false
}

// SetSections gets a reference to the given SectionsConfig and assigns it to the Sections field.
func (o *RenderConfig) SetSections(v SectionsConfig) {
	o.Sections = &v
}

func (o RenderConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RenderConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Actions) {
		toSerialize["actions"] = o.Actions
	}
	if !IsNil(o.AdditionalConfigs) {
		toSerialize["additionalConfigs"] = o.AdditionalConfigs
	}
	if !IsNil(o.Body) {
		toSerialize["body"] = o.Body
	}
	if !IsNil(o.Contact) {
		toSerialize["contact"] = o.Contact
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	if !IsNil(o.FilePath) {
		toSerialize["filePath"] = o.FilePath
	}
	if !IsNil(o.Footer) {
		toSerialize["footer"] = o.Footer
	}
	if !IsNil(o.Related) {
		toSerialize["related"] = o.Related
	}
	if !IsNil(o.Icon) {
		toSerialize["icon"] = o.Icon
	}
	if !IsNil(o.Interactions) {
		toSerialize["interactions"] = o.Interactions
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Preview) {
		toSerialize["preview"] = o.Preview
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Sections) {
		toSerialize["sections"] = o.Sections
	}
	return toSerialize, nil
}

type NullableRenderConfig struct {
	value *RenderConfig
	isSet bool
}

func (v NullableRenderConfig) Get() *RenderConfig {
	return v.value
}

func (v *NullableRenderConfig) Set(val *RenderConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableRenderConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableRenderConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRenderConfig(val *RenderConfig) *NullableRenderConfig {
	return &NullableRenderConfig{value: val, isSet: true}
}

func (v NullableRenderConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRenderConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


