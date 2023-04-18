# DatasourceSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Unique identifier of datasource instance to which this config applies. | 
**DisplayName** | Pointer to **string** | The user-friendly instance label to display. If omitted, falls back to the title-cased &#x60;name&#x60;. | [optional] 
**DatasourceCategory** | Pointer to **string** | The type of this datasource. It is an important signal for relevance and must be specified and cannot be UNCATEGORIZED. | [optional] [default to "UNCATEGORIZED"]
**UrlRegex** | Pointer to **string** | Regular expression that matches URLs of documents of the datasource instance. The behavior for multiple matches is non-deterministic. **Note: urlRegex is a required field for non-entity datasources (ie. datasources where isEntityDatasource is false). Please add a regex as specific as possible to this datasource instance.** | [optional] 
**IconUrl** | Pointer to **string** | The URL to an image to be displayed as an icon for this datasource instance. Must have a transparency mask. SVG are recommended over PNG. Public, scio-authenticated and Base64 encoded data URLs are all valid (but not third-party-authenticated URLs). | [optional] 
**ObjectDefinitions** | Pointer to [**[]ObjectDefinition**](ObjectDefinition.md) | The list of top-level &#x60;objectType&#x60;s for the datasource. | [optional] 
**SuggestionText** | Pointer to **string** | Example text for what to search for in this datasource | [optional] 
**HomeUrl** | Pointer to **string** | The URL of the landing page for this datasource instance. Should point to the most useful page for users, not the company marketing page. | [optional] 
**CrawlerSeedUrls** | Pointer to **[]string** | This only applies to WEB_CRAWL and BROWSER_CRAWL datasources. Defines the seed urls for crawling. | [optional] 
**IconDarkUrl** | Pointer to **string** | The URL to an image to be displayed as an icon for this datasource instance in dark mode. Must have a transparency mask. SVG are recommended over PNG. Public, scio-authenticated and Base64 encoded data URLs are all valid (but not third-party-authenticated URLs). | [optional] 
**HideBuiltInFacets** | Pointer to **[]string** | List of built-in facet types that should be hidden for the datasource. | [optional] 
**CanonicalizingURLRegex** | Pointer to [**[]CanonicalizingRegexType**](CanonicalizingRegexType.md) | A list of regular expressions to apply to an arbitrary URL to transform it into a canonical URL for this datasource instance. Regexes are to be applied in the order specified in this list. | [optional] 
**CanonicalizingTitleRegex** | Pointer to [**[]CanonicalizingRegexType**](CanonicalizingRegexType.md) | A list of regular expressions to apply to an arbitrary title to transform it into a title that will be displayed in the search results | [optional] 
**RedlistTitleRegex** | Pointer to **string** | A regex that identifies titles that should not be indexed | [optional] 
**ConnectorType** | Pointer to [**ConnectorType**](ConnectorType.md) |  | [optional] 
**Quicklinks** | Pointer to [**[]Quicklink**](Quicklink.md) | List of actions for this datasource instance that will show up in autocomplete and app card, e.g. \&quot;Create new issue\&quot; for jira | [optional] 
**RenderConfigPreset** | Pointer to **string** | The name of a render config to use for displaying results from this datasource. Any well known datasource name may be used to render the same as that source, e.g. &#x60;web&#x60; or &#x60;gdrive&#x60;. | [optional] 
**Aliases** | Pointer to **[]string** | Aliases that can be used as &#x60;app&#x60; operator-values. | [optional] 
**IsOnPrem** | Pointer to **bool** | Whether or not this datasource is hosted on-premise. | [optional] 
**TrustUrlRegexForViewActivity** | Pointer to **bool** | True if browser activity is able to report the correct URL for VIEW events. Set this to true if the URLs reported by Chrome are constant throughout each page load. Set this to false if the page has Javascript that modifies the URL during or after the load. | [optional] [default to true]
**IncludeUtmSource** | Pointer to **bool** | If true, a utm_source query param will be added to outbound links to this datasource within Glean. | [optional] 
**DatasourceName** | Pointer to **string** | The (non-unique) name of the datasource of which this config is an instance, e.g. \&quot;jira\&quot;. | [optional] 
**InstanceOnlyName** | Pointer to **string** | The instance of the datasource for this particular config, e.g. \&quot;onprem\&quot;. | [optional] 
**InstanceDescription** | Pointer to **string** | A human readable string identifying this instance as compared to its peers, e.g. \&quot;github.com/askscio\&quot; or \&quot;github.askscio.com\&quot; | [optional] 
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

### NewDatasourceSchema

`func NewDatasourceSchema(name string, ) *DatasourceSchema`

NewDatasourceSchema instantiates a new DatasourceSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatasourceSchemaWithDefaults

`func NewDatasourceSchemaWithDefaults() *DatasourceSchema`

NewDatasourceSchemaWithDefaults instantiates a new DatasourceSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DatasourceSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DatasourceSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DatasourceSchema) SetName(v string)`

SetName sets Name field to given value.


### GetDisplayName

`func (o *DatasourceSchema) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *DatasourceSchema) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *DatasourceSchema) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *DatasourceSchema) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDatasourceCategory

`func (o *DatasourceSchema) GetDatasourceCategory() string`

GetDatasourceCategory returns the DatasourceCategory field if non-nil, zero value otherwise.

### GetDatasourceCategoryOk

`func (o *DatasourceSchema) GetDatasourceCategoryOk() (*string, bool)`

GetDatasourceCategoryOk returns a tuple with the DatasourceCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasourceCategory

`func (o *DatasourceSchema) SetDatasourceCategory(v string)`

SetDatasourceCategory sets DatasourceCategory field to given value.

### HasDatasourceCategory

`func (o *DatasourceSchema) HasDatasourceCategory() bool`

HasDatasourceCategory returns a boolean if a field has been set.

### GetUrlRegex

`func (o *DatasourceSchema) GetUrlRegex() string`

GetUrlRegex returns the UrlRegex field if non-nil, zero value otherwise.

### GetUrlRegexOk

`func (o *DatasourceSchema) GetUrlRegexOk() (*string, bool)`

GetUrlRegexOk returns a tuple with the UrlRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlRegex

`func (o *DatasourceSchema) SetUrlRegex(v string)`

SetUrlRegex sets UrlRegex field to given value.

### HasUrlRegex

`func (o *DatasourceSchema) HasUrlRegex() bool`

HasUrlRegex returns a boolean if a field has been set.

### GetIconUrl

`func (o *DatasourceSchema) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *DatasourceSchema) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *DatasourceSchema) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *DatasourceSchema) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### GetObjectDefinitions

`func (o *DatasourceSchema) GetObjectDefinitions() []ObjectDefinition`

GetObjectDefinitions returns the ObjectDefinitions field if non-nil, zero value otherwise.

### GetObjectDefinitionsOk

`func (o *DatasourceSchema) GetObjectDefinitionsOk() (*[]ObjectDefinition, bool)`

GetObjectDefinitionsOk returns a tuple with the ObjectDefinitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectDefinitions

`func (o *DatasourceSchema) SetObjectDefinitions(v []ObjectDefinition)`

SetObjectDefinitions sets ObjectDefinitions field to given value.

### HasObjectDefinitions

`func (o *DatasourceSchema) HasObjectDefinitions() bool`

HasObjectDefinitions returns a boolean if a field has been set.

### GetSuggestionText

`func (o *DatasourceSchema) GetSuggestionText() string`

GetSuggestionText returns the SuggestionText field if non-nil, zero value otherwise.

### GetSuggestionTextOk

`func (o *DatasourceSchema) GetSuggestionTextOk() (*string, bool)`

GetSuggestionTextOk returns a tuple with the SuggestionText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestionText

`func (o *DatasourceSchema) SetSuggestionText(v string)`

SetSuggestionText sets SuggestionText field to given value.

### HasSuggestionText

`func (o *DatasourceSchema) HasSuggestionText() bool`

HasSuggestionText returns a boolean if a field has been set.

### GetHomeUrl

`func (o *DatasourceSchema) GetHomeUrl() string`

GetHomeUrl returns the HomeUrl field if non-nil, zero value otherwise.

### GetHomeUrlOk

`func (o *DatasourceSchema) GetHomeUrlOk() (*string, bool)`

GetHomeUrlOk returns a tuple with the HomeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeUrl

`func (o *DatasourceSchema) SetHomeUrl(v string)`

SetHomeUrl sets HomeUrl field to given value.

### HasHomeUrl

`func (o *DatasourceSchema) HasHomeUrl() bool`

HasHomeUrl returns a boolean if a field has been set.

### GetCrawlerSeedUrls

`func (o *DatasourceSchema) GetCrawlerSeedUrls() []string`

GetCrawlerSeedUrls returns the CrawlerSeedUrls field if non-nil, zero value otherwise.

### GetCrawlerSeedUrlsOk

`func (o *DatasourceSchema) GetCrawlerSeedUrlsOk() (*[]string, bool)`

GetCrawlerSeedUrlsOk returns a tuple with the CrawlerSeedUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrawlerSeedUrls

`func (o *DatasourceSchema) SetCrawlerSeedUrls(v []string)`

SetCrawlerSeedUrls sets CrawlerSeedUrls field to given value.

### HasCrawlerSeedUrls

`func (o *DatasourceSchema) HasCrawlerSeedUrls() bool`

HasCrawlerSeedUrls returns a boolean if a field has been set.

### GetIconDarkUrl

`func (o *DatasourceSchema) GetIconDarkUrl() string`

GetIconDarkUrl returns the IconDarkUrl field if non-nil, zero value otherwise.

### GetIconDarkUrlOk

`func (o *DatasourceSchema) GetIconDarkUrlOk() (*string, bool)`

GetIconDarkUrlOk returns a tuple with the IconDarkUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconDarkUrl

`func (o *DatasourceSchema) SetIconDarkUrl(v string)`

SetIconDarkUrl sets IconDarkUrl field to given value.

### HasIconDarkUrl

`func (o *DatasourceSchema) HasIconDarkUrl() bool`

HasIconDarkUrl returns a boolean if a field has been set.

### GetHideBuiltInFacets

`func (o *DatasourceSchema) GetHideBuiltInFacets() []string`

GetHideBuiltInFacets returns the HideBuiltInFacets field if non-nil, zero value otherwise.

### GetHideBuiltInFacetsOk

`func (o *DatasourceSchema) GetHideBuiltInFacetsOk() (*[]string, bool)`

GetHideBuiltInFacetsOk returns a tuple with the HideBuiltInFacets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideBuiltInFacets

`func (o *DatasourceSchema) SetHideBuiltInFacets(v []string)`

SetHideBuiltInFacets sets HideBuiltInFacets field to given value.

### HasHideBuiltInFacets

`func (o *DatasourceSchema) HasHideBuiltInFacets() bool`

HasHideBuiltInFacets returns a boolean if a field has been set.

### GetCanonicalizingURLRegex

`func (o *DatasourceSchema) GetCanonicalizingURLRegex() []CanonicalizingRegexType`

GetCanonicalizingURLRegex returns the CanonicalizingURLRegex field if non-nil, zero value otherwise.

### GetCanonicalizingURLRegexOk

`func (o *DatasourceSchema) GetCanonicalizingURLRegexOk() (*[]CanonicalizingRegexType, bool)`

GetCanonicalizingURLRegexOk returns a tuple with the CanonicalizingURLRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalizingURLRegex

`func (o *DatasourceSchema) SetCanonicalizingURLRegex(v []CanonicalizingRegexType)`

SetCanonicalizingURLRegex sets CanonicalizingURLRegex field to given value.

### HasCanonicalizingURLRegex

`func (o *DatasourceSchema) HasCanonicalizingURLRegex() bool`

HasCanonicalizingURLRegex returns a boolean if a field has been set.

### GetCanonicalizingTitleRegex

`func (o *DatasourceSchema) GetCanonicalizingTitleRegex() []CanonicalizingRegexType`

GetCanonicalizingTitleRegex returns the CanonicalizingTitleRegex field if non-nil, zero value otherwise.

### GetCanonicalizingTitleRegexOk

`func (o *DatasourceSchema) GetCanonicalizingTitleRegexOk() (*[]CanonicalizingRegexType, bool)`

GetCanonicalizingTitleRegexOk returns a tuple with the CanonicalizingTitleRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalizingTitleRegex

`func (o *DatasourceSchema) SetCanonicalizingTitleRegex(v []CanonicalizingRegexType)`

SetCanonicalizingTitleRegex sets CanonicalizingTitleRegex field to given value.

### HasCanonicalizingTitleRegex

`func (o *DatasourceSchema) HasCanonicalizingTitleRegex() bool`

HasCanonicalizingTitleRegex returns a boolean if a field has been set.

### GetRedlistTitleRegex

`func (o *DatasourceSchema) GetRedlistTitleRegex() string`

GetRedlistTitleRegex returns the RedlistTitleRegex field if non-nil, zero value otherwise.

### GetRedlistTitleRegexOk

`func (o *DatasourceSchema) GetRedlistTitleRegexOk() (*string, bool)`

GetRedlistTitleRegexOk returns a tuple with the RedlistTitleRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedlistTitleRegex

`func (o *DatasourceSchema) SetRedlistTitleRegex(v string)`

SetRedlistTitleRegex sets RedlistTitleRegex field to given value.

### HasRedlistTitleRegex

`func (o *DatasourceSchema) HasRedlistTitleRegex() bool`

HasRedlistTitleRegex returns a boolean if a field has been set.

### GetConnectorType

`func (o *DatasourceSchema) GetConnectorType() ConnectorType`

GetConnectorType returns the ConnectorType field if non-nil, zero value otherwise.

### GetConnectorTypeOk

`func (o *DatasourceSchema) GetConnectorTypeOk() (*ConnectorType, bool)`

GetConnectorTypeOk returns a tuple with the ConnectorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorType

`func (o *DatasourceSchema) SetConnectorType(v ConnectorType)`

SetConnectorType sets ConnectorType field to given value.

### HasConnectorType

`func (o *DatasourceSchema) HasConnectorType() bool`

HasConnectorType returns a boolean if a field has been set.

### GetQuicklinks

`func (o *DatasourceSchema) GetQuicklinks() []Quicklink`

GetQuicklinks returns the Quicklinks field if non-nil, zero value otherwise.

### GetQuicklinksOk

`func (o *DatasourceSchema) GetQuicklinksOk() (*[]Quicklink, bool)`

GetQuicklinksOk returns a tuple with the Quicklinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuicklinks

`func (o *DatasourceSchema) SetQuicklinks(v []Quicklink)`

SetQuicklinks sets Quicklinks field to given value.

### HasQuicklinks

`func (o *DatasourceSchema) HasQuicklinks() bool`

HasQuicklinks returns a boolean if a field has been set.

### GetRenderConfigPreset

`func (o *DatasourceSchema) GetRenderConfigPreset() string`

GetRenderConfigPreset returns the RenderConfigPreset field if non-nil, zero value otherwise.

### GetRenderConfigPresetOk

`func (o *DatasourceSchema) GetRenderConfigPresetOk() (*string, bool)`

GetRenderConfigPresetOk returns a tuple with the RenderConfigPreset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenderConfigPreset

`func (o *DatasourceSchema) SetRenderConfigPreset(v string)`

SetRenderConfigPreset sets RenderConfigPreset field to given value.

### HasRenderConfigPreset

`func (o *DatasourceSchema) HasRenderConfigPreset() bool`

HasRenderConfigPreset returns a boolean if a field has been set.

### GetAliases

`func (o *DatasourceSchema) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *DatasourceSchema) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *DatasourceSchema) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *DatasourceSchema) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetIsOnPrem

`func (o *DatasourceSchema) GetIsOnPrem() bool`

GetIsOnPrem returns the IsOnPrem field if non-nil, zero value otherwise.

### GetIsOnPremOk

`func (o *DatasourceSchema) GetIsOnPremOk() (*bool, bool)`

GetIsOnPremOk returns a tuple with the IsOnPrem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOnPrem

`func (o *DatasourceSchema) SetIsOnPrem(v bool)`

SetIsOnPrem sets IsOnPrem field to given value.

### HasIsOnPrem

`func (o *DatasourceSchema) HasIsOnPrem() bool`

HasIsOnPrem returns a boolean if a field has been set.

### GetTrustUrlRegexForViewActivity

`func (o *DatasourceSchema) GetTrustUrlRegexForViewActivity() bool`

GetTrustUrlRegexForViewActivity returns the TrustUrlRegexForViewActivity field if non-nil, zero value otherwise.

### GetTrustUrlRegexForViewActivityOk

`func (o *DatasourceSchema) GetTrustUrlRegexForViewActivityOk() (*bool, bool)`

GetTrustUrlRegexForViewActivityOk returns a tuple with the TrustUrlRegexForViewActivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustUrlRegexForViewActivity

`func (o *DatasourceSchema) SetTrustUrlRegexForViewActivity(v bool)`

SetTrustUrlRegexForViewActivity sets TrustUrlRegexForViewActivity field to given value.

### HasTrustUrlRegexForViewActivity

`func (o *DatasourceSchema) HasTrustUrlRegexForViewActivity() bool`

HasTrustUrlRegexForViewActivity returns a boolean if a field has been set.

### GetIncludeUtmSource

`func (o *DatasourceSchema) GetIncludeUtmSource() bool`

GetIncludeUtmSource returns the IncludeUtmSource field if non-nil, zero value otherwise.

### GetIncludeUtmSourceOk

`func (o *DatasourceSchema) GetIncludeUtmSourceOk() (*bool, bool)`

GetIncludeUtmSourceOk returns a tuple with the IncludeUtmSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeUtmSource

`func (o *DatasourceSchema) SetIncludeUtmSource(v bool)`

SetIncludeUtmSource sets IncludeUtmSource field to given value.

### HasIncludeUtmSource

`func (o *DatasourceSchema) HasIncludeUtmSource() bool`

HasIncludeUtmSource returns a boolean if a field has been set.

### GetDatasourceName

`func (o *DatasourceSchema) GetDatasourceName() string`

GetDatasourceName returns the DatasourceName field if non-nil, zero value otherwise.

### GetDatasourceNameOk

`func (o *DatasourceSchema) GetDatasourceNameOk() (*string, bool)`

GetDatasourceNameOk returns a tuple with the DatasourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasourceName

`func (o *DatasourceSchema) SetDatasourceName(v string)`

SetDatasourceName sets DatasourceName field to given value.

### HasDatasourceName

`func (o *DatasourceSchema) HasDatasourceName() bool`

HasDatasourceName returns a boolean if a field has been set.

### GetInstanceOnlyName

`func (o *DatasourceSchema) GetInstanceOnlyName() string`

GetInstanceOnlyName returns the InstanceOnlyName field if non-nil, zero value otherwise.

### GetInstanceOnlyNameOk

`func (o *DatasourceSchema) GetInstanceOnlyNameOk() (*string, bool)`

GetInstanceOnlyNameOk returns a tuple with the InstanceOnlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceOnlyName

`func (o *DatasourceSchema) SetInstanceOnlyName(v string)`

SetInstanceOnlyName sets InstanceOnlyName field to given value.

### HasInstanceOnlyName

`func (o *DatasourceSchema) HasInstanceOnlyName() bool`

HasInstanceOnlyName returns a boolean if a field has been set.

### GetInstanceDescription

`func (o *DatasourceSchema) GetInstanceDescription() string`

GetInstanceDescription returns the InstanceDescription field if non-nil, zero value otherwise.

### GetInstanceDescriptionOk

`func (o *DatasourceSchema) GetInstanceDescriptionOk() (*string, bool)`

GetInstanceDescriptionOk returns a tuple with the InstanceDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceDescription

`func (o *DatasourceSchema) SetInstanceDescription(v string)`

SetInstanceDescription sets InstanceDescription field to given value.

### HasInstanceDescription

`func (o *DatasourceSchema) HasInstanceDescription() bool`

HasInstanceDescription returns a boolean if a field has been set.

### GetCalendarEnabled

`func (o *DatasourceSchema) GetCalendarEnabled() bool`

GetCalendarEnabled returns the CalendarEnabled field if non-nil, zero value otherwise.

### GetCalendarEnabledOk

`func (o *DatasourceSchema) GetCalendarEnabledOk() (*bool, bool)`

GetCalendarEnabledOk returns a tuple with the CalendarEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalendarEnabled

`func (o *DatasourceSchema) SetCalendarEnabled(v bool)`

SetCalendarEnabled sets CalendarEnabled field to given value.

### HasCalendarEnabled

`func (o *DatasourceSchema) HasCalendarEnabled() bool`

HasCalendarEnabled returns a boolean if a field has been set.

### GetCollectBodyInActivity

`func (o *DatasourceSchema) GetCollectBodyInActivity() bool`

GetCollectBodyInActivity returns the CollectBodyInActivity field if non-nil, zero value otherwise.

### GetCollectBodyInActivityOk

`func (o *DatasourceSchema) GetCollectBodyInActivityOk() (*bool, bool)`

GetCollectBodyInActivityOk returns a tuple with the CollectBodyInActivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectBodyInActivity

`func (o *DatasourceSchema) SetCollectBodyInActivity(v bool)`

SetCollectBodyInActivity sets CollectBodyInActivity field to given value.

### HasCollectBodyInActivity

`func (o *DatasourceSchema) HasCollectBodyInActivity() bool`

HasCollectBodyInActivity returns a boolean if a field has been set.

### GetCrawlHomeUrl

`func (o *DatasourceSchema) GetCrawlHomeUrl() bool`

GetCrawlHomeUrl returns the CrawlHomeUrl field if non-nil, zero value otherwise.

### GetCrawlHomeUrlOk

`func (o *DatasourceSchema) GetCrawlHomeUrlOk() (*bool, bool)`

GetCrawlHomeUrlOk returns a tuple with the CrawlHomeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrawlHomeUrl

`func (o *DatasourceSchema) SetCrawlHomeUrl(v bool)`

SetCrawlHomeUrl sets CrawlHomeUrl field to given value.

### HasCrawlHomeUrl

`func (o *DatasourceSchema) HasCrawlHomeUrl() bool`

HasCrawlHomeUrl returns a boolean if a field has been set.

### GetDatasourceDescription

`func (o *DatasourceSchema) GetDatasourceDescription() string`

GetDatasourceDescription returns the DatasourceDescription field if non-nil, zero value otherwise.

### GetDatasourceDescriptionOk

`func (o *DatasourceSchema) GetDatasourceDescriptionOk() (*string, bool)`

GetDatasourceDescriptionOk returns a tuple with the DatasourceDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasourceDescription

`func (o *DatasourceSchema) SetDatasourceDescription(v string)`

SetDatasourceDescription sets DatasourceDescription field to given value.

### HasDatasourceDescription

`func (o *DatasourceSchema) HasDatasourceDescription() bool`

HasDatasourceDescription returns a boolean if a field has been set.

### GetDatasourceDisplayName

`func (o *DatasourceSchema) GetDatasourceDisplayName() string`

GetDatasourceDisplayName returns the DatasourceDisplayName field if non-nil, zero value otherwise.

### GetDatasourceDisplayNameOk

`func (o *DatasourceSchema) GetDatasourceDisplayNameOk() (*string, bool)`

GetDatasourceDisplayNameOk returns a tuple with the DatasourceDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasourceDisplayName

`func (o *DatasourceSchema) SetDatasourceDisplayName(v string)`

SetDatasourceDisplayName sets DatasourceDisplayName field to given value.

### HasDatasourceDisplayName

`func (o *DatasourceSchema) HasDatasourceDisplayName() bool`

HasDatasourceDisplayName returns a boolean if a field has been set.

### GetFederatedGoogleOAuth

`func (o *DatasourceSchema) GetFederatedGoogleOAuth() OAuthConfig`

GetFederatedGoogleOAuth returns the FederatedGoogleOAuth field if non-nil, zero value otherwise.

### GetFederatedGoogleOAuthOk

`func (o *DatasourceSchema) GetFederatedGoogleOAuthOk() (*OAuthConfig, bool)`

GetFederatedGoogleOAuthOk returns a tuple with the FederatedGoogleOAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederatedGoogleOAuth

`func (o *DatasourceSchema) SetFederatedGoogleOAuth(v OAuthConfig)`

SetFederatedGoogleOAuth sets FederatedGoogleOAuth field to given value.

### HasFederatedGoogleOAuth

`func (o *DatasourceSchema) HasFederatedGoogleOAuth() bool`

HasFederatedGoogleOAuth returns a boolean if a field has been set.

### GetIsEligibleForNativeReplacement

`func (o *DatasourceSchema) GetIsEligibleForNativeReplacement() bool`

GetIsEligibleForNativeReplacement returns the IsEligibleForNativeReplacement field if non-nil, zero value otherwise.

### GetIsEligibleForNativeReplacementOk

`func (o *DatasourceSchema) GetIsEligibleForNativeReplacementOk() (*bool, bool)`

GetIsEligibleForNativeReplacementOk returns a tuple with the IsEligibleForNativeReplacement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEligibleForNativeReplacement

`func (o *DatasourceSchema) SetIsEligibleForNativeReplacement(v bool)`

SetIsEligibleForNativeReplacement sets IsEligibleForNativeReplacement field to given value.

### HasIsEligibleForNativeReplacement

`func (o *DatasourceSchema) HasIsEligibleForNativeReplacement() bool`

HasIsEligibleForNativeReplacement returns a boolean if a field has been set.

### GetIsEnabled

`func (o *DatasourceSchema) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *DatasourceSchema) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *DatasourceSchema) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *DatasourceSchema) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetIsNativeReplacementAutoEnabled

`func (o *DatasourceSchema) GetIsNativeReplacementAutoEnabled() bool`

GetIsNativeReplacementAutoEnabled returns the IsNativeReplacementAutoEnabled field if non-nil, zero value otherwise.

### GetIsNativeReplacementAutoEnabledOk

`func (o *DatasourceSchema) GetIsNativeReplacementAutoEnabledOk() (*bool, bool)`

GetIsNativeReplacementAutoEnabledOk returns a tuple with the IsNativeReplacementAutoEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNativeReplacementAutoEnabled

`func (o *DatasourceSchema) SetIsNativeReplacementAutoEnabled(v bool)`

SetIsNativeReplacementAutoEnabled sets IsNativeReplacementAutoEnabled field to given value.

### HasIsNativeReplacementAutoEnabled

`func (o *DatasourceSchema) HasIsNativeReplacementAutoEnabled() bool`

HasIsNativeReplacementAutoEnabled returns a boolean if a field has been set.

### GetIsSearchable

`func (o *DatasourceSchema) GetIsSearchable() bool`

GetIsSearchable returns the IsSearchable field if non-nil, zero value otherwise.

### GetIsSearchableOk

`func (o *DatasourceSchema) GetIsSearchableOk() (*bool, bool)`

GetIsSearchableOk returns a tuple with the IsSearchable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSearchable

`func (o *DatasourceSchema) SetIsSearchable(v bool)`

SetIsSearchable sets IsSearchable field to given value.

### HasIsSearchable

`func (o *DatasourceSchema) HasIsSearchable() bool`

HasIsSearchable returns a boolean if a field has been set.

### GetMentionsEnabled

`func (o *DatasourceSchema) GetMentionsEnabled() bool`

GetMentionsEnabled returns the MentionsEnabled field if non-nil, zero value otherwise.

### GetMentionsEnabledOk

`func (o *DatasourceSchema) GetMentionsEnabledOk() (*bool, bool)`

GetMentionsEnabledOk returns a tuple with the MentionsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentionsEnabled

`func (o *DatasourceSchema) SetMentionsEnabled(v bool)`

SetMentionsEnabled sets MentionsEnabled field to given value.

### HasMentionsEnabled

`func (o *DatasourceSchema) HasMentionsEnabled() bool`

HasMentionsEnabled returns a boolean if a field has been set.

### GetPrivateAuth

`func (o *DatasourceSchema) GetPrivateAuth() PrivateAuth`

GetPrivateAuth returns the PrivateAuth field if non-nil, zero value otherwise.

### GetPrivateAuthOk

`func (o *DatasourceSchema) GetPrivateAuthOk() (*PrivateAuth, bool)`

GetPrivateAuthOk returns a tuple with the PrivateAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateAuth

`func (o *DatasourceSchema) SetPrivateAuth(v PrivateAuth)`

SetPrivateAuth sets PrivateAuth field to given value.

### HasPrivateAuth

`func (o *DatasourceSchema) HasPrivateAuth() bool`

HasPrivateAuth returns a boolean if a field has been set.

### GetSearchInputSelectors

`func (o *DatasourceSchema) GetSearchInputSelectors() []string`

GetSearchInputSelectors returns the SearchInputSelectors field if non-nil, zero value otherwise.

### GetSearchInputSelectorsOk

`func (o *DatasourceSchema) GetSearchInputSelectorsOk() (*[]string, bool)`

GetSearchInputSelectorsOk returns a tuple with the SearchInputSelectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchInputSelectors

`func (o *DatasourceSchema) SetSearchInputSelectors(v []string)`

SetSearchInputSelectors sets SearchInputSelectors field to given value.

### HasSearchInputSelectors

`func (o *DatasourceSchema) HasSearchInputSelectors() bool`

HasSearchInputSelectors returns a boolean if a field has been set.

### GetSupportsCalendar

`func (o *DatasourceSchema) GetSupportsCalendar() bool`

GetSupportsCalendar returns the SupportsCalendar field if non-nil, zero value otherwise.

### GetSupportsCalendarOk

`func (o *DatasourceSchema) GetSupportsCalendarOk() (*bool, bool)`

GetSupportsCalendarOk returns a tuple with the SupportsCalendar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsCalendar

`func (o *DatasourceSchema) SetSupportsCalendar(v bool)`

SetSupportsCalendar sets SupportsCalendar field to given value.

### HasSupportsCalendar

`func (o *DatasourceSchema) HasSupportsCalendar() bool`

HasSupportsCalendar returns a boolean if a field has been set.

### GetUrlRegexVersion

`func (o *DatasourceSchema) GetUrlRegexVersion() int32`

GetUrlRegexVersion returns the UrlRegexVersion field if non-nil, zero value otherwise.

### GetUrlRegexVersionOk

`func (o *DatasourceSchema) GetUrlRegexVersionOk() (*int32, bool)`

GetUrlRegexVersionOk returns a tuple with the UrlRegexVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlRegexVersion

`func (o *DatasourceSchema) SetUrlRegexVersion(v int32)`

SetUrlRegexVersion sets UrlRegexVersion field to given value.

### HasUrlRegexVersion

`func (o *DatasourceSchema) HasUrlRegexVersion() bool`

HasUrlRegexVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


