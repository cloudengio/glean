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

// checks if the GetCollectionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCollectionRequest{}

// GetCollectionRequest struct for GetCollectionRequest
type GetCollectionRequest struct {
	// The id of the collection to be retrieved.
	Id int32 `json:"id"`
	// Whether or not to include the Collection Items in this Collection. Only request if absolutely required, as this is expensive.
	WithItems *bool `json:"withItems,omitempty"`
	// Whether or not to include the top level Collection in this Collection's hierarchy.
	WithHierarchy *bool `json:"withHierarchy,omitempty"`
	// The datasource allowed in the collection returned.
	AllowedDatasource *string `json:"allowedDatasource,omitempty"`
}

// NewGetCollectionRequest instantiates a new GetCollectionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCollectionRequest(id int32) *GetCollectionRequest {
	this := GetCollectionRequest{}
	this.Id = id
	return &this
}

// NewGetCollectionRequestWithDefaults instantiates a new GetCollectionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCollectionRequestWithDefaults() *GetCollectionRequest {
	this := GetCollectionRequest{}
	return &this
}

// GetId returns the Id field value
func (o *GetCollectionRequest) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GetCollectionRequest) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GetCollectionRequest) SetId(v int32) {
	o.Id = v
}

// GetWithItems returns the WithItems field value if set, zero value otherwise.
func (o *GetCollectionRequest) GetWithItems() bool {
	if o == nil || IsNil(o.WithItems) {
		var ret bool
		return ret
	}
	return *o.WithItems
}

// GetWithItemsOk returns a tuple with the WithItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCollectionRequest) GetWithItemsOk() (*bool, bool) {
	if o == nil || IsNil(o.WithItems) {
		return nil, false
	}
	return o.WithItems, true
}

// HasWithItems returns a boolean if a field has been set.
func (o *GetCollectionRequest) HasWithItems() bool {
	if o != nil && !IsNil(o.WithItems) {
		return true
	}

	return false
}

// SetWithItems gets a reference to the given bool and assigns it to the WithItems field.
func (o *GetCollectionRequest) SetWithItems(v bool) {
	o.WithItems = &v
}

// GetWithHierarchy returns the WithHierarchy field value if set, zero value otherwise.
func (o *GetCollectionRequest) GetWithHierarchy() bool {
	if o == nil || IsNil(o.WithHierarchy) {
		var ret bool
		return ret
	}
	return *o.WithHierarchy
}

// GetWithHierarchyOk returns a tuple with the WithHierarchy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCollectionRequest) GetWithHierarchyOk() (*bool, bool) {
	if o == nil || IsNil(o.WithHierarchy) {
		return nil, false
	}
	return o.WithHierarchy, true
}

// HasWithHierarchy returns a boolean if a field has been set.
func (o *GetCollectionRequest) HasWithHierarchy() bool {
	if o != nil && !IsNil(o.WithHierarchy) {
		return true
	}

	return false
}

// SetWithHierarchy gets a reference to the given bool and assigns it to the WithHierarchy field.
func (o *GetCollectionRequest) SetWithHierarchy(v bool) {
	o.WithHierarchy = &v
}

// GetAllowedDatasource returns the AllowedDatasource field value if set, zero value otherwise.
func (o *GetCollectionRequest) GetAllowedDatasource() string {
	if o == nil || IsNil(o.AllowedDatasource) {
		var ret string
		return ret
	}
	return *o.AllowedDatasource
}

// GetAllowedDatasourceOk returns a tuple with the AllowedDatasource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCollectionRequest) GetAllowedDatasourceOk() (*string, bool) {
	if o == nil || IsNil(o.AllowedDatasource) {
		return nil, false
	}
	return o.AllowedDatasource, true
}

// HasAllowedDatasource returns a boolean if a field has been set.
func (o *GetCollectionRequest) HasAllowedDatasource() bool {
	if o != nil && !IsNil(o.AllowedDatasource) {
		return true
	}

	return false
}

// SetAllowedDatasource gets a reference to the given string and assigns it to the AllowedDatasource field.
func (o *GetCollectionRequest) SetAllowedDatasource(v string) {
	o.AllowedDatasource = &v
}

func (o GetCollectionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCollectionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.WithItems) {
		toSerialize["withItems"] = o.WithItems
	}
	if !IsNil(o.WithHierarchy) {
		toSerialize["withHierarchy"] = o.WithHierarchy
	}
	if !IsNil(o.AllowedDatasource) {
		toSerialize["allowedDatasource"] = o.AllowedDatasource
	}
	return toSerialize, nil
}

type NullableGetCollectionRequest struct {
	value *GetCollectionRequest
	isSet bool
}

func (v NullableGetCollectionRequest) Get() *GetCollectionRequest {
	return v.value
}

func (v *NullableGetCollectionRequest) Set(val *GetCollectionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCollectionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCollectionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCollectionRequest(val *GetCollectionRequest) *NullableGetCollectionRequest {
	return &NullableGetCollectionRequest{value: val, isSet: true}
}

func (v NullableGetCollectionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCollectionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


