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

// checks if the StructuredResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StructuredResult{}

// StructuredResult struct for StructuredResult
type StructuredResult struct {
	Snippets []SearchResultSnippet `json:"snippets,omitempty"`
	Person *Person `json:"person,omitempty"`
	Customer *Customer `json:"customer,omitempty"`
	Team *Team `json:"team,omitempty"`
	CustomEntity *CustomEntity `json:"customEntity,omitempty"`
	Answer *Answer `json:"answer,omitempty"`
	ExtractedQnA *ExtractedQnA `json:"extractedQnA,omitempty"`
	App *AppResult `json:"app,omitempty"`
	Collection *Collection `json:"collection,omitempty"`
	AnswerBoard *AnswerBoard `json:"answerBoard,omitempty"`
	Code *Code `json:"code,omitempty"`
	Shortcut *Shortcut `json:"shortcut,omitempty"`
	QuerySuggestions *QuerySuggestionList `json:"querySuggestions,omitempty"`
	// A list of documents related to this structured result.
	RelatedDocuments []RelatedDocuments `json:"relatedDocuments,omitempty"`
	// Debug details for this result if debug is enabled.
	DebugInfo *string `json:"debugInfo,omitempty"`
	// An opaque token that represents this particular result in this particular query. To be used for /feedback reporting.
	TrackingToken *string `json:"trackingToken,omitempty"`
	// The level of visual distinction that should be given to a result. HERO - A high-confidence result that should feature prominently on the page. PROMOTED - May not be the best result but should be given additional visual distinction. STANDARD - Should not be distinct from any other results. TODO: Deprecate and use prominence field only in SearchResult. 
	Prominence *string `json:"prominence,omitempty"`
	// Source context for this result. Possible values depend on the result type.
	Source *string `json:"source,omitempty"`
}

// NewStructuredResult instantiates a new StructuredResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStructuredResult() *StructuredResult {
	this := StructuredResult{}
	return &this
}

// NewStructuredResultWithDefaults instantiates a new StructuredResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStructuredResultWithDefaults() *StructuredResult {
	this := StructuredResult{}
	return &this
}

// GetSnippets returns the Snippets field value if set, zero value otherwise.
func (o *StructuredResult) GetSnippets() []SearchResultSnippet {
	if o == nil || IsNil(o.Snippets) {
		var ret []SearchResultSnippet
		return ret
	}
	return o.Snippets
}

// GetSnippetsOk returns a tuple with the Snippets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuredResult) GetSnippetsOk() ([]SearchResultSnippet, bool) {
	if o == nil || IsNil(o.Snippets) {
		return nil, false
	}
	return o.Snippets, true
}

// HasSnippets returns a boolean if a field has been set.
func (o *StructuredResult) HasSnippets() bool {
	if o != nil && !IsNil(o.Snippets) {
		return true
	}

	return false
}

// SetSnippets gets a reference to the given []SearchResultSnippet and assigns it to the Snippets field.
func (o *StructuredResult) SetSnippets(v []SearchResultSnippet) {
	o.Snippets = v
}

// GetPerson returns the Person field value if set, zero value otherwise.
func (o *StructuredResult) GetPerson() Person {
	if o == nil || IsNil(o.Person) {
		var ret Person
		return ret
	}
	return *o.Person
}

// GetPersonOk returns a tuple with the Person field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuredResult) GetPersonOk() (*Person, bool) {
	if o == nil || IsNil(o.Person) {
		return nil, false
	}
	return o.Person, true
}

// HasPerson returns a boolean if a field has been set.
func (o *StructuredResult) HasPerson() bool {
	if o != nil && !IsNil(o.Person) {
		return true
	}

	return false
}

// SetPerson gets a reference to the given Person and assigns it to the Person field.
func (o *StructuredResult) SetPerson(v Person) {
	o.Person = &v
}

// GetCustomer returns the Customer field value if set, zero value otherwise.
func (o *StructuredResult) GetCustomer() Customer {
	if o == nil || IsNil(o.Customer) {
		var ret Customer
		return ret
	}
	return *o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuredResult) GetCustomerOk() (*Customer, bool) {
	if o == nil || IsNil(o.Customer) {
		return nil, false
	}
	return o.Customer, true
}

// HasCustomer returns a boolean if a field has been set.
func (o *StructuredResult) HasCustomer() bool {
	if o != nil && !IsNil(o.Customer) {
		return true
	}

	return false
}

// SetCustomer gets a reference to the given Customer and assigns it to the Customer field.
func (o *StructuredResult) SetCustomer(v Customer) {
	o.Customer = &v
}

// GetTeam returns the Team field value if set, zero value otherwise.
func (o *StructuredResult) GetTeam() Team {
	if o == nil || IsNil(o.Team) {
		var ret Team
		return ret
	}
	return *o.Team
}

// GetTeamOk returns a tuple with the Team field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuredResult) GetTeamOk() (*Team, bool) {
	if o == nil || IsNil(o.Team) {
		return nil, false
	}
	return o.Team, true
}

// HasTeam returns a boolean if a field has been set.
func (o *StructuredResult) HasTeam() bool {
	if o != nil && !IsNil(o.Team) {
		return true
	}

	return false
}

// SetTeam gets a reference to the given Team and assigns it to the Team field.
func (o *StructuredResult) SetTeam(v Team) {
	o.Team = &v
}

// GetCustomEntity returns the CustomEntity field value if set, zero value otherwise.
func (o *StructuredResult) GetCustomEntity() CustomEntity {
	if o == nil || IsNil(o.CustomEntity) {
		var ret CustomEntity
		return ret
	}
	return *o.CustomEntity
}

// GetCustomEntityOk returns a tuple with the CustomEntity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuredResult) GetCustomEntityOk() (*CustomEntity, bool) {
	if o == nil || IsNil(o.CustomEntity) {
		return nil, false
	}
	return o.CustomEntity, true
}

// HasCustomEntity returns a boolean if a field has been set.
func (o *StructuredResult) HasCustomEntity() bool {
	if o != nil && !IsNil(o.CustomEntity) {
		return true
	}

	return false
}

// SetCustomEntity gets a reference to the given CustomEntity and assigns it to the CustomEntity field.
func (o *StructuredResult) SetCustomEntity(v CustomEntity) {
	o.CustomEntity = &v
}

// GetAnswer returns the Answer field value if set, zero value otherwise.
func (o *StructuredResult) GetAnswer() Answer {
	if o == nil || IsNil(o.Answer) {
		var ret Answer
		return ret
	}
	return *o.Answer
}

// GetAnswerOk returns a tuple with the Answer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuredResult) GetAnswerOk() (*Answer, bool) {
	if o == nil || IsNil(o.Answer) {
		return nil, false
	}
	return o.Answer, true
}

// HasAnswer returns a boolean if a field has been set.
func (o *StructuredResult) HasAnswer() bool {
	if o != nil && !IsNil(o.Answer) {
		return true
	}

	return false
}

// SetAnswer gets a reference to the given Answer and assigns it to the Answer field.
func (o *StructuredResult) SetAnswer(v Answer) {
	o.Answer = &v
}

// GetExtractedQnA returns the ExtractedQnA field value if set, zero value otherwise.
func (o *StructuredResult) GetExtractedQnA() ExtractedQnA {
	if o == nil || IsNil(o.ExtractedQnA) {
		var ret ExtractedQnA
		return ret
	}
	return *o.ExtractedQnA
}

// GetExtractedQnAOk returns a tuple with the ExtractedQnA field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuredResult) GetExtractedQnAOk() (*ExtractedQnA, bool) {
	if o == nil || IsNil(o.ExtractedQnA) {
		return nil, false
	}
	return o.ExtractedQnA, true
}

// HasExtractedQnA returns a boolean if a field has been set.
func (o *StructuredResult) HasExtractedQnA() bool {
	if o != nil && !IsNil(o.ExtractedQnA) {
		return true
	}

	return false
}

// SetExtractedQnA gets a reference to the given ExtractedQnA and assigns it to the ExtractedQnA field.
func (o *StructuredResult) SetExtractedQnA(v ExtractedQnA) {
	o.ExtractedQnA = &v
}

// GetApp returns the App field value if set, zero value otherwise.
func (o *StructuredResult) GetApp() AppResult {
	if o == nil || IsNil(o.App) {
		var ret AppResult
		return ret
	}
	return *o.App
}

// GetAppOk returns a tuple with the App field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuredResult) GetAppOk() (*AppResult, bool) {
	if o == nil || IsNil(o.App) {
		return nil, false
	}
	return o.App, true
}

// HasApp returns a boolean if a field has been set.
func (o *StructuredResult) HasApp() bool {
	if o != nil && !IsNil(o.App) {
		return true
	}

	return false
}

// SetApp gets a reference to the given AppResult and assigns it to the App field.
func (o *StructuredResult) SetApp(v AppResult) {
	o.App = &v
}

// GetCollection returns the Collection field value if set, zero value otherwise.
func (o *StructuredResult) GetCollection() Collection {
	if o == nil || IsNil(o.Collection) {
		var ret Collection
		return ret
	}
	return *o.Collection
}

// GetCollectionOk returns a tuple with the Collection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuredResult) GetCollectionOk() (*Collection, bool) {
	if o == nil || IsNil(o.Collection) {
		return nil, false
	}
	return o.Collection, true
}

// HasCollection returns a boolean if a field has been set.
func (o *StructuredResult) HasCollection() bool {
	if o != nil && !IsNil(o.Collection) {
		return true
	}

	return false
}

// SetCollection gets a reference to the given Collection and assigns it to the Collection field.
func (o *StructuredResult) SetCollection(v Collection) {
	o.Collection = &v
}

// GetAnswerBoard returns the AnswerBoard field value if set, zero value otherwise.
func (o *StructuredResult) GetAnswerBoard() AnswerBoard {
	if o == nil || IsNil(o.AnswerBoard) {
		var ret AnswerBoard
		return ret
	}
	return *o.AnswerBoard
}

// GetAnswerBoardOk returns a tuple with the AnswerBoard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuredResult) GetAnswerBoardOk() (*AnswerBoard, bool) {
	if o == nil || IsNil(o.AnswerBoard) {
		return nil, false
	}
	return o.AnswerBoard, true
}

// HasAnswerBoard returns a boolean if a field has been set.
func (o *StructuredResult) HasAnswerBoard() bool {
	if o != nil && !IsNil(o.AnswerBoard) {
		return true
	}

	return false
}

// SetAnswerBoard gets a reference to the given AnswerBoard and assigns it to the AnswerBoard field.
func (o *StructuredResult) SetAnswerBoard(v AnswerBoard) {
	o.AnswerBoard = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *StructuredResult) GetCode() Code {
	if o == nil || IsNil(o.Code) {
		var ret Code
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuredResult) GetCodeOk() (*Code, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *StructuredResult) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given Code and assigns it to the Code field.
func (o *StructuredResult) SetCode(v Code) {
	o.Code = &v
}

// GetShortcut returns the Shortcut field value if set, zero value otherwise.
func (o *StructuredResult) GetShortcut() Shortcut {
	if o == nil || IsNil(o.Shortcut) {
		var ret Shortcut
		return ret
	}
	return *o.Shortcut
}

// GetShortcutOk returns a tuple with the Shortcut field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuredResult) GetShortcutOk() (*Shortcut, bool) {
	if o == nil || IsNil(o.Shortcut) {
		return nil, false
	}
	return o.Shortcut, true
}

// HasShortcut returns a boolean if a field has been set.
func (o *StructuredResult) HasShortcut() bool {
	if o != nil && !IsNil(o.Shortcut) {
		return true
	}

	return false
}

// SetShortcut gets a reference to the given Shortcut and assigns it to the Shortcut field.
func (o *StructuredResult) SetShortcut(v Shortcut) {
	o.Shortcut = &v
}

// GetQuerySuggestions returns the QuerySuggestions field value if set, zero value otherwise.
func (o *StructuredResult) GetQuerySuggestions() QuerySuggestionList {
	if o == nil || IsNil(o.QuerySuggestions) {
		var ret QuerySuggestionList
		return ret
	}
	return *o.QuerySuggestions
}

// GetQuerySuggestionsOk returns a tuple with the QuerySuggestions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuredResult) GetQuerySuggestionsOk() (*QuerySuggestionList, bool) {
	if o == nil || IsNil(o.QuerySuggestions) {
		return nil, false
	}
	return o.QuerySuggestions, true
}

// HasQuerySuggestions returns a boolean if a field has been set.
func (o *StructuredResult) HasQuerySuggestions() bool {
	if o != nil && !IsNil(o.QuerySuggestions) {
		return true
	}

	return false
}

// SetQuerySuggestions gets a reference to the given QuerySuggestionList and assigns it to the QuerySuggestions field.
func (o *StructuredResult) SetQuerySuggestions(v QuerySuggestionList) {
	o.QuerySuggestions = &v
}

// GetRelatedDocuments returns the RelatedDocuments field value if set, zero value otherwise.
func (o *StructuredResult) GetRelatedDocuments() []RelatedDocuments {
	if o == nil || IsNil(o.RelatedDocuments) {
		var ret []RelatedDocuments
		return ret
	}
	return o.RelatedDocuments
}

// GetRelatedDocumentsOk returns a tuple with the RelatedDocuments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuredResult) GetRelatedDocumentsOk() ([]RelatedDocuments, bool) {
	if o == nil || IsNil(o.RelatedDocuments) {
		return nil, false
	}
	return o.RelatedDocuments, true
}

// HasRelatedDocuments returns a boolean if a field has been set.
func (o *StructuredResult) HasRelatedDocuments() bool {
	if o != nil && !IsNil(o.RelatedDocuments) {
		return true
	}

	return false
}

// SetRelatedDocuments gets a reference to the given []RelatedDocuments and assigns it to the RelatedDocuments field.
func (o *StructuredResult) SetRelatedDocuments(v []RelatedDocuments) {
	o.RelatedDocuments = v
}

// GetDebugInfo returns the DebugInfo field value if set, zero value otherwise.
func (o *StructuredResult) GetDebugInfo() string {
	if o == nil || IsNil(o.DebugInfo) {
		var ret string
		return ret
	}
	return *o.DebugInfo
}

// GetDebugInfoOk returns a tuple with the DebugInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuredResult) GetDebugInfoOk() (*string, bool) {
	if o == nil || IsNil(o.DebugInfo) {
		return nil, false
	}
	return o.DebugInfo, true
}

// HasDebugInfo returns a boolean if a field has been set.
func (o *StructuredResult) HasDebugInfo() bool {
	if o != nil && !IsNil(o.DebugInfo) {
		return true
	}

	return false
}

// SetDebugInfo gets a reference to the given string and assigns it to the DebugInfo field.
func (o *StructuredResult) SetDebugInfo(v string) {
	o.DebugInfo = &v
}

// GetTrackingToken returns the TrackingToken field value if set, zero value otherwise.
func (o *StructuredResult) GetTrackingToken() string {
	if o == nil || IsNil(o.TrackingToken) {
		var ret string
		return ret
	}
	return *o.TrackingToken
}

// GetTrackingTokenOk returns a tuple with the TrackingToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuredResult) GetTrackingTokenOk() (*string, bool) {
	if o == nil || IsNil(o.TrackingToken) {
		return nil, false
	}
	return o.TrackingToken, true
}

// HasTrackingToken returns a boolean if a field has been set.
func (o *StructuredResult) HasTrackingToken() bool {
	if o != nil && !IsNil(o.TrackingToken) {
		return true
	}

	return false
}

// SetTrackingToken gets a reference to the given string and assigns it to the TrackingToken field.
func (o *StructuredResult) SetTrackingToken(v string) {
	o.TrackingToken = &v
}

// GetProminence returns the Prominence field value if set, zero value otherwise.
func (o *StructuredResult) GetProminence() string {
	if o == nil || IsNil(o.Prominence) {
		var ret string
		return ret
	}
	return *o.Prominence
}

// GetProminenceOk returns a tuple with the Prominence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuredResult) GetProminenceOk() (*string, bool) {
	if o == nil || IsNil(o.Prominence) {
		return nil, false
	}
	return o.Prominence, true
}

// HasProminence returns a boolean if a field has been set.
func (o *StructuredResult) HasProminence() bool {
	if o != nil && !IsNil(o.Prominence) {
		return true
	}

	return false
}

// SetProminence gets a reference to the given string and assigns it to the Prominence field.
func (o *StructuredResult) SetProminence(v string) {
	o.Prominence = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *StructuredResult) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuredResult) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *StructuredResult) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *StructuredResult) SetSource(v string) {
	o.Source = &v
}

func (o StructuredResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StructuredResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Snippets) {
		toSerialize["snippets"] = o.Snippets
	}
	if !IsNil(o.Person) {
		toSerialize["person"] = o.Person
	}
	if !IsNil(o.Customer) {
		toSerialize["customer"] = o.Customer
	}
	if !IsNil(o.Team) {
		toSerialize["team"] = o.Team
	}
	if !IsNil(o.CustomEntity) {
		toSerialize["customEntity"] = o.CustomEntity
	}
	if !IsNil(o.Answer) {
		toSerialize["answer"] = o.Answer
	}
	if !IsNil(o.ExtractedQnA) {
		toSerialize["extractedQnA"] = o.ExtractedQnA
	}
	if !IsNil(o.App) {
		toSerialize["app"] = o.App
	}
	if !IsNil(o.Collection) {
		toSerialize["collection"] = o.Collection
	}
	if !IsNil(o.AnswerBoard) {
		toSerialize["answerBoard"] = o.AnswerBoard
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Shortcut) {
		toSerialize["shortcut"] = o.Shortcut
	}
	if !IsNil(o.QuerySuggestions) {
		toSerialize["querySuggestions"] = o.QuerySuggestions
	}
	if !IsNil(o.RelatedDocuments) {
		toSerialize["relatedDocuments"] = o.RelatedDocuments
	}
	if !IsNil(o.DebugInfo) {
		toSerialize["debugInfo"] = o.DebugInfo
	}
	if !IsNil(o.TrackingToken) {
		toSerialize["trackingToken"] = o.TrackingToken
	}
	if !IsNil(o.Prominence) {
		toSerialize["prominence"] = o.Prominence
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	return toSerialize, nil
}

type NullableStructuredResult struct {
	value *StructuredResult
	isSet bool
}

func (v NullableStructuredResult) Get() *StructuredResult {
	return v.value
}

func (v *NullableStructuredResult) Set(val *StructuredResult) {
	v.value = val
	v.isSet = true
}

func (v NullableStructuredResult) IsSet() bool {
	return v.isSet
}

func (v *NullableStructuredResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStructuredResult(val *StructuredResult) *NullableStructuredResult {
	return &NullableStructuredResult{value: val, isSet: true}
}

func (v NullableStructuredResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStructuredResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


