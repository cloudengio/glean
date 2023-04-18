# DatumRoleConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**LinkSource** | Pointer to [**DatumLinkSource**](DatumLinkSource.md) |  | [optional] 

## Methods

### NewDatumRoleConfig

`func NewDatumRoleConfig() *DatumRoleConfig`

NewDatumRoleConfig instantiates a new DatumRoleConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatumRoleConfigWithDefaults

`func NewDatumRoleConfigWithDefaults() *DatumRoleConfig`

NewDatumRoleConfigWithDefaults instantiates a new DatumRoleConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *DatumRoleConfig) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *DatumRoleConfig) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *DatumRoleConfig) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *DatumRoleConfig) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetRole

`func (o *DatumRoleConfig) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *DatumRoleConfig) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *DatumRoleConfig) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *DatumRoleConfig) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetLinkSource

`func (o *DatumRoleConfig) GetLinkSource() DatumLinkSource`

GetLinkSource returns the LinkSource field if non-nil, zero value otherwise.

### GetLinkSourceOk

`func (o *DatumRoleConfig) GetLinkSourceOk() (*DatumLinkSource, bool)`

GetLinkSourceOk returns a tuple with the LinkSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkSource

`func (o *DatumRoleConfig) SetLinkSource(v DatumLinkSource)`

SetLinkSource sets LinkSource field to given value.

### HasLinkSource

`func (o *DatumRoleConfig) HasLinkSource() bool`

HasLinkSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


