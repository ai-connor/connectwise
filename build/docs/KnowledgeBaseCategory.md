# KnowledgeBaseCategory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  Max length: 50; | 
**Location** | Pointer to [**SystemLocationReference**](SystemLocationReference.md) |  | [optional] 
**Department** | Pointer to [**SystemDepartmentReference**](SystemDepartmentReference.md) |  | [optional] 
**Approver** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewKnowledgeBaseCategory

`func NewKnowledgeBaseCategory(name string, ) *KnowledgeBaseCategory`

NewKnowledgeBaseCategory instantiates a new KnowledgeBaseCategory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKnowledgeBaseCategoryWithDefaults

`func NewKnowledgeBaseCategoryWithDefaults() *KnowledgeBaseCategory`

NewKnowledgeBaseCategoryWithDefaults instantiates a new KnowledgeBaseCategory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KnowledgeBaseCategory) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KnowledgeBaseCategory) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KnowledgeBaseCategory) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KnowledgeBaseCategory) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *KnowledgeBaseCategory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KnowledgeBaseCategory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KnowledgeBaseCategory) SetName(v string)`

SetName sets Name field to given value.


### GetLocation

`func (o *KnowledgeBaseCategory) GetLocation() SystemLocationReference`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *KnowledgeBaseCategory) GetLocationOk() (*SystemLocationReference, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *KnowledgeBaseCategory) SetLocation(v SystemLocationReference)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *KnowledgeBaseCategory) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetDepartment

`func (o *KnowledgeBaseCategory) GetDepartment() SystemDepartmentReference`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *KnowledgeBaseCategory) GetDepartmentOk() (*SystemDepartmentReference, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *KnowledgeBaseCategory) SetDepartment(v SystemDepartmentReference)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *KnowledgeBaseCategory) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### GetApprover

`func (o *KnowledgeBaseCategory) GetApprover() MemberReference`

GetApprover returns the Approver field if non-nil, zero value otherwise.

### GetApproverOk

`func (o *KnowledgeBaseCategory) GetApproverOk() (*MemberReference, bool)`

GetApproverOk returns a tuple with the Approver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprover

`func (o *KnowledgeBaseCategory) SetApprover(v MemberReference)`

SetApprover sets Approver field to given value.

### HasApprover

`func (o *KnowledgeBaseCategory) HasApprover() bool`

HasApprover returns a boolean if a field has been set.

### GetInfo

`func (o *KnowledgeBaseCategory) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *KnowledgeBaseCategory) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *KnowledgeBaseCategory) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *KnowledgeBaseCategory) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


