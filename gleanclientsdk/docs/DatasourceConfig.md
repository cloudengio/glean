# DatasourceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | The human-readable name of the datasource | [optional] 
**HomeUrl** | Pointer to **string** | The fallback homepage for all instances of this datasource. Ex for slack: https://apps.slack.com/client | [optional] 
**IconUrl** | Pointer to **string** | The URL to an image to be displayed as an icon for this generic datasource. Must have a transparency mask. SVG are recommended over PNG. Public, glean-authenticated and Base64 encoded data URLs are all valid (but not third-party-authenticated URLs). | [optional] 
**Id** | Pointer to **string** | Id of the datasource. So happens to be lowercase name like \&quot;slack\&quot;, but could become opaque in the future. Don&#39;t rely on the id&#39;s format. | [optional] 
**Instances** | Pointer to **[]interface{}** | UNIMPLEMENTED. You&#39;re probably wanting DatasourceSchema, which are misnamed instances that are slated for replacement with this DatasourceConfig.instances field | [optional] 
**SupportsMultipleInstances** | Pointer to **bool** | Whether the datasource allows for multiple instances. Ex: \&quot;Jira (on prem)\&quot; and \&quot;Jira (cloud)\&quot; | [optional] 
**MultipleInstanceSetupEnabled** | Pointer to **bool** | Whether the datasource allows for automatic multiple instance setup. Currently only supported by GitHub.\&quot; | [optional] 

## Methods

### NewDatasourceConfig

`func NewDatasourceConfig() *DatasourceConfig`

NewDatasourceConfig instantiates a new DatasourceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatasourceConfigWithDefaults

`func NewDatasourceConfigWithDefaults() *DatasourceConfig`

NewDatasourceConfigWithDefaults instantiates a new DatasourceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *DatasourceConfig) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *DatasourceConfig) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *DatasourceConfig) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *DatasourceConfig) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetHomeUrl

`func (o *DatasourceConfig) GetHomeUrl() string`

GetHomeUrl returns the HomeUrl field if non-nil, zero value otherwise.

### GetHomeUrlOk

`func (o *DatasourceConfig) GetHomeUrlOk() (*string, bool)`

GetHomeUrlOk returns a tuple with the HomeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeUrl

`func (o *DatasourceConfig) SetHomeUrl(v string)`

SetHomeUrl sets HomeUrl field to given value.

### HasHomeUrl

`func (o *DatasourceConfig) HasHomeUrl() bool`

HasHomeUrl returns a boolean if a field has been set.

### GetIconUrl

`func (o *DatasourceConfig) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *DatasourceConfig) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *DatasourceConfig) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *DatasourceConfig) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### GetId

`func (o *DatasourceConfig) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DatasourceConfig) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DatasourceConfig) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DatasourceConfig) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInstances

`func (o *DatasourceConfig) GetInstances() []interface{}`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *DatasourceConfig) GetInstancesOk() (*[]interface{}, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *DatasourceConfig) SetInstances(v []interface{})`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *DatasourceConfig) HasInstances() bool`

HasInstances returns a boolean if a field has been set.

### GetSupportsMultipleInstances

`func (o *DatasourceConfig) GetSupportsMultipleInstances() bool`

GetSupportsMultipleInstances returns the SupportsMultipleInstances field if non-nil, zero value otherwise.

### GetSupportsMultipleInstancesOk

`func (o *DatasourceConfig) GetSupportsMultipleInstancesOk() (*bool, bool)`

GetSupportsMultipleInstancesOk returns a tuple with the SupportsMultipleInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsMultipleInstances

`func (o *DatasourceConfig) SetSupportsMultipleInstances(v bool)`

SetSupportsMultipleInstances sets SupportsMultipleInstances field to given value.

### HasSupportsMultipleInstances

`func (o *DatasourceConfig) HasSupportsMultipleInstances() bool`

HasSupportsMultipleInstances returns a boolean if a field has been set.

### GetMultipleInstanceSetupEnabled

`func (o *DatasourceConfig) GetMultipleInstanceSetupEnabled() bool`

GetMultipleInstanceSetupEnabled returns the MultipleInstanceSetupEnabled field if non-nil, zero value otherwise.

### GetMultipleInstanceSetupEnabledOk

`func (o *DatasourceConfig) GetMultipleInstanceSetupEnabledOk() (*bool, bool)`

GetMultipleInstanceSetupEnabledOk returns a tuple with the MultipleInstanceSetupEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleInstanceSetupEnabled

`func (o *DatasourceConfig) SetMultipleInstanceSetupEnabled(v bool)`

SetMultipleInstanceSetupEnabled sets MultipleInstanceSetupEnabled field to given value.

### HasMultipleInstanceSetupEnabled

`func (o *DatasourceConfig) HasMultipleInstanceSetupEnabled() bool`

HasMultipleInstanceSetupEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


