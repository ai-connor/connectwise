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


// KPIsAPIService KPIsAPI service
type KPIsAPIService service

type ApiGetSystemKpisRequest struct {
	ctx context.Context
	ApiService *KPIsAPIService
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
func (r ApiGetSystemKpisRequest) Conditions(conditions string) ApiGetSystemKpisRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetSystemKpisRequest) ChildConditions(childConditions string) ApiGetSystemKpisRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetSystemKpisRequest) CustomFieldConditions(customFieldConditions string) ApiGetSystemKpisRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetSystemKpisRequest) OrderBy(orderBy string) ApiGetSystemKpisRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetSystemKpisRequest) Fields(fields string) ApiGetSystemKpisRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetSystemKpisRequest) Page(page int32) ApiGetSystemKpisRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetSystemKpisRequest) PageSize(pageSize int32) ApiGetSystemKpisRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetSystemKpisRequest) PageId(pageId int32) ApiGetSystemKpisRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetSystemKpisRequest) ClientId(clientId string) ApiGetSystemKpisRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetSystemKpisRequest) Execute() ([]KPI, *http.Response, error) {
	return r.ApiService.GetSystemKpisExecute(r)
}

/*
GetSystemKpis Get List of KPI

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetSystemKpisRequest
*/
func (a *KPIsAPIService) GetSystemKpis(ctx context.Context) ApiGetSystemKpisRequest {
	return ApiGetSystemKpisRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []KPI
func (a *KPIsAPIService) GetSystemKpisExecute(r ApiGetSystemKpisRequest) ([]KPI, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []KPI
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KPIsAPIService.GetSystemKpis")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/kpis"

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

type ApiGetSystemKpisByIdRequest struct {
	ctx context.Context
	ApiService *KPIsAPIService
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
func (r ApiGetSystemKpisByIdRequest) Conditions(conditions string) ApiGetSystemKpisByIdRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetSystemKpisByIdRequest) ChildConditions(childConditions string) ApiGetSystemKpisByIdRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetSystemKpisByIdRequest) CustomFieldConditions(customFieldConditions string) ApiGetSystemKpisByIdRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetSystemKpisByIdRequest) OrderBy(orderBy string) ApiGetSystemKpisByIdRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetSystemKpisByIdRequest) Fields(fields string) ApiGetSystemKpisByIdRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetSystemKpisByIdRequest) Page(page int32) ApiGetSystemKpisByIdRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetSystemKpisByIdRequest) PageSize(pageSize int32) ApiGetSystemKpisByIdRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetSystemKpisByIdRequest) PageId(pageId int32) ApiGetSystemKpisByIdRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetSystemKpisByIdRequest) ClientId(clientId string) ApiGetSystemKpisByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetSystemKpisByIdRequest) Execute() (*KPI, *http.Response, error) {
	return r.ApiService.GetSystemKpisByIdExecute(r)
}

/*
GetSystemKpisById Get KPI

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id kpiId
 @return ApiGetSystemKpisByIdRequest
*/
func (a *KPIsAPIService) GetSystemKpisById(ctx context.Context, id int32) ApiGetSystemKpisByIdRequest {
	return ApiGetSystemKpisByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return KPI
func (a *KPIsAPIService) GetSystemKpisByIdExecute(r ApiGetSystemKpisByIdRequest) (*KPI, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KPI
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KPIsAPIService.GetSystemKpisById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/kpis/{id}"
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

type ApiGetSystemKpisCountRequest struct {
	ctx context.Context
	ApiService *KPIsAPIService
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
func (r ApiGetSystemKpisCountRequest) Conditions(conditions string) ApiGetSystemKpisCountRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetSystemKpisCountRequest) ChildConditions(childConditions string) ApiGetSystemKpisCountRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetSystemKpisCountRequest) CustomFieldConditions(customFieldConditions string) ApiGetSystemKpisCountRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetSystemKpisCountRequest) OrderBy(orderBy string) ApiGetSystemKpisCountRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetSystemKpisCountRequest) Fields(fields string) ApiGetSystemKpisCountRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetSystemKpisCountRequest) Page(page int32) ApiGetSystemKpisCountRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetSystemKpisCountRequest) PageSize(pageSize int32) ApiGetSystemKpisCountRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetSystemKpisCountRequest) PageId(pageId int32) ApiGetSystemKpisCountRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetSystemKpisCountRequest) ClientId(clientId string) ApiGetSystemKpisCountRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetSystemKpisCountRequest) Execute() (*Count, *http.Response, error) {
	return r.ApiService.GetSystemKpisCountExecute(r)
}

/*
GetSystemKpisCount Get Count of KPI

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetSystemKpisCountRequest
*/
func (a *KPIsAPIService) GetSystemKpisCount(ctx context.Context) ApiGetSystemKpisCountRequest {
	return ApiGetSystemKpisCountRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Count
func (a *KPIsAPIService) GetSystemKpisCountExecute(r ApiGetSystemKpisCountRequest) (*Count, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Count
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KPIsAPIService.GetSystemKpisCount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/kpis/count"

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
