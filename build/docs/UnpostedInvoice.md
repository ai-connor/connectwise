# UnpostedInvoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**BillingLogId** | Pointer to **NullableInt32** |  | [optional] 
**LocationId** | Pointer to **NullableInt32** |  | [optional] 
**Location** | Pointer to [**OwnerLevelReference**](OwnerLevelReference.md) |  | [optional] 
**DepartmentId** | Pointer to **NullableInt32** |  | [optional] 
**Department** | Pointer to [**SystemDepartmentReference**](SystemDepartmentReference.md) |  | [optional] 
**Company** | Pointer to [**CompanyReference**](CompanyReference.md) |  | [optional] 
**AccountNumber** | Pointer to **string** |  | [optional] 
**BillToCompany** | Pointer to [**CompanyReference**](CompanyReference.md) |  | [optional] 
**BillToSite** | Pointer to [**SiteReference**](SiteReference.md) |  | [optional] 
**ShipToCompany** | Pointer to [**CompanyReference**](CompanyReference.md) |  | [optional] 
**ShipToSite** | Pointer to [**SiteReference**](SiteReference.md) |  | [optional] 
**InvoiceNumber** | Pointer to **string** |  | [optional] 
**InvoiceDate** | Pointer to **string** |  | [optional] 
**InvoiceType** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**BillingTerms** | Pointer to [**BillingTermsReference**](BillingTermsReference.md) |  | [optional] 
**DueDays** | Pointer to **string** |  | [optional] 
**DueDate** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to [**CurrencyReference**](CurrencyReference.md) |  | [optional] 
**SubTotal** | Pointer to **NullableFloat64** |  | [optional] 
**Total** | Pointer to **NullableFloat64** |  | [optional] 
**HasTime** | Pointer to **bool** |  | [optional] 
**HasExpenses** | Pointer to **bool** |  | [optional] 
**HasProducts** | Pointer to **bool** |  | [optional] 
**InvoiceTaxableFlag** | Pointer to **NullableBool** |  | [optional] 
**TaxCode** | Pointer to [**TaxCodeReference**](TaxCodeReference.md) |  | [optional] 
**AvalaraTaxFlag** | Pointer to **NullableBool** | Used to determine if Avalara tax is enabled. | [optional] 
**ItemTaxableFlag** | Pointer to **NullableBool** |  | [optional] 
**SalesTaxAmount** | Pointer to **NullableFloat64** |  | [optional] 
**StateTaxFlag** | Pointer to **NullableBool** | Set to true if transaction is taxable at the state level. | [optional] 
**StateTaxXref** | Pointer to **string** |  | [optional] 
**StateTaxAmount** | Pointer to **NullableFloat64** |  | [optional] 
**CountyTaxFlag** | Pointer to **NullableBool** | Set to true if transaction is taxable at the county level. | [optional] 
**CountyTaxXref** | Pointer to **string** |  | [optional] 
**CountyTaxAmount** | Pointer to **NullableFloat64** |  | [optional] 
**CityTaxFlag** | Pointer to **NullableBool** | Set to true if transaction is taxable at the city level. | [optional] 
**CityTaxXref** | Pointer to **string** |  | [optional] 
**CityTaxAmount** | Pointer to **NullableFloat64** |  | [optional] 
**CountryTaxFlag** | Pointer to **NullableBool** | Set to true if transaction is taxable at the country level. | [optional] 
**CountryTaxXref** | Pointer to **string** |  | [optional] 
**CountryTaxAmount** | Pointer to **NullableFloat64** |  | [optional] 
**CompositeTaxFlag** | Pointer to **NullableBool** | Set to true if transaction is taxable at the composite level. | [optional] 
**CompositeTaxXref** | Pointer to **string** |  | [optional] 
**CompositeTaxAmount** | Pointer to **NullableFloat64** |  | [optional] 
**LevelSixTaxFlag** | Pointer to **NullableBool** | Set to true if transaction is taxable at level six. | [optional] 
**LevelSixTaxXref** | Pointer to **string** |  | [optional] 
**LevelSixTaxAmount** | Pointer to **NullableFloat64** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**DateClosed** | Pointer to **string** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewUnpostedInvoice

`func NewUnpostedInvoice() *UnpostedInvoice`

NewUnpostedInvoice instantiates a new UnpostedInvoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnpostedInvoiceWithDefaults

`func NewUnpostedInvoiceWithDefaults() *UnpostedInvoice`

NewUnpostedInvoiceWithDefaults instantiates a new UnpostedInvoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UnpostedInvoice) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnpostedInvoice) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnpostedInvoice) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *UnpostedInvoice) HasId() bool`

HasId returns a boolean if a field has been set.

### GetBillingLogId

`func (o *UnpostedInvoice) GetBillingLogId() int32`

GetBillingLogId returns the BillingLogId field if non-nil, zero value otherwise.

### GetBillingLogIdOk

`func (o *UnpostedInvoice) GetBillingLogIdOk() (*int32, bool)`

GetBillingLogIdOk returns a tuple with the BillingLogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingLogId

`func (o *UnpostedInvoice) SetBillingLogId(v int32)`

SetBillingLogId sets BillingLogId field to given value.

### HasBillingLogId

`func (o *UnpostedInvoice) HasBillingLogId() bool`

HasBillingLogId returns a boolean if a field has been set.

### SetBillingLogIdNil

`func (o *UnpostedInvoice) SetBillingLogIdNil(b bool)`

 SetBillingLogIdNil sets the value for BillingLogId to be an explicit nil

### UnsetBillingLogId
`func (o *UnpostedInvoice) UnsetBillingLogId()`

UnsetBillingLogId ensures that no value is present for BillingLogId, not even an explicit nil
### GetLocationId

`func (o *UnpostedInvoice) GetLocationId() int32`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *UnpostedInvoice) GetLocationIdOk() (*int32, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *UnpostedInvoice) SetLocationId(v int32)`

SetLocationId sets LocationId field to given value.

### HasLocationId

`func (o *UnpostedInvoice) HasLocationId() bool`

HasLocationId returns a boolean if a field has been set.

### SetLocationIdNil

`func (o *UnpostedInvoice) SetLocationIdNil(b bool)`

 SetLocationIdNil sets the value for LocationId to be an explicit nil

### UnsetLocationId
`func (o *UnpostedInvoice) UnsetLocationId()`

UnsetLocationId ensures that no value is present for LocationId, not even an explicit nil
### GetLocation

`func (o *UnpostedInvoice) GetLocation() OwnerLevelReference`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *UnpostedInvoice) GetLocationOk() (*OwnerLevelReference, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *UnpostedInvoice) SetLocation(v OwnerLevelReference)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *UnpostedInvoice) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetDepartmentId

`func (o *UnpostedInvoice) GetDepartmentId() int32`

GetDepartmentId returns the DepartmentId field if non-nil, zero value otherwise.

### GetDepartmentIdOk

`func (o *UnpostedInvoice) GetDepartmentIdOk() (*int32, bool)`

GetDepartmentIdOk returns a tuple with the DepartmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartmentId

`func (o *UnpostedInvoice) SetDepartmentId(v int32)`

SetDepartmentId sets DepartmentId field to given value.

### HasDepartmentId

`func (o *UnpostedInvoice) HasDepartmentId() bool`

HasDepartmentId returns a boolean if a field has been set.

### SetDepartmentIdNil

`func (o *UnpostedInvoice) SetDepartmentIdNil(b bool)`

 SetDepartmentIdNil sets the value for DepartmentId to be an explicit nil

### UnsetDepartmentId
`func (o *UnpostedInvoice) UnsetDepartmentId()`

UnsetDepartmentId ensures that no value is present for DepartmentId, not even an explicit nil
### GetDepartment

`func (o *UnpostedInvoice) GetDepartment() SystemDepartmentReference`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *UnpostedInvoice) GetDepartmentOk() (*SystemDepartmentReference, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *UnpostedInvoice) SetDepartment(v SystemDepartmentReference)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *UnpostedInvoice) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### GetCompany

`func (o *UnpostedInvoice) GetCompany() CompanyReference`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *UnpostedInvoice) GetCompanyOk() (*CompanyReference, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *UnpostedInvoice) SetCompany(v CompanyReference)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *UnpostedInvoice) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetAccountNumber

`func (o *UnpostedInvoice) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *UnpostedInvoice) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *UnpostedInvoice) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *UnpostedInvoice) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetBillToCompany

`func (o *UnpostedInvoice) GetBillToCompany() CompanyReference`

GetBillToCompany returns the BillToCompany field if non-nil, zero value otherwise.

### GetBillToCompanyOk

`func (o *UnpostedInvoice) GetBillToCompanyOk() (*CompanyReference, bool)`

GetBillToCompanyOk returns a tuple with the BillToCompany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillToCompany

`func (o *UnpostedInvoice) SetBillToCompany(v CompanyReference)`

SetBillToCompany sets BillToCompany field to given value.

### HasBillToCompany

`func (o *UnpostedInvoice) HasBillToCompany() bool`

HasBillToCompany returns a boolean if a field has been set.

### GetBillToSite

`func (o *UnpostedInvoice) GetBillToSite() SiteReference`

GetBillToSite returns the BillToSite field if non-nil, zero value otherwise.

### GetBillToSiteOk

`func (o *UnpostedInvoice) GetBillToSiteOk() (*SiteReference, bool)`

GetBillToSiteOk returns a tuple with the BillToSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillToSite

`func (o *UnpostedInvoice) SetBillToSite(v SiteReference)`

SetBillToSite sets BillToSite field to given value.

### HasBillToSite

`func (o *UnpostedInvoice) HasBillToSite() bool`

HasBillToSite returns a boolean if a field has been set.

### GetShipToCompany

`func (o *UnpostedInvoice) GetShipToCompany() CompanyReference`

GetShipToCompany returns the ShipToCompany field if non-nil, zero value otherwise.

### GetShipToCompanyOk

`func (o *UnpostedInvoice) GetShipToCompanyOk() (*CompanyReference, bool)`

GetShipToCompanyOk returns a tuple with the ShipToCompany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipToCompany

`func (o *UnpostedInvoice) SetShipToCompany(v CompanyReference)`

SetShipToCompany sets ShipToCompany field to given value.

### HasShipToCompany

`func (o *UnpostedInvoice) HasShipToCompany() bool`

HasShipToCompany returns a boolean if a field has been set.

### GetShipToSite

`func (o *UnpostedInvoice) GetShipToSite() SiteReference`

GetShipToSite returns the ShipToSite field if non-nil, zero value otherwise.

### GetShipToSiteOk

`func (o *UnpostedInvoice) GetShipToSiteOk() (*SiteReference, bool)`

GetShipToSiteOk returns a tuple with the ShipToSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipToSite

`func (o *UnpostedInvoice) SetShipToSite(v SiteReference)`

SetShipToSite sets ShipToSite field to given value.

### HasShipToSite

`func (o *UnpostedInvoice) HasShipToSite() bool`

HasShipToSite returns a boolean if a field has been set.

### GetInvoiceNumber

`func (o *UnpostedInvoice) GetInvoiceNumber() string`

GetInvoiceNumber returns the InvoiceNumber field if non-nil, zero value otherwise.

### GetInvoiceNumberOk

`func (o *UnpostedInvoice) GetInvoiceNumberOk() (*string, bool)`

GetInvoiceNumberOk returns a tuple with the InvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceNumber

`func (o *UnpostedInvoice) SetInvoiceNumber(v string)`

SetInvoiceNumber sets InvoiceNumber field to given value.

### HasInvoiceNumber

`func (o *UnpostedInvoice) HasInvoiceNumber() bool`

HasInvoiceNumber returns a boolean if a field has been set.

### GetInvoiceDate

`func (o *UnpostedInvoice) GetInvoiceDate() string`

GetInvoiceDate returns the InvoiceDate field if non-nil, zero value otherwise.

### GetInvoiceDateOk

`func (o *UnpostedInvoice) GetInvoiceDateOk() (*string, bool)`

GetInvoiceDateOk returns a tuple with the InvoiceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceDate

`func (o *UnpostedInvoice) SetInvoiceDate(v string)`

SetInvoiceDate sets InvoiceDate field to given value.

### HasInvoiceDate

`func (o *UnpostedInvoice) HasInvoiceDate() bool`

HasInvoiceDate returns a boolean if a field has been set.

### GetInvoiceType

`func (o *UnpostedInvoice) GetInvoiceType() string`

GetInvoiceType returns the InvoiceType field if non-nil, zero value otherwise.

### GetInvoiceTypeOk

`func (o *UnpostedInvoice) GetInvoiceTypeOk() (*string, bool)`

GetInvoiceTypeOk returns a tuple with the InvoiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceType

`func (o *UnpostedInvoice) SetInvoiceType(v string)`

SetInvoiceType sets InvoiceType field to given value.

### HasInvoiceType

`func (o *UnpostedInvoice) HasInvoiceType() bool`

HasInvoiceType returns a boolean if a field has been set.

### SetInvoiceTypeNil

`func (o *UnpostedInvoice) SetInvoiceTypeNil(b bool)`

 SetInvoiceTypeNil sets the value for InvoiceType to be an explicit nil

### UnsetInvoiceType
`func (o *UnpostedInvoice) UnsetInvoiceType()`

UnsetInvoiceType ensures that no value is present for InvoiceType, not even an explicit nil
### GetDescription

`func (o *UnpostedInvoice) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UnpostedInvoice) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UnpostedInvoice) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UnpostedInvoice) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetBillingTerms

`func (o *UnpostedInvoice) GetBillingTerms() BillingTermsReference`

GetBillingTerms returns the BillingTerms field if non-nil, zero value otherwise.

### GetBillingTermsOk

`func (o *UnpostedInvoice) GetBillingTermsOk() (*BillingTermsReference, bool)`

GetBillingTermsOk returns a tuple with the BillingTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingTerms

`func (o *UnpostedInvoice) SetBillingTerms(v BillingTermsReference)`

SetBillingTerms sets BillingTerms field to given value.

### HasBillingTerms

`func (o *UnpostedInvoice) HasBillingTerms() bool`

HasBillingTerms returns a boolean if a field has been set.

### GetDueDays

`func (o *UnpostedInvoice) GetDueDays() string`

GetDueDays returns the DueDays field if non-nil, zero value otherwise.

### GetDueDaysOk

`func (o *UnpostedInvoice) GetDueDaysOk() (*string, bool)`

GetDueDaysOk returns a tuple with the DueDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDays

`func (o *UnpostedInvoice) SetDueDays(v string)`

SetDueDays sets DueDays field to given value.

### HasDueDays

`func (o *UnpostedInvoice) HasDueDays() bool`

HasDueDays returns a boolean if a field has been set.

### GetDueDate

`func (o *UnpostedInvoice) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *UnpostedInvoice) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *UnpostedInvoice) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *UnpostedInvoice) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetCurrency

`func (o *UnpostedInvoice) GetCurrency() CurrencyReference`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnpostedInvoice) GetCurrencyOk() (*CurrencyReference, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnpostedInvoice) SetCurrency(v CurrencyReference)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnpostedInvoice) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetSubTotal

`func (o *UnpostedInvoice) GetSubTotal() float64`

GetSubTotal returns the SubTotal field if non-nil, zero value otherwise.

### GetSubTotalOk

`func (o *UnpostedInvoice) GetSubTotalOk() (*float64, bool)`

GetSubTotalOk returns a tuple with the SubTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubTotal

`func (o *UnpostedInvoice) SetSubTotal(v float64)`

SetSubTotal sets SubTotal field to given value.

### HasSubTotal

`func (o *UnpostedInvoice) HasSubTotal() bool`

HasSubTotal returns a boolean if a field has been set.

### SetSubTotalNil

`func (o *UnpostedInvoice) SetSubTotalNil(b bool)`

 SetSubTotalNil sets the value for SubTotal to be an explicit nil

### UnsetSubTotal
`func (o *UnpostedInvoice) UnsetSubTotal()`

UnsetSubTotal ensures that no value is present for SubTotal, not even an explicit nil
### GetTotal

`func (o *UnpostedInvoice) GetTotal() float64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *UnpostedInvoice) GetTotalOk() (*float64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *UnpostedInvoice) SetTotal(v float64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *UnpostedInvoice) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### SetTotalNil

`func (o *UnpostedInvoice) SetTotalNil(b bool)`

 SetTotalNil sets the value for Total to be an explicit nil

### UnsetTotal
`func (o *UnpostedInvoice) UnsetTotal()`

UnsetTotal ensures that no value is present for Total, not even an explicit nil
### GetHasTime

`func (o *UnpostedInvoice) GetHasTime() bool`

GetHasTime returns the HasTime field if non-nil, zero value otherwise.

### GetHasTimeOk

`func (o *UnpostedInvoice) GetHasTimeOk() (*bool, bool)`

GetHasTimeOk returns a tuple with the HasTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasTime

`func (o *UnpostedInvoice) SetHasTime(v bool)`

SetHasTime sets HasTime field to given value.

### HasHasTime

`func (o *UnpostedInvoice) HasHasTime() bool`

HasHasTime returns a boolean if a field has been set.

### GetHasExpenses

`func (o *UnpostedInvoice) GetHasExpenses() bool`

GetHasExpenses returns the HasExpenses field if non-nil, zero value otherwise.

### GetHasExpensesOk

`func (o *UnpostedInvoice) GetHasExpensesOk() (*bool, bool)`

GetHasExpensesOk returns a tuple with the HasExpenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasExpenses

`func (o *UnpostedInvoice) SetHasExpenses(v bool)`

SetHasExpenses sets HasExpenses field to given value.

### HasHasExpenses

`func (o *UnpostedInvoice) HasHasExpenses() bool`

HasHasExpenses returns a boolean if a field has been set.

### GetHasProducts

`func (o *UnpostedInvoice) GetHasProducts() bool`

GetHasProducts returns the HasProducts field if non-nil, zero value otherwise.

### GetHasProductsOk

`func (o *UnpostedInvoice) GetHasProductsOk() (*bool, bool)`

GetHasProductsOk returns a tuple with the HasProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasProducts

`func (o *UnpostedInvoice) SetHasProducts(v bool)`

SetHasProducts sets HasProducts field to given value.

### HasHasProducts

`func (o *UnpostedInvoice) HasHasProducts() bool`

HasHasProducts returns a boolean if a field has been set.

### GetInvoiceTaxableFlag

`func (o *UnpostedInvoice) GetInvoiceTaxableFlag() bool`

GetInvoiceTaxableFlag returns the InvoiceTaxableFlag field if non-nil, zero value otherwise.

### GetInvoiceTaxableFlagOk

`func (o *UnpostedInvoice) GetInvoiceTaxableFlagOk() (*bool, bool)`

GetInvoiceTaxableFlagOk returns a tuple with the InvoiceTaxableFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceTaxableFlag

`func (o *UnpostedInvoice) SetInvoiceTaxableFlag(v bool)`

SetInvoiceTaxableFlag sets InvoiceTaxableFlag field to given value.

### HasInvoiceTaxableFlag

`func (o *UnpostedInvoice) HasInvoiceTaxableFlag() bool`

HasInvoiceTaxableFlag returns a boolean if a field has been set.

### SetInvoiceTaxableFlagNil

`func (o *UnpostedInvoice) SetInvoiceTaxableFlagNil(b bool)`

 SetInvoiceTaxableFlagNil sets the value for InvoiceTaxableFlag to be an explicit nil

### UnsetInvoiceTaxableFlag
`func (o *UnpostedInvoice) UnsetInvoiceTaxableFlag()`

UnsetInvoiceTaxableFlag ensures that no value is present for InvoiceTaxableFlag, not even an explicit nil
### GetTaxCode

`func (o *UnpostedInvoice) GetTaxCode() TaxCodeReference`

GetTaxCode returns the TaxCode field if non-nil, zero value otherwise.

### GetTaxCodeOk

`func (o *UnpostedInvoice) GetTaxCodeOk() (*TaxCodeReference, bool)`

GetTaxCodeOk returns a tuple with the TaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCode

`func (o *UnpostedInvoice) SetTaxCode(v TaxCodeReference)`

SetTaxCode sets TaxCode field to given value.

### HasTaxCode

`func (o *UnpostedInvoice) HasTaxCode() bool`

HasTaxCode returns a boolean if a field has been set.

### GetAvalaraTaxFlag

`func (o *UnpostedInvoice) GetAvalaraTaxFlag() bool`

GetAvalaraTaxFlag returns the AvalaraTaxFlag field if non-nil, zero value otherwise.

### GetAvalaraTaxFlagOk

`func (o *UnpostedInvoice) GetAvalaraTaxFlagOk() (*bool, bool)`

GetAvalaraTaxFlagOk returns a tuple with the AvalaraTaxFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvalaraTaxFlag

`func (o *UnpostedInvoice) SetAvalaraTaxFlag(v bool)`

SetAvalaraTaxFlag sets AvalaraTaxFlag field to given value.

### HasAvalaraTaxFlag

`func (o *UnpostedInvoice) HasAvalaraTaxFlag() bool`

HasAvalaraTaxFlag returns a boolean if a field has been set.

### SetAvalaraTaxFlagNil

`func (o *UnpostedInvoice) SetAvalaraTaxFlagNil(b bool)`

 SetAvalaraTaxFlagNil sets the value for AvalaraTaxFlag to be an explicit nil

### UnsetAvalaraTaxFlag
`func (o *UnpostedInvoice) UnsetAvalaraTaxFlag()`

UnsetAvalaraTaxFlag ensures that no value is present for AvalaraTaxFlag, not even an explicit nil
### GetItemTaxableFlag

`func (o *UnpostedInvoice) GetItemTaxableFlag() bool`

GetItemTaxableFlag returns the ItemTaxableFlag field if non-nil, zero value otherwise.

### GetItemTaxableFlagOk

`func (o *UnpostedInvoice) GetItemTaxableFlagOk() (*bool, bool)`

GetItemTaxableFlagOk returns a tuple with the ItemTaxableFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemTaxableFlag

`func (o *UnpostedInvoice) SetItemTaxableFlag(v bool)`

SetItemTaxableFlag sets ItemTaxableFlag field to given value.

### HasItemTaxableFlag

`func (o *UnpostedInvoice) HasItemTaxableFlag() bool`

HasItemTaxableFlag returns a boolean if a field has been set.

### SetItemTaxableFlagNil

`func (o *UnpostedInvoice) SetItemTaxableFlagNil(b bool)`

 SetItemTaxableFlagNil sets the value for ItemTaxableFlag to be an explicit nil

### UnsetItemTaxableFlag
`func (o *UnpostedInvoice) UnsetItemTaxableFlag()`

UnsetItemTaxableFlag ensures that no value is present for ItemTaxableFlag, not even an explicit nil
### GetSalesTaxAmount

`func (o *UnpostedInvoice) GetSalesTaxAmount() float64`

GetSalesTaxAmount returns the SalesTaxAmount field if non-nil, zero value otherwise.

### GetSalesTaxAmountOk

`func (o *UnpostedInvoice) GetSalesTaxAmountOk() (*float64, bool)`

GetSalesTaxAmountOk returns a tuple with the SalesTaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxAmount

`func (o *UnpostedInvoice) SetSalesTaxAmount(v float64)`

SetSalesTaxAmount sets SalesTaxAmount field to given value.

### HasSalesTaxAmount

`func (o *UnpostedInvoice) HasSalesTaxAmount() bool`

HasSalesTaxAmount returns a boolean if a field has been set.

### SetSalesTaxAmountNil

`func (o *UnpostedInvoice) SetSalesTaxAmountNil(b bool)`

 SetSalesTaxAmountNil sets the value for SalesTaxAmount to be an explicit nil

### UnsetSalesTaxAmount
`func (o *UnpostedInvoice) UnsetSalesTaxAmount()`

UnsetSalesTaxAmount ensures that no value is present for SalesTaxAmount, not even an explicit nil
### GetStateTaxFlag

`func (o *UnpostedInvoice) GetStateTaxFlag() bool`

GetStateTaxFlag returns the StateTaxFlag field if non-nil, zero value otherwise.

### GetStateTaxFlagOk

`func (o *UnpostedInvoice) GetStateTaxFlagOk() (*bool, bool)`

GetStateTaxFlagOk returns a tuple with the StateTaxFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateTaxFlag

`func (o *UnpostedInvoice) SetStateTaxFlag(v bool)`

SetStateTaxFlag sets StateTaxFlag field to given value.

### HasStateTaxFlag

`func (o *UnpostedInvoice) HasStateTaxFlag() bool`

HasStateTaxFlag returns a boolean if a field has been set.

### SetStateTaxFlagNil

`func (o *UnpostedInvoice) SetStateTaxFlagNil(b bool)`

 SetStateTaxFlagNil sets the value for StateTaxFlag to be an explicit nil

### UnsetStateTaxFlag
`func (o *UnpostedInvoice) UnsetStateTaxFlag()`

UnsetStateTaxFlag ensures that no value is present for StateTaxFlag, not even an explicit nil
### GetStateTaxXref

`func (o *UnpostedInvoice) GetStateTaxXref() string`

GetStateTaxXref returns the StateTaxXref field if non-nil, zero value otherwise.

### GetStateTaxXrefOk

`func (o *UnpostedInvoice) GetStateTaxXrefOk() (*string, bool)`

GetStateTaxXrefOk returns a tuple with the StateTaxXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateTaxXref

`func (o *UnpostedInvoice) SetStateTaxXref(v string)`

SetStateTaxXref sets StateTaxXref field to given value.

### HasStateTaxXref

`func (o *UnpostedInvoice) HasStateTaxXref() bool`

HasStateTaxXref returns a boolean if a field has been set.

### GetStateTaxAmount

`func (o *UnpostedInvoice) GetStateTaxAmount() float64`

GetStateTaxAmount returns the StateTaxAmount field if non-nil, zero value otherwise.

### GetStateTaxAmountOk

`func (o *UnpostedInvoice) GetStateTaxAmountOk() (*float64, bool)`

GetStateTaxAmountOk returns a tuple with the StateTaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateTaxAmount

`func (o *UnpostedInvoice) SetStateTaxAmount(v float64)`

SetStateTaxAmount sets StateTaxAmount field to given value.

### HasStateTaxAmount

`func (o *UnpostedInvoice) HasStateTaxAmount() bool`

HasStateTaxAmount returns a boolean if a field has been set.

### SetStateTaxAmountNil

`func (o *UnpostedInvoice) SetStateTaxAmountNil(b bool)`

 SetStateTaxAmountNil sets the value for StateTaxAmount to be an explicit nil

### UnsetStateTaxAmount
`func (o *UnpostedInvoice) UnsetStateTaxAmount()`

UnsetStateTaxAmount ensures that no value is present for StateTaxAmount, not even an explicit nil
### GetCountyTaxFlag

`func (o *UnpostedInvoice) GetCountyTaxFlag() bool`

GetCountyTaxFlag returns the CountyTaxFlag field if non-nil, zero value otherwise.

### GetCountyTaxFlagOk

`func (o *UnpostedInvoice) GetCountyTaxFlagOk() (*bool, bool)`

GetCountyTaxFlagOk returns a tuple with the CountyTaxFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountyTaxFlag

`func (o *UnpostedInvoice) SetCountyTaxFlag(v bool)`

SetCountyTaxFlag sets CountyTaxFlag field to given value.

### HasCountyTaxFlag

`func (o *UnpostedInvoice) HasCountyTaxFlag() bool`

HasCountyTaxFlag returns a boolean if a field has been set.

### SetCountyTaxFlagNil

`func (o *UnpostedInvoice) SetCountyTaxFlagNil(b bool)`

 SetCountyTaxFlagNil sets the value for CountyTaxFlag to be an explicit nil

### UnsetCountyTaxFlag
`func (o *UnpostedInvoice) UnsetCountyTaxFlag()`

UnsetCountyTaxFlag ensures that no value is present for CountyTaxFlag, not even an explicit nil
### GetCountyTaxXref

`func (o *UnpostedInvoice) GetCountyTaxXref() string`

GetCountyTaxXref returns the CountyTaxXref field if non-nil, zero value otherwise.

### GetCountyTaxXrefOk

`func (o *UnpostedInvoice) GetCountyTaxXrefOk() (*string, bool)`

GetCountyTaxXrefOk returns a tuple with the CountyTaxXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountyTaxXref

`func (o *UnpostedInvoice) SetCountyTaxXref(v string)`

SetCountyTaxXref sets CountyTaxXref field to given value.

### HasCountyTaxXref

`func (o *UnpostedInvoice) HasCountyTaxXref() bool`

HasCountyTaxXref returns a boolean if a field has been set.

### GetCountyTaxAmount

`func (o *UnpostedInvoice) GetCountyTaxAmount() float64`

GetCountyTaxAmount returns the CountyTaxAmount field if non-nil, zero value otherwise.

### GetCountyTaxAmountOk

`func (o *UnpostedInvoice) GetCountyTaxAmountOk() (*float64, bool)`

GetCountyTaxAmountOk returns a tuple with the CountyTaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountyTaxAmount

`func (o *UnpostedInvoice) SetCountyTaxAmount(v float64)`

SetCountyTaxAmount sets CountyTaxAmount field to given value.

### HasCountyTaxAmount

`func (o *UnpostedInvoice) HasCountyTaxAmount() bool`

HasCountyTaxAmount returns a boolean if a field has been set.

### SetCountyTaxAmountNil

`func (o *UnpostedInvoice) SetCountyTaxAmountNil(b bool)`

 SetCountyTaxAmountNil sets the value for CountyTaxAmount to be an explicit nil

### UnsetCountyTaxAmount
`func (o *UnpostedInvoice) UnsetCountyTaxAmount()`

UnsetCountyTaxAmount ensures that no value is present for CountyTaxAmount, not even an explicit nil
### GetCityTaxFlag

`func (o *UnpostedInvoice) GetCityTaxFlag() bool`

GetCityTaxFlag returns the CityTaxFlag field if non-nil, zero value otherwise.

### GetCityTaxFlagOk

`func (o *UnpostedInvoice) GetCityTaxFlagOk() (*bool, bool)`

GetCityTaxFlagOk returns a tuple with the CityTaxFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityTaxFlag

`func (o *UnpostedInvoice) SetCityTaxFlag(v bool)`

SetCityTaxFlag sets CityTaxFlag field to given value.

### HasCityTaxFlag

`func (o *UnpostedInvoice) HasCityTaxFlag() bool`

HasCityTaxFlag returns a boolean if a field has been set.

### SetCityTaxFlagNil

`func (o *UnpostedInvoice) SetCityTaxFlagNil(b bool)`

 SetCityTaxFlagNil sets the value for CityTaxFlag to be an explicit nil

### UnsetCityTaxFlag
`func (o *UnpostedInvoice) UnsetCityTaxFlag()`

UnsetCityTaxFlag ensures that no value is present for CityTaxFlag, not even an explicit nil
### GetCityTaxXref

`func (o *UnpostedInvoice) GetCityTaxXref() string`

GetCityTaxXref returns the CityTaxXref field if non-nil, zero value otherwise.

### GetCityTaxXrefOk

`func (o *UnpostedInvoice) GetCityTaxXrefOk() (*string, bool)`

GetCityTaxXrefOk returns a tuple with the CityTaxXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityTaxXref

`func (o *UnpostedInvoice) SetCityTaxXref(v string)`

SetCityTaxXref sets CityTaxXref field to given value.

### HasCityTaxXref

`func (o *UnpostedInvoice) HasCityTaxXref() bool`

HasCityTaxXref returns a boolean if a field has been set.

### GetCityTaxAmount

`func (o *UnpostedInvoice) GetCityTaxAmount() float64`

GetCityTaxAmount returns the CityTaxAmount field if non-nil, zero value otherwise.

### GetCityTaxAmountOk

`func (o *UnpostedInvoice) GetCityTaxAmountOk() (*float64, bool)`

GetCityTaxAmountOk returns a tuple with the CityTaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityTaxAmount

`func (o *UnpostedInvoice) SetCityTaxAmount(v float64)`

SetCityTaxAmount sets CityTaxAmount field to given value.

### HasCityTaxAmount

`func (o *UnpostedInvoice) HasCityTaxAmount() bool`

HasCityTaxAmount returns a boolean if a field has been set.

### SetCityTaxAmountNil

`func (o *UnpostedInvoice) SetCityTaxAmountNil(b bool)`

 SetCityTaxAmountNil sets the value for CityTaxAmount to be an explicit nil

### UnsetCityTaxAmount
`func (o *UnpostedInvoice) UnsetCityTaxAmount()`

UnsetCityTaxAmount ensures that no value is present for CityTaxAmount, not even an explicit nil
### GetCountryTaxFlag

`func (o *UnpostedInvoice) GetCountryTaxFlag() bool`

GetCountryTaxFlag returns the CountryTaxFlag field if non-nil, zero value otherwise.

### GetCountryTaxFlagOk

`func (o *UnpostedInvoice) GetCountryTaxFlagOk() (*bool, bool)`

GetCountryTaxFlagOk returns a tuple with the CountryTaxFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryTaxFlag

`func (o *UnpostedInvoice) SetCountryTaxFlag(v bool)`

SetCountryTaxFlag sets CountryTaxFlag field to given value.

### HasCountryTaxFlag

`func (o *UnpostedInvoice) HasCountryTaxFlag() bool`

HasCountryTaxFlag returns a boolean if a field has been set.

### SetCountryTaxFlagNil

`func (o *UnpostedInvoice) SetCountryTaxFlagNil(b bool)`

 SetCountryTaxFlagNil sets the value for CountryTaxFlag to be an explicit nil

### UnsetCountryTaxFlag
`func (o *UnpostedInvoice) UnsetCountryTaxFlag()`

UnsetCountryTaxFlag ensures that no value is present for CountryTaxFlag, not even an explicit nil
### GetCountryTaxXref

`func (o *UnpostedInvoice) GetCountryTaxXref() string`

GetCountryTaxXref returns the CountryTaxXref field if non-nil, zero value otherwise.

### GetCountryTaxXrefOk

`func (o *UnpostedInvoice) GetCountryTaxXrefOk() (*string, bool)`

GetCountryTaxXrefOk returns a tuple with the CountryTaxXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryTaxXref

`func (o *UnpostedInvoice) SetCountryTaxXref(v string)`

SetCountryTaxXref sets CountryTaxXref field to given value.

### HasCountryTaxXref

`func (o *UnpostedInvoice) HasCountryTaxXref() bool`

HasCountryTaxXref returns a boolean if a field has been set.

### GetCountryTaxAmount

`func (o *UnpostedInvoice) GetCountryTaxAmount() float64`

GetCountryTaxAmount returns the CountryTaxAmount field if non-nil, zero value otherwise.

### GetCountryTaxAmountOk

`func (o *UnpostedInvoice) GetCountryTaxAmountOk() (*float64, bool)`

GetCountryTaxAmountOk returns a tuple with the CountryTaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryTaxAmount

`func (o *UnpostedInvoice) SetCountryTaxAmount(v float64)`

SetCountryTaxAmount sets CountryTaxAmount field to given value.

### HasCountryTaxAmount

`func (o *UnpostedInvoice) HasCountryTaxAmount() bool`

HasCountryTaxAmount returns a boolean if a field has been set.

### SetCountryTaxAmountNil

`func (o *UnpostedInvoice) SetCountryTaxAmountNil(b bool)`

 SetCountryTaxAmountNil sets the value for CountryTaxAmount to be an explicit nil

### UnsetCountryTaxAmount
`func (o *UnpostedInvoice) UnsetCountryTaxAmount()`

UnsetCountryTaxAmount ensures that no value is present for CountryTaxAmount, not even an explicit nil
### GetCompositeTaxFlag

`func (o *UnpostedInvoice) GetCompositeTaxFlag() bool`

GetCompositeTaxFlag returns the CompositeTaxFlag field if non-nil, zero value otherwise.

### GetCompositeTaxFlagOk

`func (o *UnpostedInvoice) GetCompositeTaxFlagOk() (*bool, bool)`

GetCompositeTaxFlagOk returns a tuple with the CompositeTaxFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompositeTaxFlag

`func (o *UnpostedInvoice) SetCompositeTaxFlag(v bool)`

SetCompositeTaxFlag sets CompositeTaxFlag field to given value.

### HasCompositeTaxFlag

`func (o *UnpostedInvoice) HasCompositeTaxFlag() bool`

HasCompositeTaxFlag returns a boolean if a field has been set.

### SetCompositeTaxFlagNil

`func (o *UnpostedInvoice) SetCompositeTaxFlagNil(b bool)`

 SetCompositeTaxFlagNil sets the value for CompositeTaxFlag to be an explicit nil

### UnsetCompositeTaxFlag
`func (o *UnpostedInvoice) UnsetCompositeTaxFlag()`

UnsetCompositeTaxFlag ensures that no value is present for CompositeTaxFlag, not even an explicit nil
### GetCompositeTaxXref

`func (o *UnpostedInvoice) GetCompositeTaxXref() string`

GetCompositeTaxXref returns the CompositeTaxXref field if non-nil, zero value otherwise.

### GetCompositeTaxXrefOk

`func (o *UnpostedInvoice) GetCompositeTaxXrefOk() (*string, bool)`

GetCompositeTaxXrefOk returns a tuple with the CompositeTaxXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompositeTaxXref

`func (o *UnpostedInvoice) SetCompositeTaxXref(v string)`

SetCompositeTaxXref sets CompositeTaxXref field to given value.

### HasCompositeTaxXref

`func (o *UnpostedInvoice) HasCompositeTaxXref() bool`

HasCompositeTaxXref returns a boolean if a field has been set.

### GetCompositeTaxAmount

`func (o *UnpostedInvoice) GetCompositeTaxAmount() float64`

GetCompositeTaxAmount returns the CompositeTaxAmount field if non-nil, zero value otherwise.

### GetCompositeTaxAmountOk

`func (o *UnpostedInvoice) GetCompositeTaxAmountOk() (*float64, bool)`

GetCompositeTaxAmountOk returns a tuple with the CompositeTaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompositeTaxAmount

`func (o *UnpostedInvoice) SetCompositeTaxAmount(v float64)`

SetCompositeTaxAmount sets CompositeTaxAmount field to given value.

### HasCompositeTaxAmount

`func (o *UnpostedInvoice) HasCompositeTaxAmount() bool`

HasCompositeTaxAmount returns a boolean if a field has been set.

### SetCompositeTaxAmountNil

`func (o *UnpostedInvoice) SetCompositeTaxAmountNil(b bool)`

 SetCompositeTaxAmountNil sets the value for CompositeTaxAmount to be an explicit nil

### UnsetCompositeTaxAmount
`func (o *UnpostedInvoice) UnsetCompositeTaxAmount()`

UnsetCompositeTaxAmount ensures that no value is present for CompositeTaxAmount, not even an explicit nil
### GetLevelSixTaxFlag

`func (o *UnpostedInvoice) GetLevelSixTaxFlag() bool`

GetLevelSixTaxFlag returns the LevelSixTaxFlag field if non-nil, zero value otherwise.

### GetLevelSixTaxFlagOk

`func (o *UnpostedInvoice) GetLevelSixTaxFlagOk() (*bool, bool)`

GetLevelSixTaxFlagOk returns a tuple with the LevelSixTaxFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelSixTaxFlag

`func (o *UnpostedInvoice) SetLevelSixTaxFlag(v bool)`

SetLevelSixTaxFlag sets LevelSixTaxFlag field to given value.

### HasLevelSixTaxFlag

`func (o *UnpostedInvoice) HasLevelSixTaxFlag() bool`

HasLevelSixTaxFlag returns a boolean if a field has been set.

### SetLevelSixTaxFlagNil

`func (o *UnpostedInvoice) SetLevelSixTaxFlagNil(b bool)`

 SetLevelSixTaxFlagNil sets the value for LevelSixTaxFlag to be an explicit nil

### UnsetLevelSixTaxFlag
`func (o *UnpostedInvoice) UnsetLevelSixTaxFlag()`

UnsetLevelSixTaxFlag ensures that no value is present for LevelSixTaxFlag, not even an explicit nil
### GetLevelSixTaxXref

`func (o *UnpostedInvoice) GetLevelSixTaxXref() string`

GetLevelSixTaxXref returns the LevelSixTaxXref field if non-nil, zero value otherwise.

### GetLevelSixTaxXrefOk

`func (o *UnpostedInvoice) GetLevelSixTaxXrefOk() (*string, bool)`

GetLevelSixTaxXrefOk returns a tuple with the LevelSixTaxXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelSixTaxXref

`func (o *UnpostedInvoice) SetLevelSixTaxXref(v string)`

SetLevelSixTaxXref sets LevelSixTaxXref field to given value.

### HasLevelSixTaxXref

`func (o *UnpostedInvoice) HasLevelSixTaxXref() bool`

HasLevelSixTaxXref returns a boolean if a field has been set.

### GetLevelSixTaxAmount

`func (o *UnpostedInvoice) GetLevelSixTaxAmount() float64`

GetLevelSixTaxAmount returns the LevelSixTaxAmount field if non-nil, zero value otherwise.

### GetLevelSixTaxAmountOk

`func (o *UnpostedInvoice) GetLevelSixTaxAmountOk() (*float64, bool)`

GetLevelSixTaxAmountOk returns a tuple with the LevelSixTaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelSixTaxAmount

`func (o *UnpostedInvoice) SetLevelSixTaxAmount(v float64)`

SetLevelSixTaxAmount sets LevelSixTaxAmount field to given value.

### HasLevelSixTaxAmount

`func (o *UnpostedInvoice) HasLevelSixTaxAmount() bool`

HasLevelSixTaxAmount returns a boolean if a field has been set.

### SetLevelSixTaxAmountNil

`func (o *UnpostedInvoice) SetLevelSixTaxAmountNil(b bool)`

 SetLevelSixTaxAmountNil sets the value for LevelSixTaxAmount to be an explicit nil

### UnsetLevelSixTaxAmount
`func (o *UnpostedInvoice) UnsetLevelSixTaxAmount()`

UnsetLevelSixTaxAmount ensures that no value is present for LevelSixTaxAmount, not even an explicit nil
### GetCreatedBy

`func (o *UnpostedInvoice) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *UnpostedInvoice) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *UnpostedInvoice) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *UnpostedInvoice) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDateClosed

`func (o *UnpostedInvoice) GetDateClosed() string`

GetDateClosed returns the DateClosed field if non-nil, zero value otherwise.

### GetDateClosedOk

`func (o *UnpostedInvoice) GetDateClosedOk() (*string, bool)`

GetDateClosedOk returns a tuple with the DateClosed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateClosed

`func (o *UnpostedInvoice) SetDateClosed(v string)`

SetDateClosed sets DateClosed field to given value.

### HasDateClosed

`func (o *UnpostedInvoice) HasDateClosed() bool`

HasDateClosed returns a boolean if a field has been set.

### GetInfo

`func (o *UnpostedInvoice) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *UnpostedInvoice) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *UnpostedInvoice) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *UnpostedInvoice) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


