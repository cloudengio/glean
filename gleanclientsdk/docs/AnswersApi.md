# \AnswersApi

All URIs are relative to *https://domain-be.glean.com/rest/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Createanswer**](AnswersApi.md#Createanswer) | **Post** /createanswer | Create answer
[**Deleteanswer**](AnswersApi.md#Deleteanswer) | **Post** /deleteanswer | Delete answer
[**Editanswer**](AnswersApi.md#Editanswer) | **Post** /editanswer | Edit answer
[**Listanswers**](AnswersApi.md#Listanswers) | **Post** /listanswers | List answers created by the authed user



## Createanswer

> Answer Createanswer(ctx).Payload(payload).Actas(actas).ClientVersion(clientVersion).Domain(domain).Eids(eids).Execute()

Create answer



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
    payload := *openapiclient.NewCreateAnswerRequest(*openapiclient.NewAnswerCreationData()) // CreateAnswerRequest | CreateAnswer request
    actas := "actas_example" // string | Email of another user to act as for debugging purposes. Requires sufficient permissions. (optional)
    clientVersion := "clientVersion_example" // string | The version of the client making the request. (optional)
    domain := "domain_example" // string | The domain of the top-level page in which the client is being run. For embedded search and NSR, it will be the domain of the embedding page. For sidebar, it will be the domain of the embedding page if it's a recognized workplace app. For NTP and app.glean.com requests, it will be empty. (optional)
    eids := []int64{int64(123)} // []int64 | List of experiment ids to force for incoming request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AnswersApi.Createanswer(context.Background()).Payload(payload).Actas(actas).ClientVersion(clientVersion).Domain(domain).Eids(eids).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnswersApi.Createanswer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Createanswer`: Answer
    fmt.Fprintf(os.Stdout, "Response from `AnswersApi.Createanswer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateanswerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**CreateAnswerRequest**](CreateAnswerRequest.md) | CreateAnswer request | 
 **actas** | **string** | Email of another user to act as for debugging purposes. Requires sufficient permissions. | 
 **clientVersion** | **string** | The version of the client making the request. | 
 **domain** | **string** | The domain of the top-level page in which the client is being run. For embedded search and NSR, it will be the domain of the embedding page. For sidebar, it will be the domain of the embedding page if it&#39;s a recognized workplace app. For NTP and app.glean.com requests, it will be empty. | 
 **eids** | **[]int64** | List of experiment ids to force for incoming request. | 

### Return type

[**Answer**](Answer.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Deleteanswer

> Deleteanswer(ctx).Payload(payload).Actas(actas).ClientVersion(clientVersion).Domain(domain).Eids(eids).Execute()

Delete answer



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
    payload := *openapiclient.NewDeleteAnswerRequest(int32(3)) // DeleteAnswerRequest | DeleteAnswer request
    actas := "actas_example" // string | Email of another user to act as for debugging purposes. Requires sufficient permissions. (optional)
    clientVersion := "clientVersion_example" // string | The version of the client making the request. (optional)
    domain := "domain_example" // string | The domain of the top-level page in which the client is being run. For embedded search and NSR, it will be the domain of the embedding page. For sidebar, it will be the domain of the embedding page if it's a recognized workplace app. For NTP and app.glean.com requests, it will be empty. (optional)
    eids := []int64{int64(123)} // []int64 | List of experiment ids to force for incoming request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AnswersApi.Deleteanswer(context.Background()).Payload(payload).Actas(actas).ClientVersion(clientVersion).Domain(domain).Eids(eids).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnswersApi.Deleteanswer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteanswerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**DeleteAnswerRequest**](DeleteAnswerRequest.md) | DeleteAnswer request | 
 **actas** | **string** | Email of another user to act as for debugging purposes. Requires sufficient permissions. | 
 **clientVersion** | **string** | The version of the client making the request. | 
 **domain** | **string** | The domain of the top-level page in which the client is being run. For embedded search and NSR, it will be the domain of the embedding page. For sidebar, it will be the domain of the embedding page if it&#39;s a recognized workplace app. For NTP and app.glean.com requests, it will be empty. | 
 **eids** | **[]int64** | List of experiment ids to force for incoming request. | 

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


## Editanswer

> Answer Editanswer(ctx).Payload(payload).Actas(actas).ClientVersion(clientVersion).Domain(domain).Eids(eids).Execute()

Edit answer



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
    payload := *openapiclient.NewEditAnswerRequest(int32(3)) // EditAnswerRequest | EditAnswer request
    actas := "actas_example" // string | Email of another user to act as for debugging purposes. Requires sufficient permissions. (optional)
    clientVersion := "clientVersion_example" // string | The version of the client making the request. (optional)
    domain := "domain_example" // string | The domain of the top-level page in which the client is being run. For embedded search and NSR, it will be the domain of the embedding page. For sidebar, it will be the domain of the embedding page if it's a recognized workplace app. For NTP and app.glean.com requests, it will be empty. (optional)
    eids := []int64{int64(123)} // []int64 | List of experiment ids to force for incoming request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AnswersApi.Editanswer(context.Background()).Payload(payload).Actas(actas).ClientVersion(clientVersion).Domain(domain).Eids(eids).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnswersApi.Editanswer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Editanswer`: Answer
    fmt.Fprintf(os.Stdout, "Response from `AnswersApi.Editanswer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEditanswerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**EditAnswerRequest**](EditAnswerRequest.md) | EditAnswer request | 
 **actas** | **string** | Email of another user to act as for debugging purposes. Requires sufficient permissions. | 
 **clientVersion** | **string** | The version of the client making the request. | 
 **domain** | **string** | The domain of the top-level page in which the client is being run. For embedded search and NSR, it will be the domain of the embedding page. For sidebar, it will be the domain of the embedding page if it&#39;s a recognized workplace app. For NTP and app.glean.com requests, it will be empty. | 
 **eids** | **[]int64** | List of experiment ids to force for incoming request. | 

### Return type

[**Answer**](Answer.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Listanswers

> ListAnswersResponse Listanswers(ctx).Payload(payload).Actas(actas).ClientVersion(clientVersion).Domain(domain).Eids(eids).Execute()

List answers created by the authed user



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
    payload := *openapiclient.NewListAnswersRequest() // ListAnswersRequest | ListAnswers request
    actas := "actas_example" // string | Email of another user to act as for debugging purposes. Requires sufficient permissions. (optional)
    clientVersion := "clientVersion_example" // string | The version of the client making the request. (optional)
    domain := "domain_example" // string | The domain of the top-level page in which the client is being run. For embedded search and NSR, it will be the domain of the embedding page. For sidebar, it will be the domain of the embedding page if it's a recognized workplace app. For NTP and app.glean.com requests, it will be empty. (optional)
    eids := []int64{int64(123)} // []int64 | List of experiment ids to force for incoming request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AnswersApi.Listanswers(context.Background()).Payload(payload).Actas(actas).ClientVersion(clientVersion).Domain(domain).Eids(eids).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnswersApi.Listanswers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Listanswers`: ListAnswersResponse
    fmt.Fprintf(os.Stdout, "Response from `AnswersApi.Listanswers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListanswersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**ListAnswersRequest**](ListAnswersRequest.md) | ListAnswers request | 
 **actas** | **string** | Email of another user to act as for debugging purposes. Requires sufficient permissions. | 
 **clientVersion** | **string** | The version of the client making the request. | 
 **domain** | **string** | The domain of the top-level page in which the client is being run. For embedded search and NSR, it will be the domain of the embedding page. For sidebar, it will be the domain of the embedding page if it&#39;s a recognized workplace app. For NTP and app.glean.com requests, it will be empty. | 
 **eids** | **[]int64** | List of experiment ids to force for incoming request. | 

### Return type

[**ListAnswersResponse**](ListAnswersResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

