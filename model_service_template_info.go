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

// checks if the ServiceTemplateInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceTemplateInfo{}

// ServiceTemplateInfo struct for ServiceTemplateInfo
type ServiceTemplateInfo struct {
	Id *int32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	TemplateFlag NullableBool `json:"templateFlag,omitempty"`
	Info *map[string]string `json:"_info,omitempty"`
}

// NewServiceTemplateInfo instantiates a new ServiceTemplateInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceTemplateInfo() *ServiceTemplateInfo {
	this := ServiceTemplateInfo{}
	return &this
}

// NewServiceTemplateInfoWithDefaults instantiates a new ServiceTemplateInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceTemplateInfoWithDefaults() *ServiceTemplateInfo {
	this := ServiceTemplateInfo{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ServiceTemplateInfo) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceTemplateInfo) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ServiceTemplateInfo) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ServiceTemplateInfo) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ServiceTemplateInfo) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceTemplateInfo) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ServiceTemplateInfo) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ServiceTemplateInfo) SetName(v string) {
	o.Name = &v
}

// GetTemplateFlag returns the TemplateFlag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServiceTemplateInfo) GetTemplateFlag() bool {
	if o == nil || IsNil(o.TemplateFlag.Get()) {
		var ret bool
		return ret
	}
	return *o.TemplateFlag.Get()
}

// GetTemplateFlagOk returns a tuple with the TemplateFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceTemplateInfo) GetTemplateFlagOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.TemplateFlag.Get(), o.TemplateFlag.IsSet()
}

// HasTemplateFlag returns a boolean if a field has been set.
func (o *ServiceTemplateInfo) HasTemplateFlag() bool {
	if o != nil && o.TemplateFlag.IsSet() {
		return true
	}

	return false
}

// SetTemplateFlag gets a reference to the given NullableBool and assigns it to the TemplateFlag field.
func (o *ServiceTemplateInfo) SetTemplateFlag(v bool) {
	o.TemplateFlag.Set(&v)
}
// SetTemplateFlagNil sets the value for TemplateFlag to be an explicit nil
func (o *ServiceTemplateInfo) SetTemplateFlagNil() {
	o.TemplateFlag.Set(nil)
}

// UnsetTemplateFlag ensures that no value is present for TemplateFlag, not even an explicit nil
func (o *ServiceTemplateInfo) UnsetTemplateFlag() {
	o.TemplateFlag.Unset()
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *ServiceTemplateInfo) GetInfo() map[string]string {
	if o == nil || IsNil(o.Info) {
		var ret map[string]string
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceTemplateInfo) GetInfoOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Info) {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *ServiceTemplateInfo) HasInfo() bool {
	if o != nil && !IsNil(o.Info) {
		return true
	}

	return false
}

// SetInfo gets a reference to the given map[string]string and assigns it to the Info field.
func (o *ServiceTemplateInfo) SetInfo(v map[string]string) {
	o.Info = &v
}

func (o ServiceTemplateInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceTemplateInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.TemplateFlag.IsSet() {
		toSerialize["templateFlag"] = o.TemplateFlag.Get()
	}
	if !IsNil(o.Info) {
		toSerialize["_info"] = o.Info
	}
	return toSerialize, nil
}

type NullableServiceTemplateInfo struct {
	value *ServiceTemplateInfo
	isSet bool
}

func (v NullableServiceTemplateInfo) Get() *ServiceTemplateInfo {
	return v.value
}

func (v *NullableServiceTemplateInfo) Set(val *ServiceTemplateInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceTemplateInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceTemplateInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceTemplateInfo(val *ServiceTemplateInfo) *NullableServiceTemplateInfo {
	return &NullableServiceTemplateInfo{value: val, isSet: true}
}

func (v NullableServiceTemplateInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceTemplateInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


