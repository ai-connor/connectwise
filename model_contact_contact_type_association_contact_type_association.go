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

// checks if the ContactContactTypeAssociationContactTypeAssociation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContactContactTypeAssociationContactTypeAssociation{}

// ContactContactTypeAssociationContactTypeAssociation struct for ContactContactTypeAssociationContactTypeAssociation
type ContactContactTypeAssociationContactTypeAssociation struct {
	Id *int32 `json:"id,omitempty"`
	Type ContactTypeReference `json:"type"`
	Contact ContactReference `json:"contact"`
	Info *map[string]string `json:"_info,omitempty"`
}

type _ContactContactTypeAssociationContactTypeAssociation ContactContactTypeAssociationContactTypeAssociation

// NewContactContactTypeAssociationContactTypeAssociation instantiates a new ContactContactTypeAssociationContactTypeAssociation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContactContactTypeAssociationContactTypeAssociation(type_ ContactTypeReference, contact ContactReference) *ContactContactTypeAssociationContactTypeAssociation {
	this := ContactContactTypeAssociationContactTypeAssociation{}
	this.Type = type_
	this.Contact = contact
	return &this
}

// NewContactContactTypeAssociationContactTypeAssociationWithDefaults instantiates a new ContactContactTypeAssociationContactTypeAssociation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContactContactTypeAssociationContactTypeAssociationWithDefaults() *ContactContactTypeAssociationContactTypeAssociation {
	this := ContactContactTypeAssociationContactTypeAssociation{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ContactContactTypeAssociationContactTypeAssociation) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactContactTypeAssociationContactTypeAssociation) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ContactContactTypeAssociationContactTypeAssociation) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ContactContactTypeAssociationContactTypeAssociation) SetId(v int32) {
	o.Id = &v
}

// GetType returns the Type field value
func (o *ContactContactTypeAssociationContactTypeAssociation) GetType() ContactTypeReference {
	if o == nil {
		var ret ContactTypeReference
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ContactContactTypeAssociationContactTypeAssociation) GetTypeOk() (*ContactTypeReference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ContactContactTypeAssociationContactTypeAssociation) SetType(v ContactTypeReference) {
	o.Type = v
}

// GetContact returns the Contact field value
func (o *ContactContactTypeAssociationContactTypeAssociation) GetContact() ContactReference {
	if o == nil {
		var ret ContactReference
		return ret
	}

	return o.Contact
}

// GetContactOk returns a tuple with the Contact field value
// and a boolean to check if the value has been set.
func (o *ContactContactTypeAssociationContactTypeAssociation) GetContactOk() (*ContactReference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Contact, true
}

// SetContact sets field value
func (o *ContactContactTypeAssociationContactTypeAssociation) SetContact(v ContactReference) {
	o.Contact = v
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *ContactContactTypeAssociationContactTypeAssociation) GetInfo() map[string]string {
	if o == nil || IsNil(o.Info) {
		var ret map[string]string
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactContactTypeAssociationContactTypeAssociation) GetInfoOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Info) {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *ContactContactTypeAssociationContactTypeAssociation) HasInfo() bool {
	if o != nil && !IsNil(o.Info) {
		return true
	}

	return false
}

// SetInfo gets a reference to the given map[string]string and assigns it to the Info field.
func (o *ContactContactTypeAssociationContactTypeAssociation) SetInfo(v map[string]string) {
	o.Info = &v
}

func (o ContactContactTypeAssociationContactTypeAssociation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContactContactTypeAssociationContactTypeAssociation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["type"] = o.Type
	toSerialize["contact"] = o.Contact
	if !IsNil(o.Info) {
		toSerialize["_info"] = o.Info
	}
	return toSerialize, nil
}

func (o *ContactContactTypeAssociationContactTypeAssociation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"contact",
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

	varContactContactTypeAssociationContactTypeAssociation := _ContactContactTypeAssociationContactTypeAssociation{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varContactContactTypeAssociationContactTypeAssociation)

	if err != nil {
		return err
	}

	*o = ContactContactTypeAssociationContactTypeAssociation(varContactContactTypeAssociationContactTypeAssociation)

	return err
}

type NullableContactContactTypeAssociationContactTypeAssociation struct {
	value *ContactContactTypeAssociationContactTypeAssociation
	isSet bool
}

func (v NullableContactContactTypeAssociationContactTypeAssociation) Get() *ContactContactTypeAssociationContactTypeAssociation {
	return v.value
}

func (v *NullableContactContactTypeAssociationContactTypeAssociation) Set(val *ContactContactTypeAssociationContactTypeAssociation) {
	v.value = val
	v.isSet = true
}

func (v NullableContactContactTypeAssociationContactTypeAssociation) IsSet() bool {
	return v.isSet
}

func (v *NullableContactContactTypeAssociationContactTypeAssociation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContactContactTypeAssociationContactTypeAssociation(val *ContactContactTypeAssociationContactTypeAssociation) *NullableContactContactTypeAssociationContactTypeAssociation {
	return &NullableContactContactTypeAssociationContactTypeAssociation{value: val, isSet: true}
}

func (v NullableContactContactTypeAssociationContactTypeAssociation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContactContactTypeAssociationContactTypeAssociation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


