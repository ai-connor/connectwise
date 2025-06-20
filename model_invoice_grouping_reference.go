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

// checks if the InvoiceGroupingReference type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvoiceGroupingReference{}

// InvoiceGroupingReference struct for InvoiceGroupingReference
type InvoiceGroupingReference struct {
	Id NullableInt32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	ShowPriceFlag *bool `json:"showPriceFlag,omitempty"`
	ShowSubItemsFlag *bool `json:"showSubItemsFlag,omitempty"`
	GroupParentChildAdditionsFlag *bool `json:"groupParentChildAdditionsFlag,omitempty"`
	Info *map[string]string `json:"_info,omitempty"`
}

// NewInvoiceGroupingReference instantiates a new InvoiceGroupingReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoiceGroupingReference() *InvoiceGroupingReference {
	this := InvoiceGroupingReference{}
	return &this
}

// NewInvoiceGroupingReferenceWithDefaults instantiates a new InvoiceGroupingReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceGroupingReferenceWithDefaults() *InvoiceGroupingReference {
	this := InvoiceGroupingReference{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvoiceGroupingReference) GetId() int32 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int32
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceGroupingReference) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *InvoiceGroupingReference) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt32 and assigns it to the Id field.
func (o *InvoiceGroupingReference) SetId(v int32) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *InvoiceGroupingReference) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *InvoiceGroupingReference) UnsetId() {
	o.Id.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InvoiceGroupingReference) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceGroupingReference) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InvoiceGroupingReference) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InvoiceGroupingReference) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InvoiceGroupingReference) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceGroupingReference) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InvoiceGroupingReference) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InvoiceGroupingReference) SetDescription(v string) {
	o.Description = &v
}

// GetShowPriceFlag returns the ShowPriceFlag field value if set, zero value otherwise.
func (o *InvoiceGroupingReference) GetShowPriceFlag() bool {
	if o == nil || IsNil(o.ShowPriceFlag) {
		var ret bool
		return ret
	}
	return *o.ShowPriceFlag
}

// GetShowPriceFlagOk returns a tuple with the ShowPriceFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceGroupingReference) GetShowPriceFlagOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowPriceFlag) {
		return nil, false
	}
	return o.ShowPriceFlag, true
}

// HasShowPriceFlag returns a boolean if a field has been set.
func (o *InvoiceGroupingReference) HasShowPriceFlag() bool {
	if o != nil && !IsNil(o.ShowPriceFlag) {
		return true
	}

	return false
}

// SetShowPriceFlag gets a reference to the given bool and assigns it to the ShowPriceFlag field.
func (o *InvoiceGroupingReference) SetShowPriceFlag(v bool) {
	o.ShowPriceFlag = &v
}

// GetShowSubItemsFlag returns the ShowSubItemsFlag field value if set, zero value otherwise.
func (o *InvoiceGroupingReference) GetShowSubItemsFlag() bool {
	if o == nil || IsNil(o.ShowSubItemsFlag) {
		var ret bool
		return ret
	}
	return *o.ShowSubItemsFlag
}

// GetShowSubItemsFlagOk returns a tuple with the ShowSubItemsFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceGroupingReference) GetShowSubItemsFlagOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowSubItemsFlag) {
		return nil, false
	}
	return o.ShowSubItemsFlag, true
}

// HasShowSubItemsFlag returns a boolean if a field has been set.
func (o *InvoiceGroupingReference) HasShowSubItemsFlag() bool {
	if o != nil && !IsNil(o.ShowSubItemsFlag) {
		return true
	}

	return false
}

// SetShowSubItemsFlag gets a reference to the given bool and assigns it to the ShowSubItemsFlag field.
func (o *InvoiceGroupingReference) SetShowSubItemsFlag(v bool) {
	o.ShowSubItemsFlag = &v
}

// GetGroupParentChildAdditionsFlag returns the GroupParentChildAdditionsFlag field value if set, zero value otherwise.
func (o *InvoiceGroupingReference) GetGroupParentChildAdditionsFlag() bool {
	if o == nil || IsNil(o.GroupParentChildAdditionsFlag) {
		var ret bool
		return ret
	}
	return *o.GroupParentChildAdditionsFlag
}

// GetGroupParentChildAdditionsFlagOk returns a tuple with the GroupParentChildAdditionsFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceGroupingReference) GetGroupParentChildAdditionsFlagOk() (*bool, bool) {
	if o == nil || IsNil(o.GroupParentChildAdditionsFlag) {
		return nil, false
	}
	return o.GroupParentChildAdditionsFlag, true
}

// HasGroupParentChildAdditionsFlag returns a boolean if a field has been set.
func (o *InvoiceGroupingReference) HasGroupParentChildAdditionsFlag() bool {
	if o != nil && !IsNil(o.GroupParentChildAdditionsFlag) {
		return true
	}

	return false
}

// SetGroupParentChildAdditionsFlag gets a reference to the given bool and assigns it to the GroupParentChildAdditionsFlag field.
func (o *InvoiceGroupingReference) SetGroupParentChildAdditionsFlag(v bool) {
	o.GroupParentChildAdditionsFlag = &v
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *InvoiceGroupingReference) GetInfo() map[string]string {
	if o == nil || IsNil(o.Info) {
		var ret map[string]string
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceGroupingReference) GetInfoOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Info) {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *InvoiceGroupingReference) HasInfo() bool {
	if o != nil && !IsNil(o.Info) {
		return true
	}

	return false
}

// SetInfo gets a reference to the given map[string]string and assigns it to the Info field.
func (o *InvoiceGroupingReference) SetInfo(v map[string]string) {
	o.Info = &v
}

func (o InvoiceGroupingReference) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvoiceGroupingReference) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ShowPriceFlag) {
		toSerialize["showPriceFlag"] = o.ShowPriceFlag
	}
	if !IsNil(o.ShowSubItemsFlag) {
		toSerialize["showSubItemsFlag"] = o.ShowSubItemsFlag
	}
	if !IsNil(o.GroupParentChildAdditionsFlag) {
		toSerialize["groupParentChildAdditionsFlag"] = o.GroupParentChildAdditionsFlag
	}
	if !IsNil(o.Info) {
		toSerialize["_info"] = o.Info
	}
	return toSerialize, nil
}

type NullableInvoiceGroupingReference struct {
	value *InvoiceGroupingReference
	isSet bool
}

func (v NullableInvoiceGroupingReference) Get() *InvoiceGroupingReference {
	return v.value
}

func (v *NullableInvoiceGroupingReference) Set(val *InvoiceGroupingReference) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceGroupingReference) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceGroupingReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceGroupingReference(val *InvoiceGroupingReference) *NullableInvoiceGroupingReference {
	return &NullableInvoiceGroupingReference{value: val, isSet: true}
}

func (v NullableInvoiceGroupingReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceGroupingReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


