# BoardType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  Max length: 50; | 
**Category** | Pointer to **NullableString** |  | [optional] 
**DefaultFlag** | Pointer to **NullableBool** |  | [optional] 
**InactiveFlag** | Pointer to **NullableBool** |  | [optional] 
**RequestForChangeFlag** | Pointer to **NullableBool** |  | [optional] 
**IntegrationXref** | Pointer to **string** |  Max length: 50; | [optional] 
**SkillCategory** | Pointer to [**SkillCategoryReference**](SkillCategoryReference.md) |  | [optional] 
**Skill** | Pointer to [**SkillReference**](SkillReference.md) |  | [optional] 
**Board** | Pointer to [**BoardReference**](BoardReference.md) |  | [optional] 
**Location** | Pointer to [**SystemLocationReference**](SystemLocationReference.md) |  | [optional] 
**Department** | Pointer to [**SystemDepartmentReference**](SystemDepartmentReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewBoardType

`func NewBoardType(name string, ) *BoardType`

NewBoardType instantiates a new BoardType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBoardTypeWithDefaults

`func NewBoardTypeWithDefaults() *BoardType`

NewBoardTypeWithDefaults instantiates a new BoardType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BoardType) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BoardType) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BoardType) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *BoardType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *BoardType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BoardType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BoardType) SetName(v string)`

SetName sets Name field to given value.


### GetCategory

`func (o *BoardType) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *BoardType) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *BoardType) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *BoardType) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *BoardType) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *BoardType) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetDefaultFlag

`func (o *BoardType) GetDefaultFlag() bool`

GetDefaultFlag returns the DefaultFlag field if non-nil, zero value otherwise.

### GetDefaultFlagOk

`func (o *BoardType) GetDefaultFlagOk() (*bool, bool)`

GetDefaultFlagOk returns a tuple with the DefaultFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFlag

`func (o *BoardType) SetDefaultFlag(v bool)`

SetDefaultFlag sets DefaultFlag field to given value.

### HasDefaultFlag

`func (o *BoardType) HasDefaultFlag() bool`

HasDefaultFlag returns a boolean if a field has been set.

### SetDefaultFlagNil

`func (o *BoardType) SetDefaultFlagNil(b bool)`

 SetDefaultFlagNil sets the value for DefaultFlag to be an explicit nil

### UnsetDefaultFlag
`func (o *BoardType) UnsetDefaultFlag()`

UnsetDefaultFlag ensures that no value is present for DefaultFlag, not even an explicit nil
### GetInactiveFlag

`func (o *BoardType) GetInactiveFlag() bool`

GetInactiveFlag returns the InactiveFlag field if non-nil, zero value otherwise.

### GetInactiveFlagOk

`func (o *BoardType) GetInactiveFlagOk() (*bool, bool)`

GetInactiveFlagOk returns a tuple with the InactiveFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveFlag

`func (o *BoardType) SetInactiveFlag(v bool)`

SetInactiveFlag sets InactiveFlag field to given value.

### HasInactiveFlag

`func (o *BoardType) HasInactiveFlag() bool`

HasInactiveFlag returns a boolean if a field has been set.

### SetInactiveFlagNil

`func (o *BoardType) SetInactiveFlagNil(b bool)`

 SetInactiveFlagNil sets the value for InactiveFlag to be an explicit nil

### UnsetInactiveFlag
`func (o *BoardType) UnsetInactiveFlag()`

UnsetInactiveFlag ensures that no value is present for InactiveFlag, not even an explicit nil
### GetRequestForChangeFlag

`func (o *BoardType) GetRequestForChangeFlag() bool`

GetRequestForChangeFlag returns the RequestForChangeFlag field if non-nil, zero value otherwise.

### GetRequestForChangeFlagOk

`func (o *BoardType) GetRequestForChangeFlagOk() (*bool, bool)`

GetRequestForChangeFlagOk returns a tuple with the RequestForChangeFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestForChangeFlag

`func (o *BoardType) SetRequestForChangeFlag(v bool)`

SetRequestForChangeFlag sets RequestForChangeFlag field to given value.

### HasRequestForChangeFlag

`func (o *BoardType) HasRequestForChangeFlag() bool`

HasRequestForChangeFlag returns a boolean if a field has been set.

### SetRequestForChangeFlagNil

`func (o *BoardType) SetRequestForChangeFlagNil(b bool)`

 SetRequestForChangeFlagNil sets the value for RequestForChangeFlag to be an explicit nil

### UnsetRequestForChangeFlag
`func (o *BoardType) UnsetRequestForChangeFlag()`

UnsetRequestForChangeFlag ensures that no value is present for RequestForChangeFlag, not even an explicit nil
### GetIntegrationXref

`func (o *BoardType) GetIntegrationXref() string`

GetIntegrationXref returns the IntegrationXref field if non-nil, zero value otherwise.

### GetIntegrationXrefOk

`func (o *BoardType) GetIntegrationXrefOk() (*string, bool)`

GetIntegrationXrefOk returns a tuple with the IntegrationXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationXref

`func (o *BoardType) SetIntegrationXref(v string)`

SetIntegrationXref sets IntegrationXref field to given value.

### HasIntegrationXref

`func (o *BoardType) HasIntegrationXref() bool`

HasIntegrationXref returns a boolean if a field has been set.

### GetSkillCategory

`func (o *BoardType) GetSkillCategory() SkillCategoryReference`

GetSkillCategory returns the SkillCategory field if non-nil, zero value otherwise.

### GetSkillCategoryOk

`func (o *BoardType) GetSkillCategoryOk() (*SkillCategoryReference, bool)`

GetSkillCategoryOk returns a tuple with the SkillCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkillCategory

`func (o *BoardType) SetSkillCategory(v SkillCategoryReference)`

SetSkillCategory sets SkillCategory field to given value.

### HasSkillCategory

`func (o *BoardType) HasSkillCategory() bool`

HasSkillCategory returns a boolean if a field has been set.

### GetSkill

`func (o *BoardType) GetSkill() SkillReference`

GetSkill returns the Skill field if non-nil, zero value otherwise.

### GetSkillOk

`func (o *BoardType) GetSkillOk() (*SkillReference, bool)`

GetSkillOk returns a tuple with the Skill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkill

`func (o *BoardType) SetSkill(v SkillReference)`

SetSkill sets Skill field to given value.

### HasSkill

`func (o *BoardType) HasSkill() bool`

HasSkill returns a boolean if a field has been set.

### GetBoard

`func (o *BoardType) GetBoard() BoardReference`

GetBoard returns the Board field if non-nil, zero value otherwise.

### GetBoardOk

`func (o *BoardType) GetBoardOk() (*BoardReference, bool)`

GetBoardOk returns a tuple with the Board field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoard

`func (o *BoardType) SetBoard(v BoardReference)`

SetBoard sets Board field to given value.

### HasBoard

`func (o *BoardType) HasBoard() bool`

HasBoard returns a boolean if a field has been set.

### GetLocation

`func (o *BoardType) GetLocation() SystemLocationReference`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *BoardType) GetLocationOk() (*SystemLocationReference, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *BoardType) SetLocation(v SystemLocationReference)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *BoardType) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetDepartment

`func (o *BoardType) GetDepartment() SystemDepartmentReference`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *BoardType) GetDepartmentOk() (*SystemDepartmentReference, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *BoardType) SetDepartment(v SystemDepartmentReference)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *BoardType) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### GetInfo

`func (o *BoardType) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *BoardType) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *BoardType) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *BoardType) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


