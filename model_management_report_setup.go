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

// checks if the ManagementReportSetup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagementReportSetup{}

// ManagementReportSetup struct for ManagementReportSetup
type ManagementReportSetup struct {
	Id *int32 `json:"id,omitempty"`
	ScheduledReportDisabledFlag bool `json:"scheduledReportDisabledFlag"`
	Info *map[string]string `json:"_info,omitempty"`
}

type _ManagementReportSetup ManagementReportSetup

// NewManagementReportSetup instantiates a new ManagementReportSetup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagementReportSetup(scheduledReportDisabledFlag bool) *ManagementReportSetup {
	this := ManagementReportSetup{}
	this.ScheduledReportDisabledFlag = scheduledReportDisabledFlag
	return &this
}

// NewManagementReportSetupWithDefaults instantiates a new ManagementReportSetup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagementReportSetupWithDefaults() *ManagementReportSetup {
	this := ManagementReportSetup{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ManagementReportSetup) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementReportSetup) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ManagementReportSetup) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ManagementReportSetup) SetId(v int32) {
	o.Id = &v
}

// GetScheduledReportDisabledFlag returns the ScheduledReportDisabledFlag field value
func (o *ManagementReportSetup) GetScheduledReportDisabledFlag() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ScheduledReportDisabledFlag
}

// GetScheduledReportDisabledFlagOk returns a tuple with the ScheduledReportDisabledFlag field value
// and a boolean to check if the value has been set.
func (o *ManagementReportSetup) GetScheduledReportDisabledFlagOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScheduledReportDisabledFlag, true
}

// SetScheduledReportDisabledFlag sets field value
func (o *ManagementReportSetup) SetScheduledReportDisabledFlag(v bool) {
	o.ScheduledReportDisabledFlag = v
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *ManagementReportSetup) GetInfo() map[string]string {
	if o == nil || IsNil(o.Info) {
		var ret map[string]string
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementReportSetup) GetInfoOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Info) {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *ManagementReportSetup) HasInfo() bool {
	if o != nil && !IsNil(o.Info) {
		return true
	}

	return false
}

// SetInfo gets a reference to the given map[string]string and assigns it to the Info field.
func (o *ManagementReportSetup) SetInfo(v map[string]string) {
	o.Info = &v
}

func (o ManagementReportSetup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagementReportSetup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["scheduledReportDisabledFlag"] = o.ScheduledReportDisabledFlag
	if !IsNil(o.Info) {
		toSerialize["_info"] = o.Info
	}
	return toSerialize, nil
}

func (o *ManagementReportSetup) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"scheduledReportDisabledFlag",
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

	varManagementReportSetup := _ManagementReportSetup{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varManagementReportSetup)

	if err != nil {
		return err
	}

	*o = ManagementReportSetup(varManagementReportSetup)

	return err
}

type NullableManagementReportSetup struct {
	value *ManagementReportSetup
	isSet bool
}

func (v NullableManagementReportSetup) Get() *ManagementReportSetup {
	return v.value
}

func (v *NullableManagementReportSetup) Set(val *ManagementReportSetup) {
	v.value = val
	v.isSet = true
}

func (v NullableManagementReportSetup) IsSet() bool {
	return v.isSet
}

func (v *NullableManagementReportSetup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagementReportSetup(val *ManagementReportSetup) *NullableManagementReportSetup {
	return &NullableManagementReportSetup{value: val, isSet: true}
}

func (v NullableManagementReportSetup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagementReportSetup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


