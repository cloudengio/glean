# CollectionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The unique ID of the collection. | 
**CreateTime** | Pointer to **time.Time** |  | [optional] 
**UpdateTime** | Pointer to **time.Time** |  | [optional] 
**Creator** | Pointer to [**Person**](Person.md) |  | [optional] 
**UpdatedBy** | Pointer to [**Person**](Person.md) |  | [optional] 
**ItemCount** | Pointer to **int32** | The number of items currently in the Collection. Separated from the actual items so we can grab the count without items. | [optional] 
**ChildCount** | Pointer to **int32** | The number of children Collections. Separated from the actual children so we can grab the count without children. | [optional] 
**Items** | Pointer to [**[]CollectionItem**](CollectionItem.md) | The items in this Collection. | [optional] 
**PinMetadata** | Pointer to [**CollectionPinnedMetadata**](CollectionPinnedMetadata.md) |  | [optional] 
**Shortcuts** | Pointer to **[]string** | The names of the shortcuts (Go Links) that point to this Collection. | [optional] 
**Children** | Pointer to [**[]Collection**](Collection.md) | The children Collections of this Collection. | [optional] 
**Roles** | Pointer to [**[]UserRoleSpecification**](UserRoleSpecification.md) | A list of user roles for the collection. | [optional] 

## Methods

### NewCollectionAllOf

`func NewCollectionAllOf(id int32, ) *CollectionAllOf`

NewCollectionAllOf instantiates a new CollectionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionAllOfWithDefaults

`func NewCollectionAllOfWithDefaults() *CollectionAllOf`

NewCollectionAllOfWithDefaults instantiates a new CollectionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CollectionAllOf) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CollectionAllOf) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CollectionAllOf) SetId(v int32)`

SetId sets Id field to given value.


### GetCreateTime

`func (o *CollectionAllOf) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *CollectionAllOf) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *CollectionAllOf) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *CollectionAllOf) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetUpdateTime

`func (o *CollectionAllOf) GetUpdateTime() time.Time`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *CollectionAllOf) GetUpdateTimeOk() (*time.Time, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *CollectionAllOf) SetUpdateTime(v time.Time)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *CollectionAllOf) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetCreator

`func (o *CollectionAllOf) GetCreator() Person`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *CollectionAllOf) GetCreatorOk() (*Person, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *CollectionAllOf) SetCreator(v Person)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *CollectionAllOf) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *CollectionAllOf) GetUpdatedBy() Person`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *CollectionAllOf) GetUpdatedByOk() (*Person, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *CollectionAllOf) SetUpdatedBy(v Person)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *CollectionAllOf) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetItemCount

`func (o *CollectionAllOf) GetItemCount() int32`

GetItemCount returns the ItemCount field if non-nil, zero value otherwise.

### GetItemCountOk

`func (o *CollectionAllOf) GetItemCountOk() (*int32, bool)`

GetItemCountOk returns a tuple with the ItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemCount

`func (o *CollectionAllOf) SetItemCount(v int32)`

SetItemCount sets ItemCount field to given value.

### HasItemCount

`func (o *CollectionAllOf) HasItemCount() bool`

HasItemCount returns a boolean if a field has been set.

### GetChildCount

`func (o *CollectionAllOf) GetChildCount() int32`

GetChildCount returns the ChildCount field if non-nil, zero value otherwise.

### GetChildCountOk

`func (o *CollectionAllOf) GetChildCountOk() (*int32, bool)`

GetChildCountOk returns a tuple with the ChildCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildCount

`func (o *CollectionAllOf) SetChildCount(v int32)`

SetChildCount sets ChildCount field to given value.

### HasChildCount

`func (o *CollectionAllOf) HasChildCount() bool`

HasChildCount returns a boolean if a field has been set.

### GetItems

`func (o *CollectionAllOf) GetItems() []CollectionItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CollectionAllOf) GetItemsOk() (*[]CollectionItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CollectionAllOf) SetItems(v []CollectionItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *CollectionAllOf) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetPinMetadata

`func (o *CollectionAllOf) GetPinMetadata() CollectionPinnedMetadata`

GetPinMetadata returns the PinMetadata field if non-nil, zero value otherwise.

### GetPinMetadataOk

`func (o *CollectionAllOf) GetPinMetadataOk() (*CollectionPinnedMetadata, bool)`

GetPinMetadataOk returns a tuple with the PinMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinMetadata

`func (o *CollectionAllOf) SetPinMetadata(v CollectionPinnedMetadata)`

SetPinMetadata sets PinMetadata field to given value.

### HasPinMetadata

`func (o *CollectionAllOf) HasPinMetadata() bool`

HasPinMetadata returns a boolean if a field has been set.

### GetShortcuts

`func (o *CollectionAllOf) GetShortcuts() []string`

GetShortcuts returns the Shortcuts field if non-nil, zero value otherwise.

### GetShortcutsOk

`func (o *CollectionAllOf) GetShortcutsOk() (*[]string, bool)`

GetShortcutsOk returns a tuple with the Shortcuts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortcuts

`func (o *CollectionAllOf) SetShortcuts(v []string)`

SetShortcuts sets Shortcuts field to given value.

### HasShortcuts

`func (o *CollectionAllOf) HasShortcuts() bool`

HasShortcuts returns a boolean if a field has been set.

### GetChildren

`func (o *CollectionAllOf) GetChildren() []Collection`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *CollectionAllOf) GetChildrenOk() (*[]Collection, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *CollectionAllOf) SetChildren(v []Collection)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *CollectionAllOf) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetRoles

`func (o *CollectionAllOf) GetRoles() []UserRoleSpecification`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *CollectionAllOf) GetRolesOk() (*[]UserRoleSpecification, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *CollectionAllOf) SetRoles(v []UserRoleSpecification)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *CollectionAllOf) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


