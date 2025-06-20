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

// checks if the ConfigurationType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigurationType{}

// ConfigurationType struct for ConfigurationType
type ConfigurationType struct {
	Id *int32 `json:"id,omitempty"`
	//  Max length: 50;
	Name string `json:"name"`
	InactiveFlag NullableBool `json:"inactiveFlag,omitempty"`
	SystemFlag NullableBool `json:"systemFlag,omitempty"`
	Info *map[string]string `json:"_info,omitempty"`
}

type _ConfigurationType ConfigurationType

// NewConfigurationType instantiates a new ConfigurationType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigurationType(name string) *ConfigurationType {
	this := ConfigurationType{}
	this.Name = name
	return &this
}

// NewConfigurationTypeWithDefaults instantiates a new ConfigurationType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigurationTypeWithDefaults() *ConfigurationType {
	this := ConfigurationType{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ConfigurationType) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationType) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ConfigurationType) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ConfigurationType) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *ConfigurationType) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ConfigurationType) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ConfigurationType) SetName(v string) {
	o.Name = v
}

// GetInactiveFlag returns the InactiveFlag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConfigurationType) GetInactiveFlag() bool {
	if o == nil || IsNil(o.InactiveFlag.Get()) {
		var ret bool
		return ret
	}
	return *o.InactiveFlag.Get()
}

// GetInactiveFlagOk returns a tuple with the InactiveFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfigurationType) GetInactiveFlagOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.InactiveFlag.Get(), o.InactiveFlag.IsSet()
}

// HasInactiveFlag returns a boolean if a field has been set.
func (o *ConfigurationType) HasInactiveFlag() bool {
	if o != nil && o.InactiveFlag.IsSet() {
		return true
	}

	return false
}

// SetInactiveFlag gets a reference to the given NullableBool and assigns it to the InactiveFlag field.
func (o *ConfigurationType) SetInactiveFlag(v bool) {
	o.InactiveFlag.Set(&v)
}
// SetInactiveFlagNil sets the value for InactiveFlag to be an explicit nil
func (o *ConfigurationType) SetInactiveFlagNil() {
	o.InactiveFlag.Set(nil)
}

// UnsetInactiveFlag ensures that no value is present for InactiveFlag, not even an explicit nil
func (o *ConfigurationType) UnsetInactiveFlag() {
	o.InactiveFlag.Unset()
}

// GetSystemFlag returns the SystemFlag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConfigurationType) GetSystemFlag() bool {
	if o == nil || IsNil(o.SystemFlag.Get()) {
		var ret bool
		return ret
	}
	return *o.SystemFlag.Get()
}

// GetSystemFlagOk returns a tuple with the SystemFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfigurationType) GetSystemFlagOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.SystemFlag.Get(), o.SystemFlag.IsSet()
}

// HasSystemFlag returns a boolean if a field has been set.
func (o *ConfigurationType) HasSystemFlag() bool {
	if o != nil && o.SystemFlag.IsSet() {
		return true
	}

	return false
}

// SetSystemFlag gets a reference to the given NullableBool and assigns it to the SystemFlag field.
func (o *ConfigurationType) SetSystemFlag(v bool) {
	o.SystemFlag.Set(&v)
}
// SetSystemFlagNil sets the value for SystemFlag to be an explicit nil
func (o *ConfigurationType) SetSystemFlagNil() {
	o.SystemFlag.Set(nil)
}

// UnsetSystemFlag ensures that no value is present for SystemFlag, not even an explicit nil
func (o *ConfigurationType) UnsetSystemFlag() {
	o.SystemFlag.Unset()
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *ConfigurationType) GetInfo() map[string]string {
	if o == nil || IsNil(o.Info) {
		var ret map[string]string
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationType) GetInfoOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Info) {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *ConfigurationType) HasInfo() bool {
	if o != nil && !IsNil(o.Info) {
		return true
	}

	return false
}

// SetInfo gets a reference to the given map[string]string and assigns it to the Info field.
func (o *ConfigurationType) SetInfo(v map[string]string) {
	o.Info = &v
}

func (o ConfigurationType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigurationType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	if o.InactiveFlag.IsSet() {
		toSerialize["inactiveFlag"] = o.InactiveFlag.Get()
	}
	if o.SystemFlag.IsSet() {
		toSerialize["systemFlag"] = o.SystemFlag.Get()
	}
	if !IsNil(o.Info) {
		toSerialize["_info"] = o.Info
	}
	return toSerialize, nil
}

func (o *ConfigurationType) UnmarshalJSON(data []byte) (err error) {
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

	varConfigurationType := _ConfigurationType{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varConfigurationType)

	if err != nil {
		return err
	}

	*o = ConfigurationType(varConfigurationType)

	return err
}

type NullableConfigurationType struct {
	value *ConfigurationType
	isSet bool
}

func (v NullableConfigurationType) Get() *ConfigurationType {
	return v.value
}

func (v *NullableConfigurationType) Set(val *ConfigurationType) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigurationType) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigurationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigurationType(val *ConfigurationType) *NullableConfigurationType {
	return &NullableConfigurationType{value: val, isSet: true}
}

func (v NullableConfigurationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigurationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


