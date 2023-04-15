# ExternalDocumentsProxyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Status of the request, i.e. OK or ERROR | [optional] 
**Version** | Pointer to **float32** | Version of the response | [optional] 
**Documents** | Pointer to [**[]ExternalDocumentProxy**](ExternalDocumentProxy.md) |  | [optional] 

## Methods

### NewExternalDocumentsProxyResponse

`func NewExternalDocumentsProxyResponse() *ExternalDocumentsProxyResponse`

NewExternalDocumentsProxyResponse instantiates a new ExternalDocumentsProxyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalDocumentsProxyResponseWithDefaults

`func NewExternalDocumentsProxyResponseWithDefaults() *ExternalDocumentsProxyResponse`

NewExternalDocumentsProxyResponseWithDefaults instantiates a new ExternalDocumentsProxyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ExternalDocumentsProxyResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExternalDocumentsProxyResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExternalDocumentsProxyResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ExternalDocumentsProxyResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVersion

`func (o *ExternalDocumentsProxyResponse) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ExternalDocumentsProxyResponse) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ExternalDocumentsProxyResponse) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ExternalDocumentsProxyResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetDocuments

`func (o *ExternalDocumentsProxyResponse) GetDocuments() []ExternalDocumentProxy`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *ExternalDocumentsProxyResponse) GetDocumentsOk() (*[]ExternalDocumentProxy, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *ExternalDocumentsProxyResponse) SetDocuments(v []ExternalDocumentProxy)`

SetDocuments sets Documents field to given value.

### HasDocuments

`func (o *ExternalDocumentsProxyResponse) HasDocuments() bool`

HasDocuments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


