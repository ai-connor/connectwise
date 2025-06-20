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

// checks if the GLExportInventoryTransferOffset type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GLExportInventoryTransferOffset{}

// GLExportInventoryTransferOffset struct for GLExportInventoryTransferOffset
type GLExportInventoryTransferOffset struct {
	Id NullableInt32 `json:"id,omitempty"`
	DocumentType *string `json:"documentType,omitempty"`
	DocumentDate *string `json:"documentDate,omitempty"`
	AccountNumber *string `json:"accountNumber,omitempty"`
	GlClass *string `json:"glClass,omitempty"`
	Total NullableFloat64 `json:"total,omitempty"`
	Memo *string `json:"memo,omitempty"`
	Description *string `json:"description,omitempty"`
	GlTypeId *string `json:"glTypeId,omitempty"`
}

// NewGLExportInventoryTransferOffset instantiates a new GLExportInventoryTransferOffset object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGLExportInventoryTransferOffset() *GLExportInventoryTransferOffset {
	this := GLExportInventoryTransferOffset{}
	return &this
}

// NewGLExportInventoryTransferOffsetWithDefaults instantiates a new GLExportInventoryTransferOffset object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGLExportInventoryTransferOffsetWithDefaults() *GLExportInventoryTransferOffset {
	this := GLExportInventoryTransferOffset{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GLExportInventoryTransferOffset) GetId() int32 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int32
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GLExportInventoryTransferOffset) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *GLExportInventoryTransferOffset) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt32 and assigns it to the Id field.
func (o *GLExportInventoryTransferOffset) SetId(v int32) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *GLExportInventoryTransferOffset) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *GLExportInventoryTransferOffset) UnsetId() {
	o.Id.Unset()
}

// GetDocumentType returns the DocumentType field value if set, zero value otherwise.
func (o *GLExportInventoryTransferOffset) GetDocumentType() string {
	if o == nil || IsNil(o.DocumentType) {
		var ret string
		return ret
	}
	return *o.DocumentType
}

// GetDocumentTypeOk returns a tuple with the DocumentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GLExportInventoryTransferOffset) GetDocumentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DocumentType) {
		return nil, false
	}
	return o.DocumentType, true
}

// HasDocumentType returns a boolean if a field has been set.
func (o *GLExportInventoryTransferOffset) HasDocumentType() bool {
	if o != nil && !IsNil(o.DocumentType) {
		return true
	}

	return false
}

// SetDocumentType gets a reference to the given string and assigns it to the DocumentType field.
func (o *GLExportInventoryTransferOffset) SetDocumentType(v string) {
	o.DocumentType = &v
}

// GetDocumentDate returns the DocumentDate field value if set, zero value otherwise.
func (o *GLExportInventoryTransferOffset) GetDocumentDate() string {
	if o == nil || IsNil(o.DocumentDate) {
		var ret string
		return ret
	}
	return *o.DocumentDate
}

// GetDocumentDateOk returns a tuple with the DocumentDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GLExportInventoryTransferOffset) GetDocumentDateOk() (*string, bool) {
	if o == nil || IsNil(o.DocumentDate) {
		return nil, false
	}
	return o.DocumentDate, true
}

// HasDocumentDate returns a boolean if a field has been set.
func (o *GLExportInventoryTransferOffset) HasDocumentDate() bool {
	if o != nil && !IsNil(o.DocumentDate) {
		return true
	}

	return false
}

// SetDocumentDate gets a reference to the given string and assigns it to the DocumentDate field.
func (o *GLExportInventoryTransferOffset) SetDocumentDate(v string) {
	o.DocumentDate = &v
}

// GetAccountNumber returns the AccountNumber field value if set, zero value otherwise.
func (o *GLExportInventoryTransferOffset) GetAccountNumber() string {
	if o == nil || IsNil(o.AccountNumber) {
		var ret string
		return ret
	}
	return *o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GLExportInventoryTransferOffset) GetAccountNumberOk() (*string, bool) {
	if o == nil || IsNil(o.AccountNumber) {
		return nil, false
	}
	return o.AccountNumber, true
}

// HasAccountNumber returns a boolean if a field has been set.
func (o *GLExportInventoryTransferOffset) HasAccountNumber() bool {
	if o != nil && !IsNil(o.AccountNumber) {
		return true
	}

	return false
}

// SetAccountNumber gets a reference to the given string and assigns it to the AccountNumber field.
func (o *GLExportInventoryTransferOffset) SetAccountNumber(v string) {
	o.AccountNumber = &v
}

// GetGlClass returns the GlClass field value if set, zero value otherwise.
func (o *GLExportInventoryTransferOffset) GetGlClass() string {
	if o == nil || IsNil(o.GlClass) {
		var ret string
		return ret
	}
	return *o.GlClass
}

// GetGlClassOk returns a tuple with the GlClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GLExportInventoryTransferOffset) GetGlClassOk() (*string, bool) {
	if o == nil || IsNil(o.GlClass) {
		return nil, false
	}
	return o.GlClass, true
}

// HasGlClass returns a boolean if a field has been set.
func (o *GLExportInventoryTransferOffset) HasGlClass() bool {
	if o != nil && !IsNil(o.GlClass) {
		return true
	}

	return false
}

// SetGlClass gets a reference to the given string and assigns it to the GlClass field.
func (o *GLExportInventoryTransferOffset) SetGlClass(v string) {
	o.GlClass = &v
}

// GetTotal returns the Total field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GLExportInventoryTransferOffset) GetTotal() float64 {
	if o == nil || IsNil(o.Total.Get()) {
		var ret float64
		return ret
	}
	return *o.Total.Get()
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GLExportInventoryTransferOffset) GetTotalOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Total.Get(), o.Total.IsSet()
}

// HasTotal returns a boolean if a field has been set.
func (o *GLExportInventoryTransferOffset) HasTotal() bool {
	if o != nil && o.Total.IsSet() {
		return true
	}

	return false
}

// SetTotal gets a reference to the given NullableFloat64 and assigns it to the Total field.
func (o *GLExportInventoryTransferOffset) SetTotal(v float64) {
	o.Total.Set(&v)
}
// SetTotalNil sets the value for Total to be an explicit nil
func (o *GLExportInventoryTransferOffset) SetTotalNil() {
	o.Total.Set(nil)
}

// UnsetTotal ensures that no value is present for Total, not even an explicit nil
func (o *GLExportInventoryTransferOffset) UnsetTotal() {
	o.Total.Unset()
}

// GetMemo returns the Memo field value if set, zero value otherwise.
func (o *GLExportInventoryTransferOffset) GetMemo() string {
	if o == nil || IsNil(o.Memo) {
		var ret string
		return ret
	}
	return *o.Memo
}

// GetMemoOk returns a tuple with the Memo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GLExportInventoryTransferOffset) GetMemoOk() (*string, bool) {
	if o == nil || IsNil(o.Memo) {
		return nil, false
	}
	return o.Memo, true
}

// HasMemo returns a boolean if a field has been set.
func (o *GLExportInventoryTransferOffset) HasMemo() bool {
	if o != nil && !IsNil(o.Memo) {
		return true
	}

	return false
}

// SetMemo gets a reference to the given string and assigns it to the Memo field.
func (o *GLExportInventoryTransferOffset) SetMemo(v string) {
	o.Memo = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GLExportInventoryTransferOffset) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GLExportInventoryTransferOffset) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GLExportInventoryTransferOffset) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GLExportInventoryTransferOffset) SetDescription(v string) {
	o.Description = &v
}

// GetGlTypeId returns the GlTypeId field value if set, zero value otherwise.
func (o *GLExportInventoryTransferOffset) GetGlTypeId() string {
	if o == nil || IsNil(o.GlTypeId) {
		var ret string
		return ret
	}
	return *o.GlTypeId
}

// GetGlTypeIdOk returns a tuple with the GlTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GLExportInventoryTransferOffset) GetGlTypeIdOk() (*string, bool) {
	if o == nil || IsNil(o.GlTypeId) {
		return nil, false
	}
	return o.GlTypeId, true
}

// HasGlTypeId returns a boolean if a field has been set.
func (o *GLExportInventoryTransferOffset) HasGlTypeId() bool {
	if o != nil && !IsNil(o.GlTypeId) {
		return true
	}

	return false
}

// SetGlTypeId gets a reference to the given string and assigns it to the GlTypeId field.
func (o *GLExportInventoryTransferOffset) SetGlTypeId(v string) {
	o.GlTypeId = &v
}

func (o GLExportInventoryTransferOffset) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GLExportInventoryTransferOffset) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if !IsNil(o.DocumentType) {
		toSerialize["documentType"] = o.DocumentType
	}
	if !IsNil(o.DocumentDate) {
		toSerialize["documentDate"] = o.DocumentDate
	}
	if !IsNil(o.AccountNumber) {
		toSerialize["accountNumber"] = o.AccountNumber
	}
	if !IsNil(o.GlClass) {
		toSerialize["glClass"] = o.GlClass
	}
	if o.Total.IsSet() {
		toSerialize["total"] = o.Total.Get()
	}
	if !IsNil(o.Memo) {
		toSerialize["memo"] = o.Memo
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.GlTypeId) {
		toSerialize["glTypeId"] = o.GlTypeId
	}
	return toSerialize, nil
}

type NullableGLExportInventoryTransferOffset struct {
	value *GLExportInventoryTransferOffset
	isSet bool
}

func (v NullableGLExportInventoryTransferOffset) Get() *GLExportInventoryTransferOffset {
	return v.value
}

func (v *NullableGLExportInventoryTransferOffset) Set(val *GLExportInventoryTransferOffset) {
	v.value = val
	v.isSet = true
}

func (v NullableGLExportInventoryTransferOffset) IsSet() bool {
	return v.isSet
}

func (v *NullableGLExportInventoryTransferOffset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGLExportInventoryTransferOffset(val *GLExportInventoryTransferOffset) *NullableGLExportInventoryTransferOffset {
	return &NullableGLExportInventoryTransferOffset{value: val, isSet: true}
}

func (v NullableGLExportInventoryTransferOffset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGLExportInventoryTransferOffset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


