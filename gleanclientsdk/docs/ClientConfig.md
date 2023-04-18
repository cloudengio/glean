# ClientConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BadVersions** | Pointer to **[]string** | Known bad client versions that should force update themselves | [optional] 
**FeatureFlags** | Pointer to **map[string]bool** |  | [optional] 
**FeedPeopleCelebrationsEnabled** | Pointer to **bool** | Whether people celebrations is enabled or not for the instance | [optional] 
**FeedSuggestedEnabled** | Pointer to **bool** | Whether the suggested feed is enabled | [optional] 
**FeedTrendingEnabled** | Pointer to **bool** | Whether the trending feed is enabled | [optional] 
**FeedRecentsEnabled** | Pointer to **bool** | Whether the recents feed is enabled | [optional] 
**FeedMentionsEnabled** | Pointer to **bool** | Whether the mentions feed is enabled | [optional] 
**BoolValues** | Pointer to **map[string]bool** | A map of {string, boolean} pairs representing flags that globally guard conditional features. Omitted flags mean the client should use its default state. | [optional] 
**CompanyDisplayName** | Pointer to **string** | The user-facing name of the company owning the deployment | [optional] 
**CustomSerpMarkdown** | Pointer to **string** | A markdown string to be displayed on the search results page. Useful for outlinks to help pages. | [optional] 
**FeedbackFrameSrc** | Pointer to **string** | URL of a frame to be displayed for the user to give direct feedback to their company.  A query string parameter named &#x60;context&#x60; is appended to the URL by the client. Its value is a JSON payload containing:  &#x60;&#x60;&#x60; {   datasource: string, // The current datasource tab   department: string, // The user’s department   email: string,      // The user’s email   query: string       // The most recent search query, if any } &#x60;&#x60;&#x60;  | [optional] 
**OnboardingQuery** | Pointer to **string** | A demonstrative query to show during new user onboarding | [optional] 
**IsOrgChartLinkVisible** | Pointer to **bool** | Determines whether the org chart link in the Directory panel is visible to all users. | [optional] 
**IsOrgChartAccessible** | Pointer to **bool** | Determines whether the org chart is accessible to all users, regardless of link visibility. Org chart can be accessible even if the org chart link in Directory is not visible. | [optional] 
**IsPeopleSetup** | Pointer to **bool** | Whether or not people data has been set up. | [optional] 
**WebAppUrl** | Pointer to **string** | URL the company uses to access the web app | [optional] 
**GaTrackingIds** | Pointer to **[]string** | A list of additional GA data stream tracking IDs to which client events should be sent. | [optional] 
**Themes** | Pointer to [**Themes**](Themes.md) |  | [optional] 
**Brandings** | Pointer to [**ClientConfigBrandings**](ClientConfigBrandings.md) |  | [optional] 
**GreetingFormat** | Pointer to **string** | Describes how to format the web app greeting. Possible format options include \\%t - timely greeting \\%n - the user&#39;s first name | [optional] 
**SsoCompanyProvider** | Pointer to **string** | SSO provider used by the company | [optional] 

## Methods

### NewClientConfig

`func NewClientConfig() *ClientConfig`

NewClientConfig instantiates a new ClientConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientConfigWithDefaults

`func NewClientConfigWithDefaults() *ClientConfig`

NewClientConfigWithDefaults instantiates a new ClientConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBadVersions

`func (o *ClientConfig) GetBadVersions() []string`

GetBadVersions returns the BadVersions field if non-nil, zero value otherwise.

### GetBadVersionsOk

`func (o *ClientConfig) GetBadVersionsOk() (*[]string, bool)`

GetBadVersionsOk returns a tuple with the BadVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadVersions

`func (o *ClientConfig) SetBadVersions(v []string)`

SetBadVersions sets BadVersions field to given value.

### HasBadVersions

`func (o *ClientConfig) HasBadVersions() bool`

HasBadVersions returns a boolean if a field has been set.

### GetFeatureFlags

`func (o *ClientConfig) GetFeatureFlags() map[string]bool`

GetFeatureFlags returns the FeatureFlags field if non-nil, zero value otherwise.

### GetFeatureFlagsOk

`func (o *ClientConfig) GetFeatureFlagsOk() (*map[string]bool, bool)`

GetFeatureFlagsOk returns a tuple with the FeatureFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureFlags

`func (o *ClientConfig) SetFeatureFlags(v map[string]bool)`

SetFeatureFlags sets FeatureFlags field to given value.

### HasFeatureFlags

`func (o *ClientConfig) HasFeatureFlags() bool`

HasFeatureFlags returns a boolean if a field has been set.

### GetFeedPeopleCelebrationsEnabled

`func (o *ClientConfig) GetFeedPeopleCelebrationsEnabled() bool`

GetFeedPeopleCelebrationsEnabled returns the FeedPeopleCelebrationsEnabled field if non-nil, zero value otherwise.

### GetFeedPeopleCelebrationsEnabledOk

`func (o *ClientConfig) GetFeedPeopleCelebrationsEnabledOk() (*bool, bool)`

GetFeedPeopleCelebrationsEnabledOk returns a tuple with the FeedPeopleCelebrationsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedPeopleCelebrationsEnabled

`func (o *ClientConfig) SetFeedPeopleCelebrationsEnabled(v bool)`

SetFeedPeopleCelebrationsEnabled sets FeedPeopleCelebrationsEnabled field to given value.

### HasFeedPeopleCelebrationsEnabled

`func (o *ClientConfig) HasFeedPeopleCelebrationsEnabled() bool`

HasFeedPeopleCelebrationsEnabled returns a boolean if a field has been set.

### GetFeedSuggestedEnabled

`func (o *ClientConfig) GetFeedSuggestedEnabled() bool`

GetFeedSuggestedEnabled returns the FeedSuggestedEnabled field if non-nil, zero value otherwise.

### GetFeedSuggestedEnabledOk

`func (o *ClientConfig) GetFeedSuggestedEnabledOk() (*bool, bool)`

GetFeedSuggestedEnabledOk returns a tuple with the FeedSuggestedEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedSuggestedEnabled

`func (o *ClientConfig) SetFeedSuggestedEnabled(v bool)`

SetFeedSuggestedEnabled sets FeedSuggestedEnabled field to given value.

### HasFeedSuggestedEnabled

`func (o *ClientConfig) HasFeedSuggestedEnabled() bool`

HasFeedSuggestedEnabled returns a boolean if a field has been set.

### GetFeedTrendingEnabled

`func (o *ClientConfig) GetFeedTrendingEnabled() bool`

GetFeedTrendingEnabled returns the FeedTrendingEnabled field if non-nil, zero value otherwise.

### GetFeedTrendingEnabledOk

`func (o *ClientConfig) GetFeedTrendingEnabledOk() (*bool, bool)`

GetFeedTrendingEnabledOk returns a tuple with the FeedTrendingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedTrendingEnabled

`func (o *ClientConfig) SetFeedTrendingEnabled(v bool)`

SetFeedTrendingEnabled sets FeedTrendingEnabled field to given value.

### HasFeedTrendingEnabled

`func (o *ClientConfig) HasFeedTrendingEnabled() bool`

HasFeedTrendingEnabled returns a boolean if a field has been set.

### GetFeedRecentsEnabled

`func (o *ClientConfig) GetFeedRecentsEnabled() bool`

GetFeedRecentsEnabled returns the FeedRecentsEnabled field if non-nil, zero value otherwise.

### GetFeedRecentsEnabledOk

`func (o *ClientConfig) GetFeedRecentsEnabledOk() (*bool, bool)`

GetFeedRecentsEnabledOk returns a tuple with the FeedRecentsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedRecentsEnabled

`func (o *ClientConfig) SetFeedRecentsEnabled(v bool)`

SetFeedRecentsEnabled sets FeedRecentsEnabled field to given value.

### HasFeedRecentsEnabled

`func (o *ClientConfig) HasFeedRecentsEnabled() bool`

HasFeedRecentsEnabled returns a boolean if a field has been set.

### GetFeedMentionsEnabled

`func (o *ClientConfig) GetFeedMentionsEnabled() bool`

GetFeedMentionsEnabled returns the FeedMentionsEnabled field if non-nil, zero value otherwise.

### GetFeedMentionsEnabledOk

`func (o *ClientConfig) GetFeedMentionsEnabledOk() (*bool, bool)`

GetFeedMentionsEnabledOk returns a tuple with the FeedMentionsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedMentionsEnabled

`func (o *ClientConfig) SetFeedMentionsEnabled(v bool)`

SetFeedMentionsEnabled sets FeedMentionsEnabled field to given value.

### HasFeedMentionsEnabled

`func (o *ClientConfig) HasFeedMentionsEnabled() bool`

HasFeedMentionsEnabled returns a boolean if a field has been set.

### GetBoolValues

`func (o *ClientConfig) GetBoolValues() map[string]bool`

GetBoolValues returns the BoolValues field if non-nil, zero value otherwise.

### GetBoolValuesOk

`func (o *ClientConfig) GetBoolValuesOk() (*map[string]bool, bool)`

GetBoolValuesOk returns a tuple with the BoolValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoolValues

`func (o *ClientConfig) SetBoolValues(v map[string]bool)`

SetBoolValues sets BoolValues field to given value.

### HasBoolValues

`func (o *ClientConfig) HasBoolValues() bool`

HasBoolValues returns a boolean if a field has been set.

### GetCompanyDisplayName

`func (o *ClientConfig) GetCompanyDisplayName() string`

GetCompanyDisplayName returns the CompanyDisplayName field if non-nil, zero value otherwise.

### GetCompanyDisplayNameOk

`func (o *ClientConfig) GetCompanyDisplayNameOk() (*string, bool)`

GetCompanyDisplayNameOk returns a tuple with the CompanyDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyDisplayName

`func (o *ClientConfig) SetCompanyDisplayName(v string)`

SetCompanyDisplayName sets CompanyDisplayName field to given value.

### HasCompanyDisplayName

`func (o *ClientConfig) HasCompanyDisplayName() bool`

HasCompanyDisplayName returns a boolean if a field has been set.

### GetCustomSerpMarkdown

`func (o *ClientConfig) GetCustomSerpMarkdown() string`

GetCustomSerpMarkdown returns the CustomSerpMarkdown field if non-nil, zero value otherwise.

### GetCustomSerpMarkdownOk

`func (o *ClientConfig) GetCustomSerpMarkdownOk() (*string, bool)`

GetCustomSerpMarkdownOk returns a tuple with the CustomSerpMarkdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSerpMarkdown

`func (o *ClientConfig) SetCustomSerpMarkdown(v string)`

SetCustomSerpMarkdown sets CustomSerpMarkdown field to given value.

### HasCustomSerpMarkdown

`func (o *ClientConfig) HasCustomSerpMarkdown() bool`

HasCustomSerpMarkdown returns a boolean if a field has been set.

### GetFeedbackFrameSrc

`func (o *ClientConfig) GetFeedbackFrameSrc() string`

GetFeedbackFrameSrc returns the FeedbackFrameSrc field if non-nil, zero value otherwise.

### GetFeedbackFrameSrcOk

`func (o *ClientConfig) GetFeedbackFrameSrcOk() (*string, bool)`

GetFeedbackFrameSrcOk returns a tuple with the FeedbackFrameSrc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedbackFrameSrc

`func (o *ClientConfig) SetFeedbackFrameSrc(v string)`

SetFeedbackFrameSrc sets FeedbackFrameSrc field to given value.

### HasFeedbackFrameSrc

`func (o *ClientConfig) HasFeedbackFrameSrc() bool`

HasFeedbackFrameSrc returns a boolean if a field has been set.

### GetOnboardingQuery

`func (o *ClientConfig) GetOnboardingQuery() string`

GetOnboardingQuery returns the OnboardingQuery field if non-nil, zero value otherwise.

### GetOnboardingQueryOk

`func (o *ClientConfig) GetOnboardingQueryOk() (*string, bool)`

GetOnboardingQueryOk returns a tuple with the OnboardingQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnboardingQuery

`func (o *ClientConfig) SetOnboardingQuery(v string)`

SetOnboardingQuery sets OnboardingQuery field to given value.

### HasOnboardingQuery

`func (o *ClientConfig) HasOnboardingQuery() bool`

HasOnboardingQuery returns a boolean if a field has been set.

### GetIsOrgChartLinkVisible

`func (o *ClientConfig) GetIsOrgChartLinkVisible() bool`

GetIsOrgChartLinkVisible returns the IsOrgChartLinkVisible field if non-nil, zero value otherwise.

### GetIsOrgChartLinkVisibleOk

`func (o *ClientConfig) GetIsOrgChartLinkVisibleOk() (*bool, bool)`

GetIsOrgChartLinkVisibleOk returns a tuple with the IsOrgChartLinkVisible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOrgChartLinkVisible

`func (o *ClientConfig) SetIsOrgChartLinkVisible(v bool)`

SetIsOrgChartLinkVisible sets IsOrgChartLinkVisible field to given value.

### HasIsOrgChartLinkVisible

`func (o *ClientConfig) HasIsOrgChartLinkVisible() bool`

HasIsOrgChartLinkVisible returns a boolean if a field has been set.

### GetIsOrgChartAccessible

`func (o *ClientConfig) GetIsOrgChartAccessible() bool`

GetIsOrgChartAccessible returns the IsOrgChartAccessible field if non-nil, zero value otherwise.

### GetIsOrgChartAccessibleOk

`func (o *ClientConfig) GetIsOrgChartAccessibleOk() (*bool, bool)`

GetIsOrgChartAccessibleOk returns a tuple with the IsOrgChartAccessible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOrgChartAccessible

`func (o *ClientConfig) SetIsOrgChartAccessible(v bool)`

SetIsOrgChartAccessible sets IsOrgChartAccessible field to given value.

### HasIsOrgChartAccessible

`func (o *ClientConfig) HasIsOrgChartAccessible() bool`

HasIsOrgChartAccessible returns a boolean if a field has been set.

### GetIsPeopleSetup

`func (o *ClientConfig) GetIsPeopleSetup() bool`

GetIsPeopleSetup returns the IsPeopleSetup field if non-nil, zero value otherwise.

### GetIsPeopleSetupOk

`func (o *ClientConfig) GetIsPeopleSetupOk() (*bool, bool)`

GetIsPeopleSetupOk returns a tuple with the IsPeopleSetup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPeopleSetup

`func (o *ClientConfig) SetIsPeopleSetup(v bool)`

SetIsPeopleSetup sets IsPeopleSetup field to given value.

### HasIsPeopleSetup

`func (o *ClientConfig) HasIsPeopleSetup() bool`

HasIsPeopleSetup returns a boolean if a field has been set.

### GetWebAppUrl

`func (o *ClientConfig) GetWebAppUrl() string`

GetWebAppUrl returns the WebAppUrl field if non-nil, zero value otherwise.

### GetWebAppUrlOk

`func (o *ClientConfig) GetWebAppUrlOk() (*string, bool)`

GetWebAppUrlOk returns a tuple with the WebAppUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebAppUrl

`func (o *ClientConfig) SetWebAppUrl(v string)`

SetWebAppUrl sets WebAppUrl field to given value.

### HasWebAppUrl

`func (o *ClientConfig) HasWebAppUrl() bool`

HasWebAppUrl returns a boolean if a field has been set.

### GetGaTrackingIds

`func (o *ClientConfig) GetGaTrackingIds() []string`

GetGaTrackingIds returns the GaTrackingIds field if non-nil, zero value otherwise.

### GetGaTrackingIdsOk

`func (o *ClientConfig) GetGaTrackingIdsOk() (*[]string, bool)`

GetGaTrackingIdsOk returns a tuple with the GaTrackingIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGaTrackingIds

`func (o *ClientConfig) SetGaTrackingIds(v []string)`

SetGaTrackingIds sets GaTrackingIds field to given value.

### HasGaTrackingIds

`func (o *ClientConfig) HasGaTrackingIds() bool`

HasGaTrackingIds returns a boolean if a field has been set.

### GetThemes

`func (o *ClientConfig) GetThemes() Themes`

GetThemes returns the Themes field if non-nil, zero value otherwise.

### GetThemesOk

`func (o *ClientConfig) GetThemesOk() (*Themes, bool)`

GetThemesOk returns a tuple with the Themes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThemes

`func (o *ClientConfig) SetThemes(v Themes)`

SetThemes sets Themes field to given value.

### HasThemes

`func (o *ClientConfig) HasThemes() bool`

HasThemes returns a boolean if a field has been set.

### GetBrandings

`func (o *ClientConfig) GetBrandings() ClientConfigBrandings`

GetBrandings returns the Brandings field if non-nil, zero value otherwise.

### GetBrandingsOk

`func (o *ClientConfig) GetBrandingsOk() (*ClientConfigBrandings, bool)`

GetBrandingsOk returns a tuple with the Brandings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandings

`func (o *ClientConfig) SetBrandings(v ClientConfigBrandings)`

SetBrandings sets Brandings field to given value.

### HasBrandings

`func (o *ClientConfig) HasBrandings() bool`

HasBrandings returns a boolean if a field has been set.

### GetGreetingFormat

`func (o *ClientConfig) GetGreetingFormat() string`

GetGreetingFormat returns the GreetingFormat field if non-nil, zero value otherwise.

### GetGreetingFormatOk

`func (o *ClientConfig) GetGreetingFormatOk() (*string, bool)`

GetGreetingFormatOk returns a tuple with the GreetingFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGreetingFormat

`func (o *ClientConfig) SetGreetingFormat(v string)`

SetGreetingFormat sets GreetingFormat field to given value.

### HasGreetingFormat

`func (o *ClientConfig) HasGreetingFormat() bool`

HasGreetingFormat returns a boolean if a field has been set.

### GetSsoCompanyProvider

`func (o *ClientConfig) GetSsoCompanyProvider() string`

GetSsoCompanyProvider returns the SsoCompanyProvider field if non-nil, zero value otherwise.

### GetSsoCompanyProviderOk

`func (o *ClientConfig) GetSsoCompanyProviderOk() (*string, bool)`

GetSsoCompanyProviderOk returns a tuple with the SsoCompanyProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoCompanyProvider

`func (o *ClientConfig) SetSsoCompanyProvider(v string)`

SetSsoCompanyProvider sets SsoCompanyProvider field to given value.

### HasSsoCompanyProvider

`func (o *ClientConfig) HasSsoCompanyProvider() bool`

HasSsoCompanyProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


