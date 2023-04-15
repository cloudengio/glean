# PinCollectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | Whether to pin or unpin | [default to "PIN"]
**Data** | Pointer to [**CollectionPinMetadata**](CollectionPinMetadata.md) |  | [optional] 
**Id** | **int32** | DEPRECTED - The id of the collection to be pinned. Insted use data.id. | 
**Category** | [**CollectionPinnableCategories**](CollectionPinnableCategories.md) |  | 

## Methods

### NewPinCollectionRequest

`func NewPinCollectionRequest(action string, id int32, category CollectionPinnableCategories, ) *PinCollectionRequest`

NewPinCollectionRequest instantiates a new PinCollectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPinCollectionRequestWithDefaults

`func NewPinCollectionRequestWithDefaults() *PinCollectionRequest`

NewPinCollectionRequestWithDefaults instantiates a new PinCollectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *PinCollectionRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PinCollectionRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PinCollectionRequest) SetAction(v string)`

SetAction sets Action field to given value.


### GetData

`func (o *PinCollectionRequest) GetData() CollectionPinMetadata`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PinCollectionRequest) GetDataOk() (*CollectionPinMetadata, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PinCollectionRequest) SetData(v CollectionPinMetadata)`

SetData sets Data field to given value.

### HasData

`func (o *PinCollectionRequest) HasData() bool`

HasData returns a boolean if a field has been set.

### GetId

`func (o *PinCollectionRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PinCollectionRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PinCollectionRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetCategory

`func (o *PinCollectionRequest) GetCategory() CollectionPinnableCategories`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *PinCollectionRequest) GetCategoryOk() (*CollectionPinnableCategories, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *PinCollectionRequest) SetCategory(v CollectionPinnableCategories)`

SetCategory sets Category field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


