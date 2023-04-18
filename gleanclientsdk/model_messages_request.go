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

// checks if the MessagesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessagesRequest{}

// MessagesRequest struct for MessagesRequest
type MessagesRequest struct {
	// Type of the id in the incoming request.
	IdType string `json:"idType"`
	// ID corresponding to the requested idType. Note that channel and threads are represented by the underlying datasource's ID and conversations are represented by their document's ID.
	Id string `json:"id"`
	// Id for the for the workspace in case of multiple workspaces.
	WorkspaceId *string `json:"workspaceId,omitempty"`
	// The direction of the results asked with respect to the reference timestamp. Missing field defaults to OLDER.
	Direction *string `json:"direction,omitempty"`
	// Timestamp in millis of the reference message.
	TimestampMillis int64 `json:"timestampMillis"`
	// Whether to include root message in response.
	IncludeRootMessage *bool `json:"includeRootMessage,omitempty"`
	// The type of the data source. Missing field defaults to SLACK.
	Datasource *string `json:"datasource,omitempty"`
	// The datasource instance display name from which the document was extracted. This is used for appinstance facet filter for datasources that support multiple instances.
	DatasourceInstanceDisplayName *string `json:"datasourceInstanceDisplayName,omitempty"`
	// Debug only search params to to change the flow of control in request handling. It will be passed around service boundaries/endpoints. For more details, https://docs.google.com/document/d/1e6taTfWUL8KNUC9de8kmmG2MG2L6cTx4ulOJfAshDTM/edit. Requires sufficient permissions.
	Sc *string `json:"sc,omitempty"`
}

// NewMessagesRequest instantiates a new MessagesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessagesRequest(idType string, id string, timestampMillis int64) *MessagesRequest {
	this := MessagesRequest{}
	this.IdType = idType
	this.Id = id
	this.TimestampMillis = timestampMillis
	return &this
}

// NewMessagesRequestWithDefaults instantiates a new MessagesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessagesRequestWithDefaults() *MessagesRequest {
	this := MessagesRequest{}
	return &this
}

// GetIdType returns the IdType field value
func (o *MessagesRequest) GetIdType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdType
}

// GetIdTypeOk returns a tuple with the IdType field value
// and a boolean to check if the value has been set.
func (o *MessagesRequest) GetIdTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdType, true
}

// SetIdType sets field value
func (o *MessagesRequest) SetIdType(v string) {
	o.IdType = v
}

// GetId returns the Id field value
func (o *MessagesRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MessagesRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MessagesRequest) SetId(v string) {
	o.Id = v
}

// GetWorkspaceId returns the WorkspaceId field value if set, zero value otherwise.
func (o *MessagesRequest) GetWorkspaceId() string {
	if o == nil || IsNil(o.WorkspaceId) {
		var ret string
		return ret
	}
	return *o.WorkspaceId
}

// GetWorkspaceIdOk returns a tuple with the WorkspaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessagesRequest) GetWorkspaceIdOk() (*string, bool) {
	if o == nil || IsNil(o.WorkspaceId) {
		return nil, false
	}
	return o.WorkspaceId, true
}

// HasWorkspaceId returns a boolean if a field has been set.
func (o *MessagesRequest) HasWorkspaceId() bool {
	if o != nil && !IsNil(o.WorkspaceId) {
		return true
	}

	return false
}

// SetWorkspaceId gets a reference to the given string and assigns it to the WorkspaceId field.
func (o *MessagesRequest) SetWorkspaceId(v string) {
	o.WorkspaceId = &v
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *MessagesRequest) GetDirection() string {
	if o == nil || IsNil(o.Direction) {
		var ret string
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessagesRequest) GetDirectionOk() (*string, bool) {
	if o == nil || IsNil(o.Direction) {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *MessagesRequest) HasDirection() bool {
	if o != nil && !IsNil(o.Direction) {
		return true
	}

	return false
}

// SetDirection gets a reference to the given string and assigns it to the Direction field.
func (o *MessagesRequest) SetDirection(v string) {
	o.Direction = &v
}

// GetTimestampMillis returns the TimestampMillis field value
func (o *MessagesRequest) GetTimestampMillis() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TimestampMillis
}

// GetTimestampMillisOk returns a tuple with the TimestampMillis field value
// and a boolean to check if the value has been set.
func (o *MessagesRequest) GetTimestampMillisOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimestampMillis, true
}

// SetTimestampMillis sets field value
func (o *MessagesRequest) SetTimestampMillis(v int64) {
	o.TimestampMillis = v
}

// GetIncludeRootMessage returns the IncludeRootMessage field value if set, zero value otherwise.
func (o *MessagesRequest) GetIncludeRootMessage() bool {
	if o == nil || IsNil(o.IncludeRootMessage) {
		var ret bool
		return ret
	}
	return *o.IncludeRootMessage
}

// GetIncludeRootMessageOk returns a tuple with the IncludeRootMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessagesRequest) GetIncludeRootMessageOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeRootMessage) {
		return nil, false
	}
	return o.IncludeRootMessage, true
}

// HasIncludeRootMessage returns a boolean if a field has been set.
func (o *MessagesRequest) HasIncludeRootMessage() bool {
	if o != nil && !IsNil(o.IncludeRootMessage) {
		return true
	}

	return false
}

// SetIncludeRootMessage gets a reference to the given bool and assigns it to the IncludeRootMessage field.
func (o *MessagesRequest) SetIncludeRootMessage(v bool) {
	o.IncludeRootMessage = &v
}

// GetDatasource returns the Datasource field value if set, zero value otherwise.
func (o *MessagesRequest) GetDatasource() string {
	if o == nil || IsNil(o.Datasource) {
		var ret string
		return ret
	}
	return *o.Datasource
}

// GetDatasourceOk returns a tuple with the Datasource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessagesRequest) GetDatasourceOk() (*string, bool) {
	if o == nil || IsNil(o.Datasource) {
		return nil, false
	}
	return o.Datasource, true
}

// HasDatasource returns a boolean if a field has been set.
func (o *MessagesRequest) HasDatasource() bool {
	if o != nil && !IsNil(o.Datasource) {
		return true
	}

	return false
}

// SetDatasource gets a reference to the given string and assigns it to the Datasource field.
func (o *MessagesRequest) SetDatasource(v string) {
	o.Datasource = &v
}

// GetDatasourceInstanceDisplayName returns the DatasourceInstanceDisplayName field value if set, zero value otherwise.
func (o *MessagesRequest) GetDatasourceInstanceDisplayName() string {
	if o == nil || IsNil(o.DatasourceInstanceDisplayName) {
		var ret string
		return ret
	}
	return *o.DatasourceInstanceDisplayName
}

// GetDatasourceInstanceDisplayNameOk returns a tuple with the DatasourceInstanceDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessagesRequest) GetDatasourceInstanceDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DatasourceInstanceDisplayName) {
		return nil, false
	}
	return o.DatasourceInstanceDisplayName, true
}

// HasDatasourceInstanceDisplayName returns a boolean if a field has been set.
func (o *MessagesRequest) HasDatasourceInstanceDisplayName() bool {
	if o != nil && !IsNil(o.DatasourceInstanceDisplayName) {
		return true
	}

	return false
}

// SetDatasourceInstanceDisplayName gets a reference to the given string and assigns it to the DatasourceInstanceDisplayName field.
func (o *MessagesRequest) SetDatasourceInstanceDisplayName(v string) {
	o.DatasourceInstanceDisplayName = &v
}

// GetSc returns the Sc field value if set, zero value otherwise.
func (o *MessagesRequest) GetSc() string {
	if o == nil || IsNil(o.Sc) {
		var ret string
		return ret
	}
	return *o.Sc
}

// GetScOk returns a tuple with the Sc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessagesRequest) GetScOk() (*string, bool) {
	if o == nil || IsNil(o.Sc) {
		return nil, false
	}
	return o.Sc, true
}

// HasSc returns a boolean if a field has been set.
func (o *MessagesRequest) HasSc() bool {
	if o != nil && !IsNil(o.Sc) {
		return true
	}

	return false
}

// SetSc gets a reference to the given string and assigns it to the Sc field.
func (o *MessagesRequest) SetSc(v string) {
	o.Sc = &v
}

func (o MessagesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessagesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["idType"] = o.IdType
	toSerialize["id"] = o.Id
	if !IsNil(o.WorkspaceId) {
		toSerialize["workspaceId"] = o.WorkspaceId
	}
	if !IsNil(o.Direction) {
		toSerialize["direction"] = o.Direction
	}
	toSerialize["timestampMillis"] = o.TimestampMillis
	if !IsNil(o.IncludeRootMessage) {
		toSerialize["includeRootMessage"] = o.IncludeRootMessage
	}
	if !IsNil(o.Datasource) {
		toSerialize["datasource"] = o.Datasource
	}
	if !IsNil(o.DatasourceInstanceDisplayName) {
		toSerialize["datasourceInstanceDisplayName"] = o.DatasourceInstanceDisplayName
	}
	if !IsNil(o.Sc) {
		toSerialize["sc"] = o.Sc
	}
	return toSerialize, nil
}

type NullableMessagesRequest struct {
	value *MessagesRequest
	isSet bool
}

func (v NullableMessagesRequest) Get() *MessagesRequest {
	return v.value
}

func (v *NullableMessagesRequest) Set(val *MessagesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMessagesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMessagesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessagesRequest(val *MessagesRequest) *NullableMessagesRequest {
	return &NullableMessagesRequest{value: val, isSet: true}
}

func (v NullableMessagesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessagesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


