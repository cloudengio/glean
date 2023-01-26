# \PeopleApi

All URIs are relative to *https://domain-be.glean.com/api/index/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkindexemployeesPost**](PeopleApi.md#BulkindexemployeesPost) | **Post** /bulkindexemployees | Bulk index employees
[**BulkindexteamsPost**](PeopleApi.md#BulkindexteamsPost) | **Post** /bulkindexteams | Bulk index teams
[**DeleteemployeePost**](PeopleApi.md#DeleteemployeePost) | **Post** /deleteemployee | Delete employee
[**IndexemployeePost**](PeopleApi.md#IndexemployeePost) | **Post** /indexemployee | Index employee
[**IndexemployeelistPost**](PeopleApi.md#IndexemployeelistPost) | **Post** /indexemployeelist | Bulk index employees
[**IndexteamPost**](PeopleApi.md#IndexteamPost) | **Post** /indexteam | Index team



## BulkindexemployeesPost

> BulkindexemployeesPost(ctx).BulkIndexEmployeesRequest(bulkIndexEmployeesRequest).Execute()

Bulk index employees



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
    bulkIndexEmployeesRequest := *openapiclient.NewBulkIndexEmployeesRequest("UploadId_example", []openapiclient.EmployeeInfoDefinition{*openapiclient.NewEmployeeInfoDefinition("Email_example", "Department_example")}) // BulkIndexEmployeesRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PeopleApi.BulkindexemployeesPost(context.Background()).BulkIndexEmployeesRequest(bulkIndexEmployeesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PeopleApi.BulkindexemployeesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkindexemployeesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bulkIndexEmployeesRequest** | [**BulkIndexEmployeesRequest**](BulkIndexEmployeesRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json; charset=UTF-8
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkindexteamsPost

> BulkindexteamsPost(ctx).BulkIndexTeamsRequest(bulkIndexTeamsRequest).Execute()

Bulk index teams



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
    bulkIndexTeamsRequest := *openapiclient.NewBulkIndexTeamsRequest("UploadId_example", []openapiclient.TeamInfoDefinition{*openapiclient.NewTeamInfoDefinition("Id_example", "Name_example", []openapiclient.TeamMember{*openapiclient.NewTeamMember("Email_example")})}) // BulkIndexTeamsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PeopleApi.BulkindexteamsPost(context.Background()).BulkIndexTeamsRequest(bulkIndexTeamsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PeopleApi.BulkindexteamsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkindexteamsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bulkIndexTeamsRequest** | [**BulkIndexTeamsRequest**](BulkIndexTeamsRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json; charset=UTF-8
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteemployeePost

> DeleteemployeePost(ctx).DeleteEmployeeRequest(deleteEmployeeRequest).Execute()

Delete employee



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
    deleteEmployeeRequest := *openapiclient.NewDeleteEmployeeRequest("EmployeeEmail_example") // DeleteEmployeeRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PeopleApi.DeleteemployeePost(context.Background()).DeleteEmployeeRequest(deleteEmployeeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PeopleApi.DeleteemployeePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteemployeePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteEmployeeRequest** | [**DeleteEmployeeRequest**](DeleteEmployeeRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json; charset=UTF-8
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IndexemployeePost

> IndexemployeePost(ctx).IndexEmployeeRequest(indexEmployeeRequest).Execute()

Index employee



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
    indexEmployeeRequest := *openapiclient.NewIndexEmployeeRequest(*openapiclient.NewEmployeeInfoDefinition("Email_example", "Department_example")) // IndexEmployeeRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PeopleApi.IndexemployeePost(context.Background()).IndexEmployeeRequest(indexEmployeeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PeopleApi.IndexemployeePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIndexemployeePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **indexEmployeeRequest** | [**IndexEmployeeRequest**](IndexEmployeeRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json; charset=UTF-8
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IndexemployeelistPost

> IndexemployeelistPost(ctx).IndexEmployeeListRequest(indexEmployeeListRequest).Execute()

Bulk index employees



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
    indexEmployeeListRequest := *openapiclient.NewIndexEmployeeListRequest() // IndexEmployeeListRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PeopleApi.IndexemployeelistPost(context.Background()).IndexEmployeeListRequest(indexEmployeeListRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PeopleApi.IndexemployeelistPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIndexemployeelistPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **indexEmployeeListRequest** | [**IndexEmployeeListRequest**](IndexEmployeeListRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json; charset=UTF-8
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IndexteamPost

> IndexteamPost(ctx).IndexTeamRequest(indexTeamRequest).Execute()

Index team



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
    indexTeamRequest := *openapiclient.NewIndexTeamRequest(*openapiclient.NewTeamInfoDefinition("Id_example", "Name_example", []openapiclient.TeamMember{*openapiclient.NewTeamMember("Email_example")})) // IndexTeamRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PeopleApi.IndexteamPost(context.Background()).IndexTeamRequest(indexTeamRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PeopleApi.IndexteamPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIndexteamPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **indexTeamRequest** | [**IndexTeamRequest**](IndexTeamRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json; charset=UTF-8
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

