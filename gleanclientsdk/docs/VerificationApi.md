# \VerificationApi

All URIs are relative to *https://domain-be.glean.com/rest/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Addverificationreminder**](VerificationApi.md#Addverificationreminder) | **Post** /addverificationreminder | Create verification reminder
[**Listverifications**](VerificationApi.md#Listverifications) | **Post** /listverifications | List verifications
[**Verify**](VerificationApi.md#Verify) | **Post** /verify | Update verification



## Addverificationreminder

> Verification Addverificationreminder(ctx).Payload(payload).ClientVersion(clientVersion).Domain(domain).Eids(eids).Execute()

Create verification reminder



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
    payload := *openapiclient.NewReminderRequest("DocumentId_example") // ReminderRequest | Details about the reminder.
    clientVersion := "clientVersion_example" // string | The version of the client making the request. (optional)
    domain := "domain_example" // string | The domain of the top-level page in which the client is being run. For embedded search and NSR, it will be the domain of the embedding page. For sidebar, it will be the domain of the embedding page if it's a recognized workplace app. For NTP and app.glean.com requests, it will be empty. (optional)
    eids := []int64{int64(123)} // []int64 | List of experiment ids to force for incoming request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VerificationApi.Addverificationreminder(context.Background()).Payload(payload).ClientVersion(clientVersion).Domain(domain).Eids(eids).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VerificationApi.Addverificationreminder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Addverificationreminder`: Verification
    fmt.Fprintf(os.Stdout, "Response from `VerificationApi.Addverificationreminder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddverificationreminderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**ReminderRequest**](ReminderRequest.md) | Details about the reminder. | 
 **clientVersion** | **string** | The version of the client making the request. | 
 **domain** | **string** | The domain of the top-level page in which the client is being run. For embedded search and NSR, it will be the domain of the embedding page. For sidebar, it will be the domain of the embedding page if it&#39;s a recognized workplace app. For NTP and app.glean.com requests, it will be empty. | 
 **eids** | **[]int64** | List of experiment ids to force for incoming request. | 

### Return type

[**Verification**](Verification.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Listverifications

> VerificationFeed Listverifications(ctx).Count(count).Actas(actas).ClientVersion(clientVersion).Domain(domain).Eids(eids).Execute()

List verifications



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
    count := int32(56) // int32 | Maximum number of documents to return (optional)
    actas := "actas_example" // string | Email of another user to act as for debugging purposes. Requires sufficient permissions. (optional)
    clientVersion := "clientVersion_example" // string | The version of the client making the request. (optional)
    domain := "domain_example" // string | The domain of the top-level page in which the client is being run. For embedded search and NSR, it will be the domain of the embedding page. For sidebar, it will be the domain of the embedding page if it's a recognized workplace app. For NTP and app.glean.com requests, it will be empty. (optional)
    eids := []int64{int64(123)} // []int64 | List of experiment ids to force for incoming request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VerificationApi.Listverifications(context.Background()).Count(count).Actas(actas).ClientVersion(clientVersion).Domain(domain).Eids(eids).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VerificationApi.Listverifications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Listverifications`: VerificationFeed
    fmt.Fprintf(os.Stdout, "Response from `VerificationApi.Listverifications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListverificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **int32** | Maximum number of documents to return | 
 **actas** | **string** | Email of another user to act as for debugging purposes. Requires sufficient permissions. | 
 **clientVersion** | **string** | The version of the client making the request. | 
 **domain** | **string** | The domain of the top-level page in which the client is being run. For embedded search and NSR, it will be the domain of the embedding page. For sidebar, it will be the domain of the embedding page if it&#39;s a recognized workplace app. For NTP and app.glean.com requests, it will be empty. | 
 **eids** | **[]int64** | List of experiment ids to force for incoming request. | 

### Return type

[**VerificationFeed**](VerificationFeed.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Verify

> Verification Verify(ctx).Payload(payload).ClientVersion(clientVersion).Domain(domain).Eids(eids).Execute()

Update verification



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
    payload := *openapiclient.NewVerifyRequest("DocumentId_example") // VerifyRequest | Details about the verification request.
    clientVersion := "clientVersion_example" // string | The version of the client making the request. (optional)
    domain := "domain_example" // string | The domain of the top-level page in which the client is being run. For embedded search and NSR, it will be the domain of the embedding page. For sidebar, it will be the domain of the embedding page if it's a recognized workplace app. For NTP and app.glean.com requests, it will be empty. (optional)
    eids := []int64{int64(123)} // []int64 | List of experiment ids to force for incoming request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VerificationApi.Verify(context.Background()).Payload(payload).ClientVersion(clientVersion).Domain(domain).Eids(eids).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VerificationApi.Verify``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Verify`: Verification
    fmt.Fprintf(os.Stdout, "Response from `VerificationApi.Verify`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**VerifyRequest**](VerifyRequest.md) | Details about the verification request. | 
 **clientVersion** | **string** | The version of the client making the request. | 
 **domain** | **string** | The domain of the top-level page in which the client is being run. For embedded search and NSR, it will be the domain of the embedding page. For sidebar, it will be the domain of the embedding page if it&#39;s a recognized workplace app. For NTP and app.glean.com requests, it will be empty. | 
 **eids** | **[]int64** | List of experiment ids to force for incoming request. | 

### Return type

[**Verification**](Verification.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

