# StatusLabelConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** |  | [optional] 
**OmitMissingValues** | Pointer to **bool** |  | [optional] 
**UseDefault** | Pointer to **bool** |  | [optional] 
**Values** | Pointer to [**[]LabelConfig**](LabelConfig.md) |  | [optional] 

## Methods

### NewStatusLabelConfig

`func NewStatusLabelConfig() *StatusLabelConfig`

NewStatusLabelConfig instantiates a new StatusLabelConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusLabelConfigWithDefaults

`func NewStatusLabelConfigWithDefaults() *StatusLabelConfig`

NewStatusLabelConfigWithDefaults instantiates a new StatusLabelConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *StatusLabelConfig) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *StatusLabelConfig) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *StatusLabelConfig) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *StatusLabelConfig) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetOmitMissingValues

`func (o *StatusLabelConfig) GetOmitMissingValues() bool`

GetOmitMissingValues returns the OmitMissingValues field if non-nil, zero value otherwise.

### GetOmitMissingValuesOk

`func (o *StatusLabelConfig) GetOmitMissingValuesOk() (*bool, bool)`

GetOmitMissingValuesOk returns a tuple with the OmitMissingValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmitMissingValues

`func (o *StatusLabelConfig) SetOmitMissingValues(v bool)`

SetOmitMissingValues sets OmitMissingValues field to given value.

### HasOmitMissingValues

`func (o *StatusLabelConfig) HasOmitMissingValues() bool`

HasOmitMissingValues returns a boolean if a field has been set.

### GetUseDefault

`func (o *StatusLabelConfig) GetUseDefault() bool`

GetUseDefault returns the UseDefault field if non-nil, zero value otherwise.

### GetUseDefaultOk

`func (o *StatusLabelConfig) GetUseDefaultOk() (*bool, bool)`

GetUseDefaultOk returns a tuple with the UseDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDefault

`func (o *StatusLabelConfig) SetUseDefault(v bool)`

SetUseDefault sets UseDefault field to given value.

### HasUseDefault

`func (o *StatusLabelConfig) HasUseDefault() bool`

HasUseDefault returns a boolean if a field has been set.

### GetValues

`func (o *StatusLabelConfig) GetValues() []LabelConfig`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *StatusLabelConfig) GetValuesOk() (*[]LabelConfig, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *StatusLabelConfig) SetValues(v []LabelConfig)`

SetValues sets Values field to given value.

### HasValues

`func (o *StatusLabelConfig) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


