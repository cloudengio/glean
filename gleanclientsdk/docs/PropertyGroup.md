# PropertyGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The unique identifier of the group. | [optional] 
**DisplayLabel** | Pointer to **string** | The user-friendly group label to display. | [optional] 

## Methods

### NewPropertyGroup

`func NewPropertyGroup() *PropertyGroup`

NewPropertyGroup instantiates a new PropertyGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPropertyGroupWithDefaults

`func NewPropertyGroupWithDefaults() *PropertyGroup`

NewPropertyGroupWithDefaults instantiates a new PropertyGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PropertyGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PropertyGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PropertyGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PropertyGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayLabel

`func (o *PropertyGroup) GetDisplayLabel() string`

GetDisplayLabel returns the DisplayLabel field if non-nil, zero value otherwise.

### GetDisplayLabelOk

`func (o *PropertyGroup) GetDisplayLabelOk() (*string, bool)`

GetDisplayLabelOk returns a tuple with the DisplayLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayLabel

`func (o *PropertyGroup) SetDisplayLabel(v string)`

SetDisplayLabel sets DisplayLabel field to given value.

### HasDisplayLabel

`func (o *PropertyGroup) HasDisplayLabel() bool`

HasDisplayLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


