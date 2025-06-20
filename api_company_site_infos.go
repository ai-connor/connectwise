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


// CompanySiteInfosAPIService CompanySiteInfosAPI service
type CompanySiteInfosAPIService service

type ApiGetCompanyCompaniesByParentIdSitesByIdInfoRequest struct {
	ctx context.Context
	ApiService *CompanySiteInfosAPIService
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
func (r ApiGetCompanyCompaniesByParentIdSitesByIdInfoRequest) Conditions(conditions string) ApiGetCompanyCompaniesByParentIdSitesByIdInfoRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetCompanyCompaniesByParentIdSitesByIdInfoRequest) ChildConditions(childConditions string) ApiGetCompanyCompaniesByParentIdSitesByIdInfoRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetCompanyCompaniesByParentIdSitesByIdInfoRequest) CustomFieldConditions(customFieldConditions string) ApiGetCompanyCompaniesByParentIdSitesByIdInfoRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetCompanyCompaniesByParentIdSitesByIdInfoRequest) OrderBy(orderBy string) ApiGetCompanyCompaniesByParentIdSitesByIdInfoRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetCompanyCompaniesByParentIdSitesByIdInfoRequest) Fields(fields string) ApiGetCompanyCompaniesByParentIdSitesByIdInfoRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetCompanyCompaniesByParentIdSitesByIdInfoRequest) Page(page int32) ApiGetCompanyCompaniesByParentIdSitesByIdInfoRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetCompanyCompaniesByParentIdSitesByIdInfoRequest) PageSize(pageSize int32) ApiGetCompanyCompaniesByParentIdSitesByIdInfoRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetCompanyCompaniesByParentIdSitesByIdInfoRequest) PageId(pageId int32) ApiGetCompanyCompaniesByParentIdSitesByIdInfoRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetCompanyCompaniesByParentIdSitesByIdInfoRequest) ClientId(clientId string) ApiGetCompanyCompaniesByParentIdSitesByIdInfoRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetCompanyCompaniesByParentIdSitesByIdInfoRequest) Execute() (*CompanySiteInfo, *http.Response, error) {
	return r.ApiService.GetCompanyCompaniesByParentIdSitesByIdInfoExecute(r)
}

/*
GetCompanyCompaniesByParentIdSitesByIdInfo Get CompanySiteInfos

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id siteId
 @param parentId companyId
 @return ApiGetCompanyCompaniesByParentIdSitesByIdInfoRequest
*/
func (a *CompanySiteInfosAPIService) GetCompanyCompaniesByParentIdSitesByIdInfo(ctx context.Context, id int32, parentId int32) ApiGetCompanyCompaniesByParentIdSitesByIdInfoRequest {
	return ApiGetCompanyCompaniesByParentIdSitesByIdInfoRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
		parentId: parentId,
	}
}

// Execute executes the request
//  @return CompanySiteInfo
func (a *CompanySiteInfosAPIService) GetCompanyCompaniesByParentIdSitesByIdInfoExecute(r ApiGetCompanyCompaniesByParentIdSitesByIdInfoRequest) (*CompanySiteInfo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CompanySiteInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CompanySiteInfosAPIService.GetCompanyCompaniesByParentIdSitesByIdInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/company/companies/{parentId}/sites/{id}/info"
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

type ApiGetCompanyCompaniesByParentIdSitesInfoRequest struct {
	ctx context.Context
	ApiService *CompanySiteInfosAPIService
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
func (r ApiGetCompanyCompaniesByParentIdSitesInfoRequest) Conditions(conditions string) ApiGetCompanyCompaniesByParentIdSitesInfoRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetCompanyCompaniesByParentIdSitesInfoRequest) ChildConditions(childConditions string) ApiGetCompanyCompaniesByParentIdSitesInfoRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetCompanyCompaniesByParentIdSitesInfoRequest) CustomFieldConditions(customFieldConditions string) ApiGetCompanyCompaniesByParentIdSitesInfoRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetCompanyCompaniesByParentIdSitesInfoRequest) OrderBy(orderBy string) ApiGetCompanyCompaniesByParentIdSitesInfoRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetCompanyCompaniesByParentIdSitesInfoRequest) Fields(fields string) ApiGetCompanyCompaniesByParentIdSitesInfoRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetCompanyCompaniesByParentIdSitesInfoRequest) Page(page int32) ApiGetCompanyCompaniesByParentIdSitesInfoRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetCompanyCompaniesByParentIdSitesInfoRequest) PageSize(pageSize int32) ApiGetCompanyCompaniesByParentIdSitesInfoRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetCompanyCompaniesByParentIdSitesInfoRequest) PageId(pageId int32) ApiGetCompanyCompaniesByParentIdSitesInfoRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetCompanyCompaniesByParentIdSitesInfoRequest) ClientId(clientId string) ApiGetCompanyCompaniesByParentIdSitesInfoRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetCompanyCompaniesByParentIdSitesInfoRequest) Execute() ([]CompanySiteInfo, *http.Response, error) {
	return r.ApiService.GetCompanyCompaniesByParentIdSitesInfoExecute(r)
}

/*
GetCompanyCompaniesByParentIdSitesInfo Get List of CompanySiteInfos

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param parentId companyId
 @return ApiGetCompanyCompaniesByParentIdSitesInfoRequest
*/
func (a *CompanySiteInfosAPIService) GetCompanyCompaniesByParentIdSitesInfo(ctx context.Context, parentId int32) ApiGetCompanyCompaniesByParentIdSitesInfoRequest {
	return ApiGetCompanyCompaniesByParentIdSitesInfoRequest{
		ApiService: a,
		ctx: ctx,
		parentId: parentId,
	}
}

// Execute executes the request
//  @return []CompanySiteInfo
func (a *CompanySiteInfosAPIService) GetCompanyCompaniesByParentIdSitesInfoExecute(r ApiGetCompanyCompaniesByParentIdSitesInfoRequest) ([]CompanySiteInfo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []CompanySiteInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CompanySiteInfosAPIService.GetCompanyCompaniesByParentIdSitesInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/company/companies/{parentId}/sites/info"
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

type ApiGetCompanyCompaniesByParentIdSitesInfoCountRequest struct {
	ctx context.Context
	ApiService *CompanySiteInfosAPIService
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
func (r ApiGetCompanyCompaniesByParentIdSitesInfoCountRequest) Conditions(conditions string) ApiGetCompanyCompaniesByParentIdSitesInfoCountRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetCompanyCompaniesByParentIdSitesInfoCountRequest) ChildConditions(childConditions string) ApiGetCompanyCompaniesByParentIdSitesInfoCountRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetCompanyCompaniesByParentIdSitesInfoCountRequest) CustomFieldConditions(customFieldConditions string) ApiGetCompanyCompaniesByParentIdSitesInfoCountRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetCompanyCompaniesByParentIdSitesInfoCountRequest) OrderBy(orderBy string) ApiGetCompanyCompaniesByParentIdSitesInfoCountRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetCompanyCompaniesByParentIdSitesInfoCountRequest) Fields(fields string) ApiGetCompanyCompaniesByParentIdSitesInfoCountRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetCompanyCompaniesByParentIdSitesInfoCountRequest) Page(page int32) ApiGetCompanyCompaniesByParentIdSitesInfoCountRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetCompanyCompaniesByParentIdSitesInfoCountRequest) PageSize(pageSize int32) ApiGetCompanyCompaniesByParentIdSitesInfoCountRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetCompanyCompaniesByParentIdSitesInfoCountRequest) PageId(pageId int32) ApiGetCompanyCompaniesByParentIdSitesInfoCountRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetCompanyCompaniesByParentIdSitesInfoCountRequest) ClientId(clientId string) ApiGetCompanyCompaniesByParentIdSitesInfoCountRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetCompanyCompaniesByParentIdSitesInfoCountRequest) Execute() (*Count, *http.Response, error) {
	return r.ApiService.GetCompanyCompaniesByParentIdSitesInfoCountExecute(r)
}

/*
GetCompanyCompaniesByParentIdSitesInfoCount Get Count of CompanySite

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param parentId companyId
 @return ApiGetCompanyCompaniesByParentIdSitesInfoCountRequest
*/
func (a *CompanySiteInfosAPIService) GetCompanyCompaniesByParentIdSitesInfoCount(ctx context.Context, parentId int32) ApiGetCompanyCompaniesByParentIdSitesInfoCountRequest {
	return ApiGetCompanyCompaniesByParentIdSitesInfoCountRequest{
		ApiService: a,
		ctx: ctx,
		parentId: parentId,
	}
}

// Execute executes the request
//  @return Count
func (a *CompanySiteInfosAPIService) GetCompanyCompaniesByParentIdSitesInfoCountExecute(r ApiGetCompanyCompaniesByParentIdSitesInfoCountRequest) (*Count, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Count
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CompanySiteInfosAPIService.GetCompanyCompaniesByParentIdSitesInfoCount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/company/companies/{parentId}/sites/info/count"
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
