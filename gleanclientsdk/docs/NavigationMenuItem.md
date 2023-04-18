# NavigationMenuItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Icon** | [**IconConfig**](IconConfig.md) |  | 
**SelectedIcon** | [**IconConfig**](IconConfig.md) |  | 
**MenuTitle** | **string** | Title to use for the menu item | 
**SectionTitle** | Pointer to **string** | Title to show on the open section (default to menu title if empty) | [optional] 
**DefaultLink** | [**NavigationLink**](NavigationLink.md) |  | 
**LinkGroups** | [**[]NavigationLinkGroup**](NavigationLinkGroup.md) | Links within the group | 
**ClientSideChecks** | Pointer to [**[]ClientSideCheck**](ClientSideCheck.md) | Client side checks to be performed before render | [optional] 

## Methods

### NewNavigationMenuItem

`func NewNavigationMenuItem(icon IconConfig, selectedIcon IconConfig, menuTitle string, defaultLink NavigationLink, linkGroups []NavigationLinkGroup, ) *NavigationMenuItem`

NewNavigationMenuItem instantiates a new NavigationMenuItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNavigationMenuItemWithDefaults

`func NewNavigationMenuItemWithDefaults() *NavigationMenuItem`

NewNavigationMenuItemWithDefaults instantiates a new NavigationMenuItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIcon

`func (o *NavigationMenuItem) GetIcon() IconConfig`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *NavigationMenuItem) GetIconOk() (*IconConfig, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *NavigationMenuItem) SetIcon(v IconConfig)`

SetIcon sets Icon field to given value.


### GetSelectedIcon

`func (o *NavigationMenuItem) GetSelectedIcon() IconConfig`

GetSelectedIcon returns the SelectedIcon field if non-nil, zero value otherwise.

### GetSelectedIconOk

`func (o *NavigationMenuItem) GetSelectedIconOk() (*IconConfig, bool)`

GetSelectedIconOk returns a tuple with the SelectedIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedIcon

`func (o *NavigationMenuItem) SetSelectedIcon(v IconConfig)`

SetSelectedIcon sets SelectedIcon field to given value.


### GetMenuTitle

`func (o *NavigationMenuItem) GetMenuTitle() string`

GetMenuTitle returns the MenuTitle field if non-nil, zero value otherwise.

### GetMenuTitleOk

`func (o *NavigationMenuItem) GetMenuTitleOk() (*string, bool)`

GetMenuTitleOk returns a tuple with the MenuTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMenuTitle

`func (o *NavigationMenuItem) SetMenuTitle(v string)`

SetMenuTitle sets MenuTitle field to given value.


### GetSectionTitle

`func (o *NavigationMenuItem) GetSectionTitle() string`

GetSectionTitle returns the SectionTitle field if non-nil, zero value otherwise.

### GetSectionTitleOk

`func (o *NavigationMenuItem) GetSectionTitleOk() (*string, bool)`

GetSectionTitleOk returns a tuple with the SectionTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionTitle

`func (o *NavigationMenuItem) SetSectionTitle(v string)`

SetSectionTitle sets SectionTitle field to given value.

### HasSectionTitle

`func (o *NavigationMenuItem) HasSectionTitle() bool`

HasSectionTitle returns a boolean if a field has been set.

### GetDefaultLink

`func (o *NavigationMenuItem) GetDefaultLink() NavigationLink`

GetDefaultLink returns the DefaultLink field if non-nil, zero value otherwise.

### GetDefaultLinkOk

`func (o *NavigationMenuItem) GetDefaultLinkOk() (*NavigationLink, bool)`

GetDefaultLinkOk returns a tuple with the DefaultLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultLink

`func (o *NavigationMenuItem) SetDefaultLink(v NavigationLink)`

SetDefaultLink sets DefaultLink field to given value.


### GetLinkGroups

`func (o *NavigationMenuItem) GetLinkGroups() []NavigationLinkGroup`

GetLinkGroups returns the LinkGroups field if non-nil, zero value otherwise.

### GetLinkGroupsOk

`func (o *NavigationMenuItem) GetLinkGroupsOk() (*[]NavigationLinkGroup, bool)`

GetLinkGroupsOk returns a tuple with the LinkGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkGroups

`func (o *NavigationMenuItem) SetLinkGroups(v []NavigationLinkGroup)`

SetLinkGroups sets LinkGroups field to given value.


### GetClientSideChecks

`func (o *NavigationMenuItem) GetClientSideChecks() []ClientSideCheck`

GetClientSideChecks returns the ClientSideChecks field if non-nil, zero value otherwise.

### GetClientSideChecksOk

`func (o *NavigationMenuItem) GetClientSideChecksOk() (*[]ClientSideCheck, bool)`

GetClientSideChecksOk returns a tuple with the ClientSideChecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSideChecks

`func (o *NavigationMenuItem) SetClientSideChecks(v []ClientSideCheck)`

SetClientSideChecks sets ClientSideChecks field to given value.

### HasClientSideChecks

`func (o *NavigationMenuItem) HasClientSideChecks() bool`

HasClientSideChecks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


