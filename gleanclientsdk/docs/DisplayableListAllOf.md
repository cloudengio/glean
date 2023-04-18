# DisplayableListAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | unique identifier for a list, not unique amongst all objects | [optional] 
**SourceId** | Pointer to **string** | unstructured identifier for the source to render (id, url, query) | [optional] 
**Config** | Pointer to [**DisplayableListConfig**](DisplayableListConfig.md) |  | [optional] 

## Methods

### NewDisplayableListAllOf

`func NewDisplayableListAllOf() *DisplayableListAllOf`

NewDisplayableListAllOf instantiates a new DisplayableListAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDisplayableListAllOfWithDefaults

`func NewDisplayableListAllOfWithDefaults() *DisplayableListAllOf`

NewDisplayableListAllOfWithDefaults instantiates a new DisplayableListAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DisplayableListAllOf) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DisplayableListAllOf) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DisplayableListAllOf) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *DisplayableListAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSourceId

`func (o *DisplayableListAllOf) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *DisplayableListAllOf) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *DisplayableListAllOf) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *DisplayableListAllOf) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetConfig

`func (o *DisplayableListAllOf) GetConfig() DisplayableListConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *DisplayableListAllOf) GetConfigOk() (*DisplayableListConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *DisplayableListAllOf) SetConfig(v DisplayableListConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *DisplayableListAllOf) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


