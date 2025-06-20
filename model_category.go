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

// checks if the Category type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Category{}

// Category struct for Category
type Category struct {
	Id *int32 `json:"id,omitempty"`
	//  Max length: 50;
	Name string `json:"name"`
	InactiveFlag NullableBool `json:"inactiveFlag,omitempty"`
	//  Max length: 10;
	PriceLevelXref *string `json:"priceLevelXref,omitempty"`
	//  Max length: 50;
	IntegrationXref *string `json:"integrationXref,omitempty"`
	LocationIds []int32 `json:"locationIds,omitempty"`
	DefaultFlag NullableBool `json:"defaultFlag,omitempty"`
	AddAllLocations NullableBool `json:"addAllLocations,omitempty"`
	RemoveAllLocations NullableBool `json:"removeAllLocations,omitempty"`
	Info *map[string]string `json:"_info,omitempty"`
}

type _Category Category

// NewCategory instantiates a new Category object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCategory(name string) *Category {
	this := Category{}
	this.Name = name
	return &this
}

// NewCategoryWithDefaults instantiates a new Category object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCategoryWithDefaults() *Category {
	this := Category{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Category) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Category) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Category) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Category) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *Category) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Category) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Category) SetName(v string) {
	o.Name = v
}

// GetInactiveFlag returns the InactiveFlag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Category) GetInactiveFlag() bool {
	if o == nil || IsNil(o.InactiveFlag.Get()) {
		var ret bool
		return ret
	}
	return *o.InactiveFlag.Get()
}

// GetInactiveFlagOk returns a tuple with the InactiveFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Category) GetInactiveFlagOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.InactiveFlag.Get(), o.InactiveFlag.IsSet()
}

// HasInactiveFlag returns a boolean if a field has been set.
func (o *Category) HasInactiveFlag() bool {
	if o != nil && o.InactiveFlag.IsSet() {
		return true
	}

	return false
}

// SetInactiveFlag gets a reference to the given NullableBool and assigns it to the InactiveFlag field.
func (o *Category) SetInactiveFlag(v bool) {
	o.InactiveFlag.Set(&v)
}
// SetInactiveFlagNil sets the value for InactiveFlag to be an explicit nil
func (o *Category) SetInactiveFlagNil() {
	o.InactiveFlag.Set(nil)
}

// UnsetInactiveFlag ensures that no value is present for InactiveFlag, not even an explicit nil
func (o *Category) UnsetInactiveFlag() {
	o.InactiveFlag.Unset()
}

// GetPriceLevelXref returns the PriceLevelXref field value if set, zero value otherwise.
func (o *Category) GetPriceLevelXref() string {
	if o == nil || IsNil(o.PriceLevelXref) {
		var ret string
		return ret
	}
	return *o.PriceLevelXref
}

// GetPriceLevelXrefOk returns a tuple with the PriceLevelXref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Category) GetPriceLevelXrefOk() (*string, bool) {
	if o == nil || IsNil(o.PriceLevelXref) {
		return nil, false
	}
	return o.PriceLevelXref, true
}

// HasPriceLevelXref returns a boolean if a field has been set.
func (o *Category) HasPriceLevelXref() bool {
	if o != nil && !IsNil(o.PriceLevelXref) {
		return true
	}

	return false
}

// SetPriceLevelXref gets a reference to the given string and assigns it to the PriceLevelXref field.
func (o *Category) SetPriceLevelXref(v string) {
	o.PriceLevelXref = &v
}

// GetIntegrationXref returns the IntegrationXref field value if set, zero value otherwise.
func (o *Category) GetIntegrationXref() string {
	if o == nil || IsNil(o.IntegrationXref) {
		var ret string
		return ret
	}
	return *o.IntegrationXref
}

// GetIntegrationXrefOk returns a tuple with the IntegrationXref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Category) GetIntegrationXrefOk() (*string, bool) {
	if o == nil || IsNil(o.IntegrationXref) {
		return nil, false
	}
	return o.IntegrationXref, true
}

// HasIntegrationXref returns a boolean if a field has been set.
func (o *Category) HasIntegrationXref() bool {
	if o != nil && !IsNil(o.IntegrationXref) {
		return true
	}

	return false
}

// SetIntegrationXref gets a reference to the given string and assigns it to the IntegrationXref field.
func (o *Category) SetIntegrationXref(v string) {
	o.IntegrationXref = &v
}

// GetLocationIds returns the LocationIds field value if set, zero value otherwise.
func (o *Category) GetLocationIds() []int32 {
	if o == nil || IsNil(o.LocationIds) {
		var ret []int32
		return ret
	}
	return o.LocationIds
}

// GetLocationIdsOk returns a tuple with the LocationIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Category) GetLocationIdsOk() ([]int32, bool) {
	if o == nil || IsNil(o.LocationIds) {
		return nil, false
	}
	return o.LocationIds, true
}

// HasLocationIds returns a boolean if a field has been set.
func (o *Category) HasLocationIds() bool {
	if o != nil && !IsNil(o.LocationIds) {
		return true
	}

	return false
}

// SetLocationIds gets a reference to the given []int32 and assigns it to the LocationIds field.
func (o *Category) SetLocationIds(v []int32) {
	o.LocationIds = v
}

// GetDefaultFlag returns the DefaultFlag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Category) GetDefaultFlag() bool {
	if o == nil || IsNil(o.DefaultFlag.Get()) {
		var ret bool
		return ret
	}
	return *o.DefaultFlag.Get()
}

// GetDefaultFlagOk returns a tuple with the DefaultFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Category) GetDefaultFlagOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultFlag.Get(), o.DefaultFlag.IsSet()
}

// HasDefaultFlag returns a boolean if a field has been set.
func (o *Category) HasDefaultFlag() bool {
	if o != nil && o.DefaultFlag.IsSet() {
		return true
	}

	return false
}

// SetDefaultFlag gets a reference to the given NullableBool and assigns it to the DefaultFlag field.
func (o *Category) SetDefaultFlag(v bool) {
	o.DefaultFlag.Set(&v)
}
// SetDefaultFlagNil sets the value for DefaultFlag to be an explicit nil
func (o *Category) SetDefaultFlagNil() {
	o.DefaultFlag.Set(nil)
}

// UnsetDefaultFlag ensures that no value is present for DefaultFlag, not even an explicit nil
func (o *Category) UnsetDefaultFlag() {
	o.DefaultFlag.Unset()
}

// GetAddAllLocations returns the AddAllLocations field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Category) GetAddAllLocations() bool {
	if o == nil || IsNil(o.AddAllLocations.Get()) {
		var ret bool
		return ret
	}
	return *o.AddAllLocations.Get()
}

// GetAddAllLocationsOk returns a tuple with the AddAllLocations field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Category) GetAddAllLocationsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AddAllLocations.Get(), o.AddAllLocations.IsSet()
}

// HasAddAllLocations returns a boolean if a field has been set.
func (o *Category) HasAddAllLocations() bool {
	if o != nil && o.AddAllLocations.IsSet() {
		return true
	}

	return false
}

// SetAddAllLocations gets a reference to the given NullableBool and assigns it to the AddAllLocations field.
func (o *Category) SetAddAllLocations(v bool) {
	o.AddAllLocations.Set(&v)
}
// SetAddAllLocationsNil sets the value for AddAllLocations to be an explicit nil
func (o *Category) SetAddAllLocationsNil() {
	o.AddAllLocations.Set(nil)
}

// UnsetAddAllLocations ensures that no value is present for AddAllLocations, not even an explicit nil
func (o *Category) UnsetAddAllLocations() {
	o.AddAllLocations.Unset()
}

// GetRemoveAllLocations returns the RemoveAllLocations field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Category) GetRemoveAllLocations() bool {
	if o == nil || IsNil(o.RemoveAllLocations.Get()) {
		var ret bool
		return ret
	}
	return *o.RemoveAllLocations.Get()
}

// GetRemoveAllLocationsOk returns a tuple with the RemoveAllLocations field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Category) GetRemoveAllLocationsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.RemoveAllLocations.Get(), o.RemoveAllLocations.IsSet()
}

// HasRemoveAllLocations returns a boolean if a field has been set.
func (o *Category) HasRemoveAllLocations() bool {
	if o != nil && o.RemoveAllLocations.IsSet() {
		return true
	}

	return false
}

// SetRemoveAllLocations gets a reference to the given NullableBool and assigns it to the RemoveAllLocations field.
func (o *Category) SetRemoveAllLocations(v bool) {
	o.RemoveAllLocations.Set(&v)
}
// SetRemoveAllLocationsNil sets the value for RemoveAllLocations to be an explicit nil
func (o *Category) SetRemoveAllLocationsNil() {
	o.RemoveAllLocations.Set(nil)
}

// UnsetRemoveAllLocations ensures that no value is present for RemoveAllLocations, not even an explicit nil
func (o *Category) UnsetRemoveAllLocations() {
	o.RemoveAllLocations.Unset()
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *Category) GetInfo() map[string]string {
	if o == nil || IsNil(o.Info) {
		var ret map[string]string
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Category) GetInfoOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Info) {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *Category) HasInfo() bool {
	if o != nil && !IsNil(o.Info) {
		return true
	}

	return false
}

// SetInfo gets a reference to the given map[string]string and assigns it to the Info field.
func (o *Category) SetInfo(v map[string]string) {
	o.Info = &v
}

func (o Category) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Category) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	if o.InactiveFlag.IsSet() {
		toSerialize["inactiveFlag"] = o.InactiveFlag.Get()
	}
	if !IsNil(o.PriceLevelXref) {
		toSerialize["priceLevelXref"] = o.PriceLevelXref
	}
	if !IsNil(o.IntegrationXref) {
		toSerialize["integrationXref"] = o.IntegrationXref
	}
	if !IsNil(o.LocationIds) {
		toSerialize["locationIds"] = o.LocationIds
	}
	if o.DefaultFlag.IsSet() {
		toSerialize["defaultFlag"] = o.DefaultFlag.Get()
	}
	if o.AddAllLocations.IsSet() {
		toSerialize["addAllLocations"] = o.AddAllLocations.Get()
	}
	if o.RemoveAllLocations.IsSet() {
		toSerialize["removeAllLocations"] = o.RemoveAllLocations.Get()
	}
	if !IsNil(o.Info) {
		toSerialize["_info"] = o.Info
	}
	return toSerialize, nil
}

func (o *Category) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varCategory := _Category{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCategory)

	if err != nil {
		return err
	}

	*o = Category(varCategory)

	return err
}

type NullableCategory struct {
	value *Category
	isSet bool
}

func (v NullableCategory) Get() *Category {
	return v.value
}

func (v *NullableCategory) Set(val *Category) {
	v.value = val
	v.isSet = true
}

func (v NullableCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCategory(val *Category) *NullableCategory {
	return &NullableCategory{value: val, isSet: true}
}

func (v NullableCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


