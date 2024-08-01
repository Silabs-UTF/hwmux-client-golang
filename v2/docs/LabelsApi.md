# \LabelsApi

All URIs are relative to *https://hwmux.silabs.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LabelsCreate**](LabelsApi.md#LabelsCreate) | **Post** /api/labels/ | 
[**LabelsDestroy**](LabelsApi.md#LabelsDestroy) | **Delete** /api/labels/{id}/ | 
[**LabelsList**](LabelsApi.md#LabelsList) | **Get** /api/labels/ | 
[**LabelsPartialUpdate**](LabelsApi.md#LabelsPartialUpdate) | **Patch** /api/labels/{id}/ | 
[**LabelsPermissionsPartialUpdate**](LabelsApi.md#LabelsPermissionsPartialUpdate) | **Patch** /api/labels/{id}/permissions/ | 
[**LabelsPermissionsRetrieve**](LabelsApi.md#LabelsPermissionsRetrieve) | **Get** /api/labels/{id}/permissions/ | 
[**LabelsPermissionsUpdate**](LabelsApi.md#LabelsPermissionsUpdate) | **Put** /api/labels/{id}/permissions/ | 
[**LabelsReserveAnyUpdate**](LabelsApi.md#LabelsReserveAnyUpdate) | **Put** /api/labels/{id}/reserve_any/ | 
[**LabelsRetrieve**](LabelsApi.md#LabelsRetrieve) | **Get** /api/labels/{id}/ | 
[**LabelsUpdate**](LabelsApi.md#LabelsUpdate) | **Put** /api/labels/{id}/ | 



## LabelsCreate

> LabelSerializerWithPermissions LabelsCreate(ctx).LabelSerializerWithPermissions(labelSerializerWithPermissions).Execute()





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
    labelSerializerWithPermissions := *openapiclient.NewLabelSerializerWithPermissions(int32(123), []int32{int32(123)}, "Name_example") // LabelSerializerWithPermissions | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LabelsApi.LabelsCreate(context.Background()).LabelSerializerWithPermissions(labelSerializerWithPermissions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabelsApi.LabelsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LabelsCreate`: LabelSerializerWithPermissions
    fmt.Fprintf(os.Stdout, "Response from `LabelsApi.LabelsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLabelsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **labelSerializerWithPermissions** | [**LabelSerializerWithPermissions**](LabelSerializerWithPermissions.md) |  | 

### Return type

[**LabelSerializerWithPermissions**](LabelSerializerWithPermissions.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LabelsDestroy

> LabelsDestroy(ctx, id).Execute()





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
    id := int32(56) // int32 | A unique integer value identifying this device group label.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LabelsApi.LabelsDestroy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabelsApi.LabelsDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this device group label. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLabelsDestroyRequest struct via the builder pattern


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


## LabelsList

> PaginatedLabelList LabelsList(ctx).DeviceGroups(deviceGroups).ForceMv(forceMv).IncludePermissionGroups(includePermissionGroups).Name(name).Ordering(ordering).Page(page).Search(search).Source(source).Execute()





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
    deviceGroups := []int32{int32(123)} // []int32 |  (optional)
    forceMv := "forceMv_example" // string | Specifies the behavior of the force_mv attribute. DEFAULT : Default condition which retrieve data from the Materialized view unless ax exception occurs where we fall back to postgres. ON : Force the data to be retrieved from the Materialized view. OFF : Force the data to be retrieved from the Postgres database (optional) (default to "DEFAULT")
    includePermissionGroups := true // bool | If set to true, the permission groups associated with this resource will be included in the response. Defaults to false. (optional) (default to false)
    name := "name_example" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    search := "search_example" // string | A search term. (optional)
    source := "source_example" // string | * `TERRAFORM` - Terraform * `XML` - XML * `UI` - UI * `OTHER` - Other (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LabelsApi.LabelsList(context.Background()).DeviceGroups(deviceGroups).ForceMv(forceMv).IncludePermissionGroups(includePermissionGroups).Name(name).Ordering(ordering).Page(page).Search(search).Source(source).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabelsApi.LabelsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LabelsList`: PaginatedLabelList
    fmt.Fprintf(os.Stdout, "Response from `LabelsApi.LabelsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLabelsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceGroups** | **[]int32** |  | 
 **forceMv** | **string** | Specifies the behavior of the force_mv attribute. DEFAULT : Default condition which retrieve data from the Materialized view unless ax exception occurs where we fall back to postgres. ON : Force the data to be retrieved from the Materialized view. OFF : Force the data to be retrieved from the Postgres database | [default to &quot;DEFAULT&quot;]
 **includePermissionGroups** | **bool** | If set to true, the permission groups associated with this resource will be included in the response. Defaults to false. | [default to false]
 **name** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **search** | **string** | A search term. | 
 **source** | **string** | * &#x60;TERRAFORM&#x60; - Terraform * &#x60;XML&#x60; - XML * &#x60;UI&#x60; - UI * &#x60;OTHER&#x60; - Other | 

### Return type

[**PaginatedLabelList**](PaginatedLabelList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LabelsPartialUpdate

> LabelSerializerWithPermissions LabelsPartialUpdate(ctx, id).PatchedLabelSerializerWithPermissions(patchedLabelSerializerWithPermissions).Execute()





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
    id := int32(56) // int32 | A unique integer value identifying this device group label.
    patchedLabelSerializerWithPermissions := *openapiclient.NewPatchedLabelSerializerWithPermissions() // PatchedLabelSerializerWithPermissions |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LabelsApi.LabelsPartialUpdate(context.Background(), id).PatchedLabelSerializerWithPermissions(patchedLabelSerializerWithPermissions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabelsApi.LabelsPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LabelsPartialUpdate`: LabelSerializerWithPermissions
    fmt.Fprintf(os.Stdout, "Response from `LabelsApi.LabelsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this device group label. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLabelsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedLabelSerializerWithPermissions** | [**PatchedLabelSerializerWithPermissions**](PatchedLabelSerializerWithPermissions.md) |  | 

### Return type

[**LabelSerializerWithPermissions**](LabelSerializerWithPermissions.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LabelsPermissionsPartialUpdate

> ObjectPermissions LabelsPermissionsPartialUpdate(ctx, id).PatchedObjectPermissions(patchedObjectPermissions).Execute()





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
    id := int32(56) // int32 | A unique integer value identifying this device group label.
    patchedObjectPermissions := *openapiclient.NewPatchedObjectPermissions() // PatchedObjectPermissions |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LabelsApi.LabelsPermissionsPartialUpdate(context.Background(), id).PatchedObjectPermissions(patchedObjectPermissions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabelsApi.LabelsPermissionsPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LabelsPermissionsPartialUpdate`: ObjectPermissions
    fmt.Fprintf(os.Stdout, "Response from `LabelsApi.LabelsPermissionsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this device group label. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLabelsPermissionsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedObjectPermissions** | [**PatchedObjectPermissions**](PatchedObjectPermissions.md) |  | 

### Return type

[**ObjectPermissions**](ObjectPermissions.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LabelsPermissionsRetrieve

> ObjectPermissions LabelsPermissionsRetrieve(ctx, id).Execute()





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
    id := int32(56) // int32 | A unique integer value identifying this device group label.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LabelsApi.LabelsPermissionsRetrieve(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabelsApi.LabelsPermissionsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LabelsPermissionsRetrieve`: ObjectPermissions
    fmt.Fprintf(os.Stdout, "Response from `LabelsApi.LabelsPermissionsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this device group label. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLabelsPermissionsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ObjectPermissions**](ObjectPermissions.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LabelsPermissionsUpdate

> ObjectPermissions LabelsPermissionsUpdate(ctx, id).ObjectPermissions(objectPermissions).Execute()





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
    id := int32(56) // int32 | A unique integer value identifying this device group label.
    objectPermissions := *openapiclient.NewObjectPermissions(map[string]interface{}{"key": interface{}(123)}, map[string]interface{}{"key": interface{}(123)}) // ObjectPermissions | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LabelsApi.LabelsPermissionsUpdate(context.Background(), id).ObjectPermissions(objectPermissions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabelsApi.LabelsPermissionsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LabelsPermissionsUpdate`: ObjectPermissions
    fmt.Fprintf(os.Stdout, "Response from `LabelsApi.LabelsPermissionsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this device group label. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLabelsPermissionsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **objectPermissions** | [**ObjectPermissions**](ObjectPermissions.md) |  | 

### Return type

[**ObjectPermissions**](ObjectPermissions.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LabelsReserveAnyUpdate

> DeviceGroup LabelsReserveAnyUpdate(ctx, id).Details(details).Execute()





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
    id := int32(56) // int32 | A unique integer value identifying this device group label.
    details := "details_example" // string | Additional information such as the jenkins job URL (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LabelsApi.LabelsReserveAnyUpdate(context.Background(), id).Details(details).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabelsApi.LabelsReserveAnyUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LabelsReserveAnyUpdate`: DeviceGroup
    fmt.Fprintf(os.Stdout, "Response from `LabelsApi.LabelsReserveAnyUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this device group label. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLabelsReserveAnyUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **details** | **string** | Additional information such as the jenkins job URL | 

### Return type

[**DeviceGroup**](DeviceGroup.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LabelsRetrieve

> Label LabelsRetrieve(ctx, id).IncludePermissionGroups(includePermissionGroups).Execute()





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
    id := int32(56) // int32 | A unique integer value identifying this device group label.
    includePermissionGroups := true // bool | If set to true, the permission groups associated with this resource will be included in the response. Defaults to false. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LabelsApi.LabelsRetrieve(context.Background(), id).IncludePermissionGroups(includePermissionGroups).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabelsApi.LabelsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LabelsRetrieve`: Label
    fmt.Fprintf(os.Stdout, "Response from `LabelsApi.LabelsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this device group label. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLabelsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includePermissionGroups** | **bool** | If set to true, the permission groups associated with this resource will be included in the response. Defaults to false. | [default to false]

### Return type

[**Label**](Label.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LabelsUpdate

> LabelSerializerWithPermissions LabelsUpdate(ctx, id).LabelSerializerWithPermissions(labelSerializerWithPermissions).Execute()





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
    id := int32(56) // int32 | A unique integer value identifying this device group label.
    labelSerializerWithPermissions := *openapiclient.NewLabelSerializerWithPermissions(int32(123), []int32{int32(123)}, "Name_example") // LabelSerializerWithPermissions | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LabelsApi.LabelsUpdate(context.Background(), id).LabelSerializerWithPermissions(labelSerializerWithPermissions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabelsApi.LabelsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LabelsUpdate`: LabelSerializerWithPermissions
    fmt.Fprintf(os.Stdout, "Response from `LabelsApi.LabelsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this device group label. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLabelsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **labelSerializerWithPermissions** | [**LabelSerializerWithPermissions**](LabelSerializerWithPermissions.md) |  | 

### Return type

[**LabelSerializerWithPermissions**](LabelSerializerWithPermissions.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

