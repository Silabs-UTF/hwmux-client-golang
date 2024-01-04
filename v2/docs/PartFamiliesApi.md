# \PartFamiliesAPI

All URIs are relative to *https://hwmux.silabs.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PartFamiliesCreate**](PartFamiliesAPI.md#PartFamiliesCreate) | **Post** /api/part-families/ | 
[**PartFamiliesDestroy**](PartFamiliesAPI.md#PartFamiliesDestroy) | **Delete** /api/part-families/{name}/ | 
[**PartFamiliesList**](PartFamiliesAPI.md#PartFamiliesList) | **Get** /api/part-families/ | 
[**PartFamiliesPartialUpdate**](PartFamiliesAPI.md#PartFamiliesPartialUpdate) | **Patch** /api/part-families/{name}/ | 
[**PartFamiliesPartsCreate**](PartFamiliesAPI.md#PartFamiliesPartsCreate) | **Post** /api/part-families/{part_family_pk}/parts/ | 
[**PartFamiliesPartsDestroy**](PartFamiliesAPI.md#PartFamiliesPartsDestroy) | **Delete** /api/part-families/{part_family_pk}/parts/{id}/ | 
[**PartFamiliesPartsList**](PartFamiliesAPI.md#PartFamiliesPartsList) | **Get** /api/part-families/{part_family_pk}/parts/ | 
[**PartFamiliesPartsPartialUpdate**](PartFamiliesAPI.md#PartFamiliesPartsPartialUpdate) | **Patch** /api/part-families/{part_family_pk}/parts/{id}/ | 
[**PartFamiliesPartsRetrieve**](PartFamiliesAPI.md#PartFamiliesPartsRetrieve) | **Get** /api/part-families/{part_family_pk}/parts/{id}/ | 
[**PartFamiliesPartsUpdate**](PartFamiliesAPI.md#PartFamiliesPartsUpdate) | **Put** /api/part-families/{part_family_pk}/parts/{id}/ | 
[**PartFamiliesRetrieve**](PartFamiliesAPI.md#PartFamiliesRetrieve) | **Get** /api/part-families/{name}/ | 
[**PartFamiliesUpdate**](PartFamiliesAPI.md#PartFamiliesUpdate) | **Put** /api/part-families/{name}/ | 



## PartFamiliesCreate

> PartFamily PartFamiliesCreate(ctx).PartFamily(partFamily).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Silabs-UTF/hwmux-client-golang/v2"
)

func main() {
    partFamily := *openapiclient.NewPartFamily("Name_example") // PartFamily | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartFamiliesAPI.PartFamiliesCreate(context.Background()).PartFamily(partFamily).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartFamiliesAPI.PartFamiliesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PartFamiliesCreate`: PartFamily
    fmt.Fprintf(os.Stdout, "Response from `PartFamiliesAPI.PartFamiliesCreate`: %v\n", resp)
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
    openapiclient "github.com/Silabs-UTF/hwmux-client-golang/v2"
)

func main() {
    name := "name_example" // string | A unique value identifying this part family.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PartFamiliesAPI.PartFamiliesDestroy(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartFamiliesAPI.PartFamiliesDestroy``: %v\n", err)
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
    openapiclient "github.com/Silabs-UTF/hwmux-client-golang/v2"
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
    resp, r, err := apiClient.PartFamiliesAPI.PartFamiliesList(context.Background()).Description(description).Name(name).Ordering(ordering).Page(page).RegexPattern(regexPattern).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartFamiliesAPI.PartFamiliesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PartFamiliesList`: PaginatedPartFamilyList
    fmt.Fprintf(os.Stdout, "Response from `PartFamiliesAPI.PartFamiliesList`: %v\n", resp)
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
    openapiclient "github.com/Silabs-UTF/hwmux-client-golang/v2"
)

func main() {
    name := "name_example" // string | A unique value identifying this part family.
    patchedPartFamily := *openapiclient.NewPatchedPartFamily() // PatchedPartFamily |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartFamiliesAPI.PartFamiliesPartialUpdate(context.Background(), name).PatchedPartFamily(patchedPartFamily).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartFamiliesAPI.PartFamiliesPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PartFamiliesPartialUpdate`: PartFamily
    fmt.Fprintf(os.Stdout, "Response from `PartFamiliesAPI.PartFamiliesPartialUpdate`: %v\n", resp)
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
    openapiclient "github.com/Silabs-UTF/hwmux-client-golang/v2"
)

func main() {
    partFamilyPk := "partFamilyPk_example" // string | 
    part := *openapiclient.NewPart("PartNo_example", *openapiclient.NewPartFamily("Name_example"), "BoardNo_example") // Part | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartFamiliesAPI.PartFamiliesPartsCreate(context.Background(), partFamilyPk).Part(part).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartFamiliesAPI.PartFamiliesPartsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PartFamiliesPartsCreate`: Part
    fmt.Fprintf(os.Stdout, "Response from `PartFamiliesAPI.PartFamiliesPartsCreate`: %v\n", resp)
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
    openapiclient "github.com/Silabs-UTF/hwmux-client-golang/v2"
)

func main() {
    id := "id_example" // string | 
    partFamilyPk := "partFamilyPk_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PartFamiliesAPI.PartFamiliesPartsDestroy(context.Background(), id, partFamilyPk).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartFamiliesAPI.PartFamiliesPartsDestroy``: %v\n", err)
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
    openapiclient "github.com/Silabs-UTF/hwmux-client-golang/v2"
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
    resp, r, err := apiClient.PartFamiliesAPI.PartFamiliesPartsList(context.Background(), partFamilyPk).BoardNo(boardNo).ChipNo(chipNo).Ordering(ordering).Page(page).PartFamily(partFamily).PartNo(partNo).Revision(revision).Search(search).Variant(variant).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartFamiliesAPI.PartFamiliesPartsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PartFamiliesPartsList`: PaginatedPartList
    fmt.Fprintf(os.Stdout, "Response from `PartFamiliesAPI.PartFamiliesPartsList`: %v\n", resp)
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
    openapiclient "github.com/Silabs-UTF/hwmux-client-golang/v2"
)

func main() {
    id := "id_example" // string | 
    partFamilyPk := "partFamilyPk_example" // string | 
    patchedPart := *openapiclient.NewPatchedPart() // PatchedPart |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartFamiliesAPI.PartFamiliesPartsPartialUpdate(context.Background(), id, partFamilyPk).PatchedPart(patchedPart).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartFamiliesAPI.PartFamiliesPartsPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PartFamiliesPartsPartialUpdate`: Part
    fmt.Fprintf(os.Stdout, "Response from `PartFamiliesAPI.PartFamiliesPartsPartialUpdate`: %v\n", resp)
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
    openapiclient "github.com/Silabs-UTF/hwmux-client-golang/v2"
)

func main() {
    id := "id_example" // string | 
    partFamilyPk := "partFamilyPk_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartFamiliesAPI.PartFamiliesPartsRetrieve(context.Background(), id, partFamilyPk).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartFamiliesAPI.PartFamiliesPartsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PartFamiliesPartsRetrieve`: Part
    fmt.Fprintf(os.Stdout, "Response from `PartFamiliesAPI.PartFamiliesPartsRetrieve`: %v\n", resp)
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
    openapiclient "github.com/Silabs-UTF/hwmux-client-golang/v2"
)

func main() {
    id := "id_example" // string | 
    partFamilyPk := "partFamilyPk_example" // string | 
    part := *openapiclient.NewPart("PartNo_example", *openapiclient.NewPartFamily("Name_example"), "BoardNo_example") // Part | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartFamiliesAPI.PartFamiliesPartsUpdate(context.Background(), id, partFamilyPk).Part(part).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartFamiliesAPI.PartFamiliesPartsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PartFamiliesPartsUpdate`: Part
    fmt.Fprintf(os.Stdout, "Response from `PartFamiliesAPI.PartFamiliesPartsUpdate`: %v\n", resp)
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
    openapiclient "github.com/Silabs-UTF/hwmux-client-golang/v2"
)

func main() {
    name := "name_example" // string | A unique value identifying this part family.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartFamiliesAPI.PartFamiliesRetrieve(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartFamiliesAPI.PartFamiliesRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PartFamiliesRetrieve`: PartFamily
    fmt.Fprintf(os.Stdout, "Response from `PartFamiliesAPI.PartFamiliesRetrieve`: %v\n", resp)
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
    openapiclient "github.com/Silabs-UTF/hwmux-client-golang/v2"
)

func main() {
    name := "name_example" // string | A unique value identifying this part family.
    partFamily := *openapiclient.NewPartFamily("Name_example") // PartFamily | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartFamiliesAPI.PartFamiliesUpdate(context.Background(), name).PartFamily(partFamily).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartFamiliesAPI.PartFamiliesUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PartFamiliesUpdate`: PartFamily
    fmt.Fprintf(os.Stdout, "Response from `PartFamiliesAPI.PartFamiliesUpdate`: %v\n", resp)
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

