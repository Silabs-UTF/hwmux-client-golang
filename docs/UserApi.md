# \UserApi

All URIs are relative to *https://hwmux.silabs.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserCreate**](UserApi.md#UserCreate) | **Post** /api/user/ | 
[**UserCurrentRetrieve**](UserApi.md#UserCurrentRetrieve) | **Get** /api/user/current/ | 
[**UserCurrentTokenCreate**](UserApi.md#UserCurrentTokenCreate) | **Post** /api/user/current/token/ | 
[**UserCurrentTokenRetrieve**](UserApi.md#UserCurrentTokenRetrieve) | **Get** /api/user/current/token/ | 
[**UserDestroy**](UserApi.md#UserDestroy) | **Delete** /api/user/{username}/ | 
[**UserList**](UserApi.md#UserList) | **Get** /api/user/ | 
[**UserPartialUpdate**](UserApi.md#UserPartialUpdate) | **Patch** /api/user/{username}/ | 
[**UserRetrieve**](UserApi.md#UserRetrieve) | **Get** /api/user/{username}/ | 
[**UserUpdate**](UserApi.md#UserUpdate) | **Put** /api/user/{username}/ | 



## UserCreate

> LoggedInUser UserCreate(ctx).LoggedInUser(loggedInUser).Execute()



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
    loggedInUser := *openapiclient.NewLoggedInUser("Username_example", []string{"Groups_example"}) // LoggedInUser | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.UserCreate(context.Background()).LoggedInUser(loggedInUser).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.UserCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserCreate`: LoggedInUser
    fmt.Fprintf(os.Stdout, "Response from `UserApi.UserCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **loggedInUser** | [**LoggedInUser**](LoggedInUser.md) |  | 

### Return type

[**LoggedInUser**](LoggedInUser.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserCurrentRetrieve

> LoggedInUser UserCurrentRetrieve(ctx).Execute()



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.UserCurrentRetrieve(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.UserCurrentRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserCurrentRetrieve`: LoggedInUser
    fmt.Fprintf(os.Stdout, "Response from `UserApi.UserCurrentRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUserCurrentRetrieveRequest struct via the builder pattern


### Return type

[**LoggedInUser**](LoggedInUser.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserCurrentTokenCreate

> Token UserCurrentTokenCreate(ctx).Execute()





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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.UserCurrentTokenCreate(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.UserCurrentTokenCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserCurrentTokenCreate`: Token
    fmt.Fprintf(os.Stdout, "Response from `UserApi.UserCurrentTokenCreate`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUserCurrentTokenCreateRequest struct via the builder pattern


### Return type

[**Token**](Token.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserCurrentTokenRetrieve

> Token UserCurrentTokenRetrieve(ctx).Execute()





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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.UserCurrentTokenRetrieve(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.UserCurrentTokenRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserCurrentTokenRetrieve`: Token
    fmt.Fprintf(os.Stdout, "Response from `UserApi.UserCurrentTokenRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUserCurrentTokenRetrieveRequest struct via the builder pattern


### Return type

[**Token**](Token.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserDestroy

> UserDestroy(ctx, username).Execute()



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
    username := "username_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.UserDestroy(context.Background(), username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.UserDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserDestroyRequest struct via the builder pattern


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


## UserList

> PaginatedLoggedInUserList UserList(ctx).Email(email).FirstName(firstName).LastName(lastName).Ordering(ordering).Page(page).Search(search).UserGroup(userGroup).Username(username).Execute()



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
    email := "email_example" // string |  (optional)
    firstName := "firstName_example" // string |  (optional)
    lastName := "lastName_example" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    search := "search_example" // string | A search term. (optional)
    userGroup := "userGroup_example" // string |  (optional)
    username := "username_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.UserList(context.Background()).Email(email).FirstName(firstName).LastName(lastName).Ordering(ordering).Page(page).Search(search).UserGroup(userGroup).Username(username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.UserList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserList`: PaginatedLoggedInUserList
    fmt.Fprintf(os.Stdout, "Response from `UserApi.UserList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** |  | 
 **firstName** | **string** |  | 
 **lastName** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **search** | **string** | A search term. | 
 **userGroup** | **string** |  | 
 **username** | **string** |  | 

### Return type

[**PaginatedLoggedInUserList**](PaginatedLoggedInUserList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserPartialUpdate

> LoggedInUser UserPartialUpdate(ctx, username).PatchedLoggedInUser(patchedLoggedInUser).Execute()



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
    username := "username_example" // string | 
    patchedLoggedInUser := *openapiclient.NewPatchedLoggedInUser() // PatchedLoggedInUser |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.UserPartialUpdate(context.Background(), username).PatchedLoggedInUser(patchedLoggedInUser).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.UserPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserPartialUpdate`: LoggedInUser
    fmt.Fprintf(os.Stdout, "Response from `UserApi.UserPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedLoggedInUser** | [**PatchedLoggedInUser**](PatchedLoggedInUser.md) |  | 

### Return type

[**LoggedInUser**](LoggedInUser.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserRetrieve

> LoggedInUser UserRetrieve(ctx, username).Execute()



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
    username := "username_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.UserRetrieve(context.Background(), username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.UserRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserRetrieve`: LoggedInUser
    fmt.Fprintf(os.Stdout, "Response from `UserApi.UserRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LoggedInUser**](LoggedInUser.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserUpdate

> LoggedInUser UserUpdate(ctx, username).LoggedInUser(loggedInUser).Execute()



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
    username := "username_example" // string | 
    loggedInUser := *openapiclient.NewLoggedInUser("Username_example", []string{"Groups_example"}) // LoggedInUser | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.UserUpdate(context.Background(), username).LoggedInUser(loggedInUser).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.UserUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserUpdate`: LoggedInUser
    fmt.Fprintf(os.Stdout, "Response from `UserApi.UserUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **loggedInUser** | [**LoggedInUser**](LoggedInUser.md) |  | 

### Return type

[**LoggedInUser**](LoggedInUser.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

