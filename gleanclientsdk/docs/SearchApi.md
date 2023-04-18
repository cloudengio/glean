# \SearchApi

All URIs are relative to *https://domain-be.glean.com/rest/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Autocomplete**](SearchApi.md#Autocomplete) | **Post** /autocomplete | Autocomplete a query
[**Listentities**](SearchApi.md#Listentities) | **Post** /listentities | List entities
[**People**](SearchApi.md#People) | **Post** /people | Read people details
[**Search**](SearchApi.md#Search) | **Post** /search | Search the index



## Autocomplete

> AutocompleteResponse Autocomplete(ctx).Payload(payload).Actas(actas).ClientVersion(clientVersion).Domain(domain).Eids(eids).Execute()

Autocomplete a query



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
    payload := *openapiclient.NewAutocompleteRequest() // AutocompleteRequest | Autocomplete request
    actas := "actas_example" // string | Email of another user to act as for debugging purposes. Requires sufficient permissions. (optional)
    clientVersion := "clientVersion_example" // string | The version of the client making the request. (optional)
    domain := "domain_example" // string | The domain of the top-level page in which the client is being run. For embedded search and NSR, it will be the domain of the embedding page. For sidebar, it will be the domain of the embedding page if it's a recognized workplace app. For NTP and app.glean.com requests, it will be empty. (optional)
    eids := []int64{int64(123)} // []int64 | List of experiment ids to force for incoming request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchApi.Autocomplete(context.Background()).Payload(payload).Actas(actas).ClientVersion(clientVersion).Domain(domain).Eids(eids).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.Autocomplete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Autocomplete`: AutocompleteResponse
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.Autocomplete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAutocompleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**AutocompleteRequest**](AutocompleteRequest.md) | Autocomplete request | 
 **actas** | **string** | Email of another user to act as for debugging purposes. Requires sufficient permissions. | 
 **clientVersion** | **string** | The version of the client making the request. | 
 **domain** | **string** | The domain of the top-level page in which the client is being run. For embedded search and NSR, it will be the domain of the embedding page. For sidebar, it will be the domain of the embedding page if it&#39;s a recognized workplace app. For NTP and app.glean.com requests, it will be empty. | 
 **eids** | **[]int64** | List of experiment ids to force for incoming request. | 

### Return type

[**AutocompleteResponse**](AutocompleteResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Listentities

> ListEntitiesResponse Listentities(ctx).Payload(payload).Actas(actas).ClientVersion(clientVersion).Domain(domain).Eids(eids).Execute()

List entities



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
    payload := *openapiclient.NewListEntitiesRequest() // ListEntitiesRequest | List people request
    actas := "actas_example" // string | Email of another user to act as for debugging purposes. Requires sufficient permissions. (optional)
    clientVersion := "clientVersion_example" // string | The version of the client making the request. (optional)
    domain := "domain_example" // string | The domain of the top-level page in which the client is being run. For embedded search and NSR, it will be the domain of the embedding page. For sidebar, it will be the domain of the embedding page if it's a recognized workplace app. For NTP and app.glean.com requests, it will be empty. (optional)
    eids := []int64{int64(123)} // []int64 | List of experiment ids to force for incoming request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchApi.Listentities(context.Background()).Payload(payload).Actas(actas).ClientVersion(clientVersion).Domain(domain).Eids(eids).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.Listentities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Listentities`: ListEntitiesResponse
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.Listentities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListentitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**ListEntitiesRequest**](ListEntitiesRequest.md) | List people request | 
 **actas** | **string** | Email of another user to act as for debugging purposes. Requires sufficient permissions. | 
 **clientVersion** | **string** | The version of the client making the request. | 
 **domain** | **string** | The domain of the top-level page in which the client is being run. For embedded search and NSR, it will be the domain of the embedding page. For sidebar, it will be the domain of the embedding page if it&#39;s a recognized workplace app. For NTP and app.glean.com requests, it will be empty. | 
 **eids** | **[]int64** | List of experiment ids to force for incoming request. | 

### Return type

[**ListEntitiesResponse**](ListEntitiesResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## People

> PeopleResponse People(ctx).Payload(payload).Actas(actas).ClientVersion(clientVersion).Domain(domain).Eids(eids).Execute()

Read people details



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
    payload := *openapiclient.NewPeopleRequest() // PeopleRequest | People request
    actas := "actas_example" // string | Email of another user to act as for debugging purposes. Requires sufficient permissions. (optional)
    clientVersion := "clientVersion_example" // string | The version of the client making the request. (optional)
    domain := "domain_example" // string | The domain of the top-level page in which the client is being run. For embedded search and NSR, it will be the domain of the embedding page. For sidebar, it will be the domain of the embedding page if it's a recognized workplace app. For NTP and app.glean.com requests, it will be empty. (optional)
    eids := []int64{int64(123)} // []int64 | List of experiment ids to force for incoming request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchApi.People(context.Background()).Payload(payload).Actas(actas).ClientVersion(clientVersion).Domain(domain).Eids(eids).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.People``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `People`: PeopleResponse
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.People`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPeopleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**PeopleRequest**](PeopleRequest.md) | People request | 
 **actas** | **string** | Email of another user to act as for debugging purposes. Requires sufficient permissions. | 
 **clientVersion** | **string** | The version of the client making the request. | 
 **domain** | **string** | The domain of the top-level page in which the client is being run. For embedded search and NSR, it will be the domain of the embedding page. For sidebar, it will be the domain of the embedding page if it&#39;s a recognized workplace app. For NTP and app.glean.com requests, it will be empty. | 
 **eids** | **[]int64** | List of experiment ids to force for incoming request. | 

### Return type

[**PeopleResponse**](PeopleResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Search

> SearchResponse Search(ctx).Actas(actas).ClientVersion(clientVersion).Domain(domain).Eids(eids).Payload(payload).Execute()

Search the index



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
    actas := "actas_example" // string | Email of another user to act as for debugging purposes. Requires sufficient permissions. (optional)
    clientVersion := "clientVersion_example" // string | The version of the client making the request. (optional)
    domain := "domain_example" // string | The domain of the top-level page in which the client is being run. For embedded search and NSR, it will be the domain of the embedding page. For sidebar, it will be the domain of the embedding page if it's a recognized workplace app. For NTP and app.glean.com requests, it will be empty. (optional)
    eids := []int64{int64(123)} // []int64 | List of experiment ids to force for incoming request. (optional)
    payload := *openapiclient.NewSearchRequest(*openapiclient.NewSearchRequestSourceInfo("Modality_example")) // SearchRequest | Search request (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchApi.Search(context.Background()).Actas(actas).ClientVersion(clientVersion).Domain(domain).Eids(eids).Payload(payload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.Search``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Search`: SearchResponse
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.Search`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **actas** | **string** | Email of another user to act as for debugging purposes. Requires sufficient permissions. | 
 **clientVersion** | **string** | The version of the client making the request. | 
 **domain** | **string** | The domain of the top-level page in which the client is being run. For embedded search and NSR, it will be the domain of the embedding page. For sidebar, it will be the domain of the embedding page if it&#39;s a recognized workplace app. For NTP and app.glean.com requests, it will be empty. | 
 **eids** | **[]int64** | List of experiment ids to force for incoming request. | 
 **payload** | [**SearchRequest**](SearchRequest.md) | Search request | 

### Return type

[**SearchResponse**](SearchResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

