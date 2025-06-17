# InvoiceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Invoice** | Pointer to [**Invoice**](Invoice.md) |  | [optional] 
**InvoiceTemplate** | Pointer to [**InvoiceTemplate**](InvoiceTemplate.md) |  | [optional] 
**Products** | Pointer to [**[]ProductItem**](ProductItem.md) |  | [optional] 
**BundledComponentsInfo** | Pointer to [**[]ProductComponent**](ProductComponent.md) |  | [optional] 
**Expenses** | Pointer to [**[]ExpenseEntry**](ExpenseEntry.md) |  | [optional] 
**TimeEntries** | Pointer to [**[]TimeEntry**](TimeEntry.md) |  | [optional] 
**Logo** | Pointer to [**DocumentInfo**](DocumentInfo.md) |  | [optional] 
**BillingSetup** | Pointer to [**BillingSetup**](BillingSetup.md) |  | [optional] 
**AgreementBillingInfo** | Pointer to [**[]AgreementBillingInfo**](AgreementBillingInfo.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewInvoiceInfo

`func NewInvoiceInfo() *InvoiceInfo`

NewInvoiceInfo instantiates a new InvoiceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceInfoWithDefaults

`func NewInvoiceInfoWithDefaults() *InvoiceInfo`

NewInvoiceInfoWithDefaults instantiates a new InvoiceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InvoiceInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvoiceInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvoiceInfo) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *InvoiceInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInvoice

`func (o *InvoiceInfo) GetInvoice() Invoice`

GetInvoice returns the Invoice field if non-nil, zero value otherwise.

### GetInvoiceOk

`func (o *InvoiceInfo) GetInvoiceOk() (*Invoice, bool)`

GetInvoiceOk returns a tuple with the Invoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoice

`func (o *InvoiceInfo) SetInvoice(v Invoice)`

SetInvoice sets Invoice field to given value.

### HasInvoice

`func (o *InvoiceInfo) HasInvoice() bool`

HasInvoice returns a boolean if a field has been set.

### GetInvoiceTemplate

`func (o *InvoiceInfo) GetInvoiceTemplate() InvoiceTemplate`

GetInvoiceTemplate returns the InvoiceTemplate field if non-nil, zero value otherwise.

### GetInvoiceTemplateOk

`func (o *InvoiceInfo) GetInvoiceTemplateOk() (*InvoiceTemplate, bool)`

GetInvoiceTemplateOk returns a tuple with the InvoiceTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceTemplate

`func (o *InvoiceInfo) SetInvoiceTemplate(v InvoiceTemplate)`

SetInvoiceTemplate sets InvoiceTemplate field to given value.

### HasInvoiceTemplate

`func (o *InvoiceInfo) HasInvoiceTemplate() bool`

HasInvoiceTemplate returns a boolean if a field has been set.

### GetProducts

`func (o *InvoiceInfo) GetProducts() []ProductItem`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *InvoiceInfo) GetProductsOk() (*[]ProductItem, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *InvoiceInfo) SetProducts(v []ProductItem)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *InvoiceInfo) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetBundledComponentsInfo

`func (o *InvoiceInfo) GetBundledComponentsInfo() []ProductComponent`

GetBundledComponentsInfo returns the BundledComponentsInfo field if non-nil, zero value otherwise.

### GetBundledComponentsInfoOk

`func (o *InvoiceInfo) GetBundledComponentsInfoOk() (*[]ProductComponent, bool)`

GetBundledComponentsInfoOk returns a tuple with the BundledComponentsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundledComponentsInfo

`func (o *InvoiceInfo) SetBundledComponentsInfo(v []ProductComponent)`

SetBundledComponentsInfo sets BundledComponentsInfo field to given value.

### HasBundledComponentsInfo

`func (o *InvoiceInfo) HasBundledComponentsInfo() bool`

HasBundledComponentsInfo returns a boolean if a field has been set.

### GetExpenses

`func (o *InvoiceInfo) GetExpenses() []ExpenseEntry`

GetExpenses returns the Expenses field if non-nil, zero value otherwise.

### GetExpensesOk

`func (o *InvoiceInfo) GetExpensesOk() (*[]ExpenseEntry, bool)`

GetExpensesOk returns a tuple with the Expenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenses

`func (o *InvoiceInfo) SetExpenses(v []ExpenseEntry)`

SetExpenses sets Expenses field to given value.

### HasExpenses

`func (o *InvoiceInfo) HasExpenses() bool`

HasExpenses returns a boolean if a field has been set.

### GetTimeEntries

`func (o *InvoiceInfo) GetTimeEntries() []TimeEntry`

GetTimeEntries returns the TimeEntries field if non-nil, zero value otherwise.

### GetTimeEntriesOk

`func (o *InvoiceInfo) GetTimeEntriesOk() (*[]TimeEntry, bool)`

GetTimeEntriesOk returns a tuple with the TimeEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeEntries

`func (o *InvoiceInfo) SetTimeEntries(v []TimeEntry)`

SetTimeEntries sets TimeEntries field to given value.

### HasTimeEntries

`func (o *InvoiceInfo) HasTimeEntries() bool`

HasTimeEntries returns a boolean if a field has been set.

### GetLogo

`func (o *InvoiceInfo) GetLogo() DocumentInfo`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *InvoiceInfo) GetLogoOk() (*DocumentInfo, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *InvoiceInfo) SetLogo(v DocumentInfo)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *InvoiceInfo) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetBillingSetup

`func (o *InvoiceInfo) GetBillingSetup() BillingSetup`

GetBillingSetup returns the BillingSetup field if non-nil, zero value otherwise.

### GetBillingSetupOk

`func (o *InvoiceInfo) GetBillingSetupOk() (*BillingSetup, bool)`

GetBillingSetupOk returns a tuple with the BillingSetup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingSetup

`func (o *InvoiceInfo) SetBillingSetup(v BillingSetup)`

SetBillingSetup sets BillingSetup field to given value.

### HasBillingSetup

`func (o *InvoiceInfo) HasBillingSetup() bool`

HasBillingSetup returns a boolean if a field has been set.

### GetAgreementBillingInfo

`func (o *InvoiceInfo) GetAgreementBillingInfo() []AgreementBillingInfo`

GetAgreementBillingInfo returns the AgreementBillingInfo field if non-nil, zero value otherwise.

### GetAgreementBillingInfoOk

`func (o *InvoiceInfo) GetAgreementBillingInfoOk() (*[]AgreementBillingInfo, bool)`

GetAgreementBillingInfoOk returns a tuple with the AgreementBillingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreementBillingInfo

`func (o *InvoiceInfo) SetAgreementBillingInfo(v []AgreementBillingInfo)`

SetAgreementBillingInfo sets AgreementBillingInfo field to given value.

### HasAgreementBillingInfo

`func (o *InvoiceInfo) HasAgreementBillingInfo() bool`

HasAgreementBillingInfo returns a boolean if a field has been set.

### GetInfo

`func (o *InvoiceInfo) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *InvoiceInfo) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *InvoiceInfo) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *InvoiceInfo) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


