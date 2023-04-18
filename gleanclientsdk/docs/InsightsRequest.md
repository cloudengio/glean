# InsightsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Categories** | **[]string** | Categories of data requested. Request can include single or multiple types. | 
**Departments** | Pointer to **[]string** | Departments that the data is requested for. If the empty, corresponds to whole company. | [optional] 
**DayRange** | Pointer to [**Period**](Period.md) |  | [optional] 

## Methods

### NewInsightsRequest

`func NewInsightsRequest(categories []string, ) *InsightsRequest`

NewInsightsRequest instantiates a new InsightsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInsightsRequestWithDefaults

`func NewInsightsRequestWithDefaults() *InsightsRequest`

NewInsightsRequestWithDefaults instantiates a new InsightsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategories

`func (o *InsightsRequest) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *InsightsRequest) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *InsightsRequest) SetCategories(v []string)`

SetCategories sets Categories field to given value.


### GetDepartments

`func (o *InsightsRequest) GetDepartments() []string`

GetDepartments returns the Departments field if non-nil, zero value otherwise.

### GetDepartmentsOk

`func (o *InsightsRequest) GetDepartmentsOk() (*[]string, bool)`

GetDepartmentsOk returns a tuple with the Departments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartments

`func (o *InsightsRequest) SetDepartments(v []string)`

SetDepartments sets Departments field to given value.

### HasDepartments

`func (o *InsightsRequest) HasDepartments() bool`

HasDepartments returns a boolean if a field has been set.

### GetDayRange

`func (o *InsightsRequest) GetDayRange() Period`

GetDayRange returns the DayRange field if non-nil, zero value otherwise.

### GetDayRangeOk

`func (o *InsightsRequest) GetDayRangeOk() (*Period, bool)`

GetDayRangeOk returns a tuple with the DayRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayRange

`func (o *InsightsRequest) SetDayRange(v Period)`

SetDayRange sets DayRange field to given value.

### HasDayRange

`func (o *InsightsRequest) HasDayRange() bool`

HasDayRange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


