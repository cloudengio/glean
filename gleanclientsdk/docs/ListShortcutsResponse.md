# ListShortcutsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Shortcuts** | [**[]Shortcut**](Shortcut.md) | List of all shortcuts accessible to the user | 

## Methods

### NewListShortcutsResponse

`func NewListShortcutsResponse(shortcuts []Shortcut, ) *ListShortcutsResponse`

NewListShortcutsResponse instantiates a new ListShortcutsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListShortcutsResponseWithDefaults

`func NewListShortcutsResponseWithDefaults() *ListShortcutsResponse`

NewListShortcutsResponseWithDefaults instantiates a new ListShortcutsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShortcuts

`func (o *ListShortcutsResponse) GetShortcuts() []Shortcut`

GetShortcuts returns the Shortcuts field if non-nil, zero value otherwise.

### GetShortcutsOk

`func (o *ListShortcutsResponse) GetShortcutsOk() (*[]Shortcut, bool)`

GetShortcutsOk returns a tuple with the Shortcuts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortcuts

`func (o *ListShortcutsResponse) SetShortcuts(v []Shortcut)`

SetShortcuts sets Shortcuts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


