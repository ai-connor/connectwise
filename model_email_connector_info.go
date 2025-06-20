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

// checks if the EmailConnectorInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailConnectorInfo{}

// EmailConnectorInfo struct for EmailConnectorInfo
type EmailConnectorInfo struct {
	Id *int32 `json:"id,omitempty"`
	DefaultCompany *CompanyReference `json:"defaultCompany,omitempty"`
	ImapSetup *ImapSetupReference `json:"imapSetup,omitempty"`
	Info *map[string]string `json:"_info,omitempty"`
}

// NewEmailConnectorInfo instantiates a new EmailConnectorInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailConnectorInfo() *EmailConnectorInfo {
	this := EmailConnectorInfo{}
	return &this
}

// NewEmailConnectorInfoWithDefaults instantiates a new EmailConnectorInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailConnectorInfoWithDefaults() *EmailConnectorInfo {
	this := EmailConnectorInfo{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EmailConnectorInfo) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailConnectorInfo) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EmailConnectorInfo) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *EmailConnectorInfo) SetId(v int32) {
	o.Id = &v
}

// GetDefaultCompany returns the DefaultCompany field value if set, zero value otherwise.
func (o *EmailConnectorInfo) GetDefaultCompany() CompanyReference {
	if o == nil || IsNil(o.DefaultCompany) {
		var ret CompanyReference
		return ret
	}
	return *o.DefaultCompany
}

// GetDefaultCompanyOk returns a tuple with the DefaultCompany field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailConnectorInfo) GetDefaultCompanyOk() (*CompanyReference, bool) {
	if o == nil || IsNil(o.DefaultCompany) {
		return nil, false
	}
	return o.DefaultCompany, true
}

// HasDefaultCompany returns a boolean if a field has been set.
func (o *EmailConnectorInfo) HasDefaultCompany() bool {
	if o != nil && !IsNil(o.DefaultCompany) {
		return true
	}

	return false
}

// SetDefaultCompany gets a reference to the given CompanyReference and assigns it to the DefaultCompany field.
func (o *EmailConnectorInfo) SetDefaultCompany(v CompanyReference) {
	o.DefaultCompany = &v
}

// GetImapSetup returns the ImapSetup field value if set, zero value otherwise.
func (o *EmailConnectorInfo) GetImapSetup() ImapSetupReference {
	if o == nil || IsNil(o.ImapSetup) {
		var ret ImapSetupReference
		return ret
	}
	return *o.ImapSetup
}

// GetImapSetupOk returns a tuple with the ImapSetup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailConnectorInfo) GetImapSetupOk() (*ImapSetupReference, bool) {
	if o == nil || IsNil(o.ImapSetup) {
		return nil, false
	}
	return o.ImapSetup, true
}

// HasImapSetup returns a boolean if a field has been set.
func (o *EmailConnectorInfo) HasImapSetup() bool {
	if o != nil && !IsNil(o.ImapSetup) {
		return true
	}

	return false
}

// SetImapSetup gets a reference to the given ImapSetupReference and assigns it to the ImapSetup field.
func (o *EmailConnectorInfo) SetImapSetup(v ImapSetupReference) {
	o.ImapSetup = &v
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *EmailConnectorInfo) GetInfo() map[string]string {
	if o == nil || IsNil(o.Info) {
		var ret map[string]string
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailConnectorInfo) GetInfoOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Info) {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *EmailConnectorInfo) HasInfo() bool {
	if o != nil && !IsNil(o.Info) {
		return true
	}

	return false
}

// SetInfo gets a reference to the given map[string]string and assigns it to the Info field.
func (o *EmailConnectorInfo) SetInfo(v map[string]string) {
	o.Info = &v
}

func (o EmailConnectorInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailConnectorInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.DefaultCompany) {
		toSerialize["defaultCompany"] = o.DefaultCompany
	}
	if !IsNil(o.ImapSetup) {
		toSerialize["imapSetup"] = o.ImapSetup
	}
	if !IsNil(o.Info) {
		toSerialize["_info"] = o.Info
	}
	return toSerialize, nil
}

type NullableEmailConnectorInfo struct {
	value *EmailConnectorInfo
	isSet bool
}

func (v NullableEmailConnectorInfo) Get() *EmailConnectorInfo {
	return v.value
}

func (v *NullableEmailConnectorInfo) Set(val *EmailConnectorInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailConnectorInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailConnectorInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailConnectorInfo(val *EmailConnectorInfo) *NullableEmailConnectorInfo {
	return &NullableEmailConnectorInfo{value: val, isSet: true}
}

func (v NullableEmailConnectorInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailConnectorInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


