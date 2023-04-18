# ConfigResponseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatasourceConfigs** | Pointer to [**[]DatasourceConfig**](DatasourceConfig.md) | List of datasource \&quot;classes\&quot;. Will contain a list of datasource instances, but if you need instances for now, continue to use ConfigResponse.datasources (misnamed instances) | [optional] 
**Datasources** | Pointer to [**[]DatasourceSchema**](DatasourceSchema.md) | List of configurations for datasource instances that may be displayed. | [optional] 
**UserSettings** | Pointer to [**UserSettings**](UserSettings.md) |  | [optional] 
**ServerBuildVersion** | Pointer to **string** | Build versions to be rendered in debug mode. | [optional] 
**ClientConfig** | Pointer to [**ClientConfig**](ClientConfig.md) |  | [optional] 
**Permissions** | Pointer to [**Permissions**](Permissions.md) |  | [optional] 
**SessionInfo** | Pointer to [**SessionInfo**](SessionInfo.md) |  | [optional] 
**RenderConfigs** | Pointer to [**map[string]map[string]RenderConfig**](map.md) |  | [optional] 
**Navigation** | Pointer to [**Navigation**](Navigation.md) |  | [optional] 

## Methods

### NewConfigResponseAllOf

`func NewConfigResponseAllOf() *ConfigResponseAllOf`

NewConfigResponseAllOf instantiates a new ConfigResponseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigResponseAllOfWithDefaults

`func NewConfigResponseAllOfWithDefaults() *ConfigResponseAllOf`

NewConfigResponseAllOfWithDefaults instantiates a new ConfigResponseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatasourceConfigs

`func (o *ConfigResponseAllOf) GetDatasourceConfigs() []DatasourceConfig`

GetDatasourceConfigs returns the DatasourceConfigs field if non-nil, zero value otherwise.

### GetDatasourceConfigsOk

`func (o *ConfigResponseAllOf) GetDatasourceConfigsOk() (*[]DatasourceConfig, bool)`

GetDatasourceConfigsOk returns a tuple with the DatasourceConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasourceConfigs

`func (o *ConfigResponseAllOf) SetDatasourceConfigs(v []DatasourceConfig)`

SetDatasourceConfigs sets DatasourceConfigs field to given value.

### HasDatasourceConfigs

`func (o *ConfigResponseAllOf) HasDatasourceConfigs() bool`

HasDatasourceConfigs returns a boolean if a field has been set.

### GetDatasources

`func (o *ConfigResponseAllOf) GetDatasources() []DatasourceSchema`

GetDatasources returns the Datasources field if non-nil, zero value otherwise.

### GetDatasourcesOk

`func (o *ConfigResponseAllOf) GetDatasourcesOk() (*[]DatasourceSchema, bool)`

GetDatasourcesOk returns a tuple with the Datasources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasources

`func (o *ConfigResponseAllOf) SetDatasources(v []DatasourceSchema)`

SetDatasources sets Datasources field to given value.

### HasDatasources

`func (o *ConfigResponseAllOf) HasDatasources() bool`

HasDatasources returns a boolean if a field has been set.

### GetUserSettings

`func (o *ConfigResponseAllOf) GetUserSettings() UserSettings`

GetUserSettings returns the UserSettings field if non-nil, zero value otherwise.

### GetUserSettingsOk

`func (o *ConfigResponseAllOf) GetUserSettingsOk() (*UserSettings, bool)`

GetUserSettingsOk returns a tuple with the UserSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSettings

`func (o *ConfigResponseAllOf) SetUserSettings(v UserSettings)`

SetUserSettings sets UserSettings field to given value.

### HasUserSettings

`func (o *ConfigResponseAllOf) HasUserSettings() bool`

HasUserSettings returns a boolean if a field has been set.

### GetServerBuildVersion

`func (o *ConfigResponseAllOf) GetServerBuildVersion() string`

GetServerBuildVersion returns the ServerBuildVersion field if non-nil, zero value otherwise.

### GetServerBuildVersionOk

`func (o *ConfigResponseAllOf) GetServerBuildVersionOk() (*string, bool)`

GetServerBuildVersionOk returns a tuple with the ServerBuildVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerBuildVersion

`func (o *ConfigResponseAllOf) SetServerBuildVersion(v string)`

SetServerBuildVersion sets ServerBuildVersion field to given value.

### HasServerBuildVersion

`func (o *ConfigResponseAllOf) HasServerBuildVersion() bool`

HasServerBuildVersion returns a boolean if a field has been set.

### GetClientConfig

`func (o *ConfigResponseAllOf) GetClientConfig() ClientConfig`

GetClientConfig returns the ClientConfig field if non-nil, zero value otherwise.

### GetClientConfigOk

`func (o *ConfigResponseAllOf) GetClientConfigOk() (*ClientConfig, bool)`

GetClientConfigOk returns a tuple with the ClientConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientConfig

`func (o *ConfigResponseAllOf) SetClientConfig(v ClientConfig)`

SetClientConfig sets ClientConfig field to given value.

### HasClientConfig

`func (o *ConfigResponseAllOf) HasClientConfig() bool`

HasClientConfig returns a boolean if a field has been set.

### GetPermissions

`func (o *ConfigResponseAllOf) GetPermissions() Permissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *ConfigResponseAllOf) GetPermissionsOk() (*Permissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *ConfigResponseAllOf) SetPermissions(v Permissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *ConfigResponseAllOf) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetSessionInfo

`func (o *ConfigResponseAllOf) GetSessionInfo() SessionInfo`

GetSessionInfo returns the SessionInfo field if non-nil, zero value otherwise.

### GetSessionInfoOk

`func (o *ConfigResponseAllOf) GetSessionInfoOk() (*SessionInfo, bool)`

GetSessionInfoOk returns a tuple with the SessionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionInfo

`func (o *ConfigResponseAllOf) SetSessionInfo(v SessionInfo)`

SetSessionInfo sets SessionInfo field to given value.

### HasSessionInfo

`func (o *ConfigResponseAllOf) HasSessionInfo() bool`

HasSessionInfo returns a boolean if a field has been set.

### GetRenderConfigs

`func (o *ConfigResponseAllOf) GetRenderConfigs() map[string]map[string]RenderConfig`

GetRenderConfigs returns the RenderConfigs field if non-nil, zero value otherwise.

### GetRenderConfigsOk

`func (o *ConfigResponseAllOf) GetRenderConfigsOk() (*map[string]map[string]RenderConfig, bool)`

GetRenderConfigsOk returns a tuple with the RenderConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenderConfigs

`func (o *ConfigResponseAllOf) SetRenderConfigs(v map[string]map[string]RenderConfig)`

SetRenderConfigs sets RenderConfigs field to given value.

### HasRenderConfigs

`func (o *ConfigResponseAllOf) HasRenderConfigs() bool`

HasRenderConfigs returns a boolean if a field has been set.

### GetNavigation

`func (o *ConfigResponseAllOf) GetNavigation() Navigation`

GetNavigation returns the Navigation field if non-nil, zero value otherwise.

### GetNavigationOk

`func (o *ConfigResponseAllOf) GetNavigationOk() (*Navigation, bool)`

GetNavigationOk returns a tuple with the Navigation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNavigation

`func (o *ConfigResponseAllOf) SetNavigation(v Navigation)`

SetNavigation sets Navigation field to given value.

### HasNavigation

`func (o *ConfigResponseAllOf) HasNavigation() bool`

HasNavigation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


