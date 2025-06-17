# SecurityRoleSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**AddLevel** | Pointer to **NullableString** |  | [optional] 
**EditLevel** | Pointer to **NullableString** |  | [optional] 
**DeleteLevel** | Pointer to **NullableString** |  | [optional] 
**InquireLevel** | Pointer to **NullableString** |  | [optional] 
**ModuleFunctionName** | Pointer to **string** |  | [optional] 
**ModuleFunctionDescription** | Pointer to **string** |  | [optional] 
**MyAllFlag** | Pointer to **NullableBool** |  | [optional] 
**ModuleFunctionIdentifier** | Pointer to **string** |  | [optional] 
**ReportFlag** | Pointer to **NullableBool** |  | [optional] 
**RestrictFlag** | Pointer to **NullableBool** |  | [optional] 
**CustomFlag** | Pointer to **NullableBool** |  | [optional] 
**ModuleDescription** | Pointer to **string** |  | [optional] 
**ModuleIdentifier** | Pointer to **string** |  | [optional] 
**ModuleName** | Pointer to **string** |  | [optional] 
**SortOrder** | Pointer to **NullableInt32** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewSecurityRoleSetting

`func NewSecurityRoleSetting() *SecurityRoleSetting`

NewSecurityRoleSetting instantiates a new SecurityRoleSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityRoleSettingWithDefaults

`func NewSecurityRoleSettingWithDefaults() *SecurityRoleSetting`

NewSecurityRoleSettingWithDefaults instantiates a new SecurityRoleSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SecurityRoleSetting) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecurityRoleSetting) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecurityRoleSetting) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SecurityRoleSetting) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAddLevel

`func (o *SecurityRoleSetting) GetAddLevel() string`

GetAddLevel returns the AddLevel field if non-nil, zero value otherwise.

### GetAddLevelOk

`func (o *SecurityRoleSetting) GetAddLevelOk() (*string, bool)`

GetAddLevelOk returns a tuple with the AddLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddLevel

`func (o *SecurityRoleSetting) SetAddLevel(v string)`

SetAddLevel sets AddLevel field to given value.

### HasAddLevel

`func (o *SecurityRoleSetting) HasAddLevel() bool`

HasAddLevel returns a boolean if a field has been set.

### SetAddLevelNil

`func (o *SecurityRoleSetting) SetAddLevelNil(b bool)`

 SetAddLevelNil sets the value for AddLevel to be an explicit nil

### UnsetAddLevel
`func (o *SecurityRoleSetting) UnsetAddLevel()`

UnsetAddLevel ensures that no value is present for AddLevel, not even an explicit nil
### GetEditLevel

`func (o *SecurityRoleSetting) GetEditLevel() string`

GetEditLevel returns the EditLevel field if non-nil, zero value otherwise.

### GetEditLevelOk

`func (o *SecurityRoleSetting) GetEditLevelOk() (*string, bool)`

GetEditLevelOk returns a tuple with the EditLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditLevel

`func (o *SecurityRoleSetting) SetEditLevel(v string)`

SetEditLevel sets EditLevel field to given value.

### HasEditLevel

`func (o *SecurityRoleSetting) HasEditLevel() bool`

HasEditLevel returns a boolean if a field has been set.

### SetEditLevelNil

`func (o *SecurityRoleSetting) SetEditLevelNil(b bool)`

 SetEditLevelNil sets the value for EditLevel to be an explicit nil

### UnsetEditLevel
`func (o *SecurityRoleSetting) UnsetEditLevel()`

UnsetEditLevel ensures that no value is present for EditLevel, not even an explicit nil
### GetDeleteLevel

`func (o *SecurityRoleSetting) GetDeleteLevel() string`

GetDeleteLevel returns the DeleteLevel field if non-nil, zero value otherwise.

### GetDeleteLevelOk

`func (o *SecurityRoleSetting) GetDeleteLevelOk() (*string, bool)`

GetDeleteLevelOk returns a tuple with the DeleteLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteLevel

`func (o *SecurityRoleSetting) SetDeleteLevel(v string)`

SetDeleteLevel sets DeleteLevel field to given value.

### HasDeleteLevel

`func (o *SecurityRoleSetting) HasDeleteLevel() bool`

HasDeleteLevel returns a boolean if a field has been set.

### SetDeleteLevelNil

`func (o *SecurityRoleSetting) SetDeleteLevelNil(b bool)`

 SetDeleteLevelNil sets the value for DeleteLevel to be an explicit nil

### UnsetDeleteLevel
`func (o *SecurityRoleSetting) UnsetDeleteLevel()`

UnsetDeleteLevel ensures that no value is present for DeleteLevel, not even an explicit nil
### GetInquireLevel

`func (o *SecurityRoleSetting) GetInquireLevel() string`

GetInquireLevel returns the InquireLevel field if non-nil, zero value otherwise.

### GetInquireLevelOk

`func (o *SecurityRoleSetting) GetInquireLevelOk() (*string, bool)`

GetInquireLevelOk returns a tuple with the InquireLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInquireLevel

`func (o *SecurityRoleSetting) SetInquireLevel(v string)`

SetInquireLevel sets InquireLevel field to given value.

### HasInquireLevel

`func (o *SecurityRoleSetting) HasInquireLevel() bool`

HasInquireLevel returns a boolean if a field has been set.

### SetInquireLevelNil

`func (o *SecurityRoleSetting) SetInquireLevelNil(b bool)`

 SetInquireLevelNil sets the value for InquireLevel to be an explicit nil

### UnsetInquireLevel
`func (o *SecurityRoleSetting) UnsetInquireLevel()`

UnsetInquireLevel ensures that no value is present for InquireLevel, not even an explicit nil
### GetModuleFunctionName

`func (o *SecurityRoleSetting) GetModuleFunctionName() string`

GetModuleFunctionName returns the ModuleFunctionName field if non-nil, zero value otherwise.

### GetModuleFunctionNameOk

`func (o *SecurityRoleSetting) GetModuleFunctionNameOk() (*string, bool)`

GetModuleFunctionNameOk returns a tuple with the ModuleFunctionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleFunctionName

`func (o *SecurityRoleSetting) SetModuleFunctionName(v string)`

SetModuleFunctionName sets ModuleFunctionName field to given value.

### HasModuleFunctionName

`func (o *SecurityRoleSetting) HasModuleFunctionName() bool`

HasModuleFunctionName returns a boolean if a field has been set.

### GetModuleFunctionDescription

`func (o *SecurityRoleSetting) GetModuleFunctionDescription() string`

GetModuleFunctionDescription returns the ModuleFunctionDescription field if non-nil, zero value otherwise.

### GetModuleFunctionDescriptionOk

`func (o *SecurityRoleSetting) GetModuleFunctionDescriptionOk() (*string, bool)`

GetModuleFunctionDescriptionOk returns a tuple with the ModuleFunctionDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleFunctionDescription

`func (o *SecurityRoleSetting) SetModuleFunctionDescription(v string)`

SetModuleFunctionDescription sets ModuleFunctionDescription field to given value.

### HasModuleFunctionDescription

`func (o *SecurityRoleSetting) HasModuleFunctionDescription() bool`

HasModuleFunctionDescription returns a boolean if a field has been set.

### GetMyAllFlag

`func (o *SecurityRoleSetting) GetMyAllFlag() bool`

GetMyAllFlag returns the MyAllFlag field if non-nil, zero value otherwise.

### GetMyAllFlagOk

`func (o *SecurityRoleSetting) GetMyAllFlagOk() (*bool, bool)`

GetMyAllFlagOk returns a tuple with the MyAllFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMyAllFlag

`func (o *SecurityRoleSetting) SetMyAllFlag(v bool)`

SetMyAllFlag sets MyAllFlag field to given value.

### HasMyAllFlag

`func (o *SecurityRoleSetting) HasMyAllFlag() bool`

HasMyAllFlag returns a boolean if a field has been set.

### SetMyAllFlagNil

`func (o *SecurityRoleSetting) SetMyAllFlagNil(b bool)`

 SetMyAllFlagNil sets the value for MyAllFlag to be an explicit nil

### UnsetMyAllFlag
`func (o *SecurityRoleSetting) UnsetMyAllFlag()`

UnsetMyAllFlag ensures that no value is present for MyAllFlag, not even an explicit nil
### GetModuleFunctionIdentifier

`func (o *SecurityRoleSetting) GetModuleFunctionIdentifier() string`

GetModuleFunctionIdentifier returns the ModuleFunctionIdentifier field if non-nil, zero value otherwise.

### GetModuleFunctionIdentifierOk

`func (o *SecurityRoleSetting) GetModuleFunctionIdentifierOk() (*string, bool)`

GetModuleFunctionIdentifierOk returns a tuple with the ModuleFunctionIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleFunctionIdentifier

`func (o *SecurityRoleSetting) SetModuleFunctionIdentifier(v string)`

SetModuleFunctionIdentifier sets ModuleFunctionIdentifier field to given value.

### HasModuleFunctionIdentifier

`func (o *SecurityRoleSetting) HasModuleFunctionIdentifier() bool`

HasModuleFunctionIdentifier returns a boolean if a field has been set.

### GetReportFlag

`func (o *SecurityRoleSetting) GetReportFlag() bool`

GetReportFlag returns the ReportFlag field if non-nil, zero value otherwise.

### GetReportFlagOk

`func (o *SecurityRoleSetting) GetReportFlagOk() (*bool, bool)`

GetReportFlagOk returns a tuple with the ReportFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportFlag

`func (o *SecurityRoleSetting) SetReportFlag(v bool)`

SetReportFlag sets ReportFlag field to given value.

### HasReportFlag

`func (o *SecurityRoleSetting) HasReportFlag() bool`

HasReportFlag returns a boolean if a field has been set.

### SetReportFlagNil

`func (o *SecurityRoleSetting) SetReportFlagNil(b bool)`

 SetReportFlagNil sets the value for ReportFlag to be an explicit nil

### UnsetReportFlag
`func (o *SecurityRoleSetting) UnsetReportFlag()`

UnsetReportFlag ensures that no value is present for ReportFlag, not even an explicit nil
### GetRestrictFlag

`func (o *SecurityRoleSetting) GetRestrictFlag() bool`

GetRestrictFlag returns the RestrictFlag field if non-nil, zero value otherwise.

### GetRestrictFlagOk

`func (o *SecurityRoleSetting) GetRestrictFlagOk() (*bool, bool)`

GetRestrictFlagOk returns a tuple with the RestrictFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictFlag

`func (o *SecurityRoleSetting) SetRestrictFlag(v bool)`

SetRestrictFlag sets RestrictFlag field to given value.

### HasRestrictFlag

`func (o *SecurityRoleSetting) HasRestrictFlag() bool`

HasRestrictFlag returns a boolean if a field has been set.

### SetRestrictFlagNil

`func (o *SecurityRoleSetting) SetRestrictFlagNil(b bool)`

 SetRestrictFlagNil sets the value for RestrictFlag to be an explicit nil

### UnsetRestrictFlag
`func (o *SecurityRoleSetting) UnsetRestrictFlag()`

UnsetRestrictFlag ensures that no value is present for RestrictFlag, not even an explicit nil
### GetCustomFlag

`func (o *SecurityRoleSetting) GetCustomFlag() bool`

GetCustomFlag returns the CustomFlag field if non-nil, zero value otherwise.

### GetCustomFlagOk

`func (o *SecurityRoleSetting) GetCustomFlagOk() (*bool, bool)`

GetCustomFlagOk returns a tuple with the CustomFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFlag

`func (o *SecurityRoleSetting) SetCustomFlag(v bool)`

SetCustomFlag sets CustomFlag field to given value.

### HasCustomFlag

`func (o *SecurityRoleSetting) HasCustomFlag() bool`

HasCustomFlag returns a boolean if a field has been set.

### SetCustomFlagNil

`func (o *SecurityRoleSetting) SetCustomFlagNil(b bool)`

 SetCustomFlagNil sets the value for CustomFlag to be an explicit nil

### UnsetCustomFlag
`func (o *SecurityRoleSetting) UnsetCustomFlag()`

UnsetCustomFlag ensures that no value is present for CustomFlag, not even an explicit nil
### GetModuleDescription

`func (o *SecurityRoleSetting) GetModuleDescription() string`

GetModuleDescription returns the ModuleDescription field if non-nil, zero value otherwise.

### GetModuleDescriptionOk

`func (o *SecurityRoleSetting) GetModuleDescriptionOk() (*string, bool)`

GetModuleDescriptionOk returns a tuple with the ModuleDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleDescription

`func (o *SecurityRoleSetting) SetModuleDescription(v string)`

SetModuleDescription sets ModuleDescription field to given value.

### HasModuleDescription

`func (o *SecurityRoleSetting) HasModuleDescription() bool`

HasModuleDescription returns a boolean if a field has been set.

### GetModuleIdentifier

`func (o *SecurityRoleSetting) GetModuleIdentifier() string`

GetModuleIdentifier returns the ModuleIdentifier field if non-nil, zero value otherwise.

### GetModuleIdentifierOk

`func (o *SecurityRoleSetting) GetModuleIdentifierOk() (*string, bool)`

GetModuleIdentifierOk returns a tuple with the ModuleIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleIdentifier

`func (o *SecurityRoleSetting) SetModuleIdentifier(v string)`

SetModuleIdentifier sets ModuleIdentifier field to given value.

### HasModuleIdentifier

`func (o *SecurityRoleSetting) HasModuleIdentifier() bool`

HasModuleIdentifier returns a boolean if a field has been set.

### GetModuleName

`func (o *SecurityRoleSetting) GetModuleName() string`

GetModuleName returns the ModuleName field if non-nil, zero value otherwise.

### GetModuleNameOk

`func (o *SecurityRoleSetting) GetModuleNameOk() (*string, bool)`

GetModuleNameOk returns a tuple with the ModuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleName

`func (o *SecurityRoleSetting) SetModuleName(v string)`

SetModuleName sets ModuleName field to given value.

### HasModuleName

`func (o *SecurityRoleSetting) HasModuleName() bool`

HasModuleName returns a boolean if a field has been set.

### GetSortOrder

`func (o *SecurityRoleSetting) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *SecurityRoleSetting) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *SecurityRoleSetting) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *SecurityRoleSetting) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### SetSortOrderNil

`func (o *SecurityRoleSetting) SetSortOrderNil(b bool)`

 SetSortOrderNil sets the value for SortOrder to be an explicit nil

### UnsetSortOrder
`func (o *SecurityRoleSetting) UnsetSortOrder()`

UnsetSortOrder ensures that no value is present for SortOrder, not even an explicit nil
### GetInfo

`func (o *SecurityRoleSetting) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *SecurityRoleSetting) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *SecurityRoleSetting) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *SecurityRoleSetting) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


