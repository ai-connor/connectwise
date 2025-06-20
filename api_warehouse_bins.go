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


// WarehouseBinsAPIService WarehouseBinsAPI service
type WarehouseBinsAPIService service

type ApiDeleteProcurementWarehouseBinsByIdRequest struct {
	ctx context.Context
	ApiService *WarehouseBinsAPIService
	id int32
	clientId *string
}

// 
func (r ApiDeleteProcurementWarehouseBinsByIdRequest) ClientId(clientId string) ApiDeleteProcurementWarehouseBinsByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiDeleteProcurementWarehouseBinsByIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteProcurementWarehouseBinsByIdExecute(r)
}

/*
DeleteProcurementWarehouseBinsById Delete WarehouseBin

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id warehouseBinId
 @return ApiDeleteProcurementWarehouseBinsByIdRequest
*/
func (a *WarehouseBinsAPIService) DeleteProcurementWarehouseBinsById(ctx context.Context, id int32) ApiDeleteProcurementWarehouseBinsByIdRequest {
	return ApiDeleteProcurementWarehouseBinsByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *WarehouseBinsAPIService) DeleteProcurementWarehouseBinsByIdExecute(r ApiDeleteProcurementWarehouseBinsByIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WarehouseBinsAPIService.DeleteProcurementWarehouseBinsById")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/procurement/warehouseBins/{id}"
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

type ApiGetProcurementWarehouseBinsRequest struct {
	ctx context.Context
	ApiService *WarehouseBinsAPIService
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
func (r ApiGetProcurementWarehouseBinsRequest) Conditions(conditions string) ApiGetProcurementWarehouseBinsRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetProcurementWarehouseBinsRequest) ChildConditions(childConditions string) ApiGetProcurementWarehouseBinsRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetProcurementWarehouseBinsRequest) CustomFieldConditions(customFieldConditions string) ApiGetProcurementWarehouseBinsRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetProcurementWarehouseBinsRequest) OrderBy(orderBy string) ApiGetProcurementWarehouseBinsRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetProcurementWarehouseBinsRequest) Fields(fields string) ApiGetProcurementWarehouseBinsRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetProcurementWarehouseBinsRequest) Page(page int32) ApiGetProcurementWarehouseBinsRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetProcurementWarehouseBinsRequest) PageSize(pageSize int32) ApiGetProcurementWarehouseBinsRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetProcurementWarehouseBinsRequest) PageId(pageId int32) ApiGetProcurementWarehouseBinsRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetProcurementWarehouseBinsRequest) ClientId(clientId string) ApiGetProcurementWarehouseBinsRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetProcurementWarehouseBinsRequest) Execute() ([]WarehouseBin, *http.Response, error) {
	return r.ApiService.GetProcurementWarehouseBinsExecute(r)
}

/*
GetProcurementWarehouseBins Get List of WarehouseBin

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetProcurementWarehouseBinsRequest
*/
func (a *WarehouseBinsAPIService) GetProcurementWarehouseBins(ctx context.Context) ApiGetProcurementWarehouseBinsRequest {
	return ApiGetProcurementWarehouseBinsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []WarehouseBin
func (a *WarehouseBinsAPIService) GetProcurementWarehouseBinsExecute(r ApiGetProcurementWarehouseBinsRequest) ([]WarehouseBin, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []WarehouseBin
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WarehouseBinsAPIService.GetProcurementWarehouseBins")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/procurement/warehouseBins"

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

type ApiGetProcurementWarehouseBinsByIdRequest struct {
	ctx context.Context
	ApiService *WarehouseBinsAPIService
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
func (r ApiGetProcurementWarehouseBinsByIdRequest) Conditions(conditions string) ApiGetProcurementWarehouseBinsByIdRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetProcurementWarehouseBinsByIdRequest) ChildConditions(childConditions string) ApiGetProcurementWarehouseBinsByIdRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetProcurementWarehouseBinsByIdRequest) CustomFieldConditions(customFieldConditions string) ApiGetProcurementWarehouseBinsByIdRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetProcurementWarehouseBinsByIdRequest) OrderBy(orderBy string) ApiGetProcurementWarehouseBinsByIdRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetProcurementWarehouseBinsByIdRequest) Fields(fields string) ApiGetProcurementWarehouseBinsByIdRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetProcurementWarehouseBinsByIdRequest) Page(page int32) ApiGetProcurementWarehouseBinsByIdRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetProcurementWarehouseBinsByIdRequest) PageSize(pageSize int32) ApiGetProcurementWarehouseBinsByIdRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetProcurementWarehouseBinsByIdRequest) PageId(pageId int32) ApiGetProcurementWarehouseBinsByIdRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetProcurementWarehouseBinsByIdRequest) ClientId(clientId string) ApiGetProcurementWarehouseBinsByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetProcurementWarehouseBinsByIdRequest) Execute() (*WarehouseBin, *http.Response, error) {
	return r.ApiService.GetProcurementWarehouseBinsByIdExecute(r)
}

/*
GetProcurementWarehouseBinsById Get WarehouseBin

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id warehouseBinId
 @return ApiGetProcurementWarehouseBinsByIdRequest
*/
func (a *WarehouseBinsAPIService) GetProcurementWarehouseBinsById(ctx context.Context, id int32) ApiGetProcurementWarehouseBinsByIdRequest {
	return ApiGetProcurementWarehouseBinsByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return WarehouseBin
func (a *WarehouseBinsAPIService) GetProcurementWarehouseBinsByIdExecute(r ApiGetProcurementWarehouseBinsByIdRequest) (*WarehouseBin, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *WarehouseBin
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WarehouseBinsAPIService.GetProcurementWarehouseBinsById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/procurement/warehouseBins/{id}"
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

type ApiGetProcurementWarehouseBinsCountRequest struct {
	ctx context.Context
	ApiService *WarehouseBinsAPIService
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
func (r ApiGetProcurementWarehouseBinsCountRequest) Conditions(conditions string) ApiGetProcurementWarehouseBinsCountRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetProcurementWarehouseBinsCountRequest) ChildConditions(childConditions string) ApiGetProcurementWarehouseBinsCountRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetProcurementWarehouseBinsCountRequest) CustomFieldConditions(customFieldConditions string) ApiGetProcurementWarehouseBinsCountRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetProcurementWarehouseBinsCountRequest) OrderBy(orderBy string) ApiGetProcurementWarehouseBinsCountRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetProcurementWarehouseBinsCountRequest) Fields(fields string) ApiGetProcurementWarehouseBinsCountRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetProcurementWarehouseBinsCountRequest) Page(page int32) ApiGetProcurementWarehouseBinsCountRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetProcurementWarehouseBinsCountRequest) PageSize(pageSize int32) ApiGetProcurementWarehouseBinsCountRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetProcurementWarehouseBinsCountRequest) PageId(pageId int32) ApiGetProcurementWarehouseBinsCountRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetProcurementWarehouseBinsCountRequest) ClientId(clientId string) ApiGetProcurementWarehouseBinsCountRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetProcurementWarehouseBinsCountRequest) Execute() (*Count, *http.Response, error) {
	return r.ApiService.GetProcurementWarehouseBinsCountExecute(r)
}

/*
GetProcurementWarehouseBinsCount Get Count of WarehouseBin

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetProcurementWarehouseBinsCountRequest
*/
func (a *WarehouseBinsAPIService) GetProcurementWarehouseBinsCount(ctx context.Context) ApiGetProcurementWarehouseBinsCountRequest {
	return ApiGetProcurementWarehouseBinsCountRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Count
func (a *WarehouseBinsAPIService) GetProcurementWarehouseBinsCountExecute(r ApiGetProcurementWarehouseBinsCountRequest) (*Count, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Count
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WarehouseBinsAPIService.GetProcurementWarehouseBinsCount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/procurement/warehouseBins/count"

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

type ApiPatchProcurementWarehouseBinsByIdRequest struct {
	ctx context.Context
	ApiService *WarehouseBinsAPIService
	id int32
	patchOperation *[]PatchOperation
	clientId *string
}

// List of PatchOperation
func (r ApiPatchProcurementWarehouseBinsByIdRequest) PatchOperation(patchOperation []PatchOperation) ApiPatchProcurementWarehouseBinsByIdRequest {
	r.patchOperation = &patchOperation
	return r
}

// 
func (r ApiPatchProcurementWarehouseBinsByIdRequest) ClientId(clientId string) ApiPatchProcurementWarehouseBinsByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiPatchProcurementWarehouseBinsByIdRequest) Execute() (*WarehouseBin, *http.Response, error) {
	return r.ApiService.PatchProcurementWarehouseBinsByIdExecute(r)
}

/*
PatchProcurementWarehouseBinsById Patch WarehouseBin

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id warehouseBinId
 @return ApiPatchProcurementWarehouseBinsByIdRequest
*/
func (a *WarehouseBinsAPIService) PatchProcurementWarehouseBinsById(ctx context.Context, id int32) ApiPatchProcurementWarehouseBinsByIdRequest {
	return ApiPatchProcurementWarehouseBinsByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return WarehouseBin
func (a *WarehouseBinsAPIService) PatchProcurementWarehouseBinsByIdExecute(r ApiPatchProcurementWarehouseBinsByIdRequest) (*WarehouseBin, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *WarehouseBin
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WarehouseBinsAPIService.PatchProcurementWarehouseBinsById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/procurement/warehouseBins/{id}"
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

type ApiPostProcurementWarehouseBinsRequest struct {
	ctx context.Context
	ApiService *WarehouseBinsAPIService
	warehouseBin *WarehouseBin
	clientId *string
}

// warehouseBin
func (r ApiPostProcurementWarehouseBinsRequest) WarehouseBin(warehouseBin WarehouseBin) ApiPostProcurementWarehouseBinsRequest {
	r.warehouseBin = &warehouseBin
	return r
}

// 
func (r ApiPostProcurementWarehouseBinsRequest) ClientId(clientId string) ApiPostProcurementWarehouseBinsRequest {
	r.clientId = &clientId
	return r
}

func (r ApiPostProcurementWarehouseBinsRequest) Execute() (*WarehouseBin, *http.Response, error) {
	return r.ApiService.PostProcurementWarehouseBinsExecute(r)
}

/*
PostProcurementWarehouseBins Post WarehouseBin

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPostProcurementWarehouseBinsRequest
*/
func (a *WarehouseBinsAPIService) PostProcurementWarehouseBins(ctx context.Context) ApiPostProcurementWarehouseBinsRequest {
	return ApiPostProcurementWarehouseBinsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return WarehouseBin
func (a *WarehouseBinsAPIService) PostProcurementWarehouseBinsExecute(r ApiPostProcurementWarehouseBinsRequest) (*WarehouseBin, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *WarehouseBin
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WarehouseBinsAPIService.PostProcurementWarehouseBins")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/procurement/warehouseBins"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.warehouseBin == nil {
		return localVarReturnValue, nil, reportError("warehouseBin is required and must be specified")
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
	localVarPostBody = r.warehouseBin
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

type ApiPutProcurementWarehouseBinsByIdRequest struct {
	ctx context.Context
	ApiService *WarehouseBinsAPIService
	id int32
	warehouseBin *WarehouseBin
	clientId *string
}

// warehouseBin
func (r ApiPutProcurementWarehouseBinsByIdRequest) WarehouseBin(warehouseBin WarehouseBin) ApiPutProcurementWarehouseBinsByIdRequest {
	r.warehouseBin = &warehouseBin
	return r
}

// 
func (r ApiPutProcurementWarehouseBinsByIdRequest) ClientId(clientId string) ApiPutProcurementWarehouseBinsByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiPutProcurementWarehouseBinsByIdRequest) Execute() (*WarehouseBin, *http.Response, error) {
	return r.ApiService.PutProcurementWarehouseBinsByIdExecute(r)
}

/*
PutProcurementWarehouseBinsById Put WarehouseBin

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id warehouseBinId
 @return ApiPutProcurementWarehouseBinsByIdRequest
*/
func (a *WarehouseBinsAPIService) PutProcurementWarehouseBinsById(ctx context.Context, id int32) ApiPutProcurementWarehouseBinsByIdRequest {
	return ApiPutProcurementWarehouseBinsByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return WarehouseBin
func (a *WarehouseBinsAPIService) PutProcurementWarehouseBinsByIdExecute(r ApiPutProcurementWarehouseBinsByIdRequest) (*WarehouseBin, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *WarehouseBin
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WarehouseBinsAPIService.PutProcurementWarehouseBinsById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/procurement/warehouseBins/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.warehouseBin == nil {
		return localVarReturnValue, nil, reportError("warehouseBin is required and must be specified")
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
	localVarPostBody = r.warehouseBin
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
