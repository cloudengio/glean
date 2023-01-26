# SharedDatasourceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Unique identifier of datasource instance to which this config applies. | 
**SuggestionText** | Pointer to **string** | Example text for what to search for in this datasource | [optional] 
**DisplayName** | Pointer to **string** | The user-friendly instance label to display. If omitted, falls back to the title-cased &#x60;name&#x60;. | [optional] 
**HomeUrl** | Pointer to **string** | The URL of the landing page for this datasource instance. Should point to the most useful page for users, not the company marketing page. | [optional] 
**CrawlerSeedUrls** | Pointer to **[]string** | This only applies to WEB_CRAWL and BROWSER_CRAWL datasources. Defines the seed urls for crawling. | [optional] 
**IconDarkUrl** | Pointer to **string** | The URL to an image to be displayed as an icon for this datasource instance in dark mode. Must have a transparency mask. SVG are recommended over PNG. Public, scio-authenticated and Base64 encoded data URLs are all valid (but not third-party-authenticated URLs). | [optional] 
**IconUrl** | Pointer to **string** | The URL to an image to be displayed as an icon for this datasource instance. Must have a transparency mask. SVG are recommended over PNG. Public, scio-authenticated and Base64 encoded data URLs are all valid (but not third-party-authenticated URLs). | [optional] 
**ObjectDefinitions** | Pointer to [**[]ObjectDefinition**](ObjectDefinition.md) | The list of top-level &#x60;objectType&#x60;s for the datasource. | [optional] 
**HideBuiltInFacets** | Pointer to **[]string** | List of built-in facet types that should be hidden for the datasource. | [optional] 
**UrlRegex** | Pointer to **string** | Regular expression that matches URLs of documents of the datasource instance. The behavior for multiple matches is non-deterministic. **Note: urlRegex is a required field for non-entity datasources (ie. datasources where isEntityDatasource is false). Please add a regex as specific as possible to this datasource instance.** | [optional] 
**CanonicalizingURLRegex** | Pointer to [**[]CanonicalizingRegexType**](CanonicalizingRegexType.md) | A list of regular expressions to apply to an arbitrary URL to transform it into a canonical URL for this datasource instance. Regexes are to be applied in the order specified in this list. | [optional] 
**CanonicalizingTitleRegex** | Pointer to [**[]CanonicalizingRegexType**](CanonicalizingRegexType.md) | A list of regular expressions to apply to an arbitrary title to transform it into a title that will be displayed in the search results | [optional] 
**RedlistTitleRegex** | Pointer to **string** | A regex that identifies titles that should not be indexed | [optional] 
**ConnectorType** | Pointer to [**ConnectorType**](ConnectorType.md) |  | [optional] 
**Quicklinks** | Pointer to [**[]Quicklink**](Quicklink.md) | List of actions for this datasource instance that will show up in autocomplete and app card, e.g. \&quot;Create new issue\&quot; for jira | [optional] 
**RenderConfigPreset** | Pointer to **string** | The name of a render config to use for displaying results from this datasource. Any well known datasource name may be used to render the same as that source, e.g. &#x60;web&#x60; or &#x60;gdrive&#x60;. | [optional] 
**Aliases** | Pointer to **[]string** | Aliases that can be used as &#x60;app&#x60; operator-values. | [optional] 
**DatasourceCategory** | Pointer to **string** | The type of this datasource. It is an important signal for relevance and must be specified and cannot be UNCATEGORIZED. | [optional] [default to "UNCATEGORIZED"]
**IsOnPrem** | Pointer to **bool** | Whether or not this datasource is hosted on-premise. | [optional] 
**TrustUrlRegexForViewActivity** | Pointer to **bool** | True if browser activity is able to report the correct URL for VIEW events. Set this to true if the URLs reported by Chrome are constant throughout each page load. Set this to false if the page has Javascript that modifies the URL during or after the load. | [optional] [default to true]
**IncludeUtmSource** | Pointer to **bool** | If true, a utm_source query param will be added to outbound links to this datasource within Glean. | [optional] 
**DatasourceName** | Pointer to **string** | The (non-unique) name of the datasource of which this config is an instance, e.g. \&quot;jira\&quot;. | [optional] 
**InstanceOnlyName** | Pointer to **string** | The instance of the datasource for this particular config, e.g. \&quot;onprem\&quot;. | [optional] 
**InstanceDescription** | Pointer to **string** | A human readable string identifying this instance as compared to its peers, e.g. \&quot;github.com/askscio\&quot; or \&quot;github.askscio.com\&quot; | [optional] 

## Methods

### NewSharedDatasourceConfig

`func NewSharedDatasourceConfig(name string, ) *SharedDatasourceConfig`

NewSharedDatasourceConfig instantiates a new SharedDatasourceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharedDatasourceConfigWithDefaults

`func NewSharedDatasourceConfigWithDefaults() *SharedDatasourceConfig`

NewSharedDatasourceConfigWithDefaults instantiates a new SharedDatasourceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SharedDatasourceConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SharedDatasourceConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SharedDatasourceConfig) SetName(v string)`

SetName sets Name field to given value.


### GetSuggestionText

`func (o *SharedDatasourceConfig) GetSuggestionText() string`

GetSuggestionText returns the SuggestionText field if non-nil, zero value otherwise.

### GetSuggestionTextOk

`func (o *SharedDatasourceConfig) GetSuggestionTextOk() (*string, bool)`

GetSuggestionTextOk returns a tuple with the SuggestionText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestionText

`func (o *SharedDatasourceConfig) SetSuggestionText(v string)`

SetSuggestionText sets SuggestionText field to given value.

### HasSuggestionText

`func (o *SharedDatasourceConfig) HasSuggestionText() bool`

HasSuggestionText returns a boolean if a field has been set.

### GetDisplayName

`func (o *SharedDatasourceConfig) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *SharedDatasourceConfig) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *SharedDatasourceConfig) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *SharedDatasourceConfig) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetHomeUrl

`func (o *SharedDatasourceConfig) GetHomeUrl() string`

GetHomeUrl returns the HomeUrl field if non-nil, zero value otherwise.

### GetHomeUrlOk

`func (o *SharedDatasourceConfig) GetHomeUrlOk() (*string, bool)`

GetHomeUrlOk returns a tuple with the HomeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeUrl

`func (o *SharedDatasourceConfig) SetHomeUrl(v string)`

SetHomeUrl sets HomeUrl field to given value.

### HasHomeUrl

`func (o *SharedDatasourceConfig) HasHomeUrl() bool`

HasHomeUrl returns a boolean if a field has been set.

### GetCrawlerSeedUrls

`func (o *SharedDatasourceConfig) GetCrawlerSeedUrls() []string`

GetCrawlerSeedUrls returns the CrawlerSeedUrls field if non-nil, zero value otherwise.

### GetCrawlerSeedUrlsOk

`func (o *SharedDatasourceConfig) GetCrawlerSeedUrlsOk() (*[]string, bool)`

GetCrawlerSeedUrlsOk returns a tuple with the CrawlerSeedUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrawlerSeedUrls

`func (o *SharedDatasourceConfig) SetCrawlerSeedUrls(v []string)`

SetCrawlerSeedUrls sets CrawlerSeedUrls field to given value.

### HasCrawlerSeedUrls

`func (o *SharedDatasourceConfig) HasCrawlerSeedUrls() bool`

HasCrawlerSeedUrls returns a boolean if a field has been set.

### GetIconDarkUrl

`func (o *SharedDatasourceConfig) GetIconDarkUrl() string`

GetIconDarkUrl returns the IconDarkUrl field if non-nil, zero value otherwise.

### GetIconDarkUrlOk

`func (o *SharedDatasourceConfig) GetIconDarkUrlOk() (*string, bool)`

GetIconDarkUrlOk returns a tuple with the IconDarkUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconDarkUrl

`func (o *SharedDatasourceConfig) SetIconDarkUrl(v string)`

SetIconDarkUrl sets IconDarkUrl field to given value.

### HasIconDarkUrl

`func (o *SharedDatasourceConfig) HasIconDarkUrl() bool`

HasIconDarkUrl returns a boolean if a field has been set.

### GetIconUrl

`func (o *SharedDatasourceConfig) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *SharedDatasourceConfig) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *SharedDatasourceConfig) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *SharedDatasourceConfig) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### GetObjectDefinitions

`func (o *SharedDatasourceConfig) GetObjectDefinitions() []ObjectDefinition`

GetObjectDefinitions returns the ObjectDefinitions field if non-nil, zero value otherwise.

### GetObjectDefinitionsOk

`func (o *SharedDatasourceConfig) GetObjectDefinitionsOk() (*[]ObjectDefinition, bool)`

GetObjectDefinitionsOk returns a tuple with the ObjectDefinitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectDefinitions

`func (o *SharedDatasourceConfig) SetObjectDefinitions(v []ObjectDefinition)`

SetObjectDefinitions sets ObjectDefinitions field to given value.

### HasObjectDefinitions

`func (o *SharedDatasourceConfig) HasObjectDefinitions() bool`

HasObjectDefinitions returns a boolean if a field has been set.

### GetHideBuiltInFacets

`func (o *SharedDatasourceConfig) GetHideBuiltInFacets() []string`

GetHideBuiltInFacets returns the HideBuiltInFacets field if non-nil, zero value otherwise.

### GetHideBuiltInFacetsOk

`func (o *SharedDatasourceConfig) GetHideBuiltInFacetsOk() (*[]string, bool)`

GetHideBuiltInFacetsOk returns a tuple with the HideBuiltInFacets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideBuiltInFacets

`func (o *SharedDatasourceConfig) SetHideBuiltInFacets(v []string)`

SetHideBuiltInFacets sets HideBuiltInFacets field to given value.

### HasHideBuiltInFacets

`func (o *SharedDatasourceConfig) HasHideBuiltInFacets() bool`

HasHideBuiltInFacets returns a boolean if a field has been set.

### GetUrlRegex

`func (o *SharedDatasourceConfig) GetUrlRegex() string`

GetUrlRegex returns the UrlRegex field if non-nil, zero value otherwise.

### GetUrlRegexOk

`func (o *SharedDatasourceConfig) GetUrlRegexOk() (*string, bool)`

GetUrlRegexOk returns a tuple with the UrlRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlRegex

`func (o *SharedDatasourceConfig) SetUrlRegex(v string)`

SetUrlRegex sets UrlRegex field to given value.

### HasUrlRegex

`func (o *SharedDatasourceConfig) HasUrlRegex() bool`

HasUrlRegex returns a boolean if a field has been set.

### GetCanonicalizingURLRegex

`func (o *SharedDatasourceConfig) GetCanonicalizingURLRegex() []CanonicalizingRegexType`

GetCanonicalizingURLRegex returns the CanonicalizingURLRegex field if non-nil, zero value otherwise.

### GetCanonicalizingURLRegexOk

`func (o *SharedDatasourceConfig) GetCanonicalizingURLRegexOk() (*[]CanonicalizingRegexType, bool)`

GetCanonicalizingURLRegexOk returns a tuple with the CanonicalizingURLRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalizingURLRegex

`func (o *SharedDatasourceConfig) SetCanonicalizingURLRegex(v []CanonicalizingRegexType)`

SetCanonicalizingURLRegex sets CanonicalizingURLRegex field to given value.

### HasCanonicalizingURLRegex

`func (o *SharedDatasourceConfig) HasCanonicalizingURLRegex() bool`

HasCanonicalizingURLRegex returns a boolean if a field has been set.

### GetCanonicalizingTitleRegex

`func (o *SharedDatasourceConfig) GetCanonicalizingTitleRegex() []CanonicalizingRegexType`

GetCanonicalizingTitleRegex returns the CanonicalizingTitleRegex field if non-nil, zero value otherwise.

### GetCanonicalizingTitleRegexOk

`func (o *SharedDatasourceConfig) GetCanonicalizingTitleRegexOk() (*[]CanonicalizingRegexType, bool)`

GetCanonicalizingTitleRegexOk returns a tuple with the CanonicalizingTitleRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalizingTitleRegex

`func (o *SharedDatasourceConfig) SetCanonicalizingTitleRegex(v []CanonicalizingRegexType)`

SetCanonicalizingTitleRegex sets CanonicalizingTitleRegex field to given value.

### HasCanonicalizingTitleRegex

`func (o *SharedDatasourceConfig) HasCanonicalizingTitleRegex() bool`

HasCanonicalizingTitleRegex returns a boolean if a field has been set.

### GetRedlistTitleRegex

`func (o *SharedDatasourceConfig) GetRedlistTitleRegex() string`

GetRedlistTitleRegex returns the RedlistTitleRegex field if non-nil, zero value otherwise.

### GetRedlistTitleRegexOk

`func (o *SharedDatasourceConfig) GetRedlistTitleRegexOk() (*string, bool)`

GetRedlistTitleRegexOk returns a tuple with the RedlistTitleRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedlistTitleRegex

`func (o *SharedDatasourceConfig) SetRedlistTitleRegex(v string)`

SetRedlistTitleRegex sets RedlistTitleRegex field to given value.

### HasRedlistTitleRegex

`func (o *SharedDatasourceConfig) HasRedlistTitleRegex() bool`

HasRedlistTitleRegex returns a boolean if a field has been set.

### GetConnectorType

`func (o *SharedDatasourceConfig) GetConnectorType() ConnectorType`

GetConnectorType returns the ConnectorType field if non-nil, zero value otherwise.

### GetConnectorTypeOk

`func (o *SharedDatasourceConfig) GetConnectorTypeOk() (*ConnectorType, bool)`

GetConnectorTypeOk returns a tuple with the ConnectorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorType

`func (o *SharedDatasourceConfig) SetConnectorType(v ConnectorType)`

SetConnectorType sets ConnectorType field to given value.

### HasConnectorType

`func (o *SharedDatasourceConfig) HasConnectorType() bool`

HasConnectorType returns a boolean if a field has been set.

### GetQuicklinks

`func (o *SharedDatasourceConfig) GetQuicklinks() []Quicklink`

GetQuicklinks returns the Quicklinks field if non-nil, zero value otherwise.

### GetQuicklinksOk

`func (o *SharedDatasourceConfig) GetQuicklinksOk() (*[]Quicklink, bool)`

GetQuicklinksOk returns a tuple with the Quicklinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuicklinks

`func (o *SharedDatasourceConfig) SetQuicklinks(v []Quicklink)`

SetQuicklinks sets Quicklinks field to given value.

### HasQuicklinks

`func (o *SharedDatasourceConfig) HasQuicklinks() bool`

HasQuicklinks returns a boolean if a field has been set.

### GetRenderConfigPreset

`func (o *SharedDatasourceConfig) GetRenderConfigPreset() string`

GetRenderConfigPreset returns the RenderConfigPreset field if non-nil, zero value otherwise.

### GetRenderConfigPresetOk

`func (o *SharedDatasourceConfig) GetRenderConfigPresetOk() (*string, bool)`

GetRenderConfigPresetOk returns a tuple with the RenderConfigPreset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenderConfigPreset

`func (o *SharedDatasourceConfig) SetRenderConfigPreset(v string)`

SetRenderConfigPreset sets RenderConfigPreset field to given value.

### HasRenderConfigPreset

`func (o *SharedDatasourceConfig) HasRenderConfigPreset() bool`

HasRenderConfigPreset returns a boolean if a field has been set.

### GetAliases

`func (o *SharedDatasourceConfig) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *SharedDatasourceConfig) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *SharedDatasourceConfig) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *SharedDatasourceConfig) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetDatasourceCategory

`func (o *SharedDatasourceConfig) GetDatasourceCategory() string`

GetDatasourceCategory returns the DatasourceCategory field if non-nil, zero value otherwise.

### GetDatasourceCategoryOk

`func (o *SharedDatasourceConfig) GetDatasourceCategoryOk() (*string, bool)`

GetDatasourceCategoryOk returns a tuple with the DatasourceCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasourceCategory

`func (o *SharedDatasourceConfig) SetDatasourceCategory(v string)`

SetDatasourceCategory sets DatasourceCategory field to given value.

### HasDatasourceCategory

`func (o *SharedDatasourceConfig) HasDatasourceCategory() bool`

HasDatasourceCategory returns a boolean if a field has been set.

### GetIsOnPrem

`func (o *SharedDatasourceConfig) GetIsOnPrem() bool`

GetIsOnPrem returns the IsOnPrem field if non-nil, zero value otherwise.

### GetIsOnPremOk

`func (o *SharedDatasourceConfig) GetIsOnPremOk() (*bool, bool)`

GetIsOnPremOk returns a tuple with the IsOnPrem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOnPrem

`func (o *SharedDatasourceConfig) SetIsOnPrem(v bool)`

SetIsOnPrem sets IsOnPrem field to given value.

### HasIsOnPrem

`func (o *SharedDatasourceConfig) HasIsOnPrem() bool`

HasIsOnPrem returns a boolean if a field has been set.

### GetTrustUrlRegexForViewActivity

`func (o *SharedDatasourceConfig) GetTrustUrlRegexForViewActivity() bool`

GetTrustUrlRegexForViewActivity returns the TrustUrlRegexForViewActivity field if non-nil, zero value otherwise.

### GetTrustUrlRegexForViewActivityOk

`func (o *SharedDatasourceConfig) GetTrustUrlRegexForViewActivityOk() (*bool, bool)`

GetTrustUrlRegexForViewActivityOk returns a tuple with the TrustUrlRegexForViewActivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustUrlRegexForViewActivity

`func (o *SharedDatasourceConfig) SetTrustUrlRegexForViewActivity(v bool)`

SetTrustUrlRegexForViewActivity sets TrustUrlRegexForViewActivity field to given value.

### HasTrustUrlRegexForViewActivity

`func (o *SharedDatasourceConfig) HasTrustUrlRegexForViewActivity() bool`

HasTrustUrlRegexForViewActivity returns a boolean if a field has been set.

### GetIncludeUtmSource

`func (o *SharedDatasourceConfig) GetIncludeUtmSource() bool`

GetIncludeUtmSource returns the IncludeUtmSource field if non-nil, zero value otherwise.

### GetIncludeUtmSourceOk

`func (o *SharedDatasourceConfig) GetIncludeUtmSourceOk() (*bool, bool)`

GetIncludeUtmSourceOk returns a tuple with the IncludeUtmSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeUtmSource

`func (o *SharedDatasourceConfig) SetIncludeUtmSource(v bool)`

SetIncludeUtmSource sets IncludeUtmSource field to given value.

### HasIncludeUtmSource

`func (o *SharedDatasourceConfig) HasIncludeUtmSource() bool`

HasIncludeUtmSource returns a boolean if a field has been set.

### GetDatasourceName

`func (o *SharedDatasourceConfig) GetDatasourceName() string`

GetDatasourceName returns the DatasourceName field if non-nil, zero value otherwise.

### GetDatasourceNameOk

`func (o *SharedDatasourceConfig) GetDatasourceNameOk() (*string, bool)`

GetDatasourceNameOk returns a tuple with the DatasourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasourceName

`func (o *SharedDatasourceConfig) SetDatasourceName(v string)`

SetDatasourceName sets DatasourceName field to given value.

### HasDatasourceName

`func (o *SharedDatasourceConfig) HasDatasourceName() bool`

HasDatasourceName returns a boolean if a field has been set.

### GetInstanceOnlyName

`func (o *SharedDatasourceConfig) GetInstanceOnlyName() string`

GetInstanceOnlyName returns the InstanceOnlyName field if non-nil, zero value otherwise.

### GetInstanceOnlyNameOk

`func (o *SharedDatasourceConfig) GetInstanceOnlyNameOk() (*string, bool)`

GetInstanceOnlyNameOk returns a tuple with the InstanceOnlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceOnlyName

`func (o *SharedDatasourceConfig) SetInstanceOnlyName(v string)`

SetInstanceOnlyName sets InstanceOnlyName field to given value.

### HasInstanceOnlyName

`func (o *SharedDatasourceConfig) HasInstanceOnlyName() bool`

HasInstanceOnlyName returns a boolean if a field has been set.

### GetInstanceDescription

`func (o *SharedDatasourceConfig) GetInstanceDescription() string`

GetInstanceDescription returns the InstanceDescription field if non-nil, zero value otherwise.

### GetInstanceDescriptionOk

`func (o *SharedDatasourceConfig) GetInstanceDescriptionOk() (*string, bool)`

GetInstanceDescriptionOk returns a tuple with the InstanceDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceDescription

`func (o *SharedDatasourceConfig) SetInstanceDescription(v string)`

SetInstanceDescription sets InstanceDescription field to given value.

### HasInstanceDescription

`func (o *SharedDatasourceConfig) HasInstanceDescription() bool`

HasInstanceDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


