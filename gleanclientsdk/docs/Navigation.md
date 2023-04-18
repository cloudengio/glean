# Navigation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Panels** | [**[]NavigationMenuPanel**](NavigationMenuPanel.md) | Contains a collection of (top-level) menu links | 

## Methods

### NewNavigation

`func NewNavigation(panels []NavigationMenuPanel, ) *Navigation`

NewNavigation instantiates a new Navigation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNavigationWithDefaults

`func NewNavigationWithDefaults() *Navigation`

NewNavigationWithDefaults instantiates a new Navigation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPanels

`func (o *Navigation) GetPanels() []NavigationMenuPanel`

GetPanels returns the Panels field if non-nil, zero value otherwise.

### GetPanelsOk

`func (o *Navigation) GetPanelsOk() (*[]NavigationMenuPanel, bool)`

GetPanelsOk returns a tuple with the Panels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPanels

`func (o *Navigation) SetPanels(v []NavigationMenuPanel)`

SetPanels sets Panels field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


