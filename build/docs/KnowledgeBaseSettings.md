# KnowledgeBaseSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**RequireApproval** | **NullableBool** |  | 
**DefaultApprover** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewKnowledgeBaseSettings

`func NewKnowledgeBaseSettings(requireApproval NullableBool, ) *KnowledgeBaseSettings`

NewKnowledgeBaseSettings instantiates a new KnowledgeBaseSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKnowledgeBaseSettingsWithDefaults

`func NewKnowledgeBaseSettingsWithDefaults() *KnowledgeBaseSettings`

NewKnowledgeBaseSettingsWithDefaults instantiates a new KnowledgeBaseSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KnowledgeBaseSettings) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KnowledgeBaseSettings) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KnowledgeBaseSettings) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KnowledgeBaseSettings) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRequireApproval

`func (o *KnowledgeBaseSettings) GetRequireApproval() bool`

GetRequireApproval returns the RequireApproval field if non-nil, zero value otherwise.

### GetRequireApprovalOk

`func (o *KnowledgeBaseSettings) GetRequireApprovalOk() (*bool, bool)`

GetRequireApprovalOk returns a tuple with the RequireApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireApproval

`func (o *KnowledgeBaseSettings) SetRequireApproval(v bool)`

SetRequireApproval sets RequireApproval field to given value.


### SetRequireApprovalNil

`func (o *KnowledgeBaseSettings) SetRequireApprovalNil(b bool)`

 SetRequireApprovalNil sets the value for RequireApproval to be an explicit nil

### UnsetRequireApproval
`func (o *KnowledgeBaseSettings) UnsetRequireApproval()`

UnsetRequireApproval ensures that no value is present for RequireApproval, not even an explicit nil
### GetDefaultApprover

`func (o *KnowledgeBaseSettings) GetDefaultApprover() MemberReference`

GetDefaultApprover returns the DefaultApprover field if non-nil, zero value otherwise.

### GetDefaultApproverOk

`func (o *KnowledgeBaseSettings) GetDefaultApproverOk() (*MemberReference, bool)`

GetDefaultApproverOk returns a tuple with the DefaultApprover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultApprover

`func (o *KnowledgeBaseSettings) SetDefaultApprover(v MemberReference)`

SetDefaultApprover sets DefaultApprover field to given value.

### HasDefaultApprover

`func (o *KnowledgeBaseSettings) HasDefaultApprover() bool`

HasDefaultApprover returns a boolean if a field has been set.

### GetInfo

`func (o *KnowledgeBaseSettings) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *KnowledgeBaseSettings) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *KnowledgeBaseSettings) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *KnowledgeBaseSettings) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


