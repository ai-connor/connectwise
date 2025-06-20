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

// checks if the GLExportTransactionTaxLevel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GLExportTransactionTaxLevel{}

// GLExportTransactionTaxLevel struct for GLExportTransactionTaxLevel
type GLExportTransactionTaxLevel struct {
	TaxAmount NullableFloat64 `json:"taxAmount,omitempty"`
	TaxableAmount NullableFloat64 `json:"taxableAmount,omitempty"`
	TaxCodeXref *string `json:"taxCodeXref,omitempty"`
	TaxLevel *int32 `json:"taxLevel,omitempty"`
}

// NewGLExportTransactionTaxLevel instantiates a new GLExportTransactionTaxLevel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGLExportTransactionTaxLevel() *GLExportTransactionTaxLevel {
	this := GLExportTransactionTaxLevel{}
	return &this
}

// NewGLExportTransactionTaxLevelWithDefaults instantiates a new GLExportTransactionTaxLevel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGLExportTransactionTaxLevelWithDefaults() *GLExportTransactionTaxLevel {
	this := GLExportTransactionTaxLevel{}
	return &this
}

// GetTaxAmount returns the TaxAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GLExportTransactionTaxLevel) GetTaxAmount() float64 {
	if o == nil || IsNil(o.TaxAmount.Get()) {
		var ret float64
		return ret
	}
	return *o.TaxAmount.Get()
}

// GetTaxAmountOk returns a tuple with the TaxAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GLExportTransactionTaxLevel) GetTaxAmountOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.TaxAmount.Get(), o.TaxAmount.IsSet()
}

// HasTaxAmount returns a boolean if a field has been set.
func (o *GLExportTransactionTaxLevel) HasTaxAmount() bool {
	if o != nil && o.TaxAmount.IsSet() {
		return true
	}

	return false
}

// SetTaxAmount gets a reference to the given NullableFloat64 and assigns it to the TaxAmount field.
func (o *GLExportTransactionTaxLevel) SetTaxAmount(v float64) {
	o.TaxAmount.Set(&v)
}
// SetTaxAmountNil sets the value for TaxAmount to be an explicit nil
func (o *GLExportTransactionTaxLevel) SetTaxAmountNil() {
	o.TaxAmount.Set(nil)
}

// UnsetTaxAmount ensures that no value is present for TaxAmount, not even an explicit nil
func (o *GLExportTransactionTaxLevel) UnsetTaxAmount() {
	o.TaxAmount.Unset()
}

// GetTaxableAmount returns the TaxableAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GLExportTransactionTaxLevel) GetTaxableAmount() float64 {
	if o == nil || IsNil(o.TaxableAmount.Get()) {
		var ret float64
		return ret
	}
	return *o.TaxableAmount.Get()
}

// GetTaxableAmountOk returns a tuple with the TaxableAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GLExportTransactionTaxLevel) GetTaxableAmountOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.TaxableAmount.Get(), o.TaxableAmount.IsSet()
}

// HasTaxableAmount returns a boolean if a field has been set.
func (o *GLExportTransactionTaxLevel) HasTaxableAmount() bool {
	if o != nil && o.TaxableAmount.IsSet() {
		return true
	}

	return false
}

// SetTaxableAmount gets a reference to the given NullableFloat64 and assigns it to the TaxableAmount field.
func (o *GLExportTransactionTaxLevel) SetTaxableAmount(v float64) {
	o.TaxableAmount.Set(&v)
}
// SetTaxableAmountNil sets the value for TaxableAmount to be an explicit nil
func (o *GLExportTransactionTaxLevel) SetTaxableAmountNil() {
	o.TaxableAmount.Set(nil)
}

// UnsetTaxableAmount ensures that no value is present for TaxableAmount, not even an explicit nil
func (o *GLExportTransactionTaxLevel) UnsetTaxableAmount() {
	o.TaxableAmount.Unset()
}

// GetTaxCodeXref returns the TaxCodeXref field value if set, zero value otherwise.
func (o *GLExportTransactionTaxLevel) GetTaxCodeXref() string {
	if o == nil || IsNil(o.TaxCodeXref) {
		var ret string
		return ret
	}
	return *o.TaxCodeXref
}

// GetTaxCodeXrefOk returns a tuple with the TaxCodeXref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GLExportTransactionTaxLevel) GetTaxCodeXrefOk() (*string, bool) {
	if o == nil || IsNil(o.TaxCodeXref) {
		return nil, false
	}
	return o.TaxCodeXref, true
}

// HasTaxCodeXref returns a boolean if a field has been set.
func (o *GLExportTransactionTaxLevel) HasTaxCodeXref() bool {
	if o != nil && !IsNil(o.TaxCodeXref) {
		return true
	}

	return false
}

// SetTaxCodeXref gets a reference to the given string and assigns it to the TaxCodeXref field.
func (o *GLExportTransactionTaxLevel) SetTaxCodeXref(v string) {
	o.TaxCodeXref = &v
}

// GetTaxLevel returns the TaxLevel field value if set, zero value otherwise.
func (o *GLExportTransactionTaxLevel) GetTaxLevel() int32 {
	if o == nil || IsNil(o.TaxLevel) {
		var ret int32
		return ret
	}
	return *o.TaxLevel
}

// GetTaxLevelOk returns a tuple with the TaxLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GLExportTransactionTaxLevel) GetTaxLevelOk() (*int32, bool) {
	if o == nil || IsNil(o.TaxLevel) {
		return nil, false
	}
	return o.TaxLevel, true
}

// HasTaxLevel returns a boolean if a field has been set.
func (o *GLExportTransactionTaxLevel) HasTaxLevel() bool {
	if o != nil && !IsNil(o.TaxLevel) {
		return true
	}

	return false
}

// SetTaxLevel gets a reference to the given int32 and assigns it to the TaxLevel field.
func (o *GLExportTransactionTaxLevel) SetTaxLevel(v int32) {
	o.TaxLevel = &v
}

func (o GLExportTransactionTaxLevel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GLExportTransactionTaxLevel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.TaxAmount.IsSet() {
		toSerialize["taxAmount"] = o.TaxAmount.Get()
	}
	if o.TaxableAmount.IsSet() {
		toSerialize["taxableAmount"] = o.TaxableAmount.Get()
	}
	if !IsNil(o.TaxCodeXref) {
		toSerialize["taxCodeXref"] = o.TaxCodeXref
	}
	if !IsNil(o.TaxLevel) {
		toSerialize["taxLevel"] = o.TaxLevel
	}
	return toSerialize, nil
}

type NullableGLExportTransactionTaxLevel struct {
	value *GLExportTransactionTaxLevel
	isSet bool
}

func (v NullableGLExportTransactionTaxLevel) Get() *GLExportTransactionTaxLevel {
	return v.value
}

func (v *NullableGLExportTransactionTaxLevel) Set(val *GLExportTransactionTaxLevel) {
	v.value = val
	v.isSet = true
}

func (v NullableGLExportTransactionTaxLevel) IsSet() bool {
	return v.isSet
}

func (v *NullableGLExportTransactionTaxLevel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGLExportTransactionTaxLevel(val *GLExportTransactionTaxLevel) *NullableGLExportTransactionTaxLevel {
	return &NullableGLExportTransactionTaxLevel{value: val, isSet: true}
}

func (v NullableGLExportTransactionTaxLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGLExportTransactionTaxLevel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


