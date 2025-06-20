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

// checks if the PricingDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PricingDetail{}

// PricingDetail struct for PricingDetail
type PricingDetail struct {
	Id *int32 `json:"id,omitempty"`
	Product *CatalogItemReference `json:"product,omitempty"`
	Category *ProductCategoryReference `json:"category,omitempty"`
	SubCategory *ProductSubCategoryReference `json:"subCategory,omitempty"`
	StartDate time.Time `json:"startDate"`
	EndDate *time.Time `json:"endDate,omitempty"`
	NoEndDate *bool `json:"noEndDate,omitempty"`
	Info *map[string]string `json:"_info,omitempty"`
}

type _PricingDetail PricingDetail

// NewPricingDetail instantiates a new PricingDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPricingDetail(startDate time.Time) *PricingDetail {
	this := PricingDetail{}
	this.StartDate = startDate
	return &this
}

// NewPricingDetailWithDefaults instantiates a new PricingDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPricingDetailWithDefaults() *PricingDetail {
	this := PricingDetail{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PricingDetail) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricingDetail) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PricingDetail) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *PricingDetail) SetId(v int32) {
	o.Id = &v
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *PricingDetail) GetProduct() CatalogItemReference {
	if o == nil || IsNil(o.Product) {
		var ret CatalogItemReference
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricingDetail) GetProductOk() (*CatalogItemReference, bool) {
	if o == nil || IsNil(o.Product) {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *PricingDetail) HasProduct() bool {
	if o != nil && !IsNil(o.Product) {
		return true
	}

	return false
}

// SetProduct gets a reference to the given CatalogItemReference and assigns it to the Product field.
func (o *PricingDetail) SetProduct(v CatalogItemReference) {
	o.Product = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *PricingDetail) GetCategory() ProductCategoryReference {
	if o == nil || IsNil(o.Category) {
		var ret ProductCategoryReference
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricingDetail) GetCategoryOk() (*ProductCategoryReference, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *PricingDetail) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given ProductCategoryReference and assigns it to the Category field.
func (o *PricingDetail) SetCategory(v ProductCategoryReference) {
	o.Category = &v
}

// GetSubCategory returns the SubCategory field value if set, zero value otherwise.
func (o *PricingDetail) GetSubCategory() ProductSubCategoryReference {
	if o == nil || IsNil(o.SubCategory) {
		var ret ProductSubCategoryReference
		return ret
	}
	return *o.SubCategory
}

// GetSubCategoryOk returns a tuple with the SubCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricingDetail) GetSubCategoryOk() (*ProductSubCategoryReference, bool) {
	if o == nil || IsNil(o.SubCategory) {
		return nil, false
	}
	return o.SubCategory, true
}

// HasSubCategory returns a boolean if a field has been set.
func (o *PricingDetail) HasSubCategory() bool {
	if o != nil && !IsNil(o.SubCategory) {
		return true
	}

	return false
}

// SetSubCategory gets a reference to the given ProductSubCategoryReference and assigns it to the SubCategory field.
func (o *PricingDetail) SetSubCategory(v ProductSubCategoryReference) {
	o.SubCategory = &v
}

// GetStartDate returns the StartDate field value
func (o *PricingDetail) GetStartDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value
// and a boolean to check if the value has been set.
func (o *PricingDetail) GetStartDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartDate, true
}

// SetStartDate sets field value
func (o *PricingDetail) SetStartDate(v time.Time) {
	o.StartDate = v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *PricingDetail) GetEndDate() time.Time {
	if o == nil || IsNil(o.EndDate) {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricingDetail) GetEndDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *PricingDetail) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *PricingDetail) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetNoEndDate returns the NoEndDate field value if set, zero value otherwise.
func (o *PricingDetail) GetNoEndDate() bool {
	if o == nil || IsNil(o.NoEndDate) {
		var ret bool
		return ret
	}
	return *o.NoEndDate
}

// GetNoEndDateOk returns a tuple with the NoEndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricingDetail) GetNoEndDateOk() (*bool, bool) {
	if o == nil || IsNil(o.NoEndDate) {
		return nil, false
	}
	return o.NoEndDate, true
}

// HasNoEndDate returns a boolean if a field has been set.
func (o *PricingDetail) HasNoEndDate() bool {
	if o != nil && !IsNil(o.NoEndDate) {
		return true
	}

	return false
}

// SetNoEndDate gets a reference to the given bool and assigns it to the NoEndDate field.
func (o *PricingDetail) SetNoEndDate(v bool) {
	o.NoEndDate = &v
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *PricingDetail) GetInfo() map[string]string {
	if o == nil || IsNil(o.Info) {
		var ret map[string]string
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricingDetail) GetInfoOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Info) {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *PricingDetail) HasInfo() bool {
	if o != nil && !IsNil(o.Info) {
		return true
	}

	return false
}

// SetInfo gets a reference to the given map[string]string and assigns it to the Info field.
func (o *PricingDetail) SetInfo(v map[string]string) {
	o.Info = &v
}

func (o PricingDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PricingDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Product) {
		toSerialize["product"] = o.Product
	}
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !IsNil(o.SubCategory) {
		toSerialize["subCategory"] = o.SubCategory
	}
	toSerialize["startDate"] = o.StartDate
	if !IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	if !IsNil(o.NoEndDate) {
		toSerialize["noEndDate"] = o.NoEndDate
	}
	if !IsNil(o.Info) {
		toSerialize["_info"] = o.Info
	}
	return toSerialize, nil
}

func (o *PricingDetail) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"startDate",
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

	varPricingDetail := _PricingDetail{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPricingDetail)

	if err != nil {
		return err
	}

	*o = PricingDetail(varPricingDetail)

	return err
}

type NullablePricingDetail struct {
	value *PricingDetail
	isSet bool
}

func (v NullablePricingDetail) Get() *PricingDetail {
	return v.value
}

func (v *NullablePricingDetail) Set(val *PricingDetail) {
	v.value = val
	v.isSet = true
}

func (v NullablePricingDetail) IsSet() bool {
	return v.isSet
}

func (v *NullablePricingDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePricingDetail(val *PricingDetail) *NullablePricingDetail {
	return &NullablePricingDetail{value: val, isSet: true}
}

func (v NullablePricingDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePricingDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


