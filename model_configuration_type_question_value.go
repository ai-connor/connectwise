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

// checks if the ConfigurationTypeQuestionValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigurationTypeQuestionValue{}

// ConfigurationTypeQuestionValue struct for ConfigurationTypeQuestionValue
type ConfigurationTypeQuestionValue struct {
	Id *int32 `json:"id,omitempty"`
	ConfigurationType *ConfigurationTypeReference `json:"configurationType,omitempty"`
	Question *ConfigurationTypeQuestionReference `json:"question,omitempty"`
	//  Max length: 1000;
	Value string `json:"value"`
	DefaultFlag NullableBool `json:"defaultFlag,omitempty"`
	InactiveFlag NullableBool `json:"inactiveFlag,omitempty"`
	Info *map[string]string `json:"_info,omitempty"`
}

type _ConfigurationTypeQuestionValue ConfigurationTypeQuestionValue

// NewConfigurationTypeQuestionValue instantiates a new ConfigurationTypeQuestionValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigurationTypeQuestionValue(value string) *ConfigurationTypeQuestionValue {
	this := ConfigurationTypeQuestionValue{}
	this.Value = value
	return &this
}

// NewConfigurationTypeQuestionValueWithDefaults instantiates a new ConfigurationTypeQuestionValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigurationTypeQuestionValueWithDefaults() *ConfigurationTypeQuestionValue {
	this := ConfigurationTypeQuestionValue{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ConfigurationTypeQuestionValue) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationTypeQuestionValue) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ConfigurationTypeQuestionValue) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ConfigurationTypeQuestionValue) SetId(v int32) {
	o.Id = &v
}

// GetConfigurationType returns the ConfigurationType field value if set, zero value otherwise.
func (o *ConfigurationTypeQuestionValue) GetConfigurationType() ConfigurationTypeReference {
	if o == nil || IsNil(o.ConfigurationType) {
		var ret ConfigurationTypeReference
		return ret
	}
	return *o.ConfigurationType
}

// GetConfigurationTypeOk returns a tuple with the ConfigurationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationTypeQuestionValue) GetConfigurationTypeOk() (*ConfigurationTypeReference, bool) {
	if o == nil || IsNil(o.ConfigurationType) {
		return nil, false
	}
	return o.ConfigurationType, true
}

// HasConfigurationType returns a boolean if a field has been set.
func (o *ConfigurationTypeQuestionValue) HasConfigurationType() bool {
	if o != nil && !IsNil(o.ConfigurationType) {
		return true
	}

	return false
}

// SetConfigurationType gets a reference to the given ConfigurationTypeReference and assigns it to the ConfigurationType field.
func (o *ConfigurationTypeQuestionValue) SetConfigurationType(v ConfigurationTypeReference) {
	o.ConfigurationType = &v
}

// GetQuestion returns the Question field value if set, zero value otherwise.
func (o *ConfigurationTypeQuestionValue) GetQuestion() ConfigurationTypeQuestionReference {
	if o == nil || IsNil(o.Question) {
		var ret ConfigurationTypeQuestionReference
		return ret
	}
	return *o.Question
}

// GetQuestionOk returns a tuple with the Question field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationTypeQuestionValue) GetQuestionOk() (*ConfigurationTypeQuestionReference, bool) {
	if o == nil || IsNil(o.Question) {
		return nil, false
	}
	return o.Question, true
}

// HasQuestion returns a boolean if a field has been set.
func (o *ConfigurationTypeQuestionValue) HasQuestion() bool {
	if o != nil && !IsNil(o.Question) {
		return true
	}

	return false
}

// SetQuestion gets a reference to the given ConfigurationTypeQuestionReference and assigns it to the Question field.
func (o *ConfigurationTypeQuestionValue) SetQuestion(v ConfigurationTypeQuestionReference) {
	o.Question = &v
}

// GetValue returns the Value field value
func (o *ConfigurationTypeQuestionValue) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ConfigurationTypeQuestionValue) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ConfigurationTypeQuestionValue) SetValue(v string) {
	o.Value = v
}

// GetDefaultFlag returns the DefaultFlag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConfigurationTypeQuestionValue) GetDefaultFlag() bool {
	if o == nil || IsNil(o.DefaultFlag.Get()) {
		var ret bool
		return ret
	}
	return *o.DefaultFlag.Get()
}

// GetDefaultFlagOk returns a tuple with the DefaultFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfigurationTypeQuestionValue) GetDefaultFlagOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultFlag.Get(), o.DefaultFlag.IsSet()
}

// HasDefaultFlag returns a boolean if a field has been set.
func (o *ConfigurationTypeQuestionValue) HasDefaultFlag() bool {
	if o != nil && o.DefaultFlag.IsSet() {
		return true
	}

	return false
}

// SetDefaultFlag gets a reference to the given NullableBool and assigns it to the DefaultFlag field.
func (o *ConfigurationTypeQuestionValue) SetDefaultFlag(v bool) {
	o.DefaultFlag.Set(&v)
}
// SetDefaultFlagNil sets the value for DefaultFlag to be an explicit nil
func (o *ConfigurationTypeQuestionValue) SetDefaultFlagNil() {
	o.DefaultFlag.Set(nil)
}

// UnsetDefaultFlag ensures that no value is present for DefaultFlag, not even an explicit nil
func (o *ConfigurationTypeQuestionValue) UnsetDefaultFlag() {
	o.DefaultFlag.Unset()
}

// GetInactiveFlag returns the InactiveFlag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConfigurationTypeQuestionValue) GetInactiveFlag() bool {
	if o == nil || IsNil(o.InactiveFlag.Get()) {
		var ret bool
		return ret
	}
	return *o.InactiveFlag.Get()
}

// GetInactiveFlagOk returns a tuple with the InactiveFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfigurationTypeQuestionValue) GetInactiveFlagOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.InactiveFlag.Get(), o.InactiveFlag.IsSet()
}

// HasInactiveFlag returns a boolean if a field has been set.
func (o *ConfigurationTypeQuestionValue) HasInactiveFlag() bool {
	if o != nil && o.InactiveFlag.IsSet() {
		return true
	}

	return false
}

// SetInactiveFlag gets a reference to the given NullableBool and assigns it to the InactiveFlag field.
func (o *ConfigurationTypeQuestionValue) SetInactiveFlag(v bool) {
	o.InactiveFlag.Set(&v)
}
// SetInactiveFlagNil sets the value for InactiveFlag to be an explicit nil
func (o *ConfigurationTypeQuestionValue) SetInactiveFlagNil() {
	o.InactiveFlag.Set(nil)
}

// UnsetInactiveFlag ensures that no value is present for InactiveFlag, not even an explicit nil
func (o *ConfigurationTypeQuestionValue) UnsetInactiveFlag() {
	o.InactiveFlag.Unset()
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *ConfigurationTypeQuestionValue) GetInfo() map[string]string {
	if o == nil || IsNil(o.Info) {
		var ret map[string]string
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationTypeQuestionValue) GetInfoOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Info) {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *ConfigurationTypeQuestionValue) HasInfo() bool {
	if o != nil && !IsNil(o.Info) {
		return true
	}

	return false
}

// SetInfo gets a reference to the given map[string]string and assigns it to the Info field.
func (o *ConfigurationTypeQuestionValue) SetInfo(v map[string]string) {
	o.Info = &v
}

func (o ConfigurationTypeQuestionValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigurationTypeQuestionValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ConfigurationType) {
		toSerialize["configurationType"] = o.ConfigurationType
	}
	if !IsNil(o.Question) {
		toSerialize["question"] = o.Question
	}
	toSerialize["value"] = o.Value
	if o.DefaultFlag.IsSet() {
		toSerialize["defaultFlag"] = o.DefaultFlag.Get()
	}
	if o.InactiveFlag.IsSet() {
		toSerialize["inactiveFlag"] = o.InactiveFlag.Get()
	}
	if !IsNil(o.Info) {
		toSerialize["_info"] = o.Info
	}
	return toSerialize, nil
}

func (o *ConfigurationTypeQuestionValue) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"value",
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

	varConfigurationTypeQuestionValue := _ConfigurationTypeQuestionValue{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varConfigurationTypeQuestionValue)

	if err != nil {
		return err
	}

	*o = ConfigurationTypeQuestionValue(varConfigurationTypeQuestionValue)

	return err
}

type NullableConfigurationTypeQuestionValue struct {
	value *ConfigurationTypeQuestionValue
	isSet bool
}

func (v NullableConfigurationTypeQuestionValue) Get() *ConfigurationTypeQuestionValue {
	return v.value
}

func (v *NullableConfigurationTypeQuestionValue) Set(val *ConfigurationTypeQuestionValue) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigurationTypeQuestionValue) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigurationTypeQuestionValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigurationTypeQuestionValue(val *ConfigurationTypeQuestionValue) *NullableConfigurationTypeQuestionValue {
	return &NullableConfigurationTypeQuestionValue{value: val, isSet: true}
}

func (v NullableConfigurationTypeQuestionValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigurationTypeQuestionValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


