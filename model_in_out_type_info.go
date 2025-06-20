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

// checks if the InOutTypeInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InOutTypeInfo{}

// InOutTypeInfo struct for InOutTypeInfo
type InOutTypeInfo struct {
	Id *int32 `json:"id,omitempty"`
	Description *string `json:"description,omitempty"`
	Info *map[string]string `json:"_info,omitempty"`
}

// NewInOutTypeInfo instantiates a new InOutTypeInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInOutTypeInfo() *InOutTypeInfo {
	this := InOutTypeInfo{}
	return &this
}

// NewInOutTypeInfoWithDefaults instantiates a new InOutTypeInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInOutTypeInfoWithDefaults() *InOutTypeInfo {
	this := InOutTypeInfo{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InOutTypeInfo) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InOutTypeInfo) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InOutTypeInfo) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *InOutTypeInfo) SetId(v int32) {
	o.Id = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InOutTypeInfo) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InOutTypeInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InOutTypeInfo) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InOutTypeInfo) SetDescription(v string) {
	o.Description = &v
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *InOutTypeInfo) GetInfo() map[string]string {
	if o == nil || IsNil(o.Info) {
		var ret map[string]string
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InOutTypeInfo) GetInfoOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Info) {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *InOutTypeInfo) HasInfo() bool {
	if o != nil && !IsNil(o.Info) {
		return true
	}

	return false
}

// SetInfo gets a reference to the given map[string]string and assigns it to the Info field.
func (o *InOutTypeInfo) SetInfo(v map[string]string) {
	o.Info = &v
}

func (o InOutTypeInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InOutTypeInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Info) {
		toSerialize["_info"] = o.Info
	}
	return toSerialize, nil
}

type NullableInOutTypeInfo struct {
	value *InOutTypeInfo
	isSet bool
}

func (v NullableInOutTypeInfo) Get() *InOutTypeInfo {
	return v.value
}

func (v *NullableInOutTypeInfo) Set(val *InOutTypeInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableInOutTypeInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableInOutTypeInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInOutTypeInfo(val *InOutTypeInfo) *NullableInOutTypeInfo {
	return &NullableInOutTypeInfo{value: val, isSet: true}
}

func (v NullableInOutTypeInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInOutTypeInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


