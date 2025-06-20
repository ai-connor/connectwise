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


// InOutBoardsAPIService InOutBoardsAPI service
type InOutBoardsAPIService service

type ApiDeleteSystemInOutBoardsByIdRequest struct {
	ctx context.Context
	ApiService *InOutBoardsAPIService
	id int32
	clientId *string
}

// 
func (r ApiDeleteSystemInOutBoardsByIdRequest) ClientId(clientId string) ApiDeleteSystemInOutBoardsByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiDeleteSystemInOutBoardsByIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteSystemInOutBoardsByIdExecute(r)
}

/*
DeleteSystemInOutBoardsById Delete InOutBoard

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id inOutBoardId
 @return ApiDeleteSystemInOutBoardsByIdRequest
*/
func (a *InOutBoardsAPIService) DeleteSystemInOutBoardsById(ctx context.Context, id int32) ApiDeleteSystemInOutBoardsByIdRequest {
	return ApiDeleteSystemInOutBoardsByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *InOutBoardsAPIService) DeleteSystemInOutBoardsByIdExecute(r ApiDeleteSystemInOutBoardsByIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InOutBoardsAPIService.DeleteSystemInOutBoardsById")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/inOutBoards/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

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
	localVarHTTPHeaderAccepts := []string{}

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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetSystemInOutBoardsRequest struct {
	ctx context.Context
	ApiService *InOutBoardsAPIService
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
func (r ApiGetSystemInOutBoardsRequest) Conditions(conditions string) ApiGetSystemInOutBoardsRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetSystemInOutBoardsRequest) ChildConditions(childConditions string) ApiGetSystemInOutBoardsRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetSystemInOutBoardsRequest) CustomFieldConditions(customFieldConditions string) ApiGetSystemInOutBoardsRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetSystemInOutBoardsRequest) OrderBy(orderBy string) ApiGetSystemInOutBoardsRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetSystemInOutBoardsRequest) Fields(fields string) ApiGetSystemInOutBoardsRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetSystemInOutBoardsRequest) Page(page int32) ApiGetSystemInOutBoardsRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetSystemInOutBoardsRequest) PageSize(pageSize int32) ApiGetSystemInOutBoardsRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetSystemInOutBoardsRequest) PageId(pageId int32) ApiGetSystemInOutBoardsRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetSystemInOutBoardsRequest) ClientId(clientId string) ApiGetSystemInOutBoardsRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetSystemInOutBoardsRequest) Execute() ([]InOutBoard, *http.Response, error) {
	return r.ApiService.GetSystemInOutBoardsExecute(r)
}

/*
GetSystemInOutBoards Get List of InOutBoard

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetSystemInOutBoardsRequest
*/
func (a *InOutBoardsAPIService) GetSystemInOutBoards(ctx context.Context) ApiGetSystemInOutBoardsRequest {
	return ApiGetSystemInOutBoardsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []InOutBoard
func (a *InOutBoardsAPIService) GetSystemInOutBoardsExecute(r ApiGetSystemInOutBoardsRequest) ([]InOutBoard, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []InOutBoard
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InOutBoardsAPIService.GetSystemInOutBoards")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/inOutBoards"

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

type ApiGetSystemInOutBoardsByIdRequest struct {
	ctx context.Context
	ApiService *InOutBoardsAPIService
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
func (r ApiGetSystemInOutBoardsByIdRequest) Conditions(conditions string) ApiGetSystemInOutBoardsByIdRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetSystemInOutBoardsByIdRequest) ChildConditions(childConditions string) ApiGetSystemInOutBoardsByIdRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetSystemInOutBoardsByIdRequest) CustomFieldConditions(customFieldConditions string) ApiGetSystemInOutBoardsByIdRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetSystemInOutBoardsByIdRequest) OrderBy(orderBy string) ApiGetSystemInOutBoardsByIdRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetSystemInOutBoardsByIdRequest) Fields(fields string) ApiGetSystemInOutBoardsByIdRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetSystemInOutBoardsByIdRequest) Page(page int32) ApiGetSystemInOutBoardsByIdRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetSystemInOutBoardsByIdRequest) PageSize(pageSize int32) ApiGetSystemInOutBoardsByIdRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetSystemInOutBoardsByIdRequest) PageId(pageId int32) ApiGetSystemInOutBoardsByIdRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetSystemInOutBoardsByIdRequest) ClientId(clientId string) ApiGetSystemInOutBoardsByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetSystemInOutBoardsByIdRequest) Execute() (*InOutBoard, *http.Response, error) {
	return r.ApiService.GetSystemInOutBoardsByIdExecute(r)
}

/*
GetSystemInOutBoardsById Get InOutBoard

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id inOutBoardId
 @return ApiGetSystemInOutBoardsByIdRequest
*/
func (a *InOutBoardsAPIService) GetSystemInOutBoardsById(ctx context.Context, id int32) ApiGetSystemInOutBoardsByIdRequest {
	return ApiGetSystemInOutBoardsByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return InOutBoard
func (a *InOutBoardsAPIService) GetSystemInOutBoardsByIdExecute(r ApiGetSystemInOutBoardsByIdRequest) (*InOutBoard, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InOutBoard
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InOutBoardsAPIService.GetSystemInOutBoardsById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/inOutBoards/{id}"
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

type ApiGetSystemInOutBoardsCountRequest struct {
	ctx context.Context
	ApiService *InOutBoardsAPIService
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
func (r ApiGetSystemInOutBoardsCountRequest) Conditions(conditions string) ApiGetSystemInOutBoardsCountRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetSystemInOutBoardsCountRequest) ChildConditions(childConditions string) ApiGetSystemInOutBoardsCountRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetSystemInOutBoardsCountRequest) CustomFieldConditions(customFieldConditions string) ApiGetSystemInOutBoardsCountRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetSystemInOutBoardsCountRequest) OrderBy(orderBy string) ApiGetSystemInOutBoardsCountRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetSystemInOutBoardsCountRequest) Fields(fields string) ApiGetSystemInOutBoardsCountRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetSystemInOutBoardsCountRequest) Page(page int32) ApiGetSystemInOutBoardsCountRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetSystemInOutBoardsCountRequest) PageSize(pageSize int32) ApiGetSystemInOutBoardsCountRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetSystemInOutBoardsCountRequest) PageId(pageId int32) ApiGetSystemInOutBoardsCountRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetSystemInOutBoardsCountRequest) ClientId(clientId string) ApiGetSystemInOutBoardsCountRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetSystemInOutBoardsCountRequest) Execute() (*Count, *http.Response, error) {
	return r.ApiService.GetSystemInOutBoardsCountExecute(r)
}

/*
GetSystemInOutBoardsCount Get Count of InOutBoard

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetSystemInOutBoardsCountRequest
*/
func (a *InOutBoardsAPIService) GetSystemInOutBoardsCount(ctx context.Context) ApiGetSystemInOutBoardsCountRequest {
	return ApiGetSystemInOutBoardsCountRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Count
func (a *InOutBoardsAPIService) GetSystemInOutBoardsCountExecute(r ApiGetSystemInOutBoardsCountRequest) (*Count, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Count
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InOutBoardsAPIService.GetSystemInOutBoardsCount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/inOutBoards/count"

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

type ApiPatchSystemInOutBoardsByIdRequest struct {
	ctx context.Context
	ApiService *InOutBoardsAPIService
	id int32
	patchOperation *[]PatchOperation
	clientId *string
}

// List of PatchOperation
func (r ApiPatchSystemInOutBoardsByIdRequest) PatchOperation(patchOperation []PatchOperation) ApiPatchSystemInOutBoardsByIdRequest {
	r.patchOperation = &patchOperation
	return r
}

// 
func (r ApiPatchSystemInOutBoardsByIdRequest) ClientId(clientId string) ApiPatchSystemInOutBoardsByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiPatchSystemInOutBoardsByIdRequest) Execute() (*InOutBoard, *http.Response, error) {
	return r.ApiService.PatchSystemInOutBoardsByIdExecute(r)
}

/*
PatchSystemInOutBoardsById Patch InOutBoard

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id inOutBoardId
 @return ApiPatchSystemInOutBoardsByIdRequest
*/
func (a *InOutBoardsAPIService) PatchSystemInOutBoardsById(ctx context.Context, id int32) ApiPatchSystemInOutBoardsByIdRequest {
	return ApiPatchSystemInOutBoardsByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return InOutBoard
func (a *InOutBoardsAPIService) PatchSystemInOutBoardsByIdExecute(r ApiPatchSystemInOutBoardsByIdRequest) (*InOutBoard, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InOutBoard
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InOutBoardsAPIService.PatchSystemInOutBoardsById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/inOutBoards/{id}"
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

type ApiPostSystemInOutBoardsRequest struct {
	ctx context.Context
	ApiService *InOutBoardsAPIService
	inOutBoard *InOutBoard
	clientId *string
}

// inOutBoard
func (r ApiPostSystemInOutBoardsRequest) InOutBoard(inOutBoard InOutBoard) ApiPostSystemInOutBoardsRequest {
	r.inOutBoard = &inOutBoard
	return r
}

// 
func (r ApiPostSystemInOutBoardsRequest) ClientId(clientId string) ApiPostSystemInOutBoardsRequest {
	r.clientId = &clientId
	return r
}

func (r ApiPostSystemInOutBoardsRequest) Execute() (*InOutBoard, *http.Response, error) {
	return r.ApiService.PostSystemInOutBoardsExecute(r)
}

/*
PostSystemInOutBoards Post InOutBoard

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPostSystemInOutBoardsRequest
*/
func (a *InOutBoardsAPIService) PostSystemInOutBoards(ctx context.Context) ApiPostSystemInOutBoardsRequest {
	return ApiPostSystemInOutBoardsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return InOutBoard
func (a *InOutBoardsAPIService) PostSystemInOutBoardsExecute(r ApiPostSystemInOutBoardsRequest) (*InOutBoard, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InOutBoard
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InOutBoardsAPIService.PostSystemInOutBoards")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/inOutBoards"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.inOutBoard == nil {
		return localVarReturnValue, nil, reportError("inOutBoard is required and must be specified")
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
	localVarPostBody = r.inOutBoard
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

type ApiPutSystemInOutBoardsByIdRequest struct {
	ctx context.Context
	ApiService *InOutBoardsAPIService
	id int32
	inOutBoard *InOutBoard
	clientId *string
}

// inOutBoard
func (r ApiPutSystemInOutBoardsByIdRequest) InOutBoard(inOutBoard InOutBoard) ApiPutSystemInOutBoardsByIdRequest {
	r.inOutBoard = &inOutBoard
	return r
}

// 
func (r ApiPutSystemInOutBoardsByIdRequest) ClientId(clientId string) ApiPutSystemInOutBoardsByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiPutSystemInOutBoardsByIdRequest) Execute() (*InOutBoard, *http.Response, error) {
	return r.ApiService.PutSystemInOutBoardsByIdExecute(r)
}

/*
PutSystemInOutBoardsById Put InOutBoard

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id inOutBoardId
 @return ApiPutSystemInOutBoardsByIdRequest
*/
func (a *InOutBoardsAPIService) PutSystemInOutBoardsById(ctx context.Context, id int32) ApiPutSystemInOutBoardsByIdRequest {
	return ApiPutSystemInOutBoardsByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return InOutBoard
func (a *InOutBoardsAPIService) PutSystemInOutBoardsByIdExecute(r ApiPutSystemInOutBoardsByIdRequest) (*InOutBoard, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InOutBoard
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InOutBoardsAPIService.PutSystemInOutBoardsById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/inOutBoards/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.inOutBoard == nil {
		return localVarReturnValue, nil, reportError("inOutBoard is required and must be specified")
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
	localVarPostBody = r.inOutBoard
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
