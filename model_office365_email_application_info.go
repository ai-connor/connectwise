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

// checks if the Office365EmailApplicationInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Office365EmailApplicationInfo{}

// Office365EmailApplicationInfo struct for Office365EmailApplicationInfo
type Office365EmailApplicationInfo struct {
	Id *int32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Info *map[string]string `json:"_info,omitempty"`
}

// NewOffice365EmailApplicationInfo instantiates a new Office365EmailApplicationInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOffice365EmailApplicationInfo() *Office365EmailApplicationInfo {
	this := Office365EmailApplicationInfo{}
	return &this
}

// NewOffice365EmailApplicationInfoWithDefaults instantiates a new Office365EmailApplicationInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOffice365EmailApplicationInfoWithDefaults() *Office365EmailApplicationInfo {
	this := Office365EmailApplicationInfo{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Office365EmailApplicationInfo) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Office365EmailApplicationInfo) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Office365EmailApplicationInfo) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Office365EmailApplicationInfo) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Office365EmailApplicationInfo) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Office365EmailApplicationInfo) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Office365EmailApplicationInfo) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Office365EmailApplicationInfo) SetName(v string) {
	o.Name = &v
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *Office365EmailApplicationInfo) GetInfo() map[string]string {
	if o == nil || IsNil(o.Info) {
		var ret map[string]string
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Office365EmailApplicationInfo) GetInfoOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Info) {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *Office365EmailApplicationInfo) HasInfo() bool {
	if o != nil && !IsNil(o.Info) {
		return true
	}

	return false
}

// SetInfo gets a reference to the given map[string]string and assigns it to the Info field.
func (o *Office365EmailApplicationInfo) SetInfo(v map[string]string) {
	o.Info = &v
}

func (o Office365EmailApplicationInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Office365EmailApplicationInfo) ToMap() (map[string]interface{}, error) {
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

type NullableOffice365EmailApplicationInfo struct {
	value *Office365EmailApplicationInfo
	isSet bool
}

func (v NullableOffice365EmailApplicationInfo) Get() *Office365EmailApplicationInfo {
	return v.value
}

func (v *NullableOffice365EmailApplicationInfo) Set(val *Office365EmailApplicationInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableOffice365EmailApplicationInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableOffice365EmailApplicationInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOffice365EmailApplicationInfo(val *Office365EmailApplicationInfo) *NullableOffice365EmailApplicationInfo {
	return &NullableOffice365EmailApplicationInfo{value: val, isSet: true}
}

func (v NullableOffice365EmailApplicationInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOffice365EmailApplicationInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


