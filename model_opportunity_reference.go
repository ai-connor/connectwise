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

// checks if the OpportunityReference type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpportunityReference{}

// OpportunityReference struct for OpportunityReference
type OpportunityReference struct {
	Id NullableInt32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Info *map[string]string `json:"_info,omitempty"`
}

// NewOpportunityReference instantiates a new OpportunityReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpportunityReference() *OpportunityReference {
	this := OpportunityReference{}
	return &this
}

// NewOpportunityReferenceWithDefaults instantiates a new OpportunityReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpportunityReferenceWithDefaults() *OpportunityReference {
	this := OpportunityReference{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpportunityReference) GetId() int32 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int32
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpportunityReference) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *OpportunityReference) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt32 and assigns it to the Id field.
func (o *OpportunityReference) SetId(v int32) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *OpportunityReference) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *OpportunityReference) UnsetId() {
	o.Id.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OpportunityReference) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpportunityReference) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OpportunityReference) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OpportunityReference) SetName(v string) {
	o.Name = &v
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *OpportunityReference) GetInfo() map[string]string {
	if o == nil || IsNil(o.Info) {
		var ret map[string]string
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpportunityReference) GetInfoOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Info) {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *OpportunityReference) HasInfo() bool {
	if o != nil && !IsNil(o.Info) {
		return true
	}

	return false
}

// SetInfo gets a reference to the given map[string]string and assigns it to the Info field.
func (o *OpportunityReference) SetInfo(v map[string]string) {
	o.Info = &v
}

func (o OpportunityReference) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpportunityReference) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Info) {
		toSerialize["_info"] = o.Info
	}
	return toSerialize, nil
}

type NullableOpportunityReference struct {
	value *OpportunityReference
	isSet bool
}

func (v NullableOpportunityReference) Get() *OpportunityReference {
	return v.value
}

func (v *NullableOpportunityReference) Set(val *OpportunityReference) {
	v.value = val
	v.isSet = true
}

func (v NullableOpportunityReference) IsSet() bool {
	return v.isSet
}

func (v *NullableOpportunityReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpportunityReference(val *OpportunityReference) *NullableOpportunityReference {
	return &NullableOpportunityReference{value: val, isSet: true}
}

func (v NullableOpportunityReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpportunityReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


