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

// checks if the GLExportExpenseBill type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GLExportExpenseBill{}

// GLExportExpenseBill struct for GLExportExpenseBill
type GLExportExpenseBill struct {
	Id NullableInt32 `json:"id,omitempty"`
	DocumentDate *string `json:"documentDate,omitempty"`
	DocumentType *string `json:"documentType,omitempty"`
	DocumentNumber *string `json:"documentNumber,omitempty"`
	Memo *string `json:"memo,omitempty"`
	GlClass *string `json:"glClass,omitempty"`
	ApAccountNumber *string `json:"apAccountNumber,omitempty"`
	Member *MemberReference `json:"member,omitempty"`
	VendorNumber *string `json:"vendorNumber,omitempty"`
	Currency *CurrencyReference `json:"currency,omitempty"`
	Total NullableFloat64 `json:"total,omitempty"`
	Detail []GLExportExpenseBillDetail `json:"detail,omitempty"`
}

// NewGLExportExpenseBill instantiates a new GLExportExpenseBill object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGLExportExpenseBill() *GLExportExpenseBill {
	this := GLExportExpenseBill{}
	return &this
}

// NewGLExportExpenseBillWithDefaults instantiates a new GLExportExpenseBill object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGLExportExpenseBillWithDefaults() *GLExportExpenseBill {
	this := GLExportExpenseBill{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GLExportExpenseBill) GetId() int32 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int32
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GLExportExpenseBill) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *GLExportExpenseBill) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt32 and assigns it to the Id field.
func (o *GLExportExpenseBill) SetId(v int32) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *GLExportExpenseBill) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *GLExportExpenseBill) UnsetId() {
	o.Id.Unset()
}

// GetDocumentDate returns the DocumentDate field value if set, zero value otherwise.
func (o *GLExportExpenseBill) GetDocumentDate() string {
	if o == nil || IsNil(o.DocumentDate) {
		var ret string
		return ret
	}
	return *o.DocumentDate
}

// GetDocumentDateOk returns a tuple with the DocumentDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GLExportExpenseBill) GetDocumentDateOk() (*string, bool) {
	if o == nil || IsNil(o.DocumentDate) {
		return nil, false
	}
	return o.DocumentDate, true
}

// HasDocumentDate returns a boolean if a field has been set.
func (o *GLExportExpenseBill) HasDocumentDate() bool {
	if o != nil && !IsNil(o.DocumentDate) {
		return true
	}

	return false
}

// SetDocumentDate gets a reference to the given string and assigns it to the DocumentDate field.
func (o *GLExportExpenseBill) SetDocumentDate(v string) {
	o.DocumentDate = &v
}

// GetDocumentType returns the DocumentType field value if set, zero value otherwise.
func (o *GLExportExpenseBill) GetDocumentType() string {
	if o == nil || IsNil(o.DocumentType) {
		var ret string
		return ret
	}
	return *o.DocumentType
}

// GetDocumentTypeOk returns a tuple with the DocumentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GLExportExpenseBill) GetDocumentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DocumentType) {
		return nil, false
	}
	return o.DocumentType, true
}

// HasDocumentType returns a boolean if a field has been set.
func (o *GLExportExpenseBill) HasDocumentType() bool {
	if o != nil && !IsNil(o.DocumentType) {
		return true
	}

	return false
}

// SetDocumentType gets a reference to the given string and assigns it to the DocumentType field.
func (o *GLExportExpenseBill) SetDocumentType(v string) {
	o.DocumentType = &v
}

// GetDocumentNumber returns the DocumentNumber field value if set, zero value otherwise.
func (o *GLExportExpenseBill) GetDocumentNumber() string {
	if o == nil || IsNil(o.DocumentNumber) {
		var ret string
		return ret
	}
	return *o.DocumentNumber
}

// GetDocumentNumberOk returns a tuple with the DocumentNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GLExportExpenseBill) GetDocumentNumberOk() (*string, bool) {
	if o == nil || IsNil(o.DocumentNumber) {
		return nil, false
	}
	return o.DocumentNumber, true
}

// HasDocumentNumber returns a boolean if a field has been set.
func (o *GLExportExpenseBill) HasDocumentNumber() bool {
	if o != nil && !IsNil(o.DocumentNumber) {
		return true
	}

	return false
}

// SetDocumentNumber gets a reference to the given string and assigns it to the DocumentNumber field.
func (o *GLExportExpenseBill) SetDocumentNumber(v string) {
	o.DocumentNumber = &v
}

// GetMemo returns the Memo field value if set, zero value otherwise.
func (o *GLExportExpenseBill) GetMemo() string {
	if o == nil || IsNil(o.Memo) {
		var ret string
		return ret
	}
	return *o.Memo
}

// GetMemoOk returns a tuple with the Memo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GLExportExpenseBill) GetMemoOk() (*string, bool) {
	if o == nil || IsNil(o.Memo) {
		return nil, false
	}
	return o.Memo, true
}

// HasMemo returns a boolean if a field has been set.
func (o *GLExportExpenseBill) HasMemo() bool {
	if o != nil && !IsNil(o.Memo) {
		return true
	}

	return false
}

// SetMemo gets a reference to the given string and assigns it to the Memo field.
func (o *GLExportExpenseBill) SetMemo(v string) {
	o.Memo = &v
}

// GetGlClass returns the GlClass field value if set, zero value otherwise.
func (o *GLExportExpenseBill) GetGlClass() string {
	if o == nil || IsNil(o.GlClass) {
		var ret string
		return ret
	}
	return *o.GlClass
}

// GetGlClassOk returns a tuple with the GlClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GLExportExpenseBill) GetGlClassOk() (*string, bool) {
	if o == nil || IsNil(o.GlClass) {
		return nil, false
	}
	return o.GlClass, true
}

// HasGlClass returns a boolean if a field has been set.
func (o *GLExportExpenseBill) HasGlClass() bool {
	if o != nil && !IsNil(o.GlClass) {
		return true
	}

	return false
}

// SetGlClass gets a reference to the given string and assigns it to the GlClass field.
func (o *GLExportExpenseBill) SetGlClass(v string) {
	o.GlClass = &v
}

// GetApAccountNumber returns the ApAccountNumber field value if set, zero value otherwise.
func (o *GLExportExpenseBill) GetApAccountNumber() string {
	if o == nil || IsNil(o.ApAccountNumber) {
		var ret string
		return ret
	}
	return *o.ApAccountNumber
}

// GetApAccountNumberOk returns a tuple with the ApAccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GLExportExpenseBill) GetApAccountNumberOk() (*string, bool) {
	if o == nil || IsNil(o.ApAccountNumber) {
		return nil, false
	}
	return o.ApAccountNumber, true
}

// HasApAccountNumber returns a boolean if a field has been set.
func (o *GLExportExpenseBill) HasApAccountNumber() bool {
	if o != nil && !IsNil(o.ApAccountNumber) {
		return true
	}

	return false
}

// SetApAccountNumber gets a reference to the given string and assigns it to the ApAccountNumber field.
func (o *GLExportExpenseBill) SetApAccountNumber(v string) {
	o.ApAccountNumber = &v
}

// GetMember returns the Member field value if set, zero value otherwise.
func (o *GLExportExpenseBill) GetMember() MemberReference {
	if o == nil || IsNil(o.Member) {
		var ret MemberReference
		return ret
	}
	return *o.Member
}

// GetMemberOk returns a tuple with the Member field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GLExportExpenseBill) GetMemberOk() (*MemberReference, bool) {
	if o == nil || IsNil(o.Member) {
		return nil, false
	}
	return o.Member, true
}

// HasMember returns a boolean if a field has been set.
func (o *GLExportExpenseBill) HasMember() bool {
	if o != nil && !IsNil(o.Member) {
		return true
	}

	return false
}

// SetMember gets a reference to the given MemberReference and assigns it to the Member field.
func (o *GLExportExpenseBill) SetMember(v MemberReference) {
	o.Member = &v
}

// GetVendorNumber returns the VendorNumber field value if set, zero value otherwise.
func (o *GLExportExpenseBill) GetVendorNumber() string {
	if o == nil || IsNil(o.VendorNumber) {
		var ret string
		return ret
	}
	return *o.VendorNumber
}

// GetVendorNumberOk returns a tuple with the VendorNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GLExportExpenseBill) GetVendorNumberOk() (*string, bool) {
	if o == nil || IsNil(o.VendorNumber) {
		return nil, false
	}
	return o.VendorNumber, true
}

// HasVendorNumber returns a boolean if a field has been set.
func (o *GLExportExpenseBill) HasVendorNumber() bool {
	if o != nil && !IsNil(o.VendorNumber) {
		return true
	}

	return false
}

// SetVendorNumber gets a reference to the given string and assigns it to the VendorNumber field.
func (o *GLExportExpenseBill) SetVendorNumber(v string) {
	o.VendorNumber = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *GLExportExpenseBill) GetCurrency() CurrencyReference {
	if o == nil || IsNil(o.Currency) {
		var ret CurrencyReference
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GLExportExpenseBill) GetCurrencyOk() (*CurrencyReference, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *GLExportExpenseBill) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given CurrencyReference and assigns it to the Currency field.
func (o *GLExportExpenseBill) SetCurrency(v CurrencyReference) {
	o.Currency = &v
}

// GetTotal returns the Total field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GLExportExpenseBill) GetTotal() float64 {
	if o == nil || IsNil(o.Total.Get()) {
		var ret float64
		return ret
	}
	return *o.Total.Get()
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GLExportExpenseBill) GetTotalOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Total.Get(), o.Total.IsSet()
}

// HasTotal returns a boolean if a field has been set.
func (o *GLExportExpenseBill) HasTotal() bool {
	if o != nil && o.Total.IsSet() {
		return true
	}

	return false
}

// SetTotal gets a reference to the given NullableFloat64 and assigns it to the Total field.
func (o *GLExportExpenseBill) SetTotal(v float64) {
	o.Total.Set(&v)
}
// SetTotalNil sets the value for Total to be an explicit nil
func (o *GLExportExpenseBill) SetTotalNil() {
	o.Total.Set(nil)
}

// UnsetTotal ensures that no value is present for Total, not even an explicit nil
func (o *GLExportExpenseBill) UnsetTotal() {
	o.Total.Unset()
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *GLExportExpenseBill) GetDetail() []GLExportExpenseBillDetail {
	if o == nil || IsNil(o.Detail) {
		var ret []GLExportExpenseBillDetail
		return ret
	}
	return o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GLExportExpenseBill) GetDetailOk() ([]GLExportExpenseBillDetail, bool) {
	if o == nil || IsNil(o.Detail) {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *GLExportExpenseBill) HasDetail() bool {
	if o != nil && !IsNil(o.Detail) {
		return true
	}

	return false
}

// SetDetail gets a reference to the given []GLExportExpenseBillDetail and assigns it to the Detail field.
func (o *GLExportExpenseBill) SetDetail(v []GLExportExpenseBillDetail) {
	o.Detail = v
}

func (o GLExportExpenseBill) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GLExportExpenseBill) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if !IsNil(o.DocumentDate) {
		toSerialize["documentDate"] = o.DocumentDate
	}
	if !IsNil(o.DocumentType) {
		toSerialize["documentType"] = o.DocumentType
	}
	if !IsNil(o.DocumentNumber) {
		toSerialize["documentNumber"] = o.DocumentNumber
	}
	if !IsNil(o.Memo) {
		toSerialize["memo"] = o.Memo
	}
	if !IsNil(o.GlClass) {
		toSerialize["glClass"] = o.GlClass
	}
	if !IsNil(o.ApAccountNumber) {
		toSerialize["apAccountNumber"] = o.ApAccountNumber
	}
	if !IsNil(o.Member) {
		toSerialize["member"] = o.Member
	}
	if !IsNil(o.VendorNumber) {
		toSerialize["vendorNumber"] = o.VendorNumber
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if o.Total.IsSet() {
		toSerialize["total"] = o.Total.Get()
	}
	if !IsNil(o.Detail) {
		toSerialize["detail"] = o.Detail
	}
	return toSerialize, nil
}

type NullableGLExportExpenseBill struct {
	value *GLExportExpenseBill
	isSet bool
}

func (v NullableGLExportExpenseBill) Get() *GLExportExpenseBill {
	return v.value
}

func (v *NullableGLExportExpenseBill) Set(val *GLExportExpenseBill) {
	v.value = val
	v.isSet = true
}

func (v NullableGLExportExpenseBill) IsSet() bool {
	return v.isSet
}

func (v *NullableGLExportExpenseBill) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGLExportExpenseBill(val *GLExportExpenseBill) *NullableGLExportExpenseBill {
	return &NullableGLExportExpenseBill{value: val, isSet: true}
}

func (v NullableGLExportExpenseBill) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGLExportExpenseBill) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


