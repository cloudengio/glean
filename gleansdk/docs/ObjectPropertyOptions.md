# ObjectPropertyOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubobjectProperties** | Pointer to [**[]PropertyDefinition**](PropertyDefinition.md) | The properties of the sub-object. These properties represent a nested object. For example, if this property represents a postal address, the subobjectProperties might be named street, city, and state. | [optional] 

## Methods

### NewObjectPropertyOptions

`func NewObjectPropertyOptions() *ObjectPropertyOptions`

NewObjectPropertyOptions instantiates a new ObjectPropertyOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectPropertyOptionsWithDefaults

`func NewObjectPropertyOptionsWithDefaults() *ObjectPropertyOptions`

NewObjectPropertyOptionsWithDefaults instantiates a new ObjectPropertyOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubobjectProperties

`func (o *ObjectPropertyOptions) GetSubobjectProperties() []PropertyDefinition`

GetSubobjectProperties returns the SubobjectProperties field if non-nil, zero value otherwise.

### GetSubobjectPropertiesOk

`func (o *ObjectPropertyOptions) GetSubobjectPropertiesOk() (*[]PropertyDefinition, bool)`

GetSubobjectPropertiesOk returns a tuple with the SubobjectProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubobjectProperties

`func (o *ObjectPropertyOptions) SetSubobjectProperties(v []PropertyDefinition)`

SetSubobjectProperties sets SubobjectProperties field to given value.

### HasSubobjectProperties

`func (o *ObjectPropertyOptions) HasSubobjectProperties() bool`

HasSubobjectProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


