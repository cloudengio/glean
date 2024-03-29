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
	"time"
)

// checks if the AnswerAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnswerAllOf{}

// AnswerAllOf struct for AnswerAllOf
type AnswerAllOf struct {
	CombinedAnswerText *StructuredText `json:"combinedAnswerText,omitempty"`
	Likes *AnswerLikes `json:"likes,omitempty"`
	// DEPRECATED - use roles instead. User's role on the specific answer.
	// Deprecated
	UserRole *string `json:"userRole,omitempty"`
	Author *Person `json:"author,omitempty"`
	// The time the answer was created in ISO format (ISO 8601).
	CreateTime *time.Time `json:"createTime,omitempty"`
	// The time the answer was last updated in ISO format (ISO 8601).
	UpdateTime *time.Time `json:"updateTime,omitempty"`
	UpdatedBy *Person `json:"updatedBy,omitempty"`
	Verification *Verification `json:"verification,omitempty"`
	Board *AnswerBoard `json:"board,omitempty"`
	// The collections to which the answer belongs.
	Collections []Collection `json:"collections,omitempty"`
	// The document's document_category(.proto).
	DocumentCategory *string `json:"documentCategory,omitempty"`
	SourceDocument *Document `json:"sourceDocument,omitempty"`
}

// NewAnswerAllOf instantiates a new AnswerAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnswerAllOf() *AnswerAllOf {
	this := AnswerAllOf{}
	return &this
}

// NewAnswerAllOfWithDefaults instantiates a new AnswerAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnswerAllOfWithDefaults() *AnswerAllOf {
	this := AnswerAllOf{}
	return &this
}

// GetCombinedAnswerText returns the CombinedAnswerText field value if set, zero value otherwise.
func (o *AnswerAllOf) GetCombinedAnswerText() StructuredText {
	if o == nil || IsNil(o.CombinedAnswerText) {
		var ret StructuredText
		return ret
	}
	return *o.CombinedAnswerText
}

// GetCombinedAnswerTextOk returns a tuple with the CombinedAnswerText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnswerAllOf) GetCombinedAnswerTextOk() (*StructuredText, bool) {
	if o == nil || IsNil(o.CombinedAnswerText) {
		return nil, false
	}
	return o.CombinedAnswerText, true
}

// HasCombinedAnswerText returns a boolean if a field has been set.
func (o *AnswerAllOf) HasCombinedAnswerText() bool {
	if o != nil && !IsNil(o.CombinedAnswerText) {
		return true
	}

	return false
}

// SetCombinedAnswerText gets a reference to the given StructuredText and assigns it to the CombinedAnswerText field.
func (o *AnswerAllOf) SetCombinedAnswerText(v StructuredText) {
	o.CombinedAnswerText = &v
}

// GetLikes returns the Likes field value if set, zero value otherwise.
func (o *AnswerAllOf) GetLikes() AnswerLikes {
	if o == nil || IsNil(o.Likes) {
		var ret AnswerLikes
		return ret
	}
	return *o.Likes
}

// GetLikesOk returns a tuple with the Likes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnswerAllOf) GetLikesOk() (*AnswerLikes, bool) {
	if o == nil || IsNil(o.Likes) {
		return nil, false
	}
	return o.Likes, true
}

// HasLikes returns a boolean if a field has been set.
func (o *AnswerAllOf) HasLikes() bool {
	if o != nil && !IsNil(o.Likes) {
		return true
	}

	return false
}

// SetLikes gets a reference to the given AnswerLikes and assigns it to the Likes field.
func (o *AnswerAllOf) SetLikes(v AnswerLikes) {
	o.Likes = &v
}

// GetUserRole returns the UserRole field value if set, zero value otherwise.
// Deprecated
func (o *AnswerAllOf) GetUserRole() string {
	if o == nil || IsNil(o.UserRole) {
		var ret string
		return ret
	}
	return *o.UserRole
}

// GetUserRoleOk returns a tuple with the UserRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *AnswerAllOf) GetUserRoleOk() (*string, bool) {
	if o == nil || IsNil(o.UserRole) {
		return nil, false
	}
	return o.UserRole, true
}

// HasUserRole returns a boolean if a field has been set.
func (o *AnswerAllOf) HasUserRole() bool {
	if o != nil && !IsNil(o.UserRole) {
		return true
	}

	return false
}

// SetUserRole gets a reference to the given string and assigns it to the UserRole field.
// Deprecated
func (o *AnswerAllOf) SetUserRole(v string) {
	o.UserRole = &v
}

// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *AnswerAllOf) GetAuthor() Person {
	if o == nil || IsNil(o.Author) {
		var ret Person
		return ret
	}
	return *o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnswerAllOf) GetAuthorOk() (*Person, bool) {
	if o == nil || IsNil(o.Author) {
		return nil, false
	}
	return o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *AnswerAllOf) HasAuthor() bool {
	if o != nil && !IsNil(o.Author) {
		return true
	}

	return false
}

// SetAuthor gets a reference to the given Person and assigns it to the Author field.
func (o *AnswerAllOf) SetAuthor(v Person) {
	o.Author = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *AnswerAllOf) GetCreateTime() time.Time {
	if o == nil || IsNil(o.CreateTime) {
		var ret time.Time
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnswerAllOf) GetCreateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *AnswerAllOf) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given time.Time and assigns it to the CreateTime field.
func (o *AnswerAllOf) SetCreateTime(v time.Time) {
	o.CreateTime = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *AnswerAllOf) GetUpdateTime() time.Time {
	if o == nil || IsNil(o.UpdateTime) {
		var ret time.Time
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnswerAllOf) GetUpdateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *AnswerAllOf) HasUpdateTime() bool {
	if o != nil && !IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given time.Time and assigns it to the UpdateTime field.
func (o *AnswerAllOf) SetUpdateTime(v time.Time) {
	o.UpdateTime = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *AnswerAllOf) GetUpdatedBy() Person {
	if o == nil || IsNil(o.UpdatedBy) {
		var ret Person
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnswerAllOf) GetUpdatedByOk() (*Person, bool) {
	if o == nil || IsNil(o.UpdatedBy) {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *AnswerAllOf) HasUpdatedBy() bool {
	if o != nil && !IsNil(o.UpdatedBy) {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given Person and assigns it to the UpdatedBy field.
func (o *AnswerAllOf) SetUpdatedBy(v Person) {
	o.UpdatedBy = &v
}

// GetVerification returns the Verification field value if set, zero value otherwise.
func (o *AnswerAllOf) GetVerification() Verification {
	if o == nil || IsNil(o.Verification) {
		var ret Verification
		return ret
	}
	return *o.Verification
}

// GetVerificationOk returns a tuple with the Verification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnswerAllOf) GetVerificationOk() (*Verification, bool) {
	if o == nil || IsNil(o.Verification) {
		return nil, false
	}
	return o.Verification, true
}

// HasVerification returns a boolean if a field has been set.
func (o *AnswerAllOf) HasVerification() bool {
	if o != nil && !IsNil(o.Verification) {
		return true
	}

	return false
}

// SetVerification gets a reference to the given Verification and assigns it to the Verification field.
func (o *AnswerAllOf) SetVerification(v Verification) {
	o.Verification = &v
}

// GetBoard returns the Board field value if set, zero value otherwise.
func (o *AnswerAllOf) GetBoard() AnswerBoard {
	if o == nil || IsNil(o.Board) {
		var ret AnswerBoard
		return ret
	}
	return *o.Board
}

// GetBoardOk returns a tuple with the Board field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnswerAllOf) GetBoardOk() (*AnswerBoard, bool) {
	if o == nil || IsNil(o.Board) {
		return nil, false
	}
	return o.Board, true
}

// HasBoard returns a boolean if a field has been set.
func (o *AnswerAllOf) HasBoard() bool {
	if o != nil && !IsNil(o.Board) {
		return true
	}

	return false
}

// SetBoard gets a reference to the given AnswerBoard and assigns it to the Board field.
func (o *AnswerAllOf) SetBoard(v AnswerBoard) {
	o.Board = &v
}

// GetCollections returns the Collections field value if set, zero value otherwise.
func (o *AnswerAllOf) GetCollections() []Collection {
	if o == nil || IsNil(o.Collections) {
		var ret []Collection
		return ret
	}
	return o.Collections
}

// GetCollectionsOk returns a tuple with the Collections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnswerAllOf) GetCollectionsOk() ([]Collection, bool) {
	if o == nil || IsNil(o.Collections) {
		return nil, false
	}
	return o.Collections, true
}

// HasCollections returns a boolean if a field has been set.
func (o *AnswerAllOf) HasCollections() bool {
	if o != nil && !IsNil(o.Collections) {
		return true
	}

	return false
}

// SetCollections gets a reference to the given []Collection and assigns it to the Collections field.
func (o *AnswerAllOf) SetCollections(v []Collection) {
	o.Collections = v
}

// GetDocumentCategory returns the DocumentCategory field value if set, zero value otherwise.
func (o *AnswerAllOf) GetDocumentCategory() string {
	if o == nil || IsNil(o.DocumentCategory) {
		var ret string
		return ret
	}
	return *o.DocumentCategory
}

// GetDocumentCategoryOk returns a tuple with the DocumentCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnswerAllOf) GetDocumentCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.DocumentCategory) {
		return nil, false
	}
	return o.DocumentCategory, true
}

// HasDocumentCategory returns a boolean if a field has been set.
func (o *AnswerAllOf) HasDocumentCategory() bool {
	if o != nil && !IsNil(o.DocumentCategory) {
		return true
	}

	return false
}

// SetDocumentCategory gets a reference to the given string and assigns it to the DocumentCategory field.
func (o *AnswerAllOf) SetDocumentCategory(v string) {
	o.DocumentCategory = &v
}

// GetSourceDocument returns the SourceDocument field value if set, zero value otherwise.
func (o *AnswerAllOf) GetSourceDocument() Document {
	if o == nil || IsNil(o.SourceDocument) {
		var ret Document
		return ret
	}
	return *o.SourceDocument
}

// GetSourceDocumentOk returns a tuple with the SourceDocument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnswerAllOf) GetSourceDocumentOk() (*Document, bool) {
	if o == nil || IsNil(o.SourceDocument) {
		return nil, false
	}
	return o.SourceDocument, true
}

// HasSourceDocument returns a boolean if a field has been set.
func (o *AnswerAllOf) HasSourceDocument() bool {
	if o != nil && !IsNil(o.SourceDocument) {
		return true
	}

	return false
}

// SetSourceDocument gets a reference to the given Document and assigns it to the SourceDocument field.
func (o *AnswerAllOf) SetSourceDocument(v Document) {
	o.SourceDocument = &v
}

func (o AnswerAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnswerAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CombinedAnswerText) {
		toSerialize["combinedAnswerText"] = o.CombinedAnswerText
	}
	if !IsNil(o.Likes) {
		toSerialize["likes"] = o.Likes
	}
	if !IsNil(o.UserRole) {
		toSerialize["userRole"] = o.UserRole
	}
	if !IsNil(o.Author) {
		toSerialize["author"] = o.Author
	}
	if !IsNil(o.CreateTime) {
		toSerialize["createTime"] = o.CreateTime
	}
	if !IsNil(o.UpdateTime) {
		toSerialize["updateTime"] = o.UpdateTime
	}
	if !IsNil(o.UpdatedBy) {
		toSerialize["updatedBy"] = o.UpdatedBy
	}
	if !IsNil(o.Verification) {
		toSerialize["verification"] = o.Verification
	}
	if !IsNil(o.Board) {
		toSerialize["board"] = o.Board
	}
	if !IsNil(o.Collections) {
		toSerialize["collections"] = o.Collections
	}
	if !IsNil(o.DocumentCategory) {
		toSerialize["documentCategory"] = o.DocumentCategory
	}
	if !IsNil(o.SourceDocument) {
		toSerialize["sourceDocument"] = o.SourceDocument
	}
	return toSerialize, nil
}

type NullableAnswerAllOf struct {
	value *AnswerAllOf
	isSet bool
}

func (v NullableAnswerAllOf) Get() *AnswerAllOf {
	return v.value
}

func (v *NullableAnswerAllOf) Set(val *AnswerAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAnswerAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAnswerAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnswerAllOf(val *AnswerAllOf) *NullableAnswerAllOf {
	return &NullableAnswerAllOf{value: val, isSet: true}
}

func (v NullableAnswerAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnswerAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


