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

// checks if the CorporateStructureLevel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CorporateStructureLevel{}

// CorporateStructureLevel struct for CorporateStructureLevel
type CorporateStructureLevel struct {
	Id *int32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewCorporateStructureLevel instantiates a new CorporateStructureLevel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCorporateStructureLevel() *CorporateStructureLevel {
	this := CorporateStructureLevel{}
	return &this
}

// NewCorporateStructureLevelWithDefaults instantiates a new CorporateStructureLevel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCorporateStructureLevelWithDefaults() *CorporateStructureLevel {
	this := CorporateStructureLevel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CorporateStructureLevel) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CorporateStructureLevel) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CorporateStructureLevel) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *CorporateStructureLevel) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CorporateStructureLevel) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CorporateStructureLevel) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CorporateStructureLevel) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CorporateStructureLevel) SetName(v string) {
	o.Name = &v
}

func (o CorporateStructureLevel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CorporateStructureLevel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableCorporateStructureLevel struct {
	value *CorporateStructureLevel
	isSet bool
}

func (v NullableCorporateStructureLevel) Get() *CorporateStructureLevel {
	return v.value
}

func (v *NullableCorporateStructureLevel) Set(val *CorporateStructureLevel) {
	v.value = val
	v.isSet = true
}

func (v NullableCorporateStructureLevel) IsSet() bool {
	return v.isSet
}

func (v *NullableCorporateStructureLevel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCorporateStructureLevel(val *CorporateStructureLevel) *NullableCorporateStructureLevel {
	return &NullableCorporateStructureLevel{value: val, isSet: true}
}

func (v NullableCorporateStructureLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCorporateStructureLevel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


