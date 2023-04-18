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

// checks if the PrivateAuth type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrivateAuth{}

// PrivateAuth struct for PrivateAuth
type PrivateAuth struct {
	// The authorization status of the user.
	Status *string `json:"status,omitempty"`
	// Whether the datasource is authorized through domain delegation.
	IsDomainDelegated *bool `json:"isDomainDelegated,omitempty"`
	// The subtitle of the results auth prompt, required when the datasource has results even if the user has not authed privately.
	ResultsAuthPromptSubtitle *string `json:"resultsAuthPromptSubtitle,omitempty"`
	// The title of the auth prompt shown to the user
	Title *string `json:"title,omitempty"`
	// The content of the auth prompt shown to the user
	Content *string `json:"content,omitempty"`
	// FAQs to show when asking the user to authorize the datasource
	Questions [][]string `json:"questions,omitempty"`
	// The relative path of the url to take the user to to authorize the datasource
	AuthUrlRelativePath *string `json:"authUrlRelativePath,omitempty"`
}

// NewPrivateAuth instantiates a new PrivateAuth object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrivateAuth() *PrivateAuth {
	this := PrivateAuth{}
	return &this
}

// NewPrivateAuthWithDefaults instantiates a new PrivateAuth object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrivateAuthWithDefaults() *PrivateAuth {
	this := PrivateAuth{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PrivateAuth) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateAuth) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PrivateAuth) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *PrivateAuth) SetStatus(v string) {
	o.Status = &v
}

// GetIsDomainDelegated returns the IsDomainDelegated field value if set, zero value otherwise.
func (o *PrivateAuth) GetIsDomainDelegated() bool {
	if o == nil || IsNil(o.IsDomainDelegated) {
		var ret bool
		return ret
	}
	return *o.IsDomainDelegated
}

// GetIsDomainDelegatedOk returns a tuple with the IsDomainDelegated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateAuth) GetIsDomainDelegatedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDomainDelegated) {
		return nil, false
	}
	return o.IsDomainDelegated, true
}

// HasIsDomainDelegated returns a boolean if a field has been set.
func (o *PrivateAuth) HasIsDomainDelegated() bool {
	if o != nil && !IsNil(o.IsDomainDelegated) {
		return true
	}

	return false
}

// SetIsDomainDelegated gets a reference to the given bool and assigns it to the IsDomainDelegated field.
func (o *PrivateAuth) SetIsDomainDelegated(v bool) {
	o.IsDomainDelegated = &v
}

// GetResultsAuthPromptSubtitle returns the ResultsAuthPromptSubtitle field value if set, zero value otherwise.
func (o *PrivateAuth) GetResultsAuthPromptSubtitle() string {
	if o == nil || IsNil(o.ResultsAuthPromptSubtitle) {
		var ret string
		return ret
	}
	return *o.ResultsAuthPromptSubtitle
}

// GetResultsAuthPromptSubtitleOk returns a tuple with the ResultsAuthPromptSubtitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateAuth) GetResultsAuthPromptSubtitleOk() (*string, bool) {
	if o == nil || IsNil(o.ResultsAuthPromptSubtitle) {
		return nil, false
	}
	return o.ResultsAuthPromptSubtitle, true
}

// HasResultsAuthPromptSubtitle returns a boolean if a field has been set.
func (o *PrivateAuth) HasResultsAuthPromptSubtitle() bool {
	if o != nil && !IsNil(o.ResultsAuthPromptSubtitle) {
		return true
	}

	return false
}

// SetResultsAuthPromptSubtitle gets a reference to the given string and assigns it to the ResultsAuthPromptSubtitle field.
func (o *PrivateAuth) SetResultsAuthPromptSubtitle(v string) {
	o.ResultsAuthPromptSubtitle = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *PrivateAuth) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateAuth) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *PrivateAuth) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *PrivateAuth) SetTitle(v string) {
	o.Title = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *PrivateAuth) GetContent() string {
	if o == nil || IsNil(o.Content) {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateAuth) GetContentOk() (*string, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *PrivateAuth) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *PrivateAuth) SetContent(v string) {
	o.Content = &v
}

// GetQuestions returns the Questions field value if set, zero value otherwise.
func (o *PrivateAuth) GetQuestions() [][]string {
	if o == nil || IsNil(o.Questions) {
		var ret [][]string
		return ret
	}
	return o.Questions
}

// GetQuestionsOk returns a tuple with the Questions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateAuth) GetQuestionsOk() ([][]string, bool) {
	if o == nil || IsNil(o.Questions) {
		return nil, false
	}
	return o.Questions, true
}

// HasQuestions returns a boolean if a field has been set.
func (o *PrivateAuth) HasQuestions() bool {
	if o != nil && !IsNil(o.Questions) {
		return true
	}

	return false
}

// SetQuestions gets a reference to the given [][]string and assigns it to the Questions field.
func (o *PrivateAuth) SetQuestions(v [][]string) {
	o.Questions = v
}

// GetAuthUrlRelativePath returns the AuthUrlRelativePath field value if set, zero value otherwise.
func (o *PrivateAuth) GetAuthUrlRelativePath() string {
	if o == nil || IsNil(o.AuthUrlRelativePath) {
		var ret string
		return ret
	}
	return *o.AuthUrlRelativePath
}

// GetAuthUrlRelativePathOk returns a tuple with the AuthUrlRelativePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateAuth) GetAuthUrlRelativePathOk() (*string, bool) {
	if o == nil || IsNil(o.AuthUrlRelativePath) {
		return nil, false
	}
	return o.AuthUrlRelativePath, true
}

// HasAuthUrlRelativePath returns a boolean if a field has been set.
func (o *PrivateAuth) HasAuthUrlRelativePath() bool {
	if o != nil && !IsNil(o.AuthUrlRelativePath) {
		return true
	}

	return false
}

// SetAuthUrlRelativePath gets a reference to the given string and assigns it to the AuthUrlRelativePath field.
func (o *PrivateAuth) SetAuthUrlRelativePath(v string) {
	o.AuthUrlRelativePath = &v
}

func (o PrivateAuth) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrivateAuth) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.IsDomainDelegated) {
		toSerialize["isDomainDelegated"] = o.IsDomainDelegated
	}
	if !IsNil(o.ResultsAuthPromptSubtitle) {
		toSerialize["resultsAuthPromptSubtitle"] = o.ResultsAuthPromptSubtitle
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	if !IsNil(o.Questions) {
		toSerialize["questions"] = o.Questions
	}
	if !IsNil(o.AuthUrlRelativePath) {
		toSerialize["authUrlRelativePath"] = o.AuthUrlRelativePath
	}
	return toSerialize, nil
}

type NullablePrivateAuth struct {
	value *PrivateAuth
	isSet bool
}

func (v NullablePrivateAuth) Get() *PrivateAuth {
	return v.value
}

func (v *NullablePrivateAuth) Set(val *PrivateAuth) {
	v.value = val
	v.isSet = true
}

func (v NullablePrivateAuth) IsSet() bool {
	return v.isSet
}

func (v *NullablePrivateAuth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrivateAuth(val *PrivateAuth) *NullablePrivateAuth {
	return &NullablePrivateAuth{value: val, isSet: true}
}

func (v NullablePrivateAuth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrivateAuth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


