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


// DepartmentInfosAPIService DepartmentInfosAPI service
type DepartmentInfosAPIService service

type ApiGetSystemInfoDepartmentsRequest struct {
	ctx context.Context
	ApiService *DepartmentInfosAPIService
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
func (r ApiGetSystemInfoDepartmentsRequest) Conditions(conditions string) ApiGetSystemInfoDepartmentsRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetSystemInfoDepartmentsRequest) ChildConditions(childConditions string) ApiGetSystemInfoDepartmentsRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetSystemInfoDepartmentsRequest) CustomFieldConditions(customFieldConditions string) ApiGetSystemInfoDepartmentsRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetSystemInfoDepartmentsRequest) OrderBy(orderBy string) ApiGetSystemInfoDepartmentsRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetSystemInfoDepartmentsRequest) Fields(fields string) ApiGetSystemInfoDepartmentsRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetSystemInfoDepartmentsRequest) Page(page int32) ApiGetSystemInfoDepartmentsRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetSystemInfoDepartmentsRequest) PageSize(pageSize int32) ApiGetSystemInfoDepartmentsRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetSystemInfoDepartmentsRequest) PageId(pageId int32) ApiGetSystemInfoDepartmentsRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetSystemInfoDepartmentsRequest) ClientId(clientId string) ApiGetSystemInfoDepartmentsRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetSystemInfoDepartmentsRequest) Execute() ([]DepartmentInfo, *http.Response, error) {
	return r.ApiService.GetSystemInfoDepartmentsExecute(r)
}

/*
GetSystemInfoDepartments Get List of DepartmentInfo

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetSystemInfoDepartmentsRequest
*/
func (a *DepartmentInfosAPIService) GetSystemInfoDepartments(ctx context.Context) ApiGetSystemInfoDepartmentsRequest {
	return ApiGetSystemInfoDepartmentsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []DepartmentInfo
func (a *DepartmentInfosAPIService) GetSystemInfoDepartmentsExecute(r ApiGetSystemInfoDepartmentsRequest) ([]DepartmentInfo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []DepartmentInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DepartmentInfosAPIService.GetSystemInfoDepartments")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/info/departments"

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

type ApiGetSystemInfoDepartmentsByIdRequest struct {
	ctx context.Context
	ApiService *DepartmentInfosAPIService
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
func (r ApiGetSystemInfoDepartmentsByIdRequest) Conditions(conditions string) ApiGetSystemInfoDepartmentsByIdRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetSystemInfoDepartmentsByIdRequest) ChildConditions(childConditions string) ApiGetSystemInfoDepartmentsByIdRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetSystemInfoDepartmentsByIdRequest) CustomFieldConditions(customFieldConditions string) ApiGetSystemInfoDepartmentsByIdRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetSystemInfoDepartmentsByIdRequest) OrderBy(orderBy string) ApiGetSystemInfoDepartmentsByIdRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetSystemInfoDepartmentsByIdRequest) Fields(fields string) ApiGetSystemInfoDepartmentsByIdRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetSystemInfoDepartmentsByIdRequest) Page(page int32) ApiGetSystemInfoDepartmentsByIdRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetSystemInfoDepartmentsByIdRequest) PageSize(pageSize int32) ApiGetSystemInfoDepartmentsByIdRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetSystemInfoDepartmentsByIdRequest) PageId(pageId int32) ApiGetSystemInfoDepartmentsByIdRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetSystemInfoDepartmentsByIdRequest) ClientId(clientId string) ApiGetSystemInfoDepartmentsByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetSystemInfoDepartmentsByIdRequest) Execute() (*DepartmentInfo, *http.Response, error) {
	return r.ApiService.GetSystemInfoDepartmentsByIdExecute(r)
}

/*
GetSystemInfoDepartmentsById Get DepartmentInfo

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id departmentId
 @return ApiGetSystemInfoDepartmentsByIdRequest
*/
func (a *DepartmentInfosAPIService) GetSystemInfoDepartmentsById(ctx context.Context, id int32) ApiGetSystemInfoDepartmentsByIdRequest {
	return ApiGetSystemInfoDepartmentsByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return DepartmentInfo
func (a *DepartmentInfosAPIService) GetSystemInfoDepartmentsByIdExecute(r ApiGetSystemInfoDepartmentsByIdRequest) (*DepartmentInfo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DepartmentInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DepartmentInfosAPIService.GetSystemInfoDepartmentsById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/info/departments/{id}"
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

type ApiGetSystemInfoDepartmentsCountRequest struct {
	ctx context.Context
	ApiService *DepartmentInfosAPIService
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
func (r ApiGetSystemInfoDepartmentsCountRequest) Conditions(conditions string) ApiGetSystemInfoDepartmentsCountRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetSystemInfoDepartmentsCountRequest) ChildConditions(childConditions string) ApiGetSystemInfoDepartmentsCountRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetSystemInfoDepartmentsCountRequest) CustomFieldConditions(customFieldConditions string) ApiGetSystemInfoDepartmentsCountRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetSystemInfoDepartmentsCountRequest) OrderBy(orderBy string) ApiGetSystemInfoDepartmentsCountRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetSystemInfoDepartmentsCountRequest) Fields(fields string) ApiGetSystemInfoDepartmentsCountRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetSystemInfoDepartmentsCountRequest) Page(page int32) ApiGetSystemInfoDepartmentsCountRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetSystemInfoDepartmentsCountRequest) PageSize(pageSize int32) ApiGetSystemInfoDepartmentsCountRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetSystemInfoDepartmentsCountRequest) PageId(pageId int32) ApiGetSystemInfoDepartmentsCountRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetSystemInfoDepartmentsCountRequest) ClientId(clientId string) ApiGetSystemInfoDepartmentsCountRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetSystemInfoDepartmentsCountRequest) Execute() (*Count, *http.Response, error) {
	return r.ApiService.GetSystemInfoDepartmentsCountExecute(r)
}

/*
GetSystemInfoDepartmentsCount Get Count of DepartmentInfo

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetSystemInfoDepartmentsCountRequest
*/
func (a *DepartmentInfosAPIService) GetSystemInfoDepartmentsCount(ctx context.Context) ApiGetSystemInfoDepartmentsCountRequest {
	return ApiGetSystemInfoDepartmentsCountRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Count
func (a *DepartmentInfosAPIService) GetSystemInfoDepartmentsCountExecute(r ApiGetSystemInfoDepartmentsCountRequest) (*Count, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Count
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DepartmentInfosAPIService.GetSystemInfoDepartmentsCount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/info/departments/count"

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
