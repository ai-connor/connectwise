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


// CompanyGroupsAPIService CompanyGroupsAPI service
type CompanyGroupsAPIService service

type ApiDeleteCompanyCompaniesByParentIdGroupsByIdRequest struct {
	ctx context.Context
	ApiService *CompanyGroupsAPIService
	id int32
	parentId int32
	clientId *string
}

// 
func (r ApiDeleteCompanyCompaniesByParentIdGroupsByIdRequest) ClientId(clientId string) ApiDeleteCompanyCompaniesByParentIdGroupsByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiDeleteCompanyCompaniesByParentIdGroupsByIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteCompanyCompaniesByParentIdGroupsByIdExecute(r)
}

/*
DeleteCompanyCompaniesByParentIdGroupsById Delete CompanyGroup

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id groupId
 @param parentId companyId
 @return ApiDeleteCompanyCompaniesByParentIdGroupsByIdRequest
*/
func (a *CompanyGroupsAPIService) DeleteCompanyCompaniesByParentIdGroupsById(ctx context.Context, id int32, parentId int32) ApiDeleteCompanyCompaniesByParentIdGroupsByIdRequest {
	return ApiDeleteCompanyCompaniesByParentIdGroupsByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
		parentId: parentId,
	}
}

// Execute executes the request
func (a *CompanyGroupsAPIService) DeleteCompanyCompaniesByParentIdGroupsByIdExecute(r ApiDeleteCompanyCompaniesByParentIdGroupsByIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CompanyGroupsAPIService.DeleteCompanyCompaniesByParentIdGroupsById")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/company/companies/{parentId}/groups/{id}"
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

type ApiGetCompanyCompaniesByParentIdGroupsRequest struct {
	ctx context.Context
	ApiService *CompanyGroupsAPIService
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
func (r ApiGetCompanyCompaniesByParentIdGroupsRequest) Conditions(conditions string) ApiGetCompanyCompaniesByParentIdGroupsRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetCompanyCompaniesByParentIdGroupsRequest) ChildConditions(childConditions string) ApiGetCompanyCompaniesByParentIdGroupsRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetCompanyCompaniesByParentIdGroupsRequest) CustomFieldConditions(customFieldConditions string) ApiGetCompanyCompaniesByParentIdGroupsRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetCompanyCompaniesByParentIdGroupsRequest) OrderBy(orderBy string) ApiGetCompanyCompaniesByParentIdGroupsRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetCompanyCompaniesByParentIdGroupsRequest) Fields(fields string) ApiGetCompanyCompaniesByParentIdGroupsRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetCompanyCompaniesByParentIdGroupsRequest) Page(page int32) ApiGetCompanyCompaniesByParentIdGroupsRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetCompanyCompaniesByParentIdGroupsRequest) PageSize(pageSize int32) ApiGetCompanyCompaniesByParentIdGroupsRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetCompanyCompaniesByParentIdGroupsRequest) PageId(pageId int32) ApiGetCompanyCompaniesByParentIdGroupsRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetCompanyCompaniesByParentIdGroupsRequest) ClientId(clientId string) ApiGetCompanyCompaniesByParentIdGroupsRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetCompanyCompaniesByParentIdGroupsRequest) Execute() ([]CompanyGroup, *http.Response, error) {
	return r.ApiService.GetCompanyCompaniesByParentIdGroupsExecute(r)
}

/*
GetCompanyCompaniesByParentIdGroups Get List of CompanyGroup

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param parentId companyId
 @return ApiGetCompanyCompaniesByParentIdGroupsRequest
*/
func (a *CompanyGroupsAPIService) GetCompanyCompaniesByParentIdGroups(ctx context.Context, parentId int32) ApiGetCompanyCompaniesByParentIdGroupsRequest {
	return ApiGetCompanyCompaniesByParentIdGroupsRequest{
		ApiService: a,
		ctx: ctx,
		parentId: parentId,
	}
}

// Execute executes the request
//  @return []CompanyGroup
func (a *CompanyGroupsAPIService) GetCompanyCompaniesByParentIdGroupsExecute(r ApiGetCompanyCompaniesByParentIdGroupsRequest) ([]CompanyGroup, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []CompanyGroup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CompanyGroupsAPIService.GetCompanyCompaniesByParentIdGroups")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/company/companies/{parentId}/groups"
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

type ApiGetCompanyCompaniesByParentIdGroupsByIdRequest struct {
	ctx context.Context
	ApiService *CompanyGroupsAPIService
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
func (r ApiGetCompanyCompaniesByParentIdGroupsByIdRequest) Conditions(conditions string) ApiGetCompanyCompaniesByParentIdGroupsByIdRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetCompanyCompaniesByParentIdGroupsByIdRequest) ChildConditions(childConditions string) ApiGetCompanyCompaniesByParentIdGroupsByIdRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetCompanyCompaniesByParentIdGroupsByIdRequest) CustomFieldConditions(customFieldConditions string) ApiGetCompanyCompaniesByParentIdGroupsByIdRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetCompanyCompaniesByParentIdGroupsByIdRequest) OrderBy(orderBy string) ApiGetCompanyCompaniesByParentIdGroupsByIdRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetCompanyCompaniesByParentIdGroupsByIdRequest) Fields(fields string) ApiGetCompanyCompaniesByParentIdGroupsByIdRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetCompanyCompaniesByParentIdGroupsByIdRequest) Page(page int32) ApiGetCompanyCompaniesByParentIdGroupsByIdRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetCompanyCompaniesByParentIdGroupsByIdRequest) PageSize(pageSize int32) ApiGetCompanyCompaniesByParentIdGroupsByIdRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetCompanyCompaniesByParentIdGroupsByIdRequest) PageId(pageId int32) ApiGetCompanyCompaniesByParentIdGroupsByIdRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetCompanyCompaniesByParentIdGroupsByIdRequest) ClientId(clientId string) ApiGetCompanyCompaniesByParentIdGroupsByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetCompanyCompaniesByParentIdGroupsByIdRequest) Execute() (*CompanyGroup, *http.Response, error) {
	return r.ApiService.GetCompanyCompaniesByParentIdGroupsByIdExecute(r)
}

/*
GetCompanyCompaniesByParentIdGroupsById Get CompanyGroup

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id groupId
 @param parentId companyId
 @return ApiGetCompanyCompaniesByParentIdGroupsByIdRequest
*/
func (a *CompanyGroupsAPIService) GetCompanyCompaniesByParentIdGroupsById(ctx context.Context, id int32, parentId int32) ApiGetCompanyCompaniesByParentIdGroupsByIdRequest {
	return ApiGetCompanyCompaniesByParentIdGroupsByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
		parentId: parentId,
	}
}

// Execute executes the request
//  @return CompanyGroup
func (a *CompanyGroupsAPIService) GetCompanyCompaniesByParentIdGroupsByIdExecute(r ApiGetCompanyCompaniesByParentIdGroupsByIdRequest) (*CompanyGroup, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CompanyGroup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CompanyGroupsAPIService.GetCompanyCompaniesByParentIdGroupsById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/company/companies/{parentId}/groups/{id}"
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

type ApiGetCompanyCompaniesByParentIdGroupsCountRequest struct {
	ctx context.Context
	ApiService *CompanyGroupsAPIService
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
func (r ApiGetCompanyCompaniesByParentIdGroupsCountRequest) Conditions(conditions string) ApiGetCompanyCompaniesByParentIdGroupsCountRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetCompanyCompaniesByParentIdGroupsCountRequest) ChildConditions(childConditions string) ApiGetCompanyCompaniesByParentIdGroupsCountRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetCompanyCompaniesByParentIdGroupsCountRequest) CustomFieldConditions(customFieldConditions string) ApiGetCompanyCompaniesByParentIdGroupsCountRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetCompanyCompaniesByParentIdGroupsCountRequest) OrderBy(orderBy string) ApiGetCompanyCompaniesByParentIdGroupsCountRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetCompanyCompaniesByParentIdGroupsCountRequest) Fields(fields string) ApiGetCompanyCompaniesByParentIdGroupsCountRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetCompanyCompaniesByParentIdGroupsCountRequest) Page(page int32) ApiGetCompanyCompaniesByParentIdGroupsCountRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetCompanyCompaniesByParentIdGroupsCountRequest) PageSize(pageSize int32) ApiGetCompanyCompaniesByParentIdGroupsCountRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetCompanyCompaniesByParentIdGroupsCountRequest) PageId(pageId int32) ApiGetCompanyCompaniesByParentIdGroupsCountRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetCompanyCompaniesByParentIdGroupsCountRequest) ClientId(clientId string) ApiGetCompanyCompaniesByParentIdGroupsCountRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetCompanyCompaniesByParentIdGroupsCountRequest) Execute() (*Count, *http.Response, error) {
	return r.ApiService.GetCompanyCompaniesByParentIdGroupsCountExecute(r)
}

/*
GetCompanyCompaniesByParentIdGroupsCount Get Count of CompanyGroup

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param parentId companyId
 @return ApiGetCompanyCompaniesByParentIdGroupsCountRequest
*/
func (a *CompanyGroupsAPIService) GetCompanyCompaniesByParentIdGroupsCount(ctx context.Context, parentId int32) ApiGetCompanyCompaniesByParentIdGroupsCountRequest {
	return ApiGetCompanyCompaniesByParentIdGroupsCountRequest{
		ApiService: a,
		ctx: ctx,
		parentId: parentId,
	}
}

// Execute executes the request
//  @return Count
func (a *CompanyGroupsAPIService) GetCompanyCompaniesByParentIdGroupsCountExecute(r ApiGetCompanyCompaniesByParentIdGroupsCountRequest) (*Count, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Count
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CompanyGroupsAPIService.GetCompanyCompaniesByParentIdGroupsCount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/company/companies/{parentId}/groups/count"
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

type ApiPatchCompanyCompaniesByParentIdGroupsByIdRequest struct {
	ctx context.Context
	ApiService *CompanyGroupsAPIService
	id int32
	parentId int32
	patchOperation *[]PatchOperation
	clientId *string
}

// List of PatchOperation
func (r ApiPatchCompanyCompaniesByParentIdGroupsByIdRequest) PatchOperation(patchOperation []PatchOperation) ApiPatchCompanyCompaniesByParentIdGroupsByIdRequest {
	r.patchOperation = &patchOperation
	return r
}

// 
func (r ApiPatchCompanyCompaniesByParentIdGroupsByIdRequest) ClientId(clientId string) ApiPatchCompanyCompaniesByParentIdGroupsByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiPatchCompanyCompaniesByParentIdGroupsByIdRequest) Execute() (*CompanyGroup, *http.Response, error) {
	return r.ApiService.PatchCompanyCompaniesByParentIdGroupsByIdExecute(r)
}

/*
PatchCompanyCompaniesByParentIdGroupsById Patch CompanyGroup

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id groupId
 @param parentId companyId
 @return ApiPatchCompanyCompaniesByParentIdGroupsByIdRequest
*/
func (a *CompanyGroupsAPIService) PatchCompanyCompaniesByParentIdGroupsById(ctx context.Context, id int32, parentId int32) ApiPatchCompanyCompaniesByParentIdGroupsByIdRequest {
	return ApiPatchCompanyCompaniesByParentIdGroupsByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
		parentId: parentId,
	}
}

// Execute executes the request
//  @return CompanyGroup
func (a *CompanyGroupsAPIService) PatchCompanyCompaniesByParentIdGroupsByIdExecute(r ApiPatchCompanyCompaniesByParentIdGroupsByIdRequest) (*CompanyGroup, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CompanyGroup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CompanyGroupsAPIService.PatchCompanyCompaniesByParentIdGroupsById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/company/companies/{parentId}/groups/{id}"
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

type ApiPostCompanyCompaniesByParentIdGroupsRequest struct {
	ctx context.Context
	ApiService *CompanyGroupsAPIService
	parentId int32
	companyGroup *CompanyGroup
	clientId *string
}

// companyGroup
func (r ApiPostCompanyCompaniesByParentIdGroupsRequest) CompanyGroup(companyGroup CompanyGroup) ApiPostCompanyCompaniesByParentIdGroupsRequest {
	r.companyGroup = &companyGroup
	return r
}

// 
func (r ApiPostCompanyCompaniesByParentIdGroupsRequest) ClientId(clientId string) ApiPostCompanyCompaniesByParentIdGroupsRequest {
	r.clientId = &clientId
	return r
}

func (r ApiPostCompanyCompaniesByParentIdGroupsRequest) Execute() (*CompanyGroup, *http.Response, error) {
	return r.ApiService.PostCompanyCompaniesByParentIdGroupsExecute(r)
}

/*
PostCompanyCompaniesByParentIdGroups Post CompanyGroup

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param parentId companyId
 @return ApiPostCompanyCompaniesByParentIdGroupsRequest
*/
func (a *CompanyGroupsAPIService) PostCompanyCompaniesByParentIdGroups(ctx context.Context, parentId int32) ApiPostCompanyCompaniesByParentIdGroupsRequest {
	return ApiPostCompanyCompaniesByParentIdGroupsRequest{
		ApiService: a,
		ctx: ctx,
		parentId: parentId,
	}
}

// Execute executes the request
//  @return CompanyGroup
func (a *CompanyGroupsAPIService) PostCompanyCompaniesByParentIdGroupsExecute(r ApiPostCompanyCompaniesByParentIdGroupsRequest) (*CompanyGroup, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CompanyGroup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CompanyGroupsAPIService.PostCompanyCompaniesByParentIdGroups")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/company/companies/{parentId}/groups"
	localVarPath = strings.Replace(localVarPath, "{"+"parentId"+"}", url.PathEscape(parameterValueToString(r.parentId, "parentId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.companyGroup == nil {
		return localVarReturnValue, nil, reportError("companyGroup is required and must be specified")
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
	localVarPostBody = r.companyGroup
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

type ApiPutCompanyCompaniesByParentIdGroupsByIdRequest struct {
	ctx context.Context
	ApiService *CompanyGroupsAPIService
	id int32
	parentId int32
	companyGroup *CompanyGroup
	clientId *string
}

// companyGroup
func (r ApiPutCompanyCompaniesByParentIdGroupsByIdRequest) CompanyGroup(companyGroup CompanyGroup) ApiPutCompanyCompaniesByParentIdGroupsByIdRequest {
	r.companyGroup = &companyGroup
	return r
}

// 
func (r ApiPutCompanyCompaniesByParentIdGroupsByIdRequest) ClientId(clientId string) ApiPutCompanyCompaniesByParentIdGroupsByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiPutCompanyCompaniesByParentIdGroupsByIdRequest) Execute() (*CompanyGroup, *http.Response, error) {
	return r.ApiService.PutCompanyCompaniesByParentIdGroupsByIdExecute(r)
}

/*
PutCompanyCompaniesByParentIdGroupsById Put CompanyGroup

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id groupId
 @param parentId companyId
 @return ApiPutCompanyCompaniesByParentIdGroupsByIdRequest
*/
func (a *CompanyGroupsAPIService) PutCompanyCompaniesByParentIdGroupsById(ctx context.Context, id int32, parentId int32) ApiPutCompanyCompaniesByParentIdGroupsByIdRequest {
	return ApiPutCompanyCompaniesByParentIdGroupsByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
		parentId: parentId,
	}
}

// Execute executes the request
//  @return CompanyGroup
func (a *CompanyGroupsAPIService) PutCompanyCompaniesByParentIdGroupsByIdExecute(r ApiPutCompanyCompaniesByParentIdGroupsByIdRequest) (*CompanyGroup, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CompanyGroup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CompanyGroupsAPIService.PutCompanyCompaniesByParentIdGroupsById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/company/companies/{parentId}/groups/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"parentId"+"}", url.PathEscape(parameterValueToString(r.parentId, "parentId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.companyGroup == nil {
		return localVarReturnValue, nil, reportError("companyGroup is required and must be specified")
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
	localVarPostBody = r.companyGroup
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
