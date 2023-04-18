# DatumLinkConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**LinkSource** | Pointer to [**DatumLinkSource**](DatumLinkSource.md) |  | [optional] 

## Methods

### NewDatumLinkConfig

`func NewDatumLinkConfig() *DatumLinkConfig`

NewDatumLinkConfig instantiates a new DatumLinkConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatumLinkConfigWithDefaults

`func NewDatumLinkConfigWithDefaults() *DatumLinkConfig`

NewDatumLinkConfigWithDefaults instantiates a new DatumLinkConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *DatumLinkConfig) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *DatumLinkConfig) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *DatumLinkConfig) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *DatumLinkConfig) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetRole

`func (o *DatumLinkConfig) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *DatumLinkConfig) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *DatumLinkConfig) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *DatumLinkConfig) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetLinkSource

`func (o *DatumLinkConfig) GetLinkSource() DatumLinkSource`

GetLinkSource returns the LinkSource field if non-nil, zero value otherwise.

### GetLinkSourceOk

`func (o *DatumLinkConfig) GetLinkSourceOk() (*DatumLinkSource, bool)`

GetLinkSourceOk returns a tuple with the LinkSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkSource

`func (o *DatumLinkConfig) SetLinkSource(v DatumLinkSource)`

SetLinkSource sets LinkSource field to given value.

### HasLinkSource

`func (o *DatumLinkConfig) HasLinkSource() bool`

HasLinkSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


