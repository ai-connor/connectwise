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


// ReportCardInfosAPIService ReportCardInfosAPI service
type ReportCardInfosAPIService service

type ApiGetSystemReportCardsByIdInfoRequest struct {
	ctx context.Context
	ApiService *ReportCardInfosAPIService
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
func (r ApiGetSystemReportCardsByIdInfoRequest) Conditions(conditions string) ApiGetSystemReportCardsByIdInfoRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetSystemReportCardsByIdInfoRequest) ChildConditions(childConditions string) ApiGetSystemReportCardsByIdInfoRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetSystemReportCardsByIdInfoRequest) CustomFieldConditions(customFieldConditions string) ApiGetSystemReportCardsByIdInfoRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetSystemReportCardsByIdInfoRequest) OrderBy(orderBy string) ApiGetSystemReportCardsByIdInfoRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetSystemReportCardsByIdInfoRequest) Fields(fields string) ApiGetSystemReportCardsByIdInfoRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetSystemReportCardsByIdInfoRequest) Page(page int32) ApiGetSystemReportCardsByIdInfoRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetSystemReportCardsByIdInfoRequest) PageSize(pageSize int32) ApiGetSystemReportCardsByIdInfoRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetSystemReportCardsByIdInfoRequest) PageId(pageId int32) ApiGetSystemReportCardsByIdInfoRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetSystemReportCardsByIdInfoRequest) ClientId(clientId string) ApiGetSystemReportCardsByIdInfoRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetSystemReportCardsByIdInfoRequest) Execute() (*ReportCardInfo, *http.Response, error) {
	return r.ApiService.GetSystemReportCardsByIdInfoExecute(r)
}

/*
GetSystemReportCardsByIdInfo Get ReportCardInfo

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id reportCardId
 @return ApiGetSystemReportCardsByIdInfoRequest
*/
func (a *ReportCardInfosAPIService) GetSystemReportCardsByIdInfo(ctx context.Context, id int32) ApiGetSystemReportCardsByIdInfoRequest {
	return ApiGetSystemReportCardsByIdInfoRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ReportCardInfo
func (a *ReportCardInfosAPIService) GetSystemReportCardsByIdInfoExecute(r ApiGetSystemReportCardsByIdInfoRequest) (*ReportCardInfo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ReportCardInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReportCardInfosAPIService.GetSystemReportCardsByIdInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/reportCards/{id}/info"
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

type ApiGetSystemReportCardsInfoRequest struct {
	ctx context.Context
	ApiService *ReportCardInfosAPIService
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
func (r ApiGetSystemReportCardsInfoRequest) Conditions(conditions string) ApiGetSystemReportCardsInfoRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetSystemReportCardsInfoRequest) ChildConditions(childConditions string) ApiGetSystemReportCardsInfoRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetSystemReportCardsInfoRequest) CustomFieldConditions(customFieldConditions string) ApiGetSystemReportCardsInfoRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetSystemReportCardsInfoRequest) OrderBy(orderBy string) ApiGetSystemReportCardsInfoRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetSystemReportCardsInfoRequest) Fields(fields string) ApiGetSystemReportCardsInfoRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetSystemReportCardsInfoRequest) Page(page int32) ApiGetSystemReportCardsInfoRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetSystemReportCardsInfoRequest) PageSize(pageSize int32) ApiGetSystemReportCardsInfoRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetSystemReportCardsInfoRequest) PageId(pageId int32) ApiGetSystemReportCardsInfoRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetSystemReportCardsInfoRequest) ClientId(clientId string) ApiGetSystemReportCardsInfoRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetSystemReportCardsInfoRequest) Execute() ([]ReportCardInfo, *http.Response, error) {
	return r.ApiService.GetSystemReportCardsInfoExecute(r)
}

/*
GetSystemReportCardsInfo Get List of ReportCardInfo

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetSystemReportCardsInfoRequest
*/
func (a *ReportCardInfosAPIService) GetSystemReportCardsInfo(ctx context.Context) ApiGetSystemReportCardsInfoRequest {
	return ApiGetSystemReportCardsInfoRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []ReportCardInfo
func (a *ReportCardInfosAPIService) GetSystemReportCardsInfoExecute(r ApiGetSystemReportCardsInfoRequest) ([]ReportCardInfo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ReportCardInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReportCardInfosAPIService.GetSystemReportCardsInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/reportCards/info"

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

type ApiGetSystemReportCardsInfoCountRequest struct {
	ctx context.Context
	ApiService *ReportCardInfosAPIService
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
func (r ApiGetSystemReportCardsInfoCountRequest) Conditions(conditions string) ApiGetSystemReportCardsInfoCountRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetSystemReportCardsInfoCountRequest) ChildConditions(childConditions string) ApiGetSystemReportCardsInfoCountRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetSystemReportCardsInfoCountRequest) CustomFieldConditions(customFieldConditions string) ApiGetSystemReportCardsInfoCountRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetSystemReportCardsInfoCountRequest) OrderBy(orderBy string) ApiGetSystemReportCardsInfoCountRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetSystemReportCardsInfoCountRequest) Fields(fields string) ApiGetSystemReportCardsInfoCountRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetSystemReportCardsInfoCountRequest) Page(page int32) ApiGetSystemReportCardsInfoCountRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetSystemReportCardsInfoCountRequest) PageSize(pageSize int32) ApiGetSystemReportCardsInfoCountRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetSystemReportCardsInfoCountRequest) PageId(pageId int32) ApiGetSystemReportCardsInfoCountRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetSystemReportCardsInfoCountRequest) ClientId(clientId string) ApiGetSystemReportCardsInfoCountRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetSystemReportCardsInfoCountRequest) Execute() (*Count, *http.Response, error) {
	return r.ApiService.GetSystemReportCardsInfoCountExecute(r)
}

/*
GetSystemReportCardsInfoCount Get Count of ReportCardInfo

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetSystemReportCardsInfoCountRequest
*/
func (a *ReportCardInfosAPIService) GetSystemReportCardsInfoCount(ctx context.Context) ApiGetSystemReportCardsInfoCountRequest {
	return ApiGetSystemReportCardsInfoCountRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Count
func (a *ReportCardInfosAPIService) GetSystemReportCardsInfoCountExecute(r ApiGetSystemReportCardsInfoCountRequest) (*Count, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Count
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReportCardInfosAPIService.GetSystemReportCardsInfoCount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/reportCards/info/count"

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
