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


// ActivitiesAPIService ActivitiesAPI service
type ActivitiesAPIService service

type ApiDeleteSalesActivitiesByIdRequest struct {
	ctx context.Context
	ApiService *ActivitiesAPIService
	id int32
	clientId *string
}

// 
func (r ApiDeleteSalesActivitiesByIdRequest) ClientId(clientId string) ApiDeleteSalesActivitiesByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiDeleteSalesActivitiesByIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteSalesActivitiesByIdExecute(r)
}

/*
DeleteSalesActivitiesById Delete Activity

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id activityId
 @return ApiDeleteSalesActivitiesByIdRequest
*/
func (a *ActivitiesAPIService) DeleteSalesActivitiesById(ctx context.Context, id int32) ApiDeleteSalesActivitiesByIdRequest {
	return ApiDeleteSalesActivitiesByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *ActivitiesAPIService) DeleteSalesActivitiesByIdExecute(r ApiDeleteSalesActivitiesByIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ActivitiesAPIService.DeleteSalesActivitiesById")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sales/activities/{id}"
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

type ApiGetSalesActivitiesRequest struct {
	ctx context.Context
	ApiService *ActivitiesAPIService
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
func (r ApiGetSalesActivitiesRequest) Conditions(conditions string) ApiGetSalesActivitiesRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetSalesActivitiesRequest) ChildConditions(childConditions string) ApiGetSalesActivitiesRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetSalesActivitiesRequest) CustomFieldConditions(customFieldConditions string) ApiGetSalesActivitiesRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetSalesActivitiesRequest) OrderBy(orderBy string) ApiGetSalesActivitiesRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetSalesActivitiesRequest) Fields(fields string) ApiGetSalesActivitiesRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetSalesActivitiesRequest) Page(page int32) ApiGetSalesActivitiesRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetSalesActivitiesRequest) PageSize(pageSize int32) ApiGetSalesActivitiesRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetSalesActivitiesRequest) PageId(pageId int32) ApiGetSalesActivitiesRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetSalesActivitiesRequest) ClientId(clientId string) ApiGetSalesActivitiesRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetSalesActivitiesRequest) Execute() ([]Activity, *http.Response, error) {
	return r.ApiService.GetSalesActivitiesExecute(r)
}

/*
GetSalesActivities Get List of Activity

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetSalesActivitiesRequest
*/
func (a *ActivitiesAPIService) GetSalesActivities(ctx context.Context) ApiGetSalesActivitiesRequest {
	return ApiGetSalesActivitiesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []Activity
func (a *ActivitiesAPIService) GetSalesActivitiesExecute(r ApiGetSalesActivitiesRequest) ([]Activity, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Activity
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ActivitiesAPIService.GetSalesActivities")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sales/activities"

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

type ApiGetSalesActivitiesByIdRequest struct {
	ctx context.Context
	ApiService *ActivitiesAPIService
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
func (r ApiGetSalesActivitiesByIdRequest) Conditions(conditions string) ApiGetSalesActivitiesByIdRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetSalesActivitiesByIdRequest) ChildConditions(childConditions string) ApiGetSalesActivitiesByIdRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetSalesActivitiesByIdRequest) CustomFieldConditions(customFieldConditions string) ApiGetSalesActivitiesByIdRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetSalesActivitiesByIdRequest) OrderBy(orderBy string) ApiGetSalesActivitiesByIdRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetSalesActivitiesByIdRequest) Fields(fields string) ApiGetSalesActivitiesByIdRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetSalesActivitiesByIdRequest) Page(page int32) ApiGetSalesActivitiesByIdRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetSalesActivitiesByIdRequest) PageSize(pageSize int32) ApiGetSalesActivitiesByIdRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetSalesActivitiesByIdRequest) PageId(pageId int32) ApiGetSalesActivitiesByIdRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetSalesActivitiesByIdRequest) ClientId(clientId string) ApiGetSalesActivitiesByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetSalesActivitiesByIdRequest) Execute() (*Activity, *http.Response, error) {
	return r.ApiService.GetSalesActivitiesByIdExecute(r)
}

/*
GetSalesActivitiesById Get Activity

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id activityId
 @return ApiGetSalesActivitiesByIdRequest
*/
func (a *ActivitiesAPIService) GetSalesActivitiesById(ctx context.Context, id int32) ApiGetSalesActivitiesByIdRequest {
	return ApiGetSalesActivitiesByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return Activity
func (a *ActivitiesAPIService) GetSalesActivitiesByIdExecute(r ApiGetSalesActivitiesByIdRequest) (*Activity, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Activity
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ActivitiesAPIService.GetSalesActivitiesById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sales/activities/{id}"
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

type ApiGetSalesActivitiesCountRequest struct {
	ctx context.Context
	ApiService *ActivitiesAPIService
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
func (r ApiGetSalesActivitiesCountRequest) Conditions(conditions string) ApiGetSalesActivitiesCountRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetSalesActivitiesCountRequest) ChildConditions(childConditions string) ApiGetSalesActivitiesCountRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetSalesActivitiesCountRequest) CustomFieldConditions(customFieldConditions string) ApiGetSalesActivitiesCountRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetSalesActivitiesCountRequest) OrderBy(orderBy string) ApiGetSalesActivitiesCountRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetSalesActivitiesCountRequest) Fields(fields string) ApiGetSalesActivitiesCountRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetSalesActivitiesCountRequest) Page(page int32) ApiGetSalesActivitiesCountRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetSalesActivitiesCountRequest) PageSize(pageSize int32) ApiGetSalesActivitiesCountRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetSalesActivitiesCountRequest) PageId(pageId int32) ApiGetSalesActivitiesCountRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetSalesActivitiesCountRequest) ClientId(clientId string) ApiGetSalesActivitiesCountRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetSalesActivitiesCountRequest) Execute() (*Count, *http.Response, error) {
	return r.ApiService.GetSalesActivitiesCountExecute(r)
}

/*
GetSalesActivitiesCount Get Count of Activity

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetSalesActivitiesCountRequest
*/
func (a *ActivitiesAPIService) GetSalesActivitiesCount(ctx context.Context) ApiGetSalesActivitiesCountRequest {
	return ApiGetSalesActivitiesCountRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Count
func (a *ActivitiesAPIService) GetSalesActivitiesCountExecute(r ApiGetSalesActivitiesCountRequest) (*Count, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Count
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ActivitiesAPIService.GetSalesActivitiesCount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sales/activities/count"

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

type ApiPatchSalesActivitiesByIdRequest struct {
	ctx context.Context
	ApiService *ActivitiesAPIService
	id int32
	patchOperation *[]PatchOperation
	clientId *string
}

// List of PatchOperation
func (r ApiPatchSalesActivitiesByIdRequest) PatchOperation(patchOperation []PatchOperation) ApiPatchSalesActivitiesByIdRequest {
	r.patchOperation = &patchOperation
	return r
}

// 
func (r ApiPatchSalesActivitiesByIdRequest) ClientId(clientId string) ApiPatchSalesActivitiesByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiPatchSalesActivitiesByIdRequest) Execute() (*Activity, *http.Response, error) {
	return r.ApiService.PatchSalesActivitiesByIdExecute(r)
}

/*
PatchSalesActivitiesById Patch Activity

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id activityId
 @return ApiPatchSalesActivitiesByIdRequest
*/
func (a *ActivitiesAPIService) PatchSalesActivitiesById(ctx context.Context, id int32) ApiPatchSalesActivitiesByIdRequest {
	return ApiPatchSalesActivitiesByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return Activity
func (a *ActivitiesAPIService) PatchSalesActivitiesByIdExecute(r ApiPatchSalesActivitiesByIdRequest) (*Activity, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Activity
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ActivitiesAPIService.PatchSalesActivitiesById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sales/activities/{id}"
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

type ApiPostSalesActivitiesRequest struct {
	ctx context.Context
	ApiService *ActivitiesAPIService
	activity *Activity
	clientId *string
}

// activity
func (r ApiPostSalesActivitiesRequest) Activity(activity Activity) ApiPostSalesActivitiesRequest {
	r.activity = &activity
	return r
}

// 
func (r ApiPostSalesActivitiesRequest) ClientId(clientId string) ApiPostSalesActivitiesRequest {
	r.clientId = &clientId
	return r
}

func (r ApiPostSalesActivitiesRequest) Execute() (*Activity, *http.Response, error) {
	return r.ApiService.PostSalesActivitiesExecute(r)
}

/*
PostSalesActivities Post Activity

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPostSalesActivitiesRequest
*/
func (a *ActivitiesAPIService) PostSalesActivities(ctx context.Context) ApiPostSalesActivitiesRequest {
	return ApiPostSalesActivitiesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Activity
func (a *ActivitiesAPIService) PostSalesActivitiesExecute(r ApiPostSalesActivitiesRequest) (*Activity, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Activity
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ActivitiesAPIService.PostSalesActivities")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sales/activities"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.activity == nil {
		return localVarReturnValue, nil, reportError("activity is required and must be specified")
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
	localVarPostBody = r.activity
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

type ApiPutSalesActivitiesByIdRequest struct {
	ctx context.Context
	ApiService *ActivitiesAPIService
	id int32
	activity *Activity
	clientId *string
}

// activity
func (r ApiPutSalesActivitiesByIdRequest) Activity(activity Activity) ApiPutSalesActivitiesByIdRequest {
	r.activity = &activity
	return r
}

// 
func (r ApiPutSalesActivitiesByIdRequest) ClientId(clientId string) ApiPutSalesActivitiesByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiPutSalesActivitiesByIdRequest) Execute() (*Activity, *http.Response, error) {
	return r.ApiService.PutSalesActivitiesByIdExecute(r)
}

/*
PutSalesActivitiesById Put Activity

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id activityId
 @return ApiPutSalesActivitiesByIdRequest
*/
func (a *ActivitiesAPIService) PutSalesActivitiesById(ctx context.Context, id int32) ApiPutSalesActivitiesByIdRequest {
	return ApiPutSalesActivitiesByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return Activity
func (a *ActivitiesAPIService) PutSalesActivitiesByIdExecute(r ApiPutSalesActivitiesByIdRequest) (*Activity, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Activity
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ActivitiesAPIService.PutSalesActivitiesById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sales/activities/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.activity == nil {
		return localVarReturnValue, nil, reportError("activity is required and must be specified")
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
	localVarPostBody = r.activity
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
