# SectionsConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | Pointer to [**[]DatumCondition**](DatumCondition.md) |  | [optional] 

## Methods

### NewSectionsConfig

`func NewSectionsConfig() *SectionsConfig`

NewSectionsConfig instantiates a new SectionsConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSectionsConfigWithDefaults

`func NewSectionsConfigWithDefaults() *SectionsConfig`

NewSectionsConfigWithDefaults instantiates a new SectionsConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditions

`func (o *SectionsConfig) GetConditions() []DatumCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *SectionsConfig) GetConditionsOk() (*[]DatumCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *SectionsConfig) SetConditions(v []DatumCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *SectionsConfig) HasConditions() bool`

HasConditions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


