# ProcessAllDocumentsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datasource** | Pointer to **string** | If provided, process documents only for this custom datasource. Otherwise all uploaded documents are processed. | [optional] 

## Methods

### NewProcessAllDocumentsRequest

`func NewProcessAllDocumentsRequest() *ProcessAllDocumentsRequest`

NewProcessAllDocumentsRequest instantiates a new ProcessAllDocumentsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessAllDocumentsRequestWithDefaults

`func NewProcessAllDocumentsRequestWithDefaults() *ProcessAllDocumentsRequest`

NewProcessAllDocumentsRequestWithDefaults instantiates a new ProcessAllDocumentsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatasource

`func (o *ProcessAllDocumentsRequest) GetDatasource() string`

GetDatasource returns the Datasource field if non-nil, zero value otherwise.

### GetDatasourceOk

`func (o *ProcessAllDocumentsRequest) GetDatasourceOk() (*string, bool)`

GetDatasourceOk returns a tuple with the Datasource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasource

`func (o *ProcessAllDocumentsRequest) SetDatasource(v string)`

SetDatasource sets Datasource field to given value.

### HasDatasource

`func (o *ProcessAllDocumentsRequest) HasDatasource() bool`

HasDatasource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


