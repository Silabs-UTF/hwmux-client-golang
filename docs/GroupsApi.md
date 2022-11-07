# \GroupsApi

All URIs are relative to *https://hwmux.silabs.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupsAvailableList**](GroupsApi.md#GroupsAvailableList) | **Get** /api/groups/available/ | 
[**GroupsCreate**](GroupsApi.md#GroupsCreate) | **Post** /api/groups/ | 
[**GroupsCreateWithDevicesCreate**](GroupsApi.md#GroupsCreateWithDevicesCreate) | **Post** /api/groups/create-with-devices | 
[**GroupsDestroy**](GroupsApi.md#GroupsDestroy) | **Delete** /api/groups/{id}/ | 
[**GroupsList**](GroupsApi.md#GroupsList) | **Get** /api/groups/ | 
[**GroupsMyList**](GroupsApi.md#GroupsMyList) | **Get** /api/groups/my/ | 
[**GroupsPartialUpdate**](GroupsApi.md#GroupsPartialUpdate) | **Patch** /api/groups/{id}/ | 
[**GroupsReleaseByNameUpdate**](GroupsApi.md#GroupsReleaseByNameUpdate) | **Put** /api/groups/{group_name}/release_by_name/ | 
[**GroupsReleaseUpdate**](GroupsApi.md#GroupsReleaseUpdate) | **Put** /api/groups/{id}/release/ | 
[**GroupsReserveByNameUpdate**](GroupsApi.md#GroupsReserveByNameUpdate) | **Put** /api/groups/{group_name}/reserve_by_name/ | 
[**GroupsReserveUpdate**](GroupsApi.md#GroupsReserveUpdate) | **Put** /api/groups/{id}/reserve/ | 
[**GroupsRetrieve**](GroupsApi.md#GroupsRetrieve) | **Get** /api/groups/{id}/ | 
[**GroupsUpdate**](GroupsApi.md#GroupsUpdate) | **Put** /api/groups/{id}/ | 



## GroupsAvailableList

> PaginatedDeviceGroupList GroupsAvailableList(ctx).Devices(devices).Name(name).Ordering(ordering).Page(page).Search(search).Execute()





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
    devices := []int32{int32(123)} // []int32 |  (optional)
    name := "name_example" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    search := "search_example" // string | A search term. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.GroupsAvailableList(context.Background()).Devices(devices).Name(name).Ordering(ordering).Page(page).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GroupsAvailableList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsAvailableList`: PaginatedDeviceGroupList
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GroupsAvailableList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGroupsAvailableListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **devices** | **[]int32** |  | 
 **name** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **search** | **string** | A search term. | 

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


## GroupsCreate

> DeviceGroupSerializerWithDevicePk GroupsCreate(ctx).DeviceGroupSerializerWithDevicePk(deviceGroupSerializerWithDevicePk).Execute()





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
    deviceGroupSerializerWithDevicePk := *openapiclient.NewDeviceGroupSerializerWithDevicePk(int32(123), false, "Name_example") // DeviceGroupSerializerWithDevicePk | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.GroupsCreate(context.Background()).DeviceGroupSerializerWithDevicePk(deviceGroupSerializerWithDevicePk).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GroupsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsCreate`: DeviceGroupSerializerWithDevicePk
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GroupsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGroupsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceGroupSerializerWithDevicePk** | [**DeviceGroupSerializerWithDevicePk**](DeviceGroupSerializerWithDevicePk.md) |  | 

### Return type

[**DeviceGroupSerializerWithDevicePk**](DeviceGroupSerializerWithDevicePk.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsCreateWithDevicesCreate

> DeviceGroup GroupsCreateWithDevicesCreate(ctx).NestedDeviceGroup(nestedDeviceGroup).Execute()





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
    nestedDeviceGroup := *openapiclient.NewNestedDeviceGroup("Name_example") // NestedDeviceGroup | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.GroupsCreateWithDevicesCreate(context.Background()).NestedDeviceGroup(nestedDeviceGroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GroupsCreateWithDevicesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsCreateWithDevicesCreate`: DeviceGroup
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GroupsCreateWithDevicesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGroupsCreateWithDevicesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nestedDeviceGroup** | [**NestedDeviceGroup**](NestedDeviceGroup.md) |  | 

### Return type

[**DeviceGroup**](DeviceGroup.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsDestroy

> GroupsDestroy(ctx, id).Execute()





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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.GroupsDestroy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GroupsDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this device group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsDestroyRequest struct via the builder pattern


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


## GroupsList

> PaginatedDeviceGroupList GroupsList(ctx).Devices(devices).Name(name).Ordering(ordering).Page(page).Search(search).Execute()





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
    devices := []int32{int32(123)} // []int32 |  (optional)
    name := "name_example" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    search := "search_example" // string | A search term. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.GroupsList(context.Background()).Devices(devices).Name(name).Ordering(ordering).Page(page).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GroupsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsList`: PaginatedDeviceGroupList
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GroupsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGroupsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **devices** | **[]int32** |  | 
 **name** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **search** | **string** | A search term. | 

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


## GroupsMyList

> PaginatedDeviceGroupList GroupsMyList(ctx).Devices(devices).Name(name).Ordering(ordering).Page(page).Search(search).Execute()





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
    devices := []int32{int32(123)} // []int32 |  (optional)
    name := "name_example" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    search := "search_example" // string | A search term. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.GroupsMyList(context.Background()).Devices(devices).Name(name).Ordering(ordering).Page(page).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GroupsMyList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsMyList`: PaginatedDeviceGroupList
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GroupsMyList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGroupsMyListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **devices** | **[]int32** |  | 
 **name** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **search** | **string** | A search term. | 

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


## GroupsPartialUpdate

> DeviceGroupSerializerWithDevicePk GroupsPartialUpdate(ctx, id).PatchedDeviceGroupSerializerWithDevicePk(patchedDeviceGroupSerializerWithDevicePk).Execute()





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
    patchedDeviceGroupSerializerWithDevicePk := *openapiclient.NewPatchedDeviceGroupSerializerWithDevicePk() // PatchedDeviceGroupSerializerWithDevicePk |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.GroupsPartialUpdate(context.Background(), id).PatchedDeviceGroupSerializerWithDevicePk(patchedDeviceGroupSerializerWithDevicePk).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GroupsPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsPartialUpdate`: DeviceGroupSerializerWithDevicePk
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GroupsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this device group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedDeviceGroupSerializerWithDevicePk** | [**PatchedDeviceGroupSerializerWithDevicePk**](PatchedDeviceGroupSerializerWithDevicePk.md) |  | 

### Return type

[**DeviceGroupSerializerWithDevicePk**](DeviceGroupSerializerWithDevicePk.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsReleaseByNameUpdate

> DeviceGroup GroupsReleaseByNameUpdate(ctx, groupName).Execute()



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
    groupName := "groupName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.GroupsReleaseByNameUpdate(context.Background(), groupName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GroupsReleaseByNameUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsReleaseByNameUpdate`: DeviceGroup
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GroupsReleaseByNameUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsReleaseByNameUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GroupsReleaseUpdate

> DeviceGroup GroupsReleaseUpdate(ctx, id).Execute()





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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.GroupsReleaseUpdate(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GroupsReleaseUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsReleaseUpdate`: DeviceGroup
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GroupsReleaseUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this device group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsReleaseUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GroupsReserveByNameUpdate

> DeviceGroup GroupsReserveByNameUpdate(ctx, groupName).Details(details).Execute()



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
    groupName := "groupName_example" // string | 
    details := "details_example" // string | Additional information such as the jenkins job URL (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.GroupsReserveByNameUpdate(context.Background(), groupName).Details(details).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GroupsReserveByNameUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsReserveByNameUpdate`: DeviceGroup
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GroupsReserveByNameUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsReserveByNameUpdateRequest struct via the builder pattern


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


## GroupsReserveUpdate

> DeviceGroup GroupsReserveUpdate(ctx, id).Details(details).Execute()





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
    details := "details_example" // string | Additional information such as the jenkins job URL (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.GroupsReserveUpdate(context.Background(), id).Details(details).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GroupsReserveUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsReserveUpdate`: DeviceGroup
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GroupsReserveUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this device group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsReserveUpdateRequest struct via the builder pattern


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


## GroupsRetrieve

> DeviceGroup GroupsRetrieve(ctx, id).Execute()





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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.GroupsRetrieve(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GroupsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsRetrieve`: DeviceGroup
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GroupsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this device group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GroupsUpdate

> DeviceGroupSerializerWithDevicePk GroupsUpdate(ctx, id).DeviceGroupSerializerWithDevicePk(deviceGroupSerializerWithDevicePk).Execute()





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
    deviceGroupSerializerWithDevicePk := *openapiclient.NewDeviceGroupSerializerWithDevicePk(int32(123), false, "Name_example") // DeviceGroupSerializerWithDevicePk | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.GroupsUpdate(context.Background(), id).DeviceGroupSerializerWithDevicePk(deviceGroupSerializerWithDevicePk).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GroupsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsUpdate`: DeviceGroupSerializerWithDevicePk
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GroupsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this device group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deviceGroupSerializerWithDevicePk** | [**DeviceGroupSerializerWithDevicePk**](DeviceGroupSerializerWithDevicePk.md) |  | 

### Return type

[**DeviceGroupSerializerWithDevicePk**](DeviceGroupSerializerWithDevicePk.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

