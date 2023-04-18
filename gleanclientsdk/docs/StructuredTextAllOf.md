# StructuredTextAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StructuredList** | Pointer to [**[]StructuredTextItem**](StructuredTextItem.md) | An array of objects each of which contains either a string or a link which optionally corresponds to a document. | [optional] 

## Methods

### NewStructuredTextAllOf

`func NewStructuredTextAllOf() *StructuredTextAllOf`

NewStructuredTextAllOf instantiates a new StructuredTextAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStructuredTextAllOfWithDefaults

`func NewStructuredTextAllOfWithDefaults() *StructuredTextAllOf`

NewStructuredTextAllOfWithDefaults instantiates a new StructuredTextAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStructuredList

`func (o *StructuredTextAllOf) GetStructuredList() []StructuredTextItem`

GetStructuredList returns the StructuredList field if non-nil, zero value otherwise.

### GetStructuredListOk

`func (o *StructuredTextAllOf) GetStructuredListOk() (*[]StructuredTextItem, bool)`

GetStructuredListOk returns a tuple with the StructuredList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructuredList

`func (o *StructuredTextAllOf) SetStructuredList(v []StructuredTextItem)`

SetStructuredList sets StructuredList field to given value.

### HasStructuredList

`func (o *StructuredTextAllOf) HasStructuredList() bool`

HasStructuredList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


