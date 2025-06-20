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

// checks if the Warehouse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Warehouse{}

// Warehouse struct for Warehouse
type Warehouse struct {
	Id *int32 `json:"id,omitempty"`
	//  Max length: 50;
	Name string `json:"name"`
	Company *CompanyReference `json:"company,omitempty"`
	Location SystemLocationReference `json:"location"`
	Contact *ContactReference `json:"contact,omitempty"`
	Department SystemDepartmentReference `json:"department"`
	Manager *MemberReference `json:"manager,omitempty"`
	Site *SiteReference `json:"site,omitempty"`
	//  Max length: 10;
	LocationXref *string `json:"locationXref,omitempty"`
	LocationDefaultFlag NullableBool `json:"locationDefaultFlag,omitempty"`
	OverallDefaultFlag NullableBool `json:"overallDefaultFlag,omitempty"`
	InactiveFlag NullableBool `json:"inactiveFlag,omitempty"`
	LockedFlag NullableBool `json:"lockedFlag,omitempty"`
	Currency *CurrencyReference `json:"currency,omitempty"`
	Info *map[string]string `json:"_info,omitempty"`
}

type _Warehouse Warehouse

// NewWarehouse instantiates a new Warehouse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWarehouse(name string, location SystemLocationReference, department SystemDepartmentReference) *Warehouse {
	this := Warehouse{}
	this.Name = name
	this.Location = location
	this.Department = department
	return &this
}

// NewWarehouseWithDefaults instantiates a new Warehouse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWarehouseWithDefaults() *Warehouse {
	this := Warehouse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Warehouse) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Warehouse) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Warehouse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Warehouse) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *Warehouse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Warehouse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Warehouse) SetName(v string) {
	o.Name = v
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *Warehouse) GetCompany() CompanyReference {
	if o == nil || IsNil(o.Company) {
		var ret CompanyReference
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Warehouse) GetCompanyOk() (*CompanyReference, bool) {
	if o == nil || IsNil(o.Company) {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *Warehouse) HasCompany() bool {
	if o != nil && !IsNil(o.Company) {
		return true
	}

	return false
}

// SetCompany gets a reference to the given CompanyReference and assigns it to the Company field.
func (o *Warehouse) SetCompany(v CompanyReference) {
	o.Company = &v
}

// GetLocation returns the Location field value
func (o *Warehouse) GetLocation() SystemLocationReference {
	if o == nil {
		var ret SystemLocationReference
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *Warehouse) GetLocationOk() (*SystemLocationReference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *Warehouse) SetLocation(v SystemLocationReference) {
	o.Location = v
}

// GetContact returns the Contact field value if set, zero value otherwise.
func (o *Warehouse) GetContact() ContactReference {
	if o == nil || IsNil(o.Contact) {
		var ret ContactReference
		return ret
	}
	return *o.Contact
}

// GetContactOk returns a tuple with the Contact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Warehouse) GetContactOk() (*ContactReference, bool) {
	if o == nil || IsNil(o.Contact) {
		return nil, false
	}
	return o.Contact, true
}

// HasContact returns a boolean if a field has been set.
func (o *Warehouse) HasContact() bool {
	if o != nil && !IsNil(o.Contact) {
		return true
	}

	return false
}

// SetContact gets a reference to the given ContactReference and assigns it to the Contact field.
func (o *Warehouse) SetContact(v ContactReference) {
	o.Contact = &v
}

// GetDepartment returns the Department field value
func (o *Warehouse) GetDepartment() SystemDepartmentReference {
	if o == nil {
		var ret SystemDepartmentReference
		return ret
	}

	return o.Department
}

// GetDepartmentOk returns a tuple with the Department field value
// and a boolean to check if the value has been set.
func (o *Warehouse) GetDepartmentOk() (*SystemDepartmentReference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Department, true
}

// SetDepartment sets field value
func (o *Warehouse) SetDepartment(v SystemDepartmentReference) {
	o.Department = v
}

// GetManager returns the Manager field value if set, zero value otherwise.
func (o *Warehouse) GetManager() MemberReference {
	if o == nil || IsNil(o.Manager) {
		var ret MemberReference
		return ret
	}
	return *o.Manager
}

// GetManagerOk returns a tuple with the Manager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Warehouse) GetManagerOk() (*MemberReference, bool) {
	if o == nil || IsNil(o.Manager) {
		return nil, false
	}
	return o.Manager, true
}

// HasManager returns a boolean if a field has been set.
func (o *Warehouse) HasManager() bool {
	if o != nil && !IsNil(o.Manager) {
		return true
	}

	return false
}

// SetManager gets a reference to the given MemberReference and assigns it to the Manager field.
func (o *Warehouse) SetManager(v MemberReference) {
	o.Manager = &v
}

// GetSite returns the Site field value if set, zero value otherwise.
func (o *Warehouse) GetSite() SiteReference {
	if o == nil || IsNil(o.Site) {
		var ret SiteReference
		return ret
	}
	return *o.Site
}

// GetSiteOk returns a tuple with the Site field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Warehouse) GetSiteOk() (*SiteReference, bool) {
	if o == nil || IsNil(o.Site) {
		return nil, false
	}
	return o.Site, true
}

// HasSite returns a boolean if a field has been set.
func (o *Warehouse) HasSite() bool {
	if o != nil && !IsNil(o.Site) {
		return true
	}

	return false
}

// SetSite gets a reference to the given SiteReference and assigns it to the Site field.
func (o *Warehouse) SetSite(v SiteReference) {
	o.Site = &v
}

// GetLocationXref returns the LocationXref field value if set, zero value otherwise.
func (o *Warehouse) GetLocationXref() string {
	if o == nil || IsNil(o.LocationXref) {
		var ret string
		return ret
	}
	return *o.LocationXref
}

// GetLocationXrefOk returns a tuple with the LocationXref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Warehouse) GetLocationXrefOk() (*string, bool) {
	if o == nil || IsNil(o.LocationXref) {
		return nil, false
	}
	return o.LocationXref, true
}

// HasLocationXref returns a boolean if a field has been set.
func (o *Warehouse) HasLocationXref() bool {
	if o != nil && !IsNil(o.LocationXref) {
		return true
	}

	return false
}

// SetLocationXref gets a reference to the given string and assigns it to the LocationXref field.
func (o *Warehouse) SetLocationXref(v string) {
	o.LocationXref = &v
}

// GetLocationDefaultFlag returns the LocationDefaultFlag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Warehouse) GetLocationDefaultFlag() bool {
	if o == nil || IsNil(o.LocationDefaultFlag.Get()) {
		var ret bool
		return ret
	}
	return *o.LocationDefaultFlag.Get()
}

// GetLocationDefaultFlagOk returns a tuple with the LocationDefaultFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Warehouse) GetLocationDefaultFlagOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.LocationDefaultFlag.Get(), o.LocationDefaultFlag.IsSet()
}

// HasLocationDefaultFlag returns a boolean if a field has been set.
func (o *Warehouse) HasLocationDefaultFlag() bool {
	if o != nil && o.LocationDefaultFlag.IsSet() {
		return true
	}

	return false
}

// SetLocationDefaultFlag gets a reference to the given NullableBool and assigns it to the LocationDefaultFlag field.
func (o *Warehouse) SetLocationDefaultFlag(v bool) {
	o.LocationDefaultFlag.Set(&v)
}
// SetLocationDefaultFlagNil sets the value for LocationDefaultFlag to be an explicit nil
func (o *Warehouse) SetLocationDefaultFlagNil() {
	o.LocationDefaultFlag.Set(nil)
}

// UnsetLocationDefaultFlag ensures that no value is present for LocationDefaultFlag, not even an explicit nil
func (o *Warehouse) UnsetLocationDefaultFlag() {
	o.LocationDefaultFlag.Unset()
}

// GetOverallDefaultFlag returns the OverallDefaultFlag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Warehouse) GetOverallDefaultFlag() bool {
	if o == nil || IsNil(o.OverallDefaultFlag.Get()) {
		var ret bool
		return ret
	}
	return *o.OverallDefaultFlag.Get()
}

// GetOverallDefaultFlagOk returns a tuple with the OverallDefaultFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Warehouse) GetOverallDefaultFlagOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.OverallDefaultFlag.Get(), o.OverallDefaultFlag.IsSet()
}

// HasOverallDefaultFlag returns a boolean if a field has been set.
func (o *Warehouse) HasOverallDefaultFlag() bool {
	if o != nil && o.OverallDefaultFlag.IsSet() {
		return true
	}

	return false
}

// SetOverallDefaultFlag gets a reference to the given NullableBool and assigns it to the OverallDefaultFlag field.
func (o *Warehouse) SetOverallDefaultFlag(v bool) {
	o.OverallDefaultFlag.Set(&v)
}
// SetOverallDefaultFlagNil sets the value for OverallDefaultFlag to be an explicit nil
func (o *Warehouse) SetOverallDefaultFlagNil() {
	o.OverallDefaultFlag.Set(nil)
}

// UnsetOverallDefaultFlag ensures that no value is present for OverallDefaultFlag, not even an explicit nil
func (o *Warehouse) UnsetOverallDefaultFlag() {
	o.OverallDefaultFlag.Unset()
}

// GetInactiveFlag returns the InactiveFlag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Warehouse) GetInactiveFlag() bool {
	if o == nil || IsNil(o.InactiveFlag.Get()) {
		var ret bool
		return ret
	}
	return *o.InactiveFlag.Get()
}

// GetInactiveFlagOk returns a tuple with the InactiveFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Warehouse) GetInactiveFlagOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.InactiveFlag.Get(), o.InactiveFlag.IsSet()
}

// HasInactiveFlag returns a boolean if a field has been set.
func (o *Warehouse) HasInactiveFlag() bool {
	if o != nil && o.InactiveFlag.IsSet() {
		return true
	}

	return false
}

// SetInactiveFlag gets a reference to the given NullableBool and assigns it to the InactiveFlag field.
func (o *Warehouse) SetInactiveFlag(v bool) {
	o.InactiveFlag.Set(&v)
}
// SetInactiveFlagNil sets the value for InactiveFlag to be an explicit nil
func (o *Warehouse) SetInactiveFlagNil() {
	o.InactiveFlag.Set(nil)
}

// UnsetInactiveFlag ensures that no value is present for InactiveFlag, not even an explicit nil
func (o *Warehouse) UnsetInactiveFlag() {
	o.InactiveFlag.Unset()
}

// GetLockedFlag returns the LockedFlag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Warehouse) GetLockedFlag() bool {
	if o == nil || IsNil(o.LockedFlag.Get()) {
		var ret bool
		return ret
	}
	return *o.LockedFlag.Get()
}

// GetLockedFlagOk returns a tuple with the LockedFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Warehouse) GetLockedFlagOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.LockedFlag.Get(), o.LockedFlag.IsSet()
}

// HasLockedFlag returns a boolean if a field has been set.
func (o *Warehouse) HasLockedFlag() bool {
	if o != nil && o.LockedFlag.IsSet() {
		return true
	}

	return false
}

// SetLockedFlag gets a reference to the given NullableBool and assigns it to the LockedFlag field.
func (o *Warehouse) SetLockedFlag(v bool) {
	o.LockedFlag.Set(&v)
}
// SetLockedFlagNil sets the value for LockedFlag to be an explicit nil
func (o *Warehouse) SetLockedFlagNil() {
	o.LockedFlag.Set(nil)
}

// UnsetLockedFlag ensures that no value is present for LockedFlag, not even an explicit nil
func (o *Warehouse) UnsetLockedFlag() {
	o.LockedFlag.Unset()
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *Warehouse) GetCurrency() CurrencyReference {
	if o == nil || IsNil(o.Currency) {
		var ret CurrencyReference
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Warehouse) GetCurrencyOk() (*CurrencyReference, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *Warehouse) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given CurrencyReference and assigns it to the Currency field.
func (o *Warehouse) SetCurrency(v CurrencyReference) {
	o.Currency = &v
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *Warehouse) GetInfo() map[string]string {
	if o == nil || IsNil(o.Info) {
		var ret map[string]string
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Warehouse) GetInfoOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Info) {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *Warehouse) HasInfo() bool {
	if o != nil && !IsNil(o.Info) {
		return true
	}

	return false
}

// SetInfo gets a reference to the given map[string]string and assigns it to the Info field.
func (o *Warehouse) SetInfo(v map[string]string) {
	o.Info = &v
}

func (o Warehouse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Warehouse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Company) {
		toSerialize["company"] = o.Company
	}
	toSerialize["location"] = o.Location
	if !IsNil(o.Contact) {
		toSerialize["contact"] = o.Contact
	}
	toSerialize["department"] = o.Department
	if !IsNil(o.Manager) {
		toSerialize["manager"] = o.Manager
	}
	if !IsNil(o.Site) {
		toSerialize["site"] = o.Site
	}
	if !IsNil(o.LocationXref) {
		toSerialize["locationXref"] = o.LocationXref
	}
	if o.LocationDefaultFlag.IsSet() {
		toSerialize["locationDefaultFlag"] = o.LocationDefaultFlag.Get()
	}
	if o.OverallDefaultFlag.IsSet() {
		toSerialize["overallDefaultFlag"] = o.OverallDefaultFlag.Get()
	}
	if o.InactiveFlag.IsSet() {
		toSerialize["inactiveFlag"] = o.InactiveFlag.Get()
	}
	if o.LockedFlag.IsSet() {
		toSerialize["lockedFlag"] = o.LockedFlag.Get()
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.Info) {
		toSerialize["_info"] = o.Info
	}
	return toSerialize, nil
}

func (o *Warehouse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"location",
		"department",
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

	varWarehouse := _Warehouse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWarehouse)

	if err != nil {
		return err
	}

	*o = Warehouse(varWarehouse)

	return err
}

type NullableWarehouse struct {
	value *Warehouse
	isSet bool
}

func (v NullableWarehouse) Get() *Warehouse {
	return v.value
}

func (v *NullableWarehouse) Set(val *Warehouse) {
	v.value = val
	v.isSet = true
}

func (v NullableWarehouse) IsSet() bool {
	return v.isSet
}

func (v *NullableWarehouse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWarehouse(val *Warehouse) *NullableWarehouse {
	return &NullableWarehouse{value: val, isSet: true}
}

func (v NullableWarehouse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWarehouse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


