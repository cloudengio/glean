# \DatasourcesApi

All URIs are relative to *https://domain-be.glean.com/api/index/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdddatasourcePost**](DatasourcesApi.md#AdddatasourcePost) | **Post** /adddatasource | Add datasource
[**GetdatasourceconfigPost**](DatasourcesApi.md#GetdatasourceconfigPost) | **Post** /getdatasourceconfig | Get datasource config



## AdddatasourcePost

> AdddatasourcePost(ctx).CustomDatasourceConfig(customDatasourceConfig).Execute()

Add datasource



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
    customDatasourceConfig := *openapiclient.NewCustomDatasourceConfig("Name_example") // CustomDatasourceConfig | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatasourcesApi.AdddatasourcePost(context.Background()).CustomDatasourceConfig(customDatasourceConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatasourcesApi.AdddatasourcePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdddatasourcePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customDatasourceConfig** | [**CustomDatasourceConfig**](CustomDatasourceConfig.md) |  | 

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


## GetdatasourceconfigPost

> CustomDatasourceConfig GetdatasourceconfigPost(ctx).GetDatasourceConfigRequest(getDatasourceConfigRequest).Execute()

Get datasource config



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
    getDatasourceConfigRequest := *openapiclient.NewGetDatasourceConfigRequest() // GetDatasourceConfigRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatasourcesApi.GetdatasourceconfigPost(context.Background()).GetDatasourceConfigRequest(getDatasourceConfigRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatasourcesApi.GetdatasourceconfigPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetdatasourceconfigPost`: CustomDatasourceConfig
    fmt.Fprintf(os.Stdout, "Response from `DatasourcesApi.GetdatasourceconfigPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetdatasourceconfigPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getDatasourceConfigRequest** | [**GetDatasourceConfigRequest**](GetDatasourceConfigRequest.md) |  | 

### Return type

[**CustomDatasourceConfig**](CustomDatasourceConfig.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

