# \OutboxApi

All URIs are relative to *https://hwmux.silabs.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OutboxList**](OutboxApi.md#OutboxList) | **Get** /api/outbox/ | 
[**OutboxRetrieve**](OutboxApi.md#OutboxRetrieve) | **Get** /api/outbox/{id}/ | 



## OutboxList

> PaginatedOutboxList OutboxList(ctx).Ordering(ordering).Page(page).ReservationId(reservationId).SagaId(sagaId).Search(search).Execute()



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
    reservationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
    sagaId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
    search := "search_example" // string | A search term. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OutboxApi.OutboxList(context.Background()).Ordering(ordering).Page(page).ReservationId(reservationId).SagaId(sagaId).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OutboxApi.OutboxList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OutboxList`: PaginatedOutboxList
    fmt.Fprintf(os.Stdout, "Response from `OutboxApi.OutboxList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOutboxListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **reservationId** | **string** |  | 
 **sagaId** | **string** |  | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedOutboxList**](PaginatedOutboxList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OutboxRetrieve

> Outbox OutboxRetrieve(ctx, id).Execute()



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Reservation Outbox.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OutboxApi.OutboxRetrieve(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OutboxApi.OutboxRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OutboxRetrieve`: Outbox
    fmt.Fprintf(os.Stdout, "Response from `OutboxApi.OutboxRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this Reservation Outbox. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOutboxRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Outbox**](Outbox.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

