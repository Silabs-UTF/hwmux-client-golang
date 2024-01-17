# \TokenAuthApi

All URIs are relative to *https://hwmux.silabs.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TokenAuthCreate**](TokenAuthApi.md#TokenAuthCreate) | **Post** /api-token-auth/ | 



## TokenAuthCreate

> AuthToken TokenAuthCreate(ctx).Username(username).Password(password).Token(token).Execute()



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
    password := "password_example" // string | 
    token := "token_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TokenAuthApi.TokenAuthCreate(context.Background()).Username(username).Password(password).Token(token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokenAuthApi.TokenAuthCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TokenAuthCreate`: AuthToken
    fmt.Fprintf(os.Stdout, "Response from `TokenAuthApi.TokenAuthCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTokenAuthCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **string** |  | 
 **password** | **string** |  | 
 **token** | **string** |  | 

### Return type

[**AuthToken**](AuthToken.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, multipart/form-data, application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

