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


// UserDefinedFieldInfosAPIService UserDefinedFieldInfosAPI service
type UserDefinedFieldInfosAPIService service

type ApiGetSystemUserDefinedFieldsByIdInfoRequest struct {
	ctx context.Context
	ApiService *UserDefinedFieldInfosAPIService
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
func (r ApiGetSystemUserDefinedFieldsByIdInfoRequest) Conditions(conditions string) ApiGetSystemUserDefinedFieldsByIdInfoRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetSystemUserDefinedFieldsByIdInfoRequest) ChildConditions(childConditions string) ApiGetSystemUserDefinedFieldsByIdInfoRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetSystemUserDefinedFieldsByIdInfoRequest) CustomFieldConditions(customFieldConditions string) ApiGetSystemUserDefinedFieldsByIdInfoRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetSystemUserDefinedFieldsByIdInfoRequest) OrderBy(orderBy string) ApiGetSystemUserDefinedFieldsByIdInfoRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetSystemUserDefinedFieldsByIdInfoRequest) Fields(fields string) ApiGetSystemUserDefinedFieldsByIdInfoRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetSystemUserDefinedFieldsByIdInfoRequest) Page(page int32) ApiGetSystemUserDefinedFieldsByIdInfoRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetSystemUserDefinedFieldsByIdInfoRequest) PageSize(pageSize int32) ApiGetSystemUserDefinedFieldsByIdInfoRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetSystemUserDefinedFieldsByIdInfoRequest) PageId(pageId int32) ApiGetSystemUserDefinedFieldsByIdInfoRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetSystemUserDefinedFieldsByIdInfoRequest) ClientId(clientId string) ApiGetSystemUserDefinedFieldsByIdInfoRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetSystemUserDefinedFieldsByIdInfoRequest) Execute() (*UserDefinedFieldInfo, *http.Response, error) {
	return r.ApiService.GetSystemUserDefinedFieldsByIdInfoExecute(r)
}

/*
GetSystemUserDefinedFieldsByIdInfo Get UserDefinedFieldInfos

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id UserDefinedFieldInfoId
 @return ApiGetSystemUserDefinedFieldsByIdInfoRequest
*/
func (a *UserDefinedFieldInfosAPIService) GetSystemUserDefinedFieldsByIdInfo(ctx context.Context, id int32) ApiGetSystemUserDefinedFieldsByIdInfoRequest {
	return ApiGetSystemUserDefinedFieldsByIdInfoRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return UserDefinedFieldInfo
func (a *UserDefinedFieldInfosAPIService) GetSystemUserDefinedFieldsByIdInfoExecute(r ApiGetSystemUserDefinedFieldsByIdInfoRequest) (*UserDefinedFieldInfo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UserDefinedFieldInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserDefinedFieldInfosAPIService.GetSystemUserDefinedFieldsByIdInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/userDefinedFields/{id}/info"
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

type ApiGetSystemUserDefinedFieldsInfoRequest struct {
	ctx context.Context
	ApiService *UserDefinedFieldInfosAPIService
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
func (r ApiGetSystemUserDefinedFieldsInfoRequest) Conditions(conditions string) ApiGetSystemUserDefinedFieldsInfoRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetSystemUserDefinedFieldsInfoRequest) ChildConditions(childConditions string) ApiGetSystemUserDefinedFieldsInfoRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetSystemUserDefinedFieldsInfoRequest) CustomFieldConditions(customFieldConditions string) ApiGetSystemUserDefinedFieldsInfoRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetSystemUserDefinedFieldsInfoRequest) OrderBy(orderBy string) ApiGetSystemUserDefinedFieldsInfoRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetSystemUserDefinedFieldsInfoRequest) Fields(fields string) ApiGetSystemUserDefinedFieldsInfoRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetSystemUserDefinedFieldsInfoRequest) Page(page int32) ApiGetSystemUserDefinedFieldsInfoRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetSystemUserDefinedFieldsInfoRequest) PageSize(pageSize int32) ApiGetSystemUserDefinedFieldsInfoRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetSystemUserDefinedFieldsInfoRequest) PageId(pageId int32) ApiGetSystemUserDefinedFieldsInfoRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetSystemUserDefinedFieldsInfoRequest) ClientId(clientId string) ApiGetSystemUserDefinedFieldsInfoRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetSystemUserDefinedFieldsInfoRequest) Execute() ([]UserDefinedFieldInfo, *http.Response, error) {
	return r.ApiService.GetSystemUserDefinedFieldsInfoExecute(r)
}

/*
GetSystemUserDefinedFieldsInfo Get List of UserDefinedFieldInfos

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetSystemUserDefinedFieldsInfoRequest
*/
func (a *UserDefinedFieldInfosAPIService) GetSystemUserDefinedFieldsInfo(ctx context.Context) ApiGetSystemUserDefinedFieldsInfoRequest {
	return ApiGetSystemUserDefinedFieldsInfoRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []UserDefinedFieldInfo
func (a *UserDefinedFieldInfosAPIService) GetSystemUserDefinedFieldsInfoExecute(r ApiGetSystemUserDefinedFieldsInfoRequest) ([]UserDefinedFieldInfo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []UserDefinedFieldInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserDefinedFieldInfosAPIService.GetSystemUserDefinedFieldsInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/userDefinedFields/info"

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

type ApiGetSystemUserDefinedFieldsInfoCountRequest struct {
	ctx context.Context
	ApiService *UserDefinedFieldInfosAPIService
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
func (r ApiGetSystemUserDefinedFieldsInfoCountRequest) Conditions(conditions string) ApiGetSystemUserDefinedFieldsInfoCountRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetSystemUserDefinedFieldsInfoCountRequest) ChildConditions(childConditions string) ApiGetSystemUserDefinedFieldsInfoCountRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetSystemUserDefinedFieldsInfoCountRequest) CustomFieldConditions(customFieldConditions string) ApiGetSystemUserDefinedFieldsInfoCountRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetSystemUserDefinedFieldsInfoCountRequest) OrderBy(orderBy string) ApiGetSystemUserDefinedFieldsInfoCountRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetSystemUserDefinedFieldsInfoCountRequest) Fields(fields string) ApiGetSystemUserDefinedFieldsInfoCountRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetSystemUserDefinedFieldsInfoCountRequest) Page(page int32) ApiGetSystemUserDefinedFieldsInfoCountRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetSystemUserDefinedFieldsInfoCountRequest) PageSize(pageSize int32) ApiGetSystemUserDefinedFieldsInfoCountRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetSystemUserDefinedFieldsInfoCountRequest) PageId(pageId int32) ApiGetSystemUserDefinedFieldsInfoCountRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetSystemUserDefinedFieldsInfoCountRequest) ClientId(clientId string) ApiGetSystemUserDefinedFieldsInfoCountRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetSystemUserDefinedFieldsInfoCountRequest) Execute() (*Count, *http.Response, error) {
	return r.ApiService.GetSystemUserDefinedFieldsInfoCountExecute(r)
}

/*
GetSystemUserDefinedFieldsInfoCount Get Count of UserDefinedFieldInfos

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetSystemUserDefinedFieldsInfoCountRequest
*/
func (a *UserDefinedFieldInfosAPIService) GetSystemUserDefinedFieldsInfoCount(ctx context.Context) ApiGetSystemUserDefinedFieldsInfoCountRequest {
	return ApiGetSystemUserDefinedFieldsInfoCountRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Count
func (a *UserDefinedFieldInfosAPIService) GetSystemUserDefinedFieldsInfoCountExecute(r ApiGetSystemUserDefinedFieldsInfoCountRequest) (*Count, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Count
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserDefinedFieldInfosAPIService.GetSystemUserDefinedFieldsInfoCount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/userDefinedFields/info/count"

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
