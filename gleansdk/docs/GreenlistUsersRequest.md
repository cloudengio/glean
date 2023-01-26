# GreenlistUsersRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datasource** | **string** | Datasource which needs to be made visible to users specified in the &#x60;emails&#x60; field. | 
**Emails** | **[]string** | The emails of the beta users | 

## Methods

### NewGreenlistUsersRequest

`func NewGreenlistUsersRequest(datasource string, emails []string, ) *GreenlistUsersRequest`

NewGreenlistUsersRequest instantiates a new GreenlistUsersRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGreenlistUsersRequestWithDefaults

`func NewGreenlistUsersRequestWithDefaults() *GreenlistUsersRequest`

NewGreenlistUsersRequestWithDefaults instantiates a new GreenlistUsersRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatasource

`func (o *GreenlistUsersRequest) GetDatasource() string`

GetDatasource returns the Datasource field if non-nil, zero value otherwise.

### GetDatasourceOk

`func (o *GreenlistUsersRequest) GetDatasourceOk() (*string, bool)`

GetDatasourceOk returns a tuple with the Datasource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasource

`func (o *GreenlistUsersRequest) SetDatasource(v string)`

SetDatasource sets Datasource field to given value.


### GetEmails

`func (o *GreenlistUsersRequest) GetEmails() []string`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *GreenlistUsersRequest) GetEmailsOk() (*[]string, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *GreenlistUsersRequest) SetEmails(v []string)`

SetEmails sets Emails field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


