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

// checks if the ActivityEventParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActivityEventParams{}

// ActivityEventParams struct for ActivityEventParams
type ActivityEventParams struct {
	// The HTML content of the page body.
	BodyContent *string `json:"bodyContent,omitempty"`
	// The full datasource instance name inferred from the URL of the event
	DatasourceInstance *string `json:"datasourceInstance,omitempty"`
	// The datasource without the instance inferred from the URL of the event
	Datasource *string `json:"datasource,omitempty"`
	// The instance only name of the datasource instance, e.g. 1 for jira_1, inferred from the URL of the event
	InstanceOnlyName *string `json:"instanceOnlyName,omitempty"`
	// Length in seconds of the activity. For VIEWS, this represents the amount the page was visible in the foreground.
	Duration *int32 `json:"duration,omitempty"`
	// The user's search query associated with a SEARCH.
	Query *string `json:"query,omitempty"`
	// The referring URL of the VIEW or SEARCH.
	Referrer *string `json:"referrer,omitempty"`
	// The page title associated with the URL of the event
	Title *string `json:"title,omitempty"`
	// Indicates that the params are incomplete and more params may be sent with the same action+timestamp+url in the future. This is used for sending the duration when a VIEW is finished.
	Truncated *bool `json:"truncated,omitempty"`
}

// NewActivityEventParams instantiates a new ActivityEventParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActivityEventParams() *ActivityEventParams {
	this := ActivityEventParams{}
	return &this
}

// NewActivityEventParamsWithDefaults instantiates a new ActivityEventParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActivityEventParamsWithDefaults() *ActivityEventParams {
	this := ActivityEventParams{}
	return &this
}

// GetBodyContent returns the BodyContent field value if set, zero value otherwise.
func (o *ActivityEventParams) GetBodyContent() string {
	if o == nil || IsNil(o.BodyContent) {
		var ret string
		return ret
	}
	return *o.BodyContent
}

// GetBodyContentOk returns a tuple with the BodyContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityEventParams) GetBodyContentOk() (*string, bool) {
	if o == nil || IsNil(o.BodyContent) {
		return nil, false
	}
	return o.BodyContent, true
}

// HasBodyContent returns a boolean if a field has been set.
func (o *ActivityEventParams) HasBodyContent() bool {
	if o != nil && !IsNil(o.BodyContent) {
		return true
	}

	return false
}

// SetBodyContent gets a reference to the given string and assigns it to the BodyContent field.
func (o *ActivityEventParams) SetBodyContent(v string) {
	o.BodyContent = &v
}

// GetDatasourceInstance returns the DatasourceInstance field value if set, zero value otherwise.
func (o *ActivityEventParams) GetDatasourceInstance() string {
	if o == nil || IsNil(o.DatasourceInstance) {
		var ret string
		return ret
	}
	return *o.DatasourceInstance
}

// GetDatasourceInstanceOk returns a tuple with the DatasourceInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityEventParams) GetDatasourceInstanceOk() (*string, bool) {
	if o == nil || IsNil(o.DatasourceInstance) {
		return nil, false
	}
	return o.DatasourceInstance, true
}

// HasDatasourceInstance returns a boolean if a field has been set.
func (o *ActivityEventParams) HasDatasourceInstance() bool {
	if o != nil && !IsNil(o.DatasourceInstance) {
		return true
	}

	return false
}

// SetDatasourceInstance gets a reference to the given string and assigns it to the DatasourceInstance field.
func (o *ActivityEventParams) SetDatasourceInstance(v string) {
	o.DatasourceInstance = &v
}

// GetDatasource returns the Datasource field value if set, zero value otherwise.
func (o *ActivityEventParams) GetDatasource() string {
	if o == nil || IsNil(o.Datasource) {
		var ret string
		return ret
	}
	return *o.Datasource
}

// GetDatasourceOk returns a tuple with the Datasource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityEventParams) GetDatasourceOk() (*string, bool) {
	if o == nil || IsNil(o.Datasource) {
		return nil, false
	}
	return o.Datasource, true
}

// HasDatasource returns a boolean if a field has been set.
func (o *ActivityEventParams) HasDatasource() bool {
	if o != nil && !IsNil(o.Datasource) {
		return true
	}

	return false
}

// SetDatasource gets a reference to the given string and assigns it to the Datasource field.
func (o *ActivityEventParams) SetDatasource(v string) {
	o.Datasource = &v
}

// GetInstanceOnlyName returns the InstanceOnlyName field value if set, zero value otherwise.
func (o *ActivityEventParams) GetInstanceOnlyName() string {
	if o == nil || IsNil(o.InstanceOnlyName) {
		var ret string
		return ret
	}
	return *o.InstanceOnlyName
}

// GetInstanceOnlyNameOk returns a tuple with the InstanceOnlyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityEventParams) GetInstanceOnlyNameOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceOnlyName) {
		return nil, false
	}
	return o.InstanceOnlyName, true
}

// HasInstanceOnlyName returns a boolean if a field has been set.
func (o *ActivityEventParams) HasInstanceOnlyName() bool {
	if o != nil && !IsNil(o.InstanceOnlyName) {
		return true
	}

	return false
}

// SetInstanceOnlyName gets a reference to the given string and assigns it to the InstanceOnlyName field.
func (o *ActivityEventParams) SetInstanceOnlyName(v string) {
	o.InstanceOnlyName = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *ActivityEventParams) GetDuration() int32 {
	if o == nil || IsNil(o.Duration) {
		var ret int32
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityEventParams) GetDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *ActivityEventParams) HasDuration() bool {
	if o != nil && !IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int32 and assigns it to the Duration field.
func (o *ActivityEventParams) SetDuration(v int32) {
	o.Duration = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *ActivityEventParams) GetQuery() string {
	if o == nil || IsNil(o.Query) {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityEventParams) GetQueryOk() (*string, bool) {
	if o == nil || IsNil(o.Query) {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *ActivityEventParams) HasQuery() bool {
	if o != nil && !IsNil(o.Query) {
		return true
	}

	return false
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *ActivityEventParams) SetQuery(v string) {
	o.Query = &v
}

// GetReferrer returns the Referrer field value if set, zero value otherwise.
func (o *ActivityEventParams) GetReferrer() string {
	if o == nil || IsNil(o.Referrer) {
		var ret string
		return ret
	}
	return *o.Referrer
}

// GetReferrerOk returns a tuple with the Referrer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityEventParams) GetReferrerOk() (*string, bool) {
	if o == nil || IsNil(o.Referrer) {
		return nil, false
	}
	return o.Referrer, true
}

// HasReferrer returns a boolean if a field has been set.
func (o *ActivityEventParams) HasReferrer() bool {
	if o != nil && !IsNil(o.Referrer) {
		return true
	}

	return false
}

// SetReferrer gets a reference to the given string and assigns it to the Referrer field.
func (o *ActivityEventParams) SetReferrer(v string) {
	o.Referrer = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ActivityEventParams) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityEventParams) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ActivityEventParams) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ActivityEventParams) SetTitle(v string) {
	o.Title = &v
}

// GetTruncated returns the Truncated field value if set, zero value otherwise.
func (o *ActivityEventParams) GetTruncated() bool {
	if o == nil || IsNil(o.Truncated) {
		var ret bool
		return ret
	}
	return *o.Truncated
}

// GetTruncatedOk returns a tuple with the Truncated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityEventParams) GetTruncatedOk() (*bool, bool) {
	if o == nil || IsNil(o.Truncated) {
		return nil, false
	}
	return o.Truncated, true
}

// HasTruncated returns a boolean if a field has been set.
func (o *ActivityEventParams) HasTruncated() bool {
	if o != nil && !IsNil(o.Truncated) {
		return true
	}

	return false
}

// SetTruncated gets a reference to the given bool and assigns it to the Truncated field.
func (o *ActivityEventParams) SetTruncated(v bool) {
	o.Truncated = &v
}

func (o ActivityEventParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActivityEventParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BodyContent) {
		toSerialize["bodyContent"] = o.BodyContent
	}
	if !IsNil(o.DatasourceInstance) {
		toSerialize["datasourceInstance"] = o.DatasourceInstance
	}
	if !IsNil(o.Datasource) {
		toSerialize["datasource"] = o.Datasource
	}
	if !IsNil(o.InstanceOnlyName) {
		toSerialize["instanceOnlyName"] = o.InstanceOnlyName
	}
	if !IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if !IsNil(o.Query) {
		toSerialize["query"] = o.Query
	}
	if !IsNil(o.Referrer) {
		toSerialize["referrer"] = o.Referrer
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Truncated) {
		toSerialize["truncated"] = o.Truncated
	}
	return toSerialize, nil
}

type NullableActivityEventParams struct {
	value *ActivityEventParams
	isSet bool
}

func (v NullableActivityEventParams) Get() *ActivityEventParams {
	return v.value
}

func (v *NullableActivityEventParams) Set(val *ActivityEventParams) {
	v.value = val
	v.isSet = true
}

func (v NullableActivityEventParams) IsSet() bool {
	return v.isSet
}

func (v *NullableActivityEventParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActivityEventParams(val *ActivityEventParams) *NullableActivityEventParams {
	return &NullableActivityEventParams{value: val, isSet: true}
}

func (v NullableActivityEventParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActivityEventParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


