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


// CatalogInventoriesAPIService CatalogInventoriesAPI service
type CatalogInventoriesAPIService service

type ApiGetProcurementCatalogByParentIdInventoryRequest struct {
	ctx context.Context
	ApiService *CatalogInventoriesAPIService
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
func (r ApiGetProcurementCatalogByParentIdInventoryRequest) Conditions(conditions string) ApiGetProcurementCatalogByParentIdInventoryRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetProcurementCatalogByParentIdInventoryRequest) ChildConditions(childConditions string) ApiGetProcurementCatalogByParentIdInventoryRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetProcurementCatalogByParentIdInventoryRequest) CustomFieldConditions(customFieldConditions string) ApiGetProcurementCatalogByParentIdInventoryRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetProcurementCatalogByParentIdInventoryRequest) OrderBy(orderBy string) ApiGetProcurementCatalogByParentIdInventoryRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetProcurementCatalogByParentIdInventoryRequest) Fields(fields string) ApiGetProcurementCatalogByParentIdInventoryRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetProcurementCatalogByParentIdInventoryRequest) Page(page int32) ApiGetProcurementCatalogByParentIdInventoryRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetProcurementCatalogByParentIdInventoryRequest) PageSize(pageSize int32) ApiGetProcurementCatalogByParentIdInventoryRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetProcurementCatalogByParentIdInventoryRequest) PageId(pageId int32) ApiGetProcurementCatalogByParentIdInventoryRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetProcurementCatalogByParentIdInventoryRequest) ClientId(clientId string) ApiGetProcurementCatalogByParentIdInventoryRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetProcurementCatalogByParentIdInventoryRequest) Execute() ([]CatalogInventory, *http.Response, error) {
	return r.ApiService.GetProcurementCatalogByParentIdInventoryExecute(r)
}

/*
GetProcurementCatalogByParentIdInventory Get List of CatalogInventory

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param parentId catalogId
 @return ApiGetProcurementCatalogByParentIdInventoryRequest
*/
func (a *CatalogInventoriesAPIService) GetProcurementCatalogByParentIdInventory(ctx context.Context, parentId int32) ApiGetProcurementCatalogByParentIdInventoryRequest {
	return ApiGetProcurementCatalogByParentIdInventoryRequest{
		ApiService: a,
		ctx: ctx,
		parentId: parentId,
	}
}

// Execute executes the request
//  @return []CatalogInventory
func (a *CatalogInventoriesAPIService) GetProcurementCatalogByParentIdInventoryExecute(r ApiGetProcurementCatalogByParentIdInventoryRequest) ([]CatalogInventory, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []CatalogInventory
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CatalogInventoriesAPIService.GetProcurementCatalogByParentIdInventory")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/procurement/catalog/{parentId}/inventory"
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

type ApiGetProcurementCatalogByParentIdInventoryByIdRequest struct {
	ctx context.Context
	ApiService *CatalogInventoriesAPIService
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
func (r ApiGetProcurementCatalogByParentIdInventoryByIdRequest) Conditions(conditions string) ApiGetProcurementCatalogByParentIdInventoryByIdRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetProcurementCatalogByParentIdInventoryByIdRequest) ChildConditions(childConditions string) ApiGetProcurementCatalogByParentIdInventoryByIdRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetProcurementCatalogByParentIdInventoryByIdRequest) CustomFieldConditions(customFieldConditions string) ApiGetProcurementCatalogByParentIdInventoryByIdRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetProcurementCatalogByParentIdInventoryByIdRequest) OrderBy(orderBy string) ApiGetProcurementCatalogByParentIdInventoryByIdRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetProcurementCatalogByParentIdInventoryByIdRequest) Fields(fields string) ApiGetProcurementCatalogByParentIdInventoryByIdRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetProcurementCatalogByParentIdInventoryByIdRequest) Page(page int32) ApiGetProcurementCatalogByParentIdInventoryByIdRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetProcurementCatalogByParentIdInventoryByIdRequest) PageSize(pageSize int32) ApiGetProcurementCatalogByParentIdInventoryByIdRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetProcurementCatalogByParentIdInventoryByIdRequest) PageId(pageId int32) ApiGetProcurementCatalogByParentIdInventoryByIdRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetProcurementCatalogByParentIdInventoryByIdRequest) ClientId(clientId string) ApiGetProcurementCatalogByParentIdInventoryByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetProcurementCatalogByParentIdInventoryByIdRequest) Execute() (*CatalogInventory, *http.Response, error) {
	return r.ApiService.GetProcurementCatalogByParentIdInventoryByIdExecute(r)
}

/*
GetProcurementCatalogByParentIdInventoryById Get CatalogInventory

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id inventoryId
 @param parentId catalogId
 @return ApiGetProcurementCatalogByParentIdInventoryByIdRequest
*/
func (a *CatalogInventoriesAPIService) GetProcurementCatalogByParentIdInventoryById(ctx context.Context, id int32, parentId int32) ApiGetProcurementCatalogByParentIdInventoryByIdRequest {
	return ApiGetProcurementCatalogByParentIdInventoryByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
		parentId: parentId,
	}
}

// Execute executes the request
//  @return CatalogInventory
func (a *CatalogInventoriesAPIService) GetProcurementCatalogByParentIdInventoryByIdExecute(r ApiGetProcurementCatalogByParentIdInventoryByIdRequest) (*CatalogInventory, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CatalogInventory
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CatalogInventoriesAPIService.GetProcurementCatalogByParentIdInventoryById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/procurement/catalog/{parentId}/inventory/{id}"
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

type ApiGetProcurementCatalogByParentIdInventoryCountRequest struct {
	ctx context.Context
	ApiService *CatalogInventoriesAPIService
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
func (r ApiGetProcurementCatalogByParentIdInventoryCountRequest) Conditions(conditions string) ApiGetProcurementCatalogByParentIdInventoryCountRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetProcurementCatalogByParentIdInventoryCountRequest) ChildConditions(childConditions string) ApiGetProcurementCatalogByParentIdInventoryCountRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetProcurementCatalogByParentIdInventoryCountRequest) CustomFieldConditions(customFieldConditions string) ApiGetProcurementCatalogByParentIdInventoryCountRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetProcurementCatalogByParentIdInventoryCountRequest) OrderBy(orderBy string) ApiGetProcurementCatalogByParentIdInventoryCountRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetProcurementCatalogByParentIdInventoryCountRequest) Fields(fields string) ApiGetProcurementCatalogByParentIdInventoryCountRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetProcurementCatalogByParentIdInventoryCountRequest) Page(page int32) ApiGetProcurementCatalogByParentIdInventoryCountRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetProcurementCatalogByParentIdInventoryCountRequest) PageSize(pageSize int32) ApiGetProcurementCatalogByParentIdInventoryCountRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetProcurementCatalogByParentIdInventoryCountRequest) PageId(pageId int32) ApiGetProcurementCatalogByParentIdInventoryCountRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetProcurementCatalogByParentIdInventoryCountRequest) ClientId(clientId string) ApiGetProcurementCatalogByParentIdInventoryCountRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetProcurementCatalogByParentIdInventoryCountRequest) Execute() (*Count, *http.Response, error) {
	return r.ApiService.GetProcurementCatalogByParentIdInventoryCountExecute(r)
}

/*
GetProcurementCatalogByParentIdInventoryCount Get Count of CatalogInventory

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param parentId catalogId
 @return ApiGetProcurementCatalogByParentIdInventoryCountRequest
*/
func (a *CatalogInventoriesAPIService) GetProcurementCatalogByParentIdInventoryCount(ctx context.Context, parentId int32) ApiGetProcurementCatalogByParentIdInventoryCountRequest {
	return ApiGetProcurementCatalogByParentIdInventoryCountRequest{
		ApiService: a,
		ctx: ctx,
		parentId: parentId,
	}
}

// Execute executes the request
//  @return Count
func (a *CatalogInventoriesAPIService) GetProcurementCatalogByParentIdInventoryCountExecute(r ApiGetProcurementCatalogByParentIdInventoryCountRequest) (*Count, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Count
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CatalogInventoriesAPIService.GetProcurementCatalogByParentIdInventoryCount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/procurement/catalog/{parentId}/inventory/count"
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
