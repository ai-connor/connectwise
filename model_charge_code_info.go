/*
Connectwise Manage Public Endpoints

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2025.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cwapi

import (
	"encoding/json"
)

// checks if the ChargeCodeInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChargeCodeInfo{}

// ChargeCodeInfo struct for ChargeCodeInfo
type ChargeCodeInfo struct {
	Id *int32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Location *SystemLocationReference `json:"location,omitempty"`
	Department *SystemDepartmentReference `json:"department,omitempty"`
	ExpenseEntryFlag NullableBool `json:"expenseEntryFlag,omitempty"`
	AllowAllExpenseTypeFlag NullableBool `json:"allowAllExpenseTypeFlag,omitempty"`
	TimeEntryFlag NullableBool `json:"timeEntryFlag,omitempty"`
	WorkType *WorkTypeReference `json:"workType,omitempty"`
	WorkRole *WorkRoleReference `json:"workRole,omitempty"`
	ExpenseTypeIds []int32 `json:"expenseTypeIds,omitempty"`
	Info *map[string]string `json:"_info,omitempty"`
}

// NewChargeCodeInfo instantiates a new ChargeCodeInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChargeCodeInfo() *ChargeCodeInfo {
	this := ChargeCodeInfo{}
	return &this
}

// NewChargeCodeInfoWithDefaults instantiates a new ChargeCodeInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChargeCodeInfoWithDefaults() *ChargeCodeInfo {
	this := ChargeCodeInfo{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ChargeCodeInfo) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeCodeInfo) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ChargeCodeInfo) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ChargeCodeInfo) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ChargeCodeInfo) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeCodeInfo) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ChargeCodeInfo) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ChargeCodeInfo) SetName(v string) {
	o.Name = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *ChargeCodeInfo) GetLocation() SystemLocationReference {
	if o == nil || IsNil(o.Location) {
		var ret SystemLocationReference
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeCodeInfo) GetLocationOk() (*SystemLocationReference, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *ChargeCodeInfo) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given SystemLocationReference and assigns it to the Location field.
func (o *ChargeCodeInfo) SetLocation(v SystemLocationReference) {
	o.Location = &v
}

// GetDepartment returns the Department field value if set, zero value otherwise.
func (o *ChargeCodeInfo) GetDepartment() SystemDepartmentReference {
	if o == nil || IsNil(o.Department) {
		var ret SystemDepartmentReference
		return ret
	}
	return *o.Department
}

// GetDepartmentOk returns a tuple with the Department field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeCodeInfo) GetDepartmentOk() (*SystemDepartmentReference, bool) {
	if o == nil || IsNil(o.Department) {
		return nil, false
	}
	return o.Department, true
}

// HasDepartment returns a boolean if a field has been set.
func (o *ChargeCodeInfo) HasDepartment() bool {
	if o != nil && !IsNil(o.Department) {
		return true
	}

	return false
}

// SetDepartment gets a reference to the given SystemDepartmentReference and assigns it to the Department field.
func (o *ChargeCodeInfo) SetDepartment(v SystemDepartmentReference) {
	o.Department = &v
}

// GetExpenseEntryFlag returns the ExpenseEntryFlag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChargeCodeInfo) GetExpenseEntryFlag() bool {
	if o == nil || IsNil(o.ExpenseEntryFlag.Get()) {
		var ret bool
		return ret
	}
	return *o.ExpenseEntryFlag.Get()
}

// GetExpenseEntryFlagOk returns a tuple with the ExpenseEntryFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChargeCodeInfo) GetExpenseEntryFlagOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpenseEntryFlag.Get(), o.ExpenseEntryFlag.IsSet()
}

// HasExpenseEntryFlag returns a boolean if a field has been set.
func (o *ChargeCodeInfo) HasExpenseEntryFlag() bool {
	if o != nil && o.ExpenseEntryFlag.IsSet() {
		return true
	}

	return false
}

// SetExpenseEntryFlag gets a reference to the given NullableBool and assigns it to the ExpenseEntryFlag field.
func (o *ChargeCodeInfo) SetExpenseEntryFlag(v bool) {
	o.ExpenseEntryFlag.Set(&v)
}
// SetExpenseEntryFlagNil sets the value for ExpenseEntryFlag to be an explicit nil
func (o *ChargeCodeInfo) SetExpenseEntryFlagNil() {
	o.ExpenseEntryFlag.Set(nil)
}

// UnsetExpenseEntryFlag ensures that no value is present for ExpenseEntryFlag, not even an explicit nil
func (o *ChargeCodeInfo) UnsetExpenseEntryFlag() {
	o.ExpenseEntryFlag.Unset()
}

// GetAllowAllExpenseTypeFlag returns the AllowAllExpenseTypeFlag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChargeCodeInfo) GetAllowAllExpenseTypeFlag() bool {
	if o == nil || IsNil(o.AllowAllExpenseTypeFlag.Get()) {
		var ret bool
		return ret
	}
	return *o.AllowAllExpenseTypeFlag.Get()
}

// GetAllowAllExpenseTypeFlagOk returns a tuple with the AllowAllExpenseTypeFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChargeCodeInfo) GetAllowAllExpenseTypeFlagOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AllowAllExpenseTypeFlag.Get(), o.AllowAllExpenseTypeFlag.IsSet()
}

// HasAllowAllExpenseTypeFlag returns a boolean if a field has been set.
func (o *ChargeCodeInfo) HasAllowAllExpenseTypeFlag() bool {
	if o != nil && o.AllowAllExpenseTypeFlag.IsSet() {
		return true
	}

	return false
}

// SetAllowAllExpenseTypeFlag gets a reference to the given NullableBool and assigns it to the AllowAllExpenseTypeFlag field.
func (o *ChargeCodeInfo) SetAllowAllExpenseTypeFlag(v bool) {
	o.AllowAllExpenseTypeFlag.Set(&v)
}
// SetAllowAllExpenseTypeFlagNil sets the value for AllowAllExpenseTypeFlag to be an explicit nil
func (o *ChargeCodeInfo) SetAllowAllExpenseTypeFlagNil() {
	o.AllowAllExpenseTypeFlag.Set(nil)
}

// UnsetAllowAllExpenseTypeFlag ensures that no value is present for AllowAllExpenseTypeFlag, not even an explicit nil
func (o *ChargeCodeInfo) UnsetAllowAllExpenseTypeFlag() {
	o.AllowAllExpenseTypeFlag.Unset()
}

// GetTimeEntryFlag returns the TimeEntryFlag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChargeCodeInfo) GetTimeEntryFlag() bool {
	if o == nil || IsNil(o.TimeEntryFlag.Get()) {
		var ret bool
		return ret
	}
	return *o.TimeEntryFlag.Get()
}

// GetTimeEntryFlagOk returns a tuple with the TimeEntryFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChargeCodeInfo) GetTimeEntryFlagOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.TimeEntryFlag.Get(), o.TimeEntryFlag.IsSet()
}

// HasTimeEntryFlag returns a boolean if a field has been set.
func (o *ChargeCodeInfo) HasTimeEntryFlag() bool {
	if o != nil && o.TimeEntryFlag.IsSet() {
		return true
	}

	return false
}

// SetTimeEntryFlag gets a reference to the given NullableBool and assigns it to the TimeEntryFlag field.
func (o *ChargeCodeInfo) SetTimeEntryFlag(v bool) {
	o.TimeEntryFlag.Set(&v)
}
// SetTimeEntryFlagNil sets the value for TimeEntryFlag to be an explicit nil
func (o *ChargeCodeInfo) SetTimeEntryFlagNil() {
	o.TimeEntryFlag.Set(nil)
}

// UnsetTimeEntryFlag ensures that no value is present for TimeEntryFlag, not even an explicit nil
func (o *ChargeCodeInfo) UnsetTimeEntryFlag() {
	o.TimeEntryFlag.Unset()
}

// GetWorkType returns the WorkType field value if set, zero value otherwise.
func (o *ChargeCodeInfo) GetWorkType() WorkTypeReference {
	if o == nil || IsNil(o.WorkType) {
		var ret WorkTypeReference
		return ret
	}
	return *o.WorkType
}

// GetWorkTypeOk returns a tuple with the WorkType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeCodeInfo) GetWorkTypeOk() (*WorkTypeReference, bool) {
	if o == nil || IsNil(o.WorkType) {
		return nil, false
	}
	return o.WorkType, true
}

// HasWorkType returns a boolean if a field has been set.
func (o *ChargeCodeInfo) HasWorkType() bool {
	if o != nil && !IsNil(o.WorkType) {
		return true
	}

	return false
}

// SetWorkType gets a reference to the given WorkTypeReference and assigns it to the WorkType field.
func (o *ChargeCodeInfo) SetWorkType(v WorkTypeReference) {
	o.WorkType = &v
}

// GetWorkRole returns the WorkRole field value if set, zero value otherwise.
func (o *ChargeCodeInfo) GetWorkRole() WorkRoleReference {
	if o == nil || IsNil(o.WorkRole) {
		var ret WorkRoleReference
		return ret
	}
	return *o.WorkRole
}

// GetWorkRoleOk returns a tuple with the WorkRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeCodeInfo) GetWorkRoleOk() (*WorkRoleReference, bool) {
	if o == nil || IsNil(o.WorkRole) {
		return nil, false
	}
	return o.WorkRole, true
}

// HasWorkRole returns a boolean if a field has been set.
func (o *ChargeCodeInfo) HasWorkRole() bool {
	if o != nil && !IsNil(o.WorkRole) {
		return true
	}

	return false
}

// SetWorkRole gets a reference to the given WorkRoleReference and assigns it to the WorkRole field.
func (o *ChargeCodeInfo) SetWorkRole(v WorkRoleReference) {
	o.WorkRole = &v
}

// GetExpenseTypeIds returns the ExpenseTypeIds field value if set, zero value otherwise.
func (o *ChargeCodeInfo) GetExpenseTypeIds() []int32 {
	if o == nil || IsNil(o.ExpenseTypeIds) {
		var ret []int32
		return ret
	}
	return o.ExpenseTypeIds
}

// GetExpenseTypeIdsOk returns a tuple with the ExpenseTypeIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeCodeInfo) GetExpenseTypeIdsOk() ([]int32, bool) {
	if o == nil || IsNil(o.ExpenseTypeIds) {
		return nil, false
	}
	return o.ExpenseTypeIds, true
}

// HasExpenseTypeIds returns a boolean if a field has been set.
func (o *ChargeCodeInfo) HasExpenseTypeIds() bool {
	if o != nil && !IsNil(o.ExpenseTypeIds) {
		return true
	}

	return false
}

// SetExpenseTypeIds gets a reference to the given []int32 and assigns it to the ExpenseTypeIds field.
func (o *ChargeCodeInfo) SetExpenseTypeIds(v []int32) {
	o.ExpenseTypeIds = v
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *ChargeCodeInfo) GetInfo() map[string]string {
	if o == nil || IsNil(o.Info) {
		var ret map[string]string
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeCodeInfo) GetInfoOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Info) {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *ChargeCodeInfo) HasInfo() bool {
	if o != nil && !IsNil(o.Info) {
		return true
	}

	return false
}

// SetInfo gets a reference to the given map[string]string and assigns it to the Info field.
func (o *ChargeCodeInfo) SetInfo(v map[string]string) {
	o.Info = &v
}

func (o ChargeCodeInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChargeCodeInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !IsNil(o.Department) {
		toSerialize["department"] = o.Department
	}
	if o.ExpenseEntryFlag.IsSet() {
		toSerialize["expenseEntryFlag"] = o.ExpenseEntryFlag.Get()
	}
	if o.AllowAllExpenseTypeFlag.IsSet() {
		toSerialize["allowAllExpenseTypeFlag"] = o.AllowAllExpenseTypeFlag.Get()
	}
	if o.TimeEntryFlag.IsSet() {
		toSerialize["timeEntryFlag"] = o.TimeEntryFlag.Get()
	}
	if !IsNil(o.WorkType) {
		toSerialize["workType"] = o.WorkType
	}
	if !IsNil(o.WorkRole) {
		toSerialize["workRole"] = o.WorkRole
	}
	if !IsNil(o.ExpenseTypeIds) {
		toSerialize["expenseTypeIds"] = o.ExpenseTypeIds
	}
	if !IsNil(o.Info) {
		toSerialize["_info"] = o.Info
	}
	return toSerialize, nil
}

type NullableChargeCodeInfo struct {
	value *ChargeCodeInfo
	isSet bool
}

func (v NullableChargeCodeInfo) Get() *ChargeCodeInfo {
	return v.value
}

func (v *NullableChargeCodeInfo) Set(val *ChargeCodeInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableChargeCodeInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableChargeCodeInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChargeCodeInfo(val *ChargeCodeInfo) *NullableChargeCodeInfo {
	return &NullableChargeCodeInfo{value: val, isSet: true}
}

func (v NullableChargeCodeInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChargeCodeInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


