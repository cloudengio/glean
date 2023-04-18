# CollectionMutablePropertiesAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Icon** | Pointer to **string** | The emoji icon of this Collection. | [optional] 
**AdminLocked** | Pointer to **bool** | Indicates whether edits are allowed for everyone or only admins. | [optional] 
**ParentId** | Pointer to **int32** | The parent of this Collection, or 0 if it&#39;s a top-level Collection. | [optional] 
**Thumbnail** | Pointer to [**Thumbnail**](Thumbnail.md) |  | [optional] 
**AllowedDatasource** | Pointer to **string** | The datasource type this collection can hold. | [optional] 

## Methods

### NewCollectionMutablePropertiesAllOf

`func NewCollectionMutablePropertiesAllOf() *CollectionMutablePropertiesAllOf`

NewCollectionMutablePropertiesAllOf instantiates a new CollectionMutablePropertiesAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionMutablePropertiesAllOfWithDefaults

`func NewCollectionMutablePropertiesAllOfWithDefaults() *CollectionMutablePropertiesAllOf`

NewCollectionMutablePropertiesAllOfWithDefaults instantiates a new CollectionMutablePropertiesAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIcon

`func (o *CollectionMutablePropertiesAllOf) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *CollectionMutablePropertiesAllOf) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *CollectionMutablePropertiesAllOf) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *CollectionMutablePropertiesAllOf) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetAdminLocked

`func (o *CollectionMutablePropertiesAllOf) GetAdminLocked() bool`

GetAdminLocked returns the AdminLocked field if non-nil, zero value otherwise.

### GetAdminLockedOk

`func (o *CollectionMutablePropertiesAllOf) GetAdminLockedOk() (*bool, bool)`

GetAdminLockedOk returns a tuple with the AdminLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminLocked

`func (o *CollectionMutablePropertiesAllOf) SetAdminLocked(v bool)`

SetAdminLocked sets AdminLocked field to given value.

### HasAdminLocked

`func (o *CollectionMutablePropertiesAllOf) HasAdminLocked() bool`

HasAdminLocked returns a boolean if a field has been set.

### GetParentId

`func (o *CollectionMutablePropertiesAllOf) GetParentId() int32`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *CollectionMutablePropertiesAllOf) GetParentIdOk() (*int32, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *CollectionMutablePropertiesAllOf) SetParentId(v int32)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *CollectionMutablePropertiesAllOf) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetThumbnail

`func (o *CollectionMutablePropertiesAllOf) GetThumbnail() Thumbnail`

GetThumbnail returns the Thumbnail field if non-nil, zero value otherwise.

### GetThumbnailOk

`func (o *CollectionMutablePropertiesAllOf) GetThumbnailOk() (*Thumbnail, bool)`

GetThumbnailOk returns a tuple with the Thumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnail

`func (o *CollectionMutablePropertiesAllOf) SetThumbnail(v Thumbnail)`

SetThumbnail sets Thumbnail field to given value.

### HasThumbnail

`func (o *CollectionMutablePropertiesAllOf) HasThumbnail() bool`

HasThumbnail returns a boolean if a field has been set.

### GetAllowedDatasource

`func (o *CollectionMutablePropertiesAllOf) GetAllowedDatasource() string`

GetAllowedDatasource returns the AllowedDatasource field if non-nil, zero value otherwise.

### GetAllowedDatasourceOk

`func (o *CollectionMutablePropertiesAllOf) GetAllowedDatasourceOk() (*string, bool)`

GetAllowedDatasourceOk returns a tuple with the AllowedDatasource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedDatasource

`func (o *CollectionMutablePropertiesAllOf) SetAllowedDatasource(v string)`

SetAllowedDatasource sets AllowedDatasource field to given value.

### HasAllowedDatasource

`func (o *CollectionMutablePropertiesAllOf) HasAllowedDatasource() bool`

HasAllowedDatasource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


