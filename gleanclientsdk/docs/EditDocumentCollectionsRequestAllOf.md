# EditDocumentCollectionsRequestAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentId** | Pointer to **string** | The document ID of the item being added to or removed from collections if it&#39;s a Glean-indexed document. | [optional] 
**Url** | Pointer to **string** | The url of the item being added to or removed from collections. | [optional] 
**Name** | Pointer to **string** | Custom title of the document if adding a non-indexed URL. | [optional] 
**Description** | Pointer to **string** | The description of this CollectionItem. | [optional] 

## Methods

### NewEditDocumentCollectionsRequestAllOf

`func NewEditDocumentCollectionsRequestAllOf() *EditDocumentCollectionsRequestAllOf`

NewEditDocumentCollectionsRequestAllOf instantiates a new EditDocumentCollectionsRequestAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditDocumentCollectionsRequestAllOfWithDefaults

`func NewEditDocumentCollectionsRequestAllOfWithDefaults() *EditDocumentCollectionsRequestAllOf`

NewEditDocumentCollectionsRequestAllOfWithDefaults instantiates a new EditDocumentCollectionsRequestAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentId

`func (o *EditDocumentCollectionsRequestAllOf) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *EditDocumentCollectionsRequestAllOf) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *EditDocumentCollectionsRequestAllOf) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *EditDocumentCollectionsRequestAllOf) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### GetUrl

`func (o *EditDocumentCollectionsRequestAllOf) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *EditDocumentCollectionsRequestAllOf) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *EditDocumentCollectionsRequestAllOf) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *EditDocumentCollectionsRequestAllOf) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetName

`func (o *EditDocumentCollectionsRequestAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EditDocumentCollectionsRequestAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EditDocumentCollectionsRequestAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EditDocumentCollectionsRequestAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *EditDocumentCollectionsRequestAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EditDocumentCollectionsRequestAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EditDocumentCollectionsRequestAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EditDocumentCollectionsRequestAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


