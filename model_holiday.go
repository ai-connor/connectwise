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

// checks if the Holiday type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Holiday{}

// Holiday struct for Holiday
type Holiday struct {
	Id *int32 `json:"id,omitempty"`
	//  Max length: 50;
	Name string `json:"name"`
	// Can be set to false to set a holiday for specific hours (Defaults to True).
	AllDayFlag NullableBool `json:"allDayFlag,omitempty"`
	Date string `json:"date"`
	TimeStart *string `json:"timeStart,omitempty"`
	TimeEnd *string `json:"timeEnd,omitempty"`
	HolidayList *HolidayListReference `json:"holidayList,omitempty"`
	Info *map[string]string `json:"_info,omitempty"`
}

type _Holiday Holiday

// NewHoliday instantiates a new Holiday object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHoliday(name string, date string) *Holiday {
	this := Holiday{}
	this.Name = name
	this.Date = date
	return &this
}

// NewHolidayWithDefaults instantiates a new Holiday object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHolidayWithDefaults() *Holiday {
	this := Holiday{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Holiday) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Holiday) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Holiday) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Holiday) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *Holiday) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Holiday) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Holiday) SetName(v string) {
	o.Name = v
}

// GetAllDayFlag returns the AllDayFlag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Holiday) GetAllDayFlag() bool {
	if o == nil || IsNil(o.AllDayFlag.Get()) {
		var ret bool
		return ret
	}
	return *o.AllDayFlag.Get()
}

// GetAllDayFlagOk returns a tuple with the AllDayFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Holiday) GetAllDayFlagOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AllDayFlag.Get(), o.AllDayFlag.IsSet()
}

// HasAllDayFlag returns a boolean if a field has been set.
func (o *Holiday) HasAllDayFlag() bool {
	if o != nil && o.AllDayFlag.IsSet() {
		return true
	}

	return false
}

// SetAllDayFlag gets a reference to the given NullableBool and assigns it to the AllDayFlag field.
func (o *Holiday) SetAllDayFlag(v bool) {
	o.AllDayFlag.Set(&v)
}
// SetAllDayFlagNil sets the value for AllDayFlag to be an explicit nil
func (o *Holiday) SetAllDayFlagNil() {
	o.AllDayFlag.Set(nil)
}

// UnsetAllDayFlag ensures that no value is present for AllDayFlag, not even an explicit nil
func (o *Holiday) UnsetAllDayFlag() {
	o.AllDayFlag.Unset()
}

// GetDate returns the Date field value
func (o *Holiday) GetDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *Holiday) GetDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *Holiday) SetDate(v string) {
	o.Date = v
}

// GetTimeStart returns the TimeStart field value if set, zero value otherwise.
func (o *Holiday) GetTimeStart() string {
	if o == nil || IsNil(o.TimeStart) {
		var ret string
		return ret
	}
	return *o.TimeStart
}

// GetTimeStartOk returns a tuple with the TimeStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Holiday) GetTimeStartOk() (*string, bool) {
	if o == nil || IsNil(o.TimeStart) {
		return nil, false
	}
	return o.TimeStart, true
}

// HasTimeStart returns a boolean if a field has been set.
func (o *Holiday) HasTimeStart() bool {
	if o != nil && !IsNil(o.TimeStart) {
		return true
	}

	return false
}

// SetTimeStart gets a reference to the given string and assigns it to the TimeStart field.
func (o *Holiday) SetTimeStart(v string) {
	o.TimeStart = &v
}

// GetTimeEnd returns the TimeEnd field value if set, zero value otherwise.
func (o *Holiday) GetTimeEnd() string {
	if o == nil || IsNil(o.TimeEnd) {
		var ret string
		return ret
	}
	return *o.TimeEnd
}

// GetTimeEndOk returns a tuple with the TimeEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Holiday) GetTimeEndOk() (*string, bool) {
	if o == nil || IsNil(o.TimeEnd) {
		return nil, false
	}
	return o.TimeEnd, true
}

// HasTimeEnd returns a boolean if a field has been set.
func (o *Holiday) HasTimeEnd() bool {
	if o != nil && !IsNil(o.TimeEnd) {
		return true
	}

	return false
}

// SetTimeEnd gets a reference to the given string and assigns it to the TimeEnd field.
func (o *Holiday) SetTimeEnd(v string) {
	o.TimeEnd = &v
}

// GetHolidayList returns the HolidayList field value if set, zero value otherwise.
func (o *Holiday) GetHolidayList() HolidayListReference {
	if o == nil || IsNil(o.HolidayList) {
		var ret HolidayListReference
		return ret
	}
	return *o.HolidayList
}

// GetHolidayListOk returns a tuple with the HolidayList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Holiday) GetHolidayListOk() (*HolidayListReference, bool) {
	if o == nil || IsNil(o.HolidayList) {
		return nil, false
	}
	return o.HolidayList, true
}

// HasHolidayList returns a boolean if a field has been set.
func (o *Holiday) HasHolidayList() bool {
	if o != nil && !IsNil(o.HolidayList) {
		return true
	}

	return false
}

// SetHolidayList gets a reference to the given HolidayListReference and assigns it to the HolidayList field.
func (o *Holiday) SetHolidayList(v HolidayListReference) {
	o.HolidayList = &v
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *Holiday) GetInfo() map[string]string {
	if o == nil || IsNil(o.Info) {
		var ret map[string]string
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Holiday) GetInfoOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Info) {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *Holiday) HasInfo() bool {
	if o != nil && !IsNil(o.Info) {
		return true
	}

	return false
}

// SetInfo gets a reference to the given map[string]string and assigns it to the Info field.
func (o *Holiday) SetInfo(v map[string]string) {
	o.Info = &v
}

func (o Holiday) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Holiday) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	if o.AllDayFlag.IsSet() {
		toSerialize["allDayFlag"] = o.AllDayFlag.Get()
	}
	toSerialize["date"] = o.Date
	if !IsNil(o.TimeStart) {
		toSerialize["timeStart"] = o.TimeStart
	}
	if !IsNil(o.TimeEnd) {
		toSerialize["timeEnd"] = o.TimeEnd
	}
	if !IsNil(o.HolidayList) {
		toSerialize["holidayList"] = o.HolidayList
	}
	if !IsNil(o.Info) {
		toSerialize["_info"] = o.Info
	}
	return toSerialize, nil
}

func (o *Holiday) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"date",
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

	varHoliday := _Holiday{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHoliday)

	if err != nil {
		return err
	}

	*o = Holiday(varHoliday)

	return err
}

type NullableHoliday struct {
	value *Holiday
	isSet bool
}

func (v NullableHoliday) Get() *Holiday {
	return v.value
}

func (v *NullableHoliday) Set(val *Holiday) {
	v.value = val
	v.isSet = true
}

func (v NullableHoliday) IsSet() bool {
	return v.isSet
}

func (v *NullableHoliday) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHoliday(val *Holiday) *NullableHoliday {
	return &NullableHoliday{value: val, isSet: true}
}

func (v NullableHoliday) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHoliday) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


