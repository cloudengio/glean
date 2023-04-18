# \ImagesApi

All URIs are relative to *https://domain-be.glean.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Images**](ImagesApi.md#Images) | **Get** /images | Get images
[**Uploadimage**](ImagesApi.md#Uploadimage) | **Post** /uploadimage | upload images



## Images

> *os.File Images(ctx).Key(key).Type_(type_).Id(id).Ds(ds).Cid(cid).ClientVersion(clientVersion).Domain(domain).Eids(eids).Execute()

Get images



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
    key := "key_example" // string | Primary key for the image being asked. The key is returned by the API when an image is uploaded. If key is used, other parameters should not be used. (optional)
    type_ := openapiclient.ImageType("USER") // ImageType | The type of image requested. Supported values are listed in ImageMetadata.type enum. (optional)
    id := "id_example" // string | ID, if a specific entity/type is requested. The id may have different meaning for each type. for user, it is user id, for UGC, it is the id of the content, and so on. (optional)
    ds := "ds_example" // string | A specific datasource for which an image is requested for. The ds may have different meaning for each type and can also be empty for some. (optional)
    cid := "cid_example" // string | Content id to differentitate multiple images that can have the same type and datasource e.g. thumnail or image from content of UGC. It can also be empty. (optional)
    clientVersion := "clientVersion_example" // string | The version of the client making the request. (optional)
    domain := "domain_example" // string | The domain of the top-level page in which the client is being run. For embedded search and NSR, it will be the domain of the embedding page. For sidebar, it will be the domain of the embedding page if it's a recognized workplace app. For NTP and app.glean.com requests, it will be empty. (optional)
    eids := []int64{int64(123)} // []int64 | List of experiment ids to force for incoming request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImagesApi.Images(context.Background()).Key(key).Type_(type_).Id(id).Ds(ds).Cid(cid).ClientVersion(clientVersion).Domain(domain).Eids(eids).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImagesApi.Images``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Images`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `ImagesApi.Images`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **key** | **string** | Primary key for the image being asked. The key is returned by the API when an image is uploaded. If key is used, other parameters should not be used. | 
 **type_** | [**ImageType**](ImageType.md) | The type of image requested. Supported values are listed in ImageMetadata.type enum. | 
 **id** | **string** | ID, if a specific entity/type is requested. The id may have different meaning for each type. for user, it is user id, for UGC, it is the id of the content, and so on. | 
 **ds** | **string** | A specific datasource for which an image is requested for. The ds may have different meaning for each type and can also be empty for some. | 
 **cid** | **string** | Content id to differentitate multiple images that can have the same type and datasource e.g. thumnail or image from content of UGC. It can also be empty. | 
 **clientVersion** | **string** | The version of the client making the request. | 
 **domain** | **string** | The domain of the top-level page in which the client is being run. For embedded search and NSR, it will be the domain of the embedding page. For sidebar, it will be the domain of the embedding page if it&#39;s a recognized workplace app. For NTP and app.glean.com requests, it will be empty. | 
 **eids** | **[]int64** | List of experiment ids to force for incoming request. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[actAsBearerToken](../README.md#actAsBearerToken), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Uploadimage

> UploadImageResponse Uploadimage(ctx).Payload(payload).Type_(type_).Id(id).Ds(ds).Cid(cid).Eids(eids).ClientVersion(clientVersion).Domain(domain).Actas(actas).Execute()

upload images



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
    payload := os.NewFile(1234, "some_file") // *os.File | Content and metadata for the image. Content is in post body, metada is in url.
    type_ := openapiclient.ImageType("USER") // ImageType | The type of image requested. Supported values are listed in ImageMetadata.type enum. (optional)
    id := "id_example" // string | ID, if a specific entity/type is requested. The id may have different meaning for each type. For USER, it is user id For UGC, it is the id of the content For ICON, the doctype. (optional)
    ds := "ds_example" // string | A specific datasource for which an image is requested for. The ds may have different meaning for each type and can also be empty for some. For USER, it is empty or datasource the icon is asked for. For UGC, it is the UGC datasource. For ICON, it is datasource instance the icon is asked for. (optional)
    cid := "cid_example" // string | Content id to differentitate multiple images that can have the same type and datasource e.g. thumnail or image from content of UGC. It can also be empty. (optional)
    eids := []int64{int64(123)} // []int64 | List of experiment ids to force for incoming request. (optional)
    clientVersion := "clientVersion_example" // string | The version of the client making the request. (optional)
    domain := "domain_example" // string | The domain of the top-level page in which the client is being run. For embedded search and NSR, it will be the domain of the embedding page. For sidebar, it will be the domain of the embedding page if it's a recognized workplace app. For NTP and app.glean.com requests, it will be empty. (optional)
    actas := "actas_example" // string | Email of another user to act as for debugging purposes. Requires sufficient permissions. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImagesApi.Uploadimage(context.Background()).Payload(payload).Type_(type_).Id(id).Ds(ds).Cid(cid).Eids(eids).ClientVersion(clientVersion).Domain(domain).Actas(actas).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImagesApi.Uploadimage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Uploadimage`: UploadImageResponse
    fmt.Fprintf(os.Stdout, "Response from `ImagesApi.Uploadimage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadimageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | ***os.File** | Content and metadata for the image. Content is in post body, metada is in url. | 
 **type_** | [**ImageType**](ImageType.md) | The type of image requested. Supported values are listed in ImageMetadata.type enum. | 
 **id** | **string** | ID, if a specific entity/type is requested. The id may have different meaning for each type. For USER, it is user id For UGC, it is the id of the content For ICON, the doctype. | 
 **ds** | **string** | A specific datasource for which an image is requested for. The ds may have different meaning for each type and can also be empty for some. For USER, it is empty or datasource the icon is asked for. For UGC, it is the UGC datasource. For ICON, it is datasource instance the icon is asked for. | 
 **cid** | **string** | Content id to differentitate multiple images that can have the same type and datasource e.g. thumnail or image from content of UGC. It can also be empty. | 
 **eids** | **[]int64** | List of experiment ids to force for incoming request. | 
 **clientVersion** | **string** | The version of the client making the request. | 
 **domain** | **string** | The domain of the top-level page in which the client is being run. For embedded search and NSR, it will be the domain of the embedding page. For sidebar, it will be the domain of the embedding page if it&#39;s a recognized workplace app. For NTP and app.glean.com requests, it will be empty. | 
 **actas** | **string** | Email of another user to act as for debugging purposes. Requires sufficient permissions. | 

### Return type

[**UploadImageResponse**](UploadImageResponse.md)

### Authorization

[actAsBearerToken](../README.md#actAsBearerToken), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: image/*
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

