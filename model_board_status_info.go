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

// checks if the BoardStatusInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BoardStatusInfo{}

// BoardStatusInfo struct for BoardStatusInfo
type BoardStatusInfo struct {
	Id *int32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	SortOrder NullableInt32 `json:"sortOrder,omitempty"`
	DefaultFlag NullableBool `json:"defaultFlag,omitempty"`
	InactiveFlag NullableBool `json:"inactiveFlag,omitempty"`
	ClosedFlag NullableBool `json:"closedFlag,omitempty"`
	Info *map[string]string `json:"_info,omitempty"`
}

// NewBoardStatusInfo instantiates a new BoardStatusInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBoardStatusInfo() *BoardStatusInfo {
	this := BoardStatusInfo{}
	return &this
}

// NewBoardStatusInfoWithDefaults instantiates a new BoardStatusInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBoardStatusInfoWithDefaults() *BoardStatusInfo {
	this := BoardStatusInfo{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BoardStatusInfo) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BoardStatusInfo) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BoardStatusInfo) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *BoardStatusInfo) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BoardStatusInfo) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BoardStatusInfo) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BoardStatusInfo) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BoardStatusInfo) SetName(v string) {
	o.Name = &v
}

// GetSortOrder returns the SortOrder field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BoardStatusInfo) GetSortOrder() int32 {
	if o == nil || IsNil(o.SortOrder.Get()) {
		var ret int32
		return ret
	}
	return *o.SortOrder.Get()
}

// GetSortOrderOk returns a tuple with the SortOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BoardStatusInfo) GetSortOrderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SortOrder.Get(), o.SortOrder.IsSet()
}

// HasSortOrder returns a boolean if a field has been set.
func (o *BoardStatusInfo) HasSortOrder() bool {
	if o != nil && o.SortOrder.IsSet() {
		return true
	}

	return false
}

// SetSortOrder gets a reference to the given NullableInt32 and assigns it to the SortOrder field.
func (o *BoardStatusInfo) SetSortOrder(v int32) {
	o.SortOrder.Set(&v)
}
// SetSortOrderNil sets the value for SortOrder to be an explicit nil
func (o *BoardStatusInfo) SetSortOrderNil() {
	o.SortOrder.Set(nil)
}

// UnsetSortOrder ensures that no value is present for SortOrder, not even an explicit nil
func (o *BoardStatusInfo) UnsetSortOrder() {
	o.SortOrder.Unset()
}

// GetDefaultFlag returns the DefaultFlag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BoardStatusInfo) GetDefaultFlag() bool {
	if o == nil || IsNil(o.DefaultFlag.Get()) {
		var ret bool
		return ret
	}
	return *o.DefaultFlag.Get()
}

// GetDefaultFlagOk returns a tuple with the DefaultFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BoardStatusInfo) GetDefaultFlagOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultFlag.Get(), o.DefaultFlag.IsSet()
}

// HasDefaultFlag returns a boolean if a field has been set.
func (o *BoardStatusInfo) HasDefaultFlag() bool {
	if o != nil && o.DefaultFlag.IsSet() {
		return true
	}

	return false
}

// SetDefaultFlag gets a reference to the given NullableBool and assigns it to the DefaultFlag field.
func (o *BoardStatusInfo) SetDefaultFlag(v bool) {
	o.DefaultFlag.Set(&v)
}
// SetDefaultFlagNil sets the value for DefaultFlag to be an explicit nil
func (o *BoardStatusInfo) SetDefaultFlagNil() {
	o.DefaultFlag.Set(nil)
}

// UnsetDefaultFlag ensures that no value is present for DefaultFlag, not even an explicit nil
func (o *BoardStatusInfo) UnsetDefaultFlag() {
	o.DefaultFlag.Unset()
}

// GetInactiveFlag returns the InactiveFlag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BoardStatusInfo) GetInactiveFlag() bool {
	if o == nil || IsNil(o.InactiveFlag.Get()) {
		var ret bool
		return ret
	}
	return *o.InactiveFlag.Get()
}

// GetInactiveFlagOk returns a tuple with the InactiveFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BoardStatusInfo) GetInactiveFlagOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.InactiveFlag.Get(), o.InactiveFlag.IsSet()
}

// HasInactiveFlag returns a boolean if a field has been set.
func (o *BoardStatusInfo) HasInactiveFlag() bool {
	if o != nil && o.InactiveFlag.IsSet() {
		return true
	}

	return false
}

// SetInactiveFlag gets a reference to the given NullableBool and assigns it to the InactiveFlag field.
func (o *BoardStatusInfo) SetInactiveFlag(v bool) {
	o.InactiveFlag.Set(&v)
}
// SetInactiveFlagNil sets the value for InactiveFlag to be an explicit nil
func (o *BoardStatusInfo) SetInactiveFlagNil() {
	o.InactiveFlag.Set(nil)
}

// UnsetInactiveFlag ensures that no value is present for InactiveFlag, not even an explicit nil
func (o *BoardStatusInfo) UnsetInactiveFlag() {
	o.InactiveFlag.Unset()
}

// GetClosedFlag returns the ClosedFlag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BoardStatusInfo) GetClosedFlag() bool {
	if o == nil || IsNil(o.ClosedFlag.Get()) {
		var ret bool
		return ret
	}
	return *o.ClosedFlag.Get()
}

// GetClosedFlagOk returns a tuple with the ClosedFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BoardStatusInfo) GetClosedFlagOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClosedFlag.Get(), o.ClosedFlag.IsSet()
}

// HasClosedFlag returns a boolean if a field has been set.
func (o *BoardStatusInfo) HasClosedFlag() bool {
	if o != nil && o.ClosedFlag.IsSet() {
		return true
	}

	return false
}

// SetClosedFlag gets a reference to the given NullableBool and assigns it to the ClosedFlag field.
func (o *BoardStatusInfo) SetClosedFlag(v bool) {
	o.ClosedFlag.Set(&v)
}
// SetClosedFlagNil sets the value for ClosedFlag to be an explicit nil
func (o *BoardStatusInfo) SetClosedFlagNil() {
	o.ClosedFlag.Set(nil)
}

// UnsetClosedFlag ensures that no value is present for ClosedFlag, not even an explicit nil
func (o *BoardStatusInfo) UnsetClosedFlag() {
	o.ClosedFlag.Unset()
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *BoardStatusInfo) GetInfo() map[string]string {
	if o == nil || IsNil(o.Info) {
		var ret map[string]string
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BoardStatusInfo) GetInfoOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Info) {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *BoardStatusInfo) HasInfo() bool {
	if o != nil && !IsNil(o.Info) {
		return true
	}

	return false
}

// SetInfo gets a reference to the given map[string]string and assigns it to the Info field.
func (o *BoardStatusInfo) SetInfo(v map[string]string) {
	o.Info = &v
}

func (o BoardStatusInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BoardStatusInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.SortOrder.IsSet() {
		toSerialize["sortOrder"] = o.SortOrder.Get()
	}
	if o.DefaultFlag.IsSet() {
		toSerialize["defaultFlag"] = o.DefaultFlag.Get()
	}
	if o.InactiveFlag.IsSet() {
		toSerialize["inactiveFlag"] = o.InactiveFlag.Get()
	}
	if o.ClosedFlag.IsSet() {
		toSerialize["closedFlag"] = o.ClosedFlag.Get()
	}
	if !IsNil(o.Info) {
		toSerialize["_info"] = o.Info
	}
	return toSerialize, nil
}

type NullableBoardStatusInfo struct {
	value *BoardStatusInfo
	isSet bool
}

func (v NullableBoardStatusInfo) Get() *BoardStatusInfo {
	return v.value
}

func (v *NullableBoardStatusInfo) Set(val *BoardStatusInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBoardStatusInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBoardStatusInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBoardStatusInfo(val *BoardStatusInfo) *NullableBoardStatusInfo {
	return &NullableBoardStatusInfo{value: val, isSet: true}
}

func (v NullableBoardStatusInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBoardStatusInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


