# PortalSecuritySetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**FunctionIdentifier** | Pointer to **string** |  | [optional] 
**FunctionDescription** | Pointer to **string** |  | [optional] 
**LevelOne** | Pointer to **NullableBool** |  | [optional] 
**LevelTwo** | Pointer to **NullableBool** |  | [optional] 
**LevelThree** | Pointer to **NullableBool** |  | [optional] 
**LevelFour** | Pointer to **NullableBool** |  | [optional] 
**LevelFive** | Pointer to **NullableBool** |  | [optional] 
**LevelSix** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewPortalSecuritySetting

`func NewPortalSecuritySetting() *PortalSecuritySetting`

NewPortalSecuritySetting instantiates a new PortalSecuritySetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortalSecuritySettingWithDefaults

`func NewPortalSecuritySettingWithDefaults() *PortalSecuritySetting`

NewPortalSecuritySettingWithDefaults instantiates a new PortalSecuritySetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PortalSecuritySetting) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PortalSecuritySetting) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PortalSecuritySetting) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PortalSecuritySetting) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFunctionIdentifier

`func (o *PortalSecuritySetting) GetFunctionIdentifier() string`

GetFunctionIdentifier returns the FunctionIdentifier field if non-nil, zero value otherwise.

### GetFunctionIdentifierOk

`func (o *PortalSecuritySetting) GetFunctionIdentifierOk() (*string, bool)`

GetFunctionIdentifierOk returns a tuple with the FunctionIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionIdentifier

`func (o *PortalSecuritySetting) SetFunctionIdentifier(v string)`

SetFunctionIdentifier sets FunctionIdentifier field to given value.

### HasFunctionIdentifier

`func (o *PortalSecuritySetting) HasFunctionIdentifier() bool`

HasFunctionIdentifier returns a boolean if a field has been set.

### GetFunctionDescription

`func (o *PortalSecuritySetting) GetFunctionDescription() string`

GetFunctionDescription returns the FunctionDescription field if non-nil, zero value otherwise.

### GetFunctionDescriptionOk

`func (o *PortalSecuritySetting) GetFunctionDescriptionOk() (*string, bool)`

GetFunctionDescriptionOk returns a tuple with the FunctionDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionDescription

`func (o *PortalSecuritySetting) SetFunctionDescription(v string)`

SetFunctionDescription sets FunctionDescription field to given value.

### HasFunctionDescription

`func (o *PortalSecuritySetting) HasFunctionDescription() bool`

HasFunctionDescription returns a boolean if a field has been set.

### GetLevelOne

`func (o *PortalSecuritySetting) GetLevelOne() bool`

GetLevelOne returns the LevelOne field if non-nil, zero value otherwise.

### GetLevelOneOk

`func (o *PortalSecuritySetting) GetLevelOneOk() (*bool, bool)`

GetLevelOneOk returns a tuple with the LevelOne field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelOne

`func (o *PortalSecuritySetting) SetLevelOne(v bool)`

SetLevelOne sets LevelOne field to given value.

### HasLevelOne

`func (o *PortalSecuritySetting) HasLevelOne() bool`

HasLevelOne returns a boolean if a field has been set.

### SetLevelOneNil

`func (o *PortalSecuritySetting) SetLevelOneNil(b bool)`

 SetLevelOneNil sets the value for LevelOne to be an explicit nil

### UnsetLevelOne
`func (o *PortalSecuritySetting) UnsetLevelOne()`

UnsetLevelOne ensures that no value is present for LevelOne, not even an explicit nil
### GetLevelTwo

`func (o *PortalSecuritySetting) GetLevelTwo() bool`

GetLevelTwo returns the LevelTwo field if non-nil, zero value otherwise.

### GetLevelTwoOk

`func (o *PortalSecuritySetting) GetLevelTwoOk() (*bool, bool)`

GetLevelTwoOk returns a tuple with the LevelTwo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelTwo

`func (o *PortalSecuritySetting) SetLevelTwo(v bool)`

SetLevelTwo sets LevelTwo field to given value.

### HasLevelTwo

`func (o *PortalSecuritySetting) HasLevelTwo() bool`

HasLevelTwo returns a boolean if a field has been set.

### SetLevelTwoNil

`func (o *PortalSecuritySetting) SetLevelTwoNil(b bool)`

 SetLevelTwoNil sets the value for LevelTwo to be an explicit nil

### UnsetLevelTwo
`func (o *PortalSecuritySetting) UnsetLevelTwo()`

UnsetLevelTwo ensures that no value is present for LevelTwo, not even an explicit nil
### GetLevelThree

`func (o *PortalSecuritySetting) GetLevelThree() bool`

GetLevelThree returns the LevelThree field if non-nil, zero value otherwise.

### GetLevelThreeOk

`func (o *PortalSecuritySetting) GetLevelThreeOk() (*bool, bool)`

GetLevelThreeOk returns a tuple with the LevelThree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelThree

`func (o *PortalSecuritySetting) SetLevelThree(v bool)`

SetLevelThree sets LevelThree field to given value.

### HasLevelThree

`func (o *PortalSecuritySetting) HasLevelThree() bool`

HasLevelThree returns a boolean if a field has been set.

### SetLevelThreeNil

`func (o *PortalSecuritySetting) SetLevelThreeNil(b bool)`

 SetLevelThreeNil sets the value for LevelThree to be an explicit nil

### UnsetLevelThree
`func (o *PortalSecuritySetting) UnsetLevelThree()`

UnsetLevelThree ensures that no value is present for LevelThree, not even an explicit nil
### GetLevelFour

`func (o *PortalSecuritySetting) GetLevelFour() bool`

GetLevelFour returns the LevelFour field if non-nil, zero value otherwise.

### GetLevelFourOk

`func (o *PortalSecuritySetting) GetLevelFourOk() (*bool, bool)`

GetLevelFourOk returns a tuple with the LevelFour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelFour

`func (o *PortalSecuritySetting) SetLevelFour(v bool)`

SetLevelFour sets LevelFour field to given value.

### HasLevelFour

`func (o *PortalSecuritySetting) HasLevelFour() bool`

HasLevelFour returns a boolean if a field has been set.

### SetLevelFourNil

`func (o *PortalSecuritySetting) SetLevelFourNil(b bool)`

 SetLevelFourNil sets the value for LevelFour to be an explicit nil

### UnsetLevelFour
`func (o *PortalSecuritySetting) UnsetLevelFour()`

UnsetLevelFour ensures that no value is present for LevelFour, not even an explicit nil
### GetLevelFive

`func (o *PortalSecuritySetting) GetLevelFive() bool`

GetLevelFive returns the LevelFive field if non-nil, zero value otherwise.

### GetLevelFiveOk

`func (o *PortalSecuritySetting) GetLevelFiveOk() (*bool, bool)`

GetLevelFiveOk returns a tuple with the LevelFive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelFive

`func (o *PortalSecuritySetting) SetLevelFive(v bool)`

SetLevelFive sets LevelFive field to given value.

### HasLevelFive

`func (o *PortalSecuritySetting) HasLevelFive() bool`

HasLevelFive returns a boolean if a field has been set.

### SetLevelFiveNil

`func (o *PortalSecuritySetting) SetLevelFiveNil(b bool)`

 SetLevelFiveNil sets the value for LevelFive to be an explicit nil

### UnsetLevelFive
`func (o *PortalSecuritySetting) UnsetLevelFive()`

UnsetLevelFive ensures that no value is present for LevelFive, not even an explicit nil
### GetLevelSix

`func (o *PortalSecuritySetting) GetLevelSix() bool`

GetLevelSix returns the LevelSix field if non-nil, zero value otherwise.

### GetLevelSixOk

`func (o *PortalSecuritySetting) GetLevelSixOk() (*bool, bool)`

GetLevelSixOk returns a tuple with the LevelSix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelSix

`func (o *PortalSecuritySetting) SetLevelSix(v bool)`

SetLevelSix sets LevelSix field to given value.

### HasLevelSix

`func (o *PortalSecuritySetting) HasLevelSix() bool`

HasLevelSix returns a boolean if a field has been set.

### SetLevelSixNil

`func (o *PortalSecuritySetting) SetLevelSixNil(b bool)`

 SetLevelSixNil sets the value for LevelSix to be an explicit nil

### UnsetLevelSix
`func (o *PortalSecuritySetting) UnsetLevelSix()`

UnsetLevelSix ensures that no value is present for LevelSix, not even an explicit nil
### GetInfo

`func (o *PortalSecuritySetting) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *PortalSecuritySetting) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *PortalSecuritySetting) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *PortalSecuritySetting) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


