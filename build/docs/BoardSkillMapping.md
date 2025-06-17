# BoardSkillMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Type** | [**ServiceTypeReference**](ServiceTypeReference.md) |  | 
**SubType** | Pointer to [**ServiceSubTypeReference**](ServiceSubTypeReference.md) |  | [optional] 
**Item** | Pointer to [**ServiceItemReference**](ServiceItemReference.md) |  | [optional] 
**SkillCategory** | [**SkillCategoryReference**](SkillCategoryReference.md) |  | 
**Skill** | [**SkillReference**](SkillReference.md) |  | 
**Board** | Pointer to [**BoardReference**](BoardReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewBoardSkillMapping

`func NewBoardSkillMapping(type_ ServiceTypeReference, skillCategory SkillCategoryReference, skill SkillReference, ) *BoardSkillMapping`

NewBoardSkillMapping instantiates a new BoardSkillMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBoardSkillMappingWithDefaults

`func NewBoardSkillMappingWithDefaults() *BoardSkillMapping`

NewBoardSkillMappingWithDefaults instantiates a new BoardSkillMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BoardSkillMapping) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BoardSkillMapping) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BoardSkillMapping) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *BoardSkillMapping) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *BoardSkillMapping) GetType() ServiceTypeReference`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BoardSkillMapping) GetTypeOk() (*ServiceTypeReference, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BoardSkillMapping) SetType(v ServiceTypeReference)`

SetType sets Type field to given value.


### GetSubType

`func (o *BoardSkillMapping) GetSubType() ServiceSubTypeReference`

GetSubType returns the SubType field if non-nil, zero value otherwise.

### GetSubTypeOk

`func (o *BoardSkillMapping) GetSubTypeOk() (*ServiceSubTypeReference, bool)`

GetSubTypeOk returns a tuple with the SubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubType

`func (o *BoardSkillMapping) SetSubType(v ServiceSubTypeReference)`

SetSubType sets SubType field to given value.

### HasSubType

`func (o *BoardSkillMapping) HasSubType() bool`

HasSubType returns a boolean if a field has been set.

### GetItem

`func (o *BoardSkillMapping) GetItem() ServiceItemReference`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *BoardSkillMapping) GetItemOk() (*ServiceItemReference, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *BoardSkillMapping) SetItem(v ServiceItemReference)`

SetItem sets Item field to given value.

### HasItem

`func (o *BoardSkillMapping) HasItem() bool`

HasItem returns a boolean if a field has been set.

### GetSkillCategory

`func (o *BoardSkillMapping) GetSkillCategory() SkillCategoryReference`

GetSkillCategory returns the SkillCategory field if non-nil, zero value otherwise.

### GetSkillCategoryOk

`func (o *BoardSkillMapping) GetSkillCategoryOk() (*SkillCategoryReference, bool)`

GetSkillCategoryOk returns a tuple with the SkillCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkillCategory

`func (o *BoardSkillMapping) SetSkillCategory(v SkillCategoryReference)`

SetSkillCategory sets SkillCategory field to given value.


### GetSkill

`func (o *BoardSkillMapping) GetSkill() SkillReference`

GetSkill returns the Skill field if non-nil, zero value otherwise.

### GetSkillOk

`func (o *BoardSkillMapping) GetSkillOk() (*SkillReference, bool)`

GetSkillOk returns a tuple with the Skill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkill

`func (o *BoardSkillMapping) SetSkill(v SkillReference)`

SetSkill sets Skill field to given value.


### GetBoard

`func (o *BoardSkillMapping) GetBoard() BoardReference`

GetBoard returns the Board field if non-nil, zero value otherwise.

### GetBoardOk

`func (o *BoardSkillMapping) GetBoardOk() (*BoardReference, bool)`

GetBoardOk returns a tuple with the Board field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoard

`func (o *BoardSkillMapping) SetBoard(v BoardReference)`

SetBoard sets Board field to given value.

### HasBoard

`func (o *BoardSkillMapping) HasBoard() bool`

HasBoard returns a boolean if a field has been set.

### GetInfo

`func (o *BoardSkillMapping) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *BoardSkillMapping) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *BoardSkillMapping) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *BoardSkillMapping) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


