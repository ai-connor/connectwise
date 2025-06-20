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

// checks if the ConvertToProject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConvertToProject{}

// ConvertToProject struct for ConvertToProject
type ConvertToProject struct {
	Id *int32 `json:"id,omitempty"`
	RecordType NullableString `json:"recordType,omitempty"`
	Project *ProjectReference `json:"project,omitempty"`
	Phase ProjectPhaseReference `json:"phase"`
	WbsCode string `json:"wbsCode"`
}

type _ConvertToProject ConvertToProject

// NewConvertToProject instantiates a new ConvertToProject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConvertToProject(phase ProjectPhaseReference, wbsCode string) *ConvertToProject {
	this := ConvertToProject{}
	this.Phase = phase
	this.WbsCode = wbsCode
	return &this
}

// NewConvertToProjectWithDefaults instantiates a new ConvertToProject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConvertToProjectWithDefaults() *ConvertToProject {
	this := ConvertToProject{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ConvertToProject) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvertToProject) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ConvertToProject) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ConvertToProject) SetId(v int32) {
	o.Id = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConvertToProject) GetRecordType() string {
	if o == nil || IsNil(o.RecordType.Get()) {
		var ret string
		return ret
	}
	return *o.RecordType.Get()
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConvertToProject) GetRecordTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RecordType.Get(), o.RecordType.IsSet()
}

// HasRecordType returns a boolean if a field has been set.
func (o *ConvertToProject) HasRecordType() bool {
	if o != nil && o.RecordType.IsSet() {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given NullableString and assigns it to the RecordType field.
func (o *ConvertToProject) SetRecordType(v string) {
	o.RecordType.Set(&v)
}
// SetRecordTypeNil sets the value for RecordType to be an explicit nil
func (o *ConvertToProject) SetRecordTypeNil() {
	o.RecordType.Set(nil)
}

// UnsetRecordType ensures that no value is present for RecordType, not even an explicit nil
func (o *ConvertToProject) UnsetRecordType() {
	o.RecordType.Unset()
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *ConvertToProject) GetProject() ProjectReference {
	if o == nil || IsNil(o.Project) {
		var ret ProjectReference
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvertToProject) GetProjectOk() (*ProjectReference, bool) {
	if o == nil || IsNil(o.Project) {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *ConvertToProject) HasProject() bool {
	if o != nil && !IsNil(o.Project) {
		return true
	}

	return false
}

// SetProject gets a reference to the given ProjectReference and assigns it to the Project field.
func (o *ConvertToProject) SetProject(v ProjectReference) {
	o.Project = &v
}

// GetPhase returns the Phase field value
func (o *ConvertToProject) GetPhase() ProjectPhaseReference {
	if o == nil {
		var ret ProjectPhaseReference
		return ret
	}

	return o.Phase
}

// GetPhaseOk returns a tuple with the Phase field value
// and a boolean to check if the value has been set.
func (o *ConvertToProject) GetPhaseOk() (*ProjectPhaseReference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Phase, true
}

// SetPhase sets field value
func (o *ConvertToProject) SetPhase(v ProjectPhaseReference) {
	o.Phase = v
}

// GetWbsCode returns the WbsCode field value
func (o *ConvertToProject) GetWbsCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WbsCode
}

// GetWbsCodeOk returns a tuple with the WbsCode field value
// and a boolean to check if the value has been set.
func (o *ConvertToProject) GetWbsCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WbsCode, true
}

// SetWbsCode sets field value
func (o *ConvertToProject) SetWbsCode(v string) {
	o.WbsCode = v
}

func (o ConvertToProject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConvertToProject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.RecordType.IsSet() {
		toSerialize["recordType"] = o.RecordType.Get()
	}
	if !IsNil(o.Project) {
		toSerialize["project"] = o.Project
	}
	toSerialize["phase"] = o.Phase
	toSerialize["wbsCode"] = o.WbsCode
	return toSerialize, nil
}

func (o *ConvertToProject) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"phase",
		"wbsCode",
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

	varConvertToProject := _ConvertToProject{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varConvertToProject)

	if err != nil {
		return err
	}

	*o = ConvertToProject(varConvertToProject)

	return err
}

type NullableConvertToProject struct {
	value *ConvertToProject
	isSet bool
}

func (v NullableConvertToProject) Get() *ConvertToProject {
	return v.value
}

func (v *NullableConvertToProject) Set(val *ConvertToProject) {
	v.value = val
	v.isSet = true
}

func (v NullableConvertToProject) IsSet() bool {
	return v.isSet
}

func (v *NullableConvertToProject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConvertToProject(val *ConvertToProject) *NullableConvertToProject {
	return &NullableConvertToProject{value: val, isSet: true}
}

func (v NullableConvertToProject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConvertToProject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


