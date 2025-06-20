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
)


// AuditTrailAPIService AuditTrailAPI service
type AuditTrailAPIService service

type ApiGetSystemAudittrailRequest struct {
	ctx context.Context
	ApiService *AuditTrailAPIService
	type_ *string
	id *int32
	deviceIdentifier *string
	xrefRecId *int32
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

// type
func (r ApiGetSystemAudittrailRequest) Type_(type_ string) ApiGetSystemAudittrailRequest {
	r.type_ = &type_
	return r
}

// id
func (r ApiGetSystemAudittrailRequest) Id(id int32) ApiGetSystemAudittrailRequest {
	r.id = &id
	return r
}

// deviceIdentifier
func (r ApiGetSystemAudittrailRequest) DeviceIdentifier(deviceIdentifier string) ApiGetSystemAudittrailRequest {
	r.deviceIdentifier = &deviceIdentifier
	return r
}

// xrefRecId
func (r ApiGetSystemAudittrailRequest) XrefRecId(xrefRecId int32) ApiGetSystemAudittrailRequest {
	r.xrefRecId = &xrefRecId
	return r
}

// 
func (r ApiGetSystemAudittrailRequest) Conditions(conditions string) ApiGetSystemAudittrailRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetSystemAudittrailRequest) ChildConditions(childConditions string) ApiGetSystemAudittrailRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetSystemAudittrailRequest) CustomFieldConditions(customFieldConditions string) ApiGetSystemAudittrailRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetSystemAudittrailRequest) OrderBy(orderBy string) ApiGetSystemAudittrailRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetSystemAudittrailRequest) Fields(fields string) ApiGetSystemAudittrailRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetSystemAudittrailRequest) Page(page int32) ApiGetSystemAudittrailRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetSystemAudittrailRequest) PageSize(pageSize int32) ApiGetSystemAudittrailRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetSystemAudittrailRequest) PageId(pageId int32) ApiGetSystemAudittrailRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetSystemAudittrailRequest) ClientId(clientId string) ApiGetSystemAudittrailRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetSystemAudittrailRequest) Execute() ([]AuditTrailEntry, *http.Response, error) {
	return r.ApiService.GetSystemAudittrailExecute(r)
}

/*
GetSystemAudittrail Get List of AuditTrailEntry

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetSystemAudittrailRequest
*/
func (a *AuditTrailAPIService) GetSystemAudittrail(ctx context.Context) ApiGetSystemAudittrailRequest {
	return ApiGetSystemAudittrailRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []AuditTrailEntry
func (a *AuditTrailAPIService) GetSystemAudittrailExecute(r ApiGetSystemAudittrailRequest) ([]AuditTrailEntry, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []AuditTrailEntry
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuditTrailAPIService.GetSystemAudittrail")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/audittrail"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.type_ != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "form", "")
	}
	if r.id != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "id", r.id, "form", "")
	}
	if r.deviceIdentifier != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "deviceIdentifier", r.deviceIdentifier, "form", "")
	}
	if r.xrefRecId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "xrefRecId", r.xrefRecId, "form", "")
	}
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

type ApiGetSystemAudittrailCountRequest struct {
	ctx context.Context
	ApiService *AuditTrailAPIService
	type_ *string
	id *int32
	deviceIdentifier *string
	xrefRecId *int32
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

// type
func (r ApiGetSystemAudittrailCountRequest) Type_(type_ string) ApiGetSystemAudittrailCountRequest {
	r.type_ = &type_
	return r
}

// id
func (r ApiGetSystemAudittrailCountRequest) Id(id int32) ApiGetSystemAudittrailCountRequest {
	r.id = &id
	return r
}

// deviceIdentifier
func (r ApiGetSystemAudittrailCountRequest) DeviceIdentifier(deviceIdentifier string) ApiGetSystemAudittrailCountRequest {
	r.deviceIdentifier = &deviceIdentifier
	return r
}

// xrefRecId
func (r ApiGetSystemAudittrailCountRequest) XrefRecId(xrefRecId int32) ApiGetSystemAudittrailCountRequest {
	r.xrefRecId = &xrefRecId
	return r
}

// 
func (r ApiGetSystemAudittrailCountRequest) Conditions(conditions string) ApiGetSystemAudittrailCountRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetSystemAudittrailCountRequest) ChildConditions(childConditions string) ApiGetSystemAudittrailCountRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetSystemAudittrailCountRequest) CustomFieldConditions(customFieldConditions string) ApiGetSystemAudittrailCountRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetSystemAudittrailCountRequest) OrderBy(orderBy string) ApiGetSystemAudittrailCountRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetSystemAudittrailCountRequest) Fields(fields string) ApiGetSystemAudittrailCountRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetSystemAudittrailCountRequest) Page(page int32) ApiGetSystemAudittrailCountRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetSystemAudittrailCountRequest) PageSize(pageSize int32) ApiGetSystemAudittrailCountRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetSystemAudittrailCountRequest) PageId(pageId int32) ApiGetSystemAudittrailCountRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetSystemAudittrailCountRequest) ClientId(clientId string) ApiGetSystemAudittrailCountRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetSystemAudittrailCountRequest) Execute() (*Count, *http.Response, error) {
	return r.ApiService.GetSystemAudittrailCountExecute(r)
}

/*
GetSystemAudittrailCount Get Count of AuditTrailEntry

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetSystemAudittrailCountRequest
*/
func (a *AuditTrailAPIService) GetSystemAudittrailCount(ctx context.Context) ApiGetSystemAudittrailCountRequest {
	return ApiGetSystemAudittrailCountRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Count
func (a *AuditTrailAPIService) GetSystemAudittrailCountExecute(r ApiGetSystemAudittrailCountRequest) (*Count, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Count
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuditTrailAPIService.GetSystemAudittrailCount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/audittrail/count"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.type_ != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "form", "")
	}
	if r.id != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "id", r.id, "form", "")
	}
	if r.deviceIdentifier != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "deviceIdentifier", r.deviceIdentifier, "form", "")
	}
	if r.xrefRecId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "xrefRecId", r.xrefRecId, "form", "")
	}
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
