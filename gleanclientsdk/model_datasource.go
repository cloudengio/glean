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

// checks if the Datasource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Datasource{}

// Datasource A datasource. Idea: only place fields that belong to a datasource and not any of its particular instances here. 
type Datasource struct {
	// The required id of the datasource. Example: zendesk
	Id *string `json:"id,omitempty"`
	// The functionality provided by the datasource, such as providing searchable content or SSO access.
	Capabilities []DatasourceCapability `json:"capabilities,omitempty"`
}

// NewDatasource instantiates a new Datasource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatasource() *Datasource {
	this := Datasource{}
	return &this
}

// NewDatasourceWithDefaults instantiates a new Datasource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatasourceWithDefaults() *Datasource {
	this := Datasource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Datasource) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Datasource) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Datasource) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Datasource) SetId(v string) {
	o.Id = &v
}

// GetCapabilities returns the Capabilities field value if set, zero value otherwise.
func (o *Datasource) GetCapabilities() []DatasourceCapability {
	if o == nil || IsNil(o.Capabilities) {
		var ret []DatasourceCapability
		return ret
	}
	return o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Datasource) GetCapabilitiesOk() ([]DatasourceCapability, bool) {
	if o == nil || IsNil(o.Capabilities) {
		return nil, false
	}
	return o.Capabilities, true
}

// HasCapabilities returns a boolean if a field has been set.
func (o *Datasource) HasCapabilities() bool {
	if o != nil && !IsNil(o.Capabilities) {
		return true
	}

	return false
}

// SetCapabilities gets a reference to the given []DatasourceCapability and assigns it to the Capabilities field.
func (o *Datasource) SetCapabilities(v []DatasourceCapability) {
	o.Capabilities = v
}

func (o Datasource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Datasource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Capabilities) {
		toSerialize["capabilities"] = o.Capabilities
	}
	return toSerialize, nil
}

type NullableDatasource struct {
	value *Datasource
	isSet bool
}

func (v NullableDatasource) Get() *Datasource {
	return v.value
}

func (v *NullableDatasource) Set(val *Datasource) {
	v.value = val
	v.isSet = true
}

func (v NullableDatasource) IsSet() bool {
	return v.isSet
}

func (v *NullableDatasource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatasource(val *Datasource) *NullableDatasource {
	return &NullableDatasource{value: val, isSet: true}
}

func (v NullableDatasource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatasource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


