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

// checks if the MemberAccrual type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MemberAccrual{}

// MemberAccrual struct for MemberAccrual
type MemberAccrual struct {
	Id *int32 `json:"id,omitempty"`
	AccrualType NullableString `json:"accrualType"`
	Year NullableInt32 `json:"year"`
	Hours NullableFloat64 `json:"hours"`
	Reason string `json:"reason"`
	Member *MemberReference `json:"member,omitempty"`
	Info *map[string]string `json:"_info,omitempty"`
}

type _MemberAccrual MemberAccrual

// NewMemberAccrual instantiates a new MemberAccrual object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMemberAccrual(accrualType NullableString, year NullableInt32, hours NullableFloat64, reason string) *MemberAccrual {
	this := MemberAccrual{}
	this.AccrualType = accrualType
	this.Year = year
	this.Hours = hours
	this.Reason = reason
	return &this
}

// NewMemberAccrualWithDefaults instantiates a new MemberAccrual object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMemberAccrualWithDefaults() *MemberAccrual {
	this := MemberAccrual{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MemberAccrual) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberAccrual) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MemberAccrual) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *MemberAccrual) SetId(v int32) {
	o.Id = &v
}

// GetAccrualType returns the AccrualType field value
// If the value is explicit nil, the zero value for string will be returned
func (o *MemberAccrual) GetAccrualType() string {
	if o == nil || o.AccrualType.Get() == nil {
		var ret string
		return ret
	}

	return *o.AccrualType.Get()
}

// GetAccrualTypeOk returns a tuple with the AccrualType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MemberAccrual) GetAccrualTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccrualType.Get(), o.AccrualType.IsSet()
}

// SetAccrualType sets field value
func (o *MemberAccrual) SetAccrualType(v string) {
	o.AccrualType.Set(&v)
}

// GetYear returns the Year field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *MemberAccrual) GetYear() int32 {
	if o == nil || o.Year.Get() == nil {
		var ret int32
		return ret
	}

	return *o.Year.Get()
}

// GetYearOk returns a tuple with the Year field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MemberAccrual) GetYearOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Year.Get(), o.Year.IsSet()
}

// SetYear sets field value
func (o *MemberAccrual) SetYear(v int32) {
	o.Year.Set(&v)
}

// GetHours returns the Hours field value
// If the value is explicit nil, the zero value for float64 will be returned
func (o *MemberAccrual) GetHours() float64 {
	if o == nil || o.Hours.Get() == nil {
		var ret float64
		return ret
	}

	return *o.Hours.Get()
}

// GetHoursOk returns a tuple with the Hours field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MemberAccrual) GetHoursOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Hours.Get(), o.Hours.IsSet()
}

// SetHours sets field value
func (o *MemberAccrual) SetHours(v float64) {
	o.Hours.Set(&v)
}

// GetReason returns the Reason field value
func (o *MemberAccrual) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *MemberAccrual) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *MemberAccrual) SetReason(v string) {
	o.Reason = v
}

// GetMember returns the Member field value if set, zero value otherwise.
func (o *MemberAccrual) GetMember() MemberReference {
	if o == nil || IsNil(o.Member) {
		var ret MemberReference
		return ret
	}
	return *o.Member
}

// GetMemberOk returns a tuple with the Member field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberAccrual) GetMemberOk() (*MemberReference, bool) {
	if o == nil || IsNil(o.Member) {
		return nil, false
	}
	return o.Member, true
}

// HasMember returns a boolean if a field has been set.
func (o *MemberAccrual) HasMember() bool {
	if o != nil && !IsNil(o.Member) {
		return true
	}

	return false
}

// SetMember gets a reference to the given MemberReference and assigns it to the Member field.
func (o *MemberAccrual) SetMember(v MemberReference) {
	o.Member = &v
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *MemberAccrual) GetInfo() map[string]string {
	if o == nil || IsNil(o.Info) {
		var ret map[string]string
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberAccrual) GetInfoOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Info) {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *MemberAccrual) HasInfo() bool {
	if o != nil && !IsNil(o.Info) {
		return true
	}

	return false
}

// SetInfo gets a reference to the given map[string]string and assigns it to the Info field.
func (o *MemberAccrual) SetInfo(v map[string]string) {
	o.Info = &v
}

func (o MemberAccrual) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MemberAccrual) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["accrualType"] = o.AccrualType.Get()
	toSerialize["year"] = o.Year.Get()
	toSerialize["hours"] = o.Hours.Get()
	toSerialize["reason"] = o.Reason
	if !IsNil(o.Member) {
		toSerialize["member"] = o.Member
	}
	if !IsNil(o.Info) {
		toSerialize["_info"] = o.Info
	}
	return toSerialize, nil
}

func (o *MemberAccrual) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"accrualType",
		"year",
		"hours",
		"reason",
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

	varMemberAccrual := _MemberAccrual{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMemberAccrual)

	if err != nil {
		return err
	}

	*o = MemberAccrual(varMemberAccrual)

	return err
}

type NullableMemberAccrual struct {
	value *MemberAccrual
	isSet bool
}

func (v NullableMemberAccrual) Get() *MemberAccrual {
	return v.value
}

func (v *NullableMemberAccrual) Set(val *MemberAccrual) {
	v.value = val
	v.isSet = true
}

func (v NullableMemberAccrual) IsSet() bool {
	return v.isSet
}

func (v *NullableMemberAccrual) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMemberAccrual(val *MemberAccrual) *NullableMemberAccrual {
	return &NullableMemberAccrual{value: val, isSet: true}
}

func (v NullableMemberAccrual) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMemberAccrual) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


