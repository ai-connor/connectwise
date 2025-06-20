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

// checks if the BillingStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillingStatus{}

// BillingStatus struct for BillingStatus
type BillingStatus struct {
	Id *int32 `json:"id,omitempty"`
	//  Max length: 30;
	Name string `json:"name"`
	SortOrder NullableInt32 `json:"sortOrder,omitempty"`
	DefaultFlag NullableBool `json:"defaultFlag,omitempty"`
	ClosedFlag NullableBool `json:"closedFlag,omitempty"`
	InactiveFlag NullableBool `json:"inactiveFlag,omitempty"`
	SentFlag NullableBool `json:"sentFlag,omitempty"`
	Info *map[string]string `json:"_info,omitempty"`
}

type _BillingStatus BillingStatus

// NewBillingStatus instantiates a new BillingStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingStatus(name string) *BillingStatus {
	this := BillingStatus{}
	this.Name = name
	return &this
}

// NewBillingStatusWithDefaults instantiates a new BillingStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingStatusWithDefaults() *BillingStatus {
	this := BillingStatus{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BillingStatus) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingStatus) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BillingStatus) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *BillingStatus) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *BillingStatus) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BillingStatus) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BillingStatus) SetName(v string) {
	o.Name = v
}

// GetSortOrder returns the SortOrder field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingStatus) GetSortOrder() int32 {
	if o == nil || IsNil(o.SortOrder.Get()) {
		var ret int32
		return ret
	}
	return *o.SortOrder.Get()
}

// GetSortOrderOk returns a tuple with the SortOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingStatus) GetSortOrderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SortOrder.Get(), o.SortOrder.IsSet()
}

// HasSortOrder returns a boolean if a field has been set.
func (o *BillingStatus) HasSortOrder() bool {
	if o != nil && o.SortOrder.IsSet() {
		return true
	}

	return false
}

// SetSortOrder gets a reference to the given NullableInt32 and assigns it to the SortOrder field.
func (o *BillingStatus) SetSortOrder(v int32) {
	o.SortOrder.Set(&v)
}
// SetSortOrderNil sets the value for SortOrder to be an explicit nil
func (o *BillingStatus) SetSortOrderNil() {
	o.SortOrder.Set(nil)
}

// UnsetSortOrder ensures that no value is present for SortOrder, not even an explicit nil
func (o *BillingStatus) UnsetSortOrder() {
	o.SortOrder.Unset()
}

// GetDefaultFlag returns the DefaultFlag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingStatus) GetDefaultFlag() bool {
	if o == nil || IsNil(o.DefaultFlag.Get()) {
		var ret bool
		return ret
	}
	return *o.DefaultFlag.Get()
}

// GetDefaultFlagOk returns a tuple with the DefaultFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingStatus) GetDefaultFlagOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultFlag.Get(), o.DefaultFlag.IsSet()
}

// HasDefaultFlag returns a boolean if a field has been set.
func (o *BillingStatus) HasDefaultFlag() bool {
	if o != nil && o.DefaultFlag.IsSet() {
		return true
	}

	return false
}

// SetDefaultFlag gets a reference to the given NullableBool and assigns it to the DefaultFlag field.
func (o *BillingStatus) SetDefaultFlag(v bool) {
	o.DefaultFlag.Set(&v)
}
// SetDefaultFlagNil sets the value for DefaultFlag to be an explicit nil
func (o *BillingStatus) SetDefaultFlagNil() {
	o.DefaultFlag.Set(nil)
}

// UnsetDefaultFlag ensures that no value is present for DefaultFlag, not even an explicit nil
func (o *BillingStatus) UnsetDefaultFlag() {
	o.DefaultFlag.Unset()
}

// GetClosedFlag returns the ClosedFlag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingStatus) GetClosedFlag() bool {
	if o == nil || IsNil(o.ClosedFlag.Get()) {
		var ret bool
		return ret
	}
	return *o.ClosedFlag.Get()
}

// GetClosedFlagOk returns a tuple with the ClosedFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingStatus) GetClosedFlagOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClosedFlag.Get(), o.ClosedFlag.IsSet()
}

// HasClosedFlag returns a boolean if a field has been set.
func (o *BillingStatus) HasClosedFlag() bool {
	if o != nil && o.ClosedFlag.IsSet() {
		return true
	}

	return false
}

// SetClosedFlag gets a reference to the given NullableBool and assigns it to the ClosedFlag field.
func (o *BillingStatus) SetClosedFlag(v bool) {
	o.ClosedFlag.Set(&v)
}
// SetClosedFlagNil sets the value for ClosedFlag to be an explicit nil
func (o *BillingStatus) SetClosedFlagNil() {
	o.ClosedFlag.Set(nil)
}

// UnsetClosedFlag ensures that no value is present for ClosedFlag, not even an explicit nil
func (o *BillingStatus) UnsetClosedFlag() {
	o.ClosedFlag.Unset()
}

// GetInactiveFlag returns the InactiveFlag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingStatus) GetInactiveFlag() bool {
	if o == nil || IsNil(o.InactiveFlag.Get()) {
		var ret bool
		return ret
	}
	return *o.InactiveFlag.Get()
}

// GetInactiveFlagOk returns a tuple with the InactiveFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingStatus) GetInactiveFlagOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.InactiveFlag.Get(), o.InactiveFlag.IsSet()
}

// HasInactiveFlag returns a boolean if a field has been set.
func (o *BillingStatus) HasInactiveFlag() bool {
	if o != nil && o.InactiveFlag.IsSet() {
		return true
	}

	return false
}

// SetInactiveFlag gets a reference to the given NullableBool and assigns it to the InactiveFlag field.
func (o *BillingStatus) SetInactiveFlag(v bool) {
	o.InactiveFlag.Set(&v)
}
// SetInactiveFlagNil sets the value for InactiveFlag to be an explicit nil
func (o *BillingStatus) SetInactiveFlagNil() {
	o.InactiveFlag.Set(nil)
}

// UnsetInactiveFlag ensures that no value is present for InactiveFlag, not even an explicit nil
func (o *BillingStatus) UnsetInactiveFlag() {
	o.InactiveFlag.Unset()
}

// GetSentFlag returns the SentFlag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingStatus) GetSentFlag() bool {
	if o == nil || IsNil(o.SentFlag.Get()) {
		var ret bool
		return ret
	}
	return *o.SentFlag.Get()
}

// GetSentFlagOk returns a tuple with the SentFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingStatus) GetSentFlagOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.SentFlag.Get(), o.SentFlag.IsSet()
}

// HasSentFlag returns a boolean if a field has been set.
func (o *BillingStatus) HasSentFlag() bool {
	if o != nil && o.SentFlag.IsSet() {
		return true
	}

	return false
}

// SetSentFlag gets a reference to the given NullableBool and assigns it to the SentFlag field.
func (o *BillingStatus) SetSentFlag(v bool) {
	o.SentFlag.Set(&v)
}
// SetSentFlagNil sets the value for SentFlag to be an explicit nil
func (o *BillingStatus) SetSentFlagNil() {
	o.SentFlag.Set(nil)
}

// UnsetSentFlag ensures that no value is present for SentFlag, not even an explicit nil
func (o *BillingStatus) UnsetSentFlag() {
	o.SentFlag.Unset()
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *BillingStatus) GetInfo() map[string]string {
	if o == nil || IsNil(o.Info) {
		var ret map[string]string
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingStatus) GetInfoOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Info) {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *BillingStatus) HasInfo() bool {
	if o != nil && !IsNil(o.Info) {
		return true
	}

	return false
}

// SetInfo gets a reference to the given map[string]string and assigns it to the Info field.
func (o *BillingStatus) SetInfo(v map[string]string) {
	o.Info = &v
}

func (o BillingStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillingStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	if o.SortOrder.IsSet() {
		toSerialize["sortOrder"] = o.SortOrder.Get()
	}
	if o.DefaultFlag.IsSet() {
		toSerialize["defaultFlag"] = o.DefaultFlag.Get()
	}
	if o.ClosedFlag.IsSet() {
		toSerialize["closedFlag"] = o.ClosedFlag.Get()
	}
	if o.InactiveFlag.IsSet() {
		toSerialize["inactiveFlag"] = o.InactiveFlag.Get()
	}
	if o.SentFlag.IsSet() {
		toSerialize["sentFlag"] = o.SentFlag.Get()
	}
	if !IsNil(o.Info) {
		toSerialize["_info"] = o.Info
	}
	return toSerialize, nil
}

func (o *BillingStatus) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varBillingStatus := _BillingStatus{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBillingStatus)

	if err != nil {
		return err
	}

	*o = BillingStatus(varBillingStatus)

	return err
}

type NullableBillingStatus struct {
	value *BillingStatus
	isSet bool
}

func (v NullableBillingStatus) Get() *BillingStatus {
	return v.value
}

func (v *NullableBillingStatus) Set(val *BillingStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingStatus(val *BillingStatus) *NullableBillingStatus {
	return &NullableBillingStatus{value: val, isSet: true}
}

func (v NullableBillingStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


