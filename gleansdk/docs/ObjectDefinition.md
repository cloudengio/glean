# ObjectDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Unique identifier for this &#x60;DocumentMetadata.objectType&#x60;. If omitted, this definition is used as a default for all unmatched &#x60;DocumentMetadata.objectType&#x60;s in this datasource. | [optional] 
**DisplayLabel** | Pointer to **string** | The user-friendly name of the object for display. | [optional] 
**DocCategory** | Pointer to **string** | The document category of this object type. | [optional] 
**PropertyDefinitions** | Pointer to [**[]PropertyDefinition**](PropertyDefinition.md) |  | [optional] 
**PropertyGroups** | Pointer to [**[]PropertyGroup**](PropertyGroup.md) | A list of &#x60;PropertyGroup&#x60;s belonging to this object type, which will group properties to be displayed together in the UI. | [optional] 

## Methods

### NewObjectDefinition

`func NewObjectDefinition() *ObjectDefinition`

NewObjectDefinition instantiates a new ObjectDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectDefinitionWithDefaults

`func NewObjectDefinitionWithDefaults() *ObjectDefinition`

NewObjectDefinitionWithDefaults instantiates a new ObjectDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ObjectDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ObjectDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ObjectDefinition) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ObjectDefinition) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayLabel

`func (o *ObjectDefinition) GetDisplayLabel() string`

GetDisplayLabel returns the DisplayLabel field if non-nil, zero value otherwise.

### GetDisplayLabelOk

`func (o *ObjectDefinition) GetDisplayLabelOk() (*string, bool)`

GetDisplayLabelOk returns a tuple with the DisplayLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayLabel

`func (o *ObjectDefinition) SetDisplayLabel(v string)`

SetDisplayLabel sets DisplayLabel field to given value.

### HasDisplayLabel

`func (o *ObjectDefinition) HasDisplayLabel() bool`

HasDisplayLabel returns a boolean if a field has been set.

### GetDocCategory

`func (o *ObjectDefinition) GetDocCategory() string`

GetDocCategory returns the DocCategory field if non-nil, zero value otherwise.

### GetDocCategoryOk

`func (o *ObjectDefinition) GetDocCategoryOk() (*string, bool)`

GetDocCategoryOk returns a tuple with the DocCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocCategory

`func (o *ObjectDefinition) SetDocCategory(v string)`

SetDocCategory sets DocCategory field to given value.

### HasDocCategory

`func (o *ObjectDefinition) HasDocCategory() bool`

HasDocCategory returns a boolean if a field has been set.

### GetPropertyDefinitions

`func (o *ObjectDefinition) GetPropertyDefinitions() []PropertyDefinition`

GetPropertyDefinitions returns the PropertyDefinitions field if non-nil, zero value otherwise.

### GetPropertyDefinitionsOk

`func (o *ObjectDefinition) GetPropertyDefinitionsOk() (*[]PropertyDefinition, bool)`

GetPropertyDefinitionsOk returns a tuple with the PropertyDefinitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyDefinitions

`func (o *ObjectDefinition) SetPropertyDefinitions(v []PropertyDefinition)`

SetPropertyDefinitions sets PropertyDefinitions field to given value.

### HasPropertyDefinitions

`func (o *ObjectDefinition) HasPropertyDefinitions() bool`

HasPropertyDefinitions returns a boolean if a field has been set.

### GetPropertyGroups

`func (o *ObjectDefinition) GetPropertyGroups() []PropertyGroup`

GetPropertyGroups returns the PropertyGroups field if non-nil, zero value otherwise.

### GetPropertyGroupsOk

`func (o *ObjectDefinition) GetPropertyGroupsOk() (*[]PropertyGroup, bool)`

GetPropertyGroupsOk returns a tuple with the PropertyGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyGroups

`func (o *ObjectDefinition) SetPropertyGroups(v []PropertyGroup)`

SetPropertyGroups sets PropertyGroups field to given value.

### HasPropertyGroups

`func (o *ObjectDefinition) HasPropertyGroups() bool`

HasPropertyGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


