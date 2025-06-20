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

// checks if the GenericNameIdDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GenericNameIdDTO{}

// GenericNameIdDTO struct for GenericNameIdDTO
type GenericNameIdDTO struct {
	Id *int32 `json:"id,omitempty"`
	Tag *string `json:"tag,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewGenericNameIdDTO instantiates a new GenericNameIdDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenericNameIdDTO() *GenericNameIdDTO {
	this := GenericNameIdDTO{}
	return &this
}

// NewGenericNameIdDTOWithDefaults instantiates a new GenericNameIdDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenericNameIdDTOWithDefaults() *GenericNameIdDTO {
	this := GenericNameIdDTO{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GenericNameIdDTO) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericNameIdDTO) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GenericNameIdDTO) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *GenericNameIdDTO) SetId(v int32) {
	o.Id = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *GenericNameIdDTO) GetTag() string {
	if o == nil || IsNil(o.Tag) {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericNameIdDTO) GetTagOk() (*string, bool) {
	if o == nil || IsNil(o.Tag) {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *GenericNameIdDTO) HasTag() bool {
	if o != nil && !IsNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *GenericNameIdDTO) SetTag(v string) {
	o.Tag = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GenericNameIdDTO) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericNameIdDTO) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GenericNameIdDTO) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GenericNameIdDTO) SetName(v string) {
	o.Name = &v
}

func (o GenericNameIdDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GenericNameIdDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableGenericNameIdDTO struct {
	value *GenericNameIdDTO
	isSet bool
}

func (v NullableGenericNameIdDTO) Get() *GenericNameIdDTO {
	return v.value
}

func (v *NullableGenericNameIdDTO) Set(val *GenericNameIdDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableGenericNameIdDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableGenericNameIdDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenericNameIdDTO(val *GenericNameIdDTO) *NullableGenericNameIdDTO {
	return &NullableGenericNameIdDTO{value: val, isSet: true}
}

func (v NullableGenericNameIdDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenericNameIdDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


