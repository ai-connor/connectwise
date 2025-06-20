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


// WorkflowTriggersAPIService WorkflowTriggersAPI service
type WorkflowTriggersAPIService service

type ApiGetSystemWorkflowsByParentIdTriggersRequest struct {
	ctx context.Context
	ApiService *WorkflowTriggersAPIService
	parentId int32
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
func (r ApiGetSystemWorkflowsByParentIdTriggersRequest) Conditions(conditions string) ApiGetSystemWorkflowsByParentIdTriggersRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetSystemWorkflowsByParentIdTriggersRequest) ChildConditions(childConditions string) ApiGetSystemWorkflowsByParentIdTriggersRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetSystemWorkflowsByParentIdTriggersRequest) CustomFieldConditions(customFieldConditions string) ApiGetSystemWorkflowsByParentIdTriggersRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetSystemWorkflowsByParentIdTriggersRequest) OrderBy(orderBy string) ApiGetSystemWorkflowsByParentIdTriggersRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetSystemWorkflowsByParentIdTriggersRequest) Fields(fields string) ApiGetSystemWorkflowsByParentIdTriggersRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetSystemWorkflowsByParentIdTriggersRequest) Page(page int32) ApiGetSystemWorkflowsByParentIdTriggersRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetSystemWorkflowsByParentIdTriggersRequest) PageSize(pageSize int32) ApiGetSystemWorkflowsByParentIdTriggersRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetSystemWorkflowsByParentIdTriggersRequest) PageId(pageId int32) ApiGetSystemWorkflowsByParentIdTriggersRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetSystemWorkflowsByParentIdTriggersRequest) ClientId(clientId string) ApiGetSystemWorkflowsByParentIdTriggersRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetSystemWorkflowsByParentIdTriggersRequest) Execute() ([]WorkflowTrigger, *http.Response, error) {
	return r.ApiService.GetSystemWorkflowsByParentIdTriggersExecute(r)
}

/*
GetSystemWorkflowsByParentIdTriggers Get List of WorkflowTrigger

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param parentId workflowId
 @return ApiGetSystemWorkflowsByParentIdTriggersRequest
*/
func (a *WorkflowTriggersAPIService) GetSystemWorkflowsByParentIdTriggers(ctx context.Context, parentId int32) ApiGetSystemWorkflowsByParentIdTriggersRequest {
	return ApiGetSystemWorkflowsByParentIdTriggersRequest{
		ApiService: a,
		ctx: ctx,
		parentId: parentId,
	}
}

// Execute executes the request
//  @return []WorkflowTrigger
func (a *WorkflowTriggersAPIService) GetSystemWorkflowsByParentIdTriggersExecute(r ApiGetSystemWorkflowsByParentIdTriggersRequest) ([]WorkflowTrigger, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []WorkflowTrigger
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorkflowTriggersAPIService.GetSystemWorkflowsByParentIdTriggers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/workflows/{parentId}/triggers"
	localVarPath = strings.Replace(localVarPath, "{"+"parentId"+"}", url.PathEscape(parameterValueToString(r.parentId, "parentId")), -1)

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

type ApiGetSystemWorkflowsByParentIdTriggersCountRequest struct {
	ctx context.Context
	ApiService *WorkflowTriggersAPIService
	parentId int32
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
func (r ApiGetSystemWorkflowsByParentIdTriggersCountRequest) Conditions(conditions string) ApiGetSystemWorkflowsByParentIdTriggersCountRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetSystemWorkflowsByParentIdTriggersCountRequest) ChildConditions(childConditions string) ApiGetSystemWorkflowsByParentIdTriggersCountRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetSystemWorkflowsByParentIdTriggersCountRequest) CustomFieldConditions(customFieldConditions string) ApiGetSystemWorkflowsByParentIdTriggersCountRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetSystemWorkflowsByParentIdTriggersCountRequest) OrderBy(orderBy string) ApiGetSystemWorkflowsByParentIdTriggersCountRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetSystemWorkflowsByParentIdTriggersCountRequest) Fields(fields string) ApiGetSystemWorkflowsByParentIdTriggersCountRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetSystemWorkflowsByParentIdTriggersCountRequest) Page(page int32) ApiGetSystemWorkflowsByParentIdTriggersCountRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetSystemWorkflowsByParentIdTriggersCountRequest) PageSize(pageSize int32) ApiGetSystemWorkflowsByParentIdTriggersCountRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetSystemWorkflowsByParentIdTriggersCountRequest) PageId(pageId int32) ApiGetSystemWorkflowsByParentIdTriggersCountRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetSystemWorkflowsByParentIdTriggersCountRequest) ClientId(clientId string) ApiGetSystemWorkflowsByParentIdTriggersCountRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetSystemWorkflowsByParentIdTriggersCountRequest) Execute() (*Count, *http.Response, error) {
	return r.ApiService.GetSystemWorkflowsByParentIdTriggersCountExecute(r)
}

/*
GetSystemWorkflowsByParentIdTriggersCount Get Count of WorkflowTrigger

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param parentId workflowId
 @return ApiGetSystemWorkflowsByParentIdTriggersCountRequest
*/
func (a *WorkflowTriggersAPIService) GetSystemWorkflowsByParentIdTriggersCount(ctx context.Context, parentId int32) ApiGetSystemWorkflowsByParentIdTriggersCountRequest {
	return ApiGetSystemWorkflowsByParentIdTriggersCountRequest{
		ApiService: a,
		ctx: ctx,
		parentId: parentId,
	}
}

// Execute executes the request
//  @return Count
func (a *WorkflowTriggersAPIService) GetSystemWorkflowsByParentIdTriggersCountExecute(r ApiGetSystemWorkflowsByParentIdTriggersCountRequest) (*Count, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Count
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorkflowTriggersAPIService.GetSystemWorkflowsByParentIdTriggersCount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/workflows/{parentId}/triggers/count"
	localVarPath = strings.Replace(localVarPath, "{"+"parentId"+"}", url.PathEscape(parameterValueToString(r.parentId, "parentId")), -1)

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

type ApiGetSystemWorkflowsTriggersRequest struct {
	ctx context.Context
	ApiService *WorkflowTriggersAPIService
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
func (r ApiGetSystemWorkflowsTriggersRequest) Conditions(conditions string) ApiGetSystemWorkflowsTriggersRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetSystemWorkflowsTriggersRequest) ChildConditions(childConditions string) ApiGetSystemWorkflowsTriggersRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetSystemWorkflowsTriggersRequest) CustomFieldConditions(customFieldConditions string) ApiGetSystemWorkflowsTriggersRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetSystemWorkflowsTriggersRequest) OrderBy(orderBy string) ApiGetSystemWorkflowsTriggersRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetSystemWorkflowsTriggersRequest) Fields(fields string) ApiGetSystemWorkflowsTriggersRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetSystemWorkflowsTriggersRequest) Page(page int32) ApiGetSystemWorkflowsTriggersRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetSystemWorkflowsTriggersRequest) PageSize(pageSize int32) ApiGetSystemWorkflowsTriggersRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetSystemWorkflowsTriggersRequest) PageId(pageId int32) ApiGetSystemWorkflowsTriggersRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetSystemWorkflowsTriggersRequest) ClientId(clientId string) ApiGetSystemWorkflowsTriggersRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetSystemWorkflowsTriggersRequest) Execute() ([]WorkflowTrigger, *http.Response, error) {
	return r.ApiService.GetSystemWorkflowsTriggersExecute(r)
}

/*
GetSystemWorkflowsTriggers Get List of WorkflowTrigger

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetSystemWorkflowsTriggersRequest
*/
func (a *WorkflowTriggersAPIService) GetSystemWorkflowsTriggers(ctx context.Context) ApiGetSystemWorkflowsTriggersRequest {
	return ApiGetSystemWorkflowsTriggersRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []WorkflowTrigger
func (a *WorkflowTriggersAPIService) GetSystemWorkflowsTriggersExecute(r ApiGetSystemWorkflowsTriggersRequest) ([]WorkflowTrigger, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []WorkflowTrigger
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorkflowTriggersAPIService.GetSystemWorkflowsTriggers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/workflows/triggers"

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
