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

// checks if the MemberOffice365 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MemberOffice365{}

// MemberOffice365 struct for MemberOffice365
type MemberOffice365 struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewMemberOffice365 instantiates a new MemberOffice365 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMemberOffice365() *MemberOffice365 {
	this := MemberOffice365{}
	return &this
}

// NewMemberOffice365WithDefaults instantiates a new MemberOffice365 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMemberOffice365WithDefaults() *MemberOffice365 {
	this := MemberOffice365{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MemberOffice365) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberOffice365) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MemberOffice365) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MemberOffice365) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MemberOffice365) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberOffice365) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MemberOffice365) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MemberOffice365) SetName(v string) {
	o.Name = &v
}

func (o MemberOffice365) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MemberOffice365) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableMemberOffice365 struct {
	value *MemberOffice365
	isSet bool
}

func (v NullableMemberOffice365) Get() *MemberOffice365 {
	return v.value
}

func (v *NullableMemberOffice365) Set(val *MemberOffice365) {
	v.value = val
	v.isSet = true
}

func (v NullableMemberOffice365) IsSet() bool {
	return v.isSet
}

func (v *NullableMemberOffice365) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMemberOffice365(val *MemberOffice365) *NullableMemberOffice365 {
	return &NullableMemberOffice365{value: val, isSet: true}
}

func (v NullableMemberOffice365) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMemberOffice365) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


