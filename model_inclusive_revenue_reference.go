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

// checks if the InclusiveRevenueReference type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InclusiveRevenueReference{}

// InclusiveRevenueReference struct for InclusiveRevenueReference
type InclusiveRevenueReference struct {
	Id NullableInt32 `json:"id,omitempty"`
	Revenue NullableFloat64 `json:"revenue,omitempty"`
	Cost NullableFloat64 `json:"cost,omitempty"`
	Margin NullableFloat64 `json:"margin,omitempty"`
	Percentage NullableFloat64 `json:"percentage,omitempty"`
	Info *map[string]string `json:"_info,omitempty"`
}

// NewInclusiveRevenueReference instantiates a new InclusiveRevenueReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInclusiveRevenueReference() *InclusiveRevenueReference {
	this := InclusiveRevenueReference{}
	return &this
}

// NewInclusiveRevenueReferenceWithDefaults instantiates a new InclusiveRevenueReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInclusiveRevenueReferenceWithDefaults() *InclusiveRevenueReference {
	this := InclusiveRevenueReference{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InclusiveRevenueReference) GetId() int32 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int32
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InclusiveRevenueReference) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *InclusiveRevenueReference) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt32 and assigns it to the Id field.
func (o *InclusiveRevenueReference) SetId(v int32) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *InclusiveRevenueReference) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *InclusiveRevenueReference) UnsetId() {
	o.Id.Unset()
}

// GetRevenue returns the Revenue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InclusiveRevenueReference) GetRevenue() float64 {
	if o == nil || IsNil(o.Revenue.Get()) {
		var ret float64
		return ret
	}
	return *o.Revenue.Get()
}

// GetRevenueOk returns a tuple with the Revenue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InclusiveRevenueReference) GetRevenueOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Revenue.Get(), o.Revenue.IsSet()
}

// HasRevenue returns a boolean if a field has been set.
func (o *InclusiveRevenueReference) HasRevenue() bool {
	if o != nil && o.Revenue.IsSet() {
		return true
	}

	return false
}

// SetRevenue gets a reference to the given NullableFloat64 and assigns it to the Revenue field.
func (o *InclusiveRevenueReference) SetRevenue(v float64) {
	o.Revenue.Set(&v)
}
// SetRevenueNil sets the value for Revenue to be an explicit nil
func (o *InclusiveRevenueReference) SetRevenueNil() {
	o.Revenue.Set(nil)
}

// UnsetRevenue ensures that no value is present for Revenue, not even an explicit nil
func (o *InclusiveRevenueReference) UnsetRevenue() {
	o.Revenue.Unset()
}

// GetCost returns the Cost field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InclusiveRevenueReference) GetCost() float64 {
	if o == nil || IsNil(o.Cost.Get()) {
		var ret float64
		return ret
	}
	return *o.Cost.Get()
}

// GetCostOk returns a tuple with the Cost field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InclusiveRevenueReference) GetCostOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cost.Get(), o.Cost.IsSet()
}

// HasCost returns a boolean if a field has been set.
func (o *InclusiveRevenueReference) HasCost() bool {
	if o != nil && o.Cost.IsSet() {
		return true
	}

	return false
}

// SetCost gets a reference to the given NullableFloat64 and assigns it to the Cost field.
func (o *InclusiveRevenueReference) SetCost(v float64) {
	o.Cost.Set(&v)
}
// SetCostNil sets the value for Cost to be an explicit nil
func (o *InclusiveRevenueReference) SetCostNil() {
	o.Cost.Set(nil)
}

// UnsetCost ensures that no value is present for Cost, not even an explicit nil
func (o *InclusiveRevenueReference) UnsetCost() {
	o.Cost.Unset()
}

// GetMargin returns the Margin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InclusiveRevenueReference) GetMargin() float64 {
	if o == nil || IsNil(o.Margin.Get()) {
		var ret float64
		return ret
	}
	return *o.Margin.Get()
}

// GetMarginOk returns a tuple with the Margin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InclusiveRevenueReference) GetMarginOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Margin.Get(), o.Margin.IsSet()
}

// HasMargin returns a boolean if a field has been set.
func (o *InclusiveRevenueReference) HasMargin() bool {
	if o != nil && o.Margin.IsSet() {
		return true
	}

	return false
}

// SetMargin gets a reference to the given NullableFloat64 and assigns it to the Margin field.
func (o *InclusiveRevenueReference) SetMargin(v float64) {
	o.Margin.Set(&v)
}
// SetMarginNil sets the value for Margin to be an explicit nil
func (o *InclusiveRevenueReference) SetMarginNil() {
	o.Margin.Set(nil)
}

// UnsetMargin ensures that no value is present for Margin, not even an explicit nil
func (o *InclusiveRevenueReference) UnsetMargin() {
	o.Margin.Unset()
}

// GetPercentage returns the Percentage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InclusiveRevenueReference) GetPercentage() float64 {
	if o == nil || IsNil(o.Percentage.Get()) {
		var ret float64
		return ret
	}
	return *o.Percentage.Get()
}

// GetPercentageOk returns a tuple with the Percentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InclusiveRevenueReference) GetPercentageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Percentage.Get(), o.Percentage.IsSet()
}

// HasPercentage returns a boolean if a field has been set.
func (o *InclusiveRevenueReference) HasPercentage() bool {
	if o != nil && o.Percentage.IsSet() {
		return true
	}

	return false
}

// SetPercentage gets a reference to the given NullableFloat64 and assigns it to the Percentage field.
func (o *InclusiveRevenueReference) SetPercentage(v float64) {
	o.Percentage.Set(&v)
}
// SetPercentageNil sets the value for Percentage to be an explicit nil
func (o *InclusiveRevenueReference) SetPercentageNil() {
	o.Percentage.Set(nil)
}

// UnsetPercentage ensures that no value is present for Percentage, not even an explicit nil
func (o *InclusiveRevenueReference) UnsetPercentage() {
	o.Percentage.Unset()
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *InclusiveRevenueReference) GetInfo() map[string]string {
	if o == nil || IsNil(o.Info) {
		var ret map[string]string
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InclusiveRevenueReference) GetInfoOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Info) {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *InclusiveRevenueReference) HasInfo() bool {
	if o != nil && !IsNil(o.Info) {
		return true
	}

	return false
}

// SetInfo gets a reference to the given map[string]string and assigns it to the Info field.
func (o *InclusiveRevenueReference) SetInfo(v map[string]string) {
	o.Info = &v
}

func (o InclusiveRevenueReference) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InclusiveRevenueReference) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Revenue.IsSet() {
		toSerialize["revenue"] = o.Revenue.Get()
	}
	if o.Cost.IsSet() {
		toSerialize["cost"] = o.Cost.Get()
	}
	if o.Margin.IsSet() {
		toSerialize["margin"] = o.Margin.Get()
	}
	if o.Percentage.IsSet() {
		toSerialize["percentage"] = o.Percentage.Get()
	}
	if !IsNil(o.Info) {
		toSerialize["_info"] = o.Info
	}
	return toSerialize, nil
}

type NullableInclusiveRevenueReference struct {
	value *InclusiveRevenueReference
	isSet bool
}

func (v NullableInclusiveRevenueReference) Get() *InclusiveRevenueReference {
	return v.value
}

func (v *NullableInclusiveRevenueReference) Set(val *InclusiveRevenueReference) {
	v.value = val
	v.isSet = true
}

func (v NullableInclusiveRevenueReference) IsSet() bool {
	return v.isSet
}

func (v *NullableInclusiveRevenueReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInclusiveRevenueReference(val *InclusiveRevenueReference) *NullableInclusiveRevenueReference {
	return &NullableInclusiveRevenueReference{value: val, isSet: true}
}

func (v NullableInclusiveRevenueReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInclusiveRevenueReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


