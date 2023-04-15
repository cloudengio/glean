# CustomDatasourceConfig

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
**IdentityDatasourceName** | Pointer to **string** | If the datasource uses another datasource for identity info, then the name of the datasource. The identity datasource must exist already. | [optional] 
**ProductAccessGroup** | Pointer to **string** | If the datasource uses a specific product access group, then the name of that group. | [optional] 
**IsUserReferencedByEmail** | Pointer to **bool** | whether email is used to reference users in document ACLs and in group memberships. | [optional] 
**IsEntityDatasource** | Pointer to **bool** | True if this datasource is used to push custom entities. | [optional] [default to false]
**IsTestDatasource** | Pointer to **bool** | True if this datasource will be used for testing purpose only. Documents from such a datasource wouldn&#39;t have any effect on search rankings. | [optional] [default to false]

## Methods

### NewCustomDatasourceConfig

`func NewCustomDatasourceConfig(name string, ) *CustomDatasourceConfig`

NewCustomDatasourceConfig instantiates a new CustomDatasourceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomDatasourceConfigWithDefaults

`func NewCustomDatasourceConfigWithDefaults() *CustomDatasourceConfig`

NewCustomDatasourceConfigWithDefaults instantiates a new CustomDatasourceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CustomDatasourceConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomDatasourceConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomDatasourceConfig) SetName(v string)`

SetName sets Name field to given value.


### GetDisplayName

`func (o *CustomDatasourceConfig) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CustomDatasourceConfig) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CustomDatasourceConfig) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CustomDatasourceConfig) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDatasourceCategory

`func (o *CustomDatasourceConfig) GetDatasourceCategory() string`

GetDatasourceCategory returns the DatasourceCategory field if non-nil, zero value otherwise.

### GetDatasourceCategoryOk

`func (o *CustomDatasourceConfig) GetDatasourceCategoryOk() (*string, bool)`

GetDatasourceCategoryOk returns a tuple with the DatasourceCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasourceCategory

`func (o *CustomDatasourceConfig) SetDatasourceCategory(v string)`

SetDatasourceCategory sets DatasourceCategory field to given value.

### HasDatasourceCategory

`func (o *CustomDatasourceConfig) HasDatasourceCategory() bool`

HasDatasourceCategory returns a boolean if a field has been set.

### GetUrlRegex

`func (o *CustomDatasourceConfig) GetUrlRegex() string`

GetUrlRegex returns the UrlRegex field if non-nil, zero value otherwise.

### GetUrlRegexOk

`func (o *CustomDatasourceConfig) GetUrlRegexOk() (*string, bool)`

GetUrlRegexOk returns a tuple with the UrlRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlRegex

`func (o *CustomDatasourceConfig) SetUrlRegex(v string)`

SetUrlRegex sets UrlRegex field to given value.

### HasUrlRegex

`func (o *CustomDatasourceConfig) HasUrlRegex() bool`

HasUrlRegex returns a boolean if a field has been set.

### GetIconUrl

`func (o *CustomDatasourceConfig) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *CustomDatasourceConfig) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *CustomDatasourceConfig) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *CustomDatasourceConfig) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### GetObjectDefinitions

`func (o *CustomDatasourceConfig) GetObjectDefinitions() []ObjectDefinition`

GetObjectDefinitions returns the ObjectDefinitions field if non-nil, zero value otherwise.

### GetObjectDefinitionsOk

`func (o *CustomDatasourceConfig) GetObjectDefinitionsOk() (*[]ObjectDefinition, bool)`

GetObjectDefinitionsOk returns a tuple with the ObjectDefinitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectDefinitions

`func (o *CustomDatasourceConfig) SetObjectDefinitions(v []ObjectDefinition)`

SetObjectDefinitions sets ObjectDefinitions field to given value.

### HasObjectDefinitions

`func (o *CustomDatasourceConfig) HasObjectDefinitions() bool`

HasObjectDefinitions returns a boolean if a field has been set.

### GetSuggestionText

`func (o *CustomDatasourceConfig) GetSuggestionText() string`

GetSuggestionText returns the SuggestionText field if non-nil, zero value otherwise.

### GetSuggestionTextOk

`func (o *CustomDatasourceConfig) GetSuggestionTextOk() (*string, bool)`

GetSuggestionTextOk returns a tuple with the SuggestionText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestionText

`func (o *CustomDatasourceConfig) SetSuggestionText(v string)`

SetSuggestionText sets SuggestionText field to given value.

### HasSuggestionText

`func (o *CustomDatasourceConfig) HasSuggestionText() bool`

HasSuggestionText returns a boolean if a field has been set.

### GetHomeUrl

`func (o *CustomDatasourceConfig) GetHomeUrl() string`

GetHomeUrl returns the HomeUrl field if non-nil, zero value otherwise.

### GetHomeUrlOk

`func (o *CustomDatasourceConfig) GetHomeUrlOk() (*string, bool)`

GetHomeUrlOk returns a tuple with the HomeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeUrl

`func (o *CustomDatasourceConfig) SetHomeUrl(v string)`

SetHomeUrl sets HomeUrl field to given value.

### HasHomeUrl

`func (o *CustomDatasourceConfig) HasHomeUrl() bool`

HasHomeUrl returns a boolean if a field has been set.

### GetCrawlerSeedUrls

`func (o *CustomDatasourceConfig) GetCrawlerSeedUrls() []string`

GetCrawlerSeedUrls returns the CrawlerSeedUrls field if non-nil, zero value otherwise.

### GetCrawlerSeedUrlsOk

`func (o *CustomDatasourceConfig) GetCrawlerSeedUrlsOk() (*[]string, bool)`

GetCrawlerSeedUrlsOk returns a tuple with the CrawlerSeedUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrawlerSeedUrls

`func (o *CustomDatasourceConfig) SetCrawlerSeedUrls(v []string)`

SetCrawlerSeedUrls sets CrawlerSeedUrls field to given value.

### HasCrawlerSeedUrls

`func (o *CustomDatasourceConfig) HasCrawlerSeedUrls() bool`

HasCrawlerSeedUrls returns a boolean if a field has been set.

### GetIconDarkUrl

`func (o *CustomDatasourceConfig) GetIconDarkUrl() string`

GetIconDarkUrl returns the IconDarkUrl field if non-nil, zero value otherwise.

### GetIconDarkUrlOk

`func (o *CustomDatasourceConfig) GetIconDarkUrlOk() (*string, bool)`

GetIconDarkUrlOk returns a tuple with the IconDarkUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconDarkUrl

`func (o *CustomDatasourceConfig) SetIconDarkUrl(v string)`

SetIconDarkUrl sets IconDarkUrl field to given value.

### HasIconDarkUrl

`func (o *CustomDatasourceConfig) HasIconDarkUrl() bool`

HasIconDarkUrl returns a boolean if a field has been set.

### GetHideBuiltInFacets

`func (o *CustomDatasourceConfig) GetHideBuiltInFacets() []string`

GetHideBuiltInFacets returns the HideBuiltInFacets field if non-nil, zero value otherwise.

### GetHideBuiltInFacetsOk

`func (o *CustomDatasourceConfig) GetHideBuiltInFacetsOk() (*[]string, bool)`

GetHideBuiltInFacetsOk returns a tuple with the HideBuiltInFacets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideBuiltInFacets

`func (o *CustomDatasourceConfig) SetHideBuiltInFacets(v []string)`

SetHideBuiltInFacets sets HideBuiltInFacets field to given value.

### HasHideBuiltInFacets

`func (o *CustomDatasourceConfig) HasHideBuiltInFacets() bool`

HasHideBuiltInFacets returns a boolean if a field has been set.

### GetCanonicalizingURLRegex

`func (o *CustomDatasourceConfig) GetCanonicalizingURLRegex() []CanonicalizingRegexType`

GetCanonicalizingURLRegex returns the CanonicalizingURLRegex field if non-nil, zero value otherwise.

### GetCanonicalizingURLRegexOk

`func (o *CustomDatasourceConfig) GetCanonicalizingURLRegexOk() (*[]CanonicalizingRegexType, bool)`

GetCanonicalizingURLRegexOk returns a tuple with the CanonicalizingURLRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalizingURLRegex

`func (o *CustomDatasourceConfig) SetCanonicalizingURLRegex(v []CanonicalizingRegexType)`

SetCanonicalizingURLRegex sets CanonicalizingURLRegex field to given value.

### HasCanonicalizingURLRegex

`func (o *CustomDatasourceConfig) HasCanonicalizingURLRegex() bool`

HasCanonicalizingURLRegex returns a boolean if a field has been set.

### GetCanonicalizingTitleRegex

`func (o *CustomDatasourceConfig) GetCanonicalizingTitleRegex() []CanonicalizingRegexType`

GetCanonicalizingTitleRegex returns the CanonicalizingTitleRegex field if non-nil, zero value otherwise.

### GetCanonicalizingTitleRegexOk

`func (o *CustomDatasourceConfig) GetCanonicalizingTitleRegexOk() (*[]CanonicalizingRegexType, bool)`

GetCanonicalizingTitleRegexOk returns a tuple with the CanonicalizingTitleRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalizingTitleRegex

`func (o *CustomDatasourceConfig) SetCanonicalizingTitleRegex(v []CanonicalizingRegexType)`

SetCanonicalizingTitleRegex sets CanonicalizingTitleRegex field to given value.

### HasCanonicalizingTitleRegex

`func (o *CustomDatasourceConfig) HasCanonicalizingTitleRegex() bool`

HasCanonicalizingTitleRegex returns a boolean if a field has been set.

### GetRedlistTitleRegex

`func (o *CustomDatasourceConfig) GetRedlistTitleRegex() string`

GetRedlistTitleRegex returns the RedlistTitleRegex field if non-nil, zero value otherwise.

### GetRedlistTitleRegexOk

`func (o *CustomDatasourceConfig) GetRedlistTitleRegexOk() (*string, bool)`

GetRedlistTitleRegexOk returns a tuple with the RedlistTitleRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedlistTitleRegex

`func (o *CustomDatasourceConfig) SetRedlistTitleRegex(v string)`

SetRedlistTitleRegex sets RedlistTitleRegex field to given value.

### HasRedlistTitleRegex

`func (o *CustomDatasourceConfig) HasRedlistTitleRegex() bool`

HasRedlistTitleRegex returns a boolean if a field has been set.

### GetConnectorType

`func (o *CustomDatasourceConfig) GetConnectorType() ConnectorType`

GetConnectorType returns the ConnectorType field if non-nil, zero value otherwise.

### GetConnectorTypeOk

`func (o *CustomDatasourceConfig) GetConnectorTypeOk() (*ConnectorType, bool)`

GetConnectorTypeOk returns a tuple with the ConnectorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorType

`func (o *CustomDatasourceConfig) SetConnectorType(v ConnectorType)`

SetConnectorType sets ConnectorType field to given value.

### HasConnectorType

`func (o *CustomDatasourceConfig) HasConnectorType() bool`

HasConnectorType returns a boolean if a field has been set.

### GetQuicklinks

`func (o *CustomDatasourceConfig) GetQuicklinks() []Quicklink`

GetQuicklinks returns the Quicklinks field if non-nil, zero value otherwise.

### GetQuicklinksOk

`func (o *CustomDatasourceConfig) GetQuicklinksOk() (*[]Quicklink, bool)`

GetQuicklinksOk returns a tuple with the Quicklinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuicklinks

`func (o *CustomDatasourceConfig) SetQuicklinks(v []Quicklink)`

SetQuicklinks sets Quicklinks field to given value.

### HasQuicklinks

`func (o *CustomDatasourceConfig) HasQuicklinks() bool`

HasQuicklinks returns a boolean if a field has been set.

### GetRenderConfigPreset

`func (o *CustomDatasourceConfig) GetRenderConfigPreset() string`

GetRenderConfigPreset returns the RenderConfigPreset field if non-nil, zero value otherwise.

### GetRenderConfigPresetOk

`func (o *CustomDatasourceConfig) GetRenderConfigPresetOk() (*string, bool)`

GetRenderConfigPresetOk returns a tuple with the RenderConfigPreset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenderConfigPreset

`func (o *CustomDatasourceConfig) SetRenderConfigPreset(v string)`

SetRenderConfigPreset sets RenderConfigPreset field to given value.

### HasRenderConfigPreset

`func (o *CustomDatasourceConfig) HasRenderConfigPreset() bool`

HasRenderConfigPreset returns a boolean if a field has been set.

### GetAliases

`func (o *CustomDatasourceConfig) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *CustomDatasourceConfig) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *CustomDatasourceConfig) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *CustomDatasourceConfig) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetIsOnPrem

`func (o *CustomDatasourceConfig) GetIsOnPrem() bool`

GetIsOnPrem returns the IsOnPrem field if non-nil, zero value otherwise.

### GetIsOnPremOk

`func (o *CustomDatasourceConfig) GetIsOnPremOk() (*bool, bool)`

GetIsOnPremOk returns a tuple with the IsOnPrem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOnPrem

`func (o *CustomDatasourceConfig) SetIsOnPrem(v bool)`

SetIsOnPrem sets IsOnPrem field to given value.

### HasIsOnPrem

`func (o *CustomDatasourceConfig) HasIsOnPrem() bool`

HasIsOnPrem returns a boolean if a field has been set.

### GetTrustUrlRegexForViewActivity

`func (o *CustomDatasourceConfig) GetTrustUrlRegexForViewActivity() bool`

GetTrustUrlRegexForViewActivity returns the TrustUrlRegexForViewActivity field if non-nil, zero value otherwise.

### GetTrustUrlRegexForViewActivityOk

`func (o *CustomDatasourceConfig) GetTrustUrlRegexForViewActivityOk() (*bool, bool)`

GetTrustUrlRegexForViewActivityOk returns a tuple with the TrustUrlRegexForViewActivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustUrlRegexForViewActivity

`func (o *CustomDatasourceConfig) SetTrustUrlRegexForViewActivity(v bool)`

SetTrustUrlRegexForViewActivity sets TrustUrlRegexForViewActivity field to given value.

### HasTrustUrlRegexForViewActivity

`func (o *CustomDatasourceConfig) HasTrustUrlRegexForViewActivity() bool`

HasTrustUrlRegexForViewActivity returns a boolean if a field has been set.

### GetIncludeUtmSource

`func (o *CustomDatasourceConfig) GetIncludeUtmSource() bool`

GetIncludeUtmSource returns the IncludeUtmSource field if non-nil, zero value otherwise.

### GetIncludeUtmSourceOk

`func (o *CustomDatasourceConfig) GetIncludeUtmSourceOk() (*bool, bool)`

GetIncludeUtmSourceOk returns a tuple with the IncludeUtmSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeUtmSource

`func (o *CustomDatasourceConfig) SetIncludeUtmSource(v bool)`

SetIncludeUtmSource sets IncludeUtmSource field to given value.

### HasIncludeUtmSource

`func (o *CustomDatasourceConfig) HasIncludeUtmSource() bool`

HasIncludeUtmSource returns a boolean if a field has been set.

### GetIdentityDatasourceName

`func (o *CustomDatasourceConfig) GetIdentityDatasourceName() string`

GetIdentityDatasourceName returns the IdentityDatasourceName field if non-nil, zero value otherwise.

### GetIdentityDatasourceNameOk

`func (o *CustomDatasourceConfig) GetIdentityDatasourceNameOk() (*string, bool)`

GetIdentityDatasourceNameOk returns a tuple with the IdentityDatasourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityDatasourceName

`func (o *CustomDatasourceConfig) SetIdentityDatasourceName(v string)`

SetIdentityDatasourceName sets IdentityDatasourceName field to given value.

### HasIdentityDatasourceName

`func (o *CustomDatasourceConfig) HasIdentityDatasourceName() bool`

HasIdentityDatasourceName returns a boolean if a field has been set.

### GetProductAccessGroup

`func (o *CustomDatasourceConfig) GetProductAccessGroup() string`

GetProductAccessGroup returns the ProductAccessGroup field if non-nil, zero value otherwise.

### GetProductAccessGroupOk

`func (o *CustomDatasourceConfig) GetProductAccessGroupOk() (*string, bool)`

GetProductAccessGroupOk returns a tuple with the ProductAccessGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductAccessGroup

`func (o *CustomDatasourceConfig) SetProductAccessGroup(v string)`

SetProductAccessGroup sets ProductAccessGroup field to given value.

### HasProductAccessGroup

`func (o *CustomDatasourceConfig) HasProductAccessGroup() bool`

HasProductAccessGroup returns a boolean if a field has been set.

### GetIsUserReferencedByEmail

`func (o *CustomDatasourceConfig) GetIsUserReferencedByEmail() bool`

GetIsUserReferencedByEmail returns the IsUserReferencedByEmail field if non-nil, zero value otherwise.

### GetIsUserReferencedByEmailOk

`func (o *CustomDatasourceConfig) GetIsUserReferencedByEmailOk() (*bool, bool)`

GetIsUserReferencedByEmailOk returns a tuple with the IsUserReferencedByEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUserReferencedByEmail

`func (o *CustomDatasourceConfig) SetIsUserReferencedByEmail(v bool)`

SetIsUserReferencedByEmail sets IsUserReferencedByEmail field to given value.

### HasIsUserReferencedByEmail

`func (o *CustomDatasourceConfig) HasIsUserReferencedByEmail() bool`

HasIsUserReferencedByEmail returns a boolean if a field has been set.

### GetIsEntityDatasource

`func (o *CustomDatasourceConfig) GetIsEntityDatasource() bool`

GetIsEntityDatasource returns the IsEntityDatasource field if non-nil, zero value otherwise.

### GetIsEntityDatasourceOk

`func (o *CustomDatasourceConfig) GetIsEntityDatasourceOk() (*bool, bool)`

GetIsEntityDatasourceOk returns a tuple with the IsEntityDatasource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEntityDatasource

`func (o *CustomDatasourceConfig) SetIsEntityDatasource(v bool)`

SetIsEntityDatasource sets IsEntityDatasource field to given value.

### HasIsEntityDatasource

`func (o *CustomDatasourceConfig) HasIsEntityDatasource() bool`

HasIsEntityDatasource returns a boolean if a field has been set.

### GetIsTestDatasource

`func (o *CustomDatasourceConfig) GetIsTestDatasource() bool`

GetIsTestDatasource returns the IsTestDatasource field if non-nil, zero value otherwise.

### GetIsTestDatasourceOk

`func (o *CustomDatasourceConfig) GetIsTestDatasourceOk() (*bool, bool)`

GetIsTestDatasourceOk returns a tuple with the IsTestDatasource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTestDatasource

`func (o *CustomDatasourceConfig) SetIsTestDatasource(v bool)`

SetIsTestDatasource sets IsTestDatasource field to given value.

### HasIsTestDatasource

`func (o *CustomDatasourceConfig) HasIsTestDatasource() bool`

HasIsTestDatasource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


