# \PartsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PartsList**](PartsApi.md#PartsList) | **Get** /api/parts/ | 
[**PartsRetrieve**](PartsApi.md#PartsRetrieve) | **Get** /api/parts/{part_no}/ | 



## PartsList

> PaginatedPartList PartsList(ctx).BoardNo(boardNo).ChipNo(chipNo).Ordering(ordering).Page(page).PartFamily(partFamily).PartNo(partNo).Revision(revision).Search(search).Variant(variant).Execute()



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
    boardNo := "boardNo_example" // string |  (optional)
    chipNo := "chipNo_example" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    partFamily := "partFamily_example" // string |  (optional)
    partNo := "partNo_example" // string |  (optional)
    revision := "revision_example" // string |  (optional)
    search := "search_example" // string | A search term. (optional)
    variant := "variant_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartsApi.PartsList(context.Background()).BoardNo(boardNo).ChipNo(chipNo).Ordering(ordering).Page(page).PartFamily(partFamily).PartNo(partNo).Revision(revision).Search(search).Variant(variant).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartsApi.PartsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PartsList`: PaginatedPartList
    fmt.Fprintf(os.Stdout, "Response from `PartsApi.PartsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPartsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **boardNo** | **string** |  | 
 **chipNo** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **partFamily** | **string** |  | 
 **partNo** | **string** |  | 
 **revision** | **string** |  | 
 **search** | **string** | A search term. | 
 **variant** | **string** |  | 

### Return type

[**PaginatedPartList**](PaginatedPartList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartsRetrieve

> Part PartsRetrieve(ctx, partNo).Execute()



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
    partNo := "partNo_example" // string | A unique value identifying this part.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartsApi.PartsRetrieve(context.Background(), partNo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartsApi.PartsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PartsRetrieve`: Part
    fmt.Fprintf(os.Stdout, "Response from `PartsApi.PartsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**partNo** | **string** | A unique value identifying this part. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Part**](Part.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

