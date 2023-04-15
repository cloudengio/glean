# \ActivityApi

All URIs are relative to *https://domain-be.glean.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Activity**](ActivityApi.md#Activity) | **Post** /activity | Report activity of enterprise app usage
[**Ca**](ActivityApi.md#Ca) | **Post** /ca | Report client events such as clicks, views, impressions or any other product interactions.
[**Customfeedback**](ActivityApi.md#Customfeedback) | **Post** /custom_feedback | Reports arbitrary customer-defined feedback
[**Feedback**](ActivityApi.md#Feedback) | **Post** /feedback | Report user events such as clicks or views



## Activity

> Activity(ctx).Payload(payload).Execute()

Report activity of enterprise app usage



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    payload := *openapiclient.NewActivity([]openapiclient.ActivityEvent{*openapiclient.NewActivityEvent("Action_example", "Source_example", time.Now(), "Url_example")}) // Activity | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ActivityApi.Activity(context.Background()).Payload(payload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.Activity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiActivityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**Activity**](Activity.md) |  | 

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


## Ca

> Ca(ctx).ClientVersion(clientVersion).Domain(domain).Eids(eids).Payload(payload).Execute()

Report client events such as clicks, views, impressions or any other product interactions.



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
    clientVersion := "clientVersion_example" // string | The version of the client making the request. (optional)
    domain := "domain_example" // string | The domain of the top-level page in which the client is being run. For embedded search and NSR, it will be the domain of the embedding page. For sidebar, it will be the domain of the embedding page if it's a recognized workplace app. For NTP and app.glean.com requests, it will be empty. (optional)
    eids := []int64{int64(123)} // []int64 | List of experiment ids to force for incoming request. (optional)
    payload := *openapiclient.NewClientAnalyticsRequest() // ClientAnalyticsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ActivityApi.Ca(context.Background()).ClientVersion(clientVersion).Domain(domain).Eids(eids).Payload(payload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.Ca``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientVersion** | **string** | The version of the client making the request. | 
 **domain** | **string** | The domain of the top-level page in which the client is being run. For embedded search and NSR, it will be the domain of the embedding page. For sidebar, it will be the domain of the embedding page if it&#39;s a recognized workplace app. For NTP and app.glean.com requests, it will be empty. | 
 **eids** | **[]int64** | List of experiment ids to force for incoming request. | 
 **payload** | [**ClientAnalyticsRequest**](ClientAnalyticsRequest.md) |  | 

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


## Customfeedback

> Customfeedback(ctx).Payload(payload).Execute()

Reports arbitrary customer-defined feedback



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
    payload := interface{}({"arbitrary":"jsonData","for":"myCompany","comment":"Glean is awesome!"}) // interface{} | Arbitrary customer-defined JSON (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ActivityApi.Customfeedback(context.Background()).Payload(payload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.Customfeedback``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomfeedbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | **interface{}** | Arbitrary customer-defined JSON | 

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


## Feedback

> Feedback(ctx).Feedback(feedback).Actas(actas).ClientVersion(clientVersion).Domain(domain).Payload(payload).Execute()

Report user events such as clicks or views



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
    feedback := "feedback_example" // string | A URL encoded versions of Feedback. This is useful for requests. (optional)
    actas := "actas_example" // string | Email of another user to act as for debugging purposes. Requires sufficient permissions. (optional)
    clientVersion := "clientVersion_example" // string | The version of the client making the request. (optional)
    domain := "domain_example" // string | The domain of the top-level page in which the client is being run. For embedded search and NSR, it will be the domain of the embedding page. For sidebar, it will be the domain of the embedding page if it's a recognized workplace app. For NTP and app.glean.com requests, it will be empty. (optional)
    payload := *openapiclient.NewFeedback([]string{"TrackingTokens_example"}, "Event_example") // Feedback |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ActivityApi.Feedback(context.Background()).Feedback(feedback).Actas(actas).ClientVersion(clientVersion).Domain(domain).Payload(payload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.Feedback``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFeedbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **feedback** | **string** | A URL encoded versions of Feedback. This is useful for requests. | 
 **actas** | **string** | Email of another user to act as for debugging purposes. Requires sufficient permissions. | 
 **clientVersion** | **string** | The version of the client making the request. | 
 **domain** | **string** | The domain of the top-level page in which the client is being run. For embedded search and NSR, it will be the domain of the embedding page. For sidebar, it will be the domain of the embedding page if it&#39;s a recognized workplace app. For NTP and app.glean.com requests, it will be empty. | 
 **payload** | [**Feedback**](Feedback.md) |  | 

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

