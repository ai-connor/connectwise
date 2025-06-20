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


// AuthAnvilsAPIService AuthAnvilsAPI service
type AuthAnvilsAPIService service

type ApiGetSystemAuthAnvilsRequest struct {
	ctx context.Context
	ApiService *AuthAnvilsAPIService
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
func (r ApiGetSystemAuthAnvilsRequest) Conditions(conditions string) ApiGetSystemAuthAnvilsRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetSystemAuthAnvilsRequest) ChildConditions(childConditions string) ApiGetSystemAuthAnvilsRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetSystemAuthAnvilsRequest) CustomFieldConditions(customFieldConditions string) ApiGetSystemAuthAnvilsRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetSystemAuthAnvilsRequest) OrderBy(orderBy string) ApiGetSystemAuthAnvilsRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetSystemAuthAnvilsRequest) Fields(fields string) ApiGetSystemAuthAnvilsRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetSystemAuthAnvilsRequest) Page(page int32) ApiGetSystemAuthAnvilsRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetSystemAuthAnvilsRequest) PageSize(pageSize int32) ApiGetSystemAuthAnvilsRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetSystemAuthAnvilsRequest) PageId(pageId int32) ApiGetSystemAuthAnvilsRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetSystemAuthAnvilsRequest) ClientId(clientId string) ApiGetSystemAuthAnvilsRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetSystemAuthAnvilsRequest) Execute() ([]AuthAnvil, *http.Response, error) {
	return r.ApiService.GetSystemAuthAnvilsExecute(r)
}

/*
GetSystemAuthAnvils Get List of ConnectWise.Apis.v3_0.v2015_3.System.AuthAnvil.AuthAnvil

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetSystemAuthAnvilsRequest
*/
func (a *AuthAnvilsAPIService) GetSystemAuthAnvils(ctx context.Context) ApiGetSystemAuthAnvilsRequest {
	return ApiGetSystemAuthAnvilsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []AuthAnvil
func (a *AuthAnvilsAPIService) GetSystemAuthAnvilsExecute(r ApiGetSystemAuthAnvilsRequest) ([]AuthAnvil, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []AuthAnvil
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthAnvilsAPIService.GetSystemAuthAnvils")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/authAnvils"

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

type ApiGetSystemAuthAnvilsByIdRequest struct {
	ctx context.Context
	ApiService *AuthAnvilsAPIService
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
func (r ApiGetSystemAuthAnvilsByIdRequest) Conditions(conditions string) ApiGetSystemAuthAnvilsByIdRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetSystemAuthAnvilsByIdRequest) ChildConditions(childConditions string) ApiGetSystemAuthAnvilsByIdRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetSystemAuthAnvilsByIdRequest) CustomFieldConditions(customFieldConditions string) ApiGetSystemAuthAnvilsByIdRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetSystemAuthAnvilsByIdRequest) OrderBy(orderBy string) ApiGetSystemAuthAnvilsByIdRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetSystemAuthAnvilsByIdRequest) Fields(fields string) ApiGetSystemAuthAnvilsByIdRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetSystemAuthAnvilsByIdRequest) Page(page int32) ApiGetSystemAuthAnvilsByIdRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetSystemAuthAnvilsByIdRequest) PageSize(pageSize int32) ApiGetSystemAuthAnvilsByIdRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetSystemAuthAnvilsByIdRequest) PageId(pageId int32) ApiGetSystemAuthAnvilsByIdRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetSystemAuthAnvilsByIdRequest) ClientId(clientId string) ApiGetSystemAuthAnvilsByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetSystemAuthAnvilsByIdRequest) Execute() (*AuthAnvil, *http.Response, error) {
	return r.ApiService.GetSystemAuthAnvilsByIdExecute(r)
}

/*
GetSystemAuthAnvilsById Get ConnectWise.Apis.v3_0.v2015_3.System.AuthAnvil.AuthAnvil

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id authAnvilId
 @return ApiGetSystemAuthAnvilsByIdRequest
*/
func (a *AuthAnvilsAPIService) GetSystemAuthAnvilsById(ctx context.Context, id int32) ApiGetSystemAuthAnvilsByIdRequest {
	return ApiGetSystemAuthAnvilsByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return AuthAnvil
func (a *AuthAnvilsAPIService) GetSystemAuthAnvilsByIdExecute(r ApiGetSystemAuthAnvilsByIdRequest) (*AuthAnvil, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AuthAnvil
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthAnvilsAPIService.GetSystemAuthAnvilsById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/authAnvils/{id}"
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

type ApiGetSystemAuthAnvilsCountRequest struct {
	ctx context.Context
	ApiService *AuthAnvilsAPIService
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
func (r ApiGetSystemAuthAnvilsCountRequest) Conditions(conditions string) ApiGetSystemAuthAnvilsCountRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetSystemAuthAnvilsCountRequest) ChildConditions(childConditions string) ApiGetSystemAuthAnvilsCountRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetSystemAuthAnvilsCountRequest) CustomFieldConditions(customFieldConditions string) ApiGetSystemAuthAnvilsCountRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetSystemAuthAnvilsCountRequest) OrderBy(orderBy string) ApiGetSystemAuthAnvilsCountRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetSystemAuthAnvilsCountRequest) Fields(fields string) ApiGetSystemAuthAnvilsCountRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetSystemAuthAnvilsCountRequest) Page(page int32) ApiGetSystemAuthAnvilsCountRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetSystemAuthAnvilsCountRequest) PageSize(pageSize int32) ApiGetSystemAuthAnvilsCountRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetSystemAuthAnvilsCountRequest) PageId(pageId int32) ApiGetSystemAuthAnvilsCountRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetSystemAuthAnvilsCountRequest) ClientId(clientId string) ApiGetSystemAuthAnvilsCountRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetSystemAuthAnvilsCountRequest) Execute() (*Count, *http.Response, error) {
	return r.ApiService.GetSystemAuthAnvilsCountExecute(r)
}

/*
GetSystemAuthAnvilsCount Get Count of SuccessResponse

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetSystemAuthAnvilsCountRequest
*/
func (a *AuthAnvilsAPIService) GetSystemAuthAnvilsCount(ctx context.Context) ApiGetSystemAuthAnvilsCountRequest {
	return ApiGetSystemAuthAnvilsCountRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Count
func (a *AuthAnvilsAPIService) GetSystemAuthAnvilsCountExecute(r ApiGetSystemAuthAnvilsCountRequest) (*Count, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Count
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthAnvilsAPIService.GetSystemAuthAnvilsCount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/authAnvils/count"

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

type ApiGetSystemAuthAnvilsTestConnectionRequest struct {
	ctx context.Context
	ApiService *AuthAnvilsAPIService
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
func (r ApiGetSystemAuthAnvilsTestConnectionRequest) Conditions(conditions string) ApiGetSystemAuthAnvilsTestConnectionRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetSystemAuthAnvilsTestConnectionRequest) ChildConditions(childConditions string) ApiGetSystemAuthAnvilsTestConnectionRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetSystemAuthAnvilsTestConnectionRequest) CustomFieldConditions(customFieldConditions string) ApiGetSystemAuthAnvilsTestConnectionRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetSystemAuthAnvilsTestConnectionRequest) OrderBy(orderBy string) ApiGetSystemAuthAnvilsTestConnectionRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetSystemAuthAnvilsTestConnectionRequest) Fields(fields string) ApiGetSystemAuthAnvilsTestConnectionRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetSystemAuthAnvilsTestConnectionRequest) Page(page int32) ApiGetSystemAuthAnvilsTestConnectionRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetSystemAuthAnvilsTestConnectionRequest) PageSize(pageSize int32) ApiGetSystemAuthAnvilsTestConnectionRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetSystemAuthAnvilsTestConnectionRequest) PageId(pageId int32) ApiGetSystemAuthAnvilsTestConnectionRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetSystemAuthAnvilsTestConnectionRequest) ClientId(clientId string) ApiGetSystemAuthAnvilsTestConnectionRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetSystemAuthAnvilsTestConnectionRequest) Execute() (*SuccessResponse, *http.Response, error) {
	return r.ApiService.GetSystemAuthAnvilsTestConnectionExecute(r)
}

/*
GetSystemAuthAnvilsTestConnection Get SuccessResponse

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetSystemAuthAnvilsTestConnectionRequest
*/
func (a *AuthAnvilsAPIService) GetSystemAuthAnvilsTestConnection(ctx context.Context) ApiGetSystemAuthAnvilsTestConnectionRequest {
	return ApiGetSystemAuthAnvilsTestConnectionRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SuccessResponse
func (a *AuthAnvilsAPIService) GetSystemAuthAnvilsTestConnectionExecute(r ApiGetSystemAuthAnvilsTestConnectionRequest) (*SuccessResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SuccessResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthAnvilsAPIService.GetSystemAuthAnvilsTestConnection")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/authAnvils/testConnection"

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

type ApiPatchSystemAuthAnvilsByIdRequest struct {
	ctx context.Context
	ApiService *AuthAnvilsAPIService
	id int32
	patchOperation *[]PatchOperation
	clientId *string
}

// List of PatchOperation
func (r ApiPatchSystemAuthAnvilsByIdRequest) PatchOperation(patchOperation []PatchOperation) ApiPatchSystemAuthAnvilsByIdRequest {
	r.patchOperation = &patchOperation
	return r
}

// 
func (r ApiPatchSystemAuthAnvilsByIdRequest) ClientId(clientId string) ApiPatchSystemAuthAnvilsByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiPatchSystemAuthAnvilsByIdRequest) Execute() (*AuthAnvil, *http.Response, error) {
	return r.ApiService.PatchSystemAuthAnvilsByIdExecute(r)
}

/*
PatchSystemAuthAnvilsById Patch ConnectWise.Apis.v3_0.v2015_3.System.AuthAnvil.AuthAnvil

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id authAnvilId
 @return ApiPatchSystemAuthAnvilsByIdRequest
*/
func (a *AuthAnvilsAPIService) PatchSystemAuthAnvilsById(ctx context.Context, id int32) ApiPatchSystemAuthAnvilsByIdRequest {
	return ApiPatchSystemAuthAnvilsByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return AuthAnvil
func (a *AuthAnvilsAPIService) PatchSystemAuthAnvilsByIdExecute(r ApiPatchSystemAuthAnvilsByIdRequest) (*AuthAnvil, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AuthAnvil
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthAnvilsAPIService.PatchSystemAuthAnvilsById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/authAnvils/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.patchOperation == nil {
		return localVarReturnValue, nil, reportError("patchOperation is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.patchOperation
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

type ApiPutSystemAuthAnvilsByIdRequest struct {
	ctx context.Context
	ApiService *AuthAnvilsAPIService
	id int32
	authAnvil *AuthAnvil
	clientId *string
}

// authAnvil
func (r ApiPutSystemAuthAnvilsByIdRequest) AuthAnvil(authAnvil AuthAnvil) ApiPutSystemAuthAnvilsByIdRequest {
	r.authAnvil = &authAnvil
	return r
}

// 
func (r ApiPutSystemAuthAnvilsByIdRequest) ClientId(clientId string) ApiPutSystemAuthAnvilsByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiPutSystemAuthAnvilsByIdRequest) Execute() (*AuthAnvil, *http.Response, error) {
	return r.ApiService.PutSystemAuthAnvilsByIdExecute(r)
}

/*
PutSystemAuthAnvilsById Put ConnectWise.Apis.v3_0.v2015_3.System.AuthAnvil.AuthAnvil

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id authAnvilId
 @return ApiPutSystemAuthAnvilsByIdRequest
*/
func (a *AuthAnvilsAPIService) PutSystemAuthAnvilsById(ctx context.Context, id int32) ApiPutSystemAuthAnvilsByIdRequest {
	return ApiPutSystemAuthAnvilsByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return AuthAnvil
func (a *AuthAnvilsAPIService) PutSystemAuthAnvilsByIdExecute(r ApiPutSystemAuthAnvilsByIdRequest) (*AuthAnvil, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AuthAnvil
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthAnvilsAPIService.PutSystemAuthAnvilsById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/authAnvils/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.authAnvil == nil {
		return localVarReturnValue, nil, reportError("authAnvil is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.authAnvil
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
