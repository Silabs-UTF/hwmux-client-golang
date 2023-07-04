# \PartFamiliesApi

All URIs are relative to *https://hwmux.silabs.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PartFamiliesCreate**](PartFamiliesApi.md#PartFamiliesCreate) | **Post** /api/part-families/ | 
[**PartFamiliesDestroy**](PartFamiliesApi.md#PartFamiliesDestroy) | **Delete** /api/part-families/{name}/ | 
[**PartFamiliesList**](PartFamiliesApi.md#PartFamiliesList) | **Get** /api/part-families/ | 
[**PartFamiliesPartialUpdate**](PartFamiliesApi.md#PartFamiliesPartialUpdate) | **Patch** /api/part-families/{name}/ | 
[**PartFamiliesPartsCreate**](PartFamiliesApi.md#PartFamiliesPartsCreate) | **Post** /api/part-families/{part_family_pk}/parts/ | 
[**PartFamiliesPartsDestroy**](PartFamiliesApi.md#PartFamiliesPartsDestroy) | **Delete** /api/part-families/{part_family_pk}/parts/{id}/ | 
[**PartFamiliesPartsList**](PartFamiliesApi.md#PartFamiliesPartsList) | **Get** /api/part-families/{part_family_pk}/parts/ | 
[**PartFamiliesPartsPartialUpdate**](PartFamiliesApi.md#PartFamiliesPartsPartialUpdate) | **Patch** /api/part-families/{part_family_pk}/parts/{id}/ | 
[**PartFamiliesPartsRetrieve**](PartFamiliesApi.md#PartFamiliesPartsRetrieve) | **Get** /api/part-families/{part_family_pk}/parts/{id}/ | 
[**PartFamiliesPartsUpdate**](PartFamiliesApi.md#PartFamiliesPartsUpdate) | **Put** /api/part-families/{part_family_pk}/parts/{id}/ | 
[**PartFamiliesRetrieve**](PartFamiliesApi.md#PartFamiliesRetrieve) | **Get** /api/part-families/{name}/ | 
[**PartFamiliesUpdate**](PartFamiliesApi.md#PartFamiliesUpdate) | **Put** /api/part-families/{name}/ | 



## PartFamiliesCreate

> PartFamily PartFamiliesCreate(ctx).PartFamily(partFamily).Execute()



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
    partFamily := *openapiclient.NewPartFamily("Name_example") // PartFamily | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartFamiliesApi.PartFamiliesCreate(context.Background()).PartFamily(partFamily).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartFamiliesApi.PartFamiliesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PartFamiliesCreate`: PartFamily
    fmt.Fprintf(os.Stdout, "Response from `PartFamiliesApi.PartFamiliesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPartFamiliesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **partFamily** | [**PartFamily**](PartFamily.md) |  | 

### Return type

[**PartFamily**](PartFamily.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartFamiliesDestroy

> PartFamiliesDestroy(ctx, name).Execute()



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
    name := "name_example" // string | A unique value identifying this part family.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartFamiliesApi.PartFamiliesDestroy(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartFamiliesApi.PartFamiliesDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | A unique value identifying this part family. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartFamiliesDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartFamiliesList

> PaginatedPartFamilyList PartFamiliesList(ctx).Description(description).Name(name).Ordering(ordering).Page(page).RegexPattern(regexPattern).Search(search).Execute()



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
    description := "description_example" // string |  (optional)
    name := "name_example" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    regexPattern := "regexPattern_example" // string |  (optional)
    search := "search_example" // string | A search term. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartFamiliesApi.PartFamiliesList(context.Background()).Description(description).Name(name).Ordering(ordering).Page(page).RegexPattern(regexPattern).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartFamiliesApi.PartFamiliesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PartFamiliesList`: PaginatedPartFamilyList
    fmt.Fprintf(os.Stdout, "Response from `PartFamiliesApi.PartFamiliesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPartFamiliesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **description** | **string** |  | 
 **name** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **regexPattern** | **string** |  | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedPartFamilyList**](PaginatedPartFamilyList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartFamiliesPartialUpdate

> PartFamily PartFamiliesPartialUpdate(ctx, name).PatchedPartFamily(patchedPartFamily).Execute()



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
    name := "name_example" // string | A unique value identifying this part family.
    patchedPartFamily := *openapiclient.NewPatchedPartFamily() // PatchedPartFamily |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartFamiliesApi.PartFamiliesPartialUpdate(context.Background(), name).PatchedPartFamily(patchedPartFamily).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartFamiliesApi.PartFamiliesPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PartFamiliesPartialUpdate`: PartFamily
    fmt.Fprintf(os.Stdout, "Response from `PartFamiliesApi.PartFamiliesPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | A unique value identifying this part family. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartFamiliesPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedPartFamily** | [**PatchedPartFamily**](PatchedPartFamily.md) |  | 

### Return type

[**PartFamily**](PartFamily.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartFamiliesPartsCreate

> Part PartFamiliesPartsCreate(ctx, partFamilyPk).Part(part).Execute()



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
    partFamilyPk := "partFamilyPk_example" // string | 
    part := *openapiclient.NewPart("PartNo_example", *openapiclient.NewPartFamily("Name_example"), "BoardNo_example") // Part | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartFamiliesApi.PartFamiliesPartsCreate(context.Background(), partFamilyPk).Part(part).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartFamiliesApi.PartFamiliesPartsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PartFamiliesPartsCreate`: Part
    fmt.Fprintf(os.Stdout, "Response from `PartFamiliesApi.PartFamiliesPartsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**partFamilyPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartFamiliesPartsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **part** | [**Part**](Part.md) |  | 

### Return type

[**Part**](Part.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartFamiliesPartsDestroy

> PartFamiliesPartsDestroy(ctx, id, partFamilyPk).Execute()



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
    id := "id_example" // string | 
    partFamilyPk := "partFamilyPk_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartFamiliesApi.PartFamiliesPartsDestroy(context.Background(), id, partFamilyPk).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartFamiliesApi.PartFamiliesPartsDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**partFamilyPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartFamiliesPartsDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartFamiliesPartsList

> PaginatedPartList PartFamiliesPartsList(ctx, partFamilyPk).BoardNo(boardNo).ChipNo(chipNo).Ordering(ordering).Page(page).PartFamily(partFamily).PartNo(partNo).Revision(revision).Search(search).Variant(variant).Execute()



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
    partFamilyPk := "partFamilyPk_example" // string | 
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
    resp, r, err := apiClient.PartFamiliesApi.PartFamiliesPartsList(context.Background(), partFamilyPk).BoardNo(boardNo).ChipNo(chipNo).Ordering(ordering).Page(page).PartFamily(partFamily).PartNo(partNo).Revision(revision).Search(search).Variant(variant).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartFamiliesApi.PartFamiliesPartsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PartFamiliesPartsList`: PaginatedPartList
    fmt.Fprintf(os.Stdout, "Response from `PartFamiliesApi.PartFamiliesPartsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**partFamilyPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartFamiliesPartsListRequest struct via the builder pattern


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


## PartFamiliesPartsPartialUpdate

> Part PartFamiliesPartsPartialUpdate(ctx, id, partFamilyPk).PatchedPart(patchedPart).Execute()



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
    id := "id_example" // string | 
    partFamilyPk := "partFamilyPk_example" // string | 
    patchedPart := *openapiclient.NewPatchedPart() // PatchedPart |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartFamiliesApi.PartFamiliesPartsPartialUpdate(context.Background(), id, partFamilyPk).PatchedPart(patchedPart).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartFamiliesApi.PartFamiliesPartsPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PartFamiliesPartsPartialUpdate`: Part
    fmt.Fprintf(os.Stdout, "Response from `PartFamiliesApi.PartFamiliesPartsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**partFamilyPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartFamiliesPartsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedPart** | [**PatchedPart**](PatchedPart.md) |  | 

### Return type

[**Part**](Part.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartFamiliesPartsRetrieve

> Part PartFamiliesPartsRetrieve(ctx, id, partFamilyPk).Execute()



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
    id := "id_example" // string | 
    partFamilyPk := "partFamilyPk_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartFamiliesApi.PartFamiliesPartsRetrieve(context.Background(), id, partFamilyPk).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartFamiliesApi.PartFamiliesPartsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PartFamiliesPartsRetrieve`: Part
    fmt.Fprintf(os.Stdout, "Response from `PartFamiliesApi.PartFamiliesPartsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**partFamilyPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartFamiliesPartsRetrieveRequest struct via the builder pattern


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


## PartFamiliesPartsUpdate

> Part PartFamiliesPartsUpdate(ctx, id, partFamilyPk).Part(part).Execute()



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
    id := "id_example" // string | 
    partFamilyPk := "partFamilyPk_example" // string | 
    part := *openapiclient.NewPart("PartNo_example", *openapiclient.NewPartFamily("Name_example"), "BoardNo_example") // Part | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartFamiliesApi.PartFamiliesPartsUpdate(context.Background(), id, partFamilyPk).Part(part).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartFamiliesApi.PartFamiliesPartsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PartFamiliesPartsUpdate`: Part
    fmt.Fprintf(os.Stdout, "Response from `PartFamiliesApi.PartFamiliesPartsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**partFamilyPk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartFamiliesPartsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **part** | [**Part**](Part.md) |  | 

### Return type

[**Part**](Part.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartFamiliesRetrieve

> PartFamily PartFamiliesRetrieve(ctx, name).Execute()



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
    name := "name_example" // string | A unique value identifying this part family.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartFamiliesApi.PartFamiliesRetrieve(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartFamiliesApi.PartFamiliesRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PartFamiliesRetrieve`: PartFamily
    fmt.Fprintf(os.Stdout, "Response from `PartFamiliesApi.PartFamiliesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | A unique value identifying this part family. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartFamiliesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PartFamily**](PartFamily.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartFamiliesUpdate

> PartFamily PartFamiliesUpdate(ctx, name).PartFamily(partFamily).Execute()



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
    name := "name_example" // string | A unique value identifying this part family.
    partFamily := *openapiclient.NewPartFamily("Name_example") // PartFamily | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartFamiliesApi.PartFamiliesUpdate(context.Background(), name).PartFamily(partFamily).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartFamiliesApi.PartFamiliesUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PartFamiliesUpdate`: PartFamily
    fmt.Fprintf(os.Stdout, "Response from `PartFamiliesApi.PartFamiliesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | A unique value identifying this part family. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartFamiliesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **partFamily** | [**PartFamily**](PartFamily.md) |  | 

### Return type

[**PartFamily**](PartFamily.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

