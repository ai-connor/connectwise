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

// checks if the LocationWorkRole type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LocationWorkRole{}

// LocationWorkRole struct for LocationWorkRole
type LocationWorkRole struct {
	Id *int32 `json:"id,omitempty"`
	Location *SystemLocationReference `json:"location,omitempty"`
	WorkRole *WorkRoleReference `json:"workRole,omitempty"`
	WorkRoleInactiveFlag NullableBool `json:"workRoleInactiveFlag,omitempty"`
	Info *map[string]string `json:"_info,omitempty"`
}

// NewLocationWorkRole instantiates a new LocationWorkRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocationWorkRole() *LocationWorkRole {
	this := LocationWorkRole{}
	return &this
}

// NewLocationWorkRoleWithDefaults instantiates a new LocationWorkRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationWorkRoleWithDefaults() *LocationWorkRole {
	this := LocationWorkRole{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LocationWorkRole) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationWorkRole) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LocationWorkRole) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *LocationWorkRole) SetId(v int32) {
	o.Id = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *LocationWorkRole) GetLocation() SystemLocationReference {
	if o == nil || IsNil(o.Location) {
		var ret SystemLocationReference
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationWorkRole) GetLocationOk() (*SystemLocationReference, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *LocationWorkRole) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given SystemLocationReference and assigns it to the Location field.
func (o *LocationWorkRole) SetLocation(v SystemLocationReference) {
	o.Location = &v
}

// GetWorkRole returns the WorkRole field value if set, zero value otherwise.
func (o *LocationWorkRole) GetWorkRole() WorkRoleReference {
	if o == nil || IsNil(o.WorkRole) {
		var ret WorkRoleReference
		return ret
	}
	return *o.WorkRole
}

// GetWorkRoleOk returns a tuple with the WorkRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationWorkRole) GetWorkRoleOk() (*WorkRoleReference, bool) {
	if o == nil || IsNil(o.WorkRole) {
		return nil, false
	}
	return o.WorkRole, true
}

// HasWorkRole returns a boolean if a field has been set.
func (o *LocationWorkRole) HasWorkRole() bool {
	if o != nil && !IsNil(o.WorkRole) {
		return true
	}

	return false
}

// SetWorkRole gets a reference to the given WorkRoleReference and assigns it to the WorkRole field.
func (o *LocationWorkRole) SetWorkRole(v WorkRoleReference) {
	o.WorkRole = &v
}

// GetWorkRoleInactiveFlag returns the WorkRoleInactiveFlag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LocationWorkRole) GetWorkRoleInactiveFlag() bool {
	if o == nil || IsNil(o.WorkRoleInactiveFlag.Get()) {
		var ret bool
		return ret
	}
	return *o.WorkRoleInactiveFlag.Get()
}

// GetWorkRoleInactiveFlagOk returns a tuple with the WorkRoleInactiveFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LocationWorkRole) GetWorkRoleInactiveFlagOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.WorkRoleInactiveFlag.Get(), o.WorkRoleInactiveFlag.IsSet()
}

// HasWorkRoleInactiveFlag returns a boolean if a field has been set.
func (o *LocationWorkRole) HasWorkRoleInactiveFlag() bool {
	if o != nil && o.WorkRoleInactiveFlag.IsSet() {
		return true
	}

	return false
}

// SetWorkRoleInactiveFlag gets a reference to the given NullableBool and assigns it to the WorkRoleInactiveFlag field.
func (o *LocationWorkRole) SetWorkRoleInactiveFlag(v bool) {
	o.WorkRoleInactiveFlag.Set(&v)
}
// SetWorkRoleInactiveFlagNil sets the value for WorkRoleInactiveFlag to be an explicit nil
func (o *LocationWorkRole) SetWorkRoleInactiveFlagNil() {
	o.WorkRoleInactiveFlag.Set(nil)
}

// UnsetWorkRoleInactiveFlag ensures that no value is present for WorkRoleInactiveFlag, not even an explicit nil
func (o *LocationWorkRole) UnsetWorkRoleInactiveFlag() {
	o.WorkRoleInactiveFlag.Unset()
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *LocationWorkRole) GetInfo() map[string]string {
	if o == nil || IsNil(o.Info) {
		var ret map[string]string
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationWorkRole) GetInfoOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Info) {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *LocationWorkRole) HasInfo() bool {
	if o != nil && !IsNil(o.Info) {
		return true
	}

	return false
}

// SetInfo gets a reference to the given map[string]string and assigns it to the Info field.
func (o *LocationWorkRole) SetInfo(v map[string]string) {
	o.Info = &v
}

func (o LocationWorkRole) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LocationWorkRole) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !IsNil(o.WorkRole) {
		toSerialize["workRole"] = o.WorkRole
	}
	if o.WorkRoleInactiveFlag.IsSet() {
		toSerialize["workRoleInactiveFlag"] = o.WorkRoleInactiveFlag.Get()
	}
	if !IsNil(o.Info) {
		toSerialize["_info"] = o.Info
	}
	return toSerialize, nil
}

type NullableLocationWorkRole struct {
	value *LocationWorkRole
	isSet bool
}

func (v NullableLocationWorkRole) Get() *LocationWorkRole {
	return v.value
}

func (v *NullableLocationWorkRole) Set(val *LocationWorkRole) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationWorkRole) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationWorkRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationWorkRole(val *LocationWorkRole) *NullableLocationWorkRole {
	return &NullableLocationWorkRole{value: val, isSet: true}
}

func (v NullableLocationWorkRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocationWorkRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


