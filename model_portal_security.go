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

// checks if the PortalSecurity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PortalSecurity{}

// PortalSecurity struct for PortalSecurity
type PortalSecurity struct {
	Identifier *string `json:"identifier,omitempty"`
	Enabled NullableBool `json:"enabled,omitempty"`
}

// NewPortalSecurity instantiates a new PortalSecurity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortalSecurity() *PortalSecurity {
	this := PortalSecurity{}
	return &this
}

// NewPortalSecurityWithDefaults instantiates a new PortalSecurity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortalSecurityWithDefaults() *PortalSecurity {
	this := PortalSecurity{}
	return &this
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *PortalSecurity) GetIdentifier() string {
	if o == nil || IsNil(o.Identifier) {
		var ret string
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortalSecurity) GetIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.Identifier) {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *PortalSecurity) HasIdentifier() bool {
	if o != nil && !IsNil(o.Identifier) {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given string and assigns it to the Identifier field.
func (o *PortalSecurity) SetIdentifier(v string) {
	o.Identifier = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PortalSecurity) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled.Get()) {
		var ret bool
		return ret
	}
	return *o.Enabled.Get()
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PortalSecurity) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Enabled.Get(), o.Enabled.IsSet()
}

// HasEnabled returns a boolean if a field has been set.
func (o *PortalSecurity) HasEnabled() bool {
	if o != nil && o.Enabled.IsSet() {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given NullableBool and assigns it to the Enabled field.
func (o *PortalSecurity) SetEnabled(v bool) {
	o.Enabled.Set(&v)
}
// SetEnabledNil sets the value for Enabled to be an explicit nil
func (o *PortalSecurity) SetEnabledNil() {
	o.Enabled.Set(nil)
}

// UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
func (o *PortalSecurity) UnsetEnabled() {
	o.Enabled.Unset()
}

func (o PortalSecurity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PortalSecurity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Identifier) {
		toSerialize["identifier"] = o.Identifier
	}
	if o.Enabled.IsSet() {
		toSerialize["enabled"] = o.Enabled.Get()
	}
	return toSerialize, nil
}

type NullablePortalSecurity struct {
	value *PortalSecurity
	isSet bool
}

func (v NullablePortalSecurity) Get() *PortalSecurity {
	return v.value
}

func (v *NullablePortalSecurity) Set(val *PortalSecurity) {
	v.value = val
	v.isSet = true
}

func (v NullablePortalSecurity) IsSet() bool {
	return v.isSet
}

func (v *NullablePortalSecurity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortalSecurity(val *PortalSecurity) *NullablePortalSecurity {
	return &NullablePortalSecurity{value: val, isSet: true}
}

func (v NullablePortalSecurity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortalSecurity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


