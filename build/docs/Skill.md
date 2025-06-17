# Skill

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  Max length: 50; | 
**Category** | [**SkillCategoryReference**](SkillCategoryReference.md) |  | 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewSkill

`func NewSkill(name string, category SkillCategoryReference, ) *Skill`

NewSkill instantiates a new Skill object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkillWithDefaults

`func NewSkillWithDefaults() *Skill`

NewSkillWithDefaults instantiates a new Skill object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Skill) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Skill) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Skill) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Skill) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Skill) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Skill) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Skill) SetName(v string)`

SetName sets Name field to given value.


### GetCategory

`func (o *Skill) GetCategory() SkillCategoryReference`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *Skill) GetCategoryOk() (*SkillCategoryReference, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *Skill) SetCategory(v SkillCategoryReference)`

SetCategory sets Category field to given value.


### GetInfo

`func (o *Skill) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *Skill) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *Skill) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *Skill) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


