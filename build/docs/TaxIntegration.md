# TaxIntegration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaxIntegrationType** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**AccountNumber** | Pointer to **string** |  Max length: 50; | [optional] 
**LicenseKey** | Pointer to **string** |  Max length: 50; | [optional] 
**ServiceUrl** | Pointer to **string** |  Max length: 250; | [optional] 
**CompanyCode** | Pointer to **string** |  Max length: 50; | [optional] 
**TimeTaxCode** | Pointer to **string** |  Max length: 50; | [optional] 
**ExpenseTaxCode** | Pointer to **string** |  Max length: 50; | [optional] 
**ProductTaxCode** | Pointer to **string** |  Max length: 50; | [optional] 
**InvoiceAmountTaxCode** | Pointer to **string** |  Max length: 50; | [optional] 
**EnabledFlag** | Pointer to **NullableBool** |  | [optional] 
**CommitTransactionsFlag** | Pointer to **NullableBool** |  | [optional] 
**SalesInvoiceFlag** | Pointer to **NullableBool** |  | [optional] 
**FreightTaxCode** | Pointer to **string** |  Max length: 50; | [optional] 
**AccountingIntegrationFlag** | Pointer to **NullableBool** |  | [optional] 
**TaxLineFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewTaxIntegration

`func NewTaxIntegration() *TaxIntegration`

NewTaxIntegration instantiates a new TaxIntegration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxIntegrationWithDefaults

`func NewTaxIntegrationWithDefaults() *TaxIntegration`

NewTaxIntegrationWithDefaults instantiates a new TaxIntegration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaxIntegrationType

`func (o *TaxIntegration) GetTaxIntegrationType() string`

GetTaxIntegrationType returns the TaxIntegrationType field if non-nil, zero value otherwise.

### GetTaxIntegrationTypeOk

`func (o *TaxIntegration) GetTaxIntegrationTypeOk() (*string, bool)`

GetTaxIntegrationTypeOk returns a tuple with the TaxIntegrationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxIntegrationType

`func (o *TaxIntegration) SetTaxIntegrationType(v string)`

SetTaxIntegrationType sets TaxIntegrationType field to given value.

### HasTaxIntegrationType

`func (o *TaxIntegration) HasTaxIntegrationType() bool`

HasTaxIntegrationType returns a boolean if a field has been set.

### GetId

`func (o *TaxIntegration) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaxIntegration) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaxIntegration) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TaxIntegration) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccountNumber

`func (o *TaxIntegration) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *TaxIntegration) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *TaxIntegration) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *TaxIntegration) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetLicenseKey

`func (o *TaxIntegration) GetLicenseKey() string`

GetLicenseKey returns the LicenseKey field if non-nil, zero value otherwise.

### GetLicenseKeyOk

`func (o *TaxIntegration) GetLicenseKeyOk() (*string, bool)`

GetLicenseKeyOk returns a tuple with the LicenseKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseKey

`func (o *TaxIntegration) SetLicenseKey(v string)`

SetLicenseKey sets LicenseKey field to given value.

### HasLicenseKey

`func (o *TaxIntegration) HasLicenseKey() bool`

HasLicenseKey returns a boolean if a field has been set.

### GetServiceUrl

`func (o *TaxIntegration) GetServiceUrl() string`

GetServiceUrl returns the ServiceUrl field if non-nil, zero value otherwise.

### GetServiceUrlOk

`func (o *TaxIntegration) GetServiceUrlOk() (*string, bool)`

GetServiceUrlOk returns a tuple with the ServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUrl

`func (o *TaxIntegration) SetServiceUrl(v string)`

SetServiceUrl sets ServiceUrl field to given value.

### HasServiceUrl

`func (o *TaxIntegration) HasServiceUrl() bool`

HasServiceUrl returns a boolean if a field has been set.

### GetCompanyCode

`func (o *TaxIntegration) GetCompanyCode() string`

GetCompanyCode returns the CompanyCode field if non-nil, zero value otherwise.

### GetCompanyCodeOk

`func (o *TaxIntegration) GetCompanyCodeOk() (*string, bool)`

GetCompanyCodeOk returns a tuple with the CompanyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyCode

`func (o *TaxIntegration) SetCompanyCode(v string)`

SetCompanyCode sets CompanyCode field to given value.

### HasCompanyCode

`func (o *TaxIntegration) HasCompanyCode() bool`

HasCompanyCode returns a boolean if a field has been set.

### GetTimeTaxCode

`func (o *TaxIntegration) GetTimeTaxCode() string`

GetTimeTaxCode returns the TimeTaxCode field if non-nil, zero value otherwise.

### GetTimeTaxCodeOk

`func (o *TaxIntegration) GetTimeTaxCodeOk() (*string, bool)`

GetTimeTaxCodeOk returns a tuple with the TimeTaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeTaxCode

`func (o *TaxIntegration) SetTimeTaxCode(v string)`

SetTimeTaxCode sets TimeTaxCode field to given value.

### HasTimeTaxCode

`func (o *TaxIntegration) HasTimeTaxCode() bool`

HasTimeTaxCode returns a boolean if a field has been set.

### GetExpenseTaxCode

`func (o *TaxIntegration) GetExpenseTaxCode() string`

GetExpenseTaxCode returns the ExpenseTaxCode field if non-nil, zero value otherwise.

### GetExpenseTaxCodeOk

`func (o *TaxIntegration) GetExpenseTaxCodeOk() (*string, bool)`

GetExpenseTaxCodeOk returns a tuple with the ExpenseTaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenseTaxCode

`func (o *TaxIntegration) SetExpenseTaxCode(v string)`

SetExpenseTaxCode sets ExpenseTaxCode field to given value.

### HasExpenseTaxCode

`func (o *TaxIntegration) HasExpenseTaxCode() bool`

HasExpenseTaxCode returns a boolean if a field has been set.

### GetProductTaxCode

`func (o *TaxIntegration) GetProductTaxCode() string`

GetProductTaxCode returns the ProductTaxCode field if non-nil, zero value otherwise.

### GetProductTaxCodeOk

`func (o *TaxIntegration) GetProductTaxCodeOk() (*string, bool)`

GetProductTaxCodeOk returns a tuple with the ProductTaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTaxCode

`func (o *TaxIntegration) SetProductTaxCode(v string)`

SetProductTaxCode sets ProductTaxCode field to given value.

### HasProductTaxCode

`func (o *TaxIntegration) HasProductTaxCode() bool`

HasProductTaxCode returns a boolean if a field has been set.

### GetInvoiceAmountTaxCode

`func (o *TaxIntegration) GetInvoiceAmountTaxCode() string`

GetInvoiceAmountTaxCode returns the InvoiceAmountTaxCode field if non-nil, zero value otherwise.

### GetInvoiceAmountTaxCodeOk

`func (o *TaxIntegration) GetInvoiceAmountTaxCodeOk() (*string, bool)`

GetInvoiceAmountTaxCodeOk returns a tuple with the InvoiceAmountTaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceAmountTaxCode

`func (o *TaxIntegration) SetInvoiceAmountTaxCode(v string)`

SetInvoiceAmountTaxCode sets InvoiceAmountTaxCode field to given value.

### HasInvoiceAmountTaxCode

`func (o *TaxIntegration) HasInvoiceAmountTaxCode() bool`

HasInvoiceAmountTaxCode returns a boolean if a field has been set.

### GetEnabledFlag

`func (o *TaxIntegration) GetEnabledFlag() bool`

GetEnabledFlag returns the EnabledFlag field if non-nil, zero value otherwise.

### GetEnabledFlagOk

`func (o *TaxIntegration) GetEnabledFlagOk() (*bool, bool)`

GetEnabledFlagOk returns a tuple with the EnabledFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledFlag

`func (o *TaxIntegration) SetEnabledFlag(v bool)`

SetEnabledFlag sets EnabledFlag field to given value.

### HasEnabledFlag

`func (o *TaxIntegration) HasEnabledFlag() bool`

HasEnabledFlag returns a boolean if a field has been set.

### SetEnabledFlagNil

`func (o *TaxIntegration) SetEnabledFlagNil(b bool)`

 SetEnabledFlagNil sets the value for EnabledFlag to be an explicit nil

### UnsetEnabledFlag
`func (o *TaxIntegration) UnsetEnabledFlag()`

UnsetEnabledFlag ensures that no value is present for EnabledFlag, not even an explicit nil
### GetCommitTransactionsFlag

`func (o *TaxIntegration) GetCommitTransactionsFlag() bool`

GetCommitTransactionsFlag returns the CommitTransactionsFlag field if non-nil, zero value otherwise.

### GetCommitTransactionsFlagOk

`func (o *TaxIntegration) GetCommitTransactionsFlagOk() (*bool, bool)`

GetCommitTransactionsFlagOk returns a tuple with the CommitTransactionsFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitTransactionsFlag

`func (o *TaxIntegration) SetCommitTransactionsFlag(v bool)`

SetCommitTransactionsFlag sets CommitTransactionsFlag field to given value.

### HasCommitTransactionsFlag

`func (o *TaxIntegration) HasCommitTransactionsFlag() bool`

HasCommitTransactionsFlag returns a boolean if a field has been set.

### SetCommitTransactionsFlagNil

`func (o *TaxIntegration) SetCommitTransactionsFlagNil(b bool)`

 SetCommitTransactionsFlagNil sets the value for CommitTransactionsFlag to be an explicit nil

### UnsetCommitTransactionsFlag
`func (o *TaxIntegration) UnsetCommitTransactionsFlag()`

UnsetCommitTransactionsFlag ensures that no value is present for CommitTransactionsFlag, not even an explicit nil
### GetSalesInvoiceFlag

`func (o *TaxIntegration) GetSalesInvoiceFlag() bool`

GetSalesInvoiceFlag returns the SalesInvoiceFlag field if non-nil, zero value otherwise.

### GetSalesInvoiceFlagOk

`func (o *TaxIntegration) GetSalesInvoiceFlagOk() (*bool, bool)`

GetSalesInvoiceFlagOk returns a tuple with the SalesInvoiceFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesInvoiceFlag

`func (o *TaxIntegration) SetSalesInvoiceFlag(v bool)`

SetSalesInvoiceFlag sets SalesInvoiceFlag field to given value.

### HasSalesInvoiceFlag

`func (o *TaxIntegration) HasSalesInvoiceFlag() bool`

HasSalesInvoiceFlag returns a boolean if a field has been set.

### SetSalesInvoiceFlagNil

`func (o *TaxIntegration) SetSalesInvoiceFlagNil(b bool)`

 SetSalesInvoiceFlagNil sets the value for SalesInvoiceFlag to be an explicit nil

### UnsetSalesInvoiceFlag
`func (o *TaxIntegration) UnsetSalesInvoiceFlag()`

UnsetSalesInvoiceFlag ensures that no value is present for SalesInvoiceFlag, not even an explicit nil
### GetFreightTaxCode

`func (o *TaxIntegration) GetFreightTaxCode() string`

GetFreightTaxCode returns the FreightTaxCode field if non-nil, zero value otherwise.

### GetFreightTaxCodeOk

`func (o *TaxIntegration) GetFreightTaxCodeOk() (*string, bool)`

GetFreightTaxCodeOk returns a tuple with the FreightTaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreightTaxCode

`func (o *TaxIntegration) SetFreightTaxCode(v string)`

SetFreightTaxCode sets FreightTaxCode field to given value.

### HasFreightTaxCode

`func (o *TaxIntegration) HasFreightTaxCode() bool`

HasFreightTaxCode returns a boolean if a field has been set.

### GetAccountingIntegrationFlag

`func (o *TaxIntegration) GetAccountingIntegrationFlag() bool`

GetAccountingIntegrationFlag returns the AccountingIntegrationFlag field if non-nil, zero value otherwise.

### GetAccountingIntegrationFlagOk

`func (o *TaxIntegration) GetAccountingIntegrationFlagOk() (*bool, bool)`

GetAccountingIntegrationFlagOk returns a tuple with the AccountingIntegrationFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountingIntegrationFlag

`func (o *TaxIntegration) SetAccountingIntegrationFlag(v bool)`

SetAccountingIntegrationFlag sets AccountingIntegrationFlag field to given value.

### HasAccountingIntegrationFlag

`func (o *TaxIntegration) HasAccountingIntegrationFlag() bool`

HasAccountingIntegrationFlag returns a boolean if a field has been set.

### SetAccountingIntegrationFlagNil

`func (o *TaxIntegration) SetAccountingIntegrationFlagNil(b bool)`

 SetAccountingIntegrationFlagNil sets the value for AccountingIntegrationFlag to be an explicit nil

### UnsetAccountingIntegrationFlag
`func (o *TaxIntegration) UnsetAccountingIntegrationFlag()`

UnsetAccountingIntegrationFlag ensures that no value is present for AccountingIntegrationFlag, not even an explicit nil
### GetTaxLineFlag

`func (o *TaxIntegration) GetTaxLineFlag() bool`

GetTaxLineFlag returns the TaxLineFlag field if non-nil, zero value otherwise.

### GetTaxLineFlagOk

`func (o *TaxIntegration) GetTaxLineFlagOk() (*bool, bool)`

GetTaxLineFlagOk returns a tuple with the TaxLineFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxLineFlag

`func (o *TaxIntegration) SetTaxLineFlag(v bool)`

SetTaxLineFlag sets TaxLineFlag field to given value.

### HasTaxLineFlag

`func (o *TaxIntegration) HasTaxLineFlag() bool`

HasTaxLineFlag returns a boolean if a field has been set.

### SetTaxLineFlagNil

`func (o *TaxIntegration) SetTaxLineFlagNil(b bool)`

 SetTaxLineFlagNil sets the value for TaxLineFlag to be an explicit nil

### UnsetTaxLineFlag
`func (o *TaxIntegration) UnsetTaxLineFlag()`

UnsetTaxLineFlag ensures that no value is present for TaxLineFlag, not even an explicit nil
### GetInfo

`func (o *TaxIntegration) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *TaxIntegration) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *TaxIntegration) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *TaxIntegration) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


