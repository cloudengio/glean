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

// checks if the EmailRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailRequest{}

// EmailRequest A request to send email[s] to the specified users
type EmailRequest struct {
	EmailTemplate EmailTemplate `json:"emailTemplate"`
	// The people to send emails to
	Recipients []Person `json:"recipients,omitempty"`
	RecipientFilters *PeopleFilters `json:"recipientFilters,omitempty"`
	// Name of the company.
	CompanyName *string `json:"companyName,omitempty"`
	// The instance ID of the datasource (if any)
	DatasourceInstance *string `json:"datasourceInstance,omitempty"`
	// DEPRECATED - Name of the person at the deployment that triggered this email.
	// Deprecated
	Sender *string `json:"sender,omitempty"`
	// The people who triggered this email
	Senders []Person `json:"senders,omitempty"`
	// Computed list of recipients to send the email to. Joined against the recipients list.
	ComputedRecipientType *string `json:"computedRecipientType,omitempty"`
	// The URL of the client triggering the request, as received in the ClientConfig
	WebAppUrl *string `json:"webAppUrl,omitempty"`
	// The documents this email request refers to
	Documents []Document `json:"documents,omitempty"`
	// Reasons this email request was sent. Will be shown directly to end user.
	Reasons []string `json:"reasons,omitempty"`
	FeedbackPayload *EmailRequestFeedbackPayload `json:"feedbackPayload,omitempty"`
}

// NewEmailRequest instantiates a new EmailRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailRequest(emailTemplate EmailTemplate) *EmailRequest {
	this := EmailRequest{}
	this.EmailTemplate = emailTemplate
	return &this
}

// NewEmailRequestWithDefaults instantiates a new EmailRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailRequestWithDefaults() *EmailRequest {
	this := EmailRequest{}
	return &this
}

// GetEmailTemplate returns the EmailTemplate field value
func (o *EmailRequest) GetEmailTemplate() EmailTemplate {
	if o == nil {
		var ret EmailTemplate
		return ret
	}

	return o.EmailTemplate
}

// GetEmailTemplateOk returns a tuple with the EmailTemplate field value
// and a boolean to check if the value has been set.
func (o *EmailRequest) GetEmailTemplateOk() (*EmailTemplate, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EmailTemplate, true
}

// SetEmailTemplate sets field value
func (o *EmailRequest) SetEmailTemplate(v EmailTemplate) {
	o.EmailTemplate = v
}

// GetRecipients returns the Recipients field value if set, zero value otherwise.
func (o *EmailRequest) GetRecipients() []Person {
	if o == nil || IsNil(o.Recipients) {
		var ret []Person
		return ret
	}
	return o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailRequest) GetRecipientsOk() ([]Person, bool) {
	if o == nil || IsNil(o.Recipients) {
		return nil, false
	}
	return o.Recipients, true
}

// HasRecipients returns a boolean if a field has been set.
func (o *EmailRequest) HasRecipients() bool {
	if o != nil && !IsNil(o.Recipients) {
		return true
	}

	return false
}

// SetRecipients gets a reference to the given []Person and assigns it to the Recipients field.
func (o *EmailRequest) SetRecipients(v []Person) {
	o.Recipients = v
}

// GetRecipientFilters returns the RecipientFilters field value if set, zero value otherwise.
func (o *EmailRequest) GetRecipientFilters() PeopleFilters {
	if o == nil || IsNil(o.RecipientFilters) {
		var ret PeopleFilters
		return ret
	}
	return *o.RecipientFilters
}

// GetRecipientFiltersOk returns a tuple with the RecipientFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailRequest) GetRecipientFiltersOk() (*PeopleFilters, bool) {
	if o == nil || IsNil(o.RecipientFilters) {
		return nil, false
	}
	return o.RecipientFilters, true
}

// HasRecipientFilters returns a boolean if a field has been set.
func (o *EmailRequest) HasRecipientFilters() bool {
	if o != nil && !IsNil(o.RecipientFilters) {
		return true
	}

	return false
}

// SetRecipientFilters gets a reference to the given PeopleFilters and assigns it to the RecipientFilters field.
func (o *EmailRequest) SetRecipientFilters(v PeopleFilters) {
	o.RecipientFilters = &v
}

// GetCompanyName returns the CompanyName field value if set, zero value otherwise.
func (o *EmailRequest) GetCompanyName() string {
	if o == nil || IsNil(o.CompanyName) {
		var ret string
		return ret
	}
	return *o.CompanyName
}

// GetCompanyNameOk returns a tuple with the CompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailRequest) GetCompanyNameOk() (*string, bool) {
	if o == nil || IsNil(o.CompanyName) {
		return nil, false
	}
	return o.CompanyName, true
}

// HasCompanyName returns a boolean if a field has been set.
func (o *EmailRequest) HasCompanyName() bool {
	if o != nil && !IsNil(o.CompanyName) {
		return true
	}

	return false
}

// SetCompanyName gets a reference to the given string and assigns it to the CompanyName field.
func (o *EmailRequest) SetCompanyName(v string) {
	o.CompanyName = &v
}

// GetDatasourceInstance returns the DatasourceInstance field value if set, zero value otherwise.
func (o *EmailRequest) GetDatasourceInstance() string {
	if o == nil || IsNil(o.DatasourceInstance) {
		var ret string
		return ret
	}
	return *o.DatasourceInstance
}

// GetDatasourceInstanceOk returns a tuple with the DatasourceInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailRequest) GetDatasourceInstanceOk() (*string, bool) {
	if o == nil || IsNil(o.DatasourceInstance) {
		return nil, false
	}
	return o.DatasourceInstance, true
}

// HasDatasourceInstance returns a boolean if a field has been set.
func (o *EmailRequest) HasDatasourceInstance() bool {
	if o != nil && !IsNil(o.DatasourceInstance) {
		return true
	}

	return false
}

// SetDatasourceInstance gets a reference to the given string and assigns it to the DatasourceInstance field.
func (o *EmailRequest) SetDatasourceInstance(v string) {
	o.DatasourceInstance = &v
}

// GetSender returns the Sender field value if set, zero value otherwise.
// Deprecated
func (o *EmailRequest) GetSender() string {
	if o == nil || IsNil(o.Sender) {
		var ret string
		return ret
	}
	return *o.Sender
}

// GetSenderOk returns a tuple with the Sender field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *EmailRequest) GetSenderOk() (*string, bool) {
	if o == nil || IsNil(o.Sender) {
		return nil, false
	}
	return o.Sender, true
}

// HasSender returns a boolean if a field has been set.
func (o *EmailRequest) HasSender() bool {
	if o != nil && !IsNil(o.Sender) {
		return true
	}

	return false
}

// SetSender gets a reference to the given string and assigns it to the Sender field.
// Deprecated
func (o *EmailRequest) SetSender(v string) {
	o.Sender = &v
}

// GetSenders returns the Senders field value if set, zero value otherwise.
func (o *EmailRequest) GetSenders() []Person {
	if o == nil || IsNil(o.Senders) {
		var ret []Person
		return ret
	}
	return o.Senders
}

// GetSendersOk returns a tuple with the Senders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailRequest) GetSendersOk() ([]Person, bool) {
	if o == nil || IsNil(o.Senders) {
		return nil, false
	}
	return o.Senders, true
}

// HasSenders returns a boolean if a field has been set.
func (o *EmailRequest) HasSenders() bool {
	if o != nil && !IsNil(o.Senders) {
		return true
	}

	return false
}

// SetSenders gets a reference to the given []Person and assigns it to the Senders field.
func (o *EmailRequest) SetSenders(v []Person) {
	o.Senders = v
}

// GetComputedRecipientType returns the ComputedRecipientType field value if set, zero value otherwise.
func (o *EmailRequest) GetComputedRecipientType() string {
	if o == nil || IsNil(o.ComputedRecipientType) {
		var ret string
		return ret
	}
	return *o.ComputedRecipientType
}

// GetComputedRecipientTypeOk returns a tuple with the ComputedRecipientType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailRequest) GetComputedRecipientTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ComputedRecipientType) {
		return nil, false
	}
	return o.ComputedRecipientType, true
}

// HasComputedRecipientType returns a boolean if a field has been set.
func (o *EmailRequest) HasComputedRecipientType() bool {
	if o != nil && !IsNil(o.ComputedRecipientType) {
		return true
	}

	return false
}

// SetComputedRecipientType gets a reference to the given string and assigns it to the ComputedRecipientType field.
func (o *EmailRequest) SetComputedRecipientType(v string) {
	o.ComputedRecipientType = &v
}

// GetWebAppUrl returns the WebAppUrl field value if set, zero value otherwise.
func (o *EmailRequest) GetWebAppUrl() string {
	if o == nil || IsNil(o.WebAppUrl) {
		var ret string
		return ret
	}
	return *o.WebAppUrl
}

// GetWebAppUrlOk returns a tuple with the WebAppUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailRequest) GetWebAppUrlOk() (*string, bool) {
	if o == nil || IsNil(o.WebAppUrl) {
		return nil, false
	}
	return o.WebAppUrl, true
}

// HasWebAppUrl returns a boolean if a field has been set.
func (o *EmailRequest) HasWebAppUrl() bool {
	if o != nil && !IsNil(o.WebAppUrl) {
		return true
	}

	return false
}

// SetWebAppUrl gets a reference to the given string and assigns it to the WebAppUrl field.
func (o *EmailRequest) SetWebAppUrl(v string) {
	o.WebAppUrl = &v
}

// GetDocuments returns the Documents field value if set, zero value otherwise.
func (o *EmailRequest) GetDocuments() []Document {
	if o == nil || IsNil(o.Documents) {
		var ret []Document
		return ret
	}
	return o.Documents
}

// GetDocumentsOk returns a tuple with the Documents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailRequest) GetDocumentsOk() ([]Document, bool) {
	if o == nil || IsNil(o.Documents) {
		return nil, false
	}
	return o.Documents, true
}

// HasDocuments returns a boolean if a field has been set.
func (o *EmailRequest) HasDocuments() bool {
	if o != nil && !IsNil(o.Documents) {
		return true
	}

	return false
}

// SetDocuments gets a reference to the given []Document and assigns it to the Documents field.
func (o *EmailRequest) SetDocuments(v []Document) {
	o.Documents = v
}

// GetReasons returns the Reasons field value if set, zero value otherwise.
func (o *EmailRequest) GetReasons() []string {
	if o == nil || IsNil(o.Reasons) {
		var ret []string
		return ret
	}
	return o.Reasons
}

// GetReasonsOk returns a tuple with the Reasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailRequest) GetReasonsOk() ([]string, bool) {
	if o == nil || IsNil(o.Reasons) {
		return nil, false
	}
	return o.Reasons, true
}

// HasReasons returns a boolean if a field has been set.
func (o *EmailRequest) HasReasons() bool {
	if o != nil && !IsNil(o.Reasons) {
		return true
	}

	return false
}

// SetReasons gets a reference to the given []string and assigns it to the Reasons field.
func (o *EmailRequest) SetReasons(v []string) {
	o.Reasons = v
}

// GetFeedbackPayload returns the FeedbackPayload field value if set, zero value otherwise.
func (o *EmailRequest) GetFeedbackPayload() EmailRequestFeedbackPayload {
	if o == nil || IsNil(o.FeedbackPayload) {
		var ret EmailRequestFeedbackPayload
		return ret
	}
	return *o.FeedbackPayload
}

// GetFeedbackPayloadOk returns a tuple with the FeedbackPayload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailRequest) GetFeedbackPayloadOk() (*EmailRequestFeedbackPayload, bool) {
	if o == nil || IsNil(o.FeedbackPayload) {
		return nil, false
	}
	return o.FeedbackPayload, true
}

// HasFeedbackPayload returns a boolean if a field has been set.
func (o *EmailRequest) HasFeedbackPayload() bool {
	if o != nil && !IsNil(o.FeedbackPayload) {
		return true
	}

	return false
}

// SetFeedbackPayload gets a reference to the given EmailRequestFeedbackPayload and assigns it to the FeedbackPayload field.
func (o *EmailRequest) SetFeedbackPayload(v EmailRequestFeedbackPayload) {
	o.FeedbackPayload = &v
}

func (o EmailRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["emailTemplate"] = o.EmailTemplate
	if !IsNil(o.Recipients) {
		toSerialize["recipients"] = o.Recipients
	}
	if !IsNil(o.RecipientFilters) {
		toSerialize["recipientFilters"] = o.RecipientFilters
	}
	if !IsNil(o.CompanyName) {
		toSerialize["companyName"] = o.CompanyName
	}
	if !IsNil(o.DatasourceInstance) {
		toSerialize["datasourceInstance"] = o.DatasourceInstance
	}
	if !IsNil(o.Sender) {
		toSerialize["sender"] = o.Sender
	}
	if !IsNil(o.Senders) {
		toSerialize["senders"] = o.Senders
	}
	if !IsNil(o.ComputedRecipientType) {
		toSerialize["computedRecipientType"] = o.ComputedRecipientType
	}
	if !IsNil(o.WebAppUrl) {
		toSerialize["webAppUrl"] = o.WebAppUrl
	}
	if !IsNil(o.Documents) {
		toSerialize["documents"] = o.Documents
	}
	if !IsNil(o.Reasons) {
		toSerialize["reasons"] = o.Reasons
	}
	if !IsNil(o.FeedbackPayload) {
		toSerialize["feedbackPayload"] = o.FeedbackPayload
	}
	return toSerialize, nil
}

type NullableEmailRequest struct {
	value *EmailRequest
	isSet bool
}

func (v NullableEmailRequest) Get() *EmailRequest {
	return v.value
}

func (v *NullableEmailRequest) Set(val *EmailRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailRequest(val *EmailRequest) *NullableEmailRequest {
	return &NullableEmailRequest{value: val, isSet: true}
}

func (v NullableEmailRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


