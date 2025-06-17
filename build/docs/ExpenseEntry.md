# ExpenseEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ExpenseReport** | Pointer to [**ExpenseReportReference**](ExpenseReportReference.md) |  | [optional] 
**Company** | Pointer to [**CompanyReference**](CompanyReference.md) |  | [optional] 
**ChargeToId** | Pointer to **NullableInt32** |  | [optional] 
**ChargeToType** | Pointer to **NullableString** | Gets or sets             company or chargeToType is required. | [optional] 
**Type** | [**ExpenseTypeReference**](ExpenseTypeReference.md) |  | 
**Member** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**PaymentMethod** | Pointer to [**PaymentMethodReference**](PaymentMethodReference.md) |  | [optional] 
**Classification** | Pointer to [**ClassificationReference**](ClassificationReference.md) |  | [optional] 
**Amount** | **NullableFloat64** |  | 
**BillableOption** | Pointer to **NullableString** |  | [optional] 
**Date** | **time.Time** |  | 
**LocationId** | Pointer to **NullableInt32** |  | [optional] 
**BusinessUnitId** | Pointer to **NullableInt32** |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**Agreement** | Pointer to [**AgreementReference**](AgreementReference.md) |  | [optional] 
**InvoiceAmount** | Pointer to **NullableFloat64** |  | [optional] 
**MobileGuid** | Pointer to **NullableString** |  | [optional] 
**Taxes** | Pointer to [**[]ExpenseTax**](ExpenseTax.md) |  | [optional] 
**Invoice** | Pointer to [**InvoiceReference**](InvoiceReference.md) |  | [optional] 
**Currency** | Pointer to [**CurrencyReference**](CurrencyReference.md) |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**BillAmount** | Pointer to **NullableFloat64** |  | [optional] 
**AgreementAmount** | Pointer to **NullableFloat64** |  | [optional] 
**OdometerStart** | Pointer to **NullableFloat64** |  | [optional] 
**OdometerEnd** | Pointer to **NullableFloat64** |  | [optional] 
**Ticket** | Pointer to [**TicketReference**](TicketReference.md) |  | [optional] 
**Project** | Pointer to [**ProjectReference**](ProjectReference.md) |  | [optional] 
**Phase** | Pointer to [**ProjectPhaseReference**](ProjectPhaseReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 
**CustomFields** | Pointer to [**[]CustomFieldValue**](CustomFieldValue.md) |  | [optional] 

## Methods

### NewExpenseEntry

`func NewExpenseEntry(type_ ExpenseTypeReference, amount NullableFloat64, date time.Time, ) *ExpenseEntry`

NewExpenseEntry instantiates a new ExpenseEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpenseEntryWithDefaults

`func NewExpenseEntryWithDefaults() *ExpenseEntry`

NewExpenseEntryWithDefaults instantiates a new ExpenseEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExpenseEntry) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExpenseEntry) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExpenseEntry) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ExpenseEntry) HasId() bool`

HasId returns a boolean if a field has been set.

### GetExpenseReport

`func (o *ExpenseEntry) GetExpenseReport() ExpenseReportReference`

GetExpenseReport returns the ExpenseReport field if non-nil, zero value otherwise.

### GetExpenseReportOk

`func (o *ExpenseEntry) GetExpenseReportOk() (*ExpenseReportReference, bool)`

GetExpenseReportOk returns a tuple with the ExpenseReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenseReport

`func (o *ExpenseEntry) SetExpenseReport(v ExpenseReportReference)`

SetExpenseReport sets ExpenseReport field to given value.

### HasExpenseReport

`func (o *ExpenseEntry) HasExpenseReport() bool`

HasExpenseReport returns a boolean if a field has been set.

### GetCompany

`func (o *ExpenseEntry) GetCompany() CompanyReference`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *ExpenseEntry) GetCompanyOk() (*CompanyReference, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *ExpenseEntry) SetCompany(v CompanyReference)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *ExpenseEntry) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetChargeToId

`func (o *ExpenseEntry) GetChargeToId() int32`

GetChargeToId returns the ChargeToId field if non-nil, zero value otherwise.

### GetChargeToIdOk

`func (o *ExpenseEntry) GetChargeToIdOk() (*int32, bool)`

GetChargeToIdOk returns a tuple with the ChargeToId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeToId

`func (o *ExpenseEntry) SetChargeToId(v int32)`

SetChargeToId sets ChargeToId field to given value.

### HasChargeToId

`func (o *ExpenseEntry) HasChargeToId() bool`

HasChargeToId returns a boolean if a field has been set.

### SetChargeToIdNil

`func (o *ExpenseEntry) SetChargeToIdNil(b bool)`

 SetChargeToIdNil sets the value for ChargeToId to be an explicit nil

### UnsetChargeToId
`func (o *ExpenseEntry) UnsetChargeToId()`

UnsetChargeToId ensures that no value is present for ChargeToId, not even an explicit nil
### GetChargeToType

`func (o *ExpenseEntry) GetChargeToType() string`

GetChargeToType returns the ChargeToType field if non-nil, zero value otherwise.

### GetChargeToTypeOk

`func (o *ExpenseEntry) GetChargeToTypeOk() (*string, bool)`

GetChargeToTypeOk returns a tuple with the ChargeToType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeToType

`func (o *ExpenseEntry) SetChargeToType(v string)`

SetChargeToType sets ChargeToType field to given value.

### HasChargeToType

`func (o *ExpenseEntry) HasChargeToType() bool`

HasChargeToType returns a boolean if a field has been set.

### SetChargeToTypeNil

`func (o *ExpenseEntry) SetChargeToTypeNil(b bool)`

 SetChargeToTypeNil sets the value for ChargeToType to be an explicit nil

### UnsetChargeToType
`func (o *ExpenseEntry) UnsetChargeToType()`

UnsetChargeToType ensures that no value is present for ChargeToType, not even an explicit nil
### GetType

`func (o *ExpenseEntry) GetType() ExpenseTypeReference`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExpenseEntry) GetTypeOk() (*ExpenseTypeReference, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExpenseEntry) SetType(v ExpenseTypeReference)`

SetType sets Type field to given value.


### GetMember

`func (o *ExpenseEntry) GetMember() MemberReference`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *ExpenseEntry) GetMemberOk() (*MemberReference, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *ExpenseEntry) SetMember(v MemberReference)`

SetMember sets Member field to given value.

### HasMember

`func (o *ExpenseEntry) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *ExpenseEntry) GetPaymentMethod() PaymentMethodReference`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *ExpenseEntry) GetPaymentMethodOk() (*PaymentMethodReference, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *ExpenseEntry) SetPaymentMethod(v PaymentMethodReference)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *ExpenseEntry) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetClassification

`func (o *ExpenseEntry) GetClassification() ClassificationReference`

GetClassification returns the Classification field if non-nil, zero value otherwise.

### GetClassificationOk

`func (o *ExpenseEntry) GetClassificationOk() (*ClassificationReference, bool)`

GetClassificationOk returns a tuple with the Classification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassification

`func (o *ExpenseEntry) SetClassification(v ClassificationReference)`

SetClassification sets Classification field to given value.

### HasClassification

`func (o *ExpenseEntry) HasClassification() bool`

HasClassification returns a boolean if a field has been set.

### GetAmount

`func (o *ExpenseEntry) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ExpenseEntry) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ExpenseEntry) SetAmount(v float64)`

SetAmount sets Amount field to given value.


### SetAmountNil

`func (o *ExpenseEntry) SetAmountNil(b bool)`

 SetAmountNil sets the value for Amount to be an explicit nil

### UnsetAmount
`func (o *ExpenseEntry) UnsetAmount()`

UnsetAmount ensures that no value is present for Amount, not even an explicit nil
### GetBillableOption

`func (o *ExpenseEntry) GetBillableOption() string`

GetBillableOption returns the BillableOption field if non-nil, zero value otherwise.

### GetBillableOptionOk

`func (o *ExpenseEntry) GetBillableOptionOk() (*string, bool)`

GetBillableOptionOk returns a tuple with the BillableOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillableOption

`func (o *ExpenseEntry) SetBillableOption(v string)`

SetBillableOption sets BillableOption field to given value.

### HasBillableOption

`func (o *ExpenseEntry) HasBillableOption() bool`

HasBillableOption returns a boolean if a field has been set.

### SetBillableOptionNil

`func (o *ExpenseEntry) SetBillableOptionNil(b bool)`

 SetBillableOptionNil sets the value for BillableOption to be an explicit nil

### UnsetBillableOption
`func (o *ExpenseEntry) UnsetBillableOption()`

UnsetBillableOption ensures that no value is present for BillableOption, not even an explicit nil
### GetDate

`func (o *ExpenseEntry) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *ExpenseEntry) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *ExpenseEntry) SetDate(v time.Time)`

SetDate sets Date field to given value.


### GetLocationId

`func (o *ExpenseEntry) GetLocationId() int32`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *ExpenseEntry) GetLocationIdOk() (*int32, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *ExpenseEntry) SetLocationId(v int32)`

SetLocationId sets LocationId field to given value.

### HasLocationId

`func (o *ExpenseEntry) HasLocationId() bool`

HasLocationId returns a boolean if a field has been set.

### SetLocationIdNil

`func (o *ExpenseEntry) SetLocationIdNil(b bool)`

 SetLocationIdNil sets the value for LocationId to be an explicit nil

### UnsetLocationId
`func (o *ExpenseEntry) UnsetLocationId()`

UnsetLocationId ensures that no value is present for LocationId, not even an explicit nil
### GetBusinessUnitId

`func (o *ExpenseEntry) GetBusinessUnitId() int32`

GetBusinessUnitId returns the BusinessUnitId field if non-nil, zero value otherwise.

### GetBusinessUnitIdOk

`func (o *ExpenseEntry) GetBusinessUnitIdOk() (*int32, bool)`

GetBusinessUnitIdOk returns a tuple with the BusinessUnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessUnitId

`func (o *ExpenseEntry) SetBusinessUnitId(v int32)`

SetBusinessUnitId sets BusinessUnitId field to given value.

### HasBusinessUnitId

`func (o *ExpenseEntry) HasBusinessUnitId() bool`

HasBusinessUnitId returns a boolean if a field has been set.

### SetBusinessUnitIdNil

`func (o *ExpenseEntry) SetBusinessUnitIdNil(b bool)`

 SetBusinessUnitIdNil sets the value for BusinessUnitId to be an explicit nil

### UnsetBusinessUnitId
`func (o *ExpenseEntry) UnsetBusinessUnitId()`

UnsetBusinessUnitId ensures that no value is present for BusinessUnitId, not even an explicit nil
### GetNotes

`func (o *ExpenseEntry) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *ExpenseEntry) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *ExpenseEntry) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *ExpenseEntry) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetAgreement

`func (o *ExpenseEntry) GetAgreement() AgreementReference`

GetAgreement returns the Agreement field if non-nil, zero value otherwise.

### GetAgreementOk

`func (o *ExpenseEntry) GetAgreementOk() (*AgreementReference, bool)`

GetAgreementOk returns a tuple with the Agreement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreement

`func (o *ExpenseEntry) SetAgreement(v AgreementReference)`

SetAgreement sets Agreement field to given value.

### HasAgreement

`func (o *ExpenseEntry) HasAgreement() bool`

HasAgreement returns a boolean if a field has been set.

### GetInvoiceAmount

`func (o *ExpenseEntry) GetInvoiceAmount() float64`

GetInvoiceAmount returns the InvoiceAmount field if non-nil, zero value otherwise.

### GetInvoiceAmountOk

`func (o *ExpenseEntry) GetInvoiceAmountOk() (*float64, bool)`

GetInvoiceAmountOk returns a tuple with the InvoiceAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceAmount

`func (o *ExpenseEntry) SetInvoiceAmount(v float64)`

SetInvoiceAmount sets InvoiceAmount field to given value.

### HasInvoiceAmount

`func (o *ExpenseEntry) HasInvoiceAmount() bool`

HasInvoiceAmount returns a boolean if a field has been set.

### SetInvoiceAmountNil

`func (o *ExpenseEntry) SetInvoiceAmountNil(b bool)`

 SetInvoiceAmountNil sets the value for InvoiceAmount to be an explicit nil

### UnsetInvoiceAmount
`func (o *ExpenseEntry) UnsetInvoiceAmount()`

UnsetInvoiceAmount ensures that no value is present for InvoiceAmount, not even an explicit nil
### GetMobileGuid

`func (o *ExpenseEntry) GetMobileGuid() string`

GetMobileGuid returns the MobileGuid field if non-nil, zero value otherwise.

### GetMobileGuidOk

`func (o *ExpenseEntry) GetMobileGuidOk() (*string, bool)`

GetMobileGuidOk returns a tuple with the MobileGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileGuid

`func (o *ExpenseEntry) SetMobileGuid(v string)`

SetMobileGuid sets MobileGuid field to given value.

### HasMobileGuid

`func (o *ExpenseEntry) HasMobileGuid() bool`

HasMobileGuid returns a boolean if a field has been set.

### SetMobileGuidNil

`func (o *ExpenseEntry) SetMobileGuidNil(b bool)`

 SetMobileGuidNil sets the value for MobileGuid to be an explicit nil

### UnsetMobileGuid
`func (o *ExpenseEntry) UnsetMobileGuid()`

UnsetMobileGuid ensures that no value is present for MobileGuid, not even an explicit nil
### GetTaxes

`func (o *ExpenseEntry) GetTaxes() []ExpenseTax`

GetTaxes returns the Taxes field if non-nil, zero value otherwise.

### GetTaxesOk

`func (o *ExpenseEntry) GetTaxesOk() (*[]ExpenseTax, bool)`

GetTaxesOk returns a tuple with the Taxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxes

`func (o *ExpenseEntry) SetTaxes(v []ExpenseTax)`

SetTaxes sets Taxes field to given value.

### HasTaxes

`func (o *ExpenseEntry) HasTaxes() bool`

HasTaxes returns a boolean if a field has been set.

### GetInvoice

`func (o *ExpenseEntry) GetInvoice() InvoiceReference`

GetInvoice returns the Invoice field if non-nil, zero value otherwise.

### GetInvoiceOk

`func (o *ExpenseEntry) GetInvoiceOk() (*InvoiceReference, bool)`

GetInvoiceOk returns a tuple with the Invoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoice

`func (o *ExpenseEntry) SetInvoice(v InvoiceReference)`

SetInvoice sets Invoice field to given value.

### HasInvoice

`func (o *ExpenseEntry) HasInvoice() bool`

HasInvoice returns a boolean if a field has been set.

### GetCurrency

`func (o *ExpenseEntry) GetCurrency() CurrencyReference`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ExpenseEntry) GetCurrencyOk() (*CurrencyReference, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ExpenseEntry) SetCurrency(v CurrencyReference)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ExpenseEntry) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetStatus

`func (o *ExpenseEntry) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExpenseEntry) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExpenseEntry) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ExpenseEntry) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ExpenseEntry) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ExpenseEntry) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetBillAmount

`func (o *ExpenseEntry) GetBillAmount() float64`

GetBillAmount returns the BillAmount field if non-nil, zero value otherwise.

### GetBillAmountOk

`func (o *ExpenseEntry) GetBillAmountOk() (*float64, bool)`

GetBillAmountOk returns a tuple with the BillAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillAmount

`func (o *ExpenseEntry) SetBillAmount(v float64)`

SetBillAmount sets BillAmount field to given value.

### HasBillAmount

`func (o *ExpenseEntry) HasBillAmount() bool`

HasBillAmount returns a boolean if a field has been set.

### SetBillAmountNil

`func (o *ExpenseEntry) SetBillAmountNil(b bool)`

 SetBillAmountNil sets the value for BillAmount to be an explicit nil

### UnsetBillAmount
`func (o *ExpenseEntry) UnsetBillAmount()`

UnsetBillAmount ensures that no value is present for BillAmount, not even an explicit nil
### GetAgreementAmount

`func (o *ExpenseEntry) GetAgreementAmount() float64`

GetAgreementAmount returns the AgreementAmount field if non-nil, zero value otherwise.

### GetAgreementAmountOk

`func (o *ExpenseEntry) GetAgreementAmountOk() (*float64, bool)`

GetAgreementAmountOk returns a tuple with the AgreementAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreementAmount

`func (o *ExpenseEntry) SetAgreementAmount(v float64)`

SetAgreementAmount sets AgreementAmount field to given value.

### HasAgreementAmount

`func (o *ExpenseEntry) HasAgreementAmount() bool`

HasAgreementAmount returns a boolean if a field has been set.

### SetAgreementAmountNil

`func (o *ExpenseEntry) SetAgreementAmountNil(b bool)`

 SetAgreementAmountNil sets the value for AgreementAmount to be an explicit nil

### UnsetAgreementAmount
`func (o *ExpenseEntry) UnsetAgreementAmount()`

UnsetAgreementAmount ensures that no value is present for AgreementAmount, not even an explicit nil
### GetOdometerStart

`func (o *ExpenseEntry) GetOdometerStart() float64`

GetOdometerStart returns the OdometerStart field if non-nil, zero value otherwise.

### GetOdometerStartOk

`func (o *ExpenseEntry) GetOdometerStartOk() (*float64, bool)`

GetOdometerStartOk returns a tuple with the OdometerStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdometerStart

`func (o *ExpenseEntry) SetOdometerStart(v float64)`

SetOdometerStart sets OdometerStart field to given value.

### HasOdometerStart

`func (o *ExpenseEntry) HasOdometerStart() bool`

HasOdometerStart returns a boolean if a field has been set.

### SetOdometerStartNil

`func (o *ExpenseEntry) SetOdometerStartNil(b bool)`

 SetOdometerStartNil sets the value for OdometerStart to be an explicit nil

### UnsetOdometerStart
`func (o *ExpenseEntry) UnsetOdometerStart()`

UnsetOdometerStart ensures that no value is present for OdometerStart, not even an explicit nil
### GetOdometerEnd

`func (o *ExpenseEntry) GetOdometerEnd() float64`

GetOdometerEnd returns the OdometerEnd field if non-nil, zero value otherwise.

### GetOdometerEndOk

`func (o *ExpenseEntry) GetOdometerEndOk() (*float64, bool)`

GetOdometerEndOk returns a tuple with the OdometerEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdometerEnd

`func (o *ExpenseEntry) SetOdometerEnd(v float64)`

SetOdometerEnd sets OdometerEnd field to given value.

### HasOdometerEnd

`func (o *ExpenseEntry) HasOdometerEnd() bool`

HasOdometerEnd returns a boolean if a field has been set.

### SetOdometerEndNil

`func (o *ExpenseEntry) SetOdometerEndNil(b bool)`

 SetOdometerEndNil sets the value for OdometerEnd to be an explicit nil

### UnsetOdometerEnd
`func (o *ExpenseEntry) UnsetOdometerEnd()`

UnsetOdometerEnd ensures that no value is present for OdometerEnd, not even an explicit nil
### GetTicket

`func (o *ExpenseEntry) GetTicket() TicketReference`

GetTicket returns the Ticket field if non-nil, zero value otherwise.

### GetTicketOk

`func (o *ExpenseEntry) GetTicketOk() (*TicketReference, bool)`

GetTicketOk returns a tuple with the Ticket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicket

`func (o *ExpenseEntry) SetTicket(v TicketReference)`

SetTicket sets Ticket field to given value.

### HasTicket

`func (o *ExpenseEntry) HasTicket() bool`

HasTicket returns a boolean if a field has been set.

### GetProject

`func (o *ExpenseEntry) GetProject() ProjectReference`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ExpenseEntry) GetProjectOk() (*ProjectReference, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ExpenseEntry) SetProject(v ProjectReference)`

SetProject sets Project field to given value.

### HasProject

`func (o *ExpenseEntry) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetPhase

`func (o *ExpenseEntry) GetPhase() ProjectPhaseReference`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *ExpenseEntry) GetPhaseOk() (*ProjectPhaseReference, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *ExpenseEntry) SetPhase(v ProjectPhaseReference)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *ExpenseEntry) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetInfo

`func (o *ExpenseEntry) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ExpenseEntry) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ExpenseEntry) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ExpenseEntry) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetCustomFields

`func (o *ExpenseEntry) GetCustomFields() []CustomFieldValue`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ExpenseEntry) GetCustomFieldsOk() (*[]CustomFieldValue, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ExpenseEntry) SetCustomFields(v []CustomFieldValue)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ExpenseEntry) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


