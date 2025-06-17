# ChargeCodeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Location** | Pointer to [**SystemLocationReference**](SystemLocationReference.md) |  | [optional] 
**Department** | Pointer to [**SystemDepartmentReference**](SystemDepartmentReference.md) |  | [optional] 
**ExpenseEntryFlag** | Pointer to **NullableBool** |  | [optional] 
**AllowAllExpenseTypeFlag** | Pointer to **NullableBool** |  | [optional] 
**TimeEntryFlag** | Pointer to **NullableBool** |  | [optional] 
**WorkType** | Pointer to [**WorkTypeReference**](WorkTypeReference.md) |  | [optional] 
**WorkRole** | Pointer to [**WorkRoleReference**](WorkRoleReference.md) |  | [optional] 
**ExpenseTypeIds** | Pointer to **[]int32** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewChargeCodeInfo

`func NewChargeCodeInfo() *ChargeCodeInfo`

NewChargeCodeInfo instantiates a new ChargeCodeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChargeCodeInfoWithDefaults

`func NewChargeCodeInfoWithDefaults() *ChargeCodeInfo`

NewChargeCodeInfoWithDefaults instantiates a new ChargeCodeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ChargeCodeInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChargeCodeInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChargeCodeInfo) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ChargeCodeInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ChargeCodeInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ChargeCodeInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ChargeCodeInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ChargeCodeInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLocation

`func (o *ChargeCodeInfo) GetLocation() SystemLocationReference`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ChargeCodeInfo) GetLocationOk() (*SystemLocationReference, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ChargeCodeInfo) SetLocation(v SystemLocationReference)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ChargeCodeInfo) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetDepartment

`func (o *ChargeCodeInfo) GetDepartment() SystemDepartmentReference`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *ChargeCodeInfo) GetDepartmentOk() (*SystemDepartmentReference, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *ChargeCodeInfo) SetDepartment(v SystemDepartmentReference)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *ChargeCodeInfo) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### GetExpenseEntryFlag

`func (o *ChargeCodeInfo) GetExpenseEntryFlag() bool`

GetExpenseEntryFlag returns the ExpenseEntryFlag field if non-nil, zero value otherwise.

### GetExpenseEntryFlagOk

`func (o *ChargeCodeInfo) GetExpenseEntryFlagOk() (*bool, bool)`

GetExpenseEntryFlagOk returns a tuple with the ExpenseEntryFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenseEntryFlag

`func (o *ChargeCodeInfo) SetExpenseEntryFlag(v bool)`

SetExpenseEntryFlag sets ExpenseEntryFlag field to given value.

### HasExpenseEntryFlag

`func (o *ChargeCodeInfo) HasExpenseEntryFlag() bool`

HasExpenseEntryFlag returns a boolean if a field has been set.

### SetExpenseEntryFlagNil

`func (o *ChargeCodeInfo) SetExpenseEntryFlagNil(b bool)`

 SetExpenseEntryFlagNil sets the value for ExpenseEntryFlag to be an explicit nil

### UnsetExpenseEntryFlag
`func (o *ChargeCodeInfo) UnsetExpenseEntryFlag()`

UnsetExpenseEntryFlag ensures that no value is present for ExpenseEntryFlag, not even an explicit nil
### GetAllowAllExpenseTypeFlag

`func (o *ChargeCodeInfo) GetAllowAllExpenseTypeFlag() bool`

GetAllowAllExpenseTypeFlag returns the AllowAllExpenseTypeFlag field if non-nil, zero value otherwise.

### GetAllowAllExpenseTypeFlagOk

`func (o *ChargeCodeInfo) GetAllowAllExpenseTypeFlagOk() (*bool, bool)`

GetAllowAllExpenseTypeFlagOk returns a tuple with the AllowAllExpenseTypeFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAllExpenseTypeFlag

`func (o *ChargeCodeInfo) SetAllowAllExpenseTypeFlag(v bool)`

SetAllowAllExpenseTypeFlag sets AllowAllExpenseTypeFlag field to given value.

### HasAllowAllExpenseTypeFlag

`func (o *ChargeCodeInfo) HasAllowAllExpenseTypeFlag() bool`

HasAllowAllExpenseTypeFlag returns a boolean if a field has been set.

### SetAllowAllExpenseTypeFlagNil

`func (o *ChargeCodeInfo) SetAllowAllExpenseTypeFlagNil(b bool)`

 SetAllowAllExpenseTypeFlagNil sets the value for AllowAllExpenseTypeFlag to be an explicit nil

### UnsetAllowAllExpenseTypeFlag
`func (o *ChargeCodeInfo) UnsetAllowAllExpenseTypeFlag()`

UnsetAllowAllExpenseTypeFlag ensures that no value is present for AllowAllExpenseTypeFlag, not even an explicit nil
### GetTimeEntryFlag

`func (o *ChargeCodeInfo) GetTimeEntryFlag() bool`

GetTimeEntryFlag returns the TimeEntryFlag field if non-nil, zero value otherwise.

### GetTimeEntryFlagOk

`func (o *ChargeCodeInfo) GetTimeEntryFlagOk() (*bool, bool)`

GetTimeEntryFlagOk returns a tuple with the TimeEntryFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeEntryFlag

`func (o *ChargeCodeInfo) SetTimeEntryFlag(v bool)`

SetTimeEntryFlag sets TimeEntryFlag field to given value.

### HasTimeEntryFlag

`func (o *ChargeCodeInfo) HasTimeEntryFlag() bool`

HasTimeEntryFlag returns a boolean if a field has been set.

### SetTimeEntryFlagNil

`func (o *ChargeCodeInfo) SetTimeEntryFlagNil(b bool)`

 SetTimeEntryFlagNil sets the value for TimeEntryFlag to be an explicit nil

### UnsetTimeEntryFlag
`func (o *ChargeCodeInfo) UnsetTimeEntryFlag()`

UnsetTimeEntryFlag ensures that no value is present for TimeEntryFlag, not even an explicit nil
### GetWorkType

`func (o *ChargeCodeInfo) GetWorkType() WorkTypeReference`

GetWorkType returns the WorkType field if non-nil, zero value otherwise.

### GetWorkTypeOk

`func (o *ChargeCodeInfo) GetWorkTypeOk() (*WorkTypeReference, bool)`

GetWorkTypeOk returns a tuple with the WorkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkType

`func (o *ChargeCodeInfo) SetWorkType(v WorkTypeReference)`

SetWorkType sets WorkType field to given value.

### HasWorkType

`func (o *ChargeCodeInfo) HasWorkType() bool`

HasWorkType returns a boolean if a field has been set.

### GetWorkRole

`func (o *ChargeCodeInfo) GetWorkRole() WorkRoleReference`

GetWorkRole returns the WorkRole field if non-nil, zero value otherwise.

### GetWorkRoleOk

`func (o *ChargeCodeInfo) GetWorkRoleOk() (*WorkRoleReference, bool)`

GetWorkRoleOk returns a tuple with the WorkRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkRole

`func (o *ChargeCodeInfo) SetWorkRole(v WorkRoleReference)`

SetWorkRole sets WorkRole field to given value.

### HasWorkRole

`func (o *ChargeCodeInfo) HasWorkRole() bool`

HasWorkRole returns a boolean if a field has been set.

### GetExpenseTypeIds

`func (o *ChargeCodeInfo) GetExpenseTypeIds() []int32`

GetExpenseTypeIds returns the ExpenseTypeIds field if non-nil, zero value otherwise.

### GetExpenseTypeIdsOk

`func (o *ChargeCodeInfo) GetExpenseTypeIdsOk() (*[]int32, bool)`

GetExpenseTypeIdsOk returns a tuple with the ExpenseTypeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenseTypeIds

`func (o *ChargeCodeInfo) SetExpenseTypeIds(v []int32)`

SetExpenseTypeIds sets ExpenseTypeIds field to given value.

### HasExpenseTypeIds

`func (o *ChargeCodeInfo) HasExpenseTypeIds() bool`

HasExpenseTypeIds returns a boolean if a field has been set.

### GetInfo

`func (o *ChargeCodeInfo) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ChargeCodeInfo) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ChargeCodeInfo) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ChargeCodeInfo) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


