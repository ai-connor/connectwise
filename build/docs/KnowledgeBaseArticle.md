# KnowledgeBaseArticle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Title** | **string** |  | 
**Issue** | **string** |  | 
**Resolution** | **string** |  | 
**LocationId** | Pointer to **NullableInt32** |  | [optional] 
**BusinessUnitId** | Pointer to **NullableInt32** |  | [optional] 
**Board** | Pointer to [**BoardReference**](BoardReference.md) |  | [optional] 
**CategoryId** | Pointer to **NullableInt32** |  | [optional] 
**SubCategoryId** | Pointer to **NullableInt32** |  | [optional] 
**DateCreated** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewKnowledgeBaseArticle

`func NewKnowledgeBaseArticle(title string, issue string, resolution string, ) *KnowledgeBaseArticle`

NewKnowledgeBaseArticle instantiates a new KnowledgeBaseArticle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKnowledgeBaseArticleWithDefaults

`func NewKnowledgeBaseArticleWithDefaults() *KnowledgeBaseArticle`

NewKnowledgeBaseArticleWithDefaults instantiates a new KnowledgeBaseArticle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KnowledgeBaseArticle) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KnowledgeBaseArticle) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KnowledgeBaseArticle) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KnowledgeBaseArticle) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *KnowledgeBaseArticle) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *KnowledgeBaseArticle) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *KnowledgeBaseArticle) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetIssue

`func (o *KnowledgeBaseArticle) GetIssue() string`

GetIssue returns the Issue field if non-nil, zero value otherwise.

### GetIssueOk

`func (o *KnowledgeBaseArticle) GetIssueOk() (*string, bool)`

GetIssueOk returns a tuple with the Issue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssue

`func (o *KnowledgeBaseArticle) SetIssue(v string)`

SetIssue sets Issue field to given value.


### GetResolution

`func (o *KnowledgeBaseArticle) GetResolution() string`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *KnowledgeBaseArticle) GetResolutionOk() (*string, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *KnowledgeBaseArticle) SetResolution(v string)`

SetResolution sets Resolution field to given value.


### GetLocationId

`func (o *KnowledgeBaseArticle) GetLocationId() int32`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *KnowledgeBaseArticle) GetLocationIdOk() (*int32, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *KnowledgeBaseArticle) SetLocationId(v int32)`

SetLocationId sets LocationId field to given value.

### HasLocationId

`func (o *KnowledgeBaseArticle) HasLocationId() bool`

HasLocationId returns a boolean if a field has been set.

### SetLocationIdNil

`func (o *KnowledgeBaseArticle) SetLocationIdNil(b bool)`

 SetLocationIdNil sets the value for LocationId to be an explicit nil

### UnsetLocationId
`func (o *KnowledgeBaseArticle) UnsetLocationId()`

UnsetLocationId ensures that no value is present for LocationId, not even an explicit nil
### GetBusinessUnitId

`func (o *KnowledgeBaseArticle) GetBusinessUnitId() int32`

GetBusinessUnitId returns the BusinessUnitId field if non-nil, zero value otherwise.

### GetBusinessUnitIdOk

`func (o *KnowledgeBaseArticle) GetBusinessUnitIdOk() (*int32, bool)`

GetBusinessUnitIdOk returns a tuple with the BusinessUnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessUnitId

`func (o *KnowledgeBaseArticle) SetBusinessUnitId(v int32)`

SetBusinessUnitId sets BusinessUnitId field to given value.

### HasBusinessUnitId

`func (o *KnowledgeBaseArticle) HasBusinessUnitId() bool`

HasBusinessUnitId returns a boolean if a field has been set.

### SetBusinessUnitIdNil

`func (o *KnowledgeBaseArticle) SetBusinessUnitIdNil(b bool)`

 SetBusinessUnitIdNil sets the value for BusinessUnitId to be an explicit nil

### UnsetBusinessUnitId
`func (o *KnowledgeBaseArticle) UnsetBusinessUnitId()`

UnsetBusinessUnitId ensures that no value is present for BusinessUnitId, not even an explicit nil
### GetBoard

`func (o *KnowledgeBaseArticle) GetBoard() BoardReference`

GetBoard returns the Board field if non-nil, zero value otherwise.

### GetBoardOk

`func (o *KnowledgeBaseArticle) GetBoardOk() (*BoardReference, bool)`

GetBoardOk returns a tuple with the Board field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoard

`func (o *KnowledgeBaseArticle) SetBoard(v BoardReference)`

SetBoard sets Board field to given value.

### HasBoard

`func (o *KnowledgeBaseArticle) HasBoard() bool`

HasBoard returns a boolean if a field has been set.

### GetCategoryId

`func (o *KnowledgeBaseArticle) GetCategoryId() int32`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *KnowledgeBaseArticle) GetCategoryIdOk() (*int32, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *KnowledgeBaseArticle) SetCategoryId(v int32)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *KnowledgeBaseArticle) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.

### SetCategoryIdNil

`func (o *KnowledgeBaseArticle) SetCategoryIdNil(b bool)`

 SetCategoryIdNil sets the value for CategoryId to be an explicit nil

### UnsetCategoryId
`func (o *KnowledgeBaseArticle) UnsetCategoryId()`

UnsetCategoryId ensures that no value is present for CategoryId, not even an explicit nil
### GetSubCategoryId

`func (o *KnowledgeBaseArticle) GetSubCategoryId() int32`

GetSubCategoryId returns the SubCategoryId field if non-nil, zero value otherwise.

### GetSubCategoryIdOk

`func (o *KnowledgeBaseArticle) GetSubCategoryIdOk() (*int32, bool)`

GetSubCategoryIdOk returns a tuple with the SubCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubCategoryId

`func (o *KnowledgeBaseArticle) SetSubCategoryId(v int32)`

SetSubCategoryId sets SubCategoryId field to given value.

### HasSubCategoryId

`func (o *KnowledgeBaseArticle) HasSubCategoryId() bool`

HasSubCategoryId returns a boolean if a field has been set.

### SetSubCategoryIdNil

`func (o *KnowledgeBaseArticle) SetSubCategoryIdNil(b bool)`

 SetSubCategoryIdNil sets the value for SubCategoryId to be an explicit nil

### UnsetSubCategoryId
`func (o *KnowledgeBaseArticle) UnsetSubCategoryId()`

UnsetSubCategoryId ensures that no value is present for SubCategoryId, not even an explicit nil
### GetDateCreated

`func (o *KnowledgeBaseArticle) GetDateCreated() string`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *KnowledgeBaseArticle) GetDateCreatedOk() (*string, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *KnowledgeBaseArticle) SetDateCreated(v string)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *KnowledgeBaseArticle) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetCreatedBy

`func (o *KnowledgeBaseArticle) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *KnowledgeBaseArticle) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *KnowledgeBaseArticle) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *KnowledgeBaseArticle) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetInfo

`func (o *KnowledgeBaseArticle) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *KnowledgeBaseArticle) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *KnowledgeBaseArticle) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *KnowledgeBaseArticle) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


