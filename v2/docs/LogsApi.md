# \LogsApi

All URIs are relative to *https://hwmux.silabs.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LogsDestroy**](LogsApi.md#LogsDestroy) | **Delete** /api/logs/{id} | 
[**LogsList**](LogsApi.md#LogsList) | **Get** /api/logs/ | 
[**LogsRetrieve**](LogsApi.md#LogsRetrieve) | **Get** /api/logs/{id}/ | 



## LogsDestroy

> LogsDestroy(ctx, id).Execute()



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogsApi.LogsDestroy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsApi.LogsDestroy``: %v\n", err)
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

> PaginatedLogList LogsList(ctx).CausedBy(causedBy).DateAfter(dateAfter).DateBefore(dateBefore).Datetime(datetime).Details(details).Device(device).DeviceSnOrName(deviceSnOrName).DeviceGroup(deviceGroup).DeviceGroupName(deviceGroupName).DeviceGroupLabel(deviceGroupLabel).DeviceGroupLabelName(deviceGroupLabelName).Event(event).Ordering(ordering).Owner(owner).Page(page).ResourceName(resourceName).ResourceNameList(resourceNameList).ResourceType(resourceType).Search(search).Status(status).Execute()



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
    causedBy := int32(56) // int32 |  (optional)
    dateAfter := time.Now() // time.Time |  (optional)
    dateBefore := time.Now() // time.Time |  (optional)
    datetime := time.Now() // time.Time |  (optional)
    details := "details_example" // string |  (optional)
    device := []int32{int32(123)} // []int32 |  (optional)
    deviceSnOrName := "deviceSnOrName_example" // string |  (optional)
    deviceGroup := []int32{int32(123)} // []int32 |  (optional)
    deviceGroupName := "deviceGroupName_example" // string |  (optional)
    deviceGroupLabel := []int32{int32(123)} // []int32 |  (optional)
    deviceGroupLabelName := "deviceGroupLabelName_example" // string |  (optional)
    event := "event_example" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    owner := "owner_example" // string | Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    resourceName := "resourceName_example" // string |  (optional)
    resourceNameList := "resourceNameList_example" // string |  (optional)
    resourceType := "resourceType_example" // string |  (optional)
    search := "search_example" // string | A search term. (optional)
    status := "status_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogsApi.LogsList(context.Background()).CausedBy(causedBy).DateAfter(dateAfter).DateBefore(dateBefore).Datetime(datetime).Details(details).Device(device).DeviceSnOrName(deviceSnOrName).DeviceGroup(deviceGroup).DeviceGroupName(deviceGroupName).DeviceGroupLabel(deviceGroupLabel).DeviceGroupLabelName(deviceGroupLabelName).Event(event).Ordering(ordering).Owner(owner).Page(page).ResourceName(resourceName).ResourceNameList(resourceNameList).ResourceType(resourceType).Search(search).Status(status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsApi.LogsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LogsList`: PaginatedLogList
    fmt.Fprintf(os.Stdout, "Response from `LogsApi.LogsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLogsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **causedBy** | **int32** |  | 
 **dateAfter** | **time.Time** |  | 
 **dateBefore** | **time.Time** |  | 
 **datetime** | **time.Time** |  | 
 **details** | **string** |  | 
 **device** | **[]int32** |  | 
 **deviceSnOrName** | **string** |  | 
 **deviceGroup** | **[]int32** |  | 
 **deviceGroupName** | **string** |  | 
 **deviceGroupLabel** | **[]int32** |  | 
 **deviceGroupLabelName** | **string** |  | 
 **event** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **owner** | **string** | Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **resourceName** | **string** |  | 
 **resourceNameList** | **string** |  | 
 **resourceType** | **string** |  | 
 **search** | **string** | A search term. | 
 **status** | **string** |  | 

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
    openapiclient "./openapi"
)

func main() {
    id := int32(56) // int32 | A unique integer value identifying this log.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogsApi.LogsRetrieve(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsApi.LogsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LogsRetrieve`: Log
    fmt.Fprintf(os.Stdout, "Response from `LogsApi.LogsRetrieve`: %v\n", resp)
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

