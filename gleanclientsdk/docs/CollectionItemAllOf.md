# CollectionItemAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionId** | **int32** | The Collection ID of the Collection that this CollectionItem belongs in. | 
**DocumentId** | Pointer to **string** | If this CollectionItem is a Glean-indexed document, the document ID of that document. | [optional] 
**Url** | Pointer to **string** | The URL of this CollectionItem. | [optional] 
**ItemId** | Pointer to **string** | Unique identifier for the item within the collection it belongs to. | [optional] 
**CreatedBy** | Pointer to [**Person**](Person.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | Unix timestamp for when the item was first added (in seconds since epoch UTC). | [optional] 
**Document** | Pointer to [**Document**](Document.md) |  | [optional] 
**Shortcut** | Pointer to [**Shortcut**](Shortcut.md) |  | [optional] 
**Collection** | Pointer to [**Collection**](Collection.md) |  | [optional] 
**ItemType** | **string** |  | 

## Methods

### NewCollectionItemAllOf

`func NewCollectionItemAllOf(collectionId int32, itemType string, ) *CollectionItemAllOf`

NewCollectionItemAllOf instantiates a new CollectionItemAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionItemAllOfWithDefaults

`func NewCollectionItemAllOfWithDefaults() *CollectionItemAllOf`

NewCollectionItemAllOfWithDefaults instantiates a new CollectionItemAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionId

`func (o *CollectionItemAllOf) GetCollectionId() int32`

GetCollectionId returns the CollectionId field if non-nil, zero value otherwise.

### GetCollectionIdOk

`func (o *CollectionItemAllOf) GetCollectionIdOk() (*int32, bool)`

GetCollectionIdOk returns a tuple with the CollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionId

`func (o *CollectionItemAllOf) SetCollectionId(v int32)`

SetCollectionId sets CollectionId field to given value.


### GetDocumentId

`func (o *CollectionItemAllOf) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *CollectionItemAllOf) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *CollectionItemAllOf) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *CollectionItemAllOf) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### GetUrl

`func (o *CollectionItemAllOf) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CollectionItemAllOf) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CollectionItemAllOf) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CollectionItemAllOf) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetItemId

`func (o *CollectionItemAllOf) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *CollectionItemAllOf) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *CollectionItemAllOf) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *CollectionItemAllOf) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetCreatedBy

`func (o *CollectionItemAllOf) GetCreatedBy() Person`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *CollectionItemAllOf) GetCreatedByOk() (*Person, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *CollectionItemAllOf) SetCreatedBy(v Person)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *CollectionItemAllOf) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CollectionItemAllOf) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CollectionItemAllOf) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CollectionItemAllOf) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CollectionItemAllOf) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDocument

`func (o *CollectionItemAllOf) GetDocument() Document`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *CollectionItemAllOf) GetDocumentOk() (*Document, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *CollectionItemAllOf) SetDocument(v Document)`

SetDocument sets Document field to given value.

### HasDocument

`func (o *CollectionItemAllOf) HasDocument() bool`

HasDocument returns a boolean if a field has been set.

### GetShortcut

`func (o *CollectionItemAllOf) GetShortcut() Shortcut`

GetShortcut returns the Shortcut field if non-nil, zero value otherwise.

### GetShortcutOk

`func (o *CollectionItemAllOf) GetShortcutOk() (*Shortcut, bool)`

GetShortcutOk returns a tuple with the Shortcut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortcut

`func (o *CollectionItemAllOf) SetShortcut(v Shortcut)`

SetShortcut sets Shortcut field to given value.

### HasShortcut

`func (o *CollectionItemAllOf) HasShortcut() bool`

HasShortcut returns a boolean if a field has been set.

### GetCollection

`func (o *CollectionItemAllOf) GetCollection() Collection`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *CollectionItemAllOf) GetCollectionOk() (*Collection, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *CollectionItemAllOf) SetCollection(v Collection)`

SetCollection sets Collection field to given value.

### HasCollection

`func (o *CollectionItemAllOf) HasCollection() bool`

HasCollection returns a boolean if a field has been set.

### GetItemType

`func (o *CollectionItemAllOf) GetItemType() string`

GetItemType returns the ItemType field if non-nil, zero value otherwise.

### GetItemTypeOk

`func (o *CollectionItemAllOf) GetItemTypeOk() (*string, bool)`

GetItemTypeOk returns a tuple with the ItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemType

`func (o *CollectionItemAllOf) SetItemType(v string)`

SetItemType sets ItemType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


