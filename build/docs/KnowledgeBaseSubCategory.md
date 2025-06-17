# KnowledgeBaseSubCategory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  Max length: 50; | 
**Category** | [**KBCategoryReference**](KBCategoryReference.md) |  | 
**Location** | Pointer to [**SystemLocationReference**](SystemLocationReference.md) |  | [optional] 
**Department** | Pointer to [**SystemDepartmentReference**](SystemDepartmentReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewKnowledgeBaseSubCategory

`func NewKnowledgeBaseSubCategory(name string, category KBCategoryReference, ) *KnowledgeBaseSubCategory`

NewKnowledgeBaseSubCategory instantiates a new KnowledgeBaseSubCategory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKnowledgeBaseSubCategoryWithDefaults

`func NewKnowledgeBaseSubCategoryWithDefaults() *KnowledgeBaseSubCategory`

NewKnowledgeBaseSubCategoryWithDefaults instantiates a new KnowledgeBaseSubCategory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KnowledgeBaseSubCategory) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KnowledgeBaseSubCategory) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KnowledgeBaseSubCategory) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KnowledgeBaseSubCategory) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *KnowledgeBaseSubCategory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KnowledgeBaseSubCategory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KnowledgeBaseSubCategory) SetName(v string)`

SetName sets Name field to given value.


### GetCategory

`func (o *KnowledgeBaseSubCategory) GetCategory() KBCategoryReference`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *KnowledgeBaseSubCategory) GetCategoryOk() (*KBCategoryReference, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *KnowledgeBaseSubCategory) SetCategory(v KBCategoryReference)`

SetCategory sets Category field to given value.


### GetLocation

`func (o *KnowledgeBaseSubCategory) GetLocation() SystemLocationReference`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *KnowledgeBaseSubCategory) GetLocationOk() (*SystemLocationReference, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *KnowledgeBaseSubCategory) SetLocation(v SystemLocationReference)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *KnowledgeBaseSubCategory) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetDepartment

`func (o *KnowledgeBaseSubCategory) GetDepartment() SystemDepartmentReference`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *KnowledgeBaseSubCategory) GetDepartmentOk() (*SystemDepartmentReference, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *KnowledgeBaseSubCategory) SetDepartment(v SystemDepartmentReference)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *KnowledgeBaseSubCategory) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### GetInfo

`func (o *KnowledgeBaseSubCategory) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *KnowledgeBaseSubCategory) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *KnowledgeBaseSubCategory) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *KnowledgeBaseSubCategory) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


