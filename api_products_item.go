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


// ProductsItemAPIService ProductsItemAPI service
type ProductsItemAPIService service

type ApiDeleteProcurementProductsByIdRequest struct {
	ctx context.Context
	ApiService *ProductsItemAPIService
	id int32
	clientId *string
}

// 
func (r ApiDeleteProcurementProductsByIdRequest) ClientId(clientId string) ApiDeleteProcurementProductsByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiDeleteProcurementProductsByIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteProcurementProductsByIdExecute(r)
}

/*
DeleteProcurementProductsById Delete ProductItem

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id productId
 @return ApiDeleteProcurementProductsByIdRequest
*/
func (a *ProductsItemAPIService) DeleteProcurementProductsById(ctx context.Context, id int32) ApiDeleteProcurementProductsByIdRequest {
	return ApiDeleteProcurementProductsByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *ProductsItemAPIService) DeleteProcurementProductsByIdExecute(r ApiDeleteProcurementProductsByIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductsItemAPIService.DeleteProcurementProductsById")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/procurement/products/{id}"
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

type ApiGetProcurementProductsRequest struct {
	ctx context.Context
	ApiService *ProductsItemAPIService
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
func (r ApiGetProcurementProductsRequest) Conditions(conditions string) ApiGetProcurementProductsRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetProcurementProductsRequest) ChildConditions(childConditions string) ApiGetProcurementProductsRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetProcurementProductsRequest) CustomFieldConditions(customFieldConditions string) ApiGetProcurementProductsRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetProcurementProductsRequest) OrderBy(orderBy string) ApiGetProcurementProductsRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetProcurementProductsRequest) Fields(fields string) ApiGetProcurementProductsRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetProcurementProductsRequest) Page(page int32) ApiGetProcurementProductsRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetProcurementProductsRequest) PageSize(pageSize int32) ApiGetProcurementProductsRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetProcurementProductsRequest) PageId(pageId int32) ApiGetProcurementProductsRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetProcurementProductsRequest) ClientId(clientId string) ApiGetProcurementProductsRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetProcurementProductsRequest) Execute() ([]ProductItem, *http.Response, error) {
	return r.ApiService.GetProcurementProductsExecute(r)
}

/*
GetProcurementProducts Get List of ProductItem

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetProcurementProductsRequest
*/
func (a *ProductsItemAPIService) GetProcurementProducts(ctx context.Context) ApiGetProcurementProductsRequest {
	return ApiGetProcurementProductsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []ProductItem
func (a *ProductsItemAPIService) GetProcurementProductsExecute(r ApiGetProcurementProductsRequest) ([]ProductItem, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ProductItem
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductsItemAPIService.GetProcurementProducts")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/procurement/products"

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

type ApiGetProcurementProductsByIdRequest struct {
	ctx context.Context
	ApiService *ProductsItemAPIService
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
func (r ApiGetProcurementProductsByIdRequest) Conditions(conditions string) ApiGetProcurementProductsByIdRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetProcurementProductsByIdRequest) ChildConditions(childConditions string) ApiGetProcurementProductsByIdRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetProcurementProductsByIdRequest) CustomFieldConditions(customFieldConditions string) ApiGetProcurementProductsByIdRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetProcurementProductsByIdRequest) OrderBy(orderBy string) ApiGetProcurementProductsByIdRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetProcurementProductsByIdRequest) Fields(fields string) ApiGetProcurementProductsByIdRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetProcurementProductsByIdRequest) Page(page int32) ApiGetProcurementProductsByIdRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetProcurementProductsByIdRequest) PageSize(pageSize int32) ApiGetProcurementProductsByIdRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetProcurementProductsByIdRequest) PageId(pageId int32) ApiGetProcurementProductsByIdRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetProcurementProductsByIdRequest) ClientId(clientId string) ApiGetProcurementProductsByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetProcurementProductsByIdRequest) Execute() (*ProductItem, *http.Response, error) {
	return r.ApiService.GetProcurementProductsByIdExecute(r)
}

/*
GetProcurementProductsById Get ProductItem

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id productId
 @return ApiGetProcurementProductsByIdRequest
*/
func (a *ProductsItemAPIService) GetProcurementProductsById(ctx context.Context, id int32) ApiGetProcurementProductsByIdRequest {
	return ApiGetProcurementProductsByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ProductItem
func (a *ProductsItemAPIService) GetProcurementProductsByIdExecute(r ApiGetProcurementProductsByIdRequest) (*ProductItem, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ProductItem
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductsItemAPIService.GetProcurementProductsById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/procurement/products/{id}"
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

type ApiGetProcurementProductsCountRequest struct {
	ctx context.Context
	ApiService *ProductsItemAPIService
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
func (r ApiGetProcurementProductsCountRequest) Conditions(conditions string) ApiGetProcurementProductsCountRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetProcurementProductsCountRequest) ChildConditions(childConditions string) ApiGetProcurementProductsCountRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetProcurementProductsCountRequest) CustomFieldConditions(customFieldConditions string) ApiGetProcurementProductsCountRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetProcurementProductsCountRequest) OrderBy(orderBy string) ApiGetProcurementProductsCountRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetProcurementProductsCountRequest) Fields(fields string) ApiGetProcurementProductsCountRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetProcurementProductsCountRequest) Page(page int32) ApiGetProcurementProductsCountRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetProcurementProductsCountRequest) PageSize(pageSize int32) ApiGetProcurementProductsCountRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetProcurementProductsCountRequest) PageId(pageId int32) ApiGetProcurementProductsCountRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetProcurementProductsCountRequest) ClientId(clientId string) ApiGetProcurementProductsCountRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetProcurementProductsCountRequest) Execute() (*Count, *http.Response, error) {
	return r.ApiService.GetProcurementProductsCountExecute(r)
}

/*
GetProcurementProductsCount Get Count of ProductItem

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetProcurementProductsCountRequest
*/
func (a *ProductsItemAPIService) GetProcurementProductsCount(ctx context.Context) ApiGetProcurementProductsCountRequest {
	return ApiGetProcurementProductsCountRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Count
func (a *ProductsItemAPIService) GetProcurementProductsCountExecute(r ApiGetProcurementProductsCountRequest) (*Count, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Count
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductsItemAPIService.GetProcurementProductsCount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/procurement/products/count"

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

type ApiPatchProcurementProductsByIdRequest struct {
	ctx context.Context
	ApiService *ProductsItemAPIService
	id int32
	patchOperation *[]PatchOperation
	clientId *string
}

// List of PatchOperation
func (r ApiPatchProcurementProductsByIdRequest) PatchOperation(patchOperation []PatchOperation) ApiPatchProcurementProductsByIdRequest {
	r.patchOperation = &patchOperation
	return r
}

// 
func (r ApiPatchProcurementProductsByIdRequest) ClientId(clientId string) ApiPatchProcurementProductsByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiPatchProcurementProductsByIdRequest) Execute() (*ProductItem, *http.Response, error) {
	return r.ApiService.PatchProcurementProductsByIdExecute(r)
}

/*
PatchProcurementProductsById Patch ProductItem

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id productId
 @return ApiPatchProcurementProductsByIdRequest
*/
func (a *ProductsItemAPIService) PatchProcurementProductsById(ctx context.Context, id int32) ApiPatchProcurementProductsByIdRequest {
	return ApiPatchProcurementProductsByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ProductItem
func (a *ProductsItemAPIService) PatchProcurementProductsByIdExecute(r ApiPatchProcurementProductsByIdRequest) (*ProductItem, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ProductItem
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductsItemAPIService.PatchProcurementProductsById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/procurement/products/{id}"
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

type ApiPostProcurementProductsRequest struct {
	ctx context.Context
	ApiService *ProductsItemAPIService
	productItem *ProductItem
	clientId *string
}

// productItem
func (r ApiPostProcurementProductsRequest) ProductItem(productItem ProductItem) ApiPostProcurementProductsRequest {
	r.productItem = &productItem
	return r
}

// 
func (r ApiPostProcurementProductsRequest) ClientId(clientId string) ApiPostProcurementProductsRequest {
	r.clientId = &clientId
	return r
}

func (r ApiPostProcurementProductsRequest) Execute() (*ProductItem, *http.Response, error) {
	return r.ApiService.PostProcurementProductsExecute(r)
}

/*
PostProcurementProducts Post ProductItem

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPostProcurementProductsRequest
*/
func (a *ProductsItemAPIService) PostProcurementProducts(ctx context.Context) ApiPostProcurementProductsRequest {
	return ApiPostProcurementProductsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ProductItem
func (a *ProductsItemAPIService) PostProcurementProductsExecute(r ApiPostProcurementProductsRequest) (*ProductItem, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ProductItem
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductsItemAPIService.PostProcurementProducts")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/procurement/products"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.productItem == nil {
		return localVarReturnValue, nil, reportError("productItem is required and must be specified")
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
	localVarPostBody = r.productItem
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

type ApiPostProcurementProductsByIdDetachRequest struct {
	ctx context.Context
	ApiService *ProductsItemAPIService
	id int32
	productDetach *ProductDetach
	clientId *string
}

// detach
func (r ApiPostProcurementProductsByIdDetachRequest) ProductDetach(productDetach ProductDetach) ApiPostProcurementProductsByIdDetachRequest {
	r.productDetach = &productDetach
	return r
}

// 
func (r ApiPostProcurementProductsByIdDetachRequest) ClientId(clientId string) ApiPostProcurementProductsByIdDetachRequest {
	r.clientId = &clientId
	return r
}

func (r ApiPostProcurementProductsByIdDetachRequest) Execute() (*ProductDetach, *http.Response, error) {
	return r.ApiService.PostProcurementProductsByIdDetachExecute(r)
}

/*
PostProcurementProductsByIdDetach Post ProductDetach

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id productId
 @return ApiPostProcurementProductsByIdDetachRequest
*/
func (a *ProductsItemAPIService) PostProcurementProductsByIdDetach(ctx context.Context, id int32) ApiPostProcurementProductsByIdDetachRequest {
	return ApiPostProcurementProductsByIdDetachRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ProductDetach
func (a *ProductsItemAPIService) PostProcurementProductsByIdDetachExecute(r ApiPostProcurementProductsByIdDetachRequest) (*ProductDetach, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ProductDetach
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductsItemAPIService.PostProcurementProductsByIdDetach")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/procurement/products/{id}/detach"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.productDetach == nil {
		return localVarReturnValue, nil, reportError("productDetach is required and must be specified")
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
	localVarPostBody = r.productDetach
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

type ApiPutProcurementProductsByIdRequest struct {
	ctx context.Context
	ApiService *ProductsItemAPIService
	id int32
	productItem *ProductItem
	clientId *string
}

// productItem
func (r ApiPutProcurementProductsByIdRequest) ProductItem(productItem ProductItem) ApiPutProcurementProductsByIdRequest {
	r.productItem = &productItem
	return r
}

// 
func (r ApiPutProcurementProductsByIdRequest) ClientId(clientId string) ApiPutProcurementProductsByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiPutProcurementProductsByIdRequest) Execute() (*ProductItem, *http.Response, error) {
	return r.ApiService.PutProcurementProductsByIdExecute(r)
}

/*
PutProcurementProductsById Put ProductItem

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id productId
 @return ApiPutProcurementProductsByIdRequest
*/
func (a *ProductsItemAPIService) PutProcurementProductsById(ctx context.Context, id int32) ApiPutProcurementProductsByIdRequest {
	return ApiPutProcurementProductsByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ProductItem
func (a *ProductsItemAPIService) PutProcurementProductsByIdExecute(r ApiPutProcurementProductsByIdRequest) (*ProductItem, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ProductItem
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductsItemAPIService.PutProcurementProductsById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/procurement/products/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.productItem == nil {
		return localVarReturnValue, nil, reportError("productItem is required and must be specified")
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
	localVarPostBody = r.productItem
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
