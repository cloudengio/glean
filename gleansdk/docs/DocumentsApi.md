# \DocumentsApi

All URIs are relative to *https://domain-be.glean.com/api/index/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkindexdocumentsPost**](DocumentsApi.md#BulkindexdocumentsPost) | **Post** /bulkindexdocuments | Bulk index documents
[**DeletedocumentPost**](DocumentsApi.md#DeletedocumentPost) | **Post** /deletedocument | Delete document
[**GetdocumentcountPost**](DocumentsApi.md#GetdocumentcountPost) | **Post** /getdocumentcount | Get document count
[**GetdocumentstatusPost**](DocumentsApi.md#GetdocumentstatusPost) | **Post** /getdocumentstatus | Get document upload and indexing status
[**IndexdocumentPost**](DocumentsApi.md#IndexdocumentPost) | **Post** /indexdocument | Index document
[**ProcessalldocumentsPost**](DocumentsApi.md#ProcessalldocumentsPost) | **Post** /processalldocuments | Schedules the processing of uploaded documents



## BulkindexdocumentsPost

> BulkindexdocumentsPost(ctx).BulkIndexDocumentsRequest(bulkIndexDocumentsRequest).Execute()

Bulk index documents



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
    bulkIndexDocumentsRequest := *openapiclient.NewBulkIndexDocumentsRequest("UploadId_example", "Datasource_example", []openapiclient.DocumentDefinition{*openapiclient.NewDocumentDefinition("Datasource_example")}) // BulkIndexDocumentsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentsApi.BulkindexdocumentsPost(context.Background()).BulkIndexDocumentsRequest(bulkIndexDocumentsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsApi.BulkindexdocumentsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkindexdocumentsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bulkIndexDocumentsRequest** | [**BulkIndexDocumentsRequest**](BulkIndexDocumentsRequest.md) |  | 

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


## DeletedocumentPost

> DeletedocumentPost(ctx).DeleteDocumentRequest(deleteDocumentRequest).Execute()

Delete document



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
    deleteDocumentRequest := *openapiclient.NewDeleteDocumentRequest("Datasource_example", "ObjectType_example", "Id_example") // DeleteDocumentRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentsApi.DeletedocumentPost(context.Background()).DeleteDocumentRequest(deleteDocumentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsApi.DeletedocumentPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeletedocumentPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteDocumentRequest** | [**DeleteDocumentRequest**](DeleteDocumentRequest.md) |  | 

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


## GetdocumentcountPost

> GetDocumentCountResponse GetdocumentcountPost(ctx).GetDocumentCountRequest(getDocumentCountRequest).Execute()

Get document count



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
    getDocumentCountRequest := *openapiclient.NewGetDocumentCountRequest("Name_example") // GetDocumentCountRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentsApi.GetdocumentcountPost(context.Background()).GetDocumentCountRequest(getDocumentCountRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsApi.GetdocumentcountPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetdocumentcountPost`: GetDocumentCountResponse
    fmt.Fprintf(os.Stdout, "Response from `DocumentsApi.GetdocumentcountPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetdocumentcountPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getDocumentCountRequest** | [**GetDocumentCountRequest**](GetDocumentCountRequest.md) |  | 

### Return type

[**GetDocumentCountResponse**](GetDocumentCountResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetdocumentstatusPost

> GetDocumentStatusResponse GetdocumentstatusPost(ctx).GetDocumentStatusRequest(getDocumentStatusRequest).Execute()

Get document upload and indexing status



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
    getDocumentStatusRequest := *openapiclient.NewGetDocumentStatusRequest("Datasource_example", "ObjectType_example", "DocId_example") // GetDocumentStatusRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentsApi.GetdocumentstatusPost(context.Background()).GetDocumentStatusRequest(getDocumentStatusRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsApi.GetdocumentstatusPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetdocumentstatusPost`: GetDocumentStatusResponse
    fmt.Fprintf(os.Stdout, "Response from `DocumentsApi.GetdocumentstatusPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetdocumentstatusPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getDocumentStatusRequest** | [**GetDocumentStatusRequest**](GetDocumentStatusRequest.md) |  | 

### Return type

[**GetDocumentStatusResponse**](GetDocumentStatusResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IndexdocumentPost

> IndexdocumentPost(ctx).IndexDocumentRequest(indexDocumentRequest).Execute()

Index document



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
    indexDocumentRequest := *openapiclient.NewIndexDocumentRequest(*openapiclient.NewDocumentDefinition("Datasource_example")) // IndexDocumentRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentsApi.IndexdocumentPost(context.Background()).IndexDocumentRequest(indexDocumentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsApi.IndexdocumentPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIndexdocumentPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **indexDocumentRequest** | [**IndexDocumentRequest**](IndexDocumentRequest.md) |  | 

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


## ProcessalldocumentsPost

> ProcessalldocumentsPost(ctx).ProcessAllDocumentsRequest(processAllDocumentsRequest).Execute()

Schedules the processing of uploaded documents



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
    processAllDocumentsRequest := *openapiclient.NewProcessAllDocumentsRequest() // ProcessAllDocumentsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentsApi.ProcessalldocumentsPost(context.Background()).ProcessAllDocumentsRequest(processAllDocumentsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsApi.ProcessalldocumentsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProcessalldocumentsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processAllDocumentsRequest** | [**ProcessAllDocumentsRequest**](ProcessAllDocumentsRequest.md) |  | 

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

