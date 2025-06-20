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

// checks if the ProductRecurring type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductRecurring{}

// ProductRecurring struct for ProductRecurring
type ProductRecurring struct {
	RecurringRevenue NullableFloat64 `json:"recurringRevenue,omitempty"`
	RecurringCost NullableFloat64 `json:"recurringCost,omitempty"`
	StartDate *string `json:"startDate,omitempty"`
	// The Recurring End Date is calculated based on the             start date, number of cycles, and cycle type.
	EndDate *string `json:"endDate,omitempty"`
	BillCycleId NullableInt32 `json:"billCycleId,omitempty"`
	BillCycle *BillingCycleReference `json:"billCycle,omitempty"`
	Cycles NullableInt32 `json:"cycles,omitempty"`
	CycleType NullableString `json:"cycleType,omitempty"`
	AgreementType *AgreementTypeReference `json:"agreementType,omitempty"`
}

// NewProductRecurring instantiates a new ProductRecurring object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductRecurring() *ProductRecurring {
	this := ProductRecurring{}
	return &this
}

// NewProductRecurringWithDefaults instantiates a new ProductRecurring object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductRecurringWithDefaults() *ProductRecurring {
	this := ProductRecurring{}
	return &this
}

// GetRecurringRevenue returns the RecurringRevenue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductRecurring) GetRecurringRevenue() float64 {
	if o == nil || IsNil(o.RecurringRevenue.Get()) {
		var ret float64
		return ret
	}
	return *o.RecurringRevenue.Get()
}

// GetRecurringRevenueOk returns a tuple with the RecurringRevenue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProductRecurring) GetRecurringRevenueOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RecurringRevenue.Get(), o.RecurringRevenue.IsSet()
}

// HasRecurringRevenue returns a boolean if a field has been set.
func (o *ProductRecurring) HasRecurringRevenue() bool {
	if o != nil && o.RecurringRevenue.IsSet() {
		return true
	}

	return false
}

// SetRecurringRevenue gets a reference to the given NullableFloat64 and assigns it to the RecurringRevenue field.
func (o *ProductRecurring) SetRecurringRevenue(v float64) {
	o.RecurringRevenue.Set(&v)
}
// SetRecurringRevenueNil sets the value for RecurringRevenue to be an explicit nil
func (o *ProductRecurring) SetRecurringRevenueNil() {
	o.RecurringRevenue.Set(nil)
}

// UnsetRecurringRevenue ensures that no value is present for RecurringRevenue, not even an explicit nil
func (o *ProductRecurring) UnsetRecurringRevenue() {
	o.RecurringRevenue.Unset()
}

// GetRecurringCost returns the RecurringCost field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductRecurring) GetRecurringCost() float64 {
	if o == nil || IsNil(o.RecurringCost.Get()) {
		var ret float64
		return ret
	}
	return *o.RecurringCost.Get()
}

// GetRecurringCostOk returns a tuple with the RecurringCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProductRecurring) GetRecurringCostOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RecurringCost.Get(), o.RecurringCost.IsSet()
}

// HasRecurringCost returns a boolean if a field has been set.
func (o *ProductRecurring) HasRecurringCost() bool {
	if o != nil && o.RecurringCost.IsSet() {
		return true
	}

	return false
}

// SetRecurringCost gets a reference to the given NullableFloat64 and assigns it to the RecurringCost field.
func (o *ProductRecurring) SetRecurringCost(v float64) {
	o.RecurringCost.Set(&v)
}
// SetRecurringCostNil sets the value for RecurringCost to be an explicit nil
func (o *ProductRecurring) SetRecurringCostNil() {
	o.RecurringCost.Set(nil)
}

// UnsetRecurringCost ensures that no value is present for RecurringCost, not even an explicit nil
func (o *ProductRecurring) UnsetRecurringCost() {
	o.RecurringCost.Unset()
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *ProductRecurring) GetStartDate() string {
	if o == nil || IsNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductRecurring) GetStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *ProductRecurring) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *ProductRecurring) SetStartDate(v string) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *ProductRecurring) GetEndDate() string {
	if o == nil || IsNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductRecurring) GetEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *ProductRecurring) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *ProductRecurring) SetEndDate(v string) {
	o.EndDate = &v
}

// GetBillCycleId returns the BillCycleId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductRecurring) GetBillCycleId() int32 {
	if o == nil || IsNil(o.BillCycleId.Get()) {
		var ret int32
		return ret
	}
	return *o.BillCycleId.Get()
}

// GetBillCycleIdOk returns a tuple with the BillCycleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProductRecurring) GetBillCycleIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.BillCycleId.Get(), o.BillCycleId.IsSet()
}

// HasBillCycleId returns a boolean if a field has been set.
func (o *ProductRecurring) HasBillCycleId() bool {
	if o != nil && o.BillCycleId.IsSet() {
		return true
	}

	return false
}

// SetBillCycleId gets a reference to the given NullableInt32 and assigns it to the BillCycleId field.
func (o *ProductRecurring) SetBillCycleId(v int32) {
	o.BillCycleId.Set(&v)
}
// SetBillCycleIdNil sets the value for BillCycleId to be an explicit nil
func (o *ProductRecurring) SetBillCycleIdNil() {
	o.BillCycleId.Set(nil)
}

// UnsetBillCycleId ensures that no value is present for BillCycleId, not even an explicit nil
func (o *ProductRecurring) UnsetBillCycleId() {
	o.BillCycleId.Unset()
}

// GetBillCycle returns the BillCycle field value if set, zero value otherwise.
func (o *ProductRecurring) GetBillCycle() BillingCycleReference {
	if o == nil || IsNil(o.BillCycle) {
		var ret BillingCycleReference
		return ret
	}
	return *o.BillCycle
}

// GetBillCycleOk returns a tuple with the BillCycle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductRecurring) GetBillCycleOk() (*BillingCycleReference, bool) {
	if o == nil || IsNil(o.BillCycle) {
		return nil, false
	}
	return o.BillCycle, true
}

// HasBillCycle returns a boolean if a field has been set.
func (o *ProductRecurring) HasBillCycle() bool {
	if o != nil && !IsNil(o.BillCycle) {
		return true
	}

	return false
}

// SetBillCycle gets a reference to the given BillingCycleReference and assigns it to the BillCycle field.
func (o *ProductRecurring) SetBillCycle(v BillingCycleReference) {
	o.BillCycle = &v
}

// GetCycles returns the Cycles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductRecurring) GetCycles() int32 {
	if o == nil || IsNil(o.Cycles.Get()) {
		var ret int32
		return ret
	}
	return *o.Cycles.Get()
}

// GetCyclesOk returns a tuple with the Cycles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProductRecurring) GetCyclesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cycles.Get(), o.Cycles.IsSet()
}

// HasCycles returns a boolean if a field has been set.
func (o *ProductRecurring) HasCycles() bool {
	if o != nil && o.Cycles.IsSet() {
		return true
	}

	return false
}

// SetCycles gets a reference to the given NullableInt32 and assigns it to the Cycles field.
func (o *ProductRecurring) SetCycles(v int32) {
	o.Cycles.Set(&v)
}
// SetCyclesNil sets the value for Cycles to be an explicit nil
func (o *ProductRecurring) SetCyclesNil() {
	o.Cycles.Set(nil)
}

// UnsetCycles ensures that no value is present for Cycles, not even an explicit nil
func (o *ProductRecurring) UnsetCycles() {
	o.Cycles.Unset()
}

// GetCycleType returns the CycleType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductRecurring) GetCycleType() string {
	if o == nil || IsNil(o.CycleType.Get()) {
		var ret string
		return ret
	}
	return *o.CycleType.Get()
}

// GetCycleTypeOk returns a tuple with the CycleType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProductRecurring) GetCycleTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CycleType.Get(), o.CycleType.IsSet()
}

// HasCycleType returns a boolean if a field has been set.
func (o *ProductRecurring) HasCycleType() bool {
	if o != nil && o.CycleType.IsSet() {
		return true
	}

	return false
}

// SetCycleType gets a reference to the given NullableString and assigns it to the CycleType field.
func (o *ProductRecurring) SetCycleType(v string) {
	o.CycleType.Set(&v)
}
// SetCycleTypeNil sets the value for CycleType to be an explicit nil
func (o *ProductRecurring) SetCycleTypeNil() {
	o.CycleType.Set(nil)
}

// UnsetCycleType ensures that no value is present for CycleType, not even an explicit nil
func (o *ProductRecurring) UnsetCycleType() {
	o.CycleType.Unset()
}

// GetAgreementType returns the AgreementType field value if set, zero value otherwise.
func (o *ProductRecurring) GetAgreementType() AgreementTypeReference {
	if o == nil || IsNil(o.AgreementType) {
		var ret AgreementTypeReference
		return ret
	}
	return *o.AgreementType
}

// GetAgreementTypeOk returns a tuple with the AgreementType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductRecurring) GetAgreementTypeOk() (*AgreementTypeReference, bool) {
	if o == nil || IsNil(o.AgreementType) {
		return nil, false
	}
	return o.AgreementType, true
}

// HasAgreementType returns a boolean if a field has been set.
func (o *ProductRecurring) HasAgreementType() bool {
	if o != nil && !IsNil(o.AgreementType) {
		return true
	}

	return false
}

// SetAgreementType gets a reference to the given AgreementTypeReference and assigns it to the AgreementType field.
func (o *ProductRecurring) SetAgreementType(v AgreementTypeReference) {
	o.AgreementType = &v
}

func (o ProductRecurring) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductRecurring) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.RecurringRevenue.IsSet() {
		toSerialize["recurringRevenue"] = o.RecurringRevenue.Get()
	}
	if o.RecurringCost.IsSet() {
		toSerialize["recurringCost"] = o.RecurringCost.Get()
	}
	if !IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	if !IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	if o.BillCycleId.IsSet() {
		toSerialize["billCycleId"] = o.BillCycleId.Get()
	}
	if !IsNil(o.BillCycle) {
		toSerialize["billCycle"] = o.BillCycle
	}
	if o.Cycles.IsSet() {
		toSerialize["cycles"] = o.Cycles.Get()
	}
	if o.CycleType.IsSet() {
		toSerialize["cycleType"] = o.CycleType.Get()
	}
	if !IsNil(o.AgreementType) {
		toSerialize["agreementType"] = o.AgreementType
	}
	return toSerialize, nil
}

type NullableProductRecurring struct {
	value *ProductRecurring
	isSet bool
}

func (v NullableProductRecurring) Get() *ProductRecurring {
	return v.value
}

func (v *NullableProductRecurring) Set(val *ProductRecurring) {
	v.value = val
	v.isSet = true
}

func (v NullableProductRecurring) IsSet() bool {
	return v.isSet
}

func (v *NullableProductRecurring) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductRecurring(val *ProductRecurring) *NullableProductRecurring {
	return &NullableProductRecurring{value: val, isSet: true}
}

func (v NullableProductRecurring) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductRecurring) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


