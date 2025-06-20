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


// ConfigurationTypeQuestionInfoAPIService ConfigurationTypeQuestionInfoAPI service
type ConfigurationTypeQuestionInfoAPIService service

type ApiGetConfigurationsTypesByParentIdQuestionsByIdInfoRequest struct {
	ctx context.Context
	ApiService *ConfigurationTypeQuestionInfoAPIService
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
func (r ApiGetConfigurationsTypesByParentIdQuestionsByIdInfoRequest) Conditions(conditions string) ApiGetConfigurationsTypesByParentIdQuestionsByIdInfoRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetConfigurationsTypesByParentIdQuestionsByIdInfoRequest) ChildConditions(childConditions string) ApiGetConfigurationsTypesByParentIdQuestionsByIdInfoRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetConfigurationsTypesByParentIdQuestionsByIdInfoRequest) CustomFieldConditions(customFieldConditions string) ApiGetConfigurationsTypesByParentIdQuestionsByIdInfoRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetConfigurationsTypesByParentIdQuestionsByIdInfoRequest) OrderBy(orderBy string) ApiGetConfigurationsTypesByParentIdQuestionsByIdInfoRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetConfigurationsTypesByParentIdQuestionsByIdInfoRequest) Fields(fields string) ApiGetConfigurationsTypesByParentIdQuestionsByIdInfoRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetConfigurationsTypesByParentIdQuestionsByIdInfoRequest) Page(page int32) ApiGetConfigurationsTypesByParentIdQuestionsByIdInfoRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetConfigurationsTypesByParentIdQuestionsByIdInfoRequest) PageSize(pageSize int32) ApiGetConfigurationsTypesByParentIdQuestionsByIdInfoRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetConfigurationsTypesByParentIdQuestionsByIdInfoRequest) PageId(pageId int32) ApiGetConfigurationsTypesByParentIdQuestionsByIdInfoRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetConfigurationsTypesByParentIdQuestionsByIdInfoRequest) ClientId(clientId string) ApiGetConfigurationsTypesByParentIdQuestionsByIdInfoRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetConfigurationsTypesByParentIdQuestionsByIdInfoRequest) Execute() (*ConfigurationTypeQuestionInfo, *http.Response, error) {
	return r.ApiService.GetConfigurationsTypesByParentIdQuestionsByIdInfoExecute(r)
}

/*
GetConfigurationsTypesByParentIdQuestionsByIdInfo Get ConfigurationTypeQuestionInfo

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id ConfigurationTypeQuestionInfo
 @param parentId typeId
 @return ApiGetConfigurationsTypesByParentIdQuestionsByIdInfoRequest
*/
func (a *ConfigurationTypeQuestionInfoAPIService) GetConfigurationsTypesByParentIdQuestionsByIdInfo(ctx context.Context, id int32, parentId int32) ApiGetConfigurationsTypesByParentIdQuestionsByIdInfoRequest {
	return ApiGetConfigurationsTypesByParentIdQuestionsByIdInfoRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
		parentId: parentId,
	}
}

// Execute executes the request
//  @return ConfigurationTypeQuestionInfo
func (a *ConfigurationTypeQuestionInfoAPIService) GetConfigurationsTypesByParentIdQuestionsByIdInfoExecute(r ApiGetConfigurationsTypesByParentIdQuestionsByIdInfoRequest) (*ConfigurationTypeQuestionInfo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ConfigurationTypeQuestionInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigurationTypeQuestionInfoAPIService.GetConfigurationsTypesByParentIdQuestionsByIdInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/configurations/types/{parentId}/questions/{id}/info"
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
