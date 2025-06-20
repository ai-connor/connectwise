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

// checks if the InvoiceInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvoiceInfo{}

// InvoiceInfo struct for InvoiceInfo
type InvoiceInfo struct {
	Id *int32 `json:"id,omitempty"`
	Invoice *Invoice `json:"invoice,omitempty"`
	InvoiceTemplate *InvoiceTemplate `json:"invoiceTemplate,omitempty"`
	Products []ProductItem `json:"products,omitempty"`
	BundledComponentsInfo []ProductComponent `json:"bundledComponentsInfo,omitempty"`
	Expenses []ExpenseEntry `json:"expenses,omitempty"`
	TimeEntries []TimeEntry `json:"timeEntries,omitempty"`
	Logo *DocumentInfo `json:"logo,omitempty"`
	BillingSetup *BillingSetup `json:"billingSetup,omitempty"`
	AgreementBillingInfo []AgreementBillingInfo `json:"agreementBillingInfo,omitempty"`
	Info *map[string]string `json:"_info,omitempty"`
}

// NewInvoiceInfo instantiates a new InvoiceInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoiceInfo() *InvoiceInfo {
	this := InvoiceInfo{}
	return &this
}

// NewInvoiceInfoWithDefaults instantiates a new InvoiceInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceInfoWithDefaults() *InvoiceInfo {
	this := InvoiceInfo{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InvoiceInfo) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceInfo) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InvoiceInfo) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *InvoiceInfo) SetId(v int32) {
	o.Id = &v
}

// GetInvoice returns the Invoice field value if set, zero value otherwise.
func (o *InvoiceInfo) GetInvoice() Invoice {
	if o == nil || IsNil(o.Invoice) {
		var ret Invoice
		return ret
	}
	return *o.Invoice
}

// GetInvoiceOk returns a tuple with the Invoice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceInfo) GetInvoiceOk() (*Invoice, bool) {
	if o == nil || IsNil(o.Invoice) {
		return nil, false
	}
	return o.Invoice, true
}

// HasInvoice returns a boolean if a field has been set.
func (o *InvoiceInfo) HasInvoice() bool {
	if o != nil && !IsNil(o.Invoice) {
		return true
	}

	return false
}

// SetInvoice gets a reference to the given Invoice and assigns it to the Invoice field.
func (o *InvoiceInfo) SetInvoice(v Invoice) {
	o.Invoice = &v
}

// GetInvoiceTemplate returns the InvoiceTemplate field value if set, zero value otherwise.
func (o *InvoiceInfo) GetInvoiceTemplate() InvoiceTemplate {
	if o == nil || IsNil(o.InvoiceTemplate) {
		var ret InvoiceTemplate
		return ret
	}
	return *o.InvoiceTemplate
}

// GetInvoiceTemplateOk returns a tuple with the InvoiceTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceInfo) GetInvoiceTemplateOk() (*InvoiceTemplate, bool) {
	if o == nil || IsNil(o.InvoiceTemplate) {
		return nil, false
	}
	return o.InvoiceTemplate, true
}

// HasInvoiceTemplate returns a boolean if a field has been set.
func (o *InvoiceInfo) HasInvoiceTemplate() bool {
	if o != nil && !IsNil(o.InvoiceTemplate) {
		return true
	}

	return false
}

// SetInvoiceTemplate gets a reference to the given InvoiceTemplate and assigns it to the InvoiceTemplate field.
func (o *InvoiceInfo) SetInvoiceTemplate(v InvoiceTemplate) {
	o.InvoiceTemplate = &v
}

// GetProducts returns the Products field value if set, zero value otherwise.
func (o *InvoiceInfo) GetProducts() []ProductItem {
	if o == nil || IsNil(o.Products) {
		var ret []ProductItem
		return ret
	}
	return o.Products
}

// GetProductsOk returns a tuple with the Products field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceInfo) GetProductsOk() ([]ProductItem, bool) {
	if o == nil || IsNil(o.Products) {
		return nil, false
	}
	return o.Products, true
}

// HasProducts returns a boolean if a field has been set.
func (o *InvoiceInfo) HasProducts() bool {
	if o != nil && !IsNil(o.Products) {
		return true
	}

	return false
}

// SetProducts gets a reference to the given []ProductItem and assigns it to the Products field.
func (o *InvoiceInfo) SetProducts(v []ProductItem) {
	o.Products = v
}

// GetBundledComponentsInfo returns the BundledComponentsInfo field value if set, zero value otherwise.
func (o *InvoiceInfo) GetBundledComponentsInfo() []ProductComponent {
	if o == nil || IsNil(o.BundledComponentsInfo) {
		var ret []ProductComponent
		return ret
	}
	return o.BundledComponentsInfo
}

// GetBundledComponentsInfoOk returns a tuple with the BundledComponentsInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceInfo) GetBundledComponentsInfoOk() ([]ProductComponent, bool) {
	if o == nil || IsNil(o.BundledComponentsInfo) {
		return nil, false
	}
	return o.BundledComponentsInfo, true
}

// HasBundledComponentsInfo returns a boolean if a field has been set.
func (o *InvoiceInfo) HasBundledComponentsInfo() bool {
	if o != nil && !IsNil(o.BundledComponentsInfo) {
		return true
	}

	return false
}

// SetBundledComponentsInfo gets a reference to the given []ProductComponent and assigns it to the BundledComponentsInfo field.
func (o *InvoiceInfo) SetBundledComponentsInfo(v []ProductComponent) {
	o.BundledComponentsInfo = v
}

// GetExpenses returns the Expenses field value if set, zero value otherwise.
func (o *InvoiceInfo) GetExpenses() []ExpenseEntry {
	if o == nil || IsNil(o.Expenses) {
		var ret []ExpenseEntry
		return ret
	}
	return o.Expenses
}

// GetExpensesOk returns a tuple with the Expenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceInfo) GetExpensesOk() ([]ExpenseEntry, bool) {
	if o == nil || IsNil(o.Expenses) {
		return nil, false
	}
	return o.Expenses, true
}

// HasExpenses returns a boolean if a field has been set.
func (o *InvoiceInfo) HasExpenses() bool {
	if o != nil && !IsNil(o.Expenses) {
		return true
	}

	return false
}

// SetExpenses gets a reference to the given []ExpenseEntry and assigns it to the Expenses field.
func (o *InvoiceInfo) SetExpenses(v []ExpenseEntry) {
	o.Expenses = v
}

// GetTimeEntries returns the TimeEntries field value if set, zero value otherwise.
func (o *InvoiceInfo) GetTimeEntries() []TimeEntry {
	if o == nil || IsNil(o.TimeEntries) {
		var ret []TimeEntry
		return ret
	}
	return o.TimeEntries
}

// GetTimeEntriesOk returns a tuple with the TimeEntries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceInfo) GetTimeEntriesOk() ([]TimeEntry, bool) {
	if o == nil || IsNil(o.TimeEntries) {
		return nil, false
	}
	return o.TimeEntries, true
}

// HasTimeEntries returns a boolean if a field has been set.
func (o *InvoiceInfo) HasTimeEntries() bool {
	if o != nil && !IsNil(o.TimeEntries) {
		return true
	}

	return false
}

// SetTimeEntries gets a reference to the given []TimeEntry and assigns it to the TimeEntries field.
func (o *InvoiceInfo) SetTimeEntries(v []TimeEntry) {
	o.TimeEntries = v
}

// GetLogo returns the Logo field value if set, zero value otherwise.
func (o *InvoiceInfo) GetLogo() DocumentInfo {
	if o == nil || IsNil(o.Logo) {
		var ret DocumentInfo
		return ret
	}
	return *o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceInfo) GetLogoOk() (*DocumentInfo, bool) {
	if o == nil || IsNil(o.Logo) {
		return nil, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *InvoiceInfo) HasLogo() bool {
	if o != nil && !IsNil(o.Logo) {
		return true
	}

	return false
}

// SetLogo gets a reference to the given DocumentInfo and assigns it to the Logo field.
func (o *InvoiceInfo) SetLogo(v DocumentInfo) {
	o.Logo = &v
}

// GetBillingSetup returns the BillingSetup field value if set, zero value otherwise.
func (o *InvoiceInfo) GetBillingSetup() BillingSetup {
	if o == nil || IsNil(o.BillingSetup) {
		var ret BillingSetup
		return ret
	}
	return *o.BillingSetup
}

// GetBillingSetupOk returns a tuple with the BillingSetup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceInfo) GetBillingSetupOk() (*BillingSetup, bool) {
	if o == nil || IsNil(o.BillingSetup) {
		return nil, false
	}
	return o.BillingSetup, true
}

// HasBillingSetup returns a boolean if a field has been set.
func (o *InvoiceInfo) HasBillingSetup() bool {
	if o != nil && !IsNil(o.BillingSetup) {
		return true
	}

	return false
}

// SetBillingSetup gets a reference to the given BillingSetup and assigns it to the BillingSetup field.
func (o *InvoiceInfo) SetBillingSetup(v BillingSetup) {
	o.BillingSetup = &v
}

// GetAgreementBillingInfo returns the AgreementBillingInfo field value if set, zero value otherwise.
func (o *InvoiceInfo) GetAgreementBillingInfo() []AgreementBillingInfo {
	if o == nil || IsNil(o.AgreementBillingInfo) {
		var ret []AgreementBillingInfo
		return ret
	}
	return o.AgreementBillingInfo
}

// GetAgreementBillingInfoOk returns a tuple with the AgreementBillingInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceInfo) GetAgreementBillingInfoOk() ([]AgreementBillingInfo, bool) {
	if o == nil || IsNil(o.AgreementBillingInfo) {
		return nil, false
	}
	return o.AgreementBillingInfo, true
}

// HasAgreementBillingInfo returns a boolean if a field has been set.
func (o *InvoiceInfo) HasAgreementBillingInfo() bool {
	if o != nil && !IsNil(o.AgreementBillingInfo) {
		return true
	}

	return false
}

// SetAgreementBillingInfo gets a reference to the given []AgreementBillingInfo and assigns it to the AgreementBillingInfo field.
func (o *InvoiceInfo) SetAgreementBillingInfo(v []AgreementBillingInfo) {
	o.AgreementBillingInfo = v
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *InvoiceInfo) GetInfo() map[string]string {
	if o == nil || IsNil(o.Info) {
		var ret map[string]string
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceInfo) GetInfoOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Info) {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *InvoiceInfo) HasInfo() bool {
	if o != nil && !IsNil(o.Info) {
		return true
	}

	return false
}

// SetInfo gets a reference to the given map[string]string and assigns it to the Info field.
func (o *InvoiceInfo) SetInfo(v map[string]string) {
	o.Info = &v
}

func (o InvoiceInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvoiceInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Invoice) {
		toSerialize["invoice"] = o.Invoice
	}
	if !IsNil(o.InvoiceTemplate) {
		toSerialize["invoiceTemplate"] = o.InvoiceTemplate
	}
	if !IsNil(o.Products) {
		toSerialize["products"] = o.Products
	}
	if !IsNil(o.BundledComponentsInfo) {
		toSerialize["bundledComponentsInfo"] = o.BundledComponentsInfo
	}
	if !IsNil(o.Expenses) {
		toSerialize["expenses"] = o.Expenses
	}
	if !IsNil(o.TimeEntries) {
		toSerialize["timeEntries"] = o.TimeEntries
	}
	if !IsNil(o.Logo) {
		toSerialize["logo"] = o.Logo
	}
	if !IsNil(o.BillingSetup) {
		toSerialize["billingSetup"] = o.BillingSetup
	}
	if !IsNil(o.AgreementBillingInfo) {
		toSerialize["agreementBillingInfo"] = o.AgreementBillingInfo
	}
	if !IsNil(o.Info) {
		toSerialize["_info"] = o.Info
	}
	return toSerialize, nil
}

type NullableInvoiceInfo struct {
	value *InvoiceInfo
	isSet bool
}

func (v NullableInvoiceInfo) Get() *InvoiceInfo {
	return v.value
}

func (v *NullableInvoiceInfo) Set(val *InvoiceInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceInfo(val *InvoiceInfo) *NullableInvoiceInfo {
	return &NullableInvoiceInfo{value: val, isSet: true}
}

func (v NullableInvoiceInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


