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


// CustomReportParametersAPIService CustomReportParametersAPI service
type CustomReportParametersAPIService service

type ApiDeleteSystemCustomReportsByParentIdParametersByIdRequest struct {
	ctx context.Context
	ApiService *CustomReportParametersAPIService
	id int32
	parentId int32
	clientId *string
}

// 
func (r ApiDeleteSystemCustomReportsByParentIdParametersByIdRequest) ClientId(clientId string) ApiDeleteSystemCustomReportsByParentIdParametersByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiDeleteSystemCustomReportsByParentIdParametersByIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteSystemCustomReportsByParentIdParametersByIdExecute(r)
}

/*
DeleteSystemCustomReportsByParentIdParametersById Delete CustomReportParameter

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id parameterId
 @param parentId customReportId
 @return ApiDeleteSystemCustomReportsByParentIdParametersByIdRequest
*/
func (a *CustomReportParametersAPIService) DeleteSystemCustomReportsByParentIdParametersById(ctx context.Context, id int32, parentId int32) ApiDeleteSystemCustomReportsByParentIdParametersByIdRequest {
	return ApiDeleteSystemCustomReportsByParentIdParametersByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
		parentId: parentId,
	}
}

// Execute executes the request
func (a *CustomReportParametersAPIService) DeleteSystemCustomReportsByParentIdParametersByIdExecute(r ApiDeleteSystemCustomReportsByParentIdParametersByIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomReportParametersAPIService.DeleteSystemCustomReportsByParentIdParametersById")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/customReports/{parentId}/parameters/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"parentId"+"}", url.PathEscape(parameterValueToString(r.parentId, "parentId")), -1)

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

type ApiGetSystemCustomReportsByParentIdParametersRequest struct {
	ctx context.Context
	ApiService *CustomReportParametersAPIService
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
func (r ApiGetSystemCustomReportsByParentIdParametersRequest) Conditions(conditions string) ApiGetSystemCustomReportsByParentIdParametersRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetSystemCustomReportsByParentIdParametersRequest) ChildConditions(childConditions string) ApiGetSystemCustomReportsByParentIdParametersRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetSystemCustomReportsByParentIdParametersRequest) CustomFieldConditions(customFieldConditions string) ApiGetSystemCustomReportsByParentIdParametersRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetSystemCustomReportsByParentIdParametersRequest) OrderBy(orderBy string) ApiGetSystemCustomReportsByParentIdParametersRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetSystemCustomReportsByParentIdParametersRequest) Fields(fields string) ApiGetSystemCustomReportsByParentIdParametersRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetSystemCustomReportsByParentIdParametersRequest) Page(page int32) ApiGetSystemCustomReportsByParentIdParametersRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetSystemCustomReportsByParentIdParametersRequest) PageSize(pageSize int32) ApiGetSystemCustomReportsByParentIdParametersRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetSystemCustomReportsByParentIdParametersRequest) PageId(pageId int32) ApiGetSystemCustomReportsByParentIdParametersRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetSystemCustomReportsByParentIdParametersRequest) ClientId(clientId string) ApiGetSystemCustomReportsByParentIdParametersRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetSystemCustomReportsByParentIdParametersRequest) Execute() ([]CustomReportParameter, *http.Response, error) {
	return r.ApiService.GetSystemCustomReportsByParentIdParametersExecute(r)
}

/*
GetSystemCustomReportsByParentIdParameters Get List of CustomReportParameter

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param parentId customReportId
 @return ApiGetSystemCustomReportsByParentIdParametersRequest
*/
func (a *CustomReportParametersAPIService) GetSystemCustomReportsByParentIdParameters(ctx context.Context, parentId int32) ApiGetSystemCustomReportsByParentIdParametersRequest {
	return ApiGetSystemCustomReportsByParentIdParametersRequest{
		ApiService: a,
		ctx: ctx,
		parentId: parentId,
	}
}

// Execute executes the request
//  @return []CustomReportParameter
func (a *CustomReportParametersAPIService) GetSystemCustomReportsByParentIdParametersExecute(r ApiGetSystemCustomReportsByParentIdParametersRequest) ([]CustomReportParameter, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []CustomReportParameter
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomReportParametersAPIService.GetSystemCustomReportsByParentIdParameters")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/customReports/{parentId}/parameters"
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

type ApiGetSystemCustomReportsByParentIdParametersByIdRequest struct {
	ctx context.Context
	ApiService *CustomReportParametersAPIService
	id int32
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
func (r ApiGetSystemCustomReportsByParentIdParametersByIdRequest) Conditions(conditions string) ApiGetSystemCustomReportsByParentIdParametersByIdRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetSystemCustomReportsByParentIdParametersByIdRequest) ChildConditions(childConditions string) ApiGetSystemCustomReportsByParentIdParametersByIdRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetSystemCustomReportsByParentIdParametersByIdRequest) CustomFieldConditions(customFieldConditions string) ApiGetSystemCustomReportsByParentIdParametersByIdRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetSystemCustomReportsByParentIdParametersByIdRequest) OrderBy(orderBy string) ApiGetSystemCustomReportsByParentIdParametersByIdRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetSystemCustomReportsByParentIdParametersByIdRequest) Fields(fields string) ApiGetSystemCustomReportsByParentIdParametersByIdRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetSystemCustomReportsByParentIdParametersByIdRequest) Page(page int32) ApiGetSystemCustomReportsByParentIdParametersByIdRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetSystemCustomReportsByParentIdParametersByIdRequest) PageSize(pageSize int32) ApiGetSystemCustomReportsByParentIdParametersByIdRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetSystemCustomReportsByParentIdParametersByIdRequest) PageId(pageId int32) ApiGetSystemCustomReportsByParentIdParametersByIdRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetSystemCustomReportsByParentIdParametersByIdRequest) ClientId(clientId string) ApiGetSystemCustomReportsByParentIdParametersByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetSystemCustomReportsByParentIdParametersByIdRequest) Execute() (*CustomReportParameter, *http.Response, error) {
	return r.ApiService.GetSystemCustomReportsByParentIdParametersByIdExecute(r)
}

/*
GetSystemCustomReportsByParentIdParametersById Get CustomReportParameter

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id parameterId
 @param parentId customReportId
 @return ApiGetSystemCustomReportsByParentIdParametersByIdRequest
*/
func (a *CustomReportParametersAPIService) GetSystemCustomReportsByParentIdParametersById(ctx context.Context, id int32, parentId int32) ApiGetSystemCustomReportsByParentIdParametersByIdRequest {
	return ApiGetSystemCustomReportsByParentIdParametersByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
		parentId: parentId,
	}
}

// Execute executes the request
//  @return CustomReportParameter
func (a *CustomReportParametersAPIService) GetSystemCustomReportsByParentIdParametersByIdExecute(r ApiGetSystemCustomReportsByParentIdParametersByIdRequest) (*CustomReportParameter, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CustomReportParameter
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomReportParametersAPIService.GetSystemCustomReportsByParentIdParametersById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/customReports/{parentId}/parameters/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)
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

type ApiGetSystemCustomReportsByParentIdParametersCountRequest struct {
	ctx context.Context
	ApiService *CustomReportParametersAPIService
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
func (r ApiGetSystemCustomReportsByParentIdParametersCountRequest) Conditions(conditions string) ApiGetSystemCustomReportsByParentIdParametersCountRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetSystemCustomReportsByParentIdParametersCountRequest) ChildConditions(childConditions string) ApiGetSystemCustomReportsByParentIdParametersCountRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetSystemCustomReportsByParentIdParametersCountRequest) CustomFieldConditions(customFieldConditions string) ApiGetSystemCustomReportsByParentIdParametersCountRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetSystemCustomReportsByParentIdParametersCountRequest) OrderBy(orderBy string) ApiGetSystemCustomReportsByParentIdParametersCountRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetSystemCustomReportsByParentIdParametersCountRequest) Fields(fields string) ApiGetSystemCustomReportsByParentIdParametersCountRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetSystemCustomReportsByParentIdParametersCountRequest) Page(page int32) ApiGetSystemCustomReportsByParentIdParametersCountRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetSystemCustomReportsByParentIdParametersCountRequest) PageSize(pageSize int32) ApiGetSystemCustomReportsByParentIdParametersCountRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetSystemCustomReportsByParentIdParametersCountRequest) PageId(pageId int32) ApiGetSystemCustomReportsByParentIdParametersCountRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetSystemCustomReportsByParentIdParametersCountRequest) ClientId(clientId string) ApiGetSystemCustomReportsByParentIdParametersCountRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetSystemCustomReportsByParentIdParametersCountRequest) Execute() (*Count, *http.Response, error) {
	return r.ApiService.GetSystemCustomReportsByParentIdParametersCountExecute(r)
}

/*
GetSystemCustomReportsByParentIdParametersCount Get Count of CustomReportParameter

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param parentId customReportId
 @return ApiGetSystemCustomReportsByParentIdParametersCountRequest
*/
func (a *CustomReportParametersAPIService) GetSystemCustomReportsByParentIdParametersCount(ctx context.Context, parentId int32) ApiGetSystemCustomReportsByParentIdParametersCountRequest {
	return ApiGetSystemCustomReportsByParentIdParametersCountRequest{
		ApiService: a,
		ctx: ctx,
		parentId: parentId,
	}
}

// Execute executes the request
//  @return Count
func (a *CustomReportParametersAPIService) GetSystemCustomReportsByParentIdParametersCountExecute(r ApiGetSystemCustomReportsByParentIdParametersCountRequest) (*Count, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Count
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomReportParametersAPIService.GetSystemCustomReportsByParentIdParametersCount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/customReports/{parentId}/parameters/count"
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

type ApiPatchSystemCustomReportsByParentIdParametersByIdRequest struct {
	ctx context.Context
	ApiService *CustomReportParametersAPIService
	id int32
	parentId int32
	patchOperation *[]PatchOperation
	clientId *string
}

// List of PatchOperation
func (r ApiPatchSystemCustomReportsByParentIdParametersByIdRequest) PatchOperation(patchOperation []PatchOperation) ApiPatchSystemCustomReportsByParentIdParametersByIdRequest {
	r.patchOperation = &patchOperation
	return r
}

// 
func (r ApiPatchSystemCustomReportsByParentIdParametersByIdRequest) ClientId(clientId string) ApiPatchSystemCustomReportsByParentIdParametersByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiPatchSystemCustomReportsByParentIdParametersByIdRequest) Execute() (*CustomReportParameter, *http.Response, error) {
	return r.ApiService.PatchSystemCustomReportsByParentIdParametersByIdExecute(r)
}

/*
PatchSystemCustomReportsByParentIdParametersById Patch CustomReportParameter

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id parameterId
 @param parentId customReportId
 @return ApiPatchSystemCustomReportsByParentIdParametersByIdRequest
*/
func (a *CustomReportParametersAPIService) PatchSystemCustomReportsByParentIdParametersById(ctx context.Context, id int32, parentId int32) ApiPatchSystemCustomReportsByParentIdParametersByIdRequest {
	return ApiPatchSystemCustomReportsByParentIdParametersByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
		parentId: parentId,
	}
}

// Execute executes the request
//  @return CustomReportParameter
func (a *CustomReportParametersAPIService) PatchSystemCustomReportsByParentIdParametersByIdExecute(r ApiPatchSystemCustomReportsByParentIdParametersByIdRequest) (*CustomReportParameter, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CustomReportParameter
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomReportParametersAPIService.PatchSystemCustomReportsByParentIdParametersById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/customReports/{parentId}/parameters/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"parentId"+"}", url.PathEscape(parameterValueToString(r.parentId, "parentId")), -1)

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

type ApiPostSystemCustomReportsByParentIdParametersRequest struct {
	ctx context.Context
	ApiService *CustomReportParametersAPIService
	parentId int32
	customReportParameter *CustomReportParameter
	clientId *string
}

// customReportParameter
func (r ApiPostSystemCustomReportsByParentIdParametersRequest) CustomReportParameter(customReportParameter CustomReportParameter) ApiPostSystemCustomReportsByParentIdParametersRequest {
	r.customReportParameter = &customReportParameter
	return r
}

// 
func (r ApiPostSystemCustomReportsByParentIdParametersRequest) ClientId(clientId string) ApiPostSystemCustomReportsByParentIdParametersRequest {
	r.clientId = &clientId
	return r
}

func (r ApiPostSystemCustomReportsByParentIdParametersRequest) Execute() (*CustomReportParameter, *http.Response, error) {
	return r.ApiService.PostSystemCustomReportsByParentIdParametersExecute(r)
}

/*
PostSystemCustomReportsByParentIdParameters Post CustomReportParameter

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param parentId customReportId
 @return ApiPostSystemCustomReportsByParentIdParametersRequest
*/
func (a *CustomReportParametersAPIService) PostSystemCustomReportsByParentIdParameters(ctx context.Context, parentId int32) ApiPostSystemCustomReportsByParentIdParametersRequest {
	return ApiPostSystemCustomReportsByParentIdParametersRequest{
		ApiService: a,
		ctx: ctx,
		parentId: parentId,
	}
}

// Execute executes the request
//  @return CustomReportParameter
func (a *CustomReportParametersAPIService) PostSystemCustomReportsByParentIdParametersExecute(r ApiPostSystemCustomReportsByParentIdParametersRequest) (*CustomReportParameter, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CustomReportParameter
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomReportParametersAPIService.PostSystemCustomReportsByParentIdParameters")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/customReports/{parentId}/parameters"
	localVarPath = strings.Replace(localVarPath, "{"+"parentId"+"}", url.PathEscape(parameterValueToString(r.parentId, "parentId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.customReportParameter == nil {
		return localVarReturnValue, nil, reportError("customReportParameter is required and must be specified")
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
	localVarPostBody = r.customReportParameter
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

type ApiPutSystemCustomReportsByParentIdParametersByIdRequest struct {
	ctx context.Context
	ApiService *CustomReportParametersAPIService
	id int32
	parentId int32
	customReportParameter *CustomReportParameter
	clientId *string
}

// customReportParameter
func (r ApiPutSystemCustomReportsByParentIdParametersByIdRequest) CustomReportParameter(customReportParameter CustomReportParameter) ApiPutSystemCustomReportsByParentIdParametersByIdRequest {
	r.customReportParameter = &customReportParameter
	return r
}

// 
func (r ApiPutSystemCustomReportsByParentIdParametersByIdRequest) ClientId(clientId string) ApiPutSystemCustomReportsByParentIdParametersByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiPutSystemCustomReportsByParentIdParametersByIdRequest) Execute() (*CustomReportParameter, *http.Response, error) {
	return r.ApiService.PutSystemCustomReportsByParentIdParametersByIdExecute(r)
}

/*
PutSystemCustomReportsByParentIdParametersById Put CustomReportParameter

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id parameterId
 @param parentId customReportId
 @return ApiPutSystemCustomReportsByParentIdParametersByIdRequest
*/
func (a *CustomReportParametersAPIService) PutSystemCustomReportsByParentIdParametersById(ctx context.Context, id int32, parentId int32) ApiPutSystemCustomReportsByParentIdParametersByIdRequest {
	return ApiPutSystemCustomReportsByParentIdParametersByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
		parentId: parentId,
	}
}

// Execute executes the request
//  @return CustomReportParameter
func (a *CustomReportParametersAPIService) PutSystemCustomReportsByParentIdParametersByIdExecute(r ApiPutSystemCustomReportsByParentIdParametersByIdRequest) (*CustomReportParameter, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CustomReportParameter
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomReportParametersAPIService.PutSystemCustomReportsByParentIdParametersById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/customReports/{parentId}/parameters/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"parentId"+"}", url.PathEscape(parameterValueToString(r.parentId, "parentId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.customReportParameter == nil {
		return localVarReturnValue, nil, reportError("customReportParameter is required and must be specified")
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
	localVarPostBody = r.customReportParameter
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
