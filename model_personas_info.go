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

// checks if the PersonasInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PersonasInfo{}

// PersonasInfo struct for PersonasInfo
type PersonasInfo struct {
	Id *int32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewPersonasInfo instantiates a new PersonasInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPersonasInfo() *PersonasInfo {
	this := PersonasInfo{}
	return &this
}

// NewPersonasInfoWithDefaults instantiates a new PersonasInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPersonasInfoWithDefaults() *PersonasInfo {
	this := PersonasInfo{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PersonasInfo) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonasInfo) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PersonasInfo) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *PersonasInfo) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PersonasInfo) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonasInfo) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PersonasInfo) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PersonasInfo) SetName(v string) {
	o.Name = &v
}

func (o PersonasInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PersonasInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullablePersonasInfo struct {
	value *PersonasInfo
	isSet bool
}

func (v NullablePersonasInfo) Get() *PersonasInfo {
	return v.value
}

func (v *NullablePersonasInfo) Set(val *PersonasInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePersonasInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePersonasInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePersonasInfo(val *PersonasInfo) *NullablePersonasInfo {
	return &NullablePersonasInfo{value: val, isSet: true}
}

func (v NullablePersonasInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePersonasInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


