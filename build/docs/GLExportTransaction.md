# GLExportTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** |  | [optional] 
**GlClass** | Pointer to **string** |  | [optional] 
**GlTypeId** | Pointer to **string** |  | [optional] 
**DocumentDate** | Pointer to **string** |  | [optional] 
**DocumentNumber** | Pointer to **string** |  | [optional] 
**DocumentType** | Pointer to **string** |  | [optional] 
**Memo** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Attention** | Pointer to **string** |  | [optional] 
**SalesTerritory** | Pointer to **string** |  | [optional] 
**Company** | Pointer to [**CompanyReference**](CompanyReference.md) |  | [optional] 
**CompanyType** | Pointer to [**CompanyTypeReference**](CompanyTypeReference.md) |  | [optional] 
**CompanyAccountNumber** | Pointer to **string** |  | [optional] 
**Site** | Pointer to [**SiteReference**](SiteReference.md) |  | [optional] 
**BillingTerms** | Pointer to [**BillingTermsReference**](BillingTermsReference.md) |  | [optional] 
**BillingTermsXref** | Pointer to **string** |  | [optional] 
**DueDays** | Pointer to **NullableInt32** |  | [optional] 
**DueDate** | Pointer to **string** |  | [optional] 
**EmailDeliveryFlag** | Pointer to **NullableBool** |  | [optional] 
**PrintDeliveryFlag** | Pointer to **NullableBool** |  | [optional] 
**AgreementPrePaymentFlag** | Pointer to **NullableBool** |  | [optional] 
**AccountNumber** | Pointer to **string** |  | [optional] 
**BillingType** | Pointer to **string** |  | [optional] 
**GlEntryIds** | Pointer to **string** |  | [optional] 
**PurchaseOrder** | Pointer to [**PurchaseOrderReference**](PurchaseOrderReference.md) |  | [optional] 
**Project** | Pointer to [**ProjectReference**](ProjectReference.md) |  | [optional] 
**Currency** | Pointer to [**CurrencyReference**](CurrencyReference.md) |  | [optional] 
**Total** | Pointer to **NullableFloat64** |  | [optional] 
**SalesRepId** | Pointer to **string** |  | [optional] 
**SalesRepName** | Pointer to **string** |  | [optional] 
**Taxable** | Pointer to **NullableBool** |  | [optional] 
**TaxableTotal** | Pointer to **NullableFloat64** |  | [optional] 
**TaxCode** | Pointer to [**TaxCodeReference**](TaxCodeReference.md) |  | [optional] 
**TaxGroupRate** | Pointer to **NullableFloat64** |  | [optional] 
**PiggyBackFlag** | Pointer to **NullableBool** |  | [optional] 
**TaxAccountNumber** | Pointer to **string** |  | [optional] 
**SalesTax** | Pointer to **NullableFloat64** |  | [optional] 
**StateTax** | Pointer to **NullableFloat64** |  | [optional] 
**CountyTax** | Pointer to **NullableFloat64** |  | [optional] 
**CityTax** | Pointer to **NullableFloat64** |  | [optional] 
**TaxableAmount1** | Pointer to **NullableFloat64** |  | [optional] 
**TaxableAmount2** | Pointer to **NullableFloat64** |  | [optional] 
**TaxableAmount3** | Pointer to **NullableFloat64** |  | [optional] 
**TaxableAmount4** | Pointer to **NullableFloat64** |  | [optional] 
**TaxableAmount5** | Pointer to **NullableFloat64** |  | [optional] 
**TaxAgencyXref** | Pointer to **string** |  | [optional] 
**StateTaxXref** | Pointer to **string** |  | [optional] 
**CountyTaxXref** | Pointer to **string** |  | [optional] 
**TaxId** | Pointer to **string** |  | [optional] 
**TaxDpAppliedFlag** | Pointer to **NullableBool** |  | [optional] 
**UseAvalaraFlag** | Pointer to **NullableBool** |  | [optional] 
**SendAvalaraTaxFlag** | Pointer to **NullableBool** |  | [optional] 
**ShipToCompany** | Pointer to [**CompanyReference**](CompanyReference.md) |  | [optional] 
**ShipToCompanyAccountNumber** | Pointer to **string** |  | [optional] 
**ShipToCompanyType** | Pointer to [**CompanyTypeReference**](CompanyTypeReference.md) |  | [optional] 
**ShipToTaxId** | Pointer to **string** |  | [optional] 
**ShipSite** | Pointer to [**SiteReference**](SiteReference.md) |  | [optional] 
**ShipContact** | Pointer to **string** |  | [optional] 
**Detail** | Pointer to [**[]GLExportTransactionDetail**](GLExportTransactionDetail.md) |  | [optional] 
**TaxLevels** | Pointer to [**[]GLExportTransactionTaxLevel**](GLExportTransactionTaxLevel.md) |  | [optional] 

## Methods

### NewGLExportTransaction

`func NewGLExportTransaction() *GLExportTransaction`

NewGLExportTransaction instantiates a new GLExportTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGLExportTransactionWithDefaults

`func NewGLExportTransactionWithDefaults() *GLExportTransaction`

NewGLExportTransactionWithDefaults instantiates a new GLExportTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GLExportTransaction) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GLExportTransaction) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GLExportTransaction) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GLExportTransaction) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *GLExportTransaction) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *GLExportTransaction) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetGlClass

`func (o *GLExportTransaction) GetGlClass() string`

GetGlClass returns the GlClass field if non-nil, zero value otherwise.

### GetGlClassOk

`func (o *GLExportTransaction) GetGlClassOk() (*string, bool)`

GetGlClassOk returns a tuple with the GlClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlClass

`func (o *GLExportTransaction) SetGlClass(v string)`

SetGlClass sets GlClass field to given value.

### HasGlClass

`func (o *GLExportTransaction) HasGlClass() bool`

HasGlClass returns a boolean if a field has been set.

### GetGlTypeId

`func (o *GLExportTransaction) GetGlTypeId() string`

GetGlTypeId returns the GlTypeId field if non-nil, zero value otherwise.

### GetGlTypeIdOk

`func (o *GLExportTransaction) GetGlTypeIdOk() (*string, bool)`

GetGlTypeIdOk returns a tuple with the GlTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlTypeId

`func (o *GLExportTransaction) SetGlTypeId(v string)`

SetGlTypeId sets GlTypeId field to given value.

### HasGlTypeId

`func (o *GLExportTransaction) HasGlTypeId() bool`

HasGlTypeId returns a boolean if a field has been set.

### GetDocumentDate

`func (o *GLExportTransaction) GetDocumentDate() string`

GetDocumentDate returns the DocumentDate field if non-nil, zero value otherwise.

### GetDocumentDateOk

`func (o *GLExportTransaction) GetDocumentDateOk() (*string, bool)`

GetDocumentDateOk returns a tuple with the DocumentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentDate

`func (o *GLExportTransaction) SetDocumentDate(v string)`

SetDocumentDate sets DocumentDate field to given value.

### HasDocumentDate

`func (o *GLExportTransaction) HasDocumentDate() bool`

HasDocumentDate returns a boolean if a field has been set.

### GetDocumentNumber

`func (o *GLExportTransaction) GetDocumentNumber() string`

GetDocumentNumber returns the DocumentNumber field if non-nil, zero value otherwise.

### GetDocumentNumberOk

`func (o *GLExportTransaction) GetDocumentNumberOk() (*string, bool)`

GetDocumentNumberOk returns a tuple with the DocumentNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentNumber

`func (o *GLExportTransaction) SetDocumentNumber(v string)`

SetDocumentNumber sets DocumentNumber field to given value.

### HasDocumentNumber

`func (o *GLExportTransaction) HasDocumentNumber() bool`

HasDocumentNumber returns a boolean if a field has been set.

### GetDocumentType

`func (o *GLExportTransaction) GetDocumentType() string`

GetDocumentType returns the DocumentType field if non-nil, zero value otherwise.

### GetDocumentTypeOk

`func (o *GLExportTransaction) GetDocumentTypeOk() (*string, bool)`

GetDocumentTypeOk returns a tuple with the DocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentType

`func (o *GLExportTransaction) SetDocumentType(v string)`

SetDocumentType sets DocumentType field to given value.

### HasDocumentType

`func (o *GLExportTransaction) HasDocumentType() bool`

HasDocumentType returns a boolean if a field has been set.

### GetMemo

`func (o *GLExportTransaction) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *GLExportTransaction) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *GLExportTransaction) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *GLExportTransaction) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetDescription

`func (o *GLExportTransaction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GLExportTransaction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GLExportTransaction) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GLExportTransaction) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAttention

`func (o *GLExportTransaction) GetAttention() string`

GetAttention returns the Attention field if non-nil, zero value otherwise.

### GetAttentionOk

`func (o *GLExportTransaction) GetAttentionOk() (*string, bool)`

GetAttentionOk returns a tuple with the Attention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttention

`func (o *GLExportTransaction) SetAttention(v string)`

SetAttention sets Attention field to given value.

### HasAttention

`func (o *GLExportTransaction) HasAttention() bool`

HasAttention returns a boolean if a field has been set.

### GetSalesTerritory

`func (o *GLExportTransaction) GetSalesTerritory() string`

GetSalesTerritory returns the SalesTerritory field if non-nil, zero value otherwise.

### GetSalesTerritoryOk

`func (o *GLExportTransaction) GetSalesTerritoryOk() (*string, bool)`

GetSalesTerritoryOk returns a tuple with the SalesTerritory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTerritory

`func (o *GLExportTransaction) SetSalesTerritory(v string)`

SetSalesTerritory sets SalesTerritory field to given value.

### HasSalesTerritory

`func (o *GLExportTransaction) HasSalesTerritory() bool`

HasSalesTerritory returns a boolean if a field has been set.

### GetCompany

`func (o *GLExportTransaction) GetCompany() CompanyReference`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *GLExportTransaction) GetCompanyOk() (*CompanyReference, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *GLExportTransaction) SetCompany(v CompanyReference)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *GLExportTransaction) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetCompanyType

`func (o *GLExportTransaction) GetCompanyType() CompanyTypeReference`

GetCompanyType returns the CompanyType field if non-nil, zero value otherwise.

### GetCompanyTypeOk

`func (o *GLExportTransaction) GetCompanyTypeOk() (*CompanyTypeReference, bool)`

GetCompanyTypeOk returns a tuple with the CompanyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyType

`func (o *GLExportTransaction) SetCompanyType(v CompanyTypeReference)`

SetCompanyType sets CompanyType field to given value.

### HasCompanyType

`func (o *GLExportTransaction) HasCompanyType() bool`

HasCompanyType returns a boolean if a field has been set.

### GetCompanyAccountNumber

`func (o *GLExportTransaction) GetCompanyAccountNumber() string`

GetCompanyAccountNumber returns the CompanyAccountNumber field if non-nil, zero value otherwise.

### GetCompanyAccountNumberOk

`func (o *GLExportTransaction) GetCompanyAccountNumberOk() (*string, bool)`

GetCompanyAccountNumberOk returns a tuple with the CompanyAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyAccountNumber

`func (o *GLExportTransaction) SetCompanyAccountNumber(v string)`

SetCompanyAccountNumber sets CompanyAccountNumber field to given value.

### HasCompanyAccountNumber

`func (o *GLExportTransaction) HasCompanyAccountNumber() bool`

HasCompanyAccountNumber returns a boolean if a field has been set.

### GetSite

`func (o *GLExportTransaction) GetSite() SiteReference`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *GLExportTransaction) GetSiteOk() (*SiteReference, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *GLExportTransaction) SetSite(v SiteReference)`

SetSite sets Site field to given value.

### HasSite

`func (o *GLExportTransaction) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetBillingTerms

`func (o *GLExportTransaction) GetBillingTerms() BillingTermsReference`

GetBillingTerms returns the BillingTerms field if non-nil, zero value otherwise.

### GetBillingTermsOk

`func (o *GLExportTransaction) GetBillingTermsOk() (*BillingTermsReference, bool)`

GetBillingTermsOk returns a tuple with the BillingTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingTerms

`func (o *GLExportTransaction) SetBillingTerms(v BillingTermsReference)`

SetBillingTerms sets BillingTerms field to given value.

### HasBillingTerms

`func (o *GLExportTransaction) HasBillingTerms() bool`

HasBillingTerms returns a boolean if a field has been set.

### GetBillingTermsXref

`func (o *GLExportTransaction) GetBillingTermsXref() string`

GetBillingTermsXref returns the BillingTermsXref field if non-nil, zero value otherwise.

### GetBillingTermsXrefOk

`func (o *GLExportTransaction) GetBillingTermsXrefOk() (*string, bool)`

GetBillingTermsXrefOk returns a tuple with the BillingTermsXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingTermsXref

`func (o *GLExportTransaction) SetBillingTermsXref(v string)`

SetBillingTermsXref sets BillingTermsXref field to given value.

### HasBillingTermsXref

`func (o *GLExportTransaction) HasBillingTermsXref() bool`

HasBillingTermsXref returns a boolean if a field has been set.

### GetDueDays

`func (o *GLExportTransaction) GetDueDays() int32`

GetDueDays returns the DueDays field if non-nil, zero value otherwise.

### GetDueDaysOk

`func (o *GLExportTransaction) GetDueDaysOk() (*int32, bool)`

GetDueDaysOk returns a tuple with the DueDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDays

`func (o *GLExportTransaction) SetDueDays(v int32)`

SetDueDays sets DueDays field to given value.

### HasDueDays

`func (o *GLExportTransaction) HasDueDays() bool`

HasDueDays returns a boolean if a field has been set.

### SetDueDaysNil

`func (o *GLExportTransaction) SetDueDaysNil(b bool)`

 SetDueDaysNil sets the value for DueDays to be an explicit nil

### UnsetDueDays
`func (o *GLExportTransaction) UnsetDueDays()`

UnsetDueDays ensures that no value is present for DueDays, not even an explicit nil
### GetDueDate

`func (o *GLExportTransaction) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *GLExportTransaction) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *GLExportTransaction) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *GLExportTransaction) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetEmailDeliveryFlag

`func (o *GLExportTransaction) GetEmailDeliveryFlag() bool`

GetEmailDeliveryFlag returns the EmailDeliveryFlag field if non-nil, zero value otherwise.

### GetEmailDeliveryFlagOk

`func (o *GLExportTransaction) GetEmailDeliveryFlagOk() (*bool, bool)`

GetEmailDeliveryFlagOk returns a tuple with the EmailDeliveryFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailDeliveryFlag

`func (o *GLExportTransaction) SetEmailDeliveryFlag(v bool)`

SetEmailDeliveryFlag sets EmailDeliveryFlag field to given value.

### HasEmailDeliveryFlag

`func (o *GLExportTransaction) HasEmailDeliveryFlag() bool`

HasEmailDeliveryFlag returns a boolean if a field has been set.

### SetEmailDeliveryFlagNil

`func (o *GLExportTransaction) SetEmailDeliveryFlagNil(b bool)`

 SetEmailDeliveryFlagNil sets the value for EmailDeliveryFlag to be an explicit nil

### UnsetEmailDeliveryFlag
`func (o *GLExportTransaction) UnsetEmailDeliveryFlag()`

UnsetEmailDeliveryFlag ensures that no value is present for EmailDeliveryFlag, not even an explicit nil
### GetPrintDeliveryFlag

`func (o *GLExportTransaction) GetPrintDeliveryFlag() bool`

GetPrintDeliveryFlag returns the PrintDeliveryFlag field if non-nil, zero value otherwise.

### GetPrintDeliveryFlagOk

`func (o *GLExportTransaction) GetPrintDeliveryFlagOk() (*bool, bool)`

GetPrintDeliveryFlagOk returns a tuple with the PrintDeliveryFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrintDeliveryFlag

`func (o *GLExportTransaction) SetPrintDeliveryFlag(v bool)`

SetPrintDeliveryFlag sets PrintDeliveryFlag field to given value.

### HasPrintDeliveryFlag

`func (o *GLExportTransaction) HasPrintDeliveryFlag() bool`

HasPrintDeliveryFlag returns a boolean if a field has been set.

### SetPrintDeliveryFlagNil

`func (o *GLExportTransaction) SetPrintDeliveryFlagNil(b bool)`

 SetPrintDeliveryFlagNil sets the value for PrintDeliveryFlag to be an explicit nil

### UnsetPrintDeliveryFlag
`func (o *GLExportTransaction) UnsetPrintDeliveryFlag()`

UnsetPrintDeliveryFlag ensures that no value is present for PrintDeliveryFlag, not even an explicit nil
### GetAgreementPrePaymentFlag

`func (o *GLExportTransaction) GetAgreementPrePaymentFlag() bool`

GetAgreementPrePaymentFlag returns the AgreementPrePaymentFlag field if non-nil, zero value otherwise.

### GetAgreementPrePaymentFlagOk

`func (o *GLExportTransaction) GetAgreementPrePaymentFlagOk() (*bool, bool)`

GetAgreementPrePaymentFlagOk returns a tuple with the AgreementPrePaymentFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreementPrePaymentFlag

`func (o *GLExportTransaction) SetAgreementPrePaymentFlag(v bool)`

SetAgreementPrePaymentFlag sets AgreementPrePaymentFlag field to given value.

### HasAgreementPrePaymentFlag

`func (o *GLExportTransaction) HasAgreementPrePaymentFlag() bool`

HasAgreementPrePaymentFlag returns a boolean if a field has been set.

### SetAgreementPrePaymentFlagNil

`func (o *GLExportTransaction) SetAgreementPrePaymentFlagNil(b bool)`

 SetAgreementPrePaymentFlagNil sets the value for AgreementPrePaymentFlag to be an explicit nil

### UnsetAgreementPrePaymentFlag
`func (o *GLExportTransaction) UnsetAgreementPrePaymentFlag()`

UnsetAgreementPrePaymentFlag ensures that no value is present for AgreementPrePaymentFlag, not even an explicit nil
### GetAccountNumber

`func (o *GLExportTransaction) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *GLExportTransaction) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *GLExportTransaction) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *GLExportTransaction) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetBillingType

`func (o *GLExportTransaction) GetBillingType() string`

GetBillingType returns the BillingType field if non-nil, zero value otherwise.

### GetBillingTypeOk

`func (o *GLExportTransaction) GetBillingTypeOk() (*string, bool)`

GetBillingTypeOk returns a tuple with the BillingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingType

`func (o *GLExportTransaction) SetBillingType(v string)`

SetBillingType sets BillingType field to given value.

### HasBillingType

`func (o *GLExportTransaction) HasBillingType() bool`

HasBillingType returns a boolean if a field has been set.

### GetGlEntryIds

`func (o *GLExportTransaction) GetGlEntryIds() string`

GetGlEntryIds returns the GlEntryIds field if non-nil, zero value otherwise.

### GetGlEntryIdsOk

`func (o *GLExportTransaction) GetGlEntryIdsOk() (*string, bool)`

GetGlEntryIdsOk returns a tuple with the GlEntryIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlEntryIds

`func (o *GLExportTransaction) SetGlEntryIds(v string)`

SetGlEntryIds sets GlEntryIds field to given value.

### HasGlEntryIds

`func (o *GLExportTransaction) HasGlEntryIds() bool`

HasGlEntryIds returns a boolean if a field has been set.

### GetPurchaseOrder

`func (o *GLExportTransaction) GetPurchaseOrder() PurchaseOrderReference`

GetPurchaseOrder returns the PurchaseOrder field if non-nil, zero value otherwise.

### GetPurchaseOrderOk

`func (o *GLExportTransaction) GetPurchaseOrderOk() (*PurchaseOrderReference, bool)`

GetPurchaseOrderOk returns a tuple with the PurchaseOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOrder

`func (o *GLExportTransaction) SetPurchaseOrder(v PurchaseOrderReference)`

SetPurchaseOrder sets PurchaseOrder field to given value.

### HasPurchaseOrder

`func (o *GLExportTransaction) HasPurchaseOrder() bool`

HasPurchaseOrder returns a boolean if a field has been set.

### GetProject

`func (o *GLExportTransaction) GetProject() ProjectReference`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *GLExportTransaction) GetProjectOk() (*ProjectReference, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *GLExportTransaction) SetProject(v ProjectReference)`

SetProject sets Project field to given value.

### HasProject

`func (o *GLExportTransaction) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetCurrency

`func (o *GLExportTransaction) GetCurrency() CurrencyReference`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GLExportTransaction) GetCurrencyOk() (*CurrencyReference, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GLExportTransaction) SetCurrency(v CurrencyReference)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *GLExportTransaction) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetTotal

`func (o *GLExportTransaction) GetTotal() float64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GLExportTransaction) GetTotalOk() (*float64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GLExportTransaction) SetTotal(v float64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GLExportTransaction) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### SetTotalNil

`func (o *GLExportTransaction) SetTotalNil(b bool)`

 SetTotalNil sets the value for Total to be an explicit nil

### UnsetTotal
`func (o *GLExportTransaction) UnsetTotal()`

UnsetTotal ensures that no value is present for Total, not even an explicit nil
### GetSalesRepId

`func (o *GLExportTransaction) GetSalesRepId() string`

GetSalesRepId returns the SalesRepId field if non-nil, zero value otherwise.

### GetSalesRepIdOk

`func (o *GLExportTransaction) GetSalesRepIdOk() (*string, bool)`

GetSalesRepIdOk returns a tuple with the SalesRepId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesRepId

`func (o *GLExportTransaction) SetSalesRepId(v string)`

SetSalesRepId sets SalesRepId field to given value.

### HasSalesRepId

`func (o *GLExportTransaction) HasSalesRepId() bool`

HasSalesRepId returns a boolean if a field has been set.

### GetSalesRepName

`func (o *GLExportTransaction) GetSalesRepName() string`

GetSalesRepName returns the SalesRepName field if non-nil, zero value otherwise.

### GetSalesRepNameOk

`func (o *GLExportTransaction) GetSalesRepNameOk() (*string, bool)`

GetSalesRepNameOk returns a tuple with the SalesRepName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesRepName

`func (o *GLExportTransaction) SetSalesRepName(v string)`

SetSalesRepName sets SalesRepName field to given value.

### HasSalesRepName

`func (o *GLExportTransaction) HasSalesRepName() bool`

HasSalesRepName returns a boolean if a field has been set.

### GetTaxable

`func (o *GLExportTransaction) GetTaxable() bool`

GetTaxable returns the Taxable field if non-nil, zero value otherwise.

### GetTaxableOk

`func (o *GLExportTransaction) GetTaxableOk() (*bool, bool)`

GetTaxableOk returns a tuple with the Taxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxable

`func (o *GLExportTransaction) SetTaxable(v bool)`

SetTaxable sets Taxable field to given value.

### HasTaxable

`func (o *GLExportTransaction) HasTaxable() bool`

HasTaxable returns a boolean if a field has been set.

### SetTaxableNil

`func (o *GLExportTransaction) SetTaxableNil(b bool)`

 SetTaxableNil sets the value for Taxable to be an explicit nil

### UnsetTaxable
`func (o *GLExportTransaction) UnsetTaxable()`

UnsetTaxable ensures that no value is present for Taxable, not even an explicit nil
### GetTaxableTotal

`func (o *GLExportTransaction) GetTaxableTotal() float64`

GetTaxableTotal returns the TaxableTotal field if non-nil, zero value otherwise.

### GetTaxableTotalOk

`func (o *GLExportTransaction) GetTaxableTotalOk() (*float64, bool)`

GetTaxableTotalOk returns a tuple with the TaxableTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxableTotal

`func (o *GLExportTransaction) SetTaxableTotal(v float64)`

SetTaxableTotal sets TaxableTotal field to given value.

### HasTaxableTotal

`func (o *GLExportTransaction) HasTaxableTotal() bool`

HasTaxableTotal returns a boolean if a field has been set.

### SetTaxableTotalNil

`func (o *GLExportTransaction) SetTaxableTotalNil(b bool)`

 SetTaxableTotalNil sets the value for TaxableTotal to be an explicit nil

### UnsetTaxableTotal
`func (o *GLExportTransaction) UnsetTaxableTotal()`

UnsetTaxableTotal ensures that no value is present for TaxableTotal, not even an explicit nil
### GetTaxCode

`func (o *GLExportTransaction) GetTaxCode() TaxCodeReference`

GetTaxCode returns the TaxCode field if non-nil, zero value otherwise.

### GetTaxCodeOk

`func (o *GLExportTransaction) GetTaxCodeOk() (*TaxCodeReference, bool)`

GetTaxCodeOk returns a tuple with the TaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCode

`func (o *GLExportTransaction) SetTaxCode(v TaxCodeReference)`

SetTaxCode sets TaxCode field to given value.

### HasTaxCode

`func (o *GLExportTransaction) HasTaxCode() bool`

HasTaxCode returns a boolean if a field has been set.

### GetTaxGroupRate

`func (o *GLExportTransaction) GetTaxGroupRate() float64`

GetTaxGroupRate returns the TaxGroupRate field if non-nil, zero value otherwise.

### GetTaxGroupRateOk

`func (o *GLExportTransaction) GetTaxGroupRateOk() (*float64, bool)`

GetTaxGroupRateOk returns a tuple with the TaxGroupRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxGroupRate

`func (o *GLExportTransaction) SetTaxGroupRate(v float64)`

SetTaxGroupRate sets TaxGroupRate field to given value.

### HasTaxGroupRate

`func (o *GLExportTransaction) HasTaxGroupRate() bool`

HasTaxGroupRate returns a boolean if a field has been set.

### SetTaxGroupRateNil

`func (o *GLExportTransaction) SetTaxGroupRateNil(b bool)`

 SetTaxGroupRateNil sets the value for TaxGroupRate to be an explicit nil

### UnsetTaxGroupRate
`func (o *GLExportTransaction) UnsetTaxGroupRate()`

UnsetTaxGroupRate ensures that no value is present for TaxGroupRate, not even an explicit nil
### GetPiggyBackFlag

`func (o *GLExportTransaction) GetPiggyBackFlag() bool`

GetPiggyBackFlag returns the PiggyBackFlag field if non-nil, zero value otherwise.

### GetPiggyBackFlagOk

`func (o *GLExportTransaction) GetPiggyBackFlagOk() (*bool, bool)`

GetPiggyBackFlagOk returns a tuple with the PiggyBackFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPiggyBackFlag

`func (o *GLExportTransaction) SetPiggyBackFlag(v bool)`

SetPiggyBackFlag sets PiggyBackFlag field to given value.

### HasPiggyBackFlag

`func (o *GLExportTransaction) HasPiggyBackFlag() bool`

HasPiggyBackFlag returns a boolean if a field has been set.

### SetPiggyBackFlagNil

`func (o *GLExportTransaction) SetPiggyBackFlagNil(b bool)`

 SetPiggyBackFlagNil sets the value for PiggyBackFlag to be an explicit nil

### UnsetPiggyBackFlag
`func (o *GLExportTransaction) UnsetPiggyBackFlag()`

UnsetPiggyBackFlag ensures that no value is present for PiggyBackFlag, not even an explicit nil
### GetTaxAccountNumber

`func (o *GLExportTransaction) GetTaxAccountNumber() string`

GetTaxAccountNumber returns the TaxAccountNumber field if non-nil, zero value otherwise.

### GetTaxAccountNumberOk

`func (o *GLExportTransaction) GetTaxAccountNumberOk() (*string, bool)`

GetTaxAccountNumberOk returns a tuple with the TaxAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAccountNumber

`func (o *GLExportTransaction) SetTaxAccountNumber(v string)`

SetTaxAccountNumber sets TaxAccountNumber field to given value.

### HasTaxAccountNumber

`func (o *GLExportTransaction) HasTaxAccountNumber() bool`

HasTaxAccountNumber returns a boolean if a field has been set.

### GetSalesTax

`func (o *GLExportTransaction) GetSalesTax() float64`

GetSalesTax returns the SalesTax field if non-nil, zero value otherwise.

### GetSalesTaxOk

`func (o *GLExportTransaction) GetSalesTaxOk() (*float64, bool)`

GetSalesTaxOk returns a tuple with the SalesTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTax

`func (o *GLExportTransaction) SetSalesTax(v float64)`

SetSalesTax sets SalesTax field to given value.

### HasSalesTax

`func (o *GLExportTransaction) HasSalesTax() bool`

HasSalesTax returns a boolean if a field has been set.

### SetSalesTaxNil

`func (o *GLExportTransaction) SetSalesTaxNil(b bool)`

 SetSalesTaxNil sets the value for SalesTax to be an explicit nil

### UnsetSalesTax
`func (o *GLExportTransaction) UnsetSalesTax()`

UnsetSalesTax ensures that no value is present for SalesTax, not even an explicit nil
### GetStateTax

`func (o *GLExportTransaction) GetStateTax() float64`

GetStateTax returns the StateTax field if non-nil, zero value otherwise.

### GetStateTaxOk

`func (o *GLExportTransaction) GetStateTaxOk() (*float64, bool)`

GetStateTaxOk returns a tuple with the StateTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateTax

`func (o *GLExportTransaction) SetStateTax(v float64)`

SetStateTax sets StateTax field to given value.

### HasStateTax

`func (o *GLExportTransaction) HasStateTax() bool`

HasStateTax returns a boolean if a field has been set.

### SetStateTaxNil

`func (o *GLExportTransaction) SetStateTaxNil(b bool)`

 SetStateTaxNil sets the value for StateTax to be an explicit nil

### UnsetStateTax
`func (o *GLExportTransaction) UnsetStateTax()`

UnsetStateTax ensures that no value is present for StateTax, not even an explicit nil
### GetCountyTax

`func (o *GLExportTransaction) GetCountyTax() float64`

GetCountyTax returns the CountyTax field if non-nil, zero value otherwise.

### GetCountyTaxOk

`func (o *GLExportTransaction) GetCountyTaxOk() (*float64, bool)`

GetCountyTaxOk returns a tuple with the CountyTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountyTax

`func (o *GLExportTransaction) SetCountyTax(v float64)`

SetCountyTax sets CountyTax field to given value.

### HasCountyTax

`func (o *GLExportTransaction) HasCountyTax() bool`

HasCountyTax returns a boolean if a field has been set.

### SetCountyTaxNil

`func (o *GLExportTransaction) SetCountyTaxNil(b bool)`

 SetCountyTaxNil sets the value for CountyTax to be an explicit nil

### UnsetCountyTax
`func (o *GLExportTransaction) UnsetCountyTax()`

UnsetCountyTax ensures that no value is present for CountyTax, not even an explicit nil
### GetCityTax

`func (o *GLExportTransaction) GetCityTax() float64`

GetCityTax returns the CityTax field if non-nil, zero value otherwise.

### GetCityTaxOk

`func (o *GLExportTransaction) GetCityTaxOk() (*float64, bool)`

GetCityTaxOk returns a tuple with the CityTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityTax

`func (o *GLExportTransaction) SetCityTax(v float64)`

SetCityTax sets CityTax field to given value.

### HasCityTax

`func (o *GLExportTransaction) HasCityTax() bool`

HasCityTax returns a boolean if a field has been set.

### SetCityTaxNil

`func (o *GLExportTransaction) SetCityTaxNil(b bool)`

 SetCityTaxNil sets the value for CityTax to be an explicit nil

### UnsetCityTax
`func (o *GLExportTransaction) UnsetCityTax()`

UnsetCityTax ensures that no value is present for CityTax, not even an explicit nil
### GetTaxableAmount1

`func (o *GLExportTransaction) GetTaxableAmount1() float64`

GetTaxableAmount1 returns the TaxableAmount1 field if non-nil, zero value otherwise.

### GetTaxableAmount1Ok

`func (o *GLExportTransaction) GetTaxableAmount1Ok() (*float64, bool)`

GetTaxableAmount1Ok returns a tuple with the TaxableAmount1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxableAmount1

`func (o *GLExportTransaction) SetTaxableAmount1(v float64)`

SetTaxableAmount1 sets TaxableAmount1 field to given value.

### HasTaxableAmount1

`func (o *GLExportTransaction) HasTaxableAmount1() bool`

HasTaxableAmount1 returns a boolean if a field has been set.

### SetTaxableAmount1Nil

`func (o *GLExportTransaction) SetTaxableAmount1Nil(b bool)`

 SetTaxableAmount1Nil sets the value for TaxableAmount1 to be an explicit nil

### UnsetTaxableAmount1
`func (o *GLExportTransaction) UnsetTaxableAmount1()`

UnsetTaxableAmount1 ensures that no value is present for TaxableAmount1, not even an explicit nil
### GetTaxableAmount2

`func (o *GLExportTransaction) GetTaxableAmount2() float64`

GetTaxableAmount2 returns the TaxableAmount2 field if non-nil, zero value otherwise.

### GetTaxableAmount2Ok

`func (o *GLExportTransaction) GetTaxableAmount2Ok() (*float64, bool)`

GetTaxableAmount2Ok returns a tuple with the TaxableAmount2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxableAmount2

`func (o *GLExportTransaction) SetTaxableAmount2(v float64)`

SetTaxableAmount2 sets TaxableAmount2 field to given value.

### HasTaxableAmount2

`func (o *GLExportTransaction) HasTaxableAmount2() bool`

HasTaxableAmount2 returns a boolean if a field has been set.

### SetTaxableAmount2Nil

`func (o *GLExportTransaction) SetTaxableAmount2Nil(b bool)`

 SetTaxableAmount2Nil sets the value for TaxableAmount2 to be an explicit nil

### UnsetTaxableAmount2
`func (o *GLExportTransaction) UnsetTaxableAmount2()`

UnsetTaxableAmount2 ensures that no value is present for TaxableAmount2, not even an explicit nil
### GetTaxableAmount3

`func (o *GLExportTransaction) GetTaxableAmount3() float64`

GetTaxableAmount3 returns the TaxableAmount3 field if non-nil, zero value otherwise.

### GetTaxableAmount3Ok

`func (o *GLExportTransaction) GetTaxableAmount3Ok() (*float64, bool)`

GetTaxableAmount3Ok returns a tuple with the TaxableAmount3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxableAmount3

`func (o *GLExportTransaction) SetTaxableAmount3(v float64)`

SetTaxableAmount3 sets TaxableAmount3 field to given value.

### HasTaxableAmount3

`func (o *GLExportTransaction) HasTaxableAmount3() bool`

HasTaxableAmount3 returns a boolean if a field has been set.

### SetTaxableAmount3Nil

`func (o *GLExportTransaction) SetTaxableAmount3Nil(b bool)`

 SetTaxableAmount3Nil sets the value for TaxableAmount3 to be an explicit nil

### UnsetTaxableAmount3
`func (o *GLExportTransaction) UnsetTaxableAmount3()`

UnsetTaxableAmount3 ensures that no value is present for TaxableAmount3, not even an explicit nil
### GetTaxableAmount4

`func (o *GLExportTransaction) GetTaxableAmount4() float64`

GetTaxableAmount4 returns the TaxableAmount4 field if non-nil, zero value otherwise.

### GetTaxableAmount4Ok

`func (o *GLExportTransaction) GetTaxableAmount4Ok() (*float64, bool)`

GetTaxableAmount4Ok returns a tuple with the TaxableAmount4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxableAmount4

`func (o *GLExportTransaction) SetTaxableAmount4(v float64)`

SetTaxableAmount4 sets TaxableAmount4 field to given value.

### HasTaxableAmount4

`func (o *GLExportTransaction) HasTaxableAmount4() bool`

HasTaxableAmount4 returns a boolean if a field has been set.

### SetTaxableAmount4Nil

`func (o *GLExportTransaction) SetTaxableAmount4Nil(b bool)`

 SetTaxableAmount4Nil sets the value for TaxableAmount4 to be an explicit nil

### UnsetTaxableAmount4
`func (o *GLExportTransaction) UnsetTaxableAmount4()`

UnsetTaxableAmount4 ensures that no value is present for TaxableAmount4, not even an explicit nil
### GetTaxableAmount5

`func (o *GLExportTransaction) GetTaxableAmount5() float64`

GetTaxableAmount5 returns the TaxableAmount5 field if non-nil, zero value otherwise.

### GetTaxableAmount5Ok

`func (o *GLExportTransaction) GetTaxableAmount5Ok() (*float64, bool)`

GetTaxableAmount5Ok returns a tuple with the TaxableAmount5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxableAmount5

`func (o *GLExportTransaction) SetTaxableAmount5(v float64)`

SetTaxableAmount5 sets TaxableAmount5 field to given value.

### HasTaxableAmount5

`func (o *GLExportTransaction) HasTaxableAmount5() bool`

HasTaxableAmount5 returns a boolean if a field has been set.

### SetTaxableAmount5Nil

`func (o *GLExportTransaction) SetTaxableAmount5Nil(b bool)`

 SetTaxableAmount5Nil sets the value for TaxableAmount5 to be an explicit nil

### UnsetTaxableAmount5
`func (o *GLExportTransaction) UnsetTaxableAmount5()`

UnsetTaxableAmount5 ensures that no value is present for TaxableAmount5, not even an explicit nil
### GetTaxAgencyXref

`func (o *GLExportTransaction) GetTaxAgencyXref() string`

GetTaxAgencyXref returns the TaxAgencyXref field if non-nil, zero value otherwise.

### GetTaxAgencyXrefOk

`func (o *GLExportTransaction) GetTaxAgencyXrefOk() (*string, bool)`

GetTaxAgencyXrefOk returns a tuple with the TaxAgencyXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAgencyXref

`func (o *GLExportTransaction) SetTaxAgencyXref(v string)`

SetTaxAgencyXref sets TaxAgencyXref field to given value.

### HasTaxAgencyXref

`func (o *GLExportTransaction) HasTaxAgencyXref() bool`

HasTaxAgencyXref returns a boolean if a field has been set.

### GetStateTaxXref

`func (o *GLExportTransaction) GetStateTaxXref() string`

GetStateTaxXref returns the StateTaxXref field if non-nil, zero value otherwise.

### GetStateTaxXrefOk

`func (o *GLExportTransaction) GetStateTaxXrefOk() (*string, bool)`

GetStateTaxXrefOk returns a tuple with the StateTaxXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateTaxXref

`func (o *GLExportTransaction) SetStateTaxXref(v string)`

SetStateTaxXref sets StateTaxXref field to given value.

### HasStateTaxXref

`func (o *GLExportTransaction) HasStateTaxXref() bool`

HasStateTaxXref returns a boolean if a field has been set.

### GetCountyTaxXref

`func (o *GLExportTransaction) GetCountyTaxXref() string`

GetCountyTaxXref returns the CountyTaxXref field if non-nil, zero value otherwise.

### GetCountyTaxXrefOk

`func (o *GLExportTransaction) GetCountyTaxXrefOk() (*string, bool)`

GetCountyTaxXrefOk returns a tuple with the CountyTaxXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountyTaxXref

`func (o *GLExportTransaction) SetCountyTaxXref(v string)`

SetCountyTaxXref sets CountyTaxXref field to given value.

### HasCountyTaxXref

`func (o *GLExportTransaction) HasCountyTaxXref() bool`

HasCountyTaxXref returns a boolean if a field has been set.

### GetTaxId

`func (o *GLExportTransaction) GetTaxId() string`

GetTaxId returns the TaxId field if non-nil, zero value otherwise.

### GetTaxIdOk

`func (o *GLExportTransaction) GetTaxIdOk() (*string, bool)`

GetTaxIdOk returns a tuple with the TaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxId

`func (o *GLExportTransaction) SetTaxId(v string)`

SetTaxId sets TaxId field to given value.

### HasTaxId

`func (o *GLExportTransaction) HasTaxId() bool`

HasTaxId returns a boolean if a field has been set.

### GetTaxDpAppliedFlag

`func (o *GLExportTransaction) GetTaxDpAppliedFlag() bool`

GetTaxDpAppliedFlag returns the TaxDpAppliedFlag field if non-nil, zero value otherwise.

### GetTaxDpAppliedFlagOk

`func (o *GLExportTransaction) GetTaxDpAppliedFlagOk() (*bool, bool)`

GetTaxDpAppliedFlagOk returns a tuple with the TaxDpAppliedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxDpAppliedFlag

`func (o *GLExportTransaction) SetTaxDpAppliedFlag(v bool)`

SetTaxDpAppliedFlag sets TaxDpAppliedFlag field to given value.

### HasTaxDpAppliedFlag

`func (o *GLExportTransaction) HasTaxDpAppliedFlag() bool`

HasTaxDpAppliedFlag returns a boolean if a field has been set.

### SetTaxDpAppliedFlagNil

`func (o *GLExportTransaction) SetTaxDpAppliedFlagNil(b bool)`

 SetTaxDpAppliedFlagNil sets the value for TaxDpAppliedFlag to be an explicit nil

### UnsetTaxDpAppliedFlag
`func (o *GLExportTransaction) UnsetTaxDpAppliedFlag()`

UnsetTaxDpAppliedFlag ensures that no value is present for TaxDpAppliedFlag, not even an explicit nil
### GetUseAvalaraFlag

`func (o *GLExportTransaction) GetUseAvalaraFlag() bool`

GetUseAvalaraFlag returns the UseAvalaraFlag field if non-nil, zero value otherwise.

### GetUseAvalaraFlagOk

`func (o *GLExportTransaction) GetUseAvalaraFlagOk() (*bool, bool)`

GetUseAvalaraFlagOk returns a tuple with the UseAvalaraFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAvalaraFlag

`func (o *GLExportTransaction) SetUseAvalaraFlag(v bool)`

SetUseAvalaraFlag sets UseAvalaraFlag field to given value.

### HasUseAvalaraFlag

`func (o *GLExportTransaction) HasUseAvalaraFlag() bool`

HasUseAvalaraFlag returns a boolean if a field has been set.

### SetUseAvalaraFlagNil

`func (o *GLExportTransaction) SetUseAvalaraFlagNil(b bool)`

 SetUseAvalaraFlagNil sets the value for UseAvalaraFlag to be an explicit nil

### UnsetUseAvalaraFlag
`func (o *GLExportTransaction) UnsetUseAvalaraFlag()`

UnsetUseAvalaraFlag ensures that no value is present for UseAvalaraFlag, not even an explicit nil
### GetSendAvalaraTaxFlag

`func (o *GLExportTransaction) GetSendAvalaraTaxFlag() bool`

GetSendAvalaraTaxFlag returns the SendAvalaraTaxFlag field if non-nil, zero value otherwise.

### GetSendAvalaraTaxFlagOk

`func (o *GLExportTransaction) GetSendAvalaraTaxFlagOk() (*bool, bool)`

GetSendAvalaraTaxFlagOk returns a tuple with the SendAvalaraTaxFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendAvalaraTaxFlag

`func (o *GLExportTransaction) SetSendAvalaraTaxFlag(v bool)`

SetSendAvalaraTaxFlag sets SendAvalaraTaxFlag field to given value.

### HasSendAvalaraTaxFlag

`func (o *GLExportTransaction) HasSendAvalaraTaxFlag() bool`

HasSendAvalaraTaxFlag returns a boolean if a field has been set.

### SetSendAvalaraTaxFlagNil

`func (o *GLExportTransaction) SetSendAvalaraTaxFlagNil(b bool)`

 SetSendAvalaraTaxFlagNil sets the value for SendAvalaraTaxFlag to be an explicit nil

### UnsetSendAvalaraTaxFlag
`func (o *GLExportTransaction) UnsetSendAvalaraTaxFlag()`

UnsetSendAvalaraTaxFlag ensures that no value is present for SendAvalaraTaxFlag, not even an explicit nil
### GetShipToCompany

`func (o *GLExportTransaction) GetShipToCompany() CompanyReference`

GetShipToCompany returns the ShipToCompany field if non-nil, zero value otherwise.

### GetShipToCompanyOk

`func (o *GLExportTransaction) GetShipToCompanyOk() (*CompanyReference, bool)`

GetShipToCompanyOk returns a tuple with the ShipToCompany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipToCompany

`func (o *GLExportTransaction) SetShipToCompany(v CompanyReference)`

SetShipToCompany sets ShipToCompany field to given value.

### HasShipToCompany

`func (o *GLExportTransaction) HasShipToCompany() bool`

HasShipToCompany returns a boolean if a field has been set.

### GetShipToCompanyAccountNumber

`func (o *GLExportTransaction) GetShipToCompanyAccountNumber() string`

GetShipToCompanyAccountNumber returns the ShipToCompanyAccountNumber field if non-nil, zero value otherwise.

### GetShipToCompanyAccountNumberOk

`func (o *GLExportTransaction) GetShipToCompanyAccountNumberOk() (*string, bool)`

GetShipToCompanyAccountNumberOk returns a tuple with the ShipToCompanyAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipToCompanyAccountNumber

`func (o *GLExportTransaction) SetShipToCompanyAccountNumber(v string)`

SetShipToCompanyAccountNumber sets ShipToCompanyAccountNumber field to given value.

### HasShipToCompanyAccountNumber

`func (o *GLExportTransaction) HasShipToCompanyAccountNumber() bool`

HasShipToCompanyAccountNumber returns a boolean if a field has been set.

### GetShipToCompanyType

`func (o *GLExportTransaction) GetShipToCompanyType() CompanyTypeReference`

GetShipToCompanyType returns the ShipToCompanyType field if non-nil, zero value otherwise.

### GetShipToCompanyTypeOk

`func (o *GLExportTransaction) GetShipToCompanyTypeOk() (*CompanyTypeReference, bool)`

GetShipToCompanyTypeOk returns a tuple with the ShipToCompanyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipToCompanyType

`func (o *GLExportTransaction) SetShipToCompanyType(v CompanyTypeReference)`

SetShipToCompanyType sets ShipToCompanyType field to given value.

### HasShipToCompanyType

`func (o *GLExportTransaction) HasShipToCompanyType() bool`

HasShipToCompanyType returns a boolean if a field has been set.

### GetShipToTaxId

`func (o *GLExportTransaction) GetShipToTaxId() string`

GetShipToTaxId returns the ShipToTaxId field if non-nil, zero value otherwise.

### GetShipToTaxIdOk

`func (o *GLExportTransaction) GetShipToTaxIdOk() (*string, bool)`

GetShipToTaxIdOk returns a tuple with the ShipToTaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipToTaxId

`func (o *GLExportTransaction) SetShipToTaxId(v string)`

SetShipToTaxId sets ShipToTaxId field to given value.

### HasShipToTaxId

`func (o *GLExportTransaction) HasShipToTaxId() bool`

HasShipToTaxId returns a boolean if a field has been set.

### GetShipSite

`func (o *GLExportTransaction) GetShipSite() SiteReference`

GetShipSite returns the ShipSite field if non-nil, zero value otherwise.

### GetShipSiteOk

`func (o *GLExportTransaction) GetShipSiteOk() (*SiteReference, bool)`

GetShipSiteOk returns a tuple with the ShipSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipSite

`func (o *GLExportTransaction) SetShipSite(v SiteReference)`

SetShipSite sets ShipSite field to given value.

### HasShipSite

`func (o *GLExportTransaction) HasShipSite() bool`

HasShipSite returns a boolean if a field has been set.

### GetShipContact

`func (o *GLExportTransaction) GetShipContact() string`

GetShipContact returns the ShipContact field if non-nil, zero value otherwise.

### GetShipContactOk

`func (o *GLExportTransaction) GetShipContactOk() (*string, bool)`

GetShipContactOk returns a tuple with the ShipContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipContact

`func (o *GLExportTransaction) SetShipContact(v string)`

SetShipContact sets ShipContact field to given value.

### HasShipContact

`func (o *GLExportTransaction) HasShipContact() bool`

HasShipContact returns a boolean if a field has been set.

### GetDetail

`func (o *GLExportTransaction) GetDetail() []GLExportTransactionDetail`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *GLExportTransaction) GetDetailOk() (*[]GLExportTransactionDetail, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *GLExportTransaction) SetDetail(v []GLExportTransactionDetail)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *GLExportTransaction) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetTaxLevels

`func (o *GLExportTransaction) GetTaxLevels() []GLExportTransactionTaxLevel`

GetTaxLevels returns the TaxLevels field if non-nil, zero value otherwise.

### GetTaxLevelsOk

`func (o *GLExportTransaction) GetTaxLevelsOk() (*[]GLExportTransactionTaxLevel, bool)`

GetTaxLevelsOk returns a tuple with the TaxLevels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxLevels

`func (o *GLExportTransaction) SetTaxLevels(v []GLExportTransactionTaxLevel)`

SetTaxLevels sets TaxLevels field to given value.

### HasTaxLevels

`func (o *GLExportTransaction) HasTaxLevels() bool`

HasTaxLevels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


