# SearchResponseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**SearchResponseMetadata**](SearchResponseMetadata.md) |  | [optional] 
**FacetResults** | Pointer to [**[]FacetResult**](FacetResult.md) |  | [optional] 
**ResultTabs** | Pointer to [**[]ResultTab**](ResultTab.md) | All result tabs available for the current query. Populated if QUERY_METADATA is specified in the request. | [optional] 
**ResultTabIds** | Pointer to **[]string** | The unique IDs of the result tabs to which this response belongs. | [optional] 
**ResultsDescription** | Pointer to [**ResultsDescription**](ResultsDescription.md) |  | [optional] 
**RewrittenFacetFilters** | Pointer to [**[]FacetFilter**](FacetFilter.md) | The actual applied facet filters based on the operators and facetFilters in the query. Useful for mapping typed operators to visual facets. | [optional] 
**Cursor** | Pointer to **string** | Cursor that indicates the start of the next page of results. To be passed in \&quot;more\&quot; requests for this query. | [optional] 
**HasMoreResults** | Pointer to **bool** | Whether more results are available. Use cursor to retrieve them. | [optional] 

## Methods

### NewSearchResponseAllOf

`func NewSearchResponseAllOf() *SearchResponseAllOf`

NewSearchResponseAllOf instantiates a new SearchResponseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchResponseAllOfWithDefaults

`func NewSearchResponseAllOfWithDefaults() *SearchResponseAllOf`

NewSearchResponseAllOfWithDefaults instantiates a new SearchResponseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *SearchResponseAllOf) GetMetadata() SearchResponseMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SearchResponseAllOf) GetMetadataOk() (*SearchResponseMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SearchResponseAllOf) SetMetadata(v SearchResponseMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SearchResponseAllOf) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetFacetResults

`func (o *SearchResponseAllOf) GetFacetResults() []FacetResult`

GetFacetResults returns the FacetResults field if non-nil, zero value otherwise.

### GetFacetResultsOk

`func (o *SearchResponseAllOf) GetFacetResultsOk() (*[]FacetResult, bool)`

GetFacetResultsOk returns a tuple with the FacetResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacetResults

`func (o *SearchResponseAllOf) SetFacetResults(v []FacetResult)`

SetFacetResults sets FacetResults field to given value.

### HasFacetResults

`func (o *SearchResponseAllOf) HasFacetResults() bool`

HasFacetResults returns a boolean if a field has been set.

### GetResultTabs

`func (o *SearchResponseAllOf) GetResultTabs() []ResultTab`

GetResultTabs returns the ResultTabs field if non-nil, zero value otherwise.

### GetResultTabsOk

`func (o *SearchResponseAllOf) GetResultTabsOk() (*[]ResultTab, bool)`

GetResultTabsOk returns a tuple with the ResultTabs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultTabs

`func (o *SearchResponseAllOf) SetResultTabs(v []ResultTab)`

SetResultTabs sets ResultTabs field to given value.

### HasResultTabs

`func (o *SearchResponseAllOf) HasResultTabs() bool`

HasResultTabs returns a boolean if a field has been set.

### GetResultTabIds

`func (o *SearchResponseAllOf) GetResultTabIds() []string`

GetResultTabIds returns the ResultTabIds field if non-nil, zero value otherwise.

### GetResultTabIdsOk

`func (o *SearchResponseAllOf) GetResultTabIdsOk() (*[]string, bool)`

GetResultTabIdsOk returns a tuple with the ResultTabIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultTabIds

`func (o *SearchResponseAllOf) SetResultTabIds(v []string)`

SetResultTabIds sets ResultTabIds field to given value.

### HasResultTabIds

`func (o *SearchResponseAllOf) HasResultTabIds() bool`

HasResultTabIds returns a boolean if a field has been set.

### GetResultsDescription

`func (o *SearchResponseAllOf) GetResultsDescription() ResultsDescription`

GetResultsDescription returns the ResultsDescription field if non-nil, zero value otherwise.

### GetResultsDescriptionOk

`func (o *SearchResponseAllOf) GetResultsDescriptionOk() (*ResultsDescription, bool)`

GetResultsDescriptionOk returns a tuple with the ResultsDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsDescription

`func (o *SearchResponseAllOf) SetResultsDescription(v ResultsDescription)`

SetResultsDescription sets ResultsDescription field to given value.

### HasResultsDescription

`func (o *SearchResponseAllOf) HasResultsDescription() bool`

HasResultsDescription returns a boolean if a field has been set.

### GetRewrittenFacetFilters

`func (o *SearchResponseAllOf) GetRewrittenFacetFilters() []FacetFilter`

GetRewrittenFacetFilters returns the RewrittenFacetFilters field if non-nil, zero value otherwise.

### GetRewrittenFacetFiltersOk

`func (o *SearchResponseAllOf) GetRewrittenFacetFiltersOk() (*[]FacetFilter, bool)`

GetRewrittenFacetFiltersOk returns a tuple with the RewrittenFacetFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewrittenFacetFilters

`func (o *SearchResponseAllOf) SetRewrittenFacetFilters(v []FacetFilter)`

SetRewrittenFacetFilters sets RewrittenFacetFilters field to given value.

### HasRewrittenFacetFilters

`func (o *SearchResponseAllOf) HasRewrittenFacetFilters() bool`

HasRewrittenFacetFilters returns a boolean if a field has been set.

### GetCursor

`func (o *SearchResponseAllOf) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *SearchResponseAllOf) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *SearchResponseAllOf) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *SearchResponseAllOf) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### GetHasMoreResults

`func (o *SearchResponseAllOf) GetHasMoreResults() bool`

GetHasMoreResults returns the HasMoreResults field if non-nil, zero value otherwise.

### GetHasMoreResultsOk

`func (o *SearchResponseAllOf) GetHasMoreResultsOk() (*bool, bool)`

GetHasMoreResultsOk returns a tuple with the HasMoreResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMoreResults

`func (o *SearchResponseAllOf) SetHasMoreResults(v bool)`

SetHasMoreResults sets HasMoreResults field to given value.

### HasHasMoreResults

`func (o *SearchResponseAllOf) HasHasMoreResults() bool`

HasHasMoreResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


