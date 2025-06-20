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


// M365ContactSyncPropertiesAPIService M365ContactSyncPropertiesAPI service
type M365ContactSyncPropertiesAPIService service

type ApiDeleteCompanyM365contactsyncPropertyRequest struct {
	ctx context.Context
	ApiService *M365ContactSyncPropertiesAPIService
	clientId *string
}

// 
func (r ApiDeleteCompanyM365contactsyncPropertyRequest) ClientId(clientId string) ApiDeleteCompanyM365contactsyncPropertyRequest {
	r.clientId = &clientId
	return r
}

func (r ApiDeleteCompanyM365contactsyncPropertyRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteCompanyM365contactsyncPropertyExecute(r)
}

/*
DeleteCompanyM365contactsyncProperty Delete M365ContactSyncProperty

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDeleteCompanyM365contactsyncPropertyRequest
*/
func (a *M365ContactSyncPropertiesAPIService) DeleteCompanyM365contactsyncProperty(ctx context.Context) ApiDeleteCompanyM365contactsyncPropertyRequest {
	return ApiDeleteCompanyM365contactsyncPropertyRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *M365ContactSyncPropertiesAPIService) DeleteCompanyM365contactsyncPropertyExecute(r ApiDeleteCompanyM365contactsyncPropertyRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "M365ContactSyncPropertiesAPIService.DeleteCompanyM365contactsyncProperty")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/company/m365contactsync/property/"

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

type ApiGetCompanyM365contactsyncByIdPropertyRequest struct {
	ctx context.Context
	ApiService *M365ContactSyncPropertiesAPIService
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
func (r ApiGetCompanyM365contactsyncByIdPropertyRequest) Conditions(conditions string) ApiGetCompanyM365contactsyncByIdPropertyRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetCompanyM365contactsyncByIdPropertyRequest) ChildConditions(childConditions string) ApiGetCompanyM365contactsyncByIdPropertyRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetCompanyM365contactsyncByIdPropertyRequest) CustomFieldConditions(customFieldConditions string) ApiGetCompanyM365contactsyncByIdPropertyRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetCompanyM365contactsyncByIdPropertyRequest) OrderBy(orderBy string) ApiGetCompanyM365contactsyncByIdPropertyRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetCompanyM365contactsyncByIdPropertyRequest) Fields(fields string) ApiGetCompanyM365contactsyncByIdPropertyRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetCompanyM365contactsyncByIdPropertyRequest) Page(page int32) ApiGetCompanyM365contactsyncByIdPropertyRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetCompanyM365contactsyncByIdPropertyRequest) PageSize(pageSize int32) ApiGetCompanyM365contactsyncByIdPropertyRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetCompanyM365contactsyncByIdPropertyRequest) PageId(pageId int32) ApiGetCompanyM365contactsyncByIdPropertyRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetCompanyM365contactsyncByIdPropertyRequest) ClientId(clientId string) ApiGetCompanyM365contactsyncByIdPropertyRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetCompanyM365contactsyncByIdPropertyRequest) Execute() (*M365ContactSyncProperty, *http.Response, error) {
	return r.ApiService.GetCompanyM365contactsyncByIdPropertyExecute(r)
}

/*
GetCompanyM365contactsyncByIdProperty Get M365ContactSyncProperties

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id M365ContactSyncPropertyId
 @return ApiGetCompanyM365contactsyncByIdPropertyRequest
*/
func (a *M365ContactSyncPropertiesAPIService) GetCompanyM365contactsyncByIdProperty(ctx context.Context, id int32) ApiGetCompanyM365contactsyncByIdPropertyRequest {
	return ApiGetCompanyM365contactsyncByIdPropertyRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return M365ContactSyncProperty
func (a *M365ContactSyncPropertiesAPIService) GetCompanyM365contactsyncByIdPropertyExecute(r ApiGetCompanyM365contactsyncByIdPropertyRequest) (*M365ContactSyncProperty, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *M365ContactSyncProperty
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "M365ContactSyncPropertiesAPIService.GetCompanyM365contactsyncByIdProperty")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/company/m365contactsync/{id}/property"
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

type ApiGetCompanyM365contactsyncPropertyCountRequest struct {
	ctx context.Context
	ApiService *M365ContactSyncPropertiesAPIService
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
func (r ApiGetCompanyM365contactsyncPropertyCountRequest) Conditions(conditions string) ApiGetCompanyM365contactsyncPropertyCountRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetCompanyM365contactsyncPropertyCountRequest) ChildConditions(childConditions string) ApiGetCompanyM365contactsyncPropertyCountRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetCompanyM365contactsyncPropertyCountRequest) CustomFieldConditions(customFieldConditions string) ApiGetCompanyM365contactsyncPropertyCountRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetCompanyM365contactsyncPropertyCountRequest) OrderBy(orderBy string) ApiGetCompanyM365contactsyncPropertyCountRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetCompanyM365contactsyncPropertyCountRequest) Fields(fields string) ApiGetCompanyM365contactsyncPropertyCountRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetCompanyM365contactsyncPropertyCountRequest) Page(page int32) ApiGetCompanyM365contactsyncPropertyCountRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetCompanyM365contactsyncPropertyCountRequest) PageSize(pageSize int32) ApiGetCompanyM365contactsyncPropertyCountRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetCompanyM365contactsyncPropertyCountRequest) PageId(pageId int32) ApiGetCompanyM365contactsyncPropertyCountRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetCompanyM365contactsyncPropertyCountRequest) ClientId(clientId string) ApiGetCompanyM365contactsyncPropertyCountRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetCompanyM365contactsyncPropertyCountRequest) Execute() (*Count, *http.Response, error) {
	return r.ApiService.GetCompanyM365contactsyncPropertyCountExecute(r)
}

/*
GetCompanyM365contactsyncPropertyCount Get Count of M365ContactSyncProperty

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetCompanyM365contactsyncPropertyCountRequest
*/
func (a *M365ContactSyncPropertiesAPIService) GetCompanyM365contactsyncPropertyCount(ctx context.Context) ApiGetCompanyM365contactsyncPropertyCountRequest {
	return ApiGetCompanyM365contactsyncPropertyCountRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Count
func (a *M365ContactSyncPropertiesAPIService) GetCompanyM365contactsyncPropertyCountExecute(r ApiGetCompanyM365contactsyncPropertyCountRequest) (*Count, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Count
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "M365ContactSyncPropertiesAPIService.GetCompanyM365contactsyncPropertyCount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/company/m365contactsync/property/count"

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

type ApiGetCompanyM365contactsyncPropertyExcludedRequest struct {
	ctx context.Context
	ApiService *M365ContactSyncPropertiesAPIService
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
func (r ApiGetCompanyM365contactsyncPropertyExcludedRequest) Conditions(conditions string) ApiGetCompanyM365contactsyncPropertyExcludedRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetCompanyM365contactsyncPropertyExcludedRequest) ChildConditions(childConditions string) ApiGetCompanyM365contactsyncPropertyExcludedRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetCompanyM365contactsyncPropertyExcludedRequest) CustomFieldConditions(customFieldConditions string) ApiGetCompanyM365contactsyncPropertyExcludedRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetCompanyM365contactsyncPropertyExcludedRequest) OrderBy(orderBy string) ApiGetCompanyM365contactsyncPropertyExcludedRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetCompanyM365contactsyncPropertyExcludedRequest) Fields(fields string) ApiGetCompanyM365contactsyncPropertyExcludedRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetCompanyM365contactsyncPropertyExcludedRequest) Page(page int32) ApiGetCompanyM365contactsyncPropertyExcludedRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetCompanyM365contactsyncPropertyExcludedRequest) PageSize(pageSize int32) ApiGetCompanyM365contactsyncPropertyExcludedRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetCompanyM365contactsyncPropertyExcludedRequest) PageId(pageId int32) ApiGetCompanyM365contactsyncPropertyExcludedRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetCompanyM365contactsyncPropertyExcludedRequest) ClientId(clientId string) ApiGetCompanyM365contactsyncPropertyExcludedRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetCompanyM365contactsyncPropertyExcludedRequest) Execute() ([]M365ContactSyncProperty, *http.Response, error) {
	return r.ApiService.GetCompanyM365contactsyncPropertyExcludedExecute(r)
}

/*
GetCompanyM365contactsyncPropertyExcluded Get List of M365ContactSyncPropertiesExcluded

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id companyRecId
 @return ApiGetCompanyM365contactsyncPropertyExcludedRequest
*/
func (a *M365ContactSyncPropertiesAPIService) GetCompanyM365contactsyncPropertyExcluded(ctx context.Context, id int32) ApiGetCompanyM365contactsyncPropertyExcludedRequest {
	return ApiGetCompanyM365contactsyncPropertyExcludedRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return []M365ContactSyncProperty
func (a *M365ContactSyncPropertiesAPIService) GetCompanyM365contactsyncPropertyExcludedExecute(r ApiGetCompanyM365contactsyncPropertyExcludedRequest) ([]M365ContactSyncProperty, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []M365ContactSyncProperty
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "M365ContactSyncPropertiesAPIService.GetCompanyM365contactsyncPropertyExcluded")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/company/m365contactsync/property/excluded"
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

type ApiGetCompanyM365contactsyncPropertyIncludedRequest struct {
	ctx context.Context
	ApiService *M365ContactSyncPropertiesAPIService
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
func (r ApiGetCompanyM365contactsyncPropertyIncludedRequest) Conditions(conditions string) ApiGetCompanyM365contactsyncPropertyIncludedRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetCompanyM365contactsyncPropertyIncludedRequest) ChildConditions(childConditions string) ApiGetCompanyM365contactsyncPropertyIncludedRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetCompanyM365contactsyncPropertyIncludedRequest) CustomFieldConditions(customFieldConditions string) ApiGetCompanyM365contactsyncPropertyIncludedRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetCompanyM365contactsyncPropertyIncludedRequest) OrderBy(orderBy string) ApiGetCompanyM365contactsyncPropertyIncludedRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetCompanyM365contactsyncPropertyIncludedRequest) Fields(fields string) ApiGetCompanyM365contactsyncPropertyIncludedRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetCompanyM365contactsyncPropertyIncludedRequest) Page(page int32) ApiGetCompanyM365contactsyncPropertyIncludedRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetCompanyM365contactsyncPropertyIncludedRequest) PageSize(pageSize int32) ApiGetCompanyM365contactsyncPropertyIncludedRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetCompanyM365contactsyncPropertyIncludedRequest) PageId(pageId int32) ApiGetCompanyM365contactsyncPropertyIncludedRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetCompanyM365contactsyncPropertyIncludedRequest) ClientId(clientId string) ApiGetCompanyM365contactsyncPropertyIncludedRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetCompanyM365contactsyncPropertyIncludedRequest) Execute() ([]M365ContactSyncProperty, *http.Response, error) {
	return r.ApiService.GetCompanyM365contactsyncPropertyIncludedExecute(r)
}

/*
GetCompanyM365contactsyncPropertyIncluded Get List of M365ContactSyncPropertiesIncluded

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id companyRecId
 @return ApiGetCompanyM365contactsyncPropertyIncludedRequest
*/
func (a *M365ContactSyncPropertiesAPIService) GetCompanyM365contactsyncPropertyIncluded(ctx context.Context, id int32) ApiGetCompanyM365contactsyncPropertyIncludedRequest {
	return ApiGetCompanyM365contactsyncPropertyIncludedRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return []M365ContactSyncProperty
func (a *M365ContactSyncPropertiesAPIService) GetCompanyM365contactsyncPropertyIncludedExecute(r ApiGetCompanyM365contactsyncPropertyIncludedRequest) ([]M365ContactSyncProperty, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []M365ContactSyncProperty
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "M365ContactSyncPropertiesAPIService.GetCompanyM365contactsyncPropertyIncluded")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/company/m365contactsync/property/included"
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

type ApiPostCompanyM365contactsyncPropertyRequest struct {
	ctx context.Context
	ApiService *M365ContactSyncPropertiesAPIService
	m365ContactSyncProperty *M365ContactSyncProperty
	clientId *string
}

// country
func (r ApiPostCompanyM365contactsyncPropertyRequest) M365ContactSyncProperty(m365ContactSyncProperty M365ContactSyncProperty) ApiPostCompanyM365contactsyncPropertyRequest {
	r.m365ContactSyncProperty = &m365ContactSyncProperty
	return r
}

// 
func (r ApiPostCompanyM365contactsyncPropertyRequest) ClientId(clientId string) ApiPostCompanyM365contactsyncPropertyRequest {
	r.clientId = &clientId
	return r
}

func (r ApiPostCompanyM365contactsyncPropertyRequest) Execute() (*M365ContactSyncProperty, *http.Response, error) {
	return r.ApiService.PostCompanyM365contactsyncPropertyExecute(r)
}

/*
PostCompanyM365contactsyncProperty Create M365ContactSyncProperty

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPostCompanyM365contactsyncPropertyRequest
*/
func (a *M365ContactSyncPropertiesAPIService) PostCompanyM365contactsyncProperty(ctx context.Context) ApiPostCompanyM365contactsyncPropertyRequest {
	return ApiPostCompanyM365contactsyncPropertyRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return M365ContactSyncProperty
func (a *M365ContactSyncPropertiesAPIService) PostCompanyM365contactsyncPropertyExecute(r ApiPostCompanyM365contactsyncPropertyRequest) (*M365ContactSyncProperty, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *M365ContactSyncProperty
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "M365ContactSyncPropertiesAPIService.PostCompanyM365contactsyncProperty")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/company/m365contactsync/property"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.m365ContactSyncProperty == nil {
		return localVarReturnValue, nil, reportError("m365ContactSyncProperty is required and must be specified")
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
	localVarPostBody = r.m365ContactSyncProperty
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
