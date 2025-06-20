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


// MarketplaceImportsAPIService MarketplaceImportsAPI service
type MarketplaceImportsAPIService service

type ApiGetSystemMarketplaceimportGetdefinitionByIdRequest struct {
	ctx context.Context
	ApiService *MarketplaceImportsAPIService
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
func (r ApiGetSystemMarketplaceimportGetdefinitionByIdRequest) Conditions(conditions string) ApiGetSystemMarketplaceimportGetdefinitionByIdRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetSystemMarketplaceimportGetdefinitionByIdRequest) ChildConditions(childConditions string) ApiGetSystemMarketplaceimportGetdefinitionByIdRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetSystemMarketplaceimportGetdefinitionByIdRequest) CustomFieldConditions(customFieldConditions string) ApiGetSystemMarketplaceimportGetdefinitionByIdRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetSystemMarketplaceimportGetdefinitionByIdRequest) OrderBy(orderBy string) ApiGetSystemMarketplaceimportGetdefinitionByIdRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetSystemMarketplaceimportGetdefinitionByIdRequest) Fields(fields string) ApiGetSystemMarketplaceimportGetdefinitionByIdRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetSystemMarketplaceimportGetdefinitionByIdRequest) Page(page int32) ApiGetSystemMarketplaceimportGetdefinitionByIdRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetSystemMarketplaceimportGetdefinitionByIdRequest) PageSize(pageSize int32) ApiGetSystemMarketplaceimportGetdefinitionByIdRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetSystemMarketplaceimportGetdefinitionByIdRequest) PageId(pageId int32) ApiGetSystemMarketplaceimportGetdefinitionByIdRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetSystemMarketplaceimportGetdefinitionByIdRequest) ClientId(clientId string) ApiGetSystemMarketplaceimportGetdefinitionByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetSystemMarketplaceimportGetdefinitionByIdRequest) Execute() (*MarketplaceImport, *http.Response, error) {
	return r.ApiService.GetSystemMarketplaceimportGetdefinitionByIdExecute(r)
}

/*
GetSystemMarketplaceimportGetdefinitionById Get MarketplaceImport

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id getdefinitionId
 @return ApiGetSystemMarketplaceimportGetdefinitionByIdRequest
*/
func (a *MarketplaceImportsAPIService) GetSystemMarketplaceimportGetdefinitionById(ctx context.Context, id int32) ApiGetSystemMarketplaceimportGetdefinitionByIdRequest {
	return ApiGetSystemMarketplaceimportGetdefinitionByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return MarketplaceImport
func (a *MarketplaceImportsAPIService) GetSystemMarketplaceimportGetdefinitionByIdExecute(r ApiGetSystemMarketplaceimportGetdefinitionByIdRequest) (*MarketplaceImport, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MarketplaceImport
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MarketplaceImportsAPIService.GetSystemMarketplaceimportGetdefinitionById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/marketplaceimport/getdefinition/{id}"
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

type ApiPostSystemMarketplaceimportImportRequest struct {
	ctx context.Context
	ApiService *MarketplaceImportsAPIService
	marketplaceImport *MarketplaceImport
	clientId *string
}

// marketplaceImport
func (r ApiPostSystemMarketplaceimportImportRequest) MarketplaceImport(marketplaceImport MarketplaceImport) ApiPostSystemMarketplaceimportImportRequest {
	r.marketplaceImport = &marketplaceImport
	return r
}

// 
func (r ApiPostSystemMarketplaceimportImportRequest) ClientId(clientId string) ApiPostSystemMarketplaceimportImportRequest {
	r.clientId = &clientId
	return r
}

func (r ApiPostSystemMarketplaceimportImportRequest) Execute() (*MarketplaceImport, *http.Response, error) {
	return r.ApiService.PostSystemMarketplaceimportImportExecute(r)
}

/*
PostSystemMarketplaceimportImport Post MarketplaceImport

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPostSystemMarketplaceimportImportRequest
*/
func (a *MarketplaceImportsAPIService) PostSystemMarketplaceimportImport(ctx context.Context) ApiPostSystemMarketplaceimportImportRequest {
	return ApiPostSystemMarketplaceimportImportRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MarketplaceImport
func (a *MarketplaceImportsAPIService) PostSystemMarketplaceimportImportExecute(r ApiPostSystemMarketplaceimportImportRequest) (*MarketplaceImport, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MarketplaceImport
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MarketplaceImportsAPIService.PostSystemMarketplaceimportImport")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/marketplaceimport/import"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.marketplaceImport == nil {
		return localVarReturnValue, nil, reportError("marketplaceImport is required and must be specified")
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
	localVarPostBody = r.marketplaceImport
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
