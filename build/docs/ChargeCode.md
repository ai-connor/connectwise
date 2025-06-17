# ChargeCode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  Max length: 50; | 
**Company** | [**CompanyReference**](CompanyReference.md) |  | 
**Location** | Pointer to [**SystemLocationReference**](SystemLocationReference.md) |  | [optional] 
**Department** | Pointer to [**SystemDepartmentReference**](SystemDepartmentReference.md) |  | [optional] 
**BillTime** | Pointer to **NullableString** |  | [optional] 
**ExpenseEntryFlag** | Pointer to **NullableBool** |  | [optional] 
**AllowAllExpenseTypeFlag** | Pointer to **NullableBool** |  | [optional] 
**TimeEntryFlag** | Pointer to **NullableBool** |  | [optional] 
**WorkType** | Pointer to [**WorkTypeReference**](WorkTypeReference.md) |  | [optional] 
**WorkRole** | Pointer to [**WorkRoleReference**](WorkRoleReference.md) |  | [optional] 
**IntegrationXref** | Pointer to **string** |  Max length: 50; | [optional] 
**ExpenseTypeIds** | Pointer to **[]int32** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewChargeCode

`func NewChargeCode(name string, company CompanyReference, ) *ChargeCode`

NewChargeCode instantiates a new ChargeCode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChargeCodeWithDefaults

`func NewChargeCodeWithDefaults() *ChargeCode`

NewChargeCodeWithDefaults instantiates a new ChargeCode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ChargeCode) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChargeCode) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChargeCode) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ChargeCode) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ChargeCode) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ChargeCode) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ChargeCode) SetName(v string)`

SetName sets Name field to given value.


### GetCompany

`func (o *ChargeCode) GetCompany() CompanyReference`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *ChargeCode) GetCompanyOk() (*CompanyReference, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *ChargeCode) SetCompany(v CompanyReference)`

SetCompany sets Company field to given value.


### GetLocation

`func (o *ChargeCode) GetLocation() SystemLocationReference`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ChargeCode) GetLocationOk() (*SystemLocationReference, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ChargeCode) SetLocation(v SystemLocationReference)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ChargeCode) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetDepartment

`func (o *ChargeCode) GetDepartment() SystemDepartmentReference`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *ChargeCode) GetDepartmentOk() (*SystemDepartmentReference, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *ChargeCode) SetDepartment(v SystemDepartmentReference)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *ChargeCode) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### GetBillTime

`func (o *ChargeCode) GetBillTime() string`

GetBillTime returns the BillTime field if non-nil, zero value otherwise.

### GetBillTimeOk

`func (o *ChargeCode) GetBillTimeOk() (*string, bool)`

GetBillTimeOk returns a tuple with the BillTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillTime

`func (o *ChargeCode) SetBillTime(v string)`

SetBillTime sets BillTime field to given value.

### HasBillTime

`func (o *ChargeCode) HasBillTime() bool`

HasBillTime returns a boolean if a field has been set.

### SetBillTimeNil

`func (o *ChargeCode) SetBillTimeNil(b bool)`

 SetBillTimeNil sets the value for BillTime to be an explicit nil

### UnsetBillTime
`func (o *ChargeCode) UnsetBillTime()`

UnsetBillTime ensures that no value is present for BillTime, not even an explicit nil
### GetExpenseEntryFlag

`func (o *ChargeCode) GetExpenseEntryFlag() bool`

GetExpenseEntryFlag returns the ExpenseEntryFlag field if non-nil, zero value otherwise.

### GetExpenseEntryFlagOk

`func (o *ChargeCode) GetExpenseEntryFlagOk() (*bool, bool)`

GetExpenseEntryFlagOk returns a tuple with the ExpenseEntryFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenseEntryFlag

`func (o *ChargeCode) SetExpenseEntryFlag(v bool)`

SetExpenseEntryFlag sets ExpenseEntryFlag field to given value.

### HasExpenseEntryFlag

`func (o *ChargeCode) HasExpenseEntryFlag() bool`

HasExpenseEntryFlag returns a boolean if a field has been set.

### SetExpenseEntryFlagNil

`func (o *ChargeCode) SetExpenseEntryFlagNil(b bool)`

 SetExpenseEntryFlagNil sets the value for ExpenseEntryFlag to be an explicit nil

### UnsetExpenseEntryFlag
`func (o *ChargeCode) UnsetExpenseEntryFlag()`

UnsetExpenseEntryFlag ensures that no value is present for ExpenseEntryFlag, not even an explicit nil
### GetAllowAllExpenseTypeFlag

`func (o *ChargeCode) GetAllowAllExpenseTypeFlag() bool`

GetAllowAllExpenseTypeFlag returns the AllowAllExpenseTypeFlag field if non-nil, zero value otherwise.

### GetAllowAllExpenseTypeFlagOk

`func (o *ChargeCode) GetAllowAllExpenseTypeFlagOk() (*bool, bool)`

GetAllowAllExpenseTypeFlagOk returns a tuple with the AllowAllExpenseTypeFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAllExpenseTypeFlag

`func (o *ChargeCode) SetAllowAllExpenseTypeFlag(v bool)`

SetAllowAllExpenseTypeFlag sets AllowAllExpenseTypeFlag field to given value.

### HasAllowAllExpenseTypeFlag

`func (o *ChargeCode) HasAllowAllExpenseTypeFlag() bool`

HasAllowAllExpenseTypeFlag returns a boolean if a field has been set.

### SetAllowAllExpenseTypeFlagNil

`func (o *ChargeCode) SetAllowAllExpenseTypeFlagNil(b bool)`

 SetAllowAllExpenseTypeFlagNil sets the value for AllowAllExpenseTypeFlag to be an explicit nil

### UnsetAllowAllExpenseTypeFlag
`func (o *ChargeCode) UnsetAllowAllExpenseTypeFlag()`

UnsetAllowAllExpenseTypeFlag ensures that no value is present for AllowAllExpenseTypeFlag, not even an explicit nil
### GetTimeEntryFlag

`func (o *ChargeCode) GetTimeEntryFlag() bool`

GetTimeEntryFlag returns the TimeEntryFlag field if non-nil, zero value otherwise.

### GetTimeEntryFlagOk

`func (o *ChargeCode) GetTimeEntryFlagOk() (*bool, bool)`

GetTimeEntryFlagOk returns a tuple with the TimeEntryFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeEntryFlag

`func (o *ChargeCode) SetTimeEntryFlag(v bool)`

SetTimeEntryFlag sets TimeEntryFlag field to given value.

### HasTimeEntryFlag

`func (o *ChargeCode) HasTimeEntryFlag() bool`

HasTimeEntryFlag returns a boolean if a field has been set.

### SetTimeEntryFlagNil

`func (o *ChargeCode) SetTimeEntryFlagNil(b bool)`

 SetTimeEntryFlagNil sets the value for TimeEntryFlag to be an explicit nil

### UnsetTimeEntryFlag
`func (o *ChargeCode) UnsetTimeEntryFlag()`

UnsetTimeEntryFlag ensures that no value is present for TimeEntryFlag, not even an explicit nil
### GetWorkType

`func (o *ChargeCode) GetWorkType() WorkTypeReference`

GetWorkType returns the WorkType field if non-nil, zero value otherwise.

### GetWorkTypeOk

`func (o *ChargeCode) GetWorkTypeOk() (*WorkTypeReference, bool)`

GetWorkTypeOk returns a tuple with the WorkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkType

`func (o *ChargeCode) SetWorkType(v WorkTypeReference)`

SetWorkType sets WorkType field to given value.

### HasWorkType

`func (o *ChargeCode) HasWorkType() bool`

HasWorkType returns a boolean if a field has been set.

### GetWorkRole

`func (o *ChargeCode) GetWorkRole() WorkRoleReference`

GetWorkRole returns the WorkRole field if non-nil, zero value otherwise.

### GetWorkRoleOk

`func (o *ChargeCode) GetWorkRoleOk() (*WorkRoleReference, bool)`

GetWorkRoleOk returns a tuple with the WorkRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkRole

`func (o *ChargeCode) SetWorkRole(v WorkRoleReference)`

SetWorkRole sets WorkRole field to given value.

### HasWorkRole

`func (o *ChargeCode) HasWorkRole() bool`

HasWorkRole returns a boolean if a field has been set.

### GetIntegrationXref

`func (o *ChargeCode) GetIntegrationXref() string`

GetIntegrationXref returns the IntegrationXref field if non-nil, zero value otherwise.

### GetIntegrationXrefOk

`func (o *ChargeCode) GetIntegrationXrefOk() (*string, bool)`

GetIntegrationXrefOk returns a tuple with the IntegrationXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationXref

`func (o *ChargeCode) SetIntegrationXref(v string)`

SetIntegrationXref sets IntegrationXref field to given value.

### HasIntegrationXref

`func (o *ChargeCode) HasIntegrationXref() bool`

HasIntegrationXref returns a boolean if a field has been set.

### GetExpenseTypeIds

`func (o *ChargeCode) GetExpenseTypeIds() []int32`

GetExpenseTypeIds returns the ExpenseTypeIds field if non-nil, zero value otherwise.

### GetExpenseTypeIdsOk

`func (o *ChargeCode) GetExpenseTypeIdsOk() (*[]int32, bool)`

GetExpenseTypeIdsOk returns a tuple with the ExpenseTypeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenseTypeIds

`func (o *ChargeCode) SetExpenseTypeIds(v []int32)`

SetExpenseTypeIds sets ExpenseTypeIds field to given value.

### HasExpenseTypeIds

`func (o *ChargeCode) HasExpenseTypeIds() bool`

HasExpenseTypeIds returns a boolean if a field has been set.

### GetInfo

`func (o *ChargeCode) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ChargeCode) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ChargeCode) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ChargeCode) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


