/*
HW Mux Reservation System

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.37.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package hwmux

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// PartsApiService PartsApi service
type PartsApiService service

type ApiPartsListRequest struct {
	ctx context.Context
	ApiService *PartsApiService
	boardNo *string
	chipNo *string
	ordering *string
	page *int32
	partFamily *string
	partNo *string
	revision *string
	search *string
	variant *string
}

func (r ApiPartsListRequest) BoardNo(boardNo string) ApiPartsListRequest {
	r.boardNo = &boardNo
	return r
}

func (r ApiPartsListRequest) ChipNo(chipNo string) ApiPartsListRequest {
	r.chipNo = &chipNo
	return r
}

// Which field to use when ordering the results.
func (r ApiPartsListRequest) Ordering(ordering string) ApiPartsListRequest {
	r.ordering = &ordering
	return r
}

// A page number within the paginated result set.
func (r ApiPartsListRequest) Page(page int32) ApiPartsListRequest {
	r.page = &page
	return r
}

func (r ApiPartsListRequest) PartFamily(partFamily string) ApiPartsListRequest {
	r.partFamily = &partFamily
	return r
}

func (r ApiPartsListRequest) PartNo(partNo string) ApiPartsListRequest {
	r.partNo = &partNo
	return r
}

func (r ApiPartsListRequest) Revision(revision string) ApiPartsListRequest {
	r.revision = &revision
	return r
}

// A search term.
func (r ApiPartsListRequest) Search(search string) ApiPartsListRequest {
	r.search = &search
	return r
}

func (r ApiPartsListRequest) Variant(variant string) ApiPartsListRequest {
	r.variant = &variant
	return r
}

func (r ApiPartsListRequest) Execute() (*PaginatedPartList, *http.Response, error) {
	return r.ApiService.PartsListExecute(r)
}

/*
PartsList Method for PartsList

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPartsListRequest
*/
func (a *PartsApiService) PartsList(ctx context.Context) ApiPartsListRequest {
	return ApiPartsListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedPartList
func (a *PartsApiService) PartsListExecute(r ApiPartsListRequest) (*PaginatedPartList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedPartList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PartsApiService.PartsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/parts/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.boardNo != nil {
		localVarQueryParams.Add("board_no", parameterToString(*r.boardNo, ""))
	}
	if r.chipNo != nil {
		localVarQueryParams.Add("chip_no", parameterToString(*r.chipNo, ""))
	}
	if r.ordering != nil {
		localVarQueryParams.Add("ordering", parameterToString(*r.ordering, ""))
	}
	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	if r.partFamily != nil {
		localVarQueryParams.Add("part_family", parameterToString(*r.partFamily, ""))
	}
	if r.partNo != nil {
		localVarQueryParams.Add("part_no", parameterToString(*r.partNo, ""))
	}
	if r.revision != nil {
		localVarQueryParams.Add("revision", parameterToString(*r.revision, ""))
	}
	if r.search != nil {
		localVarQueryParams.Add("search", parameterToString(*r.search, ""))
	}
	if r.variant != nil {
		localVarQueryParams.Add("variant", parameterToString(*r.variant, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["tokenAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPartsRetrieveRequest struct {
	ctx context.Context
	ApiService *PartsApiService
	partNo string
}

func (r ApiPartsRetrieveRequest) Execute() (*Part, *http.Response, error) {
	return r.ApiService.PartsRetrieveExecute(r)
}

/*
PartsRetrieve Method for PartsRetrieve

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param partNo A unique value identifying this part.
 @return ApiPartsRetrieveRequest
*/
func (a *PartsApiService) PartsRetrieve(ctx context.Context, partNo string) ApiPartsRetrieveRequest {
	return ApiPartsRetrieveRequest{
		ApiService: a,
		ctx: ctx,
		partNo: partNo,
	}
}

// Execute executes the request
//  @return Part
func (a *PartsApiService) PartsRetrieveExecute(r ApiPartsRetrieveRequest) (*Part, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Part
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PartsApiService.PartsRetrieve")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/parts/{part_no}/"
	localVarPath = strings.Replace(localVarPath, "{"+"part_no"+"}", url.PathEscape(parameterToString(r.partNo, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["tokenAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
