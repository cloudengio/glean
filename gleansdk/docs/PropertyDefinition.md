# PropertyDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the property in the &#x60;DocumentMetadata&#x60; (e.g. &#39;createTime&#39;, &#39;updateTime&#39;, &#39;author&#39;, &#39;container&#39;). In the future, this will support custom properties too. | [optional] 
**DisplayLabel** | Pointer to **string** | The user friendly label for the property. | [optional] 
**DisplayLabelPlural** | Pointer to **string** | The user friendly label for the property that will be used if a plural context. | [optional] 
**PropertyType** | Pointer to **string** | The type of custom property - this governs the search and faceting behavior | [optional] 
**UiOptions** | Pointer to **string** |  | [optional] 
**HideUiFacet** | Pointer to **bool** | If true then the property will not show up as a facet in the UI. | [optional] 
**UiFacetOrder** | Pointer to **int32** | Will be used to set the order of facets in the UI, if present. If set for one facet, must be set for all non-hidden UI facets. Must take on an integer value from 1 (shown at the top) to N (shown last), where N is the number of non-hidden UI facets. These facets will be ordered below the built-in \&quot;Type\&quot; and \&quot;Tag\&quot; operators. | [optional] 
**ObjectPropertyOptions** | Pointer to [**[]ObjectPropertyOptions**](ObjectPropertyOptions.md) |  | [optional] 
**Group** | Pointer to **string** | The unique identifier of the &#x60;PropertyGroup&#x60; to which this property belongs. | [optional] 

## Methods

### NewPropertyDefinition

`func NewPropertyDefinition() *PropertyDefinition`

NewPropertyDefinition instantiates a new PropertyDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPropertyDefinitionWithDefaults

`func NewPropertyDefinitionWithDefaults() *PropertyDefinition`

NewPropertyDefinitionWithDefaults instantiates a new PropertyDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PropertyDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PropertyDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PropertyDefinition) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PropertyDefinition) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayLabel

`func (o *PropertyDefinition) GetDisplayLabel() string`

GetDisplayLabel returns the DisplayLabel field if non-nil, zero value otherwise.

### GetDisplayLabelOk

`func (o *PropertyDefinition) GetDisplayLabelOk() (*string, bool)`

GetDisplayLabelOk returns a tuple with the DisplayLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayLabel

`func (o *PropertyDefinition) SetDisplayLabel(v string)`

SetDisplayLabel sets DisplayLabel field to given value.

### HasDisplayLabel

`func (o *PropertyDefinition) HasDisplayLabel() bool`

HasDisplayLabel returns a boolean if a field has been set.

### GetDisplayLabelPlural

`func (o *PropertyDefinition) GetDisplayLabelPlural() string`

GetDisplayLabelPlural returns the DisplayLabelPlural field if non-nil, zero value otherwise.

### GetDisplayLabelPluralOk

`func (o *PropertyDefinition) GetDisplayLabelPluralOk() (*string, bool)`

GetDisplayLabelPluralOk returns a tuple with the DisplayLabelPlural field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayLabelPlural

`func (o *PropertyDefinition) SetDisplayLabelPlural(v string)`

SetDisplayLabelPlural sets DisplayLabelPlural field to given value.

### HasDisplayLabelPlural

`func (o *PropertyDefinition) HasDisplayLabelPlural() bool`

HasDisplayLabelPlural returns a boolean if a field has been set.

### GetPropertyType

`func (o *PropertyDefinition) GetPropertyType() string`

GetPropertyType returns the PropertyType field if non-nil, zero value otherwise.

### GetPropertyTypeOk

`func (o *PropertyDefinition) GetPropertyTypeOk() (*string, bool)`

GetPropertyTypeOk returns a tuple with the PropertyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyType

`func (o *PropertyDefinition) SetPropertyType(v string)`

SetPropertyType sets PropertyType field to given value.

### HasPropertyType

`func (o *PropertyDefinition) HasPropertyType() bool`

HasPropertyType returns a boolean if a field has been set.

### GetUiOptions

`func (o *PropertyDefinition) GetUiOptions() string`

GetUiOptions returns the UiOptions field if non-nil, zero value otherwise.

### GetUiOptionsOk

`func (o *PropertyDefinition) GetUiOptionsOk() (*string, bool)`

GetUiOptionsOk returns a tuple with the UiOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiOptions

`func (o *PropertyDefinition) SetUiOptions(v string)`

SetUiOptions sets UiOptions field to given value.

### HasUiOptions

`func (o *PropertyDefinition) HasUiOptions() bool`

HasUiOptions returns a boolean if a field has been set.

### GetHideUiFacet

`func (o *PropertyDefinition) GetHideUiFacet() bool`

GetHideUiFacet returns the HideUiFacet field if non-nil, zero value otherwise.

### GetHideUiFacetOk

`func (o *PropertyDefinition) GetHideUiFacetOk() (*bool, bool)`

GetHideUiFacetOk returns a tuple with the HideUiFacet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideUiFacet

`func (o *PropertyDefinition) SetHideUiFacet(v bool)`

SetHideUiFacet sets HideUiFacet field to given value.

### HasHideUiFacet

`func (o *PropertyDefinition) HasHideUiFacet() bool`

HasHideUiFacet returns a boolean if a field has been set.

### GetUiFacetOrder

`func (o *PropertyDefinition) GetUiFacetOrder() int32`

GetUiFacetOrder returns the UiFacetOrder field if non-nil, zero value otherwise.

### GetUiFacetOrderOk

`func (o *PropertyDefinition) GetUiFacetOrderOk() (*int32, bool)`

GetUiFacetOrderOk returns a tuple with the UiFacetOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiFacetOrder

`func (o *PropertyDefinition) SetUiFacetOrder(v int32)`

SetUiFacetOrder sets UiFacetOrder field to given value.

### HasUiFacetOrder

`func (o *PropertyDefinition) HasUiFacetOrder() bool`

HasUiFacetOrder returns a boolean if a field has been set.

### GetObjectPropertyOptions

`func (o *PropertyDefinition) GetObjectPropertyOptions() []ObjectPropertyOptions`

GetObjectPropertyOptions returns the ObjectPropertyOptions field if non-nil, zero value otherwise.

### GetObjectPropertyOptionsOk

`func (o *PropertyDefinition) GetObjectPropertyOptionsOk() (*[]ObjectPropertyOptions, bool)`

GetObjectPropertyOptionsOk returns a tuple with the ObjectPropertyOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectPropertyOptions

`func (o *PropertyDefinition) SetObjectPropertyOptions(v []ObjectPropertyOptions)`

SetObjectPropertyOptions sets ObjectPropertyOptions field to given value.

### HasObjectPropertyOptions

`func (o *PropertyDefinition) HasObjectPropertyOptions() bool`

HasObjectPropertyOptions returns a boolean if a field has been set.

### GetGroup

`func (o *PropertyDefinition) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *PropertyDefinition) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *PropertyDefinition) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *PropertyDefinition) HasGroup() bool`

HasGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


