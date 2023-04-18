/*
Glean Client API - Platform Preview

# Introduction These are all the APIs used by Glean to implement the Glean client. These are available as platform preview for implementing a custom client to the Glean system.  # Usage guidelines A subset of these endpoints are also in the developer ready section, which is available for public use. The rest of the endpoints are subject to prior agreement with Glean before usage. Please contact support@glean.com if you would like to use an API that is not currently available in the developer ready section. 

API version: 0.9.0
Contact: support@glean.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gleanclientsdk

import (
	"encoding/json"
)

// checks if the SharedDatasourceConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SharedDatasourceConfig{}

// SharedDatasourceConfig Structure describing shared config properties of the datasource (including multi-instance support)
type SharedDatasourceConfig struct {
	// Unique identifier of datasource instance to which this config applies.
	Name string `json:"name"`
	// The user-friendly instance label to display. If omitted, falls back to the title-cased `name`.
	DisplayName *string `json:"displayName,omitempty"`
	// The type of this datasource. It is an important signal for relevance and must be specified and cannot be UNCATEGORIZED.
	DatasourceCategory *string `json:"datasourceCategory,omitempty"`
	// Regular expression that matches URLs of documents of the datasource instance. The behavior for multiple matches is non-deterministic. **Note: urlRegex is a required field for non-entity datasources (ie. datasources where isEntityDatasource is false). Please add a regex as specific as possible to this datasource instance.**
	UrlRegex *string `json:"urlRegex,omitempty"`
	// The URL to an image to be displayed as an icon for this datasource instance. Must have a transparency mask. SVG are recommended over PNG. Public, scio-authenticated and Base64 encoded data URLs are all valid (but not third-party-authenticated URLs).
	IconUrl *string `json:"iconUrl,omitempty"`
	// The list of top-level `objectType`s for the datasource.
	ObjectDefinitions []ObjectDefinition `json:"objectDefinitions,omitempty"`
	// Example text for what to search for in this datasource
	SuggestionText *string `json:"suggestionText,omitempty"`
	// The URL of the landing page for this datasource instance. Should point to the most useful page for users, not the company marketing page.
	HomeUrl *string `json:"homeUrl,omitempty"`
	// This only applies to WEB_CRAWL and BROWSER_CRAWL datasources. Defines the seed urls for crawling.
	CrawlerSeedUrls []string `json:"crawlerSeedUrls,omitempty"`
	// The URL to an image to be displayed as an icon for this datasource instance in dark mode. Must have a transparency mask. SVG are recommended over PNG. Public, scio-authenticated and Base64 encoded data URLs are all valid (but not third-party-authenticated URLs).
	IconDarkUrl *string `json:"iconDarkUrl,omitempty"`
	// List of built-in facet types that should be hidden for the datasource.
	HideBuiltInFacets []string `json:"hideBuiltInFacets,omitempty"`
	// A list of regular expressions to apply to an arbitrary URL to transform it into a canonical URL for this datasource instance. Regexes are to be applied in the order specified in this list.
	CanonicalizingURLRegex []CanonicalizingRegexType `json:"canonicalizingURLRegex,omitempty"`
	// A list of regular expressions to apply to an arbitrary title to transform it into a title that will be displayed in the search results
	CanonicalizingTitleRegex []CanonicalizingRegexType `json:"canonicalizingTitleRegex,omitempty"`
	// A regex that identifies titles that should not be indexed
	RedlistTitleRegex *string `json:"redlistTitleRegex,omitempty"`
	ConnectorType *ConnectorType `json:"connectorType,omitempty"`
	// List of actions for this datasource instance that will show up in autocomplete and app card, e.g. \"Create new issue\" for jira
	Quicklinks []Quicklink `json:"quicklinks,omitempty"`
	// The name of a render config to use for displaying results from this datasource. Any well known datasource name may be used to render the same as that source, e.g. `web` or `gdrive`.
	RenderConfigPreset *string `json:"renderConfigPreset,omitempty"`
	// Aliases that can be used as `app` operator-values.
	Aliases []string `json:"aliases,omitempty"`
	// Whether or not this datasource is hosted on-premise.
	IsOnPrem *bool `json:"isOnPrem,omitempty"`
	// True if browser activity is able to report the correct URL for VIEW events. Set this to true if the URLs reported by Chrome are constant throughout each page load. Set this to false if the page has Javascript that modifies the URL during or after the load.
	TrustUrlRegexForViewActivity *bool `json:"trustUrlRegexForViewActivity,omitempty"`
	// If true, a utm_source query param will be added to outbound links to this datasource within Glean.
	IncludeUtmSource *bool `json:"includeUtmSource,omitempty"`
	// The (non-unique) name of the datasource of which this config is an instance, e.g. \"jira\".
	DatasourceName *string `json:"datasourceName,omitempty"`
	// The instance of the datasource for this particular config, e.g. \"onprem\".
	InstanceOnlyName *string `json:"instanceOnlyName,omitempty"`
	// A human readable string identifying this instance as compared to its peers, e.g. \"github.com/askscio\" or \"github.askscio.com\"
	InstanceDescription *string `json:"instanceDescription,omitempty"`
}

// NewSharedDatasourceConfig instantiates a new SharedDatasourceConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSharedDatasourceConfig(name string) *SharedDatasourceConfig {
	this := SharedDatasourceConfig{}
	this.Name = name
	var datasourceCategory string = "UNCATEGORIZED"
	this.DatasourceCategory = &datasourceCategory
	var trustUrlRegexForViewActivity bool = true
	this.TrustUrlRegexForViewActivity = &trustUrlRegexForViewActivity
	return &this
}

// NewSharedDatasourceConfigWithDefaults instantiates a new SharedDatasourceConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSharedDatasourceConfigWithDefaults() *SharedDatasourceConfig {
	this := SharedDatasourceConfig{}
	var datasourceCategory string = "UNCATEGORIZED"
	this.DatasourceCategory = &datasourceCategory
	var trustUrlRegexForViewActivity bool = true
	this.TrustUrlRegexForViewActivity = &trustUrlRegexForViewActivity
	return &this
}

// GetName returns the Name field value
func (o *SharedDatasourceConfig) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SharedDatasourceConfig) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SharedDatasourceConfig) SetName(v string) {
	o.Name = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *SharedDatasourceConfig) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedDatasourceConfig) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *SharedDatasourceConfig) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *SharedDatasourceConfig) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDatasourceCategory returns the DatasourceCategory field value if set, zero value otherwise.
func (o *SharedDatasourceConfig) GetDatasourceCategory() string {
	if o == nil || IsNil(o.DatasourceCategory) {
		var ret string
		return ret
	}
	return *o.DatasourceCategory
}

// GetDatasourceCategoryOk returns a tuple with the DatasourceCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedDatasourceConfig) GetDatasourceCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.DatasourceCategory) {
		return nil, false
	}
	return o.DatasourceCategory, true
}

// HasDatasourceCategory returns a boolean if a field has been set.
func (o *SharedDatasourceConfig) HasDatasourceCategory() bool {
	if o != nil && !IsNil(o.DatasourceCategory) {
		return true
	}

	return false
}

// SetDatasourceCategory gets a reference to the given string and assigns it to the DatasourceCategory field.
func (o *SharedDatasourceConfig) SetDatasourceCategory(v string) {
	o.DatasourceCategory = &v
}

// GetUrlRegex returns the UrlRegex field value if set, zero value otherwise.
func (o *SharedDatasourceConfig) GetUrlRegex() string {
	if o == nil || IsNil(o.UrlRegex) {
		var ret string
		return ret
	}
	return *o.UrlRegex
}

// GetUrlRegexOk returns a tuple with the UrlRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedDatasourceConfig) GetUrlRegexOk() (*string, bool) {
	if o == nil || IsNil(o.UrlRegex) {
		return nil, false
	}
	return o.UrlRegex, true
}

// HasUrlRegex returns a boolean if a field has been set.
func (o *SharedDatasourceConfig) HasUrlRegex() bool {
	if o != nil && !IsNil(o.UrlRegex) {
		return true
	}

	return false
}

// SetUrlRegex gets a reference to the given string and assigns it to the UrlRegex field.
func (o *SharedDatasourceConfig) SetUrlRegex(v string) {
	o.UrlRegex = &v
}

// GetIconUrl returns the IconUrl field value if set, zero value otherwise.
func (o *SharedDatasourceConfig) GetIconUrl() string {
	if o == nil || IsNil(o.IconUrl) {
		var ret string
		return ret
	}
	return *o.IconUrl
}

// GetIconUrlOk returns a tuple with the IconUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedDatasourceConfig) GetIconUrlOk() (*string, bool) {
	if o == nil || IsNil(o.IconUrl) {
		return nil, false
	}
	return o.IconUrl, true
}

// HasIconUrl returns a boolean if a field has been set.
func (o *SharedDatasourceConfig) HasIconUrl() bool {
	if o != nil && !IsNil(o.IconUrl) {
		return true
	}

	return false
}

// SetIconUrl gets a reference to the given string and assigns it to the IconUrl field.
func (o *SharedDatasourceConfig) SetIconUrl(v string) {
	o.IconUrl = &v
}

// GetObjectDefinitions returns the ObjectDefinitions field value if set, zero value otherwise.
func (o *SharedDatasourceConfig) GetObjectDefinitions() []ObjectDefinition {
	if o == nil || IsNil(o.ObjectDefinitions) {
		var ret []ObjectDefinition
		return ret
	}
	return o.ObjectDefinitions
}

// GetObjectDefinitionsOk returns a tuple with the ObjectDefinitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedDatasourceConfig) GetObjectDefinitionsOk() ([]ObjectDefinition, bool) {
	if o == nil || IsNil(o.ObjectDefinitions) {
		return nil, false
	}
	return o.ObjectDefinitions, true
}

// HasObjectDefinitions returns a boolean if a field has been set.
func (o *SharedDatasourceConfig) HasObjectDefinitions() bool {
	if o != nil && !IsNil(o.ObjectDefinitions) {
		return true
	}

	return false
}

// SetObjectDefinitions gets a reference to the given []ObjectDefinition and assigns it to the ObjectDefinitions field.
func (o *SharedDatasourceConfig) SetObjectDefinitions(v []ObjectDefinition) {
	o.ObjectDefinitions = v
}

// GetSuggestionText returns the SuggestionText field value if set, zero value otherwise.
func (o *SharedDatasourceConfig) GetSuggestionText() string {
	if o == nil || IsNil(o.SuggestionText) {
		var ret string
		return ret
	}
	return *o.SuggestionText
}

// GetSuggestionTextOk returns a tuple with the SuggestionText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedDatasourceConfig) GetSuggestionTextOk() (*string, bool) {
	if o == nil || IsNil(o.SuggestionText) {
		return nil, false
	}
	return o.SuggestionText, true
}

// HasSuggestionText returns a boolean if a field has been set.
func (o *SharedDatasourceConfig) HasSuggestionText() bool {
	if o != nil && !IsNil(o.SuggestionText) {
		return true
	}

	return false
}

// SetSuggestionText gets a reference to the given string and assigns it to the SuggestionText field.
func (o *SharedDatasourceConfig) SetSuggestionText(v string) {
	o.SuggestionText = &v
}

// GetHomeUrl returns the HomeUrl field value if set, zero value otherwise.
func (o *SharedDatasourceConfig) GetHomeUrl() string {
	if o == nil || IsNil(o.HomeUrl) {
		var ret string
		return ret
	}
	return *o.HomeUrl
}

// GetHomeUrlOk returns a tuple with the HomeUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedDatasourceConfig) GetHomeUrlOk() (*string, bool) {
	if o == nil || IsNil(o.HomeUrl) {
		return nil, false
	}
	return o.HomeUrl, true
}

// HasHomeUrl returns a boolean if a field has been set.
func (o *SharedDatasourceConfig) HasHomeUrl() bool {
	if o != nil && !IsNil(o.HomeUrl) {
		return true
	}

	return false
}

// SetHomeUrl gets a reference to the given string and assigns it to the HomeUrl field.
func (o *SharedDatasourceConfig) SetHomeUrl(v string) {
	o.HomeUrl = &v
}

// GetCrawlerSeedUrls returns the CrawlerSeedUrls field value if set, zero value otherwise.
func (o *SharedDatasourceConfig) GetCrawlerSeedUrls() []string {
	if o == nil || IsNil(o.CrawlerSeedUrls) {
		var ret []string
		return ret
	}
	return o.CrawlerSeedUrls
}

// GetCrawlerSeedUrlsOk returns a tuple with the CrawlerSeedUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedDatasourceConfig) GetCrawlerSeedUrlsOk() ([]string, bool) {
	if o == nil || IsNil(o.CrawlerSeedUrls) {
		return nil, false
	}
	return o.CrawlerSeedUrls, true
}

// HasCrawlerSeedUrls returns a boolean if a field has been set.
func (o *SharedDatasourceConfig) HasCrawlerSeedUrls() bool {
	if o != nil && !IsNil(o.CrawlerSeedUrls) {
		return true
	}

	return false
}

// SetCrawlerSeedUrls gets a reference to the given []string and assigns it to the CrawlerSeedUrls field.
func (o *SharedDatasourceConfig) SetCrawlerSeedUrls(v []string) {
	o.CrawlerSeedUrls = v
}

// GetIconDarkUrl returns the IconDarkUrl field value if set, zero value otherwise.
func (o *SharedDatasourceConfig) GetIconDarkUrl() string {
	if o == nil || IsNil(o.IconDarkUrl) {
		var ret string
		return ret
	}
	return *o.IconDarkUrl
}

// GetIconDarkUrlOk returns a tuple with the IconDarkUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedDatasourceConfig) GetIconDarkUrlOk() (*string, bool) {
	if o == nil || IsNil(o.IconDarkUrl) {
		return nil, false
	}
	return o.IconDarkUrl, true
}

// HasIconDarkUrl returns a boolean if a field has been set.
func (o *SharedDatasourceConfig) HasIconDarkUrl() bool {
	if o != nil && !IsNil(o.IconDarkUrl) {
		return true
	}

	return false
}

// SetIconDarkUrl gets a reference to the given string and assigns it to the IconDarkUrl field.
func (o *SharedDatasourceConfig) SetIconDarkUrl(v string) {
	o.IconDarkUrl = &v
}

// GetHideBuiltInFacets returns the HideBuiltInFacets field value if set, zero value otherwise.
func (o *SharedDatasourceConfig) GetHideBuiltInFacets() []string {
	if o == nil || IsNil(o.HideBuiltInFacets) {
		var ret []string
		return ret
	}
	return o.HideBuiltInFacets
}

// GetHideBuiltInFacetsOk returns a tuple with the HideBuiltInFacets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedDatasourceConfig) GetHideBuiltInFacetsOk() ([]string, bool) {
	if o == nil || IsNil(o.HideBuiltInFacets) {
		return nil, false
	}
	return o.HideBuiltInFacets, true
}

// HasHideBuiltInFacets returns a boolean if a field has been set.
func (o *SharedDatasourceConfig) HasHideBuiltInFacets() bool {
	if o != nil && !IsNil(o.HideBuiltInFacets) {
		return true
	}

	return false
}

// SetHideBuiltInFacets gets a reference to the given []string and assigns it to the HideBuiltInFacets field.
func (o *SharedDatasourceConfig) SetHideBuiltInFacets(v []string) {
	o.HideBuiltInFacets = v
}

// GetCanonicalizingURLRegex returns the CanonicalizingURLRegex field value if set, zero value otherwise.
func (o *SharedDatasourceConfig) GetCanonicalizingURLRegex() []CanonicalizingRegexType {
	if o == nil || IsNil(o.CanonicalizingURLRegex) {
		var ret []CanonicalizingRegexType
		return ret
	}
	return o.CanonicalizingURLRegex
}

// GetCanonicalizingURLRegexOk returns a tuple with the CanonicalizingURLRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedDatasourceConfig) GetCanonicalizingURLRegexOk() ([]CanonicalizingRegexType, bool) {
	if o == nil || IsNil(o.CanonicalizingURLRegex) {
		return nil, false
	}
	return o.CanonicalizingURLRegex, true
}

// HasCanonicalizingURLRegex returns a boolean if a field has been set.
func (o *SharedDatasourceConfig) HasCanonicalizingURLRegex() bool {
	if o != nil && !IsNil(o.CanonicalizingURLRegex) {
		return true
	}

	return false
}

// SetCanonicalizingURLRegex gets a reference to the given []CanonicalizingRegexType and assigns it to the CanonicalizingURLRegex field.
func (o *SharedDatasourceConfig) SetCanonicalizingURLRegex(v []CanonicalizingRegexType) {
	o.CanonicalizingURLRegex = v
}

// GetCanonicalizingTitleRegex returns the CanonicalizingTitleRegex field value if set, zero value otherwise.
func (o *SharedDatasourceConfig) GetCanonicalizingTitleRegex() []CanonicalizingRegexType {
	if o == nil || IsNil(o.CanonicalizingTitleRegex) {
		var ret []CanonicalizingRegexType
		return ret
	}
	return o.CanonicalizingTitleRegex
}

// GetCanonicalizingTitleRegexOk returns a tuple with the CanonicalizingTitleRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedDatasourceConfig) GetCanonicalizingTitleRegexOk() ([]CanonicalizingRegexType, bool) {
	if o == nil || IsNil(o.CanonicalizingTitleRegex) {
		return nil, false
	}
	return o.CanonicalizingTitleRegex, true
}

// HasCanonicalizingTitleRegex returns a boolean if a field has been set.
func (o *SharedDatasourceConfig) HasCanonicalizingTitleRegex() bool {
	if o != nil && !IsNil(o.CanonicalizingTitleRegex) {
		return true
	}

	return false
}

// SetCanonicalizingTitleRegex gets a reference to the given []CanonicalizingRegexType and assigns it to the CanonicalizingTitleRegex field.
func (o *SharedDatasourceConfig) SetCanonicalizingTitleRegex(v []CanonicalizingRegexType) {
	o.CanonicalizingTitleRegex = v
}

// GetRedlistTitleRegex returns the RedlistTitleRegex field value if set, zero value otherwise.
func (o *SharedDatasourceConfig) GetRedlistTitleRegex() string {
	if o == nil || IsNil(o.RedlistTitleRegex) {
		var ret string
		return ret
	}
	return *o.RedlistTitleRegex
}

// GetRedlistTitleRegexOk returns a tuple with the RedlistTitleRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedDatasourceConfig) GetRedlistTitleRegexOk() (*string, bool) {
	if o == nil || IsNil(o.RedlistTitleRegex) {
		return nil, false
	}
	return o.RedlistTitleRegex, true
}

// HasRedlistTitleRegex returns a boolean if a field has been set.
func (o *SharedDatasourceConfig) HasRedlistTitleRegex() bool {
	if o != nil && !IsNil(o.RedlistTitleRegex) {
		return true
	}

	return false
}

// SetRedlistTitleRegex gets a reference to the given string and assigns it to the RedlistTitleRegex field.
func (o *SharedDatasourceConfig) SetRedlistTitleRegex(v string) {
	o.RedlistTitleRegex = &v
}

// GetConnectorType returns the ConnectorType field value if set, zero value otherwise.
func (o *SharedDatasourceConfig) GetConnectorType() ConnectorType {
	if o == nil || IsNil(o.ConnectorType) {
		var ret ConnectorType
		return ret
	}
	return *o.ConnectorType
}

// GetConnectorTypeOk returns a tuple with the ConnectorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedDatasourceConfig) GetConnectorTypeOk() (*ConnectorType, bool) {
	if o == nil || IsNil(o.ConnectorType) {
		return nil, false
	}
	return o.ConnectorType, true
}

// HasConnectorType returns a boolean if a field has been set.
func (o *SharedDatasourceConfig) HasConnectorType() bool {
	if o != nil && !IsNil(o.ConnectorType) {
		return true
	}

	return false
}

// SetConnectorType gets a reference to the given ConnectorType and assigns it to the ConnectorType field.
func (o *SharedDatasourceConfig) SetConnectorType(v ConnectorType) {
	o.ConnectorType = &v
}

// GetQuicklinks returns the Quicklinks field value if set, zero value otherwise.
func (o *SharedDatasourceConfig) GetQuicklinks() []Quicklink {
	if o == nil || IsNil(o.Quicklinks) {
		var ret []Quicklink
		return ret
	}
	return o.Quicklinks
}

// GetQuicklinksOk returns a tuple with the Quicklinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedDatasourceConfig) GetQuicklinksOk() ([]Quicklink, bool) {
	if o == nil || IsNil(o.Quicklinks) {
		return nil, false
	}
	return o.Quicklinks, true
}

// HasQuicklinks returns a boolean if a field has been set.
func (o *SharedDatasourceConfig) HasQuicklinks() bool {
	if o != nil && !IsNil(o.Quicklinks) {
		return true
	}

	return false
}

// SetQuicklinks gets a reference to the given []Quicklink and assigns it to the Quicklinks field.
func (o *SharedDatasourceConfig) SetQuicklinks(v []Quicklink) {
	o.Quicklinks = v
}

// GetRenderConfigPreset returns the RenderConfigPreset field value if set, zero value otherwise.
func (o *SharedDatasourceConfig) GetRenderConfigPreset() string {
	if o == nil || IsNil(o.RenderConfigPreset) {
		var ret string
		return ret
	}
	return *o.RenderConfigPreset
}

// GetRenderConfigPresetOk returns a tuple with the RenderConfigPreset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedDatasourceConfig) GetRenderConfigPresetOk() (*string, bool) {
	if o == nil || IsNil(o.RenderConfigPreset) {
		return nil, false
	}
	return o.RenderConfigPreset, true
}

// HasRenderConfigPreset returns a boolean if a field has been set.
func (o *SharedDatasourceConfig) HasRenderConfigPreset() bool {
	if o != nil && !IsNil(o.RenderConfigPreset) {
		return true
	}

	return false
}

// SetRenderConfigPreset gets a reference to the given string and assigns it to the RenderConfigPreset field.
func (o *SharedDatasourceConfig) SetRenderConfigPreset(v string) {
	o.RenderConfigPreset = &v
}

// GetAliases returns the Aliases field value if set, zero value otherwise.
func (o *SharedDatasourceConfig) GetAliases() []string {
	if o == nil || IsNil(o.Aliases) {
		var ret []string
		return ret
	}
	return o.Aliases
}

// GetAliasesOk returns a tuple with the Aliases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedDatasourceConfig) GetAliasesOk() ([]string, bool) {
	if o == nil || IsNil(o.Aliases) {
		return nil, false
	}
	return o.Aliases, true
}

// HasAliases returns a boolean if a field has been set.
func (o *SharedDatasourceConfig) HasAliases() bool {
	if o != nil && !IsNil(o.Aliases) {
		return true
	}

	return false
}

// SetAliases gets a reference to the given []string and assigns it to the Aliases field.
func (o *SharedDatasourceConfig) SetAliases(v []string) {
	o.Aliases = v
}

// GetIsOnPrem returns the IsOnPrem field value if set, zero value otherwise.
func (o *SharedDatasourceConfig) GetIsOnPrem() bool {
	if o == nil || IsNil(o.IsOnPrem) {
		var ret bool
		return ret
	}
	return *o.IsOnPrem
}

// GetIsOnPremOk returns a tuple with the IsOnPrem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedDatasourceConfig) GetIsOnPremOk() (*bool, bool) {
	if o == nil || IsNil(o.IsOnPrem) {
		return nil, false
	}
	return o.IsOnPrem, true
}

// HasIsOnPrem returns a boolean if a field has been set.
func (o *SharedDatasourceConfig) HasIsOnPrem() bool {
	if o != nil && !IsNil(o.IsOnPrem) {
		return true
	}

	return false
}

// SetIsOnPrem gets a reference to the given bool and assigns it to the IsOnPrem field.
func (o *SharedDatasourceConfig) SetIsOnPrem(v bool) {
	o.IsOnPrem = &v
}

// GetTrustUrlRegexForViewActivity returns the TrustUrlRegexForViewActivity field value if set, zero value otherwise.
func (o *SharedDatasourceConfig) GetTrustUrlRegexForViewActivity() bool {
	if o == nil || IsNil(o.TrustUrlRegexForViewActivity) {
		var ret bool
		return ret
	}
	return *o.TrustUrlRegexForViewActivity
}

// GetTrustUrlRegexForViewActivityOk returns a tuple with the TrustUrlRegexForViewActivity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedDatasourceConfig) GetTrustUrlRegexForViewActivityOk() (*bool, bool) {
	if o == nil || IsNil(o.TrustUrlRegexForViewActivity) {
		return nil, false
	}
	return o.TrustUrlRegexForViewActivity, true
}

// HasTrustUrlRegexForViewActivity returns a boolean if a field has been set.
func (o *SharedDatasourceConfig) HasTrustUrlRegexForViewActivity() bool {
	if o != nil && !IsNil(o.TrustUrlRegexForViewActivity) {
		return true
	}

	return false
}

// SetTrustUrlRegexForViewActivity gets a reference to the given bool and assigns it to the TrustUrlRegexForViewActivity field.
func (o *SharedDatasourceConfig) SetTrustUrlRegexForViewActivity(v bool) {
	o.TrustUrlRegexForViewActivity = &v
}

// GetIncludeUtmSource returns the IncludeUtmSource field value if set, zero value otherwise.
func (o *SharedDatasourceConfig) GetIncludeUtmSource() bool {
	if o == nil || IsNil(o.IncludeUtmSource) {
		var ret bool
		return ret
	}
	return *o.IncludeUtmSource
}

// GetIncludeUtmSourceOk returns a tuple with the IncludeUtmSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedDatasourceConfig) GetIncludeUtmSourceOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeUtmSource) {
		return nil, false
	}
	return o.IncludeUtmSource, true
}

// HasIncludeUtmSource returns a boolean if a field has been set.
func (o *SharedDatasourceConfig) HasIncludeUtmSource() bool {
	if o != nil && !IsNil(o.IncludeUtmSource) {
		return true
	}

	return false
}

// SetIncludeUtmSource gets a reference to the given bool and assigns it to the IncludeUtmSource field.
func (o *SharedDatasourceConfig) SetIncludeUtmSource(v bool) {
	o.IncludeUtmSource = &v
}

// GetDatasourceName returns the DatasourceName field value if set, zero value otherwise.
func (o *SharedDatasourceConfig) GetDatasourceName() string {
	if o == nil || IsNil(o.DatasourceName) {
		var ret string
		return ret
	}
	return *o.DatasourceName
}

// GetDatasourceNameOk returns a tuple with the DatasourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedDatasourceConfig) GetDatasourceNameOk() (*string, bool) {
	if o == nil || IsNil(o.DatasourceName) {
		return nil, false
	}
	return o.DatasourceName, true
}

// HasDatasourceName returns a boolean if a field has been set.
func (o *SharedDatasourceConfig) HasDatasourceName() bool {
	if o != nil && !IsNil(o.DatasourceName) {
		return true
	}

	return false
}

// SetDatasourceName gets a reference to the given string and assigns it to the DatasourceName field.
func (o *SharedDatasourceConfig) SetDatasourceName(v string) {
	o.DatasourceName = &v
}

// GetInstanceOnlyName returns the InstanceOnlyName field value if set, zero value otherwise.
func (o *SharedDatasourceConfig) GetInstanceOnlyName() string {
	if o == nil || IsNil(o.InstanceOnlyName) {
		var ret string
		return ret
	}
	return *o.InstanceOnlyName
}

// GetInstanceOnlyNameOk returns a tuple with the InstanceOnlyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedDatasourceConfig) GetInstanceOnlyNameOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceOnlyName) {
		return nil, false
	}
	return o.InstanceOnlyName, true
}

// HasInstanceOnlyName returns a boolean if a field has been set.
func (o *SharedDatasourceConfig) HasInstanceOnlyName() bool {
	if o != nil && !IsNil(o.InstanceOnlyName) {
		return true
	}

	return false
}

// SetInstanceOnlyName gets a reference to the given string and assigns it to the InstanceOnlyName field.
func (o *SharedDatasourceConfig) SetInstanceOnlyName(v string) {
	o.InstanceOnlyName = &v
}

// GetInstanceDescription returns the InstanceDescription field value if set, zero value otherwise.
func (o *SharedDatasourceConfig) GetInstanceDescription() string {
	if o == nil || IsNil(o.InstanceDescription) {
		var ret string
		return ret
	}
	return *o.InstanceDescription
}

// GetInstanceDescriptionOk returns a tuple with the InstanceDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedDatasourceConfig) GetInstanceDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceDescription) {
		return nil, false
	}
	return o.InstanceDescription, true
}

// HasInstanceDescription returns a boolean if a field has been set.
func (o *SharedDatasourceConfig) HasInstanceDescription() bool {
	if o != nil && !IsNil(o.InstanceDescription) {
		return true
	}

	return false
}

// SetInstanceDescription gets a reference to the given string and assigns it to the InstanceDescription field.
func (o *SharedDatasourceConfig) SetInstanceDescription(v string) {
	o.InstanceDescription = &v
}

func (o SharedDatasourceConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SharedDatasourceConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.DatasourceCategory) {
		toSerialize["datasourceCategory"] = o.DatasourceCategory
	}
	if !IsNil(o.UrlRegex) {
		toSerialize["urlRegex"] = o.UrlRegex
	}
	if !IsNil(o.IconUrl) {
		toSerialize["iconUrl"] = o.IconUrl
	}
	if !IsNil(o.ObjectDefinitions) {
		toSerialize["objectDefinitions"] = o.ObjectDefinitions
	}
	if !IsNil(o.SuggestionText) {
		toSerialize["suggestionText"] = o.SuggestionText
	}
	if !IsNil(o.HomeUrl) {
		toSerialize["homeUrl"] = o.HomeUrl
	}
	if !IsNil(o.CrawlerSeedUrls) {
		toSerialize["crawlerSeedUrls"] = o.CrawlerSeedUrls
	}
	if !IsNil(o.IconDarkUrl) {
		toSerialize["iconDarkUrl"] = o.IconDarkUrl
	}
	if !IsNil(o.HideBuiltInFacets) {
		toSerialize["hideBuiltInFacets"] = o.HideBuiltInFacets
	}
	if !IsNil(o.CanonicalizingURLRegex) {
		toSerialize["canonicalizingURLRegex"] = o.CanonicalizingURLRegex
	}
	if !IsNil(o.CanonicalizingTitleRegex) {
		toSerialize["canonicalizingTitleRegex"] = o.CanonicalizingTitleRegex
	}
	if !IsNil(o.RedlistTitleRegex) {
		toSerialize["redlistTitleRegex"] = o.RedlistTitleRegex
	}
	if !IsNil(o.ConnectorType) {
		toSerialize["connectorType"] = o.ConnectorType
	}
	if !IsNil(o.Quicklinks) {
		toSerialize["quicklinks"] = o.Quicklinks
	}
	if !IsNil(o.RenderConfigPreset) {
		toSerialize["renderConfigPreset"] = o.RenderConfigPreset
	}
	if !IsNil(o.Aliases) {
		toSerialize["aliases"] = o.Aliases
	}
	if !IsNil(o.IsOnPrem) {
		toSerialize["isOnPrem"] = o.IsOnPrem
	}
	if !IsNil(o.TrustUrlRegexForViewActivity) {
		toSerialize["trustUrlRegexForViewActivity"] = o.TrustUrlRegexForViewActivity
	}
	if !IsNil(o.IncludeUtmSource) {
		toSerialize["includeUtmSource"] = o.IncludeUtmSource
	}
	if !IsNil(o.DatasourceName) {
		toSerialize["datasourceName"] = o.DatasourceName
	}
	if !IsNil(o.InstanceOnlyName) {
		toSerialize["instanceOnlyName"] = o.InstanceOnlyName
	}
	if !IsNil(o.InstanceDescription) {
		toSerialize["instanceDescription"] = o.InstanceDescription
	}
	return toSerialize, nil
}

type NullableSharedDatasourceConfig struct {
	value *SharedDatasourceConfig
	isSet bool
}

func (v NullableSharedDatasourceConfig) Get() *SharedDatasourceConfig {
	return v.value
}

func (v *NullableSharedDatasourceConfig) Set(val *SharedDatasourceConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableSharedDatasourceConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableSharedDatasourceConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSharedDatasourceConfig(val *SharedDatasourceConfig) *NullableSharedDatasourceConfig {
	return &NullableSharedDatasourceConfig{value: val, isSet: true}
}

func (v NullableSharedDatasourceConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSharedDatasourceConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


