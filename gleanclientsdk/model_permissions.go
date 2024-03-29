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

// checks if the Permissions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Permissions{}

// Permissions Describes the permissions levels that a user has for permissioned features. When the client sends this, Permissions.read and Permissions.write are the additional permissions granted to a user on top of what they have via their roles. When the server sends this, Permissions.read and Permissions.write are the complete (merged) set of permissions the user has, and Permissions.roles is just for display purposes (along with Permisisons.isModified)
type Permissions struct {
	// TODO--deprecate in favor of the read and write properties. True if the user has access to the insights dashboard
	CanInsightsDashboard bool `json:"canInsightsDashboard"`
	// TODO--deprecate in favor of the read and write properties. True if the user has access to /adminsearch
	CanAdminSearch bool `json:"canAdminSearch"`
	// TODO--deprecate in favor of the read and write properties. True if the user can administrate client API tokens with global scope
	CanAdminClientApiGlobalTokens *bool `json:"canAdminClientApiGlobalTokens,omitempty"`
	// TODO--deprecate in favor of the read and write properties. True if the user has access to data loss prevention (DLP) features
	CanDlp *bool `json:"canDlp,omitempty"`
	// TODO--deprecate in favor of the read and write properties. Define new UGC tags.
	CreateUgcTags bool `json:"createUgcTags"`
	// TODO--deprecate in favor of the read and write properties. Add and remove existing UGC tags on any document.
	EditDocumentTags bool `json:"editDocumentTags"`
	// TODO--deprecate in favor of the read and write properties. True is the user can create/update/delete announcements.
	CanCreateAnnoucements bool `json:"canCreateAnnoucements"`
	// TODO--deprecate in favor of the read and write properties. True if the user has access to the Generated Qna feature on SERP
	CanAccessGeneratedQna *bool `json:"canAccessGeneratedQna,omitempty"`
	// Describes the read permission levels that a user has for permissioned features. Key must be PermissionedFeatureOrObject
	Read *map[string][]ReadPermission `json:"read,omitempty"`
	// Describes the write permissions levels that a user has for permissioned features. Key must be PermissionedFeatureOrObject
	Write *map[string][]WritePermission `json:"write,omitempty"`
	// Describes the grant permission levels that a user has for permissioned features. Key must be PermissionedFeatureOrObject
	Grant *map[string][]GrantPermission `json:"grant,omitempty"`
	// DEPRECATED - The role for the user
	// Deprecated
	Role *string `json:"role,omitempty"`
	// The roles a user has.
	Roles []string `json:"roles,omitempty"`
	// True if this user was given additional permissions beyond what they're granted via their permission role
	IsModified *bool `json:"isModified,omitempty"`
}

// NewPermissions instantiates a new Permissions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermissions(canInsightsDashboard bool, canAdminSearch bool, createUgcTags bool, editDocumentTags bool, canCreateAnnoucements bool) *Permissions {
	this := Permissions{}
	this.CanInsightsDashboard = canInsightsDashboard
	this.CanAdminSearch = canAdminSearch
	this.CreateUgcTags = createUgcTags
	this.EditDocumentTags = editDocumentTags
	this.CanCreateAnnoucements = canCreateAnnoucements
	return &this
}

// NewPermissionsWithDefaults instantiates a new Permissions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionsWithDefaults() *Permissions {
	this := Permissions{}
	return &this
}

// GetCanInsightsDashboard returns the CanInsightsDashboard field value
func (o *Permissions) GetCanInsightsDashboard() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CanInsightsDashboard
}

// GetCanInsightsDashboardOk returns a tuple with the CanInsightsDashboard field value
// and a boolean to check if the value has been set.
func (o *Permissions) GetCanInsightsDashboardOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CanInsightsDashboard, true
}

// SetCanInsightsDashboard sets field value
func (o *Permissions) SetCanInsightsDashboard(v bool) {
	o.CanInsightsDashboard = v
}

// GetCanAdminSearch returns the CanAdminSearch field value
func (o *Permissions) GetCanAdminSearch() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CanAdminSearch
}

// GetCanAdminSearchOk returns a tuple with the CanAdminSearch field value
// and a boolean to check if the value has been set.
func (o *Permissions) GetCanAdminSearchOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CanAdminSearch, true
}

// SetCanAdminSearch sets field value
func (o *Permissions) SetCanAdminSearch(v bool) {
	o.CanAdminSearch = v
}

// GetCanAdminClientApiGlobalTokens returns the CanAdminClientApiGlobalTokens field value if set, zero value otherwise.
func (o *Permissions) GetCanAdminClientApiGlobalTokens() bool {
	if o == nil || IsNil(o.CanAdminClientApiGlobalTokens) {
		var ret bool
		return ret
	}
	return *o.CanAdminClientApiGlobalTokens
}

// GetCanAdminClientApiGlobalTokensOk returns a tuple with the CanAdminClientApiGlobalTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Permissions) GetCanAdminClientApiGlobalTokensOk() (*bool, bool) {
	if o == nil || IsNil(o.CanAdminClientApiGlobalTokens) {
		return nil, false
	}
	return o.CanAdminClientApiGlobalTokens, true
}

// HasCanAdminClientApiGlobalTokens returns a boolean if a field has been set.
func (o *Permissions) HasCanAdminClientApiGlobalTokens() bool {
	if o != nil && !IsNil(o.CanAdminClientApiGlobalTokens) {
		return true
	}

	return false
}

// SetCanAdminClientApiGlobalTokens gets a reference to the given bool and assigns it to the CanAdminClientApiGlobalTokens field.
func (o *Permissions) SetCanAdminClientApiGlobalTokens(v bool) {
	o.CanAdminClientApiGlobalTokens = &v
}

// GetCanDlp returns the CanDlp field value if set, zero value otherwise.
func (o *Permissions) GetCanDlp() bool {
	if o == nil || IsNil(o.CanDlp) {
		var ret bool
		return ret
	}
	return *o.CanDlp
}

// GetCanDlpOk returns a tuple with the CanDlp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Permissions) GetCanDlpOk() (*bool, bool) {
	if o == nil || IsNil(o.CanDlp) {
		return nil, false
	}
	return o.CanDlp, true
}

// HasCanDlp returns a boolean if a field has been set.
func (o *Permissions) HasCanDlp() bool {
	if o != nil && !IsNil(o.CanDlp) {
		return true
	}

	return false
}

// SetCanDlp gets a reference to the given bool and assigns it to the CanDlp field.
func (o *Permissions) SetCanDlp(v bool) {
	o.CanDlp = &v
}

// GetCreateUgcTags returns the CreateUgcTags field value
func (o *Permissions) GetCreateUgcTags() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CreateUgcTags
}

// GetCreateUgcTagsOk returns a tuple with the CreateUgcTags field value
// and a boolean to check if the value has been set.
func (o *Permissions) GetCreateUgcTagsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreateUgcTags, true
}

// SetCreateUgcTags sets field value
func (o *Permissions) SetCreateUgcTags(v bool) {
	o.CreateUgcTags = v
}

// GetEditDocumentTags returns the EditDocumentTags field value
func (o *Permissions) GetEditDocumentTags() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.EditDocumentTags
}

// GetEditDocumentTagsOk returns a tuple with the EditDocumentTags field value
// and a boolean to check if the value has been set.
func (o *Permissions) GetEditDocumentTagsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EditDocumentTags, true
}

// SetEditDocumentTags sets field value
func (o *Permissions) SetEditDocumentTags(v bool) {
	o.EditDocumentTags = v
}

// GetCanCreateAnnoucements returns the CanCreateAnnoucements field value
func (o *Permissions) GetCanCreateAnnoucements() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CanCreateAnnoucements
}

// GetCanCreateAnnoucementsOk returns a tuple with the CanCreateAnnoucements field value
// and a boolean to check if the value has been set.
func (o *Permissions) GetCanCreateAnnoucementsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CanCreateAnnoucements, true
}

// SetCanCreateAnnoucements sets field value
func (o *Permissions) SetCanCreateAnnoucements(v bool) {
	o.CanCreateAnnoucements = v
}

// GetCanAccessGeneratedQna returns the CanAccessGeneratedQna field value if set, zero value otherwise.
func (o *Permissions) GetCanAccessGeneratedQna() bool {
	if o == nil || IsNil(o.CanAccessGeneratedQna) {
		var ret bool
		return ret
	}
	return *o.CanAccessGeneratedQna
}

// GetCanAccessGeneratedQnaOk returns a tuple with the CanAccessGeneratedQna field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Permissions) GetCanAccessGeneratedQnaOk() (*bool, bool) {
	if o == nil || IsNil(o.CanAccessGeneratedQna) {
		return nil, false
	}
	return o.CanAccessGeneratedQna, true
}

// HasCanAccessGeneratedQna returns a boolean if a field has been set.
func (o *Permissions) HasCanAccessGeneratedQna() bool {
	if o != nil && !IsNil(o.CanAccessGeneratedQna) {
		return true
	}

	return false
}

// SetCanAccessGeneratedQna gets a reference to the given bool and assigns it to the CanAccessGeneratedQna field.
func (o *Permissions) SetCanAccessGeneratedQna(v bool) {
	o.CanAccessGeneratedQna = &v
}

// GetRead returns the Read field value if set, zero value otherwise.
func (o *Permissions) GetRead() map[string][]ReadPermission {
	if o == nil || IsNil(o.Read) {
		var ret map[string][]ReadPermission
		return ret
	}
	return *o.Read
}

// GetReadOk returns a tuple with the Read field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Permissions) GetReadOk() (*map[string][]ReadPermission, bool) {
	if o == nil || IsNil(o.Read) {
		return nil, false
	}
	return o.Read, true
}

// HasRead returns a boolean if a field has been set.
func (o *Permissions) HasRead() bool {
	if o != nil && !IsNil(o.Read) {
		return true
	}

	return false
}

// SetRead gets a reference to the given map[string][]ReadPermission and assigns it to the Read field.
func (o *Permissions) SetRead(v map[string][]ReadPermission) {
	o.Read = &v
}

// GetWrite returns the Write field value if set, zero value otherwise.
func (o *Permissions) GetWrite() map[string][]WritePermission {
	if o == nil || IsNil(o.Write) {
		var ret map[string][]WritePermission
		return ret
	}
	return *o.Write
}

// GetWriteOk returns a tuple with the Write field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Permissions) GetWriteOk() (*map[string][]WritePermission, bool) {
	if o == nil || IsNil(o.Write) {
		return nil, false
	}
	return o.Write, true
}

// HasWrite returns a boolean if a field has been set.
func (o *Permissions) HasWrite() bool {
	if o != nil && !IsNil(o.Write) {
		return true
	}

	return false
}

// SetWrite gets a reference to the given map[string][]WritePermission and assigns it to the Write field.
func (o *Permissions) SetWrite(v map[string][]WritePermission) {
	o.Write = &v
}

// GetGrant returns the Grant field value if set, zero value otherwise.
func (o *Permissions) GetGrant() map[string][]GrantPermission {
	if o == nil || IsNil(o.Grant) {
		var ret map[string][]GrantPermission
		return ret
	}
	return *o.Grant
}

// GetGrantOk returns a tuple with the Grant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Permissions) GetGrantOk() (*map[string][]GrantPermission, bool) {
	if o == nil || IsNil(o.Grant) {
		return nil, false
	}
	return o.Grant, true
}

// HasGrant returns a boolean if a field has been set.
func (o *Permissions) HasGrant() bool {
	if o != nil && !IsNil(o.Grant) {
		return true
	}

	return false
}

// SetGrant gets a reference to the given map[string][]GrantPermission and assigns it to the Grant field.
func (o *Permissions) SetGrant(v map[string][]GrantPermission) {
	o.Grant = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
// Deprecated
func (o *Permissions) GetRole() string {
	if o == nil || IsNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *Permissions) GetRoleOk() (*string, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *Permissions) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
// Deprecated
func (o *Permissions) SetRole(v string) {
	o.Role = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *Permissions) GetRoles() []string {
	if o == nil || IsNil(o.Roles) {
		var ret []string
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Permissions) GetRolesOk() ([]string, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *Permissions) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *Permissions) SetRoles(v []string) {
	o.Roles = v
}

// GetIsModified returns the IsModified field value if set, zero value otherwise.
func (o *Permissions) GetIsModified() bool {
	if o == nil || IsNil(o.IsModified) {
		var ret bool
		return ret
	}
	return *o.IsModified
}

// GetIsModifiedOk returns a tuple with the IsModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Permissions) GetIsModifiedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsModified) {
		return nil, false
	}
	return o.IsModified, true
}

// HasIsModified returns a boolean if a field has been set.
func (o *Permissions) HasIsModified() bool {
	if o != nil && !IsNil(o.IsModified) {
		return true
	}

	return false
}

// SetIsModified gets a reference to the given bool and assigns it to the IsModified field.
func (o *Permissions) SetIsModified(v bool) {
	o.IsModified = &v
}

func (o Permissions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Permissions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["canInsightsDashboard"] = o.CanInsightsDashboard
	toSerialize["canAdminSearch"] = o.CanAdminSearch
	if !IsNil(o.CanAdminClientApiGlobalTokens) {
		toSerialize["canAdminClientApiGlobalTokens"] = o.CanAdminClientApiGlobalTokens
	}
	if !IsNil(o.CanDlp) {
		toSerialize["canDlp"] = o.CanDlp
	}
	toSerialize["createUgcTags"] = o.CreateUgcTags
	toSerialize["editDocumentTags"] = o.EditDocumentTags
	toSerialize["canCreateAnnoucements"] = o.CanCreateAnnoucements
	if !IsNil(o.CanAccessGeneratedQna) {
		toSerialize["canAccessGeneratedQna"] = o.CanAccessGeneratedQna
	}
	if !IsNil(o.Read) {
		toSerialize["read"] = o.Read
	}
	if !IsNil(o.Write) {
		toSerialize["write"] = o.Write
	}
	if !IsNil(o.Grant) {
		toSerialize["grant"] = o.Grant
	}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	if !IsNil(o.IsModified) {
		toSerialize["isModified"] = o.IsModified
	}
	return toSerialize, nil
}

type NullablePermissions struct {
	value *Permissions
	isSet bool
}

func (v NullablePermissions) Get() *Permissions {
	return v.value
}

func (v *NullablePermissions) Set(val *Permissions) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissions) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissions(val *Permissions) *NullablePermissions {
	return &NullablePermissions{value: val, isSet: true}
}

func (v NullablePermissions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


