/*
Glean Client API

# Introduction These are the public APIs to enable implementing a custom client interface to the Glean system.  # Usage guidelines This API is evolving fast. Glean will provide advance notice of any planned backwards incompatible changes along with a 6-month sunset period for anything that requires developers to adopt the new versions. 

API version: 0.9.0
Contact: support@glean.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gleanclientsdk

import (
	"encoding/json"
)

// checks if the EditAnswerRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EditAnswerRequest{}

// EditAnswerRequest struct for EditAnswerRequest
type EditAnswerRequest struct {
	// The opaque id of the answer.
	Id int32 `json:"id"`
	// Internal document id of the answer. We support using the document id for cases where the client doesn't have the integer id available. If both are available, using the integer id is preferred.
	DocId *string `json:"docId,omitempty"`
	Question *string `json:"question,omitempty"`
	// Additional ways of phrasing this question.
	QuestionVariations []string `json:"questionVariations,omitempty"`
	// The plain text answer to the question.
	BodyText *string `json:"bodyText,omitempty"`
	// The parent board ID of this Answer, or 0 if it's a floating Answer.
	BoardId *int32 `json:"boardId,omitempty"`
	// Filters which restrict who should see the answer. Values are taken from the corresponding filters in people search.
	AudienceFilters []FacetFilter `json:"audienceFilters,omitempty"`
	// A list of user roles for the answer added by the owner.
	AddedRoles []UserRoleSpecification `json:"addedRoles,omitempty"`
	// A list of user roles for the answer removed by the owner.
	RemovedRoles []UserRoleSpecification `json:"removedRoles,omitempty"`
	// A list of roles for this answer explicitly granted by an owner, editor, or admin.
	Roles []UserRoleSpecification `json:"roles,omitempty"`
	SourceDocumentSpec *DocumentSpec `json:"sourceDocumentSpec,omitempty"`
	// IDs of collections to which a document is added.
	AddedCollections []int32 `json:"addedCollections,omitempty"`
	// IDs of collections from which a document is removed.
	RemovedCollections []int32 `json:"removedCollections,omitempty"`
	CombinedAnswerText *StructuredTextMutableProperties `json:"combinedAnswerText,omitempty"`
}

// NewEditAnswerRequest instantiates a new EditAnswerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditAnswerRequest(id int32) *EditAnswerRequest {
	this := EditAnswerRequest{}
	this.Id = id
	return &this
}

// NewEditAnswerRequestWithDefaults instantiates a new EditAnswerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditAnswerRequestWithDefaults() *EditAnswerRequest {
	this := EditAnswerRequest{}
	return &this
}

// GetId returns the Id field value
func (o *EditAnswerRequest) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EditAnswerRequest) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EditAnswerRequest) SetId(v int32) {
	o.Id = v
}

// GetDocId returns the DocId field value if set, zero value otherwise.
func (o *EditAnswerRequest) GetDocId() string {
	if o == nil || IsNil(o.DocId) {
		var ret string
		return ret
	}
	return *o.DocId
}

// GetDocIdOk returns a tuple with the DocId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditAnswerRequest) GetDocIdOk() (*string, bool) {
	if o == nil || IsNil(o.DocId) {
		return nil, false
	}
	return o.DocId, true
}

// HasDocId returns a boolean if a field has been set.
func (o *EditAnswerRequest) HasDocId() bool {
	if o != nil && !IsNil(o.DocId) {
		return true
	}

	return false
}

// SetDocId gets a reference to the given string and assigns it to the DocId field.
func (o *EditAnswerRequest) SetDocId(v string) {
	o.DocId = &v
}

// GetQuestion returns the Question field value if set, zero value otherwise.
func (o *EditAnswerRequest) GetQuestion() string {
	if o == nil || IsNil(o.Question) {
		var ret string
		return ret
	}
	return *o.Question
}

// GetQuestionOk returns a tuple with the Question field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditAnswerRequest) GetQuestionOk() (*string, bool) {
	if o == nil || IsNil(o.Question) {
		return nil, false
	}
	return o.Question, true
}

// HasQuestion returns a boolean if a field has been set.
func (o *EditAnswerRequest) HasQuestion() bool {
	if o != nil && !IsNil(o.Question) {
		return true
	}

	return false
}

// SetQuestion gets a reference to the given string and assigns it to the Question field.
func (o *EditAnswerRequest) SetQuestion(v string) {
	o.Question = &v
}

// GetQuestionVariations returns the QuestionVariations field value if set, zero value otherwise.
func (o *EditAnswerRequest) GetQuestionVariations() []string {
	if o == nil || IsNil(o.QuestionVariations) {
		var ret []string
		return ret
	}
	return o.QuestionVariations
}

// GetQuestionVariationsOk returns a tuple with the QuestionVariations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditAnswerRequest) GetQuestionVariationsOk() ([]string, bool) {
	if o == nil || IsNil(o.QuestionVariations) {
		return nil, false
	}
	return o.QuestionVariations, true
}

// HasQuestionVariations returns a boolean if a field has been set.
func (o *EditAnswerRequest) HasQuestionVariations() bool {
	if o != nil && !IsNil(o.QuestionVariations) {
		return true
	}

	return false
}

// SetQuestionVariations gets a reference to the given []string and assigns it to the QuestionVariations field.
func (o *EditAnswerRequest) SetQuestionVariations(v []string) {
	o.QuestionVariations = v
}

// GetBodyText returns the BodyText field value if set, zero value otherwise.
func (o *EditAnswerRequest) GetBodyText() string {
	if o == nil || IsNil(o.BodyText) {
		var ret string
		return ret
	}
	return *o.BodyText
}

// GetBodyTextOk returns a tuple with the BodyText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditAnswerRequest) GetBodyTextOk() (*string, bool) {
	if o == nil || IsNil(o.BodyText) {
		return nil, false
	}
	return o.BodyText, true
}

// HasBodyText returns a boolean if a field has been set.
func (o *EditAnswerRequest) HasBodyText() bool {
	if o != nil && !IsNil(o.BodyText) {
		return true
	}

	return false
}

// SetBodyText gets a reference to the given string and assigns it to the BodyText field.
func (o *EditAnswerRequest) SetBodyText(v string) {
	o.BodyText = &v
}

// GetBoardId returns the BoardId field value if set, zero value otherwise.
func (o *EditAnswerRequest) GetBoardId() int32 {
	if o == nil || IsNil(o.BoardId) {
		var ret int32
		return ret
	}
	return *o.BoardId
}

// GetBoardIdOk returns a tuple with the BoardId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditAnswerRequest) GetBoardIdOk() (*int32, bool) {
	if o == nil || IsNil(o.BoardId) {
		return nil, false
	}
	return o.BoardId, true
}

// HasBoardId returns a boolean if a field has been set.
func (o *EditAnswerRequest) HasBoardId() bool {
	if o != nil && !IsNil(o.BoardId) {
		return true
	}

	return false
}

// SetBoardId gets a reference to the given int32 and assigns it to the BoardId field.
func (o *EditAnswerRequest) SetBoardId(v int32) {
	o.BoardId = &v
}

// GetAudienceFilters returns the AudienceFilters field value if set, zero value otherwise.
func (o *EditAnswerRequest) GetAudienceFilters() []FacetFilter {
	if o == nil || IsNil(o.AudienceFilters) {
		var ret []FacetFilter
		return ret
	}
	return o.AudienceFilters
}

// GetAudienceFiltersOk returns a tuple with the AudienceFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditAnswerRequest) GetAudienceFiltersOk() ([]FacetFilter, bool) {
	if o == nil || IsNil(o.AudienceFilters) {
		return nil, false
	}
	return o.AudienceFilters, true
}

// HasAudienceFilters returns a boolean if a field has been set.
func (o *EditAnswerRequest) HasAudienceFilters() bool {
	if o != nil && !IsNil(o.AudienceFilters) {
		return true
	}

	return false
}

// SetAudienceFilters gets a reference to the given []FacetFilter and assigns it to the AudienceFilters field.
func (o *EditAnswerRequest) SetAudienceFilters(v []FacetFilter) {
	o.AudienceFilters = v
}

// GetAddedRoles returns the AddedRoles field value if set, zero value otherwise.
func (o *EditAnswerRequest) GetAddedRoles() []UserRoleSpecification {
	if o == nil || IsNil(o.AddedRoles) {
		var ret []UserRoleSpecification
		return ret
	}
	return o.AddedRoles
}

// GetAddedRolesOk returns a tuple with the AddedRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditAnswerRequest) GetAddedRolesOk() ([]UserRoleSpecification, bool) {
	if o == nil || IsNil(o.AddedRoles) {
		return nil, false
	}
	return o.AddedRoles, true
}

// HasAddedRoles returns a boolean if a field has been set.
func (o *EditAnswerRequest) HasAddedRoles() bool {
	if o != nil && !IsNil(o.AddedRoles) {
		return true
	}

	return false
}

// SetAddedRoles gets a reference to the given []UserRoleSpecification and assigns it to the AddedRoles field.
func (o *EditAnswerRequest) SetAddedRoles(v []UserRoleSpecification) {
	o.AddedRoles = v
}

// GetRemovedRoles returns the RemovedRoles field value if set, zero value otherwise.
func (o *EditAnswerRequest) GetRemovedRoles() []UserRoleSpecification {
	if o == nil || IsNil(o.RemovedRoles) {
		var ret []UserRoleSpecification
		return ret
	}
	return o.RemovedRoles
}

// GetRemovedRolesOk returns a tuple with the RemovedRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditAnswerRequest) GetRemovedRolesOk() ([]UserRoleSpecification, bool) {
	if o == nil || IsNil(o.RemovedRoles) {
		return nil, false
	}
	return o.RemovedRoles, true
}

// HasRemovedRoles returns a boolean if a field has been set.
func (o *EditAnswerRequest) HasRemovedRoles() bool {
	if o != nil && !IsNil(o.RemovedRoles) {
		return true
	}

	return false
}

// SetRemovedRoles gets a reference to the given []UserRoleSpecification and assigns it to the RemovedRoles field.
func (o *EditAnswerRequest) SetRemovedRoles(v []UserRoleSpecification) {
	o.RemovedRoles = v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *EditAnswerRequest) GetRoles() []UserRoleSpecification {
	if o == nil || IsNil(o.Roles) {
		var ret []UserRoleSpecification
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditAnswerRequest) GetRolesOk() ([]UserRoleSpecification, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *EditAnswerRequest) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []UserRoleSpecification and assigns it to the Roles field.
func (o *EditAnswerRequest) SetRoles(v []UserRoleSpecification) {
	o.Roles = v
}

// GetSourceDocumentSpec returns the SourceDocumentSpec field value if set, zero value otherwise.
func (o *EditAnswerRequest) GetSourceDocumentSpec() DocumentSpec {
	if o == nil || IsNil(o.SourceDocumentSpec) {
		var ret DocumentSpec
		return ret
	}
	return *o.SourceDocumentSpec
}

// GetSourceDocumentSpecOk returns a tuple with the SourceDocumentSpec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditAnswerRequest) GetSourceDocumentSpecOk() (*DocumentSpec, bool) {
	if o == nil || IsNil(o.SourceDocumentSpec) {
		return nil, false
	}
	return o.SourceDocumentSpec, true
}

// HasSourceDocumentSpec returns a boolean if a field has been set.
func (o *EditAnswerRequest) HasSourceDocumentSpec() bool {
	if o != nil && !IsNil(o.SourceDocumentSpec) {
		return true
	}

	return false
}

// SetSourceDocumentSpec gets a reference to the given DocumentSpec and assigns it to the SourceDocumentSpec field.
func (o *EditAnswerRequest) SetSourceDocumentSpec(v DocumentSpec) {
	o.SourceDocumentSpec = &v
}

// GetAddedCollections returns the AddedCollections field value if set, zero value otherwise.
func (o *EditAnswerRequest) GetAddedCollections() []int32 {
	if o == nil || IsNil(o.AddedCollections) {
		var ret []int32
		return ret
	}
	return o.AddedCollections
}

// GetAddedCollectionsOk returns a tuple with the AddedCollections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditAnswerRequest) GetAddedCollectionsOk() ([]int32, bool) {
	if o == nil || IsNil(o.AddedCollections) {
		return nil, false
	}
	return o.AddedCollections, true
}

// HasAddedCollections returns a boolean if a field has been set.
func (o *EditAnswerRequest) HasAddedCollections() bool {
	if o != nil && !IsNil(o.AddedCollections) {
		return true
	}

	return false
}

// SetAddedCollections gets a reference to the given []int32 and assigns it to the AddedCollections field.
func (o *EditAnswerRequest) SetAddedCollections(v []int32) {
	o.AddedCollections = v
}

// GetRemovedCollections returns the RemovedCollections field value if set, zero value otherwise.
func (o *EditAnswerRequest) GetRemovedCollections() []int32 {
	if o == nil || IsNil(o.RemovedCollections) {
		var ret []int32
		return ret
	}
	return o.RemovedCollections
}

// GetRemovedCollectionsOk returns a tuple with the RemovedCollections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditAnswerRequest) GetRemovedCollectionsOk() ([]int32, bool) {
	if o == nil || IsNil(o.RemovedCollections) {
		return nil, false
	}
	return o.RemovedCollections, true
}

// HasRemovedCollections returns a boolean if a field has been set.
func (o *EditAnswerRequest) HasRemovedCollections() bool {
	if o != nil && !IsNil(o.RemovedCollections) {
		return true
	}

	return false
}

// SetRemovedCollections gets a reference to the given []int32 and assigns it to the RemovedCollections field.
func (o *EditAnswerRequest) SetRemovedCollections(v []int32) {
	o.RemovedCollections = v
}

// GetCombinedAnswerText returns the CombinedAnswerText field value if set, zero value otherwise.
func (o *EditAnswerRequest) GetCombinedAnswerText() StructuredTextMutableProperties {
	if o == nil || IsNil(o.CombinedAnswerText) {
		var ret StructuredTextMutableProperties
		return ret
	}
	return *o.CombinedAnswerText
}

// GetCombinedAnswerTextOk returns a tuple with the CombinedAnswerText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditAnswerRequest) GetCombinedAnswerTextOk() (*StructuredTextMutableProperties, bool) {
	if o == nil || IsNil(o.CombinedAnswerText) {
		return nil, false
	}
	return o.CombinedAnswerText, true
}

// HasCombinedAnswerText returns a boolean if a field has been set.
func (o *EditAnswerRequest) HasCombinedAnswerText() bool {
	if o != nil && !IsNil(o.CombinedAnswerText) {
		return true
	}

	return false
}

// SetCombinedAnswerText gets a reference to the given StructuredTextMutableProperties and assigns it to the CombinedAnswerText field.
func (o *EditAnswerRequest) SetCombinedAnswerText(v StructuredTextMutableProperties) {
	o.CombinedAnswerText = &v
}

func (o EditAnswerRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EditAnswerRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.DocId) {
		toSerialize["docId"] = o.DocId
	}
	if !IsNil(o.Question) {
		toSerialize["question"] = o.Question
	}
	if !IsNil(o.QuestionVariations) {
		toSerialize["questionVariations"] = o.QuestionVariations
	}
	if !IsNil(o.BodyText) {
		toSerialize["bodyText"] = o.BodyText
	}
	if !IsNil(o.BoardId) {
		toSerialize["boardId"] = o.BoardId
	}
	if !IsNil(o.AudienceFilters) {
		toSerialize["audienceFilters"] = o.AudienceFilters
	}
	if !IsNil(o.AddedRoles) {
		toSerialize["addedRoles"] = o.AddedRoles
	}
	if !IsNil(o.RemovedRoles) {
		toSerialize["removedRoles"] = o.RemovedRoles
	}
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	if !IsNil(o.SourceDocumentSpec) {
		toSerialize["sourceDocumentSpec"] = o.SourceDocumentSpec
	}
	if !IsNil(o.AddedCollections) {
		toSerialize["addedCollections"] = o.AddedCollections
	}
	if !IsNil(o.RemovedCollections) {
		toSerialize["removedCollections"] = o.RemovedCollections
	}
	if !IsNil(o.CombinedAnswerText) {
		toSerialize["combinedAnswerText"] = o.CombinedAnswerText
	}
	return toSerialize, nil
}

type NullableEditAnswerRequest struct {
	value *EditAnswerRequest
	isSet bool
}

func (v NullableEditAnswerRequest) Get() *EditAnswerRequest {
	return v.value
}

func (v *NullableEditAnswerRequest) Set(val *EditAnswerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEditAnswerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEditAnswerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditAnswerRequest(val *EditAnswerRequest) *NullableEditAnswerRequest {
	return &NullableEditAnswerRequest{value: val, isSet: true}
}

func (v NullableEditAnswerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditAnswerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

