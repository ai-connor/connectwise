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


// OsGradeWeightsAPIService OsGradeWeightsAPI service
type OsGradeWeightsAPIService service

type ApiGetSystemOsgradeweightsRequest struct {
	ctx context.Context
	ApiService *OsGradeWeightsAPIService
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
func (r ApiGetSystemOsgradeweightsRequest) Conditions(conditions string) ApiGetSystemOsgradeweightsRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetSystemOsgradeweightsRequest) ChildConditions(childConditions string) ApiGetSystemOsgradeweightsRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetSystemOsgradeweightsRequest) CustomFieldConditions(customFieldConditions string) ApiGetSystemOsgradeweightsRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetSystemOsgradeweightsRequest) OrderBy(orderBy string) ApiGetSystemOsgradeweightsRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetSystemOsgradeweightsRequest) Fields(fields string) ApiGetSystemOsgradeweightsRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetSystemOsgradeweightsRequest) Page(page int32) ApiGetSystemOsgradeweightsRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetSystemOsgradeweightsRequest) PageSize(pageSize int32) ApiGetSystemOsgradeweightsRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetSystemOsgradeweightsRequest) PageId(pageId int32) ApiGetSystemOsgradeweightsRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetSystemOsgradeweightsRequest) ClientId(clientId string) ApiGetSystemOsgradeweightsRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetSystemOsgradeweightsRequest) Execute() ([]OsGradeWeight, *http.Response, error) {
	return r.ApiService.GetSystemOsgradeweightsExecute(r)
}

/*
GetSystemOsgradeweights Get List of OsGradeWeight

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetSystemOsgradeweightsRequest
*/
func (a *OsGradeWeightsAPIService) GetSystemOsgradeweights(ctx context.Context) ApiGetSystemOsgradeweightsRequest {
	return ApiGetSystemOsgradeweightsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []OsGradeWeight
func (a *OsGradeWeightsAPIService) GetSystemOsgradeweightsExecute(r ApiGetSystemOsgradeweightsRequest) ([]OsGradeWeight, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []OsGradeWeight
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OsGradeWeightsAPIService.GetSystemOsgradeweights")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/osgradeweights"

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

type ApiGetSystemOsgradeweightsByIdRequest struct {
	ctx context.Context
	ApiService *OsGradeWeightsAPIService
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
func (r ApiGetSystemOsgradeweightsByIdRequest) Conditions(conditions string) ApiGetSystemOsgradeweightsByIdRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetSystemOsgradeweightsByIdRequest) ChildConditions(childConditions string) ApiGetSystemOsgradeweightsByIdRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetSystemOsgradeweightsByIdRequest) CustomFieldConditions(customFieldConditions string) ApiGetSystemOsgradeweightsByIdRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetSystemOsgradeweightsByIdRequest) OrderBy(orderBy string) ApiGetSystemOsgradeweightsByIdRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetSystemOsgradeweightsByIdRequest) Fields(fields string) ApiGetSystemOsgradeweightsByIdRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetSystemOsgradeweightsByIdRequest) Page(page int32) ApiGetSystemOsgradeweightsByIdRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetSystemOsgradeweightsByIdRequest) PageSize(pageSize int32) ApiGetSystemOsgradeweightsByIdRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetSystemOsgradeweightsByIdRequest) PageId(pageId int32) ApiGetSystemOsgradeweightsByIdRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetSystemOsgradeweightsByIdRequest) ClientId(clientId string) ApiGetSystemOsgradeweightsByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetSystemOsgradeweightsByIdRequest) Execute() (*OsGradeWeight, *http.Response, error) {
	return r.ApiService.GetSystemOsgradeweightsByIdExecute(r)
}

/*
GetSystemOsgradeweightsById Get OsGradeWeight

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id osgradeweightId
 @return ApiGetSystemOsgradeweightsByIdRequest
*/
func (a *OsGradeWeightsAPIService) GetSystemOsgradeweightsById(ctx context.Context, id int32) ApiGetSystemOsgradeweightsByIdRequest {
	return ApiGetSystemOsgradeweightsByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return OsGradeWeight
func (a *OsGradeWeightsAPIService) GetSystemOsgradeweightsByIdExecute(r ApiGetSystemOsgradeweightsByIdRequest) (*OsGradeWeight, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OsGradeWeight
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OsGradeWeightsAPIService.GetSystemOsgradeweightsById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/osgradeweights/{id}"
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

type ApiGetSystemOsgradeweightsCountRequest struct {
	ctx context.Context
	ApiService *OsGradeWeightsAPIService
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
func (r ApiGetSystemOsgradeweightsCountRequest) Conditions(conditions string) ApiGetSystemOsgradeweightsCountRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetSystemOsgradeweightsCountRequest) ChildConditions(childConditions string) ApiGetSystemOsgradeweightsCountRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetSystemOsgradeweightsCountRequest) CustomFieldConditions(customFieldConditions string) ApiGetSystemOsgradeweightsCountRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetSystemOsgradeweightsCountRequest) OrderBy(orderBy string) ApiGetSystemOsgradeweightsCountRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetSystemOsgradeweightsCountRequest) Fields(fields string) ApiGetSystemOsgradeweightsCountRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetSystemOsgradeweightsCountRequest) Page(page int32) ApiGetSystemOsgradeweightsCountRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetSystemOsgradeweightsCountRequest) PageSize(pageSize int32) ApiGetSystemOsgradeweightsCountRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetSystemOsgradeweightsCountRequest) PageId(pageId int32) ApiGetSystemOsgradeweightsCountRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetSystemOsgradeweightsCountRequest) ClientId(clientId string) ApiGetSystemOsgradeweightsCountRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetSystemOsgradeweightsCountRequest) Execute() (*Count, *http.Response, error) {
	return r.ApiService.GetSystemOsgradeweightsCountExecute(r)
}

/*
GetSystemOsgradeweightsCount Get Count of OsGradeWeight

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetSystemOsgradeweightsCountRequest
*/
func (a *OsGradeWeightsAPIService) GetSystemOsgradeweightsCount(ctx context.Context) ApiGetSystemOsgradeweightsCountRequest {
	return ApiGetSystemOsgradeweightsCountRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Count
func (a *OsGradeWeightsAPIService) GetSystemOsgradeweightsCountExecute(r ApiGetSystemOsgradeweightsCountRequest) (*Count, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Count
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OsGradeWeightsAPIService.GetSystemOsgradeweightsCount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/osgradeweights/count"

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

type ApiPatchSystemOsgradeweightsByIdRequest struct {
	ctx context.Context
	ApiService *OsGradeWeightsAPIService
	id int32
	patchOperation *[]PatchOperation
	clientId *string
}

// List of PatchOperation
func (r ApiPatchSystemOsgradeweightsByIdRequest) PatchOperation(patchOperation []PatchOperation) ApiPatchSystemOsgradeweightsByIdRequest {
	r.patchOperation = &patchOperation
	return r
}

// 
func (r ApiPatchSystemOsgradeweightsByIdRequest) ClientId(clientId string) ApiPatchSystemOsgradeweightsByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiPatchSystemOsgradeweightsByIdRequest) Execute() (*OsGradeWeight, *http.Response, error) {
	return r.ApiService.PatchSystemOsgradeweightsByIdExecute(r)
}

/*
PatchSystemOsgradeweightsById Patch OsGradeWeight

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id osgradeweightId
 @return ApiPatchSystemOsgradeweightsByIdRequest
*/
func (a *OsGradeWeightsAPIService) PatchSystemOsgradeweightsById(ctx context.Context, id int32) ApiPatchSystemOsgradeweightsByIdRequest {
	return ApiPatchSystemOsgradeweightsByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return OsGradeWeight
func (a *OsGradeWeightsAPIService) PatchSystemOsgradeweightsByIdExecute(r ApiPatchSystemOsgradeweightsByIdRequest) (*OsGradeWeight, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OsGradeWeight
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OsGradeWeightsAPIService.PatchSystemOsgradeweightsById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/osgradeweights/{id}"
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

type ApiPutSystemOsgradeweightsByIdRequest struct {
	ctx context.Context
	ApiService *OsGradeWeightsAPIService
	id int32
	osGradeWeight *OsGradeWeight
	clientId *string
}

// osGradeWeight
func (r ApiPutSystemOsgradeweightsByIdRequest) OsGradeWeight(osGradeWeight OsGradeWeight) ApiPutSystemOsgradeweightsByIdRequest {
	r.osGradeWeight = &osGradeWeight
	return r
}

// 
func (r ApiPutSystemOsgradeweightsByIdRequest) ClientId(clientId string) ApiPutSystemOsgradeweightsByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiPutSystemOsgradeweightsByIdRequest) Execute() (*OsGradeWeight, *http.Response, error) {
	return r.ApiService.PutSystemOsgradeweightsByIdExecute(r)
}

/*
PutSystemOsgradeweightsById Put OsGradeWeight

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id osgradeweightId
 @return ApiPutSystemOsgradeweightsByIdRequest
*/
func (a *OsGradeWeightsAPIService) PutSystemOsgradeweightsById(ctx context.Context, id int32) ApiPutSystemOsgradeweightsByIdRequest {
	return ApiPutSystemOsgradeweightsByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return OsGradeWeight
func (a *OsGradeWeightsAPIService) PutSystemOsgradeweightsByIdExecute(r ApiPutSystemOsgradeweightsByIdRequest) (*OsGradeWeight, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OsGradeWeight
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OsGradeWeightsAPIService.PutSystemOsgradeweightsById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/osgradeweights/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.osGradeWeight == nil {
		return localVarReturnValue, nil, reportError("osGradeWeight is required and must be specified")
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
	localVarPostBody = r.osGradeWeight
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
