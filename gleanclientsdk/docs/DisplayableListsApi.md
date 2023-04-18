# \DisplayableListsApi

All URIs are relative to *https://domain-be.glean.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Createdisplayablelists**](DisplayableListsApi.md#Createdisplayablelists) | **Post** /createdisplayablelists | Create new displayable lists
[**Deletedisplayablelists**](DisplayableListsApi.md#Deletedisplayablelists) | **Post** /deletedisplayablelists | Delete displayable lists
[**Getdisplayablelists**](DisplayableListsApi.md#Getdisplayablelists) | **Post** /getdisplayablelists | Get displayable lists
[**Updatedisplayablelists**](DisplayableListsApi.md#Updatedisplayablelists) | **Post** /updatedisplayablelists | Update displayable lists



## Createdisplayablelists

> CreateDisplayableListsResponse Createdisplayablelists(ctx).Payload(payload).ClientVersion(clientVersion).Domain(domain).Eids(eids).Execute()

Create new displayable lists



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    payload := *openapiclient.NewCreateDisplayableListsRequest([]openapiclient.DisplayableList{*openapiclient.NewDisplayableList()}) // CreateDisplayableListsRequest | Create new displayable lists
    clientVersion := "clientVersion_example" // string | The version of the client making the request. (optional)
    domain := "domain_example" // string | The domain of the top-level page in which the client is being run. For embedded search and NSR, it will be the domain of the embedding page. For sidebar, it will be the domain of the embedding page if it's a recognized workplace app. For NTP and app.glean.com requests, it will be empty. (optional)
    eids := []int64{int64(123)} // []int64 | List of experiment ids to force for incoming request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DisplayableListsApi.Createdisplayablelists(context.Background()).Payload(payload).ClientVersion(clientVersion).Domain(domain).Eids(eids).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DisplayableListsApi.Createdisplayablelists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Createdisplayablelists`: CreateDisplayableListsResponse
    fmt.Fprintf(os.Stdout, "Response from `DisplayableListsApi.Createdisplayablelists`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatedisplayablelistsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**CreateDisplayableListsRequest**](CreateDisplayableListsRequest.md) | Create new displayable lists | 
 **clientVersion** | **string** | The version of the client making the request. | 
 **domain** | **string** | The domain of the top-level page in which the client is being run. For embedded search and NSR, it will be the domain of the embedding page. For sidebar, it will be the domain of the embedding page if it&#39;s a recognized workplace app. For NTP and app.glean.com requests, it will be empty. | 
 **eids** | **[]int64** | List of experiment ids to force for incoming request. | 

### Return type

[**CreateDisplayableListsResponse**](CreateDisplayableListsResponse.md)

### Authorization

[actAsBearerToken](../README.md#actAsBearerToken), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Deletedisplayablelists

> Deletedisplayablelists(ctx).Payload(payload).ClientVersion(clientVersion).Domain(domain).Eids(eids).Execute()

Delete displayable lists



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    payload := *openapiclient.NewDeleteDisplayableListsRequest([]int32{int32(123)}) // DeleteDisplayableListsRequest | Updated version of the displayable list configs.
    clientVersion := "clientVersion_example" // string | The version of the client making the request. (optional)
    domain := "domain_example" // string | The domain of the top-level page in which the client is being run. For embedded search and NSR, it will be the domain of the embedding page. For sidebar, it will be the domain of the embedding page if it's a recognized workplace app. For NTP and app.glean.com requests, it will be empty. (optional)
    eids := []int64{int64(123)} // []int64 | List of experiment ids to force for incoming request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DisplayableListsApi.Deletedisplayablelists(context.Background()).Payload(payload).ClientVersion(clientVersion).Domain(domain).Eids(eids).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DisplayableListsApi.Deletedisplayablelists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeletedisplayablelistsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**DeleteDisplayableListsRequest**](DeleteDisplayableListsRequest.md) | Updated version of the displayable list configs. | 
 **clientVersion** | **string** | The version of the client making the request. | 
 **domain** | **string** | The domain of the top-level page in which the client is being run. For embedded search and NSR, it will be the domain of the embedding page. For sidebar, it will be the domain of the embedding page if it&#39;s a recognized workplace app. For NTP and app.glean.com requests, it will be empty. | 
 **eids** | **[]int64** | List of experiment ids to force for incoming request. | 

### Return type

 (empty response body)

### Authorization

[actAsBearerToken](../README.md#actAsBearerToken), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Getdisplayablelists

> GetDisplayableListsResponse Getdisplayablelists(ctx).Payload(payload).ClientVersion(clientVersion).Domain(domain).Eids(eids).Actas(actas).Execute()

Get displayable lists



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    payload := *openapiclient.NewGetDisplayableListsRequest([]int32{int32(123)}) // GetDisplayableListsRequest | 
    clientVersion := "clientVersion_example" // string | The version of the client making the request. (optional)
    domain := "domain_example" // string | The domain of the top-level page in which the client is being run. For embedded search and NSR, it will be the domain of the embedding page. For sidebar, it will be the domain of the embedding page if it's a recognized workplace app. For NTP and app.glean.com requests, it will be empty. (optional)
    eids := []int64{int64(123)} // []int64 | List of experiment ids to force for incoming request. (optional)
    actas := "actas_example" // string | Email of another user to act as for debugging purposes. Requires sufficient permissions. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DisplayableListsApi.Getdisplayablelists(context.Background()).Payload(payload).ClientVersion(clientVersion).Domain(domain).Eids(eids).Actas(actas).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DisplayableListsApi.Getdisplayablelists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Getdisplayablelists`: GetDisplayableListsResponse
    fmt.Fprintf(os.Stdout, "Response from `DisplayableListsApi.Getdisplayablelists`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetdisplayablelistsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**GetDisplayableListsRequest**](GetDisplayableListsRequest.md) |  | 
 **clientVersion** | **string** | The version of the client making the request. | 
 **domain** | **string** | The domain of the top-level page in which the client is being run. For embedded search and NSR, it will be the domain of the embedding page. For sidebar, it will be the domain of the embedding page if it&#39;s a recognized workplace app. For NTP and app.glean.com requests, it will be empty. | 
 **eids** | **[]int64** | List of experiment ids to force for incoming request. | 
 **actas** | **string** | Email of another user to act as for debugging purposes. Requires sufficient permissions. | 

### Return type

[**GetDisplayableListsResponse**](GetDisplayableListsResponse.md)

### Authorization

[actAsBearerToken](../README.md#actAsBearerToken), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Updatedisplayablelists

> UpdateDisplayableListsResponse Updatedisplayablelists(ctx).Payload(payload).ClientVersion(clientVersion).Domain(domain).Eids(eids).Execute()

Update displayable lists



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    payload := *openapiclient.NewUpdateDisplayableListsRequest([]openapiclient.DisplayableList{*openapiclient.NewDisplayableList()}) // UpdateDisplayableListsRequest | Updated version of the displayable list configs.
    clientVersion := "clientVersion_example" // string | The version of the client making the request. (optional)
    domain := "domain_example" // string | The domain of the top-level page in which the client is being run. For embedded search and NSR, it will be the domain of the embedding page. For sidebar, it will be the domain of the embedding page if it's a recognized workplace app. For NTP and app.glean.com requests, it will be empty. (optional)
    eids := []int64{int64(123)} // []int64 | List of experiment ids to force for incoming request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DisplayableListsApi.Updatedisplayablelists(context.Background()).Payload(payload).ClientVersion(clientVersion).Domain(domain).Eids(eids).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DisplayableListsApi.Updatedisplayablelists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Updatedisplayablelists`: UpdateDisplayableListsResponse
    fmt.Fprintf(os.Stdout, "Response from `DisplayableListsApi.Updatedisplayablelists`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdatedisplayablelistsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**UpdateDisplayableListsRequest**](UpdateDisplayableListsRequest.md) | Updated version of the displayable list configs. | 
 **clientVersion** | **string** | The version of the client making the request. | 
 **domain** | **string** | The domain of the top-level page in which the client is being run. For embedded search and NSR, it will be the domain of the embedding page. For sidebar, it will be the domain of the embedding page if it&#39;s a recognized workplace app. For NTP and app.glean.com requests, it will be empty. | 
 **eids** | **[]int64** | List of experiment ids to force for incoming request. | 

### Return type

[**UpdateDisplayableListsResponse**](UpdateDisplayableListsResponse.md)

### Authorization

[actAsBearerToken](../README.md#actAsBearerToken), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

