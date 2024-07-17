# \SitesApi

All URIs are relative to *https://hwmux.silabs.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SitesCreate**](SitesApi.md#SitesCreate) | **Post** /api/sites/ | 
[**SitesDestroy**](SitesApi.md#SitesDestroy) | **Delete** /api/sites/{name}/ | 
[**SitesList**](SitesApi.md#SitesList) | **Get** /api/sites/ | 
[**SitesPartialUpdate**](SitesApi.md#SitesPartialUpdate) | **Patch** /api/sites/{name}/ | 
[**SitesRetrieve**](SitesApi.md#SitesRetrieve) | **Get** /api/sites/{name}/ | 
[**SitesRoomsCreate**](SitesApi.md#SitesRoomsCreate) | **Post** /api/sites/{site_pk}/rooms/ | 
[**SitesRoomsDestroy**](SitesApi.md#SitesRoomsDestroy) | **Delete** /api/sites/{site_pk}/rooms/{id}/ | 
[**SitesRoomsList**](SitesApi.md#SitesRoomsList) | **Get** /api/sites/{site_pk}/rooms/ | 
[**SitesRoomsPartialUpdate**](SitesApi.md#SitesRoomsPartialUpdate) | **Patch** /api/sites/{site_pk}/rooms/{id}/ | 
[**SitesRoomsRetrieve**](SitesApi.md#SitesRoomsRetrieve) | **Get** /api/sites/{site_pk}/rooms/{id}/ | 
[**SitesRoomsUpdate**](SitesApi.md#SitesRoomsUpdate) | **Put** /api/sites/{site_pk}/rooms/{id}/ | 
[**SitesUpdate**](SitesApi.md#SitesUpdate) | **Put** /api/sites/{name}/ | 



## SitesCreate

> Site SitesCreate(ctx).Site(site).Execute()



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
    site := *openapiclient.NewSite("Name_example", float64(123), float64(123)) // Site | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesApi.SitesCreate(context.Background()).Site(site).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesApi.SitesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesCreate`: Site
    fmt.Fprintf(os.Stdout, "Response from `SitesApi.SitesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSitesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **site** | [**Site**](Site.md) |  | 

### Return type

[**Site**](Site.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesDestroy

> SitesDestroy(ctx, name).Execute()



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
    name := "name_example" // string | A unique value identifying this site.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesApi.SitesDestroy(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesApi.SitesDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | A unique value identifying this site. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesDestroyRequest struct via the builder pattern


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


## SitesList

> PaginatedSiteList SitesList(ctx).Ordering(ordering).Page(page).Search(search).Execute()



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
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    search := "search_example" // string | A search term. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesApi.SitesList(context.Background()).Ordering(ordering).Page(page).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesApi.SitesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesList`: PaginatedSiteList
    fmt.Fprintf(os.Stdout, "Response from `SitesApi.SitesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSitesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedSiteList**](PaginatedSiteList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesPartialUpdate

> Site SitesPartialUpdate(ctx, name).PatchedSite(patchedSite).Execute()



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
    name := "name_example" // string | A unique value identifying this site.
    patchedSite := *openapiclient.NewPatchedSite() // PatchedSite |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesApi.SitesPartialUpdate(context.Background(), name).PatchedSite(patchedSite).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesApi.SitesPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesPartialUpdate`: Site
    fmt.Fprintf(os.Stdout, "Response from `SitesApi.SitesPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | A unique value identifying this site. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedSite** | [**PatchedSite**](PatchedSite.md) |  | 

### Return type

[**Site**](Site.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesRetrieve

> Site SitesRetrieve(ctx, name).Execute()



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
    name := "name_example" // string | A unique value identifying this site.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesApi.SitesRetrieve(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesApi.SitesRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesRetrieve`: Site
    fmt.Fprintf(os.Stdout, "Response from `SitesApi.SitesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | A unique value identifying this site. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Site**](Site.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesRoomsCreate

> Room SitesRoomsCreate(ctx, sitePk).Room(room).Execute()



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
    sitePk := "sitePk_example" // string | 
    room := *openapiclient.NewRoom("Name_example", "Site_example") // Room | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesApi.SitesRoomsCreate(context.Background(), sitePk).Room(room).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesApi.SitesRoomsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesRoomsCreate`: Room
    fmt.Fprintf(os.Stdout, "Response from `SitesApi.SitesRoomsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sitePk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesRoomsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **room** | [**Room**](Room.md) |  | 

### Return type

[**Room**](Room.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesRoomsDestroy

> SitesRoomsDestroy(ctx, id, sitePk).Execute()



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
    sitePk := "sitePk_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesApi.SitesRoomsDestroy(context.Background(), id, sitePk).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesApi.SitesRoomsDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**sitePk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesRoomsDestroyRequest struct via the builder pattern


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


## SitesRoomsList

> PaginatedRoomList SitesRoomsList(ctx, sitePk).Description(description).Name(name).Ordering(ordering).Page(page).Search(search).Site(site).Execute()



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
    sitePk := "sitePk_example" // string | 
    description := "description_example" // string |  (optional)
    name := "name_example" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    search := "search_example" // string | A search term. (optional)
    site := "site_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesApi.SitesRoomsList(context.Background(), sitePk).Description(description).Name(name).Ordering(ordering).Page(page).Search(search).Site(site).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesApi.SitesRoomsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesRoomsList`: PaginatedRoomList
    fmt.Fprintf(os.Stdout, "Response from `SitesApi.SitesRoomsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sitePk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesRoomsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **description** | **string** |  | 
 **name** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **search** | **string** | A search term. | 
 **site** | **string** |  | 

### Return type

[**PaginatedRoomList**](PaginatedRoomList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesRoomsPartialUpdate

> Room SitesRoomsPartialUpdate(ctx, id, sitePk).PatchedRoom(patchedRoom).Execute()



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
    sitePk := "sitePk_example" // string | 
    patchedRoom := *openapiclient.NewPatchedRoom() // PatchedRoom |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesApi.SitesRoomsPartialUpdate(context.Background(), id, sitePk).PatchedRoom(patchedRoom).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesApi.SitesRoomsPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesRoomsPartialUpdate`: Room
    fmt.Fprintf(os.Stdout, "Response from `SitesApi.SitesRoomsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**sitePk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesRoomsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedRoom** | [**PatchedRoom**](PatchedRoom.md) |  | 

### Return type

[**Room**](Room.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesRoomsRetrieve

> Room SitesRoomsRetrieve(ctx, id, sitePk).Execute()



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
    sitePk := "sitePk_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesApi.SitesRoomsRetrieve(context.Background(), id, sitePk).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesApi.SitesRoomsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesRoomsRetrieve`: Room
    fmt.Fprintf(os.Stdout, "Response from `SitesApi.SitesRoomsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**sitePk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesRoomsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Room**](Room.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesRoomsUpdate

> Room SitesRoomsUpdate(ctx, id, sitePk).Room(room).Execute()



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
    sitePk := "sitePk_example" // string | 
    room := *openapiclient.NewRoom("Name_example", "Site_example") // Room | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesApi.SitesRoomsUpdate(context.Background(), id, sitePk).Room(room).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesApi.SitesRoomsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesRoomsUpdate`: Room
    fmt.Fprintf(os.Stdout, "Response from `SitesApi.SitesRoomsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**sitePk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesRoomsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **room** | [**Room**](Room.md) |  | 

### Return type

[**Room**](Room.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesUpdate

> Site SitesUpdate(ctx, name).Site(site).Execute()



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
    name := "name_example" // string | A unique value identifying this site.
    site := *openapiclient.NewSite("Name_example", float64(123), float64(123)) // Site | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesApi.SitesUpdate(context.Background(), name).Site(site).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesApi.SitesUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesUpdate`: Site
    fmt.Fprintf(os.Stdout, "Response from `SitesApi.SitesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | A unique value identifying this site. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **site** | [**Site**](Site.md) |  | 

### Return type

[**Site**](Site.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

