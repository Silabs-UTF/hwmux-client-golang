# \DevicesApi

All URIs are relative to *https://hwmux.silabs.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DevicesCreate**](DevicesApi.md#DevicesCreate) | **Post** /api/devices/ | 
[**DevicesDestroy**](DevicesApi.md#DevicesDestroy) | **Delete** /api/devices/{id}/ | 
[**DevicesList**](DevicesApi.md#DevicesList) | **Get** /api/devices/ | 
[**DevicesListMyList**](DevicesApi.md#DevicesListMyList) | **Get** /api/devices/list_my/ | 
[**DevicesLocationRetrieve**](DevicesApi.md#DevicesLocationRetrieve) | **Get** /api/devices/{device_pk}/location/ | 
[**DevicesPartialUpdate**](DevicesApi.md#DevicesPartialUpdate) | **Patch** /api/devices/{id}/ | 
[**DevicesReleaseUpdate**](DevicesApi.md#DevicesReleaseUpdate) | **Put** /api/devices/{id}/release/ | 
[**DevicesReserveUpdate**](DevicesApi.md#DevicesReserveUpdate) | **Put** /api/devices/{id}/reserve/ | 
[**DevicesRetrieve**](DevicesApi.md#DevicesRetrieve) | **Get** /api/devices/{id}/ | 
[**DevicesSearchList**](DevicesApi.md#DevicesSearchList) | **Get** /api/devices/search/ | 
[**DevicesSetOfflineCreate**](DevicesApi.md#DevicesSetOfflineCreate) | **Post** /api/devices/set_offline/ | 
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

> PaginatedDeviceSerializerPublicList DevicesList(ctx).DateCreated(dateCreated).IsWstk(isWstk).LastUpdate(lastUpdate).Online(online).Ordering(ordering).Page(page).Part(part).PartFamily(partFamily).Room(room).Search(search).Site(site).SnOrName(snOrName).SnOrNameIsnull(snOrNameIsnull).Uri(uri).UriIsnull(uriIsnull).Execute()





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
    uri := "uri_example" // string |  (optional)
    uriIsnull := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.DevicesList(context.Background()).DateCreated(dateCreated).IsWstk(isWstk).LastUpdate(lastUpdate).Online(online).Ordering(ordering).Page(page).Part(part).PartFamily(partFamily).Room(room).Search(search).Site(site).SnOrName(snOrName).SnOrNameIsnull(snOrNameIsnull).Uri(uri).UriIsnull(uriIsnull).Execute()
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
 **uri** | **string** |  | 
 **uriIsnull** | **bool** |  | 

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

> PaginatedDeviceSerializerPublicList DevicesListMyList(ctx).DateCreated(dateCreated).IsWstk(isWstk).LastUpdate(lastUpdate).Online(online).Ordering(ordering).Page(page).Part(part).PartFamily(partFamily).Room(room).Search(search).Site(site).SnOrName(snOrName).SnOrNameIsnull(snOrNameIsnull).Uri(uri).UriIsnull(uriIsnull).Execute()





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
    uri := "uri_example" // string |  (optional)
    uriIsnull := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.DevicesListMyList(context.Background()).DateCreated(dateCreated).IsWstk(isWstk).LastUpdate(lastUpdate).Online(online).Ordering(ordering).Page(page).Part(part).PartFamily(partFamily).Room(room).Search(search).Site(site).SnOrName(snOrName).SnOrNameIsnull(snOrNameIsnull).Uri(uri).UriIsnull(uriIsnull).Execute()
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
 **uri** | **string** |  | 
 **uriIsnull** | **bool** |  | 

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

> DeviceSerializerPublic DevicesRetrieve(ctx, id).Execute()





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
    resp, r, err := apiClient.DevicesApi.DevicesRetrieve(context.Background(), id).Execute()
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

> PaginatedDeviceSerializerPublicList DevicesSearchList(ctx).SearchKeyValuePairs(searchKeyValuePairs).DateCreated(dateCreated).IsWstk(isWstk).LastUpdate(lastUpdate).Online(online).Ordering(ordering).Page(page).Part(part).PartFamily(partFamily).Room(room).Search(search).Site(site).SnOrName(snOrName).SnOrNameIsnull(snOrNameIsnull).Uri(uri).UriIsnull(uriIsnull).Execute()





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
    uri := "uri_example" // string |  (optional)
    uriIsnull := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DevicesApi.DevicesSearchList(context.Background()).SearchKeyValuePairs(searchKeyValuePairs).DateCreated(dateCreated).IsWstk(isWstk).LastUpdate(lastUpdate).Online(online).Ordering(ordering).Page(page).Part(part).PartFamily(partFamily).Room(room).Search(search).Site(site).SnOrName(snOrName).SnOrNameIsnull(snOrNameIsnull).Uri(uri).UriIsnull(uriIsnull).Execute()
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
 **uri** | **string** |  | 
 **uriIsnull** | **bool** |  | 

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

