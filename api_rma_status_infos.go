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


// RmaStatusInfosAPIService RmaStatusInfosAPI service
type RmaStatusInfosAPIService service

type ApiGetProcurementRmaStatusesByIdInfoRequest struct {
	ctx context.Context
	ApiService *RmaStatusInfosAPIService
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
func (r ApiGetProcurementRmaStatusesByIdInfoRequest) Conditions(conditions string) ApiGetProcurementRmaStatusesByIdInfoRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetProcurementRmaStatusesByIdInfoRequest) ChildConditions(childConditions string) ApiGetProcurementRmaStatusesByIdInfoRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetProcurementRmaStatusesByIdInfoRequest) CustomFieldConditions(customFieldConditions string) ApiGetProcurementRmaStatusesByIdInfoRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetProcurementRmaStatusesByIdInfoRequest) OrderBy(orderBy string) ApiGetProcurementRmaStatusesByIdInfoRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetProcurementRmaStatusesByIdInfoRequest) Fields(fields string) ApiGetProcurementRmaStatusesByIdInfoRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetProcurementRmaStatusesByIdInfoRequest) Page(page int32) ApiGetProcurementRmaStatusesByIdInfoRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetProcurementRmaStatusesByIdInfoRequest) PageSize(pageSize int32) ApiGetProcurementRmaStatusesByIdInfoRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetProcurementRmaStatusesByIdInfoRequest) PageId(pageId int32) ApiGetProcurementRmaStatusesByIdInfoRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetProcurementRmaStatusesByIdInfoRequest) ClientId(clientId string) ApiGetProcurementRmaStatusesByIdInfoRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetProcurementRmaStatusesByIdInfoRequest) Execute() (*RmaStatusInfo, *http.Response, error) {
	return r.ApiService.GetProcurementRmaStatusesByIdInfoExecute(r)
}

/*
GetProcurementRmaStatusesByIdInfo Get RmaStatusInfos

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id RmaStatusInfoId
 @return ApiGetProcurementRmaStatusesByIdInfoRequest
*/
func (a *RmaStatusInfosAPIService) GetProcurementRmaStatusesByIdInfo(ctx context.Context, id int32) ApiGetProcurementRmaStatusesByIdInfoRequest {
	return ApiGetProcurementRmaStatusesByIdInfoRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return RmaStatusInfo
func (a *RmaStatusInfosAPIService) GetProcurementRmaStatusesByIdInfoExecute(r ApiGetProcurementRmaStatusesByIdInfoRequest) (*RmaStatusInfo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RmaStatusInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RmaStatusInfosAPIService.GetProcurementRmaStatusesByIdInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/procurement/rmaStatuses/{id}/info"
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

type ApiGetProcurementRmaStatusesInfoRequest struct {
	ctx context.Context
	ApiService *RmaStatusInfosAPIService
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
func (r ApiGetProcurementRmaStatusesInfoRequest) Conditions(conditions string) ApiGetProcurementRmaStatusesInfoRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetProcurementRmaStatusesInfoRequest) ChildConditions(childConditions string) ApiGetProcurementRmaStatusesInfoRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetProcurementRmaStatusesInfoRequest) CustomFieldConditions(customFieldConditions string) ApiGetProcurementRmaStatusesInfoRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetProcurementRmaStatusesInfoRequest) OrderBy(orderBy string) ApiGetProcurementRmaStatusesInfoRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetProcurementRmaStatusesInfoRequest) Fields(fields string) ApiGetProcurementRmaStatusesInfoRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetProcurementRmaStatusesInfoRequest) Page(page int32) ApiGetProcurementRmaStatusesInfoRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetProcurementRmaStatusesInfoRequest) PageSize(pageSize int32) ApiGetProcurementRmaStatusesInfoRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetProcurementRmaStatusesInfoRequest) PageId(pageId int32) ApiGetProcurementRmaStatusesInfoRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetProcurementRmaStatusesInfoRequest) ClientId(clientId string) ApiGetProcurementRmaStatusesInfoRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetProcurementRmaStatusesInfoRequest) Execute() ([]RmaStatusInfo, *http.Response, error) {
	return r.ApiService.GetProcurementRmaStatusesInfoExecute(r)
}

/*
GetProcurementRmaStatusesInfo Get List of RmaStatusInfos

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetProcurementRmaStatusesInfoRequest
*/
func (a *RmaStatusInfosAPIService) GetProcurementRmaStatusesInfo(ctx context.Context) ApiGetProcurementRmaStatusesInfoRequest {
	return ApiGetProcurementRmaStatusesInfoRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []RmaStatusInfo
func (a *RmaStatusInfosAPIService) GetProcurementRmaStatusesInfoExecute(r ApiGetProcurementRmaStatusesInfoRequest) ([]RmaStatusInfo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []RmaStatusInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RmaStatusInfosAPIService.GetProcurementRmaStatusesInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/procurement/rmaStatuses/info"

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

type ApiGetProcurementRmaStatusesInfoCountRequest struct {
	ctx context.Context
	ApiService *RmaStatusInfosAPIService
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
func (r ApiGetProcurementRmaStatusesInfoCountRequest) Conditions(conditions string) ApiGetProcurementRmaStatusesInfoCountRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetProcurementRmaStatusesInfoCountRequest) ChildConditions(childConditions string) ApiGetProcurementRmaStatusesInfoCountRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetProcurementRmaStatusesInfoCountRequest) CustomFieldConditions(customFieldConditions string) ApiGetProcurementRmaStatusesInfoCountRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetProcurementRmaStatusesInfoCountRequest) OrderBy(orderBy string) ApiGetProcurementRmaStatusesInfoCountRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetProcurementRmaStatusesInfoCountRequest) Fields(fields string) ApiGetProcurementRmaStatusesInfoCountRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetProcurementRmaStatusesInfoCountRequest) Page(page int32) ApiGetProcurementRmaStatusesInfoCountRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetProcurementRmaStatusesInfoCountRequest) PageSize(pageSize int32) ApiGetProcurementRmaStatusesInfoCountRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetProcurementRmaStatusesInfoCountRequest) PageId(pageId int32) ApiGetProcurementRmaStatusesInfoCountRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetProcurementRmaStatusesInfoCountRequest) ClientId(clientId string) ApiGetProcurementRmaStatusesInfoCountRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetProcurementRmaStatusesInfoCountRequest) Execute() (*Count, *http.Response, error) {
	return r.ApiService.GetProcurementRmaStatusesInfoCountExecute(r)
}

/*
GetProcurementRmaStatusesInfoCount Get Count of RmaStatusInfos

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetProcurementRmaStatusesInfoCountRequest
*/
func (a *RmaStatusInfosAPIService) GetProcurementRmaStatusesInfoCount(ctx context.Context) ApiGetProcurementRmaStatusesInfoCountRequest {
	return ApiGetProcurementRmaStatusesInfoCountRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Count
func (a *RmaStatusInfosAPIService) GetProcurementRmaStatusesInfoCountExecute(r ApiGetProcurementRmaStatusesInfoCountRequest) (*Count, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Count
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RmaStatusInfosAPIService.GetProcurementRmaStatusesInfoCount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/procurement/rmaStatuses/info/count"

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
