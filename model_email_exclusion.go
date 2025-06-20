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

// checks if the EmailExclusion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailExclusion{}

// EmailExclusion struct for EmailExclusion
type EmailExclusion struct {
	Id *int32 `json:"id,omitempty"`
	//  Max length: 100;
	Description string `json:"description"`
	Info *map[string]string `json:"_info,omitempty"`
}

type _EmailExclusion EmailExclusion

// NewEmailExclusion instantiates a new EmailExclusion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailExclusion(description string) *EmailExclusion {
	this := EmailExclusion{}
	this.Description = description
	return &this
}

// NewEmailExclusionWithDefaults instantiates a new EmailExclusion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailExclusionWithDefaults() *EmailExclusion {
	this := EmailExclusion{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EmailExclusion) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailExclusion) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EmailExclusion) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *EmailExclusion) SetId(v int32) {
	o.Id = &v
}

// GetDescription returns the Description field value
func (o *EmailExclusion) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *EmailExclusion) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *EmailExclusion) SetDescription(v string) {
	o.Description = v
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *EmailExclusion) GetInfo() map[string]string {
	if o == nil || IsNil(o.Info) {
		var ret map[string]string
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailExclusion) GetInfoOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Info) {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *EmailExclusion) HasInfo() bool {
	if o != nil && !IsNil(o.Info) {
		return true
	}

	return false
}

// SetInfo gets a reference to the given map[string]string and assigns it to the Info field.
func (o *EmailExclusion) SetInfo(v map[string]string) {
	o.Info = &v
}

func (o EmailExclusion) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailExclusion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["description"] = o.Description
	if !IsNil(o.Info) {
		toSerialize["_info"] = o.Info
	}
	return toSerialize, nil
}

func (o *EmailExclusion) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"description",
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

	varEmailExclusion := _EmailExclusion{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEmailExclusion)

	if err != nil {
		return err
	}

	*o = EmailExclusion(varEmailExclusion)

	return err
}

type NullableEmailExclusion struct {
	value *EmailExclusion
	isSet bool
}

func (v NullableEmailExclusion) Get() *EmailExclusion {
	return v.value
}

func (v *NullableEmailExclusion) Set(val *EmailExclusion) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailExclusion) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailExclusion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailExclusion(val *EmailExclusion) *NullableEmailExclusion {
	return &NullableEmailExclusion{value: val, isSet: true}
}

func (v NullableEmailExclusion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailExclusion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


