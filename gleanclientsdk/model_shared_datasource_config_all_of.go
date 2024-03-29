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

// checks if the SharedDatasourceConfigAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SharedDatasourceConfigAllOf{}

// SharedDatasourceConfigAllOf struct for SharedDatasourceConfigAllOf
type SharedDatasourceConfigAllOf struct {
	// The (non-unique) name of the datasource of which this config is an instance, e.g. \"jira\".
	DatasourceName *string `json:"datasourceName,omitempty"`
	// The instance of the datasource for this particular config, e.g. \"onprem\".
	InstanceOnlyName *string `json:"instanceOnlyName,omitempty"`
	// A human readable string identifying this instance as compared to its peers, e.g. \"github.com/askscio\" or \"github.askscio.com\"
	InstanceDescription *string `json:"instanceDescription,omitempty"`
}

// NewSharedDatasourceConfigAllOf instantiates a new SharedDatasourceConfigAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSharedDatasourceConfigAllOf() *SharedDatasourceConfigAllOf {
	this := SharedDatasourceConfigAllOf{}
	return &this
}

// NewSharedDatasourceConfigAllOfWithDefaults instantiates a new SharedDatasourceConfigAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSharedDatasourceConfigAllOfWithDefaults() *SharedDatasourceConfigAllOf {
	this := SharedDatasourceConfigAllOf{}
	return &this
}

// GetDatasourceName returns the DatasourceName field value if set, zero value otherwise.
func (o *SharedDatasourceConfigAllOf) GetDatasourceName() string {
	if o == nil || IsNil(o.DatasourceName) {
		var ret string
		return ret
	}
	return *o.DatasourceName
}

// GetDatasourceNameOk returns a tuple with the DatasourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedDatasourceConfigAllOf) GetDatasourceNameOk() (*string, bool) {
	if o == nil || IsNil(o.DatasourceName) {
		return nil, false
	}
	return o.DatasourceName, true
}

// HasDatasourceName returns a boolean if a field has been set.
func (o *SharedDatasourceConfigAllOf) HasDatasourceName() bool {
	if o != nil && !IsNil(o.DatasourceName) {
		return true
	}

	return false
}

// SetDatasourceName gets a reference to the given string and assigns it to the DatasourceName field.
func (o *SharedDatasourceConfigAllOf) SetDatasourceName(v string) {
	o.DatasourceName = &v
}

// GetInstanceOnlyName returns the InstanceOnlyName field value if set, zero value otherwise.
func (o *SharedDatasourceConfigAllOf) GetInstanceOnlyName() string {
	if o == nil || IsNil(o.InstanceOnlyName) {
		var ret string
		return ret
	}
	return *o.InstanceOnlyName
}

// GetInstanceOnlyNameOk returns a tuple with the InstanceOnlyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedDatasourceConfigAllOf) GetInstanceOnlyNameOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceOnlyName) {
		return nil, false
	}
	return o.InstanceOnlyName, true
}

// HasInstanceOnlyName returns a boolean if a field has been set.
func (o *SharedDatasourceConfigAllOf) HasInstanceOnlyName() bool {
	if o != nil && !IsNil(o.InstanceOnlyName) {
		return true
	}

	return false
}

// SetInstanceOnlyName gets a reference to the given string and assigns it to the InstanceOnlyName field.
func (o *SharedDatasourceConfigAllOf) SetInstanceOnlyName(v string) {
	o.InstanceOnlyName = &v
}

// GetInstanceDescription returns the InstanceDescription field value if set, zero value otherwise.
func (o *SharedDatasourceConfigAllOf) GetInstanceDescription() string {
	if o == nil || IsNil(o.InstanceDescription) {
		var ret string
		return ret
	}
	return *o.InstanceDescription
}

// GetInstanceDescriptionOk returns a tuple with the InstanceDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedDatasourceConfigAllOf) GetInstanceDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceDescription) {
		return nil, false
	}
	return o.InstanceDescription, true
}

// HasInstanceDescription returns a boolean if a field has been set.
func (o *SharedDatasourceConfigAllOf) HasInstanceDescription() bool {
	if o != nil && !IsNil(o.InstanceDescription) {
		return true
	}

	return false
}

// SetInstanceDescription gets a reference to the given string and assigns it to the InstanceDescription field.
func (o *SharedDatasourceConfigAllOf) SetInstanceDescription(v string) {
	o.InstanceDescription = &v
}

func (o SharedDatasourceConfigAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SharedDatasourceConfigAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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

type NullableSharedDatasourceConfigAllOf struct {
	value *SharedDatasourceConfigAllOf
	isSet bool
}

func (v NullableSharedDatasourceConfigAllOf) Get() *SharedDatasourceConfigAllOf {
	return v.value
}

func (v *NullableSharedDatasourceConfigAllOf) Set(val *SharedDatasourceConfigAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSharedDatasourceConfigAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSharedDatasourceConfigAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSharedDatasourceConfigAllOf(val *SharedDatasourceConfigAllOf) *NullableSharedDatasourceConfigAllOf {
	return &NullableSharedDatasourceConfigAllOf{value: val, isSet: true}
}

func (v NullableSharedDatasourceConfigAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSharedDatasourceConfigAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


