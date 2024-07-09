# \PermissionsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PermissionsGroupsCreate**](PermissionsApi.md#PermissionsGroupsCreate) | **Post** /api/permissions/groups/ | 
[**PermissionsGroupsDestroy**](PermissionsApi.md#PermissionsGroupsDestroy) | **Delete** /api/permissions/groups/{name_or_id}/ | 
[**PermissionsGroupsDeviceGroupsCreate**](PermissionsApi.md#PermissionsGroupsDeviceGroupsCreate) | **Post** /api/permissions/groups/{name_or_id}/device_groups/ | 
[**PermissionsGroupsDeviceGroupsDestroy**](PermissionsApi.md#PermissionsGroupsDeviceGroupsDestroy) | **Delete** /api/permissions/groups/{name_or_id}/device_groups/{id}/ | 
[**PermissionsGroupsDeviceGroupsList**](PermissionsApi.md#PermissionsGroupsDeviceGroupsList) | **Get** /api/permissions/groups/{name_or_id}/device_groups/ | 
[**PermissionsGroupsDeviceGroupsPartialUpdate**](PermissionsApi.md#PermissionsGroupsDeviceGroupsPartialUpdate) | **Patch** /api/permissions/groups/{name_or_id}/device_groups/{id}/ | 
[**PermissionsGroupsDeviceGroupsUpdate**](PermissionsApi.md#PermissionsGroupsDeviceGroupsUpdate) | **Put** /api/permissions/groups/{name_or_id}/device_groups/{id}/ | 
[**PermissionsGroupsDevicesCreate**](PermissionsApi.md#PermissionsGroupsDevicesCreate) | **Post** /api/permissions/groups/{name_or_id}/devices/ | 
[**PermissionsGroupsDevicesDestroy**](PermissionsApi.md#PermissionsGroupsDevicesDestroy) | **Delete** /api/permissions/groups/{name_or_id}/devices/{id}/ | 
[**PermissionsGroupsDevicesList**](PermissionsApi.md#PermissionsGroupsDevicesList) | **Get** /api/permissions/groups/{name_or_id}/devices/ | 
[**PermissionsGroupsDevicesPartialUpdate**](PermissionsApi.md#PermissionsGroupsDevicesPartialUpdate) | **Patch** /api/permissions/groups/{name_or_id}/devices/{id}/ | 
[**PermissionsGroupsDevicesUpdate**](PermissionsApi.md#PermissionsGroupsDevicesUpdate) | **Put** /api/permissions/groups/{name_or_id}/devices/{id}/ | 
[**PermissionsGroupsLabelsCreate**](PermissionsApi.md#PermissionsGroupsLabelsCreate) | **Post** /api/permissions/groups/{name_or_id}/labels/ | 
[**PermissionsGroupsLabelsDestroy**](PermissionsApi.md#PermissionsGroupsLabelsDestroy) | **Delete** /api/permissions/groups/{name_or_id}/labels/{id}/ | 
[**PermissionsGroupsLabelsList**](PermissionsApi.md#PermissionsGroupsLabelsList) | **Get** /api/permissions/groups/{name_or_id}/labels/ | 
[**PermissionsGroupsLabelsPartialUpdate**](PermissionsApi.md#PermissionsGroupsLabelsPartialUpdate) | **Patch** /api/permissions/groups/{name_or_id}/labels/{id}/ | 
[**PermissionsGroupsLabelsUpdate**](PermissionsApi.md#PermissionsGroupsLabelsUpdate) | **Put** /api/permissions/groups/{name_or_id}/labels/{id}/ | 
[**PermissionsGroupsList**](PermissionsApi.md#PermissionsGroupsList) | **Get** /api/permissions/groups/ | 
[**PermissionsGroupsPartialUpdate**](PermissionsApi.md#PermissionsGroupsPartialUpdate) | **Patch** /api/permissions/groups/{name_or_id}/ | 
[**PermissionsGroupsRetrieve**](PermissionsApi.md#PermissionsGroupsRetrieve) | **Get** /api/permissions/groups/{name_or_id}/ | 
[**PermissionsGroupsUpdate**](PermissionsApi.md#PermissionsGroupsUpdate) | **Put** /api/permissions/groups/{name_or_id}/ | 
[**PermissionsGroupsUsersCreate**](PermissionsApi.md#PermissionsGroupsUsersCreate) | **Post** /api/permissions/groups/{name_or_id}/users/ | 
[**PermissionsGroupsUsersDestroy**](PermissionsApi.md#PermissionsGroupsUsersDestroy) | **Delete** /api/permissions/groups/{name_or_id}/users/{username_or_id}/ | 



## PermissionsGroupsCreate

> PermissionGroup PermissionsGroupsCreate(ctx).PermissionGroup(permissionGroup).Execute()



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
    permissionGroup := *openapiclient.NewPermissionGroup(int32(123), []string{"Permissions_example"}, "Name_example") // PermissionGroup | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PermissionsApi.PermissionsGroupsCreate(context.Background()).PermissionGroup(permissionGroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionsApi.PermissionsGroupsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PermissionsGroupsCreate`: PermissionGroup
    fmt.Fprintf(os.Stdout, "Response from `PermissionsApi.PermissionsGroupsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPermissionsGroupsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **permissionGroup** | [**PermissionGroup**](PermissionGroup.md) |  | 

### Return type

[**PermissionGroup**](PermissionGroup.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PermissionsGroupsDestroy

> PermissionsGroupsDestroy(ctx, nameOrId).Execute()



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
    nameOrId := "nameOrId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PermissionsApi.PermissionsGroupsDestroy(context.Background(), nameOrId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionsApi.PermissionsGroupsDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPermissionsGroupsDestroyRequest struct via the builder pattern


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


## PermissionsGroupsDeviceGroupsCreate

> PermissionGroup PermissionsGroupsDeviceGroupsCreate(ctx, nameOrId).ResourcePermissions(resourcePermissions).Execute()



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
    nameOrId := "nameOrId_example" // string | 
    resourcePermissions := []openapiclient.ResourcePermissions{*openapiclient.NewResourcePermissions(int32(123))} // []ResourcePermissions | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PermissionsApi.PermissionsGroupsDeviceGroupsCreate(context.Background(), nameOrId).ResourcePermissions(resourcePermissions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionsApi.PermissionsGroupsDeviceGroupsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PermissionsGroupsDeviceGroupsCreate`: PermissionGroup
    fmt.Fprintf(os.Stdout, "Response from `PermissionsApi.PermissionsGroupsDeviceGroupsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPermissionsGroupsDeviceGroupsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resourcePermissions** | [**[]ResourcePermissions**](ResourcePermissions.md) |  | 

### Return type

[**PermissionGroup**](PermissionGroup.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PermissionsGroupsDeviceGroupsDestroy

> PermissionsGroupsDeviceGroupsDestroy(ctx, id, nameOrId).Execute()



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
    id := int32(56) // int32 | A unique integer value identifying this device group.
    nameOrId := "nameOrId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PermissionsApi.PermissionsGroupsDeviceGroupsDestroy(context.Background(), id, nameOrId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionsApi.PermissionsGroupsDeviceGroupsDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this device group. | 
**nameOrId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPermissionsGroupsDeviceGroupsDestroyRequest struct via the builder pattern


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


## PermissionsGroupsDeviceGroupsList

> PaginatedDeviceGroupList PermissionsGroupsDeviceGroupsList(ctx, nameOrId).Page(page).Perms(perms).Execute()



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
    nameOrId := "nameOrId_example" // string | 
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    perms := "view" // string | Permission types. Only entities with all permissions defined here will be fetched. Defaults to all permissions. (optional) (default to "view,change,add,delete")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PermissionsApi.PermissionsGroupsDeviceGroupsList(context.Background(), nameOrId).Page(page).Perms(perms).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionsApi.PermissionsGroupsDeviceGroupsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PermissionsGroupsDeviceGroupsList`: PaginatedDeviceGroupList
    fmt.Fprintf(os.Stdout, "Response from `PermissionsApi.PermissionsGroupsDeviceGroupsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPermissionsGroupsDeviceGroupsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | A page number within the paginated result set. | 
 **perms** | **string** | Permission types. Only entities with all permissions defined here will be fetched. Defaults to all permissions. | [default to &quot;view,change,add,delete&quot;]

### Return type

[**PaginatedDeviceGroupList**](PaginatedDeviceGroupList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PermissionsGroupsDeviceGroupsPartialUpdate

> ResourcePermissions PermissionsGroupsDeviceGroupsPartialUpdate(ctx, id, nameOrId).PatchedResourcePermissions(patchedResourcePermissions).Execute()



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
    id := int32(56) // int32 | A unique integer value identifying this device group.
    nameOrId := "nameOrId_example" // string | 
    patchedResourcePermissions := *openapiclient.NewPatchedResourcePermissions() // PatchedResourcePermissions |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PermissionsApi.PermissionsGroupsDeviceGroupsPartialUpdate(context.Background(), id, nameOrId).PatchedResourcePermissions(patchedResourcePermissions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionsApi.PermissionsGroupsDeviceGroupsPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PermissionsGroupsDeviceGroupsPartialUpdate`: ResourcePermissions
    fmt.Fprintf(os.Stdout, "Response from `PermissionsApi.PermissionsGroupsDeviceGroupsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this device group. | 
**nameOrId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPermissionsGroupsDeviceGroupsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedResourcePermissions** | [**PatchedResourcePermissions**](PatchedResourcePermissions.md) |  | 

### Return type

[**ResourcePermissions**](ResourcePermissions.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PermissionsGroupsDeviceGroupsUpdate

> PermissionGroup PermissionsGroupsDeviceGroupsUpdate(ctx, id, nameOrId).ResourcePermissions(resourcePermissions).Execute()



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
    id := int32(56) // int32 | A unique integer value identifying this device group.
    nameOrId := "nameOrId_example" // string | 
    resourcePermissions := *openapiclient.NewResourcePermissions(int32(123)) // ResourcePermissions | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PermissionsApi.PermissionsGroupsDeviceGroupsUpdate(context.Background(), id, nameOrId).ResourcePermissions(resourcePermissions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionsApi.PermissionsGroupsDeviceGroupsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PermissionsGroupsDeviceGroupsUpdate`: PermissionGroup
    fmt.Fprintf(os.Stdout, "Response from `PermissionsApi.PermissionsGroupsDeviceGroupsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this device group. | 
**nameOrId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPermissionsGroupsDeviceGroupsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **resourcePermissions** | [**ResourcePermissions**](ResourcePermissions.md) |  | 

### Return type

[**PermissionGroup**](PermissionGroup.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PermissionsGroupsDevicesCreate

> PermissionGroup PermissionsGroupsDevicesCreate(ctx, nameOrId).ResourcePermissions(resourcePermissions).Execute()



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
    nameOrId := "nameOrId_example" // string | 
    resourcePermissions := []openapiclient.ResourcePermissions{*openapiclient.NewResourcePermissions(int32(123))} // []ResourcePermissions | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PermissionsApi.PermissionsGroupsDevicesCreate(context.Background(), nameOrId).ResourcePermissions(resourcePermissions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionsApi.PermissionsGroupsDevicesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PermissionsGroupsDevicesCreate`: PermissionGroup
    fmt.Fprintf(os.Stdout, "Response from `PermissionsApi.PermissionsGroupsDevicesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPermissionsGroupsDevicesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resourcePermissions** | [**[]ResourcePermissions**](ResourcePermissions.md) |  | 

### Return type

[**PermissionGroup**](PermissionGroup.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PermissionsGroupsDevicesDestroy

> PermissionsGroupsDevicesDestroy(ctx, id, nameOrId).Execute()



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
    id := int32(56) // int32 | A unique integer value identifying this device.
    nameOrId := "nameOrId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PermissionsApi.PermissionsGroupsDevicesDestroy(context.Background(), id, nameOrId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionsApi.PermissionsGroupsDevicesDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this device. | 
**nameOrId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPermissionsGroupsDevicesDestroyRequest struct via the builder pattern


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


## PermissionsGroupsDevicesList

> PaginatedDeviceSerializerPublicList PermissionsGroupsDevicesList(ctx, nameOrId).Page(page).Perms(perms).Execute()



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
    nameOrId := "nameOrId_example" // string | 
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    perms := "view" // string | Permission types. Only entities with all permissions defined here will be fetched. Defaults to all permissions. (optional) (default to "view,change,add,delete")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PermissionsApi.PermissionsGroupsDevicesList(context.Background(), nameOrId).Page(page).Perms(perms).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionsApi.PermissionsGroupsDevicesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PermissionsGroupsDevicesList`: PaginatedDeviceSerializerPublicList
    fmt.Fprintf(os.Stdout, "Response from `PermissionsApi.PermissionsGroupsDevicesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPermissionsGroupsDevicesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | A page number within the paginated result set. | 
 **perms** | **string** | Permission types. Only entities with all permissions defined here will be fetched. Defaults to all permissions. | [default to &quot;view,change,add,delete&quot;]

### Return type

[**PaginatedDeviceSerializerPublicList**](PaginatedDeviceSerializerPublicList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PermissionsGroupsDevicesPartialUpdate

> ResourcePermissions PermissionsGroupsDevicesPartialUpdate(ctx, id, nameOrId).PatchedResourcePermissions(patchedResourcePermissions).Execute()



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
    id := int32(56) // int32 | A unique integer value identifying this device.
    nameOrId := "nameOrId_example" // string | 
    patchedResourcePermissions := *openapiclient.NewPatchedResourcePermissions() // PatchedResourcePermissions |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PermissionsApi.PermissionsGroupsDevicesPartialUpdate(context.Background(), id, nameOrId).PatchedResourcePermissions(patchedResourcePermissions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionsApi.PermissionsGroupsDevicesPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PermissionsGroupsDevicesPartialUpdate`: ResourcePermissions
    fmt.Fprintf(os.Stdout, "Response from `PermissionsApi.PermissionsGroupsDevicesPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this device. | 
**nameOrId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPermissionsGroupsDevicesPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedResourcePermissions** | [**PatchedResourcePermissions**](PatchedResourcePermissions.md) |  | 

### Return type

[**ResourcePermissions**](ResourcePermissions.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PermissionsGroupsDevicesUpdate

> PermissionGroup PermissionsGroupsDevicesUpdate(ctx, id, nameOrId).ResourcePermissions(resourcePermissions).Execute()



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
    id := int32(56) // int32 | A unique integer value identifying this device.
    nameOrId := "nameOrId_example" // string | 
    resourcePermissions := *openapiclient.NewResourcePermissions(int32(123)) // ResourcePermissions | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PermissionsApi.PermissionsGroupsDevicesUpdate(context.Background(), id, nameOrId).ResourcePermissions(resourcePermissions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionsApi.PermissionsGroupsDevicesUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PermissionsGroupsDevicesUpdate`: PermissionGroup
    fmt.Fprintf(os.Stdout, "Response from `PermissionsApi.PermissionsGroupsDevicesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this device. | 
**nameOrId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPermissionsGroupsDevicesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **resourcePermissions** | [**ResourcePermissions**](ResourcePermissions.md) |  | 

### Return type

[**PermissionGroup**](PermissionGroup.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PermissionsGroupsLabelsCreate

> PermissionGroup PermissionsGroupsLabelsCreate(ctx, nameOrId).ResourcePermissions(resourcePermissions).Execute()



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
    nameOrId := "nameOrId_example" // string | 
    resourcePermissions := []openapiclient.ResourcePermissions{*openapiclient.NewResourcePermissions(int32(123))} // []ResourcePermissions | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PermissionsApi.PermissionsGroupsLabelsCreate(context.Background(), nameOrId).ResourcePermissions(resourcePermissions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionsApi.PermissionsGroupsLabelsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PermissionsGroupsLabelsCreate`: PermissionGroup
    fmt.Fprintf(os.Stdout, "Response from `PermissionsApi.PermissionsGroupsLabelsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPermissionsGroupsLabelsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resourcePermissions** | [**[]ResourcePermissions**](ResourcePermissions.md) |  | 

### Return type

[**PermissionGroup**](PermissionGroup.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PermissionsGroupsLabelsDestroy

> PermissionsGroupsLabelsDestroy(ctx, id, nameOrId).Execute()



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
    nameOrId := "nameOrId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PermissionsApi.PermissionsGroupsLabelsDestroy(context.Background(), id, nameOrId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionsApi.PermissionsGroupsLabelsDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this device group label. | 
**nameOrId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPermissionsGroupsLabelsDestroyRequest struct via the builder pattern


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


## PermissionsGroupsLabelsList

> PaginatedLabelList PermissionsGroupsLabelsList(ctx, nameOrId).Page(page).Perms(perms).Execute()



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
    nameOrId := "nameOrId_example" // string | 
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    perms := "view" // string | Permission types. Only entities with all permissions defined here will be fetched. Defaults to all permissions. (optional) (default to "view,change,add,delete")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PermissionsApi.PermissionsGroupsLabelsList(context.Background(), nameOrId).Page(page).Perms(perms).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionsApi.PermissionsGroupsLabelsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PermissionsGroupsLabelsList`: PaginatedLabelList
    fmt.Fprintf(os.Stdout, "Response from `PermissionsApi.PermissionsGroupsLabelsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPermissionsGroupsLabelsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | A page number within the paginated result set. | 
 **perms** | **string** | Permission types. Only entities with all permissions defined here will be fetched. Defaults to all permissions. | [default to &quot;view,change,add,delete&quot;]

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


## PermissionsGroupsLabelsPartialUpdate

> ResourcePermissions PermissionsGroupsLabelsPartialUpdate(ctx, id, nameOrId).PatchedResourcePermissions(patchedResourcePermissions).Execute()



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
    nameOrId := "nameOrId_example" // string | 
    patchedResourcePermissions := *openapiclient.NewPatchedResourcePermissions() // PatchedResourcePermissions |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PermissionsApi.PermissionsGroupsLabelsPartialUpdate(context.Background(), id, nameOrId).PatchedResourcePermissions(patchedResourcePermissions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionsApi.PermissionsGroupsLabelsPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PermissionsGroupsLabelsPartialUpdate`: ResourcePermissions
    fmt.Fprintf(os.Stdout, "Response from `PermissionsApi.PermissionsGroupsLabelsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this device group label. | 
**nameOrId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPermissionsGroupsLabelsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedResourcePermissions** | [**PatchedResourcePermissions**](PatchedResourcePermissions.md) |  | 

### Return type

[**ResourcePermissions**](ResourcePermissions.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PermissionsGroupsLabelsUpdate

> PermissionGroup PermissionsGroupsLabelsUpdate(ctx, id, nameOrId).ResourcePermissions(resourcePermissions).Execute()



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
    nameOrId := "nameOrId_example" // string | 
    resourcePermissions := *openapiclient.NewResourcePermissions(int32(123)) // ResourcePermissions | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PermissionsApi.PermissionsGroupsLabelsUpdate(context.Background(), id, nameOrId).ResourcePermissions(resourcePermissions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionsApi.PermissionsGroupsLabelsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PermissionsGroupsLabelsUpdate`: PermissionGroup
    fmt.Fprintf(os.Stdout, "Response from `PermissionsApi.PermissionsGroupsLabelsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this device group label. | 
**nameOrId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPermissionsGroupsLabelsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **resourcePermissions** | [**ResourcePermissions**](ResourcePermissions.md) |  | 

### Return type

[**PermissionGroup**](PermissionGroup.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PermissionsGroupsList

> PaginatedPermissionGroupList PermissionsGroupsList(ctx).Ordering(ordering).Page(page).Search(search).Execute()



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
    resp, r, err := apiClient.PermissionsApi.PermissionsGroupsList(context.Background()).Ordering(ordering).Page(page).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionsApi.PermissionsGroupsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PermissionsGroupsList`: PaginatedPermissionGroupList
    fmt.Fprintf(os.Stdout, "Response from `PermissionsApi.PermissionsGroupsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPermissionsGroupsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedPermissionGroupList**](PaginatedPermissionGroupList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PermissionsGroupsPartialUpdate

> PermissionGroup PermissionsGroupsPartialUpdate(ctx, nameOrId).PatchedPermissionGroup(patchedPermissionGroup).Execute()



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
    nameOrId := "nameOrId_example" // string | 
    patchedPermissionGroup := *openapiclient.NewPatchedPermissionGroup() // PatchedPermissionGroup |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PermissionsApi.PermissionsGroupsPartialUpdate(context.Background(), nameOrId).PatchedPermissionGroup(patchedPermissionGroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionsApi.PermissionsGroupsPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PermissionsGroupsPartialUpdate`: PermissionGroup
    fmt.Fprintf(os.Stdout, "Response from `PermissionsApi.PermissionsGroupsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPermissionsGroupsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedPermissionGroup** | [**PatchedPermissionGroup**](PatchedPermissionGroup.md) |  | 

### Return type

[**PermissionGroup**](PermissionGroup.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PermissionsGroupsRetrieve

> PermissionGroup PermissionsGroupsRetrieve(ctx, nameOrId).Execute()



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
    nameOrId := "nameOrId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PermissionsApi.PermissionsGroupsRetrieve(context.Background(), nameOrId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionsApi.PermissionsGroupsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PermissionsGroupsRetrieve`: PermissionGroup
    fmt.Fprintf(os.Stdout, "Response from `PermissionsApi.PermissionsGroupsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPermissionsGroupsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PermissionGroup**](PermissionGroup.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PermissionsGroupsUpdate

> PermissionGroup PermissionsGroupsUpdate(ctx, nameOrId).PermissionGroup(permissionGroup).Execute()



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
    nameOrId := "nameOrId_example" // string | 
    permissionGroup := *openapiclient.NewPermissionGroup(int32(123), []string{"Permissions_example"}, "Name_example") // PermissionGroup | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PermissionsApi.PermissionsGroupsUpdate(context.Background(), nameOrId).PermissionGroup(permissionGroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionsApi.PermissionsGroupsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PermissionsGroupsUpdate`: PermissionGroup
    fmt.Fprintf(os.Stdout, "Response from `PermissionsApi.PermissionsGroupsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPermissionsGroupsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **permissionGroup** | [**PermissionGroup**](PermissionGroup.md) |  | 

### Return type

[**PermissionGroup**](PermissionGroup.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PermissionsGroupsUsersCreate

> PermissionGroup PermissionsGroupsUsersCreate(ctx, nameOrId).User(user).Execute()



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
    nameOrId := "nameOrId_example" // string | 
    user := []openapiclient.User{*openapiclient.NewUser("Username_example")} // []User | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PermissionsApi.PermissionsGroupsUsersCreate(context.Background(), nameOrId).User(user).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionsApi.PermissionsGroupsUsersCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PermissionsGroupsUsersCreate`: PermissionGroup
    fmt.Fprintf(os.Stdout, "Response from `PermissionsApi.PermissionsGroupsUsersCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPermissionsGroupsUsersCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **user** | [**[]User**](User.md) |  | 

### Return type

[**PermissionGroup**](PermissionGroup.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PermissionsGroupsUsersDestroy

> PermissionsGroupsUsersDestroy(ctx, nameOrId, usernameOrId).Execute()



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
    nameOrId := "nameOrId_example" // string | 
    usernameOrId := "usernameOrId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PermissionsApi.PermissionsGroupsUsersDestroy(context.Background(), nameOrId, usernameOrId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionsApi.PermissionsGroupsUsersDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOrId** | **string** |  | 
**usernameOrId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPermissionsGroupsUsersDestroyRequest struct via the builder pattern


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

