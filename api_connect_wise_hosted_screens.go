/*
Connectwise Manage Public Endpoints

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2025.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cwapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// ConnectWiseHostedScreensAPIService ConnectWiseHostedScreensAPI service
type ConnectWiseHostedScreensAPIService service

type ApiGetSystemConnectWiseHostedScreensRequest struct {
	ctx context.Context
	ApiService *ConnectWiseHostedScreensAPIService
	conditions *string
	childConditions *string
	customFieldConditions *string
	orderBy *string
	fields *string
	page *int32
	pageSize *int32
	pageId *int32
	clientId *string
}

// 
func (r ApiGetSystemConnectWiseHostedScreensRequest) Conditions(conditions string) ApiGetSystemConnectWiseHostedScreensRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetSystemConnectWiseHostedScreensRequest) ChildConditions(childConditions string) ApiGetSystemConnectWiseHostedScreensRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetSystemConnectWiseHostedScreensRequest) CustomFieldConditions(customFieldConditions string) ApiGetSystemConnectWiseHostedScreensRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetSystemConnectWiseHostedScreensRequest) OrderBy(orderBy string) ApiGetSystemConnectWiseHostedScreensRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetSystemConnectWiseHostedScreensRequest) Fields(fields string) ApiGetSystemConnectWiseHostedScreensRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetSystemConnectWiseHostedScreensRequest) Page(page int32) ApiGetSystemConnectWiseHostedScreensRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetSystemConnectWiseHostedScreensRequest) PageSize(pageSize int32) ApiGetSystemConnectWiseHostedScreensRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetSystemConnectWiseHostedScreensRequest) PageId(pageId int32) ApiGetSystemConnectWiseHostedScreensRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetSystemConnectWiseHostedScreensRequest) ClientId(clientId string) ApiGetSystemConnectWiseHostedScreensRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetSystemConnectWiseHostedScreensRequest) Execute() ([]ConnectWiseHostedScreen, *http.Response, error) {
	return r.ApiService.GetSystemConnectWiseHostedScreensExecute(r)
}

/*
GetSystemConnectWiseHostedScreens Get List of ConnectWiseHostedScreen

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetSystemConnectWiseHostedScreensRequest
*/
func (a *ConnectWiseHostedScreensAPIService) GetSystemConnectWiseHostedScreens(ctx context.Context) ApiGetSystemConnectWiseHostedScreensRequest {
	return ApiGetSystemConnectWiseHostedScreensRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []ConnectWiseHostedScreen
func (a *ConnectWiseHostedScreensAPIService) GetSystemConnectWiseHostedScreensExecute(r ApiGetSystemConnectWiseHostedScreensRequest) ([]ConnectWiseHostedScreen, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ConnectWiseHostedScreen
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConnectWiseHostedScreensAPIService.GetSystemConnectWiseHostedScreens")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/connectWiseHostedScreens"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.conditions != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "conditions", r.conditions, "form", "")
	}
	if r.childConditions != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "childConditions", r.childConditions, "form", "")
	}
	if r.customFieldConditions != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "customFieldConditions", r.customFieldConditions, "form", "")
	}
	if r.orderBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "orderBy", r.orderBy, "form", "")
	}
	if r.fields != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields", r.fields, "form", "")
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageSize", r.pageSize, "form", "")
	}
	if r.pageId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageId", r.pageId, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.connectwise.com+json; version=2025.1"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.clientId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "clientId", r.clientId, "simple", "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiGetSystemConnectWiseHostedScreensByIdRequest struct {
	ctx context.Context
	ApiService *ConnectWiseHostedScreensAPIService
	id int32
	conditions *string
	childConditions *string
	customFieldConditions *string
	orderBy *string
	fields *string
	page *int32
	pageSize *int32
	pageId *int32
	clientId *string
}

// 
func (r ApiGetSystemConnectWiseHostedScreensByIdRequest) Conditions(conditions string) ApiGetSystemConnectWiseHostedScreensByIdRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetSystemConnectWiseHostedScreensByIdRequest) ChildConditions(childConditions string) ApiGetSystemConnectWiseHostedScreensByIdRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetSystemConnectWiseHostedScreensByIdRequest) CustomFieldConditions(customFieldConditions string) ApiGetSystemConnectWiseHostedScreensByIdRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetSystemConnectWiseHostedScreensByIdRequest) OrderBy(orderBy string) ApiGetSystemConnectWiseHostedScreensByIdRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetSystemConnectWiseHostedScreensByIdRequest) Fields(fields string) ApiGetSystemConnectWiseHostedScreensByIdRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetSystemConnectWiseHostedScreensByIdRequest) Page(page int32) ApiGetSystemConnectWiseHostedScreensByIdRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetSystemConnectWiseHostedScreensByIdRequest) PageSize(pageSize int32) ApiGetSystemConnectWiseHostedScreensByIdRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetSystemConnectWiseHostedScreensByIdRequest) PageId(pageId int32) ApiGetSystemConnectWiseHostedScreensByIdRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetSystemConnectWiseHostedScreensByIdRequest) ClientId(clientId string) ApiGetSystemConnectWiseHostedScreensByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetSystemConnectWiseHostedScreensByIdRequest) Execute() (*ConnectWiseHostedScreen, *http.Response, error) {
	return r.ApiService.GetSystemConnectWiseHostedScreensByIdExecute(r)
}

/*
GetSystemConnectWiseHostedScreensById Get ConnectWiseHostedScreen

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id connectWiseHostedScreenId
 @return ApiGetSystemConnectWiseHostedScreensByIdRequest
*/
func (a *ConnectWiseHostedScreensAPIService) GetSystemConnectWiseHostedScreensById(ctx context.Context, id int32) ApiGetSystemConnectWiseHostedScreensByIdRequest {
	return ApiGetSystemConnectWiseHostedScreensByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ConnectWiseHostedScreen
func (a *ConnectWiseHostedScreensAPIService) GetSystemConnectWiseHostedScreensByIdExecute(r ApiGetSystemConnectWiseHostedScreensByIdRequest) (*ConnectWiseHostedScreen, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ConnectWiseHostedScreen
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConnectWiseHostedScreensAPIService.GetSystemConnectWiseHostedScreensById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/connectWiseHostedScreens/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.conditions != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "conditions", r.conditions, "form", "")
	}
	if r.childConditions != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "childConditions", r.childConditions, "form", "")
	}
	if r.customFieldConditions != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "customFieldConditions", r.customFieldConditions, "form", "")
	}
	if r.orderBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "orderBy", r.orderBy, "form", "")
	}
	if r.fields != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields", r.fields, "form", "")
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageSize", r.pageSize, "form", "")
	}
	if r.pageId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageId", r.pageId, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.connectwise.com+json; version=2025.1"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.clientId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "clientId", r.clientId, "simple", "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiGetSystemConnectWiseHostedScreensCountRequest struct {
	ctx context.Context
	ApiService *ConnectWiseHostedScreensAPIService
	conditions *string
	childConditions *string
	customFieldConditions *string
	orderBy *string
	fields *string
	page *int32
	pageSize *int32
	pageId *int32
	clientId *string
}

// 
func (r ApiGetSystemConnectWiseHostedScreensCountRequest) Conditions(conditions string) ApiGetSystemConnectWiseHostedScreensCountRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetSystemConnectWiseHostedScreensCountRequest) ChildConditions(childConditions string) ApiGetSystemConnectWiseHostedScreensCountRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetSystemConnectWiseHostedScreensCountRequest) CustomFieldConditions(customFieldConditions string) ApiGetSystemConnectWiseHostedScreensCountRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetSystemConnectWiseHostedScreensCountRequest) OrderBy(orderBy string) ApiGetSystemConnectWiseHostedScreensCountRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetSystemConnectWiseHostedScreensCountRequest) Fields(fields string) ApiGetSystemConnectWiseHostedScreensCountRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetSystemConnectWiseHostedScreensCountRequest) Page(page int32) ApiGetSystemConnectWiseHostedScreensCountRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetSystemConnectWiseHostedScreensCountRequest) PageSize(pageSize int32) ApiGetSystemConnectWiseHostedScreensCountRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetSystemConnectWiseHostedScreensCountRequest) PageId(pageId int32) ApiGetSystemConnectWiseHostedScreensCountRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetSystemConnectWiseHostedScreensCountRequest) ClientId(clientId string) ApiGetSystemConnectWiseHostedScreensCountRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetSystemConnectWiseHostedScreensCountRequest) Execute() (*Count, *http.Response, error) {
	return r.ApiService.GetSystemConnectWiseHostedScreensCountExecute(r)
}

/*
GetSystemConnectWiseHostedScreensCount Get Count of ConnectWiseHostedScreen

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetSystemConnectWiseHostedScreensCountRequest
*/
func (a *ConnectWiseHostedScreensAPIService) GetSystemConnectWiseHostedScreensCount(ctx context.Context) ApiGetSystemConnectWiseHostedScreensCountRequest {
	return ApiGetSystemConnectWiseHostedScreensCountRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Count
func (a *ConnectWiseHostedScreensAPIService) GetSystemConnectWiseHostedScreensCountExecute(r ApiGetSystemConnectWiseHostedScreensCountRequest) (*Count, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Count
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConnectWiseHostedScreensAPIService.GetSystemConnectWiseHostedScreensCount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/connectWiseHostedScreens/count"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.conditions != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "conditions", r.conditions, "form", "")
	}
	if r.childConditions != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "childConditions", r.childConditions, "form", "")
	}
	if r.customFieldConditions != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "customFieldConditions", r.customFieldConditions, "form", "")
	}
	if r.orderBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "orderBy", r.orderBy, "form", "")
	}
	if r.fields != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields", r.fields, "form", "")
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageSize", r.pageSize, "form", "")
	}
	if r.pageId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageId", r.pageId, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.connectwise.com+json; version=2025.1"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.clientId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "clientId", r.clientId, "simple", "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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
