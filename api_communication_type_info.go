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


// CommunicationTypeInfoAPIService CommunicationTypeInfoAPI service
type CommunicationTypeInfoAPIService service

type ApiGetCompanyCommunicationTypesByIdInfoRequest struct {
	ctx context.Context
	ApiService *CommunicationTypeInfoAPIService
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
func (r ApiGetCompanyCommunicationTypesByIdInfoRequest) Conditions(conditions string) ApiGetCompanyCommunicationTypesByIdInfoRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetCompanyCommunicationTypesByIdInfoRequest) ChildConditions(childConditions string) ApiGetCompanyCommunicationTypesByIdInfoRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetCompanyCommunicationTypesByIdInfoRequest) CustomFieldConditions(customFieldConditions string) ApiGetCompanyCommunicationTypesByIdInfoRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetCompanyCommunicationTypesByIdInfoRequest) OrderBy(orderBy string) ApiGetCompanyCommunicationTypesByIdInfoRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetCompanyCommunicationTypesByIdInfoRequest) Fields(fields string) ApiGetCompanyCommunicationTypesByIdInfoRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetCompanyCommunicationTypesByIdInfoRequest) Page(page int32) ApiGetCompanyCommunicationTypesByIdInfoRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetCompanyCommunicationTypesByIdInfoRequest) PageSize(pageSize int32) ApiGetCompanyCommunicationTypesByIdInfoRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetCompanyCommunicationTypesByIdInfoRequest) PageId(pageId int32) ApiGetCompanyCommunicationTypesByIdInfoRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetCompanyCommunicationTypesByIdInfoRequest) ClientId(clientId string) ApiGetCompanyCommunicationTypesByIdInfoRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetCompanyCommunicationTypesByIdInfoRequest) Execute() (*CommunicationTypeInfo, *http.Response, error) {
	return r.ApiService.GetCompanyCommunicationTypesByIdInfoExecute(r)
}

/*
GetCompanyCommunicationTypesByIdInfo Get CommunicationTypeInfos

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id AddressFormatInfoId
 @return ApiGetCompanyCommunicationTypesByIdInfoRequest
*/
func (a *CommunicationTypeInfoAPIService) GetCompanyCommunicationTypesByIdInfo(ctx context.Context, id int32) ApiGetCompanyCommunicationTypesByIdInfoRequest {
	return ApiGetCompanyCommunicationTypesByIdInfoRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return CommunicationTypeInfo
func (a *CommunicationTypeInfoAPIService) GetCompanyCommunicationTypesByIdInfoExecute(r ApiGetCompanyCommunicationTypesByIdInfoRequest) (*CommunicationTypeInfo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CommunicationTypeInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CommunicationTypeInfoAPIService.GetCompanyCommunicationTypesByIdInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/company/communicationTypes/{id}/info"
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

type ApiGetCompanyCommunicationTypesInfoRequest struct {
	ctx context.Context
	ApiService *CommunicationTypeInfoAPIService
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
func (r ApiGetCompanyCommunicationTypesInfoRequest) Conditions(conditions string) ApiGetCompanyCommunicationTypesInfoRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetCompanyCommunicationTypesInfoRequest) ChildConditions(childConditions string) ApiGetCompanyCommunicationTypesInfoRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetCompanyCommunicationTypesInfoRequest) CustomFieldConditions(customFieldConditions string) ApiGetCompanyCommunicationTypesInfoRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetCompanyCommunicationTypesInfoRequest) OrderBy(orderBy string) ApiGetCompanyCommunicationTypesInfoRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetCompanyCommunicationTypesInfoRequest) Fields(fields string) ApiGetCompanyCommunicationTypesInfoRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetCompanyCommunicationTypesInfoRequest) Page(page int32) ApiGetCompanyCommunicationTypesInfoRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetCompanyCommunicationTypesInfoRequest) PageSize(pageSize int32) ApiGetCompanyCommunicationTypesInfoRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetCompanyCommunicationTypesInfoRequest) PageId(pageId int32) ApiGetCompanyCommunicationTypesInfoRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetCompanyCommunicationTypesInfoRequest) ClientId(clientId string) ApiGetCompanyCommunicationTypesInfoRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetCompanyCommunicationTypesInfoRequest) Execute() ([]CommunicationTypeInfo, *http.Response, error) {
	return r.ApiService.GetCompanyCommunicationTypesInfoExecute(r)
}

/*
GetCompanyCommunicationTypesInfo Get List of CommunicationTypeInfos

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetCompanyCommunicationTypesInfoRequest
*/
func (a *CommunicationTypeInfoAPIService) GetCompanyCommunicationTypesInfo(ctx context.Context) ApiGetCompanyCommunicationTypesInfoRequest {
	return ApiGetCompanyCommunicationTypesInfoRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []CommunicationTypeInfo
func (a *CommunicationTypeInfoAPIService) GetCompanyCommunicationTypesInfoExecute(r ApiGetCompanyCommunicationTypesInfoRequest) ([]CommunicationTypeInfo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []CommunicationTypeInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CommunicationTypeInfoAPIService.GetCompanyCommunicationTypesInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/company/communicationTypes/info"

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

type ApiGetCompanyCommunicationTypesInfoCountRequest struct {
	ctx context.Context
	ApiService *CommunicationTypeInfoAPIService
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
func (r ApiGetCompanyCommunicationTypesInfoCountRequest) Conditions(conditions string) ApiGetCompanyCommunicationTypesInfoCountRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetCompanyCommunicationTypesInfoCountRequest) ChildConditions(childConditions string) ApiGetCompanyCommunicationTypesInfoCountRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetCompanyCommunicationTypesInfoCountRequest) CustomFieldConditions(customFieldConditions string) ApiGetCompanyCommunicationTypesInfoCountRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetCompanyCommunicationTypesInfoCountRequest) OrderBy(orderBy string) ApiGetCompanyCommunicationTypesInfoCountRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetCompanyCommunicationTypesInfoCountRequest) Fields(fields string) ApiGetCompanyCommunicationTypesInfoCountRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetCompanyCommunicationTypesInfoCountRequest) Page(page int32) ApiGetCompanyCommunicationTypesInfoCountRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetCompanyCommunicationTypesInfoCountRequest) PageSize(pageSize int32) ApiGetCompanyCommunicationTypesInfoCountRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetCompanyCommunicationTypesInfoCountRequest) PageId(pageId int32) ApiGetCompanyCommunicationTypesInfoCountRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetCompanyCommunicationTypesInfoCountRequest) ClientId(clientId string) ApiGetCompanyCommunicationTypesInfoCountRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetCompanyCommunicationTypesInfoCountRequest) Execute() (*Count, *http.Response, error) {
	return r.ApiService.GetCompanyCommunicationTypesInfoCountExecute(r)
}

/*
GetCompanyCommunicationTypesInfoCount Get Count of CommunicationTypeInfos

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetCompanyCommunicationTypesInfoCountRequest
*/
func (a *CommunicationTypeInfoAPIService) GetCompanyCommunicationTypesInfoCount(ctx context.Context) ApiGetCompanyCommunicationTypesInfoCountRequest {
	return ApiGetCompanyCommunicationTypesInfoCountRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Count
func (a *CommunicationTypeInfoAPIService) GetCompanyCommunicationTypesInfoCountExecute(r ApiGetCompanyCommunicationTypesInfoCountRequest) (*Count, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Count
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CommunicationTypeInfoAPIService.GetCompanyCommunicationTypesInfoCount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/company/communicationTypes/info/count"

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
