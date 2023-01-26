# SharedDatasourceConfigAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatasourceName** | Pointer to **string** | The (non-unique) name of the datasource of which this config is an instance, e.g. \&quot;jira\&quot;. | [optional] 
**InstanceOnlyName** | Pointer to **string** | The instance of the datasource for this particular config, e.g. \&quot;onprem\&quot;. | [optional] 
**InstanceDescription** | Pointer to **string** | A human readable string identifying this instance as compared to its peers, e.g. \&quot;github.com/askscio\&quot; or \&quot;github.askscio.com\&quot; | [optional] 

## Methods

### NewSharedDatasourceConfigAllOf

`func NewSharedDatasourceConfigAllOf() *SharedDatasourceConfigAllOf`

NewSharedDatasourceConfigAllOf instantiates a new SharedDatasourceConfigAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharedDatasourceConfigAllOfWithDefaults

`func NewSharedDatasourceConfigAllOfWithDefaults() *SharedDatasourceConfigAllOf`

NewSharedDatasourceConfigAllOfWithDefaults instantiates a new SharedDatasourceConfigAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatasourceName

`func (o *SharedDatasourceConfigAllOf) GetDatasourceName() string`

GetDatasourceName returns the DatasourceName field if non-nil, zero value otherwise.

### GetDatasourceNameOk

`func (o *SharedDatasourceConfigAllOf) GetDatasourceNameOk() (*string, bool)`

GetDatasourceNameOk returns a tuple with the DatasourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasourceName

`func (o *SharedDatasourceConfigAllOf) SetDatasourceName(v string)`

SetDatasourceName sets DatasourceName field to given value.

### HasDatasourceName

`func (o *SharedDatasourceConfigAllOf) HasDatasourceName() bool`

HasDatasourceName returns a boolean if a field has been set.

### GetInstanceOnlyName

`func (o *SharedDatasourceConfigAllOf) GetInstanceOnlyName() string`

GetInstanceOnlyName returns the InstanceOnlyName field if non-nil, zero value otherwise.

### GetInstanceOnlyNameOk

`func (o *SharedDatasourceConfigAllOf) GetInstanceOnlyNameOk() (*string, bool)`

GetInstanceOnlyNameOk returns a tuple with the InstanceOnlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceOnlyName

`func (o *SharedDatasourceConfigAllOf) SetInstanceOnlyName(v string)`

SetInstanceOnlyName sets InstanceOnlyName field to given value.

### HasInstanceOnlyName

`func (o *SharedDatasourceConfigAllOf) HasInstanceOnlyName() bool`

HasInstanceOnlyName returns a boolean if a field has been set.

### GetInstanceDescription

`func (o *SharedDatasourceConfigAllOf) GetInstanceDescription() string`

GetInstanceDescription returns the InstanceDescription field if non-nil, zero value otherwise.

### GetInstanceDescriptionOk

`func (o *SharedDatasourceConfigAllOf) GetInstanceDescriptionOk() (*string, bool)`

GetInstanceDescriptionOk returns a tuple with the InstanceDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceDescription

`func (o *SharedDatasourceConfigAllOf) SetInstanceDescription(v string)`

SetInstanceDescription sets InstanceDescription field to given value.

### HasInstanceDescription

`func (o *SharedDatasourceConfigAllOf) HasInstanceDescription() bool`

HasInstanceDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


