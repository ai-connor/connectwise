# GLExportPurchaseTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**DocumentDate** | Pointer to **string** |  | [optional] 
**DocumentNumber** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Memo** | Pointer to **string** |  | [optional] 
**ApAccountNumber** | Pointer to **string** |  | [optional] 
**PurchaseDate** | Pointer to **string** |  | [optional] 
**Company** | Pointer to [**CompanyReference**](CompanyReference.md) |  | [optional] 
**CompanyType** | Pointer to [**CompanyTypeReference**](CompanyTypeReference.md) |  | [optional] 
**Contact** | Pointer to [**ContactReference**](ContactReference.md) |  | [optional] 
**Site** | Pointer to [**SiteReference**](SiteReference.md) |  | [optional] 
**PurchaseClass** | Pointer to **string** |  | [optional] 
**FreightAmount** | Pointer to **NullableFloat64** |  | [optional] 
**FreightPackingSlip** | Pointer to **string** |  | [optional] 
**PackingSlip** | Pointer to **string** |  | [optional] 
**DropshipFlag** | Pointer to **NullableBool** |  | [optional] 
**Currency** | Pointer to [**CurrencyReference**](CurrencyReference.md) |  | [optional] 
**Total** | Pointer to **NullableFloat64** |  | [optional] 
**BillingTerms** | Pointer to [**BillingTermsReference**](BillingTermsReference.md) |  | [optional] 
**BillingTermsXref** | Pointer to **string** |  | [optional] 
**DueDays** | Pointer to **NullableInt32** |  | [optional] 
**VendorNumber** | Pointer to **string** |  | [optional] 
**VendorAccountNumber** | Pointer to **string** |  | [optional] 
**VendorInvoiceDate** | Pointer to **string** |  | [optional] 
**VendorInvoiceNumber** | Pointer to **string** |  | [optional] 
**TaxAgencyXref** | Pointer to **string** |  | [optional] 
**StateTaxXref** | Pointer to **string** |  | [optional] 
**CountyTaxXref** | Pointer to **string** |  | [optional] 
**CityTaxXref** | Pointer to **string** |  | [optional] 
**ShipToCompany** | Pointer to [**CompanyReference**](CompanyReference.md) |  | [optional] 
**ShipToCompanyAccountNumber** | Pointer to **string** |  | [optional] 
**ShipToCompanyType** | Pointer to [**CompanyTypeReference**](CompanyTypeReference.md) |  | [optional] 
**ShipToContact** | Pointer to [**ContactReference**](ContactReference.md) |  | [optional] 
**ShipToSite** | Pointer to [**SiteReference**](SiteReference.md) |  | [optional] 
**ShipToTaxGroup** | Pointer to **string** |  | [optional] 
**TaxCode** | Pointer to [**TaxCodeReference**](TaxCodeReference.md) |  | [optional] 
**TaxGroupRate** | Pointer to **NullableFloat64** |  | [optional] 
**UseAvalaraTaxFlag** | Pointer to **NullableBool** |  | [optional] 
**PurchaseHeaderTaxGroup** | Pointer to **string** |  | [optional] 
**PurchaseHeaderTaxableFlag** | Pointer to **NullableBool** |  | [optional] 
**PurchaseHeaderFreightTaxableFlag** | Pointer to **NullableBool** |  | [optional] 
**TaxLevels** | Pointer to [**[]GLExportPurchaseTransactionTaxLevel**](GLExportPurchaseTransactionTaxLevel.md) |  | [optional] 
**PurchaseDetail** | Pointer to [**[]GLExportPurchaseTransactionDetail**](GLExportPurchaseTransactionDetail.md) |  | [optional] 
**PurchaseDetailTax** | Pointer to [**[]GLExportPurchaseTransactionDetailTax**](GLExportPurchaseTransactionDetailTax.md) |  | [optional] 

## Methods

### NewGLExportPurchaseTransaction

`func NewGLExportPurchaseTransaction() *GLExportPurchaseTransaction`

NewGLExportPurchaseTransaction instantiates a new GLExportPurchaseTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGLExportPurchaseTransactionWithDefaults

`func NewGLExportPurchaseTransactionWithDefaults() *GLExportPurchaseTransaction`

NewGLExportPurchaseTransactionWithDefaults instantiates a new GLExportPurchaseTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GLExportPurchaseTransaction) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GLExportPurchaseTransaction) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GLExportPurchaseTransaction) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GLExportPurchaseTransaction) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDocumentDate

`func (o *GLExportPurchaseTransaction) GetDocumentDate() string`

GetDocumentDate returns the DocumentDate field if non-nil, zero value otherwise.

### GetDocumentDateOk

`func (o *GLExportPurchaseTransaction) GetDocumentDateOk() (*string, bool)`

GetDocumentDateOk returns a tuple with the DocumentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentDate

`func (o *GLExportPurchaseTransaction) SetDocumentDate(v string)`

SetDocumentDate sets DocumentDate field to given value.

### HasDocumentDate

`func (o *GLExportPurchaseTransaction) HasDocumentDate() bool`

HasDocumentDate returns a boolean if a field has been set.

### GetDocumentNumber

`func (o *GLExportPurchaseTransaction) GetDocumentNumber() string`

GetDocumentNumber returns the DocumentNumber field if non-nil, zero value otherwise.

### GetDocumentNumberOk

`func (o *GLExportPurchaseTransaction) GetDocumentNumberOk() (*string, bool)`

GetDocumentNumberOk returns a tuple with the DocumentNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentNumber

`func (o *GLExportPurchaseTransaction) SetDocumentNumber(v string)`

SetDocumentNumber sets DocumentNumber field to given value.

### HasDocumentNumber

`func (o *GLExportPurchaseTransaction) HasDocumentNumber() bool`

HasDocumentNumber returns a boolean if a field has been set.

### GetDescription

`func (o *GLExportPurchaseTransaction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GLExportPurchaseTransaction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GLExportPurchaseTransaction) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GLExportPurchaseTransaction) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMemo

`func (o *GLExportPurchaseTransaction) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *GLExportPurchaseTransaction) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *GLExportPurchaseTransaction) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *GLExportPurchaseTransaction) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetApAccountNumber

`func (o *GLExportPurchaseTransaction) GetApAccountNumber() string`

GetApAccountNumber returns the ApAccountNumber field if non-nil, zero value otherwise.

### GetApAccountNumberOk

`func (o *GLExportPurchaseTransaction) GetApAccountNumberOk() (*string, bool)`

GetApAccountNumberOk returns a tuple with the ApAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApAccountNumber

`func (o *GLExportPurchaseTransaction) SetApAccountNumber(v string)`

SetApAccountNumber sets ApAccountNumber field to given value.

### HasApAccountNumber

`func (o *GLExportPurchaseTransaction) HasApAccountNumber() bool`

HasApAccountNumber returns a boolean if a field has been set.

### GetPurchaseDate

`func (o *GLExportPurchaseTransaction) GetPurchaseDate() string`

GetPurchaseDate returns the PurchaseDate field if non-nil, zero value otherwise.

### GetPurchaseDateOk

`func (o *GLExportPurchaseTransaction) GetPurchaseDateOk() (*string, bool)`

GetPurchaseDateOk returns a tuple with the PurchaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseDate

`func (o *GLExportPurchaseTransaction) SetPurchaseDate(v string)`

SetPurchaseDate sets PurchaseDate field to given value.

### HasPurchaseDate

`func (o *GLExportPurchaseTransaction) HasPurchaseDate() bool`

HasPurchaseDate returns a boolean if a field has been set.

### GetCompany

`func (o *GLExportPurchaseTransaction) GetCompany() CompanyReference`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *GLExportPurchaseTransaction) GetCompanyOk() (*CompanyReference, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *GLExportPurchaseTransaction) SetCompany(v CompanyReference)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *GLExportPurchaseTransaction) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetCompanyType

`func (o *GLExportPurchaseTransaction) GetCompanyType() CompanyTypeReference`

GetCompanyType returns the CompanyType field if non-nil, zero value otherwise.

### GetCompanyTypeOk

`func (o *GLExportPurchaseTransaction) GetCompanyTypeOk() (*CompanyTypeReference, bool)`

GetCompanyTypeOk returns a tuple with the CompanyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyType

`func (o *GLExportPurchaseTransaction) SetCompanyType(v CompanyTypeReference)`

SetCompanyType sets CompanyType field to given value.

### HasCompanyType

`func (o *GLExportPurchaseTransaction) HasCompanyType() bool`

HasCompanyType returns a boolean if a field has been set.

### GetContact

`func (o *GLExportPurchaseTransaction) GetContact() ContactReference`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *GLExportPurchaseTransaction) GetContactOk() (*ContactReference, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *GLExportPurchaseTransaction) SetContact(v ContactReference)`

SetContact sets Contact field to given value.

### HasContact

`func (o *GLExportPurchaseTransaction) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetSite

`func (o *GLExportPurchaseTransaction) GetSite() SiteReference`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *GLExportPurchaseTransaction) GetSiteOk() (*SiteReference, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *GLExportPurchaseTransaction) SetSite(v SiteReference)`

SetSite sets Site field to given value.

### HasSite

`func (o *GLExportPurchaseTransaction) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetPurchaseClass

`func (o *GLExportPurchaseTransaction) GetPurchaseClass() string`

GetPurchaseClass returns the PurchaseClass field if non-nil, zero value otherwise.

### GetPurchaseClassOk

`func (o *GLExportPurchaseTransaction) GetPurchaseClassOk() (*string, bool)`

GetPurchaseClassOk returns a tuple with the PurchaseClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseClass

`func (o *GLExportPurchaseTransaction) SetPurchaseClass(v string)`

SetPurchaseClass sets PurchaseClass field to given value.

### HasPurchaseClass

`func (o *GLExportPurchaseTransaction) HasPurchaseClass() bool`

HasPurchaseClass returns a boolean if a field has been set.

### GetFreightAmount

`func (o *GLExportPurchaseTransaction) GetFreightAmount() float64`

GetFreightAmount returns the FreightAmount field if non-nil, zero value otherwise.

### GetFreightAmountOk

`func (o *GLExportPurchaseTransaction) GetFreightAmountOk() (*float64, bool)`

GetFreightAmountOk returns a tuple with the FreightAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreightAmount

`func (o *GLExportPurchaseTransaction) SetFreightAmount(v float64)`

SetFreightAmount sets FreightAmount field to given value.

### HasFreightAmount

`func (o *GLExportPurchaseTransaction) HasFreightAmount() bool`

HasFreightAmount returns a boolean if a field has been set.

### SetFreightAmountNil

`func (o *GLExportPurchaseTransaction) SetFreightAmountNil(b bool)`

 SetFreightAmountNil sets the value for FreightAmount to be an explicit nil

### UnsetFreightAmount
`func (o *GLExportPurchaseTransaction) UnsetFreightAmount()`

UnsetFreightAmount ensures that no value is present for FreightAmount, not even an explicit nil
### GetFreightPackingSlip

`func (o *GLExportPurchaseTransaction) GetFreightPackingSlip() string`

GetFreightPackingSlip returns the FreightPackingSlip field if non-nil, zero value otherwise.

### GetFreightPackingSlipOk

`func (o *GLExportPurchaseTransaction) GetFreightPackingSlipOk() (*string, bool)`

GetFreightPackingSlipOk returns a tuple with the FreightPackingSlip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreightPackingSlip

`func (o *GLExportPurchaseTransaction) SetFreightPackingSlip(v string)`

SetFreightPackingSlip sets FreightPackingSlip field to given value.

### HasFreightPackingSlip

`func (o *GLExportPurchaseTransaction) HasFreightPackingSlip() bool`

HasFreightPackingSlip returns a boolean if a field has been set.

### GetPackingSlip

`func (o *GLExportPurchaseTransaction) GetPackingSlip() string`

GetPackingSlip returns the PackingSlip field if non-nil, zero value otherwise.

### GetPackingSlipOk

`func (o *GLExportPurchaseTransaction) GetPackingSlipOk() (*string, bool)`

GetPackingSlipOk returns a tuple with the PackingSlip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackingSlip

`func (o *GLExportPurchaseTransaction) SetPackingSlip(v string)`

SetPackingSlip sets PackingSlip field to given value.

### HasPackingSlip

`func (o *GLExportPurchaseTransaction) HasPackingSlip() bool`

HasPackingSlip returns a boolean if a field has been set.

### GetDropshipFlag

`func (o *GLExportPurchaseTransaction) GetDropshipFlag() bool`

GetDropshipFlag returns the DropshipFlag field if non-nil, zero value otherwise.

### GetDropshipFlagOk

`func (o *GLExportPurchaseTransaction) GetDropshipFlagOk() (*bool, bool)`

GetDropshipFlagOk returns a tuple with the DropshipFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropshipFlag

`func (o *GLExportPurchaseTransaction) SetDropshipFlag(v bool)`

SetDropshipFlag sets DropshipFlag field to given value.

### HasDropshipFlag

`func (o *GLExportPurchaseTransaction) HasDropshipFlag() bool`

HasDropshipFlag returns a boolean if a field has been set.

### SetDropshipFlagNil

`func (o *GLExportPurchaseTransaction) SetDropshipFlagNil(b bool)`

 SetDropshipFlagNil sets the value for DropshipFlag to be an explicit nil

### UnsetDropshipFlag
`func (o *GLExportPurchaseTransaction) UnsetDropshipFlag()`

UnsetDropshipFlag ensures that no value is present for DropshipFlag, not even an explicit nil
### GetCurrency

`func (o *GLExportPurchaseTransaction) GetCurrency() CurrencyReference`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GLExportPurchaseTransaction) GetCurrencyOk() (*CurrencyReference, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GLExportPurchaseTransaction) SetCurrency(v CurrencyReference)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *GLExportPurchaseTransaction) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetTotal

`func (o *GLExportPurchaseTransaction) GetTotal() float64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GLExportPurchaseTransaction) GetTotalOk() (*float64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GLExportPurchaseTransaction) SetTotal(v float64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GLExportPurchaseTransaction) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### SetTotalNil

`func (o *GLExportPurchaseTransaction) SetTotalNil(b bool)`

 SetTotalNil sets the value for Total to be an explicit nil

### UnsetTotal
`func (o *GLExportPurchaseTransaction) UnsetTotal()`

UnsetTotal ensures that no value is present for Total, not even an explicit nil
### GetBillingTerms

`func (o *GLExportPurchaseTransaction) GetBillingTerms() BillingTermsReference`

GetBillingTerms returns the BillingTerms field if non-nil, zero value otherwise.

### GetBillingTermsOk

`func (o *GLExportPurchaseTransaction) GetBillingTermsOk() (*BillingTermsReference, bool)`

GetBillingTermsOk returns a tuple with the BillingTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingTerms

`func (o *GLExportPurchaseTransaction) SetBillingTerms(v BillingTermsReference)`

SetBillingTerms sets BillingTerms field to given value.

### HasBillingTerms

`func (o *GLExportPurchaseTransaction) HasBillingTerms() bool`

HasBillingTerms returns a boolean if a field has been set.

### GetBillingTermsXref

`func (o *GLExportPurchaseTransaction) GetBillingTermsXref() string`

GetBillingTermsXref returns the BillingTermsXref field if non-nil, zero value otherwise.

### GetBillingTermsXrefOk

`func (o *GLExportPurchaseTransaction) GetBillingTermsXrefOk() (*string, bool)`

GetBillingTermsXrefOk returns a tuple with the BillingTermsXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingTermsXref

`func (o *GLExportPurchaseTransaction) SetBillingTermsXref(v string)`

SetBillingTermsXref sets BillingTermsXref field to given value.

### HasBillingTermsXref

`func (o *GLExportPurchaseTransaction) HasBillingTermsXref() bool`

HasBillingTermsXref returns a boolean if a field has been set.

### GetDueDays

`func (o *GLExportPurchaseTransaction) GetDueDays() int32`

GetDueDays returns the DueDays field if non-nil, zero value otherwise.

### GetDueDaysOk

`func (o *GLExportPurchaseTransaction) GetDueDaysOk() (*int32, bool)`

GetDueDaysOk returns a tuple with the DueDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDays

`func (o *GLExportPurchaseTransaction) SetDueDays(v int32)`

SetDueDays sets DueDays field to given value.

### HasDueDays

`func (o *GLExportPurchaseTransaction) HasDueDays() bool`

HasDueDays returns a boolean if a field has been set.

### SetDueDaysNil

`func (o *GLExportPurchaseTransaction) SetDueDaysNil(b bool)`

 SetDueDaysNil sets the value for DueDays to be an explicit nil

### UnsetDueDays
`func (o *GLExportPurchaseTransaction) UnsetDueDays()`

UnsetDueDays ensures that no value is present for DueDays, not even an explicit nil
### GetVendorNumber

`func (o *GLExportPurchaseTransaction) GetVendorNumber() string`

GetVendorNumber returns the VendorNumber field if non-nil, zero value otherwise.

### GetVendorNumberOk

`func (o *GLExportPurchaseTransaction) GetVendorNumberOk() (*string, bool)`

GetVendorNumberOk returns a tuple with the VendorNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorNumber

`func (o *GLExportPurchaseTransaction) SetVendorNumber(v string)`

SetVendorNumber sets VendorNumber field to given value.

### HasVendorNumber

`func (o *GLExportPurchaseTransaction) HasVendorNumber() bool`

HasVendorNumber returns a boolean if a field has been set.

### GetVendorAccountNumber

`func (o *GLExportPurchaseTransaction) GetVendorAccountNumber() string`

GetVendorAccountNumber returns the VendorAccountNumber field if non-nil, zero value otherwise.

### GetVendorAccountNumberOk

`func (o *GLExportPurchaseTransaction) GetVendorAccountNumberOk() (*string, bool)`

GetVendorAccountNumberOk returns a tuple with the VendorAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorAccountNumber

`func (o *GLExportPurchaseTransaction) SetVendorAccountNumber(v string)`

SetVendorAccountNumber sets VendorAccountNumber field to given value.

### HasVendorAccountNumber

`func (o *GLExportPurchaseTransaction) HasVendorAccountNumber() bool`

HasVendorAccountNumber returns a boolean if a field has been set.

### GetVendorInvoiceDate

`func (o *GLExportPurchaseTransaction) GetVendorInvoiceDate() string`

GetVendorInvoiceDate returns the VendorInvoiceDate field if non-nil, zero value otherwise.

### GetVendorInvoiceDateOk

`func (o *GLExportPurchaseTransaction) GetVendorInvoiceDateOk() (*string, bool)`

GetVendorInvoiceDateOk returns a tuple with the VendorInvoiceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorInvoiceDate

`func (o *GLExportPurchaseTransaction) SetVendorInvoiceDate(v string)`

SetVendorInvoiceDate sets VendorInvoiceDate field to given value.

### HasVendorInvoiceDate

`func (o *GLExportPurchaseTransaction) HasVendorInvoiceDate() bool`

HasVendorInvoiceDate returns a boolean if a field has been set.

### GetVendorInvoiceNumber

`func (o *GLExportPurchaseTransaction) GetVendorInvoiceNumber() string`

GetVendorInvoiceNumber returns the VendorInvoiceNumber field if non-nil, zero value otherwise.

### GetVendorInvoiceNumberOk

`func (o *GLExportPurchaseTransaction) GetVendorInvoiceNumberOk() (*string, bool)`

GetVendorInvoiceNumberOk returns a tuple with the VendorInvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorInvoiceNumber

`func (o *GLExportPurchaseTransaction) SetVendorInvoiceNumber(v string)`

SetVendorInvoiceNumber sets VendorInvoiceNumber field to given value.

### HasVendorInvoiceNumber

`func (o *GLExportPurchaseTransaction) HasVendorInvoiceNumber() bool`

HasVendorInvoiceNumber returns a boolean if a field has been set.

### GetTaxAgencyXref

`func (o *GLExportPurchaseTransaction) GetTaxAgencyXref() string`

GetTaxAgencyXref returns the TaxAgencyXref field if non-nil, zero value otherwise.

### GetTaxAgencyXrefOk

`func (o *GLExportPurchaseTransaction) GetTaxAgencyXrefOk() (*string, bool)`

GetTaxAgencyXrefOk returns a tuple with the TaxAgencyXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAgencyXref

`func (o *GLExportPurchaseTransaction) SetTaxAgencyXref(v string)`

SetTaxAgencyXref sets TaxAgencyXref field to given value.

### HasTaxAgencyXref

`func (o *GLExportPurchaseTransaction) HasTaxAgencyXref() bool`

HasTaxAgencyXref returns a boolean if a field has been set.

### GetStateTaxXref

`func (o *GLExportPurchaseTransaction) GetStateTaxXref() string`

GetStateTaxXref returns the StateTaxXref field if non-nil, zero value otherwise.

### GetStateTaxXrefOk

`func (o *GLExportPurchaseTransaction) GetStateTaxXrefOk() (*string, bool)`

GetStateTaxXrefOk returns a tuple with the StateTaxXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateTaxXref

`func (o *GLExportPurchaseTransaction) SetStateTaxXref(v string)`

SetStateTaxXref sets StateTaxXref field to given value.

### HasStateTaxXref

`func (o *GLExportPurchaseTransaction) HasStateTaxXref() bool`

HasStateTaxXref returns a boolean if a field has been set.

### GetCountyTaxXref

`func (o *GLExportPurchaseTransaction) GetCountyTaxXref() string`

GetCountyTaxXref returns the CountyTaxXref field if non-nil, zero value otherwise.

### GetCountyTaxXrefOk

`func (o *GLExportPurchaseTransaction) GetCountyTaxXrefOk() (*string, bool)`

GetCountyTaxXrefOk returns a tuple with the CountyTaxXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountyTaxXref

`func (o *GLExportPurchaseTransaction) SetCountyTaxXref(v string)`

SetCountyTaxXref sets CountyTaxXref field to given value.

### HasCountyTaxXref

`func (o *GLExportPurchaseTransaction) HasCountyTaxXref() bool`

HasCountyTaxXref returns a boolean if a field has been set.

### GetCityTaxXref

`func (o *GLExportPurchaseTransaction) GetCityTaxXref() string`

GetCityTaxXref returns the CityTaxXref field if non-nil, zero value otherwise.

### GetCityTaxXrefOk

`func (o *GLExportPurchaseTransaction) GetCityTaxXrefOk() (*string, bool)`

GetCityTaxXrefOk returns a tuple with the CityTaxXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityTaxXref

`func (o *GLExportPurchaseTransaction) SetCityTaxXref(v string)`

SetCityTaxXref sets CityTaxXref field to given value.

### HasCityTaxXref

`func (o *GLExportPurchaseTransaction) HasCityTaxXref() bool`

HasCityTaxXref returns a boolean if a field has been set.

### GetShipToCompany

`func (o *GLExportPurchaseTransaction) GetShipToCompany() CompanyReference`

GetShipToCompany returns the ShipToCompany field if non-nil, zero value otherwise.

### GetShipToCompanyOk

`func (o *GLExportPurchaseTransaction) GetShipToCompanyOk() (*CompanyReference, bool)`

GetShipToCompanyOk returns a tuple with the ShipToCompany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipToCompany

`func (o *GLExportPurchaseTransaction) SetShipToCompany(v CompanyReference)`

SetShipToCompany sets ShipToCompany field to given value.

### HasShipToCompany

`func (o *GLExportPurchaseTransaction) HasShipToCompany() bool`

HasShipToCompany returns a boolean if a field has been set.

### GetShipToCompanyAccountNumber

`func (o *GLExportPurchaseTransaction) GetShipToCompanyAccountNumber() string`

GetShipToCompanyAccountNumber returns the ShipToCompanyAccountNumber field if non-nil, zero value otherwise.

### GetShipToCompanyAccountNumberOk

`func (o *GLExportPurchaseTransaction) GetShipToCompanyAccountNumberOk() (*string, bool)`

GetShipToCompanyAccountNumberOk returns a tuple with the ShipToCompanyAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipToCompanyAccountNumber

`func (o *GLExportPurchaseTransaction) SetShipToCompanyAccountNumber(v string)`

SetShipToCompanyAccountNumber sets ShipToCompanyAccountNumber field to given value.

### HasShipToCompanyAccountNumber

`func (o *GLExportPurchaseTransaction) HasShipToCompanyAccountNumber() bool`

HasShipToCompanyAccountNumber returns a boolean if a field has been set.

### GetShipToCompanyType

`func (o *GLExportPurchaseTransaction) GetShipToCompanyType() CompanyTypeReference`

GetShipToCompanyType returns the ShipToCompanyType field if non-nil, zero value otherwise.

### GetShipToCompanyTypeOk

`func (o *GLExportPurchaseTransaction) GetShipToCompanyTypeOk() (*CompanyTypeReference, bool)`

GetShipToCompanyTypeOk returns a tuple with the ShipToCompanyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipToCompanyType

`func (o *GLExportPurchaseTransaction) SetShipToCompanyType(v CompanyTypeReference)`

SetShipToCompanyType sets ShipToCompanyType field to given value.

### HasShipToCompanyType

`func (o *GLExportPurchaseTransaction) HasShipToCompanyType() bool`

HasShipToCompanyType returns a boolean if a field has been set.

### GetShipToContact

`func (o *GLExportPurchaseTransaction) GetShipToContact() ContactReference`

GetShipToContact returns the ShipToContact field if non-nil, zero value otherwise.

### GetShipToContactOk

`func (o *GLExportPurchaseTransaction) GetShipToContactOk() (*ContactReference, bool)`

GetShipToContactOk returns a tuple with the ShipToContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipToContact

`func (o *GLExportPurchaseTransaction) SetShipToContact(v ContactReference)`

SetShipToContact sets ShipToContact field to given value.

### HasShipToContact

`func (o *GLExportPurchaseTransaction) HasShipToContact() bool`

HasShipToContact returns a boolean if a field has been set.

### GetShipToSite

`func (o *GLExportPurchaseTransaction) GetShipToSite() SiteReference`

GetShipToSite returns the ShipToSite field if non-nil, zero value otherwise.

### GetShipToSiteOk

`func (o *GLExportPurchaseTransaction) GetShipToSiteOk() (*SiteReference, bool)`

GetShipToSiteOk returns a tuple with the ShipToSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipToSite

`func (o *GLExportPurchaseTransaction) SetShipToSite(v SiteReference)`

SetShipToSite sets ShipToSite field to given value.

### HasShipToSite

`func (o *GLExportPurchaseTransaction) HasShipToSite() bool`

HasShipToSite returns a boolean if a field has been set.

### GetShipToTaxGroup

`func (o *GLExportPurchaseTransaction) GetShipToTaxGroup() string`

GetShipToTaxGroup returns the ShipToTaxGroup field if non-nil, zero value otherwise.

### GetShipToTaxGroupOk

`func (o *GLExportPurchaseTransaction) GetShipToTaxGroupOk() (*string, bool)`

GetShipToTaxGroupOk returns a tuple with the ShipToTaxGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipToTaxGroup

`func (o *GLExportPurchaseTransaction) SetShipToTaxGroup(v string)`

SetShipToTaxGroup sets ShipToTaxGroup field to given value.

### HasShipToTaxGroup

`func (o *GLExportPurchaseTransaction) HasShipToTaxGroup() bool`

HasShipToTaxGroup returns a boolean if a field has been set.

### GetTaxCode

`func (o *GLExportPurchaseTransaction) GetTaxCode() TaxCodeReference`

GetTaxCode returns the TaxCode field if non-nil, zero value otherwise.

### GetTaxCodeOk

`func (o *GLExportPurchaseTransaction) GetTaxCodeOk() (*TaxCodeReference, bool)`

GetTaxCodeOk returns a tuple with the TaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCode

`func (o *GLExportPurchaseTransaction) SetTaxCode(v TaxCodeReference)`

SetTaxCode sets TaxCode field to given value.

### HasTaxCode

`func (o *GLExportPurchaseTransaction) HasTaxCode() bool`

HasTaxCode returns a boolean if a field has been set.

### GetTaxGroupRate

`func (o *GLExportPurchaseTransaction) GetTaxGroupRate() float64`

GetTaxGroupRate returns the TaxGroupRate field if non-nil, zero value otherwise.

### GetTaxGroupRateOk

`func (o *GLExportPurchaseTransaction) GetTaxGroupRateOk() (*float64, bool)`

GetTaxGroupRateOk returns a tuple with the TaxGroupRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxGroupRate

`func (o *GLExportPurchaseTransaction) SetTaxGroupRate(v float64)`

SetTaxGroupRate sets TaxGroupRate field to given value.

### HasTaxGroupRate

`func (o *GLExportPurchaseTransaction) HasTaxGroupRate() bool`

HasTaxGroupRate returns a boolean if a field has been set.

### SetTaxGroupRateNil

`func (o *GLExportPurchaseTransaction) SetTaxGroupRateNil(b bool)`

 SetTaxGroupRateNil sets the value for TaxGroupRate to be an explicit nil

### UnsetTaxGroupRate
`func (o *GLExportPurchaseTransaction) UnsetTaxGroupRate()`

UnsetTaxGroupRate ensures that no value is present for TaxGroupRate, not even an explicit nil
### GetUseAvalaraTaxFlag

`func (o *GLExportPurchaseTransaction) GetUseAvalaraTaxFlag() bool`

GetUseAvalaraTaxFlag returns the UseAvalaraTaxFlag field if non-nil, zero value otherwise.

### GetUseAvalaraTaxFlagOk

`func (o *GLExportPurchaseTransaction) GetUseAvalaraTaxFlagOk() (*bool, bool)`

GetUseAvalaraTaxFlagOk returns a tuple with the UseAvalaraTaxFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAvalaraTaxFlag

`func (o *GLExportPurchaseTransaction) SetUseAvalaraTaxFlag(v bool)`

SetUseAvalaraTaxFlag sets UseAvalaraTaxFlag field to given value.

### HasUseAvalaraTaxFlag

`func (o *GLExportPurchaseTransaction) HasUseAvalaraTaxFlag() bool`

HasUseAvalaraTaxFlag returns a boolean if a field has been set.

### SetUseAvalaraTaxFlagNil

`func (o *GLExportPurchaseTransaction) SetUseAvalaraTaxFlagNil(b bool)`

 SetUseAvalaraTaxFlagNil sets the value for UseAvalaraTaxFlag to be an explicit nil

### UnsetUseAvalaraTaxFlag
`func (o *GLExportPurchaseTransaction) UnsetUseAvalaraTaxFlag()`

UnsetUseAvalaraTaxFlag ensures that no value is present for UseAvalaraTaxFlag, not even an explicit nil
### GetPurchaseHeaderTaxGroup

`func (o *GLExportPurchaseTransaction) GetPurchaseHeaderTaxGroup() string`

GetPurchaseHeaderTaxGroup returns the PurchaseHeaderTaxGroup field if non-nil, zero value otherwise.

### GetPurchaseHeaderTaxGroupOk

`func (o *GLExportPurchaseTransaction) GetPurchaseHeaderTaxGroupOk() (*string, bool)`

GetPurchaseHeaderTaxGroupOk returns a tuple with the PurchaseHeaderTaxGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseHeaderTaxGroup

`func (o *GLExportPurchaseTransaction) SetPurchaseHeaderTaxGroup(v string)`

SetPurchaseHeaderTaxGroup sets PurchaseHeaderTaxGroup field to given value.

### HasPurchaseHeaderTaxGroup

`func (o *GLExportPurchaseTransaction) HasPurchaseHeaderTaxGroup() bool`

HasPurchaseHeaderTaxGroup returns a boolean if a field has been set.

### GetPurchaseHeaderTaxableFlag

`func (o *GLExportPurchaseTransaction) GetPurchaseHeaderTaxableFlag() bool`

GetPurchaseHeaderTaxableFlag returns the PurchaseHeaderTaxableFlag field if non-nil, zero value otherwise.

### GetPurchaseHeaderTaxableFlagOk

`func (o *GLExportPurchaseTransaction) GetPurchaseHeaderTaxableFlagOk() (*bool, bool)`

GetPurchaseHeaderTaxableFlagOk returns a tuple with the PurchaseHeaderTaxableFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseHeaderTaxableFlag

`func (o *GLExportPurchaseTransaction) SetPurchaseHeaderTaxableFlag(v bool)`

SetPurchaseHeaderTaxableFlag sets PurchaseHeaderTaxableFlag field to given value.

### HasPurchaseHeaderTaxableFlag

`func (o *GLExportPurchaseTransaction) HasPurchaseHeaderTaxableFlag() bool`

HasPurchaseHeaderTaxableFlag returns a boolean if a field has been set.

### SetPurchaseHeaderTaxableFlagNil

`func (o *GLExportPurchaseTransaction) SetPurchaseHeaderTaxableFlagNil(b bool)`

 SetPurchaseHeaderTaxableFlagNil sets the value for PurchaseHeaderTaxableFlag to be an explicit nil

### UnsetPurchaseHeaderTaxableFlag
`func (o *GLExportPurchaseTransaction) UnsetPurchaseHeaderTaxableFlag()`

UnsetPurchaseHeaderTaxableFlag ensures that no value is present for PurchaseHeaderTaxableFlag, not even an explicit nil
### GetPurchaseHeaderFreightTaxableFlag

`func (o *GLExportPurchaseTransaction) GetPurchaseHeaderFreightTaxableFlag() bool`

GetPurchaseHeaderFreightTaxableFlag returns the PurchaseHeaderFreightTaxableFlag field if non-nil, zero value otherwise.

### GetPurchaseHeaderFreightTaxableFlagOk

`func (o *GLExportPurchaseTransaction) GetPurchaseHeaderFreightTaxableFlagOk() (*bool, bool)`

GetPurchaseHeaderFreightTaxableFlagOk returns a tuple with the PurchaseHeaderFreightTaxableFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseHeaderFreightTaxableFlag

`func (o *GLExportPurchaseTransaction) SetPurchaseHeaderFreightTaxableFlag(v bool)`

SetPurchaseHeaderFreightTaxableFlag sets PurchaseHeaderFreightTaxableFlag field to given value.

### HasPurchaseHeaderFreightTaxableFlag

`func (o *GLExportPurchaseTransaction) HasPurchaseHeaderFreightTaxableFlag() bool`

HasPurchaseHeaderFreightTaxableFlag returns a boolean if a field has been set.

### SetPurchaseHeaderFreightTaxableFlagNil

`func (o *GLExportPurchaseTransaction) SetPurchaseHeaderFreightTaxableFlagNil(b bool)`

 SetPurchaseHeaderFreightTaxableFlagNil sets the value for PurchaseHeaderFreightTaxableFlag to be an explicit nil

### UnsetPurchaseHeaderFreightTaxableFlag
`func (o *GLExportPurchaseTransaction) UnsetPurchaseHeaderFreightTaxableFlag()`

UnsetPurchaseHeaderFreightTaxableFlag ensures that no value is present for PurchaseHeaderFreightTaxableFlag, not even an explicit nil
### GetTaxLevels

`func (o *GLExportPurchaseTransaction) GetTaxLevels() []GLExportPurchaseTransactionTaxLevel`

GetTaxLevels returns the TaxLevels field if non-nil, zero value otherwise.

### GetTaxLevelsOk

`func (o *GLExportPurchaseTransaction) GetTaxLevelsOk() (*[]GLExportPurchaseTransactionTaxLevel, bool)`

GetTaxLevelsOk returns a tuple with the TaxLevels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxLevels

`func (o *GLExportPurchaseTransaction) SetTaxLevels(v []GLExportPurchaseTransactionTaxLevel)`

SetTaxLevels sets TaxLevels field to given value.

### HasTaxLevels

`func (o *GLExportPurchaseTransaction) HasTaxLevels() bool`

HasTaxLevels returns a boolean if a field has been set.

### GetPurchaseDetail

`func (o *GLExportPurchaseTransaction) GetPurchaseDetail() []GLExportPurchaseTransactionDetail`

GetPurchaseDetail returns the PurchaseDetail field if non-nil, zero value otherwise.

### GetPurchaseDetailOk

`func (o *GLExportPurchaseTransaction) GetPurchaseDetailOk() (*[]GLExportPurchaseTransactionDetail, bool)`

GetPurchaseDetailOk returns a tuple with the PurchaseDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseDetail

`func (o *GLExportPurchaseTransaction) SetPurchaseDetail(v []GLExportPurchaseTransactionDetail)`

SetPurchaseDetail sets PurchaseDetail field to given value.

### HasPurchaseDetail

`func (o *GLExportPurchaseTransaction) HasPurchaseDetail() bool`

HasPurchaseDetail returns a boolean if a field has been set.

### GetPurchaseDetailTax

`func (o *GLExportPurchaseTransaction) GetPurchaseDetailTax() []GLExportPurchaseTransactionDetailTax`

GetPurchaseDetailTax returns the PurchaseDetailTax field if non-nil, zero value otherwise.

### GetPurchaseDetailTaxOk

`func (o *GLExportPurchaseTransaction) GetPurchaseDetailTaxOk() (*[]GLExportPurchaseTransactionDetailTax, bool)`

GetPurchaseDetailTaxOk returns a tuple with the PurchaseDetailTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseDetailTax

`func (o *GLExportPurchaseTransaction) SetPurchaseDetailTax(v []GLExportPurchaseTransactionDetailTax)`

SetPurchaseDetailTax sets PurchaseDetailTax field to given value.

### HasPurchaseDetailTax

`func (o *GLExportPurchaseTransaction) HasPurchaseDetailTax() bool`

HasPurchaseDetailTax returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


