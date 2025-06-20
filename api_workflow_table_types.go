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


// WorkflowTableTypesAPIService WorkflowTableTypesAPI service
type WorkflowTableTypesAPIService service

type ApiGetSystemWorkflowsTableTypesRequest struct {
	ctx context.Context
	ApiService *WorkflowTableTypesAPIService
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
func (r ApiGetSystemWorkflowsTableTypesRequest) Conditions(conditions string) ApiGetSystemWorkflowsTableTypesRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetSystemWorkflowsTableTypesRequest) ChildConditions(childConditions string) ApiGetSystemWorkflowsTableTypesRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetSystemWorkflowsTableTypesRequest) CustomFieldConditions(customFieldConditions string) ApiGetSystemWorkflowsTableTypesRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetSystemWorkflowsTableTypesRequest) OrderBy(orderBy string) ApiGetSystemWorkflowsTableTypesRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetSystemWorkflowsTableTypesRequest) Fields(fields string) ApiGetSystemWorkflowsTableTypesRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetSystemWorkflowsTableTypesRequest) Page(page int32) ApiGetSystemWorkflowsTableTypesRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetSystemWorkflowsTableTypesRequest) PageSize(pageSize int32) ApiGetSystemWorkflowsTableTypesRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetSystemWorkflowsTableTypesRequest) PageId(pageId int32) ApiGetSystemWorkflowsTableTypesRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetSystemWorkflowsTableTypesRequest) ClientId(clientId string) ApiGetSystemWorkflowsTableTypesRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetSystemWorkflowsTableTypesRequest) Execute() ([]WorkflowTableType, *http.Response, error) {
	return r.ApiService.GetSystemWorkflowsTableTypesExecute(r)
}

/*
GetSystemWorkflowsTableTypes Get List of WorkflowTableType

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetSystemWorkflowsTableTypesRequest
*/
func (a *WorkflowTableTypesAPIService) GetSystemWorkflowsTableTypes(ctx context.Context) ApiGetSystemWorkflowsTableTypesRequest {
	return ApiGetSystemWorkflowsTableTypesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []WorkflowTableType
func (a *WorkflowTableTypesAPIService) GetSystemWorkflowsTableTypesExecute(r ApiGetSystemWorkflowsTableTypesRequest) ([]WorkflowTableType, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []WorkflowTableType
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorkflowTableTypesAPIService.GetSystemWorkflowsTableTypes")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/workflows/tableTypes"

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

type ApiGetSystemWorkflowsTableTypesByIdRequest struct {
	ctx context.Context
	ApiService *WorkflowTableTypesAPIService
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
func (r ApiGetSystemWorkflowsTableTypesByIdRequest) Conditions(conditions string) ApiGetSystemWorkflowsTableTypesByIdRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetSystemWorkflowsTableTypesByIdRequest) ChildConditions(childConditions string) ApiGetSystemWorkflowsTableTypesByIdRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetSystemWorkflowsTableTypesByIdRequest) CustomFieldConditions(customFieldConditions string) ApiGetSystemWorkflowsTableTypesByIdRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetSystemWorkflowsTableTypesByIdRequest) OrderBy(orderBy string) ApiGetSystemWorkflowsTableTypesByIdRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetSystemWorkflowsTableTypesByIdRequest) Fields(fields string) ApiGetSystemWorkflowsTableTypesByIdRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetSystemWorkflowsTableTypesByIdRequest) Page(page int32) ApiGetSystemWorkflowsTableTypesByIdRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetSystemWorkflowsTableTypesByIdRequest) PageSize(pageSize int32) ApiGetSystemWorkflowsTableTypesByIdRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetSystemWorkflowsTableTypesByIdRequest) PageId(pageId int32) ApiGetSystemWorkflowsTableTypesByIdRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetSystemWorkflowsTableTypesByIdRequest) ClientId(clientId string) ApiGetSystemWorkflowsTableTypesByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetSystemWorkflowsTableTypesByIdRequest) Execute() (*WorkflowTableType, *http.Response, error) {
	return r.ApiService.GetSystemWorkflowsTableTypesByIdExecute(r)
}

/*
GetSystemWorkflowsTableTypesById Get WorkflowTableType

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id tableTypeId
 @return ApiGetSystemWorkflowsTableTypesByIdRequest
*/
func (a *WorkflowTableTypesAPIService) GetSystemWorkflowsTableTypesById(ctx context.Context, id int32) ApiGetSystemWorkflowsTableTypesByIdRequest {
	return ApiGetSystemWorkflowsTableTypesByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return WorkflowTableType
func (a *WorkflowTableTypesAPIService) GetSystemWorkflowsTableTypesByIdExecute(r ApiGetSystemWorkflowsTableTypesByIdRequest) (*WorkflowTableType, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *WorkflowTableType
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorkflowTableTypesAPIService.GetSystemWorkflowsTableTypesById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/workflows/tableTypes/{id}"
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

type ApiGetSystemWorkflowsTableTypesByIdInfoRequest struct {
	ctx context.Context
	ApiService *WorkflowTableTypesAPIService
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
func (r ApiGetSystemWorkflowsTableTypesByIdInfoRequest) Conditions(conditions string) ApiGetSystemWorkflowsTableTypesByIdInfoRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetSystemWorkflowsTableTypesByIdInfoRequest) ChildConditions(childConditions string) ApiGetSystemWorkflowsTableTypesByIdInfoRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetSystemWorkflowsTableTypesByIdInfoRequest) CustomFieldConditions(customFieldConditions string) ApiGetSystemWorkflowsTableTypesByIdInfoRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetSystemWorkflowsTableTypesByIdInfoRequest) OrderBy(orderBy string) ApiGetSystemWorkflowsTableTypesByIdInfoRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetSystemWorkflowsTableTypesByIdInfoRequest) Fields(fields string) ApiGetSystemWorkflowsTableTypesByIdInfoRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetSystemWorkflowsTableTypesByIdInfoRequest) Page(page int32) ApiGetSystemWorkflowsTableTypesByIdInfoRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetSystemWorkflowsTableTypesByIdInfoRequest) PageSize(pageSize int32) ApiGetSystemWorkflowsTableTypesByIdInfoRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetSystemWorkflowsTableTypesByIdInfoRequest) PageId(pageId int32) ApiGetSystemWorkflowsTableTypesByIdInfoRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetSystemWorkflowsTableTypesByIdInfoRequest) ClientId(clientId string) ApiGetSystemWorkflowsTableTypesByIdInfoRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetSystemWorkflowsTableTypesByIdInfoRequest) Execute() (*WorkflowTableTypeInfo, *http.Response, error) {
	return r.ApiService.GetSystemWorkflowsTableTypesByIdInfoExecute(r)
}

/*
GetSystemWorkflowsTableTypesByIdInfo Get WorkflowTableTypeInfo

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id tableTypeId
 @return ApiGetSystemWorkflowsTableTypesByIdInfoRequest
*/
func (a *WorkflowTableTypesAPIService) GetSystemWorkflowsTableTypesByIdInfo(ctx context.Context, id int32) ApiGetSystemWorkflowsTableTypesByIdInfoRequest {
	return ApiGetSystemWorkflowsTableTypesByIdInfoRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return WorkflowTableTypeInfo
func (a *WorkflowTableTypesAPIService) GetSystemWorkflowsTableTypesByIdInfoExecute(r ApiGetSystemWorkflowsTableTypesByIdInfoRequest) (*WorkflowTableTypeInfo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *WorkflowTableTypeInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorkflowTableTypesAPIService.GetSystemWorkflowsTableTypesByIdInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/workflows/tableTypes/{id}/info"
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

type ApiGetSystemWorkflowsTableTypesCountRequest struct {
	ctx context.Context
	ApiService *WorkflowTableTypesAPIService
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
func (r ApiGetSystemWorkflowsTableTypesCountRequest) Conditions(conditions string) ApiGetSystemWorkflowsTableTypesCountRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetSystemWorkflowsTableTypesCountRequest) ChildConditions(childConditions string) ApiGetSystemWorkflowsTableTypesCountRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetSystemWorkflowsTableTypesCountRequest) CustomFieldConditions(customFieldConditions string) ApiGetSystemWorkflowsTableTypesCountRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetSystemWorkflowsTableTypesCountRequest) OrderBy(orderBy string) ApiGetSystemWorkflowsTableTypesCountRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetSystemWorkflowsTableTypesCountRequest) Fields(fields string) ApiGetSystemWorkflowsTableTypesCountRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetSystemWorkflowsTableTypesCountRequest) Page(page int32) ApiGetSystemWorkflowsTableTypesCountRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetSystemWorkflowsTableTypesCountRequest) PageSize(pageSize int32) ApiGetSystemWorkflowsTableTypesCountRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetSystemWorkflowsTableTypesCountRequest) PageId(pageId int32) ApiGetSystemWorkflowsTableTypesCountRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetSystemWorkflowsTableTypesCountRequest) ClientId(clientId string) ApiGetSystemWorkflowsTableTypesCountRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetSystemWorkflowsTableTypesCountRequest) Execute() (*Count, *http.Response, error) {
	return r.ApiService.GetSystemWorkflowsTableTypesCountExecute(r)
}

/*
GetSystemWorkflowsTableTypesCount Get Count of WorkflowTableType

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetSystemWorkflowsTableTypesCountRequest
*/
func (a *WorkflowTableTypesAPIService) GetSystemWorkflowsTableTypesCount(ctx context.Context) ApiGetSystemWorkflowsTableTypesCountRequest {
	return ApiGetSystemWorkflowsTableTypesCountRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Count
func (a *WorkflowTableTypesAPIService) GetSystemWorkflowsTableTypesCountExecute(r ApiGetSystemWorkflowsTableTypesCountRequest) (*Count, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Count
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorkflowTableTypesAPIService.GetSystemWorkflowsTableTypesCount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/workflows/tableTypes/count"

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

type ApiGetSystemWorkflowsTableTypesInfoRequest struct {
	ctx context.Context
	ApiService *WorkflowTableTypesAPIService
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
func (r ApiGetSystemWorkflowsTableTypesInfoRequest) Conditions(conditions string) ApiGetSystemWorkflowsTableTypesInfoRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetSystemWorkflowsTableTypesInfoRequest) ChildConditions(childConditions string) ApiGetSystemWorkflowsTableTypesInfoRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetSystemWorkflowsTableTypesInfoRequest) CustomFieldConditions(customFieldConditions string) ApiGetSystemWorkflowsTableTypesInfoRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetSystemWorkflowsTableTypesInfoRequest) OrderBy(orderBy string) ApiGetSystemWorkflowsTableTypesInfoRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetSystemWorkflowsTableTypesInfoRequest) Fields(fields string) ApiGetSystemWorkflowsTableTypesInfoRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetSystemWorkflowsTableTypesInfoRequest) Page(page int32) ApiGetSystemWorkflowsTableTypesInfoRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetSystemWorkflowsTableTypesInfoRequest) PageSize(pageSize int32) ApiGetSystemWorkflowsTableTypesInfoRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetSystemWorkflowsTableTypesInfoRequest) PageId(pageId int32) ApiGetSystemWorkflowsTableTypesInfoRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetSystemWorkflowsTableTypesInfoRequest) ClientId(clientId string) ApiGetSystemWorkflowsTableTypesInfoRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetSystemWorkflowsTableTypesInfoRequest) Execute() ([]WorkflowTableTypeInfo, *http.Response, error) {
	return r.ApiService.GetSystemWorkflowsTableTypesInfoExecute(r)
}

/*
GetSystemWorkflowsTableTypesInfo Get List of WorkflowTableTypeInfo

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetSystemWorkflowsTableTypesInfoRequest
*/
func (a *WorkflowTableTypesAPIService) GetSystemWorkflowsTableTypesInfo(ctx context.Context) ApiGetSystemWorkflowsTableTypesInfoRequest {
	return ApiGetSystemWorkflowsTableTypesInfoRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []WorkflowTableTypeInfo
func (a *WorkflowTableTypesAPIService) GetSystemWorkflowsTableTypesInfoExecute(r ApiGetSystemWorkflowsTableTypesInfoRequest) ([]WorkflowTableTypeInfo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []WorkflowTableTypeInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorkflowTableTypesAPIService.GetSystemWorkflowsTableTypesInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/workflows/tableTypes/info"

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

type ApiGetSystemWorkflowsTableTypesInfoCountRequest struct {
	ctx context.Context
	ApiService *WorkflowTableTypesAPIService
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
func (r ApiGetSystemWorkflowsTableTypesInfoCountRequest) Conditions(conditions string) ApiGetSystemWorkflowsTableTypesInfoCountRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetSystemWorkflowsTableTypesInfoCountRequest) ChildConditions(childConditions string) ApiGetSystemWorkflowsTableTypesInfoCountRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetSystemWorkflowsTableTypesInfoCountRequest) CustomFieldConditions(customFieldConditions string) ApiGetSystemWorkflowsTableTypesInfoCountRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetSystemWorkflowsTableTypesInfoCountRequest) OrderBy(orderBy string) ApiGetSystemWorkflowsTableTypesInfoCountRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetSystemWorkflowsTableTypesInfoCountRequest) Fields(fields string) ApiGetSystemWorkflowsTableTypesInfoCountRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetSystemWorkflowsTableTypesInfoCountRequest) Page(page int32) ApiGetSystemWorkflowsTableTypesInfoCountRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetSystemWorkflowsTableTypesInfoCountRequest) PageSize(pageSize int32) ApiGetSystemWorkflowsTableTypesInfoCountRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetSystemWorkflowsTableTypesInfoCountRequest) PageId(pageId int32) ApiGetSystemWorkflowsTableTypesInfoCountRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetSystemWorkflowsTableTypesInfoCountRequest) ClientId(clientId string) ApiGetSystemWorkflowsTableTypesInfoCountRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetSystemWorkflowsTableTypesInfoCountRequest) Execute() (*Count, *http.Response, error) {
	return r.ApiService.GetSystemWorkflowsTableTypesInfoCountExecute(r)
}

/*
GetSystemWorkflowsTableTypesInfoCount Get Count of WorkflowTableTypeInfo

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetSystemWorkflowsTableTypesInfoCountRequest
*/
func (a *WorkflowTableTypesAPIService) GetSystemWorkflowsTableTypesInfoCount(ctx context.Context) ApiGetSystemWorkflowsTableTypesInfoCountRequest {
	return ApiGetSystemWorkflowsTableTypesInfoCountRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Count
func (a *WorkflowTableTypesAPIService) GetSystemWorkflowsTableTypesInfoCountExecute(r ApiGetSystemWorkflowsTableTypesInfoCountRequest) (*Count, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Count
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorkflowTableTypesAPIService.GetSystemWorkflowsTableTypesInfoCount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/workflows/tableTypes/info/count"

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
