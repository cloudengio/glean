# \CollectionsApi

All URIs are relative to *https://domain-be.glean.com/rest/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Createcollection**](CollectionsApi.md#Createcollection) | **Post** /createcollection | Create a collection
[**Listcollections**](CollectionsApi.md#Listcollections) | **Post** /listcollections | List collections



## Createcollection

> CreateCollectionResponse Createcollection(ctx).Payload(payload).Actas(actas).ClientVersion(clientVersion).Domain(domain).Eids(eids).Execute()

Create a collection



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
    payload := *openapiclient.NewCreateCollectionRequest("Name_example") // CreateCollectionRequest | Collection content plus any additional metadata for the request.
    actas := "actas_example" // string | Email of another user to act as for debugging purposes. Requires sufficient permissions. (optional)
    clientVersion := "clientVersion_example" // string | The version of the client making the request. (optional)
    domain := "domain_example" // string | The domain of the top-level page in which the client is being run. For embedded search and NSR, it will be the domain of the embedding page. For sidebar, it will be the domain of the embedding page if it's a recognized workplace app. For NTP and app.glean.com requests, it will be empty. (optional)
    eids := []int64{int64(123)} // []int64 | List of experiment ids to force for incoming request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CollectionsApi.Createcollection(context.Background()).Payload(payload).Actas(actas).ClientVersion(clientVersion).Domain(domain).Eids(eids).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CollectionsApi.Createcollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Createcollection`: CreateCollectionResponse
    fmt.Fprintf(os.Stdout, "Response from `CollectionsApi.Createcollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatecollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**CreateCollectionRequest**](CreateCollectionRequest.md) | Collection content plus any additional metadata for the request. | 
 **actas** | **string** | Email of another user to act as for debugging purposes. Requires sufficient permissions. | 
 **clientVersion** | **string** | The version of the client making the request. | 
 **domain** | **string** | The domain of the top-level page in which the client is being run. For embedded search and NSR, it will be the domain of the embedding page. For sidebar, it will be the domain of the embedding page if it&#39;s a recognized workplace app. For NTP and app.glean.com requests, it will be empty. | 
 **eids** | **[]int64** | List of experiment ids to force for incoming request. | 

### Return type

[**CreateCollectionResponse**](CreateCollectionResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Listcollections

> ListCollectionsResponse Listcollections(ctx).Payload(payload).Actas(actas).ClientVersion(clientVersion).Domain(domain).Eids(eids).Execute()

List collections



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
    payload := *openapiclient.NewListCollectionsRequest() // ListCollectionsRequest | ListCollections request
    actas := "actas_example" // string | Email of another user to act as for debugging purposes. Requires sufficient permissions. (optional)
    clientVersion := "clientVersion_example" // string | The version of the client making the request. (optional)
    domain := "domain_example" // string | The domain of the top-level page in which the client is being run. For embedded search and NSR, it will be the domain of the embedding page. For sidebar, it will be the domain of the embedding page if it's a recognized workplace app. For NTP and app.glean.com requests, it will be empty. (optional)
    eids := []int64{int64(123)} // []int64 | List of experiment ids to force for incoming request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CollectionsApi.Listcollections(context.Background()).Payload(payload).Actas(actas).ClientVersion(clientVersion).Domain(domain).Eids(eids).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CollectionsApi.Listcollections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Listcollections`: ListCollectionsResponse
    fmt.Fprintf(os.Stdout, "Response from `CollectionsApi.Listcollections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListcollectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**ListCollectionsRequest**](ListCollectionsRequest.md) | ListCollections request | 
 **actas** | **string** | Email of another user to act as for debugging purposes. Requires sufficient permissions. | 
 **clientVersion** | **string** | The version of the client making the request. | 
 **domain** | **string** | The domain of the top-level page in which the client is being run. For embedded search and NSR, it will be the domain of the embedding page. For sidebar, it will be the domain of the embedding page if it&#39;s a recognized workplace app. For NTP and app.glean.com requests, it will be empty. | 
 **eids** | **[]int64** | List of experiment ids to force for incoming request. | 

### Return type

[**ListCollectionsResponse**](ListCollectionsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

