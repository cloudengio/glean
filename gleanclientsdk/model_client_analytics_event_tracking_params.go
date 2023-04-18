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

// checks if the ClientAnalyticsEventTrackingParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClientAnalyticsEventTrackingParams{}

// ClientAnalyticsEventTrackingParams struct for ClientAnalyticsEventTrackingParams
type ClientAnalyticsEventTrackingParams struct {
	// Unix timestamp in millis for the client event.
	Timestamp *int64 `json:"timestamp,omitempty"`
	EventName *string `json:"eventName,omitempty"`
	Attribution *string `json:"attribution,omitempty"`
	Category *string `json:"category,omitempty"`
	Datasource *string `json:"datasource,omitempty"`
	DocType *string `json:"docType,omitempty"`
	Label *string `json:"label,omitempty"`
	PagePath *string `json:"pagePath,omitempty"`
	UiElement *string `json:"uiElement,omitempty"`
	UtmSource *string `json:"utmSource,omitempty"`
	// Sample rate applicable for this event.
	Rate *float32 `json:"rate,omitempty"`
	// Session identifier.
	Stt *string `json:"stt,omitempty"`
}

// NewClientAnalyticsEventTrackingParams instantiates a new ClientAnalyticsEventTrackingParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientAnalyticsEventTrackingParams() *ClientAnalyticsEventTrackingParams {
	this := ClientAnalyticsEventTrackingParams{}
	return &this
}

// NewClientAnalyticsEventTrackingParamsWithDefaults instantiates a new ClientAnalyticsEventTrackingParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientAnalyticsEventTrackingParamsWithDefaults() *ClientAnalyticsEventTrackingParams {
	this := ClientAnalyticsEventTrackingParams{}
	return &this
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *ClientAnalyticsEventTrackingParams) GetTimestamp() int64 {
	if o == nil || IsNil(o.Timestamp) {
		var ret int64
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientAnalyticsEventTrackingParams) GetTimestampOk() (*int64, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *ClientAnalyticsEventTrackingParams) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int64 and assigns it to the Timestamp field.
func (o *ClientAnalyticsEventTrackingParams) SetTimestamp(v int64) {
	o.Timestamp = &v
}

// GetEventName returns the EventName field value if set, zero value otherwise.
func (o *ClientAnalyticsEventTrackingParams) GetEventName() string {
	if o == nil || IsNil(o.EventName) {
		var ret string
		return ret
	}
	return *o.EventName
}

// GetEventNameOk returns a tuple with the EventName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientAnalyticsEventTrackingParams) GetEventNameOk() (*string, bool) {
	if o == nil || IsNil(o.EventName) {
		return nil, false
	}
	return o.EventName, true
}

// HasEventName returns a boolean if a field has been set.
func (o *ClientAnalyticsEventTrackingParams) HasEventName() bool {
	if o != nil && !IsNil(o.EventName) {
		return true
	}

	return false
}

// SetEventName gets a reference to the given string and assigns it to the EventName field.
func (o *ClientAnalyticsEventTrackingParams) SetEventName(v string) {
	o.EventName = &v
}

// GetAttribution returns the Attribution field value if set, zero value otherwise.
func (o *ClientAnalyticsEventTrackingParams) GetAttribution() string {
	if o == nil || IsNil(o.Attribution) {
		var ret string
		return ret
	}
	return *o.Attribution
}

// GetAttributionOk returns a tuple with the Attribution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientAnalyticsEventTrackingParams) GetAttributionOk() (*string, bool) {
	if o == nil || IsNil(o.Attribution) {
		return nil, false
	}
	return o.Attribution, true
}

// HasAttribution returns a boolean if a field has been set.
func (o *ClientAnalyticsEventTrackingParams) HasAttribution() bool {
	if o != nil && !IsNil(o.Attribution) {
		return true
	}

	return false
}

// SetAttribution gets a reference to the given string and assigns it to the Attribution field.
func (o *ClientAnalyticsEventTrackingParams) SetAttribution(v string) {
	o.Attribution = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *ClientAnalyticsEventTrackingParams) GetCategory() string {
	if o == nil || IsNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientAnalyticsEventTrackingParams) GetCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *ClientAnalyticsEventTrackingParams) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *ClientAnalyticsEventTrackingParams) SetCategory(v string) {
	o.Category = &v
}

// GetDatasource returns the Datasource field value if set, zero value otherwise.
func (o *ClientAnalyticsEventTrackingParams) GetDatasource() string {
	if o == nil || IsNil(o.Datasource) {
		var ret string
		return ret
	}
	return *o.Datasource
}

// GetDatasourceOk returns a tuple with the Datasource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientAnalyticsEventTrackingParams) GetDatasourceOk() (*string, bool) {
	if o == nil || IsNil(o.Datasource) {
		return nil, false
	}
	return o.Datasource, true
}

// HasDatasource returns a boolean if a field has been set.
func (o *ClientAnalyticsEventTrackingParams) HasDatasource() bool {
	if o != nil && !IsNil(o.Datasource) {
		return true
	}

	return false
}

// SetDatasource gets a reference to the given string and assigns it to the Datasource field.
func (o *ClientAnalyticsEventTrackingParams) SetDatasource(v string) {
	o.Datasource = &v
}

// GetDocType returns the DocType field value if set, zero value otherwise.
func (o *ClientAnalyticsEventTrackingParams) GetDocType() string {
	if o == nil || IsNil(o.DocType) {
		var ret string
		return ret
	}
	return *o.DocType
}

// GetDocTypeOk returns a tuple with the DocType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientAnalyticsEventTrackingParams) GetDocTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DocType) {
		return nil, false
	}
	return o.DocType, true
}

// HasDocType returns a boolean if a field has been set.
func (o *ClientAnalyticsEventTrackingParams) HasDocType() bool {
	if o != nil && !IsNil(o.DocType) {
		return true
	}

	return false
}

// SetDocType gets a reference to the given string and assigns it to the DocType field.
func (o *ClientAnalyticsEventTrackingParams) SetDocType(v string) {
	o.DocType = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *ClientAnalyticsEventTrackingParams) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientAnalyticsEventTrackingParams) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *ClientAnalyticsEventTrackingParams) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *ClientAnalyticsEventTrackingParams) SetLabel(v string) {
	o.Label = &v
}

// GetPagePath returns the PagePath field value if set, zero value otherwise.
func (o *ClientAnalyticsEventTrackingParams) GetPagePath() string {
	if o == nil || IsNil(o.PagePath) {
		var ret string
		return ret
	}
	return *o.PagePath
}

// GetPagePathOk returns a tuple with the PagePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientAnalyticsEventTrackingParams) GetPagePathOk() (*string, bool) {
	if o == nil || IsNil(o.PagePath) {
		return nil, false
	}
	return o.PagePath, true
}

// HasPagePath returns a boolean if a field has been set.
func (o *ClientAnalyticsEventTrackingParams) HasPagePath() bool {
	if o != nil && !IsNil(o.PagePath) {
		return true
	}

	return false
}

// SetPagePath gets a reference to the given string and assigns it to the PagePath field.
func (o *ClientAnalyticsEventTrackingParams) SetPagePath(v string) {
	o.PagePath = &v
}

// GetUiElement returns the UiElement field value if set, zero value otherwise.
func (o *ClientAnalyticsEventTrackingParams) GetUiElement() string {
	if o == nil || IsNil(o.UiElement) {
		var ret string
		return ret
	}
	return *o.UiElement
}

// GetUiElementOk returns a tuple with the UiElement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientAnalyticsEventTrackingParams) GetUiElementOk() (*string, bool) {
	if o == nil || IsNil(o.UiElement) {
		return nil, false
	}
	return o.UiElement, true
}

// HasUiElement returns a boolean if a field has been set.
func (o *ClientAnalyticsEventTrackingParams) HasUiElement() bool {
	if o != nil && !IsNil(o.UiElement) {
		return true
	}

	return false
}

// SetUiElement gets a reference to the given string and assigns it to the UiElement field.
func (o *ClientAnalyticsEventTrackingParams) SetUiElement(v string) {
	o.UiElement = &v
}

// GetUtmSource returns the UtmSource field value if set, zero value otherwise.
func (o *ClientAnalyticsEventTrackingParams) GetUtmSource() string {
	if o == nil || IsNil(o.UtmSource) {
		var ret string
		return ret
	}
	return *o.UtmSource
}

// GetUtmSourceOk returns a tuple with the UtmSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientAnalyticsEventTrackingParams) GetUtmSourceOk() (*string, bool) {
	if o == nil || IsNil(o.UtmSource) {
		return nil, false
	}
	return o.UtmSource, true
}

// HasUtmSource returns a boolean if a field has been set.
func (o *ClientAnalyticsEventTrackingParams) HasUtmSource() bool {
	if o != nil && !IsNil(o.UtmSource) {
		return true
	}

	return false
}

// SetUtmSource gets a reference to the given string and assigns it to the UtmSource field.
func (o *ClientAnalyticsEventTrackingParams) SetUtmSource(v string) {
	o.UtmSource = &v
}

// GetRate returns the Rate field value if set, zero value otherwise.
func (o *ClientAnalyticsEventTrackingParams) GetRate() float32 {
	if o == nil || IsNil(o.Rate) {
		var ret float32
		return ret
	}
	return *o.Rate
}

// GetRateOk returns a tuple with the Rate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientAnalyticsEventTrackingParams) GetRateOk() (*float32, bool) {
	if o == nil || IsNil(o.Rate) {
		return nil, false
	}
	return o.Rate, true
}

// HasRate returns a boolean if a field has been set.
func (o *ClientAnalyticsEventTrackingParams) HasRate() bool {
	if o != nil && !IsNil(o.Rate) {
		return true
	}

	return false
}

// SetRate gets a reference to the given float32 and assigns it to the Rate field.
func (o *ClientAnalyticsEventTrackingParams) SetRate(v float32) {
	o.Rate = &v
}

// GetStt returns the Stt field value if set, zero value otherwise.
func (o *ClientAnalyticsEventTrackingParams) GetStt() string {
	if o == nil || IsNil(o.Stt) {
		var ret string
		return ret
	}
	return *o.Stt
}

// GetSttOk returns a tuple with the Stt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientAnalyticsEventTrackingParams) GetSttOk() (*string, bool) {
	if o == nil || IsNil(o.Stt) {
		return nil, false
	}
	return o.Stt, true
}

// HasStt returns a boolean if a field has been set.
func (o *ClientAnalyticsEventTrackingParams) HasStt() bool {
	if o != nil && !IsNil(o.Stt) {
		return true
	}

	return false
}

// SetStt gets a reference to the given string and assigns it to the Stt field.
func (o *ClientAnalyticsEventTrackingParams) SetStt(v string) {
	o.Stt = &v
}

func (o ClientAnalyticsEventTrackingParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientAnalyticsEventTrackingParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !IsNil(o.EventName) {
		toSerialize["eventName"] = o.EventName
	}
	if !IsNil(o.Attribution) {
		toSerialize["attribution"] = o.Attribution
	}
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !IsNil(o.Datasource) {
		toSerialize["datasource"] = o.Datasource
	}
	if !IsNil(o.DocType) {
		toSerialize["docType"] = o.DocType
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.PagePath) {
		toSerialize["pagePath"] = o.PagePath
	}
	if !IsNil(o.UiElement) {
		toSerialize["uiElement"] = o.UiElement
	}
	if !IsNil(o.UtmSource) {
		toSerialize["utmSource"] = o.UtmSource
	}
	if !IsNil(o.Rate) {
		toSerialize["rate"] = o.Rate
	}
	if !IsNil(o.Stt) {
		toSerialize["stt"] = o.Stt
	}
	return toSerialize, nil
}

type NullableClientAnalyticsEventTrackingParams struct {
	value *ClientAnalyticsEventTrackingParams
	isSet bool
}

func (v NullableClientAnalyticsEventTrackingParams) Get() *ClientAnalyticsEventTrackingParams {
	return v.value
}

func (v *NullableClientAnalyticsEventTrackingParams) Set(val *ClientAnalyticsEventTrackingParams) {
	v.value = val
	v.isSet = true
}

func (v NullableClientAnalyticsEventTrackingParams) IsSet() bool {
	return v.isSet
}

func (v *NullableClientAnalyticsEventTrackingParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientAnalyticsEventTrackingParams(val *ClientAnalyticsEventTrackingParams) *NullableClientAnalyticsEventTrackingParams {
	return &NullableClientAnalyticsEventTrackingParams{value: val, isSet: true}
}

func (v NullableClientAnalyticsEventTrackingParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientAnalyticsEventTrackingParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

