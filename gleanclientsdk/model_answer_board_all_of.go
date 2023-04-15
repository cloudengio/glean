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
	"time"
)

// checks if the AnswerBoardAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnswerBoardAllOf{}

// AnswerBoardAllOf struct for AnswerBoardAllOf
type AnswerBoardAllOf struct {
	// The unique ID of the Answer Board.
	Id int32 `json:"id"`
	CreateTime *time.Time `json:"createTime,omitempty"`
	UpdateTime *time.Time `json:"updateTime,omitempty"`
	Creator *Person `json:"creator,omitempty"`
	UpdatedBy *Person `json:"updatedBy,omitempty"`
	// The number of items currently in the Answer Board. Separated from the actual items so we can grab the count without items.
	ItemCount *int32 `json:"itemCount,omitempty"`
	// A list of user roles for the Answer Board.
	Roles []UserRoleSpecification `json:"roles,omitempty"`
}

// NewAnswerBoardAllOf instantiates a new AnswerBoardAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnswerBoardAllOf(id int32) *AnswerBoardAllOf {
	this := AnswerBoardAllOf{}
	this.Id = id
	return &this
}

// NewAnswerBoardAllOfWithDefaults instantiates a new AnswerBoardAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnswerBoardAllOfWithDefaults() *AnswerBoardAllOf {
	this := AnswerBoardAllOf{}
	return &this
}

// GetId returns the Id field value
func (o *AnswerBoardAllOf) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AnswerBoardAllOf) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AnswerBoardAllOf) SetId(v int32) {
	o.Id = v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *AnswerBoardAllOf) GetCreateTime() time.Time {
	if o == nil || IsNil(o.CreateTime) {
		var ret time.Time
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnswerBoardAllOf) GetCreateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *AnswerBoardAllOf) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given time.Time and assigns it to the CreateTime field.
func (o *AnswerBoardAllOf) SetCreateTime(v time.Time) {
	o.CreateTime = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *AnswerBoardAllOf) GetUpdateTime() time.Time {
	if o == nil || IsNil(o.UpdateTime) {
		var ret time.Time
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnswerBoardAllOf) GetUpdateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *AnswerBoardAllOf) HasUpdateTime() bool {
	if o != nil && !IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given time.Time and assigns it to the UpdateTime field.
func (o *AnswerBoardAllOf) SetUpdateTime(v time.Time) {
	o.UpdateTime = &v
}

// GetCreator returns the Creator field value if set, zero value otherwise.
func (o *AnswerBoardAllOf) GetCreator() Person {
	if o == nil || IsNil(o.Creator) {
		var ret Person
		return ret
	}
	return *o.Creator
}

// GetCreatorOk returns a tuple with the Creator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnswerBoardAllOf) GetCreatorOk() (*Person, bool) {
	if o == nil || IsNil(o.Creator) {
		return nil, false
	}
	return o.Creator, true
}

// HasCreator returns a boolean if a field has been set.
func (o *AnswerBoardAllOf) HasCreator() bool {
	if o != nil && !IsNil(o.Creator) {
		return true
	}

	return false
}

// SetCreator gets a reference to the given Person and assigns it to the Creator field.
func (o *AnswerBoardAllOf) SetCreator(v Person) {
	o.Creator = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *AnswerBoardAllOf) GetUpdatedBy() Person {
	if o == nil || IsNil(o.UpdatedBy) {
		var ret Person
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnswerBoardAllOf) GetUpdatedByOk() (*Person, bool) {
	if o == nil || IsNil(o.UpdatedBy) {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *AnswerBoardAllOf) HasUpdatedBy() bool {
	if o != nil && !IsNil(o.UpdatedBy) {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given Person and assigns it to the UpdatedBy field.
func (o *AnswerBoardAllOf) SetUpdatedBy(v Person) {
	o.UpdatedBy = &v
}

// GetItemCount returns the ItemCount field value if set, zero value otherwise.
func (o *AnswerBoardAllOf) GetItemCount() int32 {
	if o == nil || IsNil(o.ItemCount) {
		var ret int32
		return ret
	}
	return *o.ItemCount
}

// GetItemCountOk returns a tuple with the ItemCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnswerBoardAllOf) GetItemCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ItemCount) {
		return nil, false
	}
	return o.ItemCount, true
}

// HasItemCount returns a boolean if a field has been set.
func (o *AnswerBoardAllOf) HasItemCount() bool {
	if o != nil && !IsNil(o.ItemCount) {
		return true
	}

	return false
}

// SetItemCount gets a reference to the given int32 and assigns it to the ItemCount field.
func (o *AnswerBoardAllOf) SetItemCount(v int32) {
	o.ItemCount = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *AnswerBoardAllOf) GetRoles() []UserRoleSpecification {
	if o == nil || IsNil(o.Roles) {
		var ret []UserRoleSpecification
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnswerBoardAllOf) GetRolesOk() ([]UserRoleSpecification, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *AnswerBoardAllOf) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []UserRoleSpecification and assigns it to the Roles field.
func (o *AnswerBoardAllOf) SetRoles(v []UserRoleSpecification) {
	o.Roles = v
}

func (o AnswerBoardAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnswerBoardAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.CreateTime) {
		toSerialize["createTime"] = o.CreateTime
	}
	if !IsNil(o.UpdateTime) {
		toSerialize["updateTime"] = o.UpdateTime
	}
	if !IsNil(o.Creator) {
		toSerialize["creator"] = o.Creator
	}
	if !IsNil(o.UpdatedBy) {
		toSerialize["updatedBy"] = o.UpdatedBy
	}
	if !IsNil(o.ItemCount) {
		toSerialize["itemCount"] = o.ItemCount
	}
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	return toSerialize, nil
}

type NullableAnswerBoardAllOf struct {
	value *AnswerBoardAllOf
	isSet bool
}

func (v NullableAnswerBoardAllOf) Get() *AnswerBoardAllOf {
	return v.value
}

func (v *NullableAnswerBoardAllOf) Set(val *AnswerBoardAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAnswerBoardAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAnswerBoardAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnswerBoardAllOf(val *AnswerBoardAllOf) *NullableAnswerBoardAllOf {
	return &NullableAnswerBoardAllOf{value: val, isSet: true}
}

func (v NullableAnswerBoardAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnswerBoardAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


