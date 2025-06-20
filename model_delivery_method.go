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

// checks if the DeliveryMethod type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeliveryMethod{}

// DeliveryMethod struct for DeliveryMethod
type DeliveryMethod struct {
	Id *int32 `json:"id,omitempty"`
	//  Max length: 50;
	Name string `json:"name"`
	DefaultFlag NullableBool `json:"defaultFlag,omitempty"`
	EmailFlag NullableBool `json:"emailFlag,omitempty"`
	IntegrationEmailFlag NullableBool `json:"integrationEmailFlag,omitempty"`
	IntegrationPrintFlag NullableBool `json:"integrationPrintFlag,omitempty"`
	IntegrationActiveFlag NullableBool `json:"integrationActiveFlag,omitempty"`
	Info *map[string]string `json:"_info,omitempty"`
}

type _DeliveryMethod DeliveryMethod

// NewDeliveryMethod instantiates a new DeliveryMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeliveryMethod(name string) *DeliveryMethod {
	this := DeliveryMethod{}
	this.Name = name
	return &this
}

// NewDeliveryMethodWithDefaults instantiates a new DeliveryMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeliveryMethodWithDefaults() *DeliveryMethod {
	this := DeliveryMethod{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DeliveryMethod) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryMethod) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DeliveryMethod) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *DeliveryMethod) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *DeliveryMethod) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DeliveryMethod) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DeliveryMethod) SetName(v string) {
	o.Name = v
}

// GetDefaultFlag returns the DefaultFlag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeliveryMethod) GetDefaultFlag() bool {
	if o == nil || IsNil(o.DefaultFlag.Get()) {
		var ret bool
		return ret
	}
	return *o.DefaultFlag.Get()
}

// GetDefaultFlagOk returns a tuple with the DefaultFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeliveryMethod) GetDefaultFlagOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultFlag.Get(), o.DefaultFlag.IsSet()
}

// HasDefaultFlag returns a boolean if a field has been set.
func (o *DeliveryMethod) HasDefaultFlag() bool {
	if o != nil && o.DefaultFlag.IsSet() {
		return true
	}

	return false
}

// SetDefaultFlag gets a reference to the given NullableBool and assigns it to the DefaultFlag field.
func (o *DeliveryMethod) SetDefaultFlag(v bool) {
	o.DefaultFlag.Set(&v)
}
// SetDefaultFlagNil sets the value for DefaultFlag to be an explicit nil
func (o *DeliveryMethod) SetDefaultFlagNil() {
	o.DefaultFlag.Set(nil)
}

// UnsetDefaultFlag ensures that no value is present for DefaultFlag, not even an explicit nil
func (o *DeliveryMethod) UnsetDefaultFlag() {
	o.DefaultFlag.Unset()
}

// GetEmailFlag returns the EmailFlag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeliveryMethod) GetEmailFlag() bool {
	if o == nil || IsNil(o.EmailFlag.Get()) {
		var ret bool
		return ret
	}
	return *o.EmailFlag.Get()
}

// GetEmailFlagOk returns a tuple with the EmailFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeliveryMethod) GetEmailFlagOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.EmailFlag.Get(), o.EmailFlag.IsSet()
}

// HasEmailFlag returns a boolean if a field has been set.
func (o *DeliveryMethod) HasEmailFlag() bool {
	if o != nil && o.EmailFlag.IsSet() {
		return true
	}

	return false
}

// SetEmailFlag gets a reference to the given NullableBool and assigns it to the EmailFlag field.
func (o *DeliveryMethod) SetEmailFlag(v bool) {
	o.EmailFlag.Set(&v)
}
// SetEmailFlagNil sets the value for EmailFlag to be an explicit nil
func (o *DeliveryMethod) SetEmailFlagNil() {
	o.EmailFlag.Set(nil)
}

// UnsetEmailFlag ensures that no value is present for EmailFlag, not even an explicit nil
func (o *DeliveryMethod) UnsetEmailFlag() {
	o.EmailFlag.Unset()
}

// GetIntegrationEmailFlag returns the IntegrationEmailFlag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeliveryMethod) GetIntegrationEmailFlag() bool {
	if o == nil || IsNil(o.IntegrationEmailFlag.Get()) {
		var ret bool
		return ret
	}
	return *o.IntegrationEmailFlag.Get()
}

// GetIntegrationEmailFlagOk returns a tuple with the IntegrationEmailFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeliveryMethod) GetIntegrationEmailFlagOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IntegrationEmailFlag.Get(), o.IntegrationEmailFlag.IsSet()
}

// HasIntegrationEmailFlag returns a boolean if a field has been set.
func (o *DeliveryMethod) HasIntegrationEmailFlag() bool {
	if o != nil && o.IntegrationEmailFlag.IsSet() {
		return true
	}

	return false
}

// SetIntegrationEmailFlag gets a reference to the given NullableBool and assigns it to the IntegrationEmailFlag field.
func (o *DeliveryMethod) SetIntegrationEmailFlag(v bool) {
	o.IntegrationEmailFlag.Set(&v)
}
// SetIntegrationEmailFlagNil sets the value for IntegrationEmailFlag to be an explicit nil
func (o *DeliveryMethod) SetIntegrationEmailFlagNil() {
	o.IntegrationEmailFlag.Set(nil)
}

// UnsetIntegrationEmailFlag ensures that no value is present for IntegrationEmailFlag, not even an explicit nil
func (o *DeliveryMethod) UnsetIntegrationEmailFlag() {
	o.IntegrationEmailFlag.Unset()
}

// GetIntegrationPrintFlag returns the IntegrationPrintFlag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeliveryMethod) GetIntegrationPrintFlag() bool {
	if o == nil || IsNil(o.IntegrationPrintFlag.Get()) {
		var ret bool
		return ret
	}
	return *o.IntegrationPrintFlag.Get()
}

// GetIntegrationPrintFlagOk returns a tuple with the IntegrationPrintFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeliveryMethod) GetIntegrationPrintFlagOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IntegrationPrintFlag.Get(), o.IntegrationPrintFlag.IsSet()
}

// HasIntegrationPrintFlag returns a boolean if a field has been set.
func (o *DeliveryMethod) HasIntegrationPrintFlag() bool {
	if o != nil && o.IntegrationPrintFlag.IsSet() {
		return true
	}

	return false
}

// SetIntegrationPrintFlag gets a reference to the given NullableBool and assigns it to the IntegrationPrintFlag field.
func (o *DeliveryMethod) SetIntegrationPrintFlag(v bool) {
	o.IntegrationPrintFlag.Set(&v)
}
// SetIntegrationPrintFlagNil sets the value for IntegrationPrintFlag to be an explicit nil
func (o *DeliveryMethod) SetIntegrationPrintFlagNil() {
	o.IntegrationPrintFlag.Set(nil)
}

// UnsetIntegrationPrintFlag ensures that no value is present for IntegrationPrintFlag, not even an explicit nil
func (o *DeliveryMethod) UnsetIntegrationPrintFlag() {
	o.IntegrationPrintFlag.Unset()
}

// GetIntegrationActiveFlag returns the IntegrationActiveFlag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeliveryMethod) GetIntegrationActiveFlag() bool {
	if o == nil || IsNil(o.IntegrationActiveFlag.Get()) {
		var ret bool
		return ret
	}
	return *o.IntegrationActiveFlag.Get()
}

// GetIntegrationActiveFlagOk returns a tuple with the IntegrationActiveFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeliveryMethod) GetIntegrationActiveFlagOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IntegrationActiveFlag.Get(), o.IntegrationActiveFlag.IsSet()
}

// HasIntegrationActiveFlag returns a boolean if a field has been set.
func (o *DeliveryMethod) HasIntegrationActiveFlag() bool {
	if o != nil && o.IntegrationActiveFlag.IsSet() {
		return true
	}

	return false
}

// SetIntegrationActiveFlag gets a reference to the given NullableBool and assigns it to the IntegrationActiveFlag field.
func (o *DeliveryMethod) SetIntegrationActiveFlag(v bool) {
	o.IntegrationActiveFlag.Set(&v)
}
// SetIntegrationActiveFlagNil sets the value for IntegrationActiveFlag to be an explicit nil
func (o *DeliveryMethod) SetIntegrationActiveFlagNil() {
	o.IntegrationActiveFlag.Set(nil)
}

// UnsetIntegrationActiveFlag ensures that no value is present for IntegrationActiveFlag, not even an explicit nil
func (o *DeliveryMethod) UnsetIntegrationActiveFlag() {
	o.IntegrationActiveFlag.Unset()
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *DeliveryMethod) GetInfo() map[string]string {
	if o == nil || IsNil(o.Info) {
		var ret map[string]string
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryMethod) GetInfoOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Info) {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *DeliveryMethod) HasInfo() bool {
	if o != nil && !IsNil(o.Info) {
		return true
	}

	return false
}

// SetInfo gets a reference to the given map[string]string and assigns it to the Info field.
func (o *DeliveryMethod) SetInfo(v map[string]string) {
	o.Info = &v
}

func (o DeliveryMethod) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeliveryMethod) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	if o.DefaultFlag.IsSet() {
		toSerialize["defaultFlag"] = o.DefaultFlag.Get()
	}
	if o.EmailFlag.IsSet() {
		toSerialize["emailFlag"] = o.EmailFlag.Get()
	}
	if o.IntegrationEmailFlag.IsSet() {
		toSerialize["integrationEmailFlag"] = o.IntegrationEmailFlag.Get()
	}
	if o.IntegrationPrintFlag.IsSet() {
		toSerialize["integrationPrintFlag"] = o.IntegrationPrintFlag.Get()
	}
	if o.IntegrationActiveFlag.IsSet() {
		toSerialize["integrationActiveFlag"] = o.IntegrationActiveFlag.Get()
	}
	if !IsNil(o.Info) {
		toSerialize["_info"] = o.Info
	}
	return toSerialize, nil
}

func (o *DeliveryMethod) UnmarshalJSON(data []byte) (err error) {
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

	varDeliveryMethod := _DeliveryMethod{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDeliveryMethod)

	if err != nil {
		return err
	}

	*o = DeliveryMethod(varDeliveryMethod)

	return err
}

type NullableDeliveryMethod struct {
	value *DeliveryMethod
	isSet bool
}

func (v NullableDeliveryMethod) Get() *DeliveryMethod {
	return v.value
}

func (v *NullableDeliveryMethod) Set(val *DeliveryMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliveryMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliveryMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliveryMethod(val *DeliveryMethod) *NullableDeliveryMethod {
	return &NullableDeliveryMethod{value: val, isSet: true}
}

func (v NullableDeliveryMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliveryMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


