# SharedDatasourceConfigNoInstance

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

## Methods

### NewSharedDatasourceConfigNoInstance

`func NewSharedDatasourceConfigNoInstance(name string, ) *SharedDatasourceConfigNoInstance`

NewSharedDatasourceConfigNoInstance instantiates a new SharedDatasourceConfigNoInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharedDatasourceConfigNoInstanceWithDefaults

`func NewSharedDatasourceConfigNoInstanceWithDefaults() *SharedDatasourceConfigNoInstance`

NewSharedDatasourceConfigNoInstanceWithDefaults instantiates a new SharedDatasourceConfigNoInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SharedDatasourceConfigNoInstance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SharedDatasourceConfigNoInstance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SharedDatasourceConfigNoInstance) SetName(v string)`

SetName sets Name field to given value.


### GetDisplayName

`func (o *SharedDatasourceConfigNoInstance) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *SharedDatasourceConfigNoInstance) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *SharedDatasourceConfigNoInstance) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *SharedDatasourceConfigNoInstance) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDatasourceCategory

`func (o *SharedDatasourceConfigNoInstance) GetDatasourceCategory() string`

GetDatasourceCategory returns the DatasourceCategory field if non-nil, zero value otherwise.

### GetDatasourceCategoryOk

`func (o *SharedDatasourceConfigNoInstance) GetDatasourceCategoryOk() (*string, bool)`

GetDatasourceCategoryOk returns a tuple with the DatasourceCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasourceCategory

`func (o *SharedDatasourceConfigNoInstance) SetDatasourceCategory(v string)`

SetDatasourceCategory sets DatasourceCategory field to given value.

### HasDatasourceCategory

`func (o *SharedDatasourceConfigNoInstance) HasDatasourceCategory() bool`

HasDatasourceCategory returns a boolean if a field has been set.

### GetUrlRegex

`func (o *SharedDatasourceConfigNoInstance) GetUrlRegex() string`

GetUrlRegex returns the UrlRegex field if non-nil, zero value otherwise.

### GetUrlRegexOk

`func (o *SharedDatasourceConfigNoInstance) GetUrlRegexOk() (*string, bool)`

GetUrlRegexOk returns a tuple with the UrlRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlRegex

`func (o *SharedDatasourceConfigNoInstance) SetUrlRegex(v string)`

SetUrlRegex sets UrlRegex field to given value.

### HasUrlRegex

`func (o *SharedDatasourceConfigNoInstance) HasUrlRegex() bool`

HasUrlRegex returns a boolean if a field has been set.

### GetIconUrl

`func (o *SharedDatasourceConfigNoInstance) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *SharedDatasourceConfigNoInstance) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *SharedDatasourceConfigNoInstance) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *SharedDatasourceConfigNoInstance) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### GetObjectDefinitions

`func (o *SharedDatasourceConfigNoInstance) GetObjectDefinitions() []ObjectDefinition`

GetObjectDefinitions returns the ObjectDefinitions field if non-nil, zero value otherwise.

### GetObjectDefinitionsOk

`func (o *SharedDatasourceConfigNoInstance) GetObjectDefinitionsOk() (*[]ObjectDefinition, bool)`

GetObjectDefinitionsOk returns a tuple with the ObjectDefinitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectDefinitions

`func (o *SharedDatasourceConfigNoInstance) SetObjectDefinitions(v []ObjectDefinition)`

SetObjectDefinitions sets ObjectDefinitions field to given value.

### HasObjectDefinitions

`func (o *SharedDatasourceConfigNoInstance) HasObjectDefinitions() bool`

HasObjectDefinitions returns a boolean if a field has been set.

### GetSuggestionText

`func (o *SharedDatasourceConfigNoInstance) GetSuggestionText() string`

GetSuggestionText returns the SuggestionText field if non-nil, zero value otherwise.

### GetSuggestionTextOk

`func (o *SharedDatasourceConfigNoInstance) GetSuggestionTextOk() (*string, bool)`

GetSuggestionTextOk returns a tuple with the SuggestionText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestionText

`func (o *SharedDatasourceConfigNoInstance) SetSuggestionText(v string)`

SetSuggestionText sets SuggestionText field to given value.

### HasSuggestionText

`func (o *SharedDatasourceConfigNoInstance) HasSuggestionText() bool`

HasSuggestionText returns a boolean if a field has been set.

### GetHomeUrl

`func (o *SharedDatasourceConfigNoInstance) GetHomeUrl() string`

GetHomeUrl returns the HomeUrl field if non-nil, zero value otherwise.

### GetHomeUrlOk

`func (o *SharedDatasourceConfigNoInstance) GetHomeUrlOk() (*string, bool)`

GetHomeUrlOk returns a tuple with the HomeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeUrl

`func (o *SharedDatasourceConfigNoInstance) SetHomeUrl(v string)`

SetHomeUrl sets HomeUrl field to given value.

### HasHomeUrl

`func (o *SharedDatasourceConfigNoInstance) HasHomeUrl() bool`

HasHomeUrl returns a boolean if a field has been set.

### GetCrawlerSeedUrls

`func (o *SharedDatasourceConfigNoInstance) GetCrawlerSeedUrls() []string`

GetCrawlerSeedUrls returns the CrawlerSeedUrls field if non-nil, zero value otherwise.

### GetCrawlerSeedUrlsOk

`func (o *SharedDatasourceConfigNoInstance) GetCrawlerSeedUrlsOk() (*[]string, bool)`

GetCrawlerSeedUrlsOk returns a tuple with the CrawlerSeedUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrawlerSeedUrls

`func (o *SharedDatasourceConfigNoInstance) SetCrawlerSeedUrls(v []string)`

SetCrawlerSeedUrls sets CrawlerSeedUrls field to given value.

### HasCrawlerSeedUrls

`func (o *SharedDatasourceConfigNoInstance) HasCrawlerSeedUrls() bool`

HasCrawlerSeedUrls returns a boolean if a field has been set.

### GetIconDarkUrl

`func (o *SharedDatasourceConfigNoInstance) GetIconDarkUrl() string`

GetIconDarkUrl returns the IconDarkUrl field if non-nil, zero value otherwise.

### GetIconDarkUrlOk

`func (o *SharedDatasourceConfigNoInstance) GetIconDarkUrlOk() (*string, bool)`

GetIconDarkUrlOk returns a tuple with the IconDarkUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconDarkUrl

`func (o *SharedDatasourceConfigNoInstance) SetIconDarkUrl(v string)`

SetIconDarkUrl sets IconDarkUrl field to given value.

### HasIconDarkUrl

`func (o *SharedDatasourceConfigNoInstance) HasIconDarkUrl() bool`

HasIconDarkUrl returns a boolean if a field has been set.

### GetHideBuiltInFacets

`func (o *SharedDatasourceConfigNoInstance) GetHideBuiltInFacets() []string`

GetHideBuiltInFacets returns the HideBuiltInFacets field if non-nil, zero value otherwise.

### GetHideBuiltInFacetsOk

`func (o *SharedDatasourceConfigNoInstance) GetHideBuiltInFacetsOk() (*[]string, bool)`

GetHideBuiltInFacetsOk returns a tuple with the HideBuiltInFacets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideBuiltInFacets

`func (o *SharedDatasourceConfigNoInstance) SetHideBuiltInFacets(v []string)`

SetHideBuiltInFacets sets HideBuiltInFacets field to given value.

### HasHideBuiltInFacets

`func (o *SharedDatasourceConfigNoInstance) HasHideBuiltInFacets() bool`

HasHideBuiltInFacets returns a boolean if a field has been set.

### GetCanonicalizingURLRegex

`func (o *SharedDatasourceConfigNoInstance) GetCanonicalizingURLRegex() []CanonicalizingRegexType`

GetCanonicalizingURLRegex returns the CanonicalizingURLRegex field if non-nil, zero value otherwise.

### GetCanonicalizingURLRegexOk

`func (o *SharedDatasourceConfigNoInstance) GetCanonicalizingURLRegexOk() (*[]CanonicalizingRegexType, bool)`

GetCanonicalizingURLRegexOk returns a tuple with the CanonicalizingURLRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalizingURLRegex

`func (o *SharedDatasourceConfigNoInstance) SetCanonicalizingURLRegex(v []CanonicalizingRegexType)`

SetCanonicalizingURLRegex sets CanonicalizingURLRegex field to given value.

### HasCanonicalizingURLRegex

`func (o *SharedDatasourceConfigNoInstance) HasCanonicalizingURLRegex() bool`

HasCanonicalizingURLRegex returns a boolean if a field has been set.

### GetCanonicalizingTitleRegex

`func (o *SharedDatasourceConfigNoInstance) GetCanonicalizingTitleRegex() []CanonicalizingRegexType`

GetCanonicalizingTitleRegex returns the CanonicalizingTitleRegex field if non-nil, zero value otherwise.

### GetCanonicalizingTitleRegexOk

`func (o *SharedDatasourceConfigNoInstance) GetCanonicalizingTitleRegexOk() (*[]CanonicalizingRegexType, bool)`

GetCanonicalizingTitleRegexOk returns a tuple with the CanonicalizingTitleRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalizingTitleRegex

`func (o *SharedDatasourceConfigNoInstance) SetCanonicalizingTitleRegex(v []CanonicalizingRegexType)`

SetCanonicalizingTitleRegex sets CanonicalizingTitleRegex field to given value.

### HasCanonicalizingTitleRegex

`func (o *SharedDatasourceConfigNoInstance) HasCanonicalizingTitleRegex() bool`

HasCanonicalizingTitleRegex returns a boolean if a field has been set.

### GetRedlistTitleRegex

`func (o *SharedDatasourceConfigNoInstance) GetRedlistTitleRegex() string`

GetRedlistTitleRegex returns the RedlistTitleRegex field if non-nil, zero value otherwise.

### GetRedlistTitleRegexOk

`func (o *SharedDatasourceConfigNoInstance) GetRedlistTitleRegexOk() (*string, bool)`

GetRedlistTitleRegexOk returns a tuple with the RedlistTitleRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedlistTitleRegex

`func (o *SharedDatasourceConfigNoInstance) SetRedlistTitleRegex(v string)`

SetRedlistTitleRegex sets RedlistTitleRegex field to given value.

### HasRedlistTitleRegex

`func (o *SharedDatasourceConfigNoInstance) HasRedlistTitleRegex() bool`

HasRedlistTitleRegex returns a boolean if a field has been set.

### GetConnectorType

`func (o *SharedDatasourceConfigNoInstance) GetConnectorType() ConnectorType`

GetConnectorType returns the ConnectorType field if non-nil, zero value otherwise.

### GetConnectorTypeOk

`func (o *SharedDatasourceConfigNoInstance) GetConnectorTypeOk() (*ConnectorType, bool)`

GetConnectorTypeOk returns a tuple with the ConnectorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorType

`func (o *SharedDatasourceConfigNoInstance) SetConnectorType(v ConnectorType)`

SetConnectorType sets ConnectorType field to given value.

### HasConnectorType

`func (o *SharedDatasourceConfigNoInstance) HasConnectorType() bool`

HasConnectorType returns a boolean if a field has been set.

### GetQuicklinks

`func (o *SharedDatasourceConfigNoInstance) GetQuicklinks() []Quicklink`

GetQuicklinks returns the Quicklinks field if non-nil, zero value otherwise.

### GetQuicklinksOk

`func (o *SharedDatasourceConfigNoInstance) GetQuicklinksOk() (*[]Quicklink, bool)`

GetQuicklinksOk returns a tuple with the Quicklinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuicklinks

`func (o *SharedDatasourceConfigNoInstance) SetQuicklinks(v []Quicklink)`

SetQuicklinks sets Quicklinks field to given value.

### HasQuicklinks

`func (o *SharedDatasourceConfigNoInstance) HasQuicklinks() bool`

HasQuicklinks returns a boolean if a field has been set.

### GetRenderConfigPreset

`func (o *SharedDatasourceConfigNoInstance) GetRenderConfigPreset() string`

GetRenderConfigPreset returns the RenderConfigPreset field if non-nil, zero value otherwise.

### GetRenderConfigPresetOk

`func (o *SharedDatasourceConfigNoInstance) GetRenderConfigPresetOk() (*string, bool)`

GetRenderConfigPresetOk returns a tuple with the RenderConfigPreset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenderConfigPreset

`func (o *SharedDatasourceConfigNoInstance) SetRenderConfigPreset(v string)`

SetRenderConfigPreset sets RenderConfigPreset field to given value.

### HasRenderConfigPreset

`func (o *SharedDatasourceConfigNoInstance) HasRenderConfigPreset() bool`

HasRenderConfigPreset returns a boolean if a field has been set.

### GetAliases

`func (o *SharedDatasourceConfigNoInstance) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *SharedDatasourceConfigNoInstance) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *SharedDatasourceConfigNoInstance) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *SharedDatasourceConfigNoInstance) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetIsOnPrem

`func (o *SharedDatasourceConfigNoInstance) GetIsOnPrem() bool`

GetIsOnPrem returns the IsOnPrem field if non-nil, zero value otherwise.

### GetIsOnPremOk

`func (o *SharedDatasourceConfigNoInstance) GetIsOnPremOk() (*bool, bool)`

GetIsOnPremOk returns a tuple with the IsOnPrem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOnPrem

`func (o *SharedDatasourceConfigNoInstance) SetIsOnPrem(v bool)`

SetIsOnPrem sets IsOnPrem field to given value.

### HasIsOnPrem

`func (o *SharedDatasourceConfigNoInstance) HasIsOnPrem() bool`

HasIsOnPrem returns a boolean if a field has been set.

### GetTrustUrlRegexForViewActivity

`func (o *SharedDatasourceConfigNoInstance) GetTrustUrlRegexForViewActivity() bool`

GetTrustUrlRegexForViewActivity returns the TrustUrlRegexForViewActivity field if non-nil, zero value otherwise.

### GetTrustUrlRegexForViewActivityOk

`func (o *SharedDatasourceConfigNoInstance) GetTrustUrlRegexForViewActivityOk() (*bool, bool)`

GetTrustUrlRegexForViewActivityOk returns a tuple with the TrustUrlRegexForViewActivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustUrlRegexForViewActivity

`func (o *SharedDatasourceConfigNoInstance) SetTrustUrlRegexForViewActivity(v bool)`

SetTrustUrlRegexForViewActivity sets TrustUrlRegexForViewActivity field to given value.

### HasTrustUrlRegexForViewActivity

`func (o *SharedDatasourceConfigNoInstance) HasTrustUrlRegexForViewActivity() bool`

HasTrustUrlRegexForViewActivity returns a boolean if a field has been set.

### GetIncludeUtmSource

`func (o *SharedDatasourceConfigNoInstance) GetIncludeUtmSource() bool`

GetIncludeUtmSource returns the IncludeUtmSource field if non-nil, zero value otherwise.

### GetIncludeUtmSourceOk

`func (o *SharedDatasourceConfigNoInstance) GetIncludeUtmSourceOk() (*bool, bool)`

GetIncludeUtmSourceOk returns a tuple with the IncludeUtmSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeUtmSource

`func (o *SharedDatasourceConfigNoInstance) SetIncludeUtmSource(v bool)`

SetIncludeUtmSource sets IncludeUtmSource field to given value.

### HasIncludeUtmSource

`func (o *SharedDatasourceConfigNoInstance) HasIncludeUtmSource() bool`

HasIncludeUtmSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


