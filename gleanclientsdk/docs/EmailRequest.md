# EmailRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailTemplate** | [**EmailTemplate**](EmailTemplate.md) |  | 
**Recipients** | Pointer to [**[]Person**](Person.md) | The people to send emails to | [optional] 
**RecipientFilters** | Pointer to [**PeopleFilters**](PeopleFilters.md) |  | [optional] 
**CompanyName** | Pointer to **string** | Name of the company. | [optional] 
**DatasourceInstance** | Pointer to **string** | The instance ID of the datasource (if any) | [optional] 
**Sender** | Pointer to **string** | DEPRECATED - Name of the person at the deployment that triggered this email. | [optional] 
**Senders** | Pointer to [**[]Person**](Person.md) | The people who triggered this email | [optional] 
**ComputedRecipientType** | Pointer to **string** | Computed list of recipients to send the email to. Joined against the recipients list. | [optional] 
**WebAppUrl** | Pointer to **string** | The URL of the client triggering the request, as received in the ClientConfig | [optional] 
**Documents** | Pointer to [**[]Document**](Document.md) | The documents this email request refers to | [optional] 
**Reasons** | Pointer to **[]string** | Reasons this email request was sent. Will be shown directly to end user. | [optional] 
**FeedbackPayload** | Pointer to [**EmailRequestFeedbackPayload**](EmailRequestFeedbackPayload.md) |  | [optional] 

## Methods

### NewEmailRequest

`func NewEmailRequest(emailTemplate EmailTemplate, ) *EmailRequest`

NewEmailRequest instantiates a new EmailRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailRequestWithDefaults

`func NewEmailRequestWithDefaults() *EmailRequest`

NewEmailRequestWithDefaults instantiates a new EmailRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailTemplate

`func (o *EmailRequest) GetEmailTemplate() EmailTemplate`

GetEmailTemplate returns the EmailTemplate field if non-nil, zero value otherwise.

### GetEmailTemplateOk

`func (o *EmailRequest) GetEmailTemplateOk() (*EmailTemplate, bool)`

GetEmailTemplateOk returns a tuple with the EmailTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailTemplate

`func (o *EmailRequest) SetEmailTemplate(v EmailTemplate)`

SetEmailTemplate sets EmailTemplate field to given value.


### GetRecipients

`func (o *EmailRequest) GetRecipients() []Person`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *EmailRequest) GetRecipientsOk() (*[]Person, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *EmailRequest) SetRecipients(v []Person)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *EmailRequest) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.

### GetRecipientFilters

`func (o *EmailRequest) GetRecipientFilters() PeopleFilters`

GetRecipientFilters returns the RecipientFilters field if non-nil, zero value otherwise.

### GetRecipientFiltersOk

`func (o *EmailRequest) GetRecipientFiltersOk() (*PeopleFilters, bool)`

GetRecipientFiltersOk returns a tuple with the RecipientFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientFilters

`func (o *EmailRequest) SetRecipientFilters(v PeopleFilters)`

SetRecipientFilters sets RecipientFilters field to given value.

### HasRecipientFilters

`func (o *EmailRequest) HasRecipientFilters() bool`

HasRecipientFilters returns a boolean if a field has been set.

### GetCompanyName

`func (o *EmailRequest) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *EmailRequest) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *EmailRequest) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *EmailRequest) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetDatasourceInstance

`func (o *EmailRequest) GetDatasourceInstance() string`

GetDatasourceInstance returns the DatasourceInstance field if non-nil, zero value otherwise.

### GetDatasourceInstanceOk

`func (o *EmailRequest) GetDatasourceInstanceOk() (*string, bool)`

GetDatasourceInstanceOk returns a tuple with the DatasourceInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasourceInstance

`func (o *EmailRequest) SetDatasourceInstance(v string)`

SetDatasourceInstance sets DatasourceInstance field to given value.

### HasDatasourceInstance

`func (o *EmailRequest) HasDatasourceInstance() bool`

HasDatasourceInstance returns a boolean if a field has been set.

### GetSender

`func (o *EmailRequest) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *EmailRequest) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *EmailRequest) SetSender(v string)`

SetSender sets Sender field to given value.

### HasSender

`func (o *EmailRequest) HasSender() bool`

HasSender returns a boolean if a field has been set.

### GetSenders

`func (o *EmailRequest) GetSenders() []Person`

GetSenders returns the Senders field if non-nil, zero value otherwise.

### GetSendersOk

`func (o *EmailRequest) GetSendersOk() (*[]Person, bool)`

GetSendersOk returns a tuple with the Senders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenders

`func (o *EmailRequest) SetSenders(v []Person)`

SetSenders sets Senders field to given value.

### HasSenders

`func (o *EmailRequest) HasSenders() bool`

HasSenders returns a boolean if a field has been set.

### GetComputedRecipientType

`func (o *EmailRequest) GetComputedRecipientType() string`

GetComputedRecipientType returns the ComputedRecipientType field if non-nil, zero value otherwise.

### GetComputedRecipientTypeOk

`func (o *EmailRequest) GetComputedRecipientTypeOk() (*string, bool)`

GetComputedRecipientTypeOk returns a tuple with the ComputedRecipientType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputedRecipientType

`func (o *EmailRequest) SetComputedRecipientType(v string)`

SetComputedRecipientType sets ComputedRecipientType field to given value.

### HasComputedRecipientType

`func (o *EmailRequest) HasComputedRecipientType() bool`

HasComputedRecipientType returns a boolean if a field has been set.

### GetWebAppUrl

`func (o *EmailRequest) GetWebAppUrl() string`

GetWebAppUrl returns the WebAppUrl field if non-nil, zero value otherwise.

### GetWebAppUrlOk

`func (o *EmailRequest) GetWebAppUrlOk() (*string, bool)`

GetWebAppUrlOk returns a tuple with the WebAppUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebAppUrl

`func (o *EmailRequest) SetWebAppUrl(v string)`

SetWebAppUrl sets WebAppUrl field to given value.

### HasWebAppUrl

`func (o *EmailRequest) HasWebAppUrl() bool`

HasWebAppUrl returns a boolean if a field has been set.

### GetDocuments

`func (o *EmailRequest) GetDocuments() []Document`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *EmailRequest) GetDocumentsOk() (*[]Document, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *EmailRequest) SetDocuments(v []Document)`

SetDocuments sets Documents field to given value.

### HasDocuments

`func (o *EmailRequest) HasDocuments() bool`

HasDocuments returns a boolean if a field has been set.

### GetReasons

`func (o *EmailRequest) GetReasons() []string`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *EmailRequest) GetReasonsOk() (*[]string, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *EmailRequest) SetReasons(v []string)`

SetReasons sets Reasons field to given value.

### HasReasons

`func (o *EmailRequest) HasReasons() bool`

HasReasons returns a boolean if a field has been set.

### GetFeedbackPayload

`func (o *EmailRequest) GetFeedbackPayload() EmailRequestFeedbackPayload`

GetFeedbackPayload returns the FeedbackPayload field if non-nil, zero value otherwise.

### GetFeedbackPayloadOk

`func (o *EmailRequest) GetFeedbackPayloadOk() (*EmailRequestFeedbackPayload, bool)`

GetFeedbackPayloadOk returns a tuple with the FeedbackPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedbackPayload

`func (o *EmailRequest) SetFeedbackPayload(v EmailRequestFeedbackPayload)`

SetFeedbackPayload sets FeedbackPayload field to given value.

### HasFeedbackPayload

`func (o *EmailRequest) HasFeedbackPayload() bool`

HasFeedbackPayload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


