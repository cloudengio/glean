# GetExternalDocumentsProxyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Status of the request, i.e. OK or ERROR | [optional] 
**Version** | Pointer to **float32** | Version of the response | [optional] 
**Documents** | Pointer to [**[]ExternalDocumentProxy**](ExternalDocumentProxy.md) |  | [optional] 

## Methods

### NewGetExternalDocumentsProxyResponse

`func NewGetExternalDocumentsProxyResponse() *GetExternalDocumentsProxyResponse`

NewGetExternalDocumentsProxyResponse instantiates a new GetExternalDocumentsProxyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetExternalDocumentsProxyResponseWithDefaults

`func NewGetExternalDocumentsProxyResponseWithDefaults() *GetExternalDocumentsProxyResponse`

NewGetExternalDocumentsProxyResponseWithDefaults instantiates a new GetExternalDocumentsProxyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetExternalDocumentsProxyResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetExternalDocumentsProxyResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetExternalDocumentsProxyResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetExternalDocumentsProxyResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVersion

`func (o *GetExternalDocumentsProxyResponse) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetExternalDocumentsProxyResponse) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetExternalDocumentsProxyResponse) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *GetExternalDocumentsProxyResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetDocuments

`func (o *GetExternalDocumentsProxyResponse) GetDocuments() []ExternalDocumentProxy`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *GetExternalDocumentsProxyResponse) GetDocumentsOk() (*[]ExternalDocumentProxy, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *GetExternalDocumentsProxyResponse) SetDocuments(v []ExternalDocumentProxy)`

SetDocuments sets Documents field to given value.

### HasDocuments

`func (o *GetExternalDocumentsProxyResponse) HasDocuments() bool`

HasDocuments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


