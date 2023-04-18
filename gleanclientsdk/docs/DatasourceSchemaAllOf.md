# DatasourceSchemaAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CalendarEnabled** | Pointer to **bool** | Whether this datasource will show calendar feed events | [optional] 
**CollectBodyInActivity** | Pointer to **bool** | Whether the HTML body should be reported in ActivityEvents. | [optional] 
**CrawlHomeUrl** | Pointer to **bool** | Whether to crawl site starting at the home URL and all crawler seed URLs. | [optional] 
**DatasourceDescription** | Pointer to **string** | The support Glean provides for this datasource. | [optional] 
**DatasourceDisplayName** | Pointer to **string** | The user label for the datasource. This is distinct from the display names of the instances. | [optional] 
**FederatedGoogleOAuth** | Pointer to [**OAuthConfig**](OAuthConfig.md) |  | [optional] 
**IsEligibleForNativeReplacement** | Pointer to **bool** | Whether the datasource supports the native search replacement feature. | [optional] 
**IsEnabled** | Pointer to **bool** | Whether the datasource is set in queryapi.datasources. | [optional] 
**IsNativeReplacementAutoEnabled** | Pointer to **bool** | Make Native Search Replace opt-out rather than opt-in. | [optional] 
**IsSearchable** | Pointer to **bool** | Whether the user is able to search for results from this datasource. | [optional] 
**MentionsEnabled** | Pointer to **bool** | Whether the the datasource is configured to be used within mentions features such as the mentions feed | [optional] 
**PrivateAuth** | Pointer to [**PrivateAuth**](PrivateAuth.md) |  | [optional] 
**SearchInputSelectors** | Pointer to **[]string** | CSS selectors that identify native search input elements on the datasource&#39;s website. | [optional] 
**SupportsCalendar** | Pointer to **bool** | Whether the datasource can power the user&#39;s calendar. | [optional] 
**UrlRegexVersion** | Pointer to **int32** | The current version of the urlRegex, changes would instruct the client to clear any information based on previous versions of the urlRegex. | [optional] 

## Methods

### NewDatasourceSchemaAllOf

`func NewDatasourceSchemaAllOf() *DatasourceSchemaAllOf`

NewDatasourceSchemaAllOf instantiates a new DatasourceSchemaAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatasourceSchemaAllOfWithDefaults

`func NewDatasourceSchemaAllOfWithDefaults() *DatasourceSchemaAllOf`

NewDatasourceSchemaAllOfWithDefaults instantiates a new DatasourceSchemaAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCalendarEnabled

`func (o *DatasourceSchemaAllOf) GetCalendarEnabled() bool`

GetCalendarEnabled returns the CalendarEnabled field if non-nil, zero value otherwise.

### GetCalendarEnabledOk

`func (o *DatasourceSchemaAllOf) GetCalendarEnabledOk() (*bool, bool)`

GetCalendarEnabledOk returns a tuple with the CalendarEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalendarEnabled

`func (o *DatasourceSchemaAllOf) SetCalendarEnabled(v bool)`

SetCalendarEnabled sets CalendarEnabled field to given value.

### HasCalendarEnabled

`func (o *DatasourceSchemaAllOf) HasCalendarEnabled() bool`

HasCalendarEnabled returns a boolean if a field has been set.

### GetCollectBodyInActivity

`func (o *DatasourceSchemaAllOf) GetCollectBodyInActivity() bool`

GetCollectBodyInActivity returns the CollectBodyInActivity field if non-nil, zero value otherwise.

### GetCollectBodyInActivityOk

`func (o *DatasourceSchemaAllOf) GetCollectBodyInActivityOk() (*bool, bool)`

GetCollectBodyInActivityOk returns a tuple with the CollectBodyInActivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectBodyInActivity

`func (o *DatasourceSchemaAllOf) SetCollectBodyInActivity(v bool)`

SetCollectBodyInActivity sets CollectBodyInActivity field to given value.

### HasCollectBodyInActivity

`func (o *DatasourceSchemaAllOf) HasCollectBodyInActivity() bool`

HasCollectBodyInActivity returns a boolean if a field has been set.

### GetCrawlHomeUrl

`func (o *DatasourceSchemaAllOf) GetCrawlHomeUrl() bool`

GetCrawlHomeUrl returns the CrawlHomeUrl field if non-nil, zero value otherwise.

### GetCrawlHomeUrlOk

`func (o *DatasourceSchemaAllOf) GetCrawlHomeUrlOk() (*bool, bool)`

GetCrawlHomeUrlOk returns a tuple with the CrawlHomeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrawlHomeUrl

`func (o *DatasourceSchemaAllOf) SetCrawlHomeUrl(v bool)`

SetCrawlHomeUrl sets CrawlHomeUrl field to given value.

### HasCrawlHomeUrl

`func (o *DatasourceSchemaAllOf) HasCrawlHomeUrl() bool`

HasCrawlHomeUrl returns a boolean if a field has been set.

### GetDatasourceDescription

`func (o *DatasourceSchemaAllOf) GetDatasourceDescription() string`

GetDatasourceDescription returns the DatasourceDescription field if non-nil, zero value otherwise.

### GetDatasourceDescriptionOk

`func (o *DatasourceSchemaAllOf) GetDatasourceDescriptionOk() (*string, bool)`

GetDatasourceDescriptionOk returns a tuple with the DatasourceDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasourceDescription

`func (o *DatasourceSchemaAllOf) SetDatasourceDescription(v string)`

SetDatasourceDescription sets DatasourceDescription field to given value.

### HasDatasourceDescription

`func (o *DatasourceSchemaAllOf) HasDatasourceDescription() bool`

HasDatasourceDescription returns a boolean if a field has been set.

### GetDatasourceDisplayName

`func (o *DatasourceSchemaAllOf) GetDatasourceDisplayName() string`

GetDatasourceDisplayName returns the DatasourceDisplayName field if non-nil, zero value otherwise.

### GetDatasourceDisplayNameOk

`func (o *DatasourceSchemaAllOf) GetDatasourceDisplayNameOk() (*string, bool)`

GetDatasourceDisplayNameOk returns a tuple with the DatasourceDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasourceDisplayName

`func (o *DatasourceSchemaAllOf) SetDatasourceDisplayName(v string)`

SetDatasourceDisplayName sets DatasourceDisplayName field to given value.

### HasDatasourceDisplayName

`func (o *DatasourceSchemaAllOf) HasDatasourceDisplayName() bool`

HasDatasourceDisplayName returns a boolean if a field has been set.

### GetFederatedGoogleOAuth

`func (o *DatasourceSchemaAllOf) GetFederatedGoogleOAuth() OAuthConfig`

GetFederatedGoogleOAuth returns the FederatedGoogleOAuth field if non-nil, zero value otherwise.

### GetFederatedGoogleOAuthOk

`func (o *DatasourceSchemaAllOf) GetFederatedGoogleOAuthOk() (*OAuthConfig, bool)`

GetFederatedGoogleOAuthOk returns a tuple with the FederatedGoogleOAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederatedGoogleOAuth

`func (o *DatasourceSchemaAllOf) SetFederatedGoogleOAuth(v OAuthConfig)`

SetFederatedGoogleOAuth sets FederatedGoogleOAuth field to given value.

### HasFederatedGoogleOAuth

`func (o *DatasourceSchemaAllOf) HasFederatedGoogleOAuth() bool`

HasFederatedGoogleOAuth returns a boolean if a field has been set.

### GetIsEligibleForNativeReplacement

`func (o *DatasourceSchemaAllOf) GetIsEligibleForNativeReplacement() bool`

GetIsEligibleForNativeReplacement returns the IsEligibleForNativeReplacement field if non-nil, zero value otherwise.

### GetIsEligibleForNativeReplacementOk

`func (o *DatasourceSchemaAllOf) GetIsEligibleForNativeReplacementOk() (*bool, bool)`

GetIsEligibleForNativeReplacementOk returns a tuple with the IsEligibleForNativeReplacement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEligibleForNativeReplacement

`func (o *DatasourceSchemaAllOf) SetIsEligibleForNativeReplacement(v bool)`

SetIsEligibleForNativeReplacement sets IsEligibleForNativeReplacement field to given value.

### HasIsEligibleForNativeReplacement

`func (o *DatasourceSchemaAllOf) HasIsEligibleForNativeReplacement() bool`

HasIsEligibleForNativeReplacement returns a boolean if a field has been set.

### GetIsEnabled

`func (o *DatasourceSchemaAllOf) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *DatasourceSchemaAllOf) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *DatasourceSchemaAllOf) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *DatasourceSchemaAllOf) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetIsNativeReplacementAutoEnabled

`func (o *DatasourceSchemaAllOf) GetIsNativeReplacementAutoEnabled() bool`

GetIsNativeReplacementAutoEnabled returns the IsNativeReplacementAutoEnabled field if non-nil, zero value otherwise.

### GetIsNativeReplacementAutoEnabledOk

`func (o *DatasourceSchemaAllOf) GetIsNativeReplacementAutoEnabledOk() (*bool, bool)`

GetIsNativeReplacementAutoEnabledOk returns a tuple with the IsNativeReplacementAutoEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNativeReplacementAutoEnabled

`func (o *DatasourceSchemaAllOf) SetIsNativeReplacementAutoEnabled(v bool)`

SetIsNativeReplacementAutoEnabled sets IsNativeReplacementAutoEnabled field to given value.

### HasIsNativeReplacementAutoEnabled

`func (o *DatasourceSchemaAllOf) HasIsNativeReplacementAutoEnabled() bool`

HasIsNativeReplacementAutoEnabled returns a boolean if a field has been set.

### GetIsSearchable

`func (o *DatasourceSchemaAllOf) GetIsSearchable() bool`

GetIsSearchable returns the IsSearchable field if non-nil, zero value otherwise.

### GetIsSearchableOk

`func (o *DatasourceSchemaAllOf) GetIsSearchableOk() (*bool, bool)`

GetIsSearchableOk returns a tuple with the IsSearchable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSearchable

`func (o *DatasourceSchemaAllOf) SetIsSearchable(v bool)`

SetIsSearchable sets IsSearchable field to given value.

### HasIsSearchable

`func (o *DatasourceSchemaAllOf) HasIsSearchable() bool`

HasIsSearchable returns a boolean if a field has been set.

### GetMentionsEnabled

`func (o *DatasourceSchemaAllOf) GetMentionsEnabled() bool`

GetMentionsEnabled returns the MentionsEnabled field if non-nil, zero value otherwise.

### GetMentionsEnabledOk

`func (o *DatasourceSchemaAllOf) GetMentionsEnabledOk() (*bool, bool)`

GetMentionsEnabledOk returns a tuple with the MentionsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentionsEnabled

`func (o *DatasourceSchemaAllOf) SetMentionsEnabled(v bool)`

SetMentionsEnabled sets MentionsEnabled field to given value.

### HasMentionsEnabled

`func (o *DatasourceSchemaAllOf) HasMentionsEnabled() bool`

HasMentionsEnabled returns a boolean if a field has been set.

### GetPrivateAuth

`func (o *DatasourceSchemaAllOf) GetPrivateAuth() PrivateAuth`

GetPrivateAuth returns the PrivateAuth field if non-nil, zero value otherwise.

### GetPrivateAuthOk

`func (o *DatasourceSchemaAllOf) GetPrivateAuthOk() (*PrivateAuth, bool)`

GetPrivateAuthOk returns a tuple with the PrivateAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateAuth

`func (o *DatasourceSchemaAllOf) SetPrivateAuth(v PrivateAuth)`

SetPrivateAuth sets PrivateAuth field to given value.

### HasPrivateAuth

`func (o *DatasourceSchemaAllOf) HasPrivateAuth() bool`

HasPrivateAuth returns a boolean if a field has been set.

### GetSearchInputSelectors

`func (o *DatasourceSchemaAllOf) GetSearchInputSelectors() []string`

GetSearchInputSelectors returns the SearchInputSelectors field if non-nil, zero value otherwise.

### GetSearchInputSelectorsOk

`func (o *DatasourceSchemaAllOf) GetSearchInputSelectorsOk() (*[]string, bool)`

GetSearchInputSelectorsOk returns a tuple with the SearchInputSelectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchInputSelectors

`func (o *DatasourceSchemaAllOf) SetSearchInputSelectors(v []string)`

SetSearchInputSelectors sets SearchInputSelectors field to given value.

### HasSearchInputSelectors

`func (o *DatasourceSchemaAllOf) HasSearchInputSelectors() bool`

HasSearchInputSelectors returns a boolean if a field has been set.

### GetSupportsCalendar

`func (o *DatasourceSchemaAllOf) GetSupportsCalendar() bool`

GetSupportsCalendar returns the SupportsCalendar field if non-nil, zero value otherwise.

### GetSupportsCalendarOk

`func (o *DatasourceSchemaAllOf) GetSupportsCalendarOk() (*bool, bool)`

GetSupportsCalendarOk returns a tuple with the SupportsCalendar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsCalendar

`func (o *DatasourceSchemaAllOf) SetSupportsCalendar(v bool)`

SetSupportsCalendar sets SupportsCalendar field to given value.

### HasSupportsCalendar

`func (o *DatasourceSchemaAllOf) HasSupportsCalendar() bool`

HasSupportsCalendar returns a boolean if a field has been set.

### GetUrlRegexVersion

`func (o *DatasourceSchemaAllOf) GetUrlRegexVersion() int32`

GetUrlRegexVersion returns the UrlRegexVersion field if non-nil, zero value otherwise.

### GetUrlRegexVersionOk

`func (o *DatasourceSchemaAllOf) GetUrlRegexVersionOk() (*int32, bool)`

GetUrlRegexVersionOk returns a tuple with the UrlRegexVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlRegexVersion

`func (o *DatasourceSchemaAllOf) SetUrlRegexVersion(v int32)`

SetUrlRegexVersion sets UrlRegexVersion field to given value.

### HasUrlRegexVersion

`func (o *DatasourceSchemaAllOf) HasUrlRegexVersion() bool`

HasUrlRegexVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


