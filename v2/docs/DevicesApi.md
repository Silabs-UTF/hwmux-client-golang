# \DevicesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DevicesCreate**](DevicesApi.md#DevicesCreate) | **Post** /api/devices/ | 
[**DevicesDestroy**](DevicesApi.md#DevicesDestroy) | **Delete** /api/devices/{id}/ | 
[**DevicesList**](DevicesApi.md#DevicesList) | **Get** /api/devices/ | 
[**DevicesListMyList**](DevicesApi.md#DevicesListMyList) | **Get** /api/devices/list_my/ | 
[**DevicesLocationRetrieve**](DevicesApi.md#DevicesLocationRetrieve) | **Get** /api/devices/{device_pk}/location/ | 
[**DevicesPartialUpdate**](DevicesApi.md#DevicesPartialUpdate) | **Patch** /api/devices/{id}/ | 
[**DevicesPermissionsPartialUpdate**](DevicesApi.md#DevicesPermissionsPartialUpdate) | **Patch** /api/devices/{id}/permissions/ | 
[**DevicesPermissionsRetrieve**](DevicesApi.md#DevicesPermissionsRetrieve) | **Get** /api/devices/{id}/permissions/ | 
[**DevicesPermissionsUpdate**](DevicesApi.md#DevicesPermissionsUpdate) | **Put** /api/devices/{id}/permissions/ | 
[**DevicesReleaseUpdate**](DevicesApi.md#DevicesReleaseUpdate) | **Put** /api/devices/{id}/release/ | 
[**DevicesReserveUpdate**](DevicesApi.md#DevicesReserveUpdate) | **Put** /api/devices/{id}/reserve/ | 
[**DevicesRetrieve**](DevicesApi.md#DevicesRetrieve) | **Get** /api/devices/{id}/ | 
[**DevicesSearchList**](DevicesApi.md#DevicesSearchList) | **Get** /api/devices/search/ | 
[**DevicesSetOfflineCreate**](DevicesApi.md#DevicesSetOfflineCreate) | **Post** /api/devices/set_offline/ | 
[**DevicesStatusCreate**](DevicesApi.md#DevicesStatusCreate) | **Post** /api/devices/{id}/status/ | 
[**DevicesUpdate**](DevicesApi.md#DevicesUpdate) | **Put** /api/devices/{id}/ | 



## DevicesCreate

> WriteOnlyDevice DevicesCreate(ctx).WriteOnlyDevice(writeOnlyDevice).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    writeOnlyDevice := *openapiclient.NewWriteOnlyDevice(int32(123), "Part_example", "LocDesc_example", false, *openapiclient.NewLocationSerializerWriteOnly(int32(123), "Room_example"), time.Now(), time.Now()) // WriteOnlyDevice | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.DevicesCreate(context.Background()).WriteOnlyDevice(writeOnlyDevice).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.DevicesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DevicesCreate`: WriteOnlyDevice
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.DevicesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDevicesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **writeOnlyDevice** | [**WriteOnlyDevice**](WriteOnlyDevice.md) |  | 

### Return type

[**WriteOnlyDevice**](WriteOnlyDevice.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesDestroy

> DevicesDestroy(ctx, id).Execute()





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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.DevicesDestroy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.DevicesDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this device. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDevicesDestroyRequest struct via the builder pattern


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


## DevicesList

> PaginatedDeviceSerializerPublicList DevicesList(ctx).DateCreated(dateCreated).ForceMv(forceMv).IdIn(idIn).IncludePermissionGroups(includePermissionGroups).IsWstk(isWstk).LastUpdate(lastUpdate).Online(online).Ordering(ordering).Page(page).Part(part).PartFamily(partFamily).Room(room).Search(search).Site(site).SnOrName(snOrName).SnOrNameIsnull(snOrNameIsnull).Source(source).Status(status).Uri(uri).UriIsnull(uriIsnull).WstkPart(wstkPart).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    dateCreated := time.Now() // time.Time |  (optional)
    forceMv := "forceMv_example" // string | Specifies the behavior of the force_mv attribute. DEFAULT : Default condition which retrieve data from the Materialized view unless ax exception occurs where we fall back to postgres. ON : Force the data to be retrieved from the Materialized view. OFF : Force the data to be retrieved from the Postgres database (optional) (default to "DEFAULT")
    idIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    includePermissionGroups := true // bool | If set to true, the permission groups associated with this resource will be included in the response. Defaults to false. (optional) (default to false)
    isWstk := true // bool |  (optional)
    lastUpdate := time.Now() // time.Time |  (optional)
    online := true // bool |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    part := "part_example" // string |  (optional)
    partFamily := "partFamily_example" // string |  (optional)
    room := "room_example" // string |  (optional)
    search := "search_example" // string | A search term. (optional)
    site := "site_example" // string |  (optional)
    snOrName := "snOrName_example" // string |  (optional)
    snOrNameIsnull := true // bool |  (optional)
    source := "source_example" // string | * `TERRAFORM` - Terraform * `XML` - XML * `UI` - UI * `OTHER` - Other (optional)
    status := "status_example" // string | * `ACTIVE` - Active * `DISABLED` - Disabled * `OFFLINE` - Offline (optional)
    uri := "uri_example" // string |  (optional)
    uriIsnull := true // bool |  (optional)
    wstkPart := "wstkPart_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.DevicesList(context.Background()).DateCreated(dateCreated).ForceMv(forceMv).IdIn(idIn).IncludePermissionGroups(includePermissionGroups).IsWstk(isWstk).LastUpdate(lastUpdate).Online(online).Ordering(ordering).Page(page).Part(part).PartFamily(partFamily).Room(room).Search(search).Site(site).SnOrName(snOrName).SnOrNameIsnull(snOrNameIsnull).Source(source).Status(status).Uri(uri).UriIsnull(uriIsnull).WstkPart(wstkPart).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.DevicesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DevicesList`: PaginatedDeviceSerializerPublicList
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.DevicesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDevicesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dateCreated** | **time.Time** |  | 
 **forceMv** | **string** | Specifies the behavior of the force_mv attribute. DEFAULT : Default condition which retrieve data from the Materialized view unless ax exception occurs where we fall back to postgres. ON : Force the data to be retrieved from the Materialized view. OFF : Force the data to be retrieved from the Postgres database | [default to &quot;DEFAULT&quot;]
 **idIn** | **[]string** | Multiple values may be separated by commas. | 
 **includePermissionGroups** | **bool** | If set to true, the permission groups associated with this resource will be included in the response. Defaults to false. | [default to false]
 **isWstk** | **bool** |  | 
 **lastUpdate** | **time.Time** |  | 
 **online** | **bool** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **part** | **string** |  | 
 **partFamily** | **string** |  | 
 **room** | **string** |  | 
 **search** | **string** | A search term. | 
 **site** | **string** |  | 
 **snOrName** | **string** |  | 
 **snOrNameIsnull** | **bool** |  | 
 **source** | **string** | * &#x60;TERRAFORM&#x60; - Terraform * &#x60;XML&#x60; - XML * &#x60;UI&#x60; - UI * &#x60;OTHER&#x60; - Other | 
 **status** | **string** | * &#x60;ACTIVE&#x60; - Active * &#x60;DISABLED&#x60; - Disabled * &#x60;OFFLINE&#x60; - Offline | 
 **uri** | **string** |  | 
 **uriIsnull** | **bool** |  | 
 **wstkPart** | **string** |  | 

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


## DevicesListMyList

> PaginatedDeviceSerializerPublicList DevicesListMyList(ctx).DateCreated(dateCreated).IdIn(idIn).IncludePermissionGroups(includePermissionGroups).IsWstk(isWstk).LastUpdate(lastUpdate).Online(online).Ordering(ordering).Page(page).Part(part).PartFamily(partFamily).Room(room).Search(search).Site(site).SnOrName(snOrName).SnOrNameIsnull(snOrNameIsnull).Source(source).Status(status).Uri(uri).UriIsnull(uriIsnull).WstkPart(wstkPart).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    dateCreated := time.Now() // time.Time |  (optional)
    idIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    includePermissionGroups := true // bool | If set to true, the permission groups associated with this resource will be included in the response. Defaults to false. (optional) (default to false)
    isWstk := true // bool |  (optional)
    lastUpdate := time.Now() // time.Time |  (optional)
    online := true // bool |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    part := "part_example" // string |  (optional)
    partFamily := "partFamily_example" // string |  (optional)
    room := "room_example" // string |  (optional)
    search := "search_example" // string | A search term. (optional)
    site := "site_example" // string |  (optional)
    snOrName := "snOrName_example" // string |  (optional)
    snOrNameIsnull := true // bool |  (optional)
    source := "source_example" // string | * `TERRAFORM` - Terraform * `XML` - XML * `UI` - UI * `OTHER` - Other (optional)
    status := "status_example" // string | * `ACTIVE` - Active * `DISABLED` - Disabled * `OFFLINE` - Offline (optional)
    uri := "uri_example" // string |  (optional)
    uriIsnull := true // bool |  (optional)
    wstkPart := "wstkPart_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.DevicesListMyList(context.Background()).DateCreated(dateCreated).IdIn(idIn).IncludePermissionGroups(includePermissionGroups).IsWstk(isWstk).LastUpdate(lastUpdate).Online(online).Ordering(ordering).Page(page).Part(part).PartFamily(partFamily).Room(room).Search(search).Site(site).SnOrName(snOrName).SnOrNameIsnull(snOrNameIsnull).Source(source).Status(status).Uri(uri).UriIsnull(uriIsnull).WstkPart(wstkPart).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.DevicesListMyList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DevicesListMyList`: PaginatedDeviceSerializerPublicList
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.DevicesListMyList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDevicesListMyListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dateCreated** | **time.Time** |  | 
 **idIn** | **[]string** | Multiple values may be separated by commas. | 
 **includePermissionGroups** | **bool** | If set to true, the permission groups associated with this resource will be included in the response. Defaults to false. | [default to false]
 **isWstk** | **bool** |  | 
 **lastUpdate** | **time.Time** |  | 
 **online** | **bool** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **part** | **string** |  | 
 **partFamily** | **string** |  | 
 **room** | **string** |  | 
 **search** | **string** | A search term. | 
 **site** | **string** |  | 
 **snOrName** | **string** |  | 
 **snOrNameIsnull** | **bool** |  | 
 **source** | **string** | * &#x60;TERRAFORM&#x60; - Terraform * &#x60;XML&#x60; - XML * &#x60;UI&#x60; - UI * &#x60;OTHER&#x60; - Other | 
 **status** | **string** | * &#x60;ACTIVE&#x60; - Active * &#x60;DISABLED&#x60; - Disabled * &#x60;OFFLINE&#x60; - Offline | 
 **uri** | **string** |  | 
 **uriIsnull** | **bool** |  | 
 **wstkPart** | **string** |  | 

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


## DevicesLocationRetrieve

> Location DevicesLocationRetrieve(ctx, devicePk).Execute()



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
    devicePk := "devicePk_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.DevicesLocationRetrieve(context.Background(), devicePk).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.DevicesLocationRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DevicesLocationRetrieve`: Location
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.DevicesLocationRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**devicePk** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDevicesLocationRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Location**](Location.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesPartialUpdate

> WriteOnlyDevice DevicesPartialUpdate(ctx, id).PatchedWriteOnlyDevice(patchedWriteOnlyDevice).Execute()





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
    patchedWriteOnlyDevice := *openapiclient.NewPatchedWriteOnlyDevice() // PatchedWriteOnlyDevice |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.DevicesPartialUpdate(context.Background(), id).PatchedWriteOnlyDevice(patchedWriteOnlyDevice).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.DevicesPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DevicesPartialUpdate`: WriteOnlyDevice
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.DevicesPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this device. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDevicesPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedWriteOnlyDevice** | [**PatchedWriteOnlyDevice**](PatchedWriteOnlyDevice.md) |  | 

### Return type

[**WriteOnlyDevice**](WriteOnlyDevice.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesPermissionsPartialUpdate

> ObjectPermissions DevicesPermissionsPartialUpdate(ctx, id).PatchedObjectPermissions(patchedObjectPermissions).Execute()





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
    patchedObjectPermissions := *openapiclient.NewPatchedObjectPermissions() // PatchedObjectPermissions |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.DevicesPermissionsPartialUpdate(context.Background(), id).PatchedObjectPermissions(patchedObjectPermissions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.DevicesPermissionsPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DevicesPermissionsPartialUpdate`: ObjectPermissions
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.DevicesPermissionsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this device. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDevicesPermissionsPartialUpdateRequest struct via the builder pattern


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


## DevicesPermissionsRetrieve

> ObjectPermissions DevicesPermissionsRetrieve(ctx, id).Execute()





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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.DevicesPermissionsRetrieve(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.DevicesPermissionsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DevicesPermissionsRetrieve`: ObjectPermissions
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.DevicesPermissionsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this device. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDevicesPermissionsRetrieveRequest struct via the builder pattern


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


## DevicesPermissionsUpdate

> ObjectPermissions DevicesPermissionsUpdate(ctx, id).ObjectPermissions(objectPermissions).Execute()





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
    objectPermissions := *openapiclient.NewObjectPermissions(map[string]interface{}{"key": interface{}(123)}, map[string]interface{}{"key": interface{}(123)}) // ObjectPermissions | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.DevicesPermissionsUpdate(context.Background(), id).ObjectPermissions(objectPermissions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.DevicesPermissionsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DevicesPermissionsUpdate`: ObjectPermissions
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.DevicesPermissionsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this device. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDevicesPermissionsUpdateRequest struct via the builder pattern


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


## DevicesReleaseUpdate

> DeviceSerializerPublic DevicesReleaseUpdate(ctx, id).Execute()





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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.DevicesReleaseUpdate(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.DevicesReleaseUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DevicesReleaseUpdate`: DeviceSerializerPublic
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.DevicesReleaseUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this device. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDevicesReleaseUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeviceSerializerPublic**](DeviceSerializerPublic.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesReserveUpdate

> DeviceSerializerPublic DevicesReserveUpdate(ctx, id).Details(details).Execute()





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
    details := "details_example" // string | Additional information such as the jenkins job URL (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.DevicesReserveUpdate(context.Background(), id).Details(details).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.DevicesReserveUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DevicesReserveUpdate`: DeviceSerializerPublic
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.DevicesReserveUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this device. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDevicesReserveUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **details** | **string** | Additional information such as the jenkins job URL | 

### Return type

[**DeviceSerializerPublic**](DeviceSerializerPublic.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesRetrieve

> DeviceSerializerPublic DevicesRetrieve(ctx, id).IncludePermissionGroups(includePermissionGroups).Execute()





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
    includePermissionGroups := true // bool | If set to true, the permission groups associated with this resource will be included in the response. Defaults to false. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.DevicesRetrieve(context.Background(), id).IncludePermissionGroups(includePermissionGroups).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.DevicesRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DevicesRetrieve`: DeviceSerializerPublic
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.DevicesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this device. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDevicesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includePermissionGroups** | **bool** | If set to true, the permission groups associated with this resource will be included in the response. Defaults to false. | [default to false]

### Return type

[**DeviceSerializerPublic**](DeviceSerializerPublic.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesSearchList

> PaginatedDeviceSerializerPublicList DevicesSearchList(ctx).SearchKeyValuePairs(searchKeyValuePairs).DateCreated(dateCreated).IdIn(idIn).IsWstk(isWstk).LastUpdate(lastUpdate).Online(online).Ordering(ordering).Page(page).Part(part).PartFamily(partFamily).Room(room).Search(search).Site(site).SnOrName(snOrName).SnOrNameIsnull(snOrNameIsnull).Source(source).Status(status).Uri(uri).UriIsnull(uriIsnull).WstkPart(wstkPart).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    searchKeyValuePairs := "part_no=unknown&is_soc=True" // string | 
    dateCreated := time.Now() // time.Time |  (optional)
    idIn := []string{"Inner_example"} // []string | Multiple values may be separated by commas. (optional)
    isWstk := true // bool |  (optional)
    lastUpdate := time.Now() // time.Time |  (optional)
    online := true // bool |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    part := "part_example" // string |  (optional)
    partFamily := "partFamily_example" // string |  (optional)
    room := "room_example" // string |  (optional)
    search := "search_example" // string | A search term. (optional)
    site := "site_example" // string |  (optional)
    snOrName := "snOrName_example" // string |  (optional)
    snOrNameIsnull := true // bool |  (optional)
    source := "source_example" // string | * `TERRAFORM` - Terraform * `XML` - XML * `UI` - UI * `OTHER` - Other (optional)
    status := "status_example" // string | * `ACTIVE` - Active * `DISABLED` - Disabled * `OFFLINE` - Offline (optional)
    uri := "uri_example" // string |  (optional)
    uriIsnull := true // bool |  (optional)
    wstkPart := "wstkPart_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.DevicesSearchList(context.Background()).SearchKeyValuePairs(searchKeyValuePairs).DateCreated(dateCreated).IdIn(idIn).IsWstk(isWstk).LastUpdate(lastUpdate).Online(online).Ordering(ordering).Page(page).Part(part).PartFamily(partFamily).Room(room).Search(search).Site(site).SnOrName(snOrName).SnOrNameIsnull(snOrNameIsnull).Source(source).Status(status).Uri(uri).UriIsnull(uriIsnull).WstkPart(wstkPart).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.DevicesSearchList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DevicesSearchList`: PaginatedDeviceSerializerPublicList
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.DevicesSearchList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDevicesSearchListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchKeyValuePairs** | **string** |  | 
 **dateCreated** | **time.Time** |  | 
 **idIn** | **[]string** | Multiple values may be separated by commas. | 
 **isWstk** | **bool** |  | 
 **lastUpdate** | **time.Time** |  | 
 **online** | **bool** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **part** | **string** |  | 
 **partFamily** | **string** |  | 
 **room** | **string** |  | 
 **search** | **string** | A search term. | 
 **site** | **string** |  | 
 **snOrName** | **string** |  | 
 **snOrNameIsnull** | **bool** |  | 
 **source** | **string** | * &#x60;TERRAFORM&#x60; - Terraform * &#x60;XML&#x60; - XML * &#x60;UI&#x60; - UI * &#x60;OTHER&#x60; - Other | 
 **status** | **string** | * &#x60;ACTIVE&#x60; - Active * &#x60;DISABLED&#x60; - Disabled * &#x60;OFFLINE&#x60; - Offline | 
 **uri** | **string** |  | 
 **uriIsnull** | **bool** |  | 
 **wstkPart** | **string** |  | 

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


## DevicesSetOfflineCreate

> WriteOnlyDevice DevicesSetOfflineCreate(ctx).WriteOnlyDevice(writeOnlyDevice).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    writeOnlyDevice := *openapiclient.NewWriteOnlyDevice(int32(123), "Part_example", "LocDesc_example", false, *openapiclient.NewLocationSerializerWriteOnly(int32(123), "Room_example"), time.Now(), time.Now()) // WriteOnlyDevice | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.DevicesSetOfflineCreate(context.Background()).WriteOnlyDevice(writeOnlyDevice).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.DevicesSetOfflineCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DevicesSetOfflineCreate`: WriteOnlyDevice
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.DevicesSetOfflineCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDevicesSetOfflineCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **writeOnlyDevice** | [**WriteOnlyDevice**](WriteOnlyDevice.md) |  | 

### Return type

[**WriteOnlyDevice**](WriteOnlyDevice.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesStatusCreate

> ResourceStatusRequest DevicesStatusCreate(ctx, id).ResourceStatusRequest(resourceStatusRequest).Execute()





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
    resourceStatusRequest := *openapiclient.NewResourceStatusRequest(openapiclient.StatusEnum("ACTIVE")) // ResourceStatusRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.DevicesStatusCreate(context.Background(), id).ResourceStatusRequest(resourceStatusRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.DevicesStatusCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DevicesStatusCreate`: ResourceStatusRequest
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.DevicesStatusCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this device. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDevicesStatusCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resourceStatusRequest** | [**ResourceStatusRequest**](ResourceStatusRequest.md) |  | 

### Return type

[**ResourceStatusRequest**](ResourceStatusRequest.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesUpdate

> WriteOnlyDevice DevicesUpdate(ctx, id).WriteOnlyDevice(writeOnlyDevice).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    id := int32(56) // int32 | A unique integer value identifying this device.
    writeOnlyDevice := *openapiclient.NewWriteOnlyDevice(int32(123), "Part_example", "LocDesc_example", false, *openapiclient.NewLocationSerializerWriteOnly(int32(123), "Room_example"), time.Now(), time.Now()) // WriteOnlyDevice | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.DevicesUpdate(context.Background(), id).WriteOnlyDevice(writeOnlyDevice).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.DevicesUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DevicesUpdate`: WriteOnlyDevice
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.DevicesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this device. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDevicesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **writeOnlyDevice** | [**WriteOnlyDevice**](WriteOnlyDevice.md) |  | 

### Return type

[**WriteOnlyDevice**](WriteOnlyDevice.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

