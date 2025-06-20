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


// TimeZoneSetupInfosAPIService TimeZoneSetupInfosAPI service
type TimeZoneSetupInfosAPIService service

type ApiGetSystemTimeZoneSetupsByIdInfoRequest struct {
	ctx context.Context
	ApiService *TimeZoneSetupInfosAPIService
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
func (r ApiGetSystemTimeZoneSetupsByIdInfoRequest) Conditions(conditions string) ApiGetSystemTimeZoneSetupsByIdInfoRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetSystemTimeZoneSetupsByIdInfoRequest) ChildConditions(childConditions string) ApiGetSystemTimeZoneSetupsByIdInfoRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetSystemTimeZoneSetupsByIdInfoRequest) CustomFieldConditions(customFieldConditions string) ApiGetSystemTimeZoneSetupsByIdInfoRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetSystemTimeZoneSetupsByIdInfoRequest) OrderBy(orderBy string) ApiGetSystemTimeZoneSetupsByIdInfoRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetSystemTimeZoneSetupsByIdInfoRequest) Fields(fields string) ApiGetSystemTimeZoneSetupsByIdInfoRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetSystemTimeZoneSetupsByIdInfoRequest) Page(page int32) ApiGetSystemTimeZoneSetupsByIdInfoRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetSystemTimeZoneSetupsByIdInfoRequest) PageSize(pageSize int32) ApiGetSystemTimeZoneSetupsByIdInfoRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetSystemTimeZoneSetupsByIdInfoRequest) PageId(pageId int32) ApiGetSystemTimeZoneSetupsByIdInfoRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetSystemTimeZoneSetupsByIdInfoRequest) ClientId(clientId string) ApiGetSystemTimeZoneSetupsByIdInfoRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetSystemTimeZoneSetupsByIdInfoRequest) Execute() (*TimeZoneSetupInfo, *http.Response, error) {
	return r.ApiService.GetSystemTimeZoneSetupsByIdInfoExecute(r)
}

/*
GetSystemTimeZoneSetupsByIdInfo Get TimeZoneSetupInfos

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id TimeZoneSetupInfoId
 @return ApiGetSystemTimeZoneSetupsByIdInfoRequest
*/
func (a *TimeZoneSetupInfosAPIService) GetSystemTimeZoneSetupsByIdInfo(ctx context.Context, id int32) ApiGetSystemTimeZoneSetupsByIdInfoRequest {
	return ApiGetSystemTimeZoneSetupsByIdInfoRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return TimeZoneSetupInfo
func (a *TimeZoneSetupInfosAPIService) GetSystemTimeZoneSetupsByIdInfoExecute(r ApiGetSystemTimeZoneSetupsByIdInfoRequest) (*TimeZoneSetupInfo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TimeZoneSetupInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TimeZoneSetupInfosAPIService.GetSystemTimeZoneSetupsByIdInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/timeZoneSetups/{id}/info"
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

type ApiGetSystemTimeZoneSetupsInfoRequest struct {
	ctx context.Context
	ApiService *TimeZoneSetupInfosAPIService
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
func (r ApiGetSystemTimeZoneSetupsInfoRequest) Conditions(conditions string) ApiGetSystemTimeZoneSetupsInfoRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetSystemTimeZoneSetupsInfoRequest) ChildConditions(childConditions string) ApiGetSystemTimeZoneSetupsInfoRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetSystemTimeZoneSetupsInfoRequest) CustomFieldConditions(customFieldConditions string) ApiGetSystemTimeZoneSetupsInfoRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetSystemTimeZoneSetupsInfoRequest) OrderBy(orderBy string) ApiGetSystemTimeZoneSetupsInfoRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetSystemTimeZoneSetupsInfoRequest) Fields(fields string) ApiGetSystemTimeZoneSetupsInfoRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetSystemTimeZoneSetupsInfoRequest) Page(page int32) ApiGetSystemTimeZoneSetupsInfoRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetSystemTimeZoneSetupsInfoRequest) PageSize(pageSize int32) ApiGetSystemTimeZoneSetupsInfoRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetSystemTimeZoneSetupsInfoRequest) PageId(pageId int32) ApiGetSystemTimeZoneSetupsInfoRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetSystemTimeZoneSetupsInfoRequest) ClientId(clientId string) ApiGetSystemTimeZoneSetupsInfoRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetSystemTimeZoneSetupsInfoRequest) Execute() ([]TimeZoneSetupInfo, *http.Response, error) {
	return r.ApiService.GetSystemTimeZoneSetupsInfoExecute(r)
}

/*
GetSystemTimeZoneSetupsInfo Get List of TimeZoneSetupInfos

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetSystemTimeZoneSetupsInfoRequest
*/
func (a *TimeZoneSetupInfosAPIService) GetSystemTimeZoneSetupsInfo(ctx context.Context) ApiGetSystemTimeZoneSetupsInfoRequest {
	return ApiGetSystemTimeZoneSetupsInfoRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []TimeZoneSetupInfo
func (a *TimeZoneSetupInfosAPIService) GetSystemTimeZoneSetupsInfoExecute(r ApiGetSystemTimeZoneSetupsInfoRequest) ([]TimeZoneSetupInfo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []TimeZoneSetupInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TimeZoneSetupInfosAPIService.GetSystemTimeZoneSetupsInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/timeZoneSetups/info"

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

type ApiGetSystemTimeZoneSetupsInfoCountRequest struct {
	ctx context.Context
	ApiService *TimeZoneSetupInfosAPIService
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
func (r ApiGetSystemTimeZoneSetupsInfoCountRequest) Conditions(conditions string) ApiGetSystemTimeZoneSetupsInfoCountRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetSystemTimeZoneSetupsInfoCountRequest) ChildConditions(childConditions string) ApiGetSystemTimeZoneSetupsInfoCountRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetSystemTimeZoneSetupsInfoCountRequest) CustomFieldConditions(customFieldConditions string) ApiGetSystemTimeZoneSetupsInfoCountRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetSystemTimeZoneSetupsInfoCountRequest) OrderBy(orderBy string) ApiGetSystemTimeZoneSetupsInfoCountRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetSystemTimeZoneSetupsInfoCountRequest) Fields(fields string) ApiGetSystemTimeZoneSetupsInfoCountRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetSystemTimeZoneSetupsInfoCountRequest) Page(page int32) ApiGetSystemTimeZoneSetupsInfoCountRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetSystemTimeZoneSetupsInfoCountRequest) PageSize(pageSize int32) ApiGetSystemTimeZoneSetupsInfoCountRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetSystemTimeZoneSetupsInfoCountRequest) PageId(pageId int32) ApiGetSystemTimeZoneSetupsInfoCountRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetSystemTimeZoneSetupsInfoCountRequest) ClientId(clientId string) ApiGetSystemTimeZoneSetupsInfoCountRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetSystemTimeZoneSetupsInfoCountRequest) Execute() (*Count, *http.Response, error) {
	return r.ApiService.GetSystemTimeZoneSetupsInfoCountExecute(r)
}

/*
GetSystemTimeZoneSetupsInfoCount Get Count of TimeZoneSetupInfos

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetSystemTimeZoneSetupsInfoCountRequest
*/
func (a *TimeZoneSetupInfosAPIService) GetSystemTimeZoneSetupsInfoCount(ctx context.Context) ApiGetSystemTimeZoneSetupsInfoCountRequest {
	return ApiGetSystemTimeZoneSetupsInfoCountRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Count
func (a *TimeZoneSetupInfosAPIService) GetSystemTimeZoneSetupsInfoCountExecute(r ApiGetSystemTimeZoneSetupsInfoCountRequest) (*Count, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Count
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TimeZoneSetupInfosAPIService.GetSystemTimeZoneSetupsInfoCount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/timeZoneSetups/info/count"

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
