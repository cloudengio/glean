# DisplayableListConfigAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | Primary title for the list | [optional] 
**Enabled** | Pointer to **bool** | Whether the list should be shown to the user | [optional] 
**AudienceFilters** | Pointer to [**[]FacetFilter**](FacetFilter.md) | Filters which restrict who should should see displayable list | [optional] 
**ItemUIConfig** | Pointer to [**DisplayableListItemUIConfig**](DisplayableListItemUIConfig.md) |  | [optional] 

## Methods

### NewDisplayableListConfigAllOf

`func NewDisplayableListConfigAllOf() *DisplayableListConfigAllOf`

NewDisplayableListConfigAllOf instantiates a new DisplayableListConfigAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDisplayableListConfigAllOfWithDefaults

`func NewDisplayableListConfigAllOfWithDefaults() *DisplayableListConfigAllOf`

NewDisplayableListConfigAllOfWithDefaults instantiates a new DisplayableListConfigAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *DisplayableListConfigAllOf) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DisplayableListConfigAllOf) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DisplayableListConfigAllOf) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *DisplayableListConfigAllOf) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetEnabled

`func (o *DisplayableListConfigAllOf) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DisplayableListConfigAllOf) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DisplayableListConfigAllOf) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *DisplayableListConfigAllOf) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAudienceFilters

`func (o *DisplayableListConfigAllOf) GetAudienceFilters() []FacetFilter`

GetAudienceFilters returns the AudienceFilters field if non-nil, zero value otherwise.

### GetAudienceFiltersOk

`func (o *DisplayableListConfigAllOf) GetAudienceFiltersOk() (*[]FacetFilter, bool)`

GetAudienceFiltersOk returns a tuple with the AudienceFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudienceFilters

`func (o *DisplayableListConfigAllOf) SetAudienceFilters(v []FacetFilter)`

SetAudienceFilters sets AudienceFilters field to given value.

### HasAudienceFilters

`func (o *DisplayableListConfigAllOf) HasAudienceFilters() bool`

HasAudienceFilters returns a boolean if a field has been set.

### GetItemUIConfig

`func (o *DisplayableListConfigAllOf) GetItemUIConfig() DisplayableListItemUIConfig`

GetItemUIConfig returns the ItemUIConfig field if non-nil, zero value otherwise.

### GetItemUIConfigOk

`func (o *DisplayableListConfigAllOf) GetItemUIConfigOk() (*DisplayableListItemUIConfig, bool)`

GetItemUIConfigOk returns a tuple with the ItemUIConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemUIConfig

`func (o *DisplayableListConfigAllOf) SetItemUIConfig(v DisplayableListItemUIConfig)`

SetItemUIConfig sets ItemUIConfig field to given value.

### HasItemUIConfig

`func (o *DisplayableListConfigAllOf) HasItemUIConfig() bool`

HasItemUIConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


