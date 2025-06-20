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

// checks if the MemberLinkSsoUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MemberLinkSsoUser{}

// MemberLinkSsoUser struct for MemberLinkSsoUser
type MemberLinkSsoUser struct {
	//  Max length: 100;
	SsoUserId *string `json:"ssoUserId,omitempty"`
}

// NewMemberLinkSsoUser instantiates a new MemberLinkSsoUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMemberLinkSsoUser() *MemberLinkSsoUser {
	this := MemberLinkSsoUser{}
	return &this
}

// NewMemberLinkSsoUserWithDefaults instantiates a new MemberLinkSsoUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMemberLinkSsoUserWithDefaults() *MemberLinkSsoUser {
	this := MemberLinkSsoUser{}
	return &this
}

// GetSsoUserId returns the SsoUserId field value if set, zero value otherwise.
func (o *MemberLinkSsoUser) GetSsoUserId() string {
	if o == nil || IsNil(o.SsoUserId) {
		var ret string
		return ret
	}
	return *o.SsoUserId
}

// GetSsoUserIdOk returns a tuple with the SsoUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberLinkSsoUser) GetSsoUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.SsoUserId) {
		return nil, false
	}
	return o.SsoUserId, true
}

// HasSsoUserId returns a boolean if a field has been set.
func (o *MemberLinkSsoUser) HasSsoUserId() bool {
	if o != nil && !IsNil(o.SsoUserId) {
		return true
	}

	return false
}

// SetSsoUserId gets a reference to the given string and assigns it to the SsoUserId field.
func (o *MemberLinkSsoUser) SetSsoUserId(v string) {
	o.SsoUserId = &v
}

func (o MemberLinkSsoUser) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MemberLinkSsoUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SsoUserId) {
		toSerialize["ssoUserId"] = o.SsoUserId
	}
	return toSerialize, nil
}

type NullableMemberLinkSsoUser struct {
	value *MemberLinkSsoUser
	isSet bool
}

func (v NullableMemberLinkSsoUser) Get() *MemberLinkSsoUser {
	return v.value
}

func (v *NullableMemberLinkSsoUser) Set(val *MemberLinkSsoUser) {
	v.value = val
	v.isSet = true
}

func (v NullableMemberLinkSsoUser) IsSet() bool {
	return v.isSet
}

func (v *NullableMemberLinkSsoUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMemberLinkSsoUser(val *MemberLinkSsoUser) *NullableMemberLinkSsoUser {
	return &NullableMemberLinkSsoUser{value: val, isSet: true}
}

func (v NullableMemberLinkSsoUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMemberLinkSsoUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


