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

// checks if the VerificationMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VerificationMetadata{}

// VerificationMetadata struct for VerificationMetadata
type VerificationMetadata struct {
	LastVerifier *Person `json:"lastVerifier,omitempty"`
	// The unix timestamp of the verification (in seconds since epoch UTC).
	LastVerificationTs *int32 `json:"lastVerificationTs,omitempty"`
	// The unix timestamp of the verification expiration if applicable (in seconds since epoch UTC).
	ExpirationTs *int32 `json:"expirationTs,omitempty"`
	Document *Document `json:"document,omitempty"`
	// Info about all outstanding verification reminders for the document if exists.
	Reminders []Reminder `json:"reminders,omitempty"`
	LastReminder *Reminder `json:"lastReminder,omitempty"`
	// Number of visitors to the document during included time periods.
	VisitorCount []CountInfo `json:"visitorCount,omitempty"`
	// List of potential verifiers for the document e.g. old verifiers and/or users with view/edit permissions.
	CandidateVerifiers []Person `json:"candidateVerifiers,omitempty"`
}

// NewVerificationMetadata instantiates a new VerificationMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerificationMetadata() *VerificationMetadata {
	this := VerificationMetadata{}
	return &this
}

// NewVerificationMetadataWithDefaults instantiates a new VerificationMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerificationMetadataWithDefaults() *VerificationMetadata {
	this := VerificationMetadata{}
	return &this
}

// GetLastVerifier returns the LastVerifier field value if set, zero value otherwise.
func (o *VerificationMetadata) GetLastVerifier() Person {
	if o == nil || IsNil(o.LastVerifier) {
		var ret Person
		return ret
	}
	return *o.LastVerifier
}

// GetLastVerifierOk returns a tuple with the LastVerifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerificationMetadata) GetLastVerifierOk() (*Person, bool) {
	if o == nil || IsNil(o.LastVerifier) {
		return nil, false
	}
	return o.LastVerifier, true
}

// HasLastVerifier returns a boolean if a field has been set.
func (o *VerificationMetadata) HasLastVerifier() bool {
	if o != nil && !IsNil(o.LastVerifier) {
		return true
	}

	return false
}

// SetLastVerifier gets a reference to the given Person and assigns it to the LastVerifier field.
func (o *VerificationMetadata) SetLastVerifier(v Person) {
	o.LastVerifier = &v
}

// GetLastVerificationTs returns the LastVerificationTs field value if set, zero value otherwise.
func (o *VerificationMetadata) GetLastVerificationTs() int32 {
	if o == nil || IsNil(o.LastVerificationTs) {
		var ret int32
		return ret
	}
	return *o.LastVerificationTs
}

// GetLastVerificationTsOk returns a tuple with the LastVerificationTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerificationMetadata) GetLastVerificationTsOk() (*int32, bool) {
	if o == nil || IsNil(o.LastVerificationTs) {
		return nil, false
	}
	return o.LastVerificationTs, true
}

// HasLastVerificationTs returns a boolean if a field has been set.
func (o *VerificationMetadata) HasLastVerificationTs() bool {
	if o != nil && !IsNil(o.LastVerificationTs) {
		return true
	}

	return false
}

// SetLastVerificationTs gets a reference to the given int32 and assigns it to the LastVerificationTs field.
func (o *VerificationMetadata) SetLastVerificationTs(v int32) {
	o.LastVerificationTs = &v
}

// GetExpirationTs returns the ExpirationTs field value if set, zero value otherwise.
func (o *VerificationMetadata) GetExpirationTs() int32 {
	if o == nil || IsNil(o.ExpirationTs) {
		var ret int32
		return ret
	}
	return *o.ExpirationTs
}

// GetExpirationTsOk returns a tuple with the ExpirationTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerificationMetadata) GetExpirationTsOk() (*int32, bool) {
	if o == nil || IsNil(o.ExpirationTs) {
		return nil, false
	}
	return o.ExpirationTs, true
}

// HasExpirationTs returns a boolean if a field has been set.
func (o *VerificationMetadata) HasExpirationTs() bool {
	if o != nil && !IsNil(o.ExpirationTs) {
		return true
	}

	return false
}

// SetExpirationTs gets a reference to the given int32 and assigns it to the ExpirationTs field.
func (o *VerificationMetadata) SetExpirationTs(v int32) {
	o.ExpirationTs = &v
}

// GetDocument returns the Document field value if set, zero value otherwise.
func (o *VerificationMetadata) GetDocument() Document {
	if o == nil || IsNil(o.Document) {
		var ret Document
		return ret
	}
	return *o.Document
}

// GetDocumentOk returns a tuple with the Document field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerificationMetadata) GetDocumentOk() (*Document, bool) {
	if o == nil || IsNil(o.Document) {
		return nil, false
	}
	return o.Document, true
}

// HasDocument returns a boolean if a field has been set.
func (o *VerificationMetadata) HasDocument() bool {
	if o != nil && !IsNil(o.Document) {
		return true
	}

	return false
}

// SetDocument gets a reference to the given Document and assigns it to the Document field.
func (o *VerificationMetadata) SetDocument(v Document) {
	o.Document = &v
}

// GetReminders returns the Reminders field value if set, zero value otherwise.
func (o *VerificationMetadata) GetReminders() []Reminder {
	if o == nil || IsNil(o.Reminders) {
		var ret []Reminder
		return ret
	}
	return o.Reminders
}

// GetRemindersOk returns a tuple with the Reminders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerificationMetadata) GetRemindersOk() ([]Reminder, bool) {
	if o == nil || IsNil(o.Reminders) {
		return nil, false
	}
	return o.Reminders, true
}

// HasReminders returns a boolean if a field has been set.
func (o *VerificationMetadata) HasReminders() bool {
	if o != nil && !IsNil(o.Reminders) {
		return true
	}

	return false
}

// SetReminders gets a reference to the given []Reminder and assigns it to the Reminders field.
func (o *VerificationMetadata) SetReminders(v []Reminder) {
	o.Reminders = v
}

// GetLastReminder returns the LastReminder field value if set, zero value otherwise.
func (o *VerificationMetadata) GetLastReminder() Reminder {
	if o == nil || IsNil(o.LastReminder) {
		var ret Reminder
		return ret
	}
	return *o.LastReminder
}

// GetLastReminderOk returns a tuple with the LastReminder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerificationMetadata) GetLastReminderOk() (*Reminder, bool) {
	if o == nil || IsNil(o.LastReminder) {
		return nil, false
	}
	return o.LastReminder, true
}

// HasLastReminder returns a boolean if a field has been set.
func (o *VerificationMetadata) HasLastReminder() bool {
	if o != nil && !IsNil(o.LastReminder) {
		return true
	}

	return false
}

// SetLastReminder gets a reference to the given Reminder and assigns it to the LastReminder field.
func (o *VerificationMetadata) SetLastReminder(v Reminder) {
	o.LastReminder = &v
}

// GetVisitorCount returns the VisitorCount field value if set, zero value otherwise.
func (o *VerificationMetadata) GetVisitorCount() []CountInfo {
	if o == nil || IsNil(o.VisitorCount) {
		var ret []CountInfo
		return ret
	}
	return o.VisitorCount
}

// GetVisitorCountOk returns a tuple with the VisitorCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerificationMetadata) GetVisitorCountOk() ([]CountInfo, bool) {
	if o == nil || IsNil(o.VisitorCount) {
		return nil, false
	}
	return o.VisitorCount, true
}

// HasVisitorCount returns a boolean if a field has been set.
func (o *VerificationMetadata) HasVisitorCount() bool {
	if o != nil && !IsNil(o.VisitorCount) {
		return true
	}

	return false
}

// SetVisitorCount gets a reference to the given []CountInfo and assigns it to the VisitorCount field.
func (o *VerificationMetadata) SetVisitorCount(v []CountInfo) {
	o.VisitorCount = v
}

// GetCandidateVerifiers returns the CandidateVerifiers field value if set, zero value otherwise.
func (o *VerificationMetadata) GetCandidateVerifiers() []Person {
	if o == nil || IsNil(o.CandidateVerifiers) {
		var ret []Person
		return ret
	}
	return o.CandidateVerifiers
}

// GetCandidateVerifiersOk returns a tuple with the CandidateVerifiers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerificationMetadata) GetCandidateVerifiersOk() ([]Person, bool) {
	if o == nil || IsNil(o.CandidateVerifiers) {
		return nil, false
	}
	return o.CandidateVerifiers, true
}

// HasCandidateVerifiers returns a boolean if a field has been set.
func (o *VerificationMetadata) HasCandidateVerifiers() bool {
	if o != nil && !IsNil(o.CandidateVerifiers) {
		return true
	}

	return false
}

// SetCandidateVerifiers gets a reference to the given []Person and assigns it to the CandidateVerifiers field.
func (o *VerificationMetadata) SetCandidateVerifiers(v []Person) {
	o.CandidateVerifiers = v
}

func (o VerificationMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VerificationMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LastVerifier) {
		toSerialize["lastVerifier"] = o.LastVerifier
	}
	if !IsNil(o.LastVerificationTs) {
		toSerialize["lastVerificationTs"] = o.LastVerificationTs
	}
	if !IsNil(o.ExpirationTs) {
		toSerialize["expirationTs"] = o.ExpirationTs
	}
	if !IsNil(o.Document) {
		toSerialize["document"] = o.Document
	}
	if !IsNil(o.Reminders) {
		toSerialize["reminders"] = o.Reminders
	}
	if !IsNil(o.LastReminder) {
		toSerialize["lastReminder"] = o.LastReminder
	}
	if !IsNil(o.VisitorCount) {
		toSerialize["visitorCount"] = o.VisitorCount
	}
	if !IsNil(o.CandidateVerifiers) {
		toSerialize["candidateVerifiers"] = o.CandidateVerifiers
	}
	return toSerialize, nil
}

type NullableVerificationMetadata struct {
	value *VerificationMetadata
	isSet bool
}

func (v NullableVerificationMetadata) Get() *VerificationMetadata {
	return v.value
}

func (v *NullableVerificationMetadata) Set(val *VerificationMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableVerificationMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableVerificationMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerificationMetadata(val *VerificationMetadata) *NullableVerificationMetadata {
	return &NullableVerificationMetadata{value: val, isSet: true}
}

func (v NullableVerificationMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerificationMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

