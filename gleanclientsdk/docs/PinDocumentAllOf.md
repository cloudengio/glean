# PinDocumentAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The opaque id of the pin. | [optional] 
**PinId** | Pointer to **int32** | DEPRECATED - Prefer use of &#x60;id&#x60; string field instead | [optional] 
**DocumentId** | **string** | The document which should be a pinned result. | 
**Query** | Pointer to **string** | DEPRECATED - The query string for which the result was generated. | [optional] 
**AudienceFilters** | Pointer to [**[]FacetFilter**](FacetFilter.md) | Filters which restrict who should see the pinned document. Values are taken from the corresponding filters in people search. | [optional] 
**Attribution** | Pointer to [**Person**](Person.md) |  | [optional] 
**UpdatedBy** | Pointer to [**Person**](Person.md) |  | [optional] 
**CreateTime** | Pointer to **time.Time** |  | [optional] 
**UpdateTime** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewPinDocumentAllOf

`func NewPinDocumentAllOf(documentId string, ) *PinDocumentAllOf`

NewPinDocumentAllOf instantiates a new PinDocumentAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPinDocumentAllOfWithDefaults

`func NewPinDocumentAllOfWithDefaults() *PinDocumentAllOf`

NewPinDocumentAllOfWithDefaults instantiates a new PinDocumentAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PinDocumentAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PinDocumentAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PinDocumentAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PinDocumentAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPinId

`func (o *PinDocumentAllOf) GetPinId() int32`

GetPinId returns the PinId field if non-nil, zero value otherwise.

### GetPinIdOk

`func (o *PinDocumentAllOf) GetPinIdOk() (*int32, bool)`

GetPinIdOk returns a tuple with the PinId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinId

`func (o *PinDocumentAllOf) SetPinId(v int32)`

SetPinId sets PinId field to given value.

### HasPinId

`func (o *PinDocumentAllOf) HasPinId() bool`

HasPinId returns a boolean if a field has been set.

### GetDocumentId

`func (o *PinDocumentAllOf) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *PinDocumentAllOf) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *PinDocumentAllOf) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.


### GetQuery

`func (o *PinDocumentAllOf) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *PinDocumentAllOf) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *PinDocumentAllOf) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *PinDocumentAllOf) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetAudienceFilters

`func (o *PinDocumentAllOf) GetAudienceFilters() []FacetFilter`

GetAudienceFilters returns the AudienceFilters field if non-nil, zero value otherwise.

### GetAudienceFiltersOk

`func (o *PinDocumentAllOf) GetAudienceFiltersOk() (*[]FacetFilter, bool)`

GetAudienceFiltersOk returns a tuple with the AudienceFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudienceFilters

`func (o *PinDocumentAllOf) SetAudienceFilters(v []FacetFilter)`

SetAudienceFilters sets AudienceFilters field to given value.

### HasAudienceFilters

`func (o *PinDocumentAllOf) HasAudienceFilters() bool`

HasAudienceFilters returns a boolean if a field has been set.

### GetAttribution

`func (o *PinDocumentAllOf) GetAttribution() Person`

GetAttribution returns the Attribution field if non-nil, zero value otherwise.

### GetAttributionOk

`func (o *PinDocumentAllOf) GetAttributionOk() (*Person, bool)`

GetAttributionOk returns a tuple with the Attribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribution

`func (o *PinDocumentAllOf) SetAttribution(v Person)`

SetAttribution sets Attribution field to given value.

### HasAttribution

`func (o *PinDocumentAllOf) HasAttribution() bool`

HasAttribution returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *PinDocumentAllOf) GetUpdatedBy() Person`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *PinDocumentAllOf) GetUpdatedByOk() (*Person, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *PinDocumentAllOf) SetUpdatedBy(v Person)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *PinDocumentAllOf) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetCreateTime

`func (o *PinDocumentAllOf) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *PinDocumentAllOf) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *PinDocumentAllOf) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *PinDocumentAllOf) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetUpdateTime

`func (o *PinDocumentAllOf) GetUpdateTime() time.Time`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *PinDocumentAllOf) GetUpdateTimeOk() (*time.Time, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *PinDocumentAllOf) SetUpdateTime(v time.Time)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *PinDocumentAllOf) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


