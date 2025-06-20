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


// CatalogComponentsAPIService CatalogComponentsAPI service
type CatalogComponentsAPIService service

type ApiDeleteProcurementCatalogByParentIdComponentsByIdRequest struct {
	ctx context.Context
	ApiService *CatalogComponentsAPIService
	id int32
	parentId int32
	clientId *string
}

// 
func (r ApiDeleteProcurementCatalogByParentIdComponentsByIdRequest) ClientId(clientId string) ApiDeleteProcurementCatalogByParentIdComponentsByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiDeleteProcurementCatalogByParentIdComponentsByIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteProcurementCatalogByParentIdComponentsByIdExecute(r)
}

/*
DeleteProcurementCatalogByParentIdComponentsById Delete CatalogComponent

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id componentId
 @param parentId catalogId
 @return ApiDeleteProcurementCatalogByParentIdComponentsByIdRequest
*/
func (a *CatalogComponentsAPIService) DeleteProcurementCatalogByParentIdComponentsById(ctx context.Context, id int32, parentId int32) ApiDeleteProcurementCatalogByParentIdComponentsByIdRequest {
	return ApiDeleteProcurementCatalogByParentIdComponentsByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
		parentId: parentId,
	}
}

// Execute executes the request
func (a *CatalogComponentsAPIService) DeleteProcurementCatalogByParentIdComponentsByIdExecute(r ApiDeleteProcurementCatalogByParentIdComponentsByIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CatalogComponentsAPIService.DeleteProcurementCatalogByParentIdComponentsById")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/procurement/catalog/{parentId}/components/{id}"
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

type ApiGetProcurementCatalogByParentIdComponentsRequest struct {
	ctx context.Context
	ApiService *CatalogComponentsAPIService
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
func (r ApiGetProcurementCatalogByParentIdComponentsRequest) Conditions(conditions string) ApiGetProcurementCatalogByParentIdComponentsRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetProcurementCatalogByParentIdComponentsRequest) ChildConditions(childConditions string) ApiGetProcurementCatalogByParentIdComponentsRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetProcurementCatalogByParentIdComponentsRequest) CustomFieldConditions(customFieldConditions string) ApiGetProcurementCatalogByParentIdComponentsRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetProcurementCatalogByParentIdComponentsRequest) OrderBy(orderBy string) ApiGetProcurementCatalogByParentIdComponentsRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetProcurementCatalogByParentIdComponentsRequest) Fields(fields string) ApiGetProcurementCatalogByParentIdComponentsRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetProcurementCatalogByParentIdComponentsRequest) Page(page int32) ApiGetProcurementCatalogByParentIdComponentsRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetProcurementCatalogByParentIdComponentsRequest) PageSize(pageSize int32) ApiGetProcurementCatalogByParentIdComponentsRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetProcurementCatalogByParentIdComponentsRequest) PageId(pageId int32) ApiGetProcurementCatalogByParentIdComponentsRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetProcurementCatalogByParentIdComponentsRequest) ClientId(clientId string) ApiGetProcurementCatalogByParentIdComponentsRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetProcurementCatalogByParentIdComponentsRequest) Execute() ([]CatalogComponent, *http.Response, error) {
	return r.ApiService.GetProcurementCatalogByParentIdComponentsExecute(r)
}

/*
GetProcurementCatalogByParentIdComponents Get List of CatalogComponent

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param parentId catalogId
 @return ApiGetProcurementCatalogByParentIdComponentsRequest
*/
func (a *CatalogComponentsAPIService) GetProcurementCatalogByParentIdComponents(ctx context.Context, parentId int32) ApiGetProcurementCatalogByParentIdComponentsRequest {
	return ApiGetProcurementCatalogByParentIdComponentsRequest{
		ApiService: a,
		ctx: ctx,
		parentId: parentId,
	}
}

// Execute executes the request
//  @return []CatalogComponent
func (a *CatalogComponentsAPIService) GetProcurementCatalogByParentIdComponentsExecute(r ApiGetProcurementCatalogByParentIdComponentsRequest) ([]CatalogComponent, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []CatalogComponent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CatalogComponentsAPIService.GetProcurementCatalogByParentIdComponents")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/procurement/catalog/{parentId}/components"
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

type ApiGetProcurementCatalogByParentIdComponentsByIdRequest struct {
	ctx context.Context
	ApiService *CatalogComponentsAPIService
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
func (r ApiGetProcurementCatalogByParentIdComponentsByIdRequest) Conditions(conditions string) ApiGetProcurementCatalogByParentIdComponentsByIdRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetProcurementCatalogByParentIdComponentsByIdRequest) ChildConditions(childConditions string) ApiGetProcurementCatalogByParentIdComponentsByIdRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetProcurementCatalogByParentIdComponentsByIdRequest) CustomFieldConditions(customFieldConditions string) ApiGetProcurementCatalogByParentIdComponentsByIdRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetProcurementCatalogByParentIdComponentsByIdRequest) OrderBy(orderBy string) ApiGetProcurementCatalogByParentIdComponentsByIdRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetProcurementCatalogByParentIdComponentsByIdRequest) Fields(fields string) ApiGetProcurementCatalogByParentIdComponentsByIdRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetProcurementCatalogByParentIdComponentsByIdRequest) Page(page int32) ApiGetProcurementCatalogByParentIdComponentsByIdRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetProcurementCatalogByParentIdComponentsByIdRequest) PageSize(pageSize int32) ApiGetProcurementCatalogByParentIdComponentsByIdRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetProcurementCatalogByParentIdComponentsByIdRequest) PageId(pageId int32) ApiGetProcurementCatalogByParentIdComponentsByIdRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetProcurementCatalogByParentIdComponentsByIdRequest) ClientId(clientId string) ApiGetProcurementCatalogByParentIdComponentsByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetProcurementCatalogByParentIdComponentsByIdRequest) Execute() (*CatalogComponent, *http.Response, error) {
	return r.ApiService.GetProcurementCatalogByParentIdComponentsByIdExecute(r)
}

/*
GetProcurementCatalogByParentIdComponentsById Get CatalogComponent

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id componentId
 @param parentId catalogId
 @return ApiGetProcurementCatalogByParentIdComponentsByIdRequest
*/
func (a *CatalogComponentsAPIService) GetProcurementCatalogByParentIdComponentsById(ctx context.Context, id int32, parentId int32) ApiGetProcurementCatalogByParentIdComponentsByIdRequest {
	return ApiGetProcurementCatalogByParentIdComponentsByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
		parentId: parentId,
	}
}

// Execute executes the request
//  @return CatalogComponent
func (a *CatalogComponentsAPIService) GetProcurementCatalogByParentIdComponentsByIdExecute(r ApiGetProcurementCatalogByParentIdComponentsByIdRequest) (*CatalogComponent, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CatalogComponent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CatalogComponentsAPIService.GetProcurementCatalogByParentIdComponentsById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/procurement/catalog/{parentId}/components/{id}"
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

type ApiGetProcurementCatalogByParentIdComponentsCountRequest struct {
	ctx context.Context
	ApiService *CatalogComponentsAPIService
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
func (r ApiGetProcurementCatalogByParentIdComponentsCountRequest) Conditions(conditions string) ApiGetProcurementCatalogByParentIdComponentsCountRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetProcurementCatalogByParentIdComponentsCountRequest) ChildConditions(childConditions string) ApiGetProcurementCatalogByParentIdComponentsCountRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetProcurementCatalogByParentIdComponentsCountRequest) CustomFieldConditions(customFieldConditions string) ApiGetProcurementCatalogByParentIdComponentsCountRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetProcurementCatalogByParentIdComponentsCountRequest) OrderBy(orderBy string) ApiGetProcurementCatalogByParentIdComponentsCountRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetProcurementCatalogByParentIdComponentsCountRequest) Fields(fields string) ApiGetProcurementCatalogByParentIdComponentsCountRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetProcurementCatalogByParentIdComponentsCountRequest) Page(page int32) ApiGetProcurementCatalogByParentIdComponentsCountRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetProcurementCatalogByParentIdComponentsCountRequest) PageSize(pageSize int32) ApiGetProcurementCatalogByParentIdComponentsCountRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetProcurementCatalogByParentIdComponentsCountRequest) PageId(pageId int32) ApiGetProcurementCatalogByParentIdComponentsCountRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetProcurementCatalogByParentIdComponentsCountRequest) ClientId(clientId string) ApiGetProcurementCatalogByParentIdComponentsCountRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetProcurementCatalogByParentIdComponentsCountRequest) Execute() (*Count, *http.Response, error) {
	return r.ApiService.GetProcurementCatalogByParentIdComponentsCountExecute(r)
}

/*
GetProcurementCatalogByParentIdComponentsCount Get Count of CatalogComponent

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param parentId catalogId
 @return ApiGetProcurementCatalogByParentIdComponentsCountRequest
*/
func (a *CatalogComponentsAPIService) GetProcurementCatalogByParentIdComponentsCount(ctx context.Context, parentId int32) ApiGetProcurementCatalogByParentIdComponentsCountRequest {
	return ApiGetProcurementCatalogByParentIdComponentsCountRequest{
		ApiService: a,
		ctx: ctx,
		parentId: parentId,
	}
}

// Execute executes the request
//  @return Count
func (a *CatalogComponentsAPIService) GetProcurementCatalogByParentIdComponentsCountExecute(r ApiGetProcurementCatalogByParentIdComponentsCountRequest) (*Count, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Count
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CatalogComponentsAPIService.GetProcurementCatalogByParentIdComponentsCount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/procurement/catalog/{parentId}/components/count"
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

type ApiPatchProcurementCatalogByParentIdComponentsByIdRequest struct {
	ctx context.Context
	ApiService *CatalogComponentsAPIService
	id int32
	parentId int32
	patchOperation *[]PatchOperation
	clientId *string
}

// List of PatchOperation
func (r ApiPatchProcurementCatalogByParentIdComponentsByIdRequest) PatchOperation(patchOperation []PatchOperation) ApiPatchProcurementCatalogByParentIdComponentsByIdRequest {
	r.patchOperation = &patchOperation
	return r
}

// 
func (r ApiPatchProcurementCatalogByParentIdComponentsByIdRequest) ClientId(clientId string) ApiPatchProcurementCatalogByParentIdComponentsByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiPatchProcurementCatalogByParentIdComponentsByIdRequest) Execute() (*CatalogComponent, *http.Response, error) {
	return r.ApiService.PatchProcurementCatalogByParentIdComponentsByIdExecute(r)
}

/*
PatchProcurementCatalogByParentIdComponentsById Patch CatalogComponent

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id componentId
 @param parentId catalogId
 @return ApiPatchProcurementCatalogByParentIdComponentsByIdRequest
*/
func (a *CatalogComponentsAPIService) PatchProcurementCatalogByParentIdComponentsById(ctx context.Context, id int32, parentId int32) ApiPatchProcurementCatalogByParentIdComponentsByIdRequest {
	return ApiPatchProcurementCatalogByParentIdComponentsByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
		parentId: parentId,
	}
}

// Execute executes the request
//  @return CatalogComponent
func (a *CatalogComponentsAPIService) PatchProcurementCatalogByParentIdComponentsByIdExecute(r ApiPatchProcurementCatalogByParentIdComponentsByIdRequest) (*CatalogComponent, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CatalogComponent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CatalogComponentsAPIService.PatchProcurementCatalogByParentIdComponentsById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/procurement/catalog/{parentId}/components/{id}"
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

type ApiPostProcurementCatalogByParentIdComponentsRequest struct {
	ctx context.Context
	ApiService *CatalogComponentsAPIService
	parentId int32
	catalogComponent *CatalogComponent
	clientId *string
}

// catalogComponent
func (r ApiPostProcurementCatalogByParentIdComponentsRequest) CatalogComponent(catalogComponent CatalogComponent) ApiPostProcurementCatalogByParentIdComponentsRequest {
	r.catalogComponent = &catalogComponent
	return r
}

// 
func (r ApiPostProcurementCatalogByParentIdComponentsRequest) ClientId(clientId string) ApiPostProcurementCatalogByParentIdComponentsRequest {
	r.clientId = &clientId
	return r
}

func (r ApiPostProcurementCatalogByParentIdComponentsRequest) Execute() (*CatalogComponent, *http.Response, error) {
	return r.ApiService.PostProcurementCatalogByParentIdComponentsExecute(r)
}

/*
PostProcurementCatalogByParentIdComponents Post CatalogComponent

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param parentId catalogId
 @return ApiPostProcurementCatalogByParentIdComponentsRequest
*/
func (a *CatalogComponentsAPIService) PostProcurementCatalogByParentIdComponents(ctx context.Context, parentId int32) ApiPostProcurementCatalogByParentIdComponentsRequest {
	return ApiPostProcurementCatalogByParentIdComponentsRequest{
		ApiService: a,
		ctx: ctx,
		parentId: parentId,
	}
}

// Execute executes the request
//  @return CatalogComponent
func (a *CatalogComponentsAPIService) PostProcurementCatalogByParentIdComponentsExecute(r ApiPostProcurementCatalogByParentIdComponentsRequest) (*CatalogComponent, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CatalogComponent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CatalogComponentsAPIService.PostProcurementCatalogByParentIdComponents")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/procurement/catalog/{parentId}/components"
	localVarPath = strings.Replace(localVarPath, "{"+"parentId"+"}", url.PathEscape(parameterValueToString(r.parentId, "parentId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.catalogComponent == nil {
		return localVarReturnValue, nil, reportError("catalogComponent is required and must be specified")
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
	localVarPostBody = r.catalogComponent
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

type ApiPutProcurementCatalogByParentIdComponentsByIdRequest struct {
	ctx context.Context
	ApiService *CatalogComponentsAPIService
	id int32
	parentId int32
	catalogComponent *CatalogComponent
	clientId *string
}

// catalogComponent
func (r ApiPutProcurementCatalogByParentIdComponentsByIdRequest) CatalogComponent(catalogComponent CatalogComponent) ApiPutProcurementCatalogByParentIdComponentsByIdRequest {
	r.catalogComponent = &catalogComponent
	return r
}

// 
func (r ApiPutProcurementCatalogByParentIdComponentsByIdRequest) ClientId(clientId string) ApiPutProcurementCatalogByParentIdComponentsByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiPutProcurementCatalogByParentIdComponentsByIdRequest) Execute() (*CatalogComponent, *http.Response, error) {
	return r.ApiService.PutProcurementCatalogByParentIdComponentsByIdExecute(r)
}

/*
PutProcurementCatalogByParentIdComponentsById Put CatalogComponent

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id componentId
 @param parentId catalogId
 @return ApiPutProcurementCatalogByParentIdComponentsByIdRequest
*/
func (a *CatalogComponentsAPIService) PutProcurementCatalogByParentIdComponentsById(ctx context.Context, id int32, parentId int32) ApiPutProcurementCatalogByParentIdComponentsByIdRequest {
	return ApiPutProcurementCatalogByParentIdComponentsByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
		parentId: parentId,
	}
}

// Execute executes the request
//  @return CatalogComponent
func (a *CatalogComponentsAPIService) PutProcurementCatalogByParentIdComponentsByIdExecute(r ApiPutProcurementCatalogByParentIdComponentsByIdRequest) (*CatalogComponent, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CatalogComponent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CatalogComponentsAPIService.PutProcurementCatalogByParentIdComponentsById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/procurement/catalog/{parentId}/components/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"parentId"+"}", url.PathEscape(parameterValueToString(r.parentId, "parentId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.catalogComponent == nil {
		return localVarReturnValue, nil, reportError("catalogComponent is required and must be specified")
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
	localVarPostBody = r.catalogComponent
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
