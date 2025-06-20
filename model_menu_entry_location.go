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

// checks if the MenuEntryLocation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MenuEntryLocation{}

// MenuEntryLocation struct for MenuEntryLocation
type MenuEntryLocation struct {
	Id *int32 `json:"id,omitempty"`
	Location SystemLocationReference `json:"location"`
	MenuEntry *SystemMenuEntryReference `json:"menuEntry,omitempty"`
	Info *map[string]string `json:"_info,omitempty"`
}

type _MenuEntryLocation MenuEntryLocation

// NewMenuEntryLocation instantiates a new MenuEntryLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMenuEntryLocation(location SystemLocationReference) *MenuEntryLocation {
	this := MenuEntryLocation{}
	this.Location = location
	return &this
}

// NewMenuEntryLocationWithDefaults instantiates a new MenuEntryLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMenuEntryLocationWithDefaults() *MenuEntryLocation {
	this := MenuEntryLocation{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MenuEntryLocation) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MenuEntryLocation) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MenuEntryLocation) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *MenuEntryLocation) SetId(v int32) {
	o.Id = &v
}

// GetLocation returns the Location field value
func (o *MenuEntryLocation) GetLocation() SystemLocationReference {
	if o == nil {
		var ret SystemLocationReference
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *MenuEntryLocation) GetLocationOk() (*SystemLocationReference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *MenuEntryLocation) SetLocation(v SystemLocationReference) {
	o.Location = v
}

// GetMenuEntry returns the MenuEntry field value if set, zero value otherwise.
func (o *MenuEntryLocation) GetMenuEntry() SystemMenuEntryReference {
	if o == nil || IsNil(o.MenuEntry) {
		var ret SystemMenuEntryReference
		return ret
	}
	return *o.MenuEntry
}

// GetMenuEntryOk returns a tuple with the MenuEntry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MenuEntryLocation) GetMenuEntryOk() (*SystemMenuEntryReference, bool) {
	if o == nil || IsNil(o.MenuEntry) {
		return nil, false
	}
	return o.MenuEntry, true
}

// HasMenuEntry returns a boolean if a field has been set.
func (o *MenuEntryLocation) HasMenuEntry() bool {
	if o != nil && !IsNil(o.MenuEntry) {
		return true
	}

	return false
}

// SetMenuEntry gets a reference to the given SystemMenuEntryReference and assigns it to the MenuEntry field.
func (o *MenuEntryLocation) SetMenuEntry(v SystemMenuEntryReference) {
	o.MenuEntry = &v
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *MenuEntryLocation) GetInfo() map[string]string {
	if o == nil || IsNil(o.Info) {
		var ret map[string]string
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MenuEntryLocation) GetInfoOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Info) {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *MenuEntryLocation) HasInfo() bool {
	if o != nil && !IsNil(o.Info) {
		return true
	}

	return false
}

// SetInfo gets a reference to the given map[string]string and assigns it to the Info field.
func (o *MenuEntryLocation) SetInfo(v map[string]string) {
	o.Info = &v
}

func (o MenuEntryLocation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MenuEntryLocation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["location"] = o.Location
	if !IsNil(o.MenuEntry) {
		toSerialize["menuEntry"] = o.MenuEntry
	}
	if !IsNil(o.Info) {
		toSerialize["_info"] = o.Info
	}
	return toSerialize, nil
}

func (o *MenuEntryLocation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"location",
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

	varMenuEntryLocation := _MenuEntryLocation{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMenuEntryLocation)

	if err != nil {
		return err
	}

	*o = MenuEntryLocation(varMenuEntryLocation)

	return err
}

type NullableMenuEntryLocation struct {
	value *MenuEntryLocation
	isSet bool
}

func (v NullableMenuEntryLocation) Get() *MenuEntryLocation {
	return v.value
}

func (v *NullableMenuEntryLocation) Set(val *MenuEntryLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableMenuEntryLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableMenuEntryLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMenuEntryLocation(val *MenuEntryLocation) *NullableMenuEntryLocation {
	return &NullableMenuEntryLocation{value: val, isSet: true}
}

func (v NullableMenuEntryLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMenuEntryLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


