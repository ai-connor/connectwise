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


// IntegratorTagsAPIService IntegratorTagsAPI service
type IntegratorTagsAPIService service

type ApiDeleteSystemIntegratorTagsByIdRequest struct {
	ctx context.Context
	ApiService *IntegratorTagsAPIService
	id int32
	clientId *string
}

// 
func (r ApiDeleteSystemIntegratorTagsByIdRequest) ClientId(clientId string) ApiDeleteSystemIntegratorTagsByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiDeleteSystemIntegratorTagsByIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteSystemIntegratorTagsByIdExecute(r)
}

/*
DeleteSystemIntegratorTagsById Delete IntegratorTag

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id integratorTagId
 @return ApiDeleteSystemIntegratorTagsByIdRequest
*/
func (a *IntegratorTagsAPIService) DeleteSystemIntegratorTagsById(ctx context.Context, id int32) ApiDeleteSystemIntegratorTagsByIdRequest {
	return ApiDeleteSystemIntegratorTagsByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *IntegratorTagsAPIService) DeleteSystemIntegratorTagsByIdExecute(r ApiDeleteSystemIntegratorTagsByIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IntegratorTagsAPIService.DeleteSystemIntegratorTagsById")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/integratorTags/{id}"
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

type ApiGetSystemIntegratorTagsRequest struct {
	ctx context.Context
	ApiService *IntegratorTagsAPIService
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
func (r ApiGetSystemIntegratorTagsRequest) Conditions(conditions string) ApiGetSystemIntegratorTagsRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetSystemIntegratorTagsRequest) ChildConditions(childConditions string) ApiGetSystemIntegratorTagsRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetSystemIntegratorTagsRequest) CustomFieldConditions(customFieldConditions string) ApiGetSystemIntegratorTagsRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetSystemIntegratorTagsRequest) OrderBy(orderBy string) ApiGetSystemIntegratorTagsRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetSystemIntegratorTagsRequest) Fields(fields string) ApiGetSystemIntegratorTagsRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetSystemIntegratorTagsRequest) Page(page int32) ApiGetSystemIntegratorTagsRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetSystemIntegratorTagsRequest) PageSize(pageSize int32) ApiGetSystemIntegratorTagsRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetSystemIntegratorTagsRequest) PageId(pageId int32) ApiGetSystemIntegratorTagsRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetSystemIntegratorTagsRequest) ClientId(clientId string) ApiGetSystemIntegratorTagsRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetSystemIntegratorTagsRequest) Execute() ([]IntegratorTag, *http.Response, error) {
	return r.ApiService.GetSystemIntegratorTagsExecute(r)
}

/*
GetSystemIntegratorTags Get List of IntegratorTag

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetSystemIntegratorTagsRequest
*/
func (a *IntegratorTagsAPIService) GetSystemIntegratorTags(ctx context.Context) ApiGetSystemIntegratorTagsRequest {
	return ApiGetSystemIntegratorTagsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []IntegratorTag
func (a *IntegratorTagsAPIService) GetSystemIntegratorTagsExecute(r ApiGetSystemIntegratorTagsRequest) ([]IntegratorTag, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []IntegratorTag
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IntegratorTagsAPIService.GetSystemIntegratorTags")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/integratorTags"

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

type ApiGetSystemIntegratorTagsByIdRequest struct {
	ctx context.Context
	ApiService *IntegratorTagsAPIService
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
func (r ApiGetSystemIntegratorTagsByIdRequest) Conditions(conditions string) ApiGetSystemIntegratorTagsByIdRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetSystemIntegratorTagsByIdRequest) ChildConditions(childConditions string) ApiGetSystemIntegratorTagsByIdRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetSystemIntegratorTagsByIdRequest) CustomFieldConditions(customFieldConditions string) ApiGetSystemIntegratorTagsByIdRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetSystemIntegratorTagsByIdRequest) OrderBy(orderBy string) ApiGetSystemIntegratorTagsByIdRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetSystemIntegratorTagsByIdRequest) Fields(fields string) ApiGetSystemIntegratorTagsByIdRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetSystemIntegratorTagsByIdRequest) Page(page int32) ApiGetSystemIntegratorTagsByIdRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetSystemIntegratorTagsByIdRequest) PageSize(pageSize int32) ApiGetSystemIntegratorTagsByIdRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetSystemIntegratorTagsByIdRequest) PageId(pageId int32) ApiGetSystemIntegratorTagsByIdRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetSystemIntegratorTagsByIdRequest) ClientId(clientId string) ApiGetSystemIntegratorTagsByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetSystemIntegratorTagsByIdRequest) Execute() (*IntegratorTag, *http.Response, error) {
	return r.ApiService.GetSystemIntegratorTagsByIdExecute(r)
}

/*
GetSystemIntegratorTagsById Get IntegratorTag

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id integratorTagId
 @return ApiGetSystemIntegratorTagsByIdRequest
*/
func (a *IntegratorTagsAPIService) GetSystemIntegratorTagsById(ctx context.Context, id int32) ApiGetSystemIntegratorTagsByIdRequest {
	return ApiGetSystemIntegratorTagsByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return IntegratorTag
func (a *IntegratorTagsAPIService) GetSystemIntegratorTagsByIdExecute(r ApiGetSystemIntegratorTagsByIdRequest) (*IntegratorTag, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *IntegratorTag
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IntegratorTagsAPIService.GetSystemIntegratorTagsById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/integratorTags/{id}"
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

type ApiGetSystemIntegratorTagsCountRequest struct {
	ctx context.Context
	ApiService *IntegratorTagsAPIService
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
func (r ApiGetSystemIntegratorTagsCountRequest) Conditions(conditions string) ApiGetSystemIntegratorTagsCountRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetSystemIntegratorTagsCountRequest) ChildConditions(childConditions string) ApiGetSystemIntegratorTagsCountRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetSystemIntegratorTagsCountRequest) CustomFieldConditions(customFieldConditions string) ApiGetSystemIntegratorTagsCountRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetSystemIntegratorTagsCountRequest) OrderBy(orderBy string) ApiGetSystemIntegratorTagsCountRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetSystemIntegratorTagsCountRequest) Fields(fields string) ApiGetSystemIntegratorTagsCountRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetSystemIntegratorTagsCountRequest) Page(page int32) ApiGetSystemIntegratorTagsCountRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetSystemIntegratorTagsCountRequest) PageSize(pageSize int32) ApiGetSystemIntegratorTagsCountRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetSystemIntegratorTagsCountRequest) PageId(pageId int32) ApiGetSystemIntegratorTagsCountRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetSystemIntegratorTagsCountRequest) ClientId(clientId string) ApiGetSystemIntegratorTagsCountRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetSystemIntegratorTagsCountRequest) Execute() (*Count, *http.Response, error) {
	return r.ApiService.GetSystemIntegratorTagsCountExecute(r)
}

/*
GetSystemIntegratorTagsCount Get Count of IntegratorTag

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetSystemIntegratorTagsCountRequest
*/
func (a *IntegratorTagsAPIService) GetSystemIntegratorTagsCount(ctx context.Context) ApiGetSystemIntegratorTagsCountRequest {
	return ApiGetSystemIntegratorTagsCountRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Count
func (a *IntegratorTagsAPIService) GetSystemIntegratorTagsCountExecute(r ApiGetSystemIntegratorTagsCountRequest) (*Count, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Count
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IntegratorTagsAPIService.GetSystemIntegratorTagsCount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/integratorTags/count"

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

type ApiPatchSystemIntegratorTagsByIdRequest struct {
	ctx context.Context
	ApiService *IntegratorTagsAPIService
	id int32
	patchOperation *[]PatchOperation
	clientId *string
}

// List of PatchOperation
func (r ApiPatchSystemIntegratorTagsByIdRequest) PatchOperation(patchOperation []PatchOperation) ApiPatchSystemIntegratorTagsByIdRequest {
	r.patchOperation = &patchOperation
	return r
}

// 
func (r ApiPatchSystemIntegratorTagsByIdRequest) ClientId(clientId string) ApiPatchSystemIntegratorTagsByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiPatchSystemIntegratorTagsByIdRequest) Execute() (*IntegratorTag, *http.Response, error) {
	return r.ApiService.PatchSystemIntegratorTagsByIdExecute(r)
}

/*
PatchSystemIntegratorTagsById Patch IntegratorTag

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id integratorTagId
 @return ApiPatchSystemIntegratorTagsByIdRequest
*/
func (a *IntegratorTagsAPIService) PatchSystemIntegratorTagsById(ctx context.Context, id int32) ApiPatchSystemIntegratorTagsByIdRequest {
	return ApiPatchSystemIntegratorTagsByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return IntegratorTag
func (a *IntegratorTagsAPIService) PatchSystemIntegratorTagsByIdExecute(r ApiPatchSystemIntegratorTagsByIdRequest) (*IntegratorTag, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *IntegratorTag
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IntegratorTagsAPIService.PatchSystemIntegratorTagsById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/integratorTags/{id}"
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

type ApiPostSystemIntegratorTagsRequest struct {
	ctx context.Context
	ApiService *IntegratorTagsAPIService
	integratorTag *IntegratorTag
	clientId *string
}

// tag
func (r ApiPostSystemIntegratorTagsRequest) IntegratorTag(integratorTag IntegratorTag) ApiPostSystemIntegratorTagsRequest {
	r.integratorTag = &integratorTag
	return r
}

// 
func (r ApiPostSystemIntegratorTagsRequest) ClientId(clientId string) ApiPostSystemIntegratorTagsRequest {
	r.clientId = &clientId
	return r
}

func (r ApiPostSystemIntegratorTagsRequest) Execute() (*IntegratorTag, *http.Response, error) {
	return r.ApiService.PostSystemIntegratorTagsExecute(r)
}

/*
PostSystemIntegratorTags Post IntegratorTag

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPostSystemIntegratorTagsRequest
*/
func (a *IntegratorTagsAPIService) PostSystemIntegratorTags(ctx context.Context) ApiPostSystemIntegratorTagsRequest {
	return ApiPostSystemIntegratorTagsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return IntegratorTag
func (a *IntegratorTagsAPIService) PostSystemIntegratorTagsExecute(r ApiPostSystemIntegratorTagsRequest) (*IntegratorTag, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *IntegratorTag
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IntegratorTagsAPIService.PostSystemIntegratorTags")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/integratorTags"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.integratorTag == nil {
		return localVarReturnValue, nil, reportError("integratorTag is required and must be specified")
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
	localVarPostBody = r.integratorTag
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

type ApiPutSystemIntegratorTagsByIdRequest struct {
	ctx context.Context
	ApiService *IntegratorTagsAPIService
	id int32
	integratorTag *IntegratorTag
	clientId *string
}

// tag
func (r ApiPutSystemIntegratorTagsByIdRequest) IntegratorTag(integratorTag IntegratorTag) ApiPutSystemIntegratorTagsByIdRequest {
	r.integratorTag = &integratorTag
	return r
}

// 
func (r ApiPutSystemIntegratorTagsByIdRequest) ClientId(clientId string) ApiPutSystemIntegratorTagsByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiPutSystemIntegratorTagsByIdRequest) Execute() (*IntegratorTag, *http.Response, error) {
	return r.ApiService.PutSystemIntegratorTagsByIdExecute(r)
}

/*
PutSystemIntegratorTagsById Put IntegratorTag

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id integratorTagId
 @return ApiPutSystemIntegratorTagsByIdRequest
*/
func (a *IntegratorTagsAPIService) PutSystemIntegratorTagsById(ctx context.Context, id int32) ApiPutSystemIntegratorTagsByIdRequest {
	return ApiPutSystemIntegratorTagsByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return IntegratorTag
func (a *IntegratorTagsAPIService) PutSystemIntegratorTagsByIdExecute(r ApiPutSystemIntegratorTagsByIdRequest) (*IntegratorTag, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *IntegratorTag
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IntegratorTagsAPIService.PutSystemIntegratorTagsById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/integratorTags/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.integratorTag == nil {
		return localVarReturnValue, nil, reportError("integratorTag is required and must be specified")
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
	localVarPostBody = r.integratorTag
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
