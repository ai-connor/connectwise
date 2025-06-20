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


// ConfigurationTypeQuestionValueInfosAPIService ConfigurationTypeQuestionValueInfosAPI service
type ConfigurationTypeQuestionValueInfosAPIService service

type ApiGetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesByIdInfoRequest struct {
	ctx context.Context
	ApiService *ConfigurationTypeQuestionValueInfosAPIService
	grandparentId int32
	parentId int32
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
func (r ApiGetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesByIdInfoRequest) Conditions(conditions string) ApiGetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesByIdInfoRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesByIdInfoRequest) ChildConditions(childConditions string) ApiGetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesByIdInfoRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesByIdInfoRequest) CustomFieldConditions(customFieldConditions string) ApiGetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesByIdInfoRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesByIdInfoRequest) OrderBy(orderBy string) ApiGetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesByIdInfoRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesByIdInfoRequest) Fields(fields string) ApiGetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesByIdInfoRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesByIdInfoRequest) Page(page int32) ApiGetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesByIdInfoRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesByIdInfoRequest) PageSize(pageSize int32) ApiGetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesByIdInfoRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesByIdInfoRequest) PageId(pageId int32) ApiGetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesByIdInfoRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesByIdInfoRequest) ClientId(clientId string) ApiGetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesByIdInfoRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesByIdInfoRequest) Execute() (*ConfigurationTypeQuestionInfo, *http.Response, error) {
	return r.ApiService.GetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesByIdInfoExecute(r)
}

/*
GetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesByIdInfo Get ConfigurationTypeQuestionValueInfo

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param grandparentId ConfigurationTypeQuestionInfo
 @param parentId ConfigurationTypeQuestionInfo
 @param id ConfigurationTypeQuestionInfo
 @return ApiGetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesByIdInfoRequest
*/
func (a *ConfigurationTypeQuestionValueInfosAPIService) GetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesByIdInfo(ctx context.Context, grandparentId int32, parentId int32, id int32) ApiGetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesByIdInfoRequest {
	return ApiGetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesByIdInfoRequest{
		ApiService: a,
		ctx: ctx,
		grandparentId: grandparentId,
		parentId: parentId,
		id: id,
	}
}

// Execute executes the request
//  @return ConfigurationTypeQuestionInfo
func (a *ConfigurationTypeQuestionValueInfosAPIService) GetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesByIdInfoExecute(r ApiGetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesByIdInfoRequest) (*ConfigurationTypeQuestionInfo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ConfigurationTypeQuestionInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationTypeQuestionValueInfosAPIService.GetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesByIdInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/configurations/types/{grandparentId}/questions/{parentId}/values/{id}/info"
	localVarPath = strings.Replace(localVarPath, "{"+"grandparentId"+"}", url.PathEscape(parameterValueToString(r.grandparentId, "grandparentId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"parentId"+"}", url.PathEscape(parameterValueToString(r.parentId, "parentId")), -1)
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

type ApiGetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesInfoRequest struct {
	ctx context.Context
	ApiService *ConfigurationTypeQuestionValueInfosAPIService
	grandparentId int32
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
func (r ApiGetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesInfoRequest) Conditions(conditions string) ApiGetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesInfoRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesInfoRequest) ChildConditions(childConditions string) ApiGetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesInfoRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesInfoRequest) CustomFieldConditions(customFieldConditions string) ApiGetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesInfoRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesInfoRequest) OrderBy(orderBy string) ApiGetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesInfoRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesInfoRequest) Fields(fields string) ApiGetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesInfoRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesInfoRequest) Page(page int32) ApiGetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesInfoRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesInfoRequest) PageSize(pageSize int32) ApiGetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesInfoRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesInfoRequest) PageId(pageId int32) ApiGetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesInfoRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesInfoRequest) ClientId(clientId string) ApiGetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesInfoRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesInfoRequest) Execute() ([]ConfigurationTypeQuestionValueInfo, *http.Response, error) {
	return r.ApiService.GetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesInfoExecute(r)
}

/*
GetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesInfo Get ConfigurationTypeQuestionValueInfo

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param grandparentId ConfigurationTypeQuestionValueInfo
 @param parentId ConfigurationTypeQuestionValueInfo
 @return ApiGetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesInfoRequest
*/
func (a *ConfigurationTypeQuestionValueInfosAPIService) GetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesInfo(ctx context.Context, grandparentId int32, parentId int32) ApiGetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesInfoRequest {
	return ApiGetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesInfoRequest{
		ApiService: a,
		ctx: ctx,
		grandparentId: grandparentId,
		parentId: parentId,
	}
}

// Execute executes the request
//  @return []ConfigurationTypeQuestionValueInfo
func (a *ConfigurationTypeQuestionValueInfosAPIService) GetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesInfoExecute(r ApiGetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesInfoRequest) ([]ConfigurationTypeQuestionValueInfo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ConfigurationTypeQuestionValueInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationTypeQuestionValueInfosAPIService.GetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/configurations/types/{grandparentId}/questions/{parentId}/values/info"
	localVarPath = strings.Replace(localVarPath, "{"+"grandparentId"+"}", url.PathEscape(parameterValueToString(r.grandparentId, "grandparentId")), -1)
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

type ApiGetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesInfoCountRequest struct {
	ctx context.Context
	ApiService *ConfigurationTypeQuestionValueInfosAPIService
	grandparentId int32
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
func (r ApiGetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesInfoCountRequest) Conditions(conditions string) ApiGetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesInfoCountRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesInfoCountRequest) ChildConditions(childConditions string) ApiGetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesInfoCountRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesInfoCountRequest) CustomFieldConditions(customFieldConditions string) ApiGetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesInfoCountRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesInfoCountRequest) OrderBy(orderBy string) ApiGetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesInfoCountRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesInfoCountRequest) Fields(fields string) ApiGetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesInfoCountRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesInfoCountRequest) Page(page int32) ApiGetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesInfoCountRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesInfoCountRequest) PageSize(pageSize int32) ApiGetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesInfoCountRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesInfoCountRequest) PageId(pageId int32) ApiGetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesInfoCountRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesInfoCountRequest) ClientId(clientId string) ApiGetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesInfoCountRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesInfoCountRequest) Execute() (*Count, *http.Response, error) {
	return r.ApiService.GetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesInfoCountExecute(r)
}

/*
GetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesInfoCount Get Count of ConfigurationTypeQuestionValueInfos

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param grandparentId ConfigurationTypeQuestionValueInfo
 @param parentId ConfigurationTypeQuestionValueInfo
 @return ApiGetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesInfoCountRequest
*/
func (a *ConfigurationTypeQuestionValueInfosAPIService) GetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesInfoCount(ctx context.Context, grandparentId int32, parentId int32) ApiGetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesInfoCountRequest {
	return ApiGetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesInfoCountRequest{
		ApiService: a,
		ctx: ctx,
		grandparentId: grandparentId,
		parentId: parentId,
	}
}

// Execute executes the request
//  @return Count
func (a *ConfigurationTypeQuestionValueInfosAPIService) GetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesInfoCountExecute(r ApiGetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesInfoCountRequest) (*Count, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Count
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationTypeQuestionValueInfosAPIService.GetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesInfoCount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/configurations/types/{grandparentId}/questions/{parentId}/values/info/count"
	localVarPath = strings.Replace(localVarPath, "{"+"grandparentId"+"}", url.PathEscape(parameterValueToString(r.grandparentId, "grandparentId")), -1)
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
