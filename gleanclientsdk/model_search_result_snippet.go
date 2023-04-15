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

// checks if the SearchResultSnippet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchResultSnippet{}

// SearchResultSnippet struct for SearchResultSnippet
type SearchResultSnippet struct {
	// A matching snippet from the document. Query term matches are marked by the unicode characters uE006 and uE007.
	Snippet string `json:"snippet"`
	// The mime type of the snippets, currently either text/plain or text/html.
	MimeType *string `json:"mimeType,omitempty"`
	// A matching snippet from the document with no highlights.
	Text *string `json:"text,omitempty"`
	// Used for sorting based off the snippet's location within all_snippetable_text
	SnippetTextOrdering *int32 `json:"snippetTextOrdering,omitempty"`
	// The bolded ranges within text.
	Ranges []TextRange `json:"ranges,omitempty"`
}

// NewSearchResultSnippet instantiates a new SearchResultSnippet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchResultSnippet(snippet string) *SearchResultSnippet {
	this := SearchResultSnippet{}
	this.Snippet = snippet
	return &this
}

// NewSearchResultSnippetWithDefaults instantiates a new SearchResultSnippet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchResultSnippetWithDefaults() *SearchResultSnippet {
	this := SearchResultSnippet{}
	return &this
}

// GetSnippet returns the Snippet field value
func (o *SearchResultSnippet) GetSnippet() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Snippet
}

// GetSnippetOk returns a tuple with the Snippet field value
// and a boolean to check if the value has been set.
func (o *SearchResultSnippet) GetSnippetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Snippet, true
}

// SetSnippet sets field value
func (o *SearchResultSnippet) SetSnippet(v string) {
	o.Snippet = v
}

// GetMimeType returns the MimeType field value if set, zero value otherwise.
func (o *SearchResultSnippet) GetMimeType() string {
	if o == nil || IsNil(o.MimeType) {
		var ret string
		return ret
	}
	return *o.MimeType
}

// GetMimeTypeOk returns a tuple with the MimeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResultSnippet) GetMimeTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MimeType) {
		return nil, false
	}
	return o.MimeType, true
}

// HasMimeType returns a boolean if a field has been set.
func (o *SearchResultSnippet) HasMimeType() bool {
	if o != nil && !IsNil(o.MimeType) {
		return true
	}

	return false
}

// SetMimeType gets a reference to the given string and assigns it to the MimeType field.
func (o *SearchResultSnippet) SetMimeType(v string) {
	o.MimeType = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *SearchResultSnippet) GetText() string {
	if o == nil || IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResultSnippet) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *SearchResultSnippet) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *SearchResultSnippet) SetText(v string) {
	o.Text = &v
}

// GetSnippetTextOrdering returns the SnippetTextOrdering field value if set, zero value otherwise.
func (o *SearchResultSnippet) GetSnippetTextOrdering() int32 {
	if o == nil || IsNil(o.SnippetTextOrdering) {
		var ret int32
		return ret
	}
	return *o.SnippetTextOrdering
}

// GetSnippetTextOrderingOk returns a tuple with the SnippetTextOrdering field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResultSnippet) GetSnippetTextOrderingOk() (*int32, bool) {
	if o == nil || IsNil(o.SnippetTextOrdering) {
		return nil, false
	}
	return o.SnippetTextOrdering, true
}

// HasSnippetTextOrdering returns a boolean if a field has been set.
func (o *SearchResultSnippet) HasSnippetTextOrdering() bool {
	if o != nil && !IsNil(o.SnippetTextOrdering) {
		return true
	}

	return false
}

// SetSnippetTextOrdering gets a reference to the given int32 and assigns it to the SnippetTextOrdering field.
func (o *SearchResultSnippet) SetSnippetTextOrdering(v int32) {
	o.SnippetTextOrdering = &v
}

// GetRanges returns the Ranges field value if set, zero value otherwise.
func (o *SearchResultSnippet) GetRanges() []TextRange {
	if o == nil || IsNil(o.Ranges) {
		var ret []TextRange
		return ret
	}
	return o.Ranges
}

// GetRangesOk returns a tuple with the Ranges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResultSnippet) GetRangesOk() ([]TextRange, bool) {
	if o == nil || IsNil(o.Ranges) {
		return nil, false
	}
	return o.Ranges, true
}

// HasRanges returns a boolean if a field has been set.
func (o *SearchResultSnippet) HasRanges() bool {
	if o != nil && !IsNil(o.Ranges) {
		return true
	}

	return false
}

// SetRanges gets a reference to the given []TextRange and assigns it to the Ranges field.
func (o *SearchResultSnippet) SetRanges(v []TextRange) {
	o.Ranges = v
}

func (o SearchResultSnippet) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchResultSnippet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["snippet"] = o.Snippet
	if !IsNil(o.MimeType) {
		toSerialize["mimeType"] = o.MimeType
	}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	if !IsNil(o.SnippetTextOrdering) {
		toSerialize["snippetTextOrdering"] = o.SnippetTextOrdering
	}
	if !IsNil(o.Ranges) {
		toSerialize["ranges"] = o.Ranges
	}
	return toSerialize, nil
}

type NullableSearchResultSnippet struct {
	value *SearchResultSnippet
	isSet bool
}

func (v NullableSearchResultSnippet) Get() *SearchResultSnippet {
	return v.value
}

func (v *NullableSearchResultSnippet) Set(val *SearchResultSnippet) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchResultSnippet) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchResultSnippet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchResultSnippet(val *SearchResultSnippet) *NullableSearchResultSnippet {
	return &NullableSearchResultSnippet{value: val, isSet: true}
}

func (v NullableSearchResultSnippet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchResultSnippet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


