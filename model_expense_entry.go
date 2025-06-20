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

// checks if the ExpenseEntry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExpenseEntry{}

// ExpenseEntry struct for ExpenseEntry
type ExpenseEntry struct {
	Id *int32 `json:"id,omitempty"`
	ExpenseReport *ExpenseReportReference `json:"expenseReport,omitempty"`
	Company *CompanyReference `json:"company,omitempty"`
	ChargeToId NullableInt32 `json:"chargeToId,omitempty"`
	// Gets or sets             company or chargeToType is required.
	ChargeToType NullableString `json:"chargeToType,omitempty"`
	Type ExpenseTypeReference `json:"type"`
	Member *MemberReference `json:"member,omitempty"`
	PaymentMethod *PaymentMethodReference `json:"paymentMethod,omitempty"`
	Classification *ClassificationReference `json:"classification,omitempty"`
	Amount NullableFloat64 `json:"amount"`
	BillableOption NullableString `json:"billableOption,omitempty"`
	Date time.Time `json:"date"`
	LocationId NullableInt32 `json:"locationId,omitempty"`
	BusinessUnitId NullableInt32 `json:"businessUnitId,omitempty"`
	Notes *string `json:"notes,omitempty"`
	Agreement *AgreementReference `json:"agreement,omitempty"`
	InvoiceAmount NullableFloat64 `json:"invoiceAmount,omitempty"`
	MobileGuid NullableString `json:"mobileGuid,omitempty"`
	Taxes []ExpenseTax `json:"taxes,omitempty"`
	Invoice *InvoiceReference `json:"invoice,omitempty"`
	Currency *CurrencyReference `json:"currency,omitempty"`
	Status NullableString `json:"status,omitempty"`
	BillAmount NullableFloat64 `json:"billAmount,omitempty"`
	AgreementAmount NullableFloat64 `json:"agreementAmount,omitempty"`
	OdometerStart NullableFloat64 `json:"odometerStart,omitempty"`
	OdometerEnd NullableFloat64 `json:"odometerEnd,omitempty"`
	Ticket *TicketReference `json:"ticket,omitempty"`
	Project *ProjectReference `json:"project,omitempty"`
	Phase *ProjectPhaseReference `json:"phase,omitempty"`
	Info *map[string]string `json:"_info,omitempty"`
	CustomFields []CustomFieldValue `json:"customFields,omitempty"`
}

type _ExpenseEntry ExpenseEntry

// NewExpenseEntry instantiates a new ExpenseEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExpenseEntry(type_ ExpenseTypeReference, amount NullableFloat64, date time.Time) *ExpenseEntry {
	this := ExpenseEntry{}
	this.Type = type_
	this.Amount = amount
	this.Date = date
	return &this
}

// NewExpenseEntryWithDefaults instantiates a new ExpenseEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExpenseEntryWithDefaults() *ExpenseEntry {
	this := ExpenseEntry{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ExpenseEntry) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpenseEntry) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ExpenseEntry) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ExpenseEntry) SetId(v int32) {
	o.Id = &v
}

// GetExpenseReport returns the ExpenseReport field value if set, zero value otherwise.
func (o *ExpenseEntry) GetExpenseReport() ExpenseReportReference {
	if o == nil || IsNil(o.ExpenseReport) {
		var ret ExpenseReportReference
		return ret
	}
	return *o.ExpenseReport
}

// GetExpenseReportOk returns a tuple with the ExpenseReport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpenseEntry) GetExpenseReportOk() (*ExpenseReportReference, bool) {
	if o == nil || IsNil(o.ExpenseReport) {
		return nil, false
	}
	return o.ExpenseReport, true
}

// HasExpenseReport returns a boolean if a field has been set.
func (o *ExpenseEntry) HasExpenseReport() bool {
	if o != nil && !IsNil(o.ExpenseReport) {
		return true
	}

	return false
}

// SetExpenseReport gets a reference to the given ExpenseReportReference and assigns it to the ExpenseReport field.
func (o *ExpenseEntry) SetExpenseReport(v ExpenseReportReference) {
	o.ExpenseReport = &v
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *ExpenseEntry) GetCompany() CompanyReference {
	if o == nil || IsNil(o.Company) {
		var ret CompanyReference
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpenseEntry) GetCompanyOk() (*CompanyReference, bool) {
	if o == nil || IsNil(o.Company) {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *ExpenseEntry) HasCompany() bool {
	if o != nil && !IsNil(o.Company) {
		return true
	}

	return false
}

// SetCompany gets a reference to the given CompanyReference and assigns it to the Company field.
func (o *ExpenseEntry) SetCompany(v CompanyReference) {
	o.Company = &v
}

// GetChargeToId returns the ChargeToId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExpenseEntry) GetChargeToId() int32 {
	if o == nil || IsNil(o.ChargeToId.Get()) {
		var ret int32
		return ret
	}
	return *o.ChargeToId.Get()
}

// GetChargeToIdOk returns a tuple with the ChargeToId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExpenseEntry) GetChargeToIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChargeToId.Get(), o.ChargeToId.IsSet()
}

// HasChargeToId returns a boolean if a field has been set.
func (o *ExpenseEntry) HasChargeToId() bool {
	if o != nil && o.ChargeToId.IsSet() {
		return true
	}

	return false
}

// SetChargeToId gets a reference to the given NullableInt32 and assigns it to the ChargeToId field.
func (o *ExpenseEntry) SetChargeToId(v int32) {
	o.ChargeToId.Set(&v)
}
// SetChargeToIdNil sets the value for ChargeToId to be an explicit nil
func (o *ExpenseEntry) SetChargeToIdNil() {
	o.ChargeToId.Set(nil)
}

// UnsetChargeToId ensures that no value is present for ChargeToId, not even an explicit nil
func (o *ExpenseEntry) UnsetChargeToId() {
	o.ChargeToId.Unset()
}

// GetChargeToType returns the ChargeToType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExpenseEntry) GetChargeToType() string {
	if o == nil || IsNil(o.ChargeToType.Get()) {
		var ret string
		return ret
	}
	return *o.ChargeToType.Get()
}

// GetChargeToTypeOk returns a tuple with the ChargeToType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExpenseEntry) GetChargeToTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChargeToType.Get(), o.ChargeToType.IsSet()
}

// HasChargeToType returns a boolean if a field has been set.
func (o *ExpenseEntry) HasChargeToType() bool {
	if o != nil && o.ChargeToType.IsSet() {
		return true
	}

	return false
}

// SetChargeToType gets a reference to the given NullableString and assigns it to the ChargeToType field.
func (o *ExpenseEntry) SetChargeToType(v string) {
	o.ChargeToType.Set(&v)
}
// SetChargeToTypeNil sets the value for ChargeToType to be an explicit nil
func (o *ExpenseEntry) SetChargeToTypeNil() {
	o.ChargeToType.Set(nil)
}

// UnsetChargeToType ensures that no value is present for ChargeToType, not even an explicit nil
func (o *ExpenseEntry) UnsetChargeToType() {
	o.ChargeToType.Unset()
}

// GetType returns the Type field value
func (o *ExpenseEntry) GetType() ExpenseTypeReference {
	if o == nil {
		var ret ExpenseTypeReference
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ExpenseEntry) GetTypeOk() (*ExpenseTypeReference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ExpenseEntry) SetType(v ExpenseTypeReference) {
	o.Type = v
}

// GetMember returns the Member field value if set, zero value otherwise.
func (o *ExpenseEntry) GetMember() MemberReference {
	if o == nil || IsNil(o.Member) {
		var ret MemberReference
		return ret
	}
	return *o.Member
}

// GetMemberOk returns a tuple with the Member field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpenseEntry) GetMemberOk() (*MemberReference, bool) {
	if o == nil || IsNil(o.Member) {
		return nil, false
	}
	return o.Member, true
}

// HasMember returns a boolean if a field has been set.
func (o *ExpenseEntry) HasMember() bool {
	if o != nil && !IsNil(o.Member) {
		return true
	}

	return false
}

// SetMember gets a reference to the given MemberReference and assigns it to the Member field.
func (o *ExpenseEntry) SetMember(v MemberReference) {
	o.Member = &v
}

// GetPaymentMethod returns the PaymentMethod field value if set, zero value otherwise.
func (o *ExpenseEntry) GetPaymentMethod() PaymentMethodReference {
	if o == nil || IsNil(o.PaymentMethod) {
		var ret PaymentMethodReference
		return ret
	}
	return *o.PaymentMethod
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpenseEntry) GetPaymentMethodOk() (*PaymentMethodReference, bool) {
	if o == nil || IsNil(o.PaymentMethod) {
		return nil, false
	}
	return o.PaymentMethod, true
}

// HasPaymentMethod returns a boolean if a field has been set.
func (o *ExpenseEntry) HasPaymentMethod() bool {
	if o != nil && !IsNil(o.PaymentMethod) {
		return true
	}

	return false
}

// SetPaymentMethod gets a reference to the given PaymentMethodReference and assigns it to the PaymentMethod field.
func (o *ExpenseEntry) SetPaymentMethod(v PaymentMethodReference) {
	o.PaymentMethod = &v
}

// GetClassification returns the Classification field value if set, zero value otherwise.
func (o *ExpenseEntry) GetClassification() ClassificationReference {
	if o == nil || IsNil(o.Classification) {
		var ret ClassificationReference
		return ret
	}
	return *o.Classification
}

// GetClassificationOk returns a tuple with the Classification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpenseEntry) GetClassificationOk() (*ClassificationReference, bool) {
	if o == nil || IsNil(o.Classification) {
		return nil, false
	}
	return o.Classification, true
}

// HasClassification returns a boolean if a field has been set.
func (o *ExpenseEntry) HasClassification() bool {
	if o != nil && !IsNil(o.Classification) {
		return true
	}

	return false
}

// SetClassification gets a reference to the given ClassificationReference and assigns it to the Classification field.
func (o *ExpenseEntry) SetClassification(v ClassificationReference) {
	o.Classification = &v
}

// GetAmount returns the Amount field value
// If the value is explicit nil, the zero value for float64 will be returned
func (o *ExpenseEntry) GetAmount() float64 {
	if o == nil || o.Amount.Get() == nil {
		var ret float64
		return ret
	}

	return *o.Amount.Get()
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExpenseEntry) GetAmountOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Amount.Get(), o.Amount.IsSet()
}

// SetAmount sets field value
func (o *ExpenseEntry) SetAmount(v float64) {
	o.Amount.Set(&v)
}

// GetBillableOption returns the BillableOption field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExpenseEntry) GetBillableOption() string {
	if o == nil || IsNil(o.BillableOption.Get()) {
		var ret string
		return ret
	}
	return *o.BillableOption.Get()
}

// GetBillableOptionOk returns a tuple with the BillableOption field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExpenseEntry) GetBillableOptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BillableOption.Get(), o.BillableOption.IsSet()
}

// HasBillableOption returns a boolean if a field has been set.
func (o *ExpenseEntry) HasBillableOption() bool {
	if o != nil && o.BillableOption.IsSet() {
		return true
	}

	return false
}

// SetBillableOption gets a reference to the given NullableString and assigns it to the BillableOption field.
func (o *ExpenseEntry) SetBillableOption(v string) {
	o.BillableOption.Set(&v)
}
// SetBillableOptionNil sets the value for BillableOption to be an explicit nil
func (o *ExpenseEntry) SetBillableOptionNil() {
	o.BillableOption.Set(nil)
}

// UnsetBillableOption ensures that no value is present for BillableOption, not even an explicit nil
func (o *ExpenseEntry) UnsetBillableOption() {
	o.BillableOption.Unset()
}

// GetDate returns the Date field value
func (o *ExpenseEntry) GetDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *ExpenseEntry) GetDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *ExpenseEntry) SetDate(v time.Time) {
	o.Date = v
}

// GetLocationId returns the LocationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExpenseEntry) GetLocationId() int32 {
	if o == nil || IsNil(o.LocationId.Get()) {
		var ret int32
		return ret
	}
	return *o.LocationId.Get()
}

// GetLocationIdOk returns a tuple with the LocationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExpenseEntry) GetLocationIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.LocationId.Get(), o.LocationId.IsSet()
}

// HasLocationId returns a boolean if a field has been set.
func (o *ExpenseEntry) HasLocationId() bool {
	if o != nil && o.LocationId.IsSet() {
		return true
	}

	return false
}

// SetLocationId gets a reference to the given NullableInt32 and assigns it to the LocationId field.
func (o *ExpenseEntry) SetLocationId(v int32) {
	o.LocationId.Set(&v)
}
// SetLocationIdNil sets the value for LocationId to be an explicit nil
func (o *ExpenseEntry) SetLocationIdNil() {
	o.LocationId.Set(nil)
}

// UnsetLocationId ensures that no value is present for LocationId, not even an explicit nil
func (o *ExpenseEntry) UnsetLocationId() {
	o.LocationId.Unset()
}

// GetBusinessUnitId returns the BusinessUnitId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExpenseEntry) GetBusinessUnitId() int32 {
	if o == nil || IsNil(o.BusinessUnitId.Get()) {
		var ret int32
		return ret
	}
	return *o.BusinessUnitId.Get()
}

// GetBusinessUnitIdOk returns a tuple with the BusinessUnitId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExpenseEntry) GetBusinessUnitIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.BusinessUnitId.Get(), o.BusinessUnitId.IsSet()
}

// HasBusinessUnitId returns a boolean if a field has been set.
func (o *ExpenseEntry) HasBusinessUnitId() bool {
	if o != nil && o.BusinessUnitId.IsSet() {
		return true
	}

	return false
}

// SetBusinessUnitId gets a reference to the given NullableInt32 and assigns it to the BusinessUnitId field.
func (o *ExpenseEntry) SetBusinessUnitId(v int32) {
	o.BusinessUnitId.Set(&v)
}
// SetBusinessUnitIdNil sets the value for BusinessUnitId to be an explicit nil
func (o *ExpenseEntry) SetBusinessUnitIdNil() {
	o.BusinessUnitId.Set(nil)
}

// UnsetBusinessUnitId ensures that no value is present for BusinessUnitId, not even an explicit nil
func (o *ExpenseEntry) UnsetBusinessUnitId() {
	o.BusinessUnitId.Unset()
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *ExpenseEntry) GetNotes() string {
	if o == nil || IsNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpenseEntry) GetNotesOk() (*string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *ExpenseEntry) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *ExpenseEntry) SetNotes(v string) {
	o.Notes = &v
}

// GetAgreement returns the Agreement field value if set, zero value otherwise.
func (o *ExpenseEntry) GetAgreement() AgreementReference {
	if o == nil || IsNil(o.Agreement) {
		var ret AgreementReference
		return ret
	}
	return *o.Agreement
}

// GetAgreementOk returns a tuple with the Agreement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpenseEntry) GetAgreementOk() (*AgreementReference, bool) {
	if o == nil || IsNil(o.Agreement) {
		return nil, false
	}
	return o.Agreement, true
}

// HasAgreement returns a boolean if a field has been set.
func (o *ExpenseEntry) HasAgreement() bool {
	if o != nil && !IsNil(o.Agreement) {
		return true
	}

	return false
}

// SetAgreement gets a reference to the given AgreementReference and assigns it to the Agreement field.
func (o *ExpenseEntry) SetAgreement(v AgreementReference) {
	o.Agreement = &v
}

// GetInvoiceAmount returns the InvoiceAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExpenseEntry) GetInvoiceAmount() float64 {
	if o == nil || IsNil(o.InvoiceAmount.Get()) {
		var ret float64
		return ret
	}
	return *o.InvoiceAmount.Get()
}

// GetInvoiceAmountOk returns a tuple with the InvoiceAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExpenseEntry) GetInvoiceAmountOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.InvoiceAmount.Get(), o.InvoiceAmount.IsSet()
}

// HasInvoiceAmount returns a boolean if a field has been set.
func (o *ExpenseEntry) HasInvoiceAmount() bool {
	if o != nil && o.InvoiceAmount.IsSet() {
		return true
	}

	return false
}

// SetInvoiceAmount gets a reference to the given NullableFloat64 and assigns it to the InvoiceAmount field.
func (o *ExpenseEntry) SetInvoiceAmount(v float64) {
	o.InvoiceAmount.Set(&v)
}
// SetInvoiceAmountNil sets the value for InvoiceAmount to be an explicit nil
func (o *ExpenseEntry) SetInvoiceAmountNil() {
	o.InvoiceAmount.Set(nil)
}

// UnsetInvoiceAmount ensures that no value is present for InvoiceAmount, not even an explicit nil
func (o *ExpenseEntry) UnsetInvoiceAmount() {
	o.InvoiceAmount.Unset()
}

// GetMobileGuid returns the MobileGuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExpenseEntry) GetMobileGuid() string {
	if o == nil || IsNil(o.MobileGuid.Get()) {
		var ret string
		return ret
	}
	return *o.MobileGuid.Get()
}

// GetMobileGuidOk returns a tuple with the MobileGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExpenseEntry) GetMobileGuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MobileGuid.Get(), o.MobileGuid.IsSet()
}

// HasMobileGuid returns a boolean if a field has been set.
func (o *ExpenseEntry) HasMobileGuid() bool {
	if o != nil && o.MobileGuid.IsSet() {
		return true
	}

	return false
}

// SetMobileGuid gets a reference to the given NullableString and assigns it to the MobileGuid field.
func (o *ExpenseEntry) SetMobileGuid(v string) {
	o.MobileGuid.Set(&v)
}
// SetMobileGuidNil sets the value for MobileGuid to be an explicit nil
func (o *ExpenseEntry) SetMobileGuidNil() {
	o.MobileGuid.Set(nil)
}

// UnsetMobileGuid ensures that no value is present for MobileGuid, not even an explicit nil
func (o *ExpenseEntry) UnsetMobileGuid() {
	o.MobileGuid.Unset()
}

// GetTaxes returns the Taxes field value if set, zero value otherwise.
func (o *ExpenseEntry) GetTaxes() []ExpenseTax {
	if o == nil || IsNil(o.Taxes) {
		var ret []ExpenseTax
		return ret
	}
	return o.Taxes
}

// GetTaxesOk returns a tuple with the Taxes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpenseEntry) GetTaxesOk() ([]ExpenseTax, bool) {
	if o == nil || IsNil(o.Taxes) {
		return nil, false
	}
	return o.Taxes, true
}

// HasTaxes returns a boolean if a field has been set.
func (o *ExpenseEntry) HasTaxes() bool {
	if o != nil && !IsNil(o.Taxes) {
		return true
	}

	return false
}

// SetTaxes gets a reference to the given []ExpenseTax and assigns it to the Taxes field.
func (o *ExpenseEntry) SetTaxes(v []ExpenseTax) {
	o.Taxes = v
}

// GetInvoice returns the Invoice field value if set, zero value otherwise.
func (o *ExpenseEntry) GetInvoice() InvoiceReference {
	if o == nil || IsNil(o.Invoice) {
		var ret InvoiceReference
		return ret
	}
	return *o.Invoice
}

// GetInvoiceOk returns a tuple with the Invoice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpenseEntry) GetInvoiceOk() (*InvoiceReference, bool) {
	if o == nil || IsNil(o.Invoice) {
		return nil, false
	}
	return o.Invoice, true
}

// HasInvoice returns a boolean if a field has been set.
func (o *ExpenseEntry) HasInvoice() bool {
	if o != nil && !IsNil(o.Invoice) {
		return true
	}

	return false
}

// SetInvoice gets a reference to the given InvoiceReference and assigns it to the Invoice field.
func (o *ExpenseEntry) SetInvoice(v InvoiceReference) {
	o.Invoice = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *ExpenseEntry) GetCurrency() CurrencyReference {
	if o == nil || IsNil(o.Currency) {
		var ret CurrencyReference
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpenseEntry) GetCurrencyOk() (*CurrencyReference, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *ExpenseEntry) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given CurrencyReference and assigns it to the Currency field.
func (o *ExpenseEntry) SetCurrency(v CurrencyReference) {
	o.Currency = &v
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExpenseEntry) GetStatus() string {
	if o == nil || IsNil(o.Status.Get()) {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExpenseEntry) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *ExpenseEntry) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *ExpenseEntry) SetStatus(v string) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *ExpenseEntry) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *ExpenseEntry) UnsetStatus() {
	o.Status.Unset()
}

// GetBillAmount returns the BillAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExpenseEntry) GetBillAmount() float64 {
	if o == nil || IsNil(o.BillAmount.Get()) {
		var ret float64
		return ret
	}
	return *o.BillAmount.Get()
}

// GetBillAmountOk returns a tuple with the BillAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExpenseEntry) GetBillAmountOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.BillAmount.Get(), o.BillAmount.IsSet()
}

// HasBillAmount returns a boolean if a field has been set.
func (o *ExpenseEntry) HasBillAmount() bool {
	if o != nil && o.BillAmount.IsSet() {
		return true
	}

	return false
}

// SetBillAmount gets a reference to the given NullableFloat64 and assigns it to the BillAmount field.
func (o *ExpenseEntry) SetBillAmount(v float64) {
	o.BillAmount.Set(&v)
}
// SetBillAmountNil sets the value for BillAmount to be an explicit nil
func (o *ExpenseEntry) SetBillAmountNil() {
	o.BillAmount.Set(nil)
}

// UnsetBillAmount ensures that no value is present for BillAmount, not even an explicit nil
func (o *ExpenseEntry) UnsetBillAmount() {
	o.BillAmount.Unset()
}

// GetAgreementAmount returns the AgreementAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExpenseEntry) GetAgreementAmount() float64 {
	if o == nil || IsNil(o.AgreementAmount.Get()) {
		var ret float64
		return ret
	}
	return *o.AgreementAmount.Get()
}

// GetAgreementAmountOk returns a tuple with the AgreementAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExpenseEntry) GetAgreementAmountOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.AgreementAmount.Get(), o.AgreementAmount.IsSet()
}

// HasAgreementAmount returns a boolean if a field has been set.
func (o *ExpenseEntry) HasAgreementAmount() bool {
	if o != nil && o.AgreementAmount.IsSet() {
		return true
	}

	return false
}

// SetAgreementAmount gets a reference to the given NullableFloat64 and assigns it to the AgreementAmount field.
func (o *ExpenseEntry) SetAgreementAmount(v float64) {
	o.AgreementAmount.Set(&v)
}
// SetAgreementAmountNil sets the value for AgreementAmount to be an explicit nil
func (o *ExpenseEntry) SetAgreementAmountNil() {
	o.AgreementAmount.Set(nil)
}

// UnsetAgreementAmount ensures that no value is present for AgreementAmount, not even an explicit nil
func (o *ExpenseEntry) UnsetAgreementAmount() {
	o.AgreementAmount.Unset()
}

// GetOdometerStart returns the OdometerStart field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExpenseEntry) GetOdometerStart() float64 {
	if o == nil || IsNil(o.OdometerStart.Get()) {
		var ret float64
		return ret
	}
	return *o.OdometerStart.Get()
}

// GetOdometerStartOk returns a tuple with the OdometerStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExpenseEntry) GetOdometerStartOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.OdometerStart.Get(), o.OdometerStart.IsSet()
}

// HasOdometerStart returns a boolean if a field has been set.
func (o *ExpenseEntry) HasOdometerStart() bool {
	if o != nil && o.OdometerStart.IsSet() {
		return true
	}

	return false
}

// SetOdometerStart gets a reference to the given NullableFloat64 and assigns it to the OdometerStart field.
func (o *ExpenseEntry) SetOdometerStart(v float64) {
	o.OdometerStart.Set(&v)
}
// SetOdometerStartNil sets the value for OdometerStart to be an explicit nil
func (o *ExpenseEntry) SetOdometerStartNil() {
	o.OdometerStart.Set(nil)
}

// UnsetOdometerStart ensures that no value is present for OdometerStart, not even an explicit nil
func (o *ExpenseEntry) UnsetOdometerStart() {
	o.OdometerStart.Unset()
}

// GetOdometerEnd returns the OdometerEnd field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExpenseEntry) GetOdometerEnd() float64 {
	if o == nil || IsNil(o.OdometerEnd.Get()) {
		var ret float64
		return ret
	}
	return *o.OdometerEnd.Get()
}

// GetOdometerEndOk returns a tuple with the OdometerEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExpenseEntry) GetOdometerEndOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.OdometerEnd.Get(), o.OdometerEnd.IsSet()
}

// HasOdometerEnd returns a boolean if a field has been set.
func (o *ExpenseEntry) HasOdometerEnd() bool {
	if o != nil && o.OdometerEnd.IsSet() {
		return true
	}

	return false
}

// SetOdometerEnd gets a reference to the given NullableFloat64 and assigns it to the OdometerEnd field.
func (o *ExpenseEntry) SetOdometerEnd(v float64) {
	o.OdometerEnd.Set(&v)
}
// SetOdometerEndNil sets the value for OdometerEnd to be an explicit nil
func (o *ExpenseEntry) SetOdometerEndNil() {
	o.OdometerEnd.Set(nil)
}

// UnsetOdometerEnd ensures that no value is present for OdometerEnd, not even an explicit nil
func (o *ExpenseEntry) UnsetOdometerEnd() {
	o.OdometerEnd.Unset()
}

// GetTicket returns the Ticket field value if set, zero value otherwise.
func (o *ExpenseEntry) GetTicket() TicketReference {
	if o == nil || IsNil(o.Ticket) {
		var ret TicketReference
		return ret
	}
	return *o.Ticket
}

// GetTicketOk returns a tuple with the Ticket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpenseEntry) GetTicketOk() (*TicketReference, bool) {
	if o == nil || IsNil(o.Ticket) {
		return nil, false
	}
	return o.Ticket, true
}

// HasTicket returns a boolean if a field has been set.
func (o *ExpenseEntry) HasTicket() bool {
	if o != nil && !IsNil(o.Ticket) {
		return true
	}

	return false
}

// SetTicket gets a reference to the given TicketReference and assigns it to the Ticket field.
func (o *ExpenseEntry) SetTicket(v TicketReference) {
	o.Ticket = &v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *ExpenseEntry) GetProject() ProjectReference {
	if o == nil || IsNil(o.Project) {
		var ret ProjectReference
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpenseEntry) GetProjectOk() (*ProjectReference, bool) {
	if o == nil || IsNil(o.Project) {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *ExpenseEntry) HasProject() bool {
	if o != nil && !IsNil(o.Project) {
		return true
	}

	return false
}

// SetProject gets a reference to the given ProjectReference and assigns it to the Project field.
func (o *ExpenseEntry) SetProject(v ProjectReference) {
	o.Project = &v
}

// GetPhase returns the Phase field value if set, zero value otherwise.
func (o *ExpenseEntry) GetPhase() ProjectPhaseReference {
	if o == nil || IsNil(o.Phase) {
		var ret ProjectPhaseReference
		return ret
	}
	return *o.Phase
}

// GetPhaseOk returns a tuple with the Phase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpenseEntry) GetPhaseOk() (*ProjectPhaseReference, bool) {
	if o == nil || IsNil(o.Phase) {
		return nil, false
	}
	return o.Phase, true
}

// HasPhase returns a boolean if a field has been set.
func (o *ExpenseEntry) HasPhase() bool {
	if o != nil && !IsNil(o.Phase) {
		return true
	}

	return false
}

// SetPhase gets a reference to the given ProjectPhaseReference and assigns it to the Phase field.
func (o *ExpenseEntry) SetPhase(v ProjectPhaseReference) {
	o.Phase = &v
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *ExpenseEntry) GetInfo() map[string]string {
	if o == nil || IsNil(o.Info) {
		var ret map[string]string
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpenseEntry) GetInfoOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Info) {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *ExpenseEntry) HasInfo() bool {
	if o != nil && !IsNil(o.Info) {
		return true
	}

	return false
}

// SetInfo gets a reference to the given map[string]string and assigns it to the Info field.
func (o *ExpenseEntry) SetInfo(v map[string]string) {
	o.Info = &v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *ExpenseEntry) GetCustomFields() []CustomFieldValue {
	if o == nil || IsNil(o.CustomFields) {
		var ret []CustomFieldValue
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpenseEntry) GetCustomFieldsOk() ([]CustomFieldValue, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *ExpenseEntry) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given []CustomFieldValue and assigns it to the CustomFields field.
func (o *ExpenseEntry) SetCustomFields(v []CustomFieldValue) {
	o.CustomFields = v
}

func (o ExpenseEntry) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExpenseEntry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ExpenseReport) {
		toSerialize["expenseReport"] = o.ExpenseReport
	}
	if !IsNil(o.Company) {
		toSerialize["company"] = o.Company
	}
	if o.ChargeToId.IsSet() {
		toSerialize["chargeToId"] = o.ChargeToId.Get()
	}
	if o.ChargeToType.IsSet() {
		toSerialize["chargeToType"] = o.ChargeToType.Get()
	}
	toSerialize["type"] = o.Type
	if !IsNil(o.Member) {
		toSerialize["member"] = o.Member
	}
	if !IsNil(o.PaymentMethod) {
		toSerialize["paymentMethod"] = o.PaymentMethod
	}
	if !IsNil(o.Classification) {
		toSerialize["classification"] = o.Classification
	}
	toSerialize["amount"] = o.Amount.Get()
	if o.BillableOption.IsSet() {
		toSerialize["billableOption"] = o.BillableOption.Get()
	}
	toSerialize["date"] = o.Date
	if o.LocationId.IsSet() {
		toSerialize["locationId"] = o.LocationId.Get()
	}
	if o.BusinessUnitId.IsSet() {
		toSerialize["businessUnitId"] = o.BusinessUnitId.Get()
	}
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	if !IsNil(o.Agreement) {
		toSerialize["agreement"] = o.Agreement
	}
	if o.InvoiceAmount.IsSet() {
		toSerialize["invoiceAmount"] = o.InvoiceAmount.Get()
	}
	if o.MobileGuid.IsSet() {
		toSerialize["mobileGuid"] = o.MobileGuid.Get()
	}
	if !IsNil(o.Taxes) {
		toSerialize["taxes"] = o.Taxes
	}
	if !IsNil(o.Invoice) {
		toSerialize["invoice"] = o.Invoice
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if o.BillAmount.IsSet() {
		toSerialize["billAmount"] = o.BillAmount.Get()
	}
	if o.AgreementAmount.IsSet() {
		toSerialize["agreementAmount"] = o.AgreementAmount.Get()
	}
	if o.OdometerStart.IsSet() {
		toSerialize["odometerStart"] = o.OdometerStart.Get()
	}
	if o.OdometerEnd.IsSet() {
		toSerialize["odometerEnd"] = o.OdometerEnd.Get()
	}
	if !IsNil(o.Ticket) {
		toSerialize["ticket"] = o.Ticket
	}
	if !IsNil(o.Project) {
		toSerialize["project"] = o.Project
	}
	if !IsNil(o.Phase) {
		toSerialize["phase"] = o.Phase
	}
	if !IsNil(o.Info) {
		toSerialize["_info"] = o.Info
	}
	if !IsNil(o.CustomFields) {
		toSerialize["customFields"] = o.CustomFields
	}
	return toSerialize, nil
}

func (o *ExpenseEntry) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"amount",
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

	varExpenseEntry := _ExpenseEntry{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varExpenseEntry)

	if err != nil {
		return err
	}

	*o = ExpenseEntry(varExpenseEntry)

	return err
}

type NullableExpenseEntry struct {
	value *ExpenseEntry
	isSet bool
}

func (v NullableExpenseEntry) Get() *ExpenseEntry {
	return v.value
}

func (v *NullableExpenseEntry) Set(val *ExpenseEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableExpenseEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableExpenseEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExpenseEntry(val *ExpenseEntry) *NullableExpenseEntry {
	return &NullableExpenseEntry{value: val, isSet: true}
}

func (v NullableExpenseEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExpenseEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


