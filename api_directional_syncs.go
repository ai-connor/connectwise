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


// DirectionalSyncsAPIService DirectionalSyncsAPI service
type DirectionalSyncsAPIService service

type ApiDeleteProcurementDirectionalSyncsByIdRequest struct {
	ctx context.Context
	ApiService *DirectionalSyncsAPIService
	id int32
	clientId *string
}

// 
func (r ApiDeleteProcurementDirectionalSyncsByIdRequest) ClientId(clientId string) ApiDeleteProcurementDirectionalSyncsByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiDeleteProcurementDirectionalSyncsByIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteProcurementDirectionalSyncsByIdExecute(r)
}

/*
DeleteProcurementDirectionalSyncsById Delete DirectionalSync

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id warehousId
 @return ApiDeleteProcurementDirectionalSyncsByIdRequest
*/
func (a *DirectionalSyncsAPIService) DeleteProcurementDirectionalSyncsById(ctx context.Context, id int32) ApiDeleteProcurementDirectionalSyncsByIdRequest {
	return ApiDeleteProcurementDirectionalSyncsByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *DirectionalSyncsAPIService) DeleteProcurementDirectionalSyncsByIdExecute(r ApiDeleteProcurementDirectionalSyncsByIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DirectionalSyncsAPIService.DeleteProcurementDirectionalSyncsById")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/procurement/directionalSyncs/{id}"
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

type ApiGetProcurementDirectionalSyncsRequest struct {
	ctx context.Context
	ApiService *DirectionalSyncsAPIService
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
func (r ApiGetProcurementDirectionalSyncsRequest) Conditions(conditions string) ApiGetProcurementDirectionalSyncsRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetProcurementDirectionalSyncsRequest) ChildConditions(childConditions string) ApiGetProcurementDirectionalSyncsRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetProcurementDirectionalSyncsRequest) CustomFieldConditions(customFieldConditions string) ApiGetProcurementDirectionalSyncsRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetProcurementDirectionalSyncsRequest) OrderBy(orderBy string) ApiGetProcurementDirectionalSyncsRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetProcurementDirectionalSyncsRequest) Fields(fields string) ApiGetProcurementDirectionalSyncsRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetProcurementDirectionalSyncsRequest) Page(page int32) ApiGetProcurementDirectionalSyncsRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetProcurementDirectionalSyncsRequest) PageSize(pageSize int32) ApiGetProcurementDirectionalSyncsRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetProcurementDirectionalSyncsRequest) PageId(pageId int32) ApiGetProcurementDirectionalSyncsRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetProcurementDirectionalSyncsRequest) ClientId(clientId string) ApiGetProcurementDirectionalSyncsRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetProcurementDirectionalSyncsRequest) Execute() ([]DirectionalSync, *http.Response, error) {
	return r.ApiService.GetProcurementDirectionalSyncsExecute(r)
}

/*
GetProcurementDirectionalSyncs Get List of DirectionalSync

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetProcurementDirectionalSyncsRequest
*/
func (a *DirectionalSyncsAPIService) GetProcurementDirectionalSyncs(ctx context.Context) ApiGetProcurementDirectionalSyncsRequest {
	return ApiGetProcurementDirectionalSyncsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []DirectionalSync
func (a *DirectionalSyncsAPIService) GetProcurementDirectionalSyncsExecute(r ApiGetProcurementDirectionalSyncsRequest) ([]DirectionalSync, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []DirectionalSync
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DirectionalSyncsAPIService.GetProcurementDirectionalSyncs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/procurement/directionalSyncs"

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

type ApiGetProcurementDirectionalSyncsByIdRequest struct {
	ctx context.Context
	ApiService *DirectionalSyncsAPIService
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
func (r ApiGetProcurementDirectionalSyncsByIdRequest) Conditions(conditions string) ApiGetProcurementDirectionalSyncsByIdRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetProcurementDirectionalSyncsByIdRequest) ChildConditions(childConditions string) ApiGetProcurementDirectionalSyncsByIdRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetProcurementDirectionalSyncsByIdRequest) CustomFieldConditions(customFieldConditions string) ApiGetProcurementDirectionalSyncsByIdRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetProcurementDirectionalSyncsByIdRequest) OrderBy(orderBy string) ApiGetProcurementDirectionalSyncsByIdRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetProcurementDirectionalSyncsByIdRequest) Fields(fields string) ApiGetProcurementDirectionalSyncsByIdRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetProcurementDirectionalSyncsByIdRequest) Page(page int32) ApiGetProcurementDirectionalSyncsByIdRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetProcurementDirectionalSyncsByIdRequest) PageSize(pageSize int32) ApiGetProcurementDirectionalSyncsByIdRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetProcurementDirectionalSyncsByIdRequest) PageId(pageId int32) ApiGetProcurementDirectionalSyncsByIdRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetProcurementDirectionalSyncsByIdRequest) ClientId(clientId string) ApiGetProcurementDirectionalSyncsByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetProcurementDirectionalSyncsByIdRequest) Execute() (*DirectionalSync, *http.Response, error) {
	return r.ApiService.GetProcurementDirectionalSyncsByIdExecute(r)
}

/*
GetProcurementDirectionalSyncsById Get DirectionalSync

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id warehousId
 @return ApiGetProcurementDirectionalSyncsByIdRequest
*/
func (a *DirectionalSyncsAPIService) GetProcurementDirectionalSyncsById(ctx context.Context, id int32) ApiGetProcurementDirectionalSyncsByIdRequest {
	return ApiGetProcurementDirectionalSyncsByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return DirectionalSync
func (a *DirectionalSyncsAPIService) GetProcurementDirectionalSyncsByIdExecute(r ApiGetProcurementDirectionalSyncsByIdRequest) (*DirectionalSync, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DirectionalSync
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DirectionalSyncsAPIService.GetProcurementDirectionalSyncsById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/procurement/directionalSyncs/{id}"
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

type ApiGetProcurementDirectionalSyncsCountRequest struct {
	ctx context.Context
	ApiService *DirectionalSyncsAPIService
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
func (r ApiGetProcurementDirectionalSyncsCountRequest) Conditions(conditions string) ApiGetProcurementDirectionalSyncsCountRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetProcurementDirectionalSyncsCountRequest) ChildConditions(childConditions string) ApiGetProcurementDirectionalSyncsCountRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetProcurementDirectionalSyncsCountRequest) CustomFieldConditions(customFieldConditions string) ApiGetProcurementDirectionalSyncsCountRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetProcurementDirectionalSyncsCountRequest) OrderBy(orderBy string) ApiGetProcurementDirectionalSyncsCountRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetProcurementDirectionalSyncsCountRequest) Fields(fields string) ApiGetProcurementDirectionalSyncsCountRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetProcurementDirectionalSyncsCountRequest) Page(page int32) ApiGetProcurementDirectionalSyncsCountRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetProcurementDirectionalSyncsCountRequest) PageSize(pageSize int32) ApiGetProcurementDirectionalSyncsCountRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetProcurementDirectionalSyncsCountRequest) PageId(pageId int32) ApiGetProcurementDirectionalSyncsCountRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetProcurementDirectionalSyncsCountRequest) ClientId(clientId string) ApiGetProcurementDirectionalSyncsCountRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetProcurementDirectionalSyncsCountRequest) Execute() (*Count, *http.Response, error) {
	return r.ApiService.GetProcurementDirectionalSyncsCountExecute(r)
}

/*
GetProcurementDirectionalSyncsCount Get Count of DirectionalSync

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetProcurementDirectionalSyncsCountRequest
*/
func (a *DirectionalSyncsAPIService) GetProcurementDirectionalSyncsCount(ctx context.Context) ApiGetProcurementDirectionalSyncsCountRequest {
	return ApiGetProcurementDirectionalSyncsCountRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Count
func (a *DirectionalSyncsAPIService) GetProcurementDirectionalSyncsCountExecute(r ApiGetProcurementDirectionalSyncsCountRequest) (*Count, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Count
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DirectionalSyncsAPIService.GetProcurementDirectionalSyncsCount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/procurement/directionalSyncs/count"

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

type ApiPatchProcurementDirectionalSyncsByIdRequest struct {
	ctx context.Context
	ApiService *DirectionalSyncsAPIService
	id int32
	patchOperation *[]PatchOperation
	clientId *string
}

// List of PatchOperation
func (r ApiPatchProcurementDirectionalSyncsByIdRequest) PatchOperation(patchOperation []PatchOperation) ApiPatchProcurementDirectionalSyncsByIdRequest {
	r.patchOperation = &patchOperation
	return r
}

// 
func (r ApiPatchProcurementDirectionalSyncsByIdRequest) ClientId(clientId string) ApiPatchProcurementDirectionalSyncsByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiPatchProcurementDirectionalSyncsByIdRequest) Execute() (*DirectionalSync, *http.Response, error) {
	return r.ApiService.PatchProcurementDirectionalSyncsByIdExecute(r)
}

/*
PatchProcurementDirectionalSyncsById Patch DirectionalSync

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id warehousId
 @return ApiPatchProcurementDirectionalSyncsByIdRequest
*/
func (a *DirectionalSyncsAPIService) PatchProcurementDirectionalSyncsById(ctx context.Context, id int32) ApiPatchProcurementDirectionalSyncsByIdRequest {
	return ApiPatchProcurementDirectionalSyncsByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return DirectionalSync
func (a *DirectionalSyncsAPIService) PatchProcurementDirectionalSyncsByIdExecute(r ApiPatchProcurementDirectionalSyncsByIdRequest) (*DirectionalSync, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DirectionalSync
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DirectionalSyncsAPIService.PatchProcurementDirectionalSyncsById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/procurement/directionalSyncs/{id}"
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

type ApiPostProcurementDirectionalSyncsRequest struct {
	ctx context.Context
	ApiService *DirectionalSyncsAPIService
	directionalSync *DirectionalSync
	clientId *string
}

// directionalSync
func (r ApiPostProcurementDirectionalSyncsRequest) DirectionalSync(directionalSync DirectionalSync) ApiPostProcurementDirectionalSyncsRequest {
	r.directionalSync = &directionalSync
	return r
}

// 
func (r ApiPostProcurementDirectionalSyncsRequest) ClientId(clientId string) ApiPostProcurementDirectionalSyncsRequest {
	r.clientId = &clientId
	return r
}

func (r ApiPostProcurementDirectionalSyncsRequest) Execute() (*DirectionalSync, *http.Response, error) {
	return r.ApiService.PostProcurementDirectionalSyncsExecute(r)
}

/*
PostProcurementDirectionalSyncs Post DirectionalSync

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPostProcurementDirectionalSyncsRequest
*/
func (a *DirectionalSyncsAPIService) PostProcurementDirectionalSyncs(ctx context.Context) ApiPostProcurementDirectionalSyncsRequest {
	return ApiPostProcurementDirectionalSyncsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return DirectionalSync
func (a *DirectionalSyncsAPIService) PostProcurementDirectionalSyncsExecute(r ApiPostProcurementDirectionalSyncsRequest) (*DirectionalSync, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DirectionalSync
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DirectionalSyncsAPIService.PostProcurementDirectionalSyncs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/procurement/directionalSyncs"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.directionalSync == nil {
		return localVarReturnValue, nil, reportError("directionalSync is required and must be specified")
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
	localVarPostBody = r.directionalSync
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

type ApiPutProcurementDirectionalSyncsByIdRequest struct {
	ctx context.Context
	ApiService *DirectionalSyncsAPIService
	id int32
	directionalSync *DirectionalSync
	clientId *string
}

// directionalSync
func (r ApiPutProcurementDirectionalSyncsByIdRequest) DirectionalSync(directionalSync DirectionalSync) ApiPutProcurementDirectionalSyncsByIdRequest {
	r.directionalSync = &directionalSync
	return r
}

// 
func (r ApiPutProcurementDirectionalSyncsByIdRequest) ClientId(clientId string) ApiPutProcurementDirectionalSyncsByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiPutProcurementDirectionalSyncsByIdRequest) Execute() (*DirectionalSync, *http.Response, error) {
	return r.ApiService.PutProcurementDirectionalSyncsByIdExecute(r)
}

/*
PutProcurementDirectionalSyncsById Put DirectionalSync

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id warehousId
 @return ApiPutProcurementDirectionalSyncsByIdRequest
*/
func (a *DirectionalSyncsAPIService) PutProcurementDirectionalSyncsById(ctx context.Context, id int32) ApiPutProcurementDirectionalSyncsByIdRequest {
	return ApiPutProcurementDirectionalSyncsByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return DirectionalSync
func (a *DirectionalSyncsAPIService) PutProcurementDirectionalSyncsByIdExecute(r ApiPutProcurementDirectionalSyncsByIdRequest) (*DirectionalSync, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DirectionalSync
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DirectionalSyncsAPIService.PutProcurementDirectionalSyncsById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/procurement/directionalSyncs/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.directionalSync == nil {
		return localVarReturnValue, nil, reportError("directionalSync is required and must be specified")
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
	localVarPostBody = r.directionalSync
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
