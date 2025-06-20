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


// ActivityStopwatchesAPIService ActivityStopwatchesAPI service
type ActivityStopwatchesAPIService service

type ApiDeleteTimeActivitystopwatchesByIdRequest struct {
	ctx context.Context
	ApiService *ActivityStopwatchesAPIService
	id int32
	clientId *string
}

// 
func (r ApiDeleteTimeActivitystopwatchesByIdRequest) ClientId(clientId string) ApiDeleteTimeActivitystopwatchesByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiDeleteTimeActivitystopwatchesByIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteTimeActivitystopwatchesByIdExecute(r)
}

/*
DeleteTimeActivitystopwatchesById Delete ActivityStopwatch

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id activitystopwatcheId
 @return ApiDeleteTimeActivitystopwatchesByIdRequest
*/
func (a *ActivityStopwatchesAPIService) DeleteTimeActivitystopwatchesById(ctx context.Context, id int32) ApiDeleteTimeActivitystopwatchesByIdRequest {
	return ApiDeleteTimeActivitystopwatchesByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *ActivityStopwatchesAPIService) DeleteTimeActivitystopwatchesByIdExecute(r ApiDeleteTimeActivitystopwatchesByIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ActivityStopwatchesAPIService.DeleteTimeActivitystopwatchesById")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/time/activitystopwatches/{id}"
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

type ApiGetTimeActivitystopwatchesRequest struct {
	ctx context.Context
	ApiService *ActivityStopwatchesAPIService
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
func (r ApiGetTimeActivitystopwatchesRequest) Conditions(conditions string) ApiGetTimeActivitystopwatchesRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetTimeActivitystopwatchesRequest) ChildConditions(childConditions string) ApiGetTimeActivitystopwatchesRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetTimeActivitystopwatchesRequest) CustomFieldConditions(customFieldConditions string) ApiGetTimeActivitystopwatchesRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetTimeActivitystopwatchesRequest) OrderBy(orderBy string) ApiGetTimeActivitystopwatchesRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetTimeActivitystopwatchesRequest) Fields(fields string) ApiGetTimeActivitystopwatchesRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetTimeActivitystopwatchesRequest) Page(page int32) ApiGetTimeActivitystopwatchesRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetTimeActivitystopwatchesRequest) PageSize(pageSize int32) ApiGetTimeActivitystopwatchesRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetTimeActivitystopwatchesRequest) PageId(pageId int32) ApiGetTimeActivitystopwatchesRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetTimeActivitystopwatchesRequest) ClientId(clientId string) ApiGetTimeActivitystopwatchesRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetTimeActivitystopwatchesRequest) Execute() ([]ActivityStopwatch, *http.Response, error) {
	return r.ApiService.GetTimeActivitystopwatchesExecute(r)
}

/*
GetTimeActivitystopwatches Get List of ActivityStopwatch

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetTimeActivitystopwatchesRequest
*/
func (a *ActivityStopwatchesAPIService) GetTimeActivitystopwatches(ctx context.Context) ApiGetTimeActivitystopwatchesRequest {
	return ApiGetTimeActivitystopwatchesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []ActivityStopwatch
func (a *ActivityStopwatchesAPIService) GetTimeActivitystopwatchesExecute(r ApiGetTimeActivitystopwatchesRequest) ([]ActivityStopwatch, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ActivityStopwatch
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ActivityStopwatchesAPIService.GetTimeActivitystopwatches")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/time/activitystopwatches"

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

type ApiGetTimeActivitystopwatchesByIdRequest struct {
	ctx context.Context
	ApiService *ActivityStopwatchesAPIService
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
func (r ApiGetTimeActivitystopwatchesByIdRequest) Conditions(conditions string) ApiGetTimeActivitystopwatchesByIdRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetTimeActivitystopwatchesByIdRequest) ChildConditions(childConditions string) ApiGetTimeActivitystopwatchesByIdRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetTimeActivitystopwatchesByIdRequest) CustomFieldConditions(customFieldConditions string) ApiGetTimeActivitystopwatchesByIdRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetTimeActivitystopwatchesByIdRequest) OrderBy(orderBy string) ApiGetTimeActivitystopwatchesByIdRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetTimeActivitystopwatchesByIdRequest) Fields(fields string) ApiGetTimeActivitystopwatchesByIdRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetTimeActivitystopwatchesByIdRequest) Page(page int32) ApiGetTimeActivitystopwatchesByIdRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetTimeActivitystopwatchesByIdRequest) PageSize(pageSize int32) ApiGetTimeActivitystopwatchesByIdRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetTimeActivitystopwatchesByIdRequest) PageId(pageId int32) ApiGetTimeActivitystopwatchesByIdRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetTimeActivitystopwatchesByIdRequest) ClientId(clientId string) ApiGetTimeActivitystopwatchesByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetTimeActivitystopwatchesByIdRequest) Execute() (*ActivityStopwatch, *http.Response, error) {
	return r.ApiService.GetTimeActivitystopwatchesByIdExecute(r)
}

/*
GetTimeActivitystopwatchesById Get ActivityStopwatch

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id activitystopwatcheId
 @return ApiGetTimeActivitystopwatchesByIdRequest
*/
func (a *ActivityStopwatchesAPIService) GetTimeActivitystopwatchesById(ctx context.Context, id int32) ApiGetTimeActivitystopwatchesByIdRequest {
	return ApiGetTimeActivitystopwatchesByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ActivityStopwatch
func (a *ActivityStopwatchesAPIService) GetTimeActivitystopwatchesByIdExecute(r ApiGetTimeActivitystopwatchesByIdRequest) (*ActivityStopwatch, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ActivityStopwatch
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ActivityStopwatchesAPIService.GetTimeActivitystopwatchesById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/time/activitystopwatches/{id}"
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

type ApiGetTimeActivitystopwatchesCountRequest struct {
	ctx context.Context
	ApiService *ActivityStopwatchesAPIService
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
func (r ApiGetTimeActivitystopwatchesCountRequest) Conditions(conditions string) ApiGetTimeActivitystopwatchesCountRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetTimeActivitystopwatchesCountRequest) ChildConditions(childConditions string) ApiGetTimeActivitystopwatchesCountRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetTimeActivitystopwatchesCountRequest) CustomFieldConditions(customFieldConditions string) ApiGetTimeActivitystopwatchesCountRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetTimeActivitystopwatchesCountRequest) OrderBy(orderBy string) ApiGetTimeActivitystopwatchesCountRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetTimeActivitystopwatchesCountRequest) Fields(fields string) ApiGetTimeActivitystopwatchesCountRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetTimeActivitystopwatchesCountRequest) Page(page int32) ApiGetTimeActivitystopwatchesCountRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetTimeActivitystopwatchesCountRequest) PageSize(pageSize int32) ApiGetTimeActivitystopwatchesCountRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetTimeActivitystopwatchesCountRequest) PageId(pageId int32) ApiGetTimeActivitystopwatchesCountRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetTimeActivitystopwatchesCountRequest) ClientId(clientId string) ApiGetTimeActivitystopwatchesCountRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetTimeActivitystopwatchesCountRequest) Execute() (*Count, *http.Response, error) {
	return r.ApiService.GetTimeActivitystopwatchesCountExecute(r)
}

/*
GetTimeActivitystopwatchesCount Get Count of ActivityStopwatch

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetTimeActivitystopwatchesCountRequest
*/
func (a *ActivityStopwatchesAPIService) GetTimeActivitystopwatchesCount(ctx context.Context) ApiGetTimeActivitystopwatchesCountRequest {
	return ApiGetTimeActivitystopwatchesCountRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Count
func (a *ActivityStopwatchesAPIService) GetTimeActivitystopwatchesCountExecute(r ApiGetTimeActivitystopwatchesCountRequest) (*Count, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Count
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ActivityStopwatchesAPIService.GetTimeActivitystopwatchesCount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/time/activitystopwatches/count"

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

type ApiPatchTimeActivitystopwatchesByIdRequest struct {
	ctx context.Context
	ApiService *ActivityStopwatchesAPIService
	id int32
	patchOperation *[]PatchOperation
	clientId *string
}

// List of PatchOperation
func (r ApiPatchTimeActivitystopwatchesByIdRequest) PatchOperation(patchOperation []PatchOperation) ApiPatchTimeActivitystopwatchesByIdRequest {
	r.patchOperation = &patchOperation
	return r
}

// 
func (r ApiPatchTimeActivitystopwatchesByIdRequest) ClientId(clientId string) ApiPatchTimeActivitystopwatchesByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiPatchTimeActivitystopwatchesByIdRequest) Execute() (*ActivityStopwatch, *http.Response, error) {
	return r.ApiService.PatchTimeActivitystopwatchesByIdExecute(r)
}

/*
PatchTimeActivitystopwatchesById Patch ActivityStopwatch

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id activitystopwatcheId
 @return ApiPatchTimeActivitystopwatchesByIdRequest
*/
func (a *ActivityStopwatchesAPIService) PatchTimeActivitystopwatchesById(ctx context.Context, id int32) ApiPatchTimeActivitystopwatchesByIdRequest {
	return ApiPatchTimeActivitystopwatchesByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ActivityStopwatch
func (a *ActivityStopwatchesAPIService) PatchTimeActivitystopwatchesByIdExecute(r ApiPatchTimeActivitystopwatchesByIdRequest) (*ActivityStopwatch, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ActivityStopwatch
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ActivityStopwatchesAPIService.PatchTimeActivitystopwatchesById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/time/activitystopwatches/{id}"
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

type ApiPostTimeActivitystopwatchesRequest struct {
	ctx context.Context
	ApiService *ActivityStopwatchesAPIService
	activityStopwatch *ActivityStopwatch
	clientId *string
}

// activityStopwatch
func (r ApiPostTimeActivitystopwatchesRequest) ActivityStopwatch(activityStopwatch ActivityStopwatch) ApiPostTimeActivitystopwatchesRequest {
	r.activityStopwatch = &activityStopwatch
	return r
}

// 
func (r ApiPostTimeActivitystopwatchesRequest) ClientId(clientId string) ApiPostTimeActivitystopwatchesRequest {
	r.clientId = &clientId
	return r
}

func (r ApiPostTimeActivitystopwatchesRequest) Execute() (*ActivityStopwatch, *http.Response, error) {
	return r.ApiService.PostTimeActivitystopwatchesExecute(r)
}

/*
PostTimeActivitystopwatches Post ActivityStopwatch

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPostTimeActivitystopwatchesRequest
*/
func (a *ActivityStopwatchesAPIService) PostTimeActivitystopwatches(ctx context.Context) ApiPostTimeActivitystopwatchesRequest {
	return ApiPostTimeActivitystopwatchesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ActivityStopwatch
func (a *ActivityStopwatchesAPIService) PostTimeActivitystopwatchesExecute(r ApiPostTimeActivitystopwatchesRequest) (*ActivityStopwatch, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ActivityStopwatch
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ActivityStopwatchesAPIService.PostTimeActivitystopwatches")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/time/activitystopwatches"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.activityStopwatch == nil {
		return localVarReturnValue, nil, reportError("activityStopwatch is required and must be specified")
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
	localVarPostBody = r.activityStopwatch
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

type ApiPutTimeActivitystopwatchesByIdRequest struct {
	ctx context.Context
	ApiService *ActivityStopwatchesAPIService
	id int32
	activityStopwatch *ActivityStopwatch
	clientId *string
}

// activityStopwatch
func (r ApiPutTimeActivitystopwatchesByIdRequest) ActivityStopwatch(activityStopwatch ActivityStopwatch) ApiPutTimeActivitystopwatchesByIdRequest {
	r.activityStopwatch = &activityStopwatch
	return r
}

// 
func (r ApiPutTimeActivitystopwatchesByIdRequest) ClientId(clientId string) ApiPutTimeActivitystopwatchesByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiPutTimeActivitystopwatchesByIdRequest) Execute() (*ActivityStopwatch, *http.Response, error) {
	return r.ApiService.PutTimeActivitystopwatchesByIdExecute(r)
}

/*
PutTimeActivitystopwatchesById Put ActivityStopwatch

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id activitystopwatcheId
 @return ApiPutTimeActivitystopwatchesByIdRequest
*/
func (a *ActivityStopwatchesAPIService) PutTimeActivitystopwatchesById(ctx context.Context, id int32) ApiPutTimeActivitystopwatchesByIdRequest {
	return ApiPutTimeActivitystopwatchesByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ActivityStopwatch
func (a *ActivityStopwatchesAPIService) PutTimeActivitystopwatchesByIdExecute(r ApiPutTimeActivitystopwatchesByIdRequest) (*ActivityStopwatch, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ActivityStopwatch
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ActivityStopwatchesAPIService.PutTimeActivitystopwatchesById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/time/activitystopwatches/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.activityStopwatch == nil {
		return localVarReturnValue, nil, reportError("activityStopwatch is required and must be specified")
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
	localVarPostBody = r.activityStopwatch
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
