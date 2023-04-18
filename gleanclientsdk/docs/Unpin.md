# Unpin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PinId** | Pointer to **int32** | DEPRECATED - Prefer use of &#x60;id&#x60; | [optional] 
**Id** | Pointer to **string** | The opaque id of the pin to be unpinned. | [optional] 

## Methods

### NewUnpin

`func NewUnpin() *Unpin`

NewUnpin instantiates a new Unpin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnpinWithDefaults

`func NewUnpinWithDefaults() *Unpin`

NewUnpinWithDefaults instantiates a new Unpin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPinId

`func (o *Unpin) GetPinId() int32`

GetPinId returns the PinId field if non-nil, zero value otherwise.

### GetPinIdOk

`func (o *Unpin) GetPinIdOk() (*int32, bool)`

GetPinIdOk returns a tuple with the PinId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinId

`func (o *Unpin) SetPinId(v int32)`

SetPinId sets PinId field to given value.

### HasPinId

`func (o *Unpin) HasPinId() bool`

HasPinId returns a boolean if a field has been set.

### GetId

`func (o *Unpin) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Unpin) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Unpin) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Unpin) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


