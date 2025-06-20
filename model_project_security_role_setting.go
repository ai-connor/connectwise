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

// checks if the ProjectSecurityRoleSetting type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectSecurityRoleSetting{}

// ProjectSecurityRoleSetting struct for ProjectSecurityRoleSetting
type ProjectSecurityRoleSetting struct {
	Id *int32 `json:"id,omitempty"`
	AddLevel NullableString `json:"addLevel,omitempty"`
	EditLevel NullableString `json:"editLevel,omitempty"`
	DeleteLevel NullableString `json:"deleteLevel,omitempty"`
	InquireLevel NullableString `json:"inquireLevel,omitempty"`
	//  Max length: 50;
	ModuleIdentifier *string `json:"moduleIdentifier,omitempty"`
	MyFlag NullableBool `json:"myFlag,omitempty"`
	Info *map[string]string `json:"_info,omitempty"`
}

// NewProjectSecurityRoleSetting instantiates a new ProjectSecurityRoleSetting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectSecurityRoleSetting() *ProjectSecurityRoleSetting {
	this := ProjectSecurityRoleSetting{}
	return &this
}

// NewProjectSecurityRoleSettingWithDefaults instantiates a new ProjectSecurityRoleSetting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectSecurityRoleSettingWithDefaults() *ProjectSecurityRoleSetting {
	this := ProjectSecurityRoleSetting{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProjectSecurityRoleSetting) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectSecurityRoleSetting) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProjectSecurityRoleSetting) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ProjectSecurityRoleSetting) SetId(v int32) {
	o.Id = &v
}

// GetAddLevel returns the AddLevel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectSecurityRoleSetting) GetAddLevel() string {
	if o == nil || IsNil(o.AddLevel.Get()) {
		var ret string
		return ret
	}
	return *o.AddLevel.Get()
}

// GetAddLevelOk returns a tuple with the AddLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectSecurityRoleSetting) GetAddLevelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AddLevel.Get(), o.AddLevel.IsSet()
}

// HasAddLevel returns a boolean if a field has been set.
func (o *ProjectSecurityRoleSetting) HasAddLevel() bool {
	if o != nil && o.AddLevel.IsSet() {
		return true
	}

	return false
}

// SetAddLevel gets a reference to the given NullableString and assigns it to the AddLevel field.
func (o *ProjectSecurityRoleSetting) SetAddLevel(v string) {
	o.AddLevel.Set(&v)
}
// SetAddLevelNil sets the value for AddLevel to be an explicit nil
func (o *ProjectSecurityRoleSetting) SetAddLevelNil() {
	o.AddLevel.Set(nil)
}

// UnsetAddLevel ensures that no value is present for AddLevel, not even an explicit nil
func (o *ProjectSecurityRoleSetting) UnsetAddLevel() {
	o.AddLevel.Unset()
}

// GetEditLevel returns the EditLevel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectSecurityRoleSetting) GetEditLevel() string {
	if o == nil || IsNil(o.EditLevel.Get()) {
		var ret string
		return ret
	}
	return *o.EditLevel.Get()
}

// GetEditLevelOk returns a tuple with the EditLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectSecurityRoleSetting) GetEditLevelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EditLevel.Get(), o.EditLevel.IsSet()
}

// HasEditLevel returns a boolean if a field has been set.
func (o *ProjectSecurityRoleSetting) HasEditLevel() bool {
	if o != nil && o.EditLevel.IsSet() {
		return true
	}

	return false
}

// SetEditLevel gets a reference to the given NullableString and assigns it to the EditLevel field.
func (o *ProjectSecurityRoleSetting) SetEditLevel(v string) {
	o.EditLevel.Set(&v)
}
// SetEditLevelNil sets the value for EditLevel to be an explicit nil
func (o *ProjectSecurityRoleSetting) SetEditLevelNil() {
	o.EditLevel.Set(nil)
}

// UnsetEditLevel ensures that no value is present for EditLevel, not even an explicit nil
func (o *ProjectSecurityRoleSetting) UnsetEditLevel() {
	o.EditLevel.Unset()
}

// GetDeleteLevel returns the DeleteLevel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectSecurityRoleSetting) GetDeleteLevel() string {
	if o == nil || IsNil(o.DeleteLevel.Get()) {
		var ret string
		return ret
	}
	return *o.DeleteLevel.Get()
}

// GetDeleteLevelOk returns a tuple with the DeleteLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectSecurityRoleSetting) GetDeleteLevelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeleteLevel.Get(), o.DeleteLevel.IsSet()
}

// HasDeleteLevel returns a boolean if a field has been set.
func (o *ProjectSecurityRoleSetting) HasDeleteLevel() bool {
	if o != nil && o.DeleteLevel.IsSet() {
		return true
	}

	return false
}

// SetDeleteLevel gets a reference to the given NullableString and assigns it to the DeleteLevel field.
func (o *ProjectSecurityRoleSetting) SetDeleteLevel(v string) {
	o.DeleteLevel.Set(&v)
}
// SetDeleteLevelNil sets the value for DeleteLevel to be an explicit nil
func (o *ProjectSecurityRoleSetting) SetDeleteLevelNil() {
	o.DeleteLevel.Set(nil)
}

// UnsetDeleteLevel ensures that no value is present for DeleteLevel, not even an explicit nil
func (o *ProjectSecurityRoleSetting) UnsetDeleteLevel() {
	o.DeleteLevel.Unset()
}

// GetInquireLevel returns the InquireLevel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectSecurityRoleSetting) GetInquireLevel() string {
	if o == nil || IsNil(o.InquireLevel.Get()) {
		var ret string
		return ret
	}
	return *o.InquireLevel.Get()
}

// GetInquireLevelOk returns a tuple with the InquireLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectSecurityRoleSetting) GetInquireLevelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InquireLevel.Get(), o.InquireLevel.IsSet()
}

// HasInquireLevel returns a boolean if a field has been set.
func (o *ProjectSecurityRoleSetting) HasInquireLevel() bool {
	if o != nil && o.InquireLevel.IsSet() {
		return true
	}

	return false
}

// SetInquireLevel gets a reference to the given NullableString and assigns it to the InquireLevel field.
func (o *ProjectSecurityRoleSetting) SetInquireLevel(v string) {
	o.InquireLevel.Set(&v)
}
// SetInquireLevelNil sets the value for InquireLevel to be an explicit nil
func (o *ProjectSecurityRoleSetting) SetInquireLevelNil() {
	o.InquireLevel.Set(nil)
}

// UnsetInquireLevel ensures that no value is present for InquireLevel, not even an explicit nil
func (o *ProjectSecurityRoleSetting) UnsetInquireLevel() {
	o.InquireLevel.Unset()
}

// GetModuleIdentifier returns the ModuleIdentifier field value if set, zero value otherwise.
func (o *ProjectSecurityRoleSetting) GetModuleIdentifier() string {
	if o == nil || IsNil(o.ModuleIdentifier) {
		var ret string
		return ret
	}
	return *o.ModuleIdentifier
}

// GetModuleIdentifierOk returns a tuple with the ModuleIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectSecurityRoleSetting) GetModuleIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.ModuleIdentifier) {
		return nil, false
	}
	return o.ModuleIdentifier, true
}

// HasModuleIdentifier returns a boolean if a field has been set.
func (o *ProjectSecurityRoleSetting) HasModuleIdentifier() bool {
	if o != nil && !IsNil(o.ModuleIdentifier) {
		return true
	}

	return false
}

// SetModuleIdentifier gets a reference to the given string and assigns it to the ModuleIdentifier field.
func (o *ProjectSecurityRoleSetting) SetModuleIdentifier(v string) {
	o.ModuleIdentifier = &v
}

// GetMyFlag returns the MyFlag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectSecurityRoleSetting) GetMyFlag() bool {
	if o == nil || IsNil(o.MyFlag.Get()) {
		var ret bool
		return ret
	}
	return *o.MyFlag.Get()
}

// GetMyFlagOk returns a tuple with the MyFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectSecurityRoleSetting) GetMyFlagOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.MyFlag.Get(), o.MyFlag.IsSet()
}

// HasMyFlag returns a boolean if a field has been set.
func (o *ProjectSecurityRoleSetting) HasMyFlag() bool {
	if o != nil && o.MyFlag.IsSet() {
		return true
	}

	return false
}

// SetMyFlag gets a reference to the given NullableBool and assigns it to the MyFlag field.
func (o *ProjectSecurityRoleSetting) SetMyFlag(v bool) {
	o.MyFlag.Set(&v)
}
// SetMyFlagNil sets the value for MyFlag to be an explicit nil
func (o *ProjectSecurityRoleSetting) SetMyFlagNil() {
	o.MyFlag.Set(nil)
}

// UnsetMyFlag ensures that no value is present for MyFlag, not even an explicit nil
func (o *ProjectSecurityRoleSetting) UnsetMyFlag() {
	o.MyFlag.Unset()
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *ProjectSecurityRoleSetting) GetInfo() map[string]string {
	if o == nil || IsNil(o.Info) {
		var ret map[string]string
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectSecurityRoleSetting) GetInfoOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Info) {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *ProjectSecurityRoleSetting) HasInfo() bool {
	if o != nil && !IsNil(o.Info) {
		return true
	}

	return false
}

// SetInfo gets a reference to the given map[string]string and assigns it to the Info field.
func (o *ProjectSecurityRoleSetting) SetInfo(v map[string]string) {
	o.Info = &v
}

func (o ProjectSecurityRoleSetting) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectSecurityRoleSetting) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.AddLevel.IsSet() {
		toSerialize["addLevel"] = o.AddLevel.Get()
	}
	if o.EditLevel.IsSet() {
		toSerialize["editLevel"] = o.EditLevel.Get()
	}
	if o.DeleteLevel.IsSet() {
		toSerialize["deleteLevel"] = o.DeleteLevel.Get()
	}
	if o.InquireLevel.IsSet() {
		toSerialize["inquireLevel"] = o.InquireLevel.Get()
	}
	if !IsNil(o.ModuleIdentifier) {
		toSerialize["moduleIdentifier"] = o.ModuleIdentifier
	}
	if o.MyFlag.IsSet() {
		toSerialize["myFlag"] = o.MyFlag.Get()
	}
	if !IsNil(o.Info) {
		toSerialize["_info"] = o.Info
	}
	return toSerialize, nil
}

type NullableProjectSecurityRoleSetting struct {
	value *ProjectSecurityRoleSetting
	isSet bool
}

func (v NullableProjectSecurityRoleSetting) Get() *ProjectSecurityRoleSetting {
	return v.value
}

func (v *NullableProjectSecurityRoleSetting) Set(val *ProjectSecurityRoleSetting) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectSecurityRoleSetting) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectSecurityRoleSetting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectSecurityRoleSetting(val *ProjectSecurityRoleSetting) *NullableProjectSecurityRoleSetting {
	return &NullableProjectSecurityRoleSetting{value: val, isSet: true}
}

func (v NullableProjectSecurityRoleSetting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectSecurityRoleSetting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


