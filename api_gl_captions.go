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


// GLCaptionsAPIService GLCaptionsAPI service
type GLCaptionsAPIService service

type ApiGetFinanceGlCaptionsRequest struct {
	ctx context.Context
	ApiService *GLCaptionsAPIService
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
func (r ApiGetFinanceGlCaptionsRequest) Conditions(conditions string) ApiGetFinanceGlCaptionsRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetFinanceGlCaptionsRequest) ChildConditions(childConditions string) ApiGetFinanceGlCaptionsRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetFinanceGlCaptionsRequest) CustomFieldConditions(customFieldConditions string) ApiGetFinanceGlCaptionsRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetFinanceGlCaptionsRequest) OrderBy(orderBy string) ApiGetFinanceGlCaptionsRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetFinanceGlCaptionsRequest) Fields(fields string) ApiGetFinanceGlCaptionsRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetFinanceGlCaptionsRequest) Page(page int32) ApiGetFinanceGlCaptionsRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetFinanceGlCaptionsRequest) PageSize(pageSize int32) ApiGetFinanceGlCaptionsRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetFinanceGlCaptionsRequest) PageId(pageId int32) ApiGetFinanceGlCaptionsRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetFinanceGlCaptionsRequest) ClientId(clientId string) ApiGetFinanceGlCaptionsRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetFinanceGlCaptionsRequest) Execute() ([]GLCaption, *http.Response, error) {
	return r.ApiService.GetFinanceGlCaptionsExecute(r)
}

/*
GetFinanceGlCaptions Get List of GLCaption

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetFinanceGlCaptionsRequest
*/
func (a *GLCaptionsAPIService) GetFinanceGlCaptions(ctx context.Context) ApiGetFinanceGlCaptionsRequest {
	return ApiGetFinanceGlCaptionsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []GLCaption
func (a *GLCaptionsAPIService) GetFinanceGlCaptionsExecute(r ApiGetFinanceGlCaptionsRequest) ([]GLCaption, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GLCaption
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GLCaptionsAPIService.GetFinanceGlCaptions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/finance/glCaptions"

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

type ApiGetFinanceGlCaptionsByIdRequest struct {
	ctx context.Context
	ApiService *GLCaptionsAPIService
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
func (r ApiGetFinanceGlCaptionsByIdRequest) Conditions(conditions string) ApiGetFinanceGlCaptionsByIdRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetFinanceGlCaptionsByIdRequest) ChildConditions(childConditions string) ApiGetFinanceGlCaptionsByIdRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetFinanceGlCaptionsByIdRequest) CustomFieldConditions(customFieldConditions string) ApiGetFinanceGlCaptionsByIdRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetFinanceGlCaptionsByIdRequest) OrderBy(orderBy string) ApiGetFinanceGlCaptionsByIdRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetFinanceGlCaptionsByIdRequest) Fields(fields string) ApiGetFinanceGlCaptionsByIdRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetFinanceGlCaptionsByIdRequest) Page(page int32) ApiGetFinanceGlCaptionsByIdRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetFinanceGlCaptionsByIdRequest) PageSize(pageSize int32) ApiGetFinanceGlCaptionsByIdRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetFinanceGlCaptionsByIdRequest) PageId(pageId int32) ApiGetFinanceGlCaptionsByIdRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetFinanceGlCaptionsByIdRequest) ClientId(clientId string) ApiGetFinanceGlCaptionsByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetFinanceGlCaptionsByIdRequest) Execute() (*GLCaption, *http.Response, error) {
	return r.ApiService.GetFinanceGlCaptionsByIdExecute(r)
}

/*
GetFinanceGlCaptionsById Get GLCaption

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id glCaptionId
 @return ApiGetFinanceGlCaptionsByIdRequest
*/
func (a *GLCaptionsAPIService) GetFinanceGlCaptionsById(ctx context.Context, id int32) ApiGetFinanceGlCaptionsByIdRequest {
	return ApiGetFinanceGlCaptionsByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return GLCaption
func (a *GLCaptionsAPIService) GetFinanceGlCaptionsByIdExecute(r ApiGetFinanceGlCaptionsByIdRequest) (*GLCaption, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GLCaption
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GLCaptionsAPIService.GetFinanceGlCaptionsById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/finance/glCaptions/{id}"
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

type ApiGetFinanceGlCaptionsCountRequest struct {
	ctx context.Context
	ApiService *GLCaptionsAPIService
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
func (r ApiGetFinanceGlCaptionsCountRequest) Conditions(conditions string) ApiGetFinanceGlCaptionsCountRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetFinanceGlCaptionsCountRequest) ChildConditions(childConditions string) ApiGetFinanceGlCaptionsCountRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetFinanceGlCaptionsCountRequest) CustomFieldConditions(customFieldConditions string) ApiGetFinanceGlCaptionsCountRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetFinanceGlCaptionsCountRequest) OrderBy(orderBy string) ApiGetFinanceGlCaptionsCountRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetFinanceGlCaptionsCountRequest) Fields(fields string) ApiGetFinanceGlCaptionsCountRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetFinanceGlCaptionsCountRequest) Page(page int32) ApiGetFinanceGlCaptionsCountRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetFinanceGlCaptionsCountRequest) PageSize(pageSize int32) ApiGetFinanceGlCaptionsCountRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetFinanceGlCaptionsCountRequest) PageId(pageId int32) ApiGetFinanceGlCaptionsCountRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetFinanceGlCaptionsCountRequest) ClientId(clientId string) ApiGetFinanceGlCaptionsCountRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetFinanceGlCaptionsCountRequest) Execute() (*Count, *http.Response, error) {
	return r.ApiService.GetFinanceGlCaptionsCountExecute(r)
}

/*
GetFinanceGlCaptionsCount Get Count of GLCaption

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetFinanceGlCaptionsCountRequest
*/
func (a *GLCaptionsAPIService) GetFinanceGlCaptionsCount(ctx context.Context) ApiGetFinanceGlCaptionsCountRequest {
	return ApiGetFinanceGlCaptionsCountRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Count
func (a *GLCaptionsAPIService) GetFinanceGlCaptionsCountExecute(r ApiGetFinanceGlCaptionsCountRequest) (*Count, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Count
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GLCaptionsAPIService.GetFinanceGlCaptionsCount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/finance/glCaptions/count"

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

type ApiPatchFinanceGlCaptionsByIdRequest struct {
	ctx context.Context
	ApiService *GLCaptionsAPIService
	id int32
	patchOperation *[]PatchOperation
	clientId *string
}

// List of PatchOperation
func (r ApiPatchFinanceGlCaptionsByIdRequest) PatchOperation(patchOperation []PatchOperation) ApiPatchFinanceGlCaptionsByIdRequest {
	r.patchOperation = &patchOperation
	return r
}

// 
func (r ApiPatchFinanceGlCaptionsByIdRequest) ClientId(clientId string) ApiPatchFinanceGlCaptionsByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiPatchFinanceGlCaptionsByIdRequest) Execute() (*GLCaption, *http.Response, error) {
	return r.ApiService.PatchFinanceGlCaptionsByIdExecute(r)
}

/*
PatchFinanceGlCaptionsById Patch GLCaption

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id glCaptionId
 @return ApiPatchFinanceGlCaptionsByIdRequest
*/
func (a *GLCaptionsAPIService) PatchFinanceGlCaptionsById(ctx context.Context, id int32) ApiPatchFinanceGlCaptionsByIdRequest {
	return ApiPatchFinanceGlCaptionsByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return GLCaption
func (a *GLCaptionsAPIService) PatchFinanceGlCaptionsByIdExecute(r ApiPatchFinanceGlCaptionsByIdRequest) (*GLCaption, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GLCaption
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GLCaptionsAPIService.PatchFinanceGlCaptionsById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/finance/glCaptions/{id}"
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

type ApiPutFinanceGlCaptionsByIdRequest struct {
	ctx context.Context
	ApiService *GLCaptionsAPIService
	id int32
	gLCaption *GLCaption
	clientId *string
}

// glCaption
func (r ApiPutFinanceGlCaptionsByIdRequest) GLCaption(gLCaption GLCaption) ApiPutFinanceGlCaptionsByIdRequest {
	r.gLCaption = &gLCaption
	return r
}

// 
func (r ApiPutFinanceGlCaptionsByIdRequest) ClientId(clientId string) ApiPutFinanceGlCaptionsByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiPutFinanceGlCaptionsByIdRequest) Execute() (*GLCaption, *http.Response, error) {
	return r.ApiService.PutFinanceGlCaptionsByIdExecute(r)
}

/*
PutFinanceGlCaptionsById Put GLCaption

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id glCaptionId
 @return ApiPutFinanceGlCaptionsByIdRequest
*/
func (a *GLCaptionsAPIService) PutFinanceGlCaptionsById(ctx context.Context, id int32) ApiPutFinanceGlCaptionsByIdRequest {
	return ApiPutFinanceGlCaptionsByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return GLCaption
func (a *GLCaptionsAPIService) PutFinanceGlCaptionsByIdExecute(r ApiPutFinanceGlCaptionsByIdRequest) (*GLCaption, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GLCaption
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GLCaptionsAPIService.PutFinanceGlCaptionsById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/finance/glCaptions/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.gLCaption == nil {
		return localVarReturnValue, nil, reportError("gLCaption is required and must be specified")
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
	localVarPostBody = r.gLCaption
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
