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

// checks if the ClientConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClientConfig{}

// ClientConfig Configuration settings for a specific client deployment that are not related to any particular datasource
type ClientConfig struct {
	// Known bad client versions that should force update themselves
	BadVersions []string `json:"badVersions,omitempty"`
	FeatureFlags *map[string]bool `json:"featureFlags,omitempty"`
	// Whether people celebrations is enabled or not for the instance
	FeedPeopleCelebrationsEnabled *bool `json:"feedPeopleCelebrationsEnabled,omitempty"`
	// Whether the suggested feed is enabled
	FeedSuggestedEnabled *bool `json:"feedSuggestedEnabled,omitempty"`
	// Whether the trending feed is enabled
	FeedTrendingEnabled *bool `json:"feedTrendingEnabled,omitempty"`
	// Whether the recents feed is enabled
	FeedRecentsEnabled *bool `json:"feedRecentsEnabled,omitempty"`
	// Whether the mentions feed is enabled
	FeedMentionsEnabled *bool `json:"feedMentionsEnabled,omitempty"`
	// A map of {string, boolean} pairs representing flags that globally guard conditional features. Omitted flags mean the client should use its default state.
	BoolValues *map[string]bool `json:"boolValues,omitempty"`
	// The user-facing name of the company owning the deployment
	CompanyDisplayName *string `json:"companyDisplayName,omitempty"`
	// A markdown string to be displayed on the search results page. Useful for outlinks to help pages.
	CustomSerpMarkdown *string `json:"customSerpMarkdown,omitempty"`
	// URL of a frame to be displayed for the user to give direct feedback to their company.  A query string parameter named `context` is appended to the URL by the client. Its value is a JSON payload containing:  ``` {   datasource: string, // The current datasource tab   department: string, // The user’s department   email: string,      // The user’s email   query: string       // The most recent search query, if any } ``` 
	FeedbackFrameSrc *string `json:"feedbackFrameSrc,omitempty"`
	// A demonstrative query to show during new user onboarding
	OnboardingQuery *string `json:"onboardingQuery,omitempty"`
	// Determines whether the org chart link in the Directory panel is visible to all users.
	IsOrgChartLinkVisible *bool `json:"isOrgChartLinkVisible,omitempty"`
	// Determines whether the org chart is accessible to all users, regardless of link visibility. Org chart can be accessible even if the org chart link in Directory is not visible.
	IsOrgChartAccessible *bool `json:"isOrgChartAccessible,omitempty"`
	// Whether or not people data has been set up.
	IsPeopleSetup *bool `json:"isPeopleSetup,omitempty"`
	// URL the company uses to access the web app
	WebAppUrl *string `json:"webAppUrl,omitempty"`
	// A list of additional GA data stream tracking IDs to which client events should be sent.
	GaTrackingIds []string `json:"gaTrackingIds,omitempty"`
	Themes *Themes `json:"themes,omitempty"`
	Brandings *ClientConfigBrandings `json:"brandings,omitempty"`
	// Describes how to format the web app greeting. Possible format options include \\%t - timely greeting \\%n - the user's first name
	GreetingFormat *string `json:"greetingFormat,omitempty"`
	// SSO provider used by the company
	SsoCompanyProvider *string `json:"ssoCompanyProvider,omitempty"`
}

// NewClientConfig instantiates a new ClientConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientConfig() *ClientConfig {
	this := ClientConfig{}
	return &this
}

// NewClientConfigWithDefaults instantiates a new ClientConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientConfigWithDefaults() *ClientConfig {
	this := ClientConfig{}
	return &this
}

// GetBadVersions returns the BadVersions field value if set, zero value otherwise.
func (o *ClientConfig) GetBadVersions() []string {
	if o == nil || IsNil(o.BadVersions) {
		var ret []string
		return ret
	}
	return o.BadVersions
}

// GetBadVersionsOk returns a tuple with the BadVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientConfig) GetBadVersionsOk() ([]string, bool) {
	if o == nil || IsNil(o.BadVersions) {
		return nil, false
	}
	return o.BadVersions, true
}

// HasBadVersions returns a boolean if a field has been set.
func (o *ClientConfig) HasBadVersions() bool {
	if o != nil && !IsNil(o.BadVersions) {
		return true
	}

	return false
}

// SetBadVersions gets a reference to the given []string and assigns it to the BadVersions field.
func (o *ClientConfig) SetBadVersions(v []string) {
	o.BadVersions = v
}

// GetFeatureFlags returns the FeatureFlags field value if set, zero value otherwise.
func (o *ClientConfig) GetFeatureFlags() map[string]bool {
	if o == nil || IsNil(o.FeatureFlags) {
		var ret map[string]bool
		return ret
	}
	return *o.FeatureFlags
}

// GetFeatureFlagsOk returns a tuple with the FeatureFlags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientConfig) GetFeatureFlagsOk() (*map[string]bool, bool) {
	if o == nil || IsNil(o.FeatureFlags) {
		return nil, false
	}
	return o.FeatureFlags, true
}

// HasFeatureFlags returns a boolean if a field has been set.
func (o *ClientConfig) HasFeatureFlags() bool {
	if o != nil && !IsNil(o.FeatureFlags) {
		return true
	}

	return false
}

// SetFeatureFlags gets a reference to the given map[string]bool and assigns it to the FeatureFlags field.
func (o *ClientConfig) SetFeatureFlags(v map[string]bool) {
	o.FeatureFlags = &v
}

// GetFeedPeopleCelebrationsEnabled returns the FeedPeopleCelebrationsEnabled field value if set, zero value otherwise.
func (o *ClientConfig) GetFeedPeopleCelebrationsEnabled() bool {
	if o == nil || IsNil(o.FeedPeopleCelebrationsEnabled) {
		var ret bool
		return ret
	}
	return *o.FeedPeopleCelebrationsEnabled
}

// GetFeedPeopleCelebrationsEnabledOk returns a tuple with the FeedPeopleCelebrationsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientConfig) GetFeedPeopleCelebrationsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.FeedPeopleCelebrationsEnabled) {
		return nil, false
	}
	return o.FeedPeopleCelebrationsEnabled, true
}

// HasFeedPeopleCelebrationsEnabled returns a boolean if a field has been set.
func (o *ClientConfig) HasFeedPeopleCelebrationsEnabled() bool {
	if o != nil && !IsNil(o.FeedPeopleCelebrationsEnabled) {
		return true
	}

	return false
}

// SetFeedPeopleCelebrationsEnabled gets a reference to the given bool and assigns it to the FeedPeopleCelebrationsEnabled field.
func (o *ClientConfig) SetFeedPeopleCelebrationsEnabled(v bool) {
	o.FeedPeopleCelebrationsEnabled = &v
}

// GetFeedSuggestedEnabled returns the FeedSuggestedEnabled field value if set, zero value otherwise.
func (o *ClientConfig) GetFeedSuggestedEnabled() bool {
	if o == nil || IsNil(o.FeedSuggestedEnabled) {
		var ret bool
		return ret
	}
	return *o.FeedSuggestedEnabled
}

// GetFeedSuggestedEnabledOk returns a tuple with the FeedSuggestedEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientConfig) GetFeedSuggestedEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.FeedSuggestedEnabled) {
		return nil, false
	}
	return o.FeedSuggestedEnabled, true
}

// HasFeedSuggestedEnabled returns a boolean if a field has been set.
func (o *ClientConfig) HasFeedSuggestedEnabled() bool {
	if o != nil && !IsNil(o.FeedSuggestedEnabled) {
		return true
	}

	return false
}

// SetFeedSuggestedEnabled gets a reference to the given bool and assigns it to the FeedSuggestedEnabled field.
func (o *ClientConfig) SetFeedSuggestedEnabled(v bool) {
	o.FeedSuggestedEnabled = &v
}

// GetFeedTrendingEnabled returns the FeedTrendingEnabled field value if set, zero value otherwise.
func (o *ClientConfig) GetFeedTrendingEnabled() bool {
	if o == nil || IsNil(o.FeedTrendingEnabled) {
		var ret bool
		return ret
	}
	return *o.FeedTrendingEnabled
}

// GetFeedTrendingEnabledOk returns a tuple with the FeedTrendingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientConfig) GetFeedTrendingEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.FeedTrendingEnabled) {
		return nil, false
	}
	return o.FeedTrendingEnabled, true
}

// HasFeedTrendingEnabled returns a boolean if a field has been set.
func (o *ClientConfig) HasFeedTrendingEnabled() bool {
	if o != nil && !IsNil(o.FeedTrendingEnabled) {
		return true
	}

	return false
}

// SetFeedTrendingEnabled gets a reference to the given bool and assigns it to the FeedTrendingEnabled field.
func (o *ClientConfig) SetFeedTrendingEnabled(v bool) {
	o.FeedTrendingEnabled = &v
}

// GetFeedRecentsEnabled returns the FeedRecentsEnabled field value if set, zero value otherwise.
func (o *ClientConfig) GetFeedRecentsEnabled() bool {
	if o == nil || IsNil(o.FeedRecentsEnabled) {
		var ret bool
		return ret
	}
	return *o.FeedRecentsEnabled
}

// GetFeedRecentsEnabledOk returns a tuple with the FeedRecentsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientConfig) GetFeedRecentsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.FeedRecentsEnabled) {
		return nil, false
	}
	return o.FeedRecentsEnabled, true
}

// HasFeedRecentsEnabled returns a boolean if a field has been set.
func (o *ClientConfig) HasFeedRecentsEnabled() bool {
	if o != nil && !IsNil(o.FeedRecentsEnabled) {
		return true
	}

	return false
}

// SetFeedRecentsEnabled gets a reference to the given bool and assigns it to the FeedRecentsEnabled field.
func (o *ClientConfig) SetFeedRecentsEnabled(v bool) {
	o.FeedRecentsEnabled = &v
}

// GetFeedMentionsEnabled returns the FeedMentionsEnabled field value if set, zero value otherwise.
func (o *ClientConfig) GetFeedMentionsEnabled() bool {
	if o == nil || IsNil(o.FeedMentionsEnabled) {
		var ret bool
		return ret
	}
	return *o.FeedMentionsEnabled
}

// GetFeedMentionsEnabledOk returns a tuple with the FeedMentionsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientConfig) GetFeedMentionsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.FeedMentionsEnabled) {
		return nil, false
	}
	return o.FeedMentionsEnabled, true
}

// HasFeedMentionsEnabled returns a boolean if a field has been set.
func (o *ClientConfig) HasFeedMentionsEnabled() bool {
	if o != nil && !IsNil(o.FeedMentionsEnabled) {
		return true
	}

	return false
}

// SetFeedMentionsEnabled gets a reference to the given bool and assigns it to the FeedMentionsEnabled field.
func (o *ClientConfig) SetFeedMentionsEnabled(v bool) {
	o.FeedMentionsEnabled = &v
}

// GetBoolValues returns the BoolValues field value if set, zero value otherwise.
func (o *ClientConfig) GetBoolValues() map[string]bool {
	if o == nil || IsNil(o.BoolValues) {
		var ret map[string]bool
		return ret
	}
	return *o.BoolValues
}

// GetBoolValuesOk returns a tuple with the BoolValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientConfig) GetBoolValuesOk() (*map[string]bool, bool) {
	if o == nil || IsNil(o.BoolValues) {
		return nil, false
	}
	return o.BoolValues, true
}

// HasBoolValues returns a boolean if a field has been set.
func (o *ClientConfig) HasBoolValues() bool {
	if o != nil && !IsNil(o.BoolValues) {
		return true
	}

	return false
}

// SetBoolValues gets a reference to the given map[string]bool and assigns it to the BoolValues field.
func (o *ClientConfig) SetBoolValues(v map[string]bool) {
	o.BoolValues = &v
}

// GetCompanyDisplayName returns the CompanyDisplayName field value if set, zero value otherwise.
func (o *ClientConfig) GetCompanyDisplayName() string {
	if o == nil || IsNil(o.CompanyDisplayName) {
		var ret string
		return ret
	}
	return *o.CompanyDisplayName
}

// GetCompanyDisplayNameOk returns a tuple with the CompanyDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientConfig) GetCompanyDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.CompanyDisplayName) {
		return nil, false
	}
	return o.CompanyDisplayName, true
}

// HasCompanyDisplayName returns a boolean if a field has been set.
func (o *ClientConfig) HasCompanyDisplayName() bool {
	if o != nil && !IsNil(o.CompanyDisplayName) {
		return true
	}

	return false
}

// SetCompanyDisplayName gets a reference to the given string and assigns it to the CompanyDisplayName field.
func (o *ClientConfig) SetCompanyDisplayName(v string) {
	o.CompanyDisplayName = &v
}

// GetCustomSerpMarkdown returns the CustomSerpMarkdown field value if set, zero value otherwise.
func (o *ClientConfig) GetCustomSerpMarkdown() string {
	if o == nil || IsNil(o.CustomSerpMarkdown) {
		var ret string
		return ret
	}
	return *o.CustomSerpMarkdown
}

// GetCustomSerpMarkdownOk returns a tuple with the CustomSerpMarkdown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientConfig) GetCustomSerpMarkdownOk() (*string, bool) {
	if o == nil || IsNil(o.CustomSerpMarkdown) {
		return nil, false
	}
	return o.CustomSerpMarkdown, true
}

// HasCustomSerpMarkdown returns a boolean if a field has been set.
func (o *ClientConfig) HasCustomSerpMarkdown() bool {
	if o != nil && !IsNil(o.CustomSerpMarkdown) {
		return true
	}

	return false
}

// SetCustomSerpMarkdown gets a reference to the given string and assigns it to the CustomSerpMarkdown field.
func (o *ClientConfig) SetCustomSerpMarkdown(v string) {
	o.CustomSerpMarkdown = &v
}

// GetFeedbackFrameSrc returns the FeedbackFrameSrc field value if set, zero value otherwise.
func (o *ClientConfig) GetFeedbackFrameSrc() string {
	if o == nil || IsNil(o.FeedbackFrameSrc) {
		var ret string
		return ret
	}
	return *o.FeedbackFrameSrc
}

// GetFeedbackFrameSrcOk returns a tuple with the FeedbackFrameSrc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientConfig) GetFeedbackFrameSrcOk() (*string, bool) {
	if o == nil || IsNil(o.FeedbackFrameSrc) {
		return nil, false
	}
	return o.FeedbackFrameSrc, true
}

// HasFeedbackFrameSrc returns a boolean if a field has been set.
func (o *ClientConfig) HasFeedbackFrameSrc() bool {
	if o != nil && !IsNil(o.FeedbackFrameSrc) {
		return true
	}

	return false
}

// SetFeedbackFrameSrc gets a reference to the given string and assigns it to the FeedbackFrameSrc field.
func (o *ClientConfig) SetFeedbackFrameSrc(v string) {
	o.FeedbackFrameSrc = &v
}

// GetOnboardingQuery returns the OnboardingQuery field value if set, zero value otherwise.
func (o *ClientConfig) GetOnboardingQuery() string {
	if o == nil || IsNil(o.OnboardingQuery) {
		var ret string
		return ret
	}
	return *o.OnboardingQuery
}

// GetOnboardingQueryOk returns a tuple with the OnboardingQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientConfig) GetOnboardingQueryOk() (*string, bool) {
	if o == nil || IsNil(o.OnboardingQuery) {
		return nil, false
	}
	return o.OnboardingQuery, true
}

// HasOnboardingQuery returns a boolean if a field has been set.
func (o *ClientConfig) HasOnboardingQuery() bool {
	if o != nil && !IsNil(o.OnboardingQuery) {
		return true
	}

	return false
}

// SetOnboardingQuery gets a reference to the given string and assigns it to the OnboardingQuery field.
func (o *ClientConfig) SetOnboardingQuery(v string) {
	o.OnboardingQuery = &v
}

// GetIsOrgChartLinkVisible returns the IsOrgChartLinkVisible field value if set, zero value otherwise.
func (o *ClientConfig) GetIsOrgChartLinkVisible() bool {
	if o == nil || IsNil(o.IsOrgChartLinkVisible) {
		var ret bool
		return ret
	}
	return *o.IsOrgChartLinkVisible
}

// GetIsOrgChartLinkVisibleOk returns a tuple with the IsOrgChartLinkVisible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientConfig) GetIsOrgChartLinkVisibleOk() (*bool, bool) {
	if o == nil || IsNil(o.IsOrgChartLinkVisible) {
		return nil, false
	}
	return o.IsOrgChartLinkVisible, true
}

// HasIsOrgChartLinkVisible returns a boolean if a field has been set.
func (o *ClientConfig) HasIsOrgChartLinkVisible() bool {
	if o != nil && !IsNil(o.IsOrgChartLinkVisible) {
		return true
	}

	return false
}

// SetIsOrgChartLinkVisible gets a reference to the given bool and assigns it to the IsOrgChartLinkVisible field.
func (o *ClientConfig) SetIsOrgChartLinkVisible(v bool) {
	o.IsOrgChartLinkVisible = &v
}

// GetIsOrgChartAccessible returns the IsOrgChartAccessible field value if set, zero value otherwise.
func (o *ClientConfig) GetIsOrgChartAccessible() bool {
	if o == nil || IsNil(o.IsOrgChartAccessible) {
		var ret bool
		return ret
	}
	return *o.IsOrgChartAccessible
}

// GetIsOrgChartAccessibleOk returns a tuple with the IsOrgChartAccessible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientConfig) GetIsOrgChartAccessibleOk() (*bool, bool) {
	if o == nil || IsNil(o.IsOrgChartAccessible) {
		return nil, false
	}
	return o.IsOrgChartAccessible, true
}

// HasIsOrgChartAccessible returns a boolean if a field has been set.
func (o *ClientConfig) HasIsOrgChartAccessible() bool {
	if o != nil && !IsNil(o.IsOrgChartAccessible) {
		return true
	}

	return false
}

// SetIsOrgChartAccessible gets a reference to the given bool and assigns it to the IsOrgChartAccessible field.
func (o *ClientConfig) SetIsOrgChartAccessible(v bool) {
	o.IsOrgChartAccessible = &v
}

// GetIsPeopleSetup returns the IsPeopleSetup field value if set, zero value otherwise.
func (o *ClientConfig) GetIsPeopleSetup() bool {
	if o == nil || IsNil(o.IsPeopleSetup) {
		var ret bool
		return ret
	}
	return *o.IsPeopleSetup
}

// GetIsPeopleSetupOk returns a tuple with the IsPeopleSetup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientConfig) GetIsPeopleSetupOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPeopleSetup) {
		return nil, false
	}
	return o.IsPeopleSetup, true
}

// HasIsPeopleSetup returns a boolean if a field has been set.
func (o *ClientConfig) HasIsPeopleSetup() bool {
	if o != nil && !IsNil(o.IsPeopleSetup) {
		return true
	}

	return false
}

// SetIsPeopleSetup gets a reference to the given bool and assigns it to the IsPeopleSetup field.
func (o *ClientConfig) SetIsPeopleSetup(v bool) {
	o.IsPeopleSetup = &v
}

// GetWebAppUrl returns the WebAppUrl field value if set, zero value otherwise.
func (o *ClientConfig) GetWebAppUrl() string {
	if o == nil || IsNil(o.WebAppUrl) {
		var ret string
		return ret
	}
	return *o.WebAppUrl
}

// GetWebAppUrlOk returns a tuple with the WebAppUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientConfig) GetWebAppUrlOk() (*string, bool) {
	if o == nil || IsNil(o.WebAppUrl) {
		return nil, false
	}
	return o.WebAppUrl, true
}

// HasWebAppUrl returns a boolean if a field has been set.
func (o *ClientConfig) HasWebAppUrl() bool {
	if o != nil && !IsNil(o.WebAppUrl) {
		return true
	}

	return false
}

// SetWebAppUrl gets a reference to the given string and assigns it to the WebAppUrl field.
func (o *ClientConfig) SetWebAppUrl(v string) {
	o.WebAppUrl = &v
}

// GetGaTrackingIds returns the GaTrackingIds field value if set, zero value otherwise.
func (o *ClientConfig) GetGaTrackingIds() []string {
	if o == nil || IsNil(o.GaTrackingIds) {
		var ret []string
		return ret
	}
	return o.GaTrackingIds
}

// GetGaTrackingIdsOk returns a tuple with the GaTrackingIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientConfig) GetGaTrackingIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.GaTrackingIds) {
		return nil, false
	}
	return o.GaTrackingIds, true
}

// HasGaTrackingIds returns a boolean if a field has been set.
func (o *ClientConfig) HasGaTrackingIds() bool {
	if o != nil && !IsNil(o.GaTrackingIds) {
		return true
	}

	return false
}

// SetGaTrackingIds gets a reference to the given []string and assigns it to the GaTrackingIds field.
func (o *ClientConfig) SetGaTrackingIds(v []string) {
	o.GaTrackingIds = v
}

// GetThemes returns the Themes field value if set, zero value otherwise.
func (o *ClientConfig) GetThemes() Themes {
	if o == nil || IsNil(o.Themes) {
		var ret Themes
		return ret
	}
	return *o.Themes
}

// GetThemesOk returns a tuple with the Themes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientConfig) GetThemesOk() (*Themes, bool) {
	if o == nil || IsNil(o.Themes) {
		return nil, false
	}
	return o.Themes, true
}

// HasThemes returns a boolean if a field has been set.
func (o *ClientConfig) HasThemes() bool {
	if o != nil && !IsNil(o.Themes) {
		return true
	}

	return false
}

// SetThemes gets a reference to the given Themes and assigns it to the Themes field.
func (o *ClientConfig) SetThemes(v Themes) {
	o.Themes = &v
}

// GetBrandings returns the Brandings field value if set, zero value otherwise.
func (o *ClientConfig) GetBrandings() ClientConfigBrandings {
	if o == nil || IsNil(o.Brandings) {
		var ret ClientConfigBrandings
		return ret
	}
	return *o.Brandings
}

// GetBrandingsOk returns a tuple with the Brandings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientConfig) GetBrandingsOk() (*ClientConfigBrandings, bool) {
	if o == nil || IsNil(o.Brandings) {
		return nil, false
	}
	return o.Brandings, true
}

// HasBrandings returns a boolean if a field has been set.
func (o *ClientConfig) HasBrandings() bool {
	if o != nil && !IsNil(o.Brandings) {
		return true
	}

	return false
}

// SetBrandings gets a reference to the given ClientConfigBrandings and assigns it to the Brandings field.
func (o *ClientConfig) SetBrandings(v ClientConfigBrandings) {
	o.Brandings = &v
}

// GetGreetingFormat returns the GreetingFormat field value if set, zero value otherwise.
func (o *ClientConfig) GetGreetingFormat() string {
	if o == nil || IsNil(o.GreetingFormat) {
		var ret string
		return ret
	}
	return *o.GreetingFormat
}

// GetGreetingFormatOk returns a tuple with the GreetingFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientConfig) GetGreetingFormatOk() (*string, bool) {
	if o == nil || IsNil(o.GreetingFormat) {
		return nil, false
	}
	return o.GreetingFormat, true
}

// HasGreetingFormat returns a boolean if a field has been set.
func (o *ClientConfig) HasGreetingFormat() bool {
	if o != nil && !IsNil(o.GreetingFormat) {
		return true
	}

	return false
}

// SetGreetingFormat gets a reference to the given string and assigns it to the GreetingFormat field.
func (o *ClientConfig) SetGreetingFormat(v string) {
	o.GreetingFormat = &v
}

// GetSsoCompanyProvider returns the SsoCompanyProvider field value if set, zero value otherwise.
func (o *ClientConfig) GetSsoCompanyProvider() string {
	if o == nil || IsNil(o.SsoCompanyProvider) {
		var ret string
		return ret
	}
	return *o.SsoCompanyProvider
}

// GetSsoCompanyProviderOk returns a tuple with the SsoCompanyProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientConfig) GetSsoCompanyProviderOk() (*string, bool) {
	if o == nil || IsNil(o.SsoCompanyProvider) {
		return nil, false
	}
	return o.SsoCompanyProvider, true
}

// HasSsoCompanyProvider returns a boolean if a field has been set.
func (o *ClientConfig) HasSsoCompanyProvider() bool {
	if o != nil && !IsNil(o.SsoCompanyProvider) {
		return true
	}

	return false
}

// SetSsoCompanyProvider gets a reference to the given string and assigns it to the SsoCompanyProvider field.
func (o *ClientConfig) SetSsoCompanyProvider(v string) {
	o.SsoCompanyProvider = &v
}

func (o ClientConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BadVersions) {
		toSerialize["badVersions"] = o.BadVersions
	}
	if !IsNil(o.FeatureFlags) {
		toSerialize["featureFlags"] = o.FeatureFlags
	}
	if !IsNil(o.FeedPeopleCelebrationsEnabled) {
		toSerialize["feedPeopleCelebrationsEnabled"] = o.FeedPeopleCelebrationsEnabled
	}
	if !IsNil(o.FeedSuggestedEnabled) {
		toSerialize["feedSuggestedEnabled"] = o.FeedSuggestedEnabled
	}
	if !IsNil(o.FeedTrendingEnabled) {
		toSerialize["feedTrendingEnabled"] = o.FeedTrendingEnabled
	}
	if !IsNil(o.FeedRecentsEnabled) {
		toSerialize["feedRecentsEnabled"] = o.FeedRecentsEnabled
	}
	if !IsNil(o.FeedMentionsEnabled) {
		toSerialize["feedMentionsEnabled"] = o.FeedMentionsEnabled
	}
	if !IsNil(o.BoolValues) {
		toSerialize["boolValues"] = o.BoolValues
	}
	if !IsNil(o.CompanyDisplayName) {
		toSerialize["companyDisplayName"] = o.CompanyDisplayName
	}
	if !IsNil(o.CustomSerpMarkdown) {
		toSerialize["customSerpMarkdown"] = o.CustomSerpMarkdown
	}
	if !IsNil(o.FeedbackFrameSrc) {
		toSerialize["feedbackFrameSrc"] = o.FeedbackFrameSrc
	}
	if !IsNil(o.OnboardingQuery) {
		toSerialize["onboardingQuery"] = o.OnboardingQuery
	}
	if !IsNil(o.IsOrgChartLinkVisible) {
		toSerialize["isOrgChartLinkVisible"] = o.IsOrgChartLinkVisible
	}
	if !IsNil(o.IsOrgChartAccessible) {
		toSerialize["isOrgChartAccessible"] = o.IsOrgChartAccessible
	}
	if !IsNil(o.IsPeopleSetup) {
		toSerialize["isPeopleSetup"] = o.IsPeopleSetup
	}
	if !IsNil(o.WebAppUrl) {
		toSerialize["webAppUrl"] = o.WebAppUrl
	}
	if !IsNil(o.GaTrackingIds) {
		toSerialize["gaTrackingIds"] = o.GaTrackingIds
	}
	if !IsNil(o.Themes) {
		toSerialize["themes"] = o.Themes
	}
	if !IsNil(o.Brandings) {
		toSerialize["brandings"] = o.Brandings
	}
	if !IsNil(o.GreetingFormat) {
		toSerialize["greetingFormat"] = o.GreetingFormat
	}
	if !IsNil(o.SsoCompanyProvider) {
		toSerialize["ssoCompanyProvider"] = o.SsoCompanyProvider
	}
	return toSerialize, nil
}

type NullableClientConfig struct {
	value *ClientConfig
	isSet bool
}

func (v NullableClientConfig) Get() *ClientConfig {
	return v.value
}

func (v *NullableClientConfig) Set(val *ClientConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableClientConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableClientConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientConfig(val *ClientConfig) *NullableClientConfig {
	return &NullableClientConfig{value: val, isSet: true}
}

func (v NullableClientConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


