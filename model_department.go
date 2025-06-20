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

// checks if the Department type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Department{}

// Department struct for Department
type Department struct {
	Id *int32 `json:"id,omitempty"`
	//  Max length: 15;
	Identifier string `json:"identifier"`
	//  Max length: 50;
	Name string `json:"name"`
	Info *map[string]string `json:"_info,omitempty"`
}

type _Department Department

// NewDepartment instantiates a new Department object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDepartment(identifier string, name string) *Department {
	this := Department{}
	this.Identifier = identifier
	this.Name = name
	return &this
}

// NewDepartmentWithDefaults instantiates a new Department object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDepartmentWithDefaults() *Department {
	this := Department{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Department) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Department) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Department) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Department) SetId(v int32) {
	o.Id = &v
}

// GetIdentifier returns the Identifier field value
func (o *Department) GetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
func (o *Department) GetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identifier, true
}

// SetIdentifier sets field value
func (o *Department) SetIdentifier(v string) {
	o.Identifier = v
}

// GetName returns the Name field value
func (o *Department) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Department) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Department) SetName(v string) {
	o.Name = v
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *Department) GetInfo() map[string]string {
	if o == nil || IsNil(o.Info) {
		var ret map[string]string
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Department) GetInfoOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Info) {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *Department) HasInfo() bool {
	if o != nil && !IsNil(o.Info) {
		return true
	}

	return false
}

// SetInfo gets a reference to the given map[string]string and assigns it to the Info field.
func (o *Department) SetInfo(v map[string]string) {
	o.Info = &v
}

func (o Department) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Department) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["identifier"] = o.Identifier
	toSerialize["name"] = o.Name
	if !IsNil(o.Info) {
		toSerialize["_info"] = o.Info
	}
	return toSerialize, nil
}

func (o *Department) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"identifier",
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

	varDepartment := _Department{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDepartment)

	if err != nil {
		return err
	}

	*o = Department(varDepartment)

	return err
}

type NullableDepartment struct {
	value *Department
	isSet bool
}

func (v NullableDepartment) Get() *Department {
	return v.value
}

func (v *NullableDepartment) Set(val *Department) {
	v.value = val
	v.isSet = true
}

func (v NullableDepartment) IsSet() bool {
	return v.isSet
}

func (v *NullableDepartment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDepartment(val *Department) *NullableDepartment {
	return &NullableDepartment{value: val, isSet: true}
}

func (v NullableDepartment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDepartment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


