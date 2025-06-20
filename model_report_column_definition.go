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

// checks if the ReportColumnDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReportColumnDefinition{}

// ReportColumnDefinition struct for ReportColumnDefinition
type ReportColumnDefinition struct {
	Type *string `json:"type,omitempty"`
	IsNullable *bool `json:"isNullable,omitempty"`
	IdentityColumn *bool `json:"identityColumn,omitempty"`
}

// NewReportColumnDefinition instantiates a new ReportColumnDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportColumnDefinition() *ReportColumnDefinition {
	this := ReportColumnDefinition{}
	return &this
}

// NewReportColumnDefinitionWithDefaults instantiates a new ReportColumnDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportColumnDefinitionWithDefaults() *ReportColumnDefinition {
	this := ReportColumnDefinition{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ReportColumnDefinition) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportColumnDefinition) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ReportColumnDefinition) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ReportColumnDefinition) SetType(v string) {
	o.Type = &v
}

// GetIsNullable returns the IsNullable field value if set, zero value otherwise.
func (o *ReportColumnDefinition) GetIsNullable() bool {
	if o == nil || IsNil(o.IsNullable) {
		var ret bool
		return ret
	}
	return *o.IsNullable
}

// GetIsNullableOk returns a tuple with the IsNullable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportColumnDefinition) GetIsNullableOk() (*bool, bool) {
	if o == nil || IsNil(o.IsNullable) {
		return nil, false
	}
	return o.IsNullable, true
}

// HasIsNullable returns a boolean if a field has been set.
func (o *ReportColumnDefinition) HasIsNullable() bool {
	if o != nil && !IsNil(o.IsNullable) {
		return true
	}

	return false
}

// SetIsNullable gets a reference to the given bool and assigns it to the IsNullable field.
func (o *ReportColumnDefinition) SetIsNullable(v bool) {
	o.IsNullable = &v
}

// GetIdentityColumn returns the IdentityColumn field value if set, zero value otherwise.
func (o *ReportColumnDefinition) GetIdentityColumn() bool {
	if o == nil || IsNil(o.IdentityColumn) {
		var ret bool
		return ret
	}
	return *o.IdentityColumn
}

// GetIdentityColumnOk returns a tuple with the IdentityColumn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportColumnDefinition) GetIdentityColumnOk() (*bool, bool) {
	if o == nil || IsNil(o.IdentityColumn) {
		return nil, false
	}
	return o.IdentityColumn, true
}

// HasIdentityColumn returns a boolean if a field has been set.
func (o *ReportColumnDefinition) HasIdentityColumn() bool {
	if o != nil && !IsNil(o.IdentityColumn) {
		return true
	}

	return false
}

// SetIdentityColumn gets a reference to the given bool and assigns it to the IdentityColumn field.
func (o *ReportColumnDefinition) SetIdentityColumn(v bool) {
	o.IdentityColumn = &v
}

func (o ReportColumnDefinition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReportColumnDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.IsNullable) {
		toSerialize["isNullable"] = o.IsNullable
	}
	if !IsNil(o.IdentityColumn) {
		toSerialize["identityColumn"] = o.IdentityColumn
	}
	return toSerialize, nil
}

type NullableReportColumnDefinition struct {
	value *ReportColumnDefinition
	isSet bool
}

func (v NullableReportColumnDefinition) Get() *ReportColumnDefinition {
	return v.value
}

func (v *NullableReportColumnDefinition) Set(val *ReportColumnDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableReportColumnDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableReportColumnDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportColumnDefinition(val *ReportColumnDefinition) *NullableReportColumnDefinition {
	return &NullableReportColumnDefinition{value: val, isSet: true}
}

func (v NullableReportColumnDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportColumnDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


