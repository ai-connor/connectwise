# ProjectSecurityRoleSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**AddLevel** | Pointer to **NullableString** |  | [optional] 
**EditLevel** | Pointer to **NullableString** |  | [optional] 
**DeleteLevel** | Pointer to **NullableString** |  | [optional] 
**InquireLevel** | Pointer to **NullableString** |  | [optional] 
**ModuleIdentifier** | Pointer to **string** |  Max length: 50; | [optional] 
**MyFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewProjectSecurityRoleSetting

`func NewProjectSecurityRoleSetting() *ProjectSecurityRoleSetting`

NewProjectSecurityRoleSetting instantiates a new ProjectSecurityRoleSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectSecurityRoleSettingWithDefaults

`func NewProjectSecurityRoleSettingWithDefaults() *ProjectSecurityRoleSetting`

NewProjectSecurityRoleSettingWithDefaults instantiates a new ProjectSecurityRoleSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectSecurityRoleSetting) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectSecurityRoleSetting) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectSecurityRoleSetting) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ProjectSecurityRoleSetting) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAddLevel

`func (o *ProjectSecurityRoleSetting) GetAddLevel() string`

GetAddLevel returns the AddLevel field if non-nil, zero value otherwise.

### GetAddLevelOk

`func (o *ProjectSecurityRoleSetting) GetAddLevelOk() (*string, bool)`

GetAddLevelOk returns a tuple with the AddLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddLevel

`func (o *ProjectSecurityRoleSetting) SetAddLevel(v string)`

SetAddLevel sets AddLevel field to given value.

### HasAddLevel

`func (o *ProjectSecurityRoleSetting) HasAddLevel() bool`

HasAddLevel returns a boolean if a field has been set.

### SetAddLevelNil

`func (o *ProjectSecurityRoleSetting) SetAddLevelNil(b bool)`

 SetAddLevelNil sets the value for AddLevel to be an explicit nil

### UnsetAddLevel
`func (o *ProjectSecurityRoleSetting) UnsetAddLevel()`

UnsetAddLevel ensures that no value is present for AddLevel, not even an explicit nil
### GetEditLevel

`func (o *ProjectSecurityRoleSetting) GetEditLevel() string`

GetEditLevel returns the EditLevel field if non-nil, zero value otherwise.

### GetEditLevelOk

`func (o *ProjectSecurityRoleSetting) GetEditLevelOk() (*string, bool)`

GetEditLevelOk returns a tuple with the EditLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditLevel

`func (o *ProjectSecurityRoleSetting) SetEditLevel(v string)`

SetEditLevel sets EditLevel field to given value.

### HasEditLevel

`func (o *ProjectSecurityRoleSetting) HasEditLevel() bool`

HasEditLevel returns a boolean if a field has been set.

### SetEditLevelNil

`func (o *ProjectSecurityRoleSetting) SetEditLevelNil(b bool)`

 SetEditLevelNil sets the value for EditLevel to be an explicit nil

### UnsetEditLevel
`func (o *ProjectSecurityRoleSetting) UnsetEditLevel()`

UnsetEditLevel ensures that no value is present for EditLevel, not even an explicit nil
### GetDeleteLevel

`func (o *ProjectSecurityRoleSetting) GetDeleteLevel() string`

GetDeleteLevel returns the DeleteLevel field if non-nil, zero value otherwise.

### GetDeleteLevelOk

`func (o *ProjectSecurityRoleSetting) GetDeleteLevelOk() (*string, bool)`

GetDeleteLevelOk returns a tuple with the DeleteLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteLevel

`func (o *ProjectSecurityRoleSetting) SetDeleteLevel(v string)`

SetDeleteLevel sets DeleteLevel field to given value.

### HasDeleteLevel

`func (o *ProjectSecurityRoleSetting) HasDeleteLevel() bool`

HasDeleteLevel returns a boolean if a field has been set.

### SetDeleteLevelNil

`func (o *ProjectSecurityRoleSetting) SetDeleteLevelNil(b bool)`

 SetDeleteLevelNil sets the value for DeleteLevel to be an explicit nil

### UnsetDeleteLevel
`func (o *ProjectSecurityRoleSetting) UnsetDeleteLevel()`

UnsetDeleteLevel ensures that no value is present for DeleteLevel, not even an explicit nil
### GetInquireLevel

`func (o *ProjectSecurityRoleSetting) GetInquireLevel() string`

GetInquireLevel returns the InquireLevel field if non-nil, zero value otherwise.

### GetInquireLevelOk

`func (o *ProjectSecurityRoleSetting) GetInquireLevelOk() (*string, bool)`

GetInquireLevelOk returns a tuple with the InquireLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInquireLevel

`func (o *ProjectSecurityRoleSetting) SetInquireLevel(v string)`

SetInquireLevel sets InquireLevel field to given value.

### HasInquireLevel

`func (o *ProjectSecurityRoleSetting) HasInquireLevel() bool`

HasInquireLevel returns a boolean if a field has been set.

### SetInquireLevelNil

`func (o *ProjectSecurityRoleSetting) SetInquireLevelNil(b bool)`

 SetInquireLevelNil sets the value for InquireLevel to be an explicit nil

### UnsetInquireLevel
`func (o *ProjectSecurityRoleSetting) UnsetInquireLevel()`

UnsetInquireLevel ensures that no value is present for InquireLevel, not even an explicit nil
### GetModuleIdentifier

`func (o *ProjectSecurityRoleSetting) GetModuleIdentifier() string`

GetModuleIdentifier returns the ModuleIdentifier field if non-nil, zero value otherwise.

### GetModuleIdentifierOk

`func (o *ProjectSecurityRoleSetting) GetModuleIdentifierOk() (*string, bool)`

GetModuleIdentifierOk returns a tuple with the ModuleIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleIdentifier

`func (o *ProjectSecurityRoleSetting) SetModuleIdentifier(v string)`

SetModuleIdentifier sets ModuleIdentifier field to given value.

### HasModuleIdentifier

`func (o *ProjectSecurityRoleSetting) HasModuleIdentifier() bool`

HasModuleIdentifier returns a boolean if a field has been set.

### GetMyFlag

`func (o *ProjectSecurityRoleSetting) GetMyFlag() bool`

GetMyFlag returns the MyFlag field if non-nil, zero value otherwise.

### GetMyFlagOk

`func (o *ProjectSecurityRoleSetting) GetMyFlagOk() (*bool, bool)`

GetMyFlagOk returns a tuple with the MyFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMyFlag

`func (o *ProjectSecurityRoleSetting) SetMyFlag(v bool)`

SetMyFlag sets MyFlag field to given value.

### HasMyFlag

`func (o *ProjectSecurityRoleSetting) HasMyFlag() bool`

HasMyFlag returns a boolean if a field has been set.

### SetMyFlagNil

`func (o *ProjectSecurityRoleSetting) SetMyFlagNil(b bool)`

 SetMyFlagNil sets the value for MyFlag to be an explicit nil

### UnsetMyFlag
`func (o *ProjectSecurityRoleSetting) UnsetMyFlag()`

UnsetMyFlag ensures that no value is present for MyFlag, not even an explicit nil
### GetInfo

`func (o *ProjectSecurityRoleSetting) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ProjectSecurityRoleSetting) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ProjectSecurityRoleSetting) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ProjectSecurityRoleSetting) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


