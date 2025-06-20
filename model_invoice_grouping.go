/*
Connectwise Manage Public Endpoints

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2025.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cwapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the InvoiceGrouping type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvoiceGrouping{}

// InvoiceGrouping struct for InvoiceGrouping
type InvoiceGrouping struct {
	Id *int32 `json:"id,omitempty"`
	Name string `json:"name"`
	CustomerDescription string `json:"customerDescription"`
	InactiveFlag NullableBool `json:"inactiveFlag,omitempty"`
	ShowPriceFlag NullableBool `json:"showPriceFlag,omitempty"`
	ShowSubItemsFlag NullableBool `json:"showSubItemsFlag,omitempty"`
	GroupParentChildAdditionsFlag *bool `json:"groupParentChildAdditionsFlag,omitempty"`
	Info *map[string]string `json:"_info,omitempty"`
}

type _InvoiceGrouping InvoiceGrouping

// NewInvoiceGrouping instantiates a new InvoiceGrouping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoiceGrouping(name string, customerDescription string) *InvoiceGrouping {
	this := InvoiceGrouping{}
	this.Name = name
	this.CustomerDescription = customerDescription
	return &this
}

// NewInvoiceGroupingWithDefaults instantiates a new InvoiceGrouping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceGroupingWithDefaults() *InvoiceGrouping {
	this := InvoiceGrouping{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InvoiceGrouping) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceGrouping) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InvoiceGrouping) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *InvoiceGrouping) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *InvoiceGrouping) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InvoiceGrouping) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InvoiceGrouping) SetName(v string) {
	o.Name = v
}

// GetCustomerDescription returns the CustomerDescription field value
func (o *InvoiceGrouping) GetCustomerDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerDescription
}

// GetCustomerDescriptionOk returns a tuple with the CustomerDescription field value
// and a boolean to check if the value has been set.
func (o *InvoiceGrouping) GetCustomerDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerDescription, true
}

// SetCustomerDescription sets field value
func (o *InvoiceGrouping) SetCustomerDescription(v string) {
	o.CustomerDescription = v
}

// GetInactiveFlag returns the InactiveFlag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvoiceGrouping) GetInactiveFlag() bool {
	if o == nil || IsNil(o.InactiveFlag.Get()) {
		var ret bool
		return ret
	}
	return *o.InactiveFlag.Get()
}

// GetInactiveFlagOk returns a tuple with the InactiveFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceGrouping) GetInactiveFlagOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.InactiveFlag.Get(), o.InactiveFlag.IsSet()
}

// HasInactiveFlag returns a boolean if a field has been set.
func (o *InvoiceGrouping) HasInactiveFlag() bool {
	if o != nil && o.InactiveFlag.IsSet() {
		return true
	}

	return false
}

// SetInactiveFlag gets a reference to the given NullableBool and assigns it to the InactiveFlag field.
func (o *InvoiceGrouping) SetInactiveFlag(v bool) {
	o.InactiveFlag.Set(&v)
}
// SetInactiveFlagNil sets the value for InactiveFlag to be an explicit nil
func (o *InvoiceGrouping) SetInactiveFlagNil() {
	o.InactiveFlag.Set(nil)
}

// UnsetInactiveFlag ensures that no value is present for InactiveFlag, not even an explicit nil
func (o *InvoiceGrouping) UnsetInactiveFlag() {
	o.InactiveFlag.Unset()
}

// GetShowPriceFlag returns the ShowPriceFlag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvoiceGrouping) GetShowPriceFlag() bool {
	if o == nil || IsNil(o.ShowPriceFlag.Get()) {
		var ret bool
		return ret
	}
	return *o.ShowPriceFlag.Get()
}

// GetShowPriceFlagOk returns a tuple with the ShowPriceFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceGrouping) GetShowPriceFlagOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShowPriceFlag.Get(), o.ShowPriceFlag.IsSet()
}

// HasShowPriceFlag returns a boolean if a field has been set.
func (o *InvoiceGrouping) HasShowPriceFlag() bool {
	if o != nil && o.ShowPriceFlag.IsSet() {
		return true
	}

	return false
}

// SetShowPriceFlag gets a reference to the given NullableBool and assigns it to the ShowPriceFlag field.
func (o *InvoiceGrouping) SetShowPriceFlag(v bool) {
	o.ShowPriceFlag.Set(&v)
}
// SetShowPriceFlagNil sets the value for ShowPriceFlag to be an explicit nil
func (o *InvoiceGrouping) SetShowPriceFlagNil() {
	o.ShowPriceFlag.Set(nil)
}

// UnsetShowPriceFlag ensures that no value is present for ShowPriceFlag, not even an explicit nil
func (o *InvoiceGrouping) UnsetShowPriceFlag() {
	o.ShowPriceFlag.Unset()
}

// GetShowSubItemsFlag returns the ShowSubItemsFlag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvoiceGrouping) GetShowSubItemsFlag() bool {
	if o == nil || IsNil(o.ShowSubItemsFlag.Get()) {
		var ret bool
		return ret
	}
	return *o.ShowSubItemsFlag.Get()
}

// GetShowSubItemsFlagOk returns a tuple with the ShowSubItemsFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceGrouping) GetShowSubItemsFlagOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShowSubItemsFlag.Get(), o.ShowSubItemsFlag.IsSet()
}

// HasShowSubItemsFlag returns a boolean if a field has been set.
func (o *InvoiceGrouping) HasShowSubItemsFlag() bool {
	if o != nil && o.ShowSubItemsFlag.IsSet() {
		return true
	}

	return false
}

// SetShowSubItemsFlag gets a reference to the given NullableBool and assigns it to the ShowSubItemsFlag field.
func (o *InvoiceGrouping) SetShowSubItemsFlag(v bool) {
	o.ShowSubItemsFlag.Set(&v)
}
// SetShowSubItemsFlagNil sets the value for ShowSubItemsFlag to be an explicit nil
func (o *InvoiceGrouping) SetShowSubItemsFlagNil() {
	o.ShowSubItemsFlag.Set(nil)
}

// UnsetShowSubItemsFlag ensures that no value is present for ShowSubItemsFlag, not even an explicit nil
func (o *InvoiceGrouping) UnsetShowSubItemsFlag() {
	o.ShowSubItemsFlag.Unset()
}

// GetGroupParentChildAdditionsFlag returns the GroupParentChildAdditionsFlag field value if set, zero value otherwise.
func (o *InvoiceGrouping) GetGroupParentChildAdditionsFlag() bool {
	if o == nil || IsNil(o.GroupParentChildAdditionsFlag) {
		var ret bool
		return ret
	}
	return *o.GroupParentChildAdditionsFlag
}

// GetGroupParentChildAdditionsFlagOk returns a tuple with the GroupParentChildAdditionsFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceGrouping) GetGroupParentChildAdditionsFlagOk() (*bool, bool) {
	if o == nil || IsNil(o.GroupParentChildAdditionsFlag) {
		return nil, false
	}
	return o.GroupParentChildAdditionsFlag, true
}

// HasGroupParentChildAdditionsFlag returns a boolean if a field has been set.
func (o *InvoiceGrouping) HasGroupParentChildAdditionsFlag() bool {
	if o != nil && !IsNil(o.GroupParentChildAdditionsFlag) {
		return true
	}

	return false
}

// SetGroupParentChildAdditionsFlag gets a reference to the given bool and assigns it to the GroupParentChildAdditionsFlag field.
func (o *InvoiceGrouping) SetGroupParentChildAdditionsFlag(v bool) {
	o.GroupParentChildAdditionsFlag = &v
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *InvoiceGrouping) GetInfo() map[string]string {
	if o == nil || IsNil(o.Info) {
		var ret map[string]string
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceGrouping) GetInfoOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Info) {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *InvoiceGrouping) HasInfo() bool {
	if o != nil && !IsNil(o.Info) {
		return true
	}

	return false
}

// SetInfo gets a reference to the given map[string]string and assigns it to the Info field.
func (o *InvoiceGrouping) SetInfo(v map[string]string) {
	o.Info = &v
}

func (o InvoiceGrouping) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvoiceGrouping) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	toSerialize["customerDescription"] = o.CustomerDescription
	if o.InactiveFlag.IsSet() {
		toSerialize["inactiveFlag"] = o.InactiveFlag.Get()
	}
	if o.ShowPriceFlag.IsSet() {
		toSerialize["showPriceFlag"] = o.ShowPriceFlag.Get()
	}
	if o.ShowSubItemsFlag.IsSet() {
		toSerialize["showSubItemsFlag"] = o.ShowSubItemsFlag.Get()
	}
	if !IsNil(o.GroupParentChildAdditionsFlag) {
		toSerialize["groupParentChildAdditionsFlag"] = o.GroupParentChildAdditionsFlag
	}
	if !IsNil(o.Info) {
		toSerialize["_info"] = o.Info
	}
	return toSerialize, nil
}

func (o *InvoiceGrouping) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"customerDescription",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varInvoiceGrouping := _InvoiceGrouping{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInvoiceGrouping)

	if err != nil {
		return err
	}

	*o = InvoiceGrouping(varInvoiceGrouping)

	return err
}

type NullableInvoiceGrouping struct {
	value *InvoiceGrouping
	isSet bool
}

func (v NullableInvoiceGrouping) Get() *InvoiceGrouping {
	return v.value
}

func (v *NullableInvoiceGrouping) Set(val *InvoiceGrouping) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceGrouping) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceGrouping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceGrouping(val *InvoiceGrouping) *NullableInvoiceGrouping {
	return &NullableInvoiceGrouping{value: val, isSet: true}
}

func (v NullableInvoiceGrouping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceGrouping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


