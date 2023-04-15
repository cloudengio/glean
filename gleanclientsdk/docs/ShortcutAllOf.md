# ShortcutAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alias** | Pointer to **string** | canonical link text following go/ prefix where hyphen/underscore is removed. | [optional] 
**Title** | Pointer to **string** | Title for the Go Link | [optional] 
**Roles** | Pointer to [**[]UserRoleSpecification**](UserRoleSpecification.md) | A list of user roles for the Go Link. | [optional] 

## Methods

### NewShortcutAllOf

`func NewShortcutAllOf() *ShortcutAllOf`

NewShortcutAllOf instantiates a new ShortcutAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShortcutAllOfWithDefaults

`func NewShortcutAllOfWithDefaults() *ShortcutAllOf`

NewShortcutAllOfWithDefaults instantiates a new ShortcutAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlias

`func (o *ShortcutAllOf) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *ShortcutAllOf) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *ShortcutAllOf) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *ShortcutAllOf) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetTitle

`func (o *ShortcutAllOf) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ShortcutAllOf) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ShortcutAllOf) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ShortcutAllOf) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetRoles

`func (o *ShortcutAllOf) GetRoles() []UserRoleSpecification`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *ShortcutAllOf) GetRolesOk() (*[]UserRoleSpecification, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *ShortcutAllOf) SetRoles(v []UserRoleSpecification)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *ShortcutAllOf) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


