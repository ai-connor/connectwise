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

// checks if the RmaAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RmaAction{}

// RmaAction struct for RmaAction
type RmaAction struct {
	Id *int32 `json:"id,omitempty"`
	//  Max length: 50;
	Name string `json:"name"`
	DefaultFlag NullableBool `json:"defaultFlag,omitempty"`
	Info *map[string]string `json:"_info,omitempty"`
}

type _RmaAction RmaAction

// NewRmaAction instantiates a new RmaAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRmaAction(name string) *RmaAction {
	this := RmaAction{}
	this.Name = name
	return &this
}

// NewRmaActionWithDefaults instantiates a new RmaAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRmaActionWithDefaults() *RmaAction {
	this := RmaAction{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RmaAction) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RmaAction) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RmaAction) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *RmaAction) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *RmaAction) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RmaAction) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RmaAction) SetName(v string) {
	o.Name = v
}

// GetDefaultFlag returns the DefaultFlag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RmaAction) GetDefaultFlag() bool {
	if o == nil || IsNil(o.DefaultFlag.Get()) {
		var ret bool
		return ret
	}
	return *o.DefaultFlag.Get()
}

// GetDefaultFlagOk returns a tuple with the DefaultFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RmaAction) GetDefaultFlagOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultFlag.Get(), o.DefaultFlag.IsSet()
}

// HasDefaultFlag returns a boolean if a field has been set.
func (o *RmaAction) HasDefaultFlag() bool {
	if o != nil && o.DefaultFlag.IsSet() {
		return true
	}

	return false
}

// SetDefaultFlag gets a reference to the given NullableBool and assigns it to the DefaultFlag field.
func (o *RmaAction) SetDefaultFlag(v bool) {
	o.DefaultFlag.Set(&v)
}
// SetDefaultFlagNil sets the value for DefaultFlag to be an explicit nil
func (o *RmaAction) SetDefaultFlagNil() {
	o.DefaultFlag.Set(nil)
}

// UnsetDefaultFlag ensures that no value is present for DefaultFlag, not even an explicit nil
func (o *RmaAction) UnsetDefaultFlag() {
	o.DefaultFlag.Unset()
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *RmaAction) GetInfo() map[string]string {
	if o == nil || IsNil(o.Info) {
		var ret map[string]string
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RmaAction) GetInfoOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Info) {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *RmaAction) HasInfo() bool {
	if o != nil && !IsNil(o.Info) {
		return true
	}

	return false
}

// SetInfo gets a reference to the given map[string]string and assigns it to the Info field.
func (o *RmaAction) SetInfo(v map[string]string) {
	o.Info = &v
}

func (o RmaAction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RmaAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	if o.DefaultFlag.IsSet() {
		toSerialize["defaultFlag"] = o.DefaultFlag.Get()
	}
	if !IsNil(o.Info) {
		toSerialize["_info"] = o.Info
	}
	return toSerialize, nil
}

func (o *RmaAction) UnmarshalJSON(data []byte) (err error) {
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

	varRmaAction := _RmaAction{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRmaAction)

	if err != nil {
		return err
	}

	*o = RmaAction(varRmaAction)

	return err
}

type NullableRmaAction struct {
	value *RmaAction
	isSet bool
}

func (v NullableRmaAction) Get() *RmaAction {
	return v.value
}

func (v *NullableRmaAction) Set(val *RmaAction) {
	v.value = val
	v.isSet = true
}

func (v NullableRmaAction) IsSet() bool {
	return v.isSet
}

func (v *NullableRmaAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRmaAction(val *RmaAction) *NullableRmaAction {
	return &NullableRmaAction{value: val, isSet: true}
}

func (v NullableRmaAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRmaAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


