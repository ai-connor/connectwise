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

// checks if the TaxableProductTypeLevel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaxableProductTypeLevel{}

// TaxableProductTypeLevel struct for TaxableProductTypeLevel
type TaxableProductTypeLevel struct {
	Id *int32 `json:"id,omitempty"`
	TaxCodeLevel TaxCodeLevelReference `json:"taxCodeLevel"`
	Info *map[string]string `json:"_info,omitempty"`
}

type _TaxableProductTypeLevel TaxableProductTypeLevel

// NewTaxableProductTypeLevel instantiates a new TaxableProductTypeLevel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxableProductTypeLevel(taxCodeLevel TaxCodeLevelReference) *TaxableProductTypeLevel {
	this := TaxableProductTypeLevel{}
	this.TaxCodeLevel = taxCodeLevel
	return &this
}

// NewTaxableProductTypeLevelWithDefaults instantiates a new TaxableProductTypeLevel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxableProductTypeLevelWithDefaults() *TaxableProductTypeLevel {
	this := TaxableProductTypeLevel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TaxableProductTypeLevel) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxableProductTypeLevel) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TaxableProductTypeLevel) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *TaxableProductTypeLevel) SetId(v int32) {
	o.Id = &v
}

// GetTaxCodeLevel returns the TaxCodeLevel field value
func (o *TaxableProductTypeLevel) GetTaxCodeLevel() TaxCodeLevelReference {
	if o == nil {
		var ret TaxCodeLevelReference
		return ret
	}

	return o.TaxCodeLevel
}

// GetTaxCodeLevelOk returns a tuple with the TaxCodeLevel field value
// and a boolean to check if the value has been set.
func (o *TaxableProductTypeLevel) GetTaxCodeLevelOk() (*TaxCodeLevelReference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TaxCodeLevel, true
}

// SetTaxCodeLevel sets field value
func (o *TaxableProductTypeLevel) SetTaxCodeLevel(v TaxCodeLevelReference) {
	o.TaxCodeLevel = v
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *TaxableProductTypeLevel) GetInfo() map[string]string {
	if o == nil || IsNil(o.Info) {
		var ret map[string]string
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxableProductTypeLevel) GetInfoOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Info) {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *TaxableProductTypeLevel) HasInfo() bool {
	if o != nil && !IsNil(o.Info) {
		return true
	}

	return false
}

// SetInfo gets a reference to the given map[string]string and assigns it to the Info field.
func (o *TaxableProductTypeLevel) SetInfo(v map[string]string) {
	o.Info = &v
}

func (o TaxableProductTypeLevel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaxableProductTypeLevel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["taxCodeLevel"] = o.TaxCodeLevel
	if !IsNil(o.Info) {
		toSerialize["_info"] = o.Info
	}
	return toSerialize, nil
}

func (o *TaxableProductTypeLevel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"taxCodeLevel",
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

	varTaxableProductTypeLevel := _TaxableProductTypeLevel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTaxableProductTypeLevel)

	if err != nil {
		return err
	}

	*o = TaxableProductTypeLevel(varTaxableProductTypeLevel)

	return err
}

type NullableTaxableProductTypeLevel struct {
	value *TaxableProductTypeLevel
	isSet bool
}

func (v NullableTaxableProductTypeLevel) Get() *TaxableProductTypeLevel {
	return v.value
}

func (v *NullableTaxableProductTypeLevel) Set(val *TaxableProductTypeLevel) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxableProductTypeLevel) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxableProductTypeLevel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxableProductTypeLevel(val *TaxableProductTypeLevel) *NullableTaxableProductTypeLevel {
	return &NullableTaxableProductTypeLevel{value: val, isSet: true}
}

func (v NullableTaxableProductTypeLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaxableProductTypeLevel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


