# EditPinRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | Pointer to **string** | DEPRECATED - The query string to be set for the pin | [optional] 
**Queries** | Pointer to **[]string** | The query strings for which the pinned result will show. | [optional] 
**AudienceFilters** | Pointer to [**[]FacetFilter**](FacetFilter.md) | Filters which restrict who should see the pinned document. Values are taken from the corresponding filters in people search. | [optional] 
**PinId** | Pointer to **int32** | DEPRECATED - Prefer use of &#x60;id&#x60; | [optional] 
**Id** | Pointer to **string** | The opaque id of the pin to be edited | [optional] 

## Methods

### NewEditPinRequest

`func NewEditPinRequest() *EditPinRequest`

NewEditPinRequest instantiates a new EditPinRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditPinRequestWithDefaults

`func NewEditPinRequestWithDefaults() *EditPinRequest`

NewEditPinRequestWithDefaults instantiates a new EditPinRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *EditPinRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *EditPinRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *EditPinRequest) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *EditPinRequest) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetQueries

`func (o *EditPinRequest) GetQueries() []string`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *EditPinRequest) GetQueriesOk() (*[]string, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *EditPinRequest) SetQueries(v []string)`

SetQueries sets Queries field to given value.

### HasQueries

`func (o *EditPinRequest) HasQueries() bool`

HasQueries returns a boolean if a field has been set.

### GetAudienceFilters

`func (o *EditPinRequest) GetAudienceFilters() []FacetFilter`

GetAudienceFilters returns the AudienceFilters field if non-nil, zero value otherwise.

### GetAudienceFiltersOk

`func (o *EditPinRequest) GetAudienceFiltersOk() (*[]FacetFilter, bool)`

GetAudienceFiltersOk returns a tuple with the AudienceFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudienceFilters

`func (o *EditPinRequest) SetAudienceFilters(v []FacetFilter)`

SetAudienceFilters sets AudienceFilters field to given value.

### HasAudienceFilters

`func (o *EditPinRequest) HasAudienceFilters() bool`

HasAudienceFilters returns a boolean if a field has been set.

### GetPinId

`func (o *EditPinRequest) GetPinId() int32`

GetPinId returns the PinId field if non-nil, zero value otherwise.

### GetPinIdOk

`func (o *EditPinRequest) GetPinIdOk() (*int32, bool)`

GetPinIdOk returns a tuple with the PinId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinId

`func (o *EditPinRequest) SetPinId(v int32)`

SetPinId sets PinId field to given value.

### HasPinId

`func (o *EditPinRequest) HasPinId() bool`

HasPinId returns a boolean if a field has been set.

### GetId

`func (o *EditPinRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EditPinRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EditPinRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EditPinRequest) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


