/*
Connectwise Manage Public Endpoints

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2025.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cwapi

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the ProjectTeamMember type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectTeamMember{}

// ProjectTeamMember struct for ProjectTeamMember
type ProjectTeamMember struct {
	Id *int32 `json:"id,omitempty"`
	ProjectId NullableInt32 `json:"projectId,omitempty"`
	Hours NullableFloat64 `json:"hours,omitempty"`
	Member MemberReference `json:"member"`
	ProjectRole ProjectRoleReference `json:"projectRole"`
	WorkRole *WorkRoleReference `json:"workRole,omitempty"`
	StartDate *time.Time `json:"startDate,omitempty"`
	EndDate *time.Time `json:"endDate,omitempty"`
	Info *map[string]string `json:"_info,omitempty"`
}

type _ProjectTeamMember ProjectTeamMember

// NewProjectTeamMember instantiates a new ProjectTeamMember object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectTeamMember(member MemberReference, projectRole ProjectRoleReference) *ProjectTeamMember {
	this := ProjectTeamMember{}
	this.Member = member
	this.ProjectRole = projectRole
	return &this
}

// NewProjectTeamMemberWithDefaults instantiates a new ProjectTeamMember object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectTeamMemberWithDefaults() *ProjectTeamMember {
	this := ProjectTeamMember{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProjectTeamMember) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectTeamMember) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProjectTeamMember) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ProjectTeamMember) SetId(v int32) {
	o.Id = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectTeamMember) GetProjectId() int32 {
	if o == nil || IsNil(o.ProjectId.Get()) {
		var ret int32
		return ret
	}
	return *o.ProjectId.Get()
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectTeamMember) GetProjectIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProjectId.Get(), o.ProjectId.IsSet()
}

// HasProjectId returns a boolean if a field has been set.
func (o *ProjectTeamMember) HasProjectId() bool {
	if o != nil && o.ProjectId.IsSet() {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given NullableInt32 and assigns it to the ProjectId field.
func (o *ProjectTeamMember) SetProjectId(v int32) {
	o.ProjectId.Set(&v)
}
// SetProjectIdNil sets the value for ProjectId to be an explicit nil
func (o *ProjectTeamMember) SetProjectIdNil() {
	o.ProjectId.Set(nil)
}

// UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
func (o *ProjectTeamMember) UnsetProjectId() {
	o.ProjectId.Unset()
}

// GetHours returns the Hours field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectTeamMember) GetHours() float64 {
	if o == nil || IsNil(o.Hours.Get()) {
		var ret float64
		return ret
	}
	return *o.Hours.Get()
}

// GetHoursOk returns a tuple with the Hours field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectTeamMember) GetHoursOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Hours.Get(), o.Hours.IsSet()
}

// HasHours returns a boolean if a field has been set.
func (o *ProjectTeamMember) HasHours() bool {
	if o != nil && o.Hours.IsSet() {
		return true
	}

	return false
}

// SetHours gets a reference to the given NullableFloat64 and assigns it to the Hours field.
func (o *ProjectTeamMember) SetHours(v float64) {
	o.Hours.Set(&v)
}
// SetHoursNil sets the value for Hours to be an explicit nil
func (o *ProjectTeamMember) SetHoursNil() {
	o.Hours.Set(nil)
}

// UnsetHours ensures that no value is present for Hours, not even an explicit nil
func (o *ProjectTeamMember) UnsetHours() {
	o.Hours.Unset()
}

// GetMember returns the Member field value
func (o *ProjectTeamMember) GetMember() MemberReference {
	if o == nil {
		var ret MemberReference
		return ret
	}

	return o.Member
}

// GetMemberOk returns a tuple with the Member field value
// and a boolean to check if the value has been set.
func (o *ProjectTeamMember) GetMemberOk() (*MemberReference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Member, true
}

// SetMember sets field value
func (o *ProjectTeamMember) SetMember(v MemberReference) {
	o.Member = v
}

// GetProjectRole returns the ProjectRole field value
func (o *ProjectTeamMember) GetProjectRole() ProjectRoleReference {
	if o == nil {
		var ret ProjectRoleReference
		return ret
	}

	return o.ProjectRole
}

// GetProjectRoleOk returns a tuple with the ProjectRole field value
// and a boolean to check if the value has been set.
func (o *ProjectTeamMember) GetProjectRoleOk() (*ProjectRoleReference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectRole, true
}

// SetProjectRole sets field value
func (o *ProjectTeamMember) SetProjectRole(v ProjectRoleReference) {
	o.ProjectRole = v
}

// GetWorkRole returns the WorkRole field value if set, zero value otherwise.
func (o *ProjectTeamMember) GetWorkRole() WorkRoleReference {
	if o == nil || IsNil(o.WorkRole) {
		var ret WorkRoleReference
		return ret
	}
	return *o.WorkRole
}

// GetWorkRoleOk returns a tuple with the WorkRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectTeamMember) GetWorkRoleOk() (*WorkRoleReference, bool) {
	if o == nil || IsNil(o.WorkRole) {
		return nil, false
	}
	return o.WorkRole, true
}

// HasWorkRole returns a boolean if a field has been set.
func (o *ProjectTeamMember) HasWorkRole() bool {
	if o != nil && !IsNil(o.WorkRole) {
		return true
	}

	return false
}

// SetWorkRole gets a reference to the given WorkRoleReference and assigns it to the WorkRole field.
func (o *ProjectTeamMember) SetWorkRole(v WorkRoleReference) {
	o.WorkRole = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *ProjectTeamMember) GetStartDate() time.Time {
	if o == nil || IsNil(o.StartDate) {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectTeamMember) GetStartDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *ProjectTeamMember) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *ProjectTeamMember) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *ProjectTeamMember) GetEndDate() time.Time {
	if o == nil || IsNil(o.EndDate) {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectTeamMember) GetEndDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *ProjectTeamMember) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *ProjectTeamMember) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *ProjectTeamMember) GetInfo() map[string]string {
	if o == nil || IsNil(o.Info) {
		var ret map[string]string
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectTeamMember) GetInfoOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Info) {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *ProjectTeamMember) HasInfo() bool {
	if o != nil && !IsNil(o.Info) {
		return true
	}

	return false
}

// SetInfo gets a reference to the given map[string]string and assigns it to the Info field.
func (o *ProjectTeamMember) SetInfo(v map[string]string) {
	o.Info = &v
}

func (o ProjectTeamMember) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectTeamMember) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.ProjectId.IsSet() {
		toSerialize["projectId"] = o.ProjectId.Get()
	}
	if o.Hours.IsSet() {
		toSerialize["hours"] = o.Hours.Get()
	}
	toSerialize["member"] = o.Member
	toSerialize["projectRole"] = o.ProjectRole
	if !IsNil(o.WorkRole) {
		toSerialize["workRole"] = o.WorkRole
	}
	if !IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	if !IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	if !IsNil(o.Info) {
		toSerialize["_info"] = o.Info
	}
	return toSerialize, nil
}

func (o *ProjectTeamMember) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"member",
		"projectRole",
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

	varProjectTeamMember := _ProjectTeamMember{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProjectTeamMember)

	if err != nil {
		return err
	}

	*o = ProjectTeamMember(varProjectTeamMember)

	return err
}

type NullableProjectTeamMember struct {
	value *ProjectTeamMember
	isSet bool
}

func (v NullableProjectTeamMember) Get() *ProjectTeamMember {
	return v.value
}

func (v *NullableProjectTeamMember) Set(val *ProjectTeamMember) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectTeamMember) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectTeamMember) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectTeamMember(val *ProjectTeamMember) *NullableProjectTeamMember {
	return &NullableProjectTeamMember{value: val, isSet: true}
}

func (v NullableProjectTeamMember) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectTeamMember) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


