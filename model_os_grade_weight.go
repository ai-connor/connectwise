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

// checks if the OsGradeWeight type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OsGradeWeight{}

// OsGradeWeight struct for OsGradeWeight
type OsGradeWeight struct {
	Id *int32 `json:"id,omitempty"`
	OsGradeWeight *float64 `json:"osGradeWeight,omitempty"`
	OsName *string `json:"osName,omitempty"`
	Info *map[string]string `json:"_info,omitempty"`
}

// NewOsGradeWeight instantiates a new OsGradeWeight object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOsGradeWeight() *OsGradeWeight {
	this := OsGradeWeight{}
	return &this
}

// NewOsGradeWeightWithDefaults instantiates a new OsGradeWeight object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOsGradeWeightWithDefaults() *OsGradeWeight {
	this := OsGradeWeight{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OsGradeWeight) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsGradeWeight) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OsGradeWeight) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *OsGradeWeight) SetId(v int32) {
	o.Id = &v
}

// GetOsGradeWeight returns the OsGradeWeight field value if set, zero value otherwise.
func (o *OsGradeWeight) GetOsGradeWeight() float64 {
	if o == nil || IsNil(o.OsGradeWeight) {
		var ret float64
		return ret
	}
	return *o.OsGradeWeight
}

// GetOsGradeWeightOk returns a tuple with the OsGradeWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsGradeWeight) GetOsGradeWeightOk() (*float64, bool) {
	if o == nil || IsNil(o.OsGradeWeight) {
		return nil, false
	}
	return o.OsGradeWeight, true
}

// HasOsGradeWeight returns a boolean if a field has been set.
func (o *OsGradeWeight) HasOsGradeWeight() bool {
	if o != nil && !IsNil(o.OsGradeWeight) {
		return true
	}

	return false
}

// SetOsGradeWeight gets a reference to the given float64 and assigns it to the OsGradeWeight field.
func (o *OsGradeWeight) SetOsGradeWeight(v float64) {
	o.OsGradeWeight = &v
}

// GetOsName returns the OsName field value if set, zero value otherwise.
func (o *OsGradeWeight) GetOsName() string {
	if o == nil || IsNil(o.OsName) {
		var ret string
		return ret
	}
	return *o.OsName
}

// GetOsNameOk returns a tuple with the OsName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsGradeWeight) GetOsNameOk() (*string, bool) {
	if o == nil || IsNil(o.OsName) {
		return nil, false
	}
	return o.OsName, true
}

// HasOsName returns a boolean if a field has been set.
func (o *OsGradeWeight) HasOsName() bool {
	if o != nil && !IsNil(o.OsName) {
		return true
	}

	return false
}

// SetOsName gets a reference to the given string and assigns it to the OsName field.
func (o *OsGradeWeight) SetOsName(v string) {
	o.OsName = &v
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *OsGradeWeight) GetInfo() map[string]string {
	if o == nil || IsNil(o.Info) {
		var ret map[string]string
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsGradeWeight) GetInfoOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Info) {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *OsGradeWeight) HasInfo() bool {
	if o != nil && !IsNil(o.Info) {
		return true
	}

	return false
}

// SetInfo gets a reference to the given map[string]string and assigns it to the Info field.
func (o *OsGradeWeight) SetInfo(v map[string]string) {
	o.Info = &v
}

func (o OsGradeWeight) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OsGradeWeight) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.OsGradeWeight) {
		toSerialize["osGradeWeight"] = o.OsGradeWeight
	}
	if !IsNil(o.OsName) {
		toSerialize["osName"] = o.OsName
	}
	if !IsNil(o.Info) {
		toSerialize["_info"] = o.Info
	}
	return toSerialize, nil
}

type NullableOsGradeWeight struct {
	value *OsGradeWeight
	isSet bool
}

func (v NullableOsGradeWeight) Get() *OsGradeWeight {
	return v.value
}

func (v *NullableOsGradeWeight) Set(val *OsGradeWeight) {
	v.value = val
	v.isSet = true
}

func (v NullableOsGradeWeight) IsSet() bool {
	return v.isSet
}

func (v *NullableOsGradeWeight) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOsGradeWeight(val *OsGradeWeight) *NullableOsGradeWeight {
	return &NullableOsGradeWeight{value: val, isSet: true}
}

func (v NullableOsGradeWeight) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOsGradeWeight) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


