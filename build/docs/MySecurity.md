# MySecurity

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
**Member** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewMySecurity

`func NewMySecurity() *MySecurity`

NewMySecurity instantiates a new MySecurity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMySecurityWithDefaults

`func NewMySecurityWithDefaults() *MySecurity`

NewMySecurityWithDefaults instantiates a new MySecurity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MySecurity) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MySecurity) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MySecurity) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MySecurity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAddLevel

`func (o *MySecurity) GetAddLevel() string`

GetAddLevel returns the AddLevel field if non-nil, zero value otherwise.

### GetAddLevelOk

`func (o *MySecurity) GetAddLevelOk() (*string, bool)`

GetAddLevelOk returns a tuple with the AddLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddLevel

`func (o *MySecurity) SetAddLevel(v string)`

SetAddLevel sets AddLevel field to given value.

### HasAddLevel

`func (o *MySecurity) HasAddLevel() bool`

HasAddLevel returns a boolean if a field has been set.

### SetAddLevelNil

`func (o *MySecurity) SetAddLevelNil(b bool)`

 SetAddLevelNil sets the value for AddLevel to be an explicit nil

### UnsetAddLevel
`func (o *MySecurity) UnsetAddLevel()`

UnsetAddLevel ensures that no value is present for AddLevel, not even an explicit nil
### GetEditLevel

`func (o *MySecurity) GetEditLevel() string`

GetEditLevel returns the EditLevel field if non-nil, zero value otherwise.

### GetEditLevelOk

`func (o *MySecurity) GetEditLevelOk() (*string, bool)`

GetEditLevelOk returns a tuple with the EditLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditLevel

`func (o *MySecurity) SetEditLevel(v string)`

SetEditLevel sets EditLevel field to given value.

### HasEditLevel

`func (o *MySecurity) HasEditLevel() bool`

HasEditLevel returns a boolean if a field has been set.

### SetEditLevelNil

`func (o *MySecurity) SetEditLevelNil(b bool)`

 SetEditLevelNil sets the value for EditLevel to be an explicit nil

### UnsetEditLevel
`func (o *MySecurity) UnsetEditLevel()`

UnsetEditLevel ensures that no value is present for EditLevel, not even an explicit nil
### GetDeleteLevel

`func (o *MySecurity) GetDeleteLevel() string`

GetDeleteLevel returns the DeleteLevel field if non-nil, zero value otherwise.

### GetDeleteLevelOk

`func (o *MySecurity) GetDeleteLevelOk() (*string, bool)`

GetDeleteLevelOk returns a tuple with the DeleteLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteLevel

`func (o *MySecurity) SetDeleteLevel(v string)`

SetDeleteLevel sets DeleteLevel field to given value.

### HasDeleteLevel

`func (o *MySecurity) HasDeleteLevel() bool`

HasDeleteLevel returns a boolean if a field has been set.

### SetDeleteLevelNil

`func (o *MySecurity) SetDeleteLevelNil(b bool)`

 SetDeleteLevelNil sets the value for DeleteLevel to be an explicit nil

### UnsetDeleteLevel
`func (o *MySecurity) UnsetDeleteLevel()`

UnsetDeleteLevel ensures that no value is present for DeleteLevel, not even an explicit nil
### GetInquireLevel

`func (o *MySecurity) GetInquireLevel() string`

GetInquireLevel returns the InquireLevel field if non-nil, zero value otherwise.

### GetInquireLevelOk

`func (o *MySecurity) GetInquireLevelOk() (*string, bool)`

GetInquireLevelOk returns a tuple with the InquireLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInquireLevel

`func (o *MySecurity) SetInquireLevel(v string)`

SetInquireLevel sets InquireLevel field to given value.

### HasInquireLevel

`func (o *MySecurity) HasInquireLevel() bool`

HasInquireLevel returns a boolean if a field has been set.

### SetInquireLevelNil

`func (o *MySecurity) SetInquireLevelNil(b bool)`

 SetInquireLevelNil sets the value for InquireLevel to be an explicit nil

### UnsetInquireLevel
`func (o *MySecurity) UnsetInquireLevel()`

UnsetInquireLevel ensures that no value is present for InquireLevel, not even an explicit nil
### GetModuleFunctionName

`func (o *MySecurity) GetModuleFunctionName() string`

GetModuleFunctionName returns the ModuleFunctionName field if non-nil, zero value otherwise.

### GetModuleFunctionNameOk

`func (o *MySecurity) GetModuleFunctionNameOk() (*string, bool)`

GetModuleFunctionNameOk returns a tuple with the ModuleFunctionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleFunctionName

`func (o *MySecurity) SetModuleFunctionName(v string)`

SetModuleFunctionName sets ModuleFunctionName field to given value.

### HasModuleFunctionName

`func (o *MySecurity) HasModuleFunctionName() bool`

HasModuleFunctionName returns a boolean if a field has been set.

### GetModuleFunctionDescription

`func (o *MySecurity) GetModuleFunctionDescription() string`

GetModuleFunctionDescription returns the ModuleFunctionDescription field if non-nil, zero value otherwise.

### GetModuleFunctionDescriptionOk

`func (o *MySecurity) GetModuleFunctionDescriptionOk() (*string, bool)`

GetModuleFunctionDescriptionOk returns a tuple with the ModuleFunctionDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleFunctionDescription

`func (o *MySecurity) SetModuleFunctionDescription(v string)`

SetModuleFunctionDescription sets ModuleFunctionDescription field to given value.

### HasModuleFunctionDescription

`func (o *MySecurity) HasModuleFunctionDescription() bool`

HasModuleFunctionDescription returns a boolean if a field has been set.

### GetMyAllFlag

`func (o *MySecurity) GetMyAllFlag() bool`

GetMyAllFlag returns the MyAllFlag field if non-nil, zero value otherwise.

### GetMyAllFlagOk

`func (o *MySecurity) GetMyAllFlagOk() (*bool, bool)`

GetMyAllFlagOk returns a tuple with the MyAllFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMyAllFlag

`func (o *MySecurity) SetMyAllFlag(v bool)`

SetMyAllFlag sets MyAllFlag field to given value.

### HasMyAllFlag

`func (o *MySecurity) HasMyAllFlag() bool`

HasMyAllFlag returns a boolean if a field has been set.

### SetMyAllFlagNil

`func (o *MySecurity) SetMyAllFlagNil(b bool)`

 SetMyAllFlagNil sets the value for MyAllFlag to be an explicit nil

### UnsetMyAllFlag
`func (o *MySecurity) UnsetMyAllFlag()`

UnsetMyAllFlag ensures that no value is present for MyAllFlag, not even an explicit nil
### GetModuleFunctionIdentifier

`func (o *MySecurity) GetModuleFunctionIdentifier() string`

GetModuleFunctionIdentifier returns the ModuleFunctionIdentifier field if non-nil, zero value otherwise.

### GetModuleFunctionIdentifierOk

`func (o *MySecurity) GetModuleFunctionIdentifierOk() (*string, bool)`

GetModuleFunctionIdentifierOk returns a tuple with the ModuleFunctionIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleFunctionIdentifier

`func (o *MySecurity) SetModuleFunctionIdentifier(v string)`

SetModuleFunctionIdentifier sets ModuleFunctionIdentifier field to given value.

### HasModuleFunctionIdentifier

`func (o *MySecurity) HasModuleFunctionIdentifier() bool`

HasModuleFunctionIdentifier returns a boolean if a field has been set.

### GetReportFlag

`func (o *MySecurity) GetReportFlag() bool`

GetReportFlag returns the ReportFlag field if non-nil, zero value otherwise.

### GetReportFlagOk

`func (o *MySecurity) GetReportFlagOk() (*bool, bool)`

GetReportFlagOk returns a tuple with the ReportFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportFlag

`func (o *MySecurity) SetReportFlag(v bool)`

SetReportFlag sets ReportFlag field to given value.

### HasReportFlag

`func (o *MySecurity) HasReportFlag() bool`

HasReportFlag returns a boolean if a field has been set.

### SetReportFlagNil

`func (o *MySecurity) SetReportFlagNil(b bool)`

 SetReportFlagNil sets the value for ReportFlag to be an explicit nil

### UnsetReportFlag
`func (o *MySecurity) UnsetReportFlag()`

UnsetReportFlag ensures that no value is present for ReportFlag, not even an explicit nil
### GetRestrictFlag

`func (o *MySecurity) GetRestrictFlag() bool`

GetRestrictFlag returns the RestrictFlag field if non-nil, zero value otherwise.

### GetRestrictFlagOk

`func (o *MySecurity) GetRestrictFlagOk() (*bool, bool)`

GetRestrictFlagOk returns a tuple with the RestrictFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictFlag

`func (o *MySecurity) SetRestrictFlag(v bool)`

SetRestrictFlag sets RestrictFlag field to given value.

### HasRestrictFlag

`func (o *MySecurity) HasRestrictFlag() bool`

HasRestrictFlag returns a boolean if a field has been set.

### SetRestrictFlagNil

`func (o *MySecurity) SetRestrictFlagNil(b bool)`

 SetRestrictFlagNil sets the value for RestrictFlag to be an explicit nil

### UnsetRestrictFlag
`func (o *MySecurity) UnsetRestrictFlag()`

UnsetRestrictFlag ensures that no value is present for RestrictFlag, not even an explicit nil
### GetCustomFlag

`func (o *MySecurity) GetCustomFlag() bool`

GetCustomFlag returns the CustomFlag field if non-nil, zero value otherwise.

### GetCustomFlagOk

`func (o *MySecurity) GetCustomFlagOk() (*bool, bool)`

GetCustomFlagOk returns a tuple with the CustomFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFlag

`func (o *MySecurity) SetCustomFlag(v bool)`

SetCustomFlag sets CustomFlag field to given value.

### HasCustomFlag

`func (o *MySecurity) HasCustomFlag() bool`

HasCustomFlag returns a boolean if a field has been set.

### SetCustomFlagNil

`func (o *MySecurity) SetCustomFlagNil(b bool)`

 SetCustomFlagNil sets the value for CustomFlag to be an explicit nil

### UnsetCustomFlag
`func (o *MySecurity) UnsetCustomFlag()`

UnsetCustomFlag ensures that no value is present for CustomFlag, not even an explicit nil
### GetModuleDescription

`func (o *MySecurity) GetModuleDescription() string`

GetModuleDescription returns the ModuleDescription field if non-nil, zero value otherwise.

### GetModuleDescriptionOk

`func (o *MySecurity) GetModuleDescriptionOk() (*string, bool)`

GetModuleDescriptionOk returns a tuple with the ModuleDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleDescription

`func (o *MySecurity) SetModuleDescription(v string)`

SetModuleDescription sets ModuleDescription field to given value.

### HasModuleDescription

`func (o *MySecurity) HasModuleDescription() bool`

HasModuleDescription returns a boolean if a field has been set.

### GetModuleIdentifier

`func (o *MySecurity) GetModuleIdentifier() string`

GetModuleIdentifier returns the ModuleIdentifier field if non-nil, zero value otherwise.

### GetModuleIdentifierOk

`func (o *MySecurity) GetModuleIdentifierOk() (*string, bool)`

GetModuleIdentifierOk returns a tuple with the ModuleIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleIdentifier

`func (o *MySecurity) SetModuleIdentifier(v string)`

SetModuleIdentifier sets ModuleIdentifier field to given value.

### HasModuleIdentifier

`func (o *MySecurity) HasModuleIdentifier() bool`

HasModuleIdentifier returns a boolean if a field has been set.

### GetModuleName

`func (o *MySecurity) GetModuleName() string`

GetModuleName returns the ModuleName field if non-nil, zero value otherwise.

### GetModuleNameOk

`func (o *MySecurity) GetModuleNameOk() (*string, bool)`

GetModuleNameOk returns a tuple with the ModuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleName

`func (o *MySecurity) SetModuleName(v string)`

SetModuleName sets ModuleName field to given value.

### HasModuleName

`func (o *MySecurity) HasModuleName() bool`

HasModuleName returns a boolean if a field has been set.

### GetSortOrder

`func (o *MySecurity) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *MySecurity) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *MySecurity) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *MySecurity) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### SetSortOrderNil

`func (o *MySecurity) SetSortOrderNil(b bool)`

 SetSortOrderNil sets the value for SortOrder to be an explicit nil

### UnsetSortOrder
`func (o *MySecurity) UnsetSortOrder()`

UnsetSortOrder ensures that no value is present for SortOrder, not even an explicit nil
### GetMember

`func (o *MySecurity) GetMember() MemberReference`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *MySecurity) GetMemberOk() (*MemberReference, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *MySecurity) SetMember(v MemberReference)`

SetMember sets Member field to given value.

### HasMember

`func (o *MySecurity) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetInfo

`func (o *MySecurity) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *MySecurity) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *MySecurity) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *MySecurity) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


