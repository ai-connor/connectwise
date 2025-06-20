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

// checks if the BillingCycle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillingCycle{}

// BillingCycle struct for BillingCycle
type BillingCycle struct {
	Id *int32 `json:"id,omitempty"`
	//  Max length: 5;
	Identifier string `json:"identifier"`
	//  Max length: 50;
	Name string `json:"name"`
	DefaultFlag *bool `json:"defaultFlag,omitempty"`
	BillingOptions NullableString `json:"billingOptions"`
	Info *map[string]string `json:"_info,omitempty"`
}

type _BillingCycle BillingCycle

// NewBillingCycle instantiates a new BillingCycle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingCycle(identifier string, name string, billingOptions NullableString) *BillingCycle {
	this := BillingCycle{}
	this.Identifier = identifier
	this.Name = name
	this.BillingOptions = billingOptions
	return &this
}

// NewBillingCycleWithDefaults instantiates a new BillingCycle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingCycleWithDefaults() *BillingCycle {
	this := BillingCycle{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BillingCycle) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingCycle) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BillingCycle) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *BillingCycle) SetId(v int32) {
	o.Id = &v
}

// GetIdentifier returns the Identifier field value
func (o *BillingCycle) GetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
func (o *BillingCycle) GetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identifier, true
}

// SetIdentifier sets field value
func (o *BillingCycle) SetIdentifier(v string) {
	o.Identifier = v
}

// GetName returns the Name field value
func (o *BillingCycle) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BillingCycle) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BillingCycle) SetName(v string) {
	o.Name = v
}

// GetDefaultFlag returns the DefaultFlag field value if set, zero value otherwise.
func (o *BillingCycle) GetDefaultFlag() bool {
	if o == nil || IsNil(o.DefaultFlag) {
		var ret bool
		return ret
	}
	return *o.DefaultFlag
}

// GetDefaultFlagOk returns a tuple with the DefaultFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingCycle) GetDefaultFlagOk() (*bool, bool) {
	if o == nil || IsNil(o.DefaultFlag) {
		return nil, false
	}
	return o.DefaultFlag, true
}

// HasDefaultFlag returns a boolean if a field has been set.
func (o *BillingCycle) HasDefaultFlag() bool {
	if o != nil && !IsNil(o.DefaultFlag) {
		return true
	}

	return false
}

// SetDefaultFlag gets a reference to the given bool and assigns it to the DefaultFlag field.
func (o *BillingCycle) SetDefaultFlag(v bool) {
	o.DefaultFlag = &v
}

// GetBillingOptions returns the BillingOptions field value
// If the value is explicit nil, the zero value for string will be returned
func (o *BillingCycle) GetBillingOptions() string {
	if o == nil || o.BillingOptions.Get() == nil {
		var ret string
		return ret
	}

	return *o.BillingOptions.Get()
}

// GetBillingOptionsOk returns a tuple with the BillingOptions field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingCycle) GetBillingOptionsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BillingOptions.Get(), o.BillingOptions.IsSet()
}

// SetBillingOptions sets field value
func (o *BillingCycle) SetBillingOptions(v string) {
	o.BillingOptions.Set(&v)
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *BillingCycle) GetInfo() map[string]string {
	if o == nil || IsNil(o.Info) {
		var ret map[string]string
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingCycle) GetInfoOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Info) {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *BillingCycle) HasInfo() bool {
	if o != nil && !IsNil(o.Info) {
		return true
	}

	return false
}

// SetInfo gets a reference to the given map[string]string and assigns it to the Info field.
func (o *BillingCycle) SetInfo(v map[string]string) {
	o.Info = &v
}

func (o BillingCycle) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillingCycle) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["identifier"] = o.Identifier
	toSerialize["name"] = o.Name
	if !IsNil(o.DefaultFlag) {
		toSerialize["defaultFlag"] = o.DefaultFlag
	}
	toSerialize["billingOptions"] = o.BillingOptions.Get()
	if !IsNil(o.Info) {
		toSerialize["_info"] = o.Info
	}
	return toSerialize, nil
}

func (o *BillingCycle) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"identifier",
		"name",
		"billingOptions",
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

	varBillingCycle := _BillingCycle{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBillingCycle)

	if err != nil {
		return err
	}

	*o = BillingCycle(varBillingCycle)

	return err
}

type NullableBillingCycle struct {
	value *BillingCycle
	isSet bool
}

func (v NullableBillingCycle) Get() *BillingCycle {
	return v.value
}

func (v *NullableBillingCycle) Set(val *BillingCycle) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingCycle) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingCycle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingCycle(val *BillingCycle) *NullableBillingCycle {
	return &NullableBillingCycle{value: val, isSet: true}
}

func (v NullableBillingCycle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingCycle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


