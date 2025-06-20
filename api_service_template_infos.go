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


// ServiceTemplateInfosAPIService ServiceTemplateInfosAPI service
type ServiceTemplateInfosAPIService service

type ApiGetServiceTemplatesByIdInfoRequest struct {
	ctx context.Context
	ApiService *ServiceTemplateInfosAPIService
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
func (r ApiGetServiceTemplatesByIdInfoRequest) Conditions(conditions string) ApiGetServiceTemplatesByIdInfoRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetServiceTemplatesByIdInfoRequest) ChildConditions(childConditions string) ApiGetServiceTemplatesByIdInfoRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetServiceTemplatesByIdInfoRequest) CustomFieldConditions(customFieldConditions string) ApiGetServiceTemplatesByIdInfoRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetServiceTemplatesByIdInfoRequest) OrderBy(orderBy string) ApiGetServiceTemplatesByIdInfoRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetServiceTemplatesByIdInfoRequest) Fields(fields string) ApiGetServiceTemplatesByIdInfoRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetServiceTemplatesByIdInfoRequest) Page(page int32) ApiGetServiceTemplatesByIdInfoRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetServiceTemplatesByIdInfoRequest) PageSize(pageSize int32) ApiGetServiceTemplatesByIdInfoRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetServiceTemplatesByIdInfoRequest) PageId(pageId int32) ApiGetServiceTemplatesByIdInfoRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetServiceTemplatesByIdInfoRequest) ClientId(clientId string) ApiGetServiceTemplatesByIdInfoRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetServiceTemplatesByIdInfoRequest) Execute() (*ServiceTemplateInfo, *http.Response, error) {
	return r.ApiService.GetServiceTemplatesByIdInfoExecute(r)
}

/*
GetServiceTemplatesByIdInfo Get ServiceTemplateInfos

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id ServiceTemplateInfoId
 @return ApiGetServiceTemplatesByIdInfoRequest
*/
func (a *ServiceTemplateInfosAPIService) GetServiceTemplatesByIdInfo(ctx context.Context, id int32) ApiGetServiceTemplatesByIdInfoRequest {
	return ApiGetServiceTemplatesByIdInfoRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ServiceTemplateInfo
func (a *ServiceTemplateInfosAPIService) GetServiceTemplatesByIdInfoExecute(r ApiGetServiceTemplatesByIdInfoRequest) (*ServiceTemplateInfo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ServiceTemplateInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceTemplateInfosAPIService.GetServiceTemplatesByIdInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/templates/{id}/info"
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

type ApiGetServiceTemplatesInfoRequest struct {
	ctx context.Context
	ApiService *ServiceTemplateInfosAPIService
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
func (r ApiGetServiceTemplatesInfoRequest) Conditions(conditions string) ApiGetServiceTemplatesInfoRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetServiceTemplatesInfoRequest) ChildConditions(childConditions string) ApiGetServiceTemplatesInfoRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetServiceTemplatesInfoRequest) CustomFieldConditions(customFieldConditions string) ApiGetServiceTemplatesInfoRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetServiceTemplatesInfoRequest) OrderBy(orderBy string) ApiGetServiceTemplatesInfoRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetServiceTemplatesInfoRequest) Fields(fields string) ApiGetServiceTemplatesInfoRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetServiceTemplatesInfoRequest) Page(page int32) ApiGetServiceTemplatesInfoRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetServiceTemplatesInfoRequest) PageSize(pageSize int32) ApiGetServiceTemplatesInfoRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetServiceTemplatesInfoRequest) PageId(pageId int32) ApiGetServiceTemplatesInfoRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetServiceTemplatesInfoRequest) ClientId(clientId string) ApiGetServiceTemplatesInfoRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetServiceTemplatesInfoRequest) Execute() ([]ServiceTemplateInfo, *http.Response, error) {
	return r.ApiService.GetServiceTemplatesInfoExecute(r)
}

/*
GetServiceTemplatesInfo Get List of ServiceTemplateInfos

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetServiceTemplatesInfoRequest
*/
func (a *ServiceTemplateInfosAPIService) GetServiceTemplatesInfo(ctx context.Context) ApiGetServiceTemplatesInfoRequest {
	return ApiGetServiceTemplatesInfoRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []ServiceTemplateInfo
func (a *ServiceTemplateInfosAPIService) GetServiceTemplatesInfoExecute(r ApiGetServiceTemplatesInfoRequest) ([]ServiceTemplateInfo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ServiceTemplateInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceTemplateInfosAPIService.GetServiceTemplatesInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/templates/info"

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

type ApiGetServiceTemplatesInfoCountRequest struct {
	ctx context.Context
	ApiService *ServiceTemplateInfosAPIService
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
func (r ApiGetServiceTemplatesInfoCountRequest) Conditions(conditions string) ApiGetServiceTemplatesInfoCountRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetServiceTemplatesInfoCountRequest) ChildConditions(childConditions string) ApiGetServiceTemplatesInfoCountRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetServiceTemplatesInfoCountRequest) CustomFieldConditions(customFieldConditions string) ApiGetServiceTemplatesInfoCountRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetServiceTemplatesInfoCountRequest) OrderBy(orderBy string) ApiGetServiceTemplatesInfoCountRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetServiceTemplatesInfoCountRequest) Fields(fields string) ApiGetServiceTemplatesInfoCountRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetServiceTemplatesInfoCountRequest) Page(page int32) ApiGetServiceTemplatesInfoCountRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetServiceTemplatesInfoCountRequest) PageSize(pageSize int32) ApiGetServiceTemplatesInfoCountRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetServiceTemplatesInfoCountRequest) PageId(pageId int32) ApiGetServiceTemplatesInfoCountRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetServiceTemplatesInfoCountRequest) ClientId(clientId string) ApiGetServiceTemplatesInfoCountRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetServiceTemplatesInfoCountRequest) Execute() (*Count, *http.Response, error) {
	return r.ApiService.GetServiceTemplatesInfoCountExecute(r)
}

/*
GetServiceTemplatesInfoCount Get Count of ServiceTemplateInfo

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetServiceTemplatesInfoCountRequest
*/
func (a *ServiceTemplateInfosAPIService) GetServiceTemplatesInfoCount(ctx context.Context) ApiGetServiceTemplatesInfoCountRequest {
	return ApiGetServiceTemplatesInfoCountRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Count
func (a *ServiceTemplateInfosAPIService) GetServiceTemplatesInfoCountExecute(r ApiGetServiceTemplatesInfoCountRequest) (*Count, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Count
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceTemplateInfosAPIService.GetServiceTemplatesInfoCount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/templates/info/count"

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
