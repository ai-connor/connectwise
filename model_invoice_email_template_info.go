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

// checks if the InvoiceEmailTemplateInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvoiceEmailTemplateInfo{}

// InvoiceEmailTemplateInfo struct for InvoiceEmailTemplateInfo
type InvoiceEmailTemplateInfo struct {
	Id *int32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Info *map[string]string `json:"_info,omitempty"`
}

// NewInvoiceEmailTemplateInfo instantiates a new InvoiceEmailTemplateInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoiceEmailTemplateInfo() *InvoiceEmailTemplateInfo {
	this := InvoiceEmailTemplateInfo{}
	return &this
}

// NewInvoiceEmailTemplateInfoWithDefaults instantiates a new InvoiceEmailTemplateInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceEmailTemplateInfoWithDefaults() *InvoiceEmailTemplateInfo {
	this := InvoiceEmailTemplateInfo{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InvoiceEmailTemplateInfo) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceEmailTemplateInfo) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InvoiceEmailTemplateInfo) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *InvoiceEmailTemplateInfo) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InvoiceEmailTemplateInfo) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceEmailTemplateInfo) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InvoiceEmailTemplateInfo) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InvoiceEmailTemplateInfo) SetName(v string) {
	o.Name = &v
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *InvoiceEmailTemplateInfo) GetInfo() map[string]string {
	if o == nil || IsNil(o.Info) {
		var ret map[string]string
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceEmailTemplateInfo) GetInfoOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Info) {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *InvoiceEmailTemplateInfo) HasInfo() bool {
	if o != nil && !IsNil(o.Info) {
		return true
	}

	return false
}

// SetInfo gets a reference to the given map[string]string and assigns it to the Info field.
func (o *InvoiceEmailTemplateInfo) SetInfo(v map[string]string) {
	o.Info = &v
}

func (o InvoiceEmailTemplateInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvoiceEmailTemplateInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Info) {
		toSerialize["_info"] = o.Info
	}
	return toSerialize, nil
}

type NullableInvoiceEmailTemplateInfo struct {
	value *InvoiceEmailTemplateInfo
	isSet bool
}

func (v NullableInvoiceEmailTemplateInfo) Get() *InvoiceEmailTemplateInfo {
	return v.value
}

func (v *NullableInvoiceEmailTemplateInfo) Set(val *InvoiceEmailTemplateInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceEmailTemplateInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceEmailTemplateInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceEmailTemplateInfo(val *InvoiceEmailTemplateInfo) *NullableInvoiceEmailTemplateInfo {
	return &NullableInvoiceEmailTemplateInfo{value: val, isSet: true}
}

func (v NullableInvoiceEmailTemplateInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceEmailTemplateInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


