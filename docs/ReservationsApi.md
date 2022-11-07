# \ReservationsApi

All URIs are relative to *https://hwmux.silabs.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReservationsActionableList**](ReservationsApi.md#ReservationsActionableList) | **Get** /api/reservations/actionable/ | 
[**ReservationsActiveList**](ReservationsApi.md#ReservationsActiveList) | **Get** /api/reservations/active/ | 
[**ReservationsCancelUpdate**](ReservationsApi.md#ReservationsCancelUpdate) | **Put** /api/reservations/{id}/cancel/ | 
[**ReservationsCreate**](ReservationsApi.md#ReservationsCreate) | **Post** /api/reservations/ | 
[**ReservationsList**](ReservationsApi.md#ReservationsList) | **Get** /api/reservations/ | 
[**ReservationsMetadataPartialUpdate**](ReservationsApi.md#ReservationsMetadataPartialUpdate) | **Patch** /api/reservations/{id}/metadata/ | 
[**ReservationsMetadataUpdate**](ReservationsApi.md#ReservationsMetadataUpdate) | **Put** /api/reservations/{id}/metadata/ | 
[**ReservationsReleaseUpdate**](ReservationsApi.md#ReservationsReleaseUpdate) | **Put** /api/reservations/{id}/release/ | 
[**ReservationsRetrieve**](ReservationsApi.md#ReservationsRetrieve) | **Get** /api/reservations/{id}/ | 
[**ReservationsUpdate**](ReservationsApi.md#ReservationsUpdate) | **Put** /api/reservations/{id}/ | 



## ReservationsActionableList

> PaginatedReservationSessionSerializerReadOnlyList ReservationsActionableList(ctx).ADeviceGroups(aDeviceGroups).ADevices(aDevices).Details(details).Id(id).MaxHistory(maxHistory).Ordering(ordering).Owner(owner).Page(page).RDeviceGroupLabels(rDeviceGroupLabels).RDeviceGroups(rDeviceGroups).RDevices(rDevices).Search(search).TCompleted(tCompleted).TCompletedAfter(tCompletedAfter).TCompletedBefore(tCompletedBefore).TCompletedIsnull(tCompletedIsnull).TLeaseExpires(tLeaseExpires).TLeaseExpiresAfter(tLeaseExpiresAfter).TLeaseExpiresBefore(tLeaseExpiresBefore).TLeaseExpiresIsnull(tLeaseExpiresIsnull).TRequested(tRequested).TRequestedAfter(tRequestedAfter).TRequestedBefore(tRequestedBefore).TSatisfied(tSatisfied).TSatisfiedAfter(tSatisfiedAfter).TSatisfiedBefore(tSatisfiedBefore).TSatisfiedIsnull(tSatisfiedIsnull).UseWatchdog(useWatchdog).Execute()





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
    aDeviceGroups := []int32{int32(123)} // []int32 |  (optional)
    aDevices := []int32{int32(123)} // []int32 |  (optional)
    details := "details_example" // string |  (optional)
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
    maxHistory := time.Now() // time.Time | Filter out reservations that expired before the specified datetime. Defaults to 24 hours prior. (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    owner := int32(56) // int32 |  (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    rDeviceGroupLabels := []int32{int32(123)} // []int32 |  (optional)
    rDeviceGroups := []int32{int32(123)} // []int32 |  (optional)
    rDevices := []int32{int32(123)} // []int32 |  (optional)
    search := "search_example" // string | A search term. (optional)
    tCompleted := time.Now() // time.Time |  (optional)
    tCompletedAfter := time.Now() // time.Time |  (optional)
    tCompletedBefore := time.Now() // time.Time |  (optional)
    tCompletedIsnull := true // bool |  (optional)
    tLeaseExpires := time.Now() // time.Time |  (optional)
    tLeaseExpiresAfter := time.Now() // time.Time |  (optional)
    tLeaseExpiresBefore := time.Now() // time.Time |  (optional)
    tLeaseExpiresIsnull := true // bool |  (optional)
    tRequested := time.Now() // time.Time |  (optional)
    tRequestedAfter := time.Now() // time.Time |  (optional)
    tRequestedBefore := time.Now() // time.Time |  (optional)
    tSatisfied := time.Now() // time.Time |  (optional)
    tSatisfiedAfter := time.Now() // time.Time |  (optional)
    tSatisfiedBefore := time.Now() // time.Time |  (optional)
    tSatisfiedIsnull := true // bool |  (optional)
    useWatchdog := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReservationsApi.ReservationsActionableList(context.Background()).ADeviceGroups(aDeviceGroups).ADevices(aDevices).Details(details).Id(id).MaxHistory(maxHistory).Ordering(ordering).Owner(owner).Page(page).RDeviceGroupLabels(rDeviceGroupLabels).RDeviceGroups(rDeviceGroups).RDevices(rDevices).Search(search).TCompleted(tCompleted).TCompletedAfter(tCompletedAfter).TCompletedBefore(tCompletedBefore).TCompletedIsnull(tCompletedIsnull).TLeaseExpires(tLeaseExpires).TLeaseExpiresAfter(tLeaseExpiresAfter).TLeaseExpiresBefore(tLeaseExpiresBefore).TLeaseExpiresIsnull(tLeaseExpiresIsnull).TRequested(tRequested).TRequestedAfter(tRequestedAfter).TRequestedBefore(tRequestedBefore).TSatisfied(tSatisfied).TSatisfiedAfter(tSatisfiedAfter).TSatisfiedBefore(tSatisfiedBefore).TSatisfiedIsnull(tSatisfiedIsnull).UseWatchdog(useWatchdog).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReservationsApi.ReservationsActionableList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReservationsActionableList`: PaginatedReservationSessionSerializerReadOnlyList
    fmt.Fprintf(os.Stdout, "Response from `ReservationsApi.ReservationsActionableList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReservationsActionableListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aDeviceGroups** | **[]int32** |  | 
 **aDevices** | **[]int32** |  | 
 **details** | **string** |  | 
 **id** | **string** |  | 
 **maxHistory** | **time.Time** | Filter out reservations that expired before the specified datetime. Defaults to 24 hours prior. | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **owner** | **int32** |  | 
 **page** | **int32** | A page number within the paginated result set. | 
 **rDeviceGroupLabels** | **[]int32** |  | 
 **rDeviceGroups** | **[]int32** |  | 
 **rDevices** | **[]int32** |  | 
 **search** | **string** | A search term. | 
 **tCompleted** | **time.Time** |  | 
 **tCompletedAfter** | **time.Time** |  | 
 **tCompletedBefore** | **time.Time** |  | 
 **tCompletedIsnull** | **bool** |  | 
 **tLeaseExpires** | **time.Time** |  | 
 **tLeaseExpiresAfter** | **time.Time** |  | 
 **tLeaseExpiresBefore** | **time.Time** |  | 
 **tLeaseExpiresIsnull** | **bool** |  | 
 **tRequested** | **time.Time** |  | 
 **tRequestedAfter** | **time.Time** |  | 
 **tRequestedBefore** | **time.Time** |  | 
 **tSatisfied** | **time.Time** |  | 
 **tSatisfiedAfter** | **time.Time** |  | 
 **tSatisfiedBefore** | **time.Time** |  | 
 **tSatisfiedIsnull** | **bool** |  | 
 **useWatchdog** | **bool** |  | 

### Return type

[**PaginatedReservationSessionSerializerReadOnlyList**](PaginatedReservationSessionSerializerReadOnlyList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReservationsActiveList

> PaginatedReservationSessionSerializerReadOnlyList ReservationsActiveList(ctx).ADeviceGroups(aDeviceGroups).ADevices(aDevices).Details(details).Id(id).Ordering(ordering).Owner(owner).Page(page).RDeviceGroupLabels(rDeviceGroupLabels).RDeviceGroups(rDeviceGroups).RDevices(rDevices).Search(search).TCompleted(tCompleted).TCompletedAfter(tCompletedAfter).TCompletedBefore(tCompletedBefore).TCompletedIsnull(tCompletedIsnull).TLeaseExpires(tLeaseExpires).TLeaseExpiresAfter(tLeaseExpiresAfter).TLeaseExpiresBefore(tLeaseExpiresBefore).TLeaseExpiresIsnull(tLeaseExpiresIsnull).TRequested(tRequested).TRequestedAfter(tRequestedAfter).TRequestedBefore(tRequestedBefore).TSatisfied(tSatisfied).TSatisfiedAfter(tSatisfiedAfter).TSatisfiedBefore(tSatisfiedBefore).TSatisfiedIsnull(tSatisfiedIsnull).UseWatchdog(useWatchdog).Execute()





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
    aDeviceGroups := []int32{int32(123)} // []int32 |  (optional)
    aDevices := []int32{int32(123)} // []int32 |  (optional)
    details := "details_example" // string |  (optional)
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    owner := int32(56) // int32 |  (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    rDeviceGroupLabels := []int32{int32(123)} // []int32 |  (optional)
    rDeviceGroups := []int32{int32(123)} // []int32 |  (optional)
    rDevices := []int32{int32(123)} // []int32 |  (optional)
    search := "search_example" // string | A search term. (optional)
    tCompleted := time.Now() // time.Time |  (optional)
    tCompletedAfter := time.Now() // time.Time |  (optional)
    tCompletedBefore := time.Now() // time.Time |  (optional)
    tCompletedIsnull := true // bool |  (optional)
    tLeaseExpires := time.Now() // time.Time |  (optional)
    tLeaseExpiresAfter := time.Now() // time.Time |  (optional)
    tLeaseExpiresBefore := time.Now() // time.Time |  (optional)
    tLeaseExpiresIsnull := true // bool |  (optional)
    tRequested := time.Now() // time.Time |  (optional)
    tRequestedAfter := time.Now() // time.Time |  (optional)
    tRequestedBefore := time.Now() // time.Time |  (optional)
    tSatisfied := time.Now() // time.Time |  (optional)
    tSatisfiedAfter := time.Now() // time.Time |  (optional)
    tSatisfiedBefore := time.Now() // time.Time |  (optional)
    tSatisfiedIsnull := true // bool |  (optional)
    useWatchdog := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReservationsApi.ReservationsActiveList(context.Background()).ADeviceGroups(aDeviceGroups).ADevices(aDevices).Details(details).Id(id).Ordering(ordering).Owner(owner).Page(page).RDeviceGroupLabels(rDeviceGroupLabels).RDeviceGroups(rDeviceGroups).RDevices(rDevices).Search(search).TCompleted(tCompleted).TCompletedAfter(tCompletedAfter).TCompletedBefore(tCompletedBefore).TCompletedIsnull(tCompletedIsnull).TLeaseExpires(tLeaseExpires).TLeaseExpiresAfter(tLeaseExpiresAfter).TLeaseExpiresBefore(tLeaseExpiresBefore).TLeaseExpiresIsnull(tLeaseExpiresIsnull).TRequested(tRequested).TRequestedAfter(tRequestedAfter).TRequestedBefore(tRequestedBefore).TSatisfied(tSatisfied).TSatisfiedAfter(tSatisfiedAfter).TSatisfiedBefore(tSatisfiedBefore).TSatisfiedIsnull(tSatisfiedIsnull).UseWatchdog(useWatchdog).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReservationsApi.ReservationsActiveList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReservationsActiveList`: PaginatedReservationSessionSerializerReadOnlyList
    fmt.Fprintf(os.Stdout, "Response from `ReservationsApi.ReservationsActiveList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReservationsActiveListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aDeviceGroups** | **[]int32** |  | 
 **aDevices** | **[]int32** |  | 
 **details** | **string** |  | 
 **id** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **owner** | **int32** |  | 
 **page** | **int32** | A page number within the paginated result set. | 
 **rDeviceGroupLabels** | **[]int32** |  | 
 **rDeviceGroups** | **[]int32** |  | 
 **rDevices** | **[]int32** |  | 
 **search** | **string** | A search term. | 
 **tCompleted** | **time.Time** |  | 
 **tCompletedAfter** | **time.Time** |  | 
 **tCompletedBefore** | **time.Time** |  | 
 **tCompletedIsnull** | **bool** |  | 
 **tLeaseExpires** | **time.Time** |  | 
 **tLeaseExpiresAfter** | **time.Time** |  | 
 **tLeaseExpiresBefore** | **time.Time** |  | 
 **tLeaseExpiresIsnull** | **bool** |  | 
 **tRequested** | **time.Time** |  | 
 **tRequestedAfter** | **time.Time** |  | 
 **tRequestedBefore** | **time.Time** |  | 
 **tSatisfied** | **time.Time** |  | 
 **tSatisfiedAfter** | **time.Time** |  | 
 **tSatisfiedBefore** | **time.Time** |  | 
 **tSatisfiedIsnull** | **bool** |  | 
 **useWatchdog** | **bool** |  | 

### Return type

[**PaginatedReservationSessionSerializerReadOnlyList**](PaginatedReservationSessionSerializerReadOnlyList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReservationsCancelUpdate

> ReservationSessionSerializerReadOnly ReservationsCancelUpdate(ctx, id).Execute()





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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this reservation session.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReservationsApi.ReservationsCancelUpdate(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReservationsApi.ReservationsCancelUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReservationsCancelUpdate`: ReservationSessionSerializerReadOnly
    fmt.Fprintf(os.Stdout, "Response from `ReservationsApi.ReservationsCancelUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this reservation session. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReservationsCancelUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ReservationSessionSerializerReadOnly**](ReservationSessionSerializerReadOnly.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReservationsCreate

> ReservationSessionSerializerReadOnly ReservationsCreate(ctx).ReservationRequest(reservationRequest).Execute()



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
    reservationRequest := *openapiclient.NewReservationRequest() // ReservationRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReservationsApi.ReservationsCreate(context.Background()).ReservationRequest(reservationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReservationsApi.ReservationsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReservationsCreate`: ReservationSessionSerializerReadOnly
    fmt.Fprintf(os.Stdout, "Response from `ReservationsApi.ReservationsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReservationsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reservationRequest** | [**ReservationRequest**](ReservationRequest.md) |  | 

### Return type

[**ReservationSessionSerializerReadOnly**](ReservationSessionSerializerReadOnly.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReservationsList

> PaginatedReservationSessionSerializerReadOnlyList ReservationsList(ctx).ADeviceGroups(aDeviceGroups).ADevices(aDevices).Details(details).Id(id).Ordering(ordering).Owner(owner).Page(page).RDeviceGroupLabels(rDeviceGroupLabels).RDeviceGroups(rDeviceGroups).RDevices(rDevices).Search(search).TCompleted(tCompleted).TCompletedAfter(tCompletedAfter).TCompletedBefore(tCompletedBefore).TCompletedIsnull(tCompletedIsnull).TLeaseExpires(tLeaseExpires).TLeaseExpiresAfter(tLeaseExpiresAfter).TLeaseExpiresBefore(tLeaseExpiresBefore).TLeaseExpiresIsnull(tLeaseExpiresIsnull).TRequested(tRequested).TRequestedAfter(tRequestedAfter).TRequestedBefore(tRequestedBefore).TSatisfied(tSatisfied).TSatisfiedAfter(tSatisfiedAfter).TSatisfiedBefore(tSatisfiedBefore).TSatisfiedIsnull(tSatisfiedIsnull).UseWatchdog(useWatchdog).Execute()



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
    aDeviceGroups := []int32{int32(123)} // []int32 |  (optional)
    aDevices := []int32{int32(123)} // []int32 |  (optional)
    details := "details_example" // string |  (optional)
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    owner := int32(56) // int32 |  (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    rDeviceGroupLabels := []int32{int32(123)} // []int32 |  (optional)
    rDeviceGroups := []int32{int32(123)} // []int32 |  (optional)
    rDevices := []int32{int32(123)} // []int32 |  (optional)
    search := "search_example" // string | A search term. (optional)
    tCompleted := time.Now() // time.Time |  (optional)
    tCompletedAfter := time.Now() // time.Time |  (optional)
    tCompletedBefore := time.Now() // time.Time |  (optional)
    tCompletedIsnull := true // bool |  (optional)
    tLeaseExpires := time.Now() // time.Time |  (optional)
    tLeaseExpiresAfter := time.Now() // time.Time |  (optional)
    tLeaseExpiresBefore := time.Now() // time.Time |  (optional)
    tLeaseExpiresIsnull := true // bool |  (optional)
    tRequested := time.Now() // time.Time |  (optional)
    tRequestedAfter := time.Now() // time.Time |  (optional)
    tRequestedBefore := time.Now() // time.Time |  (optional)
    tSatisfied := time.Now() // time.Time |  (optional)
    tSatisfiedAfter := time.Now() // time.Time |  (optional)
    tSatisfiedBefore := time.Now() // time.Time |  (optional)
    tSatisfiedIsnull := true // bool |  (optional)
    useWatchdog := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReservationsApi.ReservationsList(context.Background()).ADeviceGroups(aDeviceGroups).ADevices(aDevices).Details(details).Id(id).Ordering(ordering).Owner(owner).Page(page).RDeviceGroupLabels(rDeviceGroupLabels).RDeviceGroups(rDeviceGroups).RDevices(rDevices).Search(search).TCompleted(tCompleted).TCompletedAfter(tCompletedAfter).TCompletedBefore(tCompletedBefore).TCompletedIsnull(tCompletedIsnull).TLeaseExpires(tLeaseExpires).TLeaseExpiresAfter(tLeaseExpiresAfter).TLeaseExpiresBefore(tLeaseExpiresBefore).TLeaseExpiresIsnull(tLeaseExpiresIsnull).TRequested(tRequested).TRequestedAfter(tRequestedAfter).TRequestedBefore(tRequestedBefore).TSatisfied(tSatisfied).TSatisfiedAfter(tSatisfiedAfter).TSatisfiedBefore(tSatisfiedBefore).TSatisfiedIsnull(tSatisfiedIsnull).UseWatchdog(useWatchdog).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReservationsApi.ReservationsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReservationsList`: PaginatedReservationSessionSerializerReadOnlyList
    fmt.Fprintf(os.Stdout, "Response from `ReservationsApi.ReservationsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReservationsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aDeviceGroups** | **[]int32** |  | 
 **aDevices** | **[]int32** |  | 
 **details** | **string** |  | 
 **id** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **owner** | **int32** |  | 
 **page** | **int32** | A page number within the paginated result set. | 
 **rDeviceGroupLabels** | **[]int32** |  | 
 **rDeviceGroups** | **[]int32** |  | 
 **rDevices** | **[]int32** |  | 
 **search** | **string** | A search term. | 
 **tCompleted** | **time.Time** |  | 
 **tCompletedAfter** | **time.Time** |  | 
 **tCompletedBefore** | **time.Time** |  | 
 **tCompletedIsnull** | **bool** |  | 
 **tLeaseExpires** | **time.Time** |  | 
 **tLeaseExpiresAfter** | **time.Time** |  | 
 **tLeaseExpiresBefore** | **time.Time** |  | 
 **tLeaseExpiresIsnull** | **bool** |  | 
 **tRequested** | **time.Time** |  | 
 **tRequestedAfter** | **time.Time** |  | 
 **tRequestedBefore** | **time.Time** |  | 
 **tSatisfied** | **time.Time** |  | 
 **tSatisfiedAfter** | **time.Time** |  | 
 **tSatisfiedBefore** | **time.Time** |  | 
 **tSatisfiedIsnull** | **bool** |  | 
 **useWatchdog** | **bool** |  | 

### Return type

[**PaginatedReservationSessionSerializerReadOnlyList**](PaginatedReservationSessionSerializerReadOnlyList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReservationsMetadataPartialUpdate

> ReservationSessionSerializerReadOnly ReservationsMetadataPartialUpdate(ctx, id).RequestBody(requestBody).Execute()





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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this reservation session.
    requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReservationsApi.ReservationsMetadataPartialUpdate(context.Background(), id).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReservationsApi.ReservationsMetadataPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReservationsMetadataPartialUpdate`: ReservationSessionSerializerReadOnly
    fmt.Fprintf(os.Stdout, "Response from `ReservationsApi.ReservationsMetadataPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this reservation session. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReservationsMetadataPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **map[string]interface{}** |  | 

### Return type

[**ReservationSessionSerializerReadOnly**](ReservationSessionSerializerReadOnly.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReservationsMetadataUpdate

> ReservationSessionSerializerReadOnly ReservationsMetadataUpdate(ctx, id).RequestBody(requestBody).Execute()





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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this reservation session.
    requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReservationsApi.ReservationsMetadataUpdate(context.Background(), id).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReservationsApi.ReservationsMetadataUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReservationsMetadataUpdate`: ReservationSessionSerializerReadOnly
    fmt.Fprintf(os.Stdout, "Response from `ReservationsApi.ReservationsMetadataUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this reservation session. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReservationsMetadataUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **map[string]interface{}** |  | 

### Return type

[**ReservationSessionSerializerReadOnly**](ReservationSessionSerializerReadOnly.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReservationsReleaseUpdate

> ReservationSessionSerializerReadOnly ReservationsReleaseUpdate(ctx, id).Execute()





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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this reservation session.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReservationsApi.ReservationsReleaseUpdate(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReservationsApi.ReservationsReleaseUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReservationsReleaseUpdate`: ReservationSessionSerializerReadOnly
    fmt.Fprintf(os.Stdout, "Response from `ReservationsApi.ReservationsReleaseUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this reservation session. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReservationsReleaseUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ReservationSessionSerializerReadOnly**](ReservationSessionSerializerReadOnly.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReservationsRetrieve

> ReservationSessionSerializerReadOnly ReservationsRetrieve(ctx, id).Execute()



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this reservation session.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReservationsApi.ReservationsRetrieve(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReservationsApi.ReservationsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReservationsRetrieve`: ReservationSessionSerializerReadOnly
    fmt.Fprintf(os.Stdout, "Response from `ReservationsApi.ReservationsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this reservation session. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReservationsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ReservationSessionSerializerReadOnly**](ReservationSessionSerializerReadOnly.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReservationsUpdate

> ReservationSessionSerializerReadOnly ReservationsUpdate(ctx, id).Execute()





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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this reservation session.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReservationsApi.ReservationsUpdate(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReservationsApi.ReservationsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReservationsUpdate`: ReservationSessionSerializerReadOnly
    fmt.Fprintf(os.Stdout, "Response from `ReservationsApi.ReservationsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this reservation session. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReservationsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ReservationSessionSerializerReadOnly**](ReservationSessionSerializerReadOnly.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

