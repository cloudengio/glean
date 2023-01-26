# \PermissionsApi

All URIs are relative to *https://domain-be.glean.com/api/index/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BetausersPost**](PermissionsApi.md#BetausersPost) | **Post** /betausers | Beta users
[**BulkindexgroupsPost**](PermissionsApi.md#BulkindexgroupsPost) | **Post** /bulkindexgroups | Bulk index groups
[**BulkindexmembershipsPost**](PermissionsApi.md#BulkindexmembershipsPost) | **Post** /bulkindexmemberships | Bulk index memberships for a group
[**BulkindexusersPost**](PermissionsApi.md#BulkindexusersPost) | **Post** /bulkindexusers | Bulk index users
[**CheckdocumentaccessPost**](PermissionsApi.md#CheckdocumentaccessPost) | **Post** /checkdocumentaccess | Check document access
[**DeletegroupPost**](PermissionsApi.md#DeletegroupPost) | **Post** /deletegroup | Delete group
[**DeletemembershipPost**](PermissionsApi.md#DeletemembershipPost) | **Post** /deletemembership | Delete membership
[**DeleteuserPost**](PermissionsApi.md#DeleteuserPost) | **Post** /deleteuser | Delete user
[**GetusercountPost**](PermissionsApi.md#GetusercountPost) | **Post** /getusercount | Get user count
[**IndexgroupPost**](PermissionsApi.md#IndexgroupPost) | **Post** /indexgroup | Index group
[**IndexmembershipPost**](PermissionsApi.md#IndexmembershipPost) | **Post** /indexmembership | Index membership
[**IndexuserPost**](PermissionsApi.md#IndexuserPost) | **Post** /indexuser | Index user
[**ProcessallmembershipsPost**](PermissionsApi.md#ProcessallmembershipsPost) | **Post** /processallmemberships | Schedules the processing of group memberships



## BetausersPost

> BetausersPost(ctx).GreenlistUsersRequest(greenlistUsersRequest).Execute()

Beta users



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    greenlistUsersRequest := *openapiclient.NewGreenlistUsersRequest("Datasource_example", []string{"Emails_example"}) // GreenlistUsersRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PermissionsApi.BetausersPost(context.Background()).GreenlistUsersRequest(greenlistUsersRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionsApi.BetausersPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBetausersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **greenlistUsersRequest** | [**GreenlistUsersRequest**](GreenlistUsersRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkindexgroupsPost

> BulkindexgroupsPost(ctx).BulkIndexGroupsRequest(bulkIndexGroupsRequest).Execute()

Bulk index groups



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    bulkIndexGroupsRequest := *openapiclient.NewBulkIndexGroupsRequest("UploadId_example", "Datasource_example", []openapiclient.DatasourceGroupDefinition{*openapiclient.NewDatasourceGroupDefinition("Name_example")}) // BulkIndexGroupsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PermissionsApi.BulkindexgroupsPost(context.Background()).BulkIndexGroupsRequest(bulkIndexGroupsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionsApi.BulkindexgroupsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkindexgroupsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bulkIndexGroupsRequest** | [**BulkIndexGroupsRequest**](BulkIndexGroupsRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkindexmembershipsPost

> BulkindexmembershipsPost(ctx).BulkIndexMembershipsRequest(bulkIndexMembershipsRequest).Execute()

Bulk index memberships for a group



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    bulkIndexMembershipsRequest := *openapiclient.NewBulkIndexMembershipsRequest("UploadId_example", "Datasource_example", []openapiclient.DatasourceBulkMembershipDefinition{*openapiclient.NewDatasourceBulkMembershipDefinition()}) // BulkIndexMembershipsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PermissionsApi.BulkindexmembershipsPost(context.Background()).BulkIndexMembershipsRequest(bulkIndexMembershipsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionsApi.BulkindexmembershipsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkindexmembershipsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bulkIndexMembershipsRequest** | [**BulkIndexMembershipsRequest**](BulkIndexMembershipsRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkindexusersPost

> BulkindexusersPost(ctx).BulkIndexUsersRequest(bulkIndexUsersRequest).Execute()

Bulk index users



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    bulkIndexUsersRequest := *openapiclient.NewBulkIndexUsersRequest("UploadId_example", "Datasource_example", []openapiclient.DatasourceUserDefinition{*openapiclient.NewDatasourceUserDefinition("Name_example")}) // BulkIndexUsersRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PermissionsApi.BulkindexusersPost(context.Background()).BulkIndexUsersRequest(bulkIndexUsersRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionsApi.BulkindexusersPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkindexusersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bulkIndexUsersRequest** | [**BulkIndexUsersRequest**](BulkIndexUsersRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckdocumentaccessPost

> CheckDocumentAccessResponse CheckdocumentaccessPost(ctx).CheckDocumentAccessRequest(checkDocumentAccessRequest).Execute()

Check document access



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    checkDocumentAccessRequest := *openapiclient.NewCheckDocumentAccessRequest("Datasource_example", "ObjectType_example", "DocId_example", "UserEmail_example") // CheckDocumentAccessRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PermissionsApi.CheckdocumentaccessPost(context.Background()).CheckDocumentAccessRequest(checkDocumentAccessRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionsApi.CheckdocumentaccessPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckdocumentaccessPost`: CheckDocumentAccessResponse
    fmt.Fprintf(os.Stdout, "Response from `PermissionsApi.CheckdocumentaccessPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckdocumentaccessPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **checkDocumentAccessRequest** | [**CheckDocumentAccessRequest**](CheckDocumentAccessRequest.md) |  | 

### Return type

[**CheckDocumentAccessResponse**](CheckDocumentAccessResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletegroupPost

> DeletegroupPost(ctx).DeleteGroupRequest(deleteGroupRequest).Execute()

Delete group



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    deleteGroupRequest := *openapiclient.NewDeleteGroupRequest("Datasource_example", "GroupName_example") // DeleteGroupRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PermissionsApi.DeletegroupPost(context.Background()).DeleteGroupRequest(deleteGroupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionsApi.DeletegroupPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeletegroupPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteGroupRequest** | [**DeleteGroupRequest**](DeleteGroupRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletemembershipPost

> DeletemembershipPost(ctx).DeleteMembershipRequest(deleteMembershipRequest).Execute()

Delete membership



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    deleteMembershipRequest := *openapiclient.NewDeleteMembershipRequest("Datasource_example", *openapiclient.NewDatasourceMembershipDefinition("GroupName_example")) // DeleteMembershipRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PermissionsApi.DeletemembershipPost(context.Background()).DeleteMembershipRequest(deleteMembershipRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionsApi.DeletemembershipPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeletemembershipPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteMembershipRequest** | [**DeleteMembershipRequest**](DeleteMembershipRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteuserPost

> DeleteuserPost(ctx).DeleteUserRequest(deleteUserRequest).Execute()

Delete user



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    deleteUserRequest := *openapiclient.NewDeleteUserRequest("Datasource_example", "Email_example") // DeleteUserRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PermissionsApi.DeleteuserPost(context.Background()).DeleteUserRequest(deleteUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionsApi.DeleteuserPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteuserPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteUserRequest** | [**DeleteUserRequest**](DeleteUserRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetusercountPost

> GetUserCountResponse GetusercountPost(ctx).GetUserCountRequest(getUserCountRequest).Execute()

Get user count



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    getUserCountRequest := *openapiclient.NewGetUserCountRequest() // GetUserCountRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PermissionsApi.GetusercountPost(context.Background()).GetUserCountRequest(getUserCountRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionsApi.GetusercountPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetusercountPost`: GetUserCountResponse
    fmt.Fprintf(os.Stdout, "Response from `PermissionsApi.GetusercountPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetusercountPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getUserCountRequest** | [**GetUserCountRequest**](GetUserCountRequest.md) |  | 

### Return type

[**GetUserCountResponse**](GetUserCountResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IndexgroupPost

> IndexgroupPost(ctx).IndexGroupRequest(indexGroupRequest).Execute()

Index group



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    indexGroupRequest := *openapiclient.NewIndexGroupRequest("Datasource_example", *openapiclient.NewDatasourceGroupDefinition("Name_example")) // IndexGroupRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PermissionsApi.IndexgroupPost(context.Background()).IndexGroupRequest(indexGroupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionsApi.IndexgroupPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIndexgroupPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **indexGroupRequest** | [**IndexGroupRequest**](IndexGroupRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IndexmembershipPost

> IndexmembershipPost(ctx).IndexMembershipRequest(indexMembershipRequest).Execute()

Index membership



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    indexMembershipRequest := *openapiclient.NewIndexMembershipRequest("Datasource_example", *openapiclient.NewDatasourceMembershipDefinition("GroupName_example")) // IndexMembershipRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PermissionsApi.IndexmembershipPost(context.Background()).IndexMembershipRequest(indexMembershipRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionsApi.IndexmembershipPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIndexmembershipPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **indexMembershipRequest** | [**IndexMembershipRequest**](IndexMembershipRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IndexuserPost

> IndexuserPost(ctx).IndexUserRequest(indexUserRequest).Execute()

Index user



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    indexUserRequest := *openapiclient.NewIndexUserRequest("Datasource_example", *openapiclient.NewDatasourceUserDefinition("Name_example")) // IndexUserRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PermissionsApi.IndexuserPost(context.Background()).IndexUserRequest(indexUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionsApi.IndexuserPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIndexuserPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **indexUserRequest** | [**IndexUserRequest**](IndexUserRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProcessallmembershipsPost

> ProcessallmembershipsPost(ctx).ProcessAllMembershipsRequest(processAllMembershipsRequest).Execute()

Schedules the processing of group memberships



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    processAllMembershipsRequest := *openapiclient.NewProcessAllMembershipsRequest() // ProcessAllMembershipsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PermissionsApi.ProcessallmembershipsPost(context.Background()).ProcessAllMembershipsRequest(processAllMembershipsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionsApi.ProcessallmembershipsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProcessallmembershipsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processAllMembershipsRequest** | [**ProcessAllMembershipsRequest**](ProcessAllMembershipsRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

