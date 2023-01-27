# CheckDocumentAccessRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datasource** | **string** | Datasource of document to get check access for | 
**ObjectType** | **string** | Object type of document to get check access for | 
**DocId** | **string** | ID of document to get check access for | 
**UserEmail** | **string** | Email of user to check access for | 

## Methods

### NewCheckDocumentAccessRequest

`func NewCheckDocumentAccessRequest(datasource string, objectType string, docId string, userEmail string, ) *CheckDocumentAccessRequest`

NewCheckDocumentAccessRequest instantiates a new CheckDocumentAccessRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckDocumentAccessRequestWithDefaults

`func NewCheckDocumentAccessRequestWithDefaults() *CheckDocumentAccessRequest`

NewCheckDocumentAccessRequestWithDefaults instantiates a new CheckDocumentAccessRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatasource

`func (o *CheckDocumentAccessRequest) GetDatasource() string`

GetDatasource returns the Datasource field if non-nil, zero value otherwise.

### GetDatasourceOk

`func (o *CheckDocumentAccessRequest) GetDatasourceOk() (*string, bool)`

GetDatasourceOk returns a tuple with the Datasource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasource

`func (o *CheckDocumentAccessRequest) SetDatasource(v string)`

SetDatasource sets Datasource field to given value.


### GetObjectType

`func (o *CheckDocumentAccessRequest) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CheckDocumentAccessRequest) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CheckDocumentAccessRequest) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDocId

`func (o *CheckDocumentAccessRequest) GetDocId() string`

GetDocId returns the DocId field if non-nil, zero value otherwise.

### GetDocIdOk

`func (o *CheckDocumentAccessRequest) GetDocIdOk() (*string, bool)`

GetDocIdOk returns a tuple with the DocId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocId

`func (o *CheckDocumentAccessRequest) SetDocId(v string)`

SetDocId sets DocId field to given value.


### GetUserEmail

`func (o *CheckDocumentAccessRequest) GetUserEmail() string`

GetUserEmail returns the UserEmail field if non-nil, zero value otherwise.

### GetUserEmailOk

`func (o *CheckDocumentAccessRequest) GetUserEmailOk() (*string, bool)`

GetUserEmailOk returns a tuple with the UserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEmail

`func (o *CheckDocumentAccessRequest) SetUserEmail(v string)`

SetUserEmail sets UserEmail field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


