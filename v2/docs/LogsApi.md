# \LogsAPI

All URIs are relative to *https://hwmux.silabs.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LogsDestroy**](LogsAPI.md#LogsDestroy) | **Delete** /api/logs/{id} | 
[**LogsList**](LogsAPI.md#LogsList) | **Get** /api/logs/ | 
[**LogsRetrieve**](LogsAPI.md#LogsRetrieve) | **Get** /api/logs/{id}/ | 



## LogsDestroy

> LogsDestroy(ctx, id).Execute()



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.LogsAPI.LogsDestroy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsAPI.LogsDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLogsDestroyRequest struct via the builder pattern


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


## LogsList

> PaginatedLogList LogsList(ctx).CausedBy(causedBy).DateAfter(dateAfter).DateBefore(dateBefore).Datetime(datetime).Details(details).Device(device).DeviceSnOrName(deviceSnOrName).DeviceGroup(deviceGroup).DeviceGroupName(deviceGroupName).DeviceGroupLabel(deviceGroupLabel).Event(event).Owner(owner).Page(page).Search(search).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/Silabs-UTF/hwmux-client-golang/v2"
)

func main() {
    causedBy := int32(56) // int32 |  (optional)
    dateAfter := time.Now() // string |  (optional)
    dateBefore := time.Now() // string |  (optional)
    datetime := time.Now() // time.Time |  (optional)
    details := "details_example" // string |  (optional)
    device := int32(56) // int32 |  (optional)
    deviceSnOrName := "deviceSnOrName_example" // string |  (optional)
    deviceGroup := int32(56) // int32 |  (optional)
    deviceGroupName := "deviceGroupName_example" // string |  (optional)
    deviceGroupLabel := int32(56) // int32 |  (optional)
    event := "event_example" // string |  (optional)
    owner := "owner_example" // string | Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    search := "search_example" // string | A search term. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogsAPI.LogsList(context.Background()).CausedBy(causedBy).DateAfter(dateAfter).DateBefore(dateBefore).Datetime(datetime).Details(details).Device(device).DeviceSnOrName(deviceSnOrName).DeviceGroup(deviceGroup).DeviceGroupName(deviceGroupName).DeviceGroupLabel(deviceGroupLabel).Event(event).Owner(owner).Page(page).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsAPI.LogsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LogsList`: PaginatedLogList
    fmt.Fprintf(os.Stdout, "Response from `LogsAPI.LogsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLogsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **causedBy** | **int32** |  | 
 **dateAfter** | **string** |  | 
 **dateBefore** | **string** |  | 
 **datetime** | **time.Time** |  | 
 **details** | **string** |  | 
 **device** | **int32** |  | 
 **deviceSnOrName** | **string** |  | 
 **deviceGroup** | **int32** |  | 
 **deviceGroupName** | **string** |  | 
 **deviceGroupLabel** | **int32** |  | 
 **event** | **string** |  | 
 **owner** | **string** | Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedLogList**](PaginatedLogList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LogsRetrieve

> Log LogsRetrieve(ctx, id).Execute()



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
    id := int32(56) // int32 | A unique integer value identifying this log.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogsAPI.LogsRetrieve(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsAPI.LogsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LogsRetrieve`: Log
    fmt.Fprintf(os.Stdout, "Response from `LogsAPI.LogsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this log. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLogsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Log**](Log.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

