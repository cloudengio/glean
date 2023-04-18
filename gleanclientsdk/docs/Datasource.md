# Datasource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The required id of the datasource. Example: zendesk | [optional] 
**Capabilities** | Pointer to [**[]DatasourceCapability**](DatasourceCapability.md) | The functionality provided by the datasource, such as providing searchable content or SSO access. | [optional] 

## Methods

### NewDatasource

`func NewDatasource() *Datasource`

NewDatasource instantiates a new Datasource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatasourceWithDefaults

`func NewDatasourceWithDefaults() *Datasource`

NewDatasourceWithDefaults instantiates a new Datasource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Datasource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Datasource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Datasource) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Datasource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCapabilities

`func (o *Datasource) GetCapabilities() []DatasourceCapability`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *Datasource) GetCapabilitiesOk() (*[]DatasourceCapability, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *Datasource) SetCapabilities(v []DatasourceCapability)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *Datasource) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


