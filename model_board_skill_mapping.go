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

// checks if the BoardSkillMapping type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BoardSkillMapping{}

// BoardSkillMapping struct for BoardSkillMapping
type BoardSkillMapping struct {
	Id *int32 `json:"id,omitempty"`
	Type ServiceTypeReference `json:"type"`
	SubType *ServiceSubTypeReference `json:"subType,omitempty"`
	Item *ServiceItemReference `json:"item,omitempty"`
	SkillCategory SkillCategoryReference `json:"skillCategory"`
	Skill SkillReference `json:"skill"`
	Board *BoardReference `json:"board,omitempty"`
	Info *map[string]string `json:"_info,omitempty"`
}

type _BoardSkillMapping BoardSkillMapping

// NewBoardSkillMapping instantiates a new BoardSkillMapping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBoardSkillMapping(type_ ServiceTypeReference, skillCategory SkillCategoryReference, skill SkillReference) *BoardSkillMapping {
	this := BoardSkillMapping{}
	this.Type = type_
	this.SkillCategory = skillCategory
	this.Skill = skill
	return &this
}

// NewBoardSkillMappingWithDefaults instantiates a new BoardSkillMapping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBoardSkillMappingWithDefaults() *BoardSkillMapping {
	this := BoardSkillMapping{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BoardSkillMapping) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BoardSkillMapping) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BoardSkillMapping) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *BoardSkillMapping) SetId(v int32) {
	o.Id = &v
}

// GetType returns the Type field value
func (o *BoardSkillMapping) GetType() ServiceTypeReference {
	if o == nil {
		var ret ServiceTypeReference
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BoardSkillMapping) GetTypeOk() (*ServiceTypeReference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BoardSkillMapping) SetType(v ServiceTypeReference) {
	o.Type = v
}

// GetSubType returns the SubType field value if set, zero value otherwise.
func (o *BoardSkillMapping) GetSubType() ServiceSubTypeReference {
	if o == nil || IsNil(o.SubType) {
		var ret ServiceSubTypeReference
		return ret
	}
	return *o.SubType
}

// GetSubTypeOk returns a tuple with the SubType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BoardSkillMapping) GetSubTypeOk() (*ServiceSubTypeReference, bool) {
	if o == nil || IsNil(o.SubType) {
		return nil, false
	}
	return o.SubType, true
}

// HasSubType returns a boolean if a field has been set.
func (o *BoardSkillMapping) HasSubType() bool {
	if o != nil && !IsNil(o.SubType) {
		return true
	}

	return false
}

// SetSubType gets a reference to the given ServiceSubTypeReference and assigns it to the SubType field.
func (o *BoardSkillMapping) SetSubType(v ServiceSubTypeReference) {
	o.SubType = &v
}

// GetItem returns the Item field value if set, zero value otherwise.
func (o *BoardSkillMapping) GetItem() ServiceItemReference {
	if o == nil || IsNil(o.Item) {
		var ret ServiceItemReference
		return ret
	}
	return *o.Item
}

// GetItemOk returns a tuple with the Item field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BoardSkillMapping) GetItemOk() (*ServiceItemReference, bool) {
	if o == nil || IsNil(o.Item) {
		return nil, false
	}
	return o.Item, true
}

// HasItem returns a boolean if a field has been set.
func (o *BoardSkillMapping) HasItem() bool {
	if o != nil && !IsNil(o.Item) {
		return true
	}

	return false
}

// SetItem gets a reference to the given ServiceItemReference and assigns it to the Item field.
func (o *BoardSkillMapping) SetItem(v ServiceItemReference) {
	o.Item = &v
}

// GetSkillCategory returns the SkillCategory field value
func (o *BoardSkillMapping) GetSkillCategory() SkillCategoryReference {
	if o == nil {
		var ret SkillCategoryReference
		return ret
	}

	return o.SkillCategory
}

// GetSkillCategoryOk returns a tuple with the SkillCategory field value
// and a boolean to check if the value has been set.
func (o *BoardSkillMapping) GetSkillCategoryOk() (*SkillCategoryReference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SkillCategory, true
}

// SetSkillCategory sets field value
func (o *BoardSkillMapping) SetSkillCategory(v SkillCategoryReference) {
	o.SkillCategory = v
}

// GetSkill returns the Skill field value
func (o *BoardSkillMapping) GetSkill() SkillReference {
	if o == nil {
		var ret SkillReference
		return ret
	}

	return o.Skill
}

// GetSkillOk returns a tuple with the Skill field value
// and a boolean to check if the value has been set.
func (o *BoardSkillMapping) GetSkillOk() (*SkillReference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Skill, true
}

// SetSkill sets field value
func (o *BoardSkillMapping) SetSkill(v SkillReference) {
	o.Skill = v
}

// GetBoard returns the Board field value if set, zero value otherwise.
func (o *BoardSkillMapping) GetBoard() BoardReference {
	if o == nil || IsNil(o.Board) {
		var ret BoardReference
		return ret
	}
	return *o.Board
}

// GetBoardOk returns a tuple with the Board field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BoardSkillMapping) GetBoardOk() (*BoardReference, bool) {
	if o == nil || IsNil(o.Board) {
		return nil, false
	}
	return o.Board, true
}

// HasBoard returns a boolean if a field has been set.
func (o *BoardSkillMapping) HasBoard() bool {
	if o != nil && !IsNil(o.Board) {
		return true
	}

	return false
}

// SetBoard gets a reference to the given BoardReference and assigns it to the Board field.
func (o *BoardSkillMapping) SetBoard(v BoardReference) {
	o.Board = &v
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *BoardSkillMapping) GetInfo() map[string]string {
	if o == nil || IsNil(o.Info) {
		var ret map[string]string
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BoardSkillMapping) GetInfoOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Info) {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *BoardSkillMapping) HasInfo() bool {
	if o != nil && !IsNil(o.Info) {
		return true
	}

	return false
}

// SetInfo gets a reference to the given map[string]string and assigns it to the Info field.
func (o *BoardSkillMapping) SetInfo(v map[string]string) {
	o.Info = &v
}

func (o BoardSkillMapping) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BoardSkillMapping) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["type"] = o.Type
	if !IsNil(o.SubType) {
		toSerialize["subType"] = o.SubType
	}
	if !IsNil(o.Item) {
		toSerialize["item"] = o.Item
	}
	toSerialize["skillCategory"] = o.SkillCategory
	toSerialize["skill"] = o.Skill
	if !IsNil(o.Board) {
		toSerialize["board"] = o.Board
	}
	if !IsNil(o.Info) {
		toSerialize["_info"] = o.Info
	}
	return toSerialize, nil
}

func (o *BoardSkillMapping) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"skillCategory",
		"skill",
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

	varBoardSkillMapping := _BoardSkillMapping{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBoardSkillMapping)

	if err != nil {
		return err
	}

	*o = BoardSkillMapping(varBoardSkillMapping)

	return err
}

type NullableBoardSkillMapping struct {
	value *BoardSkillMapping
	isSet bool
}

func (v NullableBoardSkillMapping) Get() *BoardSkillMapping {
	return v.value
}

func (v *NullableBoardSkillMapping) Set(val *BoardSkillMapping) {
	v.value = val
	v.isSet = true
}

func (v NullableBoardSkillMapping) IsSet() bool {
	return v.isSet
}

func (v *NullableBoardSkillMapping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBoardSkillMapping(val *BoardSkillMapping) *NullableBoardSkillMapping {
	return &NullableBoardSkillMapping{value: val, isSet: true}
}

func (v NullableBoardSkillMapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBoardSkillMapping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


