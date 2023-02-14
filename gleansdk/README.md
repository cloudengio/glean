# Package [cloudeng.io/glean/gleansdk](https://pkg.go.dev/cloudeng.io/glean/gleansdk?tab=doc)

```go
import cloudeng.io/glean/gleansdk
```


## Variables
### ContextOAuth2, ContextBasicAuth, ContextAccessToken, ContextAPIKeys, ContextHttpSignatureAuth, ContextServerIndex, ContextOperationServerIndices, ContextServerVariables, ContextOperationServerVariables
```go
// ContextOAuth2 takes an oauth2.TokenSource as authentication for the request.
ContextOAuth2 = contextKey("token")
// ContextBasicAuth takes BasicAuth as authentication for the request.
ContextBasicAuth = contextKey("basic")
// ContextAccessToken takes a string oauth2 access token as authentication for the request.
ContextAccessToken = contextKey("accesstoken")
// ContextAPIKeys takes a string apikey as authentication for the request
ContextAPIKeys = contextKey("apiKeys")
// ContextHttpSignatureAuth takes HttpSignatureAuth as authentication for the request.
ContextHttpSignatureAuth = contextKey("httpsignature")
// ContextServerIndex uses a server configuration from the index.
ContextServerIndex = contextKey("serverIndex")
// ContextOperationServerIndices uses a server configuration from the index mapping.
ContextOperationServerIndices = contextKey("serverOperationIndices")
// ContextServerVariables overrides a server configuration variables.
ContextServerVariables = contextKey("serverVariables")
// ContextOperationServerVariables overrides a server configuration variables using operation specific values.
ContextOperationServerVariables = contextKey("serverOperationVariables")

```

### AllowedConnectorTypeEnumValues
```go
AllowedConnectorTypeEnumValues = []ConnectorType{
	"API_CRAWL",
	"BROWSER_CRAWL",
	"BROWSER_HISTORY",
	"BUILTIN",
	"FEDERATED_SEARCH",
	"PUSH_API",
	"WEB_CRAWL",
	"NATIVE_HISTORY",
}

```
All allowed values of ConnectorType enum



## Functions
### Func CacheExpires
```go
func CacheExpires(r *http.Response) time.Time
```
CacheExpires helper function to determine remaining time before repeating a
request.

### Func PtrBool
```go
func PtrBool(v bool) *bool
```
PtrBool is a helper routine that returns a pointer to given boolean value.

### Func PtrFloat32
```go
func PtrFloat32(v float32) *float32
```
PtrFloat32 is a helper routine that returns a pointer to given float value.

### Func PtrFloat64
```go
func PtrFloat64(v float64) *float64
```
PtrFloat64 is a helper routine that returns a pointer to given float value.

### Func PtrInt
```go
func PtrInt(v int) *int
```
PtrInt is a helper routine that returns a pointer to given integer value.

### Func PtrInt32
```go
func PtrInt32(v int32) *int32
```
PtrInt32 is a helper routine that returns a pointer to given integer value.

### Func PtrInt64
```go
func PtrInt64(v int64) *int64
```
PtrInt64 is a helper routine that returns a pointer to given integer value.

### Func PtrString
```go
func PtrString(v string) *string
```
PtrString is a helper routine that returns a pointer to given string value.

### Func PtrTime
```go
func PtrTime(v time.Time) *time.Time
```
PtrTime is helper routine that returns a pointer to given Time value.



## Types
### Type APIClient
```go
type APIClient struct {
	DatasourcesApi *DatasourcesApiService

	DocumentsApi *DocumentsApiService

	PeopleApi *PeopleApiService

	PermissionsApi *PermissionsApiService
	// contains filtered or unexported fields
}
```
APIClient manages communication with the Glean Indexing API API v0.9.0 In
most cases there should be only one, shared, APIClient.

### Functions

```go
func NewAPIClient(cfg *Configuration) *APIClient
```
NewAPIClient creates a new API client. Requires a userAgent string
describing your application. optionally a custom http.Client to allow for
advanced features such as caching.



### Methods

```go
func (c *APIClient) GetConfig() *Configuration
```
Allow modification of underlying config for alternate implementations and
testing Caution: modifying the configuration while live can cause data races
and potentially unwanted behavior




### Type APIKey
```go
type APIKey struct {
	Key    string
	Prefix string
}
```
APIKey provides API key based authentication to a request passed via context
using ContextAPIKey


### Type APIResponse
```go
type APIResponse struct {
	*http.Response `json:"-"`
	Message        string `json:"message,omitempty"`
	// Operation is the name of the OpenAPI operation.
	Operation string `json:"operation,omitempty"`
	// RequestURL is the request URL. This value is always available, even if the
	// embedded *http.Response is nil.
	RequestURL string `json:"url,omitempty"`
	// Method is the HTTP method used for the request.  This value is always
	// available, even if the embedded *http.Response is nil.
	Method string `json:"method,omitempty"`
	// Payload holds the contents of the response body (which may be nil or empty).
	// This is provided here as the raw response.Body() reader will have already
	// been drained.
	Payload []byte `json:"-"`
}
```
APIResponse stores the API response returned by the server.

### Functions

```go
func NewAPIResponse(r *http.Response) *APIResponse
```
NewAPIResponse returns a new APIResponse object.


```go
func NewAPIResponseWithError(errorMessage string) *APIResponse
```
NewAPIResponseWithError returns a new APIResponse object with the provided
error message.




### Type AdditionalFieldDefinition
```go
type AdditionalFieldDefinition struct {
	// Key to reference this field, e.g. \"languages\".
	Key *string `json:"key,omitempty"`
	// List of type string or HypertextField.  HypertextField is defined as ``` {   anchor: string,    // Anchor text for the hypertext field.   hyperlink: string, // URL for the hypertext field. } ``` Example: ```{\"anchor\":\"Glean\",\"hyperlink\":\"https://glean.com\"}```  When OpenAPI Generator supports oneOf, we will semantically enforce this in the docs.
	Value []map[string]interface{} `json:"value,omitempty"`
}
```
AdditionalFieldDefinition Additional information about the employee.

### Functions

```go
func NewAdditionalFieldDefinition() *AdditionalFieldDefinition
```
NewAdditionalFieldDefinition instantiates a new AdditionalFieldDefinition
object This constructor will assign default values to properties that have
it defined, and makes sure properties required by API are set, but the set
of arguments will change when the set of required properties is changed


```go
func NewAdditionalFieldDefinitionWithDefaults() *AdditionalFieldDefinition
```
NewAdditionalFieldDefinitionWithDefaults instantiates a new
AdditionalFieldDefinition object This constructor will only assign default
values to properties that have it defined, but it doesn't guarantee that
properties required by API are set



### Methods

```go
func (o *AdditionalFieldDefinition) GetKey() string
```
GetKey returns the Key field value if set, zero value otherwise.


```go
func (o *AdditionalFieldDefinition) GetKeyOk() (*string, bool)
```
GetKeyOk returns a tuple with the Key field value if set, nil otherwise and
a boolean to check if the value has been set.


```go
func (o *AdditionalFieldDefinition) GetValue() []map[string]interface{}
```
GetValue returns the Value field value if set, zero value otherwise.


```go
func (o *AdditionalFieldDefinition) GetValueOk() ([]map[string]interface{}, bool)
```
GetValueOk returns a tuple with the Value field value if set, nil otherwise
and a boolean to check if the value has been set.


```go
func (o *AdditionalFieldDefinition) HasKey() bool
```
HasKey returns a boolean if a field has been set.


```go
func (o *AdditionalFieldDefinition) HasValue() bool
```
HasValue returns a boolean if a field has been set.


```go
func (o AdditionalFieldDefinition) MarshalJSON() ([]byte, error)
```


```go
func (o *AdditionalFieldDefinition) SetKey(v string)
```
SetKey gets a reference to the given string and assigns it to the Key field.


```go
func (o *AdditionalFieldDefinition) SetValue(v []map[string]interface{})
```
SetValue gets a reference to the given []map[string]interface{} and assigns
it to the Value field.




### Type ApiAdddatasourcePostRequest
```go
type ApiAdddatasourcePostRequest struct {
	ApiService *DatasourcesApiService
	// contains filtered or unexported fields
}
```

### Methods

```go
func (r ApiAdddatasourcePostRequest) CustomDatasourceConfig(customDatasourceConfig CustomDatasourceConfig) ApiAdddatasourcePostRequest
```


```go
func (r ApiAdddatasourcePostRequest) Execute() (*http.Response, error)
```




### Type ApiBetausersPostRequest
```go
type ApiBetausersPostRequest struct {
	ApiService *PermissionsApiService
	// contains filtered or unexported fields
}
```

### Methods

```go
func (r ApiBetausersPostRequest) Execute() (*http.Response, error)
```


```go
func (r ApiBetausersPostRequest) GreenlistUsersRequest(greenlistUsersRequest GreenlistUsersRequest) ApiBetausersPostRequest
```




### Type ApiBulkindexdocumentsPostRequest
```go
type ApiBulkindexdocumentsPostRequest struct {
	ApiService *DocumentsApiService
	// contains filtered or unexported fields
}
```

### Methods

```go
func (r ApiBulkindexdocumentsPostRequest) BulkIndexDocumentsRequest(bulkIndexDocumentsRequest BulkIndexDocumentsRequest) ApiBulkindexdocumentsPostRequest
```


```go
func (r ApiBulkindexdocumentsPostRequest) Execute() (*http.Response, error)
```




### Type ApiBulkindexemployeesPostRequest
```go
type ApiBulkindexemployeesPostRequest struct {
	ApiService *PeopleApiService
	// contains filtered or unexported fields
}
```

### Methods

```go
func (r ApiBulkindexemployeesPostRequest) BulkIndexEmployeesRequest(bulkIndexEmployeesRequest BulkIndexEmployeesRequest) ApiBulkindexemployeesPostRequest
```


```go
func (r ApiBulkindexemployeesPostRequest) Execute() (*http.Response, error)
```




### Type ApiBulkindexgroupsPostRequest
```go
type ApiBulkindexgroupsPostRequest struct {
	ApiService *PermissionsApiService
	// contains filtered or unexported fields
}
```

### Methods

```go
func (r ApiBulkindexgroupsPostRequest) BulkIndexGroupsRequest(bulkIndexGroupsRequest BulkIndexGroupsRequest) ApiBulkindexgroupsPostRequest
```


```go
func (r ApiBulkindexgroupsPostRequest) Execute() (*http.Response, error)
```




### Type ApiBulkindexmembershipsPostRequest
```go
type ApiBulkindexmembershipsPostRequest struct {
	ApiService *PermissionsApiService
	// contains filtered or unexported fields
}
```

### Methods

```go
func (r ApiBulkindexmembershipsPostRequest) BulkIndexMembershipsRequest(bulkIndexMembershipsRequest BulkIndexMembershipsRequest) ApiBulkindexmembershipsPostRequest
```


```go
func (r ApiBulkindexmembershipsPostRequest) Execute() (*http.Response, error)
```




### Type ApiBulkindexteamsPostRequest
```go
type ApiBulkindexteamsPostRequest struct {
	ApiService *PeopleApiService
	// contains filtered or unexported fields
}
```

### Methods

```go
func (r ApiBulkindexteamsPostRequest) BulkIndexTeamsRequest(bulkIndexTeamsRequest BulkIndexTeamsRequest) ApiBulkindexteamsPostRequest
```


```go
func (r ApiBulkindexteamsPostRequest) Execute() (*http.Response, error)
```




### Type ApiBulkindexusersPostRequest
```go
type ApiBulkindexusersPostRequest struct {
	ApiService *PermissionsApiService
	// contains filtered or unexported fields
}
```

### Methods

```go
func (r ApiBulkindexusersPostRequest) BulkIndexUsersRequest(bulkIndexUsersRequest BulkIndexUsersRequest) ApiBulkindexusersPostRequest
```


```go
func (r ApiBulkindexusersPostRequest) Execute() (*http.Response, error)
```




### Type ApiCheckdocumentaccessPostRequest
```go
type ApiCheckdocumentaccessPostRequest struct {
	ApiService *PermissionsApiService
	// contains filtered or unexported fields
}
```

### Methods

```go
func (r ApiCheckdocumentaccessPostRequest) CheckDocumentAccessRequest(checkDocumentAccessRequest CheckDocumentAccessRequest) ApiCheckdocumentaccessPostRequest
```


```go
func (r ApiCheckdocumentaccessPostRequest) Execute() (*CheckDocumentAccessResponse, *http.Response, error)
```




### Type ApiDeletedocumentPostRequest
```go
type ApiDeletedocumentPostRequest struct {
	ApiService *DocumentsApiService
	// contains filtered or unexported fields
}
```

### Methods

```go
func (r ApiDeletedocumentPostRequest) DeleteDocumentRequest(deleteDocumentRequest DeleteDocumentRequest) ApiDeletedocumentPostRequest
```


```go
func (r ApiDeletedocumentPostRequest) Execute() (*http.Response, error)
```




### Type ApiDeleteemployeePostRequest
```go
type ApiDeleteemployeePostRequest struct {
	ApiService *PeopleApiService
	// contains filtered or unexported fields
}
```

### Methods

```go
func (r ApiDeleteemployeePostRequest) DeleteEmployeeRequest(deleteEmployeeRequest DeleteEmployeeRequest) ApiDeleteemployeePostRequest
```


```go
func (r ApiDeleteemployeePostRequest) Execute() (*http.Response, error)
```




### Type ApiDeletegroupPostRequest
```go
type ApiDeletegroupPostRequest struct {
	ApiService *PermissionsApiService
	// contains filtered or unexported fields
}
```

### Methods

```go
func (r ApiDeletegroupPostRequest) DeleteGroupRequest(deleteGroupRequest DeleteGroupRequest) ApiDeletegroupPostRequest
```


```go
func (r ApiDeletegroupPostRequest) Execute() (*http.Response, error)
```




### Type ApiDeletemembershipPostRequest
```go
type ApiDeletemembershipPostRequest struct {
	ApiService *PermissionsApiService
	// contains filtered or unexported fields
}
```

### Methods

```go
func (r ApiDeletemembershipPostRequest) DeleteMembershipRequest(deleteMembershipRequest DeleteMembershipRequest) ApiDeletemembershipPostRequest
```


```go
func (r ApiDeletemembershipPostRequest) Execute() (*http.Response, error)
```




### Type ApiDeleteuserPostRequest
```go
type ApiDeleteuserPostRequest struct {
	ApiService *PermissionsApiService
	// contains filtered or unexported fields
}
```

### Methods

```go
func (r ApiDeleteuserPostRequest) DeleteUserRequest(deleteUserRequest DeleteUserRequest) ApiDeleteuserPostRequest
```


```go
func (r ApiDeleteuserPostRequest) Execute() (*http.Response, error)
```




### Type ApiGetdatasourceconfigPostRequest
```go
type ApiGetdatasourceconfigPostRequest struct {
	ApiService *DatasourcesApiService
	// contains filtered or unexported fields
}
```

### Methods

```go
func (r ApiGetdatasourceconfigPostRequest) Execute() (*CustomDatasourceConfig, *http.Response, error)
```


```go
func (r ApiGetdatasourceconfigPostRequest) GetDatasourceConfigRequest(getDatasourceConfigRequest GetDatasourceConfigRequest) ApiGetdatasourceconfigPostRequest
```




### Type ApiGetdocumentcountPostRequest
```go
type ApiGetdocumentcountPostRequest struct {
	ApiService *DocumentsApiService
	// contains filtered or unexported fields
}
```

### Methods

```go
func (r ApiGetdocumentcountPostRequest) Execute() (*GetDocumentCountResponse, *http.Response, error)
```


```go
func (r ApiGetdocumentcountPostRequest) GetDocumentCountRequest(getDocumentCountRequest GetDocumentCountRequest) ApiGetdocumentcountPostRequest
```




### Type ApiGetdocumentstatusPostRequest
```go
type ApiGetdocumentstatusPostRequest struct {
	ApiService *DocumentsApiService
	// contains filtered or unexported fields
}
```

### Methods

```go
func (r ApiGetdocumentstatusPostRequest) Execute() (*GetDocumentStatusResponse, *http.Response, error)
```


```go
func (r ApiGetdocumentstatusPostRequest) GetDocumentStatusRequest(getDocumentStatusRequest GetDocumentStatusRequest) ApiGetdocumentstatusPostRequest
```




### Type ApiGetusercountPostRequest
```go
type ApiGetusercountPostRequest struct {
	ApiService *PermissionsApiService
	// contains filtered or unexported fields
}
```

### Methods

```go
func (r ApiGetusercountPostRequest) Execute() (*GetUserCountResponse, *http.Response, error)
```


```go
func (r ApiGetusercountPostRequest) GetUserCountRequest(getUserCountRequest GetUserCountRequest) ApiGetusercountPostRequest
```




### Type ApiIndexdocumentPostRequest
```go
type ApiIndexdocumentPostRequest struct {
	ApiService *DocumentsApiService
	// contains filtered or unexported fields
}
```

### Methods

```go
func (r ApiIndexdocumentPostRequest) Execute() (*http.Response, error)
```


```go
func (r ApiIndexdocumentPostRequest) IndexDocumentRequest(indexDocumentRequest IndexDocumentRequest) ApiIndexdocumentPostRequest
```




### Type ApiIndexemployeePostRequest
```go
type ApiIndexemployeePostRequest struct {
	ApiService *PeopleApiService
	// contains filtered or unexported fields
}
```

### Methods

```go
func (r ApiIndexemployeePostRequest) Execute() (*http.Response, error)
```


```go
func (r ApiIndexemployeePostRequest) IndexEmployeeRequest(indexEmployeeRequest IndexEmployeeRequest) ApiIndexemployeePostRequest
```




### Type ApiIndexemployeelistPostRequest
```go
type ApiIndexemployeelistPostRequest struct {
	ApiService *PeopleApiService
	// contains filtered or unexported fields
}
```

### Methods

```go
func (r ApiIndexemployeelistPostRequest) Execute() (*http.Response, error)
```


```go
func (r ApiIndexemployeelistPostRequest) IndexEmployeeListRequest(indexEmployeeListRequest IndexEmployeeListRequest) ApiIndexemployeelistPostRequest
```




### Type ApiIndexgroupPostRequest
```go
type ApiIndexgroupPostRequest struct {
	ApiService *PermissionsApiService
	// contains filtered or unexported fields
}
```

### Methods

```go
func (r ApiIndexgroupPostRequest) Execute() (*http.Response, error)
```


```go
func (r ApiIndexgroupPostRequest) IndexGroupRequest(indexGroupRequest IndexGroupRequest) ApiIndexgroupPostRequest
```




### Type ApiIndexmembershipPostRequest
```go
type ApiIndexmembershipPostRequest struct {
	ApiService *PermissionsApiService
	// contains filtered or unexported fields
}
```

### Methods

```go
func (r ApiIndexmembershipPostRequest) Execute() (*http.Response, error)
```


```go
func (r ApiIndexmembershipPostRequest) IndexMembershipRequest(indexMembershipRequest IndexMembershipRequest) ApiIndexmembershipPostRequest
```




### Type ApiIndexteamPostRequest
```go
type ApiIndexteamPostRequest struct {
	ApiService *PeopleApiService
	// contains filtered or unexported fields
}
```

### Methods

```go
func (r ApiIndexteamPostRequest) Execute() (*http.Response, error)
```


```go
func (r ApiIndexteamPostRequest) IndexTeamRequest(indexTeamRequest IndexTeamRequest) ApiIndexteamPostRequest
```




### Type ApiIndexuserPostRequest
```go
type ApiIndexuserPostRequest struct {
	ApiService *PermissionsApiService
	// contains filtered or unexported fields
}
```

### Methods

```go
func (r ApiIndexuserPostRequest) Execute() (*http.Response, error)
```


```go
func (r ApiIndexuserPostRequest) IndexUserRequest(indexUserRequest IndexUserRequest) ApiIndexuserPostRequest
```




### Type ApiProcessalldocumentsPostRequest
```go
type ApiProcessalldocumentsPostRequest struct {
	ApiService *DocumentsApiService
	// contains filtered or unexported fields
}
```

### Methods

```go
func (r ApiProcessalldocumentsPostRequest) Execute() (*http.Response, error)
```


```go
func (r ApiProcessalldocumentsPostRequest) ProcessAllDocumentsRequest(processAllDocumentsRequest ProcessAllDocumentsRequest) ApiProcessalldocumentsPostRequest
```




### Type ApiProcessallmembershipsPostRequest
```go
type ApiProcessallmembershipsPostRequest struct {
	ApiService *PermissionsApiService
	// contains filtered or unexported fields
}
```

### Methods

```go
func (r ApiProcessallmembershipsPostRequest) Execute() (*http.Response, error)
```


```go
func (r ApiProcessallmembershipsPostRequest) ProcessAllMembershipsRequest(processAllMembershipsRequest ProcessAllMembershipsRequest) ApiProcessallmembershipsPostRequest
```




### Type BasicAuth
```go
type BasicAuth struct {
	UserName string `json:"userName,omitempty"`
	Password string `json:"password,omitempty"`
}
```
BasicAuth provides basic http authentication to a request passed via context
using ContextBasicAuth


### Type BulkIndexDocumentsRequest
```go
type BulkIndexDocumentsRequest struct {
	// Unique id that must be used for this instance of datasource employees upload
	UploadId string `json:"uploadId"`
	// true if this is the first page of the upload. Defaults to false
	IsFirstPage *bool `json:"isFirstPage,omitempty"`
	// true if this is the last page of the upload. Defaults to false
	IsLastPage *bool `json:"isLastPage,omitempty"`
	// Flag to discard previous upload attempts and start from scratch. Must be specified with isFirstPage=true
	ForceRestartUpload *bool `json:"forceRestartUpload,omitempty"`
	// Datasource of the documents
	Datasource string `json:"datasource"`
	// Batch of documents for the datasource
	Documents []DocumentDefinition `json:"documents"`
	// True if older documents need to be force deleted after the upload completes. Defaults to older documents being deleted asynchronously. This must only be set when `isLastPage = true`
	DisableStaleDocumentDeletionCheck *bool `json:"disableStaleDocumentDeletionCheck,omitempty"`
}
```
BulkIndexDocumentsRequest Describes the request body of the
/bulkindexdocuments API call

### Functions

```go
func NewBulkIndexDocumentsRequest(uploadId string, datasource string, documents []DocumentDefinition) *BulkIndexDocumentsRequest
```
NewBulkIndexDocumentsRequest instantiates a new BulkIndexDocumentsRequest
object This constructor will assign default values to properties that have
it defined, and makes sure properties required by API are set, but the set
of arguments will change when the set of required properties is changed


```go
func NewBulkIndexDocumentsRequestWithDefaults() *BulkIndexDocumentsRequest
```
NewBulkIndexDocumentsRequestWithDefaults instantiates a new
BulkIndexDocumentsRequest object This constructor will only assign default
values to properties that have it defined, but it doesn't guarantee that
properties required by API are set



### Methods

```go
func (o *BulkIndexDocumentsRequest) GetDatasource() string
```
GetDatasource returns the Datasource field value


```go
func (o *BulkIndexDocumentsRequest) GetDatasourceOk() (*string, bool)
```
GetDatasourceOk returns a tuple with the Datasource field value and a
boolean to check if the value has been set.


```go
func (o *BulkIndexDocumentsRequest) GetDisableStaleDocumentDeletionCheck() bool
```
GetDisableStaleDocumentDeletionCheck returns the
DisableStaleDocumentDeletionCheck field value if set, zero value otherwise.


```go
func (o *BulkIndexDocumentsRequest) GetDisableStaleDocumentDeletionCheckOk() (*bool, bool)
```
GetDisableStaleDocumentDeletionCheckOk returns a tuple with the
DisableStaleDocumentDeletionCheck field value if set, nil otherwise and a
boolean to check if the value has been set.


```go
func (o *BulkIndexDocumentsRequest) GetDocuments() []DocumentDefinition
```
GetDocuments returns the Documents field value


```go
func (o *BulkIndexDocumentsRequest) GetDocumentsOk() ([]DocumentDefinition, bool)
```
GetDocumentsOk returns a tuple with the Documents field value and a boolean
to check if the value has been set.


```go
func (o *BulkIndexDocumentsRequest) GetForceRestartUpload() bool
```
GetForceRestartUpload returns the ForceRestartUpload field value if set,
zero value otherwise.


```go
func (o *BulkIndexDocumentsRequest) GetForceRestartUploadOk() (*bool, bool)
```
GetForceRestartUploadOk returns a tuple with the ForceRestartUpload field
value if set, nil otherwise and a boolean to check if the value has been
set.


```go
func (o *BulkIndexDocumentsRequest) GetIsFirstPage() bool
```
GetIsFirstPage returns the IsFirstPage field value if set, zero value
otherwise.


```go
func (o *BulkIndexDocumentsRequest) GetIsFirstPageOk() (*bool, bool)
```
GetIsFirstPageOk returns a tuple with the IsFirstPage field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *BulkIndexDocumentsRequest) GetIsLastPage() bool
```
GetIsLastPage returns the IsLastPage field value if set, zero value
otherwise.


```go
func (o *BulkIndexDocumentsRequest) GetIsLastPageOk() (*bool, bool)
```
GetIsLastPageOk returns a tuple with the IsLastPage field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *BulkIndexDocumentsRequest) GetUploadId() string
```
GetUploadId returns the UploadId field value


```go
func (o *BulkIndexDocumentsRequest) GetUploadIdOk() (*string, bool)
```
GetUploadIdOk returns a tuple with the UploadId field value and a boolean to
check if the value has been set.


```go
func (o *BulkIndexDocumentsRequest) HasDisableStaleDocumentDeletionCheck() bool
```
HasDisableStaleDocumentDeletionCheck returns a boolean if a field has been
set.


```go
func (o *BulkIndexDocumentsRequest) HasForceRestartUpload() bool
```
HasForceRestartUpload returns a boolean if a field has been set.


```go
func (o *BulkIndexDocumentsRequest) HasIsFirstPage() bool
```
HasIsFirstPage returns a boolean if a field has been set.


```go
func (o *BulkIndexDocumentsRequest) HasIsLastPage() bool
```
HasIsLastPage returns a boolean if a field has been set.


```go
func (o BulkIndexDocumentsRequest) MarshalJSON() ([]byte, error)
```


```go
func (o *BulkIndexDocumentsRequest) SetDatasource(v string)
```
SetDatasource sets field value


```go
func (o *BulkIndexDocumentsRequest) SetDisableStaleDocumentDeletionCheck(v bool)
```
SetDisableStaleDocumentDeletionCheck gets a reference to the given bool and
assigns it to the DisableStaleDocumentDeletionCheck field.


```go
func (o *BulkIndexDocumentsRequest) SetDocuments(v []DocumentDefinition)
```
SetDocuments sets field value


```go
func (o *BulkIndexDocumentsRequest) SetForceRestartUpload(v bool)
```
SetForceRestartUpload gets a reference to the given bool and assigns it to
the ForceRestartUpload field.


```go
func (o *BulkIndexDocumentsRequest) SetIsFirstPage(v bool)
```
SetIsFirstPage gets a reference to the given bool and assigns it to the
IsFirstPage field.


```go
func (o *BulkIndexDocumentsRequest) SetIsLastPage(v bool)
```
SetIsLastPage gets a reference to the given bool and assigns it to the
IsLastPage field.


```go
func (o *BulkIndexDocumentsRequest) SetUploadId(v string)
```
SetUploadId sets field value




### Type BulkIndexDocumentsRequestAllOf
```go
type BulkIndexDocumentsRequestAllOf struct {
	// Datasource of the documents
	Datasource string `json:"datasource"`
	// Batch of documents for the datasource
	Documents []DocumentDefinition `json:"documents"`
	// True if older documents need to be force deleted after the upload completes. Defaults to older documents being deleted asynchronously. This must only be set when `isLastPage = true`
	DisableStaleDocumentDeletionCheck *bool `json:"disableStaleDocumentDeletionCheck,omitempty"`
}
```
BulkIndexDocumentsRequestAllOf struct for BulkIndexDocumentsRequestAllOf

### Functions

```go
func NewBulkIndexDocumentsRequestAllOf(datasource string, documents []DocumentDefinition) *BulkIndexDocumentsRequestAllOf
```
NewBulkIndexDocumentsRequestAllOf instantiates a new
BulkIndexDocumentsRequestAllOf object This constructor will assign default
values to properties that have it defined, and makes sure properties
required by API are set, but the set of arguments will change when the set
of required properties is changed


```go
func NewBulkIndexDocumentsRequestAllOfWithDefaults() *BulkIndexDocumentsRequestAllOf
```
NewBulkIndexDocumentsRequestAllOfWithDefaults instantiates a new
BulkIndexDocumentsRequestAllOf object This constructor will only assign
default values to properties that have it defined, but it doesn't guarantee
that properties required by API are set



### Methods

```go
func (o *BulkIndexDocumentsRequestAllOf) GetDatasource() string
```
GetDatasource returns the Datasource field value


```go
func (o *BulkIndexDocumentsRequestAllOf) GetDatasourceOk() (*string, bool)
```
GetDatasourceOk returns a tuple with the Datasource field value and a
boolean to check if the value has been set.


```go
func (o *BulkIndexDocumentsRequestAllOf) GetDisableStaleDocumentDeletionCheck() bool
```
GetDisableStaleDocumentDeletionCheck returns the
DisableStaleDocumentDeletionCheck field value if set, zero value otherwise.


```go
func (o *BulkIndexDocumentsRequestAllOf) GetDisableStaleDocumentDeletionCheckOk() (*bool, bool)
```
GetDisableStaleDocumentDeletionCheckOk returns a tuple with the
DisableStaleDocumentDeletionCheck field value if set, nil otherwise and a
boolean to check if the value has been set.


```go
func (o *BulkIndexDocumentsRequestAllOf) GetDocuments() []DocumentDefinition
```
GetDocuments returns the Documents field value


```go
func (o *BulkIndexDocumentsRequestAllOf) GetDocumentsOk() ([]DocumentDefinition, bool)
```
GetDocumentsOk returns a tuple with the Documents field value and a boolean
to check if the value has been set.


```go
func (o *BulkIndexDocumentsRequestAllOf) HasDisableStaleDocumentDeletionCheck() bool
```
HasDisableStaleDocumentDeletionCheck returns a boolean if a field has been
set.


```go
func (o BulkIndexDocumentsRequestAllOf) MarshalJSON() ([]byte, error)
```


```go
func (o *BulkIndexDocumentsRequestAllOf) SetDatasource(v string)
```
SetDatasource sets field value


```go
func (o *BulkIndexDocumentsRequestAllOf) SetDisableStaleDocumentDeletionCheck(v bool)
```
SetDisableStaleDocumentDeletionCheck gets a reference to the given bool and
assigns it to the DisableStaleDocumentDeletionCheck field.


```go
func (o *BulkIndexDocumentsRequestAllOf) SetDocuments(v []DocumentDefinition)
```
SetDocuments sets field value




### Type BulkIndexEmployeesRequest
```go
type BulkIndexEmployeesRequest struct {
	// Unique id that must be used for this instance of datasource employees upload
	UploadId string `json:"uploadId"`
	// true if this is the first page of the upload. Defaults to false
	IsFirstPage *bool `json:"isFirstPage,omitempty"`
	// true if this is the last page of the upload. Defaults to false
	IsLastPage *bool `json:"isLastPage,omitempty"`
	// Flag to discard previous upload attempts and start from scratch. Must be specified with isFirstPage=true
	ForceRestartUpload *bool `json:"forceRestartUpload,omitempty"`
	// Batch of employee information
	Employees []EmployeeInfoDefinition `json:"employees"`
}
```
BulkIndexEmployeesRequest Describes the request body of the
/bulkindexemployees API call

### Functions

```go
func NewBulkIndexEmployeesRequest(uploadId string, employees []EmployeeInfoDefinition) *BulkIndexEmployeesRequest
```
NewBulkIndexEmployeesRequest instantiates a new BulkIndexEmployeesRequest
object This constructor will assign default values to properties that have
it defined, and makes sure properties required by API are set, but the set
of arguments will change when the set of required properties is changed


```go
func NewBulkIndexEmployeesRequestWithDefaults() *BulkIndexEmployeesRequest
```
NewBulkIndexEmployeesRequestWithDefaults instantiates a new
BulkIndexEmployeesRequest object This constructor will only assign default
values to properties that have it defined, but it doesn't guarantee that
properties required by API are set



### Methods

```go
func (o *BulkIndexEmployeesRequest) GetEmployees() []EmployeeInfoDefinition
```
GetEmployees returns the Employees field value


```go
func (o *BulkIndexEmployeesRequest) GetEmployeesOk() ([]EmployeeInfoDefinition, bool)
```
GetEmployeesOk returns a tuple with the Employees field value and a boolean
to check if the value has been set.


```go
func (o *BulkIndexEmployeesRequest) GetForceRestartUpload() bool
```
GetForceRestartUpload returns the ForceRestartUpload field value if set,
zero value otherwise.


```go
func (o *BulkIndexEmployeesRequest) GetForceRestartUploadOk() (*bool, bool)
```
GetForceRestartUploadOk returns a tuple with the ForceRestartUpload field
value if set, nil otherwise and a boolean to check if the value has been
set.


```go
func (o *BulkIndexEmployeesRequest) GetIsFirstPage() bool
```
GetIsFirstPage returns the IsFirstPage field value if set, zero value
otherwise.


```go
func (o *BulkIndexEmployeesRequest) GetIsFirstPageOk() (*bool, bool)
```
GetIsFirstPageOk returns a tuple with the IsFirstPage field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *BulkIndexEmployeesRequest) GetIsLastPage() bool
```
GetIsLastPage returns the IsLastPage field value if set, zero value
otherwise.


```go
func (o *BulkIndexEmployeesRequest) GetIsLastPageOk() (*bool, bool)
```
GetIsLastPageOk returns a tuple with the IsLastPage field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *BulkIndexEmployeesRequest) GetUploadId() string
```
GetUploadId returns the UploadId field value


```go
func (o *BulkIndexEmployeesRequest) GetUploadIdOk() (*string, bool)
```
GetUploadIdOk returns a tuple with the UploadId field value and a boolean to
check if the value has been set.


```go
func (o *BulkIndexEmployeesRequest) HasForceRestartUpload() bool
```
HasForceRestartUpload returns a boolean if a field has been set.


```go
func (o *BulkIndexEmployeesRequest) HasIsFirstPage() bool
```
HasIsFirstPage returns a boolean if a field has been set.


```go
func (o *BulkIndexEmployeesRequest) HasIsLastPage() bool
```
HasIsLastPage returns a boolean if a field has been set.


```go
func (o BulkIndexEmployeesRequest) MarshalJSON() ([]byte, error)
```


```go
func (o *BulkIndexEmployeesRequest) SetEmployees(v []EmployeeInfoDefinition)
```
SetEmployees sets field value


```go
func (o *BulkIndexEmployeesRequest) SetForceRestartUpload(v bool)
```
SetForceRestartUpload gets a reference to the given bool and assigns it to
the ForceRestartUpload field.


```go
func (o *BulkIndexEmployeesRequest) SetIsFirstPage(v bool)
```
SetIsFirstPage gets a reference to the given bool and assigns it to the
IsFirstPage field.


```go
func (o *BulkIndexEmployeesRequest) SetIsLastPage(v bool)
```
SetIsLastPage gets a reference to the given bool and assigns it to the
IsLastPage field.


```go
func (o *BulkIndexEmployeesRequest) SetUploadId(v string)
```
SetUploadId sets field value




### Type BulkIndexEmployeesRequestAllOf
```go
type BulkIndexEmployeesRequestAllOf struct {
	// Batch of employee information
	Employees []EmployeeInfoDefinition `json:"employees"`
}
```
BulkIndexEmployeesRequestAllOf struct for BulkIndexEmployeesRequestAllOf

### Functions

```go
func NewBulkIndexEmployeesRequestAllOf(employees []EmployeeInfoDefinition) *BulkIndexEmployeesRequestAllOf
```
NewBulkIndexEmployeesRequestAllOf instantiates a new
BulkIndexEmployeesRequestAllOf object This constructor will assign default
values to properties that have it defined, and makes sure properties
required by API are set, but the set of arguments will change when the set
of required properties is changed


```go
func NewBulkIndexEmployeesRequestAllOfWithDefaults() *BulkIndexEmployeesRequestAllOf
```
NewBulkIndexEmployeesRequestAllOfWithDefaults instantiates a new
BulkIndexEmployeesRequestAllOf object This constructor will only assign
default values to properties that have it defined, but it doesn't guarantee
that properties required by API are set



### Methods

```go
func (o *BulkIndexEmployeesRequestAllOf) GetEmployees() []EmployeeInfoDefinition
```
GetEmployees returns the Employees field value


```go
func (o *BulkIndexEmployeesRequestAllOf) GetEmployeesOk() ([]EmployeeInfoDefinition, bool)
```
GetEmployeesOk returns a tuple with the Employees field value and a boolean
to check if the value has been set.


```go
func (o BulkIndexEmployeesRequestAllOf) MarshalJSON() ([]byte, error)
```


```go
func (o *BulkIndexEmployeesRequestAllOf) SetEmployees(v []EmployeeInfoDefinition)
```
SetEmployees sets field value




### Type BulkIndexGroupsRequest
```go
type BulkIndexGroupsRequest struct {
	// Unique id that must be used for this instance of datasource groups upload
	UploadId string `json:"uploadId"`
	// true if this is the first page of the upload. Defaults to false
	IsFirstPage *bool `json:"isFirstPage,omitempty"`
	// true if this is the last page of the upload. Defaults to false
	IsLastPage *bool `json:"isLastPage,omitempty"`
	// Flag to discard previous upload attempts and start from scratch. Must be specified with isFirstPage=true
	ForceRestartUpload *bool `json:"forceRestartUpload,omitempty"`
	// datasource of the groups
	Datasource string `json:"datasource"`
	// batch of groups for the datasource
	Groups []DatasourceGroupDefinition `json:"groups"`
}
```
BulkIndexGroupsRequest Describes the request body for the /bulkindexgroups
API call

### Functions

```go
func NewBulkIndexGroupsRequest(uploadId string, datasource string, groups []DatasourceGroupDefinition) *BulkIndexGroupsRequest
```
NewBulkIndexGroupsRequest instantiates a new BulkIndexGroupsRequest object
This constructor will assign default values to properties that have it
defined, and makes sure properties required by API are set, but the set of
arguments will change when the set of required properties is changed


```go
func NewBulkIndexGroupsRequestWithDefaults() *BulkIndexGroupsRequest
```
NewBulkIndexGroupsRequestWithDefaults instantiates a new
BulkIndexGroupsRequest object This constructor will only assign default
values to properties that have it defined, but it doesn't guarantee that
properties required by API are set



### Methods

```go
func (o *BulkIndexGroupsRequest) GetDatasource() string
```
GetDatasource returns the Datasource field value


```go
func (o *BulkIndexGroupsRequest) GetDatasourceOk() (*string, bool)
```
GetDatasourceOk returns a tuple with the Datasource field value and a
boolean to check if the value has been set.


```go
func (o *BulkIndexGroupsRequest) GetForceRestartUpload() bool
```
GetForceRestartUpload returns the ForceRestartUpload field value if set,
zero value otherwise.


```go
func (o *BulkIndexGroupsRequest) GetForceRestartUploadOk() (*bool, bool)
```
GetForceRestartUploadOk returns a tuple with the ForceRestartUpload field
value if set, nil otherwise and a boolean to check if the value has been
set.


```go
func (o *BulkIndexGroupsRequest) GetGroups() []DatasourceGroupDefinition
```
GetGroups returns the Groups field value


```go
func (o *BulkIndexGroupsRequest) GetGroupsOk() ([]DatasourceGroupDefinition, bool)
```
GetGroupsOk returns a tuple with the Groups field value and a boolean to
check if the value has been set.


```go
func (o *BulkIndexGroupsRequest) GetIsFirstPage() bool
```
GetIsFirstPage returns the IsFirstPage field value if set, zero value
otherwise.


```go
func (o *BulkIndexGroupsRequest) GetIsFirstPageOk() (*bool, bool)
```
GetIsFirstPageOk returns a tuple with the IsFirstPage field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *BulkIndexGroupsRequest) GetIsLastPage() bool
```
GetIsLastPage returns the IsLastPage field value if set, zero value
otherwise.


```go
func (o *BulkIndexGroupsRequest) GetIsLastPageOk() (*bool, bool)
```
GetIsLastPageOk returns a tuple with the IsLastPage field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *BulkIndexGroupsRequest) GetUploadId() string
```
GetUploadId returns the UploadId field value


```go
func (o *BulkIndexGroupsRequest) GetUploadIdOk() (*string, bool)
```
GetUploadIdOk returns a tuple with the UploadId field value and a boolean to
check if the value has been set.


```go
func (o *BulkIndexGroupsRequest) HasForceRestartUpload() bool
```
HasForceRestartUpload returns a boolean if a field has been set.


```go
func (o *BulkIndexGroupsRequest) HasIsFirstPage() bool
```
HasIsFirstPage returns a boolean if a field has been set.


```go
func (o *BulkIndexGroupsRequest) HasIsLastPage() bool
```
HasIsLastPage returns a boolean if a field has been set.


```go
func (o BulkIndexGroupsRequest) MarshalJSON() ([]byte, error)
```


```go
func (o *BulkIndexGroupsRequest) SetDatasource(v string)
```
SetDatasource sets field value


```go
func (o *BulkIndexGroupsRequest) SetForceRestartUpload(v bool)
```
SetForceRestartUpload gets a reference to the given bool and assigns it to
the ForceRestartUpload field.


```go
func (o *BulkIndexGroupsRequest) SetGroups(v []DatasourceGroupDefinition)
```
SetGroups sets field value


```go
func (o *BulkIndexGroupsRequest) SetIsFirstPage(v bool)
```
SetIsFirstPage gets a reference to the given bool and assigns it to the
IsFirstPage field.


```go
func (o *BulkIndexGroupsRequest) SetIsLastPage(v bool)
```
SetIsLastPage gets a reference to the given bool and assigns it to the
IsLastPage field.


```go
func (o *BulkIndexGroupsRequest) SetUploadId(v string)
```
SetUploadId sets field value




### Type BulkIndexMembershipsRequest
```go
type BulkIndexMembershipsRequest struct {
	// Unique id that must be used for this instance of datasource group memberships upload
	UploadId string `json:"uploadId"`
	// true if this is the first page of the upload. Defaults to false
	IsFirstPage *bool `json:"isFirstPage,omitempty"`
	// true if this is the last page of the upload. Defaults to false
	IsLastPage *bool `json:"isLastPage,omitempty"`
	// Flag to discard previous upload attempts and start from scratch. Must be specified with isFirstPage=true
	ForceRestartUpload *bool `json:"forceRestartUpload,omitempty"`
	// datasource of the memberships
	Datasource string `json:"datasource"`
	// group who's memberships are specified
	Group *string `json:"group,omitempty"`
	// batch of memberships for the group
	Memberships []DatasourceBulkMembershipDefinition `json:"memberships"`
}
```
BulkIndexMembershipsRequest Describes the request body for the
/bulkindexmemberships API call

### Functions

```go
func NewBulkIndexMembershipsRequest(uploadId string, datasource string, memberships []DatasourceBulkMembershipDefinition) *BulkIndexMembershipsRequest
```
NewBulkIndexMembershipsRequest instantiates a new
BulkIndexMembershipsRequest object This constructor will assign default
values to properties that have it defined, and makes sure properties
required by API are set, but the set of arguments will change when the set
of required properties is changed


```go
func NewBulkIndexMembershipsRequestWithDefaults() *BulkIndexMembershipsRequest
```
NewBulkIndexMembershipsRequestWithDefaults instantiates a new
BulkIndexMembershipsRequest object This constructor will only assign default
values to properties that have it defined, but it doesn't guarantee that
properties required by API are set



### Methods

```go
func (o *BulkIndexMembershipsRequest) GetDatasource() string
```
GetDatasource returns the Datasource field value


```go
func (o *BulkIndexMembershipsRequest) GetDatasourceOk() (*string, bool)
```
GetDatasourceOk returns a tuple with the Datasource field value and a
boolean to check if the value has been set.


```go
func (o *BulkIndexMembershipsRequest) GetForceRestartUpload() bool
```
GetForceRestartUpload returns the ForceRestartUpload field value if set,
zero value otherwise.


```go
func (o *BulkIndexMembershipsRequest) GetForceRestartUploadOk() (*bool, bool)
```
GetForceRestartUploadOk returns a tuple with the ForceRestartUpload field
value if set, nil otherwise and a boolean to check if the value has been
set.


```go
func (o *BulkIndexMembershipsRequest) GetGroup() string
```
GetGroup returns the Group field value if set, zero value otherwise.


```go
func (o *BulkIndexMembershipsRequest) GetGroupOk() (*string, bool)
```
GetGroupOk returns a tuple with the Group field value if set, nil otherwise
and a boolean to check if the value has been set.


```go
func (o *BulkIndexMembershipsRequest) GetIsFirstPage() bool
```
GetIsFirstPage returns the IsFirstPage field value if set, zero value
otherwise.


```go
func (o *BulkIndexMembershipsRequest) GetIsFirstPageOk() (*bool, bool)
```
GetIsFirstPageOk returns a tuple with the IsFirstPage field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *BulkIndexMembershipsRequest) GetIsLastPage() bool
```
GetIsLastPage returns the IsLastPage field value if set, zero value
otherwise.


```go
func (o *BulkIndexMembershipsRequest) GetIsLastPageOk() (*bool, bool)
```
GetIsLastPageOk returns a tuple with the IsLastPage field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *BulkIndexMembershipsRequest) GetMemberships() []DatasourceBulkMembershipDefinition
```
GetMemberships returns the Memberships field value


```go
func (o *BulkIndexMembershipsRequest) GetMembershipsOk() ([]DatasourceBulkMembershipDefinition, bool)
```
GetMembershipsOk returns a tuple with the Memberships field value and a
boolean to check if the value has been set.


```go
func (o *BulkIndexMembershipsRequest) GetUploadId() string
```
GetUploadId returns the UploadId field value


```go
func (o *BulkIndexMembershipsRequest) GetUploadIdOk() (*string, bool)
```
GetUploadIdOk returns a tuple with the UploadId field value and a boolean to
check if the value has been set.


```go
func (o *BulkIndexMembershipsRequest) HasForceRestartUpload() bool
```
HasForceRestartUpload returns a boolean if a field has been set.


```go
func (o *BulkIndexMembershipsRequest) HasGroup() bool
```
HasGroup returns a boolean if a field has been set.


```go
func (o *BulkIndexMembershipsRequest) HasIsFirstPage() bool
```
HasIsFirstPage returns a boolean if a field has been set.


```go
func (o *BulkIndexMembershipsRequest) HasIsLastPage() bool
```
HasIsLastPage returns a boolean if a field has been set.


```go
func (o BulkIndexMembershipsRequest) MarshalJSON() ([]byte, error)
```


```go
func (o *BulkIndexMembershipsRequest) SetDatasource(v string)
```
SetDatasource sets field value


```go
func (o *BulkIndexMembershipsRequest) SetForceRestartUpload(v bool)
```
SetForceRestartUpload gets a reference to the given bool and assigns it to
the ForceRestartUpload field.


```go
func (o *BulkIndexMembershipsRequest) SetGroup(v string)
```
SetGroup gets a reference to the given string and assigns it to the Group
field.


```go
func (o *BulkIndexMembershipsRequest) SetIsFirstPage(v bool)
```
SetIsFirstPage gets a reference to the given bool and assigns it to the
IsFirstPage field.


```go
func (o *BulkIndexMembershipsRequest) SetIsLastPage(v bool)
```
SetIsLastPage gets a reference to the given bool and assigns it to the
IsLastPage field.


```go
func (o *BulkIndexMembershipsRequest) SetMemberships(v []DatasourceBulkMembershipDefinition)
```
SetMemberships sets field value


```go
func (o *BulkIndexMembershipsRequest) SetUploadId(v string)
```
SetUploadId sets field value




### Type BulkIndexRequest
```go
type BulkIndexRequest struct {
	// Unique id that must be used for this instance of datasource employees upload
	UploadId string `json:"uploadId"`
	// true if this is the first page of the upload. Defaults to false
	IsFirstPage *bool `json:"isFirstPage,omitempty"`
	// true if this is the last page of the upload. Defaults to false
	IsLastPage *bool `json:"isLastPage,omitempty"`
	// Flag to discard previous upload attempts and start from scratch. Must be specified with isFirstPage=true
	ForceRestartUpload *bool `json:"forceRestartUpload,omitempty"`
}
```
BulkIndexRequest Describes the request body of the /bulkindexteams API call

### Functions

```go
func NewBulkIndexRequest(uploadId string) *BulkIndexRequest
```
NewBulkIndexRequest instantiates a new BulkIndexRequest object This
constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed


```go
func NewBulkIndexRequestWithDefaults() *BulkIndexRequest
```
NewBulkIndexRequestWithDefaults instantiates a new BulkIndexRequest object
This constructor will only assign default values to properties that have it
defined, but it doesn't guarantee that properties required by API are set



### Methods

```go
func (o *BulkIndexRequest) GetForceRestartUpload() bool
```
GetForceRestartUpload returns the ForceRestartUpload field value if set,
zero value otherwise.


```go
func (o *BulkIndexRequest) GetForceRestartUploadOk() (*bool, bool)
```
GetForceRestartUploadOk returns a tuple with the ForceRestartUpload field
value if set, nil otherwise and a boolean to check if the value has been
set.


```go
func (o *BulkIndexRequest) GetIsFirstPage() bool
```
GetIsFirstPage returns the IsFirstPage field value if set, zero value
otherwise.


```go
func (o *BulkIndexRequest) GetIsFirstPageOk() (*bool, bool)
```
GetIsFirstPageOk returns a tuple with the IsFirstPage field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *BulkIndexRequest) GetIsLastPage() bool
```
GetIsLastPage returns the IsLastPage field value if set, zero value
otherwise.


```go
func (o *BulkIndexRequest) GetIsLastPageOk() (*bool, bool)
```
GetIsLastPageOk returns a tuple with the IsLastPage field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *BulkIndexRequest) GetUploadId() string
```
GetUploadId returns the UploadId field value


```go
func (o *BulkIndexRequest) GetUploadIdOk() (*string, bool)
```
GetUploadIdOk returns a tuple with the UploadId field value and a boolean to
check if the value has been set.


```go
func (o *BulkIndexRequest) HasForceRestartUpload() bool
```
HasForceRestartUpload returns a boolean if a field has been set.


```go
func (o *BulkIndexRequest) HasIsFirstPage() bool
```
HasIsFirstPage returns a boolean if a field has been set.


```go
func (o *BulkIndexRequest) HasIsLastPage() bool
```
HasIsLastPage returns a boolean if a field has been set.


```go
func (o BulkIndexRequest) MarshalJSON() ([]byte, error)
```


```go
func (o *BulkIndexRequest) SetForceRestartUpload(v bool)
```
SetForceRestartUpload gets a reference to the given bool and assigns it to
the ForceRestartUpload field.


```go
func (o *BulkIndexRequest) SetIsFirstPage(v bool)
```
SetIsFirstPage gets a reference to the given bool and assigns it to the
IsFirstPage field.


```go
func (o *BulkIndexRequest) SetIsLastPage(v bool)
```
SetIsLastPage gets a reference to the given bool and assigns it to the
IsLastPage field.


```go
func (o *BulkIndexRequest) SetUploadId(v string)
```
SetUploadId sets field value




### Type BulkIndexTeamsRequest
```go
type BulkIndexTeamsRequest struct {
	// Unique id that must be used for this instance of datasource employees upload
	UploadId string `json:"uploadId"`
	// true if this is the first page of the upload. Defaults to false
	IsFirstPage *bool `json:"isFirstPage,omitempty"`
	// true if this is the last page of the upload. Defaults to false
	IsLastPage *bool `json:"isLastPage,omitempty"`
	// Flag to discard previous upload attempts and start from scratch. Must be specified with isFirstPage=true
	ForceRestartUpload *bool `json:"forceRestartUpload,omitempty"`
	// Batch of team information
	Teams []TeamInfoDefinition `json:"teams"`
}
```
BulkIndexTeamsRequest Describes the request body of the /bulkindexteams API
call

### Functions

```go
func NewBulkIndexTeamsRequest(uploadId string, teams []TeamInfoDefinition) *BulkIndexTeamsRequest
```
NewBulkIndexTeamsRequest instantiates a new BulkIndexTeamsRequest object
This constructor will assign default values to properties that have it
defined, and makes sure properties required by API are set, but the set of
arguments will change when the set of required properties is changed


```go
func NewBulkIndexTeamsRequestWithDefaults() *BulkIndexTeamsRequest
```
NewBulkIndexTeamsRequestWithDefaults instantiates a new
BulkIndexTeamsRequest object This constructor will only assign default
values to properties that have it defined, but it doesn't guarantee that
properties required by API are set



### Methods

```go
func (o *BulkIndexTeamsRequest) GetForceRestartUpload() bool
```
GetForceRestartUpload returns the ForceRestartUpload field value if set,
zero value otherwise.


```go
func (o *BulkIndexTeamsRequest) GetForceRestartUploadOk() (*bool, bool)
```
GetForceRestartUploadOk returns a tuple with the ForceRestartUpload field
value if set, nil otherwise and a boolean to check if the value has been
set.


```go
func (o *BulkIndexTeamsRequest) GetIsFirstPage() bool
```
GetIsFirstPage returns the IsFirstPage field value if set, zero value
otherwise.


```go
func (o *BulkIndexTeamsRequest) GetIsFirstPageOk() (*bool, bool)
```
GetIsFirstPageOk returns a tuple with the IsFirstPage field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *BulkIndexTeamsRequest) GetIsLastPage() bool
```
GetIsLastPage returns the IsLastPage field value if set, zero value
otherwise.


```go
func (o *BulkIndexTeamsRequest) GetIsLastPageOk() (*bool, bool)
```
GetIsLastPageOk returns a tuple with the IsLastPage field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *BulkIndexTeamsRequest) GetTeams() []TeamInfoDefinition
```
GetTeams returns the Teams field value


```go
func (o *BulkIndexTeamsRequest) GetTeamsOk() ([]TeamInfoDefinition, bool)
```
GetTeamsOk returns a tuple with the Teams field value and a boolean to check
if the value has been set.


```go
func (o *BulkIndexTeamsRequest) GetUploadId() string
```
GetUploadId returns the UploadId field value


```go
func (o *BulkIndexTeamsRequest) GetUploadIdOk() (*string, bool)
```
GetUploadIdOk returns a tuple with the UploadId field value and a boolean to
check if the value has been set.


```go
func (o *BulkIndexTeamsRequest) HasForceRestartUpload() bool
```
HasForceRestartUpload returns a boolean if a field has been set.


```go
func (o *BulkIndexTeamsRequest) HasIsFirstPage() bool
```
HasIsFirstPage returns a boolean if a field has been set.


```go
func (o *BulkIndexTeamsRequest) HasIsLastPage() bool
```
HasIsLastPage returns a boolean if a field has been set.


```go
func (o BulkIndexTeamsRequest) MarshalJSON() ([]byte, error)
```


```go
func (o *BulkIndexTeamsRequest) SetForceRestartUpload(v bool)
```
SetForceRestartUpload gets a reference to the given bool and assigns it to
the ForceRestartUpload field.


```go
func (o *BulkIndexTeamsRequest) SetIsFirstPage(v bool)
```
SetIsFirstPage gets a reference to the given bool and assigns it to the
IsFirstPage field.


```go
func (o *BulkIndexTeamsRequest) SetIsLastPage(v bool)
```
SetIsLastPage gets a reference to the given bool and assigns it to the
IsLastPage field.


```go
func (o *BulkIndexTeamsRequest) SetTeams(v []TeamInfoDefinition)
```
SetTeams sets field value


```go
func (o *BulkIndexTeamsRequest) SetUploadId(v string)
```
SetUploadId sets field value




### Type BulkIndexTeamsRequestAllOf
```go
type BulkIndexTeamsRequestAllOf struct {
	// Batch of team information
	Teams []TeamInfoDefinition `json:"teams"`
}
```
BulkIndexTeamsRequestAllOf struct for BulkIndexTeamsRequestAllOf

### Functions

```go
func NewBulkIndexTeamsRequestAllOf(teams []TeamInfoDefinition) *BulkIndexTeamsRequestAllOf
```
NewBulkIndexTeamsRequestAllOf instantiates a new BulkIndexTeamsRequestAllOf
object This constructor will assign default values to properties that have
it defined, and makes sure properties required by API are set, but the set
of arguments will change when the set of required properties is changed


```go
func NewBulkIndexTeamsRequestAllOfWithDefaults() *BulkIndexTeamsRequestAllOf
```
NewBulkIndexTeamsRequestAllOfWithDefaults instantiates a new
BulkIndexTeamsRequestAllOf object This constructor will only assign default
values to properties that have it defined, but it doesn't guarantee that
properties required by API are set



### Methods

```go
func (o *BulkIndexTeamsRequestAllOf) GetTeams() []TeamInfoDefinition
```
GetTeams returns the Teams field value


```go
func (o *BulkIndexTeamsRequestAllOf) GetTeamsOk() ([]TeamInfoDefinition, bool)
```
GetTeamsOk returns a tuple with the Teams field value and a boolean to check
if the value has been set.


```go
func (o BulkIndexTeamsRequestAllOf) MarshalJSON() ([]byte, error)
```


```go
func (o *BulkIndexTeamsRequestAllOf) SetTeams(v []TeamInfoDefinition)
```
SetTeams sets field value




### Type BulkIndexUsersRequest
```go
type BulkIndexUsersRequest struct {
	// Unique id that must be used for this instance of datasource users upload
	UploadId string `json:"uploadId"`
	// true if this is the first page of the upload. Defaults to false
	IsFirstPage *bool `json:"isFirstPage,omitempty"`
	// true if this is the last page of the upload. Defaults to false
	IsLastPage *bool `json:"isLastPage,omitempty"`
	// Flag to discard previous upload attempts and start from scratch. Must be specified with isFirstPage=true
	ForceRestartUpload *bool `json:"forceRestartUpload,omitempty"`
	// datasource of the users
	Datasource string `json:"datasource"`
	// batch of users for the datasource
	Users []DatasourceUserDefinition `json:"users"`
}
```
BulkIndexUsersRequest Describes the request body for the /bulkindexusers API
call

### Functions

```go
func NewBulkIndexUsersRequest(uploadId string, datasource string, users []DatasourceUserDefinition) *BulkIndexUsersRequest
```
NewBulkIndexUsersRequest instantiates a new BulkIndexUsersRequest object
This constructor will assign default values to properties that have it
defined, and makes sure properties required by API are set, but the set of
arguments will change when the set of required properties is changed


```go
func NewBulkIndexUsersRequestWithDefaults() *BulkIndexUsersRequest
```
NewBulkIndexUsersRequestWithDefaults instantiates a new
BulkIndexUsersRequest object This constructor will only assign default
values to properties that have it defined, but it doesn't guarantee that
properties required by API are set



### Methods

```go
func (o *BulkIndexUsersRequest) GetDatasource() string
```
GetDatasource returns the Datasource field value


```go
func (o *BulkIndexUsersRequest) GetDatasourceOk() (*string, bool)
```
GetDatasourceOk returns a tuple with the Datasource field value and a
boolean to check if the value has been set.


```go
func (o *BulkIndexUsersRequest) GetForceRestartUpload() bool
```
GetForceRestartUpload returns the ForceRestartUpload field value if set,
zero value otherwise.


```go
func (o *BulkIndexUsersRequest) GetForceRestartUploadOk() (*bool, bool)
```
GetForceRestartUploadOk returns a tuple with the ForceRestartUpload field
value if set, nil otherwise and a boolean to check if the value has been
set.


```go
func (o *BulkIndexUsersRequest) GetIsFirstPage() bool
```
GetIsFirstPage returns the IsFirstPage field value if set, zero value
otherwise.


```go
func (o *BulkIndexUsersRequest) GetIsFirstPageOk() (*bool, bool)
```
GetIsFirstPageOk returns a tuple with the IsFirstPage field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *BulkIndexUsersRequest) GetIsLastPage() bool
```
GetIsLastPage returns the IsLastPage field value if set, zero value
otherwise.


```go
func (o *BulkIndexUsersRequest) GetIsLastPageOk() (*bool, bool)
```
GetIsLastPageOk returns a tuple with the IsLastPage field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *BulkIndexUsersRequest) GetUploadId() string
```
GetUploadId returns the UploadId field value


```go
func (o *BulkIndexUsersRequest) GetUploadIdOk() (*string, bool)
```
GetUploadIdOk returns a tuple with the UploadId field value and a boolean to
check if the value has been set.


```go
func (o *BulkIndexUsersRequest) GetUsers() []DatasourceUserDefinition
```
GetUsers returns the Users field value


```go
func (o *BulkIndexUsersRequest) GetUsersOk() ([]DatasourceUserDefinition, bool)
```
GetUsersOk returns a tuple with the Users field value and a boolean to check
if the value has been set.


```go
func (o *BulkIndexUsersRequest) HasForceRestartUpload() bool
```
HasForceRestartUpload returns a boolean if a field has been set.


```go
func (o *BulkIndexUsersRequest) HasIsFirstPage() bool
```
HasIsFirstPage returns a boolean if a field has been set.


```go
func (o *BulkIndexUsersRequest) HasIsLastPage() bool
```
HasIsLastPage returns a boolean if a field has been set.


```go
func (o BulkIndexUsersRequest) MarshalJSON() ([]byte, error)
```


```go
func (o *BulkIndexUsersRequest) SetDatasource(v string)
```
SetDatasource sets field value


```go
func (o *BulkIndexUsersRequest) SetForceRestartUpload(v bool)
```
SetForceRestartUpload gets a reference to the given bool and assigns it to
the ForceRestartUpload field.


```go
func (o *BulkIndexUsersRequest) SetIsFirstPage(v bool)
```
SetIsFirstPage gets a reference to the given bool and assigns it to the
IsFirstPage field.


```go
func (o *BulkIndexUsersRequest) SetIsLastPage(v bool)
```
SetIsLastPage gets a reference to the given bool and assigns it to the
IsLastPage field.


```go
func (o *BulkIndexUsersRequest) SetUploadId(v string)
```
SetUploadId sets field value


```go
func (o *BulkIndexUsersRequest) SetUsers(v []DatasourceUserDefinition)
```
SetUsers sets field value




### Type CanonicalizingRegexType
```go
type CanonicalizingRegexType struct {
	// Regular expression to match to an arbitrary string.
	MatchRegex *string `json:"matchRegex,omitempty"`
	// Regular expression to transform into a canonical string.
	RewriteRegex *string `json:"rewriteRegex,omitempty"`
}
```
CanonicalizingRegexType Regular expression to apply to an arbitrary string
to transform it into a canonical string.

### Functions

```go
func NewCanonicalizingRegexType() *CanonicalizingRegexType
```
NewCanonicalizingRegexType instantiates a new CanonicalizingRegexType object
This constructor will assign default values to properties that have it
defined, and makes sure properties required by API are set, but the set of
arguments will change when the set of required properties is changed


```go
func NewCanonicalizingRegexTypeWithDefaults() *CanonicalizingRegexType
```
NewCanonicalizingRegexTypeWithDefaults instantiates a new
CanonicalizingRegexType object This constructor will only assign default
values to properties that have it defined, but it doesn't guarantee that
properties required by API are set



### Methods

```go
func (o *CanonicalizingRegexType) GetMatchRegex() string
```
GetMatchRegex returns the MatchRegex field value if set, zero value
otherwise.


```go
func (o *CanonicalizingRegexType) GetMatchRegexOk() (*string, bool)
```
GetMatchRegexOk returns a tuple with the MatchRegex field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *CanonicalizingRegexType) GetRewriteRegex() string
```
GetRewriteRegex returns the RewriteRegex field value if set, zero value
otherwise.


```go
func (o *CanonicalizingRegexType) GetRewriteRegexOk() (*string, bool)
```
GetRewriteRegexOk returns a tuple with the RewriteRegex field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *CanonicalizingRegexType) HasMatchRegex() bool
```
HasMatchRegex returns a boolean if a field has been set.


```go
func (o *CanonicalizingRegexType) HasRewriteRegex() bool
```
HasRewriteRegex returns a boolean if a field has been set.


```go
func (o CanonicalizingRegexType) MarshalJSON() ([]byte, error)
```


```go
func (o *CanonicalizingRegexType) SetMatchRegex(v string)
```
SetMatchRegex gets a reference to the given string and assigns it to the
MatchRegex field.


```go
func (o *CanonicalizingRegexType) SetRewriteRegex(v string)
```
SetRewriteRegex gets a reference to the given string and assigns it to the
RewriteRegex field.




### Type CheckDocumentAccessRequest
```go
type CheckDocumentAccessRequest struct {
	// Datasource of document to get check access for
	Datasource string `json:"datasource"`
	// Object type of document to get check access for
	ObjectType string `json:"objectType"`
	// ID of document to get check access for
	DocId string `json:"docId"`
	// Email of user to check access for
	UserEmail string `json:"userEmail"`
}
```
CheckDocumentAccessRequest Describes the request body of the
/checkdocumentaccess API call

### Functions

```go
func NewCheckDocumentAccessRequest(datasource string, objectType string, docId string, userEmail string) *CheckDocumentAccessRequest
```
NewCheckDocumentAccessRequest instantiates a new CheckDocumentAccessRequest
object This constructor will assign default values to properties that have
it defined, and makes sure properties required by API are set, but the set
of arguments will change when the set of required properties is changed


```go
func NewCheckDocumentAccessRequestWithDefaults() *CheckDocumentAccessRequest
```
NewCheckDocumentAccessRequestWithDefaults instantiates a new
CheckDocumentAccessRequest object This constructor will only assign default
values to properties that have it defined, but it doesn't guarantee that
properties required by API are set



### Methods

```go
func (o *CheckDocumentAccessRequest) GetDatasource() string
```
GetDatasource returns the Datasource field value


```go
func (o *CheckDocumentAccessRequest) GetDatasourceOk() (*string, bool)
```
GetDatasourceOk returns a tuple with the Datasource field value and a
boolean to check if the value has been set.


```go
func (o *CheckDocumentAccessRequest) GetDocId() string
```
GetDocId returns the DocId field value


```go
func (o *CheckDocumentAccessRequest) GetDocIdOk() (*string, bool)
```
GetDocIdOk returns a tuple with the DocId field value and a boolean to check
if the value has been set.


```go
func (o *CheckDocumentAccessRequest) GetObjectType() string
```
GetObjectType returns the ObjectType field value


```go
func (o *CheckDocumentAccessRequest) GetObjectTypeOk() (*string, bool)
```
GetObjectTypeOk returns a tuple with the ObjectType field value and a
boolean to check if the value has been set.


```go
func (o *CheckDocumentAccessRequest) GetUserEmail() string
```
GetUserEmail returns the UserEmail field value


```go
func (o *CheckDocumentAccessRequest) GetUserEmailOk() (*string, bool)
```
GetUserEmailOk returns a tuple with the UserEmail field value and a boolean
to check if the value has been set.


```go
func (o CheckDocumentAccessRequest) MarshalJSON() ([]byte, error)
```


```go
func (o *CheckDocumentAccessRequest) SetDatasource(v string)
```
SetDatasource sets field value


```go
func (o *CheckDocumentAccessRequest) SetDocId(v string)
```
SetDocId sets field value


```go
func (o *CheckDocumentAccessRequest) SetObjectType(v string)
```
SetObjectType sets field value


```go
func (o *CheckDocumentAccessRequest) SetUserEmail(v string)
```
SetUserEmail sets field value




### Type CheckDocumentAccessResponse
```go
type CheckDocumentAccessResponse struct {
	// If true, user has access to document for search
	HasAccess *bool `json:"hasAccess,omitempty"`
}
```
CheckDocumentAccessResponse Describes the response body of the
/checkdocumentaccess API call

### Functions

```go
func NewCheckDocumentAccessResponse() *CheckDocumentAccessResponse
```
NewCheckDocumentAccessResponse instantiates a new
CheckDocumentAccessResponse object This constructor will assign default
values to properties that have it defined, and makes sure properties
required by API are set, but the set of arguments will change when the set
of required properties is changed


```go
func NewCheckDocumentAccessResponseWithDefaults() *CheckDocumentAccessResponse
```
NewCheckDocumentAccessResponseWithDefaults instantiates a new
CheckDocumentAccessResponse object This constructor will only assign default
values to properties that have it defined, but it doesn't guarantee that
properties required by API are set



### Methods

```go
func (o *CheckDocumentAccessResponse) GetHasAccess() bool
```
GetHasAccess returns the HasAccess field value if set, zero value otherwise.


```go
func (o *CheckDocumentAccessResponse) GetHasAccessOk() (*bool, bool)
```
GetHasAccessOk returns a tuple with the HasAccess field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *CheckDocumentAccessResponse) HasHasAccess() bool
```
HasHasAccess returns a boolean if a field has been set.


```go
func (o CheckDocumentAccessResponse) MarshalJSON() ([]byte, error)
```


```go
func (o *CheckDocumentAccessResponse) SetHasAccess(v bool)
```
SetHasAccess gets a reference to the given bool and assigns it to the
HasAccess field.




### Type Configuration
```go
type Configuration struct {
	Host             string            `json:"host,omitempty"`
	Scheme           string            `json:"scheme,omitempty"`
	DefaultHeader    map[string]string `json:"defaultHeader,omitempty"`
	UserAgent        string            `json:"userAgent,omitempty"`
	Debug            bool              `json:"debug,omitempty"`
	Servers          ServerConfigurations
	OperationServers map[string]ServerConfigurations
	HTTPClient       *http.Client
}
```
Configuration stores the configuration of the API client

### Functions

```go
func NewConfiguration() *Configuration
```
NewConfiguration returns a new Configuration object



### Methods

```go
func (c *Configuration) AddDefaultHeader(key string, value string)
```
AddDefaultHeader adds a new HTTP header to the default header in the request


```go
func (c *Configuration) ServerURL(index int, variables map[string]string) (string, error)
```
ServerURL returns URL based on server settings


```go
func (c *Configuration) ServerURLWithContext(ctx context.Context, endpoint string) (string, error)
```
ServerURLWithContext returns a new server URL given an endpoint




### Type ConnectorType
```go
type ConnectorType string
```
ConnectorType The source from which document content was pulled, e.g.
an API crawl or browser history

### Constants
### API_CRAWL, BROWSER_CRAWL, BROWSER_HISTORY, BUILTIN, FEDERATED_SEARCH, PUSH_API, WEB_CRAWL, NATIVE_HISTORY
```go
API_CRAWL ConnectorType = "API_CRAWL"
BROWSER_CRAWL ConnectorType = "BROWSER_CRAWL"
BROWSER_HISTORY ConnectorType = "BROWSER_HISTORY"
BUILTIN ConnectorType = "BUILTIN"
FEDERATED_SEARCH ConnectorType = "FEDERATED_SEARCH"
PUSH_API ConnectorType = "PUSH_API"
WEB_CRAWL ConnectorType = "WEB_CRAWL"
NATIVE_HISTORY ConnectorType = "NATIVE_HISTORY"

```
List of ConnectorType



### Functions

```go
func NewConnectorTypeFromValue(v string) (*ConnectorType, error)
```
NewConnectorTypeFromValue returns a pointer to a valid ConnectorType for the
value passed as argument, or an error if the value passed is not allowed by
the enum



### Methods

```go
func (v ConnectorType) IsValid() bool
```
IsValid return true if the value is valid for the enum, false otherwise


```go
func (v ConnectorType) Ptr() *ConnectorType
```
Ptr returns reference to ConnectorType value


```go
func (v *ConnectorType) UnmarshalJSON(src []byte) error
```




### Type ContentDefinition
```go
type ContentDefinition struct {
	MimeType string `json:"mimeType"`
	// text content. Only one of textContent or binary content can be specified
	TextContent *string `json:"textContent,omitempty"`
	// base64 encoded binary content. Only one of textContent or binary content can be specified
	BinaryContent *string `json:"binaryContent,omitempty"`
}
```
ContentDefinition Describes text content or base64 encoded binary content

### Functions

```go
func NewContentDefinition(mimeType string) *ContentDefinition
```
NewContentDefinition instantiates a new ContentDefinition object This
constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed


```go
func NewContentDefinitionWithDefaults() *ContentDefinition
```
NewContentDefinitionWithDefaults instantiates a new ContentDefinition object
This constructor will only assign default values to properties that have it
defined, but it doesn't guarantee that properties required by API are set



### Methods

```go
func (o *ContentDefinition) GetBinaryContent() string
```
GetBinaryContent returns the BinaryContent field value if set, zero value
otherwise.


```go
func (o *ContentDefinition) GetBinaryContentOk() (*string, bool)
```
GetBinaryContentOk returns a tuple with the BinaryContent field value if
set, nil otherwise and a boolean to check if the value has been set.


```go
func (o *ContentDefinition) GetMimeType() string
```
GetMimeType returns the MimeType field value


```go
func (o *ContentDefinition) GetMimeTypeOk() (*string, bool)
```
GetMimeTypeOk returns a tuple with the MimeType field value and a boolean to
check if the value has been set.


```go
func (o *ContentDefinition) GetTextContent() string
```
GetTextContent returns the TextContent field value if set, zero value
otherwise.


```go
func (o *ContentDefinition) GetTextContentOk() (*string, bool)
```
GetTextContentOk returns a tuple with the TextContent field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *ContentDefinition) HasBinaryContent() bool
```
HasBinaryContent returns a boolean if a field has been set.


```go
func (o *ContentDefinition) HasTextContent() bool
```
HasTextContent returns a boolean if a field has been set.


```go
func (o ContentDefinition) MarshalJSON() ([]byte, error)
```


```go
func (o *ContentDefinition) SetBinaryContent(v string)
```
SetBinaryContent gets a reference to the given string and assigns it to the
BinaryContent field.


```go
func (o *ContentDefinition) SetMimeType(v string)
```
SetMimeType sets field value


```go
func (o *ContentDefinition) SetTextContent(v string)
```
SetTextContent gets a reference to the given string and assigns it to the
TextContent field.




### Type CustomDatasourceConfig
```go
type CustomDatasourceConfig struct {
	// Unique identifier of datasource instance to which this config applies.
	Name string `json:"name"`
	// Example text for what to search for in this datasource
	SuggestionText *string `json:"suggestionText,omitempty"`
	// The user-friendly instance label to display. If omitted, falls back to the title-cased `name`.
	DisplayName *string `json:"displayName,omitempty"`
	// The URL of the landing page for this datasource instance. Should point to the most useful page for users, not the company marketing page.
	HomeUrl *string `json:"homeUrl,omitempty"`
	// This only applies to WEB_CRAWL and BROWSER_CRAWL datasources. Defines the seed urls for crawling.
	CrawlerSeedUrls []string `json:"crawlerSeedUrls,omitempty"`
	// The URL to an image to be displayed as an icon for this datasource instance in dark mode. Must have a transparency mask. SVG are recommended over PNG. Public, scio-authenticated and Base64 encoded data URLs are all valid (but not third-party-authenticated URLs).
	IconDarkUrl *string `json:"iconDarkUrl,omitempty"`
	// The URL to an image to be displayed as an icon for this datasource instance. Must have a transparency mask. SVG are recommended over PNG. Public, scio-authenticated and Base64 encoded data URLs are all valid (but not third-party-authenticated URLs).
	IconUrl *string `json:"iconUrl,omitempty"`
	// The list of top-level `objectType`s for the datasource.
	ObjectDefinitions []ObjectDefinition `json:"objectDefinitions,omitempty"`
	// List of built-in facet types that should be hidden for the datasource.
	HideBuiltInFacets []string `json:"hideBuiltInFacets,omitempty"`
	// Regular expression that matches URLs of documents of the datasource instance. The behavior for multiple matches is non-deterministic. **Note: urlRegex is a required field for non-entity datasources (ie. datasources where isEntityDatasource is false). Please add a regex as specific as possible to this datasource instance.**
	UrlRegex *string `json:"urlRegex,omitempty"`
	// A list of regular expressions to apply to an arbitrary URL to transform it into a canonical URL for this datasource instance. Regexes are to be applied in the order specified in this list.
	CanonicalizingURLRegex []CanonicalizingRegexType `json:"canonicalizingURLRegex,omitempty"`
	// A list of regular expressions to apply to an arbitrary title to transform it into a title that will be displayed in the search results
	CanonicalizingTitleRegex []CanonicalizingRegexType `json:"canonicalizingTitleRegex,omitempty"`
	// A regex that identifies titles that should not be indexed
	RedlistTitleRegex *string        `json:"redlistTitleRegex,omitempty"`
	ConnectorType     *ConnectorType `json:"connectorType,omitempty"`
	// List of actions for this datasource instance that will show up in autocomplete and app card, e.g. \"Create new issue\" for jira
	Quicklinks []Quicklink `json:"quicklinks,omitempty"`
	// The name of a render config to use for displaying results from this datasource. Any well known datasource name may be used to render the same as that source, e.g. `web` or `gdrive`.
	RenderConfigPreset *string `json:"renderConfigPreset,omitempty"`
	// Aliases that can be used as `app` operator-values.
	Aliases []string `json:"aliases,omitempty"`
	// The type of this datasource. It is an important signal for relevance and must be specified and cannot be UNCATEGORIZED.
	DatasourceCategory *string `json:"datasourceCategory,omitempty"`
	// Whether or not this datasource is hosted on-premise.
	IsOnPrem *bool `json:"isOnPrem,omitempty"`
	// True if browser activity is able to report the correct URL for VIEW events. Set this to true if the URLs reported by Chrome are constant throughout each page load. Set this to false if the page has Javascript that modifies the URL during or after the load.
	TrustUrlRegexForViewActivity *bool `json:"trustUrlRegexForViewActivity,omitempty"`
	// If true, a utm_source query param will be added to outbound links to this datasource within Glean.
	IncludeUtmSource *bool `json:"includeUtmSource,omitempty"`
	// If the datasource uses another datasource for identity info, then the name of the datasource. The identity datasource must exist already.
	IdentityDatasourceName *string `json:"identityDatasourceName,omitempty"`
	// If the datasource uses a specific product access group, then the name of that group.
	ProductAccessGroup *string `json:"productAccessGroup,omitempty"`
	// whether email is used to reference users in document ACLs and in group memberships.
	IsUserReferencedByEmail *bool `json:"isUserReferencedByEmail,omitempty"`
	// True if this datasource is used to push custom entities.
	IsEntityDatasource *bool `json:"isEntityDatasource,omitempty"`
	// True if this datasource will be used for testing purpose only. Documents from such a datasource wouldn't have any effect on search rankings.
	IsTestDatasource *bool `json:"isTestDatasource,omitempty"`
}
```
CustomDatasourceConfig Structure describing config properties of a custom
datasource

### Functions

```go
func NewCustomDatasourceConfig(name string) *CustomDatasourceConfig
```
NewCustomDatasourceConfig instantiates a new CustomDatasourceConfig object
This constructor will assign default values to properties that have it
defined, and makes sure properties required by API are set, but the set of
arguments will change when the set of required properties is changed


```go
func NewCustomDatasourceConfigWithDefaults() *CustomDatasourceConfig
```
NewCustomDatasourceConfigWithDefaults instantiates a new
CustomDatasourceConfig object This constructor will only assign default
values to properties that have it defined, but it doesn't guarantee that
properties required by API are set



### Methods

```go
func (o *CustomDatasourceConfig) GetAliases() []string
```
GetAliases returns the Aliases field value if set, zero value otherwise.


```go
func (o *CustomDatasourceConfig) GetAliasesOk() ([]string, bool)
```
GetAliasesOk returns a tuple with the Aliases field value if set, nil
otherwise and a boolean to check if the value has been set.


```go
func (o *CustomDatasourceConfig) GetCanonicalizingTitleRegex() []CanonicalizingRegexType
```
GetCanonicalizingTitleRegex returns the CanonicalizingTitleRegex field value
if set, zero value otherwise.


```go
func (o *CustomDatasourceConfig) GetCanonicalizingTitleRegexOk() ([]CanonicalizingRegexType, bool)
```
GetCanonicalizingTitleRegexOk returns a tuple with the
CanonicalizingTitleRegex field value if set, nil otherwise and a boolean to
check if the value has been set.


```go
func (o *CustomDatasourceConfig) GetCanonicalizingURLRegex() []CanonicalizingRegexType
```
GetCanonicalizingURLRegex returns the CanonicalizingURLRegex field value if
set, zero value otherwise.


```go
func (o *CustomDatasourceConfig) GetCanonicalizingURLRegexOk() ([]CanonicalizingRegexType, bool)
```
GetCanonicalizingURLRegexOk returns a tuple with the CanonicalizingURLRegex
field value if set, nil otherwise and a boolean to check if the value has
been set.


```go
func (o *CustomDatasourceConfig) GetConnectorType() ConnectorType
```
GetConnectorType returns the ConnectorType field value if set, zero value
otherwise.


```go
func (o *CustomDatasourceConfig) GetConnectorTypeOk() (*ConnectorType, bool)
```
GetConnectorTypeOk returns a tuple with the ConnectorType field value if
set, nil otherwise and a boolean to check if the value has been set.


```go
func (o *CustomDatasourceConfig) GetCrawlerSeedUrls() []string
```
GetCrawlerSeedUrls returns the CrawlerSeedUrls field value if set, zero
value otherwise.


```go
func (o *CustomDatasourceConfig) GetCrawlerSeedUrlsOk() ([]string, bool)
```
GetCrawlerSeedUrlsOk returns a tuple with the CrawlerSeedUrls field value if
set, nil otherwise and a boolean to check if the value has been set.


```go
func (o *CustomDatasourceConfig) GetDatasourceCategory() string
```
GetDatasourceCategory returns the DatasourceCategory field value if set,
zero value otherwise.


```go
func (o *CustomDatasourceConfig) GetDatasourceCategoryOk() (*string, bool)
```
GetDatasourceCategoryOk returns a tuple with the DatasourceCategory field
value if set, nil otherwise and a boolean to check if the value has been
set.


```go
func (o *CustomDatasourceConfig) GetDisplayName() string
```
GetDisplayName returns the DisplayName field value if set, zero value
otherwise.


```go
func (o *CustomDatasourceConfig) GetDisplayNameOk() (*string, bool)
```
GetDisplayNameOk returns a tuple with the DisplayName field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *CustomDatasourceConfig) GetHideBuiltInFacets() []string
```
GetHideBuiltInFacets returns the HideBuiltInFacets field value if set,
zero value otherwise.


```go
func (o *CustomDatasourceConfig) GetHideBuiltInFacetsOk() ([]string, bool)
```
GetHideBuiltInFacetsOk returns a tuple with the HideBuiltInFacets field
value if set, nil otherwise and a boolean to check if the value has been
set.


```go
func (o *CustomDatasourceConfig) GetHomeUrl() string
```
GetHomeUrl returns the HomeUrl field value if set, zero value otherwise.


```go
func (o *CustomDatasourceConfig) GetHomeUrlOk() (*string, bool)
```
GetHomeUrlOk returns a tuple with the HomeUrl field value if set, nil
otherwise and a boolean to check if the value has been set.


```go
func (o *CustomDatasourceConfig) GetIconDarkUrl() string
```
GetIconDarkUrl returns the IconDarkUrl field value if set, zero value
otherwise.


```go
func (o *CustomDatasourceConfig) GetIconDarkUrlOk() (*string, bool)
```
GetIconDarkUrlOk returns a tuple with the IconDarkUrl field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *CustomDatasourceConfig) GetIconUrl() string
```
GetIconUrl returns the IconUrl field value if set, zero value otherwise.


```go
func (o *CustomDatasourceConfig) GetIconUrlOk() (*string, bool)
```
GetIconUrlOk returns a tuple with the IconUrl field value if set, nil
otherwise and a boolean to check if the value has been set.


```go
func (o *CustomDatasourceConfig) GetIdentityDatasourceName() string
```
GetIdentityDatasourceName returns the IdentityDatasourceName field value if
set, zero value otherwise.


```go
func (o *CustomDatasourceConfig) GetIdentityDatasourceNameOk() (*string, bool)
```
GetIdentityDatasourceNameOk returns a tuple with the IdentityDatasourceName
field value if set, nil otherwise and a boolean to check if the value has
been set.


```go
func (o *CustomDatasourceConfig) GetIncludeUtmSource() bool
```
GetIncludeUtmSource returns the IncludeUtmSource field value if set,
zero value otherwise.


```go
func (o *CustomDatasourceConfig) GetIncludeUtmSourceOk() (*bool, bool)
```
GetIncludeUtmSourceOk returns a tuple with the IncludeUtmSource field value
if set, nil otherwise and a boolean to check if the value has been set.


```go
func (o *CustomDatasourceConfig) GetIsEntityDatasource() bool
```
GetIsEntityDatasource returns the IsEntityDatasource field value if set,
zero value otherwise.


```go
func (o *CustomDatasourceConfig) GetIsEntityDatasourceOk() (*bool, bool)
```
GetIsEntityDatasourceOk returns a tuple with the IsEntityDatasource field
value if set, nil otherwise and a boolean to check if the value has been
set.


```go
func (o *CustomDatasourceConfig) GetIsOnPrem() bool
```
GetIsOnPrem returns the IsOnPrem field value if set, zero value otherwise.


```go
func (o *CustomDatasourceConfig) GetIsOnPremOk() (*bool, bool)
```
GetIsOnPremOk returns a tuple with the IsOnPrem field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *CustomDatasourceConfig) GetIsTestDatasource() bool
```
GetIsTestDatasource returns the IsTestDatasource field value if set,
zero value otherwise.


```go
func (o *CustomDatasourceConfig) GetIsTestDatasourceOk() (*bool, bool)
```
GetIsTestDatasourceOk returns a tuple with the IsTestDatasource field value
if set, nil otherwise and a boolean to check if the value has been set.


```go
func (o *CustomDatasourceConfig) GetIsUserReferencedByEmail() bool
```
GetIsUserReferencedByEmail returns the IsUserReferencedByEmail field value
if set, zero value otherwise.


```go
func (o *CustomDatasourceConfig) GetIsUserReferencedByEmailOk() (*bool, bool)
```
GetIsUserReferencedByEmailOk returns a tuple with the
IsUserReferencedByEmail field value if set, nil otherwise and a boolean to
check if the value has been set.


```go
func (o *CustomDatasourceConfig) GetName() string
```
GetName returns the Name field value


```go
func (o *CustomDatasourceConfig) GetNameOk() (*string, bool)
```
GetNameOk returns a tuple with the Name field value and a boolean to check
if the value has been set.


```go
func (o *CustomDatasourceConfig) GetObjectDefinitions() []ObjectDefinition
```
GetObjectDefinitions returns the ObjectDefinitions field value if set,
zero value otherwise.


```go
func (o *CustomDatasourceConfig) GetObjectDefinitionsOk() ([]ObjectDefinition, bool)
```
GetObjectDefinitionsOk returns a tuple with the ObjectDefinitions field
value if set, nil otherwise and a boolean to check if the value has been
set.


```go
func (o *CustomDatasourceConfig) GetProductAccessGroup() string
```
GetProductAccessGroup returns the ProductAccessGroup field value if set,
zero value otherwise.


```go
func (o *CustomDatasourceConfig) GetProductAccessGroupOk() (*string, bool)
```
GetProductAccessGroupOk returns a tuple with the ProductAccessGroup field
value if set, nil otherwise and a boolean to check if the value has been
set.


```go
func (o *CustomDatasourceConfig) GetQuicklinks() []Quicklink
```
GetQuicklinks returns the Quicklinks field value if set, zero value
otherwise.


```go
func (o *CustomDatasourceConfig) GetQuicklinksOk() ([]Quicklink, bool)
```
GetQuicklinksOk returns a tuple with the Quicklinks field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *CustomDatasourceConfig) GetRedlistTitleRegex() string
```
GetRedlistTitleRegex returns the RedlistTitleRegex field value if set,
zero value otherwise.


```go
func (o *CustomDatasourceConfig) GetRedlistTitleRegexOk() (*string, bool)
```
GetRedlistTitleRegexOk returns a tuple with the RedlistTitleRegex field
value if set, nil otherwise and a boolean to check if the value has been
set.


```go
func (o *CustomDatasourceConfig) GetRenderConfigPreset() string
```
GetRenderConfigPreset returns the RenderConfigPreset field value if set,
zero value otherwise.


```go
func (o *CustomDatasourceConfig) GetRenderConfigPresetOk() (*string, bool)
```
GetRenderConfigPresetOk returns a tuple with the RenderConfigPreset field
value if set, nil otherwise and a boolean to check if the value has been
set.


```go
func (o *CustomDatasourceConfig) GetSuggestionText() string
```
GetSuggestionText returns the SuggestionText field value if set, zero value
otherwise.


```go
func (o *CustomDatasourceConfig) GetSuggestionTextOk() (*string, bool)
```
GetSuggestionTextOk returns a tuple with the SuggestionText field value if
set, nil otherwise and a boolean to check if the value has been set.


```go
func (o *CustomDatasourceConfig) GetTrustUrlRegexForViewActivity() bool
```
GetTrustUrlRegexForViewActivity returns the TrustUrlRegexForViewActivity
field value if set, zero value otherwise.


```go
func (o *CustomDatasourceConfig) GetTrustUrlRegexForViewActivityOk() (*bool, bool)
```
GetTrustUrlRegexForViewActivityOk returns a tuple with the
TrustUrlRegexForViewActivity field value if set, nil otherwise and a boolean
to check if the value has been set.


```go
func (o *CustomDatasourceConfig) GetUrlRegex() string
```
GetUrlRegex returns the UrlRegex field value if set, zero value otherwise.


```go
func (o *CustomDatasourceConfig) GetUrlRegexOk() (*string, bool)
```
GetUrlRegexOk returns a tuple with the UrlRegex field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *CustomDatasourceConfig) HasAliases() bool
```
HasAliases returns a boolean if a field has been set.


```go
func (o *CustomDatasourceConfig) HasCanonicalizingTitleRegex() bool
```
HasCanonicalizingTitleRegex returns a boolean if a field has been set.


```go
func (o *CustomDatasourceConfig) HasCanonicalizingURLRegex() bool
```
HasCanonicalizingURLRegex returns a boolean if a field has been set.


```go
func (o *CustomDatasourceConfig) HasConnectorType() bool
```
HasConnectorType returns a boolean if a field has been set.


```go
func (o *CustomDatasourceConfig) HasCrawlerSeedUrls() bool
```
HasCrawlerSeedUrls returns a boolean if a field has been set.


```go
func (o *CustomDatasourceConfig) HasDatasourceCategory() bool
```
HasDatasourceCategory returns a boolean if a field has been set.


```go
func (o *CustomDatasourceConfig) HasDisplayName() bool
```
HasDisplayName returns a boolean if a field has been set.


```go
func (o *CustomDatasourceConfig) HasHideBuiltInFacets() bool
```
HasHideBuiltInFacets returns a boolean if a field has been set.


```go
func (o *CustomDatasourceConfig) HasHomeUrl() bool
```
HasHomeUrl returns a boolean if a field has been set.


```go
func (o *CustomDatasourceConfig) HasIconDarkUrl() bool
```
HasIconDarkUrl returns a boolean if a field has been set.


```go
func (o *CustomDatasourceConfig) HasIconUrl() bool
```
HasIconUrl returns a boolean if a field has been set.


```go
func (o *CustomDatasourceConfig) HasIdentityDatasourceName() bool
```
HasIdentityDatasourceName returns a boolean if a field has been set.


```go
func (o *CustomDatasourceConfig) HasIncludeUtmSource() bool
```
HasIncludeUtmSource returns a boolean if a field has been set.


```go
func (o *CustomDatasourceConfig) HasIsEntityDatasource() bool
```
HasIsEntityDatasource returns a boolean if a field has been set.


```go
func (o *CustomDatasourceConfig) HasIsOnPrem() bool
```
HasIsOnPrem returns a boolean if a field has been set.


```go
func (o *CustomDatasourceConfig) HasIsTestDatasource() bool
```
HasIsTestDatasource returns a boolean if a field has been set.


```go
func (o *CustomDatasourceConfig) HasIsUserReferencedByEmail() bool
```
HasIsUserReferencedByEmail returns a boolean if a field has been set.


```go
func (o *CustomDatasourceConfig) HasObjectDefinitions() bool
```
HasObjectDefinitions returns a boolean if a field has been set.


```go
func (o *CustomDatasourceConfig) HasProductAccessGroup() bool
```
HasProductAccessGroup returns a boolean if a field has been set.


```go
func (o *CustomDatasourceConfig) HasQuicklinks() bool
```
HasQuicklinks returns a boolean if a field has been set.


```go
func (o *CustomDatasourceConfig) HasRedlistTitleRegex() bool
```
HasRedlistTitleRegex returns a boolean if a field has been set.


```go
func (o *CustomDatasourceConfig) HasRenderConfigPreset() bool
```
HasRenderConfigPreset returns a boolean if a field has been set.


```go
func (o *CustomDatasourceConfig) HasSuggestionText() bool
```
HasSuggestionText returns a boolean if a field has been set.


```go
func (o *CustomDatasourceConfig) HasTrustUrlRegexForViewActivity() bool
```
HasTrustUrlRegexForViewActivity returns a boolean if a field has been set.


```go
func (o *CustomDatasourceConfig) HasUrlRegex() bool
```
HasUrlRegex returns a boolean if a field has been set.


```go
func (o CustomDatasourceConfig) MarshalJSON() ([]byte, error)
```


```go
func (o *CustomDatasourceConfig) SetAliases(v []string)
```
SetAliases gets a reference to the given []string and assigns it to the
Aliases field.


```go
func (o *CustomDatasourceConfig) SetCanonicalizingTitleRegex(v []CanonicalizingRegexType)
```
SetCanonicalizingTitleRegex gets a reference to the given
[]CanonicalizingRegexType and assigns it to the CanonicalizingTitleRegex
field.


```go
func (o *CustomDatasourceConfig) SetCanonicalizingURLRegex(v []CanonicalizingRegexType)
```
SetCanonicalizingURLRegex gets a reference to the given
[]CanonicalizingRegexType and assigns it to the CanonicalizingURLRegex
field.


```go
func (o *CustomDatasourceConfig) SetConnectorType(v ConnectorType)
```
SetConnectorType gets a reference to the given ConnectorType and assigns it
to the ConnectorType field.


```go
func (o *CustomDatasourceConfig) SetCrawlerSeedUrls(v []string)
```
SetCrawlerSeedUrls gets a reference to the given []string and assigns it to
the CrawlerSeedUrls field.


```go
func (o *CustomDatasourceConfig) SetDatasourceCategory(v string)
```
SetDatasourceCategory gets a reference to the given string and assigns it to
the DatasourceCategory field.


```go
func (o *CustomDatasourceConfig) SetDisplayName(v string)
```
SetDisplayName gets a reference to the given string and assigns it to the
DisplayName field.


```go
func (o *CustomDatasourceConfig) SetHideBuiltInFacets(v []string)
```
SetHideBuiltInFacets gets a reference to the given []string and assigns it
to the HideBuiltInFacets field.


```go
func (o *CustomDatasourceConfig) SetHomeUrl(v string)
```
SetHomeUrl gets a reference to the given string and assigns it to the
HomeUrl field.


```go
func (o *CustomDatasourceConfig) SetIconDarkUrl(v string)
```
SetIconDarkUrl gets a reference to the given string and assigns it to the
IconDarkUrl field.


```go
func (o *CustomDatasourceConfig) SetIconUrl(v string)
```
SetIconUrl gets a reference to the given string and assigns it to the
IconUrl field.


```go
func (o *CustomDatasourceConfig) SetIdentityDatasourceName(v string)
```
SetIdentityDatasourceName gets a reference to the given string and assigns
it to the IdentityDatasourceName field.


```go
func (o *CustomDatasourceConfig) SetIncludeUtmSource(v bool)
```
SetIncludeUtmSource gets a reference to the given bool and assigns it to the
IncludeUtmSource field.


```go
func (o *CustomDatasourceConfig) SetIsEntityDatasource(v bool)
```
SetIsEntityDatasource gets a reference to the given bool and assigns it to
the IsEntityDatasource field.


```go
func (o *CustomDatasourceConfig) SetIsOnPrem(v bool)
```
SetIsOnPrem gets a reference to the given bool and assigns it to the
IsOnPrem field.


```go
func (o *CustomDatasourceConfig) SetIsTestDatasource(v bool)
```
SetIsTestDatasource gets a reference to the given bool and assigns it to the
IsTestDatasource field.


```go
func (o *CustomDatasourceConfig) SetIsUserReferencedByEmail(v bool)
```
SetIsUserReferencedByEmail gets a reference to the given bool and assigns it
to the IsUserReferencedByEmail field.


```go
func (o *CustomDatasourceConfig) SetName(v string)
```
SetName sets field value


```go
func (o *CustomDatasourceConfig) SetObjectDefinitions(v []ObjectDefinition)
```
SetObjectDefinitions gets a reference to the given []ObjectDefinition and
assigns it to the ObjectDefinitions field.


```go
func (o *CustomDatasourceConfig) SetProductAccessGroup(v string)
```
SetProductAccessGroup gets a reference to the given string and assigns it to
the ProductAccessGroup field.


```go
func (o *CustomDatasourceConfig) SetQuicklinks(v []Quicklink)
```
SetQuicklinks gets a reference to the given []Quicklink and assigns it to
the Quicklinks field.


```go
func (o *CustomDatasourceConfig) SetRedlistTitleRegex(v string)
```
SetRedlistTitleRegex gets a reference to the given string and assigns it to
the RedlistTitleRegex field.


```go
func (o *CustomDatasourceConfig) SetRenderConfigPreset(v string)
```
SetRenderConfigPreset gets a reference to the given string and assigns it to
the RenderConfigPreset field.


```go
func (o *CustomDatasourceConfig) SetSuggestionText(v string)
```
SetSuggestionText gets a reference to the given string and assigns it to the
SuggestionText field.


```go
func (o *CustomDatasourceConfig) SetTrustUrlRegexForViewActivity(v bool)
```
SetTrustUrlRegexForViewActivity gets a reference to the given bool and
assigns it to the TrustUrlRegexForViewActivity field.


```go
func (o *CustomDatasourceConfig) SetUrlRegex(v string)
```
SetUrlRegex gets a reference to the given string and assigns it to the
UrlRegex field.




### Type CustomDatasourceConfigAllOf
```go
type CustomDatasourceConfigAllOf struct {
	// If the datasource uses another datasource for identity info, then the name of the datasource. The identity datasource must exist already.
	IdentityDatasourceName *string `json:"identityDatasourceName,omitempty"`
	// If the datasource uses a specific product access group, then the name of that group.
	ProductAccessGroup *string `json:"productAccessGroup,omitempty"`
	// whether email is used to reference users in document ACLs and in group memberships.
	IsUserReferencedByEmail *bool `json:"isUserReferencedByEmail,omitempty"`
	// True if this datasource is used to push custom entities.
	IsEntityDatasource *bool `json:"isEntityDatasource,omitempty"`
	// True if this datasource will be used for testing purpose only. Documents from such a datasource wouldn't have any effect on search rankings.
	IsTestDatasource *bool `json:"isTestDatasource,omitempty"`
}
```
CustomDatasourceConfigAllOf struct for CustomDatasourceConfigAllOf

### Functions

```go
func NewCustomDatasourceConfigAllOf() *CustomDatasourceConfigAllOf
```
NewCustomDatasourceConfigAllOf instantiates a new
CustomDatasourceConfigAllOf object This constructor will assign default
values to properties that have it defined, and makes sure properties
required by API are set, but the set of arguments will change when the set
of required properties is changed


```go
func NewCustomDatasourceConfigAllOfWithDefaults() *CustomDatasourceConfigAllOf
```
NewCustomDatasourceConfigAllOfWithDefaults instantiates a new
CustomDatasourceConfigAllOf object This constructor will only assign default
values to properties that have it defined, but it doesn't guarantee that
properties required by API are set



### Methods

```go
func (o *CustomDatasourceConfigAllOf) GetIdentityDatasourceName() string
```
GetIdentityDatasourceName returns the IdentityDatasourceName field value if
set, zero value otherwise.


```go
func (o *CustomDatasourceConfigAllOf) GetIdentityDatasourceNameOk() (*string, bool)
```
GetIdentityDatasourceNameOk returns a tuple with the IdentityDatasourceName
field value if set, nil otherwise and a boolean to check if the value has
been set.


```go
func (o *CustomDatasourceConfigAllOf) GetIsEntityDatasource() bool
```
GetIsEntityDatasource returns the IsEntityDatasource field value if set,
zero value otherwise.


```go
func (o *CustomDatasourceConfigAllOf) GetIsEntityDatasourceOk() (*bool, bool)
```
GetIsEntityDatasourceOk returns a tuple with the IsEntityDatasource field
value if set, nil otherwise and a boolean to check if the value has been
set.


```go
func (o *CustomDatasourceConfigAllOf) GetIsTestDatasource() bool
```
GetIsTestDatasource returns the IsTestDatasource field value if set,
zero value otherwise.


```go
func (o *CustomDatasourceConfigAllOf) GetIsTestDatasourceOk() (*bool, bool)
```
GetIsTestDatasourceOk returns a tuple with the IsTestDatasource field value
if set, nil otherwise and a boolean to check if the value has been set.


```go
func (o *CustomDatasourceConfigAllOf) GetIsUserReferencedByEmail() bool
```
GetIsUserReferencedByEmail returns the IsUserReferencedByEmail field value
if set, zero value otherwise.


```go
func (o *CustomDatasourceConfigAllOf) GetIsUserReferencedByEmailOk() (*bool, bool)
```
GetIsUserReferencedByEmailOk returns a tuple with the
IsUserReferencedByEmail field value if set, nil otherwise and a boolean to
check if the value has been set.


```go
func (o *CustomDatasourceConfigAllOf) GetProductAccessGroup() string
```
GetProductAccessGroup returns the ProductAccessGroup field value if set,
zero value otherwise.


```go
func (o *CustomDatasourceConfigAllOf) GetProductAccessGroupOk() (*string, bool)
```
GetProductAccessGroupOk returns a tuple with the ProductAccessGroup field
value if set, nil otherwise and a boolean to check if the value has been
set.


```go
func (o *CustomDatasourceConfigAllOf) HasIdentityDatasourceName() bool
```
HasIdentityDatasourceName returns a boolean if a field has been set.


```go
func (o *CustomDatasourceConfigAllOf) HasIsEntityDatasource() bool
```
HasIsEntityDatasource returns a boolean if a field has been set.


```go
func (o *CustomDatasourceConfigAllOf) HasIsTestDatasource() bool
```
HasIsTestDatasource returns a boolean if a field has been set.


```go
func (o *CustomDatasourceConfigAllOf) HasIsUserReferencedByEmail() bool
```
HasIsUserReferencedByEmail returns a boolean if a field has been set.


```go
func (o *CustomDatasourceConfigAllOf) HasProductAccessGroup() bool
```
HasProductAccessGroup returns a boolean if a field has been set.


```go
func (o CustomDatasourceConfigAllOf) MarshalJSON() ([]byte, error)
```


```go
func (o *CustomDatasourceConfigAllOf) SetIdentityDatasourceName(v string)
```
SetIdentityDatasourceName gets a reference to the given string and assigns
it to the IdentityDatasourceName field.


```go
func (o *CustomDatasourceConfigAllOf) SetIsEntityDatasource(v bool)
```
SetIsEntityDatasource gets a reference to the given bool and assigns it to
the IsEntityDatasource field.


```go
func (o *CustomDatasourceConfigAllOf) SetIsTestDatasource(v bool)
```
SetIsTestDatasource gets a reference to the given bool and assigns it to the
IsTestDatasource field.


```go
func (o *CustomDatasourceConfigAllOf) SetIsUserReferencedByEmail(v bool)
```
SetIsUserReferencedByEmail gets a reference to the given bool and assigns it
to the IsUserReferencedByEmail field.


```go
func (o *CustomDatasourceConfigAllOf) SetProductAccessGroup(v string)
```
SetProductAccessGroup gets a reference to the given string and assigns it to
the ProductAccessGroup field.




### Type CustomProperty
```go
type CustomProperty struct {
	Name *string `json:"name,omitempty"`
	// Must either be a string or an array of strings. An integer, boolean, etc. is not valid. When OpenAPI Generator supports `oneOf`, we can semantically enforce this.
	Value interface{} `json:"value,omitempty"`
}
```
CustomProperty Describes the custom properties of the object.

### Functions

```go
func NewCustomProperty() *CustomProperty
```
NewCustomProperty instantiates a new CustomProperty object This constructor
will assign default values to properties that have it defined, and makes
sure properties required by API are set, but the set of arguments will
change when the set of required properties is changed


```go
func NewCustomPropertyWithDefaults() *CustomProperty
```
NewCustomPropertyWithDefaults instantiates a new CustomProperty object
This constructor will only assign default values to properties that have it
defined, but it doesn't guarantee that properties required by API are set



### Methods

```go
func (o *CustomProperty) GetName() string
```
GetName returns the Name field value if set, zero value otherwise.


```go
func (o *CustomProperty) GetNameOk() (*string, bool)
```
GetNameOk returns a tuple with the Name field value if set, nil otherwise
and a boolean to check if the value has been set.


```go
func (o *CustomProperty) GetValue() interface{}
```
GetValue returns the Value field value if set, zero value otherwise (both if
not set or set to explicit null).


```go
func (o *CustomProperty) GetValueOk() (*interface{}, bool)
```
GetValueOk returns a tuple with the Value field value if set, nil otherwise
and a boolean to check if the value has been set. NOTE: If the value is an
explicit nil, `nil, true` will be returned


```go
func (o *CustomProperty) HasName() bool
```
HasName returns a boolean if a field has been set.


```go
func (o *CustomProperty) HasValue() bool
```
HasValue returns a boolean if a field has been set.


```go
func (o CustomProperty) MarshalJSON() ([]byte, error)
```


```go
func (o *CustomProperty) SetName(v string)
```
SetName gets a reference to the given string and assigns it to the Name
field.


```go
func (o *CustomProperty) SetValue(v interface{})
```
SetValue gets a reference to the given interface{} and assigns it to the
Value field.




### Type DatasourceBulkMembershipDefinition
```go
type DatasourceBulkMembershipDefinition struct {
	// If the member is a user, then the email or datasource id for the user
	MemberUserId *string `json:"memberUserId,omitempty"`
	// If the member is a group, then the name of the member group
	MemberGroupName *string `json:"memberGroupName,omitempty"`
}
```
DatasourceBulkMembershipDefinition describes the membership row of a group
in the bulk uploaded. Only one of memberUserId and memberGroupName can be
specified.

### Functions

```go
func NewDatasourceBulkMembershipDefinition() *DatasourceBulkMembershipDefinition
```
NewDatasourceBulkMembershipDefinition instantiates a new
DatasourceBulkMembershipDefinition object This constructor will assign
default values to properties that have it defined, and makes sure properties
required by API are set, but the set of arguments will change when the set
of required properties is changed


```go
func NewDatasourceBulkMembershipDefinitionWithDefaults() *DatasourceBulkMembershipDefinition
```
NewDatasourceBulkMembershipDefinitionWithDefaults instantiates a new
DatasourceBulkMembershipDefinition object This constructor will only assign
default values to properties that have it defined, but it doesn't guarantee
that properties required by API are set



### Methods

```go
func (o *DatasourceBulkMembershipDefinition) GetMemberGroupName() string
```
GetMemberGroupName returns the MemberGroupName field value if set, zero
value otherwise.


```go
func (o *DatasourceBulkMembershipDefinition) GetMemberGroupNameOk() (*string, bool)
```
GetMemberGroupNameOk returns a tuple with the MemberGroupName field value if
set, nil otherwise and a boolean to check if the value has been set.


```go
func (o *DatasourceBulkMembershipDefinition) GetMemberUserId() string
```
GetMemberUserId returns the MemberUserId field value if set, zero value
otherwise.


```go
func (o *DatasourceBulkMembershipDefinition) GetMemberUserIdOk() (*string, bool)
```
GetMemberUserIdOk returns a tuple with the MemberUserId field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *DatasourceBulkMembershipDefinition) HasMemberGroupName() bool
```
HasMemberGroupName returns a boolean if a field has been set.


```go
func (o *DatasourceBulkMembershipDefinition) HasMemberUserId() bool
```
HasMemberUserId returns a boolean if a field has been set.


```go
func (o DatasourceBulkMembershipDefinition) MarshalJSON() ([]byte, error)
```


```go
func (o *DatasourceBulkMembershipDefinition) SetMemberGroupName(v string)
```
SetMemberGroupName gets a reference to the given string and assigns it to
the MemberGroupName field.


```go
func (o *DatasourceBulkMembershipDefinition) SetMemberUserId(v string)
```
SetMemberUserId gets a reference to the given string and assigns it to the
MemberUserId field.




### Type DatasourceConfigList
```go
type DatasourceConfigList struct {
	// Datasource configuration.
	DatasourceConfig []SharedDatasourceConfig `json:"datasourceConfig"`
}
```
DatasourceConfigList List of datasource configurations.

### Functions

```go
func NewDatasourceConfigList(datasourceConfig []SharedDatasourceConfig) *DatasourceConfigList
```
NewDatasourceConfigList instantiates a new DatasourceConfigList object This
constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed


```go
func NewDatasourceConfigListWithDefaults() *DatasourceConfigList
```
NewDatasourceConfigListWithDefaults instantiates a new DatasourceConfigList
object This constructor will only assign default values to properties that
have it defined, but it doesn't guarantee that properties required by API
are set



### Methods

```go
func (o *DatasourceConfigList) GetDatasourceConfig() []SharedDatasourceConfig
```
GetDatasourceConfig returns the DatasourceConfig field value


```go
func (o *DatasourceConfigList) GetDatasourceConfigOk() ([]SharedDatasourceConfig, bool)
```
GetDatasourceConfigOk returns a tuple with the DatasourceConfig field value
and a boolean to check if the value has been set.


```go
func (o DatasourceConfigList) MarshalJSON() ([]byte, error)
```


```go
func (o *DatasourceConfigList) SetDatasourceConfig(v []SharedDatasourceConfig)
```
SetDatasourceConfig sets field value




### Type DatasourceGroupDefinition
```go
type DatasourceGroupDefinition struct {
	// name of the group. Should be unique among all groups for the datasource, and cannot have spaces.
	Name string `json:"name"`
}
```
DatasourceGroupDefinition describes a group in the datasource

### Functions

```go
func NewDatasourceGroupDefinition(name string) *DatasourceGroupDefinition
```
NewDatasourceGroupDefinition instantiates a new DatasourceGroupDefinition
object This constructor will assign default values to properties that have
it defined, and makes sure properties required by API are set, but the set
of arguments will change when the set of required properties is changed


```go
func NewDatasourceGroupDefinitionWithDefaults() *DatasourceGroupDefinition
```
NewDatasourceGroupDefinitionWithDefaults instantiates a new
DatasourceGroupDefinition object This constructor will only assign default
values to properties that have it defined, but it doesn't guarantee that
properties required by API are set



### Methods

```go
func (o *DatasourceGroupDefinition) GetName() string
```
GetName returns the Name field value


```go
func (o *DatasourceGroupDefinition) GetNameOk() (*string, bool)
```
GetNameOk returns a tuple with the Name field value and a boolean to check
if the value has been set.


```go
func (o DatasourceGroupDefinition) MarshalJSON() ([]byte, error)
```


```go
func (o *DatasourceGroupDefinition) SetName(v string)
```
SetName sets field value




### Type DatasourceMembershipDefinition
```go
type DatasourceMembershipDefinition struct {
	// The group for which the membership is specified
	GroupName string `json:"groupName"`
	// If the member is a user, then the email or datasource id for the user
	MemberUserId *string `json:"memberUserId,omitempty"`
	// If the member is a group, then the name of the member group
	MemberGroupName *string `json:"memberGroupName,omitempty"`
}
```
DatasourceMembershipDefinition describes the membership row of a group.
Only one of memberUserId and memberGroupName can be specified.

### Functions

```go
func NewDatasourceMembershipDefinition(groupName string) *DatasourceMembershipDefinition
```
NewDatasourceMembershipDefinition instantiates a new
DatasourceMembershipDefinition object This constructor will assign default
values to properties that have it defined, and makes sure properties
required by API are set, but the set of arguments will change when the set
of required properties is changed


```go
func NewDatasourceMembershipDefinitionWithDefaults() *DatasourceMembershipDefinition
```
NewDatasourceMembershipDefinitionWithDefaults instantiates a new
DatasourceMembershipDefinition object This constructor will only assign
default values to properties that have it defined, but it doesn't guarantee
that properties required by API are set



### Methods

```go
func (o *DatasourceMembershipDefinition) GetGroupName() string
```
GetGroupName returns the GroupName field value


```go
func (o *DatasourceMembershipDefinition) GetGroupNameOk() (*string, bool)
```
GetGroupNameOk returns a tuple with the GroupName field value and a boolean
to check if the value has been set.


```go
func (o *DatasourceMembershipDefinition) GetMemberGroupName() string
```
GetMemberGroupName returns the MemberGroupName field value if set, zero
value otherwise.


```go
func (o *DatasourceMembershipDefinition) GetMemberGroupNameOk() (*string, bool)
```
GetMemberGroupNameOk returns a tuple with the MemberGroupName field value if
set, nil otherwise and a boolean to check if the value has been set.


```go
func (o *DatasourceMembershipDefinition) GetMemberUserId() string
```
GetMemberUserId returns the MemberUserId field value if set, zero value
otherwise.


```go
func (o *DatasourceMembershipDefinition) GetMemberUserIdOk() (*string, bool)
```
GetMemberUserIdOk returns a tuple with the MemberUserId field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *DatasourceMembershipDefinition) HasMemberGroupName() bool
```
HasMemberGroupName returns a boolean if a field has been set.


```go
func (o *DatasourceMembershipDefinition) HasMemberUserId() bool
```
HasMemberUserId returns a boolean if a field has been set.


```go
func (o DatasourceMembershipDefinition) MarshalJSON() ([]byte, error)
```


```go
func (o *DatasourceMembershipDefinition) SetGroupName(v string)
```
SetGroupName sets field value


```go
func (o *DatasourceMembershipDefinition) SetMemberGroupName(v string)
```
SetMemberGroupName gets a reference to the given string and assigns it to
the MemberGroupName field.


```go
func (o *DatasourceMembershipDefinition) SetMemberUserId(v string)
```
SetMemberUserId gets a reference to the given string and assigns it to the
MemberUserId field.




### Type DatasourceProfile
```go
type DatasourceProfile struct {
	// The datasource the profile is of.
	Datasource string `json:"datasource"`
	// The display name of the person in the given datasource.
	Handle string `json:"handle"`
	// URL to view the user's profile.
	Url *string `json:"url,omitempty"`
	// A deep link, if available, into the datasource's native application for the user's platform (i.e. slack://...).
	NativeAppUrl *string `json:"nativeAppUrl,omitempty"`
	// For internal use only. True iff the data source profile was manually added by a user from within Glean (aka not from the original data source)
	IsUserGenerated *bool `json:"isUserGenerated,omitempty"`
}
```
DatasourceProfile struct for DatasourceProfile

### Functions

```go
func NewDatasourceProfile(datasource string, handle string) *DatasourceProfile
```
NewDatasourceProfile instantiates a new DatasourceProfile object This
constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed


```go
func NewDatasourceProfileWithDefaults() *DatasourceProfile
```
NewDatasourceProfileWithDefaults instantiates a new DatasourceProfile object
This constructor will only assign default values to properties that have it
defined, but it doesn't guarantee that properties required by API are set



### Methods

```go
func (o *DatasourceProfile) GetDatasource() string
```
GetDatasource returns the Datasource field value


```go
func (o *DatasourceProfile) GetDatasourceOk() (*string, bool)
```
GetDatasourceOk returns a tuple with the Datasource field value and a
boolean to check if the value has been set.


```go
func (o *DatasourceProfile) GetHandle() string
```
GetHandle returns the Handle field value


```go
func (o *DatasourceProfile) GetHandleOk() (*string, bool)
```
GetHandleOk returns a tuple with the Handle field value and a boolean to
check if the value has been set.


```go
func (o *DatasourceProfile) GetIsUserGenerated() bool
```
GetIsUserGenerated returns the IsUserGenerated field value if set, zero
value otherwise.


```go
func (o *DatasourceProfile) GetIsUserGeneratedOk() (*bool, bool)
```
GetIsUserGeneratedOk returns a tuple with the IsUserGenerated field value if
set, nil otherwise and a boolean to check if the value has been set.


```go
func (o *DatasourceProfile) GetNativeAppUrl() string
```
GetNativeAppUrl returns the NativeAppUrl field value if set, zero value
otherwise.


```go
func (o *DatasourceProfile) GetNativeAppUrlOk() (*string, bool)
```
GetNativeAppUrlOk returns a tuple with the NativeAppUrl field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *DatasourceProfile) GetUrl() string
```
GetUrl returns the Url field value if set, zero value otherwise.


```go
func (o *DatasourceProfile) GetUrlOk() (*string, bool)
```
GetUrlOk returns a tuple with the Url field value if set, nil otherwise and
a boolean to check if the value has been set.


```go
func (o *DatasourceProfile) HasIsUserGenerated() bool
```
HasIsUserGenerated returns a boolean if a field has been set.


```go
func (o *DatasourceProfile) HasNativeAppUrl() bool
```
HasNativeAppUrl returns a boolean if a field has been set.


```go
func (o *DatasourceProfile) HasUrl() bool
```
HasUrl returns a boolean if a field has been set.


```go
func (o DatasourceProfile) MarshalJSON() ([]byte, error)
```


```go
func (o *DatasourceProfile) SetDatasource(v string)
```
SetDatasource sets field value


```go
func (o *DatasourceProfile) SetHandle(v string)
```
SetHandle sets field value


```go
func (o *DatasourceProfile) SetIsUserGenerated(v bool)
```
SetIsUserGenerated gets a reference to the given bool and assigns it to the
IsUserGenerated field.


```go
func (o *DatasourceProfile) SetNativeAppUrl(v string)
```
SetNativeAppUrl gets a reference to the given string and assigns it to the
NativeAppUrl field.


```go
func (o *DatasourceProfile) SetUrl(v string)
```
SetUrl gets a reference to the given string and assigns it to the Url field.




### Type DatasourceUserDefinition
```go
type DatasourceUserDefinition struct {
	Email *string `json:"email,omitempty"`
	// To be supplied if the user id in the datasource is not the email
	UserId *string `json:"userId,omitempty"`
	Name   string  `json:"name"`
	// set to false if the user is a former employee or a bot
	IsActive *bool `json:"isActive,omitempty"`
}
```
DatasourceUserDefinition describes a user in the datasource

### Functions

```go
func NewDatasourceUserDefinition(name string) *DatasourceUserDefinition
```
NewDatasourceUserDefinition instantiates a new DatasourceUserDefinition
object This constructor will assign default values to properties that have
it defined, and makes sure properties required by API are set, but the set
of arguments will change when the set of required properties is changed


```go
func NewDatasourceUserDefinitionWithDefaults() *DatasourceUserDefinition
```
NewDatasourceUserDefinitionWithDefaults instantiates a new
DatasourceUserDefinition object This constructor will only assign default
values to properties that have it defined, but it doesn't guarantee that
properties required by API are set



### Methods

```go
func (o *DatasourceUserDefinition) GetEmail() string
```
GetEmail returns the Email field value if set, zero value otherwise.


```go
func (o *DatasourceUserDefinition) GetEmailOk() (*string, bool)
```
GetEmailOk returns a tuple with the Email field value if set, nil otherwise
and a boolean to check if the value has been set.


```go
func (o *DatasourceUserDefinition) GetIsActive() bool
```
GetIsActive returns the IsActive field value if set, zero value otherwise.


```go
func (o *DatasourceUserDefinition) GetIsActiveOk() (*bool, bool)
```
GetIsActiveOk returns a tuple with the IsActive field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *DatasourceUserDefinition) GetName() string
```
GetName returns the Name field value


```go
func (o *DatasourceUserDefinition) GetNameOk() (*string, bool)
```
GetNameOk returns a tuple with the Name field value and a boolean to check
if the value has been set.


```go
func (o *DatasourceUserDefinition) GetUserId() string
```
GetUserId returns the UserId field value if set, zero value otherwise.


```go
func (o *DatasourceUserDefinition) GetUserIdOk() (*string, bool)
```
GetUserIdOk returns a tuple with the UserId field value if set, nil
otherwise and a boolean to check if the value has been set.


```go
func (o *DatasourceUserDefinition) HasEmail() bool
```
HasEmail returns a boolean if a field has been set.


```go
func (o *DatasourceUserDefinition) HasIsActive() bool
```
HasIsActive returns a boolean if a field has been set.


```go
func (o *DatasourceUserDefinition) HasUserId() bool
```
HasUserId returns a boolean if a field has been set.


```go
func (o DatasourceUserDefinition) MarshalJSON() ([]byte, error)
```


```go
func (o *DatasourceUserDefinition) SetEmail(v string)
```
SetEmail gets a reference to the given string and assigns it to the Email
field.


```go
func (o *DatasourceUserDefinition) SetIsActive(v bool)
```
SetIsActive gets a reference to the given bool and assigns it to the
IsActive field.


```go
func (o *DatasourceUserDefinition) SetName(v string)
```
SetName sets field value


```go
func (o *DatasourceUserDefinition) SetUserId(v string)
```
SetUserId gets a reference to the given string and assigns it to the UserId
field.




### Type DatasourcesApiService
```go
type DatasourcesApiService service
```
DatasourcesApiService DatasourcesApi service

### Methods

```go
func (a *DatasourcesApiService) AdddatasourcePost(ctx context.Context) ApiAdddatasourcePostRequest
```
AdddatasourcePost Add datasource

API to register a custom datasource type and its schema.

    @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
    @return ApiAdddatasourcePostRequest


```go
func (a *DatasourcesApiService) AdddatasourcePostExecute(r ApiAdddatasourcePostRequest) (*http.Response, error)
```
Execute executes the request


```go
func (a *DatasourcesApiService) GetdatasourceconfigPost(ctx context.Context) ApiGetdatasourceconfigPostRequest
```
GetdatasourceconfigPost Get datasource config

Fetches the datasource config for the specified custom datasource.

    @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
    @return ApiGetdatasourceconfigPostRequest


```go
func (a *DatasourcesApiService) GetdatasourceconfigPostExecute(r ApiGetdatasourceconfigPostRequest) (*CustomDatasourceConfig, *http.Response, error)
```
Execute executes the request

    @return CustomDatasourceConfig




### Type DeleteDocumentRequest
```go
type DeleteDocumentRequest struct {
	// Version number for document for optimistic concurrency control. If absent or 0 then no version checks are done.
	Version *int64 `json:"version,omitempty"`
	// datasource of the document
	Datasource string `json:"datasource"`
	// object type of the document
	ObjectType string `json:"objectType"`
	// The id of the document
	Id string `json:"id"`
}
```
DeleteDocumentRequest Describes the request body of the /deletedocument API
call

### Functions

```go
func NewDeleteDocumentRequest(datasource string, objectType string, id string) *DeleteDocumentRequest
```
NewDeleteDocumentRequest instantiates a new DeleteDocumentRequest object
This constructor will assign default values to properties that have it
defined, and makes sure properties required by API are set, but the set of
arguments will change when the set of required properties is changed


```go
func NewDeleteDocumentRequestWithDefaults() *DeleteDocumentRequest
```
NewDeleteDocumentRequestWithDefaults instantiates a new
DeleteDocumentRequest object This constructor will only assign default
values to properties that have it defined, but it doesn't guarantee that
properties required by API are set



### Methods

```go
func (o *DeleteDocumentRequest) GetDatasource() string
```
GetDatasource returns the Datasource field value


```go
func (o *DeleteDocumentRequest) GetDatasourceOk() (*string, bool)
```
GetDatasourceOk returns a tuple with the Datasource field value and a
boolean to check if the value has been set.


```go
func (o *DeleteDocumentRequest) GetId() string
```
GetId returns the Id field value


```go
func (o *DeleteDocumentRequest) GetIdOk() (*string, bool)
```
GetIdOk returns a tuple with the Id field value and a boolean to check if
the value has been set.


```go
func (o *DeleteDocumentRequest) GetObjectType() string
```
GetObjectType returns the ObjectType field value


```go
func (o *DeleteDocumentRequest) GetObjectTypeOk() (*string, bool)
```
GetObjectTypeOk returns a tuple with the ObjectType field value and a
boolean to check if the value has been set.


```go
func (o *DeleteDocumentRequest) GetVersion() int64
```
GetVersion returns the Version field value if set, zero value otherwise.


```go
func (o *DeleteDocumentRequest) GetVersionOk() (*int64, bool)
```
GetVersionOk returns a tuple with the Version field value if set, nil
otherwise and a boolean to check if the value has been set.


```go
func (o *DeleteDocumentRequest) HasVersion() bool
```
HasVersion returns a boolean if a field has been set.


```go
func (o DeleteDocumentRequest) MarshalJSON() ([]byte, error)
```


```go
func (o *DeleteDocumentRequest) SetDatasource(v string)
```
SetDatasource sets field value


```go
func (o *DeleteDocumentRequest) SetId(v string)
```
SetId sets field value


```go
func (o *DeleteDocumentRequest) SetObjectType(v string)
```
SetObjectType sets field value


```go
func (o *DeleteDocumentRequest) SetVersion(v int64)
```
SetVersion gets a reference to the given int64 and assigns it to the Version
field.




### Type DeleteEmployeeRequest
```go
type DeleteEmployeeRequest struct {
	// Version number for document for optimistic concurrency control. If absent or 0 then no version checks are done.
	Version *int64 `json:"version,omitempty"`
	// The deleted employee's email
	EmployeeEmail string `json:"employeeEmail"`
}
```
DeleteEmployeeRequest Describes the request body of the /deleteemployee API
call

### Functions

```go
func NewDeleteEmployeeRequest(employeeEmail string) *DeleteEmployeeRequest
```
NewDeleteEmployeeRequest instantiates a new DeleteEmployeeRequest object
This constructor will assign default values to properties that have it
defined, and makes sure properties required by API are set, but the set of
arguments will change when the set of required properties is changed


```go
func NewDeleteEmployeeRequestWithDefaults() *DeleteEmployeeRequest
```
NewDeleteEmployeeRequestWithDefaults instantiates a new
DeleteEmployeeRequest object This constructor will only assign default
values to properties that have it defined, but it doesn't guarantee that
properties required by API are set



### Methods

```go
func (o *DeleteEmployeeRequest) GetEmployeeEmail() string
```
GetEmployeeEmail returns the EmployeeEmail field value


```go
func (o *DeleteEmployeeRequest) GetEmployeeEmailOk() (*string, bool)
```
GetEmployeeEmailOk returns a tuple with the EmployeeEmail field value and a
boolean to check if the value has been set.


```go
func (o *DeleteEmployeeRequest) GetVersion() int64
```
GetVersion returns the Version field value if set, zero value otherwise.


```go
func (o *DeleteEmployeeRequest) GetVersionOk() (*int64, bool)
```
GetVersionOk returns a tuple with the Version field value if set, nil
otherwise and a boolean to check if the value has been set.


```go
func (o *DeleteEmployeeRequest) HasVersion() bool
```
HasVersion returns a boolean if a field has been set.


```go
func (o DeleteEmployeeRequest) MarshalJSON() ([]byte, error)
```


```go
func (o *DeleteEmployeeRequest) SetEmployeeEmail(v string)
```
SetEmployeeEmail sets field value


```go
func (o *DeleteEmployeeRequest) SetVersion(v int64)
```
SetVersion gets a reference to the given int64 and assigns it to the Version
field.




### Type DeleteGroupRequest
```go
type DeleteGroupRequest struct {
	// Version number for document for optimistic concurrency control. If absent or 0 then no version checks are done.
	Version *int64 `json:"version,omitempty"`
	// The datasource for which the group is removed
	Datasource string `json:"datasource"`
	// the name of the group to be deleted
	GroupName string `json:"groupName"`
}
```
DeleteGroupRequest Describes the request body of the /deletegroup API call

### Functions

```go
func NewDeleteGroupRequest(datasource string, groupName string) *DeleteGroupRequest
```
NewDeleteGroupRequest instantiates a new DeleteGroupRequest object This
constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed


```go
func NewDeleteGroupRequestWithDefaults() *DeleteGroupRequest
```
NewDeleteGroupRequestWithDefaults instantiates a new DeleteGroupRequest
object This constructor will only assign default values to properties that
have it defined, but it doesn't guarantee that properties required by API
are set



### Methods

```go
func (o *DeleteGroupRequest) GetDatasource() string
```
GetDatasource returns the Datasource field value


```go
func (o *DeleteGroupRequest) GetDatasourceOk() (*string, bool)
```
GetDatasourceOk returns a tuple with the Datasource field value and a
boolean to check if the value has been set.


```go
func (o *DeleteGroupRequest) GetGroupName() string
```
GetGroupName returns the GroupName field value


```go
func (o *DeleteGroupRequest) GetGroupNameOk() (*string, bool)
```
GetGroupNameOk returns a tuple with the GroupName field value and a boolean
to check if the value has been set.


```go
func (o *DeleteGroupRequest) GetVersion() int64
```
GetVersion returns the Version field value if set, zero value otherwise.


```go
func (o *DeleteGroupRequest) GetVersionOk() (*int64, bool)
```
GetVersionOk returns a tuple with the Version field value if set, nil
otherwise and a boolean to check if the value has been set.


```go
func (o *DeleteGroupRequest) HasVersion() bool
```
HasVersion returns a boolean if a field has been set.


```go
func (o DeleteGroupRequest) MarshalJSON() ([]byte, error)
```


```go
func (o *DeleteGroupRequest) SetDatasource(v string)
```
SetDatasource sets field value


```go
func (o *DeleteGroupRequest) SetGroupName(v string)
```
SetGroupName sets field value


```go
func (o *DeleteGroupRequest) SetVersion(v int64)
```
SetVersion gets a reference to the given int64 and assigns it to the Version
field.




### Type DeleteMembershipRequest
```go
type DeleteMembershipRequest struct {
	// Version number for document for optimistic concurrency control. If absent or 0 then no version checks are done.
	Version *int64 `json:"version,omitempty"`
	// The datasource for which the membership is removed
	Datasource string                         `json:"datasource"`
	Membership DatasourceMembershipDefinition `json:"membership"`
}
```
DeleteMembershipRequest Describes the request body of the /deletemembership
API call

### Functions

```go
func NewDeleteMembershipRequest(datasource string, membership DatasourceMembershipDefinition) *DeleteMembershipRequest
```
NewDeleteMembershipRequest instantiates a new DeleteMembershipRequest object
This constructor will assign default values to properties that have it
defined, and makes sure properties required by API are set, but the set of
arguments will change when the set of required properties is changed


```go
func NewDeleteMembershipRequestWithDefaults() *DeleteMembershipRequest
```
NewDeleteMembershipRequestWithDefaults instantiates a new
DeleteMembershipRequest object This constructor will only assign default
values to properties that have it defined, but it doesn't guarantee that
properties required by API are set



### Methods

```go
func (o *DeleteMembershipRequest) GetDatasource() string
```
GetDatasource returns the Datasource field value


```go
func (o *DeleteMembershipRequest) GetDatasourceOk() (*string, bool)
```
GetDatasourceOk returns a tuple with the Datasource field value and a
boolean to check if the value has been set.


```go
func (o *DeleteMembershipRequest) GetMembership() DatasourceMembershipDefinition
```
GetMembership returns the Membership field value


```go
func (o *DeleteMembershipRequest) GetMembershipOk() (*DatasourceMembershipDefinition, bool)
```
GetMembershipOk returns a tuple with the Membership field value and a
boolean to check if the value has been set.


```go
func (o *DeleteMembershipRequest) GetVersion() int64
```
GetVersion returns the Version field value if set, zero value otherwise.


```go
func (o *DeleteMembershipRequest) GetVersionOk() (*int64, bool)
```
GetVersionOk returns a tuple with the Version field value if set, nil
otherwise and a boolean to check if the value has been set.


```go
func (o *DeleteMembershipRequest) HasVersion() bool
```
HasVersion returns a boolean if a field has been set.


```go
func (o DeleteMembershipRequest) MarshalJSON() ([]byte, error)
```


```go
func (o *DeleteMembershipRequest) SetDatasource(v string)
```
SetDatasource sets field value


```go
func (o *DeleteMembershipRequest) SetMembership(v DatasourceMembershipDefinition)
```
SetMembership sets field value


```go
func (o *DeleteMembershipRequest) SetVersion(v int64)
```
SetVersion gets a reference to the given int64 and assigns it to the Version
field.




### Type DeleteUserRequest
```go
type DeleteUserRequest struct {
	// Version number for document for optimistic concurrency control. If absent or 0 then no version checks are done.
	Version *int64 `json:"version,omitempty"`
	// The datasource for which the user is removed
	Datasource string `json:"datasource"`
	// The email of the user to be deleted
	Email string `json:"email"`
}
```
DeleteUserRequest Describes the request body of the /deleteuser API call

### Functions

```go
func NewDeleteUserRequest(datasource string, email string) *DeleteUserRequest
```
NewDeleteUserRequest instantiates a new DeleteUserRequest object This
constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed


```go
func NewDeleteUserRequestWithDefaults() *DeleteUserRequest
```
NewDeleteUserRequestWithDefaults instantiates a new DeleteUserRequest object
This constructor will only assign default values to properties that have it
defined, but it doesn't guarantee that properties required by API are set



### Methods

```go
func (o *DeleteUserRequest) GetDatasource() string
```
GetDatasource returns the Datasource field value


```go
func (o *DeleteUserRequest) GetDatasourceOk() (*string, bool)
```
GetDatasourceOk returns a tuple with the Datasource field value and a
boolean to check if the value has been set.


```go
func (o *DeleteUserRequest) GetEmail() string
```
GetEmail returns the Email field value


```go
func (o *DeleteUserRequest) GetEmailOk() (*string, bool)
```
GetEmailOk returns a tuple with the Email field value and a boolean to check
if the value has been set.


```go
func (o *DeleteUserRequest) GetVersion() int64
```
GetVersion returns the Version field value if set, zero value otherwise.


```go
func (o *DeleteUserRequest) GetVersionOk() (*int64, bool)
```
GetVersionOk returns a tuple with the Version field value if set, nil
otherwise and a boolean to check if the value has been set.


```go
func (o *DeleteUserRequest) HasVersion() bool
```
HasVersion returns a boolean if a field has been set.


```go
func (o DeleteUserRequest) MarshalJSON() ([]byte, error)
```


```go
func (o *DeleteUserRequest) SetDatasource(v string)
```
SetDatasource sets field value


```go
func (o *DeleteUserRequest) SetEmail(v string)
```
SetEmail sets field value


```go
func (o *DeleteUserRequest) SetVersion(v int64)
```
SetVersion gets a reference to the given int64 and assigns it to the Version
field.




### Type DocumentDefinition
```go
type DocumentDefinition struct {
	// Document title, in plain text, if present.
	Title *string `json:"title,omitempty"`
	// The container for the content (Folder for example for file content).
	Container  *string `json:"container,omitempty"`
	Datasource string  `json:"datasource"`
	// The type of the document (Case, KnowledgeArticle for Salesforce for example). It cannot have spaces or _
	ObjectType *string `json:"objectType,omitempty"`
	// The permalink for viewing the document.
	ViewURL *string `json:"viewURL,omitempty"`
	// The datasource specific id for the document. This should not be more than 200 characters in length.
	Id          *string                        `json:"id,omitempty"`
	Summary     *ContentDefinition             `json:"summary,omitempty"`
	Body        *ContentDefinition             `json:"body,omitempty"`
	Author      *UserReferenceDefinition       `json:"author,omitempty"`
	Owner       *UserReferenceDefinition       `json:"owner,omitempty"`
	Permissions *DocumentPermissionsDefinition `json:"permissions,omitempty"`
	// The creation time, in epoch seconds.
	CreatedAt *int64 `json:"createdAt,omitempty"`
	// The last update time, in epoch seconds.
	UpdatedAt *int64                   `json:"updatedAt,omitempty"`
	UpdatedBy *UserReferenceDefinition `json:"updatedBy,omitempty"`
	// Labels associated with the document.
	Tags         []string                        `json:"tags,omitempty"`
	Interactions *DocumentInteractionsDefinition `json:"interactions,omitempty"`
	Status       *string                         `json:"status,omitempty"`
	// Additional variations of the url that this document points to.
	AdditionalUrls []string `json:"additionalUrls,omitempty"`
	// Additional metadata properties of the document.
	CustomProperties []CustomProperty `json:"customProperties,omitempty"`
}
```
DocumentDefinition Indexable document structure

### Functions

```go
func NewDocumentDefinition(datasource string) *DocumentDefinition
```
NewDocumentDefinition instantiates a new DocumentDefinition object This
constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed


```go
func NewDocumentDefinitionWithDefaults() *DocumentDefinition
```
NewDocumentDefinitionWithDefaults instantiates a new DocumentDefinition
object This constructor will only assign default values to properties that
have it defined, but it doesn't guarantee that properties required by API
are set



### Methods

```go
func (o *DocumentDefinition) GetAdditionalUrls() []string
```
GetAdditionalUrls returns the AdditionalUrls field value if set, zero value
otherwise.


```go
func (o *DocumentDefinition) GetAdditionalUrlsOk() ([]string, bool)
```
GetAdditionalUrlsOk returns a tuple with the AdditionalUrls field value if
set, nil otherwise and a boolean to check if the value has been set.


```go
func (o *DocumentDefinition) GetAuthor() UserReferenceDefinition
```
GetAuthor returns the Author field value if set, zero value otherwise.


```go
func (o *DocumentDefinition) GetAuthorOk() (*UserReferenceDefinition, bool)
```
GetAuthorOk returns a tuple with the Author field value if set, nil
otherwise and a boolean to check if the value has been set.


```go
func (o *DocumentDefinition) GetBody() ContentDefinition
```
GetBody returns the Body field value if set, zero value otherwise.


```go
func (o *DocumentDefinition) GetBodyOk() (*ContentDefinition, bool)
```
GetBodyOk returns a tuple with the Body field value if set, nil otherwise
and a boolean to check if the value has been set.


```go
func (o *DocumentDefinition) GetContainer() string
```
GetContainer returns the Container field value if set, zero value otherwise.


```go
func (o *DocumentDefinition) GetContainerOk() (*string, bool)
```
GetContainerOk returns a tuple with the Container field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *DocumentDefinition) GetCreatedAt() int64
```
GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.


```go
func (o *DocumentDefinition) GetCreatedAtOk() (*int64, bool)
```
GetCreatedAtOk returns a tuple with the CreatedAt field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *DocumentDefinition) GetCustomProperties() []CustomProperty
```
GetCustomProperties returns the CustomProperties field value if set,
zero value otherwise.


```go
func (o *DocumentDefinition) GetCustomPropertiesOk() ([]CustomProperty, bool)
```
GetCustomPropertiesOk returns a tuple with the CustomProperties field value
if set, nil otherwise and a boolean to check if the value has been set.


```go
func (o *DocumentDefinition) GetDatasource() string
```
GetDatasource returns the Datasource field value


```go
func (o *DocumentDefinition) GetDatasourceOk() (*string, bool)
```
GetDatasourceOk returns a tuple with the Datasource field value and a
boolean to check if the value has been set.


```go
func (o *DocumentDefinition) GetId() string
```
GetId returns the Id field value if set, zero value otherwise.


```go
func (o *DocumentDefinition) GetIdOk() (*string, bool)
```
GetIdOk returns a tuple with the Id field value if set, nil otherwise and a
boolean to check if the value has been set.


```go
func (o *DocumentDefinition) GetInteractions() DocumentInteractionsDefinition
```
GetInteractions returns the Interactions field value if set, zero value
otherwise.


```go
func (o *DocumentDefinition) GetInteractionsOk() (*DocumentInteractionsDefinition, bool)
```
GetInteractionsOk returns a tuple with the Interactions field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *DocumentDefinition) GetObjectType() string
```
GetObjectType returns the ObjectType field value if set, zero value
otherwise.


```go
func (o *DocumentDefinition) GetObjectTypeOk() (*string, bool)
```
GetObjectTypeOk returns a tuple with the ObjectType field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *DocumentDefinition) GetOwner() UserReferenceDefinition
```
GetOwner returns the Owner field value if set, zero value otherwise.


```go
func (o *DocumentDefinition) GetOwnerOk() (*UserReferenceDefinition, bool)
```
GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
and a boolean to check if the value has been set.


```go
func (o *DocumentDefinition) GetPermissions() DocumentPermissionsDefinition
```
GetPermissions returns the Permissions field value if set, zero value
otherwise.


```go
func (o *DocumentDefinition) GetPermissionsOk() (*DocumentPermissionsDefinition, bool)
```
GetPermissionsOk returns a tuple with the Permissions field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *DocumentDefinition) GetStatus() string
```
GetStatus returns the Status field value if set, zero value otherwise.


```go
func (o *DocumentDefinition) GetStatusOk() (*string, bool)
```
GetStatusOk returns a tuple with the Status field value if set, nil
otherwise and a boolean to check if the value has been set.


```go
func (o *DocumentDefinition) GetSummary() ContentDefinition
```
GetSummary returns the Summary field value if set, zero value otherwise.


```go
func (o *DocumentDefinition) GetSummaryOk() (*ContentDefinition, bool)
```
GetSummaryOk returns a tuple with the Summary field value if set, nil
otherwise and a boolean to check if the value has been set.


```go
func (o *DocumentDefinition) GetTags() []string
```
GetTags returns the Tags field value if set, zero value otherwise.


```go
func (o *DocumentDefinition) GetTagsOk() ([]string, bool)
```
GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
and a boolean to check if the value has been set.


```go
func (o *DocumentDefinition) GetTitle() string
```
GetTitle returns the Title field value if set, zero value otherwise.


```go
func (o *DocumentDefinition) GetTitleOk() (*string, bool)
```
GetTitleOk returns a tuple with the Title field value if set, nil otherwise
and a boolean to check if the value has been set.


```go
func (o *DocumentDefinition) GetUpdatedAt() int64
```
GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.


```go
func (o *DocumentDefinition) GetUpdatedAtOk() (*int64, bool)
```
GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *DocumentDefinition) GetUpdatedBy() UserReferenceDefinition
```
GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.


```go
func (o *DocumentDefinition) GetUpdatedByOk() (*UserReferenceDefinition, bool)
```
GetUpdatedByOk returns a tuple with the UpdatedBy field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *DocumentDefinition) GetViewURL() string
```
GetViewURL returns the ViewURL field value if set, zero value otherwise.


```go
func (o *DocumentDefinition) GetViewURLOk() (*string, bool)
```
GetViewURLOk returns a tuple with the ViewURL field value if set, nil
otherwise and a boolean to check if the value has been set.


```go
func (o *DocumentDefinition) HasAdditionalUrls() bool
```
HasAdditionalUrls returns a boolean if a field has been set.


```go
func (o *DocumentDefinition) HasAuthor() bool
```
HasAuthor returns a boolean if a field has been set.


```go
func (o *DocumentDefinition) HasBody() bool
```
HasBody returns a boolean if a field has been set.


```go
func (o *DocumentDefinition) HasContainer() bool
```
HasContainer returns a boolean if a field has been set.


```go
func (o *DocumentDefinition) HasCreatedAt() bool
```
HasCreatedAt returns a boolean if a field has been set.


```go
func (o *DocumentDefinition) HasCustomProperties() bool
```
HasCustomProperties returns a boolean if a field has been set.


```go
func (o *DocumentDefinition) HasId() bool
```
HasId returns a boolean if a field has been set.


```go
func (o *DocumentDefinition) HasInteractions() bool
```
HasInteractions returns a boolean if a field has been set.


```go
func (o *DocumentDefinition) HasObjectType() bool
```
HasObjectType returns a boolean if a field has been set.


```go
func (o *DocumentDefinition) HasOwner() bool
```
HasOwner returns a boolean if a field has been set.


```go
func (o *DocumentDefinition) HasPermissions() bool
```
HasPermissions returns a boolean if a field has been set.


```go
func (o *DocumentDefinition) HasStatus() bool
```
HasStatus returns a boolean if a field has been set.


```go
func (o *DocumentDefinition) HasSummary() bool
```
HasSummary returns a boolean if a field has been set.


```go
func (o *DocumentDefinition) HasTags() bool
```
HasTags returns a boolean if a field has been set.


```go
func (o *DocumentDefinition) HasTitle() bool
```
HasTitle returns a boolean if a field has been set.


```go
func (o *DocumentDefinition) HasUpdatedAt() bool
```
HasUpdatedAt returns a boolean if a field has been set.


```go
func (o *DocumentDefinition) HasUpdatedBy() bool
```
HasUpdatedBy returns a boolean if a field has been set.


```go
func (o *DocumentDefinition) HasViewURL() bool
```
HasViewURL returns a boolean if a field has been set.


```go
func (o DocumentDefinition) MarshalJSON() ([]byte, error)
```


```go
func (o *DocumentDefinition) SetAdditionalUrls(v []string)
```
SetAdditionalUrls gets a reference to the given []string and assigns it to
the AdditionalUrls field.


```go
func (o *DocumentDefinition) SetAuthor(v UserReferenceDefinition)
```
SetAuthor gets a reference to the given UserReferenceDefinition and assigns
it to the Author field.


```go
func (o *DocumentDefinition) SetBody(v ContentDefinition)
```
SetBody gets a reference to the given ContentDefinition and assigns it to
the Body field.


```go
func (o *DocumentDefinition) SetContainer(v string)
```
SetContainer gets a reference to the given string and assigns it to the
Container field.


```go
func (o *DocumentDefinition) SetCreatedAt(v int64)
```
SetCreatedAt gets a reference to the given int64 and assigns it to the
CreatedAt field.


```go
func (o *DocumentDefinition) SetCustomProperties(v []CustomProperty)
```
SetCustomProperties gets a reference to the given []CustomProperty and
assigns it to the CustomProperties field.


```go
func (o *DocumentDefinition) SetDatasource(v string)
```
SetDatasource sets field value


```go
func (o *DocumentDefinition) SetId(v string)
```
SetId gets a reference to the given string and assigns it to the Id field.


```go
func (o *DocumentDefinition) SetInteractions(v DocumentInteractionsDefinition)
```
SetInteractions gets a reference to the given DocumentInteractionsDefinition
and assigns it to the Interactions field.


```go
func (o *DocumentDefinition) SetObjectType(v string)
```
SetObjectType gets a reference to the given string and assigns it to the
ObjectType field.


```go
func (o *DocumentDefinition) SetOwner(v UserReferenceDefinition)
```
SetOwner gets a reference to the given UserReferenceDefinition and assigns
it to the Owner field.


```go
func (o *DocumentDefinition) SetPermissions(v DocumentPermissionsDefinition)
```
SetPermissions gets a reference to the given DocumentPermissionsDefinition
and assigns it to the Permissions field.


```go
func (o *DocumentDefinition) SetStatus(v string)
```
SetStatus gets a reference to the given string and assigns it to the Status
field.


```go
func (o *DocumentDefinition) SetSummary(v ContentDefinition)
```
SetSummary gets a reference to the given ContentDefinition and assigns it to
the Summary field.


```go
func (o *DocumentDefinition) SetTags(v []string)
```
SetTags gets a reference to the given []string and assigns it to the Tags
field.


```go
func (o *DocumentDefinition) SetTitle(v string)
```
SetTitle gets a reference to the given string and assigns it to the Title
field.


```go
func (o *DocumentDefinition) SetUpdatedAt(v int64)
```
SetUpdatedAt gets a reference to the given int64 and assigns it to the
UpdatedAt field.


```go
func (o *DocumentDefinition) SetUpdatedBy(v UserReferenceDefinition)
```
SetUpdatedBy gets a reference to the given UserReferenceDefinition and
assigns it to the UpdatedBy field.


```go
func (o *DocumentDefinition) SetViewURL(v string)
```
SetViewURL gets a reference to the given string and assigns it to the
ViewURL field.




### Type DocumentInteractionsDefinition
```go
type DocumentInteractionsDefinition struct {
	NumViews    *int32 `json:"numViews,omitempty"`
	NumLikes    *int32 `json:"numLikes,omitempty"`
	NumComments *int32 `json:"numComments,omitempty"`
}
```
DocumentInteractionsDefinition describes the interactions on the document

### Functions

```go
func NewDocumentInteractionsDefinition() *DocumentInteractionsDefinition
```
NewDocumentInteractionsDefinition instantiates a new
DocumentInteractionsDefinition object This constructor will assign default
values to properties that have it defined, and makes sure properties
required by API are set, but the set of arguments will change when the set
of required properties is changed


```go
func NewDocumentInteractionsDefinitionWithDefaults() *DocumentInteractionsDefinition
```
NewDocumentInteractionsDefinitionWithDefaults instantiates a new
DocumentInteractionsDefinition object This constructor will only assign
default values to properties that have it defined, but it doesn't guarantee
that properties required by API are set



### Methods

```go
func (o *DocumentInteractionsDefinition) GetNumComments() int32
```
GetNumComments returns the NumComments field value if set, zero value
otherwise.


```go
func (o *DocumentInteractionsDefinition) GetNumCommentsOk() (*int32, bool)
```
GetNumCommentsOk returns a tuple with the NumComments field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *DocumentInteractionsDefinition) GetNumLikes() int32
```
GetNumLikes returns the NumLikes field value if set, zero value otherwise.


```go
func (o *DocumentInteractionsDefinition) GetNumLikesOk() (*int32, bool)
```
GetNumLikesOk returns a tuple with the NumLikes field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *DocumentInteractionsDefinition) GetNumViews() int32
```
GetNumViews returns the NumViews field value if set, zero value otherwise.


```go
func (o *DocumentInteractionsDefinition) GetNumViewsOk() (*int32, bool)
```
GetNumViewsOk returns a tuple with the NumViews field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *DocumentInteractionsDefinition) HasNumComments() bool
```
HasNumComments returns a boolean if a field has been set.


```go
func (o *DocumentInteractionsDefinition) HasNumLikes() bool
```
HasNumLikes returns a boolean if a field has been set.


```go
func (o *DocumentInteractionsDefinition) HasNumViews() bool
```
HasNumViews returns a boolean if a field has been set.


```go
func (o DocumentInteractionsDefinition) MarshalJSON() ([]byte, error)
```


```go
func (o *DocumentInteractionsDefinition) SetNumComments(v int32)
```
SetNumComments gets a reference to the given int32 and assigns it to the
NumComments field.


```go
func (o *DocumentInteractionsDefinition) SetNumLikes(v int32)
```
SetNumLikes gets a reference to the given int32 and assigns it to the
NumLikes field.


```go
func (o *DocumentInteractionsDefinition) SetNumViews(v int32)
```
SetNumViews gets a reference to the given int32 and assigns it to the
NumViews field.




### Type DocumentPermissionsDefinition
```go
type DocumentPermissionsDefinition struct {
	// List of users who can view the document
	AllowedUsers []UserReferenceDefinition `json:"allowedUsers,omitempty"`
	// List of groups that can view the document
	AllowedGroups []string `json:"allowedGroups,omitempty"`
	// List of allowed group intersections. This describes a permissions constraint of the form ((GroupA AND GroupB AND GroupC) OR (GroupX AND GroupY) OR ...
	AllowedGroupIntersections []PermissionsGroupIntersectionDefinition `json:"allowedGroupIntersections,omitempty"`
	// If true, then any Glean user can view the document
	AllowAnonymousAccess *bool `json:"allowAnonymousAccess,omitempty"`
	// If true, then any user who has an account in the datasource can view the document.
	AllowAllDatasourceUsersAccess *bool `json:"allowAllDatasourceUsersAccess,omitempty"`
}
```
DocumentPermissionsDefinition describes the access control details of the
document

### Functions

```go
func NewDocumentPermissionsDefinition() *DocumentPermissionsDefinition
```
NewDocumentPermissionsDefinition instantiates a new
DocumentPermissionsDefinition object This constructor will assign default
values to properties that have it defined, and makes sure properties
required by API are set, but the set of arguments will change when the set
of required properties is changed


```go
func NewDocumentPermissionsDefinitionWithDefaults() *DocumentPermissionsDefinition
```
NewDocumentPermissionsDefinitionWithDefaults instantiates a new
DocumentPermissionsDefinition object This constructor will only assign
default values to properties that have it defined, but it doesn't guarantee
that properties required by API are set



### Methods

```go
func (o *DocumentPermissionsDefinition) GetAllowAllDatasourceUsersAccess() bool
```
GetAllowAllDatasourceUsersAccess returns the AllowAllDatasourceUsersAccess
field value if set, zero value otherwise.


```go
func (o *DocumentPermissionsDefinition) GetAllowAllDatasourceUsersAccessOk() (*bool, bool)
```
GetAllowAllDatasourceUsersAccessOk returns a tuple with the
AllowAllDatasourceUsersAccess field value if set, nil otherwise and a
boolean to check if the value has been set.


```go
func (o *DocumentPermissionsDefinition) GetAllowAnonymousAccess() bool
```
GetAllowAnonymousAccess returns the AllowAnonymousAccess field value if set,
zero value otherwise.


```go
func (o *DocumentPermissionsDefinition) GetAllowAnonymousAccessOk() (*bool, bool)
```
GetAllowAnonymousAccessOk returns a tuple with the AllowAnonymousAccess
field value if set, nil otherwise and a boolean to check if the value has
been set.


```go
func (o *DocumentPermissionsDefinition) GetAllowedGroupIntersections() []PermissionsGroupIntersectionDefinition
```
GetAllowedGroupIntersections returns the AllowedGroupIntersections field
value if set, zero value otherwise.


```go
func (o *DocumentPermissionsDefinition) GetAllowedGroupIntersectionsOk() ([]PermissionsGroupIntersectionDefinition, bool)
```
GetAllowedGroupIntersectionsOk returns a tuple with the
AllowedGroupIntersections field value if set, nil otherwise and a boolean to
check if the value has been set.


```go
func (o *DocumentPermissionsDefinition) GetAllowedGroups() []string
```
GetAllowedGroups returns the AllowedGroups field value if set, zero value
otherwise.


```go
func (o *DocumentPermissionsDefinition) GetAllowedGroupsOk() ([]string, bool)
```
GetAllowedGroupsOk returns a tuple with the AllowedGroups field value if
set, nil otherwise and a boolean to check if the value has been set.


```go
func (o *DocumentPermissionsDefinition) GetAllowedUsers() []UserReferenceDefinition
```
GetAllowedUsers returns the AllowedUsers field value if set, zero value
otherwise.


```go
func (o *DocumentPermissionsDefinition) GetAllowedUsersOk() ([]UserReferenceDefinition, bool)
```
GetAllowedUsersOk returns a tuple with the AllowedUsers field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *DocumentPermissionsDefinition) HasAllowAllDatasourceUsersAccess() bool
```
HasAllowAllDatasourceUsersAccess returns a boolean if a field has been set.


```go
func (o *DocumentPermissionsDefinition) HasAllowAnonymousAccess() bool
```
HasAllowAnonymousAccess returns a boolean if a field has been set.


```go
func (o *DocumentPermissionsDefinition) HasAllowedGroupIntersections() bool
```
HasAllowedGroupIntersections returns a boolean if a field has been set.


```go
func (o *DocumentPermissionsDefinition) HasAllowedGroups() bool
```
HasAllowedGroups returns a boolean if a field has been set.


```go
func (o *DocumentPermissionsDefinition) HasAllowedUsers() bool
```
HasAllowedUsers returns a boolean if a field has been set.


```go
func (o DocumentPermissionsDefinition) MarshalJSON() ([]byte, error)
```


```go
func (o *DocumentPermissionsDefinition) SetAllowAllDatasourceUsersAccess(v bool)
```
SetAllowAllDatasourceUsersAccess gets a reference to the given bool and
assigns it to the AllowAllDatasourceUsersAccess field.


```go
func (o *DocumentPermissionsDefinition) SetAllowAnonymousAccess(v bool)
```
SetAllowAnonymousAccess gets a reference to the given bool and assigns it to
the AllowAnonymousAccess field.


```go
func (o *DocumentPermissionsDefinition) SetAllowedGroupIntersections(v []PermissionsGroupIntersectionDefinition)
```
SetAllowedGroupIntersections gets a reference to the given
[]PermissionsGroupIntersectionDefinition and assigns it to the
AllowedGroupIntersections field.


```go
func (o *DocumentPermissionsDefinition) SetAllowedGroups(v []string)
```
SetAllowedGroups gets a reference to the given []string and assigns it to
the AllowedGroups field.


```go
func (o *DocumentPermissionsDefinition) SetAllowedUsers(v []UserReferenceDefinition)
```
SetAllowedUsers gets a reference to the given []UserReferenceDefinition and
assigns it to the AllowedUsers field.




### Type DocumentsApiService
```go
type DocumentsApiService service
```
DocumentsApiService DocumentsApi service

### Methods

```go
func (a *DocumentsApiService) BulkindexdocumentsPost(ctx context.Context) ApiBulkindexdocumentsPostRequest
```
BulkindexdocumentsPost Bulk index documents

Replaces the documents in a datasource using paginated batch API calls.

    @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
    @return ApiBulkindexdocumentsPostRequest


```go
func (a *DocumentsApiService) BulkindexdocumentsPostExecute(r ApiBulkindexdocumentsPostRequest) (*http.Response, error)
```
Execute executes the request


```go
func (a *DocumentsApiService) DeletedocumentPost(ctx context.Context) ApiDeletedocumentPostRequest
```
DeletedocumentPost Delete document

Deletes the specified document from the index. Succeeds if document is not
present.

    @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
    @return ApiDeletedocumentPostRequest


```go
func (a *DocumentsApiService) DeletedocumentPostExecute(r ApiDeletedocumentPostRequest) (*http.Response, error)
```
Execute executes the request


```go
func (a *DocumentsApiService) GetdocumentcountPost(ctx context.Context) ApiGetdocumentcountPostRequest
```
GetdocumentcountPost Get document count

Fetches document count for the specified custom datasource.

    @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
    @return ApiGetdocumentcountPostRequest


```go
func (a *DocumentsApiService) GetdocumentcountPostExecute(r ApiGetdocumentcountPostRequest) (*GetDocumentCountResponse, *http.Response, error)
```
Execute executes the request

    @return GetDocumentCountResponse


```go
func (a *DocumentsApiService) GetdocumentstatusPost(ctx context.Context) ApiGetdocumentstatusPostRequest
```
GetdocumentstatusPost Get document upload and indexing status

Intended for debugging/validation. Fetches the current upload and indexing
status of documents.

    @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
    @return ApiGetdocumentstatusPostRequest


```go
func (a *DocumentsApiService) GetdocumentstatusPostExecute(r ApiGetdocumentstatusPostRequest) (*GetDocumentStatusResponse, *http.Response, error)
```
Execute executes the request

    @return GetDocumentStatusResponse


```go
func (a *DocumentsApiService) IndexdocumentPost(ctx context.Context) ApiIndexdocumentPostRequest
```
IndexdocumentPost Index document

Adds a document to the index or updates an existing document.

    @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
    @return ApiIndexdocumentPostRequest


```go
func (a *DocumentsApiService) IndexdocumentPostExecute(r ApiIndexdocumentPostRequest) (*http.Response, error)
```
Execute executes the request


```go
func (a *DocumentsApiService) ProcessalldocumentsPost(ctx context.Context) ApiProcessalldocumentsPostRequest
```
ProcessalldocumentsPost Schedules the processing of uploaded documents

Schedules the immediate processing of documents uploaded through the
indexing API. By default the uploaded documents will be processed
asynchronously but this API can be used to schedule processing of all
documents on demand.

    @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
    @return ApiProcessalldocumentsPostRequest


```go
func (a *DocumentsApiService) ProcessalldocumentsPostExecute(r ApiProcessalldocumentsPostRequest) (*http.Response, error)
```
Execute executes the request




### Type EmployeeAndVersionDefinition
```go
type EmployeeAndVersionDefinition struct {
	Employee *EmployeeInfoDefinition `json:"employee,omitempty"`
	// Version number for the employee object. If absent or 0 then no version checks are done
	Version *int64 `json:"version,omitempty"`
}
```
EmployeeAndVersionDefinition describes info about an employee and optional
version for that info

### Functions

```go
func NewEmployeeAndVersionDefinition() *EmployeeAndVersionDefinition
```
NewEmployeeAndVersionDefinition instantiates a new
EmployeeAndVersionDefinition object This constructor will assign default
values to properties that have it defined, and makes sure properties
required by API are set, but the set of arguments will change when the set
of required properties is changed


```go
func NewEmployeeAndVersionDefinitionWithDefaults() *EmployeeAndVersionDefinition
```
NewEmployeeAndVersionDefinitionWithDefaults instantiates a new
EmployeeAndVersionDefinition object This constructor will only assign
default values to properties that have it defined, but it doesn't guarantee
that properties required by API are set



### Methods

```go
func (o *EmployeeAndVersionDefinition) GetEmployee() EmployeeInfoDefinition
```
GetEmployee returns the Employee field value if set, zero value otherwise.


```go
func (o *EmployeeAndVersionDefinition) GetEmployeeOk() (*EmployeeInfoDefinition, bool)
```
GetEmployeeOk returns a tuple with the Employee field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *EmployeeAndVersionDefinition) GetVersion() int64
```
GetVersion returns the Version field value if set, zero value otherwise.


```go
func (o *EmployeeAndVersionDefinition) GetVersionOk() (*int64, bool)
```
GetVersionOk returns a tuple with the Version field value if set, nil
otherwise and a boolean to check if the value has been set.


```go
func (o *EmployeeAndVersionDefinition) HasEmployee() bool
```
HasEmployee returns a boolean if a field has been set.


```go
func (o *EmployeeAndVersionDefinition) HasVersion() bool
```
HasVersion returns a boolean if a field has been set.


```go
func (o EmployeeAndVersionDefinition) MarshalJSON() ([]byte, error)
```


```go
func (o *EmployeeAndVersionDefinition) SetEmployee(v EmployeeInfoDefinition)
```
SetEmployee gets a reference to the given EmployeeInfoDefinition and assigns
it to the Employee field.


```go
func (o *EmployeeAndVersionDefinition) SetVersion(v int64)
```
SetVersion gets a reference to the given int64 and assigns it to the Version
field.




### Type EmployeeInfoDefinition
```go
type EmployeeInfoDefinition struct {
	// The employee's email
	Email string `json:"email"`
	// The first name of the employee
	FirstName *string `json:"firstName,omitempty"`
	// The last name of the employee
	LastName *string `json:"lastName,omitempty"`
	// The preferred name or nickname of the employee
	PreferredName *string `json:"preferredName,omitempty"`
	// **[Advanced]** A unique universal internal identifier for the employee. This is solely used for understanding manager relationships along with `managerId`.
	Id *string `json:"id,omitempty"`
	// The employee's phone number.
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	// The employee's location (city/office name etc).
	// Deprecated
	Location           *string             `json:"location,omitempty"`
	StructuredLocation *StructuredLocation `json:"structuredLocation,omitempty"`
	// The employee's role title.
	Title *string `json:"title,omitempty"`
	// The employee's profile pic
	PhotoUrl *string `json:"photoUrl,omitempty"`
	// Typically the highest level organizational unit; generally applies to bigger companies with multiple distinct businesses.
	BusinessUnit *string `json:"businessUnit,omitempty"`
	// An organizational unit where everyone has a similar task, e.g. `Engineering`.
	Department string `json:"department"`
	// The datasource profiles of the employee, e.g. `Slack`,`Github`.
	DatasourceProfiles []DatasourceProfile `json:"datasourceProfiles,omitempty"`
	// Info about the employee's team(s)
	Teams []EmployeeTeamInfo `json:"teams,omitempty"`
	// The date when the employee started
	StartDate *string `json:"startDate,omitempty"`
	// If a former employee, the last date of employment.
	EndDate *string `json:"endDate,omitempty"`
	// Short biography or mission statement of the employee.
	Bio *string `json:"bio,omitempty"`
	// She/her, He/his or other pronoun.
	Pronoun *string `json:"pronoun,omitempty"`
	// Other names associated with the employee.
	AlsoKnownAs []string `json:"alsoKnownAs,omitempty"`
	// Link to internal company person profile.
	ProfileUrl *string `json:"profileUrl,omitempty"`
	// List of social network profiles.
	SocialNetworks []SocialNetworkDefinition `json:"socialNetworks,omitempty"`
	// The email of the employee's manager
	ManagerEmail *string `json:"managerEmail,omitempty"`
	// **[Advanced]** A unique universal internal identifier for the employee's manager. This is solely used in conjunction with `id`.
	ManagerId *string `json:"managerId,omitempty"`
	// The status of the employee, an enum of `CURRENT`, `FUTURE`, `EX`
	Status *string `json:"status,omitempty"`
	// List of additional fields with more information about the employee.
	AdditionalFields []AdditionalFieldDefinition `json:"additionalFields,omitempty"`
}
```
EmployeeInfoDefinition Describes employee info

### Functions

```go
func NewEmployeeInfoDefinition(email string, department string) *EmployeeInfoDefinition
```
NewEmployeeInfoDefinition instantiates a new EmployeeInfoDefinition object
This constructor will assign default values to properties that have it
defined, and makes sure properties required by API are set, but the set of
arguments will change when the set of required properties is changed


```go
func NewEmployeeInfoDefinitionWithDefaults() *EmployeeInfoDefinition
```
NewEmployeeInfoDefinitionWithDefaults instantiates a new
EmployeeInfoDefinition object This constructor will only assign default
values to properties that have it defined, but it doesn't guarantee that
properties required by API are set



### Methods

```go
func (o *EmployeeInfoDefinition) GetAdditionalFields() []AdditionalFieldDefinition
```
GetAdditionalFields returns the AdditionalFields field value if set,
zero value otherwise.


```go
func (o *EmployeeInfoDefinition) GetAdditionalFieldsOk() ([]AdditionalFieldDefinition, bool)
```
GetAdditionalFieldsOk returns a tuple with the AdditionalFields field value
if set, nil otherwise and a boolean to check if the value has been set.


```go
func (o *EmployeeInfoDefinition) GetAlsoKnownAs() []string
```
GetAlsoKnownAs returns the AlsoKnownAs field value if set, zero value
otherwise.


```go
func (o *EmployeeInfoDefinition) GetAlsoKnownAsOk() ([]string, bool)
```
GetAlsoKnownAsOk returns a tuple with the AlsoKnownAs field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *EmployeeInfoDefinition) GetBio() string
```
GetBio returns the Bio field value if set, zero value otherwise.


```go
func (o *EmployeeInfoDefinition) GetBioOk() (*string, bool)
```
GetBioOk returns a tuple with the Bio field value if set, nil otherwise and
a boolean to check if the value has been set.


```go
func (o *EmployeeInfoDefinition) GetBusinessUnit() string
```
GetBusinessUnit returns the BusinessUnit field value if set, zero value
otherwise.


```go
func (o *EmployeeInfoDefinition) GetBusinessUnitOk() (*string, bool)
```
GetBusinessUnitOk returns a tuple with the BusinessUnit field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *EmployeeInfoDefinition) GetDatasourceProfiles() []DatasourceProfile
```
GetDatasourceProfiles returns the DatasourceProfiles field value if set,
zero value otherwise.


```go
func (o *EmployeeInfoDefinition) GetDatasourceProfilesOk() ([]DatasourceProfile, bool)
```
GetDatasourceProfilesOk returns a tuple with the DatasourceProfiles field
value if set, nil otherwise and a boolean to check if the value has been
set.


```go
func (o *EmployeeInfoDefinition) GetDepartment() string
```
GetDepartment returns the Department field value


```go
func (o *EmployeeInfoDefinition) GetDepartmentOk() (*string, bool)
```
GetDepartmentOk returns a tuple with the Department field value and a
boolean to check if the value has been set.


```go
func (o *EmployeeInfoDefinition) GetEmail() string
```
GetEmail returns the Email field value


```go
func (o *EmployeeInfoDefinition) GetEmailOk() (*string, bool)
```
GetEmailOk returns a tuple with the Email field value and a boolean to check
if the value has been set.


```go
func (o *EmployeeInfoDefinition) GetEndDate() string
```
GetEndDate returns the EndDate field value if set, zero value otherwise.


```go
func (o *EmployeeInfoDefinition) GetEndDateOk() (*string, bool)
```
GetEndDateOk returns a tuple with the EndDate field value if set, nil
otherwise and a boolean to check if the value has been set.


```go
func (o *EmployeeInfoDefinition) GetFirstName() string
```
GetFirstName returns the FirstName field value if set, zero value otherwise.


```go
func (o *EmployeeInfoDefinition) GetFirstNameOk() (*string, bool)
```
GetFirstNameOk returns a tuple with the FirstName field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *EmployeeInfoDefinition) GetId() string
```
GetId returns the Id field value if set, zero value otherwise.


```go
func (o *EmployeeInfoDefinition) GetIdOk() (*string, bool)
```
GetIdOk returns a tuple with the Id field value if set, nil otherwise and a
boolean to check if the value has been set.


```go
func (o *EmployeeInfoDefinition) GetLastName() string
```
GetLastName returns the LastName field value if set, zero value otherwise.


```go
func (o *EmployeeInfoDefinition) GetLastNameOk() (*string, bool)
```
GetLastNameOk returns a tuple with the LastName field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *EmployeeInfoDefinition) GetLocation() string
```
GetLocation returns the Location field value if set, zero value otherwise.
Deprecated


```go
func (o *EmployeeInfoDefinition) GetLocationOk() (*string, bool)
```
GetLocationOk returns a tuple with the Location field value if set,
nil otherwise and a boolean to check if the value has been set. Deprecated


```go
func (o *EmployeeInfoDefinition) GetManagerEmail() string
```
GetManagerEmail returns the ManagerEmail field value if set, zero value
otherwise.


```go
func (o *EmployeeInfoDefinition) GetManagerEmailOk() (*string, bool)
```
GetManagerEmailOk returns a tuple with the ManagerEmail field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *EmployeeInfoDefinition) GetManagerId() string
```
GetManagerId returns the ManagerId field value if set, zero value otherwise.


```go
func (o *EmployeeInfoDefinition) GetManagerIdOk() (*string, bool)
```
GetManagerIdOk returns a tuple with the ManagerId field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *EmployeeInfoDefinition) GetPhoneNumber() string
```
GetPhoneNumber returns the PhoneNumber field value if set, zero value
otherwise.


```go
func (o *EmployeeInfoDefinition) GetPhoneNumberOk() (*string, bool)
```
GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *EmployeeInfoDefinition) GetPhotoUrl() string
```
GetPhotoUrl returns the PhotoUrl field value if set, zero value otherwise.


```go
func (o *EmployeeInfoDefinition) GetPhotoUrlOk() (*string, bool)
```
GetPhotoUrlOk returns a tuple with the PhotoUrl field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *EmployeeInfoDefinition) GetPreferredName() string
```
GetPreferredName returns the PreferredName field value if set, zero value
otherwise.


```go
func (o *EmployeeInfoDefinition) GetPreferredNameOk() (*string, bool)
```
GetPreferredNameOk returns a tuple with the PreferredName field value if
set, nil otherwise and a boolean to check if the value has been set.


```go
func (o *EmployeeInfoDefinition) GetProfileUrl() string
```
GetProfileUrl returns the ProfileUrl field value if set, zero value
otherwise.


```go
func (o *EmployeeInfoDefinition) GetProfileUrlOk() (*string, bool)
```
GetProfileUrlOk returns a tuple with the ProfileUrl field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *EmployeeInfoDefinition) GetPronoun() string
```
GetPronoun returns the Pronoun field value if set, zero value otherwise.


```go
func (o *EmployeeInfoDefinition) GetPronounOk() (*string, bool)
```
GetPronounOk returns a tuple with the Pronoun field value if set, nil
otherwise and a boolean to check if the value has been set.


```go
func (o *EmployeeInfoDefinition) GetSocialNetworks() []SocialNetworkDefinition
```
GetSocialNetworks returns the SocialNetworks field value if set, zero value
otherwise.


```go
func (o *EmployeeInfoDefinition) GetSocialNetworksOk() ([]SocialNetworkDefinition, bool)
```
GetSocialNetworksOk returns a tuple with the SocialNetworks field value if
set, nil otherwise and a boolean to check if the value has been set.


```go
func (o *EmployeeInfoDefinition) GetStartDate() string
```
GetStartDate returns the StartDate field value if set, zero value otherwise.


```go
func (o *EmployeeInfoDefinition) GetStartDateOk() (*string, bool)
```
GetStartDateOk returns a tuple with the StartDate field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *EmployeeInfoDefinition) GetStatus() string
```
GetStatus returns the Status field value if set, zero value otherwise.


```go
func (o *EmployeeInfoDefinition) GetStatusOk() (*string, bool)
```
GetStatusOk returns a tuple with the Status field value if set, nil
otherwise and a boolean to check if the value has been set.


```go
func (o *EmployeeInfoDefinition) GetStructuredLocation() StructuredLocation
```
GetStructuredLocation returns the StructuredLocation field value if set,
zero value otherwise.


```go
func (o *EmployeeInfoDefinition) GetStructuredLocationOk() (*StructuredLocation, bool)
```
GetStructuredLocationOk returns a tuple with the StructuredLocation field
value if set, nil otherwise and a boolean to check if the value has been
set.


```go
func (o *EmployeeInfoDefinition) GetTeams() []EmployeeTeamInfo
```
GetTeams returns the Teams field value if set, zero value otherwise.


```go
func (o *EmployeeInfoDefinition) GetTeamsOk() ([]EmployeeTeamInfo, bool)
```
GetTeamsOk returns a tuple with the Teams field value if set, nil otherwise
and a boolean to check if the value has been set.


```go
func (o *EmployeeInfoDefinition) GetTitle() string
```
GetTitle returns the Title field value if set, zero value otherwise.


```go
func (o *EmployeeInfoDefinition) GetTitleOk() (*string, bool)
```
GetTitleOk returns a tuple with the Title field value if set, nil otherwise
and a boolean to check if the value has been set.


```go
func (o *EmployeeInfoDefinition) HasAdditionalFields() bool
```
HasAdditionalFields returns a boolean if a field has been set.


```go
func (o *EmployeeInfoDefinition) HasAlsoKnownAs() bool
```
HasAlsoKnownAs returns a boolean if a field has been set.


```go
func (o *EmployeeInfoDefinition) HasBio() bool
```
HasBio returns a boolean if a field has been set.


```go
func (o *EmployeeInfoDefinition) HasBusinessUnit() bool
```
HasBusinessUnit returns a boolean if a field has been set.


```go
func (o *EmployeeInfoDefinition) HasDatasourceProfiles() bool
```
HasDatasourceProfiles returns a boolean if a field has been set.


```go
func (o *EmployeeInfoDefinition) HasEndDate() bool
```
HasEndDate returns a boolean if a field has been set.


```go
func (o *EmployeeInfoDefinition) HasFirstName() bool
```
HasFirstName returns a boolean if a field has been set.


```go
func (o *EmployeeInfoDefinition) HasId() bool
```
HasId returns a boolean if a field has been set.


```go
func (o *EmployeeInfoDefinition) HasLastName() bool
```
HasLastName returns a boolean if a field has been set.


```go
func (o *EmployeeInfoDefinition) HasLocation() bool
```
HasLocation returns a boolean if a field has been set.


```go
func (o *EmployeeInfoDefinition) HasManagerEmail() bool
```
HasManagerEmail returns a boolean if a field has been set.


```go
func (o *EmployeeInfoDefinition) HasManagerId() bool
```
HasManagerId returns a boolean if a field has been set.


```go
func (o *EmployeeInfoDefinition) HasPhoneNumber() bool
```
HasPhoneNumber returns a boolean if a field has been set.


```go
func (o *EmployeeInfoDefinition) HasPhotoUrl() bool
```
HasPhotoUrl returns a boolean if a field has been set.


```go
func (o *EmployeeInfoDefinition) HasPreferredName() bool
```
HasPreferredName returns a boolean if a field has been set.


```go
func (o *EmployeeInfoDefinition) HasProfileUrl() bool
```
HasProfileUrl returns a boolean if a field has been set.


```go
func (o *EmployeeInfoDefinition) HasPronoun() bool
```
HasPronoun returns a boolean if a field has been set.


```go
func (o *EmployeeInfoDefinition) HasSocialNetworks() bool
```
HasSocialNetworks returns a boolean if a field has been set.


```go
func (o *EmployeeInfoDefinition) HasStartDate() bool
```
HasStartDate returns a boolean if a field has been set.


```go
func (o *EmployeeInfoDefinition) HasStatus() bool
```
HasStatus returns a boolean if a field has been set.


```go
func (o *EmployeeInfoDefinition) HasStructuredLocation() bool
```
HasStructuredLocation returns a boolean if a field has been set.


```go
func (o *EmployeeInfoDefinition) HasTeams() bool
```
HasTeams returns a boolean if a field has been set.


```go
func (o *EmployeeInfoDefinition) HasTitle() bool
```
HasTitle returns a boolean if a field has been set.


```go
func (o EmployeeInfoDefinition) MarshalJSON() ([]byte, error)
```


```go
func (o *EmployeeInfoDefinition) SetAdditionalFields(v []AdditionalFieldDefinition)
```
SetAdditionalFields gets a reference to the given
[]AdditionalFieldDefinition and assigns it to the AdditionalFields field.


```go
func (o *EmployeeInfoDefinition) SetAlsoKnownAs(v []string)
```
SetAlsoKnownAs gets a reference to the given []string and assigns it to the
AlsoKnownAs field.


```go
func (o *EmployeeInfoDefinition) SetBio(v string)
```
SetBio gets a reference to the given string and assigns it to the Bio field.


```go
func (o *EmployeeInfoDefinition) SetBusinessUnit(v string)
```
SetBusinessUnit gets a reference to the given string and assigns it to the
BusinessUnit field.


```go
func (o *EmployeeInfoDefinition) SetDatasourceProfiles(v []DatasourceProfile)
```
SetDatasourceProfiles gets a reference to the given []DatasourceProfile and
assigns it to the DatasourceProfiles field.


```go
func (o *EmployeeInfoDefinition) SetDepartment(v string)
```
SetDepartment sets field value


```go
func (o *EmployeeInfoDefinition) SetEmail(v string)
```
SetEmail sets field value


```go
func (o *EmployeeInfoDefinition) SetEndDate(v string)
```
SetEndDate gets a reference to the given string and assigns it to the
EndDate field.


```go
func (o *EmployeeInfoDefinition) SetFirstName(v string)
```
SetFirstName gets a reference to the given string and assigns it to the
FirstName field.


```go
func (o *EmployeeInfoDefinition) SetId(v string)
```
SetId gets a reference to the given string and assigns it to the Id field.


```go
func (o *EmployeeInfoDefinition) SetLastName(v string)
```
SetLastName gets a reference to the given string and assigns it to the
LastName field.


```go
func (o *EmployeeInfoDefinition) SetLocation(v string)
```
SetLocation gets a reference to the given string and assigns it to the
Location field. Deprecated


```go
func (o *EmployeeInfoDefinition) SetManagerEmail(v string)
```
SetManagerEmail gets a reference to the given string and assigns it to the
ManagerEmail field.


```go
func (o *EmployeeInfoDefinition) SetManagerId(v string)
```
SetManagerId gets a reference to the given string and assigns it to the
ManagerId field.


```go
func (o *EmployeeInfoDefinition) SetPhoneNumber(v string)
```
SetPhoneNumber gets a reference to the given string and assigns it to the
PhoneNumber field.


```go
func (o *EmployeeInfoDefinition) SetPhotoUrl(v string)
```
SetPhotoUrl gets a reference to the given string and assigns it to the
PhotoUrl field.


```go
func (o *EmployeeInfoDefinition) SetPreferredName(v string)
```
SetPreferredName gets a reference to the given string and assigns it to the
PreferredName field.


```go
func (o *EmployeeInfoDefinition) SetProfileUrl(v string)
```
SetProfileUrl gets a reference to the given string and assigns it to the
ProfileUrl field.


```go
func (o *EmployeeInfoDefinition) SetPronoun(v string)
```
SetPronoun gets a reference to the given string and assigns it to the
Pronoun field.


```go
func (o *EmployeeInfoDefinition) SetSocialNetworks(v []SocialNetworkDefinition)
```
SetSocialNetworks gets a reference to the given []SocialNetworkDefinition
and assigns it to the SocialNetworks field.


```go
func (o *EmployeeInfoDefinition) SetStartDate(v string)
```
SetStartDate gets a reference to the given string and assigns it to the
StartDate field.


```go
func (o *EmployeeInfoDefinition) SetStatus(v string)
```
SetStatus gets a reference to the given string and assigns it to the Status
field.


```go
func (o *EmployeeInfoDefinition) SetStructuredLocation(v StructuredLocation)
```
SetStructuredLocation gets a reference to the given StructuredLocation and
assigns it to the StructuredLocation field.


```go
func (o *EmployeeInfoDefinition) SetTeams(v []EmployeeTeamInfo)
```
SetTeams gets a reference to the given []EmployeeTeamInfo and assigns it to
the Teams field.


```go
func (o *EmployeeInfoDefinition) SetTitle(v string)
```
SetTitle gets a reference to the given string and assigns it to the Title
field.




### Type EmployeeTeamInfo
```go
type EmployeeTeamInfo struct {
	// unique identifier for this team
	Id *string `json:"id,omitempty"`
	// Team name
	Name *string `json:"name,omitempty"`
	// Link to internal company team page
	Url *string `json:"url,omitempty"`
}
```
EmployeeTeamInfo Information about which team an employee belongs to

### Functions

```go
func NewEmployeeTeamInfo() *EmployeeTeamInfo
```
NewEmployeeTeamInfo instantiates a new EmployeeTeamInfo object This
constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed


```go
func NewEmployeeTeamInfoWithDefaults() *EmployeeTeamInfo
```
NewEmployeeTeamInfoWithDefaults instantiates a new EmployeeTeamInfo object
This constructor will only assign default values to properties that have it
defined, but it doesn't guarantee that properties required by API are set



### Methods

```go
func (o *EmployeeTeamInfo) GetId() string
```
GetId returns the Id field value if set, zero value otherwise.


```go
func (o *EmployeeTeamInfo) GetIdOk() (*string, bool)
```
GetIdOk returns a tuple with the Id field value if set, nil otherwise and a
boolean to check if the value has been set.


```go
func (o *EmployeeTeamInfo) GetName() string
```
GetName returns the Name field value if set, zero value otherwise.


```go
func (o *EmployeeTeamInfo) GetNameOk() (*string, bool)
```
GetNameOk returns a tuple with the Name field value if set, nil otherwise
and a boolean to check if the value has been set.


```go
func (o *EmployeeTeamInfo) GetUrl() string
```
GetUrl returns the Url field value if set, zero value otherwise.


```go
func (o *EmployeeTeamInfo) GetUrlOk() (*string, bool)
```
GetUrlOk returns a tuple with the Url field value if set, nil otherwise and
a boolean to check if the value has been set.


```go
func (o *EmployeeTeamInfo) HasId() bool
```
HasId returns a boolean if a field has been set.


```go
func (o *EmployeeTeamInfo) HasName() bool
```
HasName returns a boolean if a field has been set.


```go
func (o *EmployeeTeamInfo) HasUrl() bool
```
HasUrl returns a boolean if a field has been set.


```go
func (o EmployeeTeamInfo) MarshalJSON() ([]byte, error)
```


```go
func (o *EmployeeTeamInfo) SetId(v string)
```
SetId gets a reference to the given string and assigns it to the Id field.


```go
func (o *EmployeeTeamInfo) SetName(v string)
```
SetName gets a reference to the given string and assigns it to the Name
field.


```go
func (o *EmployeeTeamInfo) SetUrl(v string)
```
SetUrl gets a reference to the given string and assigns it to the Url field.




### Type GenericOpenAPIError
```go
type GenericOpenAPIError struct {
	// contains filtered or unexported fields
}
```
GenericOpenAPIError Provides access to the body, error and model on returned
errors.

### Methods

```go
func (e GenericOpenAPIError) Body() []byte
```
Body returns the raw bytes of the response


```go
func (e GenericOpenAPIError) Error() string
```
Error returns non-empty string if there was an error.


```go
func (e GenericOpenAPIError) Model() interface{}
```
Model returns the unpacked model of the error




### Type GetDatasourceConfigRequest
```go
type GetDatasourceConfigRequest struct {
	// Datasource name for which config is needed.
	Name *string `json:"name,omitempty"`
}
```
GetDatasourceConfigRequest Describes the request body of the
/getdatasourceconfig API call

### Functions

```go
func NewGetDatasourceConfigRequest() *GetDatasourceConfigRequest
```
NewGetDatasourceConfigRequest instantiates a new GetDatasourceConfigRequest
object This constructor will assign default values to properties that have
it defined, and makes sure properties required by API are set, but the set
of arguments will change when the set of required properties is changed


```go
func NewGetDatasourceConfigRequestWithDefaults() *GetDatasourceConfigRequest
```
NewGetDatasourceConfigRequestWithDefaults instantiates a new
GetDatasourceConfigRequest object This constructor will only assign default
values to properties that have it defined, but it doesn't guarantee that
properties required by API are set



### Methods

```go
func (o *GetDatasourceConfigRequest) GetName() string
```
GetName returns the Name field value if set, zero value otherwise.


```go
func (o *GetDatasourceConfigRequest) GetNameOk() (*string, bool)
```
GetNameOk returns a tuple with the Name field value if set, nil otherwise
and a boolean to check if the value has been set.


```go
func (o *GetDatasourceConfigRequest) HasName() bool
```
HasName returns a boolean if a field has been set.


```go
func (o GetDatasourceConfigRequest) MarshalJSON() ([]byte, error)
```


```go
func (o *GetDatasourceConfigRequest) SetName(v string)
```
SetName gets a reference to the given string and assigns it to the Name
field.




### Type GetDocumentCountRequest
```go
type GetDocumentCountRequest struct {
	// Datasource name for which document count is needed.
	Name string `json:"name"`
}
```
GetDocumentCountRequest Describes the request body of the /getdocumentcount
API call

### Functions

```go
func NewGetDocumentCountRequest(name string) *GetDocumentCountRequest
```
NewGetDocumentCountRequest instantiates a new GetDocumentCountRequest object
This constructor will assign default values to properties that have it
defined, and makes sure properties required by API are set, but the set of
arguments will change when the set of required properties is changed


```go
func NewGetDocumentCountRequestWithDefaults() *GetDocumentCountRequest
```
NewGetDocumentCountRequestWithDefaults instantiates a new
GetDocumentCountRequest object This constructor will only assign default
values to properties that have it defined, but it doesn't guarantee that
properties required by API are set



### Methods

```go
func (o *GetDocumentCountRequest) GetName() string
```
GetName returns the Name field value


```go
func (o *GetDocumentCountRequest) GetNameOk() (*string, bool)
```
GetNameOk returns a tuple with the Name field value and a boolean to check
if the value has been set.


```go
func (o GetDocumentCountRequest) MarshalJSON() ([]byte, error)
```


```go
func (o *GetDocumentCountRequest) SetName(v string)
```
SetName sets field value




### Type GetDocumentCountResponse
```go
type GetDocumentCountResponse struct {
	// Number of documents corresponding to the specified custom datasource.
	DocumentCount *int32 `json:"documentCount,omitempty"`
}
```
GetDocumentCountResponse Describes the response body of the
/getdocumentcount API call

### Functions

```go
func NewGetDocumentCountResponse() *GetDocumentCountResponse
```
NewGetDocumentCountResponse instantiates a new GetDocumentCountResponse
object This constructor will assign default values to properties that have
it defined, and makes sure properties required by API are set, but the set
of arguments will change when the set of required properties is changed


```go
func NewGetDocumentCountResponseWithDefaults() *GetDocumentCountResponse
```
NewGetDocumentCountResponseWithDefaults instantiates a new
GetDocumentCountResponse object This constructor will only assign default
values to properties that have it defined, but it doesn't guarantee that
properties required by API are set



### Methods

```go
func (o *GetDocumentCountResponse) GetDocumentCount() int32
```
GetDocumentCount returns the DocumentCount field value if set, zero value
otherwise.


```go
func (o *GetDocumentCountResponse) GetDocumentCountOk() (*int32, bool)
```
GetDocumentCountOk returns a tuple with the DocumentCount field value if
set, nil otherwise and a boolean to check if the value has been set.


```go
func (o *GetDocumentCountResponse) HasDocumentCount() bool
```
HasDocumentCount returns a boolean if a field has been set.


```go
func (o GetDocumentCountResponse) MarshalJSON() ([]byte, error)
```


```go
func (o *GetDocumentCountResponse) SetDocumentCount(v int32)
```
SetDocumentCount gets a reference to the given int32 and assigns it to the
DocumentCount field.




### Type GetDocumentStatusRequest
```go
type GetDocumentStatusRequest struct {
	// Datasource to get fetch document status for
	Datasource string `json:"datasource"`
	// Object type of the document to get the status for
	ObjectType string `json:"objectType"`
	// Document ID within the datasource to get the status for
	DocId string `json:"docId"`
}
```
GetDocumentStatusRequest Describes the request body for /getdocumentstatus
API call

### Functions

```go
func NewGetDocumentStatusRequest(datasource string, objectType string, docId string) *GetDocumentStatusRequest
```
NewGetDocumentStatusRequest instantiates a new GetDocumentStatusRequest
object This constructor will assign default values to properties that have
it defined, and makes sure properties required by API are set, but the set
of arguments will change when the set of required properties is changed


```go
func NewGetDocumentStatusRequestWithDefaults() *GetDocumentStatusRequest
```
NewGetDocumentStatusRequestWithDefaults instantiates a new
GetDocumentStatusRequest object This constructor will only assign default
values to properties that have it defined, but it doesn't guarantee that
properties required by API are set



### Methods

```go
func (o *GetDocumentStatusRequest) GetDatasource() string
```
GetDatasource returns the Datasource field value


```go
func (o *GetDocumentStatusRequest) GetDatasourceOk() (*string, bool)
```
GetDatasourceOk returns a tuple with the Datasource field value and a
boolean to check if the value has been set.


```go
func (o *GetDocumentStatusRequest) GetDocId() string
```
GetDocId returns the DocId field value


```go
func (o *GetDocumentStatusRequest) GetDocIdOk() (*string, bool)
```
GetDocIdOk returns a tuple with the DocId field value and a boolean to check
if the value has been set.


```go
func (o *GetDocumentStatusRequest) GetObjectType() string
```
GetObjectType returns the ObjectType field value


```go
func (o *GetDocumentStatusRequest) GetObjectTypeOk() (*string, bool)
```
GetObjectTypeOk returns a tuple with the ObjectType field value and a
boolean to check if the value has been set.


```go
func (o GetDocumentStatusRequest) MarshalJSON() ([]byte, error)
```


```go
func (o *GetDocumentStatusRequest) SetDatasource(v string)
```
SetDatasource sets field value


```go
func (o *GetDocumentStatusRequest) SetDocId(v string)
```
SetDocId sets field value


```go
func (o *GetDocumentStatusRequest) SetObjectType(v string)
```
SetObjectType sets field value




### Type GetDocumentStatusResponse
```go
type GetDocumentStatusResponse struct {
	// Upload status, enum of NOT_UPLOADED, UPLOADED, STATUS_UNKNOWN
	UploadStatus *string `json:"uploadStatus,omitempty"`
	// Time of last successful upload, in epoch seconds
	LastUploadedAt *int64 `json:"lastUploadedAt,omitempty"`
	// Indexing status, enum of NOT_INDEXED, INDEXED, STATUS_UNKNOWN
	IndexingStatus *string `json:"indexingStatus,omitempty"`
	// Time of last successful indexing, in epoch seconds
	LastIndexedAt *int64 `json:"lastIndexedAt,omitempty"`
}
```
GetDocumentStatusResponse Describes the response body of the
/getdocumentstatus API call

### Functions

```go
func NewGetDocumentStatusResponse() *GetDocumentStatusResponse
```
NewGetDocumentStatusResponse instantiates a new GetDocumentStatusResponse
object This constructor will assign default values to properties that have
it defined, and makes sure properties required by API are set, but the set
of arguments will change when the set of required properties is changed


```go
func NewGetDocumentStatusResponseWithDefaults() *GetDocumentStatusResponse
```
NewGetDocumentStatusResponseWithDefaults instantiates a new
GetDocumentStatusResponse object This constructor will only assign default
values to properties that have it defined, but it doesn't guarantee that
properties required by API are set



### Methods

```go
func (o *GetDocumentStatusResponse) GetIndexingStatus() string
```
GetIndexingStatus returns the IndexingStatus field value if set, zero value
otherwise.


```go
func (o *GetDocumentStatusResponse) GetIndexingStatusOk() (*string, bool)
```
GetIndexingStatusOk returns a tuple with the IndexingStatus field value if
set, nil otherwise and a boolean to check if the value has been set.


```go
func (o *GetDocumentStatusResponse) GetLastIndexedAt() int64
```
GetLastIndexedAt returns the LastIndexedAt field value if set, zero value
otherwise.


```go
func (o *GetDocumentStatusResponse) GetLastIndexedAtOk() (*int64, bool)
```
GetLastIndexedAtOk returns a tuple with the LastIndexedAt field value if
set, nil otherwise and a boolean to check if the value has been set.


```go
func (o *GetDocumentStatusResponse) GetLastUploadedAt() int64
```
GetLastUploadedAt returns the LastUploadedAt field value if set, zero value
otherwise.


```go
func (o *GetDocumentStatusResponse) GetLastUploadedAtOk() (*int64, bool)
```
GetLastUploadedAtOk returns a tuple with the LastUploadedAt field value if
set, nil otherwise and a boolean to check if the value has been set.


```go
func (o *GetDocumentStatusResponse) GetUploadStatus() string
```
GetUploadStatus returns the UploadStatus field value if set, zero value
otherwise.


```go
func (o *GetDocumentStatusResponse) GetUploadStatusOk() (*string, bool)
```
GetUploadStatusOk returns a tuple with the UploadStatus field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *GetDocumentStatusResponse) HasIndexingStatus() bool
```
HasIndexingStatus returns a boolean if a field has been set.


```go
func (o *GetDocumentStatusResponse) HasLastIndexedAt() bool
```
HasLastIndexedAt returns a boolean if a field has been set.


```go
func (o *GetDocumentStatusResponse) HasLastUploadedAt() bool
```
HasLastUploadedAt returns a boolean if a field has been set.


```go
func (o *GetDocumentStatusResponse) HasUploadStatus() bool
```
HasUploadStatus returns a boolean if a field has been set.


```go
func (o GetDocumentStatusResponse) MarshalJSON() ([]byte, error)
```


```go
func (o *GetDocumentStatusResponse) SetIndexingStatus(v string)
```
SetIndexingStatus gets a reference to the given string and assigns it to the
IndexingStatus field.


```go
func (o *GetDocumentStatusResponse) SetLastIndexedAt(v int64)
```
SetLastIndexedAt gets a reference to the given int64 and assigns it to the
LastIndexedAt field.


```go
func (o *GetDocumentStatusResponse) SetLastUploadedAt(v int64)
```
SetLastUploadedAt gets a reference to the given int64 and assigns it to the
LastUploadedAt field.


```go
func (o *GetDocumentStatusResponse) SetUploadStatus(v string)
```
SetUploadStatus gets a reference to the given string and assigns it to the
UploadStatus field.




### Type GetUserCountRequest
```go
type GetUserCountRequest struct {
	// Datasource name for which user count is needed.
	Name *string `json:"name,omitempty"`
}
```
GetUserCountRequest Describes the request body of the /getusercount API call

### Functions

```go
func NewGetUserCountRequest() *GetUserCountRequest
```
NewGetUserCountRequest instantiates a new GetUserCountRequest object This
constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed


```go
func NewGetUserCountRequestWithDefaults() *GetUserCountRequest
```
NewGetUserCountRequestWithDefaults instantiates a new GetUserCountRequest
object This constructor will only assign default values to properties that
have it defined, but it doesn't guarantee that properties required by API
are set



### Methods

```go
func (o *GetUserCountRequest) GetName() string
```
GetName returns the Name field value if set, zero value otherwise.


```go
func (o *GetUserCountRequest) GetNameOk() (*string, bool)
```
GetNameOk returns a tuple with the Name field value if set, nil otherwise
and a boolean to check if the value has been set.


```go
func (o *GetUserCountRequest) HasName() bool
```
HasName returns a boolean if a field has been set.


```go
func (o GetUserCountRequest) MarshalJSON() ([]byte, error)
```


```go
func (o *GetUserCountRequest) SetName(v string)
```
SetName gets a reference to the given string and assigns it to the Name
field.




### Type GetUserCountResponse
```go
type GetUserCountResponse struct {
	// Number of users corresponding to the specified custom datasource.
	UserCount *int32 `json:"userCount,omitempty"`
}
```
GetUserCountResponse Describes the response body of the /getusercount API
call

### Functions

```go
func NewGetUserCountResponse() *GetUserCountResponse
```
NewGetUserCountResponse instantiates a new GetUserCountResponse object This
constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed


```go
func NewGetUserCountResponseWithDefaults() *GetUserCountResponse
```
NewGetUserCountResponseWithDefaults instantiates a new GetUserCountResponse
object This constructor will only assign default values to properties that
have it defined, but it doesn't guarantee that properties required by API
are set



### Methods

```go
func (o *GetUserCountResponse) GetUserCount() int32
```
GetUserCount returns the UserCount field value if set, zero value otherwise.


```go
func (o *GetUserCountResponse) GetUserCountOk() (*int32, bool)
```
GetUserCountOk returns a tuple with the UserCount field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *GetUserCountResponse) HasUserCount() bool
```
HasUserCount returns a boolean if a field has been set.


```go
func (o GetUserCountResponse) MarshalJSON() ([]byte, error)
```


```go
func (o *GetUserCountResponse) SetUserCount(v int32)
```
SetUserCount gets a reference to the given int32 and assigns it to the
UserCount field.




### Type GreenlistUsersRequest
```go
type GreenlistUsersRequest struct {
	// Datasource which needs to be made visible to users specified in the `emails` field.
	Datasource string `json:"datasource"`
	// The emails of the beta users
	Emails []string `json:"emails"`
}
```
GreenlistUsersRequest Describes the request body of the /betausers API call

### Functions

```go
func NewGreenlistUsersRequest(datasource string, emails []string) *GreenlistUsersRequest
```
NewGreenlistUsersRequest instantiates a new GreenlistUsersRequest object
This constructor will assign default values to properties that have it
defined, and makes sure properties required by API are set, but the set of
arguments will change when the set of required properties is changed


```go
func NewGreenlistUsersRequestWithDefaults() *GreenlistUsersRequest
```
NewGreenlistUsersRequestWithDefaults instantiates a new
GreenlistUsersRequest object This constructor will only assign default
values to properties that have it defined, but it doesn't guarantee that
properties required by API are set



### Methods

```go
func (o *GreenlistUsersRequest) GetDatasource() string
```
GetDatasource returns the Datasource field value


```go
func (o *GreenlistUsersRequest) GetDatasourceOk() (*string, bool)
```
GetDatasourceOk returns a tuple with the Datasource field value and a
boolean to check if the value has been set.


```go
func (o *GreenlistUsersRequest) GetEmails() []string
```
GetEmails returns the Emails field value


```go
func (o *GreenlistUsersRequest) GetEmailsOk() ([]string, bool)
```
GetEmailsOk returns a tuple with the Emails field value and a boolean to
check if the value has been set.


```go
func (o GreenlistUsersRequest) MarshalJSON() ([]byte, error)
```


```go
func (o *GreenlistUsersRequest) SetDatasource(v string)
```
SetDatasource sets field value


```go
func (o *GreenlistUsersRequest) SetEmails(v []string)
```
SetEmails sets field value




### Type HypertextField
```go
type HypertextField struct {
	// Anchor text for the hypertext field.
	Anchor *string `json:"anchor,omitempty"`
	// URL for the hypertext field.
	Hyperlink *string `json:"hyperlink,omitempty"`
}
```
HypertextField struct for HypertextField

### Functions

```go
func NewHypertextField() *HypertextField
```
NewHypertextField instantiates a new HypertextField object This constructor
will assign default values to properties that have it defined, and makes
sure properties required by API are set, but the set of arguments will
change when the set of required properties is changed


```go
func NewHypertextFieldWithDefaults() *HypertextField
```
NewHypertextFieldWithDefaults instantiates a new HypertextField object
This constructor will only assign default values to properties that have it
defined, but it doesn't guarantee that properties required by API are set



### Methods

```go
func (o *HypertextField) GetAnchor() string
```
GetAnchor returns the Anchor field value if set, zero value otherwise.


```go
func (o *HypertextField) GetAnchorOk() (*string, bool)
```
GetAnchorOk returns a tuple with the Anchor field value if set, nil
otherwise and a boolean to check if the value has been set.


```go
func (o *HypertextField) GetHyperlink() string
```
GetHyperlink returns the Hyperlink field value if set, zero value otherwise.


```go
func (o *HypertextField) GetHyperlinkOk() (*string, bool)
```
GetHyperlinkOk returns a tuple with the Hyperlink field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *HypertextField) HasAnchor() bool
```
HasAnchor returns a boolean if a field has been set.


```go
func (o *HypertextField) HasHyperlink() bool
```
HasHyperlink returns a boolean if a field has been set.


```go
func (o HypertextField) MarshalJSON() ([]byte, error)
```


```go
func (o *HypertextField) SetAnchor(v string)
```
SetAnchor gets a reference to the given string and assigns it to the Anchor
field.


```go
func (o *HypertextField) SetHyperlink(v string)
```
SetHyperlink gets a reference to the given string and assigns it to the
Hyperlink field.




### Type IconConfig
```go
type IconConfig struct {
	Color    *string `json:"color,omitempty"`
	Key      *string `json:"key,omitempty"`
	IconType *string `json:"iconType,omitempty"`
	// The filename for iconType.GLYPH icons
	Name *string `json:"name,omitempty"`
	// The URL to an image to be displayed for iconType.URL icons
	Url *string `json:"url,omitempty"`
}
```
IconConfig Defines how to render an icon

### Functions

```go
func NewIconConfig() *IconConfig
```
NewIconConfig instantiates a new IconConfig object This constructor will
assign default values to properties that have it defined, and makes sure
properties required by API are set, but the set of arguments will change
when the set of required properties is changed


```go
func NewIconConfigWithDefaults() *IconConfig
```
NewIconConfigWithDefaults instantiates a new IconConfig object This
constructor will only assign default values to properties that have it
defined, but it doesn't guarantee that properties required by API are set



### Methods

```go
func (o *IconConfig) GetColor() string
```
GetColor returns the Color field value if set, zero value otherwise.


```go
func (o *IconConfig) GetColorOk() (*string, bool)
```
GetColorOk returns a tuple with the Color field value if set, nil otherwise
and a boolean to check if the value has been set.


```go
func (o *IconConfig) GetIconType() string
```
GetIconType returns the IconType field value if set, zero value otherwise.


```go
func (o *IconConfig) GetIconTypeOk() (*string, bool)
```
GetIconTypeOk returns a tuple with the IconType field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *IconConfig) GetKey() string
```
GetKey returns the Key field value if set, zero value otherwise.


```go
func (o *IconConfig) GetKeyOk() (*string, bool)
```
GetKeyOk returns a tuple with the Key field value if set, nil otherwise and
a boolean to check if the value has been set.


```go
func (o *IconConfig) GetName() string
```
GetName returns the Name field value if set, zero value otherwise.


```go
func (o *IconConfig) GetNameOk() (*string, bool)
```
GetNameOk returns a tuple with the Name field value if set, nil otherwise
and a boolean to check if the value has been set.


```go
func (o *IconConfig) GetUrl() string
```
GetUrl returns the Url field value if set, zero value otherwise.


```go
func (o *IconConfig) GetUrlOk() (*string, bool)
```
GetUrlOk returns a tuple with the Url field value if set, nil otherwise and
a boolean to check if the value has been set.


```go
func (o *IconConfig) HasColor() bool
```
HasColor returns a boolean if a field has been set.


```go
func (o *IconConfig) HasIconType() bool
```
HasIconType returns a boolean if a field has been set.


```go
func (o *IconConfig) HasKey() bool
```
HasKey returns a boolean if a field has been set.


```go
func (o *IconConfig) HasName() bool
```
HasName returns a boolean if a field has been set.


```go
func (o *IconConfig) HasUrl() bool
```
HasUrl returns a boolean if a field has been set.


```go
func (o IconConfig) MarshalJSON() ([]byte, error)
```


```go
func (o *IconConfig) SetColor(v string)
```
SetColor gets a reference to the given string and assigns it to the Color
field.


```go
func (o *IconConfig) SetIconType(v string)
```
SetIconType gets a reference to the given string and assigns it to the
IconType field.


```go
func (o *IconConfig) SetKey(v string)
```
SetKey gets a reference to the given string and assigns it to the Key field.


```go
func (o *IconConfig) SetName(v string)
```
SetName gets a reference to the given string and assigns it to the Name
field.


```go
func (o *IconConfig) SetUrl(v string)
```
SetUrl gets a reference to the given string and assigns it to the Url field.




### Type IndexDocumentRequest
```go
type IndexDocumentRequest struct {
	// Version number for document for optimistic concurrency control. If absent or 0 then no version checks are done.
	Version  *int64             `json:"version,omitempty"`
	Document DocumentDefinition `json:"document"`
}
```
IndexDocumentRequest Describes the request body of the /indexdocument API
call

### Functions

```go
func NewIndexDocumentRequest(document DocumentDefinition) *IndexDocumentRequest
```
NewIndexDocumentRequest instantiates a new IndexDocumentRequest object This
constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed


```go
func NewIndexDocumentRequestWithDefaults() *IndexDocumentRequest
```
NewIndexDocumentRequestWithDefaults instantiates a new IndexDocumentRequest
object This constructor will only assign default values to properties that
have it defined, but it doesn't guarantee that properties required by API
are set



### Methods

```go
func (o *IndexDocumentRequest) GetDocument() DocumentDefinition
```
GetDocument returns the Document field value


```go
func (o *IndexDocumentRequest) GetDocumentOk() (*DocumentDefinition, bool)
```
GetDocumentOk returns a tuple with the Document field value and a boolean to
check if the value has been set.


```go
func (o *IndexDocumentRequest) GetVersion() int64
```
GetVersion returns the Version field value if set, zero value otherwise.


```go
func (o *IndexDocumentRequest) GetVersionOk() (*int64, bool)
```
GetVersionOk returns a tuple with the Version field value if set, nil
otherwise and a boolean to check if the value has been set.


```go
func (o *IndexDocumentRequest) HasVersion() bool
```
HasVersion returns a boolean if a field has been set.


```go
func (o IndexDocumentRequest) MarshalJSON() ([]byte, error)
```


```go
func (o *IndexDocumentRequest) SetDocument(v DocumentDefinition)
```
SetDocument sets field value


```go
func (o *IndexDocumentRequest) SetVersion(v int64)
```
SetVersion gets a reference to the given int64 and assigns it to the Version
field.




### Type IndexEmployeeListRequest
```go
type IndexEmployeeListRequest struct {
	// List of employee info and version.
	Employees []IndexEmployeeRequest `json:"employees,omitempty"`
}
```
IndexEmployeeListRequest Describes the request body of the
/indexemployeelist API call

### Functions

```go
func NewIndexEmployeeListRequest() *IndexEmployeeListRequest
```
NewIndexEmployeeListRequest instantiates a new IndexEmployeeListRequest
object This constructor will assign default values to properties that have
it defined, and makes sure properties required by API are set, but the set
of arguments will change when the set of required properties is changed


```go
func NewIndexEmployeeListRequestWithDefaults() *IndexEmployeeListRequest
```
NewIndexEmployeeListRequestWithDefaults instantiates a new
IndexEmployeeListRequest object This constructor will only assign default
values to properties that have it defined, but it doesn't guarantee that
properties required by API are set



### Methods

```go
func (o *IndexEmployeeListRequest) GetEmployees() []IndexEmployeeRequest
```
GetEmployees returns the Employees field value if set, zero value otherwise.


```go
func (o *IndexEmployeeListRequest) GetEmployeesOk() ([]IndexEmployeeRequest, bool)
```
GetEmployeesOk returns a tuple with the Employees field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *IndexEmployeeListRequest) HasEmployees() bool
```
HasEmployees returns a boolean if a field has been set.


```go
func (o IndexEmployeeListRequest) MarshalJSON() ([]byte, error)
```


```go
func (o *IndexEmployeeListRequest) SetEmployees(v []IndexEmployeeRequest)
```
SetEmployees gets a reference to the given []IndexEmployeeRequest and
assigns it to the Employees field.




### Type IndexEmployeeRequest
```go
type IndexEmployeeRequest struct {
	Employee EmployeeInfoDefinition `json:"employee"`
	// Version number for the employee object. If absent or 0 then no version checks are done
	Version *int64 `json:"version,omitempty"`
}
```
IndexEmployeeRequest Info about an employee and optional version for that
info

### Functions

```go
func NewIndexEmployeeRequest(employee EmployeeInfoDefinition) *IndexEmployeeRequest
```
NewIndexEmployeeRequest instantiates a new IndexEmployeeRequest object This
constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed


```go
func NewIndexEmployeeRequestWithDefaults() *IndexEmployeeRequest
```
NewIndexEmployeeRequestWithDefaults instantiates a new IndexEmployeeRequest
object This constructor will only assign default values to properties that
have it defined, but it doesn't guarantee that properties required by API
are set



### Methods

```go
func (o *IndexEmployeeRequest) GetEmployee() EmployeeInfoDefinition
```
GetEmployee returns the Employee field value


```go
func (o *IndexEmployeeRequest) GetEmployeeOk() (*EmployeeInfoDefinition, bool)
```
GetEmployeeOk returns a tuple with the Employee field value and a boolean to
check if the value has been set.


```go
func (o *IndexEmployeeRequest) GetVersion() int64
```
GetVersion returns the Version field value if set, zero value otherwise.


```go
func (o *IndexEmployeeRequest) GetVersionOk() (*int64, bool)
```
GetVersionOk returns a tuple with the Version field value if set, nil
otherwise and a boolean to check if the value has been set.


```go
func (o *IndexEmployeeRequest) HasVersion() bool
```
HasVersion returns a boolean if a field has been set.


```go
func (o IndexEmployeeRequest) MarshalJSON() ([]byte, error)
```


```go
func (o *IndexEmployeeRequest) SetEmployee(v EmployeeInfoDefinition)
```
SetEmployee sets field value


```go
func (o *IndexEmployeeRequest) SetVersion(v int64)
```
SetVersion gets a reference to the given int64 and assigns it to the Version
field.




### Type IndexGroupRequest
```go
type IndexGroupRequest struct {
	// Version number for document for optimistic concurrency control. If absent or 0 then no version checks are done.
	Version *int64 `json:"version,omitempty"`
	// The datasource for which the group is added
	Datasource string                    `json:"datasource"`
	Group      DatasourceGroupDefinition `json:"group"`
}
```
IndexGroupRequest Describes the request body of the /indexgroup API call

### Functions

```go
func NewIndexGroupRequest(datasource string, group DatasourceGroupDefinition) *IndexGroupRequest
```
NewIndexGroupRequest instantiates a new IndexGroupRequest object This
constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed


```go
func NewIndexGroupRequestWithDefaults() *IndexGroupRequest
```
NewIndexGroupRequestWithDefaults instantiates a new IndexGroupRequest object
This constructor will only assign default values to properties that have it
defined, but it doesn't guarantee that properties required by API are set



### Methods

```go
func (o *IndexGroupRequest) GetDatasource() string
```
GetDatasource returns the Datasource field value


```go
func (o *IndexGroupRequest) GetDatasourceOk() (*string, bool)
```
GetDatasourceOk returns a tuple with the Datasource field value and a
boolean to check if the value has been set.


```go
func (o *IndexGroupRequest) GetGroup() DatasourceGroupDefinition
```
GetGroup returns the Group field value


```go
func (o *IndexGroupRequest) GetGroupOk() (*DatasourceGroupDefinition, bool)
```
GetGroupOk returns a tuple with the Group field value and a boolean to check
if the value has been set.


```go
func (o *IndexGroupRequest) GetVersion() int64
```
GetVersion returns the Version field value if set, zero value otherwise.


```go
func (o *IndexGroupRequest) GetVersionOk() (*int64, bool)
```
GetVersionOk returns a tuple with the Version field value if set, nil
otherwise and a boolean to check if the value has been set.


```go
func (o *IndexGroupRequest) HasVersion() bool
```
HasVersion returns a boolean if a field has been set.


```go
func (o IndexGroupRequest) MarshalJSON() ([]byte, error)
```


```go
func (o *IndexGroupRequest) SetDatasource(v string)
```
SetDatasource sets field value


```go
func (o *IndexGroupRequest) SetGroup(v DatasourceGroupDefinition)
```
SetGroup sets field value


```go
func (o *IndexGroupRequest) SetVersion(v int64)
```
SetVersion gets a reference to the given int64 and assigns it to the Version
field.




### Type IndexMembershipRequest
```go
type IndexMembershipRequest struct {
	// Version number for document for optimistic concurrency control. If absent or 0 then no version checks are done.
	Version *int64 `json:"version,omitempty"`
	// The datasource for which the membership is added
	Datasource string                         `json:"datasource"`
	Membership DatasourceMembershipDefinition `json:"membership"`
}
```
IndexMembershipRequest Describes the request body of the /indexmembership
API call

### Functions

```go
func NewIndexMembershipRequest(datasource string, membership DatasourceMembershipDefinition) *IndexMembershipRequest
```
NewIndexMembershipRequest instantiates a new IndexMembershipRequest object
This constructor will assign default values to properties that have it
defined, and makes sure properties required by API are set, but the set of
arguments will change when the set of required properties is changed


```go
func NewIndexMembershipRequestWithDefaults() *IndexMembershipRequest
```
NewIndexMembershipRequestWithDefaults instantiates a new
IndexMembershipRequest object This constructor will only assign default
values to properties that have it defined, but it doesn't guarantee that
properties required by API are set



### Methods

```go
func (o *IndexMembershipRequest) GetDatasource() string
```
GetDatasource returns the Datasource field value


```go
func (o *IndexMembershipRequest) GetDatasourceOk() (*string, bool)
```
GetDatasourceOk returns a tuple with the Datasource field value and a
boolean to check if the value has been set.


```go
func (o *IndexMembershipRequest) GetMembership() DatasourceMembershipDefinition
```
GetMembership returns the Membership field value


```go
func (o *IndexMembershipRequest) GetMembershipOk() (*DatasourceMembershipDefinition, bool)
```
GetMembershipOk returns a tuple with the Membership field value and a
boolean to check if the value has been set.


```go
func (o *IndexMembershipRequest) GetVersion() int64
```
GetVersion returns the Version field value if set, zero value otherwise.


```go
func (o *IndexMembershipRequest) GetVersionOk() (*int64, bool)
```
GetVersionOk returns a tuple with the Version field value if set, nil
otherwise and a boolean to check if the value has been set.


```go
func (o *IndexMembershipRequest) HasVersion() bool
```
HasVersion returns a boolean if a field has been set.


```go
func (o IndexMembershipRequest) MarshalJSON() ([]byte, error)
```


```go
func (o *IndexMembershipRequest) SetDatasource(v string)
```
SetDatasource sets field value


```go
func (o *IndexMembershipRequest) SetMembership(v DatasourceMembershipDefinition)
```
SetMembership sets field value


```go
func (o *IndexMembershipRequest) SetVersion(v int64)
```
SetVersion gets a reference to the given int64 and assigns it to the Version
field.




### Type IndexTeamRequest
```go
type IndexTeamRequest struct {
	Team TeamInfoDefinition `json:"team"`
	// Version number for the team object. If absent or 0 then no version checks are done
	Version *int64 `json:"version,omitempty"`
}
```
IndexTeamRequest Info about a team and optional version for that info

### Functions

```go
func NewIndexTeamRequest(team TeamInfoDefinition) *IndexTeamRequest
```
NewIndexTeamRequest instantiates a new IndexTeamRequest object This
constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed


```go
func NewIndexTeamRequestWithDefaults() *IndexTeamRequest
```
NewIndexTeamRequestWithDefaults instantiates a new IndexTeamRequest object
This constructor will only assign default values to properties that have it
defined, but it doesn't guarantee that properties required by API are set



### Methods

```go
func (o *IndexTeamRequest) GetTeam() TeamInfoDefinition
```
GetTeam returns the Team field value


```go
func (o *IndexTeamRequest) GetTeamOk() (*TeamInfoDefinition, bool)
```
GetTeamOk returns a tuple with the Team field value and a boolean to check
if the value has been set.


```go
func (o *IndexTeamRequest) GetVersion() int64
```
GetVersion returns the Version field value if set, zero value otherwise.


```go
func (o *IndexTeamRequest) GetVersionOk() (*int64, bool)
```
GetVersionOk returns a tuple with the Version field value if set, nil
otherwise and a boolean to check if the value has been set.


```go
func (o *IndexTeamRequest) HasVersion() bool
```
HasVersion returns a boolean if a field has been set.


```go
func (o IndexTeamRequest) MarshalJSON() ([]byte, error)
```


```go
func (o *IndexTeamRequest) SetTeam(v TeamInfoDefinition)
```
SetTeam sets field value


```go
func (o *IndexTeamRequest) SetVersion(v int64)
```
SetVersion gets a reference to the given int64 and assigns it to the Version
field.




### Type IndexUserRequest
```go
type IndexUserRequest struct {
	// Version number for document for optimistic concurrency control. If absent or 0 then no version checks are done.
	Version *int64 `json:"version,omitempty"`
	// The datasource for which the user is added
	Datasource string                   `json:"datasource"`
	User       DatasourceUserDefinition `json:"user"`
}
```
IndexUserRequest Describes the request body of the /indexuser API call

### Functions

```go
func NewIndexUserRequest(datasource string, user DatasourceUserDefinition) *IndexUserRequest
```
NewIndexUserRequest instantiates a new IndexUserRequest object This
constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed


```go
func NewIndexUserRequestWithDefaults() *IndexUserRequest
```
NewIndexUserRequestWithDefaults instantiates a new IndexUserRequest object
This constructor will only assign default values to properties that have it
defined, but it doesn't guarantee that properties required by API are set



### Methods

```go
func (o *IndexUserRequest) GetDatasource() string
```
GetDatasource returns the Datasource field value


```go
func (o *IndexUserRequest) GetDatasourceOk() (*string, bool)
```
GetDatasourceOk returns a tuple with the Datasource field value and a
boolean to check if the value has been set.


```go
func (o *IndexUserRequest) GetUser() DatasourceUserDefinition
```
GetUser returns the User field value


```go
func (o *IndexUserRequest) GetUserOk() (*DatasourceUserDefinition, bool)
```
GetUserOk returns a tuple with the User field value and a boolean to check
if the value has been set.


```go
func (o *IndexUserRequest) GetVersion() int64
```
GetVersion returns the Version field value if set, zero value otherwise.


```go
func (o *IndexUserRequest) GetVersionOk() (*int64, bool)
```
GetVersionOk returns a tuple with the Version field value if set, nil
otherwise and a boolean to check if the value has been set.


```go
func (o *IndexUserRequest) HasVersion() bool
```
HasVersion returns a boolean if a field has been set.


```go
func (o IndexUserRequest) MarshalJSON() ([]byte, error)
```


```go
func (o *IndexUserRequest) SetDatasource(v string)
```
SetDatasource sets field value


```go
func (o *IndexUserRequest) SetUser(v DatasourceUserDefinition)
```
SetUser sets field value


```go
func (o *IndexUserRequest) SetVersion(v int64)
```
SetVersion gets a reference to the given int64 and assigns it to the Version
field.




### Type NullableAdditionalFieldDefinition
```go
type NullableAdditionalFieldDefinition struct {
	// contains filtered or unexported fields
}
```

### Functions

```go
func NewNullableAdditionalFieldDefinition(val *AdditionalFieldDefinition) *NullableAdditionalFieldDefinition
```



### Methods

```go
func (v NullableAdditionalFieldDefinition) Get() *AdditionalFieldDefinition
```


```go
func (v NullableAdditionalFieldDefinition) IsSet() bool
```


```go
func (v NullableAdditionalFieldDefinition) MarshalJSON() ([]byte, error)
```


```go
func (v *NullableAdditionalFieldDefinition) Set(val *AdditionalFieldDefinition)
```


```go
func (v *NullableAdditionalFieldDefinition) UnmarshalJSON(src []byte) error
```


```go
func (v *NullableAdditionalFieldDefinition) Unset()
```




### Type NullableBool
```go
type NullableBool struct {
	// contains filtered or unexported fields
}
```

### Functions

```go
func NewNullableBool(val *bool) *NullableBool
```



### Methods

```go
func (v NullableBool) Get() *bool
```


```go
func (v NullableBool) IsSet() bool
```


```go
func (v NullableBool) MarshalJSON() ([]byte, error)
```


```go
func (v *NullableBool) Set(val *bool)
```


```go
func (v *NullableBool) UnmarshalJSON(src []byte) error
```


```go
func (v *NullableBool) Unset()
```




### Type NullableBulkIndexDocumentsRequest
```go
type NullableBulkIndexDocumentsRequest struct {
	// contains filtered or unexported fields
}
```

### Functions

```go
func NewNullableBulkIndexDocumentsRequest(val *BulkIndexDocumentsRequest) *NullableBulkIndexDocumentsRequest
```



### Methods

```go
func (v NullableBulkIndexDocumentsRequest) Get() *BulkIndexDocumentsRequest
```


```go
func (v NullableBulkIndexDocumentsRequest) IsSet() bool
```


```go
func (v NullableBulkIndexDocumentsRequest) MarshalJSON() ([]byte, error)
```


```go
func (v *NullableBulkIndexDocumentsRequest) Set(val *BulkIndexDocumentsRequest)
```


```go
func (v *NullableBulkIndexDocumentsRequest) UnmarshalJSON(src []byte) error
```


```go
func (v *NullableBulkIndexDocumentsRequest) Unset()
```




### Type NullableBulkIndexDocumentsRequestAllOf
```go
type NullableBulkIndexDocumentsRequestAllOf struct {
	// contains filtered or unexported fields
}
```

### Functions

```go
func NewNullableBulkIndexDocumentsRequestAllOf(val *BulkIndexDocumentsRequestAllOf) *NullableBulkIndexDocumentsRequestAllOf
```



### Methods

```go
func (v NullableBulkIndexDocumentsRequestAllOf) Get() *BulkIndexDocumentsRequestAllOf
```


```go
func (v NullableBulkIndexDocumentsRequestAllOf) IsSet() bool
```


```go
func (v NullableBulkIndexDocumentsRequestAllOf) MarshalJSON() ([]byte, error)
```


```go
func (v *NullableBulkIndexDocumentsRequestAllOf) Set(val *BulkIndexDocumentsRequestAllOf)
```


```go
func (v *NullableBulkIndexDocumentsRequestAllOf) UnmarshalJSON(src []byte) error
```


```go
func (v *NullableBulkIndexDocumentsRequestAllOf) Unset()
```




### Type NullableBulkIndexEmployeesRequest
```go
type NullableBulkIndexEmployeesRequest struct {
	// contains filtered or unexported fields
}
```

### Functions

```go
func NewNullableBulkIndexEmployeesRequest(val *BulkIndexEmployeesRequest) *NullableBulkIndexEmployeesRequest
```



### Methods

```go
func (v NullableBulkIndexEmployeesRequest) Get() *BulkIndexEmployeesRequest
```


```go
func (v NullableBulkIndexEmployeesRequest) IsSet() bool
```


```go
func (v NullableBulkIndexEmployeesRequest) MarshalJSON() ([]byte, error)
```


```go
func (v *NullableBulkIndexEmployeesRequest) Set(val *BulkIndexEmployeesRequest)
```


```go
func (v *NullableBulkIndexEmployeesRequest) UnmarshalJSON(src []byte) error
```


```go
func (v *NullableBulkIndexEmployeesRequest) Unset()
```




### Type NullableBulkIndexEmployeesRequestAllOf
```go
type NullableBulkIndexEmployeesRequestAllOf struct {
	// contains filtered or unexported fields
}
```

### Functions

```go
func NewNullableBulkIndexEmployeesRequestAllOf(val *BulkIndexEmployeesRequestAllOf) *NullableBulkIndexEmployeesRequestAllOf
```



### Methods

```go
func (v NullableBulkIndexEmployeesRequestAllOf) Get() *BulkIndexEmployeesRequestAllOf
```


```go
func (v NullableBulkIndexEmployeesRequestAllOf) IsSet() bool
```


```go
func (v NullableBulkIndexEmployeesRequestAllOf) MarshalJSON() ([]byte, error)
```


```go
func (v *NullableBulkIndexEmployeesRequestAllOf) Set(val *BulkIndexEmployeesRequestAllOf)
```


```go
func (v *NullableBulkIndexEmployeesRequestAllOf) UnmarshalJSON(src []byte) error
```


```go
func (v *NullableBulkIndexEmployeesRequestAllOf) Unset()
```




### Type NullableBulkIndexGroupsRequest
```go
type NullableBulkIndexGroupsRequest struct {
	// contains filtered or unexported fields
}
```

### Functions

```go
func NewNullableBulkIndexGroupsRequest(val *BulkIndexGroupsRequest) *NullableBulkIndexGroupsRequest
```



### Methods

```go
func (v NullableBulkIndexGroupsRequest) Get() *BulkIndexGroupsRequest
```


```go
func (v NullableBulkIndexGroupsRequest) IsSet() bool
```


```go
func (v NullableBulkIndexGroupsRequest) MarshalJSON() ([]byte, error)
```


```go
func (v *NullableBulkIndexGroupsRequest) Set(val *BulkIndexGroupsRequest)
```


```go
func (v *NullableBulkIndexGroupsRequest) UnmarshalJSON(src []byte) error
```


```go
func (v *NullableBulkIndexGroupsRequest) Unset()
```




### Type NullableBulkIndexMembershipsRequest
```go
type NullableBulkIndexMembershipsRequest struct {
	// contains filtered or unexported fields
}
```

### Functions

```go
func NewNullableBulkIndexMembershipsRequest(val *BulkIndexMembershipsRequest) *NullableBulkIndexMembershipsRequest
```



### Methods

```go
func (v NullableBulkIndexMembershipsRequest) Get() *BulkIndexMembershipsRequest
```


```go
func (v NullableBulkIndexMembershipsRequest) IsSet() bool
```


```go
func (v NullableBulkIndexMembershipsRequest) MarshalJSON() ([]byte, error)
```


```go
func (v *NullableBulkIndexMembershipsRequest) Set(val *BulkIndexMembershipsRequest)
```


```go
func (v *NullableBulkIndexMembershipsRequest) UnmarshalJSON(src []byte) error
```


```go
func (v *NullableBulkIndexMembershipsRequest) Unset()
```




### Type NullableBulkIndexRequest
```go
type NullableBulkIndexRequest struct {
	// contains filtered or unexported fields
}
```

### Functions

```go
func NewNullableBulkIndexRequest(val *BulkIndexRequest) *NullableBulkIndexRequest
```



### Methods

```go
func (v NullableBulkIndexRequest) Get() *BulkIndexRequest
```


```go
func (v NullableBulkIndexRequest) IsSet() bool
```


```go
func (v NullableBulkIndexRequest) MarshalJSON() ([]byte, error)
```


```go
func (v *NullableBulkIndexRequest) Set(val *BulkIndexRequest)
```


```go
func (v *NullableBulkIndexRequest) UnmarshalJSON(src []byte) error
```


```go
func (v *NullableBulkIndexRequest) Unset()
```




### Type NullableBulkIndexTeamsRequest
```go
type NullableBulkIndexTeamsRequest struct {
	// contains filtered or unexported fields
}
```

### Functions

```go
func NewNullableBulkIndexTeamsRequest(val *BulkIndexTeamsRequest) *NullableBulkIndexTeamsRequest
```



### Methods

```go
func (v NullableBulkIndexTeamsRequest) Get() *BulkIndexTeamsRequest
```


```go
func (v NullableBulkIndexTeamsRequest) IsSet() bool
```


```go
func (v NullableBulkIndexTeamsRequest) MarshalJSON() ([]byte, error)
```


```go
func (v *NullableBulkIndexTeamsRequest) Set(val *BulkIndexTeamsRequest)
```


```go
func (v *NullableBulkIndexTeamsRequest) UnmarshalJSON(src []byte) error
```


```go
func (v *NullableBulkIndexTeamsRequest) Unset()
```




### Type NullableBulkIndexTeamsRequestAllOf
```go
type NullableBulkIndexTeamsRequestAllOf struct {
	// contains filtered or unexported fields
}
```

### Functions

```go
func NewNullableBulkIndexTeamsRequestAllOf(val *BulkIndexTeamsRequestAllOf) *NullableBulkIndexTeamsRequestAllOf
```



### Methods

```go
func (v NullableBulkIndexTeamsRequestAllOf) Get() *BulkIndexTeamsRequestAllOf
```


```go
func (v NullableBulkIndexTeamsRequestAllOf) IsSet() bool
```


```go
func (v NullableBulkIndexTeamsRequestAllOf) MarshalJSON() ([]byte, error)
```


```go
func (v *NullableBulkIndexTeamsRequestAllOf) Set(val *BulkIndexTeamsRequestAllOf)
```


```go
func (v *NullableBulkIndexTeamsRequestAllOf) UnmarshalJSON(src []byte) error
```


```go
func (v *NullableBulkIndexTeamsRequestAllOf) Unset()
```




### Type NullableBulkIndexUsersRequest
```go
type NullableBulkIndexUsersRequest struct {
	// contains filtered or unexported fields
}
```

### Functions

```go
func NewNullableBulkIndexUsersRequest(val *BulkIndexUsersRequest) *NullableBulkIndexUsersRequest
```



### Methods

```go
func (v NullableBulkIndexUsersRequest) Get() *BulkIndexUsersRequest
```


```go
func (v NullableBulkIndexUsersRequest) IsSet() bool
```


```go
func (v NullableBulkIndexUsersRequest) MarshalJSON() ([]byte, error)
```


```go
func (v *NullableBulkIndexUsersRequest) Set(val *BulkIndexUsersRequest)
```


```go
func (v *NullableBulkIndexUsersRequest) UnmarshalJSON(src []byte) error
```


```go
func (v *NullableBulkIndexUsersRequest) Unset()
```




### Type NullableCanonicalizingRegexType
```go
type NullableCanonicalizingRegexType struct {
	// contains filtered or unexported fields
}
```

### Functions

```go
func NewNullableCanonicalizingRegexType(val *CanonicalizingRegexType) *NullableCanonicalizingRegexType
```



### Methods

```go
func (v NullableCanonicalizingRegexType) Get() *CanonicalizingRegexType
```


```go
func (v NullableCanonicalizingRegexType) IsSet() bool
```


```go
func (v NullableCanonicalizingRegexType) MarshalJSON() ([]byte, error)
```


```go
func (v *NullableCanonicalizingRegexType) Set(val *CanonicalizingRegexType)
```


```go
func (v *NullableCanonicalizingRegexType) UnmarshalJSON(src []byte) error
```


```go
func (v *NullableCanonicalizingRegexType) Unset()
```




### Type NullableCheckDocumentAccessRequest
```go
type NullableCheckDocumentAccessRequest struct {
	// contains filtered or unexported fields
}
```

### Functions

```go
func NewNullableCheckDocumentAccessRequest(val *CheckDocumentAccessRequest) *NullableCheckDocumentAccessRequest
```



### Methods

```go
func (v NullableCheckDocumentAccessRequest) Get() *CheckDocumentAccessRequest
```


```go
func (v NullableCheckDocumentAccessRequest) IsSet() bool
```


```go
func (v NullableCheckDocumentAccessRequest) MarshalJSON() ([]byte, error)
```


```go
func (v *NullableCheckDocumentAccessRequest) Set(val *CheckDocumentAccessRequest)
```


```go
func (v *NullableCheckDocumentAccessRequest) UnmarshalJSON(src []byte) error
```


```go
func (v *NullableCheckDocumentAccessRequest) Unset()
```




### Type NullableCheckDocumentAccessResponse
```go
type NullableCheckDocumentAccessResponse struct {
	// contains filtered or unexported fields
}
```

### Functions

```go
func NewNullableCheckDocumentAccessResponse(val *CheckDocumentAccessResponse) *NullableCheckDocumentAccessResponse
```



### Methods

```go
func (v NullableCheckDocumentAccessResponse) Get() *CheckDocumentAccessResponse
```


```go
func (v NullableCheckDocumentAccessResponse) IsSet() bool
```


```go
func (v NullableCheckDocumentAccessResponse) MarshalJSON() ([]byte, error)
```


```go
func (v *NullableCheckDocumentAccessResponse) Set(val *CheckDocumentAccessResponse)
```


```go
func (v *NullableCheckDocumentAccessResponse) UnmarshalJSON(src []byte) error
```


```go
func (v *NullableCheckDocumentAccessResponse) Unset()
```




### Type NullableConnectorType
```go
type NullableConnectorType struct {
	// contains filtered or unexported fields
}
```

### Functions

```go
func NewNullableConnectorType(val *ConnectorType) *NullableConnectorType
```



### Methods

```go
func (v NullableConnectorType) Get() *ConnectorType
```


```go
func (v NullableConnectorType) IsSet() bool
```


```go
func (v NullableConnectorType) MarshalJSON() ([]byte, error)
```


```go
func (v *NullableConnectorType) Set(val *ConnectorType)
```


```go
func (v *NullableConnectorType) UnmarshalJSON(src []byte) error
```


```go
func (v *NullableConnectorType) Unset()
```




### Type NullableContentDefinition
```go
type NullableContentDefinition struct {
	// contains filtered or unexported fields
}
```

### Functions

```go
func NewNullableContentDefinition(val *ContentDefinition) *NullableContentDefinition
```



### Methods

```go
func (v NullableContentDefinition) Get() *ContentDefinition
```


```go
func (v NullableContentDefinition) IsSet() bool
```


```go
func (v NullableContentDefinition) MarshalJSON() ([]byte, error)
```


```go
func (v *NullableContentDefinition) Set(val *ContentDefinition)
```


```go
func (v *NullableContentDefinition) UnmarshalJSON(src []byte) error
```


```go
func (v *NullableContentDefinition) Unset()
```




### Type NullableCustomDatasourceConfig
```go
type NullableCustomDatasourceConfig struct {
	// contains filtered or unexported fields
}
```

### Functions

```go
func NewNullableCustomDatasourceConfig(val *CustomDatasourceConfig) *NullableCustomDatasourceConfig
```



### Methods

```go
func (v NullableCustomDatasourceConfig) Get() *CustomDatasourceConfig
```


```go
func (v NullableCustomDatasourceConfig) IsSet() bool
```


```go
func (v NullableCustomDatasourceConfig) MarshalJSON() ([]byte, error)
```


```go
func (v *NullableCustomDatasourceConfig) Set(val *CustomDatasourceConfig)
```


```go
func (v *NullableCustomDatasourceConfig) UnmarshalJSON(src []byte) error
```


```go
func (v *NullableCustomDatasourceConfig) Unset()
```




### Type NullableCustomDatasourceConfigAllOf
```go
type NullableCustomDatasourceConfigAllOf struct {
	// contains filtered or unexported fields
}
```

### Functions

```go
func NewNullableCustomDatasourceConfigAllOf(val *CustomDatasourceConfigAllOf) *NullableCustomDatasourceConfigAllOf
```



### Methods

```go
func (v NullableCustomDatasourceConfigAllOf) Get() *CustomDatasourceConfigAllOf
```


```go
func (v NullableCustomDatasourceConfigAllOf) IsSet() bool
```


```go
func (v NullableCustomDatasourceConfigAllOf) MarshalJSON() ([]byte, error)
```


```go
func (v *NullableCustomDatasourceConfigAllOf) Set(val *CustomDatasourceConfigAllOf)
```


```go
func (v *NullableCustomDatasourceConfigAllOf) UnmarshalJSON(src []byte) error
```


```go
func (v *NullableCustomDatasourceConfigAllOf) Unset()
```




### Type NullableCustomProperty
```go
type NullableCustomProperty struct {
	// contains filtered or unexported fields
}
```

### Functions

```go
func NewNullableCustomProperty(val *CustomProperty) *NullableCustomProperty
```



### Methods

```go
func (v NullableCustomProperty) Get() *CustomProperty
```


```go
func (v NullableCustomProperty) IsSet() bool
```


```go
func (v NullableCustomProperty) MarshalJSON() ([]byte, error)
```


```go
func (v *NullableCustomProperty) Set(val *CustomProperty)
```


```go
func (v *NullableCustomProperty) UnmarshalJSON(src []byte) error
```


```go
func (v *NullableCustomProperty) Unset()
```




### Type NullableDatasourceBulkMembershipDefinition
```go
type NullableDatasourceBulkMembershipDefinition struct {
	// contains filtered or unexported fields
}
```

### Functions

```go
func NewNullableDatasourceBulkMembershipDefinition(val *DatasourceBulkMembershipDefinition) *NullableDatasourceBulkMembershipDefinition
```



### Methods

```go
func (v NullableDatasourceBulkMembershipDefinition) Get() *DatasourceBulkMembershipDefinition
```


```go
func (v NullableDatasourceBulkMembershipDefinition) IsSet() bool
```


```go
func (v NullableDatasourceBulkMembershipDefinition) MarshalJSON() ([]byte, error)
```


```go
func (v *NullableDatasourceBulkMembershipDefinition) Set(val *DatasourceBulkMembershipDefinition)
```


```go
func (v *NullableDatasourceBulkMembershipDefinition) UnmarshalJSON(src []byte) error
```


```go
func (v *NullableDatasourceBulkMembershipDefinition) Unset()
```




### Type NullableDatasourceConfigList
```go
type NullableDatasourceConfigList struct {
	// contains filtered or unexported fields
}
```

### Functions

```go
func NewNullableDatasourceConfigList(val *DatasourceConfigList) *NullableDatasourceConfigList
```



### Methods

```go
func (v NullableDatasourceConfigList) Get() *DatasourceConfigList
```


```go
func (v NullableDatasourceConfigList) IsSet() bool
```


```go
func (v NullableDatasourceConfigList) MarshalJSON() ([]byte, error)
```


```go
func (v *NullableDatasourceConfigList) Set(val *DatasourceConfigList)
```


```go
func (v *NullableDatasourceConfigList) UnmarshalJSON(src []byte) error
```


```go
func (v *NullableDatasourceConfigList) Unset()
```




### Type NullableDatasourceGroupDefinition
```go
type NullableDatasourceGroupDefinition struct {
	// contains filtered or unexported fields
}
```

### Functions

```go
func NewNullableDatasourceGroupDefinition(val *DatasourceGroupDefinition) *NullableDatasourceGroupDefinition
```



### Methods

```go
func (v NullableDatasourceGroupDefinition) Get() *DatasourceGroupDefinition
```


```go
func (v NullableDatasourceGroupDefinition) IsSet() bool
```


```go
func (v NullableDatasourceGroupDefinition) MarshalJSON() ([]byte, error)
```


```go
func (v *NullableDatasourceGroupDefinition) Set(val *DatasourceGroupDefinition)
```


```go
func (v *NullableDatasourceGroupDefinition) UnmarshalJSON(src []byte) error
```


```go
func (v *NullableDatasourceGroupDefinition) Unset()
```




### Type NullableDatasourceMembershipDefinition
```go
type NullableDatasourceMembershipDefinition struct {
	// contains filtered or unexported fields
}
```

### Functions

```go
func NewNullableDatasourceMembershipDefinition(val *DatasourceMembershipDefinition) *NullableDatasourceMembershipDefinition
```



### Methods

```go
func (v NullableDatasourceMembershipDefinition) Get() *DatasourceMembershipDefinition
```


```go
func (v NullableDatasourceMembershipDefinition) IsSet() bool
```


```go
func (v NullableDatasourceMembershipDefinition) MarshalJSON() ([]byte, error)
```


```go
func (v *NullableDatasourceMembershipDefinition) Set(val *DatasourceMembershipDefinition)
```


```go
func (v *NullableDatasourceMembershipDefinition) UnmarshalJSON(src []byte) error
```


```go
func (v *NullableDatasourceMembershipDefinition) Unset()
```




### Type NullableDatasourceProfile
```go
type NullableDatasourceProfile struct {
	// contains filtered or unexported fields
}
```

### Functions

```go
func NewNullableDatasourceProfile(val *DatasourceProfile) *NullableDatasourceProfile
```



### Methods

```go
func (v NullableDatasourceProfile) Get() *DatasourceProfile
```


```go
func (v NullableDatasourceProfile) IsSet() bool
```


```go
func (v NullableDatasourceProfile) MarshalJSON() ([]byte, error)
```


```go
func (v *NullableDatasourceProfile) Set(val *DatasourceProfile)
```


```go
func (v *NullableDatasourceProfile) UnmarshalJSON(src []byte) error
```


```go
func (v *NullableDatasourceProfile) Unset()
```




### Type NullableDatasourceUserDefinition
```go
type NullableDatasourceUserDefinition struct {
	// contains filtered or unexported fields
}
```

### Functions

```go
func NewNullableDatasourceUserDefinition(val *DatasourceUserDefinition) *NullableDatasourceUserDefinition
```



### Methods

```go
func (v NullableDatasourceUserDefinition) Get() *DatasourceUserDefinition
```


```go
func (v NullableDatasourceUserDefinition) IsSet() bool
```


```go
func (v NullableDatasourceUserDefinition) MarshalJSON() ([]byte, error)
```


```go
func (v *NullableDatasourceUserDefinition) Set(val *DatasourceUserDefinition)
```


```go
func (v *NullableDatasourceUserDefinition) UnmarshalJSON(src []byte) error
```


```go
func (v *NullableDatasourceUserDefinition) Unset()
```




### Type NullableDeleteDocumentRequest
```go
type NullableDeleteDocumentRequest struct {
	// contains filtered or unexported fields
}
```

### Functions

```go
func NewNullableDeleteDocumentRequest(val *DeleteDocumentRequest) *NullableDeleteDocumentRequest
```



### Methods

```go
func (v NullableDeleteDocumentRequest) Get() *DeleteDocumentRequest
```


```go
func (v NullableDeleteDocumentRequest) IsSet() bool
```


```go
func (v NullableDeleteDocumentRequest) MarshalJSON() ([]byte, error)
```


```go
func (v *NullableDeleteDocumentRequest) Set(val *DeleteDocumentRequest)
```


```go
func (v *NullableDeleteDocumentRequest) UnmarshalJSON(src []byte) error
```


```go
func (v *NullableDeleteDocumentRequest) Unset()
```




### Type NullableDeleteEmployeeRequest
```go
type NullableDeleteEmployeeRequest struct {
	// contains filtered or unexported fields
}
```

### Functions

```go
func NewNullableDeleteEmployeeRequest(val *DeleteEmployeeRequest) *NullableDeleteEmployeeRequest
```



### Methods

```go
func (v NullableDeleteEmployeeRequest) Get() *DeleteEmployeeRequest
```


```go
func (v NullableDeleteEmployeeRequest) IsSet() bool
```


```go
func (v NullableDeleteEmployeeRequest) MarshalJSON() ([]byte, error)
```


```go
func (v *NullableDeleteEmployeeRequest) Set(val *DeleteEmployeeRequest)
```


```go
func (v *NullableDeleteEmployeeRequest) UnmarshalJSON(src []byte) error
```


```go
func (v *NullableDeleteEmployeeRequest) Unset()
```




### Type NullableDeleteGroupRequest
```go
type NullableDeleteGroupRequest struct {
	// contains filtered or unexported fields
}
```

### Functions

```go
func NewNullableDeleteGroupRequest(val *DeleteGroupRequest) *NullableDeleteGroupRequest
```



### Methods

```go
func (v NullableDeleteGroupRequest) Get() *DeleteGroupRequest
```


```go
func (v NullableDeleteGroupRequest) IsSet() bool
```


```go
func (v NullableDeleteGroupRequest) MarshalJSON() ([]byte, error)
```


```go
func (v *NullableDeleteGroupRequest) Set(val *DeleteGroupRequest)
```


```go
func (v *NullableDeleteGroupRequest) UnmarshalJSON(src []byte) error
```


```go
func (v *NullableDeleteGroupRequest) Unset()
```




### Type NullableDeleteMembershipRequest
```go
type NullableDeleteMembershipRequest struct {
	// contains filtered or unexported fields
}
```

### Functions

```go
func NewNullableDeleteMembershipRequest(val *DeleteMembershipRequest) *NullableDeleteMembershipRequest
```



### Methods

```go
func (v NullableDeleteMembershipRequest) Get() *DeleteMembershipRequest
```


```go
func (v NullableDeleteMembershipRequest) IsSet() bool
```


```go
func (v NullableDeleteMembershipRequest) MarshalJSON() ([]byte, error)
```


```go
func (v *NullableDeleteMembershipRequest) Set(val *DeleteMembershipRequest)
```


```go
func (v *NullableDeleteMembershipRequest) UnmarshalJSON(src []byte) error
```


```go
func (v *NullableDeleteMembershipRequest) Unset()
```




### Type NullableDeleteUserRequest
```go
type NullableDeleteUserRequest struct {
	// contains filtered or unexported fields
}
```

### Functions

```go
func NewNullableDeleteUserRequest(val *DeleteUserRequest) *NullableDeleteUserRequest
```



### Methods

```go
func (v NullableDeleteUserRequest) Get() *DeleteUserRequest
```


```go
func (v NullableDeleteUserRequest) IsSet() bool
```


```go
func (v NullableDeleteUserRequest) MarshalJSON() ([]byte, error)
```


```go
func (v *NullableDeleteUserRequest) Set(val *DeleteUserRequest)
```


```go
func (v *NullableDeleteUserRequest) UnmarshalJSON(src []byte) error
```


```go
func (v *NullableDeleteUserRequest) Unset()
```




### Type NullableDocumentDefinition
```go
type NullableDocumentDefinition struct {
	// contains filtered or unexported fields
}
```

### Functions

```go
func NewNullableDocumentDefinition(val *DocumentDefinition) *NullableDocumentDefinition
```



### Methods

```go
func (v NullableDocumentDefinition) Get() *DocumentDefinition
```


```go
func (v NullableDocumentDefinition) IsSet() bool
```


```go
func (v NullableDocumentDefinition) MarshalJSON() ([]byte, error)
```


```go
func (v *NullableDocumentDefinition) Set(val *DocumentDefinition)
```


```go
func (v *NullableDocumentDefinition) UnmarshalJSON(src []byte) error
```


```go
func (v *NullableDocumentDefinition) Unset()
```




### Type NullableDocumentInteractionsDefinition
```go
type NullableDocumentInteractionsDefinition struct {
	// contains filtered or unexported fields
}
```

### Functions

```go
func NewNullableDocumentInteractionsDefinition(val *DocumentInteractionsDefinition) *NullableDocumentInteractionsDefinition
```



### Methods

```go
func (v NullableDocumentInteractionsDefinition) Get() *DocumentInteractionsDefinition
```


```go
func (v NullableDocumentInteractionsDefinition) IsSet() bool
```


```go
func (v NullableDocumentInteractionsDefinition) MarshalJSON() ([]byte, error)
```


```go
func (v *NullableDocumentInteractionsDefinition) Set(val *DocumentInteractionsDefinition)
```


```go
func (v *NullableDocumentInteractionsDefinition) UnmarshalJSON(src []byte) error
```


```go
func (v *NullableDocumentInteractionsDefinition) Unset()
```




### Type NullableDocumentPermissionsDefinition
```go
type NullableDocumentPermissionsDefinition struct {
	// contains filtered or unexported fields
}
```

### Functions

```go
func NewNullableDocumentPermissionsDefinition(val *DocumentPermissionsDefinition) *NullableDocumentPermissionsDefinition
```



### Methods

```go
func (v NullableDocumentPermissionsDefinition) Get() *DocumentPermissionsDefinition
```


```go
func (v NullableDocumentPermissionsDefinition) IsSet() bool
```


```go
func (v NullableDocumentPermissionsDefinition) MarshalJSON() ([]byte, error)
```


```go
func (v *NullableDocumentPermissionsDefinition) Set(val *DocumentPermissionsDefinition)
```


```go
func (v *NullableDocumentPermissionsDefinition) UnmarshalJSON(src []byte) error
```


```go
func (v *NullableDocumentPermissionsDefinition) Unset()
```




### Type NullableEmployeeAndVersionDefinition
```go
type NullableEmployeeAndVersionDefinition struct {
	// contains filtered or unexported fields
}
```

### Functions

```go
func NewNullableEmployeeAndVersionDefinition(val *EmployeeAndVersionDefinition) *NullableEmployeeAndVersionDefinition
```



### Methods

```go
func (v NullableEmployeeAndVersionDefinition) Get() *EmployeeAndVersionDefinition
```


```go
func (v NullableEmployeeAndVersionDefinition) IsSet() bool
```


```go
func (v NullableEmployeeAndVersionDefinition) MarshalJSON() ([]byte, error)
```


```go
func (v *NullableEmployeeAndVersionDefinition) Set(val *EmployeeAndVersionDefinition)
```


```go
func (v *NullableEmployeeAndVersionDefinition) UnmarshalJSON(src []byte) error
```


```go
func (v *NullableEmployeeAndVersionDefinition) Unset()
```




### Type NullableEmployeeInfoDefinition
```go
type NullableEmployeeInfoDefinition struct {
	// contains filtered or unexported fields
}
```

### Functions

```go
func NewNullableEmployeeInfoDefinition(val *EmployeeInfoDefinition) *NullableEmployeeInfoDefinition
```



### Methods

```go
func (v NullableEmployeeInfoDefinition) Get() *EmployeeInfoDefinition
```


```go
func (v NullableEmployeeInfoDefinition) IsSet() bool
```


```go
func (v NullableEmployeeInfoDefinition) MarshalJSON() ([]byte, error)
```


```go
func (v *NullableEmployeeInfoDefinition) Set(val *EmployeeInfoDefinition)
```


```go
func (v *NullableEmployeeInfoDefinition) UnmarshalJSON(src []byte) error
```


```go
func (v *NullableEmployeeInfoDefinition) Unset()
```




### Type NullableEmployeeTeamInfo
```go
type NullableEmployeeTeamInfo struct {
	// contains filtered or unexported fields
}
```

### Functions

```go
func NewNullableEmployeeTeamInfo(val *EmployeeTeamInfo) *NullableEmployeeTeamInfo
```



### Methods

```go
func (v NullableEmployeeTeamInfo) Get() *EmployeeTeamInfo
```


```go
func (v NullableEmployeeTeamInfo) IsSet() bool
```


```go
func (v NullableEmployeeTeamInfo) MarshalJSON() ([]byte, error)
```


```go
func (v *NullableEmployeeTeamInfo) Set(val *EmployeeTeamInfo)
```


```go
func (v *NullableEmployeeTeamInfo) UnmarshalJSON(src []byte) error
```


```go
func (v *NullableEmployeeTeamInfo) Unset()
```




### Type NullableFloat32
```go
type NullableFloat32 struct {
	// contains filtered or unexported fields
}
```

### Functions

```go
func NewNullableFloat32(val *float32) *NullableFloat32
```



### Methods

```go
func (v NullableFloat32) Get() *float32
```


```go
func (v NullableFloat32) IsSet() bool
```


```go
func (v NullableFloat32) MarshalJSON() ([]byte, error)
```


```go
func (v *NullableFloat32) Set(val *float32)
```


```go
func (v *NullableFloat32) UnmarshalJSON(src []byte) error
```


```go
func (v *NullableFloat32) Unset()
```




### Type NullableFloat64
```go
type NullableFloat64 struct {
	// contains filtered or unexported fields
}
```

### Functions

```go
func NewNullableFloat64(val *float64) *NullableFloat64
```



### Methods

```go
func (v NullableFloat64) Get() *float64
```


```go
func (v NullableFloat64) IsSet() bool
```


```go
func (v NullableFloat64) MarshalJSON() ([]byte, error)
```


```go
func (v *NullableFloat64) Set(val *float64)
```


```go
func (v *NullableFloat64) UnmarshalJSON(src []byte) error
```


```go
func (v *NullableFloat64) Unset()
```




### Type NullableGetDatasourceConfigRequest
```go
type NullableGetDatasourceConfigRequest struct {
	// contains filtered or unexported fields
}
```

### Functions

```go
func NewNullableGetDatasourceConfigRequest(val *GetDatasourceConfigRequest) *NullableGetDatasourceConfigRequest
```



### Methods

```go
func (v NullableGetDatasourceConfigRequest) Get() *GetDatasourceConfigRequest
```


```go
func (v NullableGetDatasourceConfigRequest) IsSet() bool
```


```go
func (v NullableGetDatasourceConfigRequest) MarshalJSON() ([]byte, error)
```


```go
func (v *NullableGetDatasourceConfigRequest) Set(val *GetDatasourceConfigRequest)
```


```go
func (v *NullableGetDatasourceConfigRequest) UnmarshalJSON(src []byte) error
```


```go
func (v *NullableGetDatasourceConfigRequest) Unset()
```




### Type NullableGetDocumentCountRequest
```go
type NullableGetDocumentCountRequest struct {
	// contains filtered or unexported fields
}
```

### Functions

```go
func NewNullableGetDocumentCountRequest(val *GetDocumentCountRequest) *NullableGetDocumentCountRequest
```



### Methods

```go
func (v NullableGetDocumentCountRequest) Get() *GetDocumentCountRequest
```


```go
func (v NullableGetDocumentCountRequest) IsSet() bool
```


```go
func (v NullableGetDocumentCountRequest) MarshalJSON() ([]byte, error)
```


```go
func (v *NullableGetDocumentCountRequest) Set(val *GetDocumentCountRequest)
```


```go
func (v *NullableGetDocumentCountRequest) UnmarshalJSON(src []byte) error
```


```go
func (v *NullableGetDocumentCountRequest) Unset()
```




### Type NullableGetDocumentCountResponse
```go
type NullableGetDocumentCountResponse struct {
	// contains filtered or unexported fields
}
```

### Functions

```go
func NewNullableGetDocumentCountResponse(val *GetDocumentCountResponse) *NullableGetDocumentCountResponse
```



### Methods

```go
func (v NullableGetDocumentCountResponse) Get() *GetDocumentCountResponse
```


```go
func (v NullableGetDocumentCountResponse) IsSet() bool
```


```go
func (v NullableGetDocumentCountResponse) MarshalJSON() ([]byte, error)
```


```go
func (v *NullableGetDocumentCountResponse) Set(val *GetDocumentCountResponse)
```


```go
func (v *NullableGetDocumentCountResponse) UnmarshalJSON(src []byte) error
```


```go
func (v *NullableGetDocumentCountResponse) Unset()
```




### Type NullableGetDocumentStatusRequest
```go
type NullableGetDocumentStatusRequest struct {
	// contains filtered or unexported fields
}
```

### Functions

```go
func NewNullableGetDocumentStatusRequest(val *GetDocumentStatusRequest) *NullableGetDocumentStatusRequest
```



### Methods

```go
func (v NullableGetDocumentStatusRequest) Get() *GetDocumentStatusRequest
```


```go
func (v NullableGetDocumentStatusRequest) IsSet() bool
```


```go
func (v NullableGetDocumentStatusRequest) MarshalJSON() ([]byte, error)
```


```go
func (v *NullableGetDocumentStatusRequest) Set(val *GetDocumentStatusRequest)
```


```go
func (v *NullableGetDocumentStatusRequest) UnmarshalJSON(src []byte) error
```


```go
func (v *NullableGetDocumentStatusRequest) Unset()
```




### Type NullableGetDocumentStatusResponse
```go
type NullableGetDocumentStatusResponse struct {
	// contains filtered or unexported fields
}
```

### Functions

```go
func NewNullableGetDocumentStatusResponse(val *GetDocumentStatusResponse) *NullableGetDocumentStatusResponse
```



### Methods

```go
func (v NullableGetDocumentStatusResponse) Get() *GetDocumentStatusResponse
```


```go
func (v NullableGetDocumentStatusResponse) IsSet() bool
```


```go
func (v NullableGetDocumentStatusResponse) MarshalJSON() ([]byte, error)
```


```go
func (v *NullableGetDocumentStatusResponse) Set(val *GetDocumentStatusResponse)
```


```go
func (v *NullableGetDocumentStatusResponse) UnmarshalJSON(src []byte) error
```


```go
func (v *NullableGetDocumentStatusResponse) Unset()
```




### Type NullableGetUserCountRequest
```go
type NullableGetUserCountRequest struct {
	// contains filtered or unexported fields
}
```

### Functions

```go
func NewNullableGetUserCountRequest(val *GetUserCountRequest) *NullableGetUserCountRequest
```



### Methods

```go
func (v NullableGetUserCountRequest) Get() *GetUserCountRequest
```


```go
func (v NullableGetUserCountRequest) IsSet() bool
```


```go
func (v NullableGetUserCountRequest) MarshalJSON() ([]byte, error)
```


```go
func (v *NullableGetUserCountRequest) Set(val *GetUserCountRequest)
```


```go
func (v *NullableGetUserCountRequest) UnmarshalJSON(src []byte) error
```


```go
func (v *NullableGetUserCountRequest) Unset()
```




### Type NullableGetUserCountResponse
```go
type NullableGetUserCountResponse struct {
	// contains filtered or unexported fields
}
```

### Functions

```go
func NewNullableGetUserCountResponse(val *GetUserCountResponse) *NullableGetUserCountResponse
```



### Methods

```go
func (v NullableGetUserCountResponse) Get() *GetUserCountResponse
```


```go
func (v NullableGetUserCountResponse) IsSet() bool
```


```go
func (v NullableGetUserCountResponse) MarshalJSON() ([]byte, error)
```


```go
func (v *NullableGetUserCountResponse) Set(val *GetUserCountResponse)
```


```go
func (v *NullableGetUserCountResponse) UnmarshalJSON(src []byte) error
```


```go
func (v *NullableGetUserCountResponse) Unset()
```




### Type NullableGreenlistUsersRequest
```go
type NullableGreenlistUsersRequest struct {
	// contains filtered or unexported fields
}
```

### Functions

```go
func NewNullableGreenlistUsersRequest(val *GreenlistUsersRequest) *NullableGreenlistUsersRequest
```



### Methods

```go
func (v NullableGreenlistUsersRequest) Get() *GreenlistUsersRequest
```


```go
func (v NullableGreenlistUsersRequest) IsSet() bool
```


```go
func (v NullableGreenlistUsersRequest) MarshalJSON() ([]byte, error)
```


```go
func (v *NullableGreenlistUsersRequest) Set(val *GreenlistUsersRequest)
```


```go
func (v *NullableGreenlistUsersRequest) UnmarshalJSON(src []byte) error
```


```go
func (v *NullableGreenlistUsersRequest) Unset()
```




### Type NullableHypertextField
```go
type NullableHypertextField struct {
	// contains filtered or unexported fields
}
```

### Functions

```go
func NewNullableHypertextField(val *HypertextField) *NullableHypertextField
```



### Methods

```go
func (v NullableHypertextField) Get() *HypertextField
```


```go
func (v NullableHypertextField) IsSet() bool
```


```go
func (v NullableHypertextField) MarshalJSON() ([]byte, error)
```


```go
func (v *NullableHypertextField) Set(val *HypertextField)
```


```go
func (v *NullableHypertextField) UnmarshalJSON(src []byte) error
```


```go
func (v *NullableHypertextField) Unset()
```




### Type NullableIconConfig
```go
type NullableIconConfig struct {
	// contains filtered or unexported fields
}
```

### Functions

```go
func NewNullableIconConfig(val *IconConfig) *NullableIconConfig
```



### Methods

```go
func (v NullableIconConfig) Get() *IconConfig
```


```go
func (v NullableIconConfig) IsSet() bool
```


```go
func (v NullableIconConfig) MarshalJSON() ([]byte, error)
```


```go
func (v *NullableIconConfig) Set(val *IconConfig)
```


```go
func (v *NullableIconConfig) UnmarshalJSON(src []byte) error
```


```go
func (v *NullableIconConfig) Unset()
```




### Type NullableIndexDocumentRequest
```go
type NullableIndexDocumentRequest struct {
	// contains filtered or unexported fields
}
```

### Functions

```go
func NewNullableIndexDocumentRequest(val *IndexDocumentRequest) *NullableIndexDocumentRequest
```



### Methods

```go
func (v NullableIndexDocumentRequest) Get() *IndexDocumentRequest
```


```go
func (v NullableIndexDocumentRequest) IsSet() bool
```


```go
func (v NullableIndexDocumentRequest) MarshalJSON() ([]byte, error)
```


```go
func (v *NullableIndexDocumentRequest) Set(val *IndexDocumentRequest)
```


```go
func (v *NullableIndexDocumentRequest) UnmarshalJSON(src []byte) error
```


```go
func (v *NullableIndexDocumentRequest) Unset()
```




### Type NullableIndexEmployeeListRequest
```go
type NullableIndexEmployeeListRequest struct {
	// contains filtered or unexported fields
}
```

### Functions

```go
func NewNullableIndexEmployeeListRequest(val *IndexEmployeeListRequest) *NullableIndexEmployeeListRequest
```



### Methods

```go
func (v NullableIndexEmployeeListRequest) Get() *IndexEmployeeListRequest
```


```go
func (v NullableIndexEmployeeListRequest) IsSet() bool
```


```go
func (v NullableIndexEmployeeListRequest) MarshalJSON() ([]byte, error)
```


```go
func (v *NullableIndexEmployeeListRequest) Set(val *IndexEmployeeListRequest)
```


```go
func (v *NullableIndexEmployeeListRequest) UnmarshalJSON(src []byte) error
```


```go
func (v *NullableIndexEmployeeListRequest) Unset()
```




### Type NullableIndexEmployeeRequest
```go
type NullableIndexEmployeeRequest struct {
	// contains filtered or unexported fields
}
```

### Functions

```go
func NewNullableIndexEmployeeRequest(val *IndexEmployeeRequest) *NullableIndexEmployeeRequest
```



### Methods

```go
func (v NullableIndexEmployeeRequest) Get() *IndexEmployeeRequest
```


```go
func (v NullableIndexEmployeeRequest) IsSet() bool
```


```go
func (v NullableIndexEmployeeRequest) MarshalJSON() ([]byte, error)
```


```go
func (v *NullableIndexEmployeeRequest) Set(val *IndexEmployeeRequest)
```


```go
func (v *NullableIndexEmployeeRequest) UnmarshalJSON(src []byte) error
```


```go
func (v *NullableIndexEmployeeRequest) Unset()
```




### Type NullableIndexGroupRequest
```go
type NullableIndexGroupRequest struct {
	// contains filtered or unexported fields
}
```

### Functions

```go
func NewNullableIndexGroupRequest(val *IndexGroupRequest) *NullableIndexGroupRequest
```



### Methods

```go
func (v NullableIndexGroupRequest) Get() *IndexGroupRequest
```


```go
func (v NullableIndexGroupRequest) IsSet() bool
```


```go
func (v NullableIndexGroupRequest) MarshalJSON() ([]byte, error)
```


```go
func (v *NullableIndexGroupRequest) Set(val *IndexGroupRequest)
```


```go
func (v *NullableIndexGroupRequest) UnmarshalJSON(src []byte) error
```


```go
func (v *NullableIndexGroupRequest) Unset()
```




### Type NullableIndexMembershipRequest
```go
type NullableIndexMembershipRequest struct {
	// contains filtered or unexported fields
}
```

### Functions

```go
func NewNullableIndexMembershipRequest(val *IndexMembershipRequest) *NullableIndexMembershipRequest
```



### Methods

```go
func (v NullableIndexMembershipRequest) Get() *IndexMembershipRequest
```


```go
func (v NullableIndexMembershipRequest) IsSet() bool
```


```go
func (v NullableIndexMembershipRequest) MarshalJSON() ([]byte, error)
```


```go
func (v *NullableIndexMembershipRequest) Set(val *IndexMembershipRequest)
```


```go
func (v *NullableIndexMembershipRequest) UnmarshalJSON(src []byte) error
```


```go
func (v *NullableIndexMembershipRequest) Unset()
```




### Type NullableIndexTeamRequest
```go
type NullableIndexTeamRequest struct {
	// contains filtered or unexported fields
}
```

### Functions

```go
func NewNullableIndexTeamRequest(val *IndexTeamRequest) *NullableIndexTeamRequest
```



### Methods

```go
func (v NullableIndexTeamRequest) Get() *IndexTeamRequest
```


```go
func (v NullableIndexTeamRequest) IsSet() bool
```


```go
func (v NullableIndexTeamRequest) MarshalJSON() ([]byte, error)
```


```go
func (v *NullableIndexTeamRequest) Set(val *IndexTeamRequest)
```


```go
func (v *NullableIndexTeamRequest) UnmarshalJSON(src []byte) error
```


```go
func (v *NullableIndexTeamRequest) Unset()
```




### Type NullableIndexUserRequest
```go
type NullableIndexUserRequest struct {
	// contains filtered or unexported fields
}
```

### Functions

```go
func NewNullableIndexUserRequest(val *IndexUserRequest) *NullableIndexUserRequest
```



### Methods

```go
func (v NullableIndexUserRequest) Get() *IndexUserRequest
```


```go
func (v NullableIndexUserRequest) IsSet() bool
```


```go
func (v NullableIndexUserRequest) MarshalJSON() ([]byte, error)
```


```go
func (v *NullableIndexUserRequest) Set(val *IndexUserRequest)
```


```go
func (v *NullableIndexUserRequest) UnmarshalJSON(src []byte) error
```


```go
func (v *NullableIndexUserRequest) Unset()
```




### Type NullableInt
```go
type NullableInt struct {
	// contains filtered or unexported fields
}
```

### Functions

```go
func NewNullableInt(val *int) *NullableInt
```



### Methods

```go
func (v NullableInt) Get() *int
```


```go
func (v NullableInt) IsSet() bool
```


```go
func (v NullableInt) MarshalJSON() ([]byte, error)
```


```go
func (v *NullableInt) Set(val *int)
```


```go
func (v *NullableInt) UnmarshalJSON(src []byte) error
```


```go
func (v *NullableInt) Unset()
```




### Type NullableInt32
```go
type NullableInt32 struct {
	// contains filtered or unexported fields
}
```

### Functions

```go
func NewNullableInt32(val *int32) *NullableInt32
```



### Methods

```go
func (v NullableInt32) Get() *int32
```


```go
func (v NullableInt32) IsSet() bool
```


```go
func (v NullableInt32) MarshalJSON() ([]byte, error)
```


```go
func (v *NullableInt32) Set(val *int32)
```


```go
func (v *NullableInt32) UnmarshalJSON(src []byte) error
```


```go
func (v *NullableInt32) Unset()
```




### Type NullableInt64
```go
type NullableInt64 struct {
	// contains filtered or unexported fields
}
```

### Functions

```go
func NewNullableInt64(val *int64) *NullableInt64
```



### Methods

```go
func (v NullableInt64) Get() *int64
```


```go
func (v NullableInt64) IsSet() bool
```


```go
func (v NullableInt64) MarshalJSON() ([]byte, error)
```


```go
func (v *NullableInt64) Set(val *int64)
```


```go
func (v *NullableInt64) UnmarshalJSON(src []byte) error
```


```go
func (v *NullableInt64) Unset()
```




### Type NullableObjectDefinition
```go
type NullableObjectDefinition struct {
	// contains filtered or unexported fields
}
```

### Functions

```go
func NewNullableObjectDefinition(val *ObjectDefinition) *NullableObjectDefinition
```



### Methods

```go
func (v NullableObjectDefinition) Get() *ObjectDefinition
```


```go
func (v NullableObjectDefinition) IsSet() bool
```


```go
func (v NullableObjectDefinition) MarshalJSON() ([]byte, error)
```


```go
func (v *NullableObjectDefinition) Set(val *ObjectDefinition)
```


```go
func (v *NullableObjectDefinition) UnmarshalJSON(src []byte) error
```


```go
func (v *NullableObjectDefinition) Unset()
```




### Type NullableObjectPropertyOptions
```go
type NullableObjectPropertyOptions struct {
	// contains filtered or unexported fields
}
```

### Functions

```go
func NewNullableObjectPropertyOptions(val *ObjectPropertyOptions) *NullableObjectPropertyOptions
```



### Methods

```go
func (v NullableObjectPropertyOptions) Get() *ObjectPropertyOptions
```


```go
func (v NullableObjectPropertyOptions) IsSet() bool
```


```go
func (v NullableObjectPropertyOptions) MarshalJSON() ([]byte, error)
```


```go
func (v *NullableObjectPropertyOptions) Set(val *ObjectPropertyOptions)
```


```go
func (v *NullableObjectPropertyOptions) UnmarshalJSON(src []byte) error
```


```go
func (v *NullableObjectPropertyOptions) Unset()
```




### Type NullablePermissionsGroupIntersectionDefinition
```go
type NullablePermissionsGroupIntersectionDefinition struct {
	// contains filtered or unexported fields
}
```

### Functions

```go
func NewNullablePermissionsGroupIntersectionDefinition(val *PermissionsGroupIntersectionDefinition) *NullablePermissionsGroupIntersectionDefinition
```



### Methods

```go
func (v NullablePermissionsGroupIntersectionDefinition) Get() *PermissionsGroupIntersectionDefinition
```


```go
func (v NullablePermissionsGroupIntersectionDefinition) IsSet() bool
```


```go
func (v NullablePermissionsGroupIntersectionDefinition) MarshalJSON() ([]byte, error)
```


```go
func (v *NullablePermissionsGroupIntersectionDefinition) Set(val *PermissionsGroupIntersectionDefinition)
```


```go
func (v *NullablePermissionsGroupIntersectionDefinition) UnmarshalJSON(src []byte) error
```


```go
func (v *NullablePermissionsGroupIntersectionDefinition) Unset()
```




### Type NullableProcessAllDocumentsRequest
```go
type NullableProcessAllDocumentsRequest struct {
	// contains filtered or unexported fields
}
```

### Functions

```go
func NewNullableProcessAllDocumentsRequest(val *ProcessAllDocumentsRequest) *NullableProcessAllDocumentsRequest
```



### Methods

```go
func (v NullableProcessAllDocumentsRequest) Get() *ProcessAllDocumentsRequest
```


```go
func (v NullableProcessAllDocumentsRequest) IsSet() bool
```


```go
func (v NullableProcessAllDocumentsRequest) MarshalJSON() ([]byte, error)
```


```go
func (v *NullableProcessAllDocumentsRequest) Set(val *ProcessAllDocumentsRequest)
```


```go
func (v *NullableProcessAllDocumentsRequest) UnmarshalJSON(src []byte) error
```


```go
func (v *NullableProcessAllDocumentsRequest) Unset()
```




### Type NullableProcessAllMembershipsRequest
```go
type NullableProcessAllMembershipsRequest struct {
	// contains filtered or unexported fields
}
```

### Functions

```go
func NewNullableProcessAllMembershipsRequest(val *ProcessAllMembershipsRequest) *NullableProcessAllMembershipsRequest
```



### Methods

```go
func (v NullableProcessAllMembershipsRequest) Get() *ProcessAllMembershipsRequest
```


```go
func (v NullableProcessAllMembershipsRequest) IsSet() bool
```


```go
func (v NullableProcessAllMembershipsRequest) MarshalJSON() ([]byte, error)
```


```go
func (v *NullableProcessAllMembershipsRequest) Set(val *ProcessAllMembershipsRequest)
```


```go
func (v *NullableProcessAllMembershipsRequest) UnmarshalJSON(src []byte) error
```


```go
func (v *NullableProcessAllMembershipsRequest) Unset()
```




### Type NullablePropertyDefinition
```go
type NullablePropertyDefinition struct {
	// contains filtered or unexported fields
}
```

### Functions

```go
func NewNullablePropertyDefinition(val *PropertyDefinition) *NullablePropertyDefinition
```



### Methods

```go
func (v NullablePropertyDefinition) Get() *PropertyDefinition
```


```go
func (v NullablePropertyDefinition) IsSet() bool
```


```go
func (v NullablePropertyDefinition) MarshalJSON() ([]byte, error)
```


```go
func (v *NullablePropertyDefinition) Set(val *PropertyDefinition)
```


```go
func (v *NullablePropertyDefinition) UnmarshalJSON(src []byte) error
```


```go
func (v *NullablePropertyDefinition) Unset()
```




### Type NullablePropertyGroup
```go
type NullablePropertyGroup struct {
	// contains filtered or unexported fields
}
```

### Functions

```go
func NewNullablePropertyGroup(val *PropertyGroup) *NullablePropertyGroup
```



### Methods

```go
func (v NullablePropertyGroup) Get() *PropertyGroup
```


```go
func (v NullablePropertyGroup) IsSet() bool
```


```go
func (v NullablePropertyGroup) MarshalJSON() ([]byte, error)
```


```go
func (v *NullablePropertyGroup) Set(val *PropertyGroup)
```


```go
func (v *NullablePropertyGroup) UnmarshalJSON(src []byte) error
```


```go
func (v *NullablePropertyGroup) Unset()
```




### Type NullableQuicklink
```go
type NullableQuicklink struct {
	// contains filtered or unexported fields
}
```

### Functions

```go
func NewNullableQuicklink(val *Quicklink) *NullableQuicklink
```



### Methods

```go
func (v NullableQuicklink) Get() *Quicklink
```


```go
func (v NullableQuicklink) IsSet() bool
```


```go
func (v NullableQuicklink) MarshalJSON() ([]byte, error)
```


```go
func (v *NullableQuicklink) Set(val *Quicklink)
```


```go
func (v *NullableQuicklink) UnmarshalJSON(src []byte) error
```


```go
func (v *NullableQuicklink) Unset()
```




### Type NullableSharedDatasourceConfig
```go
type NullableSharedDatasourceConfig struct {
	// contains filtered or unexported fields
}
```

### Functions

```go
func NewNullableSharedDatasourceConfig(val *SharedDatasourceConfig) *NullableSharedDatasourceConfig
```



### Methods

```go
func (v NullableSharedDatasourceConfig) Get() *SharedDatasourceConfig
```


```go
func (v NullableSharedDatasourceConfig) IsSet() bool
```


```go
func (v NullableSharedDatasourceConfig) MarshalJSON() ([]byte, error)
```


```go
func (v *NullableSharedDatasourceConfig) Set(val *SharedDatasourceConfig)
```


```go
func (v *NullableSharedDatasourceConfig) UnmarshalJSON(src []byte) error
```


```go
func (v *NullableSharedDatasourceConfig) Unset()
```




### Type NullableSharedDatasourceConfigAllOf
```go
type NullableSharedDatasourceConfigAllOf struct {
	// contains filtered or unexported fields
}
```

### Functions

```go
func NewNullableSharedDatasourceConfigAllOf(val *SharedDatasourceConfigAllOf) *NullableSharedDatasourceConfigAllOf
```



### Methods

```go
func (v NullableSharedDatasourceConfigAllOf) Get() *SharedDatasourceConfigAllOf
```


```go
func (v NullableSharedDatasourceConfigAllOf) IsSet() bool
```


```go
func (v NullableSharedDatasourceConfigAllOf) MarshalJSON() ([]byte, error)
```


```go
func (v *NullableSharedDatasourceConfigAllOf) Set(val *SharedDatasourceConfigAllOf)
```


```go
func (v *NullableSharedDatasourceConfigAllOf) UnmarshalJSON(src []byte) error
```


```go
func (v *NullableSharedDatasourceConfigAllOf) Unset()
```




### Type NullableSharedDatasourceConfigNoInstance
```go
type NullableSharedDatasourceConfigNoInstance struct {
	// contains filtered or unexported fields
}
```

### Functions

```go
func NewNullableSharedDatasourceConfigNoInstance(val *SharedDatasourceConfigNoInstance) *NullableSharedDatasourceConfigNoInstance
```



### Methods

```go
func (v NullableSharedDatasourceConfigNoInstance) Get() *SharedDatasourceConfigNoInstance
```


```go
func (v NullableSharedDatasourceConfigNoInstance) IsSet() bool
```


```go
func (v NullableSharedDatasourceConfigNoInstance) MarshalJSON() ([]byte, error)
```


```go
func (v *NullableSharedDatasourceConfigNoInstance) Set(val *SharedDatasourceConfigNoInstance)
```


```go
func (v *NullableSharedDatasourceConfigNoInstance) UnmarshalJSON(src []byte) error
```


```go
func (v *NullableSharedDatasourceConfigNoInstance) Unset()
```




### Type NullableSocialNetworkDefinition
```go
type NullableSocialNetworkDefinition struct {
	// contains filtered or unexported fields
}
```

### Functions

```go
func NewNullableSocialNetworkDefinition(val *SocialNetworkDefinition) *NullableSocialNetworkDefinition
```



### Methods

```go
func (v NullableSocialNetworkDefinition) Get() *SocialNetworkDefinition
```


```go
func (v NullableSocialNetworkDefinition) IsSet() bool
```


```go
func (v NullableSocialNetworkDefinition) MarshalJSON() ([]byte, error)
```


```go
func (v *NullableSocialNetworkDefinition) Set(val *SocialNetworkDefinition)
```


```go
func (v *NullableSocialNetworkDefinition) UnmarshalJSON(src []byte) error
```


```go
func (v *NullableSocialNetworkDefinition) Unset()
```




### Type NullableString
```go
type NullableString struct {
	// contains filtered or unexported fields
}
```

### Functions

```go
func NewNullableString(val *string) *NullableString
```



### Methods

```go
func (v NullableString) Get() *string
```


```go
func (v NullableString) IsSet() bool
```


```go
func (v NullableString) MarshalJSON() ([]byte, error)
```


```go
func (v *NullableString) Set(val *string)
```


```go
func (v *NullableString) UnmarshalJSON(src []byte) error
```


```go
func (v *NullableString) Unset()
```




### Type NullableStructuredLocation
```go
type NullableStructuredLocation struct {
	// contains filtered or unexported fields
}
```

### Functions

```go
func NewNullableStructuredLocation(val *StructuredLocation) *NullableStructuredLocation
```



### Methods

```go
func (v NullableStructuredLocation) Get() *StructuredLocation
```


```go
func (v NullableStructuredLocation) IsSet() bool
```


```go
func (v NullableStructuredLocation) MarshalJSON() ([]byte, error)
```


```go
func (v *NullableStructuredLocation) Set(val *StructuredLocation)
```


```go
func (v *NullableStructuredLocation) UnmarshalJSON(src []byte) error
```


```go
func (v *NullableStructuredLocation) Unset()
```




### Type NullableTeamEmail
```go
type NullableTeamEmail struct {
	// contains filtered or unexported fields
}
```

### Functions

```go
func NewNullableTeamEmail(val *TeamEmail) *NullableTeamEmail
```



### Methods

```go
func (v NullableTeamEmail) Get() *TeamEmail
```


```go
func (v NullableTeamEmail) IsSet() bool
```


```go
func (v NullableTeamEmail) MarshalJSON() ([]byte, error)
```


```go
func (v *NullableTeamEmail) Set(val *TeamEmail)
```


```go
func (v *NullableTeamEmail) UnmarshalJSON(src []byte) error
```


```go
func (v *NullableTeamEmail) Unset()
```




### Type NullableTeamInfoDefinition
```go
type NullableTeamInfoDefinition struct {
	// contains filtered or unexported fields
}
```

### Functions

```go
func NewNullableTeamInfoDefinition(val *TeamInfoDefinition) *NullableTeamInfoDefinition
```



### Methods

```go
func (v NullableTeamInfoDefinition) Get() *TeamInfoDefinition
```


```go
func (v NullableTeamInfoDefinition) IsSet() bool
```


```go
func (v NullableTeamInfoDefinition) MarshalJSON() ([]byte, error)
```


```go
func (v *NullableTeamInfoDefinition) Set(val *TeamInfoDefinition)
```


```go
func (v *NullableTeamInfoDefinition) UnmarshalJSON(src []byte) error
```


```go
func (v *NullableTeamInfoDefinition) Unset()
```




### Type NullableTeamMember
```go
type NullableTeamMember struct {
	// contains filtered or unexported fields
}
```

### Functions

```go
func NewNullableTeamMember(val *TeamMember) *NullableTeamMember
```



### Methods

```go
func (v NullableTeamMember) Get() *TeamMember
```


```go
func (v NullableTeamMember) IsSet() bool
```


```go
func (v NullableTeamMember) MarshalJSON() ([]byte, error)
```


```go
func (v *NullableTeamMember) Set(val *TeamMember)
```


```go
func (v *NullableTeamMember) UnmarshalJSON(src []byte) error
```


```go
func (v *NullableTeamMember) Unset()
```




### Type NullableTime
```go
type NullableTime struct {
	// contains filtered or unexported fields
}
```

### Functions

```go
func NewNullableTime(val *time.Time) *NullableTime
```



### Methods

```go
func (v NullableTime) Get() *time.Time
```


```go
func (v NullableTime) IsSet() bool
```


```go
func (v NullableTime) MarshalJSON() ([]byte, error)
```


```go
func (v *NullableTime) Set(val *time.Time)
```


```go
func (v *NullableTime) UnmarshalJSON(src []byte) error
```


```go
func (v *NullableTime) Unset()
```




### Type NullableUserReferenceDefinition
```go
type NullableUserReferenceDefinition struct {
	// contains filtered or unexported fields
}
```

### Functions

```go
func NewNullableUserReferenceDefinition(val *UserReferenceDefinition) *NullableUserReferenceDefinition
```



### Methods

```go
func (v NullableUserReferenceDefinition) Get() *UserReferenceDefinition
```


```go
func (v NullableUserReferenceDefinition) IsSet() bool
```


```go
func (v NullableUserReferenceDefinition) MarshalJSON() ([]byte, error)
```


```go
func (v *NullableUserReferenceDefinition) Set(val *UserReferenceDefinition)
```


```go
func (v *NullableUserReferenceDefinition) UnmarshalJSON(src []byte) error
```


```go
func (v *NullableUserReferenceDefinition) Unset()
```




### Type ObjectDefinition
```go
type ObjectDefinition struct {
	// Unique identifier for this `DocumentMetadata.objectType`. If omitted, this definition is used as a default for all unmatched `DocumentMetadata.objectType`s in this datasource.
	Name *string `json:"name,omitempty"`
	// The user-friendly name of the object for display.
	DisplayLabel *string `json:"displayLabel,omitempty"`
	// The document category of this object type.
	DocCategory         *string              `json:"docCategory,omitempty"`
	PropertyDefinitions []PropertyDefinition `json:"propertyDefinitions,omitempty"`
	// A list of `PropertyGroup`s belonging to this object type, which will group properties to be displayed together in the UI.
	PropertyGroups []PropertyGroup `json:"propertyGroups,omitempty"`
}
```
ObjectDefinition The definition for an `DocumentMetadata.objectType` within
a datasource.

### Functions

```go
func NewObjectDefinition() *ObjectDefinition
```
NewObjectDefinition instantiates a new ObjectDefinition object This
constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed


```go
func NewObjectDefinitionWithDefaults() *ObjectDefinition
```
NewObjectDefinitionWithDefaults instantiates a new ObjectDefinition object
This constructor will only assign default values to properties that have it
defined, but it doesn't guarantee that properties required by API are set



### Methods

```go
func (o *ObjectDefinition) GetDisplayLabel() string
```
GetDisplayLabel returns the DisplayLabel field value if set, zero value
otherwise.


```go
func (o *ObjectDefinition) GetDisplayLabelOk() (*string, bool)
```
GetDisplayLabelOk returns a tuple with the DisplayLabel field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *ObjectDefinition) GetDocCategory() string
```
GetDocCategory returns the DocCategory field value if set, zero value
otherwise.


```go
func (o *ObjectDefinition) GetDocCategoryOk() (*string, bool)
```
GetDocCategoryOk returns a tuple with the DocCategory field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *ObjectDefinition) GetName() string
```
GetName returns the Name field value if set, zero value otherwise.


```go
func (o *ObjectDefinition) GetNameOk() (*string, bool)
```
GetNameOk returns a tuple with the Name field value if set, nil otherwise
and a boolean to check if the value has been set.


```go
func (o *ObjectDefinition) GetPropertyDefinitions() []PropertyDefinition
```
GetPropertyDefinitions returns the PropertyDefinitions field value if set,
zero value otherwise.


```go
func (o *ObjectDefinition) GetPropertyDefinitionsOk() ([]PropertyDefinition, bool)
```
GetPropertyDefinitionsOk returns a tuple with the PropertyDefinitions field
value if set, nil otherwise and a boolean to check if the value has been
set.


```go
func (o *ObjectDefinition) GetPropertyGroups() []PropertyGroup
```
GetPropertyGroups returns the PropertyGroups field value if set, zero value
otherwise.


```go
func (o *ObjectDefinition) GetPropertyGroupsOk() ([]PropertyGroup, bool)
```
GetPropertyGroupsOk returns a tuple with the PropertyGroups field value if
set, nil otherwise and a boolean to check if the value has been set.


```go
func (o *ObjectDefinition) HasDisplayLabel() bool
```
HasDisplayLabel returns a boolean if a field has been set.


```go
func (o *ObjectDefinition) HasDocCategory() bool
```
HasDocCategory returns a boolean if a field has been set.


```go
func (o *ObjectDefinition) HasName() bool
```
HasName returns a boolean if a field has been set.


```go
func (o *ObjectDefinition) HasPropertyDefinitions() bool
```
HasPropertyDefinitions returns a boolean if a field has been set.


```go
func (o *ObjectDefinition) HasPropertyGroups() bool
```
HasPropertyGroups returns a boolean if a field has been set.


```go
func (o ObjectDefinition) MarshalJSON() ([]byte, error)
```


```go
func (o *ObjectDefinition) SetDisplayLabel(v string)
```
SetDisplayLabel gets a reference to the given string and assigns it to the
DisplayLabel field.


```go
func (o *ObjectDefinition) SetDocCategory(v string)
```
SetDocCategory gets a reference to the given string and assigns it to the
DocCategory field.


```go
func (o *ObjectDefinition) SetName(v string)
```
SetName gets a reference to the given string and assigns it to the Name
field.


```go
func (o *ObjectDefinition) SetPropertyDefinitions(v []PropertyDefinition)
```
SetPropertyDefinitions gets a reference to the given []PropertyDefinition
and assigns it to the PropertyDefinitions field.


```go
func (o *ObjectDefinition) SetPropertyGroups(v []PropertyGroup)
```
SetPropertyGroups gets a reference to the given []PropertyGroup and assigns
it to the PropertyGroups field.




### Type ObjectPropertyOptions
```go
type ObjectPropertyOptions struct {
	// The properties of the sub-object. These properties represent a nested object. For example, if this property represents a postal address, the subobjectProperties might be named street, city, and state.
	SubobjectProperties []PropertyDefinition `json:"subobjectProperties,omitempty"`
}
```
ObjectPropertyOptions Options for object properties.

### Functions

```go
func NewObjectPropertyOptions() *ObjectPropertyOptions
```
NewObjectPropertyOptions instantiates a new ObjectPropertyOptions object
This constructor will assign default values to properties that have it
defined, and makes sure properties required by API are set, but the set of
arguments will change when the set of required properties is changed


```go
func NewObjectPropertyOptionsWithDefaults() *ObjectPropertyOptions
```
NewObjectPropertyOptionsWithDefaults instantiates a new
ObjectPropertyOptions object This constructor will only assign default
values to properties that have it defined, but it doesn't guarantee that
properties required by API are set



### Methods

```go
func (o *ObjectPropertyOptions) GetSubobjectProperties() []PropertyDefinition
```
GetSubobjectProperties returns the SubobjectProperties field value if set,
zero value otherwise.


```go
func (o *ObjectPropertyOptions) GetSubobjectPropertiesOk() ([]PropertyDefinition, bool)
```
GetSubobjectPropertiesOk returns a tuple with the SubobjectProperties field
value if set, nil otherwise and a boolean to check if the value has been
set.


```go
func (o *ObjectPropertyOptions) HasSubobjectProperties() bool
```
HasSubobjectProperties returns a boolean if a field has been set.


```go
func (o ObjectPropertyOptions) MarshalJSON() ([]byte, error)
```


```go
func (o *ObjectPropertyOptions) SetSubobjectProperties(v []PropertyDefinition)
```
SetSubobjectProperties gets a reference to the given []PropertyDefinition
and assigns it to the SubobjectProperties field.




### Type PeopleApiService
```go
type PeopleApiService service
```
PeopleApiService PeopleApi service

### Methods

```go
func (a *PeopleApiService) BulkindexemployeesPost(ctx context.Context) ApiBulkindexemployeesPostRequest
```
BulkindexemployeesPost Bulk index employees

Replaces all the currently indexed employees using paginated batch API
calls.

    @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
    @return ApiBulkindexemployeesPostRequest


```go
func (a *PeopleApiService) BulkindexemployeesPostExecute(r ApiBulkindexemployeesPostRequest) (*http.Response, error)
```
Execute executes the request


```go
func (a *PeopleApiService) BulkindexteamsPost(ctx context.Context) ApiBulkindexteamsPostRequest
```
BulkindexteamsPost Bulk index teams

Replaces all the currently indexed teams using paginated batch API calls.

    @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
    @return ApiBulkindexteamsPostRequest


```go
func (a *PeopleApiService) BulkindexteamsPostExecute(r ApiBulkindexteamsPostRequest) (*http.Response, error)
```
Execute executes the request


```go
func (a *PeopleApiService) DeleteemployeePost(ctx context.Context) ApiDeleteemployeePostRequest
```
DeleteemployeePost Delete employee

Delete an employee. Silently succeeds if employee is not present.

    @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
    @return ApiDeleteemployeePostRequest


```go
func (a *PeopleApiService) DeleteemployeePostExecute(r ApiDeleteemployeePostRequest) (*http.Response, error)
```
Execute executes the request


```go
func (a *PeopleApiService) IndexemployeePost(ctx context.Context) ApiIndexemployeePostRequest
```
IndexemployeePost Index employee

Adds an employee or updates information about an employee

    @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
    @return ApiIndexemployeePostRequest


```go
func (a *PeopleApiService) IndexemployeePostExecute(r ApiIndexemployeePostRequest) (*http.Response, error)
```
Execute executes the request


```go
func (a *PeopleApiService) IndexemployeelistPost(ctx context.Context) ApiIndexemployeelistPostRequest
```
IndexemployeelistPost Bulk index employees

Bulk upload details of all the employees. This deletes all employees
uploaded in the prior batch. SOON TO BE DEPRECATED in favor of
/bulkindexemployees.

    @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
    @return ApiIndexemployeelistPostRequest

Deprecated


```go
func (a *PeopleApiService) IndexemployeelistPostExecute(r ApiIndexemployeelistPostRequest) (*http.Response, error)
```
Execute executes the request Deprecated


```go
func (a *PeopleApiService) IndexteamPost(ctx context.Context) ApiIndexteamPostRequest
```
IndexteamPost Index team

Adds a team or updates information about a team

    @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
    @return ApiIndexteamPostRequest


```go
func (a *PeopleApiService) IndexteamPostExecute(r ApiIndexteamPostRequest) (*http.Response, error)
```
Execute executes the request




### Type PermissionsApiService
```go
type PermissionsApiService service
```
PermissionsApiService PermissionsApi service

### Methods

```go
func (a *PermissionsApiService) BetausersPost(ctx context.Context) ApiBetausersPostRequest
```
BetausersPost Beta users

Allow the datasource be visible to the specified beta users. The default
behaviour is datasource being visible to all users if it is enabled and not
visible to any user if it is not enabled.

    @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
    @return ApiBetausersPostRequest


```go
func (a *PermissionsApiService) BetausersPostExecute(r ApiBetausersPostRequest) (*http.Response, error)
```
Execute executes the request


```go
func (a *PermissionsApiService) BulkindexgroupsPost(ctx context.Context) ApiBulkindexgroupsPostRequest
```
BulkindexgroupsPost Bulk index groups

Replaces the groups in a datasource using paginated batch API calls.

    @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
    @return ApiBulkindexgroupsPostRequest


```go
func (a *PermissionsApiService) BulkindexgroupsPostExecute(r ApiBulkindexgroupsPostRequest) (*http.Response, error)
```
Execute executes the request


```go
func (a *PermissionsApiService) BulkindexmembershipsPost(ctx context.Context) ApiBulkindexmembershipsPostRequest
```
BulkindexmembershipsPost Bulk index memberships for a group

Replaces the memberships for a group in a datasource using paginated batch
API calls.

    @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
    @return ApiBulkindexmembershipsPostRequest


```go
func (a *PermissionsApiService) BulkindexmembershipsPostExecute(r ApiBulkindexmembershipsPostRequest) (*http.Response, error)
```
Execute executes the request


```go
func (a *PermissionsApiService) BulkindexusersPost(ctx context.Context) ApiBulkindexusersPostRequest
```
BulkindexusersPost Bulk index users

Replaces the users in a datasource using paginated batch API calls.

    @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
    @return ApiBulkindexusersPostRequest


```go
func (a *PermissionsApiService) BulkindexusersPostExecute(r ApiBulkindexusersPostRequest) (*http.Response, error)
```
Execute executes the request


```go
func (a *PermissionsApiService) CheckdocumentaccessPost(ctx context.Context) ApiCheckdocumentaccessPostRequest
```
CheckdocumentaccessPost Check document access

Check if a given user has access to access a document in a custom datasource

    @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
    @return ApiCheckdocumentaccessPostRequest


```go
func (a *PermissionsApiService) CheckdocumentaccessPostExecute(r ApiCheckdocumentaccessPostRequest) (*CheckDocumentAccessResponse, *http.Response, error)
```
Execute executes the request

    @return CheckDocumentAccessResponse


```go
func (a *PermissionsApiService) DeletegroupPost(ctx context.Context) ApiDeletegroupPostRequest
```
DeletegroupPost Delete group

Delete group from the datasource. Silently succeeds if group is not present.

    @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
    @return ApiDeletegroupPostRequest


```go
func (a *PermissionsApiService) DeletegroupPostExecute(r ApiDeletegroupPostRequest) (*http.Response, error)
```
Execute executes the request


```go
func (a *PermissionsApiService) DeletemembershipPost(ctx context.Context) ApiDeletemembershipPostRequest
```
DeletemembershipPost Delete membership

Delete membership to a group in the specified datasource. Silently succeeds
if membership is not present.

    @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
    @return ApiDeletemembershipPostRequest


```go
func (a *PermissionsApiService) DeletemembershipPostExecute(r ApiDeletemembershipPostRequest) (*http.Response, error)
```
Execute executes the request


```go
func (a *PermissionsApiService) DeleteuserPost(ctx context.Context) ApiDeleteuserPostRequest
```
DeleteuserPost Delete user

Delete the user from the datasource. Silently succeeds if user is not
present.

    @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
    @return ApiDeleteuserPostRequest


```go
func (a *PermissionsApiService) DeleteuserPostExecute(r ApiDeleteuserPostRequest) (*http.Response, error)
```
Execute executes the request


```go
func (a *PermissionsApiService) GetusercountPost(ctx context.Context) ApiGetusercountPostRequest
```
GetusercountPost Get user count

Fetches user count for the specified custom datasource.

    @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
    @return ApiGetusercountPostRequest


```go
func (a *PermissionsApiService) GetusercountPostExecute(r ApiGetusercountPostRequest) (*GetUserCountResponse, *http.Response, error)
```
Execute executes the request

    @return GetUserCountResponse


```go
func (a *PermissionsApiService) IndexgroupPost(ctx context.Context) ApiIndexgroupPostRequest
```
IndexgroupPost Index group

Add or update a group in the datasource.

    @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
    @return ApiIndexgroupPostRequest


```go
func (a *PermissionsApiService) IndexgroupPostExecute(r ApiIndexgroupPostRequest) (*http.Response, error)
```
Execute executes the request


```go
func (a *PermissionsApiService) IndexmembershipPost(ctx context.Context) ApiIndexmembershipPostRequest
```
IndexmembershipPost Index membership

Add the memberships of a group in the datasource.

    @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
    @return ApiIndexmembershipPostRequest


```go
func (a *PermissionsApiService) IndexmembershipPostExecute(r ApiIndexmembershipPostRequest) (*http.Response, error)
```
Execute executes the request


```go
func (a *PermissionsApiService) IndexuserPost(ctx context.Context) ApiIndexuserPostRequest
```
IndexuserPost Index user

Adds a datasource user or updates an existing user.

    @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
    @return ApiIndexuserPostRequest


```go
func (a *PermissionsApiService) IndexuserPostExecute(r ApiIndexuserPostRequest) (*http.Response, error)
```
Execute executes the request


```go
func (a *PermissionsApiService) ProcessallmembershipsPost(ctx context.Context) ApiProcessallmembershipsPostRequest
```
ProcessallmembershipsPost Schedules the processing of group memberships

Schedules the immediate processing of all group memberships uploaded
through the indexing API. By default the uploaded group memberships will be
processed asynchronously but this API can be used to schedule processing of
all memberships on demand.

    @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
    @return ApiProcessallmembershipsPostRequest


```go
func (a *PermissionsApiService) ProcessallmembershipsPostExecute(r ApiProcessallmembershipsPostRequest) (*http.Response, error)
```
Execute executes the request




### Type PermissionsGroupIntersectionDefinition
```go
type PermissionsGroupIntersectionDefinition struct {
	RequiredGroups []string `json:"requiredGroups,omitempty"`
}
```
PermissionsGroupIntersectionDefinition describes a list of groups that are
all required in a permissions constraint

### Functions

```go
func NewPermissionsGroupIntersectionDefinition() *PermissionsGroupIntersectionDefinition
```
NewPermissionsGroupIntersectionDefinition instantiates a new
PermissionsGroupIntersectionDefinition object This constructor will assign
default values to properties that have it defined, and makes sure properties
required by API are set, but the set of arguments will change when the set
of required properties is changed


```go
func NewPermissionsGroupIntersectionDefinitionWithDefaults() *PermissionsGroupIntersectionDefinition
```
NewPermissionsGroupIntersectionDefinitionWithDefaults instantiates a new
PermissionsGroupIntersectionDefinition object This constructor will only
assign default values to properties that have it defined, but it doesn't
guarantee that properties required by API are set



### Methods

```go
func (o *PermissionsGroupIntersectionDefinition) GetRequiredGroups() []string
```
GetRequiredGroups returns the RequiredGroups field value if set, zero value
otherwise.


```go
func (o *PermissionsGroupIntersectionDefinition) GetRequiredGroupsOk() ([]string, bool)
```
GetRequiredGroupsOk returns a tuple with the RequiredGroups field value if
set, nil otherwise and a boolean to check if the value has been set.


```go
func (o *PermissionsGroupIntersectionDefinition) HasRequiredGroups() bool
```
HasRequiredGroups returns a boolean if a field has been set.


```go
func (o PermissionsGroupIntersectionDefinition) MarshalJSON() ([]byte, error)
```


```go
func (o *PermissionsGroupIntersectionDefinition) SetRequiredGroups(v []string)
```
SetRequiredGroups gets a reference to the given []string and assigns it to
the RequiredGroups field.




### Type ProcessAllDocumentsRequest
```go
type ProcessAllDocumentsRequest struct {
	// If provided, process documents only for this custom datasource. Otherwise all uploaded documents are processed.
	Datasource *string `json:"datasource,omitempty"`
}
```
ProcessAllDocumentsRequest Describes the request body of the
/processalldocuments API call

### Functions

```go
func NewProcessAllDocumentsRequest() *ProcessAllDocumentsRequest
```
NewProcessAllDocumentsRequest instantiates a new ProcessAllDocumentsRequest
object This constructor will assign default values to properties that have
it defined, and makes sure properties required by API are set, but the set
of arguments will change when the set of required properties is changed


```go
func NewProcessAllDocumentsRequestWithDefaults() *ProcessAllDocumentsRequest
```
NewProcessAllDocumentsRequestWithDefaults instantiates a new
ProcessAllDocumentsRequest object This constructor will only assign default
values to properties that have it defined, but it doesn't guarantee that
properties required by API are set



### Methods

```go
func (o *ProcessAllDocumentsRequest) GetDatasource() string
```
GetDatasource returns the Datasource field value if set, zero value
otherwise.


```go
func (o *ProcessAllDocumentsRequest) GetDatasourceOk() (*string, bool)
```
GetDatasourceOk returns a tuple with the Datasource field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *ProcessAllDocumentsRequest) HasDatasource() bool
```
HasDatasource returns a boolean if a field has been set.


```go
func (o ProcessAllDocumentsRequest) MarshalJSON() ([]byte, error)
```


```go
func (o *ProcessAllDocumentsRequest) SetDatasource(v string)
```
SetDatasource gets a reference to the given string and assigns it to the
Datasource field.




### Type ProcessAllMembershipsRequest
```go
type ProcessAllMembershipsRequest struct {
	// If provided, process group memberships only for this custom datasource. Otherwise all uploaded memberships are processed.
	Datasource *string `json:"datasource,omitempty"`
}
```
ProcessAllMembershipsRequest Describes the request body of the
/processallmemberships API call

### Functions

```go
func NewProcessAllMembershipsRequest() *ProcessAllMembershipsRequest
```
NewProcessAllMembershipsRequest instantiates a new
ProcessAllMembershipsRequest object This constructor will assign default
values to properties that have it defined, and makes sure properties
required by API are set, but the set of arguments will change when the set
of required properties is changed


```go
func NewProcessAllMembershipsRequestWithDefaults() *ProcessAllMembershipsRequest
```
NewProcessAllMembershipsRequestWithDefaults instantiates a new
ProcessAllMembershipsRequest object This constructor will only assign
default values to properties that have it defined, but it doesn't guarantee
that properties required by API are set



### Methods

```go
func (o *ProcessAllMembershipsRequest) GetDatasource() string
```
GetDatasource returns the Datasource field value if set, zero value
otherwise.


```go
func (o *ProcessAllMembershipsRequest) GetDatasourceOk() (*string, bool)
```
GetDatasourceOk returns a tuple with the Datasource field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *ProcessAllMembershipsRequest) HasDatasource() bool
```
HasDatasource returns a boolean if a field has been set.


```go
func (o ProcessAllMembershipsRequest) MarshalJSON() ([]byte, error)
```


```go
func (o *ProcessAllMembershipsRequest) SetDatasource(v string)
```
SetDatasource gets a reference to the given string and assigns it to the
Datasource field.




### Type PropertyDefinition
```go
type PropertyDefinition struct {
	// The name of the property in the `DocumentMetadata` (e.g. 'createTime', 'updateTime', 'author', 'container'). In the future, this will support custom properties too.
	Name *string `json:"name,omitempty"`
	// The user friendly label for the property.
	DisplayLabel *string `json:"displayLabel,omitempty"`
	// The user friendly label for the property that will be used if a plural context.
	DisplayLabelPlural *string `json:"displayLabelPlural,omitempty"`
	// The type of custom property - this governs the search and faceting behavior
	PropertyType *string `json:"propertyType,omitempty"`
	UiOptions    *string `json:"uiOptions,omitempty"`
	// If true then the property will not show up as a facet in the UI.
	HideUiFacet *bool `json:"hideUiFacet,omitempty"`
	// Will be used to set the order of facets in the UI, if present. If set for one facet, must be set for all non-hidden UI facets. Must take on an integer value from 1 (shown at the top) to N (shown last), where N is the number of non-hidden UI facets. These facets will be ordered below the built-in \"Type\" and \"Tag\" operators.
	UiFacetOrder          *int32                  `json:"uiFacetOrder,omitempty"`
	ObjectPropertyOptions []ObjectPropertyOptions `json:"objectPropertyOptions,omitempty"`
	// The unique identifier of the `PropertyGroup` to which this property belongs.
	Group *string `json:"group,omitempty"`
}
```
PropertyDefinition struct for PropertyDefinition

### Functions

```go
func NewPropertyDefinition() *PropertyDefinition
```
NewPropertyDefinition instantiates a new PropertyDefinition object This
constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed


```go
func NewPropertyDefinitionWithDefaults() *PropertyDefinition
```
NewPropertyDefinitionWithDefaults instantiates a new PropertyDefinition
object This constructor will only assign default values to properties that
have it defined, but it doesn't guarantee that properties required by API
are set



### Methods

```go
func (o *PropertyDefinition) GetDisplayLabel() string
```
GetDisplayLabel returns the DisplayLabel field value if set, zero value
otherwise.


```go
func (o *PropertyDefinition) GetDisplayLabelOk() (*string, bool)
```
GetDisplayLabelOk returns a tuple with the DisplayLabel field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *PropertyDefinition) GetDisplayLabelPlural() string
```
GetDisplayLabelPlural returns the DisplayLabelPlural field value if set,
zero value otherwise.


```go
func (o *PropertyDefinition) GetDisplayLabelPluralOk() (*string, bool)
```
GetDisplayLabelPluralOk returns a tuple with the DisplayLabelPlural field
value if set, nil otherwise and a boolean to check if the value has been
set.


```go
func (o *PropertyDefinition) GetGroup() string
```
GetGroup returns the Group field value if set, zero value otherwise.


```go
func (o *PropertyDefinition) GetGroupOk() (*string, bool)
```
GetGroupOk returns a tuple with the Group field value if set, nil otherwise
and a boolean to check if the value has been set.


```go
func (o *PropertyDefinition) GetHideUiFacet() bool
```
GetHideUiFacet returns the HideUiFacet field value if set, zero value
otherwise.


```go
func (o *PropertyDefinition) GetHideUiFacetOk() (*bool, bool)
```
GetHideUiFacetOk returns a tuple with the HideUiFacet field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *PropertyDefinition) GetName() string
```
GetName returns the Name field value if set, zero value otherwise.


```go
func (o *PropertyDefinition) GetNameOk() (*string, bool)
```
GetNameOk returns a tuple with the Name field value if set, nil otherwise
and a boolean to check if the value has been set.


```go
func (o *PropertyDefinition) GetObjectPropertyOptions() []ObjectPropertyOptions
```
GetObjectPropertyOptions returns the ObjectPropertyOptions field value if
set, zero value otherwise.


```go
func (o *PropertyDefinition) GetObjectPropertyOptionsOk() ([]ObjectPropertyOptions, bool)
```
GetObjectPropertyOptionsOk returns a tuple with the ObjectPropertyOptions
field value if set, nil otherwise and a boolean to check if the value has
been set.


```go
func (o *PropertyDefinition) GetPropertyType() string
```
GetPropertyType returns the PropertyType field value if set, zero value
otherwise.


```go
func (o *PropertyDefinition) GetPropertyTypeOk() (*string, bool)
```
GetPropertyTypeOk returns a tuple with the PropertyType field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *PropertyDefinition) GetUiFacetOrder() int32
```
GetUiFacetOrder returns the UiFacetOrder field value if set, zero value
otherwise.


```go
func (o *PropertyDefinition) GetUiFacetOrderOk() (*int32, bool)
```
GetUiFacetOrderOk returns a tuple with the UiFacetOrder field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *PropertyDefinition) GetUiOptions() string
```
GetUiOptions returns the UiOptions field value if set, zero value otherwise.


```go
func (o *PropertyDefinition) GetUiOptionsOk() (*string, bool)
```
GetUiOptionsOk returns a tuple with the UiOptions field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *PropertyDefinition) HasDisplayLabel() bool
```
HasDisplayLabel returns a boolean if a field has been set.


```go
func (o *PropertyDefinition) HasDisplayLabelPlural() bool
```
HasDisplayLabelPlural returns a boolean if a field has been set.


```go
func (o *PropertyDefinition) HasGroup() bool
```
HasGroup returns a boolean if a field has been set.


```go
func (o *PropertyDefinition) HasHideUiFacet() bool
```
HasHideUiFacet returns a boolean if a field has been set.


```go
func (o *PropertyDefinition) HasName() bool
```
HasName returns a boolean if a field has been set.


```go
func (o *PropertyDefinition) HasObjectPropertyOptions() bool
```
HasObjectPropertyOptions returns a boolean if a field has been set.


```go
func (o *PropertyDefinition) HasPropertyType() bool
```
HasPropertyType returns a boolean if a field has been set.


```go
func (o *PropertyDefinition) HasUiFacetOrder() bool
```
HasUiFacetOrder returns a boolean if a field has been set.


```go
func (o *PropertyDefinition) HasUiOptions() bool
```
HasUiOptions returns a boolean if a field has been set.


```go
func (o PropertyDefinition) MarshalJSON() ([]byte, error)
```


```go
func (o *PropertyDefinition) SetDisplayLabel(v string)
```
SetDisplayLabel gets a reference to the given string and assigns it to the
DisplayLabel field.


```go
func (o *PropertyDefinition) SetDisplayLabelPlural(v string)
```
SetDisplayLabelPlural gets a reference to the given string and assigns it to
the DisplayLabelPlural field.


```go
func (o *PropertyDefinition) SetGroup(v string)
```
SetGroup gets a reference to the given string and assigns it to the Group
field.


```go
func (o *PropertyDefinition) SetHideUiFacet(v bool)
```
SetHideUiFacet gets a reference to the given bool and assigns it to the
HideUiFacet field.


```go
func (o *PropertyDefinition) SetName(v string)
```
SetName gets a reference to the given string and assigns it to the Name
field.


```go
func (o *PropertyDefinition) SetObjectPropertyOptions(v []ObjectPropertyOptions)
```
SetObjectPropertyOptions gets a reference to the given
[]ObjectPropertyOptions and assigns it to the ObjectPropertyOptions field.


```go
func (o *PropertyDefinition) SetPropertyType(v string)
```
SetPropertyType gets a reference to the given string and assigns it to the
PropertyType field.


```go
func (o *PropertyDefinition) SetUiFacetOrder(v int32)
```
SetUiFacetOrder gets a reference to the given int32 and assigns it to the
UiFacetOrder field.


```go
func (o *PropertyDefinition) SetUiOptions(v string)
```
SetUiOptions gets a reference to the given string and assigns it to the
UiOptions field.




### Type PropertyGroup
```go
type PropertyGroup struct {
	// The unique identifier of the group.
	Name *string `json:"name,omitempty"`
	// The user-friendly group label to display.
	DisplayLabel *string `json:"displayLabel,omitempty"`
}
```
PropertyGroup A grouping for multiple PropertyDefinition. Grouped properties
will be displayed together in the UI.

### Functions

```go
func NewPropertyGroup() *PropertyGroup
```
NewPropertyGroup instantiates a new PropertyGroup object This constructor
will assign default values to properties that have it defined, and makes
sure properties required by API are set, but the set of arguments will
change when the set of required properties is changed


```go
func NewPropertyGroupWithDefaults() *PropertyGroup
```
NewPropertyGroupWithDefaults instantiates a new PropertyGroup object This
constructor will only assign default values to properties that have it
defined, but it doesn't guarantee that properties required by API are set



### Methods

```go
func (o *PropertyGroup) GetDisplayLabel() string
```
GetDisplayLabel returns the DisplayLabel field value if set, zero value
otherwise.


```go
func (o *PropertyGroup) GetDisplayLabelOk() (*string, bool)
```
GetDisplayLabelOk returns a tuple with the DisplayLabel field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *PropertyGroup) GetName() string
```
GetName returns the Name field value if set, zero value otherwise.


```go
func (o *PropertyGroup) GetNameOk() (*string, bool)
```
GetNameOk returns a tuple with the Name field value if set, nil otherwise
and a boolean to check if the value has been set.


```go
func (o *PropertyGroup) HasDisplayLabel() bool
```
HasDisplayLabel returns a boolean if a field has been set.


```go
func (o *PropertyGroup) HasName() bool
```
HasName returns a boolean if a field has been set.


```go
func (o PropertyGroup) MarshalJSON() ([]byte, error)
```


```go
func (o *PropertyGroup) SetDisplayLabel(v string)
```
SetDisplayLabel gets a reference to the given string and assigns it to the
DisplayLabel field.


```go
func (o *PropertyGroup) SetName(v string)
```
SetName gets a reference to the given string and assigns it to the Name
field.




### Type Quicklink
```go
type Quicklink struct {
	// Full action name. Used in autocomplete
	Name *string `json:"name,omitempty"`
	// Shortened name. Used in app card
	ShortName *string `json:"shortName,omitempty"`
	// The URL for the action
	Url        *string     `json:"url,omitempty"`
	IconConfig *IconConfig `json:"iconConfig,omitempty"`
	// Unique identifier of this quicklink
	Id *string `json:"id,omitempty"`
	// The scopes for which this quicklink is applicable
	Scopes []string `json:"scopes,omitempty"`
}
```
Quicklink An action for a specific datasource that will show up in
autocomplete and app card, e.g. \"Create new issue\" for jira

### Functions

```go
func NewQuicklink() *Quicklink
```
NewQuicklink instantiates a new Quicklink object This constructor will
assign default values to properties that have it defined, and makes sure
properties required by API are set, but the set of arguments will change
when the set of required properties is changed


```go
func NewQuicklinkWithDefaults() *Quicklink
```
NewQuicklinkWithDefaults instantiates a new Quicklink object This
constructor will only assign default values to properties that have it
defined, but it doesn't guarantee that properties required by API are set



### Methods

```go
func (o *Quicklink) GetIconConfig() IconConfig
```
GetIconConfig returns the IconConfig field value if set, zero value
otherwise.


```go
func (o *Quicklink) GetIconConfigOk() (*IconConfig, bool)
```
GetIconConfigOk returns a tuple with the IconConfig field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *Quicklink) GetId() string
```
GetId returns the Id field value if set, zero value otherwise.


```go
func (o *Quicklink) GetIdOk() (*string, bool)
```
GetIdOk returns a tuple with the Id field value if set, nil otherwise and a
boolean to check if the value has been set.


```go
func (o *Quicklink) GetName() string
```
GetName returns the Name field value if set, zero value otherwise.


```go
func (o *Quicklink) GetNameOk() (*string, bool)
```
GetNameOk returns a tuple with the Name field value if set, nil otherwise
and a boolean to check if the value has been set.


```go
func (o *Quicklink) GetScopes() []string
```
GetScopes returns the Scopes field value if set, zero value otherwise.


```go
func (o *Quicklink) GetScopesOk() ([]string, bool)
```
GetScopesOk returns a tuple with the Scopes field value if set, nil
otherwise and a boolean to check if the value has been set.


```go
func (o *Quicklink) GetShortName() string
```
GetShortName returns the ShortName field value if set, zero value otherwise.


```go
func (o *Quicklink) GetShortNameOk() (*string, bool)
```
GetShortNameOk returns a tuple with the ShortName field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *Quicklink) GetUrl() string
```
GetUrl returns the Url field value if set, zero value otherwise.


```go
func (o *Quicklink) GetUrlOk() (*string, bool)
```
GetUrlOk returns a tuple with the Url field value if set, nil otherwise and
a boolean to check if the value has been set.


```go
func (o *Quicklink) HasIconConfig() bool
```
HasIconConfig returns a boolean if a field has been set.


```go
func (o *Quicklink) HasId() bool
```
HasId returns a boolean if a field has been set.


```go
func (o *Quicklink) HasName() bool
```
HasName returns a boolean if a field has been set.


```go
func (o *Quicklink) HasScopes() bool
```
HasScopes returns a boolean if a field has been set.


```go
func (o *Quicklink) HasShortName() bool
```
HasShortName returns a boolean if a field has been set.


```go
func (o *Quicklink) HasUrl() bool
```
HasUrl returns a boolean if a field has been set.


```go
func (o Quicklink) MarshalJSON() ([]byte, error)
```


```go
func (o *Quicklink) SetIconConfig(v IconConfig)
```
SetIconConfig gets a reference to the given IconConfig and assigns it to the
IconConfig field.


```go
func (o *Quicklink) SetId(v string)
```
SetId gets a reference to the given string and assigns it to the Id field.


```go
func (o *Quicklink) SetName(v string)
```
SetName gets a reference to the given string and assigns it to the Name
field.


```go
func (o *Quicklink) SetScopes(v []string)
```
SetScopes gets a reference to the given []string and assigns it to the
Scopes field.


```go
func (o *Quicklink) SetShortName(v string)
```
SetShortName gets a reference to the given string and assigns it to the
ShortName field.


```go
func (o *Quicklink) SetUrl(v string)
```
SetUrl gets a reference to the given string and assigns it to the Url field.




### Type ServerConfiguration
```go
type ServerConfiguration struct {
	URL         string
	Description string
	Variables   map[string]ServerVariable
}
```
ServerConfiguration stores the information about a server


### Type ServerConfigurations
```go
type ServerConfigurations []ServerConfiguration
```
ServerConfigurations stores multiple ServerConfiguration items

### Methods

```go
func (sc ServerConfigurations) URL(index int, variables map[string]string) (string, error)
```
URL formats template on a index using given variables




### Type ServerVariable
```go
type ServerVariable struct {
	Description  string
	DefaultValue string
	EnumValues   []string
}
```
ServerVariable stores the information about a server variable


### Type SharedDatasourceConfig
```go
type SharedDatasourceConfig struct {
	// Unique identifier of datasource instance to which this config applies.
	Name string `json:"name"`
	// Example text for what to search for in this datasource
	SuggestionText *string `json:"suggestionText,omitempty"`
	// The user-friendly instance label to display. If omitted, falls back to the title-cased `name`.
	DisplayName *string `json:"displayName,omitempty"`
	// The URL of the landing page for this datasource instance. Should point to the most useful page for users, not the company marketing page.
	HomeUrl *string `json:"homeUrl,omitempty"`
	// This only applies to WEB_CRAWL and BROWSER_CRAWL datasources. Defines the seed urls for crawling.
	CrawlerSeedUrls []string `json:"crawlerSeedUrls,omitempty"`
	// The URL to an image to be displayed as an icon for this datasource instance in dark mode. Must have a transparency mask. SVG are recommended over PNG. Public, scio-authenticated and Base64 encoded data URLs are all valid (but not third-party-authenticated URLs).
	IconDarkUrl *string `json:"iconDarkUrl,omitempty"`
	// The URL to an image to be displayed as an icon for this datasource instance. Must have a transparency mask. SVG are recommended over PNG. Public, scio-authenticated and Base64 encoded data URLs are all valid (but not third-party-authenticated URLs).
	IconUrl *string `json:"iconUrl,omitempty"`
	// The list of top-level `objectType`s for the datasource.
	ObjectDefinitions []ObjectDefinition `json:"objectDefinitions,omitempty"`
	// List of built-in facet types that should be hidden for the datasource.
	HideBuiltInFacets []string `json:"hideBuiltInFacets,omitempty"`
	// Regular expression that matches URLs of documents of the datasource instance. The behavior for multiple matches is non-deterministic. **Note: urlRegex is a required field for non-entity datasources (ie. datasources where isEntityDatasource is false). Please add a regex as specific as possible to this datasource instance.**
	UrlRegex *string `json:"urlRegex,omitempty"`
	// A list of regular expressions to apply to an arbitrary URL to transform it into a canonical URL for this datasource instance. Regexes are to be applied in the order specified in this list.
	CanonicalizingURLRegex []CanonicalizingRegexType `json:"canonicalizingURLRegex,omitempty"`
	// A list of regular expressions to apply to an arbitrary title to transform it into a title that will be displayed in the search results
	CanonicalizingTitleRegex []CanonicalizingRegexType `json:"canonicalizingTitleRegex,omitempty"`
	// A regex that identifies titles that should not be indexed
	RedlistTitleRegex *string        `json:"redlistTitleRegex,omitempty"`
	ConnectorType     *ConnectorType `json:"connectorType,omitempty"`
	// List of actions for this datasource instance that will show up in autocomplete and app card, e.g. \"Create new issue\" for jira
	Quicklinks []Quicklink `json:"quicklinks,omitempty"`
	// The name of a render config to use for displaying results from this datasource. Any well known datasource name may be used to render the same as that source, e.g. `web` or `gdrive`.
	RenderConfigPreset *string `json:"renderConfigPreset,omitempty"`
	// Aliases that can be used as `app` operator-values.
	Aliases []string `json:"aliases,omitempty"`
	// The type of this datasource. It is an important signal for relevance and must be specified and cannot be UNCATEGORIZED.
	DatasourceCategory *string `json:"datasourceCategory,omitempty"`
	// Whether or not this datasource is hosted on-premise.
	IsOnPrem *bool `json:"isOnPrem,omitempty"`
	// True if browser activity is able to report the correct URL for VIEW events. Set this to true if the URLs reported by Chrome are constant throughout each page load. Set this to false if the page has Javascript that modifies the URL during or after the load.
	TrustUrlRegexForViewActivity *bool `json:"trustUrlRegexForViewActivity,omitempty"`
	// If true, a utm_source query param will be added to outbound links to this datasource within Glean.
	IncludeUtmSource *bool `json:"includeUtmSource,omitempty"`
	// The (non-unique) name of the datasource of which this config is an instance, e.g. \"jira\".
	DatasourceName *string `json:"datasourceName,omitempty"`
	// The instance of the datasource for this particular config, e.g. \"onprem\".
	InstanceOnlyName *string `json:"instanceOnlyName,omitempty"`
	// A human readable string identifying this instance as compared to its peers, e.g. \"github.com/askscio\" or \"github.askscio.com\"
	InstanceDescription *string `json:"instanceDescription,omitempty"`
}
```
SharedDatasourceConfig Structure describing shared config properties of the
datasource (including multi-instance support)

### Functions

```go
func NewSharedDatasourceConfig(name string) *SharedDatasourceConfig
```
NewSharedDatasourceConfig instantiates a new SharedDatasourceConfig object
This constructor will assign default values to properties that have it
defined, and makes sure properties required by API are set, but the set of
arguments will change when the set of required properties is changed


```go
func NewSharedDatasourceConfigWithDefaults() *SharedDatasourceConfig
```
NewSharedDatasourceConfigWithDefaults instantiates a new
SharedDatasourceConfig object This constructor will only assign default
values to properties that have it defined, but it doesn't guarantee that
properties required by API are set



### Methods

```go
func (o *SharedDatasourceConfig) GetAliases() []string
```
GetAliases returns the Aliases field value if set, zero value otherwise.


```go
func (o *SharedDatasourceConfig) GetAliasesOk() ([]string, bool)
```
GetAliasesOk returns a tuple with the Aliases field value if set, nil
otherwise and a boolean to check if the value has been set.


```go
func (o *SharedDatasourceConfig) GetCanonicalizingTitleRegex() []CanonicalizingRegexType
```
GetCanonicalizingTitleRegex returns the CanonicalizingTitleRegex field value
if set, zero value otherwise.


```go
func (o *SharedDatasourceConfig) GetCanonicalizingTitleRegexOk() ([]CanonicalizingRegexType, bool)
```
GetCanonicalizingTitleRegexOk returns a tuple with the
CanonicalizingTitleRegex field value if set, nil otherwise and a boolean to
check if the value has been set.


```go
func (o *SharedDatasourceConfig) GetCanonicalizingURLRegex() []CanonicalizingRegexType
```
GetCanonicalizingURLRegex returns the CanonicalizingURLRegex field value if
set, zero value otherwise.


```go
func (o *SharedDatasourceConfig) GetCanonicalizingURLRegexOk() ([]CanonicalizingRegexType, bool)
```
GetCanonicalizingURLRegexOk returns a tuple with the CanonicalizingURLRegex
field value if set, nil otherwise and a boolean to check if the value has
been set.


```go
func (o *SharedDatasourceConfig) GetConnectorType() ConnectorType
```
GetConnectorType returns the ConnectorType field value if set, zero value
otherwise.


```go
func (o *SharedDatasourceConfig) GetConnectorTypeOk() (*ConnectorType, bool)
```
GetConnectorTypeOk returns a tuple with the ConnectorType field value if
set, nil otherwise and a boolean to check if the value has been set.


```go
func (o *SharedDatasourceConfig) GetCrawlerSeedUrls() []string
```
GetCrawlerSeedUrls returns the CrawlerSeedUrls field value if set, zero
value otherwise.


```go
func (o *SharedDatasourceConfig) GetCrawlerSeedUrlsOk() ([]string, bool)
```
GetCrawlerSeedUrlsOk returns a tuple with the CrawlerSeedUrls field value if
set, nil otherwise and a boolean to check if the value has been set.


```go
func (o *SharedDatasourceConfig) GetDatasourceCategory() string
```
GetDatasourceCategory returns the DatasourceCategory field value if set,
zero value otherwise.


```go
func (o *SharedDatasourceConfig) GetDatasourceCategoryOk() (*string, bool)
```
GetDatasourceCategoryOk returns a tuple with the DatasourceCategory field
value if set, nil otherwise and a boolean to check if the value has been
set.


```go
func (o *SharedDatasourceConfig) GetDatasourceName() string
```
GetDatasourceName returns the DatasourceName field value if set, zero value
otherwise.


```go
func (o *SharedDatasourceConfig) GetDatasourceNameOk() (*string, bool)
```
GetDatasourceNameOk returns a tuple with the DatasourceName field value if
set, nil otherwise and a boolean to check if the value has been set.


```go
func (o *SharedDatasourceConfig) GetDisplayName() string
```
GetDisplayName returns the DisplayName field value if set, zero value
otherwise.


```go
func (o *SharedDatasourceConfig) GetDisplayNameOk() (*string, bool)
```
GetDisplayNameOk returns a tuple with the DisplayName field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *SharedDatasourceConfig) GetHideBuiltInFacets() []string
```
GetHideBuiltInFacets returns the HideBuiltInFacets field value if set,
zero value otherwise.


```go
func (o *SharedDatasourceConfig) GetHideBuiltInFacetsOk() ([]string, bool)
```
GetHideBuiltInFacetsOk returns a tuple with the HideBuiltInFacets field
value if set, nil otherwise and a boolean to check if the value has been
set.


```go
func (o *SharedDatasourceConfig) GetHomeUrl() string
```
GetHomeUrl returns the HomeUrl field value if set, zero value otherwise.


```go
func (o *SharedDatasourceConfig) GetHomeUrlOk() (*string, bool)
```
GetHomeUrlOk returns a tuple with the HomeUrl field value if set, nil
otherwise and a boolean to check if the value has been set.


```go
func (o *SharedDatasourceConfig) GetIconDarkUrl() string
```
GetIconDarkUrl returns the IconDarkUrl field value if set, zero value
otherwise.


```go
func (o *SharedDatasourceConfig) GetIconDarkUrlOk() (*string, bool)
```
GetIconDarkUrlOk returns a tuple with the IconDarkUrl field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *SharedDatasourceConfig) GetIconUrl() string
```
GetIconUrl returns the IconUrl field value if set, zero value otherwise.


```go
func (o *SharedDatasourceConfig) GetIconUrlOk() (*string, bool)
```
GetIconUrlOk returns a tuple with the IconUrl field value if set, nil
otherwise and a boolean to check if the value has been set.


```go
func (o *SharedDatasourceConfig) GetIncludeUtmSource() bool
```
GetIncludeUtmSource returns the IncludeUtmSource field value if set,
zero value otherwise.


```go
func (o *SharedDatasourceConfig) GetIncludeUtmSourceOk() (*bool, bool)
```
GetIncludeUtmSourceOk returns a tuple with the IncludeUtmSource field value
if set, nil otherwise and a boolean to check if the value has been set.


```go
func (o *SharedDatasourceConfig) GetInstanceDescription() string
```
GetInstanceDescription returns the InstanceDescription field value if set,
zero value otherwise.


```go
func (o *SharedDatasourceConfig) GetInstanceDescriptionOk() (*string, bool)
```
GetInstanceDescriptionOk returns a tuple with the InstanceDescription field
value if set, nil otherwise and a boolean to check if the value has been
set.


```go
func (o *SharedDatasourceConfig) GetInstanceOnlyName() string
```
GetInstanceOnlyName returns the InstanceOnlyName field value if set,
zero value otherwise.


```go
func (o *SharedDatasourceConfig) GetInstanceOnlyNameOk() (*string, bool)
```
GetInstanceOnlyNameOk returns a tuple with the InstanceOnlyName field value
if set, nil otherwise and a boolean to check if the value has been set.


```go
func (o *SharedDatasourceConfig) GetIsOnPrem() bool
```
GetIsOnPrem returns the IsOnPrem field value if set, zero value otherwise.


```go
func (o *SharedDatasourceConfig) GetIsOnPremOk() (*bool, bool)
```
GetIsOnPremOk returns a tuple with the IsOnPrem field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *SharedDatasourceConfig) GetName() string
```
GetName returns the Name field value


```go
func (o *SharedDatasourceConfig) GetNameOk() (*string, bool)
```
GetNameOk returns a tuple with the Name field value and a boolean to check
if the value has been set.


```go
func (o *SharedDatasourceConfig) GetObjectDefinitions() []ObjectDefinition
```
GetObjectDefinitions returns the ObjectDefinitions field value if set,
zero value otherwise.


```go
func (o *SharedDatasourceConfig) GetObjectDefinitionsOk() ([]ObjectDefinition, bool)
```
GetObjectDefinitionsOk returns a tuple with the ObjectDefinitions field
value if set, nil otherwise and a boolean to check if the value has been
set.


```go
func (o *SharedDatasourceConfig) GetQuicklinks() []Quicklink
```
GetQuicklinks returns the Quicklinks field value if set, zero value
otherwise.


```go
func (o *SharedDatasourceConfig) GetQuicklinksOk() ([]Quicklink, bool)
```
GetQuicklinksOk returns a tuple with the Quicklinks field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *SharedDatasourceConfig) GetRedlistTitleRegex() string
```
GetRedlistTitleRegex returns the RedlistTitleRegex field value if set,
zero value otherwise.


```go
func (o *SharedDatasourceConfig) GetRedlistTitleRegexOk() (*string, bool)
```
GetRedlistTitleRegexOk returns a tuple with the RedlistTitleRegex field
value if set, nil otherwise and a boolean to check if the value has been
set.


```go
func (o *SharedDatasourceConfig) GetRenderConfigPreset() string
```
GetRenderConfigPreset returns the RenderConfigPreset field value if set,
zero value otherwise.


```go
func (o *SharedDatasourceConfig) GetRenderConfigPresetOk() (*string, bool)
```
GetRenderConfigPresetOk returns a tuple with the RenderConfigPreset field
value if set, nil otherwise and a boolean to check if the value has been
set.


```go
func (o *SharedDatasourceConfig) GetSuggestionText() string
```
GetSuggestionText returns the SuggestionText field value if set, zero value
otherwise.


```go
func (o *SharedDatasourceConfig) GetSuggestionTextOk() (*string, bool)
```
GetSuggestionTextOk returns a tuple with the SuggestionText field value if
set, nil otherwise and a boolean to check if the value has been set.


```go
func (o *SharedDatasourceConfig) GetTrustUrlRegexForViewActivity() bool
```
GetTrustUrlRegexForViewActivity returns the TrustUrlRegexForViewActivity
field value if set, zero value otherwise.


```go
func (o *SharedDatasourceConfig) GetTrustUrlRegexForViewActivityOk() (*bool, bool)
```
GetTrustUrlRegexForViewActivityOk returns a tuple with the
TrustUrlRegexForViewActivity field value if set, nil otherwise and a boolean
to check if the value has been set.


```go
func (o *SharedDatasourceConfig) GetUrlRegex() string
```
GetUrlRegex returns the UrlRegex field value if set, zero value otherwise.


```go
func (o *SharedDatasourceConfig) GetUrlRegexOk() (*string, bool)
```
GetUrlRegexOk returns a tuple with the UrlRegex field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *SharedDatasourceConfig) HasAliases() bool
```
HasAliases returns a boolean if a field has been set.


```go
func (o *SharedDatasourceConfig) HasCanonicalizingTitleRegex() bool
```
HasCanonicalizingTitleRegex returns a boolean if a field has been set.


```go
func (o *SharedDatasourceConfig) HasCanonicalizingURLRegex() bool
```
HasCanonicalizingURLRegex returns a boolean if a field has been set.


```go
func (o *SharedDatasourceConfig) HasConnectorType() bool
```
HasConnectorType returns a boolean if a field has been set.


```go
func (o *SharedDatasourceConfig) HasCrawlerSeedUrls() bool
```
HasCrawlerSeedUrls returns a boolean if a field has been set.


```go
func (o *SharedDatasourceConfig) HasDatasourceCategory() bool
```
HasDatasourceCategory returns a boolean if a field has been set.


```go
func (o *SharedDatasourceConfig) HasDatasourceName() bool
```
HasDatasourceName returns a boolean if a field has been set.


```go
func (o *SharedDatasourceConfig) HasDisplayName() bool
```
HasDisplayName returns a boolean if a field has been set.


```go
func (o *SharedDatasourceConfig) HasHideBuiltInFacets() bool
```
HasHideBuiltInFacets returns a boolean if a field has been set.


```go
func (o *SharedDatasourceConfig) HasHomeUrl() bool
```
HasHomeUrl returns a boolean if a field has been set.


```go
func (o *SharedDatasourceConfig) HasIconDarkUrl() bool
```
HasIconDarkUrl returns a boolean if a field has been set.


```go
func (o *SharedDatasourceConfig) HasIconUrl() bool
```
HasIconUrl returns a boolean if a field has been set.


```go
func (o *SharedDatasourceConfig) HasIncludeUtmSource() bool
```
HasIncludeUtmSource returns a boolean if a field has been set.


```go
func (o *SharedDatasourceConfig) HasInstanceDescription() bool
```
HasInstanceDescription returns a boolean if a field has been set.


```go
func (o *SharedDatasourceConfig) HasInstanceOnlyName() bool
```
HasInstanceOnlyName returns a boolean if a field has been set.


```go
func (o *SharedDatasourceConfig) HasIsOnPrem() bool
```
HasIsOnPrem returns a boolean if a field has been set.


```go
func (o *SharedDatasourceConfig) HasObjectDefinitions() bool
```
HasObjectDefinitions returns a boolean if a field has been set.


```go
func (o *SharedDatasourceConfig) HasQuicklinks() bool
```
HasQuicklinks returns a boolean if a field has been set.


```go
func (o *SharedDatasourceConfig) HasRedlistTitleRegex() bool
```
HasRedlistTitleRegex returns a boolean if a field has been set.


```go
func (o *SharedDatasourceConfig) HasRenderConfigPreset() bool
```
HasRenderConfigPreset returns a boolean if a field has been set.


```go
func (o *SharedDatasourceConfig) HasSuggestionText() bool
```
HasSuggestionText returns a boolean if a field has been set.


```go
func (o *SharedDatasourceConfig) HasTrustUrlRegexForViewActivity() bool
```
HasTrustUrlRegexForViewActivity returns a boolean if a field has been set.


```go
func (o *SharedDatasourceConfig) HasUrlRegex() bool
```
HasUrlRegex returns a boolean if a field has been set.


```go
func (o SharedDatasourceConfig) MarshalJSON() ([]byte, error)
```


```go
func (o *SharedDatasourceConfig) SetAliases(v []string)
```
SetAliases gets a reference to the given []string and assigns it to the
Aliases field.


```go
func (o *SharedDatasourceConfig) SetCanonicalizingTitleRegex(v []CanonicalizingRegexType)
```
SetCanonicalizingTitleRegex gets a reference to the given
[]CanonicalizingRegexType and assigns it to the CanonicalizingTitleRegex
field.


```go
func (o *SharedDatasourceConfig) SetCanonicalizingURLRegex(v []CanonicalizingRegexType)
```
SetCanonicalizingURLRegex gets a reference to the given
[]CanonicalizingRegexType and assigns it to the CanonicalizingURLRegex
field.


```go
func (o *SharedDatasourceConfig) SetConnectorType(v ConnectorType)
```
SetConnectorType gets a reference to the given ConnectorType and assigns it
to the ConnectorType field.


```go
func (o *SharedDatasourceConfig) SetCrawlerSeedUrls(v []string)
```
SetCrawlerSeedUrls gets a reference to the given []string and assigns it to
the CrawlerSeedUrls field.


```go
func (o *SharedDatasourceConfig) SetDatasourceCategory(v string)
```
SetDatasourceCategory gets a reference to the given string and assigns it to
the DatasourceCategory field.


```go
func (o *SharedDatasourceConfig) SetDatasourceName(v string)
```
SetDatasourceName gets a reference to the given string and assigns it to the
DatasourceName field.


```go
func (o *SharedDatasourceConfig) SetDisplayName(v string)
```
SetDisplayName gets a reference to the given string and assigns it to the
DisplayName field.


```go
func (o *SharedDatasourceConfig) SetHideBuiltInFacets(v []string)
```
SetHideBuiltInFacets gets a reference to the given []string and assigns it
to the HideBuiltInFacets field.


```go
func (o *SharedDatasourceConfig) SetHomeUrl(v string)
```
SetHomeUrl gets a reference to the given string and assigns it to the
HomeUrl field.


```go
func (o *SharedDatasourceConfig) SetIconDarkUrl(v string)
```
SetIconDarkUrl gets a reference to the given string and assigns it to the
IconDarkUrl field.


```go
func (o *SharedDatasourceConfig) SetIconUrl(v string)
```
SetIconUrl gets a reference to the given string and assigns it to the
IconUrl field.


```go
func (o *SharedDatasourceConfig) SetIncludeUtmSource(v bool)
```
SetIncludeUtmSource gets a reference to the given bool and assigns it to the
IncludeUtmSource field.


```go
func (o *SharedDatasourceConfig) SetInstanceDescription(v string)
```
SetInstanceDescription gets a reference to the given string and assigns it
to the InstanceDescription field.


```go
func (o *SharedDatasourceConfig) SetInstanceOnlyName(v string)
```
SetInstanceOnlyName gets a reference to the given string and assigns it to
the InstanceOnlyName field.


```go
func (o *SharedDatasourceConfig) SetIsOnPrem(v bool)
```
SetIsOnPrem gets a reference to the given bool and assigns it to the
IsOnPrem field.


```go
func (o *SharedDatasourceConfig) SetName(v string)
```
SetName sets field value


```go
func (o *SharedDatasourceConfig) SetObjectDefinitions(v []ObjectDefinition)
```
SetObjectDefinitions gets a reference to the given []ObjectDefinition and
assigns it to the ObjectDefinitions field.


```go
func (o *SharedDatasourceConfig) SetQuicklinks(v []Quicklink)
```
SetQuicklinks gets a reference to the given []Quicklink and assigns it to
the Quicklinks field.


```go
func (o *SharedDatasourceConfig) SetRedlistTitleRegex(v string)
```
SetRedlistTitleRegex gets a reference to the given string and assigns it to
the RedlistTitleRegex field.


```go
func (o *SharedDatasourceConfig) SetRenderConfigPreset(v string)
```
SetRenderConfigPreset gets a reference to the given string and assigns it to
the RenderConfigPreset field.


```go
func (o *SharedDatasourceConfig) SetSuggestionText(v string)
```
SetSuggestionText gets a reference to the given string and assigns it to the
SuggestionText field.


```go
func (o *SharedDatasourceConfig) SetTrustUrlRegexForViewActivity(v bool)
```
SetTrustUrlRegexForViewActivity gets a reference to the given bool and
assigns it to the TrustUrlRegexForViewActivity field.


```go
func (o *SharedDatasourceConfig) SetUrlRegex(v string)
```
SetUrlRegex gets a reference to the given string and assigns it to the
UrlRegex field.




### Type SharedDatasourceConfigAllOf
```go
type SharedDatasourceConfigAllOf struct {
	// The (non-unique) name of the datasource of which this config is an instance, e.g. \"jira\".
	DatasourceName *string `json:"datasourceName,omitempty"`
	// The instance of the datasource for this particular config, e.g. \"onprem\".
	InstanceOnlyName *string `json:"instanceOnlyName,omitempty"`
	// A human readable string identifying this instance as compared to its peers, e.g. \"github.com/askscio\" or \"github.askscio.com\"
	InstanceDescription *string `json:"instanceDescription,omitempty"`
}
```
SharedDatasourceConfigAllOf struct for SharedDatasourceConfigAllOf

### Functions

```go
func NewSharedDatasourceConfigAllOf() *SharedDatasourceConfigAllOf
```
NewSharedDatasourceConfigAllOf instantiates a new
SharedDatasourceConfigAllOf object This constructor will assign default
values to properties that have it defined, and makes sure properties
required by API are set, but the set of arguments will change when the set
of required properties is changed


```go
func NewSharedDatasourceConfigAllOfWithDefaults() *SharedDatasourceConfigAllOf
```
NewSharedDatasourceConfigAllOfWithDefaults instantiates a new
SharedDatasourceConfigAllOf object This constructor will only assign default
values to properties that have it defined, but it doesn't guarantee that
properties required by API are set



### Methods

```go
func (o *SharedDatasourceConfigAllOf) GetDatasourceName() string
```
GetDatasourceName returns the DatasourceName field value if set, zero value
otherwise.


```go
func (o *SharedDatasourceConfigAllOf) GetDatasourceNameOk() (*string, bool)
```
GetDatasourceNameOk returns a tuple with the DatasourceName field value if
set, nil otherwise and a boolean to check if the value has been set.


```go
func (o *SharedDatasourceConfigAllOf) GetInstanceDescription() string
```
GetInstanceDescription returns the InstanceDescription field value if set,
zero value otherwise.


```go
func (o *SharedDatasourceConfigAllOf) GetInstanceDescriptionOk() (*string, bool)
```
GetInstanceDescriptionOk returns a tuple with the InstanceDescription field
value if set, nil otherwise and a boolean to check if the value has been
set.


```go
func (o *SharedDatasourceConfigAllOf) GetInstanceOnlyName() string
```
GetInstanceOnlyName returns the InstanceOnlyName field value if set,
zero value otherwise.


```go
func (o *SharedDatasourceConfigAllOf) GetInstanceOnlyNameOk() (*string, bool)
```
GetInstanceOnlyNameOk returns a tuple with the InstanceOnlyName field value
if set, nil otherwise and a boolean to check if the value has been set.


```go
func (o *SharedDatasourceConfigAllOf) HasDatasourceName() bool
```
HasDatasourceName returns a boolean if a field has been set.


```go
func (o *SharedDatasourceConfigAllOf) HasInstanceDescription() bool
```
HasInstanceDescription returns a boolean if a field has been set.


```go
func (o *SharedDatasourceConfigAllOf) HasInstanceOnlyName() bool
```
HasInstanceOnlyName returns a boolean if a field has been set.


```go
func (o SharedDatasourceConfigAllOf) MarshalJSON() ([]byte, error)
```


```go
func (o *SharedDatasourceConfigAllOf) SetDatasourceName(v string)
```
SetDatasourceName gets a reference to the given string and assigns it to the
DatasourceName field.


```go
func (o *SharedDatasourceConfigAllOf) SetInstanceDescription(v string)
```
SetInstanceDescription gets a reference to the given string and assigns it
to the InstanceDescription field.


```go
func (o *SharedDatasourceConfigAllOf) SetInstanceOnlyName(v string)
```
SetInstanceOnlyName gets a reference to the given string and assigns it to
the InstanceOnlyName field.




### Type SharedDatasourceConfigNoInstance
```go
type SharedDatasourceConfigNoInstance struct {
	// Unique identifier of datasource instance to which this config applies.
	Name string `json:"name"`
	// Example text for what to search for in this datasource
	SuggestionText *string `json:"suggestionText,omitempty"`
	// The user-friendly instance label to display. If omitted, falls back to the title-cased `name`.
	DisplayName *string `json:"displayName,omitempty"`
	// The URL of the landing page for this datasource instance. Should point to the most useful page for users, not the company marketing page.
	HomeUrl *string `json:"homeUrl,omitempty"`
	// This only applies to WEB_CRAWL and BROWSER_CRAWL datasources. Defines the seed urls for crawling.
	CrawlerSeedUrls []string `json:"crawlerSeedUrls,omitempty"`
	// The URL to an image to be displayed as an icon for this datasource instance in dark mode. Must have a transparency mask. SVG are recommended over PNG. Public, scio-authenticated and Base64 encoded data URLs are all valid (but not third-party-authenticated URLs).
	IconDarkUrl *string `json:"iconDarkUrl,omitempty"`
	// The URL to an image to be displayed as an icon for this datasource instance. Must have a transparency mask. SVG are recommended over PNG. Public, scio-authenticated and Base64 encoded data URLs are all valid (but not third-party-authenticated URLs).
	IconUrl *string `json:"iconUrl,omitempty"`
	// The list of top-level `objectType`s for the datasource.
	ObjectDefinitions []ObjectDefinition `json:"objectDefinitions,omitempty"`
	// List of built-in facet types that should be hidden for the datasource.
	HideBuiltInFacets []string `json:"hideBuiltInFacets,omitempty"`
	// Regular expression that matches URLs of documents of the datasource instance. The behavior for multiple matches is non-deterministic. **Note: urlRegex is a required field for non-entity datasources (ie. datasources where isEntityDatasource is false). Please add a regex as specific as possible to this datasource instance.**
	UrlRegex *string `json:"urlRegex,omitempty"`
	// A list of regular expressions to apply to an arbitrary URL to transform it into a canonical URL for this datasource instance. Regexes are to be applied in the order specified in this list.
	CanonicalizingURLRegex []CanonicalizingRegexType `json:"canonicalizingURLRegex,omitempty"`
	// A list of regular expressions to apply to an arbitrary title to transform it into a title that will be displayed in the search results
	CanonicalizingTitleRegex []CanonicalizingRegexType `json:"canonicalizingTitleRegex,omitempty"`
	// A regex that identifies titles that should not be indexed
	RedlistTitleRegex *string        `json:"redlistTitleRegex,omitempty"`
	ConnectorType     *ConnectorType `json:"connectorType,omitempty"`
	// List of actions for this datasource instance that will show up in autocomplete and app card, e.g. \"Create new issue\" for jira
	Quicklinks []Quicklink `json:"quicklinks,omitempty"`
	// The name of a render config to use for displaying results from this datasource. Any well known datasource name may be used to render the same as that source, e.g. `web` or `gdrive`.
	RenderConfigPreset *string `json:"renderConfigPreset,omitempty"`
	// Aliases that can be used as `app` operator-values.
	Aliases []string `json:"aliases,omitempty"`
	// The type of this datasource. It is an important signal for relevance and must be specified and cannot be UNCATEGORIZED.
	DatasourceCategory *string `json:"datasourceCategory,omitempty"`
	// Whether or not this datasource is hosted on-premise.
	IsOnPrem *bool `json:"isOnPrem,omitempty"`
	// True if browser activity is able to report the correct URL for VIEW events. Set this to true if the URLs reported by Chrome are constant throughout each page load. Set this to false if the page has Javascript that modifies the URL during or after the load.
	TrustUrlRegexForViewActivity *bool `json:"trustUrlRegexForViewActivity,omitempty"`
	// If true, a utm_source query param will be added to outbound links to this datasource within Glean.
	IncludeUtmSource *bool `json:"includeUtmSource,omitempty"`
}
```
SharedDatasourceConfigNoInstance Structure describing shared config
properties of a datasource with no multi-instance support.

### Functions

```go
func NewSharedDatasourceConfigNoInstance(name string) *SharedDatasourceConfigNoInstance
```
NewSharedDatasourceConfigNoInstance instantiates a new
SharedDatasourceConfigNoInstance object This constructor will assign
default values to properties that have it defined, and makes sure properties
required by API are set, but the set of arguments will change when the set
of required properties is changed


```go
func NewSharedDatasourceConfigNoInstanceWithDefaults() *SharedDatasourceConfigNoInstance
```
NewSharedDatasourceConfigNoInstanceWithDefaults instantiates a new
SharedDatasourceConfigNoInstance object This constructor will only assign
default values to properties that have it defined, but it doesn't guarantee
that properties required by API are set



### Methods

```go
func (o *SharedDatasourceConfigNoInstance) GetAliases() []string
```
GetAliases returns the Aliases field value if set, zero value otherwise.


```go
func (o *SharedDatasourceConfigNoInstance) GetAliasesOk() ([]string, bool)
```
GetAliasesOk returns a tuple with the Aliases field value if set, nil
otherwise and a boolean to check if the value has been set.


```go
func (o *SharedDatasourceConfigNoInstance) GetCanonicalizingTitleRegex() []CanonicalizingRegexType
```
GetCanonicalizingTitleRegex returns the CanonicalizingTitleRegex field value
if set, zero value otherwise.


```go
func (o *SharedDatasourceConfigNoInstance) GetCanonicalizingTitleRegexOk() ([]CanonicalizingRegexType, bool)
```
GetCanonicalizingTitleRegexOk returns a tuple with the
CanonicalizingTitleRegex field value if set, nil otherwise and a boolean to
check if the value has been set.


```go
func (o *SharedDatasourceConfigNoInstance) GetCanonicalizingURLRegex() []CanonicalizingRegexType
```
GetCanonicalizingURLRegex returns the CanonicalizingURLRegex field value if
set, zero value otherwise.


```go
func (o *SharedDatasourceConfigNoInstance) GetCanonicalizingURLRegexOk() ([]CanonicalizingRegexType, bool)
```
GetCanonicalizingURLRegexOk returns a tuple with the CanonicalizingURLRegex
field value if set, nil otherwise and a boolean to check if the value has
been set.


```go
func (o *SharedDatasourceConfigNoInstance) GetConnectorType() ConnectorType
```
GetConnectorType returns the ConnectorType field value if set, zero value
otherwise.


```go
func (o *SharedDatasourceConfigNoInstance) GetConnectorTypeOk() (*ConnectorType, bool)
```
GetConnectorTypeOk returns a tuple with the ConnectorType field value if
set, nil otherwise and a boolean to check if the value has been set.


```go
func (o *SharedDatasourceConfigNoInstance) GetCrawlerSeedUrls() []string
```
GetCrawlerSeedUrls returns the CrawlerSeedUrls field value if set, zero
value otherwise.


```go
func (o *SharedDatasourceConfigNoInstance) GetCrawlerSeedUrlsOk() ([]string, bool)
```
GetCrawlerSeedUrlsOk returns a tuple with the CrawlerSeedUrls field value if
set, nil otherwise and a boolean to check if the value has been set.


```go
func (o *SharedDatasourceConfigNoInstance) GetDatasourceCategory() string
```
GetDatasourceCategory returns the DatasourceCategory field value if set,
zero value otherwise.


```go
func (o *SharedDatasourceConfigNoInstance) GetDatasourceCategoryOk() (*string, bool)
```
GetDatasourceCategoryOk returns a tuple with the DatasourceCategory field
value if set, nil otherwise and a boolean to check if the value has been
set.


```go
func (o *SharedDatasourceConfigNoInstance) GetDisplayName() string
```
GetDisplayName returns the DisplayName field value if set, zero value
otherwise.


```go
func (o *SharedDatasourceConfigNoInstance) GetDisplayNameOk() (*string, bool)
```
GetDisplayNameOk returns a tuple with the DisplayName field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *SharedDatasourceConfigNoInstance) GetHideBuiltInFacets() []string
```
GetHideBuiltInFacets returns the HideBuiltInFacets field value if set,
zero value otherwise.


```go
func (o *SharedDatasourceConfigNoInstance) GetHideBuiltInFacetsOk() ([]string, bool)
```
GetHideBuiltInFacetsOk returns a tuple with the HideBuiltInFacets field
value if set, nil otherwise and a boolean to check if the value has been
set.


```go
func (o *SharedDatasourceConfigNoInstance) GetHomeUrl() string
```
GetHomeUrl returns the HomeUrl field value if set, zero value otherwise.


```go
func (o *SharedDatasourceConfigNoInstance) GetHomeUrlOk() (*string, bool)
```
GetHomeUrlOk returns a tuple with the HomeUrl field value if set, nil
otherwise and a boolean to check if the value has been set.


```go
func (o *SharedDatasourceConfigNoInstance) GetIconDarkUrl() string
```
GetIconDarkUrl returns the IconDarkUrl field value if set, zero value
otherwise.


```go
func (o *SharedDatasourceConfigNoInstance) GetIconDarkUrlOk() (*string, bool)
```
GetIconDarkUrlOk returns a tuple with the IconDarkUrl field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *SharedDatasourceConfigNoInstance) GetIconUrl() string
```
GetIconUrl returns the IconUrl field value if set, zero value otherwise.


```go
func (o *SharedDatasourceConfigNoInstance) GetIconUrlOk() (*string, bool)
```
GetIconUrlOk returns a tuple with the IconUrl field value if set, nil
otherwise and a boolean to check if the value has been set.


```go
func (o *SharedDatasourceConfigNoInstance) GetIncludeUtmSource() bool
```
GetIncludeUtmSource returns the IncludeUtmSource field value if set,
zero value otherwise.


```go
func (o *SharedDatasourceConfigNoInstance) GetIncludeUtmSourceOk() (*bool, bool)
```
GetIncludeUtmSourceOk returns a tuple with the IncludeUtmSource field value
if set, nil otherwise and a boolean to check if the value has been set.


```go
func (o *SharedDatasourceConfigNoInstance) GetIsOnPrem() bool
```
GetIsOnPrem returns the IsOnPrem field value if set, zero value otherwise.


```go
func (o *SharedDatasourceConfigNoInstance) GetIsOnPremOk() (*bool, bool)
```
GetIsOnPremOk returns a tuple with the IsOnPrem field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *SharedDatasourceConfigNoInstance) GetName() string
```
GetName returns the Name field value


```go
func (o *SharedDatasourceConfigNoInstance) GetNameOk() (*string, bool)
```
GetNameOk returns a tuple with the Name field value and a boolean to check
if the value has been set.


```go
func (o *SharedDatasourceConfigNoInstance) GetObjectDefinitions() []ObjectDefinition
```
GetObjectDefinitions returns the ObjectDefinitions field value if set,
zero value otherwise.


```go
func (o *SharedDatasourceConfigNoInstance) GetObjectDefinitionsOk() ([]ObjectDefinition, bool)
```
GetObjectDefinitionsOk returns a tuple with the ObjectDefinitions field
value if set, nil otherwise and a boolean to check if the value has been
set.


```go
func (o *SharedDatasourceConfigNoInstance) GetQuicklinks() []Quicklink
```
GetQuicklinks returns the Quicklinks field value if set, zero value
otherwise.


```go
func (o *SharedDatasourceConfigNoInstance) GetQuicklinksOk() ([]Quicklink, bool)
```
GetQuicklinksOk returns a tuple with the Quicklinks field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *SharedDatasourceConfigNoInstance) GetRedlistTitleRegex() string
```
GetRedlistTitleRegex returns the RedlistTitleRegex field value if set,
zero value otherwise.


```go
func (o *SharedDatasourceConfigNoInstance) GetRedlistTitleRegexOk() (*string, bool)
```
GetRedlistTitleRegexOk returns a tuple with the RedlistTitleRegex field
value if set, nil otherwise and a boolean to check if the value has been
set.


```go
func (o *SharedDatasourceConfigNoInstance) GetRenderConfigPreset() string
```
GetRenderConfigPreset returns the RenderConfigPreset field value if set,
zero value otherwise.


```go
func (o *SharedDatasourceConfigNoInstance) GetRenderConfigPresetOk() (*string, bool)
```
GetRenderConfigPresetOk returns a tuple with the RenderConfigPreset field
value if set, nil otherwise and a boolean to check if the value has been
set.


```go
func (o *SharedDatasourceConfigNoInstance) GetSuggestionText() string
```
GetSuggestionText returns the SuggestionText field value if set, zero value
otherwise.


```go
func (o *SharedDatasourceConfigNoInstance) GetSuggestionTextOk() (*string, bool)
```
GetSuggestionTextOk returns a tuple with the SuggestionText field value if
set, nil otherwise and a boolean to check if the value has been set.


```go
func (o *SharedDatasourceConfigNoInstance) GetTrustUrlRegexForViewActivity() bool
```
GetTrustUrlRegexForViewActivity returns the TrustUrlRegexForViewActivity
field value if set, zero value otherwise.


```go
func (o *SharedDatasourceConfigNoInstance) GetTrustUrlRegexForViewActivityOk() (*bool, bool)
```
GetTrustUrlRegexForViewActivityOk returns a tuple with the
TrustUrlRegexForViewActivity field value if set, nil otherwise and a boolean
to check if the value has been set.


```go
func (o *SharedDatasourceConfigNoInstance) GetUrlRegex() string
```
GetUrlRegex returns the UrlRegex field value if set, zero value otherwise.


```go
func (o *SharedDatasourceConfigNoInstance) GetUrlRegexOk() (*string, bool)
```
GetUrlRegexOk returns a tuple with the UrlRegex field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *SharedDatasourceConfigNoInstance) HasAliases() bool
```
HasAliases returns a boolean if a field has been set.


```go
func (o *SharedDatasourceConfigNoInstance) HasCanonicalizingTitleRegex() bool
```
HasCanonicalizingTitleRegex returns a boolean if a field has been set.


```go
func (o *SharedDatasourceConfigNoInstance) HasCanonicalizingURLRegex() bool
```
HasCanonicalizingURLRegex returns a boolean if a field has been set.


```go
func (o *SharedDatasourceConfigNoInstance) HasConnectorType() bool
```
HasConnectorType returns a boolean if a field has been set.


```go
func (o *SharedDatasourceConfigNoInstance) HasCrawlerSeedUrls() bool
```
HasCrawlerSeedUrls returns a boolean if a field has been set.


```go
func (o *SharedDatasourceConfigNoInstance) HasDatasourceCategory() bool
```
HasDatasourceCategory returns a boolean if a field has been set.


```go
func (o *SharedDatasourceConfigNoInstance) HasDisplayName() bool
```
HasDisplayName returns a boolean if a field has been set.


```go
func (o *SharedDatasourceConfigNoInstance) HasHideBuiltInFacets() bool
```
HasHideBuiltInFacets returns a boolean if a field has been set.


```go
func (o *SharedDatasourceConfigNoInstance) HasHomeUrl() bool
```
HasHomeUrl returns a boolean if a field has been set.


```go
func (o *SharedDatasourceConfigNoInstance) HasIconDarkUrl() bool
```
HasIconDarkUrl returns a boolean if a field has been set.


```go
func (o *SharedDatasourceConfigNoInstance) HasIconUrl() bool
```
HasIconUrl returns a boolean if a field has been set.


```go
func (o *SharedDatasourceConfigNoInstance) HasIncludeUtmSource() bool
```
HasIncludeUtmSource returns a boolean if a field has been set.


```go
func (o *SharedDatasourceConfigNoInstance) HasIsOnPrem() bool
```
HasIsOnPrem returns a boolean if a field has been set.


```go
func (o *SharedDatasourceConfigNoInstance) HasObjectDefinitions() bool
```
HasObjectDefinitions returns a boolean if a field has been set.


```go
func (o *SharedDatasourceConfigNoInstance) HasQuicklinks() bool
```
HasQuicklinks returns a boolean if a field has been set.


```go
func (o *SharedDatasourceConfigNoInstance) HasRedlistTitleRegex() bool
```
HasRedlistTitleRegex returns a boolean if a field has been set.


```go
func (o *SharedDatasourceConfigNoInstance) HasRenderConfigPreset() bool
```
HasRenderConfigPreset returns a boolean if a field has been set.


```go
func (o *SharedDatasourceConfigNoInstance) HasSuggestionText() bool
```
HasSuggestionText returns a boolean if a field has been set.


```go
func (o *SharedDatasourceConfigNoInstance) HasTrustUrlRegexForViewActivity() bool
```
HasTrustUrlRegexForViewActivity returns a boolean if a field has been set.


```go
func (o *SharedDatasourceConfigNoInstance) HasUrlRegex() bool
```
HasUrlRegex returns a boolean if a field has been set.


```go
func (o SharedDatasourceConfigNoInstance) MarshalJSON() ([]byte, error)
```


```go
func (o *SharedDatasourceConfigNoInstance) SetAliases(v []string)
```
SetAliases gets a reference to the given []string and assigns it to the
Aliases field.


```go
func (o *SharedDatasourceConfigNoInstance) SetCanonicalizingTitleRegex(v []CanonicalizingRegexType)
```
SetCanonicalizingTitleRegex gets a reference to the given
[]CanonicalizingRegexType and assigns it to the CanonicalizingTitleRegex
field.


```go
func (o *SharedDatasourceConfigNoInstance) SetCanonicalizingURLRegex(v []CanonicalizingRegexType)
```
SetCanonicalizingURLRegex gets a reference to the given
[]CanonicalizingRegexType and assigns it to the CanonicalizingURLRegex
field.


```go
func (o *SharedDatasourceConfigNoInstance) SetConnectorType(v ConnectorType)
```
SetConnectorType gets a reference to the given ConnectorType and assigns it
to the ConnectorType field.


```go
func (o *SharedDatasourceConfigNoInstance) SetCrawlerSeedUrls(v []string)
```
SetCrawlerSeedUrls gets a reference to the given []string and assigns it to
the CrawlerSeedUrls field.


```go
func (o *SharedDatasourceConfigNoInstance) SetDatasourceCategory(v string)
```
SetDatasourceCategory gets a reference to the given string and assigns it to
the DatasourceCategory field.


```go
func (o *SharedDatasourceConfigNoInstance) SetDisplayName(v string)
```
SetDisplayName gets a reference to the given string and assigns it to the
DisplayName field.


```go
func (o *SharedDatasourceConfigNoInstance) SetHideBuiltInFacets(v []string)
```
SetHideBuiltInFacets gets a reference to the given []string and assigns it
to the HideBuiltInFacets field.


```go
func (o *SharedDatasourceConfigNoInstance) SetHomeUrl(v string)
```
SetHomeUrl gets a reference to the given string and assigns it to the
HomeUrl field.


```go
func (o *SharedDatasourceConfigNoInstance) SetIconDarkUrl(v string)
```
SetIconDarkUrl gets a reference to the given string and assigns it to the
IconDarkUrl field.


```go
func (o *SharedDatasourceConfigNoInstance) SetIconUrl(v string)
```
SetIconUrl gets a reference to the given string and assigns it to the
IconUrl field.


```go
func (o *SharedDatasourceConfigNoInstance) SetIncludeUtmSource(v bool)
```
SetIncludeUtmSource gets a reference to the given bool and assigns it to the
IncludeUtmSource field.


```go
func (o *SharedDatasourceConfigNoInstance) SetIsOnPrem(v bool)
```
SetIsOnPrem gets a reference to the given bool and assigns it to the
IsOnPrem field.


```go
func (o *SharedDatasourceConfigNoInstance) SetName(v string)
```
SetName sets field value


```go
func (o *SharedDatasourceConfigNoInstance) SetObjectDefinitions(v []ObjectDefinition)
```
SetObjectDefinitions gets a reference to the given []ObjectDefinition and
assigns it to the ObjectDefinitions field.


```go
func (o *SharedDatasourceConfigNoInstance) SetQuicklinks(v []Quicklink)
```
SetQuicklinks gets a reference to the given []Quicklink and assigns it to
the Quicklinks field.


```go
func (o *SharedDatasourceConfigNoInstance) SetRedlistTitleRegex(v string)
```
SetRedlistTitleRegex gets a reference to the given string and assigns it to
the RedlistTitleRegex field.


```go
func (o *SharedDatasourceConfigNoInstance) SetRenderConfigPreset(v string)
```
SetRenderConfigPreset gets a reference to the given string and assigns it to
the RenderConfigPreset field.


```go
func (o *SharedDatasourceConfigNoInstance) SetSuggestionText(v string)
```
SetSuggestionText gets a reference to the given string and assigns it to the
SuggestionText field.


```go
func (o *SharedDatasourceConfigNoInstance) SetTrustUrlRegexForViewActivity(v bool)
```
SetTrustUrlRegexForViewActivity gets a reference to the given bool and
assigns it to the TrustUrlRegexForViewActivity field.


```go
func (o *SharedDatasourceConfigNoInstance) SetUrlRegex(v string)
```
SetUrlRegex gets a reference to the given string and assigns it to the
UrlRegex field.




### Type SocialNetworkDefinition
```go
type SocialNetworkDefinition struct {
	// Possible values are \"twitter\", \"linkedin\".
	Name *string `json:"name,omitempty"`
	// Human-readable profile name.
	ProfileName *string `json:"profileName,omitempty"`
	// Link to profile.
	ProfileUrl *string `json:"profileUrl,omitempty"`
}
```
SocialNetworkDefinition Employee's social network profile

### Functions

```go
func NewSocialNetworkDefinition() *SocialNetworkDefinition
```
NewSocialNetworkDefinition instantiates a new SocialNetworkDefinition object
This constructor will assign default values to properties that have it
defined, and makes sure properties required by API are set, but the set of
arguments will change when the set of required properties is changed


```go
func NewSocialNetworkDefinitionWithDefaults() *SocialNetworkDefinition
```
NewSocialNetworkDefinitionWithDefaults instantiates a new
SocialNetworkDefinition object This constructor will only assign default
values to properties that have it defined, but it doesn't guarantee that
properties required by API are set



### Methods

```go
func (o *SocialNetworkDefinition) GetName() string
```
GetName returns the Name field value if set, zero value otherwise.


```go
func (o *SocialNetworkDefinition) GetNameOk() (*string, bool)
```
GetNameOk returns a tuple with the Name field value if set, nil otherwise
and a boolean to check if the value has been set.


```go
func (o *SocialNetworkDefinition) GetProfileName() string
```
GetProfileName returns the ProfileName field value if set, zero value
otherwise.


```go
func (o *SocialNetworkDefinition) GetProfileNameOk() (*string, bool)
```
GetProfileNameOk returns a tuple with the ProfileName field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *SocialNetworkDefinition) GetProfileUrl() string
```
GetProfileUrl returns the ProfileUrl field value if set, zero value
otherwise.


```go
func (o *SocialNetworkDefinition) GetProfileUrlOk() (*string, bool)
```
GetProfileUrlOk returns a tuple with the ProfileUrl field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *SocialNetworkDefinition) HasName() bool
```
HasName returns a boolean if a field has been set.


```go
func (o *SocialNetworkDefinition) HasProfileName() bool
```
HasProfileName returns a boolean if a field has been set.


```go
func (o *SocialNetworkDefinition) HasProfileUrl() bool
```
HasProfileUrl returns a boolean if a field has been set.


```go
func (o SocialNetworkDefinition) MarshalJSON() ([]byte, error)
```


```go
func (o *SocialNetworkDefinition) SetName(v string)
```
SetName gets a reference to the given string and assigns it to the Name
field.


```go
func (o *SocialNetworkDefinition) SetProfileName(v string)
```
SetProfileName gets a reference to the given string and assigns it to the
ProfileName field.


```go
func (o *SocialNetworkDefinition) SetProfileUrl(v string)
```
SetProfileUrl gets a reference to the given string and assigns it to the
ProfileUrl field.




### Type StructuredLocation
```go
type StructuredLocation struct {
	// Desk number.
	DeskLocation *string `json:"deskLocation,omitempty"`
	// Location's timezone, e.g. UTC, PST.
	Timezone *string `json:"timezone,omitempty"`
	// Office address or name.
	Address *string `json:"address,omitempty"`
	// Name of the city.
	City *string `json:"city,omitempty"`
	// State code.
	State *string `json:"state,omitempty"`
	// Region information, e.g. NORAM, APAC.
	Region *string `json:"region,omitempty"`
	// ZIP Code for the address.
	ZipCode *string `json:"zipCode,omitempty"`
	// Country name.
	Country *string `json:"country,omitempty"`
	// Alpha-2 or Alpha-3 ISO 3166 country code, e.g. US or USA.
	CountryCode *string `json:"countryCode,omitempty"`
}
```
StructuredLocation Detailed location with information about country, state,
city etc.

### Functions

```go
func NewStructuredLocation() *StructuredLocation
```
NewStructuredLocation instantiates a new StructuredLocation object This
constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed


```go
func NewStructuredLocationWithDefaults() *StructuredLocation
```
NewStructuredLocationWithDefaults instantiates a new StructuredLocation
object This constructor will only assign default values to properties that
have it defined, but it doesn't guarantee that properties required by API
are set



### Methods

```go
func (o *StructuredLocation) GetAddress() string
```
GetAddress returns the Address field value if set, zero value otherwise.


```go
func (o *StructuredLocation) GetAddressOk() (*string, bool)
```
GetAddressOk returns a tuple with the Address field value if set, nil
otherwise and a boolean to check if the value has been set.


```go
func (o *StructuredLocation) GetCity() string
```
GetCity returns the City field value if set, zero value otherwise.


```go
func (o *StructuredLocation) GetCityOk() (*string, bool)
```
GetCityOk returns a tuple with the City field value if set, nil otherwise
and a boolean to check if the value has been set.


```go
func (o *StructuredLocation) GetCountry() string
```
GetCountry returns the Country field value if set, zero value otherwise.


```go
func (o *StructuredLocation) GetCountryCode() string
```
GetCountryCode returns the CountryCode field value if set, zero value
otherwise.


```go
func (o *StructuredLocation) GetCountryCodeOk() (*string, bool)
```
GetCountryCodeOk returns a tuple with the CountryCode field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *StructuredLocation) GetCountryOk() (*string, bool)
```
GetCountryOk returns a tuple with the Country field value if set, nil
otherwise and a boolean to check if the value has been set.


```go
func (o *StructuredLocation) GetDeskLocation() string
```
GetDeskLocation returns the DeskLocation field value if set, zero value
otherwise.


```go
func (o *StructuredLocation) GetDeskLocationOk() (*string, bool)
```
GetDeskLocationOk returns a tuple with the DeskLocation field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *StructuredLocation) GetRegion() string
```
GetRegion returns the Region field value if set, zero value otherwise.


```go
func (o *StructuredLocation) GetRegionOk() (*string, bool)
```
GetRegionOk returns a tuple with the Region field value if set, nil
otherwise and a boolean to check if the value has been set.


```go
func (o *StructuredLocation) GetState() string
```
GetState returns the State field value if set, zero value otherwise.


```go
func (o *StructuredLocation) GetStateOk() (*string, bool)
```
GetStateOk returns a tuple with the State field value if set, nil otherwise
and a boolean to check if the value has been set.


```go
func (o *StructuredLocation) GetTimezone() string
```
GetTimezone returns the Timezone field value if set, zero value otherwise.


```go
func (o *StructuredLocation) GetTimezoneOk() (*string, bool)
```
GetTimezoneOk returns a tuple with the Timezone field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *StructuredLocation) GetZipCode() string
```
GetZipCode returns the ZipCode field value if set, zero value otherwise.


```go
func (o *StructuredLocation) GetZipCodeOk() (*string, bool)
```
GetZipCodeOk returns a tuple with the ZipCode field value if set, nil
otherwise and a boolean to check if the value has been set.


```go
func (o *StructuredLocation) HasAddress() bool
```
HasAddress returns a boolean if a field has been set.


```go
func (o *StructuredLocation) HasCity() bool
```
HasCity returns a boolean if a field has been set.


```go
func (o *StructuredLocation) HasCountry() bool
```
HasCountry returns a boolean if a field has been set.


```go
func (o *StructuredLocation) HasCountryCode() bool
```
HasCountryCode returns a boolean if a field has been set.


```go
func (o *StructuredLocation) HasDeskLocation() bool
```
HasDeskLocation returns a boolean if a field has been set.


```go
func (o *StructuredLocation) HasRegion() bool
```
HasRegion returns a boolean if a field has been set.


```go
func (o *StructuredLocation) HasState() bool
```
HasState returns a boolean if a field has been set.


```go
func (o *StructuredLocation) HasTimezone() bool
```
HasTimezone returns a boolean if a field has been set.


```go
func (o *StructuredLocation) HasZipCode() bool
```
HasZipCode returns a boolean if a field has been set.


```go
func (o StructuredLocation) MarshalJSON() ([]byte, error)
```


```go
func (o *StructuredLocation) SetAddress(v string)
```
SetAddress gets a reference to the given string and assigns it to the
Address field.


```go
func (o *StructuredLocation) SetCity(v string)
```
SetCity gets a reference to the given string and assigns it to the City
field.


```go
func (o *StructuredLocation) SetCountry(v string)
```
SetCountry gets a reference to the given string and assigns it to the
Country field.


```go
func (o *StructuredLocation) SetCountryCode(v string)
```
SetCountryCode gets a reference to the given string and assigns it to the
CountryCode field.


```go
func (o *StructuredLocation) SetDeskLocation(v string)
```
SetDeskLocation gets a reference to the given string and assigns it to the
DeskLocation field.


```go
func (o *StructuredLocation) SetRegion(v string)
```
SetRegion gets a reference to the given string and assigns it to the Region
field.


```go
func (o *StructuredLocation) SetState(v string)
```
SetState gets a reference to the given string and assigns it to the State
field.


```go
func (o *StructuredLocation) SetTimezone(v string)
```
SetTimezone gets a reference to the given string and assigns it to the
Timezone field.


```go
func (o *StructuredLocation) SetZipCode(v string)
```
SetZipCode gets a reference to the given string and assigns it to the
ZipCode field.




### Type TeamEmail
```go
type TeamEmail struct {
	// An email address
	Email string `json:"email"`
	// An enum of `PRIMARY`, `SECONDARY`, `ONCALL`, `OTHER`
	Type string `json:"type"`
}
```
TeamEmail Information about a team's email

### Functions

```go
func NewTeamEmail(email string, type_ string) *TeamEmail
```
NewTeamEmail instantiates a new TeamEmail object This constructor will
assign default values to properties that have it defined, and makes sure
properties required by API are set, but the set of arguments will change
when the set of required properties is changed


```go
func NewTeamEmailWithDefaults() *TeamEmail
```
NewTeamEmailWithDefaults instantiates a new TeamEmail object This
constructor will only assign default values to properties that have it
defined, but it doesn't guarantee that properties required by API are set



### Methods

```go
func (o *TeamEmail) GetEmail() string
```
GetEmail returns the Email field value


```go
func (o *TeamEmail) GetEmailOk() (*string, bool)
```
GetEmailOk returns a tuple with the Email field value and a boolean to check
if the value has been set.


```go
func (o *TeamEmail) GetType() string
```
GetType returns the Type field value


```go
func (o *TeamEmail) GetTypeOk() (*string, bool)
```
GetTypeOk returns a tuple with the Type field value and a boolean to check
if the value has been set.


```go
func (o TeamEmail) MarshalJSON() ([]byte, error)
```


```go
func (o *TeamEmail) SetEmail(v string)
```
SetEmail sets field value


```go
func (o *TeamEmail) SetType(v string)
```
SetType sets field value




### Type TeamInfoDefinition
```go
type TeamInfoDefinition struct {
	// The unique ID of the team
	Id string `json:"id"`
	// Human-readable team name
	Name string `json:"name"`
	// The description of this team
	Description *string `json:"description,omitempty"`
	// Typically the highest level organizational unit; generally applies to bigger companies with multiple distinct businesses.
	BusinessUnit *string `json:"businessUnit,omitempty"`
	// An organizational unit where everyone has a similar task, e.g. `Engineering`.
	Department *string `json:"department,omitempty"`
	// A link to the team's photo
	PhotoUrl *string `json:"photoUrl,omitempty"`
	// A link to an external team page. If set, team results will link to it.
	ExternalLink *string `json:"externalLink,omitempty"`
	// The emails of the team
	Emails []TeamEmail `json:"emails,omitempty"`
	// The datasource profiles of the team
	DatasourceProfiles []DatasourceProfile `json:"datasourceProfiles,omitempty"`
	// The members of the team
	Members []TeamMember `json:"members"`
}
```
TeamInfoDefinition Information about an employee's team

### Functions

```go
func NewTeamInfoDefinition(id string, name string, members []TeamMember) *TeamInfoDefinition
```
NewTeamInfoDefinition instantiates a new TeamInfoDefinition object This
constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed


```go
func NewTeamInfoDefinitionWithDefaults() *TeamInfoDefinition
```
NewTeamInfoDefinitionWithDefaults instantiates a new TeamInfoDefinition
object This constructor will only assign default values to properties that
have it defined, but it doesn't guarantee that properties required by API
are set



### Methods

```go
func (o *TeamInfoDefinition) GetBusinessUnit() string
```
GetBusinessUnit returns the BusinessUnit field value if set, zero value
otherwise.


```go
func (o *TeamInfoDefinition) GetBusinessUnitOk() (*string, bool)
```
GetBusinessUnitOk returns a tuple with the BusinessUnit field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *TeamInfoDefinition) GetDatasourceProfiles() []DatasourceProfile
```
GetDatasourceProfiles returns the DatasourceProfiles field value if set,
zero value otherwise.


```go
func (o *TeamInfoDefinition) GetDatasourceProfilesOk() ([]DatasourceProfile, bool)
```
GetDatasourceProfilesOk returns a tuple with the DatasourceProfiles field
value if set, nil otherwise and a boolean to check if the value has been
set.


```go
func (o *TeamInfoDefinition) GetDepartment() string
```
GetDepartment returns the Department field value if set, zero value
otherwise.


```go
func (o *TeamInfoDefinition) GetDepartmentOk() (*string, bool)
```
GetDepartmentOk returns a tuple with the Department field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *TeamInfoDefinition) GetDescription() string
```
GetDescription returns the Description field value if set, zero value
otherwise.


```go
func (o *TeamInfoDefinition) GetDescriptionOk() (*string, bool)
```
GetDescriptionOk returns a tuple with the Description field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *TeamInfoDefinition) GetEmails() []TeamEmail
```
GetEmails returns the Emails field value if set, zero value otherwise.


```go
func (o *TeamInfoDefinition) GetEmailsOk() ([]TeamEmail, bool)
```
GetEmailsOk returns a tuple with the Emails field value if set, nil
otherwise and a boolean to check if the value has been set.


```go
func (o *TeamInfoDefinition) GetExternalLink() string
```
GetExternalLink returns the ExternalLink field value if set, zero value
otherwise.


```go
func (o *TeamInfoDefinition) GetExternalLinkOk() (*string, bool)
```
GetExternalLinkOk returns a tuple with the ExternalLink field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *TeamInfoDefinition) GetId() string
```
GetId returns the Id field value


```go
func (o *TeamInfoDefinition) GetIdOk() (*string, bool)
```
GetIdOk returns a tuple with the Id field value and a boolean to check if
the value has been set.


```go
func (o *TeamInfoDefinition) GetMembers() []TeamMember
```
GetMembers returns the Members field value


```go
func (o *TeamInfoDefinition) GetMembersOk() ([]TeamMember, bool)
```
GetMembersOk returns a tuple with the Members field value and a boolean to
check if the value has been set.


```go
func (o *TeamInfoDefinition) GetName() string
```
GetName returns the Name field value


```go
func (o *TeamInfoDefinition) GetNameOk() (*string, bool)
```
GetNameOk returns a tuple with the Name field value and a boolean to check
if the value has been set.


```go
func (o *TeamInfoDefinition) GetPhotoUrl() string
```
GetPhotoUrl returns the PhotoUrl field value if set, zero value otherwise.


```go
func (o *TeamInfoDefinition) GetPhotoUrlOk() (*string, bool)
```
GetPhotoUrlOk returns a tuple with the PhotoUrl field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *TeamInfoDefinition) HasBusinessUnit() bool
```
HasBusinessUnit returns a boolean if a field has been set.


```go
func (o *TeamInfoDefinition) HasDatasourceProfiles() bool
```
HasDatasourceProfiles returns a boolean if a field has been set.


```go
func (o *TeamInfoDefinition) HasDepartment() bool
```
HasDepartment returns a boolean if a field has been set.


```go
func (o *TeamInfoDefinition) HasDescription() bool
```
HasDescription returns a boolean if a field has been set.


```go
func (o *TeamInfoDefinition) HasEmails() bool
```
HasEmails returns a boolean if a field has been set.


```go
func (o *TeamInfoDefinition) HasExternalLink() bool
```
HasExternalLink returns a boolean if a field has been set.


```go
func (o *TeamInfoDefinition) HasPhotoUrl() bool
```
HasPhotoUrl returns a boolean if a field has been set.


```go
func (o TeamInfoDefinition) MarshalJSON() ([]byte, error)
```


```go
func (o *TeamInfoDefinition) SetBusinessUnit(v string)
```
SetBusinessUnit gets a reference to the given string and assigns it to the
BusinessUnit field.


```go
func (o *TeamInfoDefinition) SetDatasourceProfiles(v []DatasourceProfile)
```
SetDatasourceProfiles gets a reference to the given []DatasourceProfile and
assigns it to the DatasourceProfiles field.


```go
func (o *TeamInfoDefinition) SetDepartment(v string)
```
SetDepartment gets a reference to the given string and assigns it to the
Department field.


```go
func (o *TeamInfoDefinition) SetDescription(v string)
```
SetDescription gets a reference to the given string and assigns it to the
Description field.


```go
func (o *TeamInfoDefinition) SetEmails(v []TeamEmail)
```
SetEmails gets a reference to the given []TeamEmail and assigns it to the
Emails field.


```go
func (o *TeamInfoDefinition) SetExternalLink(v string)
```
SetExternalLink gets a reference to the given string and assigns it to the
ExternalLink field.


```go
func (o *TeamInfoDefinition) SetId(v string)
```
SetId sets field value


```go
func (o *TeamInfoDefinition) SetMembers(v []TeamMember)
```
SetMembers sets field value


```go
func (o *TeamInfoDefinition) SetName(v string)
```
SetName sets field value


```go
func (o *TeamInfoDefinition) SetPhotoUrl(v string)
```
SetPhotoUrl gets a reference to the given string and assigns it to the
PhotoUrl field.




### Type TeamMember
```go
type TeamMember struct {
	// The member's email
	Email string `json:"email"`
	// The member's relationship to the team, an enum of `MEMBER`, `MANAGER`, `LEAD`, `POINT_OF_CONTACT`, `OTHER`
	Relationship *string `json:"relationship,omitempty"`
	// The member's start date
	JoinDate *string `json:"join_date,omitempty"`
}
```
TeamMember Information about a team's member

### Functions

```go
func NewTeamMember(email string) *TeamMember
```
NewTeamMember instantiates a new TeamMember object This constructor will
assign default values to properties that have it defined, and makes sure
properties required by API are set, but the set of arguments will change
when the set of required properties is changed


```go
func NewTeamMemberWithDefaults() *TeamMember
```
NewTeamMemberWithDefaults instantiates a new TeamMember object This
constructor will only assign default values to properties that have it
defined, but it doesn't guarantee that properties required by API are set



### Methods

```go
func (o *TeamMember) GetEmail() string
```
GetEmail returns the Email field value


```go
func (o *TeamMember) GetEmailOk() (*string, bool)
```
GetEmailOk returns a tuple with the Email field value and a boolean to check
if the value has been set.


```go
func (o *TeamMember) GetJoinDate() string
```
GetJoinDate returns the JoinDate field value if set, zero value otherwise.


```go
func (o *TeamMember) GetJoinDateOk() (*string, bool)
```
GetJoinDateOk returns a tuple with the JoinDate field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *TeamMember) GetRelationship() string
```
GetRelationship returns the Relationship field value if set, zero value
otherwise.


```go
func (o *TeamMember) GetRelationshipOk() (*string, bool)
```
GetRelationshipOk returns a tuple with the Relationship field value if set,
nil otherwise and a boolean to check if the value has been set.


```go
func (o *TeamMember) HasJoinDate() bool
```
HasJoinDate returns a boolean if a field has been set.


```go
func (o *TeamMember) HasRelationship() bool
```
HasRelationship returns a boolean if a field has been set.


```go
func (o TeamMember) MarshalJSON() ([]byte, error)
```


```go
func (o *TeamMember) SetEmail(v string)
```
SetEmail sets field value


```go
func (o *TeamMember) SetJoinDate(v string)
```
SetJoinDate gets a reference to the given string and assigns it to the
JoinDate field.


```go
func (o *TeamMember) SetRelationship(v string)
```
SetRelationship gets a reference to the given string and assigns it to the
Relationship field.




### Type UserReferenceDefinition
```go
type UserReferenceDefinition struct {
	Email *string `json:"email,omitempty"`
	// some datasources refer to the user by the datasource user id in the document
	DatasourceUserId *string `json:"datasourceUserId,omitempty"`
	Name             *string `json:"name,omitempty"`
}
```
UserReferenceDefinition Describes how a user is referenced in a document.
The user can be referenced by email or by a datasource specific id.

### Functions

```go
func NewUserReferenceDefinition() *UserReferenceDefinition
```
NewUserReferenceDefinition instantiates a new UserReferenceDefinition object
This constructor will assign default values to properties that have it
defined, and makes sure properties required by API are set, but the set of
arguments will change when the set of required properties is changed


```go
func NewUserReferenceDefinitionWithDefaults() *UserReferenceDefinition
```
NewUserReferenceDefinitionWithDefaults instantiates a new
UserReferenceDefinition object This constructor will only assign default
values to properties that have it defined, but it doesn't guarantee that
properties required by API are set



### Methods

```go
func (o *UserReferenceDefinition) GetDatasourceUserId() string
```
GetDatasourceUserId returns the DatasourceUserId field value if set,
zero value otherwise.


```go
func (o *UserReferenceDefinition) GetDatasourceUserIdOk() (*string, bool)
```
GetDatasourceUserIdOk returns a tuple with the DatasourceUserId field value
if set, nil otherwise and a boolean to check if the value has been set.


```go
func (o *UserReferenceDefinition) GetEmail() string
```
GetEmail returns the Email field value if set, zero value otherwise.


```go
func (o *UserReferenceDefinition) GetEmailOk() (*string, bool)
```
GetEmailOk returns a tuple with the Email field value if set, nil otherwise
and a boolean to check if the value has been set.


```go
func (o *UserReferenceDefinition) GetName() string
```
GetName returns the Name field value if set, zero value otherwise.


```go
func (o *UserReferenceDefinition) GetNameOk() (*string, bool)
```
GetNameOk returns a tuple with the Name field value if set, nil otherwise
and a boolean to check if the value has been set.


```go
func (o *UserReferenceDefinition) HasDatasourceUserId() bool
```
HasDatasourceUserId returns a boolean if a field has been set.


```go
func (o *UserReferenceDefinition) HasEmail() bool
```
HasEmail returns a boolean if a field has been set.


```go
func (o *UserReferenceDefinition) HasName() bool
```
HasName returns a boolean if a field has been set.


```go
func (o UserReferenceDefinition) MarshalJSON() ([]byte, error)
```


```go
func (o *UserReferenceDefinition) SetDatasourceUserId(v string)
```
SetDatasourceUserId gets a reference to the given string and assigns it to
the DatasourceUserId field.


```go
func (o *UserReferenceDefinition) SetEmail(v string)
```
SetEmail gets a reference to the given string and assigns it to the Email
field.


```go
func (o *UserReferenceDefinition) SetName(v string)
```
SetName gets a reference to the given string and assigns it to the Name
field.







