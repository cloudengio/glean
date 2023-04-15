# NavigationMenuPanel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MenuItems** | [**[]NavigationMenuItem**](NavigationMenuItem.md) | Collection of (top-level menu items | 
**Hidden** | Pointer to **bool** | Indicates this panel itself need not be directly rendered or accessible (workspace/user settings) | [optional] 

## Methods

### NewNavigationMenuPanel

`func NewNavigationMenuPanel(menuItems []NavigationMenuItem, ) *NavigationMenuPanel`

NewNavigationMenuPanel instantiates a new NavigationMenuPanel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNavigationMenuPanelWithDefaults

`func NewNavigationMenuPanelWithDefaults() *NavigationMenuPanel`

NewNavigationMenuPanelWithDefaults instantiates a new NavigationMenuPanel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMenuItems

`func (o *NavigationMenuPanel) GetMenuItems() []NavigationMenuItem`

GetMenuItems returns the MenuItems field if non-nil, zero value otherwise.

### GetMenuItemsOk

`func (o *NavigationMenuPanel) GetMenuItemsOk() (*[]NavigationMenuItem, bool)`

GetMenuItemsOk returns a tuple with the MenuItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMenuItems

`func (o *NavigationMenuPanel) SetMenuItems(v []NavigationMenuItem)`

SetMenuItems sets MenuItems field to given value.


### GetHidden

`func (o *NavigationMenuPanel) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *NavigationMenuPanel) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *NavigationMenuPanel) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *NavigationMenuPanel) HasHidden() bool`

HasHidden returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


