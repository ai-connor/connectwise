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

// checks if the ServiceTicketLinkInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceTicketLinkInfo{}

// ServiceTicketLinkInfo struct for ServiceTicketLinkInfo
type ServiceTicketLinkInfo struct {
	Id *int32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	LinkText *string `json:"linkText,omitempty"`
	Url *string `json:"url,omitempty"`
	Info *map[string]string `json:"_info,omitempty"`
}

// NewServiceTicketLinkInfo instantiates a new ServiceTicketLinkInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceTicketLinkInfo() *ServiceTicketLinkInfo {
	this := ServiceTicketLinkInfo{}
	return &this
}

// NewServiceTicketLinkInfoWithDefaults instantiates a new ServiceTicketLinkInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceTicketLinkInfoWithDefaults() *ServiceTicketLinkInfo {
	this := ServiceTicketLinkInfo{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ServiceTicketLinkInfo) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceTicketLinkInfo) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ServiceTicketLinkInfo) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ServiceTicketLinkInfo) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ServiceTicketLinkInfo) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceTicketLinkInfo) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ServiceTicketLinkInfo) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ServiceTicketLinkInfo) SetName(v string) {
	o.Name = &v
}

// GetLinkText returns the LinkText field value if set, zero value otherwise.
func (o *ServiceTicketLinkInfo) GetLinkText() string {
	if o == nil || IsNil(o.LinkText) {
		var ret string
		return ret
	}
	return *o.LinkText
}

// GetLinkTextOk returns a tuple with the LinkText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceTicketLinkInfo) GetLinkTextOk() (*string, bool) {
	if o == nil || IsNil(o.LinkText) {
		return nil, false
	}
	return o.LinkText, true
}

// HasLinkText returns a boolean if a field has been set.
func (o *ServiceTicketLinkInfo) HasLinkText() bool {
	if o != nil && !IsNil(o.LinkText) {
		return true
	}

	return false
}

// SetLinkText gets a reference to the given string and assigns it to the LinkText field.
func (o *ServiceTicketLinkInfo) SetLinkText(v string) {
	o.LinkText = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *ServiceTicketLinkInfo) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceTicketLinkInfo) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *ServiceTicketLinkInfo) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *ServiceTicketLinkInfo) SetUrl(v string) {
	o.Url = &v
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *ServiceTicketLinkInfo) GetInfo() map[string]string {
	if o == nil || IsNil(o.Info) {
		var ret map[string]string
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceTicketLinkInfo) GetInfoOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Info) {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *ServiceTicketLinkInfo) HasInfo() bool {
	if o != nil && !IsNil(o.Info) {
		return true
	}

	return false
}

// SetInfo gets a reference to the given map[string]string and assigns it to the Info field.
func (o *ServiceTicketLinkInfo) SetInfo(v map[string]string) {
	o.Info = &v
}

func (o ServiceTicketLinkInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceTicketLinkInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.LinkText) {
		toSerialize["linkText"] = o.LinkText
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Info) {
		toSerialize["_info"] = o.Info
	}
	return toSerialize, nil
}

type NullableServiceTicketLinkInfo struct {
	value *ServiceTicketLinkInfo
	isSet bool
}

func (v NullableServiceTicketLinkInfo) Get() *ServiceTicketLinkInfo {
	return v.value
}

func (v *NullableServiceTicketLinkInfo) Set(val *ServiceTicketLinkInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceTicketLinkInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceTicketLinkInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceTicketLinkInfo(val *ServiceTicketLinkInfo) *NullableServiceTicketLinkInfo {
	return &NullableServiceTicketLinkInfo{value: val, isSet: true}
}

func (v NullableServiceTicketLinkInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceTicketLinkInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


