# DatasourceGroupDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | name of the group. Should be unique among all groups for the datasource, and cannot have spaces. | 

## Methods

### NewDatasourceGroupDefinition

`func NewDatasourceGroupDefinition(name string, ) *DatasourceGroupDefinition`

NewDatasourceGroupDefinition instantiates a new DatasourceGroupDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatasourceGroupDefinitionWithDefaults

`func NewDatasourceGroupDefinitionWithDefaults() *DatasourceGroupDefinition`

NewDatasourceGroupDefinitionWithDefaults instantiates a new DatasourceGroupDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DatasourceGroupDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DatasourceGroupDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DatasourceGroupDefinition) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


