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


// SalesQuotasAPIService SalesQuotasAPI service
type SalesQuotasAPIService service

type ApiDeleteSalesQuotasByIdRequest struct {
	ctx context.Context
	ApiService *SalesQuotasAPIService
	id int32
	clientId *string
}

// 
func (r ApiDeleteSalesQuotasByIdRequest) ClientId(clientId string) ApiDeleteSalesQuotasByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiDeleteSalesQuotasByIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteSalesQuotasByIdExecute(r)
}

/*
DeleteSalesQuotasById Delete SalesQuota

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id quotaId
 @return ApiDeleteSalesQuotasByIdRequest
*/
func (a *SalesQuotasAPIService) DeleteSalesQuotasById(ctx context.Context, id int32) ApiDeleteSalesQuotasByIdRequest {
	return ApiDeleteSalesQuotasByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *SalesQuotasAPIService) DeleteSalesQuotasByIdExecute(r ApiDeleteSalesQuotasByIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SalesQuotasAPIService.DeleteSalesQuotasById")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sales/quotas/{id}"
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

type ApiGetSalesQuotasRequest struct {
	ctx context.Context
	ApiService *SalesQuotasAPIService
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
func (r ApiGetSalesQuotasRequest) Conditions(conditions string) ApiGetSalesQuotasRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetSalesQuotasRequest) ChildConditions(childConditions string) ApiGetSalesQuotasRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetSalesQuotasRequest) CustomFieldConditions(customFieldConditions string) ApiGetSalesQuotasRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetSalesQuotasRequest) OrderBy(orderBy string) ApiGetSalesQuotasRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetSalesQuotasRequest) Fields(fields string) ApiGetSalesQuotasRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetSalesQuotasRequest) Page(page int32) ApiGetSalesQuotasRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetSalesQuotasRequest) PageSize(pageSize int32) ApiGetSalesQuotasRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetSalesQuotasRequest) PageId(pageId int32) ApiGetSalesQuotasRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetSalesQuotasRequest) ClientId(clientId string) ApiGetSalesQuotasRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetSalesQuotasRequest) Execute() ([]SalesQuota, *http.Response, error) {
	return r.ApiService.GetSalesQuotasExecute(r)
}

/*
GetSalesQuotas Get List of SalesQuota

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetSalesQuotasRequest
*/
func (a *SalesQuotasAPIService) GetSalesQuotas(ctx context.Context) ApiGetSalesQuotasRequest {
	return ApiGetSalesQuotasRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SalesQuota
func (a *SalesQuotasAPIService) GetSalesQuotasExecute(r ApiGetSalesQuotasRequest) ([]SalesQuota, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SalesQuota
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SalesQuotasAPIService.GetSalesQuotas")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sales/quotas"

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

type ApiGetSalesQuotasByIdRequest struct {
	ctx context.Context
	ApiService *SalesQuotasAPIService
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
func (r ApiGetSalesQuotasByIdRequest) Conditions(conditions string) ApiGetSalesQuotasByIdRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetSalesQuotasByIdRequest) ChildConditions(childConditions string) ApiGetSalesQuotasByIdRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetSalesQuotasByIdRequest) CustomFieldConditions(customFieldConditions string) ApiGetSalesQuotasByIdRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetSalesQuotasByIdRequest) OrderBy(orderBy string) ApiGetSalesQuotasByIdRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetSalesQuotasByIdRequest) Fields(fields string) ApiGetSalesQuotasByIdRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetSalesQuotasByIdRequest) Page(page int32) ApiGetSalesQuotasByIdRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetSalesQuotasByIdRequest) PageSize(pageSize int32) ApiGetSalesQuotasByIdRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetSalesQuotasByIdRequest) PageId(pageId int32) ApiGetSalesQuotasByIdRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetSalesQuotasByIdRequest) ClientId(clientId string) ApiGetSalesQuotasByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetSalesQuotasByIdRequest) Execute() (*SalesQuota, *http.Response, error) {
	return r.ApiService.GetSalesQuotasByIdExecute(r)
}

/*
GetSalesQuotasById Get SalesQuota

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id quotaId
 @return ApiGetSalesQuotasByIdRequest
*/
func (a *SalesQuotasAPIService) GetSalesQuotasById(ctx context.Context, id int32) ApiGetSalesQuotasByIdRequest {
	return ApiGetSalesQuotasByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return SalesQuota
func (a *SalesQuotasAPIService) GetSalesQuotasByIdExecute(r ApiGetSalesQuotasByIdRequest) (*SalesQuota, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SalesQuota
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SalesQuotasAPIService.GetSalesQuotasById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sales/quotas/{id}"
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

type ApiGetSalesQuotasCountRequest struct {
	ctx context.Context
	ApiService *SalesQuotasAPIService
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
func (r ApiGetSalesQuotasCountRequest) Conditions(conditions string) ApiGetSalesQuotasCountRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetSalesQuotasCountRequest) ChildConditions(childConditions string) ApiGetSalesQuotasCountRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetSalesQuotasCountRequest) CustomFieldConditions(customFieldConditions string) ApiGetSalesQuotasCountRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetSalesQuotasCountRequest) OrderBy(orderBy string) ApiGetSalesQuotasCountRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetSalesQuotasCountRequest) Fields(fields string) ApiGetSalesQuotasCountRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetSalesQuotasCountRequest) Page(page int32) ApiGetSalesQuotasCountRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetSalesQuotasCountRequest) PageSize(pageSize int32) ApiGetSalesQuotasCountRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetSalesQuotasCountRequest) PageId(pageId int32) ApiGetSalesQuotasCountRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetSalesQuotasCountRequest) ClientId(clientId string) ApiGetSalesQuotasCountRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetSalesQuotasCountRequest) Execute() (*Count, *http.Response, error) {
	return r.ApiService.GetSalesQuotasCountExecute(r)
}

/*
GetSalesQuotasCount Get Count of SalesQuota

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetSalesQuotasCountRequest
*/
func (a *SalesQuotasAPIService) GetSalesQuotasCount(ctx context.Context) ApiGetSalesQuotasCountRequest {
	return ApiGetSalesQuotasCountRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Count
func (a *SalesQuotasAPIService) GetSalesQuotasCountExecute(r ApiGetSalesQuotasCountRequest) (*Count, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Count
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SalesQuotasAPIService.GetSalesQuotasCount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sales/quotas/count"

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

type ApiPatchSalesQuotasByIdRequest struct {
	ctx context.Context
	ApiService *SalesQuotasAPIService
	id int32
	patchOperation *[]PatchOperation
	clientId *string
}

// List of PatchOperation
func (r ApiPatchSalesQuotasByIdRequest) PatchOperation(patchOperation []PatchOperation) ApiPatchSalesQuotasByIdRequest {
	r.patchOperation = &patchOperation
	return r
}

// 
func (r ApiPatchSalesQuotasByIdRequest) ClientId(clientId string) ApiPatchSalesQuotasByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiPatchSalesQuotasByIdRequest) Execute() (*SalesQuota, *http.Response, error) {
	return r.ApiService.PatchSalesQuotasByIdExecute(r)
}

/*
PatchSalesQuotasById Patch SalesQuota

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id quotaId
 @return ApiPatchSalesQuotasByIdRequest
*/
func (a *SalesQuotasAPIService) PatchSalesQuotasById(ctx context.Context, id int32) ApiPatchSalesQuotasByIdRequest {
	return ApiPatchSalesQuotasByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return SalesQuota
func (a *SalesQuotasAPIService) PatchSalesQuotasByIdExecute(r ApiPatchSalesQuotasByIdRequest) (*SalesQuota, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SalesQuota
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SalesQuotasAPIService.PatchSalesQuotasById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sales/quotas/{id}"
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

type ApiPostSalesQuotasRequest struct {
	ctx context.Context
	ApiService *SalesQuotasAPIService
	salesQuota *SalesQuota
	clientId *string
}

// salesQuota
func (r ApiPostSalesQuotasRequest) SalesQuota(salesQuota SalesQuota) ApiPostSalesQuotasRequest {
	r.salesQuota = &salesQuota
	return r
}

// 
func (r ApiPostSalesQuotasRequest) ClientId(clientId string) ApiPostSalesQuotasRequest {
	r.clientId = &clientId
	return r
}

func (r ApiPostSalesQuotasRequest) Execute() (*SalesQuota, *http.Response, error) {
	return r.ApiService.PostSalesQuotasExecute(r)
}

/*
PostSalesQuotas Post SalesQuota

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPostSalesQuotasRequest
*/
func (a *SalesQuotasAPIService) PostSalesQuotas(ctx context.Context) ApiPostSalesQuotasRequest {
	return ApiPostSalesQuotasRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SalesQuota
func (a *SalesQuotasAPIService) PostSalesQuotasExecute(r ApiPostSalesQuotasRequest) (*SalesQuota, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SalesQuota
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SalesQuotasAPIService.PostSalesQuotas")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sales/quotas"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.salesQuota == nil {
		return localVarReturnValue, nil, reportError("salesQuota is required and must be specified")
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
	localVarPostBody = r.salesQuota
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

type ApiPutSalesQuotasByIdRequest struct {
	ctx context.Context
	ApiService *SalesQuotasAPIService
	id int32
	salesQuota *SalesQuota
	clientId *string
}

// salesQuota
func (r ApiPutSalesQuotasByIdRequest) SalesQuota(salesQuota SalesQuota) ApiPutSalesQuotasByIdRequest {
	r.salesQuota = &salesQuota
	return r
}

// 
func (r ApiPutSalesQuotasByIdRequest) ClientId(clientId string) ApiPutSalesQuotasByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiPutSalesQuotasByIdRequest) Execute() (*SalesQuota, *http.Response, error) {
	return r.ApiService.PutSalesQuotasByIdExecute(r)
}

/*
PutSalesQuotasById Put SalesQuota

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id quotaId
 @return ApiPutSalesQuotasByIdRequest
*/
func (a *SalesQuotasAPIService) PutSalesQuotasById(ctx context.Context, id int32) ApiPutSalesQuotasByIdRequest {
	return ApiPutSalesQuotasByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return SalesQuota
func (a *SalesQuotasAPIService) PutSalesQuotasByIdExecute(r ApiPutSalesQuotasByIdRequest) (*SalesQuota, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SalesQuota
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SalesQuotasAPIService.PutSalesQuotasById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sales/quotas/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.salesQuota == nil {
		return localVarReturnValue, nil, reportError("salesQuota is required and must be specified")
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
	localVarPostBody = r.salesQuota
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
