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


// ScheduleDetailsAPIService ScheduleDetailsAPI service
type ScheduleDetailsAPIService service

type ApiGetScheduleEntriesByParentIdDetailsRequest struct {
	ctx context.Context
	ApiService *ScheduleDetailsAPIService
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
func (r ApiGetScheduleEntriesByParentIdDetailsRequest) Conditions(conditions string) ApiGetScheduleEntriesByParentIdDetailsRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetScheduleEntriesByParentIdDetailsRequest) ChildConditions(childConditions string) ApiGetScheduleEntriesByParentIdDetailsRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetScheduleEntriesByParentIdDetailsRequest) CustomFieldConditions(customFieldConditions string) ApiGetScheduleEntriesByParentIdDetailsRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetScheduleEntriesByParentIdDetailsRequest) OrderBy(orderBy string) ApiGetScheduleEntriesByParentIdDetailsRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetScheduleEntriesByParentIdDetailsRequest) Fields(fields string) ApiGetScheduleEntriesByParentIdDetailsRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetScheduleEntriesByParentIdDetailsRequest) Page(page int32) ApiGetScheduleEntriesByParentIdDetailsRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetScheduleEntriesByParentIdDetailsRequest) PageSize(pageSize int32) ApiGetScheduleEntriesByParentIdDetailsRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetScheduleEntriesByParentIdDetailsRequest) PageId(pageId int32) ApiGetScheduleEntriesByParentIdDetailsRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetScheduleEntriesByParentIdDetailsRequest) ClientId(clientId string) ApiGetScheduleEntriesByParentIdDetailsRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetScheduleEntriesByParentIdDetailsRequest) Execute() ([]ScheduleDetail, *http.Response, error) {
	return r.ApiService.GetScheduleEntriesByParentIdDetailsExecute(r)
}

/*
GetScheduleEntriesByParentIdDetails Get List of ScheduleDetail

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param parentId entryId
 @return ApiGetScheduleEntriesByParentIdDetailsRequest
*/
func (a *ScheduleDetailsAPIService) GetScheduleEntriesByParentIdDetails(ctx context.Context, parentId int32) ApiGetScheduleEntriesByParentIdDetailsRequest {
	return ApiGetScheduleEntriesByParentIdDetailsRequest{
		ApiService: a,
		ctx: ctx,
		parentId: parentId,
	}
}

// Execute executes the request
//  @return []ScheduleDetail
func (a *ScheduleDetailsAPIService) GetScheduleEntriesByParentIdDetailsExecute(r ApiGetScheduleEntriesByParentIdDetailsRequest) ([]ScheduleDetail, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ScheduleDetail
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ScheduleDetailsAPIService.GetScheduleEntriesByParentIdDetails")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/schedule/entries/{parentId}/details"
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

type ApiGetScheduleEntriesByParentIdDetailsByIdRequest struct {
	ctx context.Context
	ApiService *ScheduleDetailsAPIService
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
func (r ApiGetScheduleEntriesByParentIdDetailsByIdRequest) Conditions(conditions string) ApiGetScheduleEntriesByParentIdDetailsByIdRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetScheduleEntriesByParentIdDetailsByIdRequest) ChildConditions(childConditions string) ApiGetScheduleEntriesByParentIdDetailsByIdRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetScheduleEntriesByParentIdDetailsByIdRequest) CustomFieldConditions(customFieldConditions string) ApiGetScheduleEntriesByParentIdDetailsByIdRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetScheduleEntriesByParentIdDetailsByIdRequest) OrderBy(orderBy string) ApiGetScheduleEntriesByParentIdDetailsByIdRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetScheduleEntriesByParentIdDetailsByIdRequest) Fields(fields string) ApiGetScheduleEntriesByParentIdDetailsByIdRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetScheduleEntriesByParentIdDetailsByIdRequest) Page(page int32) ApiGetScheduleEntriesByParentIdDetailsByIdRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetScheduleEntriesByParentIdDetailsByIdRequest) PageSize(pageSize int32) ApiGetScheduleEntriesByParentIdDetailsByIdRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetScheduleEntriesByParentIdDetailsByIdRequest) PageId(pageId int32) ApiGetScheduleEntriesByParentIdDetailsByIdRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetScheduleEntriesByParentIdDetailsByIdRequest) ClientId(clientId string) ApiGetScheduleEntriesByParentIdDetailsByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetScheduleEntriesByParentIdDetailsByIdRequest) Execute() (*ScheduleDetail, *http.Response, error) {
	return r.ApiService.GetScheduleEntriesByParentIdDetailsByIdExecute(r)
}

/*
GetScheduleEntriesByParentIdDetailsById Get ScheduleDetail

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id detailId
 @param parentId entryId
 @return ApiGetScheduleEntriesByParentIdDetailsByIdRequest
*/
func (a *ScheduleDetailsAPIService) GetScheduleEntriesByParentIdDetailsById(ctx context.Context, id int32, parentId int32) ApiGetScheduleEntriesByParentIdDetailsByIdRequest {
	return ApiGetScheduleEntriesByParentIdDetailsByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
		parentId: parentId,
	}
}

// Execute executes the request
//  @return ScheduleDetail
func (a *ScheduleDetailsAPIService) GetScheduleEntriesByParentIdDetailsByIdExecute(r ApiGetScheduleEntriesByParentIdDetailsByIdRequest) (*ScheduleDetail, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ScheduleDetail
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ScheduleDetailsAPIService.GetScheduleEntriesByParentIdDetailsById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/schedule/entries/{parentId}/details/{id}"
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

type ApiGetScheduleEntriesByParentIdDetailsCountRequest struct {
	ctx context.Context
	ApiService *ScheduleDetailsAPIService
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
func (r ApiGetScheduleEntriesByParentIdDetailsCountRequest) Conditions(conditions string) ApiGetScheduleEntriesByParentIdDetailsCountRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetScheduleEntriesByParentIdDetailsCountRequest) ChildConditions(childConditions string) ApiGetScheduleEntriesByParentIdDetailsCountRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetScheduleEntriesByParentIdDetailsCountRequest) CustomFieldConditions(customFieldConditions string) ApiGetScheduleEntriesByParentIdDetailsCountRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetScheduleEntriesByParentIdDetailsCountRequest) OrderBy(orderBy string) ApiGetScheduleEntriesByParentIdDetailsCountRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetScheduleEntriesByParentIdDetailsCountRequest) Fields(fields string) ApiGetScheduleEntriesByParentIdDetailsCountRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetScheduleEntriesByParentIdDetailsCountRequest) Page(page int32) ApiGetScheduleEntriesByParentIdDetailsCountRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetScheduleEntriesByParentIdDetailsCountRequest) PageSize(pageSize int32) ApiGetScheduleEntriesByParentIdDetailsCountRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetScheduleEntriesByParentIdDetailsCountRequest) PageId(pageId int32) ApiGetScheduleEntriesByParentIdDetailsCountRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetScheduleEntriesByParentIdDetailsCountRequest) ClientId(clientId string) ApiGetScheduleEntriesByParentIdDetailsCountRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetScheduleEntriesByParentIdDetailsCountRequest) Execute() (*Count, *http.Response, error) {
	return r.ApiService.GetScheduleEntriesByParentIdDetailsCountExecute(r)
}

/*
GetScheduleEntriesByParentIdDetailsCount Get Count of ScheduleDetail

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param parentId entryId
 @return ApiGetScheduleEntriesByParentIdDetailsCountRequest
*/
func (a *ScheduleDetailsAPIService) GetScheduleEntriesByParentIdDetailsCount(ctx context.Context, parentId int32) ApiGetScheduleEntriesByParentIdDetailsCountRequest {
	return ApiGetScheduleEntriesByParentIdDetailsCountRequest{
		ApiService: a,
		ctx: ctx,
		parentId: parentId,
	}
}

// Execute executes the request
//  @return Count
func (a *ScheduleDetailsAPIService) GetScheduleEntriesByParentIdDetailsCountExecute(r ApiGetScheduleEntriesByParentIdDetailsCountRequest) (*Count, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Count
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ScheduleDetailsAPIService.GetScheduleEntriesByParentIdDetailsCount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/schedule/entries/{parentId}/details/count"
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
