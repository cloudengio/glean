# ConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ThemeKeys** | Pointer to **[]string** | A list of theme keys to include in the response. | [optional] 
**BoolKeys** | Pointer to **[]string** | A list of boolean flag keys to include in the response. | [optional] 
**SessionInfo** | Pointer to [**SessionInfo**](SessionInfo.md) |  | [optional] 

## Methods

### NewConfigRequest

`func NewConfigRequest() *ConfigRequest`

NewConfigRequest instantiates a new ConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigRequestWithDefaults

`func NewConfigRequestWithDefaults() *ConfigRequest`

NewConfigRequestWithDefaults instantiates a new ConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThemeKeys

`func (o *ConfigRequest) GetThemeKeys() []string`

GetThemeKeys returns the ThemeKeys field if non-nil, zero value otherwise.

### GetThemeKeysOk

`func (o *ConfigRequest) GetThemeKeysOk() (*[]string, bool)`

GetThemeKeysOk returns a tuple with the ThemeKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThemeKeys

`func (o *ConfigRequest) SetThemeKeys(v []string)`

SetThemeKeys sets ThemeKeys field to given value.

### HasThemeKeys

`func (o *ConfigRequest) HasThemeKeys() bool`

HasThemeKeys returns a boolean if a field has been set.

### GetBoolKeys

`func (o *ConfigRequest) GetBoolKeys() []string`

GetBoolKeys returns the BoolKeys field if non-nil, zero value otherwise.

### GetBoolKeysOk

`func (o *ConfigRequest) GetBoolKeysOk() (*[]string, bool)`

GetBoolKeysOk returns a tuple with the BoolKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoolKeys

`func (o *ConfigRequest) SetBoolKeys(v []string)`

SetBoolKeys sets BoolKeys field to given value.

### HasBoolKeys

`func (o *ConfigRequest) HasBoolKeys() bool`

HasBoolKeys returns a boolean if a field has been set.

### GetSessionInfo

`func (o *ConfigRequest) GetSessionInfo() SessionInfo`

GetSessionInfo returns the SessionInfo field if non-nil, zero value otherwise.

### GetSessionInfoOk

`func (o *ConfigRequest) GetSessionInfoOk() (*SessionInfo, bool)`

GetSessionInfoOk returns a tuple with the SessionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionInfo

`func (o *ConfigRequest) SetSessionInfo(v SessionInfo)`

SetSessionInfo sets SessionInfo field to given value.

### HasSessionInfo

`func (o *ConfigRequest) HasSessionInfo() bool`

HasSessionInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


