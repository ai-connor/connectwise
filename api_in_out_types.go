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


// InOutTypesAPIService InOutTypesAPI service
type InOutTypesAPIService service

type ApiDeleteSystemInOutTypesByIdRequest struct {
	ctx context.Context
	ApiService *InOutTypesAPIService
	id int32
	clientId *string
}

// 
func (r ApiDeleteSystemInOutTypesByIdRequest) ClientId(clientId string) ApiDeleteSystemInOutTypesByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiDeleteSystemInOutTypesByIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteSystemInOutTypesByIdExecute(r)
}

/*
DeleteSystemInOutTypesById Delete InOutType

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id inOutTypeId
 @return ApiDeleteSystemInOutTypesByIdRequest
*/
func (a *InOutTypesAPIService) DeleteSystemInOutTypesById(ctx context.Context, id int32) ApiDeleteSystemInOutTypesByIdRequest {
	return ApiDeleteSystemInOutTypesByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *InOutTypesAPIService) DeleteSystemInOutTypesByIdExecute(r ApiDeleteSystemInOutTypesByIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InOutTypesAPIService.DeleteSystemInOutTypesById")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/inOutTypes/{id}"
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

type ApiGetSystemInOutTypesRequest struct {
	ctx context.Context
	ApiService *InOutTypesAPIService
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
func (r ApiGetSystemInOutTypesRequest) Conditions(conditions string) ApiGetSystemInOutTypesRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetSystemInOutTypesRequest) ChildConditions(childConditions string) ApiGetSystemInOutTypesRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetSystemInOutTypesRequest) CustomFieldConditions(customFieldConditions string) ApiGetSystemInOutTypesRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetSystemInOutTypesRequest) OrderBy(orderBy string) ApiGetSystemInOutTypesRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetSystemInOutTypesRequest) Fields(fields string) ApiGetSystemInOutTypesRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetSystemInOutTypesRequest) Page(page int32) ApiGetSystemInOutTypesRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetSystemInOutTypesRequest) PageSize(pageSize int32) ApiGetSystemInOutTypesRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetSystemInOutTypesRequest) PageId(pageId int32) ApiGetSystemInOutTypesRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetSystemInOutTypesRequest) ClientId(clientId string) ApiGetSystemInOutTypesRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetSystemInOutTypesRequest) Execute() ([]InOutType, *http.Response, error) {
	return r.ApiService.GetSystemInOutTypesExecute(r)
}

/*
GetSystemInOutTypes Get List of InOutType

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetSystemInOutTypesRequest
*/
func (a *InOutTypesAPIService) GetSystemInOutTypes(ctx context.Context) ApiGetSystemInOutTypesRequest {
	return ApiGetSystemInOutTypesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []InOutType
func (a *InOutTypesAPIService) GetSystemInOutTypesExecute(r ApiGetSystemInOutTypesRequest) ([]InOutType, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []InOutType
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InOutTypesAPIService.GetSystemInOutTypes")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/inOutTypes"

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

type ApiGetSystemInOutTypesByIdRequest struct {
	ctx context.Context
	ApiService *InOutTypesAPIService
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
func (r ApiGetSystemInOutTypesByIdRequest) Conditions(conditions string) ApiGetSystemInOutTypesByIdRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetSystemInOutTypesByIdRequest) ChildConditions(childConditions string) ApiGetSystemInOutTypesByIdRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetSystemInOutTypesByIdRequest) CustomFieldConditions(customFieldConditions string) ApiGetSystemInOutTypesByIdRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetSystemInOutTypesByIdRequest) OrderBy(orderBy string) ApiGetSystemInOutTypesByIdRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetSystemInOutTypesByIdRequest) Fields(fields string) ApiGetSystemInOutTypesByIdRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetSystemInOutTypesByIdRequest) Page(page int32) ApiGetSystemInOutTypesByIdRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetSystemInOutTypesByIdRequest) PageSize(pageSize int32) ApiGetSystemInOutTypesByIdRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetSystemInOutTypesByIdRequest) PageId(pageId int32) ApiGetSystemInOutTypesByIdRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetSystemInOutTypesByIdRequest) ClientId(clientId string) ApiGetSystemInOutTypesByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetSystemInOutTypesByIdRequest) Execute() (*InOutType, *http.Response, error) {
	return r.ApiService.GetSystemInOutTypesByIdExecute(r)
}

/*
GetSystemInOutTypesById Get InOutType

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id inOutTypeId
 @return ApiGetSystemInOutTypesByIdRequest
*/
func (a *InOutTypesAPIService) GetSystemInOutTypesById(ctx context.Context, id int32) ApiGetSystemInOutTypesByIdRequest {
	return ApiGetSystemInOutTypesByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return InOutType
func (a *InOutTypesAPIService) GetSystemInOutTypesByIdExecute(r ApiGetSystemInOutTypesByIdRequest) (*InOutType, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InOutType
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InOutTypesAPIService.GetSystemInOutTypesById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/inOutTypes/{id}"
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

type ApiGetSystemInOutTypesCountRequest struct {
	ctx context.Context
	ApiService *InOutTypesAPIService
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
func (r ApiGetSystemInOutTypesCountRequest) Conditions(conditions string) ApiGetSystemInOutTypesCountRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetSystemInOutTypesCountRequest) ChildConditions(childConditions string) ApiGetSystemInOutTypesCountRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetSystemInOutTypesCountRequest) CustomFieldConditions(customFieldConditions string) ApiGetSystemInOutTypesCountRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetSystemInOutTypesCountRequest) OrderBy(orderBy string) ApiGetSystemInOutTypesCountRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetSystemInOutTypesCountRequest) Fields(fields string) ApiGetSystemInOutTypesCountRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetSystemInOutTypesCountRequest) Page(page int32) ApiGetSystemInOutTypesCountRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetSystemInOutTypesCountRequest) PageSize(pageSize int32) ApiGetSystemInOutTypesCountRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetSystemInOutTypesCountRequest) PageId(pageId int32) ApiGetSystemInOutTypesCountRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetSystemInOutTypesCountRequest) ClientId(clientId string) ApiGetSystemInOutTypesCountRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetSystemInOutTypesCountRequest) Execute() (*Count, *http.Response, error) {
	return r.ApiService.GetSystemInOutTypesCountExecute(r)
}

/*
GetSystemInOutTypesCount Get Count of InOutType

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetSystemInOutTypesCountRequest
*/
func (a *InOutTypesAPIService) GetSystemInOutTypesCount(ctx context.Context) ApiGetSystemInOutTypesCountRequest {
	return ApiGetSystemInOutTypesCountRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Count
func (a *InOutTypesAPIService) GetSystemInOutTypesCountExecute(r ApiGetSystemInOutTypesCountRequest) (*Count, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Count
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InOutTypesAPIService.GetSystemInOutTypesCount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/inOutTypes/count"

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

type ApiPatchSystemInOutTypesByIdRequest struct {
	ctx context.Context
	ApiService *InOutTypesAPIService
	id int32
	patchOperation *[]PatchOperation
	clientId *string
}

// List of PatchOperation
func (r ApiPatchSystemInOutTypesByIdRequest) PatchOperation(patchOperation []PatchOperation) ApiPatchSystemInOutTypesByIdRequest {
	r.patchOperation = &patchOperation
	return r
}

// 
func (r ApiPatchSystemInOutTypesByIdRequest) ClientId(clientId string) ApiPatchSystemInOutTypesByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiPatchSystemInOutTypesByIdRequest) Execute() (*InOutType, *http.Response, error) {
	return r.ApiService.PatchSystemInOutTypesByIdExecute(r)
}

/*
PatchSystemInOutTypesById Patch InOutType

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id inOutTypeId
 @return ApiPatchSystemInOutTypesByIdRequest
*/
func (a *InOutTypesAPIService) PatchSystemInOutTypesById(ctx context.Context, id int32) ApiPatchSystemInOutTypesByIdRequest {
	return ApiPatchSystemInOutTypesByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return InOutType
func (a *InOutTypesAPIService) PatchSystemInOutTypesByIdExecute(r ApiPatchSystemInOutTypesByIdRequest) (*InOutType, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InOutType
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InOutTypesAPIService.PatchSystemInOutTypesById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/inOutTypes/{id}"
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

type ApiPostSystemInOutTypesRequest struct {
	ctx context.Context
	ApiService *InOutTypesAPIService
	inOutType *InOutType
	clientId *string
}

// inOutType
func (r ApiPostSystemInOutTypesRequest) InOutType(inOutType InOutType) ApiPostSystemInOutTypesRequest {
	r.inOutType = &inOutType
	return r
}

// 
func (r ApiPostSystemInOutTypesRequest) ClientId(clientId string) ApiPostSystemInOutTypesRequest {
	r.clientId = &clientId
	return r
}

func (r ApiPostSystemInOutTypesRequest) Execute() (*InOutType, *http.Response, error) {
	return r.ApiService.PostSystemInOutTypesExecute(r)
}

/*
PostSystemInOutTypes Post InOutType

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPostSystemInOutTypesRequest
*/
func (a *InOutTypesAPIService) PostSystemInOutTypes(ctx context.Context) ApiPostSystemInOutTypesRequest {
	return ApiPostSystemInOutTypesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return InOutType
func (a *InOutTypesAPIService) PostSystemInOutTypesExecute(r ApiPostSystemInOutTypesRequest) (*InOutType, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InOutType
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InOutTypesAPIService.PostSystemInOutTypes")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/inOutTypes"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.inOutType == nil {
		return localVarReturnValue, nil, reportError("inOutType is required and must be specified")
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
	localVarPostBody = r.inOutType
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

type ApiPutSystemInOutTypesByIdRequest struct {
	ctx context.Context
	ApiService *InOutTypesAPIService
	id int32
	inOutType *InOutType
	clientId *string
}

// inOutType
func (r ApiPutSystemInOutTypesByIdRequest) InOutType(inOutType InOutType) ApiPutSystemInOutTypesByIdRequest {
	r.inOutType = &inOutType
	return r
}

// 
func (r ApiPutSystemInOutTypesByIdRequest) ClientId(clientId string) ApiPutSystemInOutTypesByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiPutSystemInOutTypesByIdRequest) Execute() (*InOutType, *http.Response, error) {
	return r.ApiService.PutSystemInOutTypesByIdExecute(r)
}

/*
PutSystemInOutTypesById Put InOutType

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id inOutTypeId
 @return ApiPutSystemInOutTypesByIdRequest
*/
func (a *InOutTypesAPIService) PutSystemInOutTypesById(ctx context.Context, id int32) ApiPutSystemInOutTypesByIdRequest {
	return ApiPutSystemInOutTypesByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return InOutType
func (a *InOutTypesAPIService) PutSystemInOutTypesByIdExecute(r ApiPutSystemInOutTypesByIdRequest) (*InOutType, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InOutType
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InOutTypesAPIService.PutSystemInOutTypesById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/inOutTypes/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.inOutType == nil {
		return localVarReturnValue, nil, reportError("inOutType is required and must be specified")
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
	localVarPostBody = r.inOutType
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
